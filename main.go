package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/*
#include <stdio.h>

void printInt(int v) {
    printf("printInt: %d\n", v);
}
 */
import "C"

func Foo() {
	fmt.Println("print 1")
	defer fmt.Println("print 2")
	fmt.Println("print 3")
}

func f(a ...int) {
	fmt.Printf("%#v\n", a)
	fmt.Println(a == nil)
}

func main() {
	v := 42
	C.printInt(C.int(v))

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Printf("%s %d\n", image.ID[7:20], image.Size)
	}
}
