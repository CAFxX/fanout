package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 string; F1 uint32; F2 int32 }
type GenS3 struct { F0 int; F1 string; F2 int }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint; F1 int64; F2 int8 }
type GenS6 struct { F0 bool; F1 uint; F2 int32 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 int; F1 int64; F2 uint8 }
type GenS10 struct { F0 int; F1 bool; F2 uint16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 int; F2 uint8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS16 struct { F0 int64; F1 uint64; F2 int16 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 string; F1 uint64 }
type GenS19 struct { F0 uint; F1 bool; F2 uint16 }
type GenS20 struct { F0 string; F1 uint8; F2 uint32 }
type GenS21 struct { F0 bool; F1 uint8; F2 int16 }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 int }
type GenS24 struct { F0 bool; F1 int32 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 uint64 }
type GenS27 struct { F0 uint64; F1 int16; F2 bool }
func genFn0(p0 string) uint {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 float64 = 0.0
        if (-32768 < 1) {
            v8 := GenS1{F0: 1, F1: -1, F2: true}
            _ = v8
        }
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 = ("hello" + "foo bar")
        _ = fi9
    }
    v10 := []int{-701, 32768, -578}
    _ = p0
    _ = v10
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2147483648
    v3 := make(map[int]string)
    var v4 float64 = 1.5
    v5 := make(map[string]int)
    switch "" {
    case "\n\t":
        fmt.Println(v2, v4)
    case "hello":
        v11 := []int{-1, -2147483648, 391, 128}
        _ = v11
    }
    if (((v2 + v2) >= 2) && ((false && (!("a" == ""))) && (false || (1 <= 1)))) {
        v12 := GenS1{F0: 42, F1: -1, F2: (65535 > 724)}
        _ = v12
    } else {
        var v13 float64 = 1.5
        if (!(true || true)) {
            if (false || (!(1 <= 255))) {
                v14 := GenS2{F0: "a", F1: v2, F2: 775}
                var v15 uint64 = 343
                _ = v14
                _ = v15
            }
            if (!true) {
                var v16 uint32 = 953
                _ = v16
            } else {
                if true {
                    switch ("a" + "hello") {
                    case "\n\t":
                        for fi17 := 0; fi17 < 6; fi17++ {
                            for fi18 := 0; fi18 < 2; fi18++ {
                                v4 -= v4
                                v19 := GenS2{F0: "foo bar", F1: 42, F2: 42}
                                switch 32768 {
                                case 42:
                                    v4 += 0.0
                                case 668:
                                    if false {
                                        if true {
                                            v20 := []int{(fi17 ^ (fi18 & -1)), (fi17 + fi17)}
                                            switch fi18 {
                                            case 128:
                                                if (42 <= 255) {
                                                    v21 := GenS3{F0: (fi18 / ((fi17) | 1)), F1: "\n\t", F2: -927}
                                                    var v22 int = fi18
                                                    _ = len(v1)
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    var v23 uint32 = (v2 / (((v2 >> ((65535) & 63))) | 1))
                                                    _ = v23
                                                }
                                            case -1:
                                                v24 := GenS3{F0: (fi18 << ((236) & 63)), F1: "", F2: (^(fi17))}
                                                fmt.Println(fi18, fi18)
                                                fmt.Println(fi17, fi17, v13)
                                                _ = v24
                                            }
                                            _ = v20
                                        }
                                    } else {
                                        if (int32(v2) == 2147483647) {
                                            var v25 string = "foo bar"
                                            _ = v25
                                        } else {
                                            v4 += 3.14159
                                            if (false && ((float32((v13 / v13)) != 1.5) && (127 >= 67))) {
                                                v2 = 4294967295
                                            } else {
                                                var v26 uint64 = 42
                                                v27 := []int{665, (fi18 >> ((106) & 63))}
                                                var v28 float32 = float32(1.5)
                                                v4 = 0.0
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := []int{}
                                            if true {
                                                var v30 float32 = (1.5 / -2.25)
                                                if false {
                                                    if len(v29) >= 2 {
                                                        v29 = v29[1:len(v29)]
                                                    }
                                                }
                                                v13 -= (3.14159 / v4)
                                                _ = v30
                                            }
                                            fmt.Println(v4)
                                            _ = v29
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            v31 := []int{-144}
                                            _ = v31
                                        case "":
                                            v2 += (v2 % (((v2 + v2)) | 1))
                                        case "foo bar":
                                            var v32 uint32 = (^(v2))
                                            _ = v32
                                        }
                                    }
                                case 127:
                                    v13 = (1.5 * (v4 - v4))
                                }
                                _ = fi18
                                _ = v19
                            }
                            var v33 string = ((((((("hello" + (("foo bar" + "hello") + "hello")) + ("\n\t" + "\n\t")) + "") + "hello") + ((("foo bar" + "\n\t") + (("" + "\n\t") + ("foo bar" + ("hello" + "")))) + (("foo bar" + ("" + "a")) + "a"))) + (("foo bar" + "") + ("a" + "a"))) + "foo bar")
                            _ = fi17
                            _ = v33
                        }
                    }
                    switch 4294967295 {
                    case -27:
                        var v34 uint = 216
                        _ = v34
                    default:
                        v35 := GenS3{F0: 0, F1: ("\n\t" + "\n\t"), F2: -1}
                        _ = v35
                    }
                } else {
                    var v36 uint = 65535
                    for fi37 := 0; fi37 < 11; fi37++ {
                        var v38 uint8 = 192
                        _ = fi37
                        _ = v38
                    }
                    _ = v36
                }
            }
            for fi39 := 0; fi39 < 14; fi39++ {
                v40 := GenS4{F0: 216, F1: -112}
                _ = fi39
                _ = v40
            }
        } else {
            switch 42 {
            case 4294967295:
                _ = len(v3)
            case -2147483648:
                v41 := v1[1]
                _ = v41
            case 42:
                for fi42 := 0; fi42 < 5; fi42++ {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        var v44 float64 = v4
                        v1[(-(fi42))] = (fi43 >> ((4294967295) & 63))
                        if (true && (!false)) {
                            var v45 float64 = v4
                            _ = v45
                        }
                        _ = fi43
                        _ = v44
                    }
                    v4 = (1.5 * 1.5)
                    _ = fi42
                }
            default:
                v46 := GenS5{F0: 4294967295, F1: -93, F2: 100}
                _ = v46
            }
            var v47 string = "foo bar"
            v48 := GenS1{F0: 996, F1: 9223372036854775807, F2: false}
            _ = v47
            _ = v48
        }
        var v49 uint64 = 4294967295
        _ = v13
        _ = v49
    }
    if false {
        v50 := GenS1{F0: 18446744073709551615, F1: 1, F2: true}
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 float32 = (0.1 + 0.0)
            for fi53 := 0; fi53 < 6; fi53++ {
                fmt.Println(v2, v52)
                switch -600 {
                case 4294967295:
                    v54 := []int{}
                    _ = v54
                }
                if false {
                    if false {
                        fmt.Println(fi53, v2, v2)
                    } else {
                        fmt.Println(v2, fi53)
                    }
                    v2 -= v2
                }
                _ = fi53
            }
            _ = fi51
            _ = v52
        }
        _ = v50
    }
    if (-728 == 32767) {
        for fi55 := 0; fi55 < 9; fi55++ {
            var v56 string = ("" + "")
            _ = fi55
            _ = v56
        }
    } else {
        v2 = v2
    }
    if false {
        var v57 int8 = -94
        var v58 uint16 = 0
        switch -964 {
        case 128:
            v59 := []int{-546, 4294967295, -9223372036854775808, 396}
            _ = v59
        case 42:
            v58 -= v58
        case 1:
            var v60 uint8 = 0
            _ = v60
        default:
            v61 := []int{127, 128, -2147483648, 225}
            _ = v61
        }
        _ = v57
        _ = v58
    } else {
        var v62 uint = 667
        if true {
            if true {
                v63 := []int{9223372036854775807, 32768, -2147483648}
                _ = v63
            } else {
                if false {
                    v64 := GenS6{F0: (54 == 127), F1: 356, F2: 42}
                    var v65 string = ("foo bar" + "\n\t")
                    if false {
                        fmt.Println(v4, v2, v2)
                        for fi66 := 0; fi66 < 8; fi66++ {
                            for fi67 := 0; fi67 < 2; fi67++ {
                                if false {
                                    v68 := GenS1{F0: 4294967295, F1: int64(v2), F2: false}
                                    var v69 int64 = 0
                                    if (v65 < v65) {
                                        switch (fi67 - (^(fi66))) {
                                        case -1:
                                            switch (^(fi67)) {
                                            case 127:
                                                var v70 uint64 = 1
                                                _ = v70
                                            }
                                        case 4294967295:
                                            switch (fi67 - (fi67 >> (((v62 + (v62 ^ uint(v2)))) & 63))) {
                                            case -2147483648:
                                                var v71 int16 = -1
                                                _ = v71
                                            case 32768:
                                                v72 := GenS4{F0: 42, F1: -53}
                                                if (int16(fi67) >= 0) {
                                                    var v73 bool = ((^(v2)) > (v2 & v2))
                                                    v74 := GenS7{F0: 65535, F1: v73, F2: v62}
                                                    _ = v73
                                                    _ = v74
                                                }
                                                if false {
                                                    fmt.Println(fi66)
                                                    delete(v1, (^(fi66)))
                                                } else {
                                                    var v75 uint8 = 127
                                                    fmt.Println(v62, v75, v62)
                                                    _ = v75
                                                }
                                                v65 = ((("foo bar" + "") + "\n\t") + (v65 + v65))
                                                fmt.Println(v2)
                                                v76 := []int{}
                                                _ = v72
                                                _ = v76
                                            }
                                        case 823:
                                            switch fi67 {
                                            case 2147483647:
                                                var v77 float32 = float32((((v4 - (v4 * 0.1)) / 1.5) - v4))
                                                _ = v77
                                            case -9223372036854775808:
                                                var v78 uint = 0
                                                _ = v78
                                            }
                                            v79 := GenS6{F0: false, F1: 0, F2: 597}
                                            var v80 uint64 = 775
                                            _ = v79
                                            _ = v80
                                        }
                                        if false {
                                            switch 9223372036854775807 {
                                            case -2147483648:
                                                v81 := GenS7{F0: 314, F1: true, F2: (v62 ^ v62)}
                                                _ = v81
                                            }
                                            v82 := []int{fi67}
                                            _ = v82
                                        } else {
                                            var v83 uint16 = 1
                                            _ = v83
                                        }
                                        if true {
                                            v4 = v4
                                            if false {
                                                var v84 int = 9223372036854775807
                                                _ = v84
                                            }
                                            var v85 bool = false
                                            _ = v85
                                        }
                                        v86 := GenS4{F0: 18446744073709551615, F1: 127}
                                        _ = v86
                                    }
                                    var v87 int32 = 0
                                    var v88 uint64 = 2
                                    fmt.Println(v69)
                                    if true {
                                        if (v65 > (v65 + (v65 + ((((v65 + v65) + v65) + v65) + v65)))) {
                                            v89 := GenS8{F0: 399, F1: 253}
                                            if true {
                                                var v90 uint32 = uint32(v62)
                                                _ = v90
                                            } else {
                                                if (1 > 127) {
                                                    if true {
                                                        if true {
                                                            switch (fi67 >> ((v62) & 63)) {
                                                            case -9223372036854775808:
                                                                fmt.Println(v88, v69)
                                                            case 607:
                                                                v91 := GenS9{F0: fi67, F1: (v69 & (^(v69))), F2: 128}
                                                                _ = v91
                                                            case 0:
                                                                var v92 float32 = -2.25
                                                                _ = v92
                                                            default:
                                                                v88 = (v88 * v88)
                                                            }
                                                        } else {
                                                            var v93 uint16 = 65535
                                                            delete(v3, (^(fi66)))
                                                            switch 42 {
                                                            case -2147483648:
                                                                v65 = v65
                                                            case 1:
                                                                var v94 uint16 = 0
                                                                _ = v94
                                                            case 846:
                                                                v95 := GenS4{F0: (v88 % (((^(v88))) | 1)), F1: -1}
                                                                _ = v95
                                                            default:
                                                                delete(v5, ("" + "\n\t"))
                                                            }
                                                            switch ("" + ((v65 + ("\n\t" + v65)) + "foo bar")) {
                                                            case "hello":
                                                                var v96 int8 = 42
                                                                _ = v96
                                                            case "":
                                                                var v97 uint16 = (v93 >> ((v62) & 63))
                                                                _ = v97
                                                            }
                                                            _ = v93
                                                        }
                                                    } else {
                                                        v98 := GenS10{F0: int(v88), F1: ((v2 / ((v2) | 1)) > (v2 ^ uint32(v62))), F2: 739}
                                                        if true {
                                                            if ((v69 >= -9223372036854775808) || ((v4 - (float64(3.14159) - float64(float32(0.1)))) >= (float64(0.0) - (float64(((1.5 + 0.0) * -2.25)) / v4)))) {
                                                                fmt.Println(fi67, fi67, v69)
                                                            }
                                                            v99 := GenS7{F0: 42, F1: false, F2: genFn0("")}
                                                            _ = v99
                                                        }
                                                        var v100 int64 = v69
                                                        _ = v100
                                                        _ = v98
                                                    }
                                                    fmt.Println(v69)
                                                    v65 = (((("a" + v65) + ("a" + (v65 + v65))) + (((v65 + v65) + v65) + v65)) + v65)
                                                }
                                                v101 := GenS11{F0: (-(fi67))}
                                                _ = v101
                                            }
                                            fmt.Println(fi67)
                                            _ = v89
                                        } else {
                                            if (113 == 32767) {
                                                v102 := []int{fi67}
                                                _ = v102
                                            } else {
                                                v103 := v3[fi67]
                                                if (!false) {
                                                    v4 -= 3.14159
                                                    fmt.Println(v2, fi66, v4)
                                                    v62 = 18446744073709551615
                                                }
                                                _ = v103
                                            }
                                        }
                                        v104 := v5["a"]
                                        var v105 string = ((v65 + v65) + v65)
                                        var v106 float32 = ((1e10 * (-2.25 - 3.14159)) * -2.25)
                                        fmt.Println(v87)
                                        switch ((v65 + ("hello" + v65)) + (("hello" + (((v65 + (("\n\t" + v65) + (v65 + v105))) + "\n\t") + v65)) + "hello")) {
                                        case "a":
                                            fmt.Println(v104, v4, v88)
                                        case "\n\t":
                                            var v107 int64 = (v69 * (v69 * 270))
                                            _ = v107
                                        default:
                                            if false {
                                                fmt.Println(v87, fi67, v88)
                                                if (0 < -104) {
                                                    v105 = ((v65 + "a") + (((v105 + (((v105 + (v65 + "hello")) + (v65 + "\n\t")) + (("" + v65) + v65))) + "a") + ((v105 + v65) + "")))
                                                    if ((v88 >> ((uint(v88)) & 63)) >= uint64(v62)) {
                                                        switch v105 {
                                                        case "":
                                                            v69 += (^(v69))
                                                        case "a":
                                                            v104 += fi66
                                                        default:
                                                            switch (-(fi67)) {
                                                            case 42:
                                                                var v108 uint16 = 255
                                                                _ = v108
                                                            case 544:
                                                                fmt.Println(fi66)
                                                            default:
                                                                var v109 bool = (false && (0.0 > (v4 - (v4 - v4))))
                                                                v109 = (-2.25 != ((v106 + (v106 - -2.25)) / v106))
                                                                fmt.Println(v2, v2)
                                                                v110 := GenS9{F0: fi66, F1: v69, F2: 255}
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                        }
                                                        fmt.Println(v69, v104, v69)
                                                    } else {
                                                        if true {
                                                            v111 := []int{-455, fi67, (v104 + (fi67 | (fi66 >> ((genFn0("hello")) & 63)))), fi67}
                                                            switch "" {
                                                            case "":
                                                                if len(v111) >= 2 {
                                                                    v111 = v111[1:len(v111)]
                                                                }
                                                            case "a":
                                                                fmt.Println(fi66)
                                                                var v112 string = "\n\t"
                                                                _ = v112
                                                            }
                                                            _ = v111
                                                        }
                                                    }
                                                    v65 = "hello"
                                                    if false {
                                                        v113 := []int{}
                                                        v69 += (v69 ^ (-(v69)))
                                                        v62 = v62
                                                        if false {
                                                            _ = len(v113)
                                                            v114 := GenS11{F0: (v104 * fi67)}
                                                            _ = v114
                                                        }
                                                        _ = v113
                                                    }
                                                } else {
                                                    v2 = v2
                                                    switch ((v105 + "") + ((v65 + v105) + v65)) {
                                                    case "hello":
                                                        v115 := GenS2{F0: (v65 + ""), F1: uint32(v88), F2: (v87 + v87)}
                                                        _ = v115
                                                    case "foo bar":
                                                        var v116 int32 = 0
                                                        v117 := GenS3{F0: fi67, F1: v65, F2: (-(fi66))}
                                                        switch "a" {
                                                        case "foo bar":
                                                            v118 := GenS12{F0: (^(v2))}
                                                            _ = v118
                                                        case "":
                                                            if true {
                                                                v62 -= (v62 & v62)
                                                                v4 = (v4 - v4)
                                                                var v119 uint16 = 255
                                                                _ = v119
                                                            }
                                                        case "a":
                                                            v69 = (-(v69))
                                                        default:
                                                            var v120 int16 = -32768
                                                            _ = v120
                                                        }
                                                        _ = v116
                                                        _ = v117
                                                    default:
                                                        fmt.Println(v65, fi66)
                                                    }
                                                    _ = len(v5)
                                                }
                                            } else {
                                                v121 := GenS13{F0: v105, F1: (v104 | fi66), F2: 0}
                                                v65 = v65
                                                _ = v121
                                            }
                                        }
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                    v4 += v4
                                    fmt.Println(v2, v88, fi66)
                                    fmt.Println(v4, v2)
                                    v122 := GenS13{F0: "", F1: (^(fi66)), F2: 1}
                                    v1[fi67] = 32767
                                    var v123 uint64 = v88
                                    v2 = v2
                                    _ = v122
                                    _ = v123
                                    _ = v68
                                    _ = v69
                                    _ = v87
                                    _ = v88
                                } else {
                                    v124 := []int{0, fi67, fi67}
                                    fmt.Println(v2, v2)
                                    v125 := GenS8{F0: v62, F1: 242}
                                    _ = v124
                                    _ = v125
                                }
                                var v126 uint32 = (v2 + (v2 % (((v2 / ((v2) | 1))) | 1)))
                                v2 = (^(v2))
                                v126 = 0
                                if (true || true) {
                                    var v127 uint64 = 1
                                    fmt.Println(v2, fi66)
                                    switch v65 {
                                    case "":
                                        v128 := GenS14{F0: fi67}
                                        _ = v128
                                    default:
                                        switch ("hello" + ((((((v65 + "foo bar") + ((v65 + v65) + v65)) + v65) + ((v65 + "") + ("hello" + v65))) + (v65 + ((v65 + v65) + ((("\n\t" + "") + ((v65 + "hello") + (v65 + ((v65 + v65) + v65)))) + (v65 + v65))))) + (("" + "\n\t") + ""))) {
                                        case "\n\t":
                                            v2 = 2147483647
                                            if (v62 != genFn0("\n\t")) {
                                                v62 = v62
                                                var v129 uint = v62
                                                _ = v129
                                            }
                                        case "foo bar":
                                            v130 := GenS15{F0: 128, F1: (v127 % ((v127) | 1)), F2: (v127 * 912)}
                                            _ = v130
                                        }
                                    }
                                    _ = v127
                                }
                                v65 = ("\n\t" + (v65 + v65))
                                var v131 int64 = 9223372036854775807
                                if (!(false && false)) {
                                    v132 := []int{(fi67 / (((fi66 << ((v62) & 63))) | 1)), fi67, fi66, fi67}
                                    var v133 int16 = -32768
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v133, v65, v133)
                                    case "hello":
                                        switch v65 {
                                        case "a":
                                            var v134 float64 = 0.0
                                            _ = v134
                                        case "foo bar":
                                            v133 = int16(v126)
                                        }
                                    default:
                                        if len(v132) >= 2 {
                                            v132 = v132[1:len(v132)]
                                        }
                                    }
                                    if ((57 < -85) || (fi66 < (fi67 / ((764) | 1)))) {
                                        var v135 int16 = (^(v133))
                                        v136 := GenS16{F0: 42, F1: uint64(v62), F2: (v135 << (((v62 + (^(v62)))) & 63))}
                                        fmt.Println(fi66)
                                        _ = v135
                                        _ = v136
                                    } else {
                                        if false {
                                            v137 := GenS2{F0: v65, F1: v2, F2: 273}
                                            if false {
                                                if true {
                                                    delete(v5, "")
                                                    v62 = genFn0("hello")
                                                    v131 = v131
                                                }
                                                if false {
                                                    v3[903] = v65
                                                    if len(v132) >= 2 {
                                                        v132 = v132[1:len(v132)]
                                                    }
                                                }
                                                v62 -= 638
                                            }
                                            _ = v137
                                        }
                                    }
                                    if ((v2 << ((v62) & 63)) != 454) {
                                        v131 = (v131 << ((uint(v126)) & 63))
                                        v65 = v65
                                        fmt.Println(v4, v4)
                                    }
                                    var v138 int64 = v131
                                    if true {
                                        if (!(609 == 1)) {
                                            v138 = (v131 & (v138 ^ v138))
                                            fmt.Println(v4, fi67)
                                        } else {
                                            v132 = append(v132, 127)
                                            if (((v126 + (v126 % ((v126) | 1))) < (v2 / ((2) | 1))) && (!(false && (v133 >= v133)))) {
                                                v131 = v131
                                                v139 := GenS10{F0: (fi66 >> ((143) & 63)), F1: false, F2: 42}
                                                fmt.Println(v131, v126)
                                                _ = v139
                                            } else {
                                                var v140 int64 = -904
                                                _ = v140
                                            }
                                            v62 -= genFn0(v65)
                                        }
                                        v62 += (v62 | 18446744073709551615)
                                    } else {
                                        fmt.Println(v133, v62, fi66)
                                        v2 = uint32(v62)
                                    }
                                    _ = v132
                                    _ = v133
                                    _ = v138
                                } else {
                                    v4 = 1e10
                                    v141 := GenS13{F0: "", F1: fi67, F2: 42}
                                    _ = v141
                                }
                                v2 += (v2 ^ 4294967295)
                                v126 += v126
                                var v142 int = (fi66 / ((-847) | 1))
                                _ = fi67
                                _ = v126
                                _ = v131
                                _ = v142
                            }
                            v143 := GenS17{F0: 127}
                            v144 := GenS18{F0: v65, F1: uint64(v62)}
                            _ = fi66
                            _ = v143
                            _ = v144
                        }
                    }
                    _ = v64
                    _ = v65
                }
                v145 := []int{42}
                _ = v145
            }
            _ = len(v5)
        }
        if (0 == 0) {
            v2 += v2
        } else {
            if (!true) {
                v62 = (v62 * (v62 % (((v62 & (v62 >> (((v62 % (((^(v62))) | 1))) & 63)))) | 1)))
            }
            v5["hello"] = 1
            if true {
                if true {
                    if (65535 == 22) {
                        for fi146 := 0; fi146 < 1; fi146++ {
                            var v147 uint = (v62 << ((255) & 63))
                            _ = fi146
                            _ = v147
                        }
                    }
                    var v148 uint16 = 42
                    fmt.Println(v62, v148)
                    _ = v148
                } else {
                    var v149 uint = (v62 << (((v62 / ((v62) | 1))) & 63))
                    _ = v149
                }
                if false {
                    if true {
                        if (!false) {
                            v150 := GenS15{F0: 42, F1: uint64(v62), F2: 0}
                            var v151 int = 2147483647
                            for fi152 := 0; fi152 < 20; fi152++ {
                                var v153 uint = (v62 - v62)
                                v2 -= 0
                                _ = fi152
                                _ = v153
                            }
                            _ = v150
                            _ = v151
                        } else {
                            for fi154 := 0; fi154 < 18; fi154++ {
                                var v155 int16 = -874
                                v4 = (1e10 / v4)
                                if ((v62 >> ((v62) & 63)) == genFn0("foo bar")) {
                                    var v156 uint32 = 2
                                    v157 := GenS19{F0: genFn0((((("a" + "foo bar") + "hello") + "\n\t") + "a")), F1: true, F2: 65535}
                                    if false {
                                        var v158 uint = (v62 >> ((v62) & 63))
                                        v62 -= uint(v156)
                                        _ = v158
                                    } else {
                                        v159 := v5["hello"]
                                        _ = v159
                                    }
                                    _ = v156
                                    _ = v157
                                } else {
                                    if true {
                                        if (("" + "\n\t") <= "") {
                                            var v160 int = (fi154 * fi154)
                                            _ = v160
                                        } else {
                                            v161 := GenS11{F0: 471}
                                            v162 := GenS5{F0: (v62 % ((87) | 1)), F1: int64(v62), F2: 15}
                                            v2 += (v2 ^ v2)
                                            _ = v161
                                            _ = v162
                                        }
                                        v163 := GenS12{F0: (v2 % (((v2 - v2)) | 1))}
                                        var v164 string = "foo bar"
                                        _ = v163
                                        _ = v164
                                    } else {
                                        fmt.Println(v155, v62, v62)
                                        for fi165 := 0; fi165 < 6; fi165++ {
                                            if true {
                                                var v166 uint64 = 42
                                                _ = v166
                                            } else {
                                                if true {
                                                    v167 := GenS20{F0: "\n\t", F1: 255, F2: 4294967295}
                                                    if (false && false) {
                                                        _ = len(v5)
                                                        var v168 int16 = (^(v155))
                                                        _ = v168
                                                    } else {
                                                        if true {
                                                            var v169 uint = v62
                                                            switch fi154 {
                                                            case -9223372036854775808:
                                                                var v170 bool = (!(fi154 >= (^(fi154))))
                                                                _ = v170
                                                            }
                                                            _ = v169
                                                        }
                                                    }
                                                    _ = v167
                                                }
                                            }
                                            switch fi165 {
                                            case 32767:
                                                v4 -= float64(-2.25)
                                            case 2147483647:
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(fi154)
                                                default:
                                                    if (!(1 >= 100)) {
                                                        fmt.Println(fi165)
                                                        v155 = (v155 | v155)
                                                        var v171 string = "hello"
                                                        _ = v171
                                                    }
                                                }
                                            }
                                            if (v2 > 244) {
                                                v4 += (0.1 * -2.25)
                                            }
                                            _ = fi165
                                        }
                                        v172 := []int{fi154, (^(fi154)), fi154}
                                        _ = v172
                                    }
                                }
                                _ = fi154
                                _ = v155
                            }
                        }
                    }
                } else {
                    var v173 int64 = 42
                    switch 0 {
                    case 1:
                        var v174 string = ((("" + ("foo bar" + "a")) + "") + "foo bar")
                        _ = v174
                    default:
                        for fi175 := 0; fi175 < 8; fi175++ {
                            if (!false) {
                                var v176 int64 = (v173 / ((v173) | 1))
                                for fi177 := 0; fi177 < 12; fi177++ {
                                    v176 = v176
                                    _ = fi177
                                }
                                _ = v176
                            }
                            v178 := []int{}
                            _ = fi175
                            _ = v178
                        }
                    }
                    v62 += v62
                    _ = v173
                }
                v179 := GenS17{F0: 127}
                _ = v179
            }
        }
        _ = v62
    }
    fmt.Println(v2)
    v180 := []int{32768}
    var v181 float32 = -2.25
    var v182 int16 = 1000
    var v183 uint64 = 918
    v182 -= -32768
    delete(v5, "\n\t")
    v182 += (v182 % ((42) | 1))
    if true {
        switch 1 {
        case -1:
            for fi184 := 0; fi184 < 9; fi184++ {
                v182 += v182
                var v185 float32 = (1e10 + v181)
                v5[("a" + "a")] = (-(fi184))
                _ = fi184
                _ = v185
            }
        case -700:
            var v186 float32 = -2.25
            _ = v186
        case 46:
            if true {
                switch ("hello" + "\n\t") {
                case "foo bar":
                    switch -304 {
                    case 4294967295:
                        var v187 bool = true
                        _ = v187
                    case -1:
                        for fi188 := 0; fi188 < 7; fi188++ {
                            v181 = 1.5
                            if len(v180) > 0 {
                                v180[0] = fi188
                            }
                            if (!((v183 + (v183 | v183)) < v183)) {
                                for fi189 := 0; fi189 < 14; fi189++ {
                                    v4 = v4
                                    v190 := GenS16{F0: 42, F1: v183, F2: v182}
                                    _ = fi189
                                    _ = v190
                                }
                            } else {
                                var v191 int32 = 323
                                _ = v191
                            }
                            _ = fi188
                        }
                    case 127:
                        v192 := GenS21{F0: true, F1: 1, F2: (^(v182))}
                        _ = v192
                    }
                case "\n\t":
                    v193 := GenS15{F0: 2, F1: (^(v183)), F2: v183}
                    _ = v193
                case "":
                    fmt.Println(v181)
                }
                var v194 uint32 = v2
                _ = v194
            } else {
                for fi195 := 0; fi195 < 19; fi195++ {
                    v181 = v181
                    var v196 uint64 = (v183 >> ((80) & 63))
                    switch (^(fi195)) {
                    case 1:
                        v197 := GenS22{F0: 2}
                        _ = v197
                    case 4294967295:
                        for fi198 := 0; fi198 < 4; fi198++ {
                            v183 = 266
                            _ = len(v180)
                            v2 = (^(v2))
                            _ = fi198
                        }
                    case -658:
                        switch fi195 {
                        case 20:
                            v181 += float32(3.14159)
                        case 32767:
                            v199 := GenS23{F0: (fi195 & (fi195 | 9223372036854775807))}
                            _ = v199
                        default:
                            v200 := v3[fi195]
                            _ = v200
                        }
                    default:
                        for fi201 := 0; fi201 < 9; fi201++ {
                            var v202 float64 = ((0.0 - float64(((((v181 + ((v181 * 1.5) / v181)) - v181) / v181) + v181))) / 1e10)
                            _ = fi201
                            _ = v202
                        }
                    }
                    _ = fi195
                    _ = v196
                }
                if (true || true) {
                    var v203 bool = false
                    if true {
                        var v204 uint = genFn0("hello")
                        fmt.Println(v204)
                        _ = v204
                    }
                    _ = v203
                }
                for fi205 := 0; fi205 < 4; fi205++ {
                    v182 = 1
                    fmt.Println(fi205, v181, v4)
                    fmt.Println(v182, v182, v181)
                    _ = fi205
                }
            }
        }
        v4 = (v4 + v4)
        for fi206 := 0; fi206 < 14; fi206++ {
            v207 := GenS24{F0: (((42 >= -128) && true) && true), F1: -105}
            switch "a" {
            case "":
                for fi208 := 0; fi208 < 17; fi208++ {
                    v180 = append(v180, (fi208 / ((int(v183)) | 1)))
                    var v209 int32 = 273
                    if (195 >= 65535) {
                        if true {
                            v182 = (^(v182))
                        } else {
                            v4 += ((v4 - (1.5 - (v4 / v4))) / (v4 / (v4 * v4)))
                        }
                    }
                    _ = fi208
                    _ = v209
                }
            case "a":
                for fi210 := 0; fi210 < 7; fi210++ {
                    switch "\n\t" {
                    case "":
                        v211 := v1[-900]
                        _ = v211
                    case "\n\t":
                        if (!((65535 > 18446744073709551615) && (fi210 == (^(fi206))))) {
                            v4 -= float64(((float32(v4) - (v181 / (float32((((v4 + v4) / (v4 * float64((((v181 / v181) / v181) * (((0.0 / (v181 / ((v181 - v181) + v181))) / v181) - (((3.14159 + float32(((1e10 + v4) / float64(v181)))) * (1e10 + 1e10)) * (float32((v4 + (v4 - 0.1))) - 1e10))))))) * v4)) * -2.25))) * 1e10))
                            fmt.Println(v182, v2)
                        } else {
                            var v212 uint64 = 508
                            fmt.Println(v4, v182, v212)
                            var v213 uint8 = 255
                            _ = v212
                            _ = v213
                        }
                    case "a":
                        var v214 bool = false
                        _ = v214
                    }
                    switch -1 {
                    case 9223372036854775807:
                        if true {
                            v2 = 42
                        }
                    default:
                        var v215 uint8 = 42
                        _ = v215
                    }
                    var v216 uint = 65535
                    _ = fi210
                    _ = v216
                }
            case "foo bar":
                if len(v180) >= 2 {
                    v180 = v180[1:len(v180)]
                }
            }
            v217 := GenS6{F0: false, F1: genFn0("foo bar"), F2: 307}
            _ = fi206
            _ = v207
            _ = v217
        }
    }
    if true {
        var v218 float32 = (3.14159 - v181)
        if true {
            v180 = append(v180, -1)
        } else {
            var v219 uint64 = (^(v183))
            _ = v219
        }
        var v220 int16 = v182
        _ = v218
        _ = v220
    }
    if false {
        v182 -= 32767
        if false {
            v221 := GenS25{F0: 42}
            var v222 string = ("hello" + "hello")
            v182 = (v182 | 32767)
            _ = v221
            _ = v222
        } else {
            var v223 int16 = 42
            var v224 uint = 42
            if false {
                if false {
                    v1[910] = 128
                } else {
                    if false {
                        _ = len(v180)
                        v223 -= (v182 | (v182 - v223))
                    }
                    for fi225 := 0; fi225 < 10; fi225++ {
                        for fi226 := 0; fi226 < 17; fi226++ {
                            v2 = (v2 * v2)
                            fmt.Println(v223)
                            if (v224 <= genFn0("hello")) {
                                switch ("hello" + ("\n\t" + ("hello" + "foo bar"))) {
                                case "foo bar":
                                    v4 -= 0.1
                                case "hello":
                                    var v227 bool = (!true)
                                    _ = v227
                                }
                            }
                            _ = fi226
                        }
                        _ = fi225
                    }
                }
            } else {
                if (519 <= 241) {
                    for fi228 := 0; fi228 < 4; fi228++ {
                        if len(v180) > 0 {
                            v180[0] = 318
                        }
                        v229 := GenS26{F0: (v183 << ((genFn0(("hello" + "\n\t"))) & 63))}
                        v181 -= (0.1 + (v181 + 0.0))
                        _ = fi228
                        _ = v229
                    }
                    fmt.Println(v182, v2, v182)
                } else {
                    var v230 int = 4294967295
                    _ = v230
                }
            }
            _ = v223
            _ = v224
        }
        var v231 uint32 = v2
        _ = v231
    }
    v180 = append(v180, 42)
    v232 := GenS27{F0: (v183 ^ (^(v183))), F1: (v182 * v182), F2: (128 > 1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v180:", len(v180), v180)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("v183:", v183)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v180
    _ = v181
    _ = v182
    _ = v183
    _ = v2
    _ = v232
    _ = v3
    _ = v4
    _ = v5
}
