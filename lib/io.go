package kyopro

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func NextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func NextStr() string {
	sc.Scan()
	return sc.Text()
}
