//#include"server.h"
#include"thread_pool.h"
#include <iostream>

int main()
{
	
	std::cout << "thread id:" << std::this_thread::get_id();
	boost::asio::io_context io_context;
	ThreadPool pool;
	//Server s(io_context);
	

	//for (size_t i = 0; i < 8; ++i) {
		//pool.enqueue([&io_context]() { io_context.run(); });
	//}

	while (1);
	return 0;
	
}