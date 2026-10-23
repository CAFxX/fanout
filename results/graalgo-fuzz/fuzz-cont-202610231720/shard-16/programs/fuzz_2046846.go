package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 bool; F1 int; F2 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 int16; F2 int }
type GenS11 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS12 struct { F0 int16; F1 int64; F2 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int; F1 int64; F2 uint }
type GenS15 struct { F0 string; F1 uint }
type GenS16 struct { F0 uint8; F1 bool }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int; F1 int8; F2 int8 }
type GenS19 struct { F0 int; F1 uint; F2 bool }

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 string = "hello"
    fmt.Println(v1, v5)
    for fi6 := 0; fi6 < 1; fi6++ {
        v1 = v1
        if (2147483648 == 877) {
            fmt.Println(v5, v1, fi6)
            v1 = (int8(fi6) < -1)
            v5 = v5
        }
        _ = fi6
    }
    switch 128 {
    case -9223372036854775808:
        v7 := []int{4294967295, 127, 32768, -1}
        _ = v7
    case 128:
        v8 := GenS1{F0: 42}
        _ = v8
    default:
        switch v5 {
        case "\n\t":
            for fi9 := 0; fi9 < 19; fi9++ {
                delete(v3, (v5 + v5))
                _ = fi9
            }
        }
    }
    switch 42 {
    case -2147483648:
        _ = len(v2)
    case -1:
        v10 := GenS1{F0: 127}
        _ = v10
    case 9223372036854775807:
        v11 := GenS1{F0: -29}
        _ = v11
    default:
        if ((159 <= 130) && v1) {
            for fi12 := 0; fi12 < 15; fi12++ {
                fmt.Println(v1, v1, fi12)
                for fi13 := 0; fi13 < 20; fi13++ {
                    if (!(808 >= 65535)) {
                        v14 := GenS1{F0: 42}
                        var v15 uint8 = 42
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint8 = 1
                        v4[(-(fi12))] = 42
                        _ = v16
                    }
                    if true {
                        v17 := GenS2{F0: -121}
                        var v18 float64 = float64(float32(1e10))
                        v18 = 0.1
                        _ = v17
                        _ = v18
                    } else {
                        v5 = (v5 + v5)
                        v5 = "hello"
                        var v19 uint64 = 448
                        _ = v19
                    }
                    _ = fi13
                }
                v20 := []int{(fi12 >> ((18446744073709551615) & 63)), (-(fi12))}
                _ = fi12
                _ = v20
            }
            v21 := v3[v5]
            _ = v21
        } else {
            if true {
                v5 = v5
            } else {
                if (2147483648 < 417) {
                    v22 := v3[v5]
                    if (v1 || v22) {
                        fmt.Println(v5, v1, v5)
                    } else {
                        if (65535 == 32767) {
                            for fi23 := 0; fi23 < 17; fi23++ {
                                var v24 float32 = float32(0.0)
                                v25 := GenS2{F0: 42}
                                for fi26 := 0; fi26 < 20; fi26++ {
                                    v24 += v24
                                    if (89 != 192) {
                                        var v27 int = fi26
                                        var v28 float64 = float64(3.14159)
                                        _ = v27
                                        _ = v28
                                    }
                                    fmt.Println(v5)
                                    _ = fi26
                                }
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            var v29 string = v5
                            _ = v29
                        }
                        var v30 uint = 65535
                        v30 += v30
                        _ = v30
                    }
                    _ = len(v3)
                    _ = v22
                }
                var v31 int64 = -9223372036854775808
                for fi32 := 0; fi32 < 7; fi32++ {
                    v33 := GenS2{F0: -1}
                    var v34 uint16 = 28
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = v31
            }
            if v1 {
                for fi35 := 0; fi35 < 5; fi35++ {
                    v1 = (v1 || v1)
                    _ = fi35
                }
                fmt.Println(v5, v1)
                if v1 {
                    if v1 {
                        if false {
                            fmt.Println(v1, v1)
                            v36 := GenS2{F0: 100}
                            _ = v36
                        }
                        var v37 uint64 = 4294967295
                        _ = v37
                    } else {
                        var v38 int32 = 0
                        v39 := GenS3{F0: 255, F1: 953}
                        v40 := GenS4{F0: 65535, F1: 100}
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    v1 = v1
                }
            } else {
                v5 = "hello"
            }
        }
    }
    var v41 int16 = 42
    var v42 int16 = 1000
    var v43 int8 = 0
    var v44 float32 = 1e10
    if (v5 >= "hello") {
        v45 := v3[((v5 + (v5 + ("" + ((v5 + v5) + (((v5 + ("" + ("foo bar" + (v5 + (v5 + "a"))))) + v5) + v5))))) + v5)]
        _ = v45
    }
    for fi46 := 0; fi46 < 6; fi46++ {
        v47 := []int{}
        _ = fi46
        _ = v47
    }
    v48 := v3[v5]
    v44 = v44
    v49 := []int{4294967295, -9223372036854775808, 32768}
    v1 = true
    var v50 uint32 = 737
    fmt.Println(v50, v1)
    if (v1 || ("" >= (v5 + v5))) {
        switch 931 {
        case 100:
            if (!(v43 <= (v43 >> ((2) & 63)))) {
                var v51 int64 = 42
                _ = v51
            }
        default:
            var v52 uint8 = 0
            _ = v52
        }
        switch 128 {
        case 26:
            for fi53 := 0; fi53 < 13; fi53++ {
                if v1 {
                    var v54 string = v5
                    var v55 int64 = 0
                    _ = v54
                    _ = v55
                }
                _ = fi53
            }
        case 42:
            if (true || ((2 != 18446744073709551615) || true)) {
                if ((float64(v44) <= 1.5) || v1) {
                    switch ("a" + "hello") {
                    case "foo bar":
                        if v1 {
                            for fi56 := 0; fi56 < 15; fi56++ {
                                if v48 {
                                    for fi57 := 0; fi57 < 15; fi57++ {
                                        var v58 int = fi56
                                        v41 = v42
                                        v59 := GenS1{F0: 1}
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = len(v49)
                                    v60 := GenS3{F0: 127, F1: 4294967295}
                                    _ = v60
                                } else {
                                    var v61 bool = (!(1 != 2))
                                    _ = v61
                                }
                                for fi62 := 0; fi62 < 6; fi62++ {
                                    v63 := GenS5{F0: 1, F1: 255}
                                    _ = fi62
                                    _ = v63
                                }
                                _ = fi56
                            }
                        }
                    case "\n\t":
                        fmt.Println(v44, v1, v5)
                    case "a":
                        if ((^(v50)) > (v50 - v50)) {
                            switch 399 {
                            case 2147483647:
                                v64 := GenS6{F0: 255, F1: 9223372036854775807}
                                _ = v64
                            case 32767:
                                var v65 uint32 = v50
                                _ = v65
                            default:
                                fmt.Println(v43)
                            }
                        }
                    default:
                        if (v48 || ((-(v43)) > -128)) {
                            v41 += v41
                            for fi66 := 0; fi66 < 16; fi66++ {
                                if (v1 && v48) {
                                    for fi67 := 0; fi67 < 20; fi67++ {
                                        v41 = v42
                                        if (937 < 127) {
                                            var v68 string = (v5 + "")
                                            _ = v68
                                        } else {
                                            var v69 uint64 = 42
                                            if (!false) {
                                                v70 := GenS7{F0: (^(fi66)), F1: fi67}
                                                switch 132 {
                                                case 244:
                                                    var v71 string = (v5 + ("hello" + v5))
                                                    _ = v71
                                                case 128:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if ((v5 + v5) == v5) {
                                                            v72 := GenS8{F0: (float64(((v44 + (0.1 + (float32(0.0) * (((v44 / v44) * 0.0) + 3.14159)))) - (3.14159 / 1.0))) > 0.1), F1: (fi66 * (fi67 / (((fi67 >> ((255) & 63))) | 1))), F2: -1}
                                                            var v73 uint64 = (^(v69))
                                                            if (-118 == -1) {
                                                                var v74 string = "a"
                                                                _ = v74
                                                            } else {
                                                                if len(v49) >= 2 {
                                                                    v49 = v49[1:len(v49)]
                                                                }
                                                            }
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            v75 := GenS8{F0: ((!false) && true), F1: fi66, F2: 0}
                                                            switch (fi66 * fi66) {
                                                            case 264:
                                                                if true {
                                                                    v1 = (!((^(v69)) >= (v69 % (((^(v69))) | 1))))
                                                                } else {
                                                                    fmt.Println(v5, v1, v5)
                                                                    if v48 {
                                                                        fmt.Println(v1, v1, v41)
                                                                        fmt.Println(v48, v69, v41)
                                                                    } else {
                                                                        v1 = false
                                                                        var v76 int16 = v41
                                                                        _ = v76
                                                                    }
                                                                }
                                                            case 0:
                                                                v1 = (v41 <= v41)
                                                            default:
                                                                var v77 int = (-(fi66))
                                                                if true {
                                                                    v78 := GenS9{F0: v43}
                                                                    v79 := GenS10{F0: (v5 + (v5 + "hello")), F1: v42, F2: 128}
                                                                    v2[32768] = ((v5 + "hello") + v5)
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                                var v80 uint32 = 1
                                                                if (215 > 65535) {
                                                                    switch (v77 / ((fi67) | 1)) {
                                                                    case -2147483648:
                                                                        fmt.Println(v50, v77)
                                                                    case 32768:
                                                                        var v81 uint = 610
                                                                        switch (v77 / ((-1) | 1)) {
                                                                        case 2147483647:
                                                                            var v82 uint16 = 65535
                                                                            _ = v82
                                                                        }
                                                                        if (v80 == 922) {
                                                                            switch v5 {
                                                                            case "a":
                                                                                v77 += (-(v77))
                                                                            }
                                                                            v83 := GenS4{F0: 0, F1: (v43 / (((^(v43))) | 1))}
                                                                            if (v81 >= (v81 << ((1) & 63))) {
                                                                                v42 = 0
                                                                                switch (v5 + v5) {
                                                                                case "hello":
                                                                                    var v84 int16 = v42
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                            _ = v83
                                                                        }
                                                                        _ = v81
                                                                    case -570:
                                                                        v85 := GenS1{F0: (v43 * 0)}
                                                                        _ = v85
                                                                    }
                                                                }
                                                                _ = v77
                                                                _ = v80
                                                            }
                                                            _ = v75
                                                        }
                                                    case "hello":
                                                        switch (fi66 - fi67) {
                                                        case 404:
                                                            fmt.Println(v1, v50)
                                                        case 195:
                                                            if (true || (242 != 27)) {
                                                                var v86 uint64 = 351
                                                                v49 = append(v49, (-(fi66)))
                                                                v50 -= v50
                                                                v87 := GenS10{F0: (v5 + (v5 + "foo bar")), F1: 485, F2: (-(fi67))}
                                                                _ = v86
                                                                _ = v87
                                                            } else {
                                                                v3["hello"] = (int64(v50) != -45)
                                                            }
                                                        case -1:
                                                            fmt.Println(v41)
                                                        }
                                                    default:
                                                        if v1 {
                                                            v3[v5] = v1
                                                            var v88 uint = 2
                                                            _ = v88
                                                        }
                                                    }
                                                case 4294967295:
                                                    fmt.Println(v50, fi66, v43)
                                                    if v1 {
                                                        v48 = (v48 || (true || (829 < 326)))
                                                        v44 += 0.0
                                                    } else {
                                                        switch fi67 {
                                                        case -705:
                                                            v89 := GenS11{F0: 100, F1: 144, F2: 28}
                                                            _ = v89
                                                        case 4294967295:
                                                            var v90 int32 = -2147483648
                                                            _ = len(v49)
                                                            v49 = append(v49, fi67)
                                                            _ = v90
                                                        case -2147483648:
                                                            v48 = ((v44 == v44) || v48)
                                                        }
                                                        if v1 {
                                                            v69 = uint64(v50)
                                                        } else {
                                                            if (179 == 127) {
                                                                if ((v44 - (v44 + (v44 / v44))) > (1e10 + (3.14159 * ((v44 - v44) + float32((0.1 * 1e10)))))) {
                                                                    var v91 uint8 = 2
                                                                    var v92 int16 = v42
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        var v93 int32 = 42
                                                        _ = v93
                                                    }
                                                    v94 := GenS9{F0: (v43 | int8(fi67))}
                                                    _ = v94
                                                default:
                                                    switch ((v5 + v5) + v5) {
                                                    case "foo bar":
                                                        if v48 {
                                                            switch fi67 {
                                                            case -127:
                                                                var v95 uint8 = 127
                                                                _ = v95
                                                            case -9223372036854775808:
                                                                var v96 int8 = v43
                                                                v97 := GenS6{F0: 828, F1: fi66}
                                                                v69 = (^(v69))
                                                                _ = v96
                                                                _ = v97
                                                            case -18:
                                                                fmt.Println(v43, fi66, v50)
                                                            }
                                                            fmt.Println(v42, fi67, v69)
                                                            if v48 {
                                                                var v98 uint16 = 2
                                                                _ = v98
                                                            } else {
                                                                if v48 {
                                                                    _ = len(v49)
                                                                    v3["a"] = true
                                                                    if v48 {
                                                                        fmt.Println(v5, v48)
                                                                    }
                                                                }
                                                                _ = len(v3)
                                                                if v1 {
                                                                    var v99 int64 = -1
                                                                    if true {
                                                                        switch v5 {
                                                                        case "foo bar":
                                                                            if true {
                                                                                var v100 int32 = 2147483647
                                                                                if v48 {
                                                                                    v101 := GenS8{F0: true, F1: fi66, F2: v100}
                                                                                    if (!v1) {
                                                                                        if (!(((v5 + "") + v5) != (v5 + ""))) {
                                                                                            if len(v49) >= 2 {
                                                                                                v49 = v49[1:len(v49)]
                                                                                            }
                                                                                            var v102 int8 = (v43 ^ 1)
                                                                                            var v103 float64 = float64(0.1)
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        }
                                                                                        v104 := GenS12{F0: v41, F1: (^(v99)), F2: (((v48 || (!(((v50 * v50) <= v50) || (42 > (v99 / ((1) | 1)))))) || (v42 != (v41 >> ((613) & 63)))) || (!(0 < 518)))}
                                                                                        v48 = (false || ((!v1) || v48))
                                                                                        _ = v104
                                                                                    } else {
                                                                                        v105 := GenS2{F0: v43}
                                                                                        fmt.Println(v44)
                                                                                        _ = v105
                                                                                    }
                                                                                    if false {
                                                                                        if false {
                                                                                            v106 := v4[fi66]
                                                                                            fmt.Println(v43)
                                                                                            v1 = v48
                                                                                            _ = v106
                                                                                        } else {
                                                                                            switch fi67 {
                                                                                            case -2147483648:
                                                                                                var v107 string = v5
                                                                                                _ = v107
                                                                                            case -1:
                                                                                                v5 = v5
                                                                                            case 736:
                                                                                                switch ((v5 + (v5 + v5)) + (v5 + (v5 + v5))) {
                                                                                                case "hello":
                                                                                                    v41 += -732
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v44, v69, v50)
                                                                                                }
                                                                                                v44 = float32(((1.5 * 0.1) + (-2.25 - 1e10)))
                                                                                            }
                                                                                            v44 = float32(1.5)
                                                                                            var v108 string = ("foo bar" + "")
                                                                                            _ = v108
                                                                                        }
                                                                                        var v109 uint64 = v69
                                                                                        var v110 uint64 = v109
                                                                                        if (v44 == (v44 / v44)) {
                                                                                            v69 -= v69
                                                                                            if v48 {
                                                                                                switch v5 {
                                                                                                case "":
                                                                                                    v4[(fi67 | fi66)] = fi67
                                                                                                case "a":
                                                                                                    v110 += (v110 >> ((42) & 63))
                                                                                                default:
                                                                                                    fmt.Println(v69, v42)
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                            }
                                                                                            var v111 int = (fi66 >> ((943) & 63))
                                                                                            _ = v111
                                                                                        } else {
                                                                                            v69 = (v109 & 415)
                                                                                            switch v5 {
                                                                                            case "":
                                                                                                switch (fi66 % ((2147483647) | 1)) {
                                                                                                case -2147483648:
                                                                                                    v112 := GenS13{F0: 100}
                                                                                                    _ = v112
                                                                                                case -239:
                                                                                                    var v113 int = 9223372036854775807
                                                                                                    _ = v113
                                                                                                default:
                                                                                                    v42 = v42
                                                                                                }
                                                                                            case "hello":
                                                                                                v109 = v69
                                                                                            case "a":
                                                                                                var v114 int8 = 127
                                                                                                _ = v114
                                                                                            default:
                                                                                                var v115 int8 = v43
                                                                                                _ = v115
                                                                                            }
                                                                                        }
                                                                                        var v116 int32 = 42
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                        _ = v116
                                                                                    }
                                                                                    v50 = v50
                                                                                    _ = v101
                                                                                } else {
                                                                                    var v117 uint16 = 65535
                                                                                    var v118 bool = (v41 >= v42)
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                }
                                                                                v119 := GenS7{F0: fi66, F1: -452}
                                                                                var v120 int64 = v99
                                                                                _ = v100
                                                                                _ = v119
                                                                                _ = v120
                                                                            }
                                                                            var v121 uint64 = v69
                                                                            _ = v121
                                                                        case "":
                                                                            var v122 bool = v48
                                                                            _ = v122
                                                                        default:
                                                                            if v1 {
                                                                                if ((v5 + v5) == "\n\t") {
                                                                                    v44 -= v44
                                                                                }
                                                                                fmt.Println(v44, v44)
                                                                                var v123 uint = 42
                                                                                _ = v123
                                                                            }
                                                                        }
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v124 uint = 29
                                                                            _ = v124
                                                                        case "foo bar":
                                                                            var v125 string = "foo bar"
                                                                            _ = v125
                                                                        case "hello":
                                                                            v41 -= v41
                                                                        default:
                                                                            v41 = (v42 ^ (v41 ^ v42))
                                                                        }
                                                                    }
                                                                    _ = v99
                                                                }
                                                            }
                                                            switch v5 {
                                                            case "\n\t":
                                                                var v126 uint8 = 0
                                                                _ = v126
                                                            case "hello":
                                                                if (0 <= 1) {
                                                                    fmt.Println(v69, v43, v41)
                                                                }
                                                            }
                                                            var v127 int64 = 0
                                                            _ = v127
                                                        }
                                                    }
                                                }
                                                _ = v70
                                            } else {
                                                v128 := GenS4{F0: 2, F1: v43}
                                                v44 = 0.1
                                                _ = v128
                                            }
                                            v129 := GenS14{F0: 0, F1: 1, F2: 18446744073709551615}
                                            _ = v129
                                            _ = v69
                                        }
                                        _ = fi67
                                    }
                                    var v130 uint32 = (v50 / ((v50) | 1))
                                    delete(v2, (fi66 / ((fi66) | 1)))
                                    _ = v130
                                }
                                for fi131 := 0; fi131 < 1; fi131++ {
                                    if (1 > -2147483648) {
                                        var v132 float64 = 3.14159
                                        _ = v132
                                    } else {
                                        v41 = -1
                                        if false {
                                            fmt.Println(fi66, v44)
                                            fmt.Println(v1, v50, v5)
                                        } else {
                                            v50 = 2147483647
                                            if len(v49) > 0 {
                                                v49[0] = (fi131 & fi131)
                                            }
                                            fmt.Println(v42, v50)
                                            var v133 uint32 = (^(v50))
                                            _ = v133
                                        }
                                        v5 = "\n\t"
                                        v134 := GenS1{F0: v43}
                                        if (v48 || (42 < 42)) {
                                            v42 = v42
                                        }
                                        _ = v134
                                    }
                                    fmt.Println(v1, v48, v5)
                                    v41 += v42
                                    _ = fi131
                                }
                                _ = fi66
                            }
                        } else {
                            if false {
                                for fi135 := 0; fi135 < 6; fi135++ {
                                    v136 := GenS11{F0: v43, F1: v50, F2: 42}
                                    _ = fi135
                                    _ = v136
                                }
                                if v48 {
                                    v137 := GenS15{F0: ("a" + v5), F1: 255}
                                    v42 -= 855
                                    _ = v137
                                }
                                for fi138 := 0; fi138 < 11; fi138++ {
                                    v42 = 32767
                                    _ = fi138
                                }
                            } else {
                                if (2 == v50) {
                                    var v139 uint8 = 42
                                    v4[32768] = 32768
                                    v140 := GenS16{F0: (^(v139)), F1: v1}
                                    _ = v139
                                    _ = v140
                                }
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                            }
                            if (v44 != (0.0 - 0.0)) {
                                if v1 {
                                    switch "a" {
                                    case "foo bar":
                                        v43 = v43
                                    default:
                                        var v141 int8 = 127
                                        _ = v141
                                    }
                                    v5 = ""
                                } else {
                                    var v142 int = 32768
                                    var v143 int32 = 0
                                    _ = v142
                                    _ = v143
                                }
                                if v1 {
                                    _ = len(v3)
                                }
                                v144 := GenS17{F0: -476}
                                _ = v144
                            }
                            v145 := GenS12{F0: 42, F1: 295, F2: false}
                            _ = v145
                        }
                    }
                }
            } else {
                v146 := GenS6{F0: 18446744073709551615, F1: 42}
                v147 := GenS18{F0: 2147483647, F1: (^(v43)), F2: (-(v43))}
                v148 := GenS13{F0: (v43 & -1)}
                _ = v146
                _ = v147
                _ = v148
            }
        }
        fmt.Println(v43)
    } else {
        var v149 float64 = (1e10 * -2.25)
        v42 += v41
        v150 := GenS19{F0: 390, F1: 255, F2: true}
        _ = v149
        _ = v150
    }
    for fi151 := 0; fi151 < 5; fi151++ {
        fmt.Println(v50)
        var v152 uint16 = 2
        var v153 string = ("a" + (v5 + v5))
        _ = fi151
        _ = v152
        _ = v153
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
