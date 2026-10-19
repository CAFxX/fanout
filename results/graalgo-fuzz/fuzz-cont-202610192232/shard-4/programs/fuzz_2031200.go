package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 string = ("" + ("hello" + "a"))
    v5 := GenS1{F0: 128, F1: (!true)}
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := v3[2147483648]
        _ = fi6
        _ = v7
    }
    var v8 int = 42
    v1[v8] = true
    _ = len(v3)
    v9 := []int{}
    v3[4294967295] = 79
    v9 = append(v9, 127)
    v10 := GenS1{F0: 255, F1: true}
    v8 = 4294967295
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println(v4, v4)
    if false {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = len(v9)
    }
    var v11 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
