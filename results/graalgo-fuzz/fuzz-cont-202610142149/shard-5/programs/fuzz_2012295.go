package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 bool = (!true)
    v2 := []uint32{2147483647, 42, 188}
    var v3 int8 = 127
    v4 := []int{-2147483648, 0, 2147483647}
    v5 := []uint32{569}
    for fi6 := 0; fi6 < 20; fi6++ {
        v3 -= v3
        _ = fi6
    }
    v7 := GenS1{F0: 0, F1: 2}
    switch ("a" + "hello") {
    case "hello":
        switch 547 {
        case -222:
            var v8 int64 = 0
            _ = v8
        case 322:
            v9 := GenS2{F0: 453, F1: 154}
            _ = v9
        case 2147483647:
            if v1 {
                _ = len(v2)
                v10 := GenS1{F0: 877, F1: 18446744073709551615}
                v11 := GenS2{F0: 1, F1: -944}
                _ = v10
                _ = v11
            } else {
                v1 = (!(("a" > "\n\t") && false))
            }
        default:
            switch "" {
            case "a":
                var v12 int32 = 0
                _ = v12
            case "hello":
                v1 = false
            default:
                v1 = true
            }
        }
    default:
        if v1 {
            v13 := GenS3{F0: 4294967295}
            fmt.Println(v3, v3)
            _ = v13
        }
    }
    v14 := GenS4{F0: (("hello" + (("a" + "a") + "foo bar")) + "foo bar")}
    v2 = append(v2, 976)
    if v1 {
        var v15 uint32 = 65535
        v16 := GenS1{F0: 255, F1: 1}
        _ = len(v2)
        _ = v15
        _ = v16
    }
    switch 42 {
    case -948:
        var v17 int8 = (v3 % ((0) | 1))
        _ = v17
    case 32767:
        fmt.Println(v3, v1)
    default:
        var v18 int8 = (v3 & 1)
        _ = v18
    }
    if v1 {
        v19 := GenS4{F0: "a"}
        _ = v19
    } else {
        v3 -= v3
        if (v1 || false) {
            if v1 {
                if false {
                    v1 = (146 < -521)
                    fmt.Println(v3, v3)
                    for fi20 := 0; fi20 < 17; fi20++ {
                        if v1 {
                            var v21 uint64 = 523
                            v3 = v3
                            v22 := GenS4{F0: "\n\t"}
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(fi20, fi20)
                            if (float32((3.14159 / 3.14159)) <= float32(float64(-2.25))) {
                                fmt.Println(v3)
                            }
                            var v23 float32 = 1e10
                            _ = v23
                        }
                        _ = fi20
                    }
                } else {
                    v4 = append(v4, -1)
                    v24 := GenS5{F0: 0}
                    _ = v24
                }
                var v25 int16 = -63
                v1 = true
                _ = v25
            }
            var v26 uint8 = 66
            fmt.Println(v26)
            _ = v26
        } else {
            fmt.Println(v1, v1)
            v3 = (v3 >> ((372) & 63))
        }
    }
    var v27 uint8 = 127
    v28 := GenS4{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
