package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint, p1 int64, p2 string) int32 {
    v4 := []int{9223372036854775807, 4294967295}
    switch "hello" {
    case "a":
        switch "\n\t" {
        case "hello":
            v4 = append(v4, 2147483647)
        default:
            switch (p2 + p2) {
            case "":
                switch 516 {
                case 127:
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(p2)
                    default:
                        v4 = append(v4, -1)
                    }
                case 0:
                    var v5 string = (p2 + p2)
                    _ = v5
                case -484:
                    _ = len(v4)
                }
            case "a":
                if false {
                    for fi6 := 0; fi6 < 18; fi6++ {
                        var v7 uint = (p0 ^ 65535)
                        _ = fi6
                        _ = v7
                    }
                    for fi8 := 0; fi8 < 18; fi8++ {
                        if true {
                            for fi9 := 0; fi9 < 11; fi9++ {
                                p1 -= p1
                                fmt.Println(fi8, fi8, p2)
                                _ = fi9
                            }
                            v10 := GenS1{F0: 127}
                            _ = v10
                        }
                        switch fi8 {
                        case -457:
                            var v11 int16 = -81
                            _ = v11
                        }
                        _ = fi8
                    }
                } else {
                    fmt.Println(p0, p0)
                    var v12 int32 = 42
                    switch -2147483648 {
                    case 102:
                        v13 := GenS1{F0: -16}
                        _ = v13
                    case 128:
                        if true {
                            var v14 uint = p0
                            p1 -= 42
                            if false {
                                var v15 uint64 = 770
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = 551
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 2147483647
                                    }
                                    switch -561 {
                                    case 848:
                                        v4 = append(v4, 127)
                                    case -437:
                                        v16 := GenS2{F0: 957, F1: 146, F2: v15}
                                        _ = v16
                                    case -655:
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = 1
                                            }
                                        }
                                    }
                                } else {
                                    p0 = v14
                                    p1 += -9223372036854775808
                                }
                                v12 = (v12 * 42)
                                _ = v15
                            }
                            _ = v14
                        }
                    }
                    _ = v12
                }
            case "hello":
                v17 := GenS3{F0: -1, F1: -2147483648}
                _ = v17
            }
        }
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(p0, p0, fi18)
            fmt.Println(fi18, fi18, fi18)
            _ = fi18
        }
    }
    if true {
        v19 := GenS4{F0: 299, F1: -865, F2: 42}
        fmt.Println(p0, p0, p1)
        _ = v19
    } else {
        v4 = append(v4, 128)
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    v1 := []int{127, 599, -549, 127}
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    v2 -= 1
    v2 += v2
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 float64 = 0.0
            _ = fi21
            _ = v22
        }
        switch "hello" {
        case "hello":
            var v23 int8 = v2
            _ = v23
        case "\n\t":
            var v24 bool = ((!false) && true)
            _ = v24
        case "":
            if len(v1) > 0 {
                v1[0] = (fi20 % ((0) | 1))
            }
        }
        v25 := GenS5{F0: 896, F1: 0}
        _ = fi20
        _ = v25
    }
    v2 += (v2 | (v2 << ((18446744073709551615) & 63)))
    var v26 uint = 42
    var v27 uint64 = 1
    var v28 string = "\n\t"
    fmt.Println(v27, v28, v28)
    v28 = (v28 + v28)
    for fi29 := 0; fi29 < 7; fi29++ {
        if (!false) {
            v30 := v3[0]
            _ = v30
        }
        v31 := GenS2{F0: (fi29 & (fi29 ^ fi29)), F1: 42, F2: 0}
        var v32 int = (fi29 % (((fi29 / ((fi29) | 1))) | 1))
        _ = fi29
        _ = v31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
