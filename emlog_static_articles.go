/*
	Use:
		go build emlog_static_articles.go
		./emlog_static_articles -url=http://example.com/index.php -c=10 -n=100

		//请将上面的example.com替换成你自己的网址，注意后面要加上index.php
		//c为同时运行的协程数，n为文章总数

		//默认在网站根目录下编译运行，建议修改os.OpenFile第一个参数，全部归到static目录下,注意先创建此目录


*/

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var home_url = flag.String("url", "http://lvwenhan.com/blog/index.php", "your url, include /index.php")
	var xiancheng = flag.Int("c", 10, "interval time /s ")
	var wenzhangzongshu = flag.Int("n", 100, "interval time /s ")
	flag.Parse()

	var wenzhangshu int = *wenzhangzongshu
	var number int = *xiancheng
	var blog_url string = *home_url

	yushu := wenzhangshu % number
	beishu := wenzhangshu / number
	fmt.Println(beishu)
	for j := 0; j < beishu; j++ {
		do_100(number*j, number, blog_url)
	}

	do_100(number*beishu, yushu+1, blog_url)

}

func do_100(begin, number int, blog_url string) {

	chs := make([]chan int, number)

	for m := 0; m < len(chs); m++ {
		aid := begin + m
		chs[m] = make(chan int)
		go action(aid, chs[m], blog_url)

	}

	for _, ch := range chs {
		<-ch
	}

}

func action(aid int, ch chan int, blog_url string) {

	blog_url = blog_url + "?post=" + strconv.Itoa(aid)

	content, err := http.Get(blog_url)

	defer content.Body.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		buf := make([]byte, 10240)
		//createfile
		f, err1 := os.OpenFile(strconv.Itoa(aid)+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err1 != nil {
			panic(err1)
			return
		}
		defer f.Close()

		for {
			n, _ := content.Body.Read(buf)
			if 0 == n {
				break
			}
			f.WriteString(string(buf[:n]))
		}
	}

	ch <- aid * 2

	fmt.Printf("%s\n", "第"+strconv.Itoa(aid)+"篇文章生成成功")

}
