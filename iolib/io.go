package iolib

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func check(e error, s string) {
	if e != nil {
		fmt.Printf("%s --> %s\n", s, e)
		panic(e)
	}
}

func GetRand(seed int64) *rand.Rand {
	var nseed rand.Source
	if seed != 0 {
		nseed = rand.NewSource(seed)
	} else {
		nseed = rand.NewSource(time.Now().UnixNano())
	}
	r := rand.New(nseed)
	return r
}

func WriteRandomBytesToDevice(r *rand.Rand, num int, device string, offset int64) {
	b := make([]byte, num)
	for i := 0; i <= num; i++ {
		b = append(b, byte(r.Float32()*255))
	}
	WriteBytesToDevice(b, device, offset)
}

func WriteBytesToDevice(data []byte, device string, offset int64) {
	f, err := os.Create(device)
	check(err, "Could not open device")
	defer f.Close()

	f.Seek(offset, 0)
	_, err = f.Write(data)
	check(err, "Could not write to device")
	f.Sync()
}
