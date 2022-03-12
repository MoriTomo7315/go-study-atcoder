package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 4096)
	sc.Buffer(buf, 2000000)
	sc.Scan()
	inputs := strings.Split(sc.Text(), "")

	sort.Strings(inputs)

	a := strings.Join(inputs, "")
	fmt.Printf("%s", a)
}
