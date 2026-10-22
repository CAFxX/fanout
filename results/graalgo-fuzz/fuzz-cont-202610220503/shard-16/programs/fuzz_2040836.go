package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 int, p1 int16, p2 string) uint64 {
    p0 = (p0 << ((2) & 63))
    var v4 bool = false
    for fi5 := 0; fi5 < 19; fi5++ {
        if (v4 || ((2 < 1) || (!true))) {
            v6 := GenS1{F0: 2147483648, F1: 0}
            _ = v6
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 126
}
func genFn1() uint64 {
    v7 := []int{-1, 53, 9223372036854775807, 729}
    v8 := GenS1{F0: 2147483648, F1: 52}
    var v9 float32 = ((3.14159 + ((-2.25 * float32(float64((0.1 - 1.5)))) * float32(3.14159))) / (float32(0.1) / 1e10))
    v7 = append(v7, -201)
    _ = v7
    _ = v8
    _ = v9
    return 110
}

func main() {
    v1 := []string{}
    v2 := []uint32{2147483647, 2147483648, 4294967295, 42}
    var v3 float32 = -2.25
    v3 += v3
    fmt.Println(v3)
    if (-84 >= -128) {
        var v10 bool = (!(1.5 >= v3))
        _ = v10
    } else {
        v3 = 1e10
    }
    var v11 uint16 = 32767
    if false {
        v12 := GenS2{F0: (true || (388 == -1)), F1: -9223372036854775808}
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS3{F0: v11, F1: 65535}
            _ = fi13
            _ = v14
        }
        switch 2147483647 {
        case 698:
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := GenS3{F0: v11, F1: 2}
                if len(v2) > 0 {
                    v2[0] = 1
                }
                _ = fi15
                _ = v16
            }
        case -1:
            v17 := GenS1{F0: 258, F1: 2}
            _ = v17
        default:
            _ = len(v1)
        }
        _ = v12
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v11 = 0
        fmt.Println(fi18, fi18)
        v11 -= v11
        _ = fi18
    }
    if len(v2) > 0 {
        v2[0] = 910
    }
    _ = len(v2)
    if true {
        v3 -= (0.1 / 1.0)
        v3 = float32(3.14159)
        _ = len(v2)
    } else {
        v19 := GenS3{F0: 0, F1: 860}
        var v20 int16 = -964
        _ = v19
        _ = v20
    }
    v11 = v11
    fmt.Println(v11, v11)
    v3 = ((float32(0.0) - 0.1) - 0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
