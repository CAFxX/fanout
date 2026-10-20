package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 42
    var v2 uint16 = 788
    v3 := make(map[int]int)
    var v4 int = -188
    var v5 int8 = 100
    v6 := GenS1{F0: 1}
    var v7 string = ""
    v7 = (v7 + "\n\t")
    v8 := GenS2{F0: (v7 + v7), F1: -475, F2: ((float64((0.1 * -2.25)) * (1e10 + (1e10 / 1.0))) < 1e10)}
    v7 = v7
    if true {
        var v9 int = (v4 % (((^(v1))) | 1))
        v2 -= (v2 << ((297) & 63))
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{fi10, v4}
        var v12 int16 = 600
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint8 = 128
            v15 := GenS3{F0: 42, F1: 850}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v16 uint32 = 2147483648
    if false {
        v4 -= v1
        var v17 uint16 = (v2 - v2)
        v18 := GenS4{F0: v7}
        _ = v17
        _ = v18
    } else {
        var v19 int64 = 0
        v20 := GenS1{F0: 0}
        _ = v19
        _ = v20
    }
    var v21 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
