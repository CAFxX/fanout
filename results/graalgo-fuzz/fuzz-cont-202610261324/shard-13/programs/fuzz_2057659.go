package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0(p0 int64, p1 uint64, p2 bool) uint32 {
    p0 -= -1
    var v5 int = 0
    var v6 int = v5
    var v7 int64 = p0
    v8 := GenS1{F0: 216, F1: 32767, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint = 975
    var v2 int8 = 100
    var v3 float32 = ((0.0 - -2.25) - (-2.25 / 0.1))
    var v4 uint16 = 42
    fmt.Println(v4)
    var v9 uint = v1
    if false {
        v10 := []int{42, 1, -1}
        var v11 int8 = (v2 >> (((v1 ^ v1)) & 63))
        switch -863 {
        case -2147483648:
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 string = ""
                v2 -= (v2 + v2)
                _ = fi12
                _ = v13
            }
        case -374:
            v10 = append(v10, 466)
        case -9223372036854775808:
            v14 := GenS1{F0: uint16(v9), F1: v4, F2: v11}
            _ = v14
        }
        _ = v10
        _ = v11
    } else {
        v2 += (v2 ^ v2)
    }
    v3 = v3
    v15 := GenS1{F0: (v4 * v4), F1: (v4 ^ (v4 & (^(v4)))), F2: int8(v4)}
    v16 := GenS1{F0: v4, F1: 585, F2: (v2 << ((v1) & 63))}
    v17 := []int{0, 0, 0, 297}
    if (0 <= (v2 % ((v2) | 1))) {
        if ((1.5 < (1.5 * 0.1)) || (970 < 2)) {
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(v2, v9)
                _ = fi18
            }
            if len(v17) > 0 {
                v17[0] = 42
            }
        } else {
            for fi19 := 0; fi19 < 8; fi19++ {
                if false {
                    var v20 uint32 = 4294967295
                    v1 += v9
                    fmt.Println(v4, v2, v20)
                    _ = v20
                } else {
                    _ = len(v17)
                    fmt.Println(v1, v3, v1)
                }
                _ = fi19
            }
            v4 += v4
        }
        switch 127 {
        case 128:
            v21 := GenS1{F0: v4, F1: v4, F2: (v2 << ((18446744073709551615) & 63))}
            _ = v21
        }
        var v22 int8 = (v2 / (((^(v2))) | 1))
        _ = v22
    }
    v23 := GenS1{F0: v4, F1: v4, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
