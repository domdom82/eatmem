package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.cloudfoundry.org/bytefmt"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Memory string `short:"m" long:"mem" description:"Memory to eat e.g. 100M, 1G, 1024KB etc." required:"true"`
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	nBytes, err := bytefmt.ToBytes(opts.Memory)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Printf("Eating up %s of memory...\n", opts.Memory)

	buf := bytes.Buffer{}

	for ; nBytes > 0; nBytes-- {
		buf.WriteByte(42)
	}

	fmt.Printf("%s allocated.\n Press ctrl-c to release it and quit.\n", opts.Memory)
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done
}
