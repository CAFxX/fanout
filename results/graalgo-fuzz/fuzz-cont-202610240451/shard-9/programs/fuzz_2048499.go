package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 1
    var v3 uint64 = 998
    var v4 int8 = 0
    fmt.Println(v3, v4)
    switch ("foo bar" + ("foo bar" + "foo bar")) {
    case "hello":
        v5 := []int{-839, 485}
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: v4, F1: "\n\t", F2: 629}
        _ = v6
    }
    v7 := []int{4294967295, -291}
    var v8 int64 = 832
    v3 += (v3 ^ 18446744073709551615)
    v9 := GenS1{F0: v4, F1: "\n\t", F2: 452}
    v2 += (^(v2))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
