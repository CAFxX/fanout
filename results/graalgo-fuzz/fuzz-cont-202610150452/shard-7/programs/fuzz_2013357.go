package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
func genFn0(p0 int) uint16 {
    p0 -= (p0 ^ 32768)
    p0 += (p0 | -9223372036854775808)
    v5 := []int{p0, p0, 127}
    if len(v5) > 0 {
        v5[0] = p0
    }
    _ = p0
    _ = v5
    return 598
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int8 = -116
    var v6 int64 = 42
    v7 := GenS1{F0: 0, F1: 65535}
    v8 := v3[4294967295]
    v9 := []int{4294967295}
    v10 := GenS1{F0: 2, F1: 4294967295}
    fmt.Println(v1)
    fmt.Println(v8, v4)
    v8 = v8
    var v11 uint = 255
    delete(v3, 523)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
