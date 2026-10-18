package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 int16, p2 uint8) uint64 {
    if (0 != (p2 >> ((1) & 63))) {
        if false {
            p0 -= uint32(p2)
            if true {
                for fi4 := 0; fi4 < 6; fi4++ {
                    if (0 > 42) {
                        var v5 bool = true
                        if (!false) {
                            fmt.Println(v5)
                            v6 := GenS1{F0: int8(p1), F1: 791}
                            _ = v6
                        }
                        _ = v5
                    } else {
                        if true {
                            var v7 int8 = 42
                            var v8 uint = 65535
                            _ = v7
                            _ = v8
                        } else {
                            if (!true) {
                                if true {
                                    for fi9 := 0; fi9 < 11; fi9++ {
                                        if (false && ((p0 % ((2147483648) | 1)) < (p0 - (p0 - 867)))) {
                                            p0 -= p0
                                            v10 := []int{}
                                            var v11 string = "hello"
                                            _ = v10
                                            _ = v11
                                        }
                                        _ = fi9
                                    }
                                }
                                v12 := []int{fi4, (fi4 * 1)}
                                _ = v12
                            }
                        }
                    }
                    v13 := GenS1{F0: 42, F1: p0}
                    _ = fi4
                    _ = v13
                }
            }
        } else {
            fmt.Println(p0, p1)
        }
        fmt.Println(p2, p2)
        v14 := GenS1{F0: 1, F1: 4294967295}
        _ = v14
    }
    var v15 bool = true
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 256
}
func genFn1(p0 int32, p1 int32, p2 uint16) int32 {
    if true {
        var v16 int64 = 210
        _ = v16
    }
    var v17 uint8 = 224
    switch 128 {
    case 128:
        v18 := GenS2{F0: 18446744073709551615}
        _ = v18
    default:
        v19 := []int{-30}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p1
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = (^(v1))
    var v3 int64 = 9223372036854775807
    if true {
        v20 := GenS2{F0: 0}
        var v21 uint8 = 42
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: -1, F1: 2147483647}
    v23 := []int{}
    _ = len(v23)
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 float64 = -2.25
        switch "foo bar" {
        case "foo bar":
            var v26 float64 = 0.0
            _ = v26
        case "":
            if len(v23) > 0 {
                v23[0] = int(v3)
            }
        default:
            switch fi24 {
            case 9223372036854775807:
                switch "a" {
                case "\n\t":
                    if false {
                        v2 = (v2 / ((2) | 1))
                        fmt.Println(v2, v3, v1)
                        v1 += genFn0(4294967295, 579, 0)
                    }
                }
            case -1:
                v27 := GenS3{F0: -705}
                _ = v27
            case 0:
                switch "" {
                case "\n\t":
                    v1 += (v2 + v1)
                case "foo bar":
                    v28 := GenS3{F0: -2147483648}
                    _ = v28
                }
            }
        }
        _ = fi24
        _ = v25
    }
    var v29 int = -941
    v29 = v29
    v30 := GenS3{F0: genFn1(2147483647, 526, 479)}
    _ = len(v23)
    var v31 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
