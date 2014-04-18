emlog_static.go
====================
###功能###
1. 新闻博客类非数据交互网站通用首页静态化
2. emlog博客系统文章页静态化
3. 使用Go语言编写，跨平台

###条件###
1. 需要有服务器管理权限，因为要编译、运行软件
2. 需要Go语言编译环境，请去[The Go Programming Language](http://golang.org/) 下载

###使用方法###
1. 两个文件单独编译:

	`
	go build common_static_homepage.go
	`
	
	`
	go build emlog_static_articles
	`  
2. common_static_homepage `通用首页`静态化 使用方法为：    
    
    `
    ./common_static_homepage -url=http://example.com/index.php -t=30    
    `
    
    t 为间隔时间，单位是秒    

3. emlog_static_articles `emlog文章页`静态化 使用方法为：    
    
    `
    ./emlog_static_articles -url=http://example.com/index.php -c=10 -n=100    
    `
    
    c 为并发数，n 为总文章数，即文章id最大值    
