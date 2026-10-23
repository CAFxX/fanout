package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 bool; F1 int; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int64; F2 bool }
type GenS6 struct { F0 uint32; F1 int32; F2 int32 }
type GenS7 struct { F0 int64; F1 int8; F2 int32 }
func genFn0(p0 string) uint {
    v7 := GenS1{F0: -4, F1: 42, F2: -1}
    p0 = "hello"
    var v8 int32 = -1
    v9 := []int{0, 128, 0, -285}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 873
}

func main() {
    v1 := []uint32{4294967295, 243, 4294967295}
    var v2 float32 = 1.5
    var v3 uint8 = 42
    v4 := []uint32{1, 0, 65535}
    var v5 uint = 255
    var v6 uint32 = 2147483647
    for fi10 := 0; fi10 < 10; fi10++ {
        if ("foo bar" != ("a" + (("hello" + "hello") + ""))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                if (127 > 255) {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        switch fi11 {
                        case 823:
                            v12 := GenS1{F0: 100, F1: 18446744073709551615, F2: 4}
                            var v13 bool = (-1 >= 9223372036854775807)
                            v2 -= (v2 * -2.25)
                            _ = v12
                            _ = v13
                        case 15:
                            v2 += (v2 - v2)
                        default:
                            v14 := GenS1{F0: -1, F1: 2, F2: -128}
                            _ = v14
                        }
                        fmt.Println(fi10)
                        v3 = (^(v3))
                        _ = fi11
                    }
                    v15 := GenS2{F0: 727}
                    var v16 uint8 = v3
                    _ = v15
                    _ = v16
                } else {
                    var v17 float64 = (1.5 + (1e10 * 1.5))
                    _ = v17
                }
                switch "foo bar" {
                case "foo bar":
                    for fi18 := 0; fi18 < 17; fi18++ {
                        switch -505 {
                        case -2147483648:
                            v19 := GenS3{F0: "\n\t", F1: 100}
                            if len(v1) > 0 {
                                v1[0] = uint32(v5)
                            }
                            _ = v19
                        }
                        if len(v4) > 0 {
                            v4[0] = 416
                        }
                        var v20 int = fi10
                        _ = fi18
                        _ = v20
                    }
                case "hello":
                    v21 := GenS4{F0: true, F1: (^(fi10)), F2: 42}
                    _ = v21
                case "":
                    v3 = (v3 + (v3 + (v3 & 42)))
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    v23 := GenS5{F0: 0, F1: -978, F2: (!true)}
                    if true {
                        v24 := GenS1{F0: 26, F1: 1, F2: 42}
                        _ = v24
                    }
                    _ = fi22
                    _ = v23
                }
            } else {
                fmt.Println(fi10, v5)
                if (0 == -720) {
                    v3 -= (^(v3))
                    v25 := GenS5{F0: 42, F1: -1, F2: (!((-128 < 42) && true))}
                    if (53 > 0) {
                        v26 := GenS6{F0: 2, F1: -1, F2: -1}
                        _ = v26
                    } else {
                        var v27 int32 = 952
                        v28 := GenS6{F0: 1, F1: v27, F2: 42}
                        _ = v27
                        _ = v28
                    }
                    _ = v25
                }
            }
        }
        _ = fi10
    }
    fmt.Println(v6, v6)
    switch "\n\t" {
    case "foo bar":
        var v29 int = 128
        _ = v29
    default:
        switch "hello" {
        case "\n\t":
            fmt.Println(v2)
        default:
            v30 := GenS1{F0: 42, F1: 220, F2: -128}
            _ = v30
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 float32 = v2
        v3 = v3
        var v33 int32 = -520
        _ = fi31
        _ = v32
        _ = v33
    }
    v3 -= v3
    for fi34 := 0; fi34 < 12; fi34++ {
        v5 += v5
        var v35 float32 = v2
        fmt.Println(v5, v5)
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        v37 := GenS1{F0: int8(v5), F1: 2, F2: 1}
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 int16 = -1
            v40 := GenS7{F0: -973, F1: 100, F2: -2147483648}
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi36
        _ = v37
    }
    var v41 float64 = 0.0
    fmt.Println(v41)
    var v42 int8 = -128
    fmt.Println(v42, v41)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
