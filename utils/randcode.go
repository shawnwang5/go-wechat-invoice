package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func ToSerialCode() string {
	// 创建一个范围在[100000, 999999]的大整数
	max := big.NewInt(999999)
	max.Sub(max, big.NewInt(100000))

	// 生成一个在范围内的随机数
	randomNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	// 将大整数转换为字符串，并在前面补充0以满足6位数的要求
	randomNumStr := fmt.Sprintf("%06d", randomNum)
	//fmt.Println(randomNumStr)
	return randomNumStr
}
