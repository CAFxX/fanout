package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 string) uint8 {
    v7 := []int{9223372036854775807}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: -9223372036854775808, F1: -621, F2: -2147483648}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 int64) bool {
    p1 = p1
    var v10 uint64 = 0
    _ = p0
    _ = p1
    _ = v10
    return (false || false)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -677
    v3 := make(map[string]int)
    var v4 uint8 = 0
    var v5 int64 = 0
    var v6 float32 = 1.5
    v11 := []int{(v2 / ((4294967295) | 1)), (v2 * v2)}
    fmt.Println(v5, v6)
    fmt.Println(v4)
    fmt.Println(v5, v5)
    if false {
        _ = len(v11)
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                v3["hello"] = (fi12 | v2)
                if (255 > 626) {
                    if ("" != "a") {
                        var v14 int32 = 42
                        fmt.Println(fi12)
                        _ = v14
                    } else {
                        v2 -= v2
                    }
                    var v15 int8 = 1
                    v16 := v1[2147483647]
                    _ = v15
                    _ = v16
                }
                fmt.Println(v6, v4)
                _ = fi13
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 float64 = ((-2.25 + float64(float32((1e10 + float64(v6))))) * 1e10)
                _ = len(v11)
                _ = fi17
                _ = v18
            }
            _ = fi12
        }
        switch v2 {
        case 703:
            v6 = (v6 / v6)
        }
    } else {
        v1[65535] = 85
    }
    if ((v5 & 42) == (v5 << ((4294967295) & 63))) {
        v19 := GenS2{F0: -32768}
        var v20 uint16 = 166
        v2 -= v2
        _ = v19
        _ = v20
    } else {
        if false {
            var v21 uint8 = (v4 + (v4 * (v4 / ((2) | 1))))
            if (32767 != 1) {
                var v22 uint64 = 0
                _ = v22
            }
            _ = v21
        } else {
            v23 := v3["a"]
            fmt.Println(v2, v5, v23)
            _ = v23
        }
    }
    switch v2 {
    case -12:
        var v24 int64 = v5
        _ = v24
    case 42:
        if (!false) {
            _ = len(v11)
        }
    case 127:
        v6 += 0.1
    default:
        v5 = v5
    }
    if (-268 <= (v2 + (v2 << ((18446744073709551615) & 63)))) {
        for fi25 := 0; fi25 < 18; fi25++ {
            _ = len(v11)
            v5 += (v5 >> ((4294967295) & 63))
            _ = fi25
        }
    }
    v5 = v5
    var v26 float32 = 1e10
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
