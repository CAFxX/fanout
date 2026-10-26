package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int) int16 {
    v6 := []int{p0, (^(p0)), p0, p0}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    p0 = 491
    var v7 uint = 18446744073709551615
    _ = p0
    _ = v6
    _ = v7
    return -245
}

func main() {
    var v1 int8 = 127
    var v2 float64 = (3.14159 * (3.14159 + 1.5))
    var v3 int = -2147483648
    var v4 int8 = (^(v1))
    var v5 float64 = 1e10
    var v8 uint32 = 38
    var v9 string = "hello"
    v5 += (((v5 - (3.14159 / (((v5 / (3.14159 - (0.1 - (((float64(float32(float64(((float32(v2) + 1.5) / 1.5)))) * (v2 / 1.0)) - (3.14159 / (0.1 - v5))) + (v2 * ((v5 * v2) + -2.25)))))) + ((v2 * v2) * v2)) - 3.14159))) / (v2 + float64((1.5 - -2.25)))) * 3.14159)
    switch (v3 | 2147483647) {
    case 787:
        v1 -= v1
    default:
        if (true && ((((float32(0.1) * 0.1) + -2.25) + (-2.25 + 1.5)) > 3.14159)) {
            var v10 int8 = v4
            _ = v10
        } else {
            v4 += (v4 + (v1 ^ v4))
        }
    }
    v3 = 128
    v11 := []int{}
    var v12 int16 = -21
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(v8, v5)
        switch (fi13 - v3) {
        case 42:
            v11 = append(v11, fi13)
        case -1:
            v8 = 0
        }
        var v14 bool = (false && (2 == 255))
        _ = fi13
        _ = v14
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v15 := GenS1{F0: 4294967295, F1: 2147483647, F2: v1}
    if len(v11) > 0 {
        v11[0] = (-(v3))
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v8 = (v8 & (v8 / ((883) | 1)))
        if (!false) {
            switch v9 {
            case "foo bar":
                switch "" {
                case "\n\t":
                    var v17 int = fi16
                    _ = v17
                case "foo bar":
                    v18 := GenS2{F0: 65535, F1: 355}
                    _ = v18
                case "":
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
            case "a":
                v9 = ("hello" + ((v9 + v9) + v9))
            }
        } else {
            if (-1 == (-(v4))) {
                v9 = ((("" + ((("a" + ("a" + (v9 + (v9 + v9)))) + v9) + (v9 + ("foo bar" + ("\n\t" + ("\n\t" + v9)))))) + v9) + (v9 + v9))
                var v19 uint16 = 972
                if (-718 != 2147483647) {
                    fmt.Println(v3, v3)
                } else {
                    fmt.Println(v9, v9, v8)
                }
                _ = v19
            } else {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                if (float32(0.1) == 1.5) {
                    if len(v11) > 0 {
                        v11[0] = 4294967295
                    }
                    if (false || ((0 >= 0) && false)) {
                        switch fi16 {
                        case 0:
                            v5 -= float64(3.14159)
                        case 1:
                            var v20 float32 = (0.1 - float32(v2))
                            _ = v20
                        default:
                            v9 = v9
                        }
                        var v21 bool = true
                        if v21 {
                            var v22 int = v3
                            _ = v22
                        }
                        _ = v21
                    }
                } else {
                    if false {
                        if (!false) {
                            fmt.Println(v3, v9, v9)
                            v9 = "\n\t"
                            var v23 uint16 = 456
                            _ = v23
                        }
                        switch (v9 + ("\n\t" + ((("a" + v9) + v9) + (("hello" + (((v9 + v9) + "") + ("foo bar" + v9))) + v9)))) {
                        case "a":
                            v1 -= v1
                        case "\n\t":
                            if true {
                                v24 := GenS1{F0: (v8 % ((42) | 1)), F1: 42, F2: 67}
                                _ = v24
                            }
                        case "":
                            v25 := GenS1{F0: (^(v8)), F1: -588, F2: (v1 << ((0) & 63))}
                            _ = v25
                        default:
                            switch (^(fi16)) {
                            case -9223372036854775808:
                                v4 -= v4
                            case -620:
                                fmt.Println(v12, v4)
                            default:
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    v2 = v5
                                    _ = fi26
                                }
                            }
                        }
                        var v27 int = -9223372036854775808
                        _ = v27
                    } else {
                        v12 -= genFn0(0)
                    }
                }
            }
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi16
    }
    v28 := GenS1{F0: 2147483648, F1: -2147483648, F2: v4}
    for fi29 := 0; fi29 < 5; fi29++ {
        var v30 uint = 834
        _ = fi29
        _ = v30
    }
    if true {
        for fi31 := 0; fi31 < 18; fi31++ {
            if false {
                if ((0.1 + 3.14159) > 1.5) {
                    for fi32 := 0; fi32 < 8; fi32++ {
                        var v33 float64 = -2.25
                        var v34 uint = 18446744073709551615
                        if false {
                            var v35 bool = true
                            v12 -= v12
                            switch fi31 {
                            case -576:
                                v1 = v1
                            case 189:
                                if true {
                                    v3 = -9223372036854775808
                                } else {
                                    v33 -= v5
                                    v11 = append(v11, (v3 | 127))
                                    var v36 string = "\n\t"
                                    _ = v36
                                }
                            case -9223372036854775808:
                                var v37 int8 = v1
                                _ = v37
                            default:
                                var v38 uint8 = 255
                                _ = v38
                            }
                            _ = v35
                        }
                        var v39 float32 = 3.14159
                        var v40 int = 1
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v39
                        _ = v40
                    }
                } else {
                    v5 += v2
                    if (v3 < v3) {
                        var v41 uint64 = 953
                        v9 = v9
                        v3 = (v3 * 9223372036854775807)
                        _ = v41
                    }
                    fmt.Println(v4, v3, fi31)
                }
                switch (("foo bar" + (v9 + v9)) + (v9 + v9)) {
                case "":
                    _ = len(v11)
                case "a":
                    var v42 uint32 = 2
                    _ = v42
                case "hello":
                    for fi43 := 0; fi43 < 8; fi43++ {
                        if (1000 >= (v12 + -1)) {
                            var v44 uint8 = 128
                            var v45 uint8 = v44
                            var v46 uint32 = 923
                            if false {
                                v12 += v12
                            }
                            v9 = v9
                            _ = v44
                            _ = v45
                            _ = v46
                        } else {
                            var v47 int32 = 931
                            v4 = v1
                            _ = v47
                        }
                        var v48 float64 = -2.25
                        fmt.Println(v4)
                        _ = fi43
                        _ = v48
                    }
                default:
                    switch fi31 {
                    case -2147483648:
                        for fi49 := 0; fi49 < 9; fi49++ {
                            var v50 uint64 = 4294967295
                            if false {
                                if (genFn0(v3) == genFn0(32767)) {
                                    v11 = append(v11, -142)
                                    v50 -= (v50 - v50)
                                    v51 := GenS3{F0: (v12 / ((32767) | 1)), F1: 1, F2: false}
                                    var v52 bool = false
                                    switch fi49 {
                                    case 4294967295:
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v2, v12, v8)
                                    default:
                                        if v52 {
                                            if (v52 || (("" > "\n\t") || (float64(1e10) > ((0.0 - v2) * v5)))) {
                                                v11 = append(v11, 35)
                                                v53 := GenS1{F0: uint32(v50), F1: 0, F2: v4}
                                                _ = v53
                                            } else {
                                                if (!false) {
                                                    switch v3 {
                                                    case 670:
                                                        v54 := GenS3{F0: v12, F1: 255, F2: (!(100 != v4))}
                                                        _ = v54
                                                    case 1:
                                                        v11 = append(v11, fi49)
                                                    case 514:
                                                        v4 -= v4
                                                    default:
                                                        v55 := GenS1{F0: (v8 & (^(v8))), F1: 1, F2: v4}
                                                        _ = v55
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        if true {
                                                            v56 := GenS2{F0: 65535, F1: (v50 % ((uint64(v8)) | 1))}
                                                            v57 := GenS3{F0: genFn0((fi31 ^ 767)), F1: 2, F2: v52}
                                                            v58 := GenS4{F0: 255}
                                                            v12 += genFn0(32768)
                                                            switch v9 {
                                                            case "":
                                                                fmt.Println(v8)
                                                            case "\n\t":
                                                                switch v9 {
                                                                case "a":
                                                                    var v59 string = v9
                                                                    _ = v59
                                                                case "\n\t":
                                                                    var v60 bool = (!(v52 || v52))
                                                                    _ = v60
                                                                }
                                                            }
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    case "\n\t":
                                                        switch v9 {
                                                        case "\n\t":
                                                            v61 := GenS2{F0: 2, F1: v50}
                                                            _ = v61
                                                        default:
                                                            if false {
                                                                v62 := GenS3{F0: (v12 >> ((18446744073709551615) & 63)), F1: uint8(v50), F2: v52}
                                                                var v63 float32 = ((1e10 / ((1e10 / 3.14159) * 1e10)) + ((((0.1 / float32(0.1)) + -2.25) - 3.14159) + (0.1 - (float32(-2.25) + (float32((3.14159 + (-2.25 * (((v2 - 0.0) * 3.14159) / v5)))) / 0.1)))))
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                        }
                                                    case "hello":
                                                        fmt.Println(v8, v8, v5)
                                                    default:
                                                        fmt.Println(v12)
                                                    }
                                                } else {
                                                    v64 := GenS5{F0: -1, F1: 135, F2: 233}
                                                    _ = v64
                                                }
                                                if false {
                                                    v9 = ("hello" + (((v9 + v9) + "hello") + v9))
                                                    v65 := GenS5{F0: 0, F1: 255, F2: 2}
                                                    _ = v65
                                                }
                                            }
                                        }
                                    }
                                    v3 = (fi31 * (v3 & (fi49 ^ (fi31 & fi31))))
                                    _ = len(v11)
                                    _ = v51
                                    _ = v52
                                }
                                if ((v12 * v12) == v12) {
                                    v12 -= v12
                                    if true {
                                        v66 := GenS2{F0: 754, F1: (v50 >> ((255) & 63))}
                                        if (!false) {
                                            _ = len(v11)
                                            if (v1 >= -57) {
                                                v67 := GenS2{F0: 681, F1: v50}
                                                _ = v67
                                            } else {
                                                if (0.0 <= 3.14159) {
                                                    switch "" {
                                                    case "hello":
                                                        var v68 uint = 42
                                                        _ = v68
                                                    }
                                                } else {
                                                    v4 = 10
                                                    fmt.Println(v4, v3, v8)
                                                }
                                                var v69 int64 = 9223372036854775807
                                                var v70 float64 = v2
                                                var v71 int16 = -1
                                                var v72 uint = 951
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                        var v73 int = (fi49 % ((128) | 1))
                                        v50 = v50
                                        if false {
                                            v74 := GenS5{F0: v12, F1: 255, F2: 128}
                                            _ = v74
                                        } else {
                                            var v75 int32 = 1
                                            v76 := GenS3{F0: v12, F1: 255, F2: true}
                                            if (false && (v2 <= (-2.25 * v2))) {
                                                _ = len(v11)
                                            } else {
                                                if false {
                                                    if (!true) {
                                                        if true {
                                                            fmt.Println(v4, fi49, v75)
                                                            v77 := GenS6{F0: 0, F1: 814, F2: 255}
                                                            var v78 float64 = (1.5 / v2)
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    }
                                                    v79 := GenS4{F0: 21}
                                                    v75 += 555
                                                    _ = v79
                                                } else {
                                                    _ = len(v11)
                                                    v73 += int(v75)
                                                    fmt.Println(v9)
                                                }
                                                var v80 string = "hello"
                                                v11 = append(v11, (v3 >> ((0) & 63)))
                                                _ = v80
                                            }
                                            _ = v75
                                            _ = v76
                                        }
                                        _ = v66
                                        _ = v73
                                    }
                                }
                            } else {
                                v4 = v4
                                fmt.Println(v50, v8, v8)
                                v12 += v12
                            }
                            _ = fi49
                            _ = v50
                        }
                    case 2147483647:
                        if false {
                            if (!false) {
                                var v81 float32 = (0.0 + (0.0 * float32((float64(1.5) * v2))))
                                switch (v9 + (v9 + ((v9 + v9) + ""))) {
                                case "":
                                    for fi82 := 0; fi82 < 18; fi82++ {
                                        var v83 int64 = -9223372036854775808
                                        var v84 int64 = (v83 | (v83 + 265))
                                        fmt.Println(v12, fi82)
                                        _ = fi82
                                        _ = v83
                                        _ = v84
                                    }
                                case "\n\t":
                                    v85 := GenS6{F0: 715, F1: 255, F2: 32767}
                                    _ = v85
                                case "hello":
                                    v12 -= v12
                                }
                                _ = v81
                            }
                            var v86 int64 = 42
                            _ = v86
                        } else {
                            if (!(0 < 1)) {
                                var v87 uint8 = 0
                                v88 := GenS7{F0: (v4 - 127)}
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        v89 := GenS6{F0: 42, F1: 32767, F2: 2}
                                        _ = v89
                                    case "":
                                        v90 := GenS4{F0: 32767}
                                        _ = v90
                                    default:
                                        v1 += 127
                                    }
                                    switch "" {
                                    case "":
                                        for fi91 := 0; fi91 < 15; fi91++ {
                                            switch fi31 {
                                            case 4294967295:
                                                var v92 uint32 = v8
                                                _ = v92
                                            case 0:
                                                if true {
                                                    v93 := GenS2{F0: 255, F1: 1}
                                                    _ = v93
                                                }
                                            case 763:
                                                fmt.Println(v5, fi91)
                                            default:
                                                if false {
                                                    if len(v11) > 0 {
                                                        v11[0] = 127
                                                    }
                                                }
                                            }
                                            _ = fi91
                                        }
                                    }
                                }
                                _ = v87
                                _ = v88
                            } else {
                                if false {
                                    var v94 uint = 2
                                    for fi95 := 0; fi95 < 7; fi95++ {
                                        if (!true) {
                                            var v96 uint = v94
                                            fmt.Println(v1)
                                            var v97 uint16 = 255
                                            _ = v96
                                            _ = v97
                                        } else {
                                            v5 += ((((float64(1e10) / float64((0.1 / 3.14159))) + v2) * 1.5) - ((((1.5 / 0.1) / (float64(0.1) + v5)) + (0.1 / v2)) / v5))
                                        }
                                        if len(v11) > 0 {
                                            v11[0] = (fi95 % (((-(fi31))) | 1))
                                        }
                                        switch v9 {
                                        case "":
                                            if (42 == 330) {
                                                var v98 float64 = v5
                                                _ = v98
                                            } else {
                                                v99 := GenS7{F0: v1}
                                                _ = v99
                                            }
                                        default:
                                            v100 := GenS6{F0: 455, F1: 65535, F2: 530}
                                            _ = v100
                                        }
                                        _ = fi95
                                    }
                                    _ = len(v11)
                                    _ = v94
                                } else {
                                    _ = len(v11)
                                    for fi101 := 0; fi101 < 9; fi101++ {
                                        v11 = append(v11, (fi101 >> ((2) & 63)))
                                        v102 := GenS1{F0: (v8 * v8), F1: 0, F2: 100}
                                        _ = fi101
                                        _ = v102
                                    }
                                }
                            }
                            if false {
                                fmt.Println(v8, v8, v1)
                            }
                            v12 = genFn0((v3 << ((80) & 63)))
                        }
                    case 249:
                        v11 = append(v11, (^(v3)))
                    default:
                        fmt.Println(v9, v2, v3)
                    }
                }
                _ = len(v11)
            }
            _ = fi31
        }
    } else {
        for fi103 := 0; fi103 < 20; fi103++ {
            fmt.Println(v2, fi103, v1)
            _ = fi103
        }
        var v104 int = v3
        _ = v104
    }
    for fi105 := 0; fi105 < 17; fi105++ {
        if (false || false) {
            switch "foo bar" {
            case "":
                var v106 int16 = (-(v12))
                _ = v106
            case "a":
                switch (v9 + ("\n\t" + v9)) {
                case "a":
                    if false {
                        if (255 >= 86) {
                            var v107 uint16 = 32767
                            _ = len(v11)
                            _ = v107
                        } else {
                            for fi108 := 0; fi108 < 15; fi108++ {
                                switch v9 {
                                case "":
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                case "a":
                                    if (v9 >= "\n\t") {
                                        var v109 float64 = v5
                                        _ = v109
                                    }
                                }
                                v3 -= (fi108 / ((fi105) | 1))
                                var v110 int64 = int64(v12)
                                _ = fi108
                                _ = v110
                            }
                        }
                    } else {
                        var v111 int = v3
                        var v112 bool = true
                        _ = v111
                        _ = v112
                    }
                case "foo bar":
                    v4 = 100
                case "":
                    v11 = append(v11, (v3 >> ((302) & 63)))
                }
            case "\n\t":
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            default:
                if true {
                    switch 32768 {
                    case 128:
                        v113 := GenS2{F0: 32767, F1: 4294967295}
                        _ = v113
                    case -2147483648:
                        var v114 uint = 1
                        _ = v114
                    default:
                        var v115 float64 = 0.1
                        _ = v115
                    }
                }
            }
        }
        _ = fi105
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
