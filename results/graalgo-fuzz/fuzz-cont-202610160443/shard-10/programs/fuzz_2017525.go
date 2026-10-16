package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 string; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 string; F2 uint32 }
type GenS6 struct { F0 uint32; F1 string; F2 int8 }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint64) uint64 {
    p0 = p0
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 += (p0 % (((p0 ^ p0)) | 1))
        _ = fi4
    }
    v5 := GenS1{F0: false}
    var v6 int8 = -128
    v7 := GenS2{F0: 2}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint8) uint16 {
    if true {
        if false {
            p0 = 35
            if (4294967295 >= -2147483648) {
                for fi8 := 0; fi8 < 18; fi8++ {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        if false {
                            var v10 string = ("a" + ("a" + "foo bar"))
                            var v11 float32 = (-2.25 - 0.1)
                            _ = v10
                            _ = v11
                        }
                        switch fi9 {
                        case -1:
                            var v12 uint8 = p0
                            _ = v12
                        case 32767:
                            v13 := GenS3{F0: 1, F1: ("hello" + ("a" + "")), F2: (fi8 & (fi9 >> ((575) & 63)))}
                            _ = v13
                        case 128:
                            fmt.Println(fi9)
                        default:
                            v14 := []int{(fi8 + fi9), fi8, 32768, 127}
                            _ = v14
                        }
                        _ = fi9
                    }
                    var v15 float32 = float32(((-2.25 / float64(3.14159)) + 3.14159))
                    p0 -= (^(p0))
                    _ = fi8
                    _ = v15
                }
            }
        } else {
            fmt.Println(p0)
            v16 := GenS4{F0: 455}
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := []int{1, 2147483647}
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    }
    fmt.Println(p0)
    if true {
        var v19 bool = (!(220 > 641))
        if true {
            p0 = (p0 & p0)
        }
        v20 := []int{-47}
        _ = v19
        _ = v20
    } else {
        p0 -= (p0 + (p0 & p0))
        for fi21 := 0; fi21 < 11; fi21++ {
            p0 += p0
            if ((((((-2.25 + float32(float64(0.0))) + ((0.0 / 1e10) * (((0.1 + 1e10) - 1.5) + (3.14159 - (0.0 * 0.0))))) / 1e10) * 3.14159) - float32(3.14159)) > 1.5) {
                p0 = (p0 * 127)
            } else {
                v22 := []int{fi21}
                switch (fi21 / ((fi21) | 1)) {
                case -2147483648:
                    if (fi21 >= fi21) {
                        switch (fi21 & fi21) {
                        case 128:
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        case -196:
                            if true {
                                if ((!false) && (-2147483648 > 1)) {
                                    p0 += (p0 >> ((65535) & 63))
                                    v23 := GenS2{F0: 42}
                                    switch (fi21 / ((fi21) | 1)) {
                                    case 0:
                                        _ = len(v22)
                                    case 128:
                                        p0 -= (p0 | p0)
                                    case 810:
                                        v22 = append(v22, (fi21 + fi21))
                                    }
                                    _ = v23
                                } else {
                                    v24 := GenS4{F0: 2}
                                    v22 = append(v22, fi21)
                                    _ = v24
                                }
                            }
                        case -2147483648:
                            var v25 int32 = -2147483648
                            _ = v25
                        default:
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        }
                    } else {
                        var v26 uint64 = 2
                        _ = v26
                    }
                case 441:
                    p0 -= 255
                }
                _ = v22
            }
            fmt.Println(fi21, fi21)
            _ = fi21
        }
        if true {
            fmt.Println(p0)
            switch ("" + "foo bar") {
            case "foo bar":
                switch 4294967295 {
                case 32768:
                    p0 += p0
                case -9223372036854775808:
                    var v27 float32 = (1.5 + -2.25)
                    _ = v27
                case 202:
                    var v28 int32 = -576
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        for fi30 := 0; fi30 < 1; fi30++ {
                            if true {
                                var v31 uint = 136
                                _ = v31
                            }
                            var v32 int64 = 9223372036854775807
                            _ = fi30
                            _ = v32
                        }
                        p0 += (p0 << ((2) & 63))
                        _ = fi29
                    }
                }
            default:
                if (("hello" <= ("foo bar" + "foo bar")) && false) {
                    p0 = (^(p0))
                }
            }
        }
    }
    var v33 int16 = 32767
    _ = p0
    _ = v33
    return 2
}

func main() {
    v1 := []uint32{745, 2147483647, 0}
    var v2 int16 = 42
    var v3 int8 = 100
    for fi34 := 0; fi34 < 8; fi34++ {
        if true {
            v35 := GenS3{F0: 42, F1: "\n\t", F2: 382}
            v36 := GenS5{F0: 42, F1: "hello", F2: 826}
            _ = v35
            _ = v36
        }
        _ = fi34
    }
    var v37 uint32 = 316
    for fi38 := 0; fi38 < 17; fi38++ {
        _ = len(v1)
        _ = fi38
    }
    var v39 uint32 = v37
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v40 := GenS2{F0: 2147483648}
    if true {
        v41 := GenS5{F0: 403, F1: "a", F2: (v39 % ((v39) | 1))}
        var v42 uint8 = uint8(v39)
        var v43 int = 562
        _ = v41
        _ = v42
        _ = v43
    }
    fmt.Println(v2, v3)
    if false {
        if ((^(v3)) != v3) {
            fmt.Println(v39, v3)
            var v44 int16 = -392
            v44 += (v44 ^ (v44 / (((-(v44))) | 1)))
            _ = v44
        } else {
            switch "\n\t" {
            case "a":
                v2 = (v2 & (v2 * (v2 ^ (v2 << ((65535) & 63)))))
            case "":
                v2 = (v2 + v2)
            }
            switch 127 {
            case -1:
                if (0 > -320) {
                    if true {
                        switch -1 {
                        case 127:
                            var v45 int64 = 0
                            _ = v45
                        default:
                            var v46 uint32 = (v37 % (((v37 / (((^(v37))) | 1))) | 1))
                            _ = v46
                        }
                    }
                    var v47 bool = ((0.1 + 1e10) == 0.0)
                    _ = v47
                }
            case 128:
                var v48 int32 = 42
                _ = v48
            case 2147483647:
                v39 = v39
            }
        }
        for fi49 := 0; fi49 < 19; fi49++ {
            v2 -= int16(fi49)
            v3 = v3
            _ = fi49
        }
        v39 = (v39 % (((v37 >> ((318) & 63))) | 1))
    }
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    fmt.Println(v3, v39)
    switch "a" {
    case "hello":
        for fi50 := 0; fi50 < 13; fi50++ {
            if len(v1) > 0 {
                v1[0] = v39
            }
            _ = fi50
        }
    default:
        var v51 int = 32768
        _ = v51
    }
    fmt.Println(v39, v39)
    var v52 int32 = -987
    var v53 int32 = v52
    if (float64(1.5) > 0.1) {
        for fi54 := 0; fi54 < 15; fi54++ {
            v55 := GenS6{F0: 42, F1: "hello", F2: v3}
            _ = fi54
            _ = v55
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v56 uint16 = 255
        _ = v56
    } else {
        for fi57 := 0; fi57 < 3; fi57++ {
            switch "hello" {
            case "a":
                if false {
                    if (true || true) {
                        v37 = 1
                    } else {
                        v58 := GenS1{F0: (-797 > (fi57 & fi57))}
                        _ = v58
                    }
                } else {
                    v37 -= v39
                }
            case "\n\t":
                var v59 uint = 1
                _ = v59
            case "foo bar":
                var v60 uint32 = v37
                _ = v60
            default:
                if false {
                    fmt.Println(v3, v39, v39)
                    v37 = 1
                    for fi61 := 0; fi61 < 20; fi61++ {
                        if ((!false) && (false && false)) {
                            fmt.Println(v37, v52)
                        } else {
                            v53 += (v52 & v53)
                            v39 = v37
                            var v62 uint = 42
                            v3 = (v3 & (v3 % (((-(v3))) | 1)))
                            _ = v62
                        }
                        fmt.Println(v37, fi57)
                        var v63 uint64 = 748
                        _ = fi61
                        _ = v63
                    }
                } else {
                    var v64 int32 = -1
                    fmt.Println(v3, v2)
                    if true {
                        var v65 uint64 = 0
                        for fi66 := 0; fi66 < 1; fi66++ {
                            v3 -= 0
                            if false {
                                var v67 string = "foo bar"
                                var v68 int = (fi66 ^ int(v3))
                                if (-128 >= v3) {
                                    if len(v1) > 0 {
                                        v1[0] = v37
                                    }
                                } else {
                                    v69 := GenS7{F0: 100, F1: 369}
                                    _ = v69
                                }
                                _ = v67
                                _ = v68
                            }
                            _ = fi66
                        }
                        if true {
                            switch fi57 {
                            case 42:
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    switch "\n\t" {
                                    case "a":
                                        var v71 int16 = (v2 << ((65535) & 63))
                                        _ = v71
                                    case "\n\t":
                                        v64 = (v64 % ((v64) | 1))
                                        v52 -= (^(v53))
                                    case "foo bar":
                                        if false {
                                            v3 = 1
                                            if true {
                                                if (int(v65) < -713) {
                                                    v65 -= (v65 - v65)
                                                    if ((468 < -20) && false) {
                                                        v37 -= v39
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        var v72 uint8 = 161
                                                        var v73 uint8 = 128
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    v39 = (v37 | v39)
                                                } else {
                                                    var v74 string = ""
                                                    var v75 int32 = 612
                                                    _ = v74
                                                    _ = v75
                                                }
                                                v52 = (v53 << ((0) & 63))
                                                v52 = (v53 * (v52 >> ((0) & 63)))
                                            } else {
                                                v76 := GenS2{F0: (v37 | 65535)}
                                                fmt.Println(v65, v65, v53)
                                                v77 := GenS4{F0: 2}
                                                _ = v76
                                                _ = v77
                                            }
                                            var v78 uint8 = 2
                                            _ = v78
                                        }
                                    default:
                                        var v79 uint8 = 127
                                        _ = v79
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v80 uint64 = (v65 % ((genFn0(genFn0(897))) | 1))
                                        _ = v80
                                    case "a":
                                        v81 := GenS8{F0: 0}
                                        _ = v81
                                    }
                                    v37 += 482
                                    _ = fi70
                                }
                            default:
                                switch "" {
                                case "foo bar":
                                    v37 = 862
                                case "hello":
                                    for fi82 := 0; fi82 < 8; fi82++ {
                                        var v83 float64 = 1e10
                                        v64 = 42
                                        if len(v1) > 0 {
                                            v1[0] = 99
                                        }
                                        v84 := GenS1{F0: true}
                                        var v85 uint32 = 42
                                        _ = fi82
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                default:
                                    for fi86 := 0; fi86 < 9; fi86++ {
                                        v53 = (v53 >> ((322) & 63))
                                        if true {
                                            v2 += v2
                                        }
                                        _ = fi86
                                    }
                                }
                            }
                        }
                        _ = v65
                    }
                    _ = v64
                }
            }
            _ = fi57
        }
    }
    var v87 int = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v40
    _ = v52
    _ = v53
    _ = v87
}
