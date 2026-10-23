package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0(p0 int16) bool {
    v7 := GenS1{F0: 958, F1: true, F2: 32767}
    if true {
        switch "\n\t" {
        case "\n\t":
            p0 -= (p0 + (p0 % ((1) | 1)))
        case "hello":
            var v8 int64 = -9223372036854775808
            _ = v8
        case "foo bar":
            fmt.Println(p0)
        }
        p0 = p0
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(p0)
            var v10 uint = 65535
            switch fi9 {
            case -206:
                v11 := GenS1{F0: 65535, F1: false, F2: 848}
                _ = v11
            case 9223372036854775807:
                var v12 uint64 = 18446744073709551615
                _ = v12
            default:
                fmt.Println(v10)
            }
            _ = fi9
            _ = v10
        }
    } else {
        p0 = p0
    }
    p0 += 478
    fmt.Println(p0)
    _ = p0
    _ = v7
    return (2 == 440)
}
func genFn1() int32 {
    var v13 string = "\n\t"
    v13 = v13
    v13 = v13
    var v14 int32 = -2147483648
    _ = v13
    _ = v14
    return (v14 / ((v14) | 1))
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]bool)
    var v3 uint32 = 42
    var v4 float64 = 0.0
    var v5 int64 = 0
    var v6 float64 = (float64(0.1) / v4)
    var v15 int32 = -195
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := []int{0, (-(fi16))}
        v15 -= (v15 + -985)
        _ = fi16
        _ = v17
    }
    v18 := []int{int(v5)}
    fmt.Println(v3, v4)
    _ = len(v18)
    v3 = (v3 ^ v3)
    v1 -= 1
    v6 = (1e10 - v6)
    if (true && (uint64(v3) != 4294967295)) {
        fmt.Println(v4, v3)
        switch -2147483648 {
        case 782:
            v19 := GenS2{F0: true, F1: v15}
            _ = v19
        case -149:
            v20 := GenS3{F0: 2147483647, F1: 1, F2: v3}
            _ = v20
        default:
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int16 = -382
                switch "hello" {
                case "foo bar":
                    if false {
                        fmt.Println(v3, v15, v5)
                        fmt.Println(v5, v22)
                    } else {
                        var v23 uint = uint(v3)
                        v22 = v22
                        _ = v23
                    }
                case "\n\t":
                    delete(v2, 0)
                case "":
                    v24 := GenS1{F0: 508, F1: (v5 >= v5), F2: 434}
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
        }
    }
    fmt.Println(v4)
    var v25 uint = 1
    v25 = (^(v25))
    switch 702 {
    case 32767:
        v3 -= (^(v3))
    case 349:
        var v26 int = 128
        _ = v26
    default:
        if (0 <= 0) {
            if true {
                v27 := GenS3{F0: v3, F1: int(v15), F2: 42}
                var v28 uint64 = 516
                _ = v27
                _ = v28
            } else {
                if true {
                    switch -605 {
                    case -2147483648:
                        v4 = -2.25
                    case 128:
                        v1 += -1
                    case 42:
                        v1 = v1
                    default:
                        switch "a" {
                        case "a":
                            _ = len(v18)
                        case "":
                            for fi29 := 0; fi29 < 16; fi29++ {
                                if len(v18) > 0 {
                                    v18[0] = (fi29 * 127)
                                }
                                _ = fi29
                            }
                        }
                    }
                    var v30 uint64 = 0
                    var v31 uint32 = v3
                    _ = v30
                    _ = v31
                }
            }
        } else {
            v25 = v25
            v4 += 1e10
            v32 := GenS1{F0: 42, F1: true, F2: 297}
            _ = v32
        }
    }
    delete(v2, v3)
    v6 += 3.14159
    switch "\n\t" {
    case "":
        v15 -= v15
    case "foo bar":
        var v33 uint32 = (v3 - 0)
        _ = v33
    default:
        v2[2] = (!true)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
