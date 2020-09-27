package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Prefix represents a ip prefix to be randomized
type Prefix struct {
	address [4]uint8
	mask    uint8
}

func (p Prefix) String() string {
	return fmt.Sprintf("%d.%d.%d.%d/%d", p.address[0], p.address[1], p.address[2], p.address[3], p.mask)
}

// MakeRandomPrefix generates a random prefix
func MakeRandomPrefix(minMask, maxMask int) Prefix {
	if minMask < 0 {
		minMask = 0
	}
	if maxMask > 32 {
		maxMask = 32
	}
	if minMask > maxMask {
		minMask = maxMask
	}
	var p Prefix
	buf := make([]byte, 4)
	rand.Read(buf)
	for i := range p.address {
		p.address[i] = uint8(buf[i])
	}
	if minMask == maxMask {
		p.mask = uint8(minMask)
	} else {
		p.mask = uint8(minMask + rand.Intn(maxMask-minMask+1))
	}
	return p
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := Prefix{address: [4]uint8{10, 2, 2, 1}, mask: 4}
	fmt.Println(a)

	for i := 0; i < 1000; i++ {
		b := MakeRandomPrefix(8, 24)
		fmt.Println(b)
	}
}
