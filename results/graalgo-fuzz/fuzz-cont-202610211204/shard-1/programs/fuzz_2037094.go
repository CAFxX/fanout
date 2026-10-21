package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
type GenS3 struct { F0 int; F1 uint; F2 uint }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint64; F1 string; F2 int8 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int32; F1 uint; F2 uint }
type GenS11 struct { F0 int8; F1 uint32 }
type GenS12 struct { F0 string; F1 uint32; F2 uint8 }

func main() {
    var v1 string = ("" + ("hello" + ("foo bar" + "")))
    var v2 uint = 18446744073709551615
    var v3 int16 = 0
    var v4 int8 = 1
    v5 := make(map[int]bool)
    v6 := []uint32{65535, 242, 48}
    v1 = "\n\t"
    v7 := GenS1{F0: 1}
    if (!false) {
        var v8 int16 = v3
        if true {
            var v9 bool = false
            if v9 {
                for fi10 := 0; fi10 < 2; fi10++ {
                    _ = len(v5)
                    _ = fi10
                }
                v11 := v5[262]
                if (false || v11) {
                    if (v3 < -1) {
                        var v12 uint8 = 1
                        var v13 float32 = 1.5
                        _ = v12
                        _ = v13
                    }
                    v11 = true
                    v14 := GenS1{F0: 1}
                    _ = v14
                }
                _ = v11
            }
            v4 = (-(v4))
            _ = v9
        } else {
            var v15 int8 = v4
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v15
        }
        _ = v8
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            switch fi16 {
            case 32768:
                if (v2 >= (v2 % ((0) | 1))) {
                    v3 = v3
                    var v17 bool = false
                    var v18 int16 = v3
                    _ = v17
                    _ = v18
                } else {
                    v3 += v3
                    v19 := GenS2{F0: true, F1: 271, F2: -1}
                    v20 := GenS1{F0: 18446744073709551615}
                    _ = v19
                    _ = v20
                }
            case 42:
                if false {
                    v21 := v5[fi16]
                    if true {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            if (!false) {
                                if true {
                                    fmt.Println(fi16, v3)
                                } else {
                                    var v23 int16 = 32767
                                    if false {
                                        var v24 uint32 = 2
                                        if (v24 <= v24) {
                                            switch "\n\t" {
                                            case "a":
                                                var v25 uint8 = 255
                                                _ = v25
                                            }
                                            fmt.Println(v1, v4)
                                        } else {
                                            fmt.Println(v3)
                                            if (v2 < v2) {
                                                var v26 uint8 = 44
                                                fmt.Println(v26, v3)
                                                _ = v26
                                            }
                                        }
                                        switch (fi22 * fi22) {
                                        case 42:
                                            var v27 float64 = 3.14159
                                            _ = v27
                                        case 4294967295:
                                            var v28 uint = (v2 % ((v2) | 1))
                                            _ = v28
                                        default:
                                            v3 = v3
                                        }
                                        _ = v24
                                    } else {
                                        var v29 float32 = ((float32(((0.0 - 1.5) / (3.14159 * 3.14159))) - 1.5) * 0.1)
                                        _ = v29
                                    }
                                    var v30 uint8 = 128
                                    _ = v23
                                    _ = v30
                                }
                                switch ("foo bar" + v1) {
                                case "":
                                    if (-269 == 712) {
                                        if v21 {
                                            if (float32(1.5) < (1e10 + 0.0)) {
                                                if v21 {
                                                    if (!v21) {
                                                        var v31 uint8 = 127
                                                        fmt.Println(v2, v1)
                                                        _ = v31
                                                    }
                                                    if (!(88 > v2)) {
                                                        _ = len(v6)
                                                        fmt.Println(v2)
                                                        if (!(v4 != v4)) {
                                                            fmt.Println(v21, v21)
                                                            if (fi16 <= (fi22 ^ 238)) {
                                                                if true {
                                                                    v32 := GenS2{F0: false, F1: 42, F2: (-(v4))}
                                                                    v33 := GenS1{F0: 18446744073709551615}
                                                                    var v34 uint8 = 139
                                                                    _ = v32
                                                                    _ = v33
                                                                    _ = v34
                                                                } else {
                                                                    var v35 int8 = (v4 & (v4 - (v4 >> ((v2) & 63))))
                                                                    _ = v35
                                                                }
                                                                var v36 float64 = -2.25
                                                                _ = v36
                                                            }
                                                            var v37 int16 = (v3 / (((v3 + 42)) | 1))
                                                            var v38 uint16 = 255
                                                            v2 = v2
                                                            var v39 bool = (v37 <= int16(v4))
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                    } else {
                                                        v40 := GenS2{F0: (94 >= (^(v4))), F1: 381, F2: v4}
                                                        if ((42 > 18446744073709551615) || (4 < 2147483647)) {
                                                            var v41 uint64 = 42
                                                            _ = v41
                                                        }
                                                        if (float32(-2.25) >= (float32(0.1) + (1.5 + 3.14159))) {
                                                            var v42 int16 = 0
                                                            var v43 uint64 = 42
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        _ = v40
                                                    }
                                                } else {
                                                    v44 := GenS1{F0: 0}
                                                    fmt.Println(fi22, v4, v4)
                                                    var v45 int32 = 42
                                                    v1 = ((("\n\t" + v1) + (("foo bar" + "foo bar") + "hello")) + (((v1 + v1) + (v1 + v1)) + v1))
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                        }
                                    } else {
                                        switch (-(fi22)) {
                                        case -815:
                                            v46 := GenS3{F0: -1, F1: v2, F2: (v2 | (v2 | v2))}
                                            _ = v46
                                        default:
                                            v4 += 42
                                        }
                                        var v47 int = fi22
                                        v1 = (v1 + v1)
                                        _ = v47
                                    }
                                }
                                v21 = (226 <= -513)
                                if v21 {
                                    if len(v6) > 0 {
                                        v6[0] = 656
                                    }
                                    v48 := GenS2{F0: (131 >= 28), F1: 18446744073709551615, F2: v4}
                                    _ = v48
                                }
                            } else {
                                var v49 float64 = (1.5 * -2.25)
                                var v50 float32 = 0.0
                                _ = v49
                                _ = v50
                            }
                            fmt.Println(v21, fi22, fi16)
                            _ = fi22
                        }
                        v21 = v21
                    }
                    v51 := GenS2{F0: false, F1: 900, F2: v4}
                    _ = v21
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 1; fi52++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v53 int8 = int8(v3)
                        v54 := GenS4{F0: 382, F1: v4}
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    fmt.Println(fi16, v3)
                    if (v4 > (-(v4))) {
                        v55 := GenS3{F0: fi16, F1: 798, F2: 96}
                        var v56 bool = false
                        var v57 int64 = 760
                        _ = v55
                        _ = v56
                        _ = v57
                    } else {
                        v58 := GenS1{F0: 2}
                        _ = v58
                    }
                }
            case 128:
                if (v1 == "") {
                    switch (fi16 ^ (fi16 % ((-554) | 1))) {
                    case 18:
                        v1 = "foo bar"
                    case 32768:
                        delete(v5, (fi16 | int(v4)))
                    case 507:
                        v59 := GenS1{F0: 4294967295}
                        _ = v59
                    }
                    _ = len(v6)
                }
            default:
                switch 32767 {
                case 128:
                    v60 := GenS5{F0: 1, F1: v1, F2: (-(v4))}
                    _ = v60
                case -769:
                    if (!true) {
                        v61 := GenS1{F0: 2}
                        _ = v61
                    } else {
                        v5[(^(fi16))] = (435 < 65535)
                        v2 -= (v2 % ((255) | 1))
                    }
                default:
                    v2 = v2
                }
            }
            _ = len(v5)
            v1 = "\n\t"
            _ = fi16
        }
        if (true || ((334 >= 1) || true)) {
            var v62 uint8 = 127
            var v63 float64 = (-2.25 / -2.25)
            _ = v62
            _ = v63
        } else {
            fmt.Println(v1, v3, v4)
            v2 -= v2
            v3 = -1
        }
        switch 128 {
        case 1:
            switch 57 {
            case 1:
                var v64 string = v1
                _ = v64
            case -2147483648:
                if false {
                    if true {
                        v65 := v5[2147483647]
                        fmt.Println(v3, v4)
                        var v66 uint = v2
                        _ = v65
                        _ = v66
                    }
                    var v67 int64 = -9223372036854775808
                    _ = v67
                }
            }
        case 74:
            v1 = v1
        case 150:
            fmt.Println(v3, v4)
        }
    }
    if (1 >= 0) {
        for fi68 := 0; fi68 < 12; fi68++ {
            v2 = v2
            v4 -= v4
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi68
        }
        var v69 float32 = -2.25
        _ = v69
    }
    switch v1 {
    case "hello":
        fmt.Println(v4, v4)
    case "":
        for fi70 := 0; fi70 < 12; fi70++ {
            v71 := GenS6{F0: v2, F1: 0}
            _ = fi70
            _ = v71
        }
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = 228
        }
    default:
        fmt.Println(v3)
    }
    var v72 int32 = 1
    var v73 string = (v1 + ("\n\t" + (v1 + (v1 + "\n\t"))))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if false {
        v74 := GenS2{F0: (!false), F1: 18446744073709551615, F2: -1}
        if false {
            if (65535 == 952) {
                if (42 > uint16(v2)) {
                    var v75 uint8 = 128
                    var v76 float64 = 1.5
                    _ = v75
                    _ = v76
                } else {
                    if false {
                        var v77 bool = (!false)
                        _ = v77
                    }
                    if true {
                        if (!(float64((((float32(1e10) - 0.0) / 1e10) * 1e10)) != float64(0.0))) {
                            for fi78 := 0; fi78 < 4; fi78++ {
                                if false {
                                    if (42 != 9223372036854775807) {
                                        v1 = ((v73 + v73) + v1)
                                        switch (fi78 >> ((255) & 63)) {
                                        case -269:
                                            var v79 string = ("" + "foo bar")
                                            _ = v79
                                        case 32768:
                                            _ = len(v5)
                                        default:
                                            if (!true) {
                                                v80 := GenS7{F0: 737}
                                                switch (("a" + v73) + v73) {
                                                case "":
                                                    fmt.Println(v4)
                                                case "a":
                                                    switch v1 {
                                                    case "foo bar":
                                                        switch fi78 {
                                                        case 42:
                                                            var v81 uint8 = 214
                                                            _ = v81
                                                        case -602:
                                                            fmt.Println(v2, v3, v73)
                                                        }
                                                    case "\n\t":
                                                        switch v73 {
                                                        case "hello":
                                                            fmt.Println(v2, v2, v2)
                                                        }
                                                    default:
                                                        v2 = v2
                                                    }
                                                case "hello":
                                                    var v82 uint = (^(v2))
                                                    _ = v82
                                                default:
                                                    switch fi78 {
                                                    case -329:
                                                        for fi83 := 0; fi83 < 14; fi83++ {
                                                            switch ("\n\t" + v1) {
                                                            case "":
                                                                v4 -= v4
                                                            case "\n\t":
                                                                var v84 uint64 = 769
                                                                _ = v84
                                                            case "a":
                                                                v85 := GenS1{F0: 59}
                                                                _ = v85
                                                            }
                                                            if true {
                                                                v86 := GenS7{F0: -708}
                                                                _ = v86
                                                            } else {
                                                                v87 := GenS8{F0: false, F1: 486}
                                                                _ = v87
                                                            }
                                                            var v88 int64 = -9223372036854775808
                                                            _ = fi83
                                                            _ = v88
                                                        }
                                                    case 128:
                                                        if true {
                                                            switch (^(fi78)) {
                                                            case 1:
                                                                fmt.Println(v3)
                                                            case 4294967295:
                                                                for fi89 := 0; fi89 < 17; fi89++ {
                                                                    v1 = v1
                                                                    var v90 int8 = (-(v4))
                                                                    _ = fi89
                                                                    _ = v90
                                                                }
                                                            case 32767:
                                                                if true {
                                                                    v91 := GenS5{F0: 18446744073709551615, F1: v1, F2: (^(v4))}
                                                                    var v92 int16 = -1
                                                                    v73 = v1
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    for fi93 := 0; fi93 < 16; fi93++ {
                                                                        fmt.Println(fi93, fi93)
                                                                        v3 = 42
                                                                        _ = fi93
                                                                    }
                                                                }
                                                            }
                                                            var v94 int = (-(fi78))
                                                            fmt.Println(v94)
                                                            _ = v94
                                                        } else {
                                                            var v95 int = -492
                                                            _ = len(v5)
                                                            v96 := GenS9{F0: v3}
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                    case 568:
                                                        for fi97 := 0; fi97 < 14; fi97++ {
                                                            _ = len(v6)
                                                            var v98 int64 = 396
                                                            _ = fi97
                                                            _ = v98
                                                        }
                                                    }
                                                }
                                                _ = v80
                                            } else {
                                                v2 += (v2 + (v2 / ((v2) | 1)))
                                                v99 := GenS10{F0: 159, F1: (^(v2)), F2: (^(v2))}
                                                _ = v99
                                            }
                                        }
                                    }
                                    v1 = v73
                                    v73 = (v73 + v1)
                                } else {
                                    v73 = v73
                                }
                                _ = fi78
                            }
                            v72 += 700
                            for fi100 := 0; fi100 < 4; fi100++ {
                                fmt.Println(v1)
                                _ = fi100
                            }
                        }
                    }
                    v2 -= v2
                }
            } else {
                for fi101 := 0; fi101 < 13; fi101++ {
                    if ((210 < 88) && (!(true && (("" + v1) >= ((v73 + "hello") + "\n\t"))))) {
                        switch (fi101 + fi101) {
                        case 1:
                            for fi102 := 0; fi102 < 20; fi102++ {
                                v5[257] = ((^(fi101)) == fi101)
                                if false {
                                    var v103 int = fi102
                                    var v104 uint32 = 0
                                    var v105 int16 = v3
                                    _ = v103
                                    _ = v104
                                    _ = v105
                                }
                                _ = fi102
                            }
                        case 313:
                            v106 := v5[772]
                            _ = v106
                        default:
                            var v107 int16 = (v3 - 0)
                            _ = v107
                        }
                        v108 := GenS10{F0: 1, F1: 834, F2: v2}
                        for fi109 := 0; fi109 < 17; fi109++ {
                            v110 := GenS6{F0: 710, F1: 2}
                            switch 9223372036854775807 {
                            case -284:
                                if (false && ((v4 >> (((v2 - 159)) & 63)) != 42)) {
                                    v111 := GenS7{F0: -9223372036854775808}
                                    if false {
                                        var v112 uint32 = 1
                                        var v113 string = v1
                                        if (!(false && ((-2.25 + (1e10 / 1e10)) != -2.25))) {
                                            switch (fi101 * -2147483648) {
                                            case 128:
                                                v73 = ""
                                            }
                                            fmt.Println(v113, v113, v4)
                                            v114 := GenS11{F0: v4, F1: (v112 >> ((616) & 63))}
                                            _ = v114
                                        }
                                        _ = v112
                                        _ = v113
                                    }
                                    var v115 float64 = ((1e10 / (3.14159 * 0.1)) * (1.5 + 1.5))
                                    _ = v111
                                    _ = v115
                                } else {
                                    var v116 bool = ((int16(fi101) != (^(v3))) || (false || true))
                                    var v117 float32 = 0.1
                                    var v118 uint64 = 1
                                    _ = v116
                                    _ = v117
                                    _ = v118
                                }
                            case 32768:
                                v1 = v73
                            }
                            fmt.Println(v4, v2)
                            _ = fi109
                            _ = v110
                        }
                        _ = v108
                    } else {
                        switch v73 {
                        case "a":
                            var v119 int16 = v3
                            _ = v119
                        }
                    }
                    v1 = (v1 + v1)
                    v2 -= (^(v2))
                    _ = fi101
                }
                if false {
                    switch v1 {
                    case "foo bar":
                        var v120 float32 = ((0.0 / (((((-2.25 * float32(((float64(1.5) * 1e10) - 1.5))) / 0.1) * 0.1) - 0.1) - -2.25)) / -2.25)
                        _ = v120
                    case "a":
                        var v121 uint16 = 602
                        _ = v121
                    }
                    switch 128 {
                    case 1:
                        v122 := v5[-9223372036854775808]
                        _ = v122
                    case -471:
                        fmt.Println(v1, v2)
                    case 890:
                        v123 := GenS7{F0: 42}
                        _ = v123
                    default:
                        switch v73 {
                        case "a":
                            switch v73 {
                            case "hello":
                                v2 = v2
                            case "foo bar":
                                var v124 uint8 = 1
                                _ = v124
                            }
                        case "foo bar":
                            for fi125 := 0; fi125 < 5; fi125++ {
                                _ = len(v6)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi125
                            }
                        }
                    }
                }
            }
        }
        _ = v74
    }
    v4 = v4
    var v126 int64 = 524
    v127 := GenS12{F0: v73, F1: 2147483648, F2: 1}
    v1 = (v73 + v73)
    var v128 string = v73
    v1 = (v1 + v1)
    delete(v5, -349)
    fmt.Println(v3, v72)
    v73 = v128
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v126
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v72
    _ = v73
}
