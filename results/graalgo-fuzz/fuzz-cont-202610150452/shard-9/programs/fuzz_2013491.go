package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 uint16, p1 uint, p2 int32) string {
    fmt.Println(p1, p1, p1)
    var v5 uint8 = uint8(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "\n\t"
}
func genFn1(p0 uint8, p1 uint, p2 string) int {
    p0 = (p0 - 24)
    if false {
        v6 := GenS1{F0: 0}
        var v7 uint16 = 65535
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 100
        p1 -= (p1 + (p1 ^ (p1 & p1)))
        _ = v8
    }
    if false {
        var v9 uint = p1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (-1 < 32767)
    var v3 float64 = 3.14159
    var v4 int32 = 1
    v10 := []int{1, 2147483647, 128}
    if (true || (v4 == 6)) {
        var v11 float32 = float32(v3)
        _ = v11
    } else {
        v3 = v3
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int64 = 9223372036854775807
            for fi14 := 0; fi14 < 4; fi14++ {
                _ = len(v1)
                _ = fi14
            }
            switch genFn1(13, 148, "foo bar") {
            case -9223372036854775808:
                for fi15 := 0; fi15 < 5; fi15++ {
                    if false {
                        var v16 uint32 = 2147483647
                        _ = v16
                    } else {
                        v2 = (v2 || v2)
                        fmt.Println(v13)
                        v2 = v2
                    }
                    if v2 {
                        v4 += -1
                    } else {
                        var v17 int32 = (v4 / (((v4 / ((v4) | 1))) | 1))
                        v3 = v3
                        _ = v17
                    }
                    _ = fi15
                }
            case 2147483647:
                v18 := GenS2{F0: v4, F1: 1, F2: -1}
                _ = v18
            case 32767:
                if (!(v13 < -1)) {
                    switch (fi12 - 32768) {
                    case 379:
                        fmt.Println(v4)
                    case -767:
                        if v2 {
                            switch 765 {
                            case 32768:
                                _ = len(v10)
                            default:
                                if (true || v2) {
                                    fmt.Println(v4, fi12)
                                }
                            }
                        } else {
                            if (false && v2) {
                                var v19 int32 = v4
                                _ = v19
                            }
                            fmt.Println(v3, v13)
                            var v20 string = "foo bar"
                            _ = v20
                        }
                    default:
                        fmt.Println(v4, v2, v4)
                    }
                    var v21 uint8 = 91
                    v10 = append(v10, 128)
                    _ = v21
                }
            }
            _ = fi12
            _ = v13
        }
        v22 := GenS1{F0: (v4 % ((v4) | 1))}
        _ = v22
    }
    if v2 {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int16 = 1
            v3 -= ((0.0 * v3) + 1.5)
            _ = fi23
            _ = v24
        }
    }
    if len(v10) > 0 {
        v10[0] = 9223372036854775807
    }
    switch 754 {
    case -178:
        var v25 uint64 = 18446744073709551615
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(v4, v4)
        var v27 int16 = 0
        _ = fi26
        _ = v27
    }
    var v28 float32 = (0.0 - 1e10)
    var v29 int64 = 0
    if true {
        fmt.Println(v3, v29)
        var v30 bool = false
        v28 += (((v28 / (0.0 * v28)) + (0.1 * v28)) / 3.14159)
        _ = v30
    } else {
        v31 := GenS3{F0: v29, F1: 127, F2: 1}
        fmt.Println(v4)
        _ = v31
    }
    switch 0 {
    case -84:
        var v32 int32 = (v4 >> ((18446744073709551615) & 63))
        _ = v32
    case 32767:
        fmt.Println(v3, v28)
    }
    var v33 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
}
