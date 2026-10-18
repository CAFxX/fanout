package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0(p0 bool, p1 uint64) int {
    p0 = p0
    var v4 uint8 = 128
    p0 = (!(128 > v4))
    _ = p0
    _ = p1
    _ = v4
    return -373
}

func main() {
    v1 := []int{478}
    var v2 string = ("hello" + "a")
    var v3 int8 = 127
    fmt.Println(v3, v2)
    v3 += v3
    if ((!false) && true) {
        v2 = v2
        if len(v1) > 0 {
            v1[0] = 127
        }
    } else {
        v5 := GenS1{F0: 2, F1: -9223372036854775808}
        if (-2.25 != 1e10) {
            v6 := GenS2{F0: 1, F1: 620}
            switch v2 {
            case "foo bar":
                v7 := GenS3{F0: "foo bar"}
                _ = v7
            case "\n\t":
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 float64 = 3.14159
                    _ = fi8
                    _ = v9
                }
            case "hello":
                if ((false && false) && ((-2.25 < -2.25) || (527 >= 255))) {
                    fmt.Println(v2, v3)
                    if (!true) {
                        var v10 uint16 = 2
                        _ = v10
                    } else {
                        var v11 int8 = (v3 - 100)
                        v2 = v2
                        _ = v11
                    }
                    switch v2 {
                    case "a":
                        v3 -= v3
                    case "\n\t":
                        var v12 bool = false
                        _ = v12
                    default:
                        if len(v1) > 0 {
                            v1[0] = 9223372036854775807
                        }
                    }
                } else {
                    var v13 int = -1
                    var v14 uint = 255
                    _ = v13
                    _ = v14
                }
            default:
                _ = len(v1)
            }
            v15 := GenS4{F0: v2}
            _ = v15
            _ = v6
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v5
    }
    if false {
        v16 := GenS5{F0: 176, F1: ((v3 / ((v3) | 1)) <= v3)}
        var v17 uint64 = 2
        var v18 uint64 = (v17 / (((^(v17))) | 1))
        _ = v16
        _ = v17
        _ = v18
    }
    if false {
        var v19 int32 = 0
        v19 = 0
        _ = v19
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    var v20 int32 = 0
    if false {
        v3 += (v3 % (((v3 & v3)) | 1))
        v3 = v3
    }
    var v21 float32 = 0.0
    var v22 bool = true
    v23 := GenS2{F0: 42, F1: 18446744073709551615}
    if len(v1) > 0 {
        v1[0] = 0
    }
    v20 = (v20 | 267)
    if (!(3.14159 != (0.1 + 3.14159))) {
        switch "\n\t" {
        case "a":
            v24 := GenS6{F0: 39, F1: ((v3 / (((v3 / ((v3) | 1))) | 1)) == (v3 >> ((0) & 63)))}
            _ = v24
        case "foo bar":
            fmt.Println(v2, v21, v21)
        default:
            var v25 uint16 = 581
            _ = v25
        }
        _ = len(v1)
        var v26 int64 = 9223372036854775807
        _ = v26
    }
    v27 := GenS3{F0: v2}
    v28 := GenS4{F0: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
}
