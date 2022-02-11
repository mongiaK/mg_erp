/*================================================================
*
*  文件名称：common.go
*  创 建 者: mongia
*  创建日期：2022年01月04日
*
================================================================*/

package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandonString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}
