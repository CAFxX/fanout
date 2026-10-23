package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 90
    v3 := make(map[uint32]int)
    var v4 string = ("foo bar" + ("hello" + ""))
    v5 := []uint32{109, 2, 1, 668}
    v2 = (v2 >> ((42) & 63))
    var v6 uint8 = 2
    v4 = "hello"
    delete(v1, 2)
    var v7 bool = true
    v8 := GenS1{F0: (v6 | v6), F1: -2147483648}
    v9 := GenS1{F0: v6, F1: 1}
    v6 = 74
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
