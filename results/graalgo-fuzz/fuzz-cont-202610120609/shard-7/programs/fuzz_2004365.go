package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -1
    var v3 uint64 = 894
    var v4 uint32 = 0
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 string = "a"
        var v7 int = (fi5 ^ 554)
        v8 := GenS1{F0: 0, F1: v4, F2: (fi5 ^ (v7 % ((v7) | 1)))}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v3, v2)
    v9 := []int{int(v4)}
    var v10 int8 = 100
    _ = len(v9)
    v11 := GenS1{F0: 106, F1: v4, F2: -235}
    var v12 string = "hello"
    v13 := v1[-819]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
