package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{42}
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 int8 = 42
    fmt.Println(v5)
    v6 := GenS1{F0: v5, F1: 18446744073709551615, F2: 1}
    fmt.Println(v5)
    v7 := GenS1{F0: v5, F1: 2, F2: 65535}
    v5 = (-(v5))
    switch 127 {
    case -1:
        var v8 int8 = v5
        _ = v8
    default:
        fmt.Println(v5)
    }
    switch ("hello" + ("\n\t" + "hello")) {
    case "":
        var v9 int32 = -1
        _ = v9
    default:
        v2 = append(v2, 42)
    }
    fmt.Println(v5)
    var v10 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
