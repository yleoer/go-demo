package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shirou/gopsutil/process"
)

func main() {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		log.Fatalf("failed to get process %s by pid %d", err, os.Getpid())
	}

	mp, _ := p.MemoryPercent()

	fmt.Println(mp)
}
