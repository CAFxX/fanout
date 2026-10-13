package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    var v1 uint8 = 0
    var v2 uint16 = 1
    var v3 int = -9223372036854775808
    var v4 uint16 = (v2 % ((0) | 1))
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 4294967295, F1: "a"}
    var v7 string = "hello"
    v8 := []int{(v3 & v3), (v3 * 674), 128, v3}
    fmt.Println(v2, v4, v7)
    if ((447 <= -1) && true) {
        var v9 float64 = ((0.0 * 1.5) / -2.25)
        _ = v9
    }
    v1 -= 130
    for fi10 := 0; fi10 < 7; fi10++ {
        _ = len(v5)
        delete(v5, 2)
        _ = fi10
    }
    v1 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
