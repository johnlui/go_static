/*
	Use:
		go build emlog_static_homepage.go
		./emlog_static_homepage -url=http://example.com/index.php -t=30

		//请将上面的example.com替换成你自己的网址，注意后面要加上index.php
		//t的单位是秒，默认为一分钟

		//默认在网站根目录下编译运行，如需改动，可直接修改os.OpenFile第一个参数


*/

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	var home_url = flag.String("url", "http://lvwenhan.com", "your url, include /index.php")
	var interval_time = flag.Int("t", 60, "interval time /s ")
	flag.Parse()
	tc := time.Tick(time.Second * time.Duration(*interval_time))
	for i := 0; i < 5256000; i++ {
		<-tc
		action(*home_url)
		fmt.Println(i)
	}
}

func action(home_url string) {

	content, err := http.Get(home_url)

	defer content.Body.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		buf := make([]byte, 10240)
		//createfile
		f, err1 := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
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

	t := time.Now().Unix()

	fmt.Printf("%s\n", "首页生成成功！ 在："+time.Unix(t, 0).String())

}
