package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	ints, err := ReadInts(bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	uf, ints := NewUF(ints[0]), ints[1:]
	var p, q int
	for len(ints) > 1 {
		p, q, ints = ints[0], ints[1], ints[2:]
		uf.union(p, q)
	}
	fmt.Println(uf)
	fmt.Printf("%d, %d connected: %v\n", 0, 1, uf.connected(0, 1))
	fmt.Printf("%d, %d connected: %v\n", 2, 3, uf.connected(2, 3))
}

type UF []int

func NewUF(n int) UF {
	var uf UF
	uf = make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
	}
	return uf
}

func (uf UF) connected(p, q int) bool {
	return uf[p] == uf[q]
}

func (uf UF) union(p, q int) {
	pid := uf[p]
	qid := uf[q]
	if pid != qid {
		for i := 0; i < len(uf); i++ {
			if uf[i] == pid {
				uf[i] = qid
			}
		}
	}
}

func (uf UF) String() string {
	var buffer bytes.Buffer
	for i, e := range uf {
		buffer.WriteString(strconv.Itoa(e))
		if i != len(uf)-1 {
			buffer.WriteString(", ")
		}
	}
	return "[" + buffer.String() + "]"
}
