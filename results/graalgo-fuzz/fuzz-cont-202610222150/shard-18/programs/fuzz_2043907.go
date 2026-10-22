package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int = -740
    v2 := make(map[uint32]uint32)
    v3 := []int{(-(v1)), (v1 / (((v1 - (v1 & (v1 >> ((0) & 63))))) | 1)), (-(v1))}
    var v4 int32 = -736
    v5 := make(map[int]string)
    if true {
        if false {
            v6 := GenS1{F0: 32767, F1: true}
            var v7 uint = 2
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int32 = (^(v4))
                if true {
                    fmt.Println(v4)
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(v4, v1)
            var v10 int16 = 42
            _ = v10
        }
    }
    v4 = v4
    var v11 uint16 = 354
    if ((-1 < -128) && true) {
        switch (v1 & -42) {
        case 1:
            var v12 uint16 = v11
            _ = v12
        default:
            v4 += (v4 - v4)
        }
        delete(v5, (v1 | v1))
    }
    v13 := GenS2{F0: 665, F1: 88}
    delete(v5, 4294967295)
    var v14 uint32 = 2
    var v15 uint16 = v11
    fmt.Println(v4, v15, v14)
    v15 -= 99
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
