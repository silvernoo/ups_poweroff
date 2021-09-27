package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"ups_poweroff/goping"
)

func main() {
	var counter uint64 = 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	interval := flag.Int("i", 3, "Interval (Second)")
	gateway := flag.String("g", "", "Gateway")
	retry := flag.Uint64("r", 10, "Retry count")
	flag.Parse()
	if len(*gateway) == 0 {
		panic("Gateway cannot be empty\n")
	}
	raddr, err := net.ResolveIPAddr("ip", *gateway)
	if err != nil {
		panic(fmt.Sprintf("Fail to resolve %s, %s\n", *gateway, err))
	}
	timer := time.NewTimer(time.Duration(*interval) * time.Second)
out:
	for {
		timer.Reset(time.Duration(*interval) * time.Second)
		select {
		case <-timer.C:
			if _, err := goping.SendICMPRequest(goping.GetICMP(uint16(1)), raddr); err != nil {
				fmt.Printf("Error: %s\n", err)
				counter++
				if counter >= *retry {
					power0ff()
					counter = 0
				}
			} else {
				counter = 0
			}
		case <-c:
			break out
		}
	}
}

func power0ff() {
	_, err := exec.Command("bash","-c", "shutdown -h now").Output()
	if err != nil {
		fmt.Println(err)
	}
}
