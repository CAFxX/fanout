package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 uint16; F2 string }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 int }
func genFn0(p0 uint, p1 uint64, p2 int16) int64 {
    v4 := GenS1{F0: "hello"}
    fmt.Println(p2, p0)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 int64 = 1
    var v2 float64 = 0.0
    v3 := make(map[string]string)
    var v5 bool = (0 != 520)
    for fi6 := 0; fi6 < 17; fi6++ {
        switch fi6 {
        case 534:
            v1 = (^(v1))
        case 42:
            var v7 int32 = 42
            _ = v7
        default:
            fmt.Println(fi6)
        }
        _ = fi6
    }
    delete(v3, "hello")
    switch "a" {
    case "":
        switch 2147483647 {
        case 0:
            for fi8 := 0; fi8 < 4; fi8++ {
                switch fi8 {
                case -348:
                    if (1 <= 2147483647) {
                        fmt.Println(v5, v5, fi8)
                        var v9 int32 = -2147483648
                        v10 := GenS1{F0: ""}
                        _ = v10
                        _ = v9
                    }
                }
                var v11 uint64 = 0
                v12 := GenS2{F0: 32767, F1: 42, F2: 2147483647}
                _ = fi8
                _ = v11
                _ = v12
            }
        default:
            var v13 float32 = 0.0
            _ = v13
        }
    case "hello":
        v1 = (v1 % ((0) | 1))
    case "a":
        if false {
            v2 = v2
            v2 = (v2 - float64(0.1))
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := GenS3{F0: 127}
                _ = fi14
                _ = v15
            }
            v1 = 106
            fmt.Println(v5, v1, v1)
        }
    }
    var v16 float64 = v2
    for fi17 := 0; fi17 < 3; fi17++ {
        if v5 {
            var v18 int64 = v1
            v1 = v1
            switch (("hello" + "a") + (("hello" + ("" + "\n\t")) + "foo bar")) {
            case "hello":
                v5 = v5
            case "a":
                v2 = (v2 * float64(0.1))
            case "\n\t":
                delete(v3, ("\n\t" + "a"))
            default:
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 int = (fi17 & (-(fi19)))
                    v21 := GenS1{F0: "\n\t"}
                    v22 := GenS1{F0: "a"}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            _ = v18
        } else {
            v2 = (v2 * v16)
        }
        _ = fi17
    }
    delete(v3, "\n\t")
    var v23 uint = 975
    switch -9223372036854775808 {
    case 0:
        for fi24 := 0; fi24 < 6; fi24++ {
            if v5 {
                fmt.Println(v1, fi24, v23)
                if (127 > int8(v23)) {
                    v5 = (!true)
                } else {
                    var v25 float32 = (-2.25 - -2.25)
                    _ = v25
                }
                v26 := GenS4{F0: 1, F1: ((-2147483648 != 0) && true)}
                _ = v26
            } else {
                var v27 uint = v23
                if (811 == 0) {
                    switch (^(fi24)) {
                    case 32767:
                        if (135 <= 0) {
                            v5 = (42 > v1)
                            if (1.5 < v2) {
                                v28 := []int{}
                                var v29 float64 = 0.0
                                _ = v28
                                _ = v29
                            } else {
                                v30 := []int{(fi24 - (^(fi24))), (fi24 << ((0) & 63))}
                                _ = v30
                            }
                        }
                    case -1:
                        v31 := v3[((("a" + "a") + (("\n\t" + "") + "\n\t")) + "\n\t")]
                        _ = v31
                    default:
                        var v32 float32 = 0.1
                        _ = v32
                    }
                    v5 = (-24 >= -45)
                }
                _ = v27
            }
            switch "foo bar" {
            case "":
                v33 := GenS2{F0: 609, F1: 255, F2: -1}
                _ = v33
            case "\n\t":
                v5 = false
            case "a":
                fmt.Println(fi24, v5)
            }
            _ = fi24
        }
    case 473:
        for fi34 := 0; fi34 < 8; fi34++ {
            v16 = 1.5
            _ = fi34
        }
    }
    switch "\n\t" {
    case "\n\t":
        if v5 {
            var v35 uint = v23
            _ = v35
        } else {
            var v36 int8 = 127
            if (!v5) {
                switch -2147483648 {
                case 0:
                    if v5 {
                        var v37 int16 = 32767
                        fmt.Println(v1, v2, v16)
                        _ = v37
                    } else {
                        v36 = (v36 | (v36 + v36))
                        var v38 int64 = v1
                        _ = v38
                    }
                }
                switch "\n\t" {
                case "":
                    v5 = (!v5)
                case "hello":
                    var v39 uint16 = 1
                    _ = v39
                default:
                    switch 2147483647 {
                    case 491:
                        v16 += 0.1
                    default:
                        for fi40 := 0; fi40 < 15; fi40++ {
                            switch "hello" {
                            case "\n\t":
                                switch "hello" {
                                case "foo bar":
                                    v41 := []int{(-(fi40)), fi40, fi40}
                                    _ = v41
                                case "":
                                    v23 = 0
                                case "\n\t":
                                    for fi42 := 0; fi42 < 9; fi42++ {
                                        v43 := GenS4{F0: 113, F1: v5}
                                        if (("a" + "\n\t") > "hello") {
                                            if (2 > 0) {
                                                v36 = 100
                                                var v44 float32 = 3.14159
                                                var v45 uint16 = 42
                                                _ = v44
                                                _ = v45
                                            } else {
                                                switch (fi42 ^ fi40) {
                                                case 127:
                                                    v23 = v23
                                                case 32768:
                                                    fmt.Println(v16)
                                                default:
                                                    v46 := []int{128, fi42, (fi42 & 127), fi42}
                                                    _ = v46
                                                }
                                            }
                                        } else {
                                            var v47 int16 = 32767
                                            var v48 string = ("a" + "\n\t")
                                            _ = v47
                                            _ = v48
                                        }
                                        if v5 {
                                            if v5 {
                                                fmt.Println(fi42)
                                                v49 := GenS1{F0: "foo bar"}
                                                fmt.Println(v36, v1, v5)
                                                _ = v49
                                            } else {
                                                var v50 uint = v23
                                                fmt.Println(v50)
                                                _ = v50
                                            }
                                            v51 := GenS1{F0: ("a" + ("a" + ("foo bar" + "")))}
                                            v5 = (!v5)
                                            v36 = (v36 << ((v23) & 63))
                                            _ = v51
                                        }
                                        _ = fi42
                                        _ = v43
                                    }
                                default:
                                    var v52 uint64 = uint64(v23)
                                    _ = v52
                                }
                            case "a":
                                v53 := GenS4{F0: (v36 << ((v23) & 63)), F1: (0.0 != -2.25)}
                                _ = v53
                            case "foo bar":
                                if true {
                                    switch fi40 {
                                    case 127:
                                        switch (fi40 - int(v23)) {
                                        case 493:
                                            fmt.Println(v16, v16, v36)
                                        case 763:
                                            for fi54 := 0; fi54 < 15; fi54++ {
                                                v55 := []int{-2147483648, (fi54 & (fi54 * (fi54 ^ fi40))), fi54}
                                                if v5 {
                                                    if (v5 && (783 > 841)) {
                                                        switch (fi54 * fi54) {
                                                        case 9223372036854775807:
                                                            switch fi54 {
                                                            case -406:
                                                                if v5 {
                                                                    v23 = (v23 ^ (v23 << (((v23 >> ((4294967295) & 63))) & 63)))
                                                                    fmt.Println(v2)
                                                                    if v5 {
                                                                        if v5 {
                                                                            var v56 float32 = float32(v16)
                                                                            fmt.Println(v16)
                                                                            fmt.Println(fi40, v23, fi40)
                                                                            _ = v56
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case 127:
                                                            v57 := GenS4{F0: -97, F1: (!v5)}
                                                            _ = v57
                                                        }
                                                        v1 = 0
                                                        if v5 {
                                                            delete(v3, ("a" + "a"))
                                                        } else {
                                                            v58 := GenS5{F0: 47, F1: 2147483647}
                                                            _ = v58
                                                        }
                                                    }
                                                    v36 -= v36
                                                    if v5 {
                                                        var v59 int = (^(fi40))
                                                        var v60 uint8 = 255
                                                        _ = len(v3)
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                } else {
                                                    var v61 int = fi54
                                                    if v5 {
                                                        var v62 float32 = 0.1
                                                        if ((v36 - (v36 / ((v36) | 1))) >= (^(v36))) {
                                                            var v63 float32 = (v62 / float32(v16))
                                                            var v64 int = 0
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            if (65535 < 0) {
                                                                v65 := GenS6{F0: 0}
                                                                _ = len(v55)
                                                                switch "\n\t" {
                                                                case "":
                                                                    v66 := GenS3{F0: (-(v36))}
                                                                    _ = v66
                                                                case "foo bar":
                                                                    if (1e10 > v2) {
                                                                        if len(v55) >= 2 {
                                                                            v55 = v55[1:len(v55)]
                                                                        }
                                                                    } else {
                                                                        var v67 uint32 = 42
                                                                        var v68 float32 = (1e10 / (-2.25 / 1e10))
                                                                        var v69 int = fi40
                                                                        _ = v67
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                case "\n\t":
                                                                    v3[""] = "a"
                                                                }
                                                                _ = v65
                                                            }
                                                            if (float64(3.14159) <= v16) {
                                                                if v5 {
                                                                    if true {
                                                                        v3["\n\t"] = ("" + ("\n\t" + ("" + "\n\t")))
                                                                    }
                                                                    var v70 uint8 = 12
                                                                    if len(v55) > 0 {
                                                                        v55[0] = (fi40 * int(v36))
                                                                    }
                                                                    _ = v70
                                                                } else {
                                                                    v16 = (1e10 - float64(1e10))
                                                                    var v71 float32 = 0.0
                                                                    _ = v71
                                                                }
                                                            } else {
                                                                if (v16 != float64(0.0)) {
                                                                    fmt.Println(v36, v23, v16)
                                                                    switch v61 {
                                                                    case 128:
                                                                        v72 := GenS5{F0: 255, F1: -1}
                                                                        _ = v72
                                                                    case 737:
                                                                        v3["hello"] = "a"
                                                                    case 681:
                                                                        switch "hello" {
                                                                        case "a":
                                                                            var v73 int16 = 908
                                                                            _ = v73
                                                                        case "\n\t":
                                                                            var v74 uint32 = 558
                                                                            _ = v74
                                                                        default:
                                                                            if len(v55) >= 2 {
                                                                                v55 = v55[1:len(v55)]
                                                                            }
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v75 uint64 = 191
                                                                        _ = v75
                                                                    }
                                                                }
                                                            }
                                                            var v76 uint = 42
                                                            _ = v76
                                                        }
                                                        if ((true && (65535 != 65535)) || ((v36 * (v36 >> (((v23 & v23)) & 63))) != (v36 ^ (v36 % (((v36 >> (((v23 + v23)) & 63))) | 1))))) {
                                                            v16 = -2.25
                                                            var v77 uint32 = 65535
                                                            v78 := GenS3{F0: 127}
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        _ = v62
                                                    }
                                                    _ = v61
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    var v79 int32 = -1
                                                    _ = v79
                                                }
                                                _ = fi54
                                                _ = v55
                                            }
                                        case 0:
                                            for fi80 := 0; fi80 < 4; fi80++ {
                                                v5 = (!((((0.0 - ((3.14159 + 3.14159) + (1.5 / (0.1 * 1.5)))) + -2.25) + 0.0) >= float32((float64(1.5) / 1e10))))
                                                var v81 int16 = -1
                                                v3["a"] = ""
                                                switch fi80 {
                                                case 0:
                                                    v5 = v5
                                                    switch "hello" {
                                                    case "a":
                                                        var v82 int64 = (v1 << ((v23) & 63))
                                                        _ = v82
                                                    default:
                                                        fmt.Println(fi80, fi40, v1)
                                                    }
                                                default:
                                                    v81 = -454
                                                    if (!(!v5)) {
                                                        var v83 int = -19
                                                        if (v5 && (42 <= 944)) {
                                                            var v84 int = v83
                                                            switch (fi80 + 32767) {
                                                            case -9223372036854775808:
                                                                var v85 int64 = (v1 + 1)
                                                                _ = v85
                                                            case -1:
                                                                switch (v83 ^ fi40) {
                                                                case 2147483647:
                                                                    var v86 uint32 = 2147483647
                                                                    _ = v86
                                                                case -9223372036854775808:
                                                                    delete(v3, ("a" + ""))
                                                                case 986:
                                                                    fmt.Println(v1, v5)
                                                                }
                                                                v84 = fi40
                                                                fmt.Println(v2, v81)
                                                            case -972:
                                                                var v87 uint8 = 223
                                                                _ = v87
                                                            default:
                                                                switch (^(fi40)) {
                                                                case 127:
                                                                    var v88 int64 = v1
                                                                    _ = v88
                                                                case -105:
                                                                    var v89 int = fi80
                                                                    _ = v89
                                                                case 4294967295:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v90 string = "hello"
                                                                        _ = v90
                                                                    case "hello":
                                                                        v91 := GenS7{F0: 0}
                                                                        _ = v91
                                                                    }
                                                                default:
                                                                    if false {
                                                                        if (!(42 >= 946)) {
                                                                            v5 = true
                                                                        } else {
                                                                            fmt.Println(v16)
                                                                            var v92 float64 = (float64(0.1) / 1.0)
                                                                            v93 := []int{fi80, 582}
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        var v94 uint8 = 128
                                                                        if true {
                                                                            if false {
                                                                                if ((v94 ^ 42) == 0) {
                                                                                    v95 := GenS3{F0: (v36 & (v36 ^ v36))}
                                                                                    v96 := GenS3{F0: (v36 & -1)}
                                                                                    var v97 uint64 = 18446744073709551615
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                    _ = v97
                                                                                } else {
                                                                                    var v98 bool = (genFn0(1, 4294967295, 0) == (v1 % ((genFn0((v23 + v23), 962, 1)) | 1)))
                                                                                    var v99 uint = (v23 % ((4294967295) | 1))
                                                                                    var v100 int = -315
                                                                                    _ = v100
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                                v101 := GenS5{F0: (v94 / ((42) | 1)), F1: 796}
                                                                                v102 := GenS8{F0: (true && true), F1: 42, F2: "\n\t"}
                                                                                _ = v101
                                                                                _ = v102
                                                                            } else {
                                                                                var v103 int8 = v36
                                                                                v23 = (v23 >> (((^(v23))) & 63))
                                                                                var v104 float32 = 0.1
                                                                                _ = v103
                                                                                _ = v104
                                                                            }
                                                                        }
                                                                        _ = v94
                                                                    } else {
                                                                        var v105 string = "foo bar"
                                                                        var v106 float32 = -2.25
                                                                        _ = len(v3)
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                }
                                                            }
                                                            _ = v84
                                                        }
                                                        fmt.Println(v83)
                                                        _ = v83
                                                    } else {
                                                        v107 := v3["a"]
                                                        v2 += 1e10
                                                        if v5 {
                                                            v5 = ((932 < 42) || v5)
                                                            v108 := []int{(fi40 >> ((v23) & 63)), fi40, fi40}
                                                            _ = v108
                                                        }
                                                        _ = v107
                                                    }
                                                }
                                                _ = fi80
                                                _ = v81
                                            }
                                        }
                                    }
                                } else {
                                    var v109 uint64 = 1
                                    fmt.Println(v1, fi40, v16)
                                    _ = v109
                                }
                            default:
                                var v110 uint = v23
                                _ = v110
                            }
                            var v111 float64 = v16
                            _ = fi40
                            _ = v111
                        }
                    }
                }
                delete(v3, ("\n\t" + (("hello" + (("\n\t" + "") + ("a" + "foo bar"))) + (("hello" + "\n\t") + "a"))))
            }
            v112 := GenS8{F0: false, F1: 42, F2: "hello"}
            _ = v112
            _ = v36
        }
    case "":
        var v113 bool = ("" == (("\n\t" + ("\n\t" + ("hello" + ("" + "foo bar")))) + "\n\t"))
        _ = v113
    case "a":
        for fi114 := 0; fi114 < 4; fi114++ {
            switch fi114 {
            case 535:
                var v115 uint16 = 711
                _ = v115
            case 4294967295:
                fmt.Println(v2, v5, v1)
            }
            v116 := GenS7{F0: 146}
            _ = fi114
            _ = v116
        }
    }
    switch 128 {
    case 451:
        for fi117 := 0; fi117 < 5; fi117++ {
            v1 = v1
            _ = fi117
        }
    case 32767:
        if v5 {
            v118 := []int{-135}
            v23 -= 838
            _ = v118
        }
    }
    switch "" {
    case "hello":
        delete(v3, ("" + "a"))
    }
    v119 := []int{0, 32767, 321}
    _ = len(v119)
    if (true && (!(-1 >= 1))) {
        for fi120 := 0; fi120 < 4; fi120++ {
            if v5 {
                fmt.Println(v2, v23, fi120)
                if len(v119) > 0 {
                    v119[0] = 175
                }
                v121 := GenS9{F0: (((("hello" + "") + "a") + "foo bar") + "hello"), F1: 2147483647}
                _ = v121
            }
            if true {
                fmt.Println(v2)
            }
            delete(v3, ("" + "foo bar"))
            _ = fi120
        }
        var v122 int = 127
        _ = v122
    } else {
        if (0.1 < (0.1 - ((v16 * float64(((((-2.25 - 1e10) * ((float32(v2) * -2.25) - -2.25)) * ((float32(-2.25) + 0.1) / 0.1)) * (((-2.25 + 1.5) * 0.1) / 1.0)))) * float64((1.5 * -2.25))))) {
            if false {
                if v5 {
                    var v123 bool = v5
                    v124 := v3[("foo bar" + "a")]
                    for fi125 := 0; fi125 < 19; fi125++ {
                        var v126 int8 = 1
                        _ = fi125
                        _ = v126
                    }
                    _ = v123
                    _ = v124
                } else {
                    v23 += (v23 - v23)
                    switch "hello" {
                    case "\n\t":
                        v16 = 0.1
                    case "a":
                        _ = len(v119)
                    case "foo bar":
                        var v127 int8 = -128
                        _ = v127
                    }
                }
                v128 := GenS10{F0: 127}
                _ = v128
            }
            if v5 {
                v129 := GenS2{F0: 883, F1: (^(v23)), F2: 825}
                fmt.Println(v1, v23)
                _ = v129
            }
        }
        if false {
            v2 += v16
            fmt.Println(v2, v16, v16)
        }
        if len(v119) >= 2 {
            v119 = v119[1:len(v119)]
        }
    }
    var v130 int8 = -9
    for fi131 := 0; fi131 < 5; fi131++ {
        var v132 uint32 = 2
        v16 = 0.0
        v132 = (v132 ^ v132)
        _ = fi131
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v119:", len(v119), v119)
    fmt.Println("v130:", v130)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v119
    _ = v130
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v5
}
