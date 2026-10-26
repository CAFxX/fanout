package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int16; F1 int16; F2 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 uint; F1 int; F2 int8 }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 uint8; F1 bool }
func genFn0() uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6)
        var v7 int8 = 42
        _ = fi6
        _ = v7
    }
    switch 0 {
    case 128:
        var v8 int16 = 1000
        _ = v8
    case 4294967295:
        var v9 uint = 18446744073709551615
        _ = v9
    case 32768:
        v10 := []int{9223372036854775807, 127, 2147483647}
        _ = v10
    default:
        if false {
            v11 := GenS1{F0: -1}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 string = (("hello" + "a") + "")
                v13 = v13
                _ = fi12
                _ = v13
            }
            if (!(-32768 != 1000)) {
                if true {
                    if (true && false) {
                        var v14 float64 = (3.14159 * 0.0)
                        _ = v14
                    }
                    var v15 int = -1
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v15 = 966
                        v15 = (v15 ^ v15)
                        _ = fi16
                    }
                    _ = v15
                } else {
                    if false {
                        v17 := []int{-661, -146, 736}
                        _ = v17
                    } else {
                        var v18 uint64 = 760
                        if (!(0 == 65535)) {
                            var v19 float64 = ((1e10 + ((1e10 - (float64(0.1) * (3.14159 + 1.5))) - 1.5)) / float64(1.5))
                            if false {
                                fmt.Println(v18, v18)
                                switch -1 {
                                case 32767:
                                    if (!(!true)) {
                                        switch 42 {
                                        case -9223372036854775808:
                                            v19 = v19
                                        case 32767:
                                            var v20 float64 = v19
                                            _ = v20
                                        case 128:
                                            if false {
                                                v21 := []int{32768}
                                                var v22 uint = uint(v18)
                                                _ = v21
                                                _ = v22
                                            } else {
                                                for fi23 := 0; fi23 < 15; fi23++ {
                                                    v19 = 1.5
                                                    v24 := GenS2{F0: 0, F1: 486, F2: 9223372036854775807}
                                                    _ = fi23
                                                    _ = v24
                                                }
                                            }
                                        default:
                                            if false {
                                                v18 = v18
                                                fmt.Println(v19)
                                            }
                                        }
                                        v25 := GenS2{F0: 42, F1: 0, F2: -9223372036854775808}
                                        _ = v25
                                    } else {
                                        if (!(true || false)) {
                                            if true {
                                                if (!false) {
                                                    var v26 string = ""
                                                    _ = v26
                                                } else {
                                                    var v27 uint16 = 32767
                                                    var v28 int32 = 0
                                                    var v29 int64 = 42
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                }
                                                for fi30 := 0; fi30 < 7; fi30++ {
                                                    if ((v18 * (v18 | 42)) <= 18446744073709551615) {
                                                        for fi31 := 0; fi31 < 13; fi31++ {
                                                            v32 := GenS3{F0: 846, F1: (v18 + (v18 * 192))}
                                                            _ = fi31
                                                            _ = v32
                                                        }
                                                        switch (-(fi30)) {
                                                        case -378:
                                                            var v33 uint8 = 128
                                                            _ = v33
                                                        case 42:
                                                            var v34 uint8 = 2
                                                            _ = v34
                                                        default:
                                                            var v35 uint = 65535
                                                            _ = v35
                                                        }
                                                    } else {
                                                        switch fi30 {
                                                        case -2147483648:
                                                            if true {
                                                                if false {
                                                                    v36 := GenS3{F0: 42, F1: 825}
                                                                    for fi37 := 0; fi37 < 11; fi37++ {
                                                                        if true {
                                                                            if true {
                                                                                var v38 int = (fi30 & fi30)
                                                                                v19 = float64(float32(((float64(3.14159) - (v19 / v19)) + v19)))
                                                                                fmt.Println(v18, fi37)
                                                                                _ = v38
                                                                            } else {
                                                                                v39 := GenS4{F0: 32767, F1: 0, F2: -1}
                                                                                var v40 uint64 = (v18 + v18)
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            v41 := []int{}
                                                                            _ = v41
                                                                        }
                                                                        v42 := []int{}
                                                                        if true {
                                                                            var v43 int64 = -1
                                                                            _ = len(v42)
                                                                            v44 := GenS2{F0: 1, F1: 65535, F2: v43}
                                                                            var v45 uint64 = (v18 & v18)
                                                                            _ = v43
                                                                            _ = v44
                                                                            _ = v45
                                                                        } else {
                                                                            var v46 bool = false
                                                                            _ = v46
                                                                        }
                                                                        _ = fi37
                                                                        _ = v42
                                                                    }
                                                                    if (false && false) {
                                                                        if (-113 <= -128) {
                                                                            if (!false) {
                                                                                switch -9223372036854775808 {
                                                                                case 9223372036854775807:
                                                                                    for fi47 := 0; fi47 < 16; fi47++ {
                                                                                        if false {
                                                                                            v19 += v19
                                                                                            switch 127 {
                                                                                            case -740:
                                                                                                if true {
                                                                                                    v48 := GenS4{F0: -1, F1: 0, F2: -13}
                                                                                                    v18 = v18
                                                                                                    _ = v48
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                v18 -= (v18 - v18)
                                                                                            case -9223372036854775808:
                                                                                                var v49 float64 = ((-2.25 / 3.14159) + float64(-2.25))
                                                                                                _ = v49
                                                                                            }
                                                                                            var v50 int32 = 1
                                                                                            _ = v50
                                                                                        }
                                                                                        if (255 > 162) {
                                                                                            v19 += v19
                                                                                            if false {
                                                                                                v51 := []int{fi30, (-(fi30))}
                                                                                                var v52 uint32 = 36
                                                                                                _ = v51
                                                                                                _ = v52
                                                                                            }
                                                                                            v19 = 1.5
                                                                                            var v53 uint = 2
                                                                                            v18 -= (v18 * 2)
                                                                                            _ = v53
                                                                                        }
                                                                                        fmt.Println(v19)
                                                                                        _ = fi47
                                                                                    }
                                                                                case -2147483648:
                                                                                    if ((v18 + 18446744073709551615) != (v18 / (((v18 << ((4294967295) & 63))) | 1))) {
                                                                                        if false {
                                                                                            v18 -= v18
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v54 := []int{(fi30 << ((2) & 63)), (fi30 - (fi30 | (fi30 + fi30))), 42, -2147483648}
                                                                            v55 := GenS2{F0: 42, F1: 21, F2: -363}
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                        fmt.Println(v18, fi30, v18)
                                                                    }
                                                                    _ = v36
                                                                }
                                                                v56 := GenS3{F0: 1, F1: v18}
                                                                v57 := []int{fi30, (-(fi30)), 127, (fi30 * (fi30 + fi30))}
                                                                _ = v56
                                                                _ = v57
                                                            } else {
                                                                fmt.Println(v19)
                                                                for fi58 := 0; fi58 < 16; fi58++ {
                                                                    v59 := []int{}
                                                                    v19 -= 0.1
                                                                    var v60 int32 = -768
                                                                    _ = fi58
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                            }
                                                        case 84:
                                                            fmt.Println(fi30)
                                                        default:
                                                            var v61 float32 = (-2.25 + 0.1)
                                                            _ = v61
                                                        }
                                                    }
                                                    for fi62 := 0; fi62 < 19; fi62++ {
                                                        v18 = v18
                                                        var v63 int16 = -985
                                                        if true {
                                                            var v64 uint16 = 148
                                                            if false {
                                                                var v65 int = fi30
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v18 -= (v18 - (v18 * v18))
                                                                    var v66 bool = true
                                                                    _ = v66
                                                                case "foo bar":
                                                                    v67 := GenS5{F0: true, F1: 107}
                                                                    v19 += 1.5
                                                                    _ = v67
                                                                case "hello":
                                                                    fmt.Println(v19, v64, v19)
                                                                }
                                                                v64 -= v64
                                                                _ = v65
                                                            }
                                                            _ = v64
                                                        } else {
                                                            if ((1 != 42) || true) {
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v68 int8 = 127
                                                                    _ = v68
                                                                default:
                                                                    if (0 > 42) {
                                                                        var v69 uint8 = 255
                                                                        v69 = (^(v69))
                                                                        _ = v69
                                                                    }
                                                                    if (!(0 < 147)) {
                                                                        fmt.Println(v63, v63)
                                                                    }
                                                                }
                                                                var v70 int64 = 19
                                                                var v71 float32 = -2.25
                                                                var v72 int8 = 127
                                                                _ = v70
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                        }
                                                        _ = fi62
                                                        _ = v63
                                                    }
                                                    switch (^(fi30)) {
                                                    case 312:
                                                        var v73 uint64 = 42
                                                        _ = v73
                                                    case 9223372036854775807:
                                                        for fi74 := 0; fi74 < 8; fi74++ {
                                                            if (!(0 != 459)) {
                                                                v18 = 42
                                                                v18 = v18
                                                            } else {
                                                                v19 -= float64((3.14159 - 0.0))
                                                            }
                                                            var v75 uint32 = 65535
                                                            _ = fi74
                                                            _ = v75
                                                        }
                                                    case -902:
                                                        if false {
                                                            switch (("a" + "\n\t") + "\n\t") {
                                                            case "":
                                                                v19 = v19
                                                            case "hello":
                                                                for fi76 := 0; fi76 < 16; fi76++ {
                                                                    if false {
                                                                        v77 := GenS6{F0: (127 >= 127), F1: 2147483648}
                                                                        v78 := GenS7{F0: fi30, F1: v18}
                                                                        v79 := []int{32768, -9223372036854775808, fi30, fi76}
                                                                        _ = v77
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                    v80 := GenS2{F0: 42, F1: 106, F2: -1}
                                                                    v18 = (v18 / ((v18) | 1))
                                                                    _ = fi76
                                                                    _ = v80
                                                                }
                                                            }
                                                            if false {
                                                                if (0 <= 275) {
                                                                    switch (fi30 % ((fi30) | 1)) {
                                                                    case -1:
                                                                        v81 := []int{}
                                                                        _ = v81
                                                                    case 0:
                                                                        var v82 float64 = 3.14159
                                                                        _ = v82
                                                                    case 4294967295:
                                                                        var v83 float64 = v19
                                                                        _ = v83
                                                                    }
                                                                } else {
                                                                    v18 += (v18 & v18)
                                                                    var v84 int = (fi30 << ((4294967295) & 63))
                                                                    _ = v84
                                                                }
                                                                if (1 != 42) {
                                                                    v85 := GenS8{F0: 32}
                                                                    v86 := []int{(fi30 & (fi30 + fi30)), fi30}
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                v87 := []int{}
                                                                _ = v87
                                                            } else {
                                                                var v88 uint32 = 42
                                                                var v89 int = (fi30 / ((127) | 1))
                                                                v88 -= (v88 ^ (^(v88)))
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            var v90 uint32 = 402
                                                            _ = v90
                                                        } else {
                                                            v19 += v19
                                                            var v91 string = "foo bar"
                                                            switch (fi30 / ((fi30) | 1)) {
                                                            case 128:
                                                                fmt.Println(v18, v18)
                                                            case -871:
                                                                for fi92 := 0; fi92 < 14; fi92++ {
                                                                    if (true && false) {
                                                                        var v93 int8 = 1
                                                                        fmt.Println(v19, fi30, v93)
                                                                        v94 := GenS9{F0: v91}
                                                                        v19 -= ((3.14159 + 0.1) + (1.5 * ((0.0 + v19) + float64((0.0 * float32(0.0))))))
                                                                        v95 := GenS10{F0: 255, F1: -1}
                                                                        v96 := []int{}
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(fi30, v19, fi30)
                                                                        } else {
                                                                            switch v91 {
                                                                            case "hello":
                                                                                fmt.Println(v91, v91, v91)
                                                                            case "foo bar":
                                                                                var v97 string = ((("hello" + v91) + (v91 + v91)) + v91)
                                                                                var v98 uint64 = 0
                                                                                _ = v97
                                                                                _ = v98
                                                                            case "":
                                                                                var v99 float32 = -2.25
                                                                                _ = v99
                                                                            }
                                                                            v100 := []int{fi92, 127}
                                                                            v18 += (v18 & v18)
                                                                            v18 = 352
                                                                            _ = v100
                                                                        }
                                                                        v18 = v18
                                                                    }
                                                                    fmt.Println(fi30, fi30, v91)
                                                                    _ = fi92
                                                                }
                                                            case 42:
                                                                var v101 int32 = -623
                                                                _ = v101
                                                            default:
                                                                for fi102 := 0; fi102 < 1; fi102++ {
                                                                    var v103 uint16 = 32767
                                                                    if (!false) {
                                                                        var v104 uint32 = 533
                                                                        _ = v104
                                                                    }
                                                                    var v105 int16 = 1000
                                                                    _ = fi102
                                                                    _ = v103
                                                                    _ = v105
                                                                }
                                                            }
                                                            _ = v91
                                                        }
                                                    default:
                                                        var v106 uint = 1
                                                        _ = v106
                                                    }
                                                    _ = fi30
                                                }
                                            }
                                            switch ("foo bar" + "") {
                                            case "\n\t":
                                                var v107 uint = 0
                                                _ = v107
                                            case "":
                                                v19 += v19
                                            case "hello":
                                                if true {
                                                    var v108 int16 = 0
                                                    if false {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v19 -= v19
                                                        case "\n\t":
                                                            for fi109 := 0; fi109 < 14; fi109++ {
                                                                v110 := []int{}
                                                                if true {
                                                                    for fi111 := 0; fi111 < 13; fi111++ {
                                                                        if (fi111 == 127) {
                                                                            v18 = 0
                                                                            if true {
                                                                                v112 := GenS9{F0: "\n\t"}
                                                                                if ((!true) && true) {
                                                                                    var v113 uint64 = v18
                                                                                    v110 = append(v110, 32768)
                                                                                    if len(v110) >= 2 {
                                                                                        v110 = v110[1:len(v110)]
                                                                                    }
                                                                                    _ = v113
                                                                                }
                                                                                _ = len(v110)
                                                                                switch fi109 {
                                                                                case 128:
                                                                                    _ = len(v110)
                                                                                    v18 = 42
                                                                                    fmt.Println(fi109, fi111, v19)
                                                                                case 42:
                                                                                    if (!(((0.1 - (3.14159 - ((-2.25 * 1.5) + float32((0.1 - v19))))) - 0.0) < ((float32((v19 / (v19 - ((((v19 + ((v19 / (v19 - (1e10 * v19))) + float64((-2.25 - (-2.25 + -2.25))))) - v19) / 1e10) * 1e10)))) + 0.1) + 1.5))) {
                                                                                        v108 += v108
                                                                                        if (!true) {
                                                                                            v18 += (^(v18))
                                                                                            if false {
                                                                                                v108 = 0
                                                                                                v114 := GenS11{F0: 817}
                                                                                                _ = v114
                                                                                            }
                                                                                            fmt.Println(v19, v18, fi111)
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v110)
                                                                                        v115 := GenS12{F0: 4294967295}
                                                                                        var v116 int = 752
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    v19 = ((float64(float32(v19)) * 1e10) + v19)
                                                                                }
                                                                                var v117 bool = (118 >= 5)
                                                                                v118 := GenS8{F0: 0}
                                                                                _ = v112
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                            var v119 float32 = 1e10
                                                                            v18 -= 42
                                                                            _ = v119
                                                                        } else {
                                                                            v108 += 881
                                                                        }
                                                                        _ = fi111
                                                                    }
                                                                } else {
                                                                    for fi120 := 0; fi120 < 16; fi120++ {
                                                                        v108 -= v108
                                                                        if len(v110) > 0 {
                                                                            v110[0] = fi120
                                                                        }
                                                                        _ = fi120
                                                                    }
                                                                    _ = len(v110)
                                                                }
                                                                _ = fi109
                                                                _ = v110
                                                            }
                                                        default:
                                                            if true {
                                                                switch "hello" {
                                                                case "":
                                                                    var v121 string = "a"
                                                                    _ = v121
                                                                case "hello":
                                                                    v108 -= 0
                                                                default:
                                                                    v122 := []int{721}
                                                                    _ = v122
                                                                }
                                                                var v123 int16 = (v108 * 0)
                                                                _ = v123
                                                            } else {
                                                                if true {
                                                                    for fi124 := 0; fi124 < 11; fi124++ {
                                                                        v125 := []int{fi124, fi124, fi124}
                                                                        _ = fi124
                                                                        _ = v125
                                                                    }
                                                                    v19 = 1e10
                                                                    v126 := []int{}
                                                                    _ = v126
                                                                }
                                                                v19 += v19
                                                            }
                                                        }
                                                        if false {
                                                            v127 := GenS7{F0: 32767, F1: 42}
                                                            _ = v127
                                                        } else {
                                                            if ((-430 < -695) || true) {
                                                                for fi128 := 0; fi128 < 10; fi128++ {
                                                                    for fi129 := 0; fi129 < 20; fi129++ {
                                                                        if (false || (!true)) {
                                                                            var v130 int64 = 744
                                                                            switch fi128 {
                                                                            case -9223372036854775808:
                                                                                fmt.Println(fi128, v19)
                                                                            case -2147483648:
                                                                                v18 = (^(v18))
                                                                                v18 -= 2
                                                                                v131 := GenS1{F0: 39}
                                                                                _ = v131
                                                                            }
                                                                            v19 = v19
                                                                            _ = v130
                                                                        } else {
                                                                            var v132 float64 = 1.5
                                                                            _ = v132
                                                                        }
                                                                        var v133 int = (fi128 << ((1) & 63))
                                                                        _ = fi129
                                                                        _ = v133
                                                                    }
                                                                    _ = fi128
                                                                }
                                                                if (v108 == v108) {
                                                                    v134 := GenS8{F0: int8(v18)}
                                                                    var v135 float32 = (1.5 / 3.14159)
                                                                    v136 := GenS8{F0: -1}
                                                                    _ = v134
                                                                    _ = v135
                                                                    _ = v136
                                                                } else {
                                                                    v137 := GenS6{F0: false, F1: 42}
                                                                    _ = v137
                                                                }
                                                            }
                                                        }
                                                        if (false || (v19 != (v19 + v19))) {
                                                            var v138 uint64 = (v18 << ((886) & 63))
                                                            switch "" {
                                                            case "foo bar":
                                                                switch "hello" {
                                                                case "a":
                                                                    fmt.Println(v18)
                                                                }
                                                            case "hello":
                                                                var v139 int16 = v108
                                                                _ = v139
                                                            case "":
                                                                if (117 >= 0) {
                                                                    if (-9223372036854775808 > 0) {
                                                                        fmt.Println(v18)
                                                                        for fi140 := 0; fi140 < 3; fi140++ {
                                                                            var v141 int8 = -123
                                                                            v138 = v138
                                                                            _ = fi140
                                                                            _ = v141
                                                                        }
                                                                        v142 := []int{971}
                                                                        _ = v142
                                                                    } else {
                                                                        if true {
                                                                            v138 = v18
                                                                        }
                                                                        fmt.Println(v18, v19)
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v138)
                                                            }
                                                            _ = v138
                                                        } else {
                                                            var v143 int16 = v108
                                                            _ = v143
                                                        }
                                                    }
                                                    fmt.Println(v19, v108)
                                                    _ = v108
                                                } else {
                                                    v144 := []int{1, -729, -393}
                                                    _ = v144
                                                }
                                            }
                                            v18 = (v18 - v18)
                                        } else {
                                            v19 = v19
                                            if true {
                                                var v145 float32 = 1.5
                                                v18 = v18
                                                _ = v145
                                            } else {
                                                v146 := []int{-164, 0, -789}
                                                _ = v146
                                            }
                                        }
                                        v147 := []int{128, 0, 42, 42}
                                        _ = v147
                                    }
                                default:
                                    for fi148 := 0; fi148 < 7; fi148++ {
                                        for fi149 := 0; fi149 < 10; fi149++ {
                                            var v150 int32 = -313
                                            if true {
                                                if false {
                                                    v19 = v19
                                                    var v151 float32 = 1.5
                                                    _ = v151
                                                }
                                            }
                                            v152 := GenS13{F0: 2, F1: 127}
                                            _ = fi149
                                            _ = v150
                                            _ = v152
                                        }
                                        v18 = (v18 << ((922) & 63))
                                        var v153 uint32 = 906
                                        _ = fi148
                                        _ = v153
                                    }
                                }
                                var v154 int16 = 1
                                _ = v154
                            }
                            _ = v19
                        } else {
                            var v155 float32 = 0.0
                            fmt.Println(v155, v155)
                            var v156 uint64 = 482
                            _ = v155
                            _ = v156
                        }
                        _ = v18
                    }
                }
                var v157 uint16 = 160
                v158 := []int{-1}
                _ = v157
                _ = v158
            } else {
                var v159 bool = ((float32((1e10 - 1e10)) <= 0.1) && ((((((("hello" + ("foo bar" + "hello")) + "") + "") + "") + "") + "") >= ("hello" + "a")))
                var v160 uint64 = 42
                for fi161 := 0; fi161 < 2; fi161++ {
                    if ((v160 << ((1) & 63)) >= v160) {
                        var v162 int64 = 0
                        v163 := GenS14{F0: 437, F1: fi161, F2: -1}
                        _ = v162
                        _ = v163
                    } else {
                        if v159 {
                            switch "\n\t" {
                            case "\n\t":
                                var v164 bool = true
                                _ = v164
                            default:
                                fmt.Println(v159, fi161, fi161)
                            }
                            v159 = (v160 < v160)
                            v165 := GenS15{F0: -32768, F1: 0}
                            _ = v165
                        }
                        var v166 int32 = 2147483647
                        if (3.14159 <= 1e10) {
                            fmt.Println(v166)
                        } else {
                            var v167 float64 = 0.0
                            v159 = v159
                            _ = v167
                        }
                        _ = v166
                    }
                    _ = fi161
                }
                _ = v159
                _ = v160
            }
        }
    }
    var v168 uint16 = 723
    for fi169 := 0; fi169 < 20; fi169++ {
        v168 -= v168
        _ = fi169
    }
    _ = v168
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ("hello" + ("foo bar" + ""))
    var v3 uint32 = 2147483648
    var v4 int16 = 0
    var v5 uint8 = 2
    if (!(102 == -128)) {
        switch "a" {
        case "hello":
            var v170 uint = 65535
            _ = v170
        default:
            v171 := GenS16{F0: 24, F1: false}
            _ = v171
        }
        fmt.Println(v4, v3, v5)
        fmt.Println(v3, v3)
    }
    v172 := []int{42, 4294967295, 42}
    v2 = "foo bar"
    switch "\n\t" {
    case "hello":
        var v173 float32 = 3.14159
        _ = v173
    case "\n\t":
        for fi174 := 0; fi174 < 4; fi174++ {
            v2 = ("" + ("a" + v2))
            if len(v172) > 0 {
                v172[0] = 4294967295
            }
            _ = fi174
        }
    case "":
        var v175 int16 = (v4 - (v4 & (^(v4))))
        _ = v175
    default:
        for fi176 := 0; fi176 < 5; fi176++ {
            var v177 string = "\n\t"
            _ = fi176
            _ = v177
        }
    }
    if len(v172) >= 2 {
        v172 = v172[1:len(v172)]
    }
    var v178 bool = true
    v178 = v178
    delete(v1, "hello")
    var v179 int64 = 0
    var v180 bool = (((((((("a" + v2) + v2) + "hello") + ((v2 + v2) + v2)) + v2) + ("" + (v2 + "hello"))) != v2) && (v178 && v178))
    v181 := GenS12{F0: 65535}
    var v182 float64 = 3.14159
    var v183 int16 = v4
    v180 = v178
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v172:", len(v172), v172)
    fmt.Println("v178:", v178)
    fmt.Println("v179:", v179)
    fmt.Println("v180:", v180)
    fmt.Println("v182:", v182)
    fmt.Println("v183:", v183)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v172
    _ = v178
    _ = v179
    _ = v180
    _ = v181
    _ = v182
    _ = v183
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
