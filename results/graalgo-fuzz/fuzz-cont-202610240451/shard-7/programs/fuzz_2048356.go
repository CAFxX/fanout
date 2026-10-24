package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0(p0 int8, p1 bool) int {
    var v5 int8 = (p0 * p0)
    if p1 {
        var v6 int32 = 2147483647
        _ = v6
    }
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 uint = 2
    var v2 uint8 = 127
    var v3 int = 0
    v4 := make(map[int]int)
    if false {
        fmt.Println(v3, v3)
    }
    var v7 int16 = -32768
    switch 2147483647 {
    case -1:
        var v8 int32 = 1
        _ = v8
    default:
        v7 += (v7 * v7)
    }
    if true {
        v9 := []int{(-(v3)), (v3 ^ v3), -2147483648, v3}
        v10 := GenS1{F0: int(v7)}
        _ = v10
        _ = v9
    } else {
        v3 = 9223372036854775807
        switch "" {
        case "hello":
            switch "hello" {
            case "hello":
                v3 = (v3 + 9223372036854775807)
            case "":
                if true {
                    v3 = (^(v3))
                    if true {
                        v11 := []int{}
                        if (!false) {
                            v12 := GenS2{F0: 190, F1: v2}
                            _ = v12
                        }
                        v13 := GenS3{F0: "\n\t", F1: false}
                        _ = v11
                        _ = v13
                    } else {
                        v14 := GenS4{F0: 32768, F1: 18446744073709551615}
                        v1 -= (v1 | (v1 - v1))
                        _ = v14
                    }
                } else {
                    switch genFn0(-128, (300 >= 1)) {
                    case 128:
                        if ((2147483647 <= 1) || false) {
                            v7 += int16(v3)
                            v4[v3] = 269
                        }
                    case 127:
                        if ((^(v2)) >= (^(v2))) {
                            v3 -= v3
                        }
                    }
                    v15 := GenS5{F0: 424, F1: int32(v3), F2: -1}
                    if false {
                        var v16 uint32 = 795
                        var v17 uint32 = v16
                        _ = v16
                        _ = v17
                    } else {
                        _ = len(v4)
                    }
                    _ = v15
                }
            default:
                if false {
                    var v18 float32 = (0.1 / float32(-2.25))
                    fmt.Println(v1, v3)
                    v19 := GenS5{F0: 2147483647, F1: -919, F2: 1}
                    _ = v18
                    _ = v19
                }
            }
        case "a":
            v20 := GenS2{F0: (^(v2)), F1: (v2 / ((1) | 1))}
            _ = v20
        case "foo bar":
            v21 := []int{v3}
            _ = v21
        default:
            v1 = (v1 % (((v1 + v1)) | 1))
        }
    }
    var v22 int64 = 1
    fmt.Println(v22)
    v1 = 247
    var v23 uint16 = 65535
    v23 += uint16(v2)
    if true {
        if false {
            if (980 == 2) {
                var v24 bool = true
                _ = v24
            } else {
                v25 := GenS2{F0: (v2 & (^(v2))), F1: v2}
                _ = v25
            }
            v7 += v7
            fmt.Println(v2)
        }
        v26 := GenS1{F0: (^(v3))}
        _ = v26
    }
    v1 += v1
    var v27 int16 = v7
    v23 += v23
    v2 = (^(v2))
    v22 = int64(v1)
    v28 := GenS4{F0: v3, F1: 2}
    if (v3 != 2147483647) {
        fmt.Println(v27)
    } else {
        fmt.Println(v22)
        if ((146 == (^(v1))) && (!(v2 == (v2 - (v2 << ((v1) & 63)))))) {
            fmt.Println(v23)
            v22 = (v22 * 670)
        }
        _ = len(v4)
    }
    var v29 int8 = -44
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v7
}
