package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    var v5 int32 = 1
    fmt.Println(v4, v4)
    v4 = "a"
    v4 = "\n\t"
    v2 = v2
    if (false || ((v4 + "") != v4)) {
        v6 := GenS1{F0: 32768}
        v7 := GenS1{F0: -2147483648}
        _ = v6
        _ = v7
    }
    v8 := []int{0, 0}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v5 -= (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
