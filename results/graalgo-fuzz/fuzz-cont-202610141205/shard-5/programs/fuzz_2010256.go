package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    var v5 float64 = -2.25
    var v6 int32 = 42
    v4 = (v1 == v1)
    v7 := GenS1{F0: 128, F1: v2, F2: 128}
    v8 := GenS1{F0: 237, F1: "\n\t", F2: 0}
    v9 := []int{}
    var v10 int8 = 100
    v3[2] = 65535
    if len(v9) > 0 {
        v9[0] = 867
    }
    var v11 bool = v4
    var v12 uint = 2
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v13 uint8 = 22
    fmt.Println(v5, v6, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
