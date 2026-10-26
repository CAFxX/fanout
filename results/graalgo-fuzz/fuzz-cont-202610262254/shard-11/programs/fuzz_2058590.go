package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 2
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 11; fi4++ {
        v2 = (v2 % ((v2) | 1))
        var v5 float32 = (float32(3.14159) * (0.0 - 0.0))
        fmt.Println(fi4, fi4)
        _ = fi4
        _ = v5
    }
    v2 = (v2 % ((v2) | 1))
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            v2 = 1
            switch "\n\t" {
            case "":
                _ = len(v1)
            case "a":
                v2 = (v2 & 38)
            default:
                if false {
                    var v7 uint64 = 123
                    _ = v7
                }
            }
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := []int{fi6, fi8, 4294967295}
                v10 := GenS1{F0: 254, F1: v2}
                v11 := GenS2{F0: 1, F1: 18446744073709551615, F2: 2147483648}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
        _ = fi6
    }
    var v12 int = 0
    switch (v12 * v12) {
    case 800:
        v12 += 1
    case 257:
        v13 := GenS1{F0: -228, F1: 255}
        _ = v13
    case 538:
        v14 := GenS1{F0: 1, F1: (v2 ^ (v2 ^ (v2 >> ((42) & 63))))}
        _ = v14
    default:
        var v15 float64 = float64(((((-2.25 - 0.1) * 3.14159) - -2.25) * 0.0))
        _ = v15
    }
    v16 := []int{32767, v12, v12}
    v12 -= v12
    v2 += 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
}
