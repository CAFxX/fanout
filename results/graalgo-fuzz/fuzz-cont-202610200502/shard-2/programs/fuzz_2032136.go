package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 uint32; F1 uint; F2 uint32 }

func main() {
    var v1 float32 = float32(1.5)
    var v2 float32 = v1
    var v3 uint8 = 255
    v4 := make(map[uint32]bool)
    v2 -= ((-2.25 + float32((0.1 / -2.25))) * ((0.0 + v2) * (-2.25 + 1.5)))
    if (602 < 65535) {
        var v5 uint8 = (v3 >> ((255) & 63))
        var v6 float64 = -2.25
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        if (100 == -58) {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 888, F1: -19, F2: int(v3)}
                v1 -= v2
                _ = fi8
                _ = v9
            }
            var v10 int64 = 825
            v11 := []int{fi7, (fi7 - 825), fi7, 346}
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    var v12 uint32 = 1
    var v13 uint64 = 0
    _ = len(v4)
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            switch fi14 {
            case -9:
                if (((0 > 0) && false) || false) {
                    var v16 uint = 42
                    _ = v16
                }
            case -9223372036854775808:
                v17 := []int{fi14}
                _ = v17
            default:
                if true {
                    fmt.Println(v3, v13, fi15)
                    fmt.Println(v2)
                } else {
                    v12 -= v12
                }
            }
            _ = fi15
        }
        _ = fi14
    }
    var v18 int64 = -9223372036854775808
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS2{F0: v12, F1: 564, F2: (v12 << ((18446744073709551615) & 63))}
            _ = fi20
            _ = v21
        }
        fmt.Println(v12, v1)
        for fi22 := 0; fi22 < 18; fi22++ {
            if true {
                if ((1e10 <= 1e10) && false) {
                    var v23 bool = false
                    _ = v23
                }
            }
            var v24 string = "a"
            _ = fi22
            _ = v24
        }
        _ = fi19
    }
    v3 -= (v3 | v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
