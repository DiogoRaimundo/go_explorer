/*
	Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
		ch := make(chan int)
		ch <- v    // Send v to channel ch.
		v := <-ch  // Receive from ch, and assign value to v.

	By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	Channels can also be buffered by providing the buffer length as the second argument to make to initialize a buffered channel.
	Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	If no goroutines are receiving, a deadlock appears (fatal error: all goroutines are asleep - deadlock!).

	Channels can be used in a range loop as slices are:
		for i := range c

	A sender can close a channel to indicate that no more values will be sent.
	Closing is only necessary when the receiver must be told there are no more values coming.
	Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.
*/

package goTour06

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci02_04(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func RunExample02_04() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	c = make(chan int, 10)
	go fibonacci02_04(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
