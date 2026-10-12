package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0() uint8 {
    var v7 uint = 65535
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint16 = 42
        _ = fi8
        _ = v9
    }
    _ = v7
    return 2
}
func genFn1() int64 {
    v10 := []int{32768, 42, 32768, -304}
    var v11 int16 = 32767
    if (false || (42 <= -9223372036854775808)) {
        var v12 string = "a"
        _ = v12
    } else {
        fmt.Println(v11)
        v13 := GenS1{F0: 1, F1: 0}
        _ = v13
    }
    _ = v10
    _ = v11
    return -9223372036854775808
}

func main() {
    var v1 int64 = -285
    var v2 int32 = 54
    v3 := []int{-601}
    var v4 int8 = 100
    v5 := []int{-2147483648, 9223372036854775807, 128, 370}
    var v6 float64 = 0.1
    var v14 uint32 = 558
    v15 := GenS1{F0: v14, F1: v14}
    if (((0 >= 1) && false) || (4294967295 == 0)) {
        var v16 uint8 = 154
        if (42 == 838) {
            for fi17 := 0; fi17 < 4; fi17++ {
                fmt.Println(v1, v16, v6)
                v2 -= 42
                _ = fi17
            }
            if (true || (0 < 127)) {
                var v18 uint32 = v14
                fmt.Println(v1, v2)
                _ = v18
            } else {
                switch 878 {
                case 32768:
                    fmt.Println(v14, v16)
                case -2147483648:
                    v16 += v16
                default:
                    switch 128 {
                    case -9223372036854775808:
                        var v19 bool = true
                        _ = v19
                    default:
                        v6 = (0.1 * 0.0)
                    }
                }
            }
            v6 += float64(1e10)
        }
        _ = v16
    }
    _ = len(v3)
    if (!false) {
        fmt.Println(v4)
    }
    if true {
        if true {
            switch int(v4) {
            case 9223372036854775807:
                switch 807 {
                case 128:
                    v14 += v14
                }
            case 127:
                var v20 uint8 = genFn0()
                _ = v20
            case -1:
                if len(v5) > 0 {
                    v5[0] = 281
                }
            default:
                v14 += 4294967295
            }
        }
        var v21 int8 = (v4 * (v4 & (v4 - v4)))
        _ = v21
    } else {
        if true {
            switch 127 {
            case 42:
                var v22 uint32 = 65535
                _ = v22
            }
        }
        var v23 int32 = int32(v1)
        _ = v23
    }
    v24 := GenS1{F0: (v14 << ((4294967295) & 63)), F1: (v14 * 4294967295)}
    v25 := GenS1{F0: (v14 ^ v14), F1: v14}
    fmt.Println(v4, v14, v4)
    fmt.Println(v6, v6, v2)
    if (!(!(737 > 216))) {
        var v26 uint8 = 42
        _ = v26
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v27 uint16 = 836
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
