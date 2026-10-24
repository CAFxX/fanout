package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }

func main() {
    v1 := []int{128, -210, 2147483647}
    var v2 bool = false
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 1, F1: 115, F2: "foo bar"}
        fmt.Println(fi4, fi4)
        v6 := GenS1{F0: -613, F1: 2, F2: "\n\t"}
        _ = fi4
        _ = v5
        _ = v6
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v7 uint = 1
    var v8 string = "hello"
    var v9 uint = v7
    fmt.Println(v2, v2, v2)
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int = (fi10 >> ((v7) & 63))
        var v12 int32 = 1
        v9 = v7
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int32 = 357
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 string = (((v8 + (v8 + (("a" + v8) + v8))) + v8) + v8)
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v8, v9)
    delete(v3, 42)
    v16 := GenS1{F0: 1, F1: 215, F2: v8}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
