package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint }
type GenS4 struct { F0 int8; F1 string; F2 uint8 }
func genFn0() uint64 {
    var v4 uint64 = 0
    var v5 float32 = 1e10
    switch 9223372036854775807 {
    case 7:
        switch 4294967295 {
        case 127:
            fmt.Println(v4, v4)
        default:
            fmt.Println(v5, v5)
        }
    case 950:
        v5 += -2.25
    case 9223372036854775807:
        var v6 int = 128
        _ = v6
    default:
        v5 -= ((3.14159 + float32(((0.0 + float64(v5)) / 0.1))) + 0.1)
    }
    var v7 int8 = 100
    _ = v4
    _ = v5
    _ = v7
    return (^(v4))
}
func genFn1(p0 uint8) uint64 {
    if false {
        var v8 bool = (100 > 0)
        var v9 int8 = 47
        _ = v8
        _ = v9
    }
    if (-1 != 637) {
        v10 := GenS1{F0: 1, F1: 42, F2: 27}
        p0 += (p0 | (p0 % ((2) | 1)))
        v11 := []int{-385, 2147483647}
        _ = v10
        _ = v11
    } else {
        var v12 float64 = -2.25
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                switch fi14 {
                case 32767:
                    var v15 float64 = -2.25
                    _ = v15
                case 127:
                    var v16 int16 = -32768
                    _ = v16
                case 19:
                    v17 := GenS2{F0: -576}
                    _ = v17
                default:
                    if (1.5 == 0.0) {
                        var v18 float64 = 3.14159
                        switch (fi14 * fi14) {
                        case 866:
                            var v19 uint16 = 0
                            _ = v19
                        case -514:
                            v20 := []int{fi14, (fi14 & fi13)}
                            _ = v20
                        case 128:
                            v21 := []int{2147483647, (fi13 >> ((559) & 63)), fi13, fi13}
                            _ = v21
                        default:
                            if false {
                                var v22 int8 = 1
                                _ = v22
                            }
                        }
                        _ = v18
                    } else {
                        fmt.Println(fi14, fi13)
                        p0 = p0
                    }
                }
                _ = fi14
            }
            fmt.Println(p0)
            _ = fi13
        }
        var v23 uint = 42
        _ = v12
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(fi24, p0)
        if false {
            v25 := []int{32767, fi24, fi24}
            _ = v25
        }
        fmt.Println(p0, fi24)
        _ = fi24
    }
    p0 -= p0
    v26 := GenS3{F0: 2147483647, F1: 255, F2: 2}
    _ = p0
    _ = v26
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 uint16 = 873
    v27 := GenS2{F0: -1}
    var v28 int32 = 2147483647
    switch "\n\t" {
    case "":
        var v29 float32 = 0.0
        _ = v29
    case "\n\t":
        v30 := GenS1{F0: (v28 % (((v28 * v28)) | 1)), F1: v28, F2: 170}
        _ = v30
    }
    if (451 >= 222) {
        if ((((1e10 * (1.5 / float32(1e10))) + 0.1) / (((1e10 + 1e10) + (((0.0 / 1.5) - 1.5) * 1e10)) * 0.1)) >= 1.5) {
            var v31 uint64 = 30
            _ = v31
        }
        fmt.Println(v28, v28)
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 int64 = -1
            _ = fi32
            _ = v33
        }
    }
    var v34 int = 128
    if (false && (true || (-1 == 42))) {
        v35 := GenS1{F0: (v28 + 136), F1: (v28 ^ (v28 * (^(v28)))), F2: 1}
        v36 := GenS4{F0: -1, F1: "\n\t", F2: 8}
        v28 = (^(v28))
        _ = v35
        _ = v36
    } else {
        var v37 int = -1
        _ = v37
    }
    v28 += (v28 % ((243) | 1))
    v38 := []int{v34}
    var v39 int16 = -50
    if (!false) {
        var v40 uint32 = 65535
        _ = v40
    }
    if false {
        var v41 uint8 = 2
        if (1 > 255) {
            v1[("hello" + ("\n\t" + "\n\t"))] = (v34 | 2147483647)
            if ((^(v39)) >= v39) {
                fmt.Println(v3)
            }
            fmt.Println(v3, v28, v41)
        } else {
            delete(v1, "\n\t")
        }
        _ = v41
    }
    var v42 float64 = 0.0
    v3 = (v3 + v3)
    v38 = append(v38, (v34 >> ((uint(v3)) & 63)))
    var v43 uint8 = 127
    var v44 uint = 437
    v44 -= 723
    fmt.Println(v28, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v42
    _ = v43
    _ = v44
}
