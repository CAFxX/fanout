package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
func genFn0(p0 int8, p1 uint) int16 {
    var v4 float32 = 1e10
    v4 -= float32(0.1)
    v5 := GenS1{F0: 255, F1: 2, F2: 1000}
    fmt.Println(p0, v4)
    v4 += v4
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -32768
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 110
    var v3 int = -545
    var v6 uint = 1
    if ((!false) && true) {
        for fi7 := 0; fi7 < 6; fi7++ {
            v3 += v3
            var v8 int = (fi7 / (((fi7 + (v3 % ((int(v6)) | 1)))) | 1))
            v6 = (v6 - (^(v6)))
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 float64 = 1e10
            if ((!false) && false) {
                var v11 bool = true
                var v12 uint8 = v2
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: 0, F1: true, F2: 0}
            _ = fi9
            _ = v10
            _ = v13
        }
    }
    v2 -= (v2 | 118)
    switch v3 {
    case 9223372036854775807:
        var v14 float32 = 3.14159
        _ = v14
    default:
        v3 = v3
    }
    fmt.Println(v2)
    fmt.Println(v2, v3)
    for fi15 := 0; fi15 < 14; fi15++ {
        if false {
            v16 := GenS2{F0: 1, F1: ((463 >= 255) || true), F2: -928}
            if true {
                var v17 float64 = 1.5
                if false {
                    v6 = 4294967295
                    var v18 uint8 = v2
                    _ = v18
                } else {
                    var v19 int64 = 0
                    _ = v19
                }
                var v20 int16 = 42
                _ = v17
                _ = v20
            } else {
                fmt.Println(v2, fi15, v3)
            }
            _ = v16
        }
        var v21 uint = 18446744073709551615
        _ = fi15
        _ = v21
    }
    var v22 uint32 = 789
    for fi23 := 0; fi23 < 18; fi23++ {
        v3 = fi23
        _ = fi23
    }
    var v24 string = "hello"
    var v25 int = (^(v3))
    v6 -= v6
    fmt.Println(v6)
    var v26 uint = v6
    v6 = v26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v6
}
