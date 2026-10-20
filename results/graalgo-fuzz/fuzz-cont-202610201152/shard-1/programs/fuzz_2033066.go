package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = 1
    var v2 bool = true
    var v3 bool = true
    var v4 uint = 480
    var v5 bool = false
    v2 = (42 == 2)
    v6 := []int{872, -1, -9223372036854775808, -2147483648}
    v4 = v4
    v7 := GenS1{F0: 1}
    v5 = v5
    if len(v6) > 0 {
        v6[0] = -1
    }
    var v8 int16 = -1
    v1 -= 308
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
