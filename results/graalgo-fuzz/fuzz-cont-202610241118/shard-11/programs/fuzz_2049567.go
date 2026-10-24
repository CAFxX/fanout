package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 uint8; F2 int64 }
type GenS4 struct { F0 int; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int; F1 uint32; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint }
type GenS7 struct { F0 uint64; F1 bool; F2 int64 }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 string }
type GenS16 struct { F0 uint; F1 int32; F2 uint }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int; F1 uint32; F2 uint16 }
type GenS19 struct { F0 bool; F1 int }
type GenS20 struct { F0 int8; F1 uint32; F2 int64 }
type GenS21 struct { F0 string; F1 string }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (1.5 + 1.5)
    var v3 uint8 = 127
    var v4 string = ("foo bar" + "foo bar")
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 18; fi6++ {
        if true {
            switch (fi6 % (((fi6 << ((4294967295) & 63))) | 1)) {
            case -237:
                if (((!(false && false)) || false) && ((true && false) && ((^(v5)) >= v5))) {
                    v7 := GenS1{F0: (fi6 - -2147483648)}
                    v8 := GenS2{F0: 233, F1: (!(((((v2 - v2) - (v2 + (v2 - v2))) + (1.5 + (v2 * v2))) < ((((v2 * float64(1.5)) + 0.0) - float64(1e10)) / v2)) || true))}
                    _ = v7
                    _ = v8
                } else {
                    switch fi6 {
                    case 42:
                        if (false && (((-2.25 - v2) + v2) == v2)) {
                            for fi9 := 0; fi9 < 19; fi9++ {
                                v10 := GenS1{F0: (fi6 - (fi9 + fi6))}
                                _ = fi9
                                _ = v10
                            }
                        }
                    case 32768:
                        for fi11 := 0; fi11 < 2; fi11++ {
                            v12 := []int{(fi6 + fi6), fi6, fi6}
                            v13 := GenS3{F0: -2147483648, F1: v3, F2: -9223372036854775808}
                            if ((v4 + ((v4 + v4) + v4)) != "") {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                fmt.Println(v2, fi6, v4)
                                v14 := GenS3{F0: fi6, F1: v3, F2: 1}
                                _ = v14
                            }
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    if true {
                        v3 -= 255
                    }
                }
            case 496:
                var v15 int = (fi6 * 1)
                _ = v15
            }
            if false {
                switch v4 {
                case "foo bar":
                    v2 += 1e10
                default:
                    var v16 float32 = (-2.25 - (0.1 - (float32((-2.25 / (((v2 - v2) * v2) / 3.14159))) * float32(1.5))))
                    _ = v16
                }
                var v17 uint = 0
                _ = v17
            }
            v4 = "foo bar"
        } else {
            if false {
                v4 = v4
                var v18 uint32 = 2
                var v19 float32 = (-2.25 + 3.14159)
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 int32 = 2147483647
                switch v4 {
                case "a":
                    if (!(32767 < 0)) {
                        v22 := GenS2{F0: 255, F1: ((1.5 - v2) == (0.0 * v2))}
                        switch (^(fi20)) {
                        case 690:
                            switch fi6 {
                            case 486:
                                if false {
                                    var v23 int = fi20
                                    if true {
                                        v24 := GenS1{F0: fi6}
                                        v25 := v1[1]
                                        _ = v24
                                        _ = v25
                                    }
                                    fmt.Println(v21, v5, v5)
                                    _ = v23
                                }
                            }
                        case 776:
                            if (!true) {
                                switch "" {
                                case "\n\t":
                                    v26 := GenS2{F0: (v3 * (v3 >> ((1) & 63))), F1: (-2.25 > ((0.0 / (-2.25 + (float32(v2) * (3.14159 + (-2.25 / 1.0))))) * 3.14159))}
                                    _ = v26
                                case "foo bar":
                                    var v27 int = fi20
                                    _ = v27
                                case "hello":
                                    v28 := []int{fi20, fi20, 9223372036854775807, (fi6 * (fi6 % ((fi20) | 1)))}
                                    _ = v28
                                }
                                v21 = -2147483648
                                switch int(v5) {
                                case -641:
                                    var v29 int16 = 32767
                                    _ = v29
                                case 127:
                                    v21 -= v21
                                case 676:
                                    v5 -= (v5 << ((255) & 63))
                                default:
                                    fmt.Println(v4)
                                }
                                v30 := GenS4{F0: fi20, F1: 0, F2: (v5 & (v5 - v5))}
                                v31 := GenS4{F0: fi20, F1: 1, F2: (v5 | 0)}
                                var v32 int32 = v21
                                v4 = ((v4 + v4) + "")
                                v33 := GenS1{F0: (fi6 - (fi20 << ((191) & 63)))}
                                _ = v30
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                if false {
                                    switch v4 {
                                    case "a":
                                        var v34 int32 = v21
                                        _ = v34
                                    case "foo bar":
                                        switch 1 {
                                        case 32768:
                                            if false {
                                                v5 = v5
                                                fmt.Println(v5, v4, v4)
                                                if (true && true) {
                                                    fmt.Println(v2, v2)
                                                    v35 := []int{fi6, (fi6 - (fi20 / ((fi20) | 1)))}
                                                    _ = v35
                                                }
                                            }
                                            v36 := GenS1{F0: 2147483647}
                                            var v37 string = ("a" + v4)
                                            if true {
                                                if (fi20 != 1) {
                                                    if (false || true) {
                                                        if (!(-704 < 42)) {
                                                            switch (v4 + "a") {
                                                            case "foo bar":
                                                                v38 := GenS5{F0: 318, F1: 267, F2: v5}
                                                                _ = v38
                                                            case "":
                                                                var v39 uint16 = 707
                                                                _ = v39
                                                            }
                                                            var v40 uint32 = v5
                                                            if true {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v41 uint8 = (v3 - v3)
                                                                    _ = v41
                                                                case "hello":
                                                                    var v42 int16 = 441
                                                                    _ = v42
                                                                default:
                                                                    v5 -= 65535
                                                                }
                                                                v43 := GenS5{F0: fi20, F1: 552, F2: v5}
                                                                _ = v43
                                                            }
                                                            _ = v40
                                                        } else {
                                                            switch (fi6 >> ((42) & 63)) {
                                                            case 128:
                                                                v5 = v5
                                                            case -2147483648:
                                                                v44 := GenS3{F0: (fi6 / ((fi20) | 1)), F1: 42, F2: -9223372036854775808}
                                                                _ = v44
                                                            }
                                                            v2 = v2
                                                        }
                                                        v21 = 0
                                                        v3 -= v3
                                                        var v45 string = (v37 + "foo bar")
                                                        _ = v45
                                                    } else {
                                                        var v46 int8 = 100
                                                        _ = v46
                                                    }
                                                    if false {
                                                        var v47 int32 = v21
                                                        var v48 int8 = -38
                                                        v5 += 1
                                                        var v49 float64 = v2
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        var v50 int = (fi20 | (fi20 & fi20))
                                                        _ = v50
                                                    }
                                                } else {
                                                    var v51 int = 42
                                                    var v52 int32 = 1
                                                    if (255 != uint8(v5)) {
                                                        var v53 int16 = 1000
                                                        v2 += ((0.1 + 3.14159) + float64(1.5))
                                                        _ = v53
                                                    }
                                                    _ = v51
                                                    _ = v52
                                                }
                                                var v54 int16 = -32768
                                                v54 -= 0
                                                v21 += 2147483647
                                                _ = v54
                                            }
                                            var v55 int64 = -1
                                            _ = v36
                                            _ = v37
                                            _ = v55
                                        case 42:
                                            var v56 int8 = -1
                                            _ = v56
                                        default:
                                            if true {
                                                switch fi6 {
                                                case 128:
                                                    fmt.Println(fi6, v5)
                                                case -216:
                                                    var v57 float32 = (3.14159 * -2.25)
                                                    _ = v57
                                                case 32767:
                                                    v58 := GenS3{F0: fi6, F1: (v3 & v3), F2: -9223372036854775808}
                                                    _ = v58
                                                }
                                                if false {
                                                    if true {
                                                        v59 := GenS4{F0: -291, F1: 242, F2: 0}
                                                        _ = v59
                                                    } else {
                                                        v5 = 505
                                                    }
                                                    fmt.Println(v2)
                                                    if false {
                                                        v60 := GenS3{F0: (fi20 ^ -582), F1: (v3 - v3), F2: 0}
                                                        if (v5 > (v5 % ((0) | 1))) {
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(v21)
                                                                v61 := GenS5{F0: fi6, F1: v5, F2: v5}
                                                                _ = v61
                                                            default:
                                                                delete(v1, fi6)
                                                            }
                                                            v62 := GenS6{F0: false, F1: 336, F2: 2}
                                                            _ = v62
                                                        } else {
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v63 int32 = (v21 * (^(v21)))
                                                                _ = v63
                                                            case "":
                                                                v64 := GenS3{F0: fi20, F1: v3, F2: -1}
                                                                _ = v64
                                                            default:
                                                                v1[fi20] = (-(fi6))
                                                            }
                                                        }
                                                        _ = v60
                                                    } else {
                                                        v65 := GenS4{F0: (fi20 + fi20), F1: 4294967295, F2: v5}
                                                        _ = v65
                                                    }
                                                } else {
                                                    v66 := GenS7{F0: 2, F1: true, F2: -604}
                                                    _ = v66
                                                }
                                                if (32767 <= 471) {
                                                    if false {
                                                        var v67 string = "a"
                                                        fmt.Println(fi6, v2, fi6)
                                                        if (621 <= 979) {
                                                            v2 = (v2 / v2)
                                                            if true {
                                                                fmt.Println(fi6, v5, v5)
                                                            }
                                                            v68 := GenS4{F0: -1, F1: 18446744073709551615, F2: v5}
                                                            _ = v68
                                                        } else {
                                                            switch fi20 {
                                                            case 9223372036854775807:
                                                                if true {
                                                                    v2 -= v2
                                                                    switch v4 {
                                                                    case "":
                                                                        v69 := GenS5{F0: -2147483648, F1: (^(v5)), F2: 65535}
                                                                        _ = v69
                                                                    case "a":
                                                                        switch 127 {
                                                                        case 372:
                                                                            var v70 float64 = (1e10 - v2)
                                                                            _ = v70
                                                                        case -341:
                                                                            v71 := GenS8{F0: 2147483648, F1: (v21 ^ (v21 >> ((0) & 63)))}
                                                                            _ = v71
                                                                        case 1:
                                                                            v72 := GenS9{F0: (v3 & (v3 - (v3 / (((^(v3))) | 1)))), F1: -9223372036854775808}
                                                                            var v73 string = "a"
                                                                            v74 := []int{fi6, (-(fi20)), fi20}
                                                                            _ = v72
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                    case "foo bar":
                                                                        v4 = v67
                                                                    default:
                                                                        v1[-9223372036854775808] = (^(fi20))
                                                                    }
                                                                } else {
                                                                    v75 := GenS10{F0: 0, F1: -95}
                                                                    if (!(2 > 242)) {
                                                                        v76 := []int{}
                                                                        _ = v76
                                                                    } else {
                                                                        v1[fi6] = 127
                                                                        v77 := GenS11{F0: "foo bar", F1: 1}
                                                                        _ = v77
                                                                    }
                                                                    switch fi20 {
                                                                    case 32768:
                                                                        var v78 int64 = -688
                                                                        _ = v78
                                                                    case 9223372036854775807:
                                                                        var v79 float64 = 1e10
                                                                        _ = v79
                                                                    default:
                                                                        v67 = (v67 + (v4 + "a"))
                                                                    }
                                                                    _ = v75
                                                                }
                                                            case -745:
                                                                switch fi20 {
                                                                case 352:
                                                                    if false {
                                                                        v21 = (v21 - (v21 + (-(v21))))
                                                                        v3 += (v3 / ((v3) | 1))
                                                                    }
                                                                default:
                                                                    if true {
                                                                        v80 := GenS12{F0: 65535}
                                                                        switch fi20 {
                                                                        case -9223372036854775808:
                                                                            var v81 int8 = -115
                                                                            _ = v81
                                                                        default:
                                                                            if true {
                                                                                v82 := GenS11{F0: "a", F1: -1}
                                                                                v83 := []int{-1, 42}
                                                                                if false {
                                                                                    v5 += v5
                                                                                    var v84 uint = 152
                                                                                    _ = v84
                                                                                }
                                                                                v85 := GenS3{F0: fi20, F1: v3, F2: 1}
                                                                                if len(v83) > 0 {
                                                                                    v83[0] = (fi20 & (fi20 ^ (-(fi20))))
                                                                                }
                                                                                _ = v82
                                                                                _ = v83
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                        _ = v80
                                                                    } else {
                                                                        v86 := v1[(fi20 & (fi20 & 42))]
                                                                        if false {
                                                                            if true {
                                                                                v21 += 1
                                                                            } else {
                                                                                var v87 int32 = -2147483648
                                                                                if ((!(1e10 != 1.5)) || true) {
                                                                                    var v88 uint8 = (v3 | v3)
                                                                                    _ = v88
                                                                                }
                                                                                _ = v87
                                                                            }
                                                                            if (v2 < v2) {
                                                                                if (!(!((^(fi20)) > v86))) {
                                                                                    fmt.Println(v3)
                                                                                    if false {
                                                                                        if (-1 == 32767) {
                                                                                            fmt.Println(v21)
                                                                                            fmt.Println(v2, fi6, fi6)
                                                                                        }
                                                                                        if true {
                                                                                            var v89 int8 = -89
                                                                                            _ = len(v1)
                                                                                            _ = v89
                                                                                        }
                                                                                    } else {
                                                                                        v90 := []int{}
                                                                                        if true {
                                                                                            v4 = "foo bar"
                                                                                            var v91 float32 = -2.25
                                                                                            var v92 uint16 = uint16(v3)
                                                                                            var v93 uint8 = v3
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                            _ = v93
                                                                                        }
                                                                                        _ = v90
                                                                                    }
                                                                                    v94 := GenS7{F0: 18446744073709551615, F1: (!true), F2: 306}
                                                                                    _ = v94
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v86, fi20)
                                                                                    var v95 float64 = (-2.25 * (0.0 - v2))
                                                                                    _ = v95
                                                                                } else {
                                                                                    v67 = v4
                                                                                    _ = len(v1)
                                                                                    if false {
                                                                                        v21 = -2147483648
                                                                                        var v96 uint16 = 255
                                                                                        _ = v96
                                                                                    }
                                                                                    v97 := GenS13{F0: -60}
                                                                                    v4 = (("hello" + (v4 + v67)) + v4)
                                                                                    if (!true) {
                                                                                        v98 := GenS14{F0: -56}
                                                                                        var v99 int32 = (v21 % ((v21) | 1))
                                                                                        v3 -= v3
                                                                                        _ = v98
                                                                                        _ = v99
                                                                                    } else {
                                                                                        var v100 uint64 = 1
                                                                                        v101 := v1[int(v5)]
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    }
                                                                                    var v102 int8 = 127
                                                                                    var v103 int8 = 1
                                                                                    if (!(1 != (v103 >> ((uint(v3)) & 63)))) {
                                                                                        v104 := GenS7{F0: 394, F1: (-955 == 779), F2: 42}
                                                                                        v102 = (v103 % ((-128) | 1))
                                                                                        _ = v104
                                                                                    }
                                                                                    if (false && (v3 > (v3 / (((v3 >> ((1) & 63))) | 1)))) {
                                                                                        if true {
                                                                                            v105 := GenS11{F0: v67, F1: 42}
                                                                                            _ = v105
                                                                                        } else {
                                                                                            v106 := []int{(^(v86)), (-(v86)), (^(v86)), v86}
                                                                                            v21 = -179
                                                                                            var v107 int64 = 266
                                                                                            var v108 uint8 = 128
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                    }
                                                                                    var v109 uint32 = (v5 >> ((16) & 63))
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                    _ = v109
                                                                                    _ = v97
                                                                                }
                                                                                v110 := GenS5{F0: (fi6 | 2147483647), F1: v5, F2: (v5 ^ (v5 + v5))}
                                                                                v111 := GenS15{F0: -309, F1: "hello"}
                                                                                v21 = v21
                                                                                _ = v110
                                                                                _ = v111
                                                                            }
                                                                            v2 += 0.1
                                                                        } else {
                                                                            var v112 uint8 = 42
                                                                            _ = v112
                                                                        }
                                                                        v113 := GenS8{F0: (v5 / ((uint32(v3)) | 1)), F1: v21}
                                                                        _ = v113
                                                                        _ = v86
                                                                    }
                                                                }
                                                            case 128:
                                                                switch ((v4 + v4) + ("hello" + "hello")) {
                                                                case "foo bar":
                                                                    fmt.Println(v21)
                                                                    switch -2147483648 {
                                                                    case 1:
                                                                        fmt.Println(v3, v3)
                                                                        var v114 string = (v67 + v67)
                                                                        v5 -= (v5 * 1)
                                                                        _ = v114
                                                                    case 123:
                                                                        switch (^(fi6)) {
                                                                        case -9223372036854775808:
                                                                            v21 = (-(v21))
                                                                        case 154:
                                                                            switch fi6 {
                                                                            case 32767:
                                                                                switch (fi6 * (-(fi20))) {
                                                                                case 4294967295:
                                                                                    fmt.Println(fi20, v4, v21)
                                                                                case 1:
                                                                                    fmt.Println(v67, fi20)
                                                                                case 128:
                                                                                    switch (fi20 + fi20) {
                                                                                    case -989:
                                                                                        var v115 uint64 = 504
                                                                                        _ = v115
                                                                                    case 1:
                                                                                        var v116 uint8 = v3
                                                                                        _ = v116
                                                                                    case -1:
                                                                                        v117 := v1[(fi6 << ((65535) & 63))]
                                                                                        _ = v117
                                                                                    default:
                                                                                        v3 = v3
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v5, v67, v2)
                                                                                }
                                                                                fmt.Println(fi6, v4, fi20)
                                                                                var v118 uint32 = (v5 << ((1) & 63))
                                                                                _ = v118
                                                                            case 0:
                                                                                v119 := GenS16{F0: 65535, F1: (v21 - v21), F2: 2}
                                                                                _ = v119
                                                                            case 42:
                                                                                var v120 uint64 = 0
                                                                                _ = v120
                                                                            default:
                                                                                var v121 int = (-(fi20))
                                                                                _ = v121
                                                                            }
                                                                        case 614:
                                                                            _ = len(v1)
                                                                        default:
                                                                            if false {
                                                                                if true {
                                                                                    var v122 bool = true
                                                                                    _ = v122
                                                                                } else {
                                                                                    switch v4 {
                                                                                    case "foo bar":
                                                                                        v123 := GenS17{F0: fi20}
                                                                                        _ = v123
                                                                                    default:
                                                                                        v124 := []int{125, fi6, 128, (-(fi20))}
                                                                                        v4 = v4
                                                                                        _ = v124
                                                                                    }
                                                                                }
                                                                                v125 := GenS18{F0: -88, F1: (^(v5)), F2: 32767}
                                                                                var v126 float32 = 1.5
                                                                                _ = v125
                                                                                _ = v126
                                                                            }
                                                                            if (0 == 255) {
                                                                                v127 := GenS17{F0: fi6}
                                                                                v128 := GenS19{F0: true, F1: 995}
                                                                                var v129 float64 = -2.25
                                                                                _ = v127
                                                                                _ = v128
                                                                                _ = v129
                                                                            } else {
                                                                                if true {
                                                                                    var v130 uint64 = 234
                                                                                    fmt.Println(fi20)
                                                                                    _ = v130
                                                                                }
                                                                                v21 = (v21 | v21)
                                                                            }
                                                                        }
                                                                    case -1:
                                                                        v5 = 42
                                                                    }
                                                                    var v131 uint32 = 2
                                                                    _ = v131
                                                                case "":
                                                                    if true {
                                                                        if (int(v5) < fi20) {
                                                                            switch -9223372036854775808 {
                                                                            case 0:
                                                                                v2 = -2.25
                                                                            default:
                                                                                switch 32767 {
                                                                                case 1:
                                                                                    v67 = v4
                                                                                    fmt.Println(v21, v67, fi20)
                                                                                    var v132 int32 = -818
                                                                                    v133 := GenS16{F0: 0, F1: (v132 / ((-451) | 1)), F2: 4294967295}
                                                                                    var v134 int64 = 0
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                case -1:
                                                                                    v5 -= 433
                                                                                case -649:
                                                                                    v21 += (v21 << ((1) & 63))
                                                                                default:
                                                                                    var v135 float32 = 1e10
                                                                                    _ = v135
                                                                                }
                                                                            }
                                                                            var v136 uint16 = 2
                                                                            var v137 int16 = 822
                                                                            _ = v136
                                                                            _ = v137
                                                                        } else {
                                                                            var v138 uint64 = 637
                                                                            v139 := []int{-554, (fi6 % ((fi6) | 1)), (-(fi20))}
                                                                            v3 -= (v3 | v3)
                                                                            v139 = append(v139, 2147483647)
                                                                            fmt.Println(v21, v4)
                                                                            if true {
                                                                                v2 += (3.14159 - (v2 + 0.0))
                                                                                switch ((("" + (((v67 + v4) + "foo bar") + (((v4 + v67) + "") + (v67 + v4)))) + ("" + v67)) + ((("a" + v4) + v4) + "hello")) {
                                                                                case "\n\t":
                                                                                    v138 = v138
                                                                                case "":
                                                                                    if (587 >= -413) {
                                                                                        _ = len(v1)
                                                                                        var v140 bool = true
                                                                                        fmt.Println(v2)
                                                                                        _ = len(v1)
                                                                                        var v141 int64 = 985
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                    }
                                                                                case "hello":
                                                                                    switch (fi6 >> ((4294967295) & 63)) {
                                                                                    case 42:
                                                                                        v142 := GenS17{F0: fi20}
                                                                                        _ = v142
                                                                                    case 2147483647:
                                                                                        v139 = append(v139, fi6)
                                                                                    case -2147483648:
                                                                                        v143 := GenS14{F0: -49}
                                                                                        _ = v143
                                                                                    default:
                                                                                        fmt.Println(fi20, v4, fi6)
                                                                                    }
                                                                                }
                                                                                if (!(!(!false))) {
                                                                                    v144 := GenS20{F0: int8(fi20), F1: (^(v5)), F2: -910}
                                                                                    v67 = (v67 + ((v4 + "\n\t") + v67))
                                                                                    v145 := GenS8{F0: (v5 / (((v5 / (((v5 & (^(v5)))) | 1))) | 1)), F1: (v21 - (v21 >> ((255) & 63)))}
                                                                                    _ = v144
                                                                                    _ = v145
                                                                                }
                                                                                var v146 uint = 570
                                                                                v147 := GenS6{F0: true, F1: v138, F2: (v146 + 1)}
                                                                                _ = v146
                                                                                _ = v147
                                                                            } else {
                                                                                switch -315 {
                                                                                case -2147483648:
                                                                                    var v148 int8 = -128
                                                                                    _ = v148
                                                                                case 127:
                                                                                    if (((3.14159 / v2) * (v2 * (v2 - v2))) == (v2 + 3.14159)) {
                                                                                        v4 = v4
                                                                                    }
                                                                                }
                                                                            }
                                                                            v149 := GenS16{F0: 0, F1: v21, F2: 2}
                                                                            var v150 uint32 = (^(v5))
                                                                            _ = v138
                                                                            _ = v139
                                                                            _ = v149
                                                                            _ = v150
                                                                        }
                                                                    }
                                                                case "a":
                                                                    if (1 == 32767) {
                                                                        switch -1 {
                                                                        case 1:
                                                                            if (!false) {
                                                                                v5 = 769
                                                                                if ((!true) && true) {
                                                                                    switch (fi20 + (fi6 % ((2147483647) | 1))) {
                                                                                    case -1:
                                                                                        v151 := GenS16{F0: 1, F1: (v21 << ((0) & 63)), F2: 1}
                                                                                        _ = v151
                                                                                    case 42:
                                                                                        var v152 uint8 = 127
                                                                                        _ = v152
                                                                                    }
                                                                                    if false {
                                                                                        v21 = -2147483648
                                                                                    } else {
                                                                                        v5 += (^(v5))
                                                                                    }
                                                                                }
                                                                                v153 := GenS7{F0: 1, F1: (1 <= 0), F2: 9223372036854775807}
                                                                                switch v4 {
                                                                                case "a":
                                                                                    v154 := []int{fi6, (fi6 * (fi20 >> ((537) & 63))), fi20, (fi20 - (fi6 + fi20))}
                                                                                    _ = v154
                                                                                }
                                                                                var v155 int32 = 42
                                                                                _ = v153
                                                                                _ = v155
                                                                            }
                                                                        case -2147483648:
                                                                            switch (fi6 & (fi20 | fi20)) {
                                                                            case 686:
                                                                                v156 := GenS15{F0: 863, F1: "a"}
                                                                                _ = v156
                                                                            }
                                                                        default:
                                                                            v157 := []int{(fi20 % ((fi6) | 1)), -1, (fi6 | fi6), (fi20 % ((-673) | 1))}
                                                                            _ = v157
                                                                        }
                                                                        v158 := GenS19{F0: false, F1: (fi6 | -2147483648)}
                                                                        _ = v158
                                                                    } else {
                                                                        delete(v1, fi20)
                                                                    }
                                                                }
                                                            }
                                                            var v159 string = v67
                                                            _ = v159
                                                        }
                                                        _ = v67
                                                    }
                                                    var v160 float64 = (1.5 * -2.25)
                                                    _ = v160
                                                }
                                            } else {
                                                if false {
                                                    fmt.Println(v3, v21, v21)
                                                    var v161 uint64 = 42
                                                    if (32767 <= 32767) {
                                                        var v162 int8 = 0
                                                        v162 = (^(v162))
                                                        var v163 uint64 = v161
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    _ = v161
                                                } else {
                                                    fmt.Println(v21, v2)
                                                    switch v4 {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    case "":
                                                        fmt.Println(v4, v4)
                                                    case "hello":
                                                        var v164 uint32 = v5
                                                        _ = v164
                                                    }
                                                    var v165 bool = (true || (1000 < int16(v5)))
                                                    v4 = "hello"
                                                    var v166 float32 = float32(v2)
                                                    v167 := GenS18{F0: -742, F1: 1, F2: 0}
                                                    _ = v165
                                                    _ = v166
                                                    _ = v167
                                                }
                                            }
                                        }
                                        fmt.Println(v2)
                                        v168 := GenS21{F0: (v4 + v4), F1: v4}
                                        _ = v168
                                    }
                                    var v169 float32 = 0.1
                                    _ = v169
                                } else {
                                    v3 += 0
                                }
                            }
                        case 2147483647:
                            v2 += (1.5 - 0.0)
                        }
                        v1[fi20] = fi6
                        _ = v22
                    }
                default:
                    var v170 uint64 = 0
                    _ = v170
                }
                fmt.Println(v4, v2)
                _ = fi20
                _ = v21
            }
        }
        _ = fi6
    }
    if false {
        fmt.Println(v3, v3)
    } else {
        var v171 int = 451
        for fi172 := 0; fi172 < 12; fi172++ {
            fmt.Println(v3, fi172, v171)
            v173 := GenS22{F0: v5}
            v174 := v1[fi172]
            _ = fi172
            _ = v173
            _ = v174
        }
        _ = v171
    }
    if true {
        v3 = v3
        for fi175 := 0; fi175 < 13; fi175++ {
            for fi176 := 0; fi176 < 2; fi176++ {
                v2 = v2
                v3 += (v3 >> ((255) & 63))
                _ = fi176
            }
            switch -99 {
            case 128:
                v3 = v3
            }
            for fi177 := 0; fi177 < 5; fi177++ {
                v178 := GenS7{F0: 42, F1: (true && false), F2: -9223372036854775808}
                _ = fi177
                _ = v178
            }
            _ = fi175
        }
    }
    if false {
        fmt.Println(v5, v2, v2)
        var v179 string = ("\n\t" + v4)
        var v180 uint8 = (^(v3))
        _ = v179
        _ = v180
    }
    for fi181 := 0; fi181 < 3; fi181++ {
        _ = len(v1)
        _ = fi181
    }
    v3 += v3
    v182 := GenS23{F0: 0, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v182
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
