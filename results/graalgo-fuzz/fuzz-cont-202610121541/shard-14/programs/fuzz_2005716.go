package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]string)
    var v3 int8 = 0
    var v4 int64 = 1
    var v5 uint64 = 42
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{}
            v2[fi7] = v6
            v10 := GenS1{F0: 495, F1: false, F2: v5}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{}
        _ = fi7
        _ = v11
    }
    if (!true) {
        fmt.Println(v5)
    }
    v12 := GenS2{F0: 0}
    switch v6 {
    case "foo bar":
        v13 := GenS3{F0: (v5 ^ 42)}
        _ = v13
    case "":
        var v14 uint8 = (v1 << ((118) & 63))
        _ = v14
    }
    v5 = v5
    v15 := []int{-985, -9223372036854775808}
    v4 = v4
    fmt.Println(v6, v6, v3)
    var v16 int64 = v4
    var v17 int8 = (v3 + (v3 % ((v3) | 1)))
    var v18 uint = 904
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
