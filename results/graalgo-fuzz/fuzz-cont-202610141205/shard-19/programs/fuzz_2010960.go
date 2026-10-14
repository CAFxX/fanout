package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint16; F1 int16; F2 int64 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 uint; F1 int32 }
func genFn0(p0 bool, p1 uint16, p2 int64) int16 {
    if false {
        p2 += -356
        v6 := GenS1{F0: (p0 || false), F1: p0}
        _ = v6
    } else {
        var v7 float32 = 0.1
        _ = v7
    }
    fmt.Println(p2, p1)
    for fi8 := 0; fi8 < 10; fi8++ {
        switch "hello" {
        case "":
            v9 := []int{fi8, 32768, fi8}
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 1000
}
func genFn1(p0 int8) int8 {
    var v10 uint64 = 1
    v11 := GenS1{F0: false, F1: (((false || true) && false) && false)}
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := []int{1, 140, (fi13 * fi12)}
            _ = fi13
            _ = v14
        }
        switch ("hello" + ("a" + (("" + ("\n\t" + ("\n\t" + "foo bar"))) + ""))) {
        case "":
            v15 := []int{fi12, (-(fi12)), 128}
            _ = v15
        case "foo bar":
            var v16 bool = (!false)
            _ = v16
        default:
            var v17 uint8 = 0
            _ = v17
        }
        _ = fi12
    }
    var v18 string = ""
    v19 := []int{0, -2147483648, 466, 42}
    _ = p0
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    return 100
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -368
    var v3 int32 = 495
    var v4 uint8 = 243
    var v5 int32 = (-(v3))
    v20 := []int{-5, -279, 85}
    v2 += (v2 * genFn0((0.1 > float32(0.0)), 249, 1))
    switch 1 {
    case 9223372036854775807:
        v5 += v3
    case 0:
        if true {
            for fi21 := 0; fi21 < 9; fi21++ {
                if (0 <= 4294967295) {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    v5 = v3
                    v4 = v4
                }
                _ = fi21
            }
            if false {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v23 := GenS1{F0: false, F1: (true && (!(-9223372036854775808 <= 42)))}
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v4 += v4
                        _ = fi24
                    }
                    _ = fi22
                    _ = v23
                }
            }
            var v25 float64 = ((1.5 - 0.1) + (0.1 + 0.0))
            _ = v25
        }
    case 491:
        var v26 float32 = (3.14159 + ((float32(0.0) - 3.14159) + ((3.14159 * float32(0.1)) - 1.5)))
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 uint32 = 4294967295
        var v29 float32 = 0.1
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        if (42 <= 18446744073709551615) {
            v30 := GenS1{F0: false, F1: (-485 >= -9223372036854775808)}
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 3; fi31++ {
            if true {
                switch (fi31 + fi31) {
                case 593:
                    for fi32 := 0; fi32 < 9; fi32++ {
                        v3 = (v3 - v3)
                        if true {
                            switch ("\n\t" + "hello") {
                            case "":
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    var v33 float64 = (1.5 - -2.25)
                                    _ = v33
                                case "hello":
                                    var v34 int32 = (v3 / ((v5) | 1))
                                    _ = v34
                                }
                            case "hello":
                                switch ("" + "hello") {
                                case "a":
                                    var v35 bool = (296 == 1)
                                    _ = v35
                                default:
                                    var v36 int32 = 1
                                    _ = v36
                                }
                            }
                            var v37 int16 = v2
                            _ = v37
                        } else {
                            fmt.Println(v3)
                            if (!(!(v3 <= v5))) {
                                if ((1 <= 18446744073709551615) || true) {
                                    v5 = v5
                                }
                                if false {
                                    var v38 uint16 = 32767
                                    _ = v38
                                }
                                switch ((("a" + ("hello" + "\n\t")) + "\n\t") + "") {
                                case "":
                                    v3 = v3
                                case "foo bar":
                                    v3 += 0
                                case "\n\t":
                                    fmt.Println(v4, fi32, v4)
                                }
                                if (true || (9 == 270)) {
                                    fmt.Println(v2, fi32, v4)
                                    if ((-89 >= 127) && false) {
                                        if false {
                                            v3 += v5
                                            var v39 int32 = v5
                                            var v40 bool = (387 < uint64(v4))
                                            if false {
                                                var v41 uint = 42
                                                v39 -= (v39 / (((v39 / ((v5) | 1))) | 1))
                                                if v40 {
                                                    var v42 float32 = 1.5
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v43 := GenS1{F0: (2 != 150), F1: false}
                                                        _ = v43
                                                    }
                                                    _ = v42
                                                }
                                                _ = v41
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                        switch "" {
                                        case "foo bar":
                                            v44 := v1["a"]
                                            _ = v44
                                        }
                                    } else {
                                        v2 = v2
                                    }
                                    switch "" {
                                    case "":
                                        v5 -= (v5 % ((1) | 1))
                                    default:
                                        v5 = (-(v3))
                                    }
                                }
                                if true {
                                    v3 += (^(v3))
                                    v45 := GenS1{F0: true, F1: (334 >= 2147483648)}
                                    if true {
                                        v2 = v2
                                    }
                                    v5 -= v3
                                    _ = v45
                                } else {
                                    fmt.Println(v4, v4)
                                }
                            } else {
                                var v46 float64 = 0.1
                                _ = v46
                            }
                        }
                        _ = fi32
                    }
                case -2147483648:
                    v2 = (^(v2))
                }
                if (false || false) {
                    if (fi31 == 32767) {
                        switch (fi31 >> ((18446744073709551615) & 63)) {
                        case 1:
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        case -2147483648:
                            var v47 uint64 = 18446744073709551615
                            _ = v47
                        }
                    }
                } else {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    v48 := GenS1{F0: (int32(v4) != v5), F1: false}
                    v49 := GenS2{F0: 42, F1: (false || (!((v2 / ((42) | 1)) == 602)))}
                    _ = v48
                    _ = v49
                }
            } else {
                if true {
                    var v50 int32 = (v5 % ((v3) | 1))
                    for fi51 := 0; fi51 < 5; fi51++ {
                        v3 += (-(v3))
                        var v52 string = ""
                        v2 -= (v2 / (((-(v2))) | 1))
                        _ = fi51
                        _ = v52
                    }
                    var v53 int = (fi31 & int(v50))
                    _ = v50
                    _ = v53
                }
                switch ("" + "") {
                case "foo bar":
                    var v54 float32 = 0.0
                    _ = v54
                case "hello":
                    switch ("" + "") {
                    case "\n\t":
                        _ = len(v1)
                    default:
                        for fi55 := 0; fi55 < 6; fi55++ {
                            v56 := GenS2{F0: v4, F1: (2 < 4294967295)}
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            v20 = append(v20, (fi55 % (((-(fi55))) | 1)))
                            _ = fi55
                            _ = v56
                        }
                    }
                default:
                    _ = len(v20)
                }
            }
            fmt.Println(v3)
            _ = fi31
        }
    }
    var v57 uint16 = 975
    if false {
        if true {
            if false {
                switch 4294967295 {
                case 4294967295:
                    var v58 float32 = 0.1
                    _ = v58
                case 32768:
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            v59 := GenS2{F0: 127, F1: (!((2 < 459) && (127 <= -2147483648)))}
                            _ = len(v20)
                            fmt.Println(v57, v57)
                            _ = v59
                        }
                    case "a":
                        switch (("" + ("" + "hello")) + "hello") {
                        case "":
                            var v60 int8 = -22
                            _ = v60
                        default:
                            delete(v1, ("\n\t" + ""))
                        }
                    default:
                        v57 = v57
                    }
                default:
                    switch "hello" {
                    case "\n\t":
                        if false {
                            for fi61 := 0; fi61 < 17; fi61++ {
                                for fi62 := 0; fi62 < 19; fi62++ {
                                    v63 := GenS2{F0: v4, F1: true}
                                    if true {
                                        v57 = v57
                                        if len(v20) > 0 {
                                            v20[0] = (fi62 / ((fi61) | 1))
                                        }
                                        fmt.Println(v3, fi62, v57)
                                        if true {
                                            var v64 uint = 360
                                            v57 = (v57 * v57)
                                            v57 -= 255
                                            _ = v64
                                        } else {
                                            fmt.Println(fi61, v57)
                                            v4 += v4
                                            v1["foo bar"] = true
                                        }
                                        var v65 int64 = 42
                                        _ = v65
                                    } else {
                                        fmt.Println(v2, v2)
                                        if true {
                                            v66 := GenS3{F0: 2147483647}
                                            if false {
                                                if ("a" > ("\n\t" + "hello")) {
                                                    if true {
                                                        v67 := GenS2{F0: v4, F1: (false || (true || false))}
                                                        _ = v67
                                                    } else {
                                                        v2 += (^(v2))
                                                    }
                                                    v57 = v57
                                                }
                                                switch -9223372036854775808 {
                                                case 0:
                                                    var v68 bool = (!false)
                                                    _ = v68
                                                default:
                                                    switch -526 {
                                                    case -397:
                                                        var v69 uint16 = v57
                                                        _ = v69
                                                    case -1:
                                                        var v70 uint16 = (v57 ^ v57)
                                                        _ = v70
                                                    default:
                                                        v71 := GenS1{F0: (1e10 < (((1e10 + ((((-2.25 * 1e10) - (0.1 - 0.0)) * float64(0.1)) - -2.25)) - (0.0 + float64(1e10))) / (((float64((0.1 + 1e10)) / float64(-2.25)) + 0.1) / -2.25))), F1: (!true)}
                                                        _ = v71
                                                    }
                                                }
                                            } else {
                                                var v72 bool = false
                                                var v73 int8 = 100
                                                _ = v72
                                                _ = v73
                                            }
                                            var v74 int32 = v3
                                            var v75 int16 = -802
                                            v76 := GenS1{F0: (false || (558 < 1)), F1: (1000 <= v2)}
                                            _ = v66
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                    }
                                    v77 := v1["hello"]
                                    _ = fi62
                                    _ = v63
                                    _ = v77
                                }
                                if (0.0 <= (-2.25 + ((0.0 + (3.14159 * 3.14159)) + (-2.25 / 0.1)))) {
                                    v78 := GenS1{F0: (!false), F1: (1e10 >= float64(0.0))}
                                    v4 = 103
                                    v79 := GenS2{F0: 225, F1: ((fi61 * (fi61 % ((fi61) | 1))) <= (fi61 | fi61))}
                                    _ = v78
                                    _ = v79
                                } else {
                                    var v80 uint8 = 0
                                    v2 += genFn0((v57 == v57), (v57 << ((uint(v4)) & 63)), 1)
                                    if (true || true) {
                                        var v81 uint = 2
                                        _ = v81
                                    } else {
                                        if ((false && (fi61 >= (fi61 << ((18446744073709551615) & 63)))) && true) {
                                            v57 = v57
                                            var v82 float64 = (0.0 / 1.5)
                                            _ = v82
                                        }
                                        fmt.Println(v5)
                                        switch fi61 {
                                        case -584:
                                            for fi83 := 0; fi83 < 4; fi83++ {
                                                v2 += v2
                                                fmt.Println(v4)
                                                _ = fi83
                                            }
                                        }
                                    }
                                    _ = v80
                                }
                                _ = fi61
                            }
                            v84 := GenS4{F0: v57, F1: 379, F2: 0}
                            v5 -= -1
                            _ = v84
                        } else {
                            v5 = v5
                            var v85 int = 127
                            var v86 float32 = 3.14159
                            _ = v85
                            _ = v86
                        }
                    case "hello":
                        var v87 int32 = v3
                        _ = v87
                    case "a":
                        v2 = (v2 & v2)
                    }
                }
                switch 286 {
                case 32768:
                    v57 += 2
                case 46:
                    if false {
                        v5 -= v3
                        if (1.5 == 3.14159) {
                            v88 := GenS5{F0: (^(v3))}
                            v89 := GenS5{F0: v3}
                            v90 := GenS6{F0: (v2 | (v2 << ((152) & 63))), F1: 65535}
                            _ = v88
                            _ = v89
                            _ = v90
                        }
                        var v91 int = 4294967295
                        _ = v91
                    } else {
                        if ((0 <= 0) || false) {
                            _ = len(v1)
                        } else {
                            var v92 float64 = 1.5
                            _ = v92
                        }
                        v93 := GenS3{F0: 19}
                        _ = v93
                    }
                case -350:
                    _ = len(v1)
                }
                if (!true) {
                    v94 := GenS5{F0: 2147483647}
                    _ = v94
                } else {
                    switch (("" + "hello") + "a") {
                    case "a":
                        v95 := GenS7{F0: v57, F1: genFn0(true, 981, 1), F2: -1}
                        _ = v95
                    case "\n\t":
                        var v96 bool = false
                        _ = v96
                    case "":
                        for fi97 := 0; fi97 < 20; fi97++ {
                            v2 = (v2 / (((v2 << ((4294967295) & 63))) | 1))
                            v98 := GenS8{F0: (v57 >> ((0) & 63)), F1: 1}
                            v99 := v1["foo bar"]
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                    }
                    var v100 uint16 = (v57 ^ (^(v57)))
                    v2 = 32767
                    _ = v100
                }
            } else {
                v4 = (v4 << ((255) & 63))
                v5 = (v3 / (((v3 % ((0) | 1))) | 1))
            }
            var v101 int32 = (v3 * int32(v2))
            for fi102 := 0; fi102 < 14; fi102++ {
                if (false || (!false)) {
                    fmt.Println(v101, v101)
                    for fi103 := 0; fi103 < 10; fi103++ {
                        fmt.Println(v3)
                        v104 := GenS5{F0: v101}
                        fmt.Println(v57)
                        if (-84 >= 217) {
                            if (true || (true && false)) {
                                v20 = append(v20, 525)
                                v5 = v101
                            } else {
                                var v105 int32 = 0
                                var v106 int32 = (v101 - (v5 | (v105 ^ (v105 >> ((255) & 63)))))
                                _ = v105
                                _ = v106
                            }
                            var v107 float32 = 1.5
                            var v108 string = ("hello" + ("hello" + ("\n\t" + ("\n\t" + ""))))
                            v109 := GenS4{F0: 1, F1: 1, F2: 1}
                            v110 := v1[v108]
                            _ = v107
                            _ = v108
                            _ = v109
                            _ = v110
                        }
                        _ = fi103
                        _ = v104
                    }
                }
                _ = fi102
            }
            _ = v101
        }
    } else {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        fmt.Println(v4, v3, v2)
        if true {
            switch -2147483648 {
            case 32768:
                for fi111 := 0; fi111 < 8; fi111++ {
                    v112 := GenS1{F0: true, F1: false}
                    v113 := GenS9{F0: v2, F1: 1}
                    _ = fi111
                    _ = v112
                    _ = v113
                }
            case 127:
                var v114 int8 = 42
                _ = v114
            case 42:
                v115 := GenS10{F0: 721, F1: (v3 << ((1) & 63))}
                _ = v115
            }
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v5)
            case "":
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            case "a":
                delete(v1, (("" + "hello") + "\n\t"))
            default:
                switch "a" {
                case "a":
                    var v116 float64 = -2.25
                    _ = v116
                default:
                    switch -9223372036854775808 {
                    case 995:
                        if true {
                            var v117 uint = 18446744073709551615
                            _ = v117
                        } else {
                            for fi118 := 0; fi118 < 11; fi118++ {
                                var v119 int16 = (^(v2))
                                _ = len(v1)
                                for fi120 := 0; fi120 < 7; fi120++ {
                                    if ("\n\t" == "") {
                                        var v121 uint = 1
                                        _ = len(v20)
                                        _ = v121
                                    } else {
                                        switch (fi118 & fi120) {
                                        case 42:
                                            v3 = (^(v5))
                                        }
                                        var v122 uint16 = (v57 * 32767)
                                        var v123 string = "\n\t"
                                        delete(v1, v123)
                                        _ = v122
                                        _ = v123
                                    }
                                    _ = fi120
                                }
                                _ = fi118
                                _ = v119
                            }
                            var v124 float64 = 1.5
                            v5 = (v3 & 2147483647)
                            _ = v124
                        }
                    case 32768:
                        var v125 int = 2147483647
                        _ = v125
                    default:
                        var v126 uint = 679
                        _ = v126
                    }
                }
            }
        }
    }
    for fi127 := 0; fi127 < 8; fi127++ {
        v128 := GenS10{F0: 2, F1: v5}
        v5 = 0
        _ = fi127
        _ = v128
    }
    var v129 uint64 = 0
    if (0 < -939) {
        v57 = v57
        switch 4294967295 {
        case 1:
            var v130 float32 = -2.25
            _ = v130
        default:
            for fi131 := 0; fi131 < 6; fi131++ {
                var v132 string = ("" + "hello")
                v133 := GenS8{F0: (v57 ^ v57), F1: 41}
                _ = fi131
                _ = v132
                _ = v133
            }
        }
    }
    v57 = v57
    v20 = append(v20, 127)
    v3 += -2147483648
    fmt.Println(v5, v3)
    if true {
        var v134 bool = (v129 != (v129 >> ((0) & 63)))
        var v135 float32 = (1e10 + ((0.1 / float32((float64(1e10) / 1.5))) + 0.0))
        _ = v134
        _ = v135
    }
    v4 += (v4 ^ v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v129
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
