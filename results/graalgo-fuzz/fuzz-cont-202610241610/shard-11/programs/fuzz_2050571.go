package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16 }
func genFn0(p0 int32) string {
    if true {
        switch "hello" {
        case "a":
            v6 := []int{-102}
            _ = v6
        case "":
            if (!(!(42 <= 18446744073709551615))) {
                var v7 uint16 = 362
                _ = v7
            }
        }
    } else {
        v8 := GenS1{F0: p0, F1: 565, F2: 4294967295}
        v9 := GenS1{F0: p0, F1: 65535, F2: 594}
        p0 += p0
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: 1}
    _ = p0
    _ = v10
    return "foo bar"
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    v11 := GenS1{F0: 569, F1: 65535, F2: 65535}
    if (("\n\t" + ("foo bar" + ("hello" + "foo bar"))) == "foo bar") {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float32 = float32(-2.25)
            fmt.Println(fi12, p0)
            _ = fi12
            _ = v13
        }
        if (0 < -216) {
            var v14 bool = false
            _ = v14
        }
    } else {
        switch "\n\t" {
        case "":
            var v15 int = 600
            _ = v15
        case "foo bar":
            v16 := GenS2{F0: false, F1: 1}
            _ = v16
        }
        p0 = 228
    }
    v17 := []int{-869, 32768, 32768}
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(p0)
        var v19 uint32 = 2
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 float64 = -2.25
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = v11
    _ = v17
    return "foo bar"
}

func main() {
    var v1 string = ""
    v2 := []string{v1, (v1 + ((v1 + v1) + (("\n\t" + (v1 + v1)) + v1)))}
    v3 := make(map[uint32]int)
    var v4 uint64 = 1
    var v5 int = -1
    fmt.Println(v5)
    v5 = 77
    v1 = ("hello" + (("\n\t" + "") + "hello"))
    delete(v3, 0)
    if (!(!false)) {
        var v22 int16 = -906
        if false {
            v23 := GenS2{F0: true, F1: -9223372036854775808}
            v24 := GenS2{F0: ((-2.25 + (((1e10 / 0.1) * (1e10 * 1e10)) - 0.0)) <= float64(((-2.25 / float32((1.5 / 1.0))) + 0.0))), F1: 1}
            var v25 int32 = 42
            _ = v23
            _ = v24
            _ = v25
        } else {
            var v26 int32 = 42
            if (v22 != int16(v5)) {
                _ = len(v2)
                var v27 int8 = 127
                _ = v27
            } else {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                if true {
                    if true {
                        var v28 int16 = (v22 * (v22 - (v22 ^ 32767)))
                        v5 = (^(v5))
                        if true {
                            for fi29 := 0; fi29 < 8; fi29++ {
                                for fi30 := 0; fi30 < 14; fi30++ {
                                    if (v1 < "foo bar") {
                                        v4 += v4
                                        if (v4 >= v4) {
                                            v31 := GenS3{F0: v1}
                                            v32 := GenS4{F0: 332}
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    fmt.Println(fi30, v22)
                                    _ = fi30
                                }
                                v22 = (v22 % ((v28) | 1))
                                v33 := GenS1{F0: -2147483648, F1: 42, F2: 2}
                                _ = fi29
                                _ = v33
                            }
                            v28 -= -32768
                            var v34 uint16 = 32767
                            _ = v34
                        } else {
                            var v35 int16 = v28
                            if ((145 != 725) || (1 != 255)) {
                                v2 = append(v2, (v1 + v1))
                                var v36 uint = 73
                                _ = v36
                            }
                            _ = v35
                        }
                        _ = v28
                    } else {
                        var v37 int16 = (v22 & 32767)
                        var v38 float64 = 0.0
                        _ = v37
                        _ = v38
                    }
                    if true {
                        for fi39 := 0; fi39 < 9; fi39++ {
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                            _ = fi39
                        }
                        var v40 uint8 = 2
                        _ = v40
                    } else {
                        v5 = (v5 | 2147483647)
                    }
                    var v41 float32 = 3.14159
                    _ = v41
                } else {
                    if (float32((-2.25 / 1.0)) <= 0.0) {
                        v1 = v1
                        if len(v2) > 0 {
                            v2[0] = v1
                        }
                        v22 -= v22
                    } else {
                        v2 = append(v2, v1)
                    }
                }
            }
            v26 = v26
            _ = v26
        }
        for fi42 := 0; fi42 < 20; fi42++ {
            var v43 uint8 = 42
            _ = fi42
            _ = v43
        }
        _ = v22
    } else {
        if true {
            var v44 float32 = 0.1
            v44 = v44
            fmt.Println(v1, v44, v44)
            _ = v44
        } else {
            if true {
                if false {
                    v45 := GenS5{F0: (v1 + ("" + v1))}
                    v46 := GenS6{F0: 408, F1: 4294967295}
                    var v47 int64 = -905
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    if false {
                        v5 = (v5 ^ -2147483648)
                        switch (v5 + v5) {
                        case 0:
                            v4 = v4
                        case -9223372036854775808:
                            switch (v1 + v1) {
                            case "\n\t":
                                fmt.Println(v5, v4)
                            case "a":
                                if (194 == 145) {
                                    v5 = (v5 >> ((1) & 63))
                                }
                            case "":
                                _ = len(v3)
                            default:
                                v1 = "\n\t"
                            }
                        default:
                            var v48 int32 = 2147483647
                            _ = v48
                        }
                    }
                }
                switch v5 {
                case 2147483647:
                    if (v5 > (v5 >> ((438) & 63))) {
                        v49 := GenS6{F0: -647, F1: 142}
                        _ = v49
                    } else {
                        v50 := GenS7{F0: 1}
                        _ = v50
                    }
                }
            } else {
                var v51 uint32 = 98
                if (0.0 <= 1.5) {
                    v2 = append(v2, v1)
                    var v52 bool = false
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 9; fi53++ {
                        v54 := GenS8{F0: -9223372036854775808}
                        var v55 uint16 = 1
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    v56 := GenS2{F0: (uint16(v4) > 931), F1: -9223372036854775808}
                    var v57 int8 = 0
                    _ = v56
                    _ = v57
                }
                fmt.Println(v5, v1)
                _ = v51
            }
            for fi58 := 0; fi58 < 4; fi58++ {
                var v59 float64 = 0.1
                switch v1 {
                case "a":
                    var v60 string = v1
                    _ = v60
                default:
                    fmt.Println(fi58)
                }
                _ = fi58
                _ = v59
            }
        }
        fmt.Println(v1, v5)
        v4 += v4
    }
    v61 := GenS4{F0: 633}
    if ((v1 + "foo bar") < v1) {
        v4 -= v4
        fmt.Println(v1, v4, v5)
        delete(v3, 65535)
    } else {
        _ = len(v3)
        var v62 uint64 = v4
        var v63 uint8 = 0
        _ = v62
        _ = v63
    }
    v5 = (v5 ^ v5)
    if false {
        switch v5 {
        case 379:
            switch v5 {
            case 128:
                var v64 float64 = 0.1
                _ = v64
            default:
                v5 += v5
            }
        case -980:
            fmt.Println(v1)
        case -9223372036854775808:
            fmt.Println(v5, v5)
        default:
            v4 += v4
        }
        if false {
            v4 = v4
        } else {
            v4 -= v4
        }
        var v65 int16 = 357
        _ = v65
    } else {
        var v66 int8 = 117
        fmt.Println(v66)
        var v67 uint32 = 42
        _ = v66
        _ = v67
    }
    if true {
        _ = len(v2)
    } else {
        v5 -= v5
        v1 = v1
        var v68 int64 = 1
        _ = v68
    }
    v69 := GenS6{F0: 54, F1: 42}
    var v70 int = (v5 >> ((0) & 63))
    var v71 bool = (42 <= 463)
    switch "\n\t" {
    case "\n\t":
        var v72 int64 = 0
        _ = v72
    case "a":
        v73 := v3[294]
        _ = v73
    case "hello":
        if false {
            var v74 float64 = (1e10 - 1.5)
            _ = v74
        }
    default:
        var v75 uint32 = 2147483647
        _ = v75
    }
    if v71 {
        v2 = append(v2, "foo bar")
    } else {
        delete(v3, 65535)
        if (!("\n\t" < "a")) {
            v70 += 576
            for fi76 := 0; fi76 < 7; fi76++ {
                v77 := GenS6{F0: 1, F1: 0}
                _ = fi76
                _ = v77
            }
            v1 = "hello"
        }
    }
    v70 -= (v5 & (v70 - v70))
    v78 := GenS9{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v69
    _ = v70
    _ = v71
    _ = v78
}
