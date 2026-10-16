package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }

func main() {
    var v1 int = 1
    v2 := make(map[int]int)
    v3 := []int{169}
    var v4 string = "a"
    v5 := make(map[uint32]string)
    var v6 int32 = 0
    var v7 bool = false
    _ = len(v3)
    v4 = ("a" + (((v4 + "") + "foo bar") + ((v4 + v4) + ("foo bar" + v4))))
    var v8 int16 = 990
    v1 = v1
    v6 -= v6
    var v9 uint64 = 86
    v10 := GenS1{F0: 255, F1: 212, F2: v1}
    if ((!v7) || (true || v7)) {
        var v11 int = v1
        _ = v11
    }
    v12 := GenS1{F0: 255, F1: 0, F2: (v1 & v1)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
