// go run -gcflags '-N -l' ./main.go
package main

/*
static int foo() {
    return 42;
}
*/
import "C"

import (
    "fmt"
    "time"
)

func foo() int { return 42 }

const N = 100000

func main() {
    var (
        t  time.Time
        dt time.Duration
    )

    t = time.Now()
    for n := 0; n < N; n++ {
        foo()
    }
    dt = time.Since(t)
    fmt.Printf("go:  %vns\n", dt.Nanoseconds()/N)

    t = time.Now()
    for n := 0; n < N; n++ {
        C.foo()
    }
    dt = time.Since(t)
    fmt.Printf("cgo: %vns\n", dt.Nanoseconds()/N)
}
