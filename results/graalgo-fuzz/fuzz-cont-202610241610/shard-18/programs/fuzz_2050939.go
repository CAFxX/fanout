package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 int; F2 uint8 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint64; F1 int }

func main() {
    var v1 uint32 = 691
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 int64 = -333
    var v5 bool = (2 > 0)
    v6 := make(map[uint32]uint32)
    var v7 float32 = 3.14159
    if (!v5) {
        switch (("a" + "\n\t") + ("foo bar" + "hello")) {
        case "":
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(fi8, v5)
                if ((v1 / ((v1) | 1)) < (^(v1))) {
                    var v9 uint8 = 255
                    fmt.Println(fi8)
                    v7 += 0.1
                    _ = v9
                } else {
                    var v10 uint = 4294967295
                    v5 = v5
                    _ = v10
                }
                _ = fi8
            }
        case "hello":
            var v11 uint16 = 255
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 16; fi12++ {
                switch "a" {
                case "a":
                    if v5 {
                        v13 := v6[v1]
                        v14 := []int{(fi12 - fi12), (fi12 ^ fi12), -9223372036854775808}
                        fmt.Println(fi12)
                        _ = v13
                        _ = v14
                    } else {
                        var v15 float32 = ((v7 / float32(-2.25)) * v7)
                        if (127 >= 255) {
                            delete(v2, "")
                            var v16 int = 128
                            switch 2147483647 {
                            case 194:
                                fmt.Println(v5, fi12, v15)
                            case -9223372036854775808:
                                v16 = int(v4)
                            }
                            _ = v16
                        } else {
                            var v17 int16 = 1
                            if true {
                                v18 := GenS1{F0: "\n\t", F1: 843, F2: 127}
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    fmt.Println(v7)
                                    _ = fi19
                                }
                                _ = v18
                            } else {
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    v21 := GenS1{F0: ((("hello" + "foo bar") + "hello") + (("\n\t" + "") + ("\n\t" + "foo bar"))), F1: fi20, F2: 120}
                                    v22 := []int{}
                                    if v5 {
                                        v22 = append(v22, 902)
                                        v4 = v4
                                        v23 := GenS2{F0: 1000}
                                        _ = v23
                                    } else {
                                        v1 = (^(v1))
                                    }
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                var v24 uint8 = 0
                                _ = v24
                            }
                            v4 -= (v4 * v4)
                            _ = v17
                        }
                        if (18446744073709551615 != 65535) {
                            if (276 != 0) {
                                v25 := GenS1{F0: "a", F1: fi12, F2: 127}
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    var v27 uint32 = v1
                                    var v28 int32 = 408
                                    v29 := []int{fi12, fi26, (fi26 << ((821) & 63)), (fi26 << ((0) & 63))}
                                    var v30 int = (fi26 ^ 9223372036854775807)
                                    v29 = append(v29, v30)
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            }
                            v7 = (1.5 + v7)
                        } else {
                            fmt.Println(v5, v5)
                        }
                        _ = v15
                    }
                case "":
                    if (true && v5) {
                        v31 := GenS2{F0: 473}
                        _ = v31
                    }
                default:
                    v5 = false
                }
                fmt.Println(v5)
                if (false && v5) {
                    delete(v6, (v1 % ((v1) | 1)))
                    v7 = (1e10 + ((v7 / (float32((1.5 - -2.25)) * (v7 - v7))) * (v7 * (v7 - 0.0))))
                    switch ("" + "a") {
                    case "\n\t":
                        var v32 int16 = -374
                        _ = v32
                    case "foo bar":
                        if (255 <= 128) {
                            if ((0.0 - (float32((1e10 + float64(((1.5 / (1.5 * (((v7 / (v7 - ((0.0 - v7) - v7))) * v7) / 1.5))) / (v7 / 1.0))))) + float32(1e10))) <= 0.0) {
                                var v33 uint16 = 65535
                                v34 := []int{fi12, 0}
                                _ = v33
                                _ = v34
                            } else {
                                var v35 uint16 = 42
                                _ = v35
                            }
                            var v36 uint = 568
                            switch "foo bar" {
                            case "\n\t":
                                if v5 {
                                    fmt.Println(v7, v5, v36)
                                } else {
                                    switch ((("a" + "a") + "\n\t") + "\n\t") {
                                    case "":
                                        var v37 uint8 = 0
                                        _ = v37
                                    default:
                                        if false {
                                            if (!((v5 && v5) || v5)) {
                                                v1 += (v1 % ((v1) | 1))
                                                switch (("\n\t" + ((("hello" + "hello") + "a") + "a")) + ("" + "foo bar")) {
                                                case "foo bar":
                                                    for fi38 := 0; fi38 < 8; fi38++ {
                                                        v39 := v2[""]
                                                        _ = fi38
                                                        _ = v39
                                                    }
                                                case "hello":
                                                    v40 := []int{int(v1), (^(fi12))}
                                                    _ = v40
                                                }
                                                var v41 uint32 = (v1 % ((v1) | 1))
                                                _ = v41
                                            }
                                            var v42 uint8 = 128
                                            v43 := GenS2{F0: 42}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            var v44 uint = v36
                                            v45 := GenS2{F0: -906}
                                            fmt.Println(v44, v7, v36)
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    v7 += float32(0.0)
                                }
                            case "foo bar":
                                if v5 {
                                    var v46 int8 = -57
                                    fmt.Println(v7)
                                    if v5 {
                                        switch "\n\t" {
                                        case "a":
                                            switch (fi12 - (fi12 << (((v36 | v36)) & 63))) {
                                            case 32768:
                                                v47 := GenS1{F0: "\n\t", F1: (fi12 - (^(fi12))), F2: -67}
                                                _ = v47
                                            default:
                                                for fi48 := 0; fi48 < 15; fi48++ {
                                                    if true {
                                                        v36 -= 18446744073709551615
                                                        var v49 int32 = 0
                                                        _ = v49
                                                    }
                                                    var v50 int32 = 1
                                                    v46 = 0
                                                    switch 885 {
                                                    case 32768:
                                                        var v51 int8 = (v46 - (v46 / ((v46) | 1)))
                                                        _ = v51
                                                    case 0:
                                                        v52 := GenS3{F0: 1, F1: 156, F2: 37}
                                                        _ = v52
                                                    case 127:
                                                        if (((float64(v7) / -2.25) + 0.0) >= 0.1) {
                                                            v53 := []int{}
                                                            _ = v53
                                                        }
                                                    default:
                                                        var v54 int16 = -1
                                                        _ = v54
                                                    }
                                                    if false {
                                                        fmt.Println(fi48, fi48)
                                                    } else {
                                                        if true {
                                                            var v55 int8 = v46
                                                            var v56 uint8 = 25
                                                            switch ("foo bar" + "a") {
                                                            case "":
                                                                v57 := GenS2{F0: -1}
                                                                _ = v57
                                                            case "\n\t":
                                                                v58 := []int{2147483647, (fi12 & (-(fi12)))}
                                                                _ = v58
                                                            case "hello":
                                                                var v59 uint64 = 112
                                                                _ = v59
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            v60 := GenS4{F0: 0, F1: (^(v36))}
                                                            _ = len(v2)
                                                            v5 = true
                                                            _ = v60
                                                        }
                                                    }
                                                    _ = fi48
                                                    _ = v50
                                                }
                                            }
                                        }
                                    }
                                    _ = v46
                                }
                            case "hello":
                                fmt.Println(v1, v1)
                            }
                            _ = v36
                        } else {
                            v4 += v4
                            var v61 int8 = 0
                            var v62 uint8 = 128
                            _ = v61
                            _ = v62
                        }
                    default:
                        if ((((("a" + ("hello" + (("a" + ("hello" + "\n\t")) + (("\n\t" + ("hello" + "")) + "foo bar")))) + "hello") + "") + "a") >= ("a" + "a")) {
                            v1 = v1
                        }
                    }
                }
                _ = fi12
            }
        default:
            fmt.Println(v7, v4)
        }
        v1 -= v1
        if true {
            if false {
                fmt.Println(v4)
                for fi63 := 0; fi63 < 11; fi63++ {
                    if false {
                        v4 = 924
                        var v64 uint = 65535
                        _ = v64
                    } else {
                        v65 := GenS3{F0: -32768, F1: -9223372036854775808, F2: 248}
                        var v66 uint8 = 136
                        for fi67 := 0; fi67 < 17; fi67++ {
                            if (0 > 6) {
                                var v68 int16 = 32767
                                switch "foo bar" {
                                case "hello":
                                    var v69 int = (fi63 - (fi67 - (fi63 >> ((4294967295) & 63))))
                                    _ = v69
                                case "":
                                    _ = len(v6)
                                default:
                                    v70 := []int{9223372036854775807, (fi67 ^ 32768), fi67, fi67}
                                    _ = v70
                                }
                                _ = v68
                            }
                            var v71 int = (fi67 % ((fi63) | 1))
                            switch "foo bar" {
                            case "\n\t":
                                v1 -= 0
                            case "foo bar":
                                var v72 int = (fi67 + fi63)
                                _ = v72
                            default:
                                if ((-2.25 + 3.14159) < -2.25) {
                                    var v73 uint = 42
                                    _ = v73
                                }
                            }
                            var v74 uint32 = (v1 >> ((4294967295) & 63))
                            _ = fi67
                            _ = v71
                            _ = v74
                        }
                        _ = v65
                        _ = v66
                    }
                    switch (fi63 - -9223372036854775808) {
                    case 42:
                        for fi75 := 0; fi75 < 3; fi75++ {
                            v1 += 0
                            fmt.Println(fi75)
                            var v76 uint8 = 0
                            _ = fi75
                            _ = v76
                        }
                    case -9223372036854775808:
                        v77 := []int{fi63, (fi63 - fi63), fi63}
                        _ = v77
                    case 481:
                        _ = len(v6)
                    default:
                        var v78 int32 = -353
                        _ = v78
                    }
                    _ = fi63
                }
            } else {
                fmt.Println(v5, v4, v7)
            }
            var v79 int64 = 9223372036854775807
            _ = v79
        } else {
            v5 = (v4 < (v4 << ((255) & 63)))
            var v80 uint32 = 65535
            for fi81 := 0; fi81 < 5; fi81++ {
                for fi82 := 0; fi82 < 18; fi82++ {
                    v7 += v7
                    v83 := []int{-955}
                    if false {
                        delete(v6, (v1 - v1))
                        delete(v3, ("foo bar" + ("\n\t" + "hello")))
                        _ = len(v83)
                    } else {
                        var v84 int64 = v4
                        fmt.Println(v80)
                        if len(v83) >= 2 {
                            v83 = v83[1:len(v83)]
                        }
                        _ = v84
                    }
                    _ = fi82
                    _ = v83
                }
                if (float32(3.14159) >= (-2.25 / float32(float64(0.1)))) {
                    var v85 string = "\n\t"
                    _ = v85
                }
                _ = fi81
            }
            _ = v80
        }
    } else {
        var v86 int = 215
        v7 = 0.0
        v4 = (v4 % ((0) | 1))
        _ = v86
    }
    fmt.Println(v4, v4, v4)
    v5 = (128 >= 42)
    var v87 uint32 = 2
    for fi88 := 0; fi88 < 3; fi88++ {
        v5 = (!(!false))
        _ = fi88
    }
    for fi89 := 0; fi89 < 18; fi89++ {
        v90 := GenS5{F0: 42}
        _ = fi89
        _ = v90
    }
    var v91 uint = 65535
    v7 += (0.1 + ((v7 - v7) * (v7 - (v7 * float32(0.0)))))
    _ = len(v2)
    fmt.Println(v5)
    for fi92 := 0; fi92 < 2; fi92++ {
        switch (fi92 | (^(fi92))) {
        case 9223372036854775807:
            v93 := GenS6{F0: 16, F1: fi92, F2: 255}
            _ = v93
        case -9223372036854775808:
            v94 := v2["foo bar"]
            _ = v94
        case 42:
            for fi95 := 0; fi95 < 1; fi95++ {
                v96 := GenS7{F0: v91, F1: v4}
                switch "\n\t" {
                case "":
                    if v5 {
                        var v97 uint16 = 65535
                        _ = v97
                    } else {
                        fmt.Println(v7, v91, v5)
                        if (!v5) {
                            var v98 float64 = (1.5 * (-2.25 + (0.1 + (0.0 - float64(v7)))))
                            if false {
                                switch "foo bar" {
                                case "":
                                    if ((v5 && ((219 != 128) && (0.0 < v7))) && (!v5)) {
                                        v5 = true
                                    } else {
                                        fmt.Println(v7, v91)
                                    }
                                case "\n\t":
                                    switch (fi95 ^ (fi92 - fi95)) {
                                    case 128:
                                        var v99 bool = true
                                        _ = v99
                                    case -171:
                                        v91 += (v91 & v91)
                                    case -734:
                                        _ = len(v6)
                                    default:
                                        var v100 uint = 2
                                        switch ("foo bar" + ("\n\t" + "a")) {
                                        case "\n\t":
                                            v101 := []int{fi95, 2147483647, (^(fi95)), (fi95 ^ 160)}
                                            if len(v101) >= 2 {
                                                v101 = v101[1:len(v101)]
                                            }
                                            if false {
                                                if (32767 > 1) {
                                                    switch fi92 {
                                                    case -851:
                                                        v5 = v5
                                                    case -1:
                                                        v102 := GenS6{F0: -128, F1: 4294967295, F2: 163}
                                                        var v103 uint16 = 0
                                                        if (!true) {
                                                            if (!((v7 - (float32(1.5) / (0.0 + (v7 - 1e10)))) == ((float32((3.14159 - v98)) + v7) * v7))) {
                                                                v100 = 1
                                                                var v104 uint = (^(v100))
                                                                _ = v104
                                                            }
                                                            v105 := GenS5{F0: 153}
                                                            v106 := GenS4{F0: -849, F1: (v91 + v100)}
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            v91 += (^(v100))
                                                        }
                                                        _ = v102
                                                        _ = v103
                                                    case 2147483647:
                                                        v7 = float32(v98)
                                                    }
                                                    v1 -= (v1 >> ((v91) & 63))
                                                } else {
                                                    _ = len(v3)
                                                    switch ("" + "a") {
                                                    case "hello":
                                                        _ = len(v2)
                                                    case "":
                                                        v5 = false
                                                    case "foo bar":
                                                        v91 = v100
                                                    default:
                                                        var v107 uint8 = 1
                                                        _ = v107
                                                    }
                                                    if v5 {
                                                        v4 = v4
                                                        v7 -= -2.25
                                                        var v108 uint8 = 255
                                                        _ = v108
                                                    } else {
                                                        v109 := v2[(("hello" + (("\n\t" + "hello") + "a")) + "a")]
                                                        v5 = true
                                                        v91 = 787
                                                        _ = v109
                                                    }
                                                }
                                            }
                                            _ = v101
                                        case "":
                                            switch fi92 {
                                            case 0:
                                                v110 := GenS7{F0: (v100 | v91), F1: (v4 >> (((v100 * (v91 >> (((v100 + v91)) & 63)))) & 63))}
                                                _ = v110
                                            case 32767:
                                                v111 := GenS6{F0: 42, F1: fi95, F2: 138}
                                                _ = v111
                                            default:
                                                if (!true) {
                                                    var v112 int16 = 856
                                                    _ = v112
                                                }
                                            }
                                        }
                                        switch (("a" + "hello") + "hello") {
                                        case "foo bar":
                                            var v113 int64 = (v4 * v4)
                                            v98 -= (((v98 - 0.0) / (v98 * v98)) / 1.5)
                                            var v114 string = ""
                                            _ = v113
                                            _ = v114
                                        case "\n\t":
                                            v115 := GenS8{F0: v5, F1: (v100 / (((v100 / (((v91 * 105)) | 1))) | 1))}
                                            _ = v115
                                        default:
                                            if (53 == 42) {
                                                v100 -= v91
                                                fmt.Println(v87)
                                            } else {
                                                v116 := GenS9{F0: 258, F1: 32767}
                                                if v5 {
                                                    fmt.Println(v1)
                                                }
                                                v1 = (v87 ^ v1)
                                                switch fi92 {
                                                case -768:
                                                    v117 := GenS2{F0: -255}
                                                    _ = v117
                                                case -795:
                                                    delete(v2, ("hello" + "hello"))
                                                }
                                                v87 += v87
                                                _ = v116
                                            }
                                        }
                                        _ = v100
                                    }
                                case "a":
                                    v118 := GenS6{F0: -1, F1: (fi92 ^ 1), F2: 1}
                                    _ = v118
                                default:
                                    var v119 uint64 = 437
                                    _ = v119
                                }
                                var v120 float64 = (v98 + ((v98 + (v98 - float64((1.5 * ((1e10 + (float32(float64(0.0)) * v7)) + (v7 * 3.14159)))))) - float64((3.14159 / v7))))
                                var v121 int32 = 423
                                v122 := v3[("foo bar" + ("" + "foo bar"))]
                                v123 := []int{fi92, -1, (fi95 & (fi95 + int(v4)))}
                                _ = v120
                                _ = v121
                                _ = v122
                                _ = v123
                            } else {
                                v124 := GenS9{F0: 4294967295, F1: (fi95 & (fi92 << ((v91) & 63)))}
                                v125 := []int{}
                                var v126 bool = true
                                _ = v124
                                _ = v125
                                _ = v126
                            }
                            v5 = (42 > 0)
                            _ = v98
                        }
                    }
                default:
                    v127 := GenS8{F0: (128 <= 1), F1: v91}
                    _ = v127
                }
                _ = fi95
                _ = v96
            }
        default:
            v128 := []int{127}
            _ = v128
        }
        v129 := GenS8{F0: (v5 || v5), F1: (v91 * v91)}
        _ = fi92
        _ = v129
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v87:", v87)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v87
    _ = v91
}
