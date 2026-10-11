package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint; F2 uint64 }

func main() {
    var v1 bool = (-1 <= -2147483648)
    var v2 int32 = 1
    var v3 int8 = -128
    var v4 int16 = 42
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint8 = 42
        v6 -= 42
        _ = fi5
        _ = v6
    }
    var v7 int = 42
    v8 := GenS1{F0: (!(128 <= 1)), F1: (v7 < v7), F2: v1}
    v9 := []int{}
    fmt.Println(v3)
    v4 -= v4
    fmt.Println(v3, v1)
    v7 -= (v7 >> ((255) & 63))
    v10 := GenS1{F0: (3.14159 < 1e10), F1: (v1 && v1), F2: (!(-2.25 != 1.5))}
    if len(v9) > 0 {
        v9[0] = 108
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v7 = (-(v7))
        v12 := GenS2{F0: 2, F1: (v4 % ((v4) | 1)), F2: 229}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if (true && v1) {
                v15 := GenS2{F0: 0, F1: 1, F2: 623}
                v16 := GenS3{F0: v2, F1: 1, F2: 2}
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: true, F1: false, F2: v1}
            _ = fi14
            _ = v17
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
