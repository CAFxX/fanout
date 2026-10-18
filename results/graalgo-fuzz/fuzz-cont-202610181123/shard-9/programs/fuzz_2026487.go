package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 bool; F1 int16; F2 int32 }
func genFn0(p0 uint32, p1 uint) int64 {
    if (-1 >= 42) {
        v4 := GenS1{F0: "foo bar", F1: 0}
        v5 := GenS2{F0: 0}
        var v6 float64 = 1.5
        _ = v4
        _ = v5
        _ = v6
    }
    fmt.Println(p0, p1)
    for fi7 := 0; fi7 < 1; fi7++ {
        switch "a" {
        case "foo bar":
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{}
                switch fi8 {
                case -9223372036854775808:
                    var v10 float32 = -2.25
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
        case "hello":
            if (-128 != 100) {
                v11 := []int{}
                var v12 uint16 = 2
                _ = v11
                _ = v12
            }
        default:
            p0 = uint32(p1)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    return 71
}
func genFn1(p0 string, p1 uint, p2 int64) int32 {
    p1 -= 4294967295
    var v13 uint8 = 128
    v13 = (v13 - v13)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 32767
    v3 := []uint32{2147483647}
    switch "foo bar" {
    case "a":
        if false {
            v2 -= 128
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                v2 += 782
                v15 := GenS3{F0: 21, F1: -9223372036854775808}
                if false {
                    var v16 uint64 = 4294967295
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            var v17 int8 = 42
            v2 -= v2
            _ = v17
        }
    case "":
        if true {
            fmt.Println(v2)
        } else {
            v2 -= (v2 | 0)
            switch "" {
            case "foo bar":
                for fi18 := 0; fi18 < 11; fi18++ {
                    v1[4294967295] = ((("hello" + ((("" + "") + "hello") + (("a" + ("foo bar" + "hello")) + "foo bar"))) + "foo bar") + ("a" + "\n\t"))
                    var v19 int64 = -628
                    if ((410 <= 1) && true) {
                        fmt.Println(v19)
                        switch "foo bar" {
                        case "foo bar":
                            switch "\n\t" {
                            case "a":
                                var v20 uint32 = 0
                                _ = v20
                            default:
                                for fi21 := 0; fi21 < 9; fi21++ {
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    switch fi21 {
                                    case -9223372036854775808:
                                        v2 += -360
                                    }
                                    _ = fi21
                                }
                            }
                        default:
                            var v22 int16 = -83
                            _ = v22
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            case "\n\t":
                var v23 int = (v2 >> ((65535) & 63))
                _ = v23
            case "":
                v2 = v2
            default:
                var v24 uint64 = 1
                _ = v24
            }
            v25 := GenS4{F0: false}
            _ = v25
        }
    case "foo bar":
        v2 += v2
    default:
        var v26 uint = 836
        _ = v26
    }
    v27 := GenS2{F0: 2147483647}
    v2 = (-(v2))
    v28 := GenS5{F0: -2147483648, F1: 42}
    v1[195] = "\n\t"
    v2 = v2
    v2 = (-(v2))
    v29 := GenS6{F0: (103 != 2), F1: -855, F2: -1}
    _ = len(v1)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
