package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func in(in chan<- error, f func() error) {
	in <- f()
}

func parallel(arr []func() error, n int, errNum int) {

	if errNum > len(arr) {
		fmt.Printf("Too many number of maximum errors. Must be greater than %d \n", len(arr))
	}

	errCh := make(chan error, len(arr))

	i := 0
	for len(errCh) < errNum {
		if runtime.NumGoroutine() <= n && i < len(arr) {
			go in(errCh, arr[i])
			i += 1
		}
	}

	close(errCh)

	fmt.Printf("Exited, num errors: %d", len(errCh))

}

func main() {
	n, errNum := 3, 5

	f1 := func() error {
		time.Sleep(time.Second * 2)
		return errors.New("error")
	}

	f2 := func() error {
		return errors.New("error")
	}

	f3 := func() error {
		time.Sleep(time.Second * 4)
		return errors.New("error")
	}

	f4 := func() error {
		time.Sleep(time.Second * 6)
		return errors.New("error")
	}

	f5 := func() error {
		time.Sleep(time.Second * 5)
		return errors.New("error")
	}

	f6 := func() error {
		time.Sleep(time.Second * 8)
		return errors.New("error")
	}

	funcs := []func() error{f1, f2, f3, f4, f5, f6}

	parallel(funcs, n, errNum)
}
