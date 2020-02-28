package main

import "fmt"

/*
假设一共有50枚金币.分给下列人：Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano,
规则如下:
名字每包含1个e或者E分1个金币
名字每包含1个i或者I分2个金币
名字每包含1个o或者O分3个金币
名字每包含1个u或者U分4个金币
计算每个人分到多少个金币,一共用了多少个金币和剩余的金币数
*/

var (
	goldNumber = 50
	nameList   = [...]string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano"}
)

func gold() {
	totalNum := 0
	for _, v := range nameList {
		peopleNum := 0
		for _, s := range v {
			if string(s) == "e" || string(s) == "E" {
				peopleNum += 1
			}
			if string(s) == "i" || string(s) == "I" {
				peopleNum += 2
			}
			if string(s) == "0" || string(s) == "O" {
				peopleNum += 3
			}
			if string(s) == "u" || string(s) == "U" {
				peopleNum += 4
			}
		}
		fmt.Printf("%s一共得到%d枚金币\n", v, peopleNum)
		totalNum += peopleNum
	}
	fmt.Printf("一共分了%d枚金币\n", totalNum)
	fmt.Printf("剩余%d枚金币\n", goldNumber-totalNum)
}

func main() {
	gold()
}
