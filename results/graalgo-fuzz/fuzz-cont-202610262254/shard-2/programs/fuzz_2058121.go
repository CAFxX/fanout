package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int64; F2 int16 }
func genFn0(p0 bool) uint {
    fmt.Println(p0)
    if (-48 >= 127) {
        v6 := []int{-2147483648, 128, 9223372036854775807, 0}
        _ = len(v6)
        if len(v6) > 0 {
            v6[0] = 32768
        }
        _ = v6
    }
    if p0 {
        p0 = p0
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = false
            p0 = p0
            _ = fi7
        }
        if p0 {
            p0 = (true || p0)
        } else {
            var v8 uint64 = 42
            _ = v8
        }
    } else {
        v9 := []int{-838, 944, -118, -897}
        v9 = append(v9, 127)
        _ = v9
    }
    _ = p0
    return 877
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 4294967295
    v3 := make(map[int]uint32)
    var v4 float64 = float64(1e10)
    var v5 int32 = 2147483647
    v10 := []int{(^(v2))}
    for fi11 := 0; fi11 < 12; fi11++ {
        v1[(fi11 | v2)] = 259
        _ = fi11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := v1[0]
        _ = len(v10)
        v14 := GenS1{F0: true, F1: -128}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = len(v10)
    var v15 int64 = -9223372036854775808
    v15 = v15
    switch (-(v2)) {
    case 127:
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                if len(v10) > 0 {
                    v10[0] = fi17
                }
                if true {
                    v2 = fi17
                    _ = len(v1)
                } else {
                    if ((v4 * v4) == 1e10) {
                        var v18 uint64 = 42
                        _ = v18
                    } else {
                        v2 = (fi17 & 32767)
                        var v19 uint64 = 42
                        var v20 int = -2147483648
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(v15, v15)
                    switch "a" {
                    case "a":
                        fmt.Println(v2)
                    case "":
                        var v21 uint16 = 32767
                        _ = v21
                    default:
                        delete(v3, 9223372036854775807)
                    }
                }
                v22 := GenS2{F0: 255}
                _ = fi17
                _ = v22
            }
            _ = fi16
        }
    case 155:
        if len(v10) > 0 {
            v10[0] = (v2 % (((v2 ^ -150)) | 1))
        }
    }
    v23 := GenS2{F0: 65535}
    v24 := GenS3{F0: (v2 * (-(v2)))}
    fmt.Println(v15, v2)
    fmt.Println(v2, v15, v4)
    v25 := GenS4{F0: 128, F1: 0, F2: -32768}
    v4 -= 0.0
    var v26 int64 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
