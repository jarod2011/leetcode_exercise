package topic7_reverse

import (
	"bytes"
	"math"
	"strconv"
)

func reverseByConvertString(x int) int {
	m := strconv.Itoa(x)
	flag := x >= 0
	var res []byte
	pos := 0
	if !flag {
		pos++
	}
	for i := len(m) - 1; i >= pos; i-- {
		res = append(res, m[i])
	}
	rev, _ := strconv.ParseInt(string(bytes.TrimPrefix(res, []byte{0x30})), 10, 64)
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}
	if flag {
		return int(rev)
	}
	return int(-int32(rev))
}
