package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int32 }
func genFn0(p0 int64) uint {
    var v7 int8 = 42
    fmt.Println(p0)
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := []int{32768, fi8}
        var v10 int = fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = (v7 / (((v7 + v7)) | 1))
    _ = p0
    _ = v7
    return 176
}

func main() {
    v1 := []int{-896, 42, 117, 4294967295}
    var v2 int16 = -853
    v3 := make(map[int]uint32)
    v4 := []int{42, -818, -571, -304}
    var v5 uint32 = 670
    var v6 uint32 = v5
    v11 := GenS1{F0: -128}
    switch ("a" + "") {
    case "a":
        v2 = (v2 / (((v2 - (v2 << ((65535) & 63)))) | 1))
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        v6 += v5
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(fi12)
        fmt.Println(v6)
        _ = fi12
    }
    v13 := v3[9223372036854775807]
    fmt.Println(v6, v5, v13)
    if false {
        var v14 int = -258
        v2 = 768
        _ = v14
    } else {
        var v15 uint16 = 42
        _ = v15
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (false || (false || false)) {
        v2 += (-(v2))
        v6 += 850
        v16 := GenS2{F0: 4294967295, F1: 2147483647}
        _ = v16
    } else {
        var v17 int = 448
        if (true || (255 < 0)) {
            for fi18 := 0; fi18 < 11; fi18++ {
                v13 = (^(v13))
                var v19 int8 = 127
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                switch (fi20 - int(v5)) {
                case 128:
                    if (!false) {
                        v3[748] = 180
                        _ = len(v3)
                        var v21 int = (v17 ^ fi20)
                        _ = v21
                    } else {
                        var v22 string = "\n\t"
                        _ = len(v3)
                        v13 = (^(v5))
                        _ = v22
                    }
                case -244:
                    var v23 uint8 = 0
                    _ = v23
                case 34:
                    v5 = 0
                default:
                    switch "" {
                    case "hello":
                        fmt.Println(v5)
                    case "foo bar":
                        if (!(1 > 1)) {
                            if true {
                                v17 -= (fi20 ^ 32767)
                            } else {
                                v24 := GenS1{F0: 100}
                                _ = v24
                            }
                        }
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        for fi25 := 0; fi25 < 1; fi25++ {
                            switch "\n\t" {
                            case "":
                                var v26 uint8 = 42
                                _ = v26
                            case "a":
                                if true {
                                    if false {
                                        var v27 uint8 = 2
                                        _ = v27
                                    } else {
                                        v5 = v13
                                    }
                                    var v28 uint16 = 144
                                    _ = v28
                                }
                            }
                            _ = fi25
                        }
                    }
                }
                v13 = 1
                _ = fi20
            }
            _ = len(v4)
        } else {
            v29 := GenS2{F0: (-(v17)), F1: -2147483648}
            v17 -= v17
            _ = len(v3)
            _ = v29
        }
        _ = v17
    }
    var v30 int32 = -1
    fmt.Println(v6, v5)
    v6 += v5
    _ = len(v3)
    fmt.Println(v13, v5)
    var v31 uint64 = 18446744073709551615
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS2{F0: -1, F1: (v30 << ((65535) & 63))}
        if (v2 <= 42) {
            fmt.Println(v31, v5, v6)
            fmt.Println(v5, fi32)
            v6 += uint32(v31)
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
