package main

import (
    "fmt"
)

func genFn0() int8 {
    var v6 bool = true
    fmt.Println(v6)
    _ = v6
    return 127
}

func main() {
    v1 := []int{334, 42, -146, 42}
    var v2 float32 = 1.5
    var v3 uint64 = 42
    var v4 int32 = 119
    var v5 int16 = 536
    var v7 int32 = (v4 >> ((0) & 63))
    fmt.Println(v3)
    var v8 int = 0
    if false {
        if true {
            v2 -= (0.0 + v2)
        }
        v7 -= (v7 % ((0) | 1))
    }
    v2 -= 1e10
    _ = len(v1)
    fmt.Println(v5)
    v4 = (-(v4))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
