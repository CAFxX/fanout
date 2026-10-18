package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint; F1 int16; F2 uint8 }
type GenS6 struct { F0 int8; F1 bool; F2 int }
type GenS7 struct { F0 int32; F1 int; F2 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS12 struct { F0 int64; F1 int32; F2 uint }
type GenS13 struct { F0 string }
func genFn0(p0 string) uint8 {
    v6 := []int{}
    var v7 int32 = -1
    if ((0.0 + float32(1.5)) >= 0.0) {
        fmt.Println(p0)
    } else {
        var v8 int16 = 32767
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 string = "hello"
                fmt.Println(p0, v8)
                _ = fi10
                _ = v11
            }
            if true {
                switch fi9 {
                case 1:
                    if false {
                        v12 := GenS1{F0: -899, F1: -1}
                        if true {
                            var v13 string = ("a" + (((p0 + "") + (("" + ((((p0 + (((("hello" + p0) + p0) + (p0 + (p0 + ""))) + p0)) + "foo bar") + ((p0 + "") + p0)) + p0)) + "a")) + p0))
                            var v14 int16 = v8
                            v7 -= v7
                            _ = v13
                            _ = v14
                        } else {
                            v8 -= v8
                            switch ((("\n\t" + p0) + ((p0 + (p0 + p0)) + (("" + p0) + p0))) + ("" + "foo bar")) {
                            case "a":
                                fmt.Println(p0)
                            case "\n\t":
                                var v15 int16 = -368
                                _ = v15
                            }
                        }
                        _ = v12
                    } else {
                        if (42 == 1) {
                            if len(v6) > 0 {
                                v6[0] = fi9
                            }
                            v7 += int32(fi9)
                        }
                        p0 = p0
                        for fi16 := 0; fi16 < 17; fi16++ {
                            fmt.Println(fi16, v7)
                            if true {
                                if (true || true) {
                                    fmt.Println(fi16)
                                    if true {
                                        p0 = p0
                                        var v17 uint64 = 1
                                        _ = v17
                                    }
                                    fmt.Println(fi9, fi16, v7)
                                } else {
                                    v7 += (v7 << ((18446744073709551615) & 63))
                                }
                                var v18 float32 = 0.1
                                fmt.Println(v7, v8)
                                if ((p0 <= (p0 + "")) && (false && false)) {
                                    switch (fi9 | fi16) {
                                    case -582:
                                        v19 := GenS1{F0: 9223372036854775807, F1: 100}
                                        _ = v19
                                    case 654:
                                        v7 += 367
                                    case 32767:
                                        _ = len(v6)
                                    default:
                                        fmt.Println(fi9, v18)
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = fi16
                                    }
                                } else {
                                    fmt.Println(fi16, fi9, fi16)
                                    v20 := GenS2{F0: 766, F1: 322, F2: -128}
                                    v8 = (^(v8))
                                    _ = v20
                                }
                                if true {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v21 uint64 = 2
                                    if false {
                                        switch fi9 {
                                        case 904:
                                            v22 := GenS2{F0: 0, F1: 817, F2: 0}
                                            _ = v22
                                        case 128:
                                            v18 += v18
                                        default:
                                            v7 = v7
                                        }
                                        v8 = -356
                                        var v23 int16 = v8
                                        _ = v23
                                    }
                                    var v24 string = p0
                                    v21 += 42
                                    if (2 < (v21 % ((v21) | 1))) {
                                        if false {
                                            if true {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            var v25 uint16 = 21
                                            _ = v25
                                        }
                                        if ((^(v21)) > v21) {
                                            var v26 int8 = 105
                                            if false {
                                                v27 := GenS3{F0: v7, F1: 1, F2: 18446744073709551615}
                                                v28 := GenS4{F0: (v21 + v21), F1: 255}
                                                var v29 uint = 42
                                                v30 := GenS5{F0: v29, F1: (v8 ^ (^(v8))), F2: 42}
                                                var v31 float32 = 1.5
                                                if true {
                                                    if true {
                                                        v18 -= v31
                                                        var v32 int8 = v26
                                                        var v33 bool = (false || (255 == 869))
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v26 -= v26
                                                        v34 := GenS2{F0: 2, F1: uint32(v21), F2: (-(v26))}
                                                        _ = v34
                                                    }
                                                    if (true && (!(!(94 < 255)))) {
                                                        var v35 int8 = -13
                                                        if true {
                                                            if false {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v35 -= (v26 * v35)
                                                                v21 = (v21 % (((v21 >> (((v29 - (v29 >> ((v29) & 63)))) & 63))) | 1))
                                                            }
                                                        }
                                                        _ = v35
                                                    } else {
                                                        v36 := GenS6{F0: int8(v29), F1: (v7 <= (v7 ^ 1)), F2: 4294967295}
                                                        v37 := GenS7{F0: v7, F1: (fi16 | (fi9 + 0)), F2: (v8 << ((v29) & 63))}
                                                        v38 := GenS2{F0: 4294967295, F1: 2147483647, F2: (-(v26))}
                                                        if len(v6) > 0 {
                                                            v6[0] = 4294967295
                                                        }
                                                        v21 += 255
                                                        if len(v6) > 0 {
                                                            v6[0] = (fi9 + (fi9 & (fi9 | fi16)))
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                            _ = v26
                                        } else {
                                            if false {
                                                if (0 == 337) {
                                                    switch p0 {
                                                    case "\n\t":
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v7 -= (v7 + (^(v7)))
                                                        case "foo bar":
                                                            switch 9223372036854775807 {
                                                            case 4294967295:
                                                                switch (v24 + ((v24 + p0) + p0)) {
                                                                case "hello":
                                                                    var v39 uint8 = 255
                                                                    _ = v39
                                                                }
                                                            case 252:
                                                                var v40 int8 = -107
                                                                _ = v40
                                                            }
                                                        case "":
                                                            if len(v6) > 0 {
                                                                v6[0] = 9223372036854775807
                                                            }
                                                        default:
                                                            if (v8 <= (v8 - v8)) {
                                                                v41 := GenS8{F0: v24}
                                                                _ = v41
                                                            }
                                                        }
                                                    }
                                                    var v42 int16 = -32768
                                                    _ = v42
                                                }
                                                v43 := GenS8{F0: p0}
                                                _ = v43
                                            } else {
                                                switch ((("foo bar" + "a") + v24) + "\n\t") {
                                                case "\n\t":
                                                    v44 := GenS9{F0: (^(v8)), F1: -9223372036854775808}
                                                    _ = v44
                                                default:
                                                    fmt.Println(v18, fi16, fi9)
                                                }
                                                if true {
                                                    v7 += (v7 / ((v7) | 1))
                                                    v21 -= v21
                                                    var v45 uint16 = 503
                                                    _ = v45
                                                }
                                                var v46 float64 = (0.0 / ((-2.25 - 1.5) * (((float64(v18) + 3.14159) - (float64((v18 - ((v18 - float32(float64(((float32(1.5) * (float32((-2.25 * float64(v18))) / v18)) + v18)))) / ((((3.14159 / 1.0) / v18) + (float32((0.1 - (3.14159 - -2.25))) / -2.25)) + float32(0.0))))) + 0.1)) - 0.0)))
                                                _ = v46
                                            }
                                            _ = len(v6)
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        fmt.Println(v18)
                                        var v47 int = (fi16 + -2147483648)
                                        v48 := GenS8{F0: v24}
                                        _ = v47
                                        _ = v48
                                    }
                                    if (true || (70 <= -128)) {
                                        v6 = append(v6, 2147483647)
                                        _ = len(v6)
                                        var v49 int32 = (v7 | -270)
                                        _ = v49
                                    } else {
                                        v21 = 2
                                        v50 := GenS10{F0: 209}
                                        _ = v50
                                    }
                                    _ = v21
                                    _ = v24
                                }
                                if (true || false) {
                                    v51 := GenS3{F0: v7, F1: 0, F2: 876}
                                    var v52 int8 = 96
                                    switch (fi9 + 127) {
                                    case 42:
                                        v6 = append(v6, (fi9 * (fi9 << ((65535) & 63))))
                                    case 649:
                                        var v53 int64 = 42
                                        _ = v53
                                    case -545:
                                        v8 -= int16(v7)
                                    default:
                                        if (1 > 42) {
                                            v52 -= (v52 << ((1) & 63))
                                        }
                                    }
                                    switch (p0 + ((p0 + "\n\t") + (p0 + ((("foo bar" + p0) + "foo bar") + p0)))) {
                                    case "hello":
                                        v6 = append(v6, fi9)
                                    case "":
                                        v54 := GenS11{F0: 255, F1: true, F2: 0}
                                        _ = v54
                                    default:
                                        var v55 bool = false
                                        _ = v55
                                    }
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v56 int32 = v7
                                    switch (fi16 << ((255) & 63)) {
                                    case -9223372036854775808:
                                        fmt.Println(v18)
                                    case 9223372036854775807:
                                        fmt.Println(v8, v8)
                                    default:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                    v57 := GenS4{F0: 2, F1: 449}
                                    _ = v56
                                    _ = v57
                                }
                                _ = v18
                            }
                            var v58 int8 = 100
                            fmt.Println(v7, fi16)
                            v7 -= -1
                            _ = fi16
                            _ = v58
                        }
                    }
                }
            }
            fmt.Println(v7, v7, v8)
            _ = fi9
        }
        _ = v8
    }
    for fi59 := 0; fi59 < 2; fi59++ {
        v7 = v7
        var v60 uint = 4294967295
        if false {
            var v61 int8 = 78
            v62 := GenS1{F0: 199, F1: (v61 % ((42) | 1))}
            _ = v61
            _ = v62
        }
        _ = fi59
        _ = v60
    }
    v7 -= v7
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string, p1 int) uint {
    var v63 int = p1
    switch p1 {
    case 102:
        p0 = "a"
    case 9223372036854775807:
        var v64 string = (("" + p0) + ("" + p0))
        _ = v64
    default:
        if false {
            p0 = ((p0 + ("a" + p0)) + "hello")
            v65 := []int{152, p1}
            p1 -= p1
            _ = v65
        } else {
            for fi66 := 0; fi66 < 10; fi66++ {
                if (286 > 178) {
                    v67 := []int{v63, (fi66 << ((65535) & 63))}
                    if len(v67) > 0 {
                        v67[0] = 32768
                    }
                    _ = v67
                }
                _ = fi66
            }
            fmt.Println(p1, v63, p1)
        }
    }
    p1 = p1
    if (31 >= 42) {
        switch 1 {
        case 429:
            for fi68 := 0; fi68 < 17; fi68++ {
                v69 := GenS12{F0: 42, F1: 2147483647, F2: 793}
                _ = fi68
                _ = v69
            }
        default:
            if (!(((false && true) && false) && false)) {
                v70 := GenS8{F0: "hello"}
                _ = v70
            }
        }
        for fi71 := 0; fi71 < 1; fi71++ {
            var v72 int32 = -2147483648
            v72 += (v72 + (-(v72)))
            _ = fi71
            _ = v72
        }
    }
    switch 32768 {
    case 32767:
        p0 = p0
    case 0:
        if (907 < 2) {
            v73 := []int{v63, p1}
            _ = v73
        } else {
            v74 := GenS2{F0: 0, F1: 4294967295, F2: 1}
            fmt.Println(p0, p0)
            _ = v74
        }
    default:
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v63
    return 4294967295
}

func main() {
    var v1 int32 = -837
    var v2 int32 = (-(v1))
    v3 := make(map[uint32]uint32)
    var v4 float32 = 0.1
    v5 := []string{"\n\t", "foo bar"}
    fmt.Println(v1, v1)
    delete(v3, 1)
    for fi75 := 0; fi75 < 16; fi75++ {
        for fi76 := 0; fi76 < 13; fi76++ {
            if (2 >= 877) {
                switch "hello" {
                case "hello":
                    fmt.Println(v4, v4)
                }
                if true {
                    if ((1e10 * 1.5) != (1.5 + 0.0)) {
                        _ = len(v3)
                        var v77 uint8 = genFn0("")
                        switch "" {
                        case "foo bar":
                            v1 = (v1 + (^(v2)))
                        case "a":
                            if false {
                                var v78 string = ""
                                _ = v78
                            } else {
                                if (!true) {
                                    if true {
                                        v2 = (v1 + (v2 << ((18446744073709551615) & 63)))
                                        v4 -= v4
                                        var v79 int16 = 568
                                        _ = v79
                                    }
                                    v80 := v3[944]
                                    fmt.Println(v4, v4)
                                    _ = v80
                                }
                                fmt.Println(v4, v1)
                            }
                        }
                        _ = v77
                    } else {
                        var v81 float64 = -2.25
                        v4 = 1e10
                        v2 -= v1
                        v82 := v3[1]
                        fmt.Println(fi75)
                        _ = v81
                        _ = v82
                    }
                } else {
                    v1 = (-(v1))
                    var v83 int16 = 32767
                    _ = v83
                }
                v2 = 336
            }
            _ = fi76
        }
        _ = fi75
    }
    var v84 string = "a"
    v1 = 2147483647
    v2 = v2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v85 float64 = -2.25
    v86 := GenS13{F0: v84}
    v1 += (v2 - v1)
    for fi87 := 0; fi87 < 20; fi87++ {
        if true {
            v85 = v85
            v88 := GenS12{F0: 0, F1: (-(v2)), F2: 1}
            _ = v88
        }
        _ = fi87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
    _ = v86
}
