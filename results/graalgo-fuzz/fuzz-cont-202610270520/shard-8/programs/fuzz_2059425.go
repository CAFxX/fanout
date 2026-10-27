package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 67
    var v3 float64 = -2.25
    var v4 uint64 = 4294967295
    var v5 int32 = -2147483648
    var v6 int32 = 0
    var v7 uint16 = v2
    var v8 uint8 = 136
    v3 -= (float64(0.1) + 3.14159)
    if (true || false) {
        v7 = v7
    } else {
        v9 := GenS1{F0: 42, F1: -1, F2: 965}
        _ = v9
    }
    var v10 uint16 = v7
    v10 += 65535
    if true {
        v7 = (v10 & (^(v2)))
        if (232 < 223) {
            var v11 uint8 = v8
            _ = v11
        } else {
            var v12 bool = true
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 int32 = (v6 << ((uint(v8)) & 63))
            v14 = (-(v5))
            switch fi13 {
            case 32768:
                var v15 int32 = v14
                _ = v15
            case 0:
                if ((v6 >= v5) && ((v3 + (v3 - (v3 - float64(((1.5 - ((1e10 * -2.25) - (1.5 * (0.1 + 1e10)))) + 3.14159))))) != (v3 / v3))) {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := v1[832]
                        _ = fi16
                        _ = v17
                    }
                    v4 -= (^(v4))
                    for fi18 := 0; fi18 < 11; fi18++ {
                        _ = len(v1)
                        fmt.Println(fi13)
                        v8 = v8
                        _ = fi18
                    }
                }
            case 1:
                v19 := v1[4294967295]
                _ = v19
            default:
                v20 := GenS2{F0: 32767, F1: 1}
                _ = v20
            }
            _ = fi13
            _ = v14
        }
    }
    v6 = v6
    for fi21 := 0; fi21 < 4; fi21++ {
        if false {
            if (v5 > (v6 << ((18446744073709551615) & 63))) {
                fmt.Println(v7, v4)
            }
        }
        _ = fi21
    }
    v6 = -1
    v5 = -2147483648
    switch 128 {
    case 0:
        fmt.Println(v2, v7, v3)
    case 9223372036854775807:
        var v22 float64 = v3
        _ = v22
    default:
        v3 += 1e10
    }
    v1[42] = 2147483647
    v4 -= 486
    switch -2147483648 {
    case 128:
        v23 := v1[540]
        _ = v23
    default:
        v24 := GenS1{F0: 255, F1: -49, F2: 0}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
