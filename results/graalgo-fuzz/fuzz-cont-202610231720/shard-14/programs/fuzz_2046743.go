package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 uint8; F1 int64; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8; F1 int16; F2 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 string; F1 uint8; F2 uint }
type GenS18 struct { F0 uint16; F1 bool; F2 int32 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 uint; F1 int64; F2 uint }
type GenS22 struct { F0 int16; F1 int8 }
type GenS23 struct { F0 bool }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 bool }
type GenS26 struct { F0 bool; F1 uint8 }
type GenS27 struct { F0 int32; F1 bool }
type GenS28 struct { F0 string; F1 uint16 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    var v4 int16 = p0
    _ = p0
    _ = v4
    return "a"
}
func genFn1(p0 uint) int8 {
    var v5 uint32 = 1
    for fi6 := 0; fi6 < 4; fi6++ {
        switch ("hello" + ("\n\t" + ("a" + "hello"))) {
        case "\n\t":
            var v7 string = ("foo bar" + (("foo bar" + "foo bar") + "hello"))
            _ = v7
        }
        if true {
            var v8 int16 = int16(p0)
            p0 = 4294967295
            _ = v8
        } else {
            var v9 uint8 = 50
            v10 := GenS1{F0: -850, F1: 255, F2: 42}
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        p0 += (p0 % ((p0) | 1))
        v12 := GenS2{F0: p0, F1: (fi11 ^ fi11)}
        if true {
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 bool = (false && true)
                    v15 := GenS3{F0: -128, F1: 42, F2: ("foo bar" + ("foo bar" + ("" + (("" + "\n\t") + "foo bar"))))}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS3{F0: -101, F1: -1, F2: "foo bar"}
                v17 := GenS4{F0: 127, F1: -850, F2: -463}
                _ = v16
                _ = v17
            }
        }
        _ = fi11
        _ = v12
    }
    if false {
        var v18 bool = false
        _ = v18
    } else {
        v19 := []int{32767, -2147483648}
        var v20 float64 = 1.5
        fmt.Println(v20)
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v3[("hello" + (("\n\t" + "a") + (("foo bar" + "hello") + "\n\t")))] = 546
    var v21 int = 678
    v22 := GenS5{F0: 0}
    if false {
        v23 := GenS2{F0: 1, F1: -883}
        v24 := GenS5{F0: 40}
        _ = v23
        _ = v24
    } else {
        if (false && true) {
            var v25 uint32 = 65535
            if true {
                v26 := GenS6{F0: (v25 | 2147483647)}
                _ = v26
            } else {
                var v27 uint = 272
                _ = v27
            }
            _ = v25
        } else {
            if (!((((true && (520 < 922)) && true) || ((1e10 - ((1e10 - (3.14159 / ((-2.25 - (0.0 - 0.1)) + (((1.5 + (0.1 - -2.25)) - 1e10) / float64((0.1 / -2.25)))))) * ((float64(1e10) / 3.14159) + (1e10 * ((0.1 + (1.5 - -2.25)) / 0.1))))) >= 1e10)) || (!false))) {
                if false {
                    var v28 int = 0
                    var v29 uint8 = 2
                    _ = v28
                    _ = v29
                } else {
                    if (42 > 926) {
                        if (-2.25 != 0.0) {
                            fmt.Println(v1, v1)
                        }
                        fmt.Println(v21, v21)
                        v30 := GenS2{F0: 65535, F1: (v21 / ((v21) | 1))}
                        _ = v30
                    } else {
                        v21 -= (v21 + 0)
                    }
                    v21 = v21
                    for fi31 := 0; fi31 < 11; fi31++ {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            switch 42 {
                            case 4294967295:
                                fmt.Println(v1, fi31, v1)
                            case -9223372036854775808:
                                fmt.Println(fi32, fi32, v21)
                            }
                            v33 := []int{(-(v21))}
                            _ = fi32
                            _ = v33
                        }
                        _ = fi31
                    }
                }
                if (false || ((-2.25 * 1e10) < 3.14159)) {
                    if true {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v35 := []int{(fi34 | 9223372036854775807), fi34, (fi34 >> ((352) & 63)), (v21 ^ v21)}
                            var v36 float32 = 3.14159
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        v1 = -917
                    }
                    v21 += v21
                    v1 += (v1 ^ (-(v1)))
                } else {
                    switch "" {
                    case "foo bar":
                        for fi37 := 0; fi37 < 3; fi37++ {
                            v38 := []int{4294967295, 32767, (fi37 << ((1) & 63)), (fi37 + 1)}
                            v21 += v21
                            _ = fi37
                            _ = v38
                        }
                    case "hello":
                        v21 = (v21 / (((v21 >> ((967) & 63))) | 1))
                    case "a":
                        _ = len(v2)
                    default:
                        if true {
                            v1 -= 1000
                        } else {
                            var v39 uint16 = 225
                            _ = v39
                        }
                    }
                }
                var v40 bool = (2 == 4294967295)
                _ = v40
            } else {
                v41 := GenS4{F0: 80, F1: -1, F2: -860}
                var v42 int16 = 1000
                v42 = v42
                _ = v41
                _ = v42
            }
        }
        v43 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 806}
        switch "a" {
        case "\n\t":
            var v44 uint = 1
            _ = v44
        case "a":
            var v45 float64 = -2.25
            _ = v45
        case "":
            v46 := []int{(v21 | (v21 / ((-145) | 1))), (v21 | v21), (v21 | v21), v21}
            _ = v46
        default:
            var v47 int64 = 42
            _ = v47
        }
        _ = v43
    }
    fmt.Println(v1)
    for fi48 := 0; fi48 < 20; fi48++ {
        v49 := []int{(v21 % (((fi48 - (-(fi48)))) | 1))}
        v50 := GenS7{F0: 1}
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        var v52 bool = false
        v3["foo bar"] = 761
        _ = fi51
        _ = v52
    }
    switch "hello" {
    case "":
        var v53 uint8 = 127
        _ = v53
    }
    for fi54 := 0; fi54 < 7; fi54++ {
        if false {
            for fi55 := 0; fi55 < 17; fi55++ {
                v1 += v1
                var v56 int16 = (-(v1))
                v1 = (v56 + (v1 & 910))
                _ = fi55
                _ = v56
            }
            var v57 bool = true
            v21 -= (fi54 % ((-1) | 1))
            _ = v57
        }
        for fi58 := 0; fi58 < 4; fi58++ {
            switch (fi58 - (fi58 / (((fi54 / ((fi58) | 1))) | 1))) {
            case -924:
                v1 = v1
            case -497:
                if false {
                    fmt.Println(fi54)
                    delete(v3, "")
                } else {
                    var v59 string = ""
                    switch (fi58 - 32768) {
                    case 241:
                        fmt.Println(fi54, v59)
                    case 127:
                        if (true || (!(((((-2.25 - ((0.0 / 1e10) + 1e10)) / -2.25) * 3.14159) > 1.5) && true))) {
                            var v60 uint8 = 44
                            _ = v60
                        } else {
                            v21 = v21
                            var v61 int16 = (^(v1))
                            v21 -= (fi58 >> ((0) & 63))
                            _ = v61
                        }
                    default:
                        if false {
                            var v62 int16 = -1
                            _ = v62
                        } else {
                            switch 32767 {
                            case -9223372036854775808:
                                fmt.Println(v1)
                            case -768:
                                fmt.Println(fi58)
                            case 9223372036854775807:
                                v63 := []int{fi54, 32768, 32767}
                                _ = v63
                            }
                            v64 := GenS6{F0: 4294967295}
                            v59 = v59
                            v1 = -82
                            switch fi54 {
                            case -16:
                                fmt.Println(v59, v1)
                            case 1:
                                var v65 int64 = -9223372036854775808
                                _ = v65
                            case -1:
                                var v66 string = (v59 + v59)
                                v1 = v1
                                v2[308] = (fi54 >> ((255) & 63))
                                _ = v66
                            }
                            v59 = "\n\t"
                            fmt.Println(fi54)
                            _ = v64
                        }
                    }
                    switch ("foo bar" + (("foo bar" + v59) + "a")) {
                    case "hello":
                        switch v21 {
                        case 388:
                            var v67 uint32 = 1
                            _ = v67
                        case -2147483648:
                            var v68 bool = false
                            _ = v68
                        case -183:
                            var v69 float64 = (1e10 - ((3.14159 + ((((-2.25 + 0.1) / ((1.5 - (1e10 * 1e10)) + 3.14159)) + ((0.0 * 3.14159) / (1.5 * (-2.25 / 0.1)))) / float64(3.14159))) + 0.1))
                            _ = v69
                        default:
                            switch v21 {
                            case 4294967295:
                                var v70 int = 575
                                _ = v70
                            }
                        }
                    case "":
                        switch v59 {
                        case "":
                            if (true || false) {
                                v71 := []int{}
                                v1 = (v1 << ((65535) & 63))
                                _ = v71
                            }
                        case "hello":
                            if true {
                                v1 = v1
                            } else {
                                v1 -= (v1 / ((9) | 1))
                            }
                        case "\n\t":
                            if false {
                                fmt.Println(v21)
                                v72 := []int{}
                                v1 -= -1
                                _ = v72
                            } else {
                                switch v59 {
                                case "hello":
                                    var v73 float64 = -2.25
                                    _ = v73
                                default:
                                    v21 = -466
                                }
                            }
                        default:
                            var v74 int64 = 848
                            _ = v74
                        }
                    case "a":
                        fmt.Println(v59, v59)
                    default:
                        v75 := []int{fi54}
                        _ = v75
                    }
                    if true {
                        if true {
                            switch fi54 {
                            case -219:
                                if ((-72 <= 106) && (int64(fi54) >= 9223372036854775807)) {
                                    var v76 int16 = v1
                                    _ = v76
                                } else {
                                    v77 := []int{(fi58 << ((65535) & 63)), (v21 ^ (fi54 % ((fi54) | 1)))}
                                    v78 := GenS1{F0: -1, F1: 400, F2: 918}
                                    _ = v77
                                    _ = v78
                                }
                            }
                            var v79 int = 4294967295
                            _ = v79
                        }
                        var v80 float32 = float32((1e10 - 1e10))
                        _ = v80
                    } else {
                        switch (fi54 * 32767) {
                        case 1:
                            var v81 uint = 1
                            if false {
                                v82 := GenS2{F0: v81, F1: fi58}
                                switch (-(fi54)) {
                                case 42:
                                    if true {
                                        switch (v59 + "foo bar") {
                                        case "a":
                                            switch v59 {
                                            case "foo bar":
                                                switch "hello" {
                                                case "hello":
                                                    v81 = (^(v81))
                                                case "foo bar":
                                                    v1 += 1000
                                                }
                                            case "hello":
                                                if (false || false) {
                                                    var v83 int16 = 0
                                                    if false {
                                                        switch (^(v21)) {
                                                        case -2147483648:
                                                            var v84 int16 = v83
                                                            _ = v84
                                                        case 9223372036854775807:
                                                            var v85 int32 = 0
                                                            v86 := []int{-2147483648, int(v85)}
                                                            v87 := GenS8{F0: ((v59 + "\n\t") + v59), F1: true}
                                                            _ = v85
                                                            _ = v86
                                                            _ = v87
                                                        default:
                                                            if (((false && false) && true) || (42 <= 127)) {
                                                                v88 := GenS9{F0: 53}
                                                                fmt.Println(v59, fi58)
                                                                _ = v88
                                                            } else {
                                                                if (true && (18446744073709551615 == (v81 + (v81 + (^(v81)))))) {
                                                                    var v89 int16 = (v1 % ((v83) | 1))
                                                                    if false {
                                                                        v90 := GenS10{F0: 37, F1: (v83 << ((0) & 63)), F2: (fi54 - (^(fi54)))}
                                                                        v91 := GenS2{F0: v81, F1: -421}
                                                                        var v92 uint64 = 46
                                                                        if (((false && (42 < 190)) && (!(true || true))) || (42 >= 0)) {
                                                                            fmt.Println(v92, v81)
                                                                            if (float32((1.5 - 1e10)) <= 0.1) {
                                                                                v93 := []int{-341}
                                                                                _ = v93
                                                                            }
                                                                            if ((v1 % (((v83 | (v89 * (^(v83))))) | 1)) < (v1 & v83)) {
                                                                                v94 := GenS1{F0: 0, F1: 33, F2: v92}
                                                                                v81 += 877
                                                                                _ = v94
                                                                            } else {
                                                                                v95 := GenS11{F0: 1}
                                                                                v92 = v92
                                                                                var v96 uint8 = 0
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                        } else {
                                                                            var v97 int8 = 1
                                                                            fmt.Println(v21)
                                                                            v98 := []int{}
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                        v92 = (^(v92))
                                                                        _ = v90
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                    _ = v89
                                                                } else {
                                                                    fmt.Println(fi54, v21)
                                                                }
                                                                v21 = v21
                                                            }
                                                        }
                                                        var v99 uint8 = 224
                                                        var v100 int8 = 0
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = v83
                                                } else {
                                                    v59 = v59
                                                    fmt.Println(v1, v21, v59)
                                                }
                                            }
                                        }
                                    } else {
                                        if false {
                                            switch 128 {
                                            case 4294967295:
                                                if false {
                                                    var v101 uint16 = 183
                                                    if true {
                                                        fmt.Println(v81)
                                                        var v102 float64 = 0.0
                                                        fmt.Println(v1)
                                                        _ = v102
                                                    }
                                                    _ = v101
                                                } else {
                                                    switch v59 {
                                                    case "foo bar":
                                                        if ((3.14159 + -2.25) > 0.0) {
                                                            v81 -= (^(v81))
                                                        }
                                                    case "\n\t":
                                                        _ = len(v2)
                                                        switch fi54 {
                                                        case 492:
                                                            v103 := GenS2{F0: v81, F1: (^(fi54))}
                                                            _ = v103
                                                        case 127:
                                                            var v104 int32 = -158
                                                            _ = v104
                                                        }
                                                        var v105 uint8 = 127
                                                        _ = v105
                                                    }
                                                    switch v59 {
                                                    case "\n\t":
                                                        _ = len(v2)
                                                    }
                                                }
                                            case 644:
                                                fmt.Println(v21, v81, v81)
                                            case 2147483647:
                                                var v106 float32 = 0.1
                                                _ = v106
                                            }
                                            v81 += v81
                                            fmt.Println(v59, v59)
                                        }
                                        if (uint64(v81) != 940) {
                                            if false {
                                                v59 = "a"
                                                if false {
                                                    v107 := v2[42]
                                                    if false {
                                                        var v108 uint16 = 255
                                                        var v109 uint64 = 0
                                                        var v110 float32 = (1.5 / 0.1)
                                                        v110 -= v110
                                                        _ = v108
                                                        _ = v109
                                                        _ = v110
                                                    }
                                                    var v111 uint64 = 18446744073709551615
                                                    _ = v107
                                                    _ = v111
                                                } else {
                                                    var v112 uint32 = 0
                                                    _ = v112
                                                }
                                                v59 = v59
                                                v113 := GenS7{F0: 0}
                                                _ = v113
                                            } else {
                                                if true {
                                                    v59 = v59
                                                    fmt.Println(v81)
                                                } else {
                                                    v1 = 871
                                                    var v114 int64 = 9223372036854775807
                                                    if (!true) {
                                                        switch 32767 {
                                                        case 32768:
                                                            var v115 float32 = 0.0
                                                            _ = v115
                                                        case 42:
                                                            var v116 int8 = 127
                                                            _ = v116
                                                        default:
                                                            v59 = v59
                                                        }
                                                    }
                                                    _ = v114
                                                }
                                                if (false || true) {
                                                    if true {
                                                        switch 2147483647 {
                                                        case 4294967295:
                                                            v81 = (v81 % ((65535) | 1))
                                                            v117 := v3[(v59 + (v59 + v59))]
                                                            v59 = (v59 + "a")
                                                            var v118 int = fi54
                                                            v119 := GenS12{F0: 1}
                                                            fmt.Println(fi58, v21, fi54)
                                                            _ = v117
                                                            _ = v118
                                                            _ = v119
                                                        case -913:
                                                            v120 := GenS1{F0: -850, F1: 42, F2: 2}
                                                            _ = v120
                                                        }
                                                        if (4294967295 < 1) {
                                                            var v121 string = v59
                                                            _ = v121
                                                        } else {
                                                            var v122 uint8 = 42
                                                            _ = v122
                                                        }
                                                    } else {
                                                        var v123 float32 = 0.1
                                                        var v124 uint = 241
                                                        v124 += (v81 % ((v124) | 1))
                                                        _ = v123
                                                        _ = v124
                                                    }
                                                    var v125 int32 = 476
                                                    _ = v125
                                                }
                                                if (genFn1((v81 | 4294967295)) != 42) {
                                                    v81 -= 0
                                                } else {
                                                    _ = len(v3)
                                                    delete(v2, 1)
                                                }
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                v126 := []int{}
                                                _ = v126
                                            case "foo bar":
                                                v127 := GenS10{F0: 148, F1: -1, F2: v21}
                                                _ = v127
                                            case "hello":
                                                v21 += fi58
                                            default:
                                                v128 := []int{fi54, (v21 / ((127) | 1)), fi54}
                                                _ = v128
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                switch v59 {
                                                case "\n\t":
                                                    var v129 uint32 = uint32(v81)
                                                    _ = v129
                                                case "hello":
                                                    v81 = 65535
                                                }
                                            case "hello":
                                                switch "foo bar" {
                                                case "a":
                                                    if false {
                                                        v130 := GenS13{F0: 128}
                                                        _ = v130
                                                    }
                                                }
                                                v21 = 0
                                                _ = len(v2)
                                            case "a":
                                                v3["\n\t"] = 42
                                            default:
                                                v131 := []int{(fi58 + (fi58 / ((32767) | 1))), fi58}
                                                _ = v131
                                            }
                                        } else {
                                            v81 = 0
                                            switch (v59 + v59) {
                                            case "foo bar":
                                                if (943 >= 183) {
                                                    var v132 int16 = (v1 + (v1 % ((int16(fi54)) | 1)))
                                                    var v133 float32 = (1.5 / -2.25)
                                                    var v134 int32 = 0
                                                    _ = v132
                                                    _ = v133
                                                    _ = v134
                                                }
                                            case "":
                                                var v135 uint = 2
                                                _ = v135
                                            case "\n\t":
                                                if true {
                                                    if false {
                                                        if (-9223372036854775808 > 1) {
                                                            var v136 uint64 = 4294967295
                                                            delete(v3, ((("" + "\n\t") + (v59 + ((v59 + (("\n\t" + "foo bar") + v59)) + (v59 + (v59 + (v59 + v59)))))) + v59))
                                                            v137 := GenS14{F0: 4294967295}
                                                            _ = v136
                                                            _ = v137
                                                        } else {
                                                            switch fi58 {
                                                            case -90:
                                                                var v138 int32 = -629
                                                                _ = v138
                                                            case 824:
                                                                if true {
                                                                    var v139 uint64 = 485
                                                                    _ = v139
                                                                }
                                                            case -2147483648:
                                                                v21 += 127
                                                            }
                                                            v81 = (v81 + (v81 & (v81 ^ (v81 << ((v81) & 63)))))
                                                        }
                                                        var v140 int64 = 376
                                                        _ = v140
                                                    }
                                                    if (42 > 4294967295) {
                                                        var v141 uint64 = 18446744073709551615
                                                        _ = v141
                                                    }
                                                } else {
                                                    if true {
                                                        v59 = v59
                                                    } else {
                                                        v142 := GenS15{F0: 4294967295, F1: 4294967295, F2: 100}
                                                        v3[(((("hello" + v59) + "\n\t") + "\n\t") + "\n\t")] = 578
                                                        _ = v142
                                                    }
                                                    switch (v59 + "\n\t") {
                                                    case "\n\t":
                                                        var v143 int16 = v1
                                                        _ = v143
                                                    case "hello":
                                                        if (true && true) {
                                                            fmt.Println(fi58)
                                                            fmt.Println(v21, fi54)
                                                            v81 -= 2
                                                        } else {
                                                            delete(v3, (("a" + v59) + (v59 + (v59 + (("a" + v59) + v59)))))
                                                        }
                                                    case "a":
                                                        var v144 uint8 = 1
                                                        _ = v144
                                                    }
                                                }
                                            default:
                                                switch v21 {
                                                case 2147483647:
                                                    if ((-1 < 42) && (v1 > v1)) {
                                                        v145 := GenS4{F0: 162, F1: 0, F2: 32767}
                                                        _ = v145
                                                    } else {
                                                        v59 = v59
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v82
                            }
                            _ = v81
                        case -2147483648:
                            if (false && (false || (v59 > ("foo bar" + (v59 + "hello"))))) {
                                v146 := []int{(v21 ^ 4294967295), fi54}
                                var v147 int64 = -371
                                _ = v146
                                _ = v147
                            } else {
                                switch (v59 + "\n\t") {
                                case "a":
                                    if (9223372036854775807 > -1) {
                                        v21 = (fi54 / ((0) | 1))
                                    }
                                }
                                v59 = ((v59 + v59) + (v59 + "foo bar"))
                            }
                        case 32768:
                            if (87 <= 68) {
                                fmt.Println(v1, v21, v21)
                            } else {
                                var v148 float32 = 3.14159
                                v149 := GenS2{F0: 65535, F1: (v21 - (-(v21)))}
                                v150 := GenS9{F0: 42}
                                _ = v148
                                _ = v149
                                _ = v150
                            }
                        default:
                            if true {
                                var v151 string = ((v59 + v59) + "")
                                v152 := GenS8{F0: (((v151 + v59) + (v59 + v59)) + "hello"), F1: true}
                                _ = v151
                                _ = v152
                            } else {
                                var v153 int64 = -545
                                _ = v153
                            }
                        }
                        v154 := GenS16{F0: -14}
                        _ = v154
                    }
                    if (2 == 303) {
                        _ = len(v2)
                    } else {
                        switch (fi54 & (v21 - -473)) {
                        case 4294967295:
                            v155 := []int{(v21 % ((fi54) | 1)), (-(fi54))}
                            if (0.0 < float32(3.14159)) {
                                v21 += (fi58 - 0)
                                if ((v59 == (("" + (v59 + v59)) + v59)) && true) {
                                    if false {
                                        var v156 string = v59
                                        var v157 uint8 = 0
                                        switch 1 {
                                        case 569:
                                            v21 -= int(v1)
                                            var v158 uint = 4294967295
                                            v159 := GenS17{F0: v59, F1: v157, F2: v158}
                                            if len(v155) > 0 {
                                                v155[0] = (-(fi58))
                                            }
                                            _ = v158
                                            _ = v159
                                        case 127:
                                            delete(v2, 1)
                                        case -9223372036854775808:
                                            if (127 < 100) {
                                                if (!true) {
                                                    var v160 float32 = 1.5
                                                    _ = v160
                                                }
                                            }
                                            v155 = append(v155, v21)
                                            if false {
                                                switch -9223372036854775808 {
                                                case -457:
                                                    v21 = fi58
                                                case 127:
                                                    if false {
                                                        v1 = (v1 + 407)
                                                        fmt.Println(v1, v156)
                                                        v59 = v59
                                                        var v161 uint64 = 2
                                                        _ = v161
                                                    } else {
                                                        switch v59 {
                                                        case "":
                                                            switch -434 {
                                                            case 32768:
                                                                if true {
                                                                    var v162 string = "hello"
                                                                    if false {
                                                                        v163 := GenS6{F0: 1}
                                                                        _ = v163
                                                                    } else {
                                                                        var v164 uint = 0
                                                                        switch ("foo bar" + (v59 + v156)) {
                                                                        case "a":
                                                                            v162 = v162
                                                                        case "":
                                                                            fmt.Println(v21)
                                                                        case "foo bar":
                                                                            var v165 int16 = v1
                                                                            _ = v165
                                                                        }
                                                                        var v166 uint64 = 0
                                                                        _ = v164
                                                                        _ = v166
                                                                    }
                                                                    v167 := GenS2{F0: 2, F1: v21}
                                                                    fmt.Println(v162, v1)
                                                                    var v168 string = v156
                                                                    v157 -= v157
                                                                    v1 -= (v1 - (-(v1)))
                                                                    _ = v162
                                                                    _ = v167
                                                                    _ = v168
                                                                }
                                                            case -574:
                                                                fmt.Println(v157)
                                                            default:
                                                                delete(v2, 230)
                                                            }
                                                        case "a":
                                                            _ = len(v2)
                                                        case "\n\t":
                                                            fmt.Println(v157, fi58, fi54)
                                                        default:
                                                            if false {
                                                                if (false || false) {
                                                                    v169 := GenS8{F0: v156, F1: false}
                                                                    _ = v169
                                                                }
                                                                switch (("" + "foo bar") + v156) {
                                                                case "\n\t":
                                                                    fmt.Println(v157, v21, v157)
                                                                case "a":
                                                                    switch fi58 {
                                                                    case 0:
                                                                        var v170 uint64 = 244
                                                                        _ = v170
                                                                    default:
                                                                        fmt.Println(v157)
                                                                    }
                                                                case "foo bar":
                                                                    if len(v155) > 0 {
                                                                        v155[0] = (v21 >> ((659) & 63))
                                                                    }
                                                                default:
                                                                    if (v157 >= (v157 % (((v157 & v157)) | 1))) {
                                                                        if true {
                                                                            v21 += (fi54 | 911)
                                                                            v157 += (^(v157))
                                                                            switch v21 {
                                                                            case 1:
                                                                                if (276 != 4294967295) {
                                                                                    v157 = (v157 & v157)
                                                                                    delete(v3, "")
                                                                                    if false {
                                                                                        var v171 int32 = -2147483648
                                                                                        var v172 float32 = (-2.25 / ((-2.25 + (0.0 * (0.1 + 1e10))) * (1.5 / 1.5)))
                                                                                        _ = v171
                                                                                        _ = v172
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            switch (fi54 | fi58) {
                                                                            case 127:
                                                                                v173 := GenS1{F0: -9223372036854775808, F1: 212, F2: 2}
                                                                                _ = v173
                                                                            case -9223372036854775808:
                                                                                v157 = v157
                                                                            case -1:
                                                                                if (1 == (v157 % ((1) | 1))) {
                                                                                    var v174 uint32 = 310
                                                                                    _ = v174
                                                                                }
                                                                            }
                                                                            _ = len(v3)
                                                                            var v175 float32 = ((-2.25 - 0.0) + 1e10)
                                                                            v157 = (v157 / (((^(v157))) | 1))
                                                                            _ = v175
                                                                        }
                                                                        if (v21 >= fi58) {
                                                                            if (128 == (fi58 - (-(fi58)))) {
                                                                                v2[2] = fi58
                                                                                var v176 int8 = -1
                                                                                _ = v176
                                                                            }
                                                                        } else {
                                                                            switch 1 {
                                                                            case 2147483647:
                                                                                if true {
                                                                                    v157 = 235
                                                                                    if len(v155) > 0 {
                                                                                        v155[0] = 42
                                                                                    }
                                                                                    switch int(v1) {
                                                                                    case 0:
                                                                                        v156 = ("hello" + (v59 + v156))
                                                                                    default:
                                                                                        fmt.Println(v21)
                                                                                    }
                                                                                }
                                                                            case 9223372036854775807:
                                                                                var v177 float64 = float64(1e10)
                                                                                if (!(int8(v1) <= genFn1(uint(v157)))) {
                                                                                    var v178 int64 = -1
                                                                                    v1 = -32768
                                                                                    v21 -= v21
                                                                                    _ = v178
                                                                                } else {
                                                                                    if len(v155) >= 2 {
                                                                                        v155 = v155[1:len(v155)]
                                                                                    }
                                                                                    v179 := GenS18{F0: 255, F1: (42 >= -1), F2: -196}
                                                                                    var v180 uint16 = 146
                                                                                    _ = v179
                                                                                    _ = v180
                                                                                }
                                                                                var v181 int = fi58
                                                                                _ = v177
                                                                                _ = v181
                                                                            case 389:
                                                                                if (42 == fi58) {
                                                                                    switch -586 {
                                                                                    case 127:
                                                                                        var v182 float64 = 3.14159
                                                                                        _ = v182
                                                                                    case 781:
                                                                                        fmt.Println(fi54)
                                                                                    case 696:
                                                                                        var v183 uint64 = 0
                                                                                        _ = v183
                                                                                    }
                                                                                    v21 = (fi58 ^ fi54)
                                                                                }
                                                                            default:
                                                                                v184 := GenS9{F0: 2147483647}
                                                                                _ = v184
                                                                            }
                                                                        }
                                                                        var v185 float32 = 3.14159
                                                                        _ = v185
                                                                    }
                                                                }
                                                                fmt.Println(v21)
                                                                var v186 float64 = ((-2.25 + (((0.0 - 1.5) + -2.25) - 3.14159)) * (1e10 + 3.14159))
                                                                var v187 string = ("" + v59)
                                                                _ = v186
                                                                _ = v187
                                                            } else {
                                                                if (18446744073709551615 != 594) {
                                                                    fmt.Println(v156)
                                                                    var v188 int8 = 41
                                                                    v59 = v59
                                                                    _ = v188
                                                                } else {
                                                                    v2[42] = (-(fi58))
                                                                    v189 := GenS16{F0: 2147483647}
                                                                    if true {
                                                                        fmt.Println(fi58)
                                                                        switch v59 {
                                                                        case "a":
                                                                            v21 = (fi58 << ((90) & 63))
                                                                        case "\n\t":
                                                                            v155 = append(v155, (fi58 % ((fi54) | 1)))
                                                                        case "":
                                                                            v190 := GenS13{F0: v157}
                                                                            _ = v190
                                                                        default:
                                                                            v59 = "a"
                                                                        }
                                                                    }
                                                                    v191 := GenS15{F0: 1, F1: 1, F2: 127}
                                                                    v192 := GenS19{F0: 115}
                                                                    _ = v189
                                                                    _ = v191
                                                                    _ = v192
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                v156 = (v59 + (v156 + v156))
                                                if len(v155) >= 2 {
                                                    v155 = v155[1:len(v155)]
                                                }
                                                switch v59 {
                                                case "":
                                                    fmt.Println(v1, fi58)
                                                }
                                            } else {
                                                v193 := GenS6{F0: 2147483648}
                                                if false {
                                                    var v194 uint64 = 2
                                                    v3[v156] = 644
                                                    _ = len(v155)
                                                    _ = v194
                                                } else {
                                                    var v195 uint32 = 255
                                                    if false {
                                                        switch fi54 {
                                                        case -9223372036854775808:
                                                            if (false || (v59 == ((v156 + "foo bar") + ""))) {
                                                                v196 := GenS6{F0: v195}
                                                                var v197 uint8 = (^(v157))
                                                                _ = v196
                                                                _ = v197
                                                            }
                                                        }
                                                        switch (((v59 + "foo bar") + "") + "\n\t") {
                                                        case "foo bar":
                                                            if (!((((1e10 - 0.0) + 1e10) + ((((0.0 / 1.0) + float64(float32((((0.0 + float64((-2.25 / 1.0))) * (-2.25 / 1.0)) - 0.0)))) + 3.14159) + 0.1)) != (-2.25 + 3.14159))) {
                                                                v3["foo bar"] = (v195 / ((v195) | 1))
                                                                if true {
                                                                    fmt.Println(fi58, v1)
                                                                }
                                                                if (!(true || (18446744073709551615 < 255))) {
                                                                    v195 = (v195 ^ 2)
                                                                }
                                                            }
                                                        case "":
                                                            switch ("foo bar" + v59) {
                                                            case "foo bar":
                                                                var v198 uint = 65535
                                                                _ = v198
                                                            case "hello":
                                                                var v199 uint = uint(v157)
                                                                _ = v199
                                                            }
                                                            v200 := GenS15{F0: 42, F1: 4294967295, F2: genFn1(676)}
                                                            var v201 string = ("" + v59)
                                                            _ = v200
                                                            _ = v201
                                                        case "\n\t":
                                                            v1 += v1
                                                        default:
                                                            v202 := GenS5{F0: 18446744073709551615}
                                                            _ = v202
                                                        }
                                                    } else {
                                                        v1 += int16(fi54)
                                                        var v203 uint8 = (v157 >> ((578) & 63))
                                                        _ = v203
                                                    }
                                                    v1 = 197
                                                    _ = v195
                                                }
                                                _ = v193
                                            }
                                        default:
                                            if false {
                                                if ((3.14159 / 1.5) >= 1.5) {
                                                    v21 = -1
                                                    var v204 uint16 = 2
                                                    _ = v204
                                                }
                                                switch 42 {
                                                case 32767:
                                                    _ = len(v2)
                                                case -2147483648:
                                                    v205 := GenS20{F0: (-(v1))}
                                                    _ = v205
                                                }
                                                switch "a" {
                                                case "":
                                                    var v206 uint64 = 1
                                                    _ = v206
                                                case "\n\t":
                                                    v156 = (("hello" + (v156 + v156)) + v59)
                                                default:
                                                    if true {
                                                        var v207 int64 = -9223372036854775808
                                                        if false {
                                                            var v208 int32 = 2147483647
                                                            v157 += v157
                                                            _ = v208
                                                        } else {
                                                            switch "" {
                                                            case "\n\t":
                                                                switch v156 {
                                                                case "":
                                                                    v209 := GenS15{F0: 4294967295, F1: 4294967295, F2: genFn1(1)}
                                                                    _ = v209
                                                                }
                                                                var v210 float32 = -2.25
                                                                v207 = (v207 % ((v207) | 1))
                                                                _ = v210
                                                            default:
                                                                var v211 int64 = v207
                                                                _ = v211
                                                            }
                                                        }
                                                        _ = v207
                                                    } else {
                                                        var v212 float32 = 1.5
                                                        v213 := GenS21{F0: 1, F1: -9223372036854775808, F2: 0}
                                                        _ = v212
                                                        _ = v213
                                                    }
                                                }
                                            }
                                        }
                                        _ = v156
                                        _ = v157
                                    } else {
                                        if ((3.14159 * float32(1e10)) > 3.14159) {
                                            var v214 float64 = (-2.25 + -2.25)
                                            _ = v214
                                        }
                                    }
                                } else {
                                    var v215 uint8 = 42
                                    fmt.Println(v215, fi58)
                                    _ = v215
                                }
                                if ((1.5 != (float32(-2.25) - (float32(1e10) + (float32(3.14159) / (((1e10 * 3.14159) - 3.14159) * (float32(1.5) * 0.1)))))) && (49 < 1)) {
                                    v59 = "foo bar"
                                    v216 := v3[v59]
                                    v21 = fi54
                                    switch fi54 {
                                    case -1:
                                        switch ((v59 + "a") + (v59 + "")) {
                                        case "hello":
                                            if false {
                                                v1 += (v1 << ((171) & 63))
                                                v216 = v216
                                                fmt.Println(fi58, v21)
                                            }
                                            v217 := GenS22{F0: (v1 ^ (v1 / (((v1 + 0)) | 1))), F1: 100}
                                            fmt.Println(v21, v59)
                                            v1 += 42
                                            _ = v217
                                        case "\n\t":
                                            v21 = 9223372036854775807
                                        default:
                                            var v218 bool = (!false)
                                            _ = v218
                                        }
                                    case 2147483647:
                                        fmt.Println(v1, v59)
                                        v219 := GenS23{F0: false}
                                        _ = v219
                                    default:
                                        v1 = (v1 >> ((2) & 63))
                                    }
                                    var v220 float32 = float32(3.14159)
                                    _ = v216
                                    _ = v220
                                }
                                if len(v155) > 0 {
                                    v155[0] = fi54
                                }
                            } else {
                                fmt.Println(fi54, v1, fi58)
                                switch 32768 {
                                case 299:
                                    if false {
                                        var v221 int16 = 1000
                                        if (255 > 65535) {
                                            switch (-(v21)) {
                                            case -2147483648:
                                                if len(v155) >= 2 {
                                                    v155 = v155[1:len(v155)]
                                                }
                                            default:
                                                var v222 float32 = 3.14159
                                                _ = v222
                                            }
                                            switch (v59 + v59) {
                                            case "":
                                                if (!false) {
                                                    v223 := GenS4{F0: 42, F1: 0, F2: (v221 * v221)}
                                                    var v224 uint16 = 0
                                                    _ = v223
                                                    _ = v224
                                                }
                                            case "a":
                                                v225 := GenS2{F0: 42, F1: (^(fi54))}
                                                _ = v225
                                            }
                                            v21 = fi54
                                        }
                                        _ = v221
                                    } else {
                                        var v226 int64 = -9223372036854775808
                                        v1 = (v1 - -1)
                                        v227 := GenS24{F0: 18446744073709551615}
                                        _ = v226
                                        _ = v227
                                    }
                                case 42:
                                    if (-963 >= 1) {
                                        var v228 float64 = (0.1 * 1.5)
                                        fmt.Println(v59, v59)
                                        _ = v228
                                    }
                                    var v229 bool = (true && false)
                                    if (-87 < 1) {
                                        v230 := GenS7{F0: 123}
                                        _ = v230
                                    }
                                    _ = v229
                                default:
                                    v231 := GenS25{F0: false}
                                    _ = v231
                                }
                                var v232 uint8 = 2
                                _ = v232
                            }
                            var v233 int = (fi58 / ((fi54) | 1))
                            v21 -= 42
                            _ = v155
                            _ = v233
                        case 42:
                            if false {
                                fmt.Println(v1, v21, fi58)
                                v234 := []int{(fi54 / (((^(fi54))) | 1)), (^(fi54)), -1, v21}
                                fmt.Println(v1, fi58)
                                v21 -= -2147483648
                                v59 = (v59 + (v59 + ("\n\t" + ((((v59 + v59) + "\n\t") + v59) + v59))))
                                _ = v234
                            } else {
                                var v235 int8 = 60
                                _ = v235
                            }
                        default:
                            var v236 uint64 = 0
                            if true {
                                v1 = (v1 + v1)
                            }
                            _ = v236
                        }
                        v21 = (fi58 * (fi58 % ((v21) | 1)))
                        var v237 bool = false
                        _ = v237
                    }
                    _ = v59
                }
            case 32767:
                switch -209 {
                case 127:
                    v21 -= (fi58 ^ (fi58 + (fi54 - (fi54 - 1))))
                case 535:
                    if (!false) {
                        var v238 int8 = -38
                        var v239 int16 = 1000
                        _ = v238
                        _ = v239
                    } else {
                        v3[""] = 1
                    }
                }
            default:
                var v240 bool = true
                _ = v240
            }
            if true {
                v21 -= (^(fi58))
            }
            var v241 int16 = v1
            _ = fi58
            _ = v241
        }
        _ = fi54
    }
    var v242 uint64 = 534
    for fi243 := 0; fi243 < 17; fi243++ {
        for fi244 := 0; fi244 < 4; fi244++ {
            switch (fi243 + (fi243 - (v21 & (v21 << ((2) & 63))))) {
            case 32767:
                v2[1] = (^(fi243))
            case 42:
                var v245 int16 = v1
                _ = v245
            case 0:
                fmt.Println(v21, v21)
            }
            switch ("" + "foo bar") {
            case "":
                switch fi243 {
                case 128:
                    var v246 int32 = -821
                    _ = v246
                case 32767:
                    v21 = -2147483648
                case 4294967295:
                    v247 := GenS25{F0: false}
                    _ = v247
                }
            case "\n\t":
                if false {
                    v248 := v3["hello"]
                    var v249 float32 = 1e10
                    _ = v248
                    _ = v249
                }
            default:
                var v250 int64 = 9223372036854775807
                _ = v250
            }
            _ = fi244
        }
        var v251 int8 = -1
        _ = fi243
        _ = v251
    }
    _ = len(v2)
    v252 := GenS26{F0: false, F1: 1}
    v1 += (^(v1))
    v242 = 1
    for fi253 := 0; fi253 < 13; fi253++ {
        for fi254 := 0; fi254 < 8; fi254++ {
            v255 := GenS27{F0: 691, F1: false}
            var v256 float32 = float32(-2.25)
            var v257 uint32 = 737
            var v258 int32 = 42
            _ = fi254
            _ = v255
            _ = v256
            _ = v257
            _ = v258
        }
        switch v21 {
        case 32768:
            v242 += v242
        case 787:
            var v259 int = v21
            _ = v259
        }
        _ = fi253
    }
    for fi260 := 0; fi260 < 8; fi260++ {
        switch ((("" + ("\n\t" + "")) + "hello") + "hello") {
        case "hello":
            for fi261 := 0; fi261 < 13; fi261++ {
                if true {
                    v262 := GenS2{F0: 986, F1: fi261}
                    v1 = -32768
                    _ = v262
                } else {
                    if false {
                        var v263 int16 = (v1 | 0)
                        v1 += 98
                        fmt.Println(v263, v1, v21)
                        _ = v263
                    }
                    v264 := GenS28{F0: "", F1: 65535}
                    v265 := GenS5{F0: (v242 ^ v242)}
                    _ = v264
                    _ = v265
                }
                _ = fi261
            }
        default:
            v266 := v3["hello"]
            _ = v266
        }
        switch (v21 | (v21 << ((426) & 63))) {
        case 4294967295:
            v267 := []int{523}
            _ = v267
        case 319:
            v1 += -943
        }
        _ = fi260
    }
    for fi268 := 0; fi268 < 1; fi268++ {
        v269 := GenS8{F0: "", F1: (false || (0 >= 0))}
        _ = fi268
        _ = v269
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v242:", v242)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v242
    _ = v252
    _ = v3
}
