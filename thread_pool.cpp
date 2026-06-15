#include "thread_pool.h"

ThreadPool::ThreadPool(ThreadPool::MaxTaskPolicy policy) :run_(true), idle_thread_num_(0), policy_(policy)
{
    monitor_thread_ = std::make_unique<std::thread>([this] {
        while (1) {
            while (run_) {
                std::unique_lock<std::mutex> lock(monitor_mutex_);
                std::cv_status s = monitor_cond_.wait_for(lock, std::chrono::seconds(kShrinkThreadPeriod));
                if (s == std::cv_status::no_timeout && !run_) {
                    return;
                }
                if (s == std::cv_status::timeout) {
                    break;
                }
            }
            shrink_to_fit();
        }
        });
    for (size_t i = 0; i < kPermanentThreadNum; ++i) {
        create_thread();
    }
}

ThreadPool::~ThreadPool()
{
    run_ = false;
    condition_.notify_all();
    monitor_cond_.notify_all();
    for (auto it = pool_.begin(); it != pool_.end(); ++it) {
        it->join();
    }
    monitor_thread_->join();
}

void ThreadPool::create_thread()
{
    std::unique_lock<std::mutex> lock(mutex_);
    auto it = pool_.insert(pool_.end(), std::thread());
    std::thread t([this, it] {
        process_task(it);
        });
    idle_thread_num_++;
    pool_.back() = std::move(t);
}

void ThreadPool::process_task(std::list<std::thread>::iterator it)
{
    idle_thread_num_--;
    while (1) {
        std::function<void()> t;
        {
            std::unique_lock<std::mutex> lock(mutex_);
            while (task_.empty() && run_)
            {
                idle_thread_num_++;
                if (idle_thread_num_ > 0 && pool_.size() > kPermanentThreadNum) {
                    dead_thread_it_.emplace_back(it);
                    return;
                }
                condition_.wait(lock);               
                idle_thread_num_--;
            }
            if (!run_)
            {
                return;
            }
            t = std::move(task_.front());
            task_.pop();
        }
        t();
    }
}

void ThreadPool::shrink_to_fit()
{
    {
        /*
        std::cout<<"| shrink |"<<std::endl;
        std::cout<<"idle threadNum"<<idle_thread_num_<<std::endl;
        std::cout<<"pool size"<<pool_.size()<<std::endl;
        std::cout<<"dead threadNum"<<dead_thread_it_.size()<<std::endl;
        std::cout<<"task size"<<task_.size()<<std::endl<<std::endl;
*/

        std::unique_lock<std::mutex> lock(mutex_);
        for (auto it : dead_thread_it_) {
            it->join();
            pool_.erase(it);
            idle_thread_num_--;
        }
        dead_thread_it_.clear();
    }
    condition_.notify_all();
}
