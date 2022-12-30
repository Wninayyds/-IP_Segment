package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ip_list []string
	ip_str := flag.String("ip", "-", "IP段")
	flag.Parse()
	var new_ip_str string
	string_slice := strings.Split(*ip_str, ".")
	ip_a_max := string_slice[len(string_slice)-2]
	num, err := strconv.Atoi(ip_a_max)
	if err != nil {
		panic(err)
	}
	if num < 255 {
		var ip int
		for ip = 0; ip < 256; ip++ {
			//fmt.Printf("ip 为 %v\n", ip)
			new_ip_str = string_slice[0] + "." + string_slice[1] + "." + string_slice[2] + "." + strconv.Itoa(ip)
			ip_list = append(ip_list, new_ip_str)
		}
	}
	//fmt.Println(ip_list)
	//file, err := os.OpenFile("./res.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0402)
	file, err := os.OpenFile("./res.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for index, res_ip := range ip_list {
		ip_list[index] = res_ip
		fmt.Printf("%v\n", ip_list[index])
		write.WriteString(ip_list[index] + "\n")
	}
	write.Flush()
}
