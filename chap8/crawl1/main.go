package main

import (
	"fmt"
	"golearn/chap5/links"
	"log"
	"os"
)

/**
    @Destription:一个并发的爬虫
    @Author:lichang
    @Date:2020/1/19
    @Time:下午8:20
*/

func craw1(url string) []string {
	/**
	    @Description:从给定的url中提取出里面的链接
	    @Parm:
	        url: 网页链接
	    @Return:
	        []string:给定的网页中的所有链接
	 */
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil{
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	
	/**
	    @Des: 将命令行参数传递到worklist channnel
	*/
	go func(){worklist <- os.Args[1:]}()

	seen := make(map[string]bool)
	for list:=range worklist{
		for _, link := range list{
			if !seen[link] {
				seen[link] = true
				go func(link string){
					worklist <- craw1(link)
				}(link)
			}
		}
	}
}
