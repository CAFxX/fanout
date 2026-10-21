package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int8; F2 int32 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint16; F1 int; F2 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
func genFn0(p0 int32, p1 uint16, p2 int32) uint32 {
    for fi4 := 0; fi4 < 1; fi4++ {
        if false {
            var v5 uint64 = 0
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 bool = ((p0 << ((18446744073709551615) & 63)) != (p2 & (p2 * p0)))
                var v8 int32 = p2
                if v7 {
                    v5 -= v5
                    fmt.Println(p0, p1)
                    switch 32768 {
                    case -1:
                        v9 := []int{fi6, (fi4 >> ((4294967295) & 63))}
                        _ = v9
                    }
                    p0 = 1
                    fmt.Println(v5, v5)
                    v10 := []int{fi4, fi4}
                    _ = v10
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        } else {
            v11 := []int{(fi4 - 128), fi4}
            p1 += (^(p1))
            _ = v11
        }
        _ = fi4
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int64 = 42
        if (2 == 255) {
            v14 := []int{fi12}
            switch (fi12 ^ (fi12 | -530)) {
            case 128:
                fmt.Println(v13)
            case 0:
                v13 += v13
            }
            _ = v14
        } else {
            v13 = (v13 * 0)
        }
        fmt.Println(p0)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 532
}
func genFn1(p0 int64) int {
    switch "" {
    case "hello":
        var v15 int16 = 42
        _ = v15
    }
    v16 := []int{0}
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint16 = 255
        v16 = append(v16, fi17)
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v16
    return 127
}

func main() {
    var v1 bool = ((false || (!false)) && (42 <= 2147483647))
    var v2 uint32 = 41
    var v3 bool = true
    var v19 float32 = 1e10
    if v3 {
        v2 = 65535
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(fi20, v3)
        _ = fi20
    }
    if (false && true) {
        var v21 float64 = (0.1 - 3.14159)
        v22 := []int{42}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        v24 := GenS1{F0: -1}
        if ((1 <= 786) || v1) {
            var v25 int16 = 42
            _ = v25
        }
        switch fi23 {
        case -2147483648:
            if v1 {
                var v26 bool = (1.5 == float64((0.0 / v19)))
                _ = v26
            } else {
                if ((1.5 - (v19 - (-2.25 * (-2.25 + v19)))) == v19) {
                    var v27 int = fi23
                    v28 := GenS1{F0: -702}
                    _ = v27
                    _ = v28
                }
            }
        case -580:
            if (!((!(-25 >= 1)) && v1)) {
                switch "" {
                case "\n\t":
                    v19 = (v19 / (v19 - (3.14159 * 1.5)))
                case "":
                    if v1 {
                        v1 = ("a" <= "")
                        fmt.Println(v19)
                        v29 := []int{(-(fi23))}
                        _ = v29
                    }
                case "foo bar":
                    v19 = -2.25
                default:
                    if ("hello" < ("" + "foo bar")) {
                        for fi30 := 0; fi30 < 8; fi30++ {
                            switch "a" {
                            case "hello":
                                switch "" {
                                case "a":
                                    v3 = ((float64(v19) + (1.5 / 0.1)) >= (1e10 * 1.5))
                                case "\n\t":
                                    v31 := GenS2{F0: (!true), F1: 2}
                                    _ = v31
                                }
                            case "a":
                                if v1 {
                                    v2 -= v2
                                    fmt.Println(v1, fi30, fi30)
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v32 bool = (false && (!(v3 || ((1.5 >= float32(0.1)) || ((v19 >= float32(0.1)) || (!(!v3)))))))
                                        _ = v32
                                    case "a":
                                        v33 := GenS1{F0: -1}
                                        _ = v33
                                    case "":
                                        v2 += v2
                                    default:
                                        fmt.Println(v1)
                                    }
                                    v34 := []int{(fi23 - (-(fi30))), fi30, 819, (fi23 + genFn1(42))}
                                    switch fi23 {
                                    case 1:
                                        var v35 int16 = 257
                                        _ = v35
                                    }
                                    switch ("" + (("a" + "a") + ((("foo bar" + "") + (("a" + "hello") + ("a" + "hello"))) + (("a" + "hello") + ("foo bar" + "a"))))) {
                                    case "foo bar":
                                        v36 := GenS2{F0: (42 != 42), F1: 18446744073709551615}
                                        _ = v36
                                    default:
                                        var v37 uint = 417
                                        _ = v37
                                    }
                                    v19 -= float32(0.0)
                                    var v38 int16 = 42
                                    v39 := GenS3{F0: -9223372036854775808, F1: 221, F2: 79}
                                    _ = v34
                                    _ = v38
                                    _ = v39
                                }
                            }
                            v40 := GenS3{F0: 42, F1: 32767, F2: uint8(v2)}
                            _ = fi30
                            _ = v40
                        }
                    } else {
                        if (2 == 42) {
                            for fi41 := 0; fi41 < 14; fi41++ {
                                var v42 int32 = 42
                                _ = fi41
                                _ = v42
                            }
                        }
                    }
                }
            }
        }
        _ = fi23
        _ = v24
    }
    switch genFn1(965) {
    case -1:
        fmt.Println(v2, v3, v3)
    case -829:
        var v43 string = "\n\t"
        _ = v43
    case 128:
        var v44 string = "hello"
        _ = v44
    default:
        switch genFn1(42) {
        case -2147483648:
            var v45 string = (("" + "a") + "foo bar")
            _ = v45
        case -610:
            switch -133 {
            case 4294967295:
                var v46 int64 = -9223372036854775808
                _ = v46
            case -576:
                v47 := []int{820, genFn1(0), -2147483648}
                _ = v47
            case 128:
                for fi48 := 0; fi48 < 7; fi48++ {
                    if v3 {
                        var v49 string = (((("\n\t" + "hello") + ((("\n\t" + ("" + "\n\t")) + "foo bar") + ("foo bar" + (((("hello" + "\n\t") + "foo bar") + "\n\t") + "")))) + ("a" + ("" + "foo bar"))) + "")
                        for fi50 := 0; fi50 < 7; fi50++ {
                            v2 = v2
                            var v51 float32 = float32((-2.25 * 1.5))
                            v3 = v1
                            _ = fi50
                            _ = v51
                        }
                        _ = v49
                    }
                    v52 := []int{}
                    v3 = v1
                    _ = fi48
                    _ = v52
                }
            }
        case 1:
            v3 = v3
        default:
            var v53 uint64 = 505
            _ = v53
        }
    }
    if false {
        if true {
            var v54 uint = 0
            for fi55 := 0; fi55 < 2; fi55++ {
                switch -866 {
                case -570:
                    switch "a" {
                    case "\n\t":
                        var v56 int16 = 32767
                        _ = v56
                    default:
                        if v3 {
                            switch "\n\t" {
                            case "":
                                var v57 uint = 75
                                _ = v57
                            }
                            if (42 <= 0) {
                                for fi58 := 0; fi58 < 18; fi58++ {
                                    var v59 int8 = 127
                                    _ = fi58
                                    _ = v59
                                }
                                if true {
                                    if ((fi55 / ((-1) | 1)) > genFn1(-9223372036854775808)) {
                                        var v60 float32 = ((v19 * v19) - 3.14159)
                                        v3 = ((v2 + 2) < (v2 % (((^(v2))) | 1)))
                                        v61 := GenS3{F0: -9223372036854775808, F1: 0, F2: 2}
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v62 := []int{(fi55 >> ((v54) & 63))}
                                        v3 = v1
                                        _ = v62
                                    }
                                }
                                v63 := GenS3{F0: 0, F1: 255, F2: 128}
                                _ = v63
                            }
                        } else {
                            v64 := []int{genFn1(-210), fi55, -93}
                            v1 = (v3 || ((v3 || (((float32(float64(v19)) * v19) + (3.14159 - 0.0)) > v19)) && v3))
                            _ = v64
                        }
                    }
                case -1:
                    if (v3 && v1) {
                        if (0.0 <= (float64(1.5) * 3.14159)) {
                            switch "\n\t" {
                            case "foo bar":
                                if (!v3) {
                                    var v65 int64 = -1
                                    if true {
                                        var v66 uint64 = 1
                                        var v67 int16 = -590
                                        _ = v66
                                        _ = v67
                                    }
                                    fmt.Println(v1)
                                    _ = v65
                                }
                            case "\n\t":
                                v1 = false
                            case "a":
                                v68 := GenS4{F0: 1, F1: 3, F2: -2147483648}
                                _ = v68
                            default:
                                if true {
                                    fmt.Println(v1)
                                    v69 := GenS5{F0: 9223372036854775807, F1: 1}
                                    _ = v69
                                }
                            }
                            var v70 int16 = -287
                            _ = v70
                        }
                    } else {
                        var v71 int32 = -2147483648
                        _ = v71
                    }
                }
                _ = fi55
            }
            v2 -= (^(v2))
            _ = v54
        } else {
            v19 = v19
        }
        fmt.Println(v1)
        var v72 float64 = 0.1
        _ = v72
    }
    fmt.Println(v19)
    for fi73 := 0; fi73 < 7; fi73++ {
        if true {
            fmt.Println(v1)
        } else {
            v19 += v19
        }
        var v74 float32 = (v19 - -2.25)
        _ = fi73
        _ = v74
    }
    switch genFn1(-584) {
    case 127:
        v75 := GenS4{F0: 255, F1: 1, F2: 1}
        _ = v75
    case 32768:
        switch "" {
        case "":
            switch "foo bar" {
            case "a":
                fmt.Println(v3)
            default:
                var v76 bool = (480 > 4294967295)
                _ = v76
            }
        default:
            if true {
                for fi77 := 0; fi77 < 13; fi77++ {
                    for fi78 := 0; fi78 < 15; fi78++ {
                        switch genFn1(48) {
                        case -422:
                            var v79 int8 = 1
                            _ = v79
                        default:
                            var v80 float64 = 1e10
                            _ = v80
                        }
                        if false {
                            v3 = ("\n\t" == "\n\t")
                            var v81 int8 = -1
                            _ = v81
                        }
                        _ = fi78
                    }
                    v82 := []int{(fi77 ^ (fi77 << ((678) & 63))), (fi77 / ((fi77) | 1))}
                    _ = fi77
                    _ = v82
                }
            }
        }
    default:
        v2 = genFn0(631, 506, -2147483648)
    }
    for fi83 := 0; fi83 < 9; fi83++ {
        for fi84 := 0; fi84 < 20; fi84++ {
            if (!(true && false)) {
                v85 := []int{}
                _ = len(v85)
                var v86 uint16 = 255
                _ = v85
                _ = v86
            } else {
                var v87 uint64 = 18446744073709551615
                _ = v87
            }
            var v88 int64 = -795
            _ = fi84
            _ = v88
        }
        if v3 {
            v89 := []int{(fi83 ^ fi83), 1}
            if v3 {
                var v90 uint = 787
                switch "foo bar" {
                case "foo bar":
                    switch fi83 {
                    case -9223372036854775808:
                        if v1 {
                            fmt.Println(v2, v90, v1)
                            v90 = v90
                            v1 = v1
                        } else {
                            v90 += v90
                            fmt.Println(v3, v19, v2)
                        }
                    default:
                        v90 = v90
                    }
                case "\n\t":
                    _ = len(v89)
                case "hello":
                    var v91 int8 = -47
                    _ = v91
                default:
                    fmt.Println(fi83, v1)
                }
                _ = v90
            } else {
                if ((-2.25 <= 0.0) && v1) {
                    switch (fi83 | fi83) {
                    case -388:
                        v92 := GenS6{F0: ("" + "hello"), F1: 4294967295}
                        _ = v92
                    case 247:
                        v2 = 930
                    case 0:
                        for fi93 := 0; fi93 < 2; fi93++ {
                            switch "" {
                            case "":
                                v3 = (v3 && (!false))
                            case "a":
                                fmt.Println(v19)
                            }
                            var v94 uint64 = 213
                            if len(v89) >= 2 {
                                v89 = v89[1:len(v89)]
                            }
                            switch "\n\t" {
                            case "foo bar":
                                if true {
                                    v89 = append(v89, (^(fi83)))
                                    fmt.Println(v3, v94)
                                    if (0 < 1) {
                                        fmt.Println(v19)
                                        v1 = v1
                                        if (v94 <= v94) {
                                            var v95 uint8 = 42
                                            _ = v95
                                        }
                                    } else {
                                        v89 = append(v89, fi93)
                                        if len(v89) >= 2 {
                                            v89 = v89[1:len(v89)]
                                        }
                                    }
                                }
                            case "":
                                if ((1 != 1) && (0 <= -540)) {
                                    if true {
                                        var v96 bool = (-77 <= -128)
                                        switch fi83 {
                                        case -548:
                                            v2 += v2
                                        default:
                                            var v97 string = "\n\t"
                                            _ = v97
                                        }
                                        _ = v96
                                    } else {
                                        var v98 float32 = -2.25
                                        v98 = (v19 / float32(float64(float32(3.14159))))
                                        _ = v98
                                    }
                                    switch fi83 {
                                    case 685:
                                        if ((fi93 ^ (fi93 - fi83)) > (fi93 - 541)) {
                                            if len(v89) > 0 {
                                                v89[0] = 4294967295
                                            }
                                            var v99 uint = 4294967295
                                            v100 := GenS4{F0: 0, F1: 0, F2: 168}
                                            _ = v100
                                            _ = v99
                                        }
                                    }
                                    _ = len(v89)
                                } else {
                                    if len(v89) >= 2 {
                                        v89 = v89[1:len(v89)]
                                    }
                                    v19 = ((v19 + v19) * 0.1)
                                }
                            default:
                                var v101 float64 = -2.25
                                _ = v101
                            }
                            switch (^(fi93)) {
                            case -276:
                                var v102 string = "\n\t"
                                _ = v102
                            case 708:
                                _ = len(v89)
                            case 42:
                                v3 = true
                            default:
                                if false {
                                    if true {
                                        fmt.Println(v1, v94, fi83)
                                        v94 = (v94 << ((2) & 63))
                                        var v103 int = genFn1(0)
                                        _ = v103
                                    } else {
                                        if v1 {
                                            v94 += (v94 - v94)
                                            v19 = (3.14159 + 1.5)
                                        } else {
                                            v104 := GenS7{F0: 42, F1: fi93, F2: 956}
                                            var v105 int = fi83
                                            _ = v104
                                            _ = v105
                                        }
                                        v89 = append(v89, genFn1(999))
                                        var v106 uint8 = 157
                                        _ = v106
                                    }
                                    v3 = v3
                                    v107 := GenS3{F0: 42, F1: 1, F2: 2}
                                    _ = v107
                                }
                            }
                            _ = fi93
                            _ = v94
                        }
                    }
                }
            }
            for fi108 := 0; fi108 < 6; fi108++ {
                v89 = append(v89, fi108)
                v3 = v1
                _ = fi108
            }
            _ = v89
        } else {
            var v109 uint32 = v2
            for fi110 := 0; fi110 < 7; fi110++ {
                v111 := GenS8{F0: 127}
                _ = fi110
                _ = v111
            }
            if (2147483647 == -1) {
                var v112 int = (fi83 >> ((18446744073709551615) & 63))
                switch (fi83 & -2147483648) {
                case -971:
                    switch ("foo bar" + "") {
                    case "hello":
                        v113 := []int{v112, (v112 / ((v112) | 1)), (v112 - genFn1(-929))}
                        _ = v113
                    case "a":
                        v114 := GenS9{F0: v3}
                        _ = v114
                    case "foo bar":
                        v112 = genFn1(153)
                    default:
                        v3 = false
                    }
                case 32767:
                    var v115 int32 = 186
                    _ = v115
                default:
                    if v3 {
                        var v116 bool = v3
                        _ = v116
                    } else {
                        for fi117 := 0; fi117 < 12; fi117++ {
                            if v1 {
                                v118 := GenS3{F0: 0, F1: 255, F2: 91}
                                var v119 float32 = ((((v19 - 0.0) + float32((1e10 - (float64(v19) * (3.14159 * -2.25))))) - v19) / v19)
                                _ = v118
                                _ = v119
                            } else {
                                v19 += v19
                                switch (fi83 - genFn1(42)) {
                                case 504:
                                    if v3 {
                                        v2 -= v2
                                    }
                                case 4294967295:
                                    v120 := []int{genFn1(-1)}
                                    _ = v120
                                default:
                                    var v121 uint64 = 413
                                    _ = v121
                                }
                                if (1 > 843) {
                                    v19 = float32(1.5)
                                }
                            }
                            _ = fi117
                        }
                        fmt.Println(v1, fi83, v3)
                    }
                }
                var v122 int32 = 2147483647
                _ = v112
                _ = v122
            } else {
                var v123 float32 = v19
                switch int(v109) {
                case 127:
                    fmt.Println(v2, v109, v19)
                case -720:
                    for fi124 := 0; fi124 < 15; fi124++ {
                        fmt.Println(v123)
                        _ = fi124
                    }
                default:
                    var v125 int = (fi83 >> ((42) & 63))
                    _ = v125
                }
                var v126 uint64 = 18446744073709551615
                _ = v123
                _ = v126
            }
            _ = v109
        }
        _ = fi83
    }
    if ((127 > 42) || ((v2 * v2) > v2)) {
        v127 := []int{32768, 362, 127, 128}
        _ = v127
    } else {
        fmt.Println(v3)
        v128 := []int{517, 42, genFn1(42)}
        _ = v128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
