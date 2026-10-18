package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int8, p1 uint8, p2 int64) bool {
    var v4 uint = 255
    var v5 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1
    var v3 uint32 = 2
    v1 = (v1 - (v1 % ((v1) | 1)))
    fmt.Println(v1, v3)
    v1 = 18446744073709551615
    var v6 int16 = 32767
    v7 := []int{128}
    v1 = (v1 - v1)
    v8 := GenS1{F0: 579}
    var v9 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
