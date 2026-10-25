package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int64, p1 uint64, p2 int32) uint64 {
    fmt.Println(p2)
    fmt.Println(p0)
    var v6 int = -9223372036854775808
    for fi7 := 0; fi7 < 13; fi7++ {
        switch "a" {
        case "a":
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 uint8 = 2
                v10 := []int{fi8, v6}
                _ = fi8
                _ = v10
                _ = v9
            }
        case "foo bar":
            var v11 uint64 = (^(p1))
            _ = v11
        default:
            fmt.Println(v6, p0)
        }
        v12 := GenS1{F0: false}
        _ = fi7
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1(p0 uint16, p1 int) bool {
    var v13 int16 = 457
    if true {
        if (-9223372036854775808 > -797) {
            var v14 int = p1
            for fi15 := 0; fi15 < 5; fi15++ {
                v13 = v13
                v13 -= v13
                _ = fi15
            }
            var v16 float32 = 1.5
            _ = v14
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return true
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    var v4 uint = 1
    var v5 bool = (!((-32768 != -536) || (uint16(v4) == 449)))
    v4 = v4
    v5 = false
    var v17 float64 = (0.0 + (1e10 / (0.1 / (0.1 + 1.5))))
    v2 = (-(v2))
    var v18 float32 = (0.0 + float32((v17 - v17)))
    fmt.Println(v2, v1)
    switch ("a" + "hello") {
    case "\n\t":
        v19 := GenS1{F0: false}
        _ = v19
    case "a":
        switch 4294967295 {
        case 32767:
            if v5 {
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 float32 = 1.5
                    _ = fi20
                    _ = v21
                }
                var v22 uint8 = 42
                _ = v22
            } else {
                if v5 {
                    var v23 uint64 = 42
                    v23 += genFn0(v2, (v23 + genFn0(v2, v23, 1)), 42)
                    var v24 uint16 = 163
                    _ = v23
                    _ = v24
                }
                v1 += v1
                var v25 uint32 = 42
                _ = v25
            }
        case 4294967295:
            fmt.Println(v5, v18, v18)
        case 128:
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 float32 = float32(v17)
                _ = fi26
                _ = v27
            }
        default:
            var v28 uint8 = 1
            _ = v28
        }
    case "foo bar":
        v29 := GenS1{F0: v5}
        _ = v29
    default:
        for fi30 := 0; fi30 < 1; fi30++ {
            v17 += (float64((v18 / ((3.14159 + -2.25) / (float32((v17 - v17)) * 0.1)))) - v17)
            fmt.Println(v2, v5, v5)
            _ = fi30
        }
    }
    var v31 int32 = -206
    v32 := []int{431, 2147483647}
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    v33 := GenS1{F0: v5}
    switch 42 {
    case 9223372036854775807:
        var v34 int8 = (v1 / ((v1) | 1))
        _ = v34
    case -2147483648:
        v3[1] = -40
    }
    var v35 float32 = (((v18 / v18) - (v18 + (3.14159 / ((1.5 * (float32((v17 * 0.1)) / v18)) - v18)))) * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
