package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 int64 = -1
    var v4 int32 = 1
    switch (("\n\t" + "a") + "a") {
    case "hello":
        var v5 int16 = 623
        _ = v5
    case "\n\t":
        var v6 uint64 = 120
        _ = v6
    case "":
        v7 := GenS1{F0: 42, F1: 128, F2: (v3 / ((1) | 1))}
        _ = v7
    }
    v4 -= (v4 % ((v4) | 1))
    v4 -= v4
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int64 = (^(v3))
        fmt.Println(v3, v9)
        var v10 uint16 = 255
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = 0
    v11 := GenS2{F0: -2147483648}
    var v12 string = ("" + "\n\t")
    v3 = v3
    delete(v1, v12)
    v4 -= (v4 % ((v4) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
