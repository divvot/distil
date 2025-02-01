package distil

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

const (
	magicEncryptNumber = 73895173
	sumOffset          = 128789

	maxUint32 = ^uint32(0)
	prefix    = "fe80:"
)

func sumChars(str string) uint32 {
	l := len(str)
	i := 0
	counter := uint32(sumOffset)

	if l == 0 {
		return counter
	}

	for ok := true; ok; ok = i < l {
		counter += uint32(str[i])
		i += 1
	}

	return counter
}

func crack(r uint32) uint32 {
	for i := range maxUint32 {
		f := uint32(magicEncryptNumber * i)
		if r == f {
			return i
		}
	}

	return 0
}

func generateIPv6() string {
	ipv6 := prefix
	for range 4 {
		ipv6 += ":" + strconv.FormatInt(rand.Int64N(0xffff), 16)
	}

	return ipv6
}

func generateIPv4() string {

	octecs := []int{10, 127, 192, 172}

	octeto1 := octecs[rand.IntN(len(octecs))]
	octeto2 := rand.IntN(256)
	octeto3 := rand.IntN(256)
	octeto4 := rand.IntN(254) + 1

	return fmt.Sprintf("%d.%d.%d.%d", octeto1, octeto2, octeto3, octeto4)
}
