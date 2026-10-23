package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int64; F2 bool }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 int8; F2 uint8 }
type GenS12 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS13 struct { F0 bool; F1 bool; F2 uint }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{65535, 398}
    v3 := make(map[string]string)
    v4 := v1[153]
    v4 += -1
    if false {
        var v5 int8 = 127
        for fi6 := 0; fi6 < 19; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: 1, F1: 2147483647, F2: 100}
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        _ = v5
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9, v4)
        v4 -= (fi9 * (v4 / ((fi9) | 1)))
        delete(v1, 378)
        _ = fi9
    }
    fmt.Println(v4)
    v4 -= v4
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS2{F0: "", F1: 42}
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 bool = ((true || (18446744073709551615 <= 2)) || false)
            var v14 float64 = (((((3.14159 - 0.1) + 0.1) / (0.1 * ((-2.25 * float64((((0.1 * (-2.25 + (0.0 - float32(3.14159)))) / ((0.0 + 0.0) + 3.14159)) - float32(0.1)))) * (1e10 - ((1e10 / 1e10) - 3.14159))))) * ((1e10 + 1.5) - (0.0 * float64((float32(3.14159) + 1.5))))) - 0.0)
            var v15 uint16 = 255
            v13 = (100 == 255)
            var v16 int16 = 0
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        fmt.Println(fi10, v4)
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "foo bar":
        v4 += 32767
    case "\n\t":
        if true {
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS3{F0: -1}
                _ = fi17
                _ = v18
            }
        } else {
            var v19 int32 = -129
            v19 = (v19 / (((v19 - (v19 / (((v19 & 0)) | 1)))) | 1))
            for fi20 := 0; fi20 < 10; fi20++ {
                v1[2147483648] = v4
                var v21 float32 = 0.1
                switch (fi20 << ((0) & 63)) {
                case 0:
                    v21 -= (-2.25 * v21)
                case 90:
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v21 = (0.1 * v21)
                        v4 = 32767
                        _ = len(v3)
                        _ = fi22
                    }
                case 571:
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    var v23 bool = (false || (1 != 2))
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v24 string = ("\n\t" + "\n\t")
    if (v23 && (!true)) {
        switch v24 {
        case "\n\t":
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint64 = 42
                var v27 uint32 = 42
                var v28 bool = (255 >= uint(v26))
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        default:
            switch (v4 * (v4 - (^(v4)))) {
            case 32767:
                for fi29 := 0; fi29 < 17; fi29++ {
                    v24 = v24
                    switch (v4 & (fi29 / (((fi29 & (-(fi29)))) | 1))) {
                    case 2147483647:
                        v24 = v24
                    case -9223372036854775808:
                        for fi30 := 0; fi30 < 11; fi30++ {
                            v4 += (fi30 << ((528) & 63))
                            v31 := GenS2{F0: v24, F1: -887}
                            _ = fi30
                            _ = v31
                        }
                    }
                    for fi32 := 0; fi32 < 18; fi32++ {
                        if (v23 || v23) {
                            if v23 {
                                fmt.Println(fi32, fi32)
                                if v23 {
                                    var v33 int32 = -2147483648
                                    _ = v33
                                } else {
                                    var v34 int = 128
                                    v35 := GenS1{F0: 1, F1: -837, F2: 1}
                                    _ = v34
                                    _ = v35
                                }
                            }
                            var v36 float64 = ((0.0 - 1.5) - (((((((1.5 / 1.0) / -2.25) - float64((3.14159 - (0.1 + 0.1)))) - ((float64((1e10 - (1.5 / 1.5))) + float64((3.14159 / 1e10))) * float64(-2.25))) * (float64(0.1) * (-2.25 / 1.5))) - 0.1) + 1.5))
                            _ = v36
                        } else {
                            var v37 uint64 = 42
                            if (-583 <= -1) {
                                if (-2147483648 >= 114) {
                                    delete(v1, 2)
                                    v4 = (fi32 & 32768)
                                    v38 := GenS3{F0: -2147483648}
                                    v39 := v1[920]
                                    var v40 uint16 = 42
                                    fmt.Println(v24, v37)
                                    var v41 int8 = 0
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    v37 += 642
                                    if len(v2) > 0 {
                                        v2[0] = 851
                                    }
                                }
                                fmt.Println(fi32, v23)
                                v24 = v24
                            }
                            v42 := GenS3{F0: 738}
                            _ = v37
                            _ = v42
                        }
                        if ((0 < 1) || v23) {
                            var v43 uint64 = 4294967295
                            v44 := GenS4{F0: 98}
                            if ((v43 + 848) < (v43 ^ v43)) {
                                v4 += fi32
                                var v45 uint = 18446744073709551615
                                var v46 uint = v45
                                var v47 uint16 = 255
                                switch fi32 {
                                case 616:
                                    _ = len(v2)
                                case 817:
                                    if v23 {
                                        _ = len(v2)
                                        var v48 float32 = 3.14159
                                        var v49 int32 = -561
                                        v48 += (1e10 / (3.14159 - float32(float64(((v48 / (float32(0.0) * v48)) / 1.5)))))
                                        v47 -= (v47 % ((v47) | 1))
                                        var v50 uint64 = (v43 ^ v43)
                                        v23 = false
                                        _ = len(v2)
                                        var v51 int32 = (v49 & int32(v50))
                                        fmt.Println(v4, fi32, v50)
                                        if false {
                                            _ = len(v2)
                                            if (!v23) {
                                                var v52 int64 = 1
                                                _ = v52
                                            } else {
                                                var v53 float32 = (0.1 * 1.5)
                                                v49 = (v51 ^ (v51 / ((0) | 1)))
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v53
                                            }
                                            switch v24 {
                                            case "a":
                                                if v23 {
                                                    v48 = (v48 - (v48 - 3.14159))
                                                    v49 -= (v49 ^ (v51 & (v49 ^ 0)))
                                                    if false {
                                                        _ = len(v3)
                                                        switch v24 {
                                                        case "\n\t":
                                                            switch fi29 {
                                                            case 9223372036854775807:
                                                                v54 := GenS5{F0: -957}
                                                                _ = v54
                                                            case 0:
                                                                if len(v2) > 0 {
                                                                    v2[0] = 2147483647
                                                                }
                                                            default:
                                                                var v55 int8 = -52
                                                                _ = v55
                                                            }
                                                        case "foo bar":
                                                            if v23 {
                                                                v56 := GenS1{F0: v47, F1: 536, F2: -18}
                                                                if v23 {
                                                                    var v57 uint32 = 65535
                                                                    v45 -= (^(v46))
                                                                    var v58 int32 = (v49 | -709)
                                                                    var v59 uint = v45
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                                fmt.Println(v23, v4)
                                                                _ = v56
                                                            } else {
                                                                v60 := GenS6{F0: v23, F1: -9223372036854775808, F2: ((v51 | 2147483647) > -1)}
                                                                v49 -= (v51 ^ -1)
                                                                _ = v60
                                                            }
                                                        case "a":
                                                            fmt.Println(v4, v46, v43)
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                        v3[(("foo bar" + v24) + v24)] = "\n\t"
                                                    }
                                                }
                                            default:
                                                var v61 uint32 = 1
                                                _ = v61
                                            }
                                            switch fi29 {
                                            case 248:
                                                if (-456 >= (v4 * (v4 * (v4 | 128)))) {
                                                    _ = len(v2)
                                                }
                                            case 918:
                                                v62 := GenS3{F0: (^(v49))}
                                                _ = v62
                                            }
                                        } else {
                                            v49 -= (v49 * (^(v51)))
                                            var v63 int8 = -1
                                            _ = v63
                                        }
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v64 := GenS6{F0: true, F1: 42, F2: (int64(v47) <= 42)}
                                        _ = v64
                                    }
                                }
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            _ = v43
                            _ = v44
                        }
                        if v23 {
                            v24 = (v24 + v24)
                            v1[0] = 2147483647
                            var v65 uint32 = 952
                            _ = v65
                        }
                        switch "hello" {
                        case "foo bar":
                            delete(v3, (("a" + (v24 + (v24 + (v24 + v24)))) + (v24 + (((v24 + (v24 + v24)) + (("hello" + "a") + ("" + v24))) + (v24 + "\n\t")))))
                        }
                        _ = fi32
                    }
                    _ = fi29
                }
            default:
                for fi66 := 0; fi66 < 6; fi66++ {
                    var v67 string = v24
                    if false {
                        v24 = v67
                        for fi68 := 0; fi68 < 15; fi68++ {
                            fmt.Println(fi66, v24, v23)
                            var v69 int32 = -974
                            _ = len(v2)
                            _ = fi68
                            _ = v69
                        }
                    }
                    v70 := GenS7{F0: 1000, F1: (0 < -1)}
                    _ = fi66
                    _ = v67
                    _ = v70
                }
            }
        }
        switch v4 {
        case 32768:
            switch v4 {
            case -1:
                for fi71 := 0; fi71 < 9; fi71++ {
                    var v72 uint16 = 42
                    v23 = (false && false)
                    if v23 {
                        var v73 int = (fi71 + (fi71 % ((fi71) | 1)))
                        _ = len(v2)
                        for fi74 := 0; fi74 < 20; fi74++ {
                            fmt.Println(v23)
                            _ = fi74
                        }
                        _ = v73
                    } else {
                        if (0 != 979) {
                            for fi75 := 0; fi75 < 19; fi75++ {
                                if (v24 >= (("hello" + "hello") + v24)) {
                                    if (0 == 65535) {
                                        if (!true) {
                                            if len(v2) > 0 {
                                                v2[0] = 2147483647
                                            }
                                            var v76 uint8 = 1
                                            _ = v76
                                        } else {
                                            switch (v24 + v24) {
                                            case "hello":
                                                if (v24 != "") {
                                                    v77 := GenS1{F0: v72, F1: 2147483647, F2: -1}
                                                    switch v24 {
                                                    case "a":
                                                        var v78 int16 = -1
                                                        switch (fi71 - (^(v4))) {
                                                        case 42:
                                                            switch -1 {
                                                            case -2147483648:
                                                                var v79 int16 = (v78 + (v78 & v78))
                                                                _ = v79
                                                            case 86:
                                                                v80 := GenS8{F0: 0, F1: ""}
                                                                _ = v80
                                                            }
                                                        case 2147483647:
                                                            switch (fi71 >> ((42) & 63)) {
                                                            case 0:
                                                                v24 = ""
                                                                if false {
                                                                    v81 := GenS9{F0: -1, F1: -1}
                                                                    _ = v81
                                                                } else {
                                                                    delete(v1, 615)
                                                                    v1[2147483648] = (v4 / (((fi75 << ((836) & 63))) | 1))
                                                                    v78 += (v78 + v78)
                                                                    v72 = v72
                                                                    v82 := GenS10{F0: 100}
                                                                    _ = v82
                                                                }
                                                            case 388:
                                                                v83 := GenS11{F0: 9223372036854775807, F1: -70, F2: 0}
                                                                _ = v83
                                                            }
                                                        default:
                                                            v78 += v78
                                                        }
                                                        _ = len(v2)
                                                        _ = v78
                                                    case "\n\t":
                                                        if true {
                                                            if true {
                                                                if ((v23 || (v23 && v23)) && v23) {
                                                                    v23 = (!(42 != 675))
                                                                    v84 := GenS3{F0: -1}
                                                                    _ = v84
                                                                } else {
                                                                    _ = len(v1)
                                                                    switch ((v24 + (v24 + v24)) + v24) {
                                                                    case "foo bar":
                                                                        v4 -= (^(fi71))
                                                                    }
                                                                }
                                                                fmt.Println(v72, v4)
                                                                v85 := GenS6{F0: (fi75 != (fi71 % (((fi71 | (v4 / (((^(fi71))) | 1)))) | 1))), F1: 990, F2: (v24 >= (v24 + (v24 + v24)))}
                                                                delete(v3, "foo bar")
                                                                if false {
                                                                    var v86 uint64 = 359
                                                                    _ = v86
                                                                } else {
                                                                    v87 := GenS2{F0: (((("a" + "a") + "a") + ("\n\t" + v24)) + ("" + (v24 + v24))), F1: 42}
                                                                    _ = v87
                                                                }
                                                                _ = v85
                                                            }
                                                            if (-1 != 42) {
                                                                var v88 int64 = 0
                                                                if (((!((148 != 2147483648) && (111 < 162))) && (v88 == (v88 % (((v88 - -9223372036854775808)) | 1)))) && false) {
                                                                    if v23 {
                                                                        v89 := v3[v24]
                                                                        v90 := GenS1{F0: 616, F1: -1, F2: 27}
                                                                        var v91 string = "a"
                                                                        _ = v89
                                                                        _ = v90
                                                                        _ = v91
                                                                    } else {
                                                                        switch 201 {
                                                                        case 127:
                                                                            v72 = 0
                                                                        }
                                                                    }
                                                                    if v23 {
                                                                        if ((!((v88 << ((4294967295) & 63)) == (^(v88)))) && v23) {
                                                                            if ((!v23) || (42 >= 18446744073709551615)) {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v92 uint32 = 4294967295
                                                                                var v93 uint32 = v92
                                                                                var v94 float32 = 0.0
                                                                                _ = v92
                                                                                _ = v93
                                                                                _ = v94
                                                                            } else {
                                                                                switch v24 {
                                                                                case "hello":
                                                                                    var v95 bool = (v23 || v23)
                                                                                    _ = v95
                                                                                case "foo bar":
                                                                                    switch -27 {
                                                                                    case 2147483647:
                                                                                        if v23 {
                                                                                            v72 = (v72 >> ((4294967295) & 63))
                                                                                            var v96 int32 = -2147483648
                                                                                            _ = v96
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        v97 := GenS5{F0: (v88 + (v88 / ((v88) | 1)))}
                                                                                        _ = v97
                                                                                    case -76:
                                                                                        v72 = v72
                                                                                    default:
                                                                                        var v98 int = 2147483647
                                                                                        _ = v98
                                                                                    }
                                                                                default:
                                                                                    v99 := GenS6{F0: (false || v23), F1: 1, F2: v23}
                                                                                    _ = v99
                                                                                }
                                                                                var v100 uint32 = 42
                                                                                if (true && false) {
                                                                                    v101 := GenS12{F0: 255, F1: v23, F2: 349}
                                                                                    _ = len(v3)
                                                                                    _ = v101
                                                                                } else {
                                                                                    v100 -= 606
                                                                                    fmt.Println(v88)
                                                                                }
                                                                                var v102 float32 = 1e10
                                                                                _ = v100
                                                                                _ = v102
                                                                            }
                                                                            fmt.Println(v24, fi71, fi75)
                                                                            v103 := GenS8{F0: 255, F1: v24}
                                                                            if v23 {
                                                                                v72 -= 65535
                                                                                var v104 string = v24
                                                                                v105 := v1[1]
                                                                                var v106 int = -1
                                                                                fmt.Println(v4, fi75)
                                                                                var v107 bool = (42 < 1)
                                                                                _ = v104
                                                                                _ = v105
                                                                                _ = v106
                                                                                _ = v107
                                                                            } else {
                                                                                fmt.Println(v23)
                                                                                switch (fi75 / ((fi75) | 1)) {
                                                                                case -2147483648:
                                                                                    var v108 float64 = -2.25
                                                                                    _ = v108
                                                                                case 4294967295:
                                                                                    if (fi75 == (-(v4))) {
                                                                                        var v109 uint = 1
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                                v110 := GenS6{F0: (v23 || ((-2.25 == -2.25) && false)), F1: v88, F2: v23}
                                                                                var v111 float32 = (0.1 / (-2.25 - 0.0))
                                                                                fmt.Println(v4)
                                                                                _ = v110
                                                                                _ = v111
                                                                            }
                                                                            if false {
                                                                                var v112 int8 = -112
                                                                                _ = v112
                                                                            } else {
                                                                                fmt.Println(v4, v72)
                                                                                v113 := GenS5{F0: v88}
                                                                                if v23 {
                                                                                    var v114 uint64 = 4294967295
                                                                                    v24 = "hello"
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        v88 += v88
                                                                                    default:
                                                                                        switch (v24 + ((v24 + "\n\t") + v24)) {
                                                                                        case "\n\t":
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                var v115 uint8 = 42
                                                                                                _ = v115
                                                                                            case "foo bar":
                                                                                                v72 = (^(v72))
                                                                                            default:
                                                                                                v24 = v24
                                                                                            }
                                                                                        case "":
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        case "hello":
                                                                                            var v116 string = ""
                                                                                            _ = v116
                                                                                        default:
                                                                                            v114 -= 42
                                                                                        }
                                                                                    }
                                                                                    _ = v114
                                                                                }
                                                                                _ = v113
                                                                            }
                                                                            _ = v103
                                                                        } else {
                                                                            v117 := GenS7{F0: 42, F1: v23}
                                                                            _ = v117
                                                                        }
                                                                        var v118 int = -1
                                                                        fmt.Println(v24)
                                                                        if true {
                                                                            switch (^(v4)) {
                                                                            case 1:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case 4294967295:
                                                                                v119 := GenS6{F0: v23, F1: (v88 | (v88 & v88)), F2: (886 != 216)}
                                                                                _ = v119
                                                                            }
                                                                            v4 = fi71
                                                                            v24 = "hello"
                                                                        }
                                                                        fmt.Println(v24, v72)
                                                                        _ = v118
                                                                    }
                                                                    v2 = append(v2, 747)
                                                                } else {
                                                                    switch (^(v4)) {
                                                                    case -9223372036854775808:
                                                                        var v120 uint64 = 484
                                                                        _ = v120
                                                                    case 128:
                                                                        fmt.Println(v72)
                                                                    default:
                                                                        v24 = v24
                                                                    }
                                                                    if (v23 && true) {
                                                                        if (((1.5 - 0.0) + ((1.5 / float64(float32(((3.14159 / 1.0) - (((3.14159 + 0.1) * 1.5) + 1.5))))) + -2.25)) <= ((-2.25 / (1.5 / 1.5)) * 1.5)) {
                                                                            switch v24 {
                                                                            case "foo bar":
                                                                                switch fi75 {
                                                                                case 623:
                                                                                    v121 := GenS9{F0: v88, F1: 9223372036854775807}
                                                                                    _ = v121
                                                                                }
                                                                            }
                                                                        }
                                                                        var v122 int8 = 127
                                                                        _ = v122
                                                                    }
                                                                }
                                                                switch v24 {
                                                                case "\n\t":
                                                                    if (128 <= 42) {
                                                                        v24 = (((v24 + "") + v24) + v24)
                                                                    }
                                                                    v123 := GenS13{F0: (128 <= 255), F1: v23, F2: 18446744073709551615}
                                                                    _ = v123
                                                                }
                                                                _ = v88
                                                            } else {
                                                                var v124 uint = 65535
                                                                if (v72 < 42) {
                                                                    v72 = v72
                                                                } else {
                                                                    if v23 {
                                                                        _ = len(v2)
                                                                        v125 := GenS10{F0: 127}
                                                                        _ = v125
                                                                    }
                                                                    v24 = "\n\t"
                                                                }
                                                                _ = v124
                                                            }
                                                        } else {
                                                            v4 += fi75
                                                            fmt.Println(v72)
                                                            var v126 uint32 = 469
                                                            _ = v126
                                                        }
                                                    case "foo bar":
                                                        v72 += (v72 | v72)
                                                    default:
                                                        var v127 float32 = 0.1
                                                        _ = v127
                                                    }
                                                    v23 = false
                                                    var v128 uint8 = 72
                                                    _ = v128
                                                    _ = v77
                                                }
                                            case "foo bar":
                                                switch (("" + v24) + ("hello" + ("\n\t" + (v24 + v24)))) {
                                                case "foo bar":
                                                    v129 := GenS5{F0: 42}
                                                    _ = v129
                                                case "hello":
                                                    v24 = v24
                                                default:
                                                    v23 = (false && v23)
                                                    switch 903 {
                                                    case 9223372036854775807:
                                                        v72 -= 2
                                                    case 261:
                                                        v2 = append(v2, 4294967295)
                                                    case -707:
                                                        if (int16(fi71) >= -1) {
                                                            switch "" {
                                                            case "hello":
                                                                v72 += (v72 & v72)
                                                            case "foo bar":
                                                                v4 += (fi75 % ((fi75) | 1))
                                                            case "\n\t":
                                                                var v130 uint64 = 1
                                                                _ = v130
                                                            default:
                                                                fmt.Println(v72, v4, v72)
                                                                v131 := GenS6{F0: ((-2.25 - 3.14159) != (((-2.25 / float64(1e10)) - (1e10 / (((0.0 - -2.25) / 0.1) + 1.5))) + -2.25)), F1: 0, F2: ((v72 - (v72 - (v72 % ((v72) | 1)))) >= v72)}
                                                                if false {
                                                                    v132 := GenS4{F0: 100}
                                                                    fmt.Println(v23, v24, v24)
                                                                    _ = v132
                                                                } else {
                                                                    var v133 int16 = -1
                                                                    fmt.Println(v24, v24)
                                                                    if v23 {
                                                                        v23 = (-352 >= 0)
                                                                        v1[2] = (fi75 >> ((4294967295) & 63))
                                                                    }
                                                                    if v23 {
                                                                        var v134 uint8 = 2
                                                                        var v135 uint16 = 32767
                                                                        v2 = append(v2, 1)
                                                                        _ = v134
                                                                        _ = v135
                                                                    } else {
                                                                        var v136 int64 = 1
                                                                        var v137 uint = 1
                                                                        _ = v136
                                                                        _ = v137
                                                                    }
                                                                    _ = v133
                                                                }
                                                                _ = v131
                                                            }
                                                            v138 := GenS7{F0: 0, F1: (42 == 547)}
                                                            v24 = "a"
                                                            _ = v138
                                                        } else {
                                                            switch fi75 {
                                                            case 32767:
                                                                v139 := v3["a"]
                                                                _ = v139
                                                            case 128:
                                                                fmt.Println(v23)
                                                            default:
                                                                v4 = -2147483648
                                                            }
                                                            v72 -= (v72 + v72)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                    switch (-(v4)) {
                                                    case 127:
                                                        var v140 int64 = 9223372036854775807
                                                        _ = v140
                                                    default:
                                                        var v141 float32 = 1e10
                                                        _ = v141
                                                    }
                                                }
                                            case "\n\t":
                                                var v142 bool = false
                                                _ = v142
                                            default:
                                                v24 = (("foo bar" + v24) + v24)
                                            }
                                        }
                                    }
                                    var v143 uint32 = 42
                                    v24 = "a"
                                    if len(v2) > 0 {
                                        v2[0] = 551
                                    }
                                    switch ("hello" + "a") {
                                    case "":
                                        var v144 int64 = -596
                                        _ = v144
                                    default:
                                        var v145 int32 = 0
                                        var v146 float32 = 0.0
                                        v147 := GenS14{F0: 0, F1: 42}
                                        _ = v145
                                        _ = v146
                                        _ = v147
                                    }
                                    _ = v143
                                } else {
                                    var v148 int8 = 42
                                    _ = v148
                                }
                                v4 = 0
                                _ = fi75
                            }
                            fmt.Println(v23, v24, v23)
                        }
                    }
                    _ = fi71
                    _ = v72
                }
            case 0:
                v149 := GenS12{F0: 255, F1: (!false), F2: 2}
                _ = v149
            case 1:
                var v150 uint16 = 1
                _ = v150
            default:
                if (1 >= 62) {
                    fmt.Println(v23)
                    v151 := GenS3{F0: 1}
                    var v152 int = (v4 + 638)
                    _ = v151
                    _ = v152
                }
            }
        case -9223372036854775808:
            v153 := GenS15{F0: 128, F1: 18446744073709551615}
            _ = v153
        case -2147483648:
            for fi154 := 0; fi154 < 10; fi154++ {
                switch fi154 {
                case 42:
                    for fi155 := 0; fi155 < 11; fi155++ {
                        v156 := v3[v24]
                        _ = fi155
                        _ = v156
                    }
                case 127:
                    for fi157 := 0; fi157 < 12; fi157++ {
                        var v158 uint64 = 42
                        fmt.Println(v4)
                        if v23 {
                            var v159 bool = v23
                            fmt.Println(v4, v159)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v159
                        } else {
                            var v160 uint64 = (v158 | (^(v158)))
                            v161 := GenS5{F0: 42}
                            switch 288 {
                            case -9223372036854775808:
                                v2 = append(v2, 2)
                            }
                            _ = v160
                            _ = v161
                        }
                        _ = fi157
                        _ = v158
                    }
                case -1:
                    v162 := GenS8{F0: 0, F1: (("foo bar" + "a") + (v24 + "\n\t"))}
                    _ = v162
                }
                var v163 int16 = 1
                _ = fi154
                _ = v163
            }
        }
        var v164 uint64 = 601
        _ = v164
    } else {
        v2 = append(v2, 2147483648)
        fmt.Println(v4)
    }
    var v165 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v165:", v165)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v165
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
