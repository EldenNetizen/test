#ifndef THREADPOOL_H
#define THREADPOOL_H

#include <mutex>
#include <future>
#include <thread>
#include <vector>
#include <list>
#include <queue>
#include <condition_variable>
#include <iostream>

class ThreadPool
{
public:
    enum class MaxTaskPolicy {
        Abort
    };
    ThreadPool(ThreadPool::MaxTaskPolicy policy = ThreadPool::MaxTaskPolicy::Abort);
    ThreadPool(ThreadPool const& pool) = delete;
    ThreadPool(ThreadPool&& pool) = delete;
    ThreadPool& operator=(ThreadPool&& pool) = delete;
    ThreadPool& operator=(const ThreadPool& pool) = delete;
    ~ThreadPool();
    template<class F, class... Args>
    auto enqueue(F&& f, Args&&... args)->std::future<decltype(f(args...))>;

private:
    void create_thread();
    void shrink_to_fit();
    void process_task(std::list<std::thread>::iterator it);
    const size_t kPermanentThreadNum = 4;
    const size_t kMaxThreadNum = 8;
    const size_t kMaxTaskNum = 8;
    const int kShrinkThreadPeriod = 10;
    std::list<std::thread> pool_;
    std::queue<std::function<void()>> task_;
    std::vector<std::list<std::thread>::iterator> dead_thread_it_;
    std::mutex mutex_;
    std::mutex monitor_mutex_;
    std::condition_variable condition_;
    std::condition_variable monitor_cond_;
    std::atomic<bool> run_;
    std::atomic<int> idle_thread_num_;
    std::unique_ptr<std::thread> monitor_thread_;
    MaxTaskPolicy policy_;
};

template<class F, class... Args>
auto ThreadPool::enqueue(F&& f, Args&&... args) -> std::future<decltype(f(args...))>
{

    {
        std::lock_guard<std::mutex> lock(mutex_);
        if (task_.size() >= kMaxTaskNum) {
            switch (policy_) {
            case ThreadPool::MaxTaskPolicy::Abort: throw "Max Task";
            }
        }
    }

    std::function<decltype(f(args...))()> func(std::bind(std::forward<F>(f), std::forward<Args>(args)...));
    std::shared_ptr<std::packaged_task<decltype(f(args...))()>> packaged_task = std::make_shared<std::packaged_task<decltype(f(args...))()>>(std::move(func));
    std::future<decltype(f(args...))> future = packaged_task->get_future();
    {
        std::unique_lock<std::mutex> lock(mutex_);
        task_.push([task = std::move(packaged_task)]{ (*task)(); });
        if (!idle_thread_num_ && pool_.size() < kMaxThreadNum) {
            lock.unlock();
            create_thread();
        }
    }
    condition_.notify_one();
    return future;
}
#endif // THREADPOOL_H
