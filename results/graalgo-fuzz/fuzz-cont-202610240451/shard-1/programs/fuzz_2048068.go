package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 bool; F1 int64; F2 uint8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 bool; F2 int16 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint64; F1 string; F2 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 string; F2 int }
type GenS13 struct { F0 int32; F1 int }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 bool; F1 uint32 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint16 }

func main() {
    var v1 bool = true
    var v2 uint8 = 1
    var v3 float32 = (1e10 + 3.14159)
    v4 := make(map[int]int)
    var v5 uint16 = 32767
    v6 := make(map[uint32]string)
    v2 += v2
    v7 := []int{127, 127, 0, 127}
    if (!v1) {
        if v1 {
            if (v5 >= 0) {
                v2 = v2
                if (1 != -2147483648) {
                    v1 = v1
                }
                v3 += v3
            } else {
                v7 = append(v7, 0)
                var v8 int16 = -568
                _ = v8
            }
            if false {
                fmt.Println(v3, v2)
            } else {
                v9 := GenS1{F0: -1, F1: v5}
                var v10 uint16 = (^(v5))
                var v11 int16 = -1
                _ = v10
                _ = v11
                _ = v9
            }
        }
        var v12 int16 = 1000
        _ = v12
    } else {
        switch int(v5) {
        case 783:
            v5 = (v5 * 0)
        }
        if (42 >= 9223372036854775807) {
            v13 := v4[708]
            _ = len(v7)
            v1 = ((v1 && v1) && true)
            _ = v13
        }
        if false {
            if v1 {
                if (3.14159 < (0.1 - 1.5)) {
                    var v14 int64 = 15
                    if v1 {
                        v15 := v4[127]
                        _ = v15
                    }
                    v2 = v2
                    _ = v14
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    if (1000 <= 1000) {
                        fmt.Println(v3)
                    }
                    if false {
                        fmt.Println(v1)
                        for fi16 := 0; fi16 < 17; fi16++ {
                            v17 := GenS1{F0: 2147483647, F1: 2}
                            _ = fi16
                            _ = v17
                        }
                    }
                }
            }
            switch 2147483647 {
            case 192:
                v1 = (2147483647 <= 2)
            case 128:
                for fi18 := 0; fi18 < 20; fi18++ {
                    v4[fi18] = fi18
                    var v19 int64 = 9223372036854775807
                    _ = fi18
                    _ = v19
                }
            default:
                var v20 int8 = 7
                _ = v20
            }
        } else {
            fmt.Println(v3)
            var v21 uint8 = (v2 >> ((587) & 63))
            switch ("\n\t" + "a") {
            case "":
                for fi22 := 0; fi22 < 9; fi22++ {
                    fmt.Println(v1, v21, v5)
                    v5 = (^(v5))
                    v23 := GenS1{F0: 2147483647, F1: (^(v5))}
                    _ = fi22
                    _ = v23
                }
            case "\n\t":
                if len(v7) > 0 {
                    v7[0] = 32768
                }
            case "foo bar":
                v3 -= float32(((1e10 * (1e10 + 0.0)) + ((float64(v3) + 1e10) * 0.1)))
            }
            _ = v21
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int16 = -1
        var v26 uint8 = v2
        delete(v4, fi24)
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 float32 = (-2.25 / (v3 + (float32((float64(float32(((1e10 - 0.0) + (((3.14159 - -2.25) - 0.0) * float64(v3))))) / 1.0)) - (v3 + -2.25))))
    switch -476 {
    case 128:
        for fi28 := 0; fi28 < 4; fi28++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            switch "" {
            case "a":
                v4[fi28] = (fi28 - 127)
            case "hello":
                var v29 int64 = -1
                _ = v29
            default:
                if false {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v1 = true
                        _ = fi30
                    }
                }
            }
            if v1 {
                fmt.Println(v1, v3)
                for fi31 := 0; fi31 < 20; fi31++ {
                    v32 := GenS2{F0: (v1 || ((("hello" + "foo bar") + "a") < ""))}
                    if (((1.5 / (-2.25 - 1.5)) / 1.5) == float64(1.5)) {
                        var v33 float64 = -2.25
                        _ = v33
                    }
                    v6[451] = "\n\t"
                    _ = fi31
                    _ = v32
                }
                v27 = (1e10 + v3)
            } else {
                v5 = 2
            }
            _ = fi28
        }
    case 32768:
        var v34 uint = 4294967295
        _ = v34
    case 103:
        v5 += v5
    default:
        switch ("" + "") {
        case "a":
            var v35 int16 = 0
            _ = v35
        default:
            v36 := GenS3{F0: v2, F1: 42}
            _ = v36
        }
    }
    v5 -= v5
    switch "" {
    case "foo bar":
        v2 -= v2
    case "hello":
        v37 := GenS2{F0: (v1 || (!(v1 || v1)))}
        _ = v37
    case "\n\t":
        v1 = true
    default:
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 float32 = v27
            for fi40 := 0; fi40 < 13; fi40++ {
                var v41 uint64 = uint64(v2)
                delete(v4, (fi38 >> ((255) & 63)))
                if (1000 <= 0) {
                    switch "foo bar" {
                    case "a":
                        var v42 int8 = 0
                        _ = v42
                    case "hello":
                        if ((v2 >> ((42) & 63)) < v2) {
                            if v1 {
                                v41 = 42
                            } else {
                                fmt.Println(v3, v2)
                                if v1 {
                                    if (!((-17 >= 100) && (-1 != 42))) {
                                        var v43 uint16 = 1
                                        v44 := GenS4{F0: 1, F1: 4294967295}
                                        switch "hello" {
                                        case "a":
                                            fmt.Println(v5)
                                        case "\n\t":
                                            v27 += 1.5
                                        default:
                                            if true {
                                                v45 := GenS2{F0: v1}
                                                _ = v45
                                            } else {
                                                var v46 uint = 89
                                                if true {
                                                    _ = len(v6)
                                                    v47 := GenS1{F0: -2147483648, F1: (v5 >> ((v46) & 63))}
                                                    _ = v47
                                                } else {
                                                    v2 += 127
                                                    v2 += 1
                                                    v48 := GenS1{F0: 42, F1: v5}
                                                    _ = v48
                                                }
                                                v6[65535] = ""
                                                _ = v46
                                            }
                                        }
                                        _ = v43
                                        _ = v44
                                    } else {
                                        v49 := GenS1{F0: -182, F1: 65535}
                                        v50 := GenS5{F0: (!((int16(v41) == 1) || (v1 && true))), F1: -9223372036854775808, F2: (^(v2))}
                                        _ = v49
                                        _ = v50
                                    }
                                    v51 := GenS1{F0: -1, F1: 32767}
                                    if v1 {
                                        v52 := GenS3{F0: v2, F1: 0}
                                        v3 = 0.0
                                        v53 := GenS4{F0: 2, F1: 0}
                                        _ = v52
                                        _ = v53
                                    } else {
                                        fmt.Println(fi40)
                                    }
                                    if true {
                                        v2 = v2
                                        if ((!v1) || v1) {
                                            var v54 uint8 = (v2 % (((v2 * (v2 << ((0) & 63)))) | 1))
                                            if len(v7) > 0 {
                                                v7[0] = 533
                                            }
                                            _ = v54
                                        } else {
                                            v3 += v3
                                            v2 += v2
                                            v5 = 255
                                        }
                                        switch "" {
                                        case "a":
                                            var v55 uint8 = v2
                                            _ = v55
                                        case "\n\t":
                                            v5 = v5
                                        }
                                        v27 = ((1.5 * 1e10) / 1e10)
                                        if ((!(uint64(v2) > (^(v41)))) && v1) {
                                            if false {
                                                fmt.Println(v3, v2, v2)
                                            } else {
                                                fmt.Println(v5)
                                                v27 = v39
                                                if (9223372036854775807 < -701) {
                                                    var v56 uint64 = v41
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    _ = v56
                                                }
                                            }
                                        }
                                    } else {
                                        v2 = 2
                                    }
                                    _ = v51
                                } else {
                                    if v1 {
                                        switch ("hello" + "foo bar") {
                                        case "hello":
                                            var v57 uint = 18446744073709551615
                                            _ = v57
                                        case "foo bar":
                                            if v1 {
                                                var v58 uint8 = v2
                                                v58 += (v2 | (v2 + (v2 | (v58 ^ (v2 / ((v2) | 1))))))
                                                _ = v58
                                            } else {
                                                v5 = v5
                                                if v1 {
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    var v59 int16 = 836
                                                    _ = v59
                                                }
                                                switch (fi38 % (((-(fi38))) | 1)) {
                                                case 917:
                                                    v41 += v41
                                                }
                                                var v60 int = (^(fi38))
                                                v61 := GenS2{F0: v1}
                                                v1 = v1
                                                _ = v60
                                                _ = v61
                                            }
                                        case "\n\t":
                                            v7 = append(v7, 32768)
                                        }
                                    } else {
                                        v62 := GenS1{F0: 0, F1: v5}
                                        _ = v62
                                    }
                                }
                            }
                            var v63 int8 = -1
                            v64 := GenS6{F0: 2147483647, F1: fi38}
                            _ = v63
                            _ = v64
                        }
                    }
                    switch (^(fi38)) {
                    case -854:
                        if len(v7) > 0 {
                            v7[0] = (fi38 / ((fi40) | 1))
                        }
                    case 4294967295:
                        if v1 {
                            v5 += v5
                        } else {
                            switch "" {
                            case "hello":
                                switch ("" + "\n\t") {
                                case "a":
                                    var v65 float64 = 1.5
                                    _ = v65
                                case "hello":
                                    switch 32768 {
                                    case 9223372036854775807:
                                        switch (fi38 >> ((255) & 63)) {
                                        case -862:
                                            delete(v4, fi40)
                                        case 42:
                                            var v66 int64 = 1
                                            _ = v66
                                        }
                                    }
                                case "\n\t":
                                    switch (fi40 & (fi38 ^ (fi38 ^ int(v2)))) {
                                    case 127:
                                        var v67 float64 = 3.14159
                                        _ = v67
                                    case -180:
                                        var v68 int64 = 9223372036854775807
                                        _ = v68
                                    case -2147483648:
                                        v3 += v27
                                    default:
                                        v69 := GenS1{F0: 2147483647, F1: (v5 | v5)}
                                        _ = v69
                                    }
                                }
                            }
                            v5 = 592
                        }
                    }
                    var v70 bool = v1
                    var v71 uint16 = 32767
                    if v1 {
                        switch ("" + "a") {
                        case "":
                            switch "\n\t" {
                            case "\n\t":
                                if (1000 <= -490) {
                                    v2 = 81
                                    var v72 int8 = 29
                                    v73 := GenS4{F0: 18446744073709551615, F1: 2147483648}
                                    _ = v72
                                    _ = v73
                                } else {
                                    if (0 > 0) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch ("hello" + ("a" + ("hello" + ""))) {
                                            case "foo bar":
                                                switch ("foo bar" + "hello") {
                                                case "":
                                                    if v70 {
                                                        fmt.Println(fi40)
                                                        v74 := GenS6{F0: 42, F1: (fi38 << ((255) & 63))}
                                                        var v75 float64 = (1.5 * 1e10)
                                                        _ = v74
                                                        _ = v75
                                                    } else {
                                                        var v76 bool = (2147483648 <= 4294967295)
                                                        var v77 int16 = 1000
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                case "a":
                                                    fmt.Println(fi38)
                                                case "hello":
                                                    var v78 uint16 = v71
                                                    _ = v78
                                                default:
                                                    fmt.Println(v39)
                                                    var v79 float64 = 0.1
                                                    _ = v79
                                                }
                                            case "\n\t":
                                                v3 = v27
                                            case "a":
                                                v7 = append(v7, (fi40 % ((fi38) | 1)))
                                            default:
                                                var v80 uint32 = 352
                                                _ = v80
                                            }
                                        case "foo bar":
                                            v71 += v5
                                        default:
                                            fmt.Println(v41, fi38, fi38)
                                        }
                                        v2 -= 128
                                    }
                                }
                            case "a":
                                var v81 int64 = 1
                                _ = v81
                            case "hello":
                                var v82 uint16 = 2
                                _ = v82
                            }
                        case "a":
                            if (v70 && v70) {
                                v5 = (v5 | uint16(v2))
                                fmt.Println(v27, v27)
                                v2 -= 128
                            } else {
                                v2 = 1
                                switch fi38 {
                                case 0:
                                    fmt.Println(v2, v1, v41)
                                case 4294967295:
                                    var v83 int8 = 0
                                    _ = v83
                                }
                            }
                        default:
                            v84 := GenS7{F0: (v2 | (v2 | 240))}
                            _ = v84
                        }
                        v39 += 0.0
                    }
                    var v85 int16 = -32768
                    if (v70 && v1) {
                        v7 = append(v7, (fi38 >> ((2) & 63)))
                        fmt.Println(v5, v3, v70)
                    } else {
                        if false {
                            var v86 uint16 = (v5 % (((v71 & 42)) | 1))
                            switch "" {
                            case "hello":
                                _ = len(v7)
                            }
                            v87 := GenS1{F0: 1, F1: v71}
                            _ = v86
                            _ = v87
                        } else {
                            if ("" < "\n\t") {
                                var v88 int32 = 641
                                var v89 uint = 2
                                if (122 <= 65535) {
                                    switch 32768 {
                                    case 630:
                                        if (-2147483648 <= (^(fi40))) {
                                            v70 = true
                                            if (v88 >= -45) {
                                                var v90 string = ""
                                                if len(v7) > 0 {
                                                    v7[0] = 575
                                                }
                                                var v91 int16 = (v85 * v85)
                                                v6[24] = v90
                                                var v92 string = v90
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                            }
                                        } else {
                                            v71 = uint16(v2)
                                        }
                                    case 128:
                                        if len(v7) > 0 {
                                            v7[0] = fi38
                                        }
                                        var v93 uint16 = (v5 & v5)
                                        _ = len(v4)
                                        v27 += v27
                                        _ = v93
                                    default:
                                        if true {
                                            v27 -= v27
                                            v89 += 42
                                            switch -1 {
                                            case 0:
                                                v94 := v6[391]
                                                _ = v94
                                            case 32768:
                                                fmt.Println(v71)
                                            default:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v95 float32 = (1e10 / (float32((((float64(1.5) * 1.5) * float64(v39)) - 3.14159)) / 3.14159))
                                                    _ = v95
                                                default:
                                                    var v96 int64 = 0
                                                    _ = v96
                                                }
                                            }
                                        } else {
                                            var v97 uint32 = 65535
                                            v98 := GenS8{F0: (v88 ^ (v88 / ((v88) | 1))), F1: (((float64(float32((0.0 * 1.5))) - ((-2.25 - 0.0) - float64(v3))) + (1e10 * (0.1 - ((((1e10 * 0.1) - -2.25) / 1e10) * (0.0 * float64(v27)))))) <= float64((v3 * v39))), F2: v85}
                                            _ = v97
                                            _ = v98
                                        }
                                    }
                                    v99 := GenS8{F0: v88, F1: true, F2: v85}
                                    v7 = append(v7, 0)
                                    switch 0 {
                                    case 2147483647:
                                        if (false && false) {
                                            var v100 float64 = 3.14159
                                            if v70 {
                                                if false {
                                                    if false {
                                                        if false {
                                                            v6[1] = "\n\t"
                                                        }
                                                        if len(v7) > 0 {
                                                            v7[0] = -2147483648
                                                        }
                                                        switch ("" + "a") {
                                                        case "":
                                                            v101 := GenS1{F0: (v88 << ((uint(v41)) & 63)), F1: v5}
                                                            _ = v101
                                                        }
                                                    } else {
                                                        var v102 int64 = 42
                                                        _ = v102
                                                    }
                                                    v103 := GenS8{F0: (v88 * 1), F1: (false || v70), F2: (v85 ^ (v85 | v85))}
                                                    switch 0 {
                                                    case -9223372036854775808:
                                                        switch (("\n\t" + "foo bar") + (("foo bar" + "a") + "hello")) {
                                                        case "\n\t":
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                        }
                                                    case 1:
                                                        if false {
                                                            v104 := v6[42]
                                                            v70 = v1
                                                            _ = len(v6)
                                                            _ = v104
                                                        } else {
                                                            v105 := GenS6{F0: (v88 & v88), F1: (fi40 & fi40)}
                                                            fmt.Println(v88)
                                                            var v106 int16 = 1
                                                            _ = v105
                                                            _ = v106
                                                        }
                                                    }
                                                    _ = v103
                                                }
                                                v70 = ((v2 >> ((v89) & 63)) == 127)
                                                var v107 float32 = (v39 + v39)
                                                v85 += v85
                                                delete(v6, uint32(v71))
                                                delete(v6, 35)
                                                _ = v107
                                            } else {
                                                if false {
                                                    v108 := GenS1{F0: v88, F1: (v5 % ((v71) | 1))}
                                                    _ = v108
                                                } else {
                                                    v88 = v88
                                                    v109 := GenS2{F0: (v1 && v1)}
                                                    switch ("foo bar" + "a") {
                                                    case "a":
                                                        v85 = (-(v85))
                                                    }
                                                    _ = v109
                                                }
                                                var v110 int32 = (v88 + v88)
                                                _ = v110
                                            }
                                            v85 = v85
                                            _ = v100
                                        }
                                    case 4294967295:
                                        v111 := GenS9{F0: (v2 - (v2 / ((v2) | 1))), F1: 42}
                                        _ = v111
                                    default:
                                        var v112 float32 = ((v27 - (0.1 / ((v3 + v39) * (v39 - v3)))) + float32(-2.25))
                                        _ = v112
                                    }
                                    _ = v99
                                } else {
                                    switch fi40 {
                                    case 42:
                                        v113 := GenS2{F0: v70}
                                        _ = v113
                                    case 692:
                                        v70 = v1
                                    }
                                }
                                _ = v88
                                _ = v89
                            } else {
                                switch fi40 {
                                case 1:
                                    switch fi40 {
                                    case 42:
                                        delete(v4, fi38)
                                    }
                                case 127:
                                    v114 := GenS10{F0: v41, F1: "\n\t", F2: 0}
                                    fmt.Println(v1, v39, v3)
                                    var v115 uint16 = v71
                                    _ = v114
                                    _ = v115
                                case 4294967295:
                                    v1 = (1 != -9223372036854775808)
                                default:
                                    v85 = 32767
                                }
                                var v116 uint16 = (v71 * (^(v71)))
                                if v1 {
                                    v117 := GenS6{F0: -3, F1: fi38}
                                    _ = v117
                                }
                                switch "a" {
                                case "foo bar":
                                    var v118 float32 = v3
                                    _ = v118
                                }
                                if false {
                                    var v119 int32 = 1
                                    var v120 uint8 = (^(v2))
                                    _ = v119
                                    _ = v120
                                } else {
                                    _ = len(v6)
                                    v121 := GenS11{F0: (v70 || (0 == 42))}
                                    v122 := GenS1{F0: 2147483647, F1: v5}
                                    _ = v121
                                    _ = v122
                                }
                                switch (-(fi40)) {
                                case 702:
                                    v71 = (^(v116))
                                    if ((("a" + "\n\t") == (("\n\t" + "foo bar") + "hello")) || (839 != 42)) {
                                        var v123 bool = (116 >= 201)
                                        _ = v123
                                    } else {
                                        var v124 string = ("hello" + "\n\t")
                                        v116 += (v71 % ((955) | 1))
                                        var v125 int = (fi40 - (^(fi40)))
                                        v126 := GenS12{F0: v2, F1: v124, F2: 127}
                                        _ = v124
                                        _ = v125
                                        _ = v126
                                    }
                                    var v127 int = 128
                                    var v128 uint64 = (^(v41))
                                    switch int(v71) {
                                    case -9223372036854775808:
                                        var v129 uint = 1
                                        _ = v129
                                    case 127:
                                        var v130 uint8 = (^(v2))
                                        _ = v130
                                    case 0:
                                        switch (fi40 % (((v127 - v127)) | 1)) {
                                        case 4294967295:
                                            v7 = append(v7, fi40)
                                        }
                                    default:
                                        v127 = -2147483648
                                    }
                                    _ = v127
                                    _ = v128
                                case -1:
                                    var v131 int64 = 735
                                    _ = v131
                                }
                                _ = v116
                            }
                            switch (("hello" + ("" + "hello")) + "\n\t") {
                            case "\n\t":
                                _ = len(v4)
                            }
                        }
                        if false {
                            var v132 uint8 = v2
                            v2 = v132
                            if ("\n\t" < "a") {
                                switch (fi38 + (fi38 / (((fi38 - 2147483647)) | 1))) {
                                case 32768:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                case 724:
                                    var v133 bool = ((-(fi38)) == fi38)
                                    _ = v133
                                }
                                v3 += v3
                            }
                            var v134 float32 = 0.0
                            v41 += v41
                            _ = v132
                            _ = v134
                        } else {
                            v85 -= 42
                        }
                        v39 = (1.5 + 1.5)
                    }
                    _ = v70
                    _ = v71
                    _ = v85
                } else {
                    switch (fi40 + 42) {
                    case 42:
                        _ = len(v6)
                    }
                    if (!v1) {
                        v135 := v6[829]
                        v135 = v135
                        v136 := GenS13{F0: 42, F1: (fi38 & 526)}
                        fmt.Println(v2)
                        _ = v135
                        _ = v136
                    } else {
                        v27 -= (((v3 * v27) / 0.1) * (1.5 + (((v3 - ((((v39 - v27) + 1e10) * v3) / 1.5)) * 0.0) * v3)))
                        if v1 {
                            var v137 int = 9223372036854775807
                            v3 -= v39
                            _ = v137
                        }
                    }
                    var v138 int32 = 890
                    _ = v138
                }
                v139 := GenS14{F0: "hello"}
                _ = len(v6)
                var v140 float32 = (float32(-2.25) / -2.25)
                _ = fi40
                _ = v139
                _ = v140
                _ = v41
            }
            var v141 int = (^(fi38))
            _ = fi38
            _ = v141
            _ = v39
        }
    }
    for fi142 := 0; fi142 < 18; fi142++ {
        if (fi142 != fi142) {
            v5 = 2
            v27 = v3
        } else {
            if v1 {
                v5 -= (v5 + v5)
                for fi143 := 0; fi143 < 12; fi143++ {
                    switch -421 {
                    case -2147483648:
                        switch "" {
                        case "foo bar":
                            if (("foo bar" + "foo bar") > "hello") {
                                switch ("hello" + "foo bar") {
                                case "":
                                    if (!(!true)) {
                                        v144 := GenS7{F0: (^(v2))}
                                        v5 = v5
                                        v145 := GenS15{F0: 2}
                                        _ = v144
                                        _ = v145
                                    }
                                case "hello":
                                    v3 -= (((1e10 + v3) - (float32(float64(1.5)) / v3)) / v3)
                                default:
                                    v146 := GenS16{F0: "foo bar", F1: "foo bar"}
                                    _ = v146
                                }
                                var v147 string = "\n\t"
                                _ = v147
                            } else {
                                switch "" {
                                case "\n\t":
                                    v148 := GenS11{F0: false}
                                    _ = v148
                                case "a":
                                    v149 := GenS17{F0: (942 < 986), F1: 529}
                                    _ = v149
                                case "foo bar":
                                    delete(v4, (fi142 & fi142))
                                }
                                v2 -= v2
                                v150 := GenS14{F0: "a"}
                                _ = v150
                            }
                        case "":
                            delete(v6, 65535)
                        case "hello":
                            var v151 bool = (595 <= 1)
                            _ = v151
                        default:
                            var v152 uint32 = 554
                            _ = v152
                        }
                    default:
                        fmt.Println(v1, v27, v2)
                    }
                    _ = fi143
                }
                v153 := GenS5{F0: true, F1: -806, F2: v2}
                _ = v153
            }
            fmt.Println(v1, v1)
            var v154 uint32 = 695
            _ = v154
        }
        if true {
            v2 = v2
            v6[183] = ("a" + "\n\t")
        } else {
            switch fi142 {
            case 42:
                fmt.Println(v3, v27)
            default:
                fmt.Println(v27)
            }
        }
        _ = fi142
    }
    fmt.Println(v2)
    v27 = v27
    if false {
        v7 = append(v7, -1)
        v155 := GenS13{F0: -1, F1: 0}
        _ = v155
    } else {
        var v156 int8 = -43
        _ = v156
    }
    fmt.Println(v5, v5)
    switch ("" + (("\n\t" + ("foo bar" + ("" + "a"))) + "")) {
    case "foo bar":
        var v157 float64 = 3.14159
        _ = v157
    case "\n\t":
        for fi158 := 0; fi158 < 3; fi158++ {
            v7 = append(v7, 2147483647)
            _ = len(v7)
            _ = fi158
        }
    case "a":
        switch ("" + "a") {
        case "":
            for fi159 := 0; fi159 < 8; fi159++ {
                switch (fi159 & fi159) {
                case -245:
                    switch "" {
                    case "":
                        for fi160 := 0; fi160 < 15; fi160++ {
                            v161 := GenS5{F0: ("foo bar" > "a"), F1: 42, F2: v2}
                            v1 = v1
                            _ = fi160
                            _ = v161
                        }
                    case "hello":
                        v162 := GenS14{F0: ""}
                        _ = v162
                    default:
                        _ = len(v4)
                    }
                }
                if len(v7) > 0 {
                    v7[0] = fi159
                }
                v2 -= v2
                _ = fi159
            }
        case "a":
            if v1 {
                if ((!true) && (0 > -1)) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    for fi163 := 0; fi163 < 19; fi163++ {
                        for fi164 := 0; fi164 < 13; fi164++ {
                            v165 := GenS11{F0: (!((v5 / ((255) | 1)) >= (v5 * (v5 ^ v5))))}
                            _ = fi164
                            _ = v165
                        }
                        for fi166 := 0; fi166 < 18; fi166++ {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            delete(v4, fi166)
                            var v167 int8 = 1
                            fmt.Println(v3, v3, v5)
                            v1 = v1
                            var v168 int32 = 2147483647
                            if false {
                                v169 := v4[-209]
                                if false {
                                    _ = len(v7)
                                    v1 = true
                                }
                                _ = v169
                            }
                            _ = fi166
                            _ = v167
                            _ = v168
                        }
                        switch ((("\n\t" + "") + "") + "foo bar") {
                        case "hello":
                            if v1 {
                                _ = len(v6)
                            }
                        case "":
                            var v170 int = fi163
                            _ = v170
                        case "a":
                            var v171 uint = 4294967295
                            _ = v171
                        default:
                            v172 := GenS10{F0: 4294967295, F1: "", F2: 0}
                            _ = v172
                        }
                        _ = fi163
                    }
                    v27 -= 3.14159
                }
                _ = len(v7)
            }
        default:
            switch (("a" + "a") + "foo bar") {
            case "a":
                if (((v5 + v5) >= v5) && (127 < -6)) {
                    v3 -= v3
                    _ = len(v7)
                } else {
                    for fi173 := 0; fi173 < 18; fi173++ {
                        if (1 < -2147483648) {
                            v6[0] = ("" + "")
                            v174 := GenS1{F0: 40, F1: v5}
                            var v175 int32 = -718
                            _ = v174
                            _ = v175
                        } else {
                            fmt.Println(v3)
                            if v1 {
                                v7 = append(v7, (fi173 + 9223372036854775807))
                                v1 = false
                                var v176 int32 = -1
                                _ = v176
                            } else {
                                for fi177 := 0; fi177 < 19; fi177++ {
                                    v178 := GenS9{F0: 8, F1: 510}
                                    _ = fi177
                                    _ = v178
                                }
                            }
                            v3 = (v27 / v3)
                        }
                        v179 := GenS18{F0: -35}
                        _ = fi173
                        _ = v179
                    }
                    fmt.Println(v2)
                    v180 := GenS19{F0: (v5 / (((v5 * v5)) | 1))}
                    _ = v180
                }
            case "":
                v7 = append(v7, 32768)
            default:
                v5 = v5
            }
        }
    }
    var v181 int32 = 42
    v182 := GenS4{F0: 858, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v181:", v181)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v181
    _ = v182
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
