package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 bool }

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 1e10
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := []int{fi4, 32767, fi5}
            _ = fi5
            _ = v6
        }
        var v7 uint64 = 42
        _ = fi4
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 32767}
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = true
        fmt.Println(v1)
        fmt.Println(v2)
        _ = fi10
        _ = v11
    }
    v1 += float32(((1e10 - (1e10 * 1.5)) + (v2 + float64(v1))))
    v1 = 0.1
    v12 := GenS1{F0: v9, F1: v9}
    fmt.Println(v9, v9, v1)
    var v13 float64 = v2
    delete(v3, 32767)
    v9 = 255
    v9 = (v9 - (v9 ^ (^(v9))))
    var v14 int8 = 0
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: -2147483648, F1: 0, F2: false}
        _ = fi15
        _ = v16
    }
    var v17 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
