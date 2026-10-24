package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int16; F1 string; F2 uint64 }
func genFn0(p0 uint64, p1 int8, p2 int8) uint8 {
    switch "\n\t" {
    case "\n\t":
        if false {
            v6 := GenS1{F0: 2, F1: (p0 * p0)}
            fmt.Println(p0, p1)
            switch 169 {
            case 980:
                p0 -= p0
            default:
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 int32 = 1
                    switch fi7 {
                    case 0:
                        var v9 float64 = 1.5
                        _ = v9
                    case 849:
                        var v10 int = (fi7 + (-(fi7)))
                        _ = v10
                    case 2147483647:
                        for fi11 := 0; fi11 < 9; fi11++ {
                            switch fi7 {
                            case -1:
                                p0 = 0
                            case 2147483647:
                                fmt.Println(fi11)
                            case 128:
                                p1 = p1
                            default:
                                var v12 uint8 = 183
                                _ = v12
                            }
                            var v13 int8 = 100
                            switch fi11 {
                            case 42:
                                v14 := []int{-874, 1, fi7, fi7}
                                _ = v14
                            }
                            _ = fi11
                            _ = v13
                        }
                    }
                    _ = fi7
                    _ = v8
                }
            }
            _ = v6
        } else {
            v15 := []int{42, 9223372036854775807, 565, 643}
            _ = v15
        }
    }
    v16 := GenS2{F0: 1, F1: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 148
}
func genFn1() uint8 {
    var v17 uint64 = 4294967295
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "hello" {
        case "\n\t":
            if false {
                v19 := GenS1{F0: 255, F1: 42}
                fmt.Println(fi18, v17)
                v20 := GenS3{F0: 100, F1: -9223372036854775808}
                _ = v19
                _ = v20
            }
        default:
            v21 := []int{(fi18 ^ fi18)}
            _ = v21
        }
        fmt.Println(v17)
        v17 = (^(v17))
        _ = fi18
    }
    _ = v17
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := []string{((("a" + ("foo bar" + ("a" + "hello"))) + "foo bar") + ""), ("foo bar" + "\n\t")}
    var v3 uint8 = 82
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v3 -= 128
    case -197:
        v22 := v4[65535]
        _ = v22
    case -155:
        v23 := GenS4{F0: ((^(v1)) == (v1 | (v1 ^ (v1 - v1)))), F1: 764}
        _ = v23
    }
    var v24 uint64 = 944
    v25 := GenS2{F0: 42, F1: 0}
    v26 := GenS5{F0: (v1 + v1), F1: "a", F2: 215}
    v24 = (v24 >> ((42) & 63))
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS5{F0: int16(v24), F1: ((((("foo bar" + "hello") + (("\n\t" + ("a" + (("foo bar" + "foo bar") + "foo bar"))) + "a")) + "foo bar") + "hello") + "\n\t"), F2: uint64(v3)}
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: -328, F1: 894}
    switch 1 {
    case 32768:
        if true {
            v1 = 693
            fmt.Println(v24, v24, v3)
        } else {
            var v30 int = 32768
            var v31 uint8 = genFn1()
            v30 += -573
            _ = v30
            _ = v31
        }
    default:
        v3 = (v3 >> ((18446744073709551615) & 63))
    }
    v3 -= v3
    var v32 uint32 = 1
    fmt.Println(v24)
    v24 += (v24 - v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
