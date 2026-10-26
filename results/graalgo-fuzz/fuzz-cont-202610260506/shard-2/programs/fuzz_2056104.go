package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint8 }
type GenS4 struct { F0 string; F1 uint16; F2 string }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 string; F2 int32 }
type GenS9 struct { F0 uint32; F1 string; F2 uint64 }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 int8; F1 int8; F2 string }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint; F1 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int32; F1 int32; F2 uint }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint32; F1 int16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int; F1 int64 }
type GenS21 struct { F0 int64; F1 uint64 }
type GenS22 struct { F0 uint16; F1 string }
type GenS23 struct { F0 int }
type GenS24 struct { F0 int32 }
type GenS25 struct { F0 bool; F1 uint8 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 bool; F1 bool; F2 int32 }
type GenS28 struct { F0 bool }
type GenS29 struct { F0 bool; F1 int32 }
type GenS30 struct { F0 int; F1 int; F2 uint64 }
type GenS31 struct { F0 string }
type GenS32 struct { F0 uint }
type GenS33 struct { F0 int64 }
type GenS34 struct { F0 uint16; F1 uint32 }
type GenS35 struct { F0 int32 }
type GenS36 struct { F0 int8; F1 uint64 }
type GenS37 struct { F0 int8; F1 int }
func genFn0(p0 uint) int64 {
    var v4 float64 = -2.25
    v5 := GenS1{F0: 42, F1: -128}
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint32 {
    if false {
        v6 := GenS2{F0: 1}
        var v7 uint = 1
        fmt.Println(v7)
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int16 = -170
            _ = fi8
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 5; fi10++ {
                if (((-2.25 / 3.14159) == (3.14159 / 0.1)) || true) {
                    if ((!(!(false && false))) && (true && (!((138 > 255) && true)))) {
                        v11 := GenS3{F0: 2, F1: "a", F2: 255}
                        v12 := GenS4{F0: "hello", F1: 255, F2: "\n\t"}
                        _ = v11
                        _ = v12
                    }
                }
                var v13 int16 = 32767
                var v14 uint16 = 956
                _ = fi10
                _ = v13
                _ = v14
            }
            var v15 float64 = 1.5
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if (255 != 128) {
                        switch (((("" + "a") + (("foo bar" + ("hello" + ("hello" + ""))) + ("foo bar" + ("foo bar" + ("foo bar" + ("" + "foo bar")))))) + "a") + (("" + "") + "")) {
                        case "foo bar":
                            v18 := []int{}
                            _ = v18
                        case "a":
                            v19 := []int{(-(fi17))}
                            _ = v19
                        default:
                            switch (fi17 / ((fi17) | 1)) {
                            case 2147483647:
                                v20 := GenS5{F0: 0, F1: 2, F2: 18446744073709551615}
                                _ = v20
                            }
                        }
                        switch fi16 {
                        case 285:
                            if (42 <= 100) {
                                v15 += v15
                                v15 += v15
                            }
                            if (!true) {
                                v21 := GenS6{F0: 196}
                                var v22 uint8 = 128
                                var v23 int64 = -9223372036854775808
                                if false {
                                    if (v22 <= v22) {
                                        fmt.Println(v23)
                                        v24 := []int{fi17, (^(fi16))}
                                        v22 = v22
                                        v22 += v22
                                        if (32767 <= -1) {
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        } else {
                                            fmt.Println(fi17)
                                            fmt.Println(fi17, v22, v23)
                                            fmt.Println(v23, fi16)
                                        }
                                        _ = v24
                                    }
                                    v25 := []int{(fi16 << ((65535) & 63))}
                                    switch 9223372036854775807 {
                                    case 716:
                                        if true {
                                            v22 += v22
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                        }
                                    case 629:
                                        switch "\n\t" {
                                        case "":
                                            var v26 int32 = 1
                                            _ = v26
                                        case "\n\t":
                                            v15 += float64(0.0)
                                        }
                                    case -2147483648:
                                        if (v22 == (^(v22))) {
                                            var v27 uint64 = 387
                                            _ = v27
                                        }
                                    }
                                    _ = v25
                                } else {
                                    fmt.Println(v23)
                                }
                                switch fi16 {
                                case -833:
                                    if false {
                                        v28 := []int{1, (fi17 & (fi16 << ((65535) & 63))), (fi17 << ((255) & 63)), (fi17 >> ((4294967295) & 63))}
                                        _ = v28
                                    } else {
                                        if false {
                                            v29 := GenS6{F0: -1}
                                            var v30 float64 = v15
                                            _ = v29
                                            _ = v30
                                        }
                                        v22 -= v22
                                        if true {
                                            v31 := []int{-742, (fi16 | (fi16 >> ((2) & 63)))}
                                            _ = v31
                                        } else {
                                            if (2147483648 >= 42) {
                                                var v32 string = ("foo bar" + "hello")
                                                _ = v32
                                            } else {
                                                if false {
                                                    v33 := []int{-838}
                                                    if false {
                                                        var v34 uint64 = 0
                                                        var v35 uint = 155
                                                        var v36 uint8 = v22
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    fmt.Println(v22)
                                                    _ = v33
                                                } else {
                                                    v37 := GenS6{F0: -1}
                                                    _ = v37
                                                }
                                            }
                                            if false {
                                                if ((!true) && (!(65535 == 682))) {
                                                    var v38 int8 = 0
                                                    v39 := []int{(-(fi17)), fi17, 9223372036854775807}
                                                    if len(v39) >= 2 {
                                                        v39 = v39[1:len(v39)]
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            fmt.Println(v15)
                                        }
                                        var v40 uint8 = v22
                                        fmt.Println(v15)
                                        var v41 uint16 = 0
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                v42 := GenS4{F0: "\n\t", F1: 334, F2: "a"}
                                if false {
                                    if (1000 != -1) {
                                        var v43 int16 = -32768
                                        _ = v43
                                    }
                                    var v44 int16 = 1000
                                    var v45 float32 = 0.0
                                    _ = v44
                                    _ = v45
                                } else {
                                    v46 := []int{(fi17 + (fi17 * fi17)), 0, 32768, (-(fi17))}
                                    var v47 int8 = -67
                                    if false {
                                        v48 := GenS1{F0: 0, F1: 1}
                                        if len(v46) > 0 {
                                            v46[0] = fi16
                                        }
                                        var v49 bool = false
                                        if v49 {
                                            v47 += (v47 & (v47 & (v47 - (v47 & v47))))
                                            var v50 uint8 = 0
                                            _ = v50
                                        } else {
                                            v22 = v22
                                        }
                                        _ = v48
                                        _ = v49
                                    }
                                    var v51 uint16 = 0
                                    v23 = -9223372036854775808
                                    _ = v46
                                    _ = v47
                                    _ = v51
                                }
                                _ = v21
                                _ = v22
                                _ = v23
                                _ = v42
                            } else {
                                switch "a" {
                                case "a":
                                    v15 = (v15 * v15)
                                }
                                v15 = 0.1
                            }
                        case 9223372036854775807:
                            var v52 uint8 = 128
                            _ = v52
                        default:
                            v15 -= v15
                        }
                        var v53 int32 = 2147483647
                        _ = v53
                    }
                    v54 := GenS3{F0: 860, F1: "a", F2: 161}
                    if false {
                        fmt.Println(fi17, v15, v15)
                        var v55 bool = (-723 != 42)
                        if true {
                            var v56 uint8 = 142
                            if ((32767 < 1000) && v55) {
                                var v57 float64 = 0.0
                                _ = v57
                            } else {
                                v58 := GenS7{F0: 42}
                                var v59 uint8 = v56
                                _ = v58
                                _ = v59
                            }
                            v55 = ("" == ("foo bar" + ("hello" + ("" + ("" + ("hello" + "a"))))))
                            _ = v56
                        }
                        _ = v55
                    }
                    v60 := GenS5{F0: 127, F1: 2, F2: 1}
                    _ = fi17
                    _ = v54
                    _ = v60
                }
                fmt.Println(fi16, v15)
                var v61 bool = (false && false)
                _ = fi16
                _ = v61
            }
            _ = v15
        } else {
            if true {
                var v62 uint64 = 4294967295
                var v63 int = -1
                switch -160 {
                case 9223372036854775807:
                    for fi64 := 0; fi64 < 3; fi64++ {
                        fmt.Println(v62, fi64, v63)
                        for fi65 := 0; fi65 < 6; fi65++ {
                            if true {
                                v66 := []int{}
                                if true {
                                    switch 32767 {
                                    case 680:
                                        if (1e10 < (1e10 + 1.5)) {
                                            v63 = -2147483648
                                            var v67 int64 = 9223372036854775807
                                            v68 := GenS8{F0: "\n\t", F1: "hello", F2: 0}
                                            _ = v67
                                            _ = v68
                                        }
                                    default:
                                        if len(v66) > 0 {
                                            v66[0] = (fi65 | (v63 & v63))
                                        }
                                    }
                                    if false {
                                        v62 += (v62 << ((18446744073709551615) & 63))
                                        v62 = (^(v62))
                                    }
                                }
                                if len(v66) >= 2 {
                                    v66 = v66[1:len(v66)]
                                }
                                _ = v66
                            } else {
                                if false {
                                    if false {
                                        switch "a" {
                                        case "foo bar":
                                            v69 := []int{-9223372036854775808, 128, (v63 / (((fi65 << ((18446744073709551615) & 63))) | 1)), (fi64 >> ((0) & 63))}
                                            _ = v69
                                        case "a":
                                            v70 := GenS4{F0: ("" + "\n\t"), F1: 791, F2: "\n\t"}
                                            _ = v70
                                        }
                                        v62 -= (v62 * v62)
                                        v71 := []int{v63, (v63 + v63), 1}
                                        fmt.Println(v62, v63)
                                        _ = v71
                                    } else {
                                        if (!(0 == uint16(v62))) {
                                            v63 = (fi64 % ((fi65) | 1))
                                            switch (("a" + "hello") + "") {
                                            case "hello":
                                                if (-32768 >= -650) {
                                                    if false {
                                                        if (!true) {
                                                            var v72 uint = 945
                                                            var v73 uint32 = 990
                                                            if true {
                                                                v74 := GenS9{F0: 4294967295, F1: "", F2: (^(v62))}
                                                                if (2147483647 != 2147483647) {
                                                                    v63 -= (fi64 << (((^(v72))) & 63))
                                                                } else {
                                                                    v73 -= uint32(v72)
                                                                    switch (fi65 | (v63 ^ 0)) {
                                                                    case 32767:
                                                                        var v75 bool = ((!((!((v73 & v73) <= v73)) && ((!(0 != -418)) || (!(!false))))) || false)
                                                                        _ = v75
                                                                    default:
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "hello":
                                                                            if true {
                                                                                v76 := []int{(-(fi65)), fi65}
                                                                                v77 := GenS10{F0: (("a" + "a") + "hello"), F1: fi65}
                                                                                v73 += v73
                                                                                _ = v76
                                                                                _ = v77
                                                                            }
                                                                        case "\n\t":
                                                                            var v78 uint8 = 207
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                    v63 = (^(v63))
                                                                }
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v79 := []int{(^(v63)), fi65, v63, -1}
                                                                    _ = v79
                                                                default:
                                                                    v73 += uint32(v72)
                                                                }
                                                                _ = v74
                                                            }
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        v80 := []int{(v63 | (v63 ^ 32768)), (fi65 ^ 68), 4294967295, 962}
                                                        if (4294967295 > 628) {
                                                            v62 += (v62 << ((535) & 63))
                                                            var v81 bool = false
                                                            _ = v81
                                                        } else {
                                                            v62 += 42
                                                            if (0 > 176) {
                                                                switch fi64 {
                                                                case 378:
                                                                    v82 := GenS1{F0: 146, F1: -128}
                                                                    _ = v82
                                                                case 539:
                                                                    if len(v80) >= 2 {
                                                                        v80 = v80[1:len(v80)]
                                                                    }
                                                                case -9223372036854775808:
                                                                    _ = len(v80)
                                                                }
                                                                v63 = fi65
                                                                v83 := GenS3{F0: 1, F1: ("hello" + ("\n\t" + "")), F2: 0}
                                                                _ = v83
                                                            }
                                                        }
                                                        _ = v80
                                                    } else {
                                                        var v84 int32 = 0
                                                        _ = v84
                                                    }
                                                }
                                            case "\n\t":
                                                var v85 uint = 1
                                                _ = v85
                                            case "":
                                                v86 := []int{(-(fi64)), (fi64 << ((2) & 63)), (v63 & fi64), v63}
                                                _ = v86
                                            default:
                                                switch 4294967295 {
                                                case -1:
                                                    switch "a" {
                                                    case "\n\t":
                                                        if false {
                                                            var v87 uint16 = 0
                                                            v88 := GenS8{F0: "a", F1: "foo bar", F2: 42}
                                                            if (2 >= 172) {
                                                                v89 := []int{fi65, fi65, (-(fi65)), 1}
                                                                v89 = append(v89, fi64)
                                                                v87 = (v87 / ((69) | 1))
                                                                _ = v89
                                                            }
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    case "a":
                                                        v62 -= (^(v62))
                                                    case "foo bar":
                                                        v62 -= v62
                                                    default:
                                                        var v90 uint = 0
                                                        _ = v90
                                                    }
                                                case 32768:
                                                    switch (fi64 * (v63 - v63)) {
                                                    case 128:
                                                        v91 := []int{}
                                                        _ = v91
                                                    case 784:
                                                        var v92 int64 = genFn0(0)
                                                        _ = v92
                                                    }
                                                default:
                                                    switch (^(v63)) {
                                                    case 4294967295:
                                                        v62 -= v62
                                                    default:
                                                        var v93 uint32 = 4294967295
                                                        _ = v93
                                                    }
                                                }
                                            }
                                        }
                                        if false {
                                            if ((1.5 / -2.25) > 0.0) {
                                                switch fi65 {
                                                case -9223372036854775808:
                                                    v62 += 1
                                                case 1:
                                                    v62 -= v62
                                                default:
                                                    if (4294967295 <= 342) {
                                                        var v94 bool = (-1 > 0)
                                                        if (v94 || v94) {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                fmt.Println(v62, fi64)
                                                            case "hello":
                                                                if v94 {
                                                                    if (-1 == -1) {
                                                                        v63 = v63
                                                                        v95 := GenS6{F0: 0}
                                                                        _ = v95
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            fmt.Println(v94, fi65, fi65)
                                                                        default:
                                                                            fmt.Println(v63, v63)
                                                                            v96 := []int{fi64}
                                                                            _ = v96
                                                                        }
                                                                        fmt.Println(v63)
                                                                        var v97 int32 = 1
                                                                        _ = v97
                                                                    }
                                                                } else {
                                                                    var v98 uint8 = 232
                                                                    v94 = (v94 && (-1 == 31))
                                                                    if v94 {
                                                                        if true {
                                                                            v62 = v62
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                var v99 uint32 = 1
                                                                                _ = v99
                                                                            case "\n\t":
                                                                                switch v63 {
                                                                                case 876:
                                                                                    fmt.Println(v63, v62)
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v63, v98, fi65)
                                                                                case 32768:
                                                                                    v100 := GenS10{F0: ((("" + "") + ("a" + "")) + "foo bar"), F1: v63}
                                                                                    _ = v100
                                                                                default:
                                                                                    v101 := GenS4{F0: ("hello" + "hello"), F1: 42, F2: ""}
                                                                                    if v94 {
                                                                                        v98 += v98
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            v62 += 4294967295
                                                                                        }
                                                                                        v94 = true
                                                                                    }
                                                                                    var v102 int32 = -643
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v103 uint64 = (v62 + v62)
                                                                                fmt.Println(v94)
                                                                                _ = v103
                                                                            } else {
                                                                                var v104 int64 = -498
                                                                                _ = v104
                                                                            }
                                                                        }
                                                                        if v94 {
                                                                            if true {
                                                                                var v105 float32 = (0.0 / (3.14159 - 1e10))
                                                                                v106 := []int{9223372036854775807, v63, (v63 ^ fi65)}
                                                                                if true {
                                                                                    if ((!v94) && false) {
                                                                                        if (-461 >= -1) {
                                                                                            var v107 uint64 = v62
                                                                                            v106 = append(v106, (fi64 >> ((1) & 63)))
                                                                                            _ = v107
                                                                                        }
                                                                                        switch (v63 & fi65) {
                                                                                        case -2147483648:
                                                                                            v98 = (^(v98))
                                                                                        case -317:
                                                                                            v106 = append(v106, -489)
                                                                                        }
                                                                                        v98 += (v98 % (((v98 % (((v98 * 39)) | 1))) | 1))
                                                                                    } else {
                                                                                        switch (-(v63)) {
                                                                                        case 2147483647:
                                                                                            if len(v106) >= 2 {
                                                                                                v106 = v106[1:len(v106)]
                                                                                            }
                                                                                        case 734:
                                                                                            v108 := GenS7{F0: 8}
                                                                                            _ = v108
                                                                                        default:
                                                                                            var v109 bool = (-767 == 0)
                                                                                            _ = v109
                                                                                        }
                                                                                        if v94 {
                                                                                            if (v94 && v94) {
                                                                                                if len(v106) > 0 {
                                                                                                    v106[0] = (-(fi64))
                                                                                                }
                                                                                            } else {
                                                                                                switch fi65 {
                                                                                                case -2147483648:
                                                                                                    if (100 >= 127) {
                                                                                                        v62 = (v62 & (v62 | 1))
                                                                                                        var v110 uint64 = v62
                                                                                                        if len(v106) >= 2 {
                                                                                                            v106 = v106[1:len(v106)]
                                                                                                        }
                                                                                                        _ = v110
                                                                                                    } else {
                                                                                                        var v111 int8 = 112
                                                                                                        _ = v111
                                                                                                    }
                                                                                                case 0:
                                                                                                    switch fi65 {
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(v105, v63, v62)
                                                                                                    case -814:
                                                                                                        var v112 uint32 = 42
                                                                                                        _ = v112
                                                                                                    case -294:
                                                                                                        var v113 uint = 0
                                                                                                        _ = v113
                                                                                                    default:
                                                                                                        _ = len(v106)
                                                                                                    }
                                                                                                default:
                                                                                                    var v114 bool = v94
                                                                                                    _ = v114
                                                                                                }
                                                                                                var v115 float32 = 0.0
                                                                                                var v116 string = (("foo bar" + "hello") + ("" + "\n\t"))
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            }
                                                                                            var v117 int = -9223372036854775808
                                                                                            _ = v117
                                                                                        } else {
                                                                                            v63 += -189
                                                                                            v118 := GenS11{F0: -62, F1: 127, F2: "a"}
                                                                                            _ = v118
                                                                                        }
                                                                                        v119 := GenS4{F0: "a", F1: 594, F2: "foo bar"}
                                                                                        _ = v119
                                                                                    }
                                                                                    if (v94 && ((v94 || true) && v94)) {
                                                                                        v105 -= v105
                                                                                        v106 = append(v106, (^(fi64)))
                                                                                    } else {
                                                                                        if v94 {
                                                                                            var v120 uint = 255
                                                                                            _ = v120
                                                                                        }
                                                                                        switch ("" + "foo bar") {
                                                                                        case "a":
                                                                                            switch fi65 {
                                                                                            case 128:
                                                                                                var v121 uint8 = v98
                                                                                                _ = v121
                                                                                            }
                                                                                        }
                                                                                        var v122 int8 = 127
                                                                                        _ = v122
                                                                                    }
                                                                                    if len(v106) >= 2 {
                                                                                        v106 = v106[1:len(v106)]
                                                                                    }
                                                                                }
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                            v123 := []int{(v63 ^ fi64), fi65, (fi64 << ((4294967295) & 63))}
                                                                            if (-1 > 100) {
                                                                                fmt.Println(fi64, v98)
                                                                            } else {
                                                                                var v124 uint32 = 2147483647
                                                                                if (876 >= 758) {
                                                                                    if ((0.1 == 1.5) || (!(fi64 < (fi65 - v63)))) {
                                                                                        v98 += (v98 * v98)
                                                                                    }
                                                                                    var v125 float64 = 0.1
                                                                                    var v126 bool = v94
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                } else {
                                                                                    fmt.Println(fi65, v94, fi65)
                                                                                    if len(v123) >= 2 {
                                                                                        v123 = v123[1:len(v123)]
                                                                                    }
                                                                                    if (false || (57 < 2)) {
                                                                                        v124 = (v124 << ((674) & 63))
                                                                                        var v127 uint8 = 203
                                                                                        _ = v127
                                                                                    } else {
                                                                                        v123 = append(v123, (fi64 & 128))
                                                                                    }
                                                                                }
                                                                                if len(v123) >= 2 {
                                                                                    v123 = v123[1:len(v123)]
                                                                                }
                                                                                _ = v124
                                                                            }
                                                                            _ = v123
                                                                        } else {
                                                                            fmt.Println(v98, v63)
                                                                            var v128 int32 = -94
                                                                            v129 := []int{fi65, 1}
                                                                            _ = v128
                                                                            _ = v129
                                                                        }
                                                                    } else {
                                                                        v98 = (^(v98))
                                                                    }
                                                                    _ = v98
                                                                }
                                                            case "a":
                                                                v130 := GenS9{F0: 42, F1: ("" + ""), F2: v62}
                                                                var v131 int8 = 0
                                                                v132 := []int{4294967295, (fi64 * (fi64 + (fi65 + v63)))}
                                                                _ = v130
                                                                _ = v131
                                                                _ = v132
                                                            }
                                                        } else {
                                                            v133 := []int{-9223372036854775808, (fi65 >> ((630) & 63)), (-(fi64))}
                                                            v63 = fi64
                                                            if (("a" + "\n\t") == "hello") {
                                                                v94 = false
                                                                _ = len(v133)
                                                                var v134 uint8 = 54
                                                                var v135 int32 = 2147483647
                                                                var v136 uint = 255
                                                                v133 = append(v133, fi64)
                                                                v134 = 241
                                                                v63 += int(v134)
                                                                var v137 int32 = v135
                                                                _ = v134
                                                                _ = v135
                                                                _ = v136
                                                                _ = v137
                                                            } else {
                                                                var v138 float32 = float32(3.14159)
                                                                switch (^(v63)) {
                                                                case 838:
                                                                    if (-117 != 42) {
                                                                        var v139 uint = 645
                                                                        _ = v139
                                                                    }
                                                                case -2147483648:
                                                                    switch (fi64 % (((v63 >> ((255) & 63))) | 1)) {
                                                                    case 32768:
                                                                        switch fi65 {
                                                                        case 128:
                                                                            v140 := GenS12{F0: (fi65 - 32767)}
                                                                            _ = v140
                                                                        case 544:
                                                                            var v141 bool = v94
                                                                            _ = v141
                                                                        }
                                                                    case 9223372036854775807:
                                                                        var v142 float64 = (1.5 * 0.1)
                                                                        _ = v142
                                                                    case -714:
                                                                        if v94 {
                                                                            v143 := GenS8{F0: ("a" + ("foo bar" + "")), F1: "foo bar", F2: 2147483647}
                                                                            v62 = (v62 * 40)
                                                                            _ = v143
                                                                        } else {
                                                                            fmt.Println(fi65, fi64, v62)
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                if v94 {
                                                                                    if false {
                                                                                        var v144 uint16 = 255
                                                                                        var v145 int = fi64
                                                                                        _ = v144
                                                                                        _ = v145
                                                                                    } else {
                                                                                        if true {
                                                                                            if ((-2147483648 <= 0) && true) {
                                                                                                var v146 int16 = -837
                                                                                                switch (v63 - (^(v63))) {
                                                                                                case 204:
                                                                                                    var v147 uint32 = 2147483647
                                                                                                    _ = v147
                                                                                                case 4294967295:
                                                                                                    v148 := GenS8{F0: "a", F1: "\n\t", F2: -2147483648}
                                                                                                    _ = v148
                                                                                                case 660:
                                                                                                    v146 += (v146 << ((0) & 63))
                                                                                                }
                                                                                                v149 := GenS13{F0: 4294967295, F1: genFn0(255)}
                                                                                                _ = v146
                                                                                                _ = v149
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v63, v138, v94)
                                                                                        }
                                                                                        if (784 <= 18446744073709551615) {
                                                                                            fmt.Println(fi65, fi64)
                                                                                            if len(v133) > 0 {
                                                                                                v133[0] = 42
                                                                                            }
                                                                                        } else {
                                                                                            v63 = (fi65 + v63)
                                                                                            if v94 {
                                                                                                v150 := GenS14{F0: -32768}
                                                                                                v151 := GenS15{F0: 0}
                                                                                                _ = v150
                                                                                                _ = v151
                                                                                            } else {
                                                                                                switch (("\n\t" + "a") + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v138, v63, v94)
                                                                                                default:
                                                                                                    v62 = 2
                                                                                                }
                                                                                            }
                                                                                            v138 = v138
                                                                                            var v152 uint = 880
                                                                                            _ = v152
                                                                                        }
                                                                                        v62 += (v62 << ((202) & 63))
                                                                                    }
                                                                                    var v153 int8 = -128
                                                                                    _ = v153
                                                                                } else {
                                                                                    v154 := GenS16{F0: 0, F1: -2147483648, F2: 18446744073709551615}
                                                                                    v94 = false
                                                                                    _ = v154
                                                                                }
                                                                            case "\n\t":
                                                                                if len(v133) > 0 {
                                                                                    v133[0] = int(v62)
                                                                                }
                                                                            default:
                                                                                v155 := GenS9{F0: 207, F1: ("" + "a"), F2: v62}
                                                                                _ = v155
                                                                            }
                                                                        }
                                                                    }
                                                                case 1:
                                                                    switch fi65 {
                                                                    case -9223372036854775808:
                                                                        var v156 uint16 = 42
                                                                        v157 := GenS17{F0: ("foo bar" + "foo bar")}
                                                                        if len(v133) > 0 {
                                                                            v133[0] = fi65
                                                                        }
                                                                        var v158 float32 = ((v138 - v138) / -2.25)
                                                                        _ = v156
                                                                        _ = v157
                                                                        _ = v158
                                                                    case 815:
                                                                        v159 := GenS6{F0: 0}
                                                                        _ = v159
                                                                    case -2147483648:
                                                                        v94 = false
                                                                    }
                                                                default:
                                                                    var v160 uint64 = (v62 << ((65535) & 63))
                                                                    _ = v160
                                                                }
                                                                var v161 int32 = -857
                                                                v133 = append(v133, fi65)
                                                                if ((v94 || (735 < 37)) && (!v94)) {
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v138 = 0.1
                                                                    }
                                                                } else {
                                                                    v162 := GenS18{F0: uint32(v62), F1: 1}
                                                                    fmt.Println(fi65, fi65, v138)
                                                                    _ = v162
                                                                }
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    switch 487 {
                                                                    case -2147483648:
                                                                        v138 = (0.1 / (v138 * 0.1))
                                                                    case 0:
                                                                        var v163 int32 = (-(v161))
                                                                        _ = v163
                                                                    case 127:
                                                                        var v164 uint8 = 127
                                                                        _ = v164
                                                                    default:
                                                                        v63 = (fi64 | v63)
                                                                        var v165 float64 = 3.14159
                                                                        _ = v165
                                                                    }
                                                                case "\n\t":
                                                                    if (!v94) {
                                                                        v166 := GenS19{F0: 507}
                                                                        v62 += v62
                                                                        _ = v166
                                                                    } else {
                                                                        if len(v133) >= 2 {
                                                                            v133 = v133[1:len(v133)]
                                                                        }
                                                                        if (fi65 >= (v63 * fi64)) {
                                                                            if ((^(v62)) != 437) {
                                                                                var v167 int8 = 1
                                                                                switch (fi65 / ((437) | 1)) {
                                                                                case -9223372036854775808:
                                                                                    v168 := GenS12{F0: fi65}
                                                                                    _ = v168
                                                                                }
                                                                                var v169 int64 = -1
                                                                                _ = v167
                                                                                _ = v169
                                                                            } else {
                                                                                v94 = (!false)
                                                                            }
                                                                            if v94 {
                                                                                fmt.Println(v94)
                                                                                var v170 uint64 = (v62 * v62)
                                                                                _ = v170
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v94, fi65, v62)
                                                                                    v171 := GenS4{F0: "\n\t", F1: 1, F2: ""}
                                                                                    var v172 uint8 = 42
                                                                                    _ = v171
                                                                                    _ = v172
                                                                                }
                                                                                var v173 int16 = int16(v161)
                                                                                var v174 int = (-(fi65))
                                                                                _ = v173
                                                                                _ = v174
                                                                            }
                                                                            v94 = v94
                                                                        } else {
                                                                            _ = len(v133)
                                                                        }
                                                                        v133 = append(v133, (fi65 / ((822) | 1)))
                                                                    }
                                                                default:
                                                                    v138 += (v138 / ((0.0 * 3.14159) + (0.1 - 3.14159)))
                                                                }
                                                                var v175 int8 = -105
                                                                _ = v138
                                                                _ = v161
                                                                _ = v175
                                                            }
                                                            _ = v133
                                                        }
                                                        _ = v94
                                                    } else {
                                                        v176 := GenS10{F0: "a", F1: (fi65 * (fi65 + fi65))}
                                                        if (2 >= uint16(v62)) {
                                                            v177 := GenS16{F0: 2147483647, F1: -1, F2: 707}
                                                            var v178 uint64 = 4294967295
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v179 := GenS3{F0: 58, F1: ("" + "hello"), F2: 128}
                                                                v180 := []int{fi64, (fi64 * (fi64 >> ((732) & 63))), fi64, (fi64 & 2147483647)}
                                                                _ = v179
                                                                _ = v180
                                                            }
                                                            _ = v177
                                                            _ = v178
                                                        }
                                                        v63 -= (fi64 / (((^(fi64))) | 1))
                                                        _ = v176
                                                    }
                                                }
                                                if true {
                                                    if false {
                                                        v181 := []int{}
                                                        var v182 string = ""
                                                        var v183 int16 = 219
                                                        var v184 uint32 = 42
                                                        _ = v181
                                                        _ = v182
                                                        _ = v183
                                                        _ = v184
                                                    } else {
                                                        if (1 == 1) {
                                                            v63 = fi64
                                                        } else {
                                                            if (1 == -32768) {
                                                                var v185 uint16 = 829
                                                                v186 := []int{(^(fi65)), fi64, fi64}
                                                                var v187 uint16 = v185
                                                                _ = v185
                                                                _ = v186
                                                                _ = v187
                                                            }
                                                            v62 = v62
                                                            fmt.Println(fi64, v63, fi65)
                                                        }
                                                        fmt.Println(v63, v63, fi65)
                                                        v62 = (^(v62))
                                                    }
                                                    switch v63 {
                                                    case 4294967295:
                                                        var v188 uint32 = 4294967295
                                                        _ = v188
                                                    }
                                                } else {
                                                    if true {
                                                        v62 += (^(v62))
                                                        fmt.Println(v62)
                                                        v189 := []int{}
                                                        _ = v189
                                                    }
                                                }
                                            } else {
                                                v190 := GenS20{F0: (fi65 * fi65), F1: genFn0(18446744073709551615)}
                                                _ = v190
                                            }
                                            switch v63 {
                                            case 127:
                                                v63 -= fi65
                                            }
                                            var v191 uint = 295
                                            _ = v191
                                        } else {
                                            if true {
                                                var v192 int64 = 9223372036854775807
                                                v192 = genFn0(98)
                                                switch v63 {
                                                case 128:
                                                    var v193 int64 = 957
                                                    _ = v193
                                                case -9223372036854775808:
                                                    var v194 int8 = 42
                                                    _ = v194
                                                default:
                                                    v192 += (v192 - int64(fi64))
                                                }
                                                _ = v192
                                            } else {
                                                v195 := []int{v63, v63}
                                                var v196 uint = 0
                                                v197 := GenS6{F0: 0}
                                                v198 := GenS13{F0: v196, F1: 9223372036854775807}
                                                _ = v195
                                                _ = v196
                                                _ = v197
                                                _ = v198
                                            }
                                            if (false && ((!true) && false)) {
                                                v199 := GenS21{F0: -9223372036854775808, F1: (v62 >> ((0) & 63))}
                                                if ("\n\t" <= ("a" + "foo bar")) {
                                                    var v200 bool = (-503 >= 1)
                                                    _ = v200
                                                } else {
                                                    v201 := []int{}
                                                    _ = v201
                                                }
                                                v202 := []int{(v63 % (((fi65 % ((v63) | 1))) | 1)), (fi64 & (v63 | (fi64 & (fi64 + -2147483648)))), v63}
                                                _ = v199
                                                _ = v202
                                            } else {
                                                if true {
                                                    v203 := []int{fi64, (v63 - 32767), (fi64 | (v63 ^ 32768)), 42}
                                                    fmt.Println(v62, fi65)
                                                    _ = v203
                                                }
                                            }
                                            if (0.1 == (float32(1.5) * 1e10)) {
                                                v204 := []int{fi65, fi64, v63}
                                                _ = v204
                                            } else {
                                                v205 := GenS1{F0: v62, F1: -1}
                                                _ = v205
                                            }
                                        }
                                    }
                                    v206 := GenS22{F0: 65535, F1: "a"}
                                    v207 := []int{(fi65 + 851), (fi65 / (((v63 / ((-9223372036854775808) | 1))) | 1)), v63, (fi65 - v63)}
                                    if ((("hello" + "\n\t") < ("\n\t" + "\n\t")) || (0.0 != 1e10)) {
                                        var v208 int = 9223372036854775807
                                        v62 -= (v62 + (^(v62)))
                                        v208 += 42
                                        _ = v208
                                    }
                                    v63 += (-(v63))
                                    var v209 uint = 2
                                    _ = v206
                                    _ = v207
                                    _ = v209
                                } else {
                                    switch (fi65 * fi65) {
                                    case 1:
                                        var v210 uint = 65535
                                        _ = v210
                                    }
                                }
                            }
                            _ = fi65
                        }
                        _ = fi64
                    }
                case 1:
                    for fi211 := 0; fi211 < 9; fi211++ {
                        for fi212 := 0; fi212 < 2; fi212++ {
                            if true {
                                var v213 string = "\n\t"
                                if (872 < -2147483648) {
                                    v214 := GenS11{F0: 42, F1: 0, F2: "a"}
                                    if ((3.14159 + 3.14159) >= 0.1) {
                                        if false {
                                            fmt.Println(v213, fi212)
                                        } else {
                                            v62 += v62
                                        }
                                        var v215 int64 = 9223372036854775807
                                        v216 := GenS20{F0: (v63 << ((530) & 63)), F1: (v215 % (((v215 ^ (v215 % ((v215) | 1)))) | 1))}
                                        _ = v215
                                        _ = v216
                                    }
                                    _ = v214
                                } else {
                                    v63 -= (fi212 << ((65535) & 63))
                                    var v217 uint32 = 1
                                    _ = v217
                                }
                                _ = v213
                            } else {
                                v218 := GenS9{F0: 516, F1: "a", F2: (v62 | v62)}
                                if false {
                                    v62 = 4294967295
                                    v63 = (fi212 >> ((1) & 63))
                                    v63 -= (fi211 + 4294967295)
                                } else {
                                    var v219 uint16 = 56
                                    _ = v219
                                }
                                var v220 uint32 = 42
                                _ = v218
                                _ = v220
                            }
                            var v221 int64 = genFn0(2)
                            v222 := GenS23{F0: (fi211 << ((uint(v62)) & 63))}
                            v62 = (v62 ^ v62)
                            _ = fi212
                            _ = v221
                            _ = v222
                        }
                        fmt.Println(v62)
                        for fi223 := 0; fi223 < 7; fi223++ {
                            var v224 bool = false
                            v225 := []int{(fi223 % ((v63) | 1)), (-(v63)), fi223}
                            v224 = (!false)
                            _ = fi223
                            _ = v224
                            _ = v225
                        }
                        _ = fi211
                    }
                case 127:
                    v63 += v63
                }
                _ = v62
                _ = v63
            } else {
                v226 := []int{1, 179}
                var v227 int32 = -315
                _ = v226
                _ = v227
            }
        }
        for fi228 := 0; fi228 < 16; fi228++ {
            var v229 int8 = -128
            for fi230 := 0; fi230 < 7; fi230++ {
                fmt.Println(fi228, fi230, fi230)
                _ = fi230
            }
            _ = fi228
            _ = v229
        }
    }
    for fi231 := 0; fi231 < 13; fi231++ {
        if false {
            var v232 uint32 = 2147483648
            _ = v232
        }
        fmt.Println(fi231)
        _ = fi231
    }
    if false {
        if ((true || false) && true) {
            switch -2147483648 {
            case 1:
                var v233 uint16 = 0
                _ = v233
            case 0:
                var v234 float32 = 3.14159
                _ = v234
            case -1:
                v235 := GenS19{F0: 4294967295}
                _ = v235
            }
            var v236 uint = 255
            _ = v236
        }
        if true {
            var v237 string = "foo bar"
            var v238 bool = true
            switch -1 {
            case 32767:
                v239 := GenS2{F0: -1}
                _ = v239
            case 9223372036854775807:
                v240 := []int{-1, -482, 537, 315}
                _ = v240
            case 42:
                fmt.Println(v238, v237)
            default:
                switch -9223372036854775808 {
                case 32767:
                    v238 = (!v238)
                case -847:
                    var v241 int8 = 100
                    _ = v241
                default:
                    v242 := []int{-9223372036854775808, 2147483647, 4294967295, 32768}
                    _ = v242
                }
            }
            _ = v237
            _ = v238
        }
        var v243 string = (("hello" + "a") + ("hello" + ("\n\t" + "")))
        _ = v243
    } else {
        var v244 uint32 = 209
        if false {
            v245 := GenS21{F0: 42, F1: 4294967295}
            if (42 != -8) {
                v246 := GenS24{F0: 42}
                v244 = 0
                if (false || false) {
                    v244 += v244
                    v244 = 2
                    var v247 int8 = 100
                    _ = v247
                } else {
                    for fi248 := 0; fi248 < 12; fi248++ {
                        if true {
                            fmt.Println(v244, v244)
                            fmt.Println(v244)
                        } else {
                            v244 = v244
                        }
                        _ = fi248
                    }
                    v249 := GenS25{F0: false, F1: 2}
                    v250 := GenS9{F0: v244, F1: "hello", F2: 550}
                    _ = v249
                    _ = v250
                }
                _ = v246
            } else {
                v251 := GenS3{F0: v244, F1: "hello", F2: 2}
                v252 := GenS24{F0: 42}
                _ = v251
                _ = v252
            }
            _ = v245
        }
        _ = v244
    }
    if (395 == 4294967295) {
        var v253 int64 = 327
        v254 := []int{-9223372036854775808, 9223372036854775807, 32768, 4294967295}
        _ = v253
        _ = v254
    } else {
        var v255 int = 726
        v256 := GenS26{F0: (!true)}
        v257 := []int{v255}
        _ = v255
        _ = v256
        _ = v257
    }
    return 2147483648
}

func main() {
    var v1 bool = (2147483648 > 0)
    var v2 float64 = 0.1
    v3 := make(map[int]bool)
    var v258 int16 = -972
    var v259 string = "a"
    v259 = "a"
    var v260 uint = 2
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch 2147483647 {
        case -1:
            for fi261 := 0; fi261 < 17; fi261++ {
                var v262 bool = true
                for fi263 := 0; fi263 < 9; fi263++ {
                    if v262 {
                        if v1 {
                            var v264 uint = 439
                            v258 += (v258 & (v258 % ((-32768) | 1)))
                            fmt.Println(v260, v259, fi263)
                            _ = v264
                        } else {
                            v265 := GenS25{F0: (1 < 2), F1: 1}
                            var v266 int = 4294967295
                            v260 = (v260 - v260)
                            var v267 int = v266
                            var v268 uint = (v260 + v260)
                            var v269 uint16 = 529
                            _ = v265
                            _ = v266
                            _ = v267
                            _ = v268
                            _ = v269
                        }
                    } else {
                        if ((fi263 >> (((v260 | (v260 - (v260 >> ((v260) & 63))))) & 63)) < fi263) {
                            if (genFn0(v260) >= 1) {
                                v2 = (v2 / ((((0.1 - ((((float64((0.1 - ((0.0 + 1e10) * (1e10 * 3.14159)))) - v2) / 3.14159) + v2) + 1e10)) + v2) - v2) + -2.25))
                            }
                            v1 = false
                            var v270 string = ((v259 + "foo bar") + v259)
                            fmt.Println(fi263, v258)
                            _ = v270
                        } else {
                            v271 := []int{(fi263 << ((v260) & 63)), fi263, fi263, (-(fi261))}
                            _ = v271
                        }
                    }
                    var v272 string = v259
                    _ = fi263
                    _ = v272
                }
                if false {
                    if v262 {
                        switch v259 {
                        case "hello":
                            var v273 int64 = -9223372036854775808
                            _ = v273
                        }
                        v274 := []int{(fi261 % (((fi261 * (fi261 * (fi261 / ((fi261) | 1))))) | 1)), (fi261 + 32767), fi261, (fi261 % ((fi261) | 1))}
                        _ = v274
                    } else {
                        var v275 string = v259
                        _ = v275
                    }
                    v1 = ((int8(v258) != 100) || v262)
                    var v276 float32 = (0.1 - (float32((v2 / (float64(1.5) - ((v2 / (1.5 - 3.14159)) / -2.25)))) - 0.0))
                    _ = v276
                } else {
                    for fi277 := 0; fi277 < 6; fi277++ {
                        if (v1 || (v259 != ("a" + "\n\t"))) {
                            fmt.Println(v2, fi277, v2)
                            var v278 uint = (v260 * v260)
                            _ = v278
                        } else {
                            v279 := GenS11{F0: -128, F1: -81, F2: ((v259 + "a") + v259)}
                            var v280 int = (fi277 | (fi261 >> ((v260) & 63)))
                            _ = v279
                            _ = v280
                        }
                        _ = fi277
                    }
                    v281 := v3[(fi261 | fi261)]
                    _ = v281
                }
                _ = fi261
                _ = v262
            }
        case 32767:
            if true {
                for fi282 := 0; fi282 < 18; fi282++ {
                    if ((!v1) && (((true || v1) && (975 >= 42)) && (0.1 <= 3.14159))) {
                        var v283 float64 = v2
                        var v284 int64 = 9223372036854775807
                        _ = v283
                        _ = v284
                    } else {
                        v285 := GenS23{F0: fi282}
                        for fi286 := 0; fi286 < 9; fi286++ {
                            fmt.Println(v2, fi282, v1)
                            v258 += (v258 | (-(v258)))
                            if (149 != -1) {
                                fmt.Println(v259)
                            } else {
                                switch (v259 + (((("foo bar" + "a") + "foo bar") + ("a" + "foo bar")) + (v259 + "\n\t"))) {
                                case "hello":
                                    var v287 uint16 = 255
                                    _ = v287
                                case "foo bar":
                                    switch 1 {
                                    case -989:
                                        var v288 uint8 = 127
                                        _ = v288
                                    case 42:
                                        v289 := GenS2{F0: 0}
                                        _ = v289
                                    case -424:
                                        fmt.Println(v258)
                                    default:
                                        v1 = v1
                                    }
                                default:
                                    delete(v3, fi282)
                                }
                                v290 := GenS27{F0: (433 != 568), F1: (!false), F2: -591}
                                v291 := GenS28{F0: v1}
                                fmt.Println(v259, v260, fi282)
                                v1 = v1
                                if v1 {
                                    v292 := GenS29{F0: v1, F1: 42}
                                    if v1 {
                                        var v293 uint = v260
                                        var v294 uint = 4294967295
                                        _ = v293
                                        _ = v294
                                    }
                                    var v295 float32 = 0.1
                                    _ = v292
                                    _ = v295
                                }
                                _ = v290
                                _ = v291
                            }
                            _ = fi286
                        }
                        _ = v285
                    }
                    _ = fi282
                }
            } else {
                var v296 bool = (v1 || v1)
                for fi297 := 0; fi297 < 3; fi297++ {
                    for fi298 := 0; fi298 < 3; fi298++ {
                        v299 := GenS30{F0: 2147483647, F1: fi298, F2: 0}
                        _ = fi298
                        _ = v299
                    }
                    if (!v296) {
                        v258 = (v258 << (((^(v260))) & 63))
                        switch v259 {
                        case "hello":
                            v300 := GenS31{F0: "a"}
                            _ = v300
                        }
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v301 := v3[fi297]
                            _ = v301
                        default:
                            switch v259 {
                            case "hello":
                                v258 -= v258
                            case "a":
                                if v296 {
                                    v2 += (v2 + v2)
                                    v258 = 42
                                    var v302 uint8 = 1
                                    _ = v302
                                } else {
                                    if (!false) {
                                        switch fi297 {
                                        case -800:
                                            v303 := GenS8{F0: ("" + ("hello" + v259)), F1: "hello", F2: 2147483647}
                                            _ = v303
                                        case -1:
                                            var v304 int = 115
                                            _ = v304
                                        case 42:
                                            v305 := []int{}
                                            _ = v305
                                        }
                                        for fi306 := 0; fi306 < 13; fi306++ {
                                            v259 = v259
                                            switch v259 {
                                            case "hello":
                                                v307 := v3[fi306]
                                                _ = v307
                                            case "a":
                                                fmt.Println(v296, v260, v1)
                                                if (!false) {
                                                    fmt.Println(v258)
                                                    switch v259 {
                                                    case "hello":
                                                        var v308 string = "\n\t"
                                                        _ = v308
                                                    case "foo bar":
                                                        v309 := []int{fi306, -9223372036854775808, (fi297 ^ fi297)}
                                                        if v1 {
                                                            if v1 {
                                                                if (42 >= 93) {
                                                                    fmt.Println(fi297, v2, fi297)
                                                                } else {
                                                                    v310 := GenS32{F0: 65535}
                                                                    switch v259 {
                                                                    case "a":
                                                                        var v311 string = ("hello" + "\n\t")
                                                                        _ = v311
                                                                    default:
                                                                        v312 := GenS31{F0: v259}
                                                                        _ = v312
                                                                    }
                                                                    _ = v310
                                                                }
                                                            }
                                                            fmt.Println(fi297, v296, fi306)
                                                        } else {
                                                            v2 = (v2 / float64(float32(((1e10 + ((((1.5 - 0.1) + 1e10) + (float64((float32(float64(0.1)) * 1e10)) + (3.14159 + float64((-2.25 - 3.14159))))) - v2)) / v2))))
                                                            fmt.Println(v296, v2, v260)
                                                            switch (v259 + (v259 + "foo bar")) {
                                                            case "\n\t":
                                                                switch v259 {
                                                                case "\n\t":
                                                                    var v313 int64 = -9223372036854775808
                                                                    _ = v313
                                                                }
                                                            case "a":
                                                                var v314 uint32 = 48
                                                                _ = v314
                                                            case "foo bar":
                                                                var v315 uint32 = 42
                                                                _ = v315
                                                            }
                                                        }
                                                        if (!true) {
                                                            v260 -= v260
                                                            v316 := GenS26{F0: (v1 && v1)}
                                                            _ = v316
                                                        } else {
                                                            var v317 uint32 = 452
                                                            v258 -= v258
                                                            var v318 uint8 = 0
                                                            v319 := GenS33{F0: genFn0(4294967295)}
                                                            v320 := GenS34{F0: 0, F1: (v317 << (((^(v260))) & 63))}
                                                            _ = v317
                                                            _ = v318
                                                            _ = v319
                                                            _ = v320
                                                        }
                                                        var v321 uint64 = 0
                                                        _ = v309
                                                        _ = v321
                                                    case "":
                                                        if (126 >= 255) {
                                                            if ((1.5 / float32((float64((3.14159 - 0.1)) / v2))) != ((0.1 * (((0.0 / (1.5 + (0.0 * 0.1))) * float32(v2)) - ((1.5 / 1.0) * (((0.0 - 3.14159) * 3.14159) / 1.0)))) * 1e10)) {
                                                                fmt.Println(v1, v258, v1)
                                                                v322 := v3[-788]
                                                                _ = v322
                                                            }
                                                        }
                                                    default:
                                                        var v323 uint8 = 2
                                                        _ = v323
                                                    }
                                                    var v324 int8 = int8(v258)
                                                    _ = v324
                                                } else {
                                                    v325 := []int{4294967295, (-(fi306)), (fi297 | fi297), (fi306 << ((903) & 63))}
                                                    v326 := GenS4{F0: v259, F1: 332, F2: ""}
                                                    _ = v325
                                                    _ = v326
                                                }
                                            case "foo bar":
                                                if true {
                                                    switch ("" + v259) {
                                                    case "foo bar":
                                                        var v327 int32 = -467
                                                        _ = v327
                                                    case "hello":
                                                        delete(v3, fi306)
                                                    }
                                                    var v328 uint64 = 261
                                                    v328 += 96
                                                    v2 -= (v2 - v2)
                                                    fmt.Println(fi306)
                                                    _ = v328
                                                } else {
                                                    v329 := GenS29{F0: (-1 == 1), F1: int32(v258)}
                                                    v260 -= v260
                                                    v296 = (!(1 == -795))
                                                    _ = v329
                                                }
                                            default:
                                                if (true && v1) {
                                                    fmt.Println(v259, v2)
                                                    v259 = v259
                                                } else {
                                                    fmt.Println(v296, v296)
                                                    v3[(fi306 % ((fi297) | 1))] = (1e10 > float32(v2))
                                                }
                                            }
                                            var v330 uint8 = 128
                                            _ = fi306
                                            _ = v330
                                        }
                                    }
                                    fmt.Println(v260, v260, v259)
                                }
                            }
                        }
                        for fi331 := 0; fi331 < 17; fi331++ {
                            fmt.Println(v296, fi297)
                            _ = fi331
                        }
                        v3[(fi297 | fi297)] = (255 != 65535)
                    }
                    v332 := GenS23{F0: -26}
                    _ = fi297
                    _ = v332
                }
                _ = v296
            }
        case 128:
            if (65535 >= 2147483648) {
                v333 := GenS25{F0: v1, F1: 1}
                _ = v333
            }
        default:
            var v334 uint8 = 0
            _ = v334
        }
    default:
        v2 -= ((0.0 / 0.1) / ((v2 * ((float64(0.1) + 1e10) / v2)) - v2))
    }
    if (!(true && (42 == -457))) {
        v258 = (v258 << ((v260) & 63))
        v2 = v2
    } else {
        v259 = ((v259 + ((v259 + "") + v259)) + v259)
        var v335 int16 = (-(v258))
        _ = v335
    }
    for fi336 := 0; fi336 < 18; fi336++ {
        var v337 int8 = -128
        var v338 int16 = v258
        _ = fi336
        _ = v337
        _ = v338
    }
    v2 -= (((float64(-2.25) - 3.14159) - float64(-2.25)) - -2.25)
    fmt.Println(v259, v258)
    if (v1 && (v1 || false)) {
        var v339 uint64 = 0
        if ((1000 <= (v258 << ((539) & 63))) && ((v339 & (v339 - 0)) != v339)) {
            var v340 float64 = v2
            if false {
                v339 = (v339 ^ (v339 | (^(v339))))
                var v341 bool = (42 <= 9223372036854775807)
                _ = v341
            }
            _ = v340
        }
        _ = v339
    }
    var v342 int32 = int32(v260)
    v343 := []int{4294967295, -685, 2147483647}
    if (!(841 >= 2147483647)) {
        var v344 float64 = float64(1.5)
        for fi345 := 0; fi345 < 12; fi345++ {
            if v1 {
                _ = len(v343)
            } else {
                var v346 uint64 = 42
                v347 := GenS35{F0: -2147483648}
                v343 = append(v343, (fi345 >> ((42) & 63)))
                _ = v346
                _ = v347
            }
            if len(v343) > 0 {
                v343[0] = (fi345 + (fi345 % (((fi345 ^ 955)) | 1)))
            }
            var v348 uint = (^(v260))
            _ = fi345
            _ = v348
        }
        var v349 bool = (1 == 712)
        _ = v344
        _ = v349
    } else {
        v350 := GenS36{F0: -29, F1: 572}
        if v1 {
            v1 = ((((2 < 2) && v1) || (!v1)) && true)
            for fi351 := 0; fi351 < 19; fi351++ {
                switch -9223372036854775808 {
                case 942:
                    if (!((v258 << ((940) & 63)) >= (v258 % ((int16(fi351)) | 1)))) {
                        for fi352 := 0; fi352 < 9; fi352++ {
                            delete(v3, fi352)
                            _ = fi352
                        }
                        v353 := GenS22{F0: 42, F1: (("\n\t" + v259) + v259)}
                        v1 = (v259 < v259)
                        _ = v353
                    }
                case 42:
                    switch (v259 + "\n\t") {
                    case "hello":
                        v354 := GenS23{F0: (fi351 ^ 0)}
                        _ = v354
                    case "\n\t":
                        if len(v343) > 0 {
                            v343[0] = (fi351 >> ((329) & 63))
                        }
                    case "foo bar":
                        if false {
                            for fi355 := 0; fi355 < 11; fi355++ {
                                var v356 uint8 = 191
                                _ = fi355
                                _ = v356
                            }
                            var v357 int8 = 0
                            for fi358 := 0; fi358 < 7; fi358++ {
                                if (!v1) {
                                    if (-1 >= genFn0(v260)) {
                                        delete(v3, (fi351 | -1))
                                        v359 := GenS13{F0: (v260 | v260), F1: int64(v357)}
                                        _ = v359
                                    } else {
                                        var v360 int16 = (-(v258))
                                        v259 = v259
                                        _ = v360
                                    }
                                    if (v1 || v1) {
                                        var v361 bool = ((((v258 / (((v258 >> ((4294967295) & 63))) | 1)) > v258) && (255 != 42)) && true)
                                        delete(v3, fi358)
                                        switch 32768 {
                                        case 0:
                                            fmt.Println(v357, fi351)
                                        case 128:
                                            v2 = v2
                                        case 32768:
                                            v362 := GenS37{F0: v357, F1: -837}
                                            _ = v362
                                        }
                                        v258 = -953
                                        _ = v361
                                    } else {
                                        var v363 float32 = -2.25
                                        v258 -= (v258 | v258)
                                        v260 = 65535
                                        _ = v363
                                    }
                                    v2 += v2
                                    v260 -= (v260 + 383)
                                }
                                if len(v343) >= 2 {
                                    v343 = v343[1:len(v343)]
                                }
                                var v364 float32 = 1e10
                                _ = fi358
                                _ = v364
                            }
                            _ = v357
                        }
                    default:
                        v2 -= (v2 * v2)
                    }
                }
                _ = fi351
            }
        } else {
            var v365 uint8 = 0
            _ = v365
        }
        _ = v350
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v258:", v258)
    fmt.Println("v259:", v259)
    fmt.Println("v260:", v260)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v342:", v342)
    fmt.Println("slicelen v343:", len(v343), v343)
    _ = v1
    _ = v2
    _ = v258
    _ = v259
    _ = v260
    _ = v3
    _ = v342
    _ = v343
}
