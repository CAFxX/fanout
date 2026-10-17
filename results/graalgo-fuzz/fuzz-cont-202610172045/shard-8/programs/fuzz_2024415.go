package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 int8) int64 {
    v4 := []int{-480}
    v5 := GenS1{F0: p0, F1: -9223372036854775808, F2: false}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = (1e10 * 0.1)
        v8 := GenS1{F0: p0, F1: 675, F2: false}
        for fi9 := 0; fi9 < 4; fi9++ {
            p0 -= p0
            fmt.Println(fi6)
            var v10 bool = ("foo bar" >= "foo bar")
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v11 := GenS1{F0: (p0 * 32767), F1: -9223372036854775808, F2: false}
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    _ = v5
    return -1
}
func genFn1() uint32 {
    var v12 float64 = 1.5
    v13 := GenS2{F0: 0, F1: 18446744073709551615}
    switch 865 {
    case 32768:
        if (!(152 < 709)) {
            v12 += 3.14159
            var v14 float32 = 1.5
            _ = v14
        }
    case 32767:
        fmt.Println(v12)
    case 42:
        v12 -= v12
    default:
        fmt.Println(v12)
    }
    _ = v12
    _ = v13
    return 137
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = (1e10 * 3.14159)
    v3 := make(map[uint32]int)
    v3[0] = -2147483648
    var v15 int16 = 42
    v15 = 0
    var v16 uint64 = 4294967295
    v17 := GenS2{F0: 42, F1: (v16 - (v16 & (v16 << ((2) & 63))))}
    v18 := []int{32767}
    _ = len(v3)
    switch -9223372036854775808 {
    case 42:
        if false {
            var v19 int64 = 42
            _ = v19
        }
    case -708:
        switch ("hello" + (("" + "a") + "")) {
        case "\n\t":
            for fi20 := 0; fi20 < 19; fi20++ {
                delete(v3, 1)
                for fi21 := 0; fi21 < 4; fi21++ {
                    v15 -= int16(v16)
                    delete(v3, 441)
                    v22 := GenS1{F0: v15, F1: 42, F2: true}
                    _ = fi21
                    _ = v22
                }
                if true {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        if false {
                            switch "foo bar" {
                            case "":
                                v16 -= (v16 % ((v16) | 1))
                                fmt.Println(fi23, v2)
                            case "\n\t":
                                if (94 < 255) {
                                    v2 -= 0.1
                                    v2 += float64(v1)
                                    var v24 int64 = genFn0(v15, 42)
                                    v1 = v1
                                    var v25 uint8 = 230
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint = 52
                                    v27 := GenS2{F0: 170, F1: v16}
                                    _ = v26
                                    _ = v27
                                }
                            }
                            v2 += float64((-2.25 / float32(v2)))
                        }
                        v28 := GenS2{F0: -9223372036854775808, F1: (v16 * v16)}
                        _ = fi23
                        _ = v28
                    }
                    switch fi20 {
                    case 26:
                        v29 := GenS2{F0: int64(v16), F1: 18446744073709551615}
                        _ = v29
                    }
                } else {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v15 += (^(v15))
                        v15 -= v15
                        _ = fi30
                    }
                }
                _ = fi20
            }
        default:
            var v31 int16 = (-(v15))
            _ = v31
        }
    }
    v16 -= v16
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := GenS3{F0: fi32}
        _ = fi32
        _ = v33
    }
    var v34 uint64 = (v16 / ((v16) | 1))
    delete(v3, 429)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
}
