package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	//指定需要抓取到的空间
	bucket = "test"
	//指定需要抓取的文件的url，必须是公网上面可以访问到的
	//指定抓取保存到空间的文件的key指
	key    = "test" + time.Now().String()
	domain = "http://oa8dpdexh.bkt.clouddn.com/"
	accKey = "你的access key"
	secKey = "你的secret key"
)

func main() {
	targetUrl := flag.String("url", "www.qiniu.com", "输入图片地址")

	flag.Parse()

	lTarget := strings.Split(*targetUrl, "/")
	key = lTarget[len(lTarget)-1]
	conf.ACCESS_KEY = accKey
	conf.SECRET_KEY = secKey
	fmt.Printf("fetching %s\n", *targetUrl)
	//new一个Bucket对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	//调用Fetch方法

	err := p.Fetch(nil, key, *targetUrl)
	if err != nil {
		fmt.Println("bucket.Fetch failed:", err)
	} else {
		fmt.Println("fetch success\nthe response url:")
		fmt.Println(domain + "/" + key)
	}

}
