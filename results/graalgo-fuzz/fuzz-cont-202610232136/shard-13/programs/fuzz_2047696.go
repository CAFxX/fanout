package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
func genFn0() uint8 {
    var v5 int = -1000
    fmt.Println(v5)
    v5 += (-(v5))
    var v6 uint64 = 42
    _ = v5
    _ = v6
    return 18
}

func main() {
    var v1 uint32 = 20
    v2 := []uint32{(v1 | 42), 2147483648, 0}
    var v3 string = "a"
    v4 := make(map[int]int)
    v3 = "\n\t"
    v1 = 2147483648
    if (((0.1 * 1.5) < 3.14159) || (!false)) {
        v7 := GenS1{F0: 70, F1: false, F2: -32768}
        v1 += (^(v1))
        var v8 int64 = 42
        _ = v7
        _ = v8
    } else {
        switch 4294967295 {
        case 42:
            v9 := GenS1{F0: 128, F1: (((!((0.1 / 1e10) <= 0.1)) && true) || false), F2: -105}
            _ = v9
        case 973:
            v1 = 2147483647
        case 1:
            for fi10 := 0; fi10 < 18; fi10++ {
                v1 += (v1 >> ((255) & 63))
                _ = fi10
            }
        default:
            fmt.Println(v3, v1)
        }
    }
    v3 = v3
    var v11 float64 = (0.1 / (3.14159 - 0.1))
    v12 := GenS2{F0: 128, F1: 42, F2: -933}
    if len(v2) > 0 {
        v2[0] = (^(v1))
    }
    if (1 < 32767) {
        if true {
            if (-108 > 42) {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := GenS2{F0: 0, F1: 320, F2: -9223372036854775808}
                    fmt.Println(v3)
                    fmt.Println(v1, v3)
                    _ = fi13
                    _ = v14
                }
            }
            var v15 float64 = ((v11 * -2.25) * v11)
            if (!(!false)) {
                for fi16 := 0; fi16 < 18; fi16++ {
                    fmt.Println(v11, v15, v1)
                    _ = fi16
                }
                _ = len(v2)
            } else {
                v17 := GenS1{F0: 1, F1: (170 != 67), F2: -677}
                _ = v17
            }
            _ = v15
        }
        v11 = (0.0 * v11)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
