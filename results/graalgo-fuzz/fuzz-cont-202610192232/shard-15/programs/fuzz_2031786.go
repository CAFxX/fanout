package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 uint32; F1 string; F2 int64 }
func genFn0() int8 {
    v7 := GenS1{F0: (0.0 >= 1.5), F1: 127, F2: 1000}
    v8 := GenS2{F0: 9223372036854775807, F1: 32767}
    v9 := GenS3{F0: (65535 != 1), F1: 1}
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int32, p1 int16, p2 string) bool {
    switch p2 {
    case "hello":
        var v10 int = -789
        _ = v10
    case "a":
        p0 -= (-(p0))
    default:
        var v11 int64 = 42
        _ = v11
    }
    p1 = 1
    _ = p0
    _ = p1
    _ = p2
    return (1e10 > float64(3.14159))
}

func main() {
    var v1 uint8 = 128
    var v2 string = "hello"
    var v3 uint = 255
    v4 := make(map[int]uint32)
    var v5 float32 = 3.14159
    v6 := []uint32{}
    var v12 int16 = 1000
    if len(v6) > 0 {
        v6[0] = 881
    }
    v13 := GenS4{F0: int(v12), F1: 98}
    v14 := GenS5{F0: 127}
    if true {
        var v15 uint = v3
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 string = v2
            if true {
                v18 := v4[fi16]
                _ = v18
            } else {
                v4[fi16] = 0
                _ = len(v4)
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v1)
        _ = v15
    } else {
        if len(v6) > 0 {
            v6[0] = 0
        }
        fmt.Println(v2, v1)
    }
    v5 += (-2.25 / (((((v5 * 1.5) + 1.5) + (0.1 * 3.14159)) * v5) - (-2.25 / 1.0)))
    var v19 uint = 374
    switch v2 {
    case "":
        v20 := GenS4{F0: 113, F1: 42}
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 3; fi21++ {
            if (0 >= 382) {
                var v22 int = (fi21 ^ 0)
                _ = v22
            } else {
                v1 = (v1 >> (((v3 / ((v3) | 1))) & 63))
                v23 := GenS6{F0: true, F1: 185}
                for fi24 := 0; fi24 < 5; fi24++ {
                    v19 += v3
                    if (42 == 916) {
                        v6 = append(v6, 4294967295)
                        if false {
                            var v25 int64 = 42
                            _ = v25
                        }
                        v26 := GenS7{F0: 0, F1: 65535}
                        _ = v26
                    } else {
                        v1 = (v1 ^ 42)
                        v27 := GenS8{F0: 916, F1: 191}
                        _ = len(v4)
                        v5 -= -2.25
                        v28 := GenS5{F0: 0}
                        _ = v27
                        _ = v28
                    }
                    v4[(fi21 * 30)] = 328
                    _ = fi24
                }
                _ = v23
            }
            var v29 uint = (v3 / ((v3) | 1))
            _ = fi21
            _ = v29
        }
    default:
        v2 = v2
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        fmt.Println(fi30)
        v31 := GenS9{F0: 42, F1: "hello", F2: 42}
        var v32 int16 = v12
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
