package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    var v1 int = 42
    v2 := make(map[uint32]string)
    var v3 int16 = 42
    var v4 float32 = -2.25
    var v5 uint32 = 453
    fmt.Println(v1, v4, v4)
    var v6 uint16 = 65535
    fmt.Println(v4)
    fmt.Println(v1, v5, v6)
    var v7 int16 = v3
    _ = len(v2)
    v3 = int16(v5)
    v8 := GenS1{F0: 128, F1: 1}
    v9 := []int{}
    v10 := GenS1{F0: 255, F1: 127}
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: 127, F1: int8(v5)}
        v13 := GenS1{F0: uint8(v6), F1: -115}
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        if ((true && (v6 > (v6 % (((v6 / (((v6 - (v6 ^ (v6 + (^(v6)))))) | 1))) | 1)))) && (fi14 < (fi14 ^ v1))) {
            v3 = (v7 << ((1) & 63))
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS2{F0: v5, F1: -1}
                if false {
                    v1 = fi14
                }
                fmt.Println(v7, fi14)
                _ = fi15
                _ = v16
            }
        }
        _ = fi14
    }
    var v17 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
