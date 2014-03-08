package main

import (
	"fmt"
	"regexp"
)

func main() {
	phash := "827ccb0eea8a706c4c34a16891f84e7b"

	hash_arr := map[string][]string{
		"^[a-f0-9]{4}$": {"CRC-16","CRC-16-CCITT","FCS-16"},
		"^[a-f0-9]{8}$": {"Adler-32","CRC-32","CRC-32B","FCS-32","GHash-32-3","GHash-32-5","FNV-132","Fletcher-32","Joaat","ELF-32","XOR-32"},
		"^[0-9a-f]{32}$": {"MD5","MD4","MD2","NTLM","LM","RAdmin v2.x","RIPEMD-128","Haval-128","Tiger-128","Snefru-128","MD5(ZipMonster)","Skein-256(128)","Skein-512(128)"},
	}
	for regex := range hash_arr {
		r, _ := regexp.Compile(regex)
		if(r.MatchString(phash)){
			fmt.Println(hash_arr[regex])
		}
	}
}
