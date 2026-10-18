package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 string; F2 uint }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 bool }
func genFn0() uint {
    switch ("foo bar" + ("" + "")) {
    case "":
        var v4 float64 = 1.5
        _ = v4
    case "hello":
        var v5 uint16 = 694
        _ = v5
    default:
        if false {
            var v6 uint = 292
            for fi7 := 0; fi7 < 7; fi7++ {
                if (!(673 >= 1)) {
                    v8 := GenS1{F0: (^(fi7)), F1: 1, F2: 231}
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v6 -= v6
                        var v10 uint = 1
                        if (!(v6 != v10)) {
                            var v11 string = (("a" + "") + "")
                            _ = v11
                        } else {
                            var v12 uint32 = 0
                            v13 := []int{1, (-(fi7))}
                            _ = v12
                            _ = v13
                        }
                        _ = fi9
                        _ = v10
                    }
                    v14 := []int{fi7, 0, (fi7 + fi7), fi7}
                    _ = v14
                    _ = v8
                }
                if false {
                    fmt.Println(fi7)
                    var v15 float32 = 1.5
                    _ = v15
                }
                switch (fi7 % (((fi7 * -9223372036854775808)) | 1)) {
                case 83:
                    v16 := GenS1{F0: (fi7 * (^(fi7))), F1: 0, F2: 128}
                    _ = v16
                case -844:
                    var v17 uint64 = 18446744073709551615
                    _ = v17
                case -175:
                    if true {
                        v18 := GenS2{F0: 42}
                        v6 = 0
                        _ = v18
                    } else {
                        var v19 float64 = 0.1
                        for fi20 := 0; fi20 < 7; fi20++ {
                            var v21 float64 = v19
                            var v22 int = (fi7 - fi20)
                            switch (^(fi20)) {
                            case 1:
                                switch (fi20 * -1) {
                                case 803:
                                    fmt.Println(v19, fi7)
                                }
                            case 2147483647:
                                if (!(!(true && false))) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v23 string = "\n\t"
                                        _ = v23
                                    case "":
                                        if (-128 == 1) {
                                            if true {
                                                if false {
                                                    fmt.Println(v19)
                                                }
                                            } else {
                                                v21 = (v19 / 0.1)
                                                v24 := GenS3{F0: -44, F1: 0, F2: ""}
                                                _ = v24
                                            }
                                        } else {
                                            switch (v22 * 101) {
                                            case 2147483647:
                                                switch "\n\t" {
                                                case "hello":
                                                    if true {
                                                        var v25 float32 = 3.14159
                                                        fmt.Println(fi7, v19)
                                                        _ = v25
                                                    } else {
                                                        v26 := []int{-1, -80, 2147483647, (v22 | (fi20 - (fi20 << ((v6) & 63))))}
                                                        _ = v26
                                                    }
                                                case "":
                                                    fmt.Println(fi7, fi20, v22)
                                                case "a":
                                                    if true {
                                                        v27 := []int{}
                                                        v21 -= v19
                                                        _ = v27
                                                    }
                                                }
                                            }
                                        }
                                    case "a":
                                        switch "" {
                                        case "":
                                            v28 := []int{9223372036854775807}
                                            _ = v28
                                        case "hello":
                                            if (true || (-9223372036854775808 < 1)) {
                                                v29 := GenS4{F0: false, F1: -1}
                                                _ = v29
                                            }
                                        }
                                    }
                                    v19 = -2.25
                                }
                            default:
                                var v30 string = (("a" + "hello") + ("foo bar" + (("a" + "hello") + "\n\t")))
                                _ = v30
                            }
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        v31 := GenS5{F0: 4294967295}
                        _ = v19
                        _ = v31
                    }
                default:
                    v32 := GenS6{F0: -62}
                    _ = v32
                }
                _ = fi7
            }
            _ = v6
        }
    }
    var v33 bool = (false || true)
    v33 = v33
    _ = v33
    return 255
}
func genFn1(p0 bool) int16 {
    switch "a" {
    case "foo bar":
        p0 = (!p0)
    }
    if p0 {
        switch "\n\t" {
        case "foo bar":
            var v34 uint16 = 1
            _ = v34
        default:
            fmt.Println(p0)
        }
        var v35 uint16 = 1
        _ = v35
    } else {
        p0 = (!(-743 != -499))
        if p0 {
            var v36 float32 = ((1e10 * 1.5) * 0.1)
            var v37 float32 = v36
            _ = v36
            _ = v37
        }
        var v38 int8 = -95
        _ = v38
    }
    p0 = (!(4294967295 != 207))
    var v39 uint = 65535
    var v40 uint16 = 65535
    _ = p0
    _ = v39
    _ = v40
    return -32768
}

func main() {
    v1 := []uint32{757, 0, 42}
    var v2 int16 = 32767
    var v3 int16 = v2
    fmt.Println(v3, v3)
    var v41 int32 = -1
    var v42 uint64 = 843
    v43 := GenS6{F0: 42}
    fmt.Println(v42, v42, v2)
    if false {
        v1 = append(v1, 42)
        if ("foo bar" != "hello") {
            if (42 < (v41 << ((1) & 63))) {
                if false {
                    fmt.Println(v3)
                    switch 32768 {
                    case 0:
                        if true {
                            fmt.Println(v42)
                            v2 = v3
                            v44 := GenS5{F0: 579}
                            _ = v44
                        } else {
                            var v45 int16 = 528
                            _ = v45
                        }
                    case 128:
                        _ = len(v1)
                    default:
                        if (0 > 1) {
                            fmt.Println(v3, v42, v41)
                        }
                    }
                }
                var v46 float64 = -2.25
                _ = v46
            }
        } else {
            if (-35 != -106) {
                var v47 uint64 = (v42 - 4294967295)
                v48 := GenS3{F0: -8, F1: v47, F2: (("a" + "\n\t") + "foo bar")}
                _ = v47
                _ = v48
            }
            switch "foo bar" {
            case "foo bar":
                v2 = (v2 ^ (-(v2)))
            case "\n\t":
                for fi49 := 0; fi49 < 4; fi49++ {
                    var v50 bool = ((42 != 0) && false)
                    _ = fi49
                    _ = v50
                }
            default:
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            }
            var v51 float32 = ((((float32(1e10) * ((float32(1.5) * (3.14159 + 1.5)) + 3.14159)) + 1e10) + -2.25) / (0.0 - (((0.1 * 1.5) / 3.14159) - (-2.25 * 1e10))))
            _ = v51
        }
    } else {
        var v52 uint = 255
        var v53 uint32 = 1
        if ((v52 - (v52 >> (((v52 / ((483) | 1))) & 63))) != genFn0()) {
            if true {
                var v54 int8 = 42
                _ = v54
            }
        } else {
            v55 := GenS7{F0: -128, F1: "a", F2: v52}
            if len(v1) > 0 {
                v1[0] = (v53 % ((1) | 1))
            }
            for fi56 := 0; fi56 < 4; fi56++ {
                if false {
                    v3 = 0
                }
                for fi57 := 0; fi57 < 11; fi57++ {
                    var v58 uint64 = (v42 + v42)
                    if true {
                        var v59 string = ("hello" + ("foo bar" + ((("foo bar" + "hello") + "a") + "foo bar")))
                        switch (fi57 | fi56) {
                        case 127:
                            if ((v58 & (v58 ^ 4294967295)) >= (v58 & v58)) {
                                if (float64(float32(3.14159)) < (-2.25 / 1e10)) {
                                    v60 := GenS4{F0: (!true), F1: 1}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch int(v2) {
                                    case 1:
                                        var v61 uint = genFn0()
                                        _ = v61
                                    case -475:
                                        switch "hello" {
                                        case "\n\t":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            fmt.Println(v3, v53)
                                        }
                                    case 9223372036854775807:
                                        var v62 uint8 = 2
                                        fmt.Println(v53, fi56)
                                        v63 := GenS3{F0: -128, F1: v58, F2: v59}
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v60
                                }
                                v41 = (v41 - v41)
                                fmt.Println(v59, v41, v52)
                            }
                        case -106:
                            var v64 int = (-(fi56))
                            _ = v64
                        }
                        fmt.Println(v59)
                        _ = v59
                    } else {
                        if (12 < 0) {
                            fmt.Println(v3)
                            v58 = 0
                            v58 += 716
                        }
                    }
                    v52 = v52
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
            _ = v55
        }
        _ = v52
        _ = v53
    }
    for fi65 := 0; fi65 < 13; fi65++ {
        fmt.Println(v41, v42, v42)
        v3 = (v3 >> ((genFn0()) & 63))
        _ = fi65
    }
    var v66 bool = (0 < 18446744073709551615)
    for fi67 := 0; fi67 < 20; fi67++ {
        if v66 {
            v68 := GenS5{F0: 602}
            for fi69 := 0; fi69 < 18; fi69++ {
                v41 += (v41 & (v41 ^ (v41 * v41)))
                var v70 string = "a"
                var v71 int8 = 127
                _ = fi69
                _ = v70
                _ = v71
            }
            for fi72 := 0; fi72 < 12; fi72++ {
                if v66 {
                    v66 = (!((!false) && (false || true)))
                }
                switch "" {
                case "\n\t":
                    if false {
                        v1 = append(v1, 0)
                        v2 -= (v2 | 167)
                        v73 := GenS8{F0: v42, F1: (v3 % ((int16(fi67)) | 1))}
                        _ = v73
                    } else {
                        v74 := GenS8{F0: (v42 * v42), F1: (v3 << ((796) & 63))}
                        if true {
                            v75 := GenS8{F0: v42, F1: genFn1((!(42 < 42)))}
                            _ = v75
                        } else {
                            fmt.Println(fi67, v2)
                            if len(v1) > 0 {
                                v1[0] = 637
                            }
                            v41 = v41
                        }
                        _ = v74
                    }
                case "":
                    _ = len(v1)
                default:
                    v2 += (^(v3))
                }
                if (192 == 255) {
                    v66 = false
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi72
            }
            _ = v68
        }
        for fi76 := 0; fi76 < 16; fi76++ {
            var v77 uint8 = 2
            var v78 int32 = v41
            fmt.Println(v77, v66)
            _ = fi76
            _ = v77
            _ = v78
        }
        _ = fi67
    }
    v66 = v66
    v79 := GenS9{F0: 934, F1: 4294967295}
    switch ("a" + "foo bar") {
    case "\n\t":
        var v80 int = 9223372036854775807
        _ = v80
    default:
        v3 += 495
    }
    for fi81 := 0; fi81 < 7; fi81++ {
        for fi82 := 0; fi82 < 14; fi82++ {
            if v66 {
                switch "a" {
                case "foo bar":
                    if (v66 && false) {
                        v41 = v41
                        v66 = ("foo bar" > "\n\t")
                        if v66 {
                            fmt.Println(fi81, v2, v66)
                            v83 := GenS10{F0: fi81}
                            var v84 uint8 = 42
                            if ((-1 > v2) || (65535 <= 65535)) {
                                v66 = (-101 >= -128)
                                fmt.Println(v42)
                            } else {
                                v84 = 0
                                fmt.Println(v42)
                            }
                            if ((^(fi82)) > fi82) {
                                fmt.Println(v3, fi82, v66)
                                if ((42 == 4294967295) || false) {
                                    if ((false && (!v66)) || ((42 >= 255) || false)) {
                                        var v85 int = 754
                                        _ = v85
                                    }
                                    var v86 float32 = float32(float64(1.5))
                                    if len(v1) > 0 {
                                        v1[0] = 827
                                    }
                                    var v87 uint32 = 835
                                    _ = v86
                                    _ = v87
                                } else {
                                    v88 := GenS11{F0: genFn0()}
                                    switch "foo bar" {
                                    case "":
                                        if v66 {
                                            fmt.Println(v41, v41, v66)
                                        } else {
                                            var v89 int16 = (v2 + v2)
                                            var v90 uint = 0
                                            v91 := GenS8{F0: 18446744073709551615, F1: genFn1(v66)}
                                            v92 := GenS12{F0: ((v90 >> (((^(v90))) & 63)) > (v90 | uint(v42)))}
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        }
                                    case "\n\t":
                                        if ((!true) || false) {
                                            var v93 uint64 = v42
                                            v93 -= (^(v42))
                                            _ = len(v1)
                                            _ = v93
                                        }
                                    case "foo bar":
                                        if ((0 >= 1) || (42 > -46)) {
                                            var v94 int64 = 1
                                            var v95 bool = (2147483648 < uint32(v84))
                                            v2 = v3
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v94
                                            _ = v95
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v84 = 42
                                            _ = len(v1)
                                        }
                                    default:
                                        if v66 {
                                            v42 = uint64(v84)
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v88
                                }
                            }
                            fmt.Println(fi82, v3, v41)
                            _ = v83
                            _ = v84
                        }
                    } else {
                        var v96 float32 = 0.0
                        _ = v96
                    }
                default:
                    switch "hello" {
                    case "foo bar":
                        v2 += v2
                    }
                }
                fmt.Println(v66)
                if len(v1) > 0 {
                    v1[0] = 301
                }
            } else {
                fmt.Println(v2, v3, v3)
            }
            v97 := GenS6{F0: 100}
            _ = fi82
            _ = v97
        }
        if (false && true) {
            fmt.Println(v3, v42)
        }
        var v98 float64 = (1e10 * 1e10)
        _ = fi81
        _ = v98
    }
    for fi99 := 0; fi99 < 13; fi99++ {
        var v100 string = ((((("hello" + "a") + "foo bar") + "foo bar") + "hello") + "a")
        switch (^(fi99)) {
        case -9223372036854775808:
            var v101 int8 = -13
            _ = v101
        }
        var v102 int16 = 0
        _ = fi99
        _ = v100
        _ = v102
    }
    v3 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v66
    _ = v79
}
