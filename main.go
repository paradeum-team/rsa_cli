package main

import (
	"bfs_cli_rsa/service/rsa"
	"fmt"
)

func main(){
	fmt.Println("======<<<<<<<>>>>>>>======begin ")

    fmt.Println("获取版本号")
	rsa.NewRsaServiceInstace().GetVersion()

	fmt.Println("创建公私钥密码对")
	rsa.NewRsaServiceInstace().CreateRsaKeyPeer()

	fmt.Println("======<<<<<<<>>>>>>>======end ")
}
