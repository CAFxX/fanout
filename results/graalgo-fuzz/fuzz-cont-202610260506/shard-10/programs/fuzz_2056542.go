package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint; F1 uint; F2 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint8; F1 string; F2 int32 }
type GenS11 struct { F0 string; F1 int8; F2 int32 }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int32) uint32 {
    var v5 uint8 = 1
    fmt.Println(p0, v5)
    var v6 float32 = float32(1e10)
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{fi8, -896}
        if false {
            switch "\n\t" {
            case "foo bar":
                if len(v9) > 0 {
                    v9[0] = (fi8 + (fi8 >> ((42) & 63)))
                }
            }
            var v10 float64 = -2.25
            var v11 int64 = 0
            _ = v10
            _ = v11
        } else {
            if len(v9) > 0 {
                v9[0] = fi8
            }
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            v9 = append(v9, fi12)
            switch ("hello" + "foo bar") {
            case "":
                v6 = ((v6 * -2.25) * (((v6 * 3.14159) + -2.25) + v6))
            }
            if true {
                if (!(!(!false))) {
                    v13 := GenS1{F0: true, F1: 9223372036854775807}
                    var v14 int64 = -9223372036854775808
                    if true {
                        if (false || true) {
                            var v15 int16 = 100
                            _ = v15
                        } else {
                            fmt.Println(v7, fi8, v6)
                            var v16 uint32 = 4294967295
                            _ = v16
                        }
                        var v17 string = (((("hello" + "") + (("foo bar" + (((("foo bar" + (("hello" + "foo bar") + "foo bar")) + (("a" + "a") + "\n\t")) + "a") + "foo bar")) + "foo bar")) + ("foo bar" + ("\n\t" + "hello"))) + "a")
                        if (128 == v5) {
                            if (!false) {
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(p0)
                                case "\n\t":
                                    v9 = append(v9, fi8)
                                default:
                                    var v18 bool = (!true)
                                    _ = v18
                                }
                            }
                        } else {
                            var v19 int32 = p0
                            _ = v19
                        }
                        _ = v17
                    }
                    _ = v13
                    _ = v14
                } else {
                    switch ((("\n\t" + "foo bar") + "a") + "foo bar") {
                    case "foo bar":
                        v7 = (^(v7))
                    case "":
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    default:
                        v20 := GenS1{F0: (539 == 65535), F1: -9223372036854775808}
                        p0 = p0
                        if (449 != 18446744073709551615) {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            v21 := GenS1{F0: (((("a" + "") + "hello") + ("" + "foo bar")) < ("foo bar" + "foo bar")), F1: -9223372036854775808}
                            if false {
                                var v22 string = ("\n\t" + "hello")
                                var v23 uint32 = 725
                                var v24 bool = false
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                if false {
                                    switch (fi12 - -975) {
                                    case -1:
                                        v25 := GenS1{F0: true, F1: 1}
                                        _ = v25
                                    default:
                                        v26 := GenS2{F0: 2147483648}
                                        _ = v26
                                    }
                                } else {
                                    v27 := GenS2{F0: uint32(v5)}
                                    if ("foo bar" < ("foo bar" + "a")) {
                                        if (v6 <= (v6 * ((v6 - v6) * 0.0))) {
                                            switch "a" {
                                            case "a":
                                                fmt.Println(v7, v5)
                                                var v28 int64 = -616
                                                switch ("a" + ("foo bar" + "foo bar")) {
                                                case "":
                                                    if (42 <= v28) {
                                                        v29 := GenS3{F0: (fi8 % ((int(p0)) | 1)), F1: 0}
                                                        fmt.Println(p0)
                                                        p0 += (p0 / (((p0 ^ (p0 >> ((2) & 63)))) | 1))
                                                        _ = v29
                                                    } else {
                                                        v6 += float32(1e10)
                                                    }
                                                case "hello":
                                                    v7 = v7
                                                }
                                                if false {
                                                    var v30 uint32 = 829
                                                    _ = v30
                                                } else {
                                                    v31 := GenS4{F0: "foo bar", F1: 0}
                                                    _ = v31
                                                }
                                                _ = v28
                                            case "hello":
                                                fmt.Println(fi12, fi8, v6)
                                            case "":
                                                v32 := GenS2{F0: uint32(v5)}
                                                v7 -= 495
                                                if (-2147483648 > 2147483647) {
                                                    p0 = (-(p0))
                                                }
                                                if true {
                                                    fmt.Println(v7)
                                                    switch ("foo bar" + (("\n\t" + "hello") + "\n\t")) {
                                                    case "hello":
                                                        v33 := GenS3{F0: (fi8 % ((32768) | 1)), F1: 73}
                                                        _ = v33
                                                    case "":
                                                        v6 -= v6
                                                    case "foo bar":
                                                        p0 = (p0 >> ((2) & 63))
                                                    }
                                                    var v34 int16 = 878
                                                    _ = v34
                                                } else {
                                                    if (-1 == 994) {
                                                        switch fi8 {
                                                        case 525:
                                                            if false {
                                                                var v35 uint32 = 339
                                                                v5 -= v5
                                                                v35 += (v35 >> ((255) & 63))
                                                                _ = v35
                                                            } else {
                                                                fmt.Println(fi12, v7, v6)
                                                            }
                                                        case -1:
                                                            switch "a" {
                                                            case "a":
                                                                var v36 int32 = (p0 << ((65535) & 63))
                                                                _ = v36
                                                            case "foo bar":
                                                                fmt.Println(fi12, v7)
                                                            default:
                                                                var v37 int = fi12
                                                                var v38 int16 = 42
                                                                switch fi8 {
                                                                case 195:
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v39 int64 = 0
                                                                        _ = v39
                                                                    case "":
                                                                        v40 := GenS5{F0: 657, F1: (^(fi12)), F2: 440}
                                                                        _ = v40
                                                                    case "a":
                                                                        var v41 int = fi8
                                                                        _ = v41
                                                                    default:
                                                                        var v42 int = v37
                                                                        _ = v42
                                                                    }
                                                                case 1:
                                                                    var v43 uint16 = 0
                                                                    _ = v43
                                                                case -1:
                                                                    if (-2.25 >= float64(v6)) {
                                                                        v38 += v38
                                                                        v38 = 1
                                                                        v44 := GenS6{F0: v5, F1: 9223372036854775807}
                                                                        _ = v44
                                                                    }
                                                                default:
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                }
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        case 32767:
                                                            var v45 int16 = 42
                                                            _ = v45
                                                        }
                                                        v46 := GenS2{F0: 966}
                                                        var v47 uint32 = 624
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    switch ("foo bar" + "a") {
                                                    case "":
                                                        v5 -= (v5 + v5)
                                                    case "\n\t":
                                                        var v48 uint16 = 2
                                                        _ = v48
                                                    }
                                                    _ = len(v9)
                                                }
                                                _ = v32
                                            default:
                                                _ = len(v9)
                                            }
                                            v7 = v7
                                        }
                                        var v49 uint16 = 481
                                        switch fi12 {
                                        case 128:
                                            v6 = 1.5
                                        default:
                                            if true {
                                                if ((fi12 != (fi8 * 50)) || (int32(v5) >= (^(p0)))) {
                                                    fmt.Println(p0)
                                                }
                                            } else {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                v50 := GenS6{F0: (v5 ^ v5), F1: -1}
                                                _ = v50
                                            }
                                        }
                                        _ = v49
                                    } else {
                                        v5 = (^(v5))
                                    }
                                    _ = v27
                                }
                                p0 += -1
                                switch (fi12 ^ fi12) {
                                case -9223372036854775808:
                                    v6 = float32(((0.0 - ((float64(v6) + 0.0) + (3.14159 / 1.0))) / (1e10 / 1.5)))
                                }
                            }
                            _ = v21
                        } else {
                            v9 = append(v9, 128)
                        }
                        if true {
                            v6 += float32((0.1 + (3.14159 * (3.14159 - 0.1))))
                        }
                        var v51 uint32 = 1
                        _ = v20
                        _ = v51
                    }
                    var v52 string = "a"
                    switch 0 {
                    case 42:
                        var v53 int8 = -1
                        _ = v53
                    case 127:
                        if (-1 != -489) {
                            var v54 uint = 1
                            v55 := GenS3{F0: (fi8 / ((int(v5)) | 1)), F1: 1}
                            _ = v54
                            _ = v55
                        } else {
                            var v56 int = fi8
                            _ = v56
                        }
                    case -873:
                        if false {
                            fmt.Println(v6, v7, fi8)
                            var v57 int16 = -1
                            _ = v57
                        } else {
                            _ = len(v9)
                            switch fi8 {
                            case 128:
                                v9 = append(v9, (^(fi12)))
                            }
                        }
                    }
                    switch (fi12 % ((32767) | 1)) {
                    case 32767:
                        var v58 bool = true
                        _ = v58
                    case 42:
                        var v59 uint32 = 0
                        _ = v59
                    case 32768:
                        v7 -= v7
                    default:
                        if len(v9) > 0 {
                            v9[0] = fi12
                        }
                    }
                    if (1 >= -9223372036854775808) {
                        if ((fi12 + (fi12 ^ fi8)) <= (-(fi8))) {
                            _ = len(v9)
                            v5 = 128
                        } else {
                            v60 := GenS5{F0: 475, F1: (fi12 & fi12), F2: 873}
                            if false {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                switch 32768 {
                                case 32768:
                                    var v61 int16 = 888
                                    if false {
                                        if false {
                                            v62 := GenS4{F0: (v52 + ((v52 + "a") + "\n\t")), F1: 42}
                                            v5 = v5
                                            var v63 int8 = 100
                                            _ = v62
                                            _ = v63
                                        }
                                    } else {
                                        v52 = ((v52 + ((v52 + (v52 + "foo bar")) + v52)) + v52)
                                    }
                                    p0 = p0
                                    _ = v61
                                default:
                                    if (!((true && true) && (-1 < 451))) {
                                        _ = len(v9)
                                        if len(v9) > 0 {
                                            v9[0] = (fi8 << ((259) & 63))
                                        }
                                        v9 = append(v9, 70)
                                    } else {
                                        var v64 uint64 = v7
                                        _ = v64
                                    }
                                }
                                if false {
                                    fmt.Println(v7)
                                    v65 := GenS7{F0: 255, F1: 255, F2: false}
                                    _ = v65
                                }
                            } else {
                                if ((v52 > ((v52 + ("\n\t" + v52)) + "a")) && true) {
                                    var v66 string = v52
                                    v7 = (v7 % ((2) | 1))
                                    switch v66 {
                                    case "a":
                                        switch v52 {
                                        case "":
                                            v52 = (("\n\t" + v66) + ("a" + "a"))
                                        case "foo bar":
                                            if true {
                                                var v67 uint = 362
                                                var v68 int16 = 1000
                                                if false {
                                                    var v69 uint8 = uint8(v7)
                                                    _ = v69
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                case "":
                                                    if (!((v68 ^ (^(v68))) < (v68 ^ v68))) {
                                                        var v70 uint16 = 65535
                                                        if true {
                                                            v71 := GenS5{F0: 686, F1: (fi8 - int(v7)), F2: 2147483647}
                                                            fmt.Println(fi8)
                                                            if (!true) {
                                                                var v72 int32 = p0
                                                                v73 := GenS5{F0: 2147483647, F1: (fi8 | int(v7)), F2: uint32(v7)}
                                                                fmt.Println(v7)
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            _ = v71
                                                        } else {
                                                            var v74 string = v52
                                                            v5 += 2
                                                            v66 = v52
                                                            _ = v74
                                                        }
                                                        v75 := GenS4{F0: v52, F1: 1}
                                                        _ = v70
                                                        _ = v75
                                                    } else {
                                                        var v76 float64 = 0.0
                                                        var v77 uint = (v67 ^ v67)
                                                        if (v77 < v77) {
                                                            v68 += (v68 / (((-(v68))) | 1))
                                                        } else {
                                                            v9 = append(v9, fi12)
                                                            var v78 int16 = (-(v68))
                                                            var v79 uint8 = v5
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    if (!(1 == 1)) {
                                                        var v80 uint32 = uint32(v5)
                                                        v7 = (v7 << (((v67 | (^(v67)))) & 63))
                                                        v68 -= (-(v68))
                                                        _ = v80
                                                    }
                                                    var v81 float64 = (-2.25 * 1e10)
                                                    var v82 int32 = (p0 << (((v67 >> (((v67 | 423)) & 63))) & 63))
                                                    _ = v81
                                                    _ = v82
                                                }
                                                fmt.Println(fi12, fi12)
                                                var v83 uint16 = 0
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                _ = v67
                                                _ = v68
                                                _ = v83
                                            } else {
                                                v9 = append(v9, (fi8 * -359))
                                                switch fi8 {
                                                case 1:
                                                    v5 -= 128
                                                case 0:
                                                    switch v66 {
                                                    case "":
                                                        v84 := GenS1{F0: ((383 > 99) || (true || ((0 >= v7) && (0 > 42)))), F1: 676}
                                                        _ = v84
                                                    case "a":
                                                        if false {
                                                            v85 := GenS8{F0: (^(fi12))}
                                                            switch fi12 {
                                                            case 2147483647:
                                                                v5 += (v5 << ((65535) & 63))
                                                            case -9223372036854775808:
                                                                var v86 uint16 = 452
                                                                _ = v86
                                                            case -871:
                                                                switch (fi12 | fi8) {
                                                                case 32768:
                                                                    var v87 uint32 = 42
                                                                    _ = v87
                                                                }
                                                            default:
                                                                v7 = (^(v7))
                                                            }
                                                            _ = v85
                                                        }
                                                    case "foo bar":
                                                        var v88 int64 = 9223372036854775807
                                                        _ = v88
                                                    default:
                                                        if false {
                                                            v7 = 4294967295
                                                            v52 = (v52 + ("hello" + v52))
                                                        }
                                                    }
                                                default:
                                                    v9 = append(v9, 1)
                                                }
                                            }
                                        case "hello":
                                            v6 = v6
                                        default:
                                            if false {
                                                var v89 uint16 = 42
                                                var v90 int32 = p0
                                                var v91 bool = false
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                            } else {
                                                v92 := GenS9{F0: 4294967295, F1: p0}
                                                _ = v92
                                            }
                                        }
                                    case "":
                                        if false {
                                            if (!false) {
                                                switch v52 {
                                                case "":
                                                    if (true || (0.0 == -2.25)) {
                                                        v93 := GenS2{F0: 2147483648}
                                                        var v94 int64 = 9223372036854775807
                                                        if (false || ((v7 < v7) && true)) {
                                                            v95 := GenS2{F0: 0}
                                                            p0 = (p0 / (((p0 << ((475) & 63))) | 1))
                                                            v96 := GenS10{F0: v5, F1: (((v66 + (((v52 + (v66 + "foo bar")) + ("\n\t" + (("\n\t" + (("" + v66) + "foo bar")) + ("\n\t" + ("" + v52))))) + v66)) + v52) + ((v52 + v66) + "a")), F2: p0}
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            var v97 uint32 = 2147483647
                                                            _ = v97
                                                        }
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    v98 := GenS7{F0: 4294967295, F1: 274, F2: (4294967295 >= 779)}
                                                    if true {
                                                        var v99 int16 = 286
                                                        var v100 int16 = v99
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        fmt.Println(v66, v66, v52)
                                                    }
                                                    _ = v98
                                                case "foo bar":
                                                    if (((!false) || false) || ((v5 - (^(v5))) < v5)) {
                                                        v101 := GenS11{F0: "a", F1: 127, F2: (p0 | (p0 % (((p0 >> ((262) & 63))) | 1)))}
                                                        p0 = 42
                                                        _ = v101
                                                    }
                                                case "hello":
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                default:
                                                    if true {
                                                        if (!false) {
                                                            if false {
                                                                if len(v9) > 0 {
                                                                    v9[0] = -164
                                                                }
                                                            } else {
                                                                var v102 float64 = 1e10
                                                                _ = v102
                                                            }
                                                            var v103 uint64 = (v7 * 392)
                                                            v7 -= (v103 / ((v103) | 1))
                                                            _ = v103
                                                        }
                                                        v104 := GenS6{F0: (^(v5)), F1: 42}
                                                        var v105 int64 = 0
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    v7 -= (v7 + (v7 << ((42) & 63)))
                                                    if false {
                                                        fmt.Println(v7)
                                                    } else {
                                                        _ = len(v9)
                                                        var v106 uint16 = 42
                                                        _ = v106
                                                    }
                                                    switch (^(fi12)) {
                                                    case -2147483648:
                                                        switch "hello" {
                                                        case "a":
                                                            if false {
                                                                switch v66 {
                                                                case "a":
                                                                    v6 -= (v6 / (v6 + 1e10))
                                                                case "\n\t":
                                                                    if (65535 < 42) {
                                                                        var v107 uint32 = 42
                                                                        var v108 int8 = 83
                                                                        v9 = append(v9, (fi12 >> ((2) & 63)))
                                                                        _ = v107
                                                                        _ = v108
                                                                    }
                                                                case "foo bar":
                                                                    var v109 uint64 = v7
                                                                    _ = v109
                                                                }
                                                            }
                                                        case "hello":
                                                            var v110 string = ("a" + v66)
                                                            _ = v110
                                                        }
                                                    case 4294967295:
                                                        v111 := GenS9{F0: 4294967295, F1: p0}
                                                        _ = v111
                                                    }
                                                    var v112 bool = ((v52 < "\n\t") && (false || ((v7 - (v7 % ((2) | 1))) <= (v7 & 240))))
                                                    v52 = v66
                                                    _ = v112
                                                }
                                                var v113 uint32 = 2
                                                _ = v113
                                            } else {
                                                var v114 int16 = int16(v7)
                                                v114 -= (-(v114))
                                                if true {
                                                    v6 += v6
                                                }
                                                _ = v114
                                            }
                                            _ = len(v9)
                                        }
                                    }
                                    _ = v66
                                }
                            }
                            _ = v60
                        }
                        var v115 int8 = -31
                        _ = v115
                    } else {
                        v116 := GenS12{F0: (p0 - p0), F1: 598}
                        v117 := GenS8{F0: 936}
                        _ = v116
                        _ = v117
                    }
                    fmt.Println(v5, v7, v5)
                    _ = v52
                }
            }
            _ = fi12
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint8, p1 uint8) int16 {
    v118 := GenS13{F0: 0, F1: 2}
    if false {
        var v119 int16 = 1000
        fmt.Println(v119, p1)
        v119 -= (v119 * v119)
        _ = v119
    } else {
        v120 := GenS8{F0: 32767}
        v121 := []int{1}
        switch (("hello" + "a") + "\n\t") {
        case "":
            if (128 <= p0) {
                var v122 float64 = (-2.25 + 3.14159)
                for fi123 := 0; fi123 < 16; fi123++ {
                    var v124 int = (fi123 >> ((4294967295) & 63))
                    v125 := GenS10{F0: p0, F1: "", F2: -2147483648}
                    var v126 uint32 = genFn0(-2147483648)
                    _ = fi123
                    _ = v124
                    _ = v125
                    _ = v126
                }
                v122 = (v122 + (v122 / (v122 / 0.1)))
                _ = v122
            }
        case "\n\t":
            p0 = p0
        default:
            p1 -= p0
        }
        _ = v120
        _ = v121
    }
    var v127 int64 = -317
    switch 2147483647 {
    case -954:
        if (1.5 == 1.5) {
            fmt.Println(v127)
            if ((32767 < 805) || (32767 == 505)) {
                var v128 int8 = -1
                fmt.Println(p0, v127)
                _ = v128
            } else {
                for fi129 := 0; fi129 < 5; fi129++ {
                    v130 := []int{(fi129 * -702)}
                    v131 := GenS8{F0: fi129}
                    _ = fi129
                    _ = v130
                    _ = v131
                }
            }
            var v132 int32 = 799
            _ = v132
        } else {
            v133 := []int{}
            _ = v133
        }
    case 1:
        for fi134 := 0; fi134 < 3; fi134++ {
            var v135 int32 = 0
            _ = fi134
            _ = v135
        }
    case 32768:
        var v136 int32 = -1
        _ = v136
    }
    var v137 int16 = -32768
    _ = p0
    _ = p1
    _ = v118
    _ = v127
    _ = v137
    return -1
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 uint64 = 336
    if (42 <= 0) {
        v1 = 42
        fmt.Println(v1)
    } else {
        fmt.Println(v4, v4)
        v1 -= 1
    }
    v3[65535] = -450
    v4 = (v4 & (v4 + (^(v4))))
    var v138 float64 = float64(3.14159)
    var v139 int16 = genFn1(0, 106)
    var v140 float32 = 0.0
    var v141 uint32 = 0
    var v142 int = 127
    v1 = 84
    v142 = (v142 % ((0) | 1))
    fmt.Println(v140, v140)
    v143 := []int{(v142 / ((v142) | 1)), (v142 << ((42) & 63)), (v142 * (v142 + (v142 / (((-(v142))) | 1))))}
    fmt.Println(v140, v139, v140)
    if (true && true) {
        for fi144 := 0; fi144 < 6; fi144++ {
            v138 += float64(v140)
            if true {
                var v145 uint64 = 18446744073709551615
                v139 += 1
                _ = v145
            } else {
                if true {
                    var v146 int16 = (v139 | -218)
                    if (!true) {
                        v146 = 1
                    } else {
                        for fi147 := 0; fi147 < 15; fi147++ {
                            v139 -= (-(v146))
                            fmt.Println(v140, fi147)
                            _ = fi147
                        }
                    }
                    _ = v146
                } else {
                    fmt.Println(v4, v1)
                    for fi148 := 0; fi148 < 17; fi148++ {
                        var v149 uint8 = 128
                        switch "hello" {
                        case "\n\t":
                            _ = len(v143)
                        case "hello":
                            if (0 == genFn1(v149, 42)) {
                                var v150 int = (fi148 | fi148)
                                _ = v150
                            }
                        }
                        _ = fi148
                        _ = v149
                    }
                    fmt.Println(v141, v138, v1)
                }
                fmt.Println(v139, v138)
                v138 = (-2.25 / ((v138 * (float64(float32((0.0 - (3.14159 / ((-2.25 - float64(1.5)) + v138))))) / v138)) - 0.0))
            }
            v151 := GenS6{F0: 42, F1: 0}
            _ = fi144
            _ = v151
        }
        var v152 int64 = 9223372036854775807
        var v153 float32 = v140
        _ = v152
        _ = v153
    }
    var v154 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v143:", len(v143), v143)
    fmt.Println("v154:", v154)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v138
    _ = v139
    _ = v140
    _ = v141
    _ = v142
    _ = v143
    _ = v154
    _ = v2
    _ = v3
    _ = v4
}
