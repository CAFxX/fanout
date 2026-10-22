package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 1
    var v3 string = "\n\t"
    var v4 uint16 = 0
    v5 := GenS1{F0: 367}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 9223372036854775807}
        for fi8 := 0; fi8 < 3; fi8++ {
            switch fi8 {
            case 536:
                v9 := []int{(fi6 | fi8), -603, (fi6 / (((fi6 << ((255) & 63))) | 1)), 32768}
                _ = v9
            default:
                v10 := []int{32768, (fi6 % ((fi6) | 1)), (fi6 >> ((2) & 63))}
                _ = v10
            }
            v4 -= 84
            v11 := GenS2{F0: -43, F1: v2}
            _ = fi8
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v3 = ""
    if (v4 > (v4 - v4)) {
        var v12 uint32 = 545
        _ = v12
    }
    v2 = v2
    v3 = "\n\t"
    v2 -= (v2 << ((900) & 63))
    fmt.Println(v3, v2, v4)
    var v13 uint8 = 2
    fmt.Println(v3)
    var v14 int = 1
    var v15 uint16 = 583
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS1{F0: 605}
        var v18 uint = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
