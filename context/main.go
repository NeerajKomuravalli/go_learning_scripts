package main

import (
	"context"
	"fmt"
	"time"
)

func handleDeadLine(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Printf("Closing the context %s\n", ctx.Value("Name"))
	}
}

func main() {
	// Cotext.Background is nil and cannot be cancelled
	ctx := context.Background()
	fmt.Println("Printing initial values of context.Background")
	fmt.Println("Done : ", ctx.Done())
	fmt.Print("Deadline : ")
	fmt.Println(ctx.Deadline())
	fmt.Println("Error : ", ctx.Err())
	fmt.Println("Value of key : 'anything' : ", ctx.Value("anything"))
	fmt.Println()
	fmt.Println("Adding key pair data to context")
	ctxWithValue := context.WithValue(ctx, "Hello", "World!")
	fmt.Printf("Key 'Hello' has value %v\n", ctxWithValue.Value("Hello"))
	fmt.Println()
	fmt.Println("With deadlines of parent and child to see their closing pattern")
	ctxWithDeadline, _ := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	ctxWithDeadline2, _ := context.WithDeadline(ctxWithDeadline, time.Now().Add(10*time.Second))
	ctxWithDeadline = context.WithValue(ctxWithDeadline, "Name", "ctxWithDeadline1")
	ctxWithDeadline2 = context.WithValue(ctxWithDeadline2, "Name", "ctxWithDeadline2")
	go handleDeadLine(ctxWithDeadline)
	go handleDeadLine(ctxWithDeadline2)
	fmt.Println("Starting a 3 second sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("Ending a 3 second sleep")
	fmt.Println()

	fmt.Println("With timeout of parent and child to see their closing pattern")
	ctxWithTimeOut, _ := context.WithTimeout(ctx, 2*time.Second)
	ctxWithTimeOut2, _ := context.WithTimeout(ctxWithDeadline, 10*time.Second)
	ctxWithTimeOut = context.WithValue(ctxWithTimeOut, "Name", "ctxWithTimeOut1")
	ctxWithTimeOut2 = context.WithValue(ctxWithTimeOut2, "Name", "ctxWithTimeOut2")
	go handleDeadLine(ctxWithTimeOut)
	go handleDeadLine(ctxWithTimeOut2)
	fmt.Println("Starting a 3 second sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("Ending a 3 second sleep")
	fmt.Println()

	fmt.Println("With deadlines of parent and child to see their closing pattern by closing before their deadline")
	ctxWithDeadline11, ctxWithDeadlineClosing11 := context.WithDeadline(ctx, time.Now().Add(20*time.Second))
	ctxWithDeadline22, _ := context.WithDeadline(ctxWithDeadline11, time.Now().Add(30*time.Second))
	ctxWithDeadline11 = context.WithValue(ctxWithDeadline11, "Name", "ctxWithDeadline11")
	ctxWithDeadline22 = context.WithValue(ctxWithDeadline22, "Name", "ctxWithDeadline22")
	go handleDeadLine(ctxWithDeadline11)
	go handleDeadLine(ctxWithDeadline22)
	ctxWithDeadlineClosing11()
	fmt.Println("Starting a 2 second sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("Ending a 2 second sleep")
	fmt.Println()
}
