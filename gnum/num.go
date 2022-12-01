package gnum

import (
	"strconv"
	"strings"
)

// NumFillZero 数字转字符串,位数不够的前面补0
func NumFillZero(n, l int) string {
	numStr := strconv.Itoa(n)
	nl := len(numStr)
	if nl >= l {
		return numStr
	}
	sb := strings.Builder{}
	for i := 0; i < (l - nl); i++ {
		sb.WriteString("0")
	}
	sb.WriteString(numStr)
	return sb.String()
}

// NumMulti 是否一个数字是否是另一个数字的整数倍
func NumMulti(n1, n2 int) bool {
	return (n1 % n2) == 0
}
