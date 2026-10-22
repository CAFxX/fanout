package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8) int8 {
    p0 = 0
    if (101 != 42) {
        v4 := GenS1{F0: 2147483647, F1: ((p0 | (p0 | (p0 % ((-128) | 1)))) != -19)}
        if true {
            v5 := []int{32767}
            _ = v5
        } else {
            v6 := GenS1{F0: -2147483648, F1: true}
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := []int{(fi7 >> ((2) & 63))}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        v9 := GenS1{F0: -1, F1: true}
        _ = v4
        _ = v9
    } else {
        var v10 string = "a"
        var v11 uint64 = 18446744073709551615
        var v12 int64 = -9223372036854775808
        _ = v10
        _ = v11
        _ = v12
    }
    if (!(!false)) {
        switch 4294967295 {
        case 128:
            p0 -= (^(p0))
        default:
            v13 := []int{1, -254}
            _ = v13
        }
        var v14 int64 = 76
        _ = v14
    }
    var v15 uint16 = 174
    _ = p0
    _ = v15
    return (p0 >> ((uint(v15)) & 63))
}

func main() {
    v1 := []int{}
    var v2 float64 = (3.14159 - 0.1)
    var v3 int8 = 1
    v3 = v3
    switch -2147483648 {
    case 62:
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(v3, v2)
            switch ("" + "hello") {
            case "foo bar":
                var v17 int64 = 9223372036854775807
                _ = v17
            case "\n\t":
                v18 := GenS2{F0: -322}
                _ = v18
            case "hello":
                if true {
                    var v19 int16 = -860
                    switch 316 {
                    case 851:
                        for fi20 := 0; fi20 < 19; fi20++ {
                            if true {
                                v21 := GenS2{F0: -1}
                                _ = len(v1)
                                _ = len(v1)
                                _ = v21
                            }
                            _ = fi20
                        }
                    case -324:
                        v2 += (0.1 * (v2 * v2))
                    }
                    _ = v19
                } else {
                    var v22 uint32 = 996
                    v22 -= 2
                    var v23 float32 = (-2.25 + 0.1)
                    _ = v22
                    _ = v23
                }
            default:
                var v24 uint8 = 2
                _ = v24
            }
            _ = fi16
        }
    case 4294967295:
        if true {
            v2 = v2
        } else {
            var v25 uint = 1
            _ = v25
        }
    case -2147483648:
        fmt.Println(v3, v3)
    default:
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS3{F0: 9223372036854775807}
            _ = fi26
            _ = v27
        }
    }
    if (false && (154 > 0)) {
        var v28 uint = 0
        _ = v28
    }
    var v29 int8 = (v3 - v3)
    fmt.Println(v29)
    v2 += v2
    fmt.Println(v29)
    if (255 <= 0) {
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 9; fi31++ {
                v29 -= 60
                fmt.Println(v29)
                _ = fi31
            }
            _ = fi30
        }
        var v32 bool = true
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
}
