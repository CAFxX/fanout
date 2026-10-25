package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int32; F1 string; F2 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 uint; F1 uint16 }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 uint32; F1 uint32; F2 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 bool; F1 uint64 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 bool; F1 int8 }
type GenS17 struct { F0 uint }
func genFn0(p0 int64, p1 int64) int8 {
    v5 := GenS1{F0: 264, F1: 621, F2: "foo bar"}
    p1 += 42
    _ = p0
    _ = p1
    _ = v5
    return -92
}

func main() {
    v1 := []uint32{332, 65535, 2147483648, 0}
    var v2 float32 = 0.0
    var v3 float64 = float64(v2)
    v4 := []int{32768, 860, -714}
    _ = len(v1)
    v3 -= v3
    v6 := GenS1{F0: 2, F1: 4294967295, F2: "foo bar"}
    if len(v1) > 0 {
        v1[0] = 11
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS2{F0: 81}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                v2 = ((0.1 / 0.1) - v2)
                switch "foo bar" {
                case "":
                    if (("hello" + ((("a" + "a") + "\n\t") + "foo bar")) < "hello") {
                        var v10 int8 = 42
                        _ = v10
                    }
                case "hello":
                    switch 32767 {
                    case 127:
                        v3 = (v3 + float64(v2))
                    case 32768:
                        var v11 uint64 = 275
                        _ = v11
                    case 179:
                        switch "" {
                        case "hello":
                            switch (("a" + "") + "") {
                            case "hello":
                                var v12 float64 = 3.14159
                                _ = v12
                            case "":
                                fmt.Println(v2)
                            }
                        case "a":
                            if (32767 == 2) {
                                var v13 uint = 4294967295
                                v3 = 1.5
                                v3 = v3
                                _ = v13
                            } else {
                                if true {
                                    v3 -= (v3 + v3)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v14 := GenS2{F0: -128}
                                    var v15 uint16 = 255
                                    _ = v14
                                    _ = v15
                                } else {
                                    fmt.Println(v3, v3, v3)
                                    if false {
                                        v16 := GenS2{F0: 0}
                                        _ = v16
                                    } else {
                                        v17 := GenS1{F0: 2147483648, F1: 4294967295, F2: ""}
                                        if (false && true) {
                                            v2 = 3.14159
                                        }
                                        _ = v17
                                    }
                                    v3 = float64(((-2.25 * (((((v2 * v2) - ((-2.25 - (1.5 / ((1.5 - 0.1) / (float32(float64(float32(((v3 * float64(v2)) - v3)))) - 3.14159)))) + 0.0)) + (((1.5 / 1e10) * (1.5 / v2)) * (((v2 / v2) * (v2 - (((0.1 / float32(((0.1 - v3) + v3))) - ((1e10 * (((0.1 - float32((v3 - -2.25))) - (float32(((1e10 + v3) - v3)) * v2)) - v2)) * v2)) * v2))) + v2))) / ((float32((1.5 - 0.0)) - v2) * 0.0)) + (v2 + v2))) / -2.25))
                                }
                                switch -875 {
                                case 862:
                                    var v18 uint8 = 128
                                    if (!false) {
                                        fmt.Println(v2, v18)
                                    } else {
                                        _ = len(v1)
                                        fmt.Println(fi9, v2)
                                        if (("" + "hello") < "") {
                                            var v19 uint16 = 452
                                            _ = v19
                                        }
                                    }
                                    v3 = (v3 - 1.5)
                                    _ = v18
                                case 32767:
                                    fmt.Println(v3, fi7, fi7)
                                case 127:
                                    if false {
                                        if len(v4) > 0 {
                                            v4[0] = 4294967295
                                        }
                                        var v20 int16 = 685
                                        var v21 uint32 = 560
                                        _ = v20
                                        _ = v21
                                    }
                                default:
                                    switch "hello" {
                                    case "hello":
                                        if (false && (909 == fi9)) {
                                            fmt.Println(fi9, fi7, v2)
                                        }
                                        if (619 == 65535) {
                                            if false {
                                                fmt.Println(fi9)
                                            } else {
                                                v22 := GenS2{F0: 1}
                                                v2 -= 3.14159
                                                _ = v22
                                            }
                                            fmt.Println(v3)
                                            var v23 int8 = -81
                                            _ = v23
                                        } else {
                                            if (true && true) {
                                                v2 = 0.1
                                            }
                                        }
                                    default:
                                        if true {
                                            if true {
                                                v24 := GenS1{F0: 1, F1: 42, F2: "foo bar"}
                                                if (!false) {
                                                    v2 += v2
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi9 + 127)
                                                    }
                                                }
                                                _ = v24
                                            } else {
                                                _ = len(v1)
                                                fmt.Println(fi7, v2)
                                            }
                                        } else {
                                            v25 := GenS1{F0: 65535, F1: 0, F2: (("a" + "") + "")}
                                            _ = v25
                                        }
                                    }
                                }
                            }
                        }
                    }
                default:
                    v26 := GenS2{F0: -5}
                    _ = v26
                }
                if (169 > 255) {
                    v2 = v2
                    v27 := GenS3{F0: 32767, F1: 865}
                    _ = v27
                } else {
                    v28 := GenS4{F0: 155}
                    if true {
                        if false {
                            v29 := GenS1{F0: 719, F1: 2, F2: "hello"}
                            fmt.Println(v2)
                            switch "hello" {
                            case "foo bar":
                                if true {
                                    var v30 bool = (65535 < 42)
                                    switch "" {
                                    case "foo bar":
                                        v31 := GenS2{F0: -60}
                                        _ = v31
                                    case "\n\t":
                                        var v32 uint8 = 128
                                        _ = v32
                                    case "a":
                                        fmt.Println(fi7, v30, fi7)
                                    default:
                                        if v30 {
                                            v33 := GenS1{F0: 65535, F1: 990, F2: "hello"}
                                            v30 = false
                                            if v30 {
                                                if len(v4) > 0 {
                                                    v4[0] = 2147483647
                                                }
                                            } else {
                                                v2 = v2
                                                v34 := GenS3{F0: fi9, F1: 0}
                                                fmt.Println(fi7, v3)
                                                _ = v34
                                            }
                                            _ = v33
                                        }
                                    }
                                    _ = v30
                                }
                            case "a":
                                if (!(false && ((3.14159 * (v2 - ((((1e10 / v2) + (v2 - (v2 * (v2 / float32(v3))))) + 1.5) + v2))) > v2))) {
                                    fmt.Println(fi7, v2, fi7)
                                    v35 := GenS5{F0: 2147483648, F1: 65535, F2: 42}
                                    var v36 int64 = int64(fi7)
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = v29
                        } else {
                            if (true || false) {
                                var v37 uint64 = 4294967295
                                _ = v37
                            }
                            if false {
                                v3 = 3.14159
                                fmt.Println(fi7, v3, v3)
                                v2 = 1.5
                            } else {
                                switch fi9 {
                                case 32768:
                                    if (!true) {
                                        if ((v3 + (v3 / 1.0)) != float64(0.1)) {
                                            switch (fi9 / ((32768) | 1)) {
                                            case 2147483647:
                                                var v38 int8 = -128
                                                _ = v38
                                            case 128:
                                                v4 = append(v4, fi9)
                                            case 42:
                                                switch fi7 {
                                                case 42:
                                                    var v39 float32 = ((((v2 / (-2.25 + float32(3.14159))) + 1.5) - v2) / (v2 / (v2 + (0.0 / ((1.5 / v2) / v2)))))
                                                    _ = v39
                                                default:
                                                    var v40 float64 = v3
                                                    _ = v40
                                                }
                                            default:
                                                v41 := GenS2{F0: 61}
                                                _ = v41
                                            }
                                        }
                                        var v42 bool = true
                                        v43 := GenS2{F0: 1}
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v44 := GenS1{F0: 2147483648, F1: 202, F2: "foo bar"}
                                        _ = v44
                                    }
                                case -94:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -2147483648:
                                    var v45 float64 = v3
                                    _ = v45
                                }
                                v3 = (((((0.0 + v3) + ((3.14159 - float64(v2)) + v3)) + v3) + 0.0) + 0.0)
                            }
                            if true {
                                var v46 float64 = v3
                                _ = v46
                            }
                            v47 := GenS4{F0: 160}
                            if true {
                                var v48 uint64 = 42
                                v49 := GenS3{F0: 0, F1: 2147483647}
                                _ = v48
                                _ = v49
                            }
                            _ = v47
                        }
                        _ = len(v4)
                    } else {
                        v3 = float64(v2)
                    }
                    v1 = append(v1, 4294967295)
                    _ = v28
                }
            } else {
                var v50 int16 = 1000
                _ = v50
            }
            var v51 string = "a"
            if ((1.5 / v2) != (1.5 + v2)) {
                if false {
                    var v52 int32 = 42
                    var v53 float32 = v2
                    _ = v52
                    _ = v53
                } else {
                    fmt.Println(v3)
                    fmt.Println(fi7)
                }
            } else {
                v54 := GenS2{F0: 0}
                if true {
                    switch v51 {
                    case "a":
                        if ((((v3 + (-2.25 / (((v3 / float64(1.5)) * -2.25) / 1.5))) + v3) / (1.5 - -2.25)) >= -2.25) {
                            v2 = (0.1 - v2)
                            if true {
                                switch fi9 {
                                case 1:
                                    var v55 uint8 = 128
                                    _ = v55
                                }
                                switch fi7 {
                                case -1:
                                    v56 := GenS2{F0: genFn0(439, 0)}
                                    _ = v56
                                default:
                                    v57 := GenS1{F0: 175, F1: 42, F2: v51}
                                    _ = v57
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v58 int = (fi7 | fi7)
                                _ = v58
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v2 = float32(-2.25)
                        } else {
                            v59 := GenS3{F0: (^(fi7)), F1: -1}
                            var v60 int64 = 403
                            _ = v59
                            _ = v60
                        }
                    case "\n\t":
                        if true {
                            var v61 bool = (557 > 1)
                            if v61 {
                                v62 := GenS4{F0: 174}
                                _ = v62
                            }
                            v61 = (441 >= 1)
                            _ = v61
                        }
                    default:
                        if (((true || false) && ((((float64(v2) / -2.25) - (v3 - float64(v2))) - ((float64(v2) + 1.5) + (v3 * v3))) != v3)) && false) {
                            var v63 bool = (-119 != 42)
                            _ = v63
                        } else {
                            fmt.Println(v2)
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = fi9
                    }
                }
                _ = v54
            }
            _ = fi9
            _ = v51
        }
        _ = fi7
        _ = v8
    }
    if ((false || (-32768 == 1)) || (2 > 65535)) {
        fmt.Println(v2)
    } else {
        if true {
            v64 := GenS6{F0: 1, F1: 42}
            _ = v64
        } else {
            switch 313 {
            case -984:
                if len(v4) > 0 {
                    v4[0] = 38
                }
            case 128:
                v65 := GenS2{F0: 100}
                _ = v65
            }
            v2 -= float32(float64(-2.25))
        }
        fmt.Println(v2, v2)
        var v66 int8 = 52
        _ = v66
    }
    for fi67 := 0; fi67 < 1; fi67++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = 1e10
        for fi68 := 0; fi68 < 20; fi68++ {
            if true {
                v2 += (((v2 * float32(v3)) * float32(v3)) - ((v2 / (v2 * v2)) / v2))
                v1 = append(v1, 65535)
            } else {
                fmt.Println(v3)
                var v69 uint64 = 4294967295
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v69
            }
            _ = fi68
        }
        _ = fi67
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        var v70 float64 = (v3 / (3.14159 / (((1.5 * v3) + 1.5) / (0.1 * (v3 + 3.14159)))))
        _ = v70
    case "\n\t":
        for fi71 := 0; fi71 < 5; fi71++ {
            for fi72 := 0; fi72 < 12; fi72++ {
                v3 = v3
                switch ("\n\t" + "\n\t") {
                case "hello":
                    switch 346 {
                    case -790:
                        fmt.Println(v2, fi71, fi72)
                    default:
                        if (true && false) {
                            v3 = 1e10
                            if (!true) {
                                v73 := GenS7{F0: -55, F1: "", F2: 2}
                                if len(v1) > 0 {
                                    v1[0] = 989
                                }
                                if ((!true) || true) {
                                    var v74 uint = 4294967295
                                    if (0 <= 128) {
                                        var v75 float32 = -2.25
                                        var v76 uint16 = 32767
                                        var v77 int64 = -1
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                    }
                                    v78 := GenS8{F0: (v74 >> ((v74) & 63))}
                                    _ = v74
                                    _ = v78
                                } else {
                                    switch (fi72 - fi72) {
                                    case 9223372036854775807:
                                        v2 -= ((v2 - v2) + v2)
                                    }
                                    if true {
                                        switch fi72 {
                                        case 32768:
                                            if (2147483647 >= 42) {
                                                fmt.Println(fi72, v3, v3)
                                            }
                                            if (1 == -1) {
                                                v4 = append(v4, (fi71 + fi71))
                                                var v79 bool = false
                                                _ = v79
                                            } else {
                                                if (-2147483648 < 1) {
                                                    v3 = (float64(0.0) - float64((((((((v2 - float32(3.14159)) - v2) + v2) * -2.25) + float32(v3)) / v2) / 1.0)))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if (true && (v3 != v3)) {
                                                        v3 = (v3 - float64(0.0))
                                                    }
                                                } else {
                                                    if false {
                                                        if false {
                                                            var v80 int = 522
                                                            _ = v80
                                                        }
                                                    }
                                                }
                                                v81 := GenS2{F0: 88}
                                                _ = v81
                                            }
                                            var v82 string = ("hello" + "")
                                            _ = v82
                                        case 1:
                                            _ = len(v4)
                                        }
                                        if (127 <= 181) {
                                            fmt.Println(v2)
                                            var v83 float32 = v2
                                            _ = v83
                                        }
                                    } else {
                                        var v84 int = -9223372036854775808
                                        v3 -= float64(3.14159)
                                        if false {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (((1 < -354) && false) && (!(1 != 723))) {
                                                    v84 = v84
                                                    v85 := GenS5{F0: 827, F1: 2147483647, F2: 736}
                                                    _ = v85
                                                } else {
                                                    if true {
                                                        var v86 uint16 = 255
                                                        if ((false && true) || (18446744073709551615 != 237)) {
                                                            if true {
                                                                if (true || true) {
                                                                    var v87 float32 = (0.1 / float32(float64(v2)))
                                                                    switch 0 {
                                                                    case 4294967295:
                                                                        var v88 float64 = v3
                                                                        _ = v88
                                                                    default:
                                                                        v86 = 32767
                                                                    }
                                                                    _ = v87
                                                                } else {
                                                                    if false {
                                                                        v86 = v86
                                                                        v86 -= (v86 ^ 47)
                                                                    } else {
                                                                        switch v84 {
                                                                        case -9223372036854775808:
                                                                            fmt.Println(fi72, fi71, fi71)
                                                                        case 285:
                                                                            v3 = (v3 / (v3 * v3))
                                                                        case 0:
                                                                            var v89 int = (-(fi71))
                                                                            _ = v89
                                                                        }
                                                                        var v90 int32 = 825
                                                                        v4 = append(v4, -808)
                                                                        _ = v90
                                                                    }
                                                                    v84 = fi72
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 0
                                                                    }
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    fmt.Println(v2)
                                                                case "foo bar":
                                                                    v1 = append(v1, 0)
                                                                }
                                                            }
                                                        }
                                                        v84 += fi71
                                                        _ = v86
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = 2147483647
                                                        }
                                                        if true {
                                                            switch fi71 {
                                                            case 2147483647:
                                                                switch (fi71 << ((2) & 63)) {
                                                                case 985:
                                                                    v1 = append(v1, 720)
                                                                case 4294967295:
                                                                    if (128 == 132) {
                                                                        v91 := GenS6{F0: 0, F1: -2147483648}
                                                                        var v92 uint16 = 32767
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                case -1:
                                                                    var v93 uint32 = 2147483647
                                                                    _ = v93
                                                                default:
                                                                    if false {
                                                                        switch fi72 {
                                                                        case 630:
                                                                            switch "" {
                                                                            case "":
                                                                                _ = len(v4)
                                                                            }
                                                                            fmt.Println(v84)
                                                                        case 9223372036854775807:
                                                                            v94 := GenS7{F0: 202, F1: "foo bar", F2: 255}
                                                                            _ = v94
                                                                        default:
                                                                            var v95 int16 = 1
                                                                            _ = v95
                                                                        }
                                                                        fmt.Println(fi71, v2)
                                                                        v96 := GenS9{F0: 1, F1: 1}
                                                                        _ = v96
                                                                    }
                                                                }
                                                            case 42:
                                                                if (false || ((42 > 65) && (false || true))) {
                                                                    if false {
                                                                        var v97 uint64 = 139
                                                                        var v98 uint32 = 436
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                }
                                                            case 0:
                                                                v99 := GenS10{F0: 874, F1: 32767}
                                                                _ = v99
                                                            default:
                                                                v3 -= 0.0
                                                            }
                                                            v3 += (v3 / v3)
                                                        } else {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                fmt.Println(fi71, fi72)
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                var v100 int8 = 127
                                                                _ = v100
                                                            default:
                                                                v2 = (float32(v3) - -2.25)
                                                            }
                                                            if (((((((v2 * (float32(0.1) + ((v2 + (((3.14159 * 0.1) - (0.1 * (v2 + v2))) / (((v2 + v2) - (0.0 / ((1.5 - (v2 - 1.5)) / (v2 / 3.14159)))) * v2))) + v2))) / 0.1) - 0.1) / v2) * (0.0 / v2)) == (float32(float64(v2)) / (((v2 - 0.0) + float32(1.5)) - ((-2.25 * 3.14159) + float32((float64(v2) - v3)))))) || (1 <= 65535)) {
                                                                if ((fi72 % ((fi72) | 1)) < (^(fi72))) {
                                                                    fmt.Println(v2)
                                                                }
                                                                v84 += fi71
                                                            }
                                                        }
                                                    }
                                                    v2 -= float32(((((v3 / float64(v2)) - v3) * ((-2.25 / (v3 + ((v3 - float64(0.0)) + 0.0))) + ((((1.5 / -2.25) / v3) * v3) * v3))) / v3))
                                                    if false {
                                                        v3 = float64(v2)
                                                        if len(v1) > 0 {
                                                            v1[0] = 583
                                                        }
                                                        fmt.Println(v84, v84, fi71)
                                                        v101 := GenS6{F0: 201, F1: 42}
                                                        _ = v101
                                                    }
                                                }
                                            case "a":
                                                switch "\n\t" {
                                                case "a":
                                                    if false {
                                                        if (0 <= 4294967295) {
                                                            var v102 int = -2147483648
                                                            _ = v102
                                                        }
                                                        v3 -= (0.0 / (v3 + ((v3 + v3) - v3)))
                                                        switch fi71 {
                                                        case -826:
                                                            _ = len(v4)
                                                        default:
                                                            v103 := GenS9{F0: -1, F1: -1}
                                                            _ = v103
                                                        }
                                                    } else {
                                                        v104 := GenS4{F0: 35}
                                                        _ = v104
                                                    }
                                                case "foo bar":
                                                    v3 -= 0.1
                                                case "hello":
                                                    v105 := GenS11{F0: 42, F1: 0}
                                                    _ = v105
                                                }
                                            }
                                            var v106 uint64 = 2
                                            var v107 int32 = 12
                                            _ = len(v1)
                                            _ = v106
                                            _ = v107
                                        }
                                        _ = v84
                                    }
                                    if ((0.0 * v2) <= (1.5 / (3.14159 + v2))) {
                                        switch ("a" + "") {
                                        case "hello":
                                            switch fi72 {
                                            case 9223372036854775807:
                                                if false {
                                                    _ = len(v4)
                                                    if (4294967295 >= 1) {
                                                        var v108 uint = 65535
                                                        _ = v108
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = 838
                                                        }
                                                        var v109 uint16 = 290
                                                        if true {
                                                            if (false && (!(-678 == 1))) {
                                                                if true {
                                                                    var v110 int8 = 42
                                                                    _ = len(v4)
                                                                    v109 -= (^(v109))
                                                                    _ = v110
                                                                }
                                                            } else {
                                                                if true {
                                                                    var v111 int64 = 9223372036854775807
                                                                    v3 = (1.5 + v3)
                                                                    _ = v111
                                                                } else {
                                                                    v112 := GenS12{F0: 0, F1: 0, F2: ((("foo bar" + ("" + "hello")) + (("" + ((("a" + ("" + ("\n\t" + "hello"))) + "") + "foo bar")) + "foo bar")) + ("a" + ""))}
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v112
                                                                }
                                                                var v113 int64 = 0
                                                                var v114 float32 = v2
                                                                var v115 string = "\n\t"
                                                                fmt.Println(v3, fi72)
                                                                _ = v113
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                        } else {
                                                            var v116 uint16 = (v109 / ((v109) | 1))
                                                            fmt.Println(v116, v2, v2)
                                                            var v117 int8 = 127
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        _ = v109
                                                    }
                                                    v118 := GenS13{F0: 160}
                                                    v119 := GenS3{F0: (-(fi72)), F1: -864}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v2 -= 3.14159
                                                    case "foo bar":
                                                        if false {
                                                            fmt.Println(fi72)
                                                            v2 -= v2
                                                        }
                                                    case "hello":
                                                        v3 += float64((3.14159 * float32(v3)))
                                                    }
                                                    _ = v118
                                                    _ = v119
                                                } else {
                                                    v2 = 1e10
                                                    v3 = (((v3 * (v3 * float64(-2.25))) * (v3 / v3)) - v3)
                                                }
                                            case 2147483647:
                                                var v120 uint8 = 150
                                                _ = v120
                                            case -1:
                                                v1 = append(v1, 554)
                                                switch "hello" {
                                                case "\n\t":
                                                    v2 -= v2
                                                }
                                                v121 := GenS14{F0: true, F1: 576}
                                                _ = v121
                                            default:
                                                switch (fi72 % (((fi71 | fi72)) | 1)) {
                                                case -1:
                                                    switch (fi71 % (((fi71 + fi71)) | 1)) {
                                                    case 32768:
                                                        if false {
                                                            v2 = v2
                                                        }
                                                    case 2147483647:
                                                        v2 = v2
                                                    default:
                                                        var v122 uint64 = 42
                                                        _ = v122
                                                    }
                                                case -2147483648:
                                                    v2 = v2
                                                }
                                            }
                                        case "a":
                                            v2 += (float32(float64(((float32((v3 * 1.5)) + (v2 / 3.14159)) * v2))) - (3.14159 - (((0.0 - v2) - v2) * 3.14159)))
                                        case "":
                                            if (-9223372036854775808 < 1) {
                                                v2 = v2
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = len(v1)
                                            }
                                        }
                                        switch ("foo bar" + ("hello" + (("foo bar" + "a") + ("foo bar" + ("\n\t" + ""))))) {
                                        case "hello":
                                            _ = len(v1)
                                        case "":
                                            switch 2147483647 {
                                            case 2147483647:
                                                if false {
                                                    var v123 float64 = v3
                                                    v123 -= ((v3 * (0.1 + (v3 - v3))) * v3)
                                                    v3 = (v123 * 1e10)
                                                    _ = v123
                                                }
                                            case 127:
                                                v3 = v3
                                            default:
                                                v124 := GenS4{F0: 1}
                                                _ = v124
                                            }
                                        }
                                        switch fi72 {
                                        case -136:
                                            var v125 string = "foo bar"
                                            _ = v125
                                        }
                                    }
                                }
                                if (true && false) {
                                    v2 = 1.5
                                    v2 += ((v2 - v2) - v2)
                                    if (643 != 995) {
                                        v3 -= -2.25
                                        fmt.Println(fi71, fi72, fi71)
                                        v2 -= (1.5 / v2)
                                    } else {
                                        v2 = 1e10
                                        fmt.Println(fi71, fi72, v2)
                                        var v126 int8 = -55
                                        _ = v126
                                    }
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483648
                                    }
                                }
                                v1 = append(v1, 2147483648)
                                _ = v73
                            }
                        }
                    }
                case "":
                    fmt.Println(fi71)
                default:
                    fmt.Println(v3, v2)
                    v3 -= (float64(v2) + v3)
                    v3 += v3
                }
                v127 := GenS7{F0: -2147483648, F1: "hello", F2: 1}
                _ = fi72
                _ = v127
            }
            if true {
                for fi128 := 0; fi128 < 12; fi128++ {
                    var v129 string = "hello"
                    _ = fi128
                    _ = v129
                }
            } else {
                fmt.Println(fi71, v3, fi71)
            }
            if (fi71 >= -1) {
                for fi130 := 0; fi130 < 8; fi130++ {
                    switch (fi130 + 2147483647) {
                    case -9223372036854775808:
                        fmt.Println(v3, fi130, fi71)
                        if false {
                            v131 := GenS15{F0: 2}
                            _ = v131
                        }
                        var v132 bool = (0 != (fi71 / ((fi71) | 1)))
                        _ = v132
                    default:
                        fmt.Println(fi71, v3)
                    }
                    _ = fi130
                }
                for fi133 := 0; fi133 < 14; fi133++ {
                    if ((-9223372036854775808 > 861) || true) {
                        if false {
                            var v134 int64 = 207
                            _ = v134
                        }
                        if (349 != 1) {
                            var v135 string = "hello"
                            v2 -= v2
                            if true {
                                _ = len(v4)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    v135 = v135
                                    var v136 string = v135
                                    _ = v136
                                } else {
                                    fmt.Println(v135)
                                    var v137 uint64 = 42
                                    var v138 int64 = -584
                                    _ = v137
                                    _ = v138
                                }
                            } else {
                                v3 = (v3 / (v3 / (float64((float32(3.14159) + ((3.14159 / v2) / -2.25))) / v3)))
                                fmt.Println(v135, v135)
                                _ = len(v1)
                            }
                            _ = v135
                        } else {
                            v139 := GenS16{F0: (!false), F1: 100}
                            _ = v139
                        }
                    }
                    v2 += float32(1e10)
                    fmt.Println(fi71)
                    if len(v1) > 0 {
                        v1[0] = 619
                    }
                    _ = fi133
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch fi71 {
                case -1:
                    v140 := GenS17{F0: 42}
                    _ = v140
                case 2147483647:
                    var v141 float32 = float32((v3 - 0.1))
                    _ = v141
                }
            }
            _ = fi71
        }
    default:
        for fi142 := 0; fi142 < 12; fi142++ {
            var v143 float64 = v3
            _ = fi142
            _ = v143
        }
    }
    v3 = (v3 / (v3 / v3))
    fmt.Println(v3, v3)
    v2 += v2
    var v144 uint8 = 164
    for fi145 := 0; fi145 < 16; fi145++ {
        fmt.Println(v144, v3, v2)
        var v146 int8 = 42
        v147 := GenS7{F0: 0, F1: "hello", F2: 1}
        _ = fi145
        _ = v146
        _ = v147
    }
    var v148 int = 4294967295
    v2 = v2
    if true {
        for fi149 := 0; fi149 < 19; fi149++ {
            v2 += float32((((v3 + v3) / (((-2.25 * v3) - v3) + float64(v2))) * v3))
            v2 = float32(v3)
            _ = fi149
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v1 = append(v1, 734)
    }
    if (false || (2147483647 < 472)) {
        v1 = append(v1, 42)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v144:", v144)
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v144
    _ = v148
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
