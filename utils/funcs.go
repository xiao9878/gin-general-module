package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 类型转换

func Float64(v interface{}) (r float64, err error) {
	if v == nil {
		r = 0
		return
	}
	str := fmt.Sprintf("%v", v)
	if str == "" {
		r = 0
		return
	}
	r, err = strconv.ParseFloat(str, 64)
	return
}

func Int(v interface{}) (r int, err error) {
	if v == nil {
		r = 0
		return
	}
	str := fmt.Sprintf("%v", v)
	if str == "" {
		r = 0
		return
	}
	return strconv.Atoi(str)
}

// 三目运算
func IF(b bool, v1 interface{}, v2 interface{}) interface{} {
	if b {
		return v1
	}
	return v2
}

/****
 * 获取0 - max 之间随机的 num 个位置
 */
func GetRandomID(max int, num int) (ids []int) {
	ids = make([]int, 0, num)
	if max < num {
		return ids
	}
	idMap := make(map[int]bool)
	var tmp int = 0
	for {
		tmp = rand.Intn(max)
		if idMap[tmp] {
			// 已经有了
			continue
		}
		idMap[tmp] = true
		ids = append(ids, tmp)
		if len(ids) >= num {
			break
		}
	}
	return
}

func String(v interface{}) string {
	if v == nil {
		return ""
	}
	str := fmt.Sprintf("%v", v)
	// 去除空格
	str = strings.Trim(str, " ")
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func Trim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

// md5加密
func MD5V(v string) string {
	h := md5.New()
	h.Write([]byte(v))
	return hex.EncodeToString(h.Sum(nil))
}
