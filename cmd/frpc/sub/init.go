package sub

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/util"
	"os"
)

func InitFrpcIni() {
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	subdomain := os.Getenv("SUBDOMAIN")
	token := os.Getenv("TOKEN")
	dirPath := "/etc/frp"
	filePath := dirPath + "/frpc.ini"

	if GetMd5String(token) != "7d7ccf9be8c6d8fe65b03a79a92872d86fb6766b" {
		log.Error("错误的token:" + token)
		os.Exit(1)
	}

	if ip == "" {
		ip = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	if subdomain == "" {
		id, _ := util.RandID()
		subdomain = id
	}

	os.Mkdir(dirPath, os.ModePerm)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	_, err = file.WriteString("[common]\nserver_addr = 124.223.30.237\nserver_port = 7000\ntoken =" + token + " \n\n[" + subdomain + " by docker]\ntype = http\nlocal_ip = " + ip + "\nlocal_port = " + port + "\nsubdomain = " + subdomain + "\n")
	log.Info("代理地址 http://" + ip + ":" + port)
	log.Info("访问地址 https://" + subdomain + ".xiw.asia")
	if err != nil {
		return
	}
}

func GetMd5String(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
