package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	fmt.Println("go + etcd demo")

	// 配置 etcd 客户端
	endpoints := []string{"localhost:2379"} // etcd 服务的地址
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close() // 确保在函数结束时关闭连接

	// 设置键值对
	key := "name"
	value := "javapub"
	putResp, err := cli.Put(context.Background(), key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Put: ", putResp.Header.Revision)

	// 获取键的值
	getResp, err := cli.Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get: ", string(getResp.Kvs[0].Value))

	// 删除键
	delResp, err := cli.Delete(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete: ", delResp.Header.Revision)
}
