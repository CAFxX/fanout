package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 uint8; F1 int8 }
func genFn0(p0 int, p1 string) uint16 {
    p0 -= 128
    fmt.Println(p0, p1)
    p0 = 32767
    for fi4 := 0; fi4 < 10; fi4++ {
        p1 = p1
        var v5 int64 = 9223372036854775807
        _ = fi4
        _ = v5
    }
    var v6 float64 = (1.5 * ((1e10 / ((0.0 + (-2.25 / (1e10 * 0.1))) + -2.25)) * -2.25))
    _ = p0
    _ = p1
    _ = v6
    return 32767
}
func genFn1(p0 int32) uint64 {
    var v7 float32 = 3.14159
    var v8 float64 = 1e10
    fmt.Println(v8, v7, p0)
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = (((0.0 - 0.1) + 1e10) * float32(1.5))
    var v2 uint8 = 85
    var v3 int32 = 0
    var v9 int = 42
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 129}
        fmt.Println(v1)
        _ = fi10
        _ = v11
    }
    v12 := []int{(^(v9))}
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v1)
        v14 := GenS2{F0: (-129 != -1), F1: v2}
        switch v9 {
        case 2147483647:
            fmt.Println(v9)
        case 1:
            fmt.Println(v9)
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v2, v9, v9)
    for fi15 := 0; fi15 < 9; fi15++ {
        if true {
            v16 := GenS1{F0: 42}
            _ = v16
        } else {
            var v17 string = ""
            v18 := GenS2{F0: (v2 >= (v2 | (v2 / ((v2) | 1)))), F1: v2}
            _ = v17
            _ = v18
        }
        if ((true || true) || false) {
            if len(v12) > 0 {
                v12[0] = (^(fi15))
            }
        } else {
            var v19 int32 = (v3 * 42)
            v19 = (^(v3))
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint32 = 212
                v22 := GenS1{F0: uint(v21)}
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v19
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            switch (v9 % (((fi15 / ((fi15) | 1))) | 1)) {
            case 9223372036854775807:
                v2 = v2
            case 32767:
                v24 := GenS3{F0: -324, F1: (v3 ^ (v3 - (v3 >> ((255) & 63)))), F2: (!true)}
                _ = v24
            case 42:
                v25 := GenS4{F0: 127, F1: -128}
                _ = v25
            default:
                switch "a" {
                case "hello":
                    if (int32(fi15) >= (v3 | (-(v3)))) {
                        var v26 int16 = -32768
                        _ = v26
                    } else {
                        fmt.Println(fi15, fi23)
                        v12 = append(v12, fi23)
                    }
                case "":
                    switch (fi15 << ((18446744073709551615) & 63)) {
                    case 4294967295:
                        var v27 int8 = 1
                        _ = v27
                    case 1:
                        var v28 int8 = -1
                        _ = v28
                    case 128:
                        _ = len(v12)
                    default:
                        v29 := GenS1{F0: 42}
                        _ = v29
                    }
                case "\n\t":
                    v12 = append(v12, 1)
                default:
                    var v30 float64 = -2.25
                    _ = v30
                }
            }
            _ = len(v12)
            _ = fi23
        }
        _ = fi15
    }
    if (!true) {
        v2 = 42
    }
    v3 = v3
    var v31 uint = 65535
    fmt.Println(v9, v9, v1)
    fmt.Println(v9, v9, v31)
    v32 := GenS1{F0: v31}
    fmt.Println(v9)
    var v33 int = v9
    if (!true) {
        v31 += (v31 / ((v31) | 1))
        if false {
            for fi34 := 0; fi34 < 5; fi34++ {
                v2 = (v2 & 100)
                _ = len(v12)
                v33 -= fi34
                _ = fi34
            }
            v9 -= v33
            if true {
                switch v9 {
                case 1:
                    fmt.Println(v31)
                case 32768:
                    v1 = ((v1 + float32((-2.25 * -2.25))) - (v1 - (float32(1.5) * v1)))
                case 127:
                    var v35 int64 = -277
                    _ = v35
                default:
                    switch (("" + "") + "\n\t") {
                    case "foo bar":
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    default:
                        switch (v9 | v33) {
                        case -159:
                            v33 += -453
                        case -9223372036854775808:
                            var v36 uint8 = 0
                            _ = v36
                        default:
                            if len(v12) > 0 {
                                v12[0] = v33
                            }
                        }
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v9
}
