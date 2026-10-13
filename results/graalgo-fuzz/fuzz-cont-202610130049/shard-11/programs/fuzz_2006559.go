package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0() uint32 {
    if (531 <= 42) {
        var v5 uint32 = 65535
        _ = v5
    }
    switch -1 {
    case 32767:
        var v6 int16 = -1
        _ = v6
    }
    v7 := GenS1{F0: 32767, F1: 255}
    v8 := []int{128, 4294967295, -9223372036854775808, -2147483648}
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int64, p1 int8, p2 bool) uint8 {
    var v9 uint64 = 0
    switch 2147483647 {
    case 4294967295:
        p2 = p2
    case 128:
        var v10 uint8 = 42
        _ = v10
    case 32768:
        v11 := GenS1{F0: 32767, F1: 993}
        _ = v11
    default:
        v12 := []int{42}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    var v1 string = "a"
    var v2 int32 = -99
    var v3 uint32 = 278
    var v4 bool = false
    if (0 >= 255) {
        v2 -= (v2 << ((187) & 63))
        fmt.Println(v4, v4)
        v2 = (v2 - v2)
    }
    v13 := []int{330}
    var v14 float32 = 0.1
    v3 += 77
    for fi15 := 0; fi15 < 14; fi15++ {
        v1 = "\n\t"
        v2 -= -1
        if v4 {
            v16 := GenS1{F0: 643, F1: 65535}
            var v17 uint32 = v3
            _ = v16
            _ = v17
        } else {
            if len(v13) > 0 {
                v13[0] = 32767
            }
            fmt.Println(v2)
            if (float32(((1.5 / 1e10) / 1.0)) > (v14 * -2.25)) {
                if (v4 && v4) {
                    var v18 float64 = 0.1
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v4 = (32767 < -1)
                        var v20 uint64 = 1
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 int64 = 9223372036854775807
                        var v23 uint = 4294967295
                        fmt.Println(v18, fi15, v18)
                        v1 = "hello"
                        var v24 bool = ((v2 - (v2 | v2)) != (-(v2)))
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v18
                }
            } else {
                var v25 float32 = (v14 - 3.14159)
                _ = v25
            }
        }
        _ = fi15
    }
    v26 := GenS1{F0: 514, F1: 1}
    v3 += 42
    for fi27 := 0; fi27 < 5; fi27++ {
        v4 = (!false)
        var v28 string = v1
        var v29 int32 = (v2 ^ v2)
        _ = fi27
        _ = v28
        _ = v29
    }
    v4 = ((396 >= 42) || v4)
    _ = len(v13)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
