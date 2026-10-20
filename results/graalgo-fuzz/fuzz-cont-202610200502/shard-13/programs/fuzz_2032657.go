package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 int8; F2 int }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 int8; F1 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int; F1 uint16; F2 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16; F1 string; F2 uint32 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int32; F1 uint32; F2 uint }
type GenS20 struct { F0 uint8; F1 uint; F2 int16 }

func main() {
    var v1 uint16 = 260
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    var v4 bool = (0.0 >= v3)
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            if (!true) {
                switch ("hello" + "foo bar") {
                case "a":
                    var v7 uint = 628
                    _ = v7
                case "hello":
                    if (!v4) {
                        v8 := GenS1{F0: 705, F1: 2, F2: "a"}
                        v9 := []int{}
                        v9 = append(v9, fi5)
                        _ = v8
                        _ = v9
                    } else {
                        v3 += (0.1 / ((v3 + v3) + v3))
                        v2[fi6] = fi6
                    }
                case "foo bar":
                    switch "" {
                    case "\n\t":
                        var v10 int64 = -421
                        _ = v10
                    case "foo bar":
                        v11 := []int{fi6}
                        _ = v11
                    default:
                        v12 := GenS2{F0: v4}
                        _ = v12
                    }
                }
                v13 := GenS2{F0: (v4 || v4)}
                _ = v13
            } else {
                var v14 bool = ((-668 > 186) && (!false))
                v15 := GenS1{F0: fi6, F1: 709, F2: ""}
                v3 = v3
                switch "a" {
                case "":
                    v16 := []int{fi6, fi6, 198, (^(fi6))}
                    _ = v16
                case "a":
                    var v17 uint32 = 1
                    _ = v17
                case "hello":
                    switch 2147483647 {
                    case 42:
                        v1 = v1
                    case 83:
                        v18 := v2[-9223372036854775808]
                        _ = v18
                    case 0:
                        switch "a" {
                        case "":
                            if v4 {
                                var v19 int8 = 42
                                v14 = (v1 >= v1)
                                var v20 int64 = 1
                                _ = v19
                                _ = v20
                            } else {
                                var v21 int = -610
                                if true {
                                    var v22 uint16 = (v1 + v1)
                                    var v23 float64 = 3.14159
                                    if v4 {
                                        var v24 bool = v4
                                        _ = v24
                                    }
                                    _ = v22
                                    _ = v23
                                }
                                switch "a" {
                                case "":
                                    v25 := GenS3{F0: 2147483648, F1: 4294967295}
                                    switch (^(v21)) {
                                    case -552:
                                        if v4 {
                                            v1 -= 0
                                        } else {
                                            v14 = true
                                        }
                                    default:
                                        var v26 bool = v14
                                        _ = v26
                                    }
                                    _ = v25
                                case "a":
                                    v27 := v2[v21]
                                    _ = v27
                                default:
                                    var v28 float32 = 0.0
                                    _ = v28
                                }
                                v4 = false
                                _ = v21
                            }
                        case "a":
                            if v14 {
                                fmt.Println(fi5, v3, v4)
                                var v29 int64 = 1
                                v3 -= float32(float64(v3))
                                _ = v29
                            } else {
                                fmt.Println(fi5)
                            }
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                v1 = (v1 % (((v1 >> ((42) & 63))) | 1))
                            case "":
                                v30 := GenS4{F0: 1}
                                _ = v30
                            default:
                                var v31 int16 = 0
                                _ = v31
                            }
                        default:
                            v3 = (v3 * (((v3 - float32((0.0 * 1e10))) - v3) / (((v3 - 0.0) / float32(1.5)) * v3)))
                        }
                    }
                }
                fmt.Println(fi6, v1)
                _ = v14
                _ = v15
            }
            _ = fi6
        }
        v32 := v2[fi5]
        var v33 int64 = -9223372036854775808
        _ = fi5
        _ = v32
        _ = v33
    }
    delete(v2, 0)
    var v34 int64 = 369
    v4 = (!v4)
    _ = len(v2)
    switch 9223372036854775807 {
    case 9223372036854775807:
        v35 := GenS5{F0: 127, F1: 42}
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := GenS2{F0: v4}
        switch ("\n\t" + "") {
        case "\n\t":
            if false {
                v38 := []int{fi36, -337, 32768, (fi36 | 66)}
                for fi39 := 0; fi39 < 13; fi39++ {
                    fmt.Println(v4)
                    v3 = v3
                    _ = fi39
                }
                v40 := GenS6{F0: -1}
                _ = v38
                _ = v40
            } else {
                var v41 bool = false
                v3 += -2.25
                _ = v41
            }
        case "":
            v4 = v4
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 bool = (0 != -1)
            var v44 uint64 = 317
            if false {
                switch fi36 {
                case -338:
                    v45 := GenS2{F0: v4}
                    _ = v45
                }
                if (v44 >= (v44 & v44)) {
                    switch -395 {
                    case 32767:
                        v34 = v34
                    case 128:
                        if v43 {
                            fmt.Println(v1, v4, v3)
                            if ((float32((3.14159 - 1.5)) - -2.25) <= -2.25) {
                                if v4 {
                                    if v4 {
                                        switch (("" + ("" + ("hello" + "foo bar"))) + "\n\t") {
                                        case "\n\t":
                                            if (0 > 255) {
                                                fmt.Println(fi36)
                                                v46 := GenS7{F0: v4, F1: 100, F2: 2147483647}
                                                v47 := GenS8{F0: v4, F1: 0}
                                                if v4 {
                                                    var v48 int64 = (v34 << ((0) & 63))
                                                    _ = v48
                                                }
                                                v3 = 1e10
                                                if (-920 <= 823) {
                                                    v1 -= 65535
                                                    v4 = v43
                                                    v49 := GenS4{F0: (v34 + (v34 << ((255) & 63)))}
                                                    if false {
                                                        if (v1 == 65535) {
                                                            fmt.Println(v44)
                                                            v43 = false
                                                            var v50 bool = (233 >= int16(v34))
                                                            _ = v50
                                                        }
                                                        v4 = (-749 == v34)
                                                        v51 := GenS3{F0: 42, F1: (v44 | v44)}
                                                        _ = v51
                                                    }
                                                    var v52 uint64 = v44
                                                    _ = v49
                                                    _ = v52
                                                }
                                                _ = v46
                                                _ = v47
                                            }
                                        case "foo bar":
                                            _ = len(v2)
                                        }
                                    } else {
                                        if true {
                                            var v53 bool = v43
                                            _ = v53
                                        }
                                    }
                                    var v54 uint16 = v1
                                    _ = v54
                                } else {
                                    var v55 int64 = v34
                                    if ((v1 | (v1 & v1)) < 32767) {
                                        switch 127 {
                                        case -1:
                                            if (false && false) {
                                                fmt.Println(v43, fi36)
                                                var v56 uint32 = 2147483648
                                                var v57 uint8 = 128
                                                _ = v56
                                                _ = v57
                                            }
                                        case 4294967295:
                                            fmt.Println(v34, v3, v55)
                                        default:
                                            switch fi42 {
                                            case -9223372036854775808:
                                                v1 = (v1 << ((274) & 63))
                                            case 1:
                                                v1 -= v1
                                            case 42:
                                                if (2 < 1) {
                                                    if true {
                                                        var v58 float64 = ((3.14159 * 3.14159) * 3.14159)
                                                        v3 -= v3
                                                        _ = v58
                                                    } else {
                                                        if (!v43) {
                                                            v59 := []int{-431, fi42}
                                                            _ = v59
                                                        }
                                                    }
                                                    fmt.Println(v44, v3)
                                                } else {
                                                    v60 := GenS9{F0: 2, F1: v1}
                                                    switch fi42 {
                                                    case 127:
                                                        switch fi42 {
                                                        case 128:
                                                            fmt.Println(v44, v4, v3)
                                                        default:
                                                            v34 = -1
                                                        }
                                                    case 1:
                                                        v34 += (v55 << ((106) & 63))
                                                    case -43:
                                                        switch ("" + "a") {
                                                        case "\n\t":
                                                            fmt.Println(v3, fi42)
                                                        case "":
                                                            if false {
                                                                v61 := GenS10{F0: 1, F1: -2147483648}
                                                                var v62 uint64 = 2
                                                                _ = v61
                                                                _ = v62
                                                            }
                                                            fmt.Println(v55, v44, v1)
                                                        default:
                                                            var v63 int = -2147483648
                                                            _ = v63
                                                        }
                                                    }
                                                    v4 = (!(!(2147483647 > 1)))
                                                    _ = v60
                                                }
                                            default:
                                                var v64 string = ("hello" + "a")
                                                _ = v64
                                            }
                                        }
                                        v65 := GenS11{F0: 265}
                                        v34 = 0
                                        var v66 float64 = (3.14159 * 3.14159)
                                        _ = v65
                                        _ = v66
                                    } else {
                                        fmt.Println(v43, v34)
                                        if true {
                                            v43 = (9223372036854775807 != v55)
                                            v2[fi42] = (fi36 + fi36)
                                            delete(v2, -2147483648)
                                        }
                                        if true {
                                            if true {
                                                switch ((("" + "foo bar") + "") + ("a" + ("foo bar" + ""))) {
                                                case "hello":
                                                    v67 := []int{fi42, fi42, fi42}
                                                    _ = v67
                                                case "\n\t":
                                                    var v68 bool = (91 <= 1)
                                                    _ = v68
                                                default:
                                                    if true {
                                                        if v43 {
                                                            v69 := GenS12{F0: (fi36 / (((^(fi42))) | 1)), F1: (^(v1)), F2: 0}
                                                            _ = v69
                                                        }
                                                        fmt.Println(fi42)
                                                        var v70 uint32 = 2147483648
                                                        _ = v70
                                                    } else {
                                                        var v71 uint64 = v44
                                                        fmt.Println(fi36, v43)
                                                        v72 := []int{}
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                }
                                                v73 := []int{2147483647, fi42}
                                                if ((((249 >= 2147483647) || ((!v43) && false)) || (2147483647 < -1)) || v43) {
                                                    var v74 uint32 = 2147483647
                                                    _ = v74
                                                } else {
                                                    fmt.Println(fi36)
                                                }
                                                v75 := GenS3{F0: 503, F1: (v44 ^ (v44 ^ 0))}
                                                _ = len(v2)
                                                _ = v73
                                                _ = v75
                                            }
                                            switch fi42 {
                                            case 10:
                                                var v76 uint = 42
                                                _ = v76
                                            }
                                        }
                                    }
                                    _ = v55
                                }
                                var v77 int32 = int32(v34)
                                _ = v77
                            }
                            var v78 int = (^(fi36))
                            v79 := GenS8{F0: (-58 < 0), F1: -1}
                            switch "foo bar" {
                            case "a":
                                var v80 int64 = v34
                                _ = v80
                            default:
                                v81 := []int{0, v78}
                                _ = v81
                            }
                            var v82 uint16 = 546
                            _ = v78
                            _ = v79
                            _ = v82
                        } else {
                            v83 := []int{(fi36 + fi42), -1}
                            _ = v83
                        }
                    default:
                        if (true && (v43 && (v44 != (v44 % ((uint64(v1)) | 1))))) {
                            switch "foo bar" {
                            case "\n\t":
                                var v84 bool = (4294967295 == 890)
                                _ = v84
                            default:
                                if v4 {
                                    if (((v34 << ((18446744073709551615) & 63)) > v34) || (!v4)) {
                                        v43 = (fi42 == fi42)
                                    } else {
                                        if ("foo bar" >= "hello") {
                                            v85 := GenS5{F0: fi36, F1: 1}
                                            var v86 int = 127
                                            var v87 int8 = -128
                                            fmt.Println(v4)
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        }
                                    }
                                    var v88 uint32 = 2
                                    var v89 uint8 = 126
                                    v90 := GenS13{F0: (("" + ("hello" + "foo bar")) + "\n\t")}
                                    v1 = v1
                                    v43 = v43
                                    var v91 uint32 = (v88 * 2147483648)
                                    var v92 float64 = 1e10
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                    _ = v92
                                } else {
                                    if v43 {
                                        fmt.Println(v4, v34, v44)
                                        fmt.Println(v43, fi42, fi42)
                                    } else {
                                        fmt.Println(v44, fi36)
                                        v4 = (4294967295 <= uint(v44))
                                        if true {
                                            if v43 {
                                                v93 := GenS6{F0: -1}
                                                var v94 uint32 = 65535
                                                _ = v93
                                                _ = v94
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    v1 = 65535
                                                default:
                                                    if ("\n\t" <= "a") {
                                                        if (!true) {
                                                            v44 = 100
                                                            switch (("" + "a") + ("" + "\n\t")) {
                                                            case "foo bar":
                                                                switch (fi36 * fi36) {
                                                                case 0:
                                                                    var v95 int8 = -123
                                                                    _ = v95
                                                                case 5:
                                                                    var v96 bool = v43
                                                                    _ = v96
                                                                default:
                                                                    v97 := GenS14{F0: 32767, F1: "hello", F2: uint32(v1)}
                                                                    _ = v97
                                                                }
                                                            case "\n\t":
                                                                switch (fi42 - (fi36 + fi36)) {
                                                                case -440:
                                                                    v98 := GenS11{F0: 2}
                                                                    _ = v98
                                                                case -745:
                                                                    if (364 > 2) {
                                                                        if v43 {
                                                                            v99 := GenS15{F0: -2147483648}
                                                                            _ = v99
                                                                        } else {
                                                                            v43 = (0 == 0)
                                                                            var v100 uint64 = (v44 >> ((944) & 63))
                                                                            fmt.Println(v100, v100)
                                                                            if true {
                                                                                v44 += (^(v44))
                                                                            }
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                case 789:
                                                                    var v101 uint = 65535
                                                                    _ = v101
                                                                }
                                                            case "hello":
                                                                if (!(!(1.5 > (-2.25 - (1.5 / float64(3.14159)))))) {
                                                                    v4 = v43
                                                                    v43 = (!true)
                                                                    v3 = 0.0
                                                                } else {
                                                                    var v102 string = "foo bar"
                                                                    var v103 float64 = (1e10 - 0.1)
                                                                    _ = v102
                                                                    _ = v103
                                                                }
                                                            }
                                                            v104 := GenS8{F0: true, F1: 127}
                                                            _ = v104
                                                        }
                                                    } else {
                                                        var v105 int8 = -23
                                                        if (-2.25 >= float64(((v3 / float32(0.1)) + v3))) {
                                                            _ = len(v2)
                                                        } else {
                                                            switch -1 {
                                                            case -590:
                                                                var v106 int8 = (^(v105))
                                                                var v107 int64 = int64(v44)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v108 string = "foo bar"
                                                                    _ = v108
                                                                case "foo bar":
                                                                    fmt.Println(v44, v107)
                                                                }
                                                                _ = v106
                                                                _ = v107
                                                            case 2147483647:
                                                                v34 += (v34 >> ((2) & 63))
                                                            default:
                                                                v109 := GenS3{F0: 4294967295, F1: 579}
                                                                _ = v109
                                                            }
                                                            if ((v105 >= (v105 ^ v105)) || false) {
                                                                fmt.Println(v44, v4)
                                                                v110 := GenS16{F0: 42}
                                                                var v111 uint16 = 2
                                                                switch (-(fi36)) {
                                                                case 127:
                                                                    if (42 < (v44 | v44)) {
                                                                        if (3.14159 < ((1.5 + 0.1) + (0.0 / ((0.1 * 1e10) / ((0.0 / -2.25) + (-2.25 / 1.5)))))) {
                                                                            v3 = (v3 - (float32((0.0 / (0.1 + 0.1))) / (v3 / v3)))
                                                                            v112 := []int{1}
                                                                            v113 := GenS8{F0: (((fi42 | 32767) < (fi36 | fi36)) && v4), F1: v105}
                                                                            _ = v112
                                                                            _ = v113
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v4)
                                                                            fmt.Println(fi42, v111)
                                                                            v105 = -30
                                                                        } else {
                                                                            if true {
                                                                                v4 = true
                                                                            } else {
                                                                                v114 := GenS17{F0: v34, F1: (v44 & (v44 >> ((219) & 63))), F2: 127}
                                                                                var v115 bool = false
                                                                                var v116 int16 = 42
                                                                                _ = v114
                                                                                _ = v115
                                                                                _ = v116
                                                                            }
                                                                            _ = len(v2)
                                                                        }
                                                                        switch fi36 {
                                                                        case -9223372036854775808:
                                                                            var v117 string = "\n\t"
                                                                            _ = v117
                                                                        }
                                                                    }
                                                                }
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                        }
                                                        switch (fi36 + fi42) {
                                                        case 9223372036854775807:
                                                            v43 = (true && (v43 && (float32(float64(v3)) != (((v3 + 1e10) - v3) / v3))))
                                                        case 32767:
                                                            v118 := GenS8{F0: ((1.5 - float64(0.0)) < ((1.5 * float64((0.0 * v3))) + 0.0)), F1: int8(v34)}
                                                            _ = v118
                                                        default:
                                                            v119 := GenS18{F0: (-(v105))}
                                                            _ = v119
                                                        }
                                                        _ = v105
                                                    }
                                                }
                                            }
                                        }
                                        fmt.Println(v3, v44, fi42)
                                        v120 := []int{-9223372036854775808, -208, fi36, fi42}
                                        _ = v120
                                    }
                                    var v121 int = -124
                                    _ = v121
                                }
                            }
                            v122 := v2[fi42]
                            _ = v122
                        }
                    }
                    var v123 float64 = ((-2.25 + -2.25) - float64(float32(3.14159)))
                    _ = v123
                } else {
                    v3 = (v3 - float32((0.1 + (float64(float32(3.14159)) - 0.0))))
                    delete(v2, 4294967295)
                }
                v124 := []int{32768, (fi36 / ((fi42) | 1)), (-(fi36)), (fi36 + 629)}
                _ = v124
            } else {
                v43 = (v3 < -2.25)
                v125 := []int{fi42, (fi36 >> ((905) & 63)), fi36}
                delete(v2, (fi42 * (fi36 + 2147483647)))
                if len(v125) >= 2 {
                    v125 = v125[1:len(v125)]
                }
                var v126 float64 = 0.1
                _ = v125
                _ = v126
            }
            v127 := GenS15{F0: -2147483648}
            var v128 uint16 = (v1 / ((v1) | 1))
            _ = fi42
            _ = v127
            _ = v128
            _ = v43
            _ = v44
        }
        _ = fi36
        _ = v37
    }
    v129 := GenS19{F0: 617, F1: 2, F2: 18446744073709551615}
    var v130 int = 32767
    v131 := GenS20{F0: 250, F1: 18446744073709551615, F2: -467}
    v4 = false
    var v132 int32 = 42
    var v133 int32 = (v132 / (((v132 * (v132 * v132))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
