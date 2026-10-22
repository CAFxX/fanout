package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int64; F1 int64; F2 uint32 }
func genFn0(p0 int16, p1 uint32, p2 bool) int8 {
    if (!(0 == -43)) {
        v5 := []int{-46, 944}
        _ = v5
    } else {
        v6 := []int{}
        _ = v6
    }
    var v7 int16 = (p0 / (((p0 >> ((18446744073709551615) & 63))) | 1))
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            p2 = p2
            _ = fi9
        }
        v7 -= v7
        v10 := GenS1{F0: p0}
        _ = fi8
        _ = v10
    }
    fmt.Println(p1, v7, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -16
}

func main() {
    v1 := []int{}
    var v2 int8 = 106
    var v3 uint32 = 0
    var v4 float64 = (1.5 - 1e10)
    switch 128 {
    case 42:
        var v11 uint64 = 1
        _ = v11
    case -2147483648:
        v3 -= 713
    default:
        if false {
            switch 2147483647 {
            case -9223372036854775808:
                var v12 int64 = 42
                _ = v12
            case 471:
                fmt.Println(v2, v2)
            default:
                var v13 int32 = 0
                _ = v13
            }
        } else {
            fmt.Println(v3, v4, v3)
        }
    }
    v3 = (^(v3))
    v2 = v2
    v2 = v2
    v14 := GenS2{F0: v3, F1: 353, F2: 127}
    v15 := GenS1{F0: 42}
    for fi16 := 0; fi16 < 10; fi16++ {
        switch (-(fi16)) {
        case 0:
            if len(v1) > 0 {
                v1[0] = fi16
            }
        case 32768:
            if true {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int16 = 86
                    _ = fi17
                    _ = v18
                }
                _ = len(v1)
                v19 := GenS1{F0: 1000}
                _ = v19
            } else {
                var v20 int = 279
                _ = v20
            }
        case -483:
            for fi21 := 0; fi21 < 8; fi21++ {
                v3 = (v3 - (v3 - (v3 << ((0) & 63))))
                var v22 int16 = 32767
                if false {
                    var v23 uint64 = 0
                    _ = v23
                } else {
                    switch ("hello" + "\n\t") {
                    case "":
                        if (fi16 == 156) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "hello" {
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = 32767
                                }
                            }
                            switch fi16 {
                            case -2147483648:
                                v2 += genFn0(v22, (v3 & v3), false)
                            case 128:
                                if true {
                                    v2 = int8(v3)
                                }
                            case 705:
                                fmt.Println(v22, v2, v22)
                            }
                        } else {
                            if (!true) {
                                if true {
                                    v24 := GenS2{F0: (v3 >> ((4294967295) & 63)), F1: 65535, F2: 165}
                                    var v25 uint16 = uint16(v3)
                                    _ = v24
                                    _ = v25
                                }
                                switch (((("foo bar" + "hello") + "hello") + "\n\t") + "\n\t") {
                                case "":
                                    var v26 int16 = v22
                                    _ = v26
                                case "foo bar":
                                    switch (fi21 % (((^(fi21))) | 1)) {
                                    case -865:
                                        v2 += v2
                                        switch "\n\t" {
                                        case "foo bar":
                                            v2 += v2
                                        }
                                        switch (^(fi21)) {
                                        case 128:
                                            v27 := GenS2{F0: v3, F1: 735, F2: 127}
                                            _ = v27
                                        default:
                                            if true {
                                                v28 := GenS3{F0: 42, F1: (536 != -595)}
                                                v3 = 15
                                                if (255 == 2) {
                                                    v3 += (v3 ^ (v3 - 65535))
                                                }
                                                _ = v28
                                            } else {
                                                fmt.Println(v22, fi16)
                                                v2 += v2
                                                fmt.Println(v22)
                                                if true {
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi16 * fi16)
                                                    }
                                                }
                                                v1 = append(v1, (fi16 * (fi16 ^ 1)))
                                            }
                                        }
                                    case -65:
                                        v29 := GenS4{F0: 1}
                                        _ = v29
                                    }
                                case "hello":
                                    _ = len(v1)
                                default:
                                    var v30 uint8 = 128
                                    _ = v30
                                }
                                var v31 int64 = 1
                                _ = v31
                            }
                            switch fi21 {
                            case 1:
                                switch fi16 {
                                case -1:
                                    if len(v1) > 0 {
                                        v1[0] = fi21
                                    }
                                default:
                                    v32 := GenS5{F0: 412, F1: (!(-80 == 1))}
                                    _ = v32
                                }
                            case 2147483647:
                                if false {
                                    var v33 uint8 = 139
                                    _ = v33
                                } else {
                                    var v34 int64 = 42
                                    v35 := GenS5{F0: 32767, F1: true}
                                    var v36 int = -900
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = len(v1)
                        }
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = 9223372036854775807
                        }
                    case "\n\t":
                        var v37 float32 = float32(v4)
                        _ = v37
                    }
                    v38 := GenS4{F0: 42}
                    _ = v38
                }
                _ = fi21
                _ = v22
            }
        default:
            var v39 int16 = 42
            _ = v39
        }
        v3 -= (v3 - (v3 << ((42) & 63)))
        switch 568 {
        case 0:
            fmt.Println(fi16, fi16)
        case -580:
            v2 = (v2 + (v2 - 1))
        case 9223372036854775807:
            for fi40 := 0; fi40 < 10; fi40++ {
                switch (("a" + "foo bar") + (("a" + "foo bar") + "hello")) {
                case "foo bar":
                    v2 -= -128
                    v4 = 1e10
                    if (!false) {
                        var v41 int64 = 42
                        if (201 <= 42) {
                            var v42 uint32 = v3
                            var v43 int8 = (-(v2))
                            v44 := GenS6{F0: v41, F1: v41, F2: (v3 - 846)}
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        var v45 uint16 = 1
                        _ = v41
                        _ = v45
                    } else {
                        fmt.Println(v4, v3)
                    }
                    v46 := GenS2{F0: (v3 - (v3 >> ((4294967295) & 63))), F1: 526, F2: 190}
                    _ = v46
                }
                fmt.Println(v4)
                var v47 uint = 4294967295
                v1 = append(v1, 9223372036854775807)
                var v48 uint32 = 2147483648
                if (false || ((!true) || (!(false && (((1e10 - v4) / v4) != v4))))) {
                    v1 = append(v1, (fi40 & fi40))
                } else {
                    v4 = (v4 / (0.0 * (((v4 - v4) / v4) / 1.5)))
                    switch 128 {
                    case 32768:
                        var v49 uint32 = 42
                        _ = v49
                    case 4294967295:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -823:
                        var v50 uint16 = 781
                        _ = v50
                    }
                    fmt.Println(v47, v3)
                }
                _ = fi40
                _ = v47
                _ = v48
            }
        default:
            v51 := GenS2{F0: (v3 | v3), F1: 644, F2: 147}
            _ = v51
        }
        _ = fi16
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    var v52 float32 = 0.0
    v53 := GenS6{F0: 1, F1: 9223372036854775807, F2: v3}
    v4 -= 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
}
