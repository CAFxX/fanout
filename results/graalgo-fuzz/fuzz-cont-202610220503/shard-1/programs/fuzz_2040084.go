package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 int8; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 string }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int16; F2 string }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 bool; F1 bool; F2 int8 }
type GenS10 struct { F0 uint; F1 int16 }

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 100
    var v3 float64 = 0.0
    var v4 int64 = 485
    var v5 float32 = 0.0
    v4 = 1
    fmt.Println(v2)
    v6 := []int{42}
    if (42 >= 42) {
        var v7 uint16 = 721
        _ = v7
    } else {
        for fi8 := 0; fi8 < 10; fi8++ {
            v3 -= 0.1
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint32 = 42
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            switch fi11 {
            case -1:
                if true {
                    var v12 float32 = 0.0
                    _ = v12
                } else {
                    v13 := GenS1{F0: 255, F1: -209, F2: "a"}
                    if len(v6) > 0 {
                        v6[0] = 9223372036854775807
                    }
                    var v14 float32 = 3.14159
                    _ = v13
                    _ = v14
                }
            case 32768:
                var v15 uint16 = 255
                _ = v15
            case -346:
                _ = len(v6)
            }
            for fi16 := 0; fi16 < 13; fi16++ {
                _ = len(v6)
                if false {
                    v17 := GenS2{F0: (v2 | v2), F1: 2, F2: -645}
                    if (128 > (fi11 * fi16)) {
                        var v18 uint64 = 42
                        if false {
                            var v19 int = -2147483648
                            if true {
                                if true {
                                    var v20 int = 0
                                    switch (v19 >> ((466) & 63)) {
                                    case 2147483647:
                                        if true {
                                            v21 := GenS3{F0: 33, F1: int64(v2), F2: ("" + (v1 + (v1 + (v1 + v1))))}
                                            _ = v21
                                        }
                                    case 42:
                                        if false {
                                            switch int(v4) {
                                            case 443:
                                                if true {
                                                    if true {
                                                        v22 := GenS1{F0: 42, F1: 42, F2: "\n\t"}
                                                        _ = v22
                                                    } else {
                                                        if false {
                                                            if ((((v2 & -128) <= v2) && (0 != 1)) && false) {
                                                                v1 = "\n\t"
                                                                var v23 bool = false
                                                                _ = v23
                                                            } else {
                                                                if (true || true) {
                                                                    _ = len(v6)
                                                                    v3 += 1e10
                                                                }
                                                                _ = len(v6)
                                                                fmt.Println(v5, v3)
                                                            }
                                                        }
                                                        if false {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                fmt.Println(v1, fi11, v19)
                                                            default:
                                                                if false {
                                                                    _ = len(v6)
                                                                } else {
                                                                    _ = len(v6)
                                                                    var v24 uint64 = v18
                                                                    _ = v24
                                                                }
                                                            }
                                                            if false {
                                                                var v25 string = v1
                                                                _ = v25
                                                            } else {
                                                                if false {
                                                                    var v26 uint16 = 32767
                                                                    _ = v26
                                                                }
                                                            }
                                                            fmt.Println(fi16)
                                                        } else {
                                                            v1 = ""
                                                            v20 += -706
                                                            v27 := GenS4{F0: v18, F1: 192}
                                                            _ = v27
                                                        }
                                                    }
                                                    v6 = append(v6, fi11)
                                                } else {
                                                    var v28 uint8 = 117
                                                    _ = v28
                                                }
                                            }
                                        }
                                    default:
                                        v1 = (v1 + v1)
                                    }
                                    _ = v20
                                } else {
                                    switch (fi16 / (((^(fi16))) | 1)) {
                                    case 32768:
                                        switch (v19 % ((v19) | 1)) {
                                        case 345:
                                            if len(v6) > 0 {
                                                v6[0] = (fi11 - -337)
                                            }
                                        case 0:
                                            var v29 uint64 = (v18 - 42)
                                            _ = v29
                                        }
                                    case -105:
                                        v18 = v18
                                    case 421:
                                        var v30 int32 = -352
                                        v31 := GenS4{F0: v18, F1: 241}
                                        v30 = 0
                                        _ = v30
                                        _ = v31
                                    default:
                                        v4 = (v4 ^ (v4 - v4))
                                    }
                                }
                            }
                            _ = v19
                        } else {
                            var v32 int8 = (v2 & 123)
                            var v33 uint16 = 700
                            _ = v32
                            _ = v33
                        }
                        var v34 string = v1
                        _ = v18
                        _ = v34
                    }
                    _ = v17
                } else {
                    v35 := GenS1{F0: 1, F1: -1, F2: (v1 + ((("" + v1) + (v1 + v1)) + ("" + v1)))}
                    fmt.Println(v3, v5, v3)
                    _ = v35
                }
                fmt.Println(v1)
                _ = fi16
            }
            _ = fi11
        }
        if ((v2 != v2) && (v4 != v4)) {
            switch -1 {
            case 2147483647:
                _ = len(v6)
            default:
                v6 = append(v6, 42)
            }
        } else {
            v5 = 1e10
            v36 := GenS2{F0: int8(v4), F1: 0, F2: 636}
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 uint8 = 128
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 int16 = -719
            switch (v1 + ("hello" + (("" + v1) + "hello"))) {
            case "hello":
                if (true || (v1 < ("\n\t" + (v1 + v1)))) {
                    var v41 int8 = v2
                    var v42 uint16 = 65535
                    if len(v6) > 0 {
                        v6[0] = 32767
                    }
                    _ = v41
                    _ = v42
                } else {
                    var v43 uint32 = 2147483648
                    v43 = v43
                    _ = v43
                }
            default:
                if ((fi39 << ((18446744073709551615) & 63)) == 32768) {
                    v44 := GenS4{F0: 0, F1: 2}
                    _ = v44
                }
            }
            if (((v40 << ((65535) & 63)) > (v40 & v40)) && (-747 <= 0)) {
                switch -2147483648 {
                case -9223372036854775808:
                    if false {
                        if (!(fi39 > (fi37 >> ((42) & 63)))) {
                            if (42 != -1) {
                                switch v1 {
                                case "foo bar":
                                    switch v1 {
                                    case "":
                                        fmt.Println(v3)
                                    case "\n\t":
                                        if (65535 != 255) {
                                            switch "" {
                                            case "hello":
                                                fmt.Println(v2, v38, fi37)
                                            case "\n\t":
                                                v45 := GenS2{F0: v2, F1: 1, F2: -1}
                                                _ = v45
                                            }
                                        } else {
                                            v46 := GenS5{F0: 248, F1: (v40 ^ v40), F2: v1}
                                            v1 = ""
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v46
                                        }
                                    }
                                case "a":
                                    switch "" {
                                    case "hello":
                                        var v47 uint32 = 629
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v47
                                    case "\n\t":
                                        var v48 float32 = v5
                                        _ = v48
                                    case "foo bar":
                                        v49 := GenS6{F0: (1 <= 4294967295), F1: 0}
                                        _ = v49
                                    default:
                                        v2 = v2
                                    }
                                }
                                v38 = (^(v38))
                            } else {
                                v4 = 0
                                var v50 uint = 235
                                var v51 bool = true
                                _ = v50
                                _ = v51
                            }
                            var v52 bool = false
                            _ = v52
                        } else {
                            v3 = v3
                            if ((1e10 < v3) || ((v4 & v4) == (v4 - v4))) {
                                v2 = v2
                                var v53 uint8 = (v38 | 2)
                                _ = v53
                            } else {
                                var v54 string = v1
                                var v55 uint8 = (v38 ^ v38)
                                if (true && (false && (true || true))) {
                                    fmt.Println(v40, v5)
                                    switch v1 {
                                    case "\n\t":
                                        if (v5 >= (v5 / float32(float64(0.1)))) {
                                            v4 -= int64(v38)
                                            var v56 int64 = v4
                                            var v57 uint8 = 20
                                            _ = v56
                                            _ = v57
                                        }
                                    case "hello":
                                        var v58 float32 = v5
                                        _ = v58
                                    case "foo bar":
                                        switch "" {
                                        case "foo bar":
                                            var v59 bool = true
                                            _ = v59
                                        case "hello":
                                            if len(v6) > 0 {
                                                v6[0] = (fi37 / ((fi37) | 1))
                                            }
                                            v3 = v3
                                        case "a":
                                            if (0 == 65535) {
                                                if false {
                                                    var v60 uint = 65535
                                                    v38 -= 128
                                                    _ = v60
                                                }
                                                v6 = append(v6, (fi37 & (^(fi37))))
                                                v61 := GenS6{F0: true, F1: 0}
                                                v62 := GenS2{F0: (v2 % ((v2) | 1)), F1: 18446744073709551615, F2: 42}
                                                _ = v61
                                                _ = v62
                                            }
                                        default:
                                            v38 -= v55
                                        }
                                    default:
                                        switch 4294967295 {
                                        case -266:
                                            switch v1 {
                                            case "\n\t":
                                                if (v2 <= 6) {
                                                    v3 = 0.0
                                                } else {
                                                    fmt.Println(fi39, v2, v54)
                                                    if ((2147483648 >= 32) && ((float64(1.5) * v3) == -2.25)) {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v38 = v38
                                                        default:
                                                            if (true && ((float64(v5) + 3.14159) != (v3 / 3.14159))) {
                                                                if len(v6) > 0 {
                                                                    v6[0] = (fi39 + fi39)
                                                                }
                                                            }
                                                        }
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    } else {
                                                        v54 = v54
                                                        var v63 uint32 = 2
                                                        _ = v63
                                                    }
                                                }
                                            case "foo bar":
                                                var v64 float32 = ((3.14159 - ((0.1 + (v5 - (v5 / v5))) - float32(v3))) / 0.1)
                                                _ = v64
                                            default:
                                                if (!((fi37 & fi37) <= fi37)) {
                                                    fmt.Println(v40)
                                                }
                                            }
                                        case 32768:
                                            var v65 uint64 = 0
                                            _ = v65
                                        }
                                    }
                                    switch 792 {
                                    case -1:
                                        if true {
                                            switch ((v54 + v1) + v1) {
                                            case "foo bar":
                                                v4 = v4
                                            default:
                                                if (v3 == (v3 + 1e10)) {
                                                    v3 = v3
                                                } else {
                                                    v66 := GenS5{F0: v55, F1: 42, F2: "hello"}
                                                    var v67 uint = 513
                                                    v54 = ((v1 + v54) + v54)
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v68 := GenS7{F0: 0, F1: (^(v38)), F2: 944}
                                    v69 := GenS2{F0: 41, F1: 117, F2: 627}
                                    var v70 int8 = (v2 ^ -128)
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                                _ = v54
                                _ = v55
                            }
                        }
                        if (0 != 445) {
                            var v71 uint16 = uint16(v38)
                            switch fi37 {
                            case -1:
                                v72 := GenS8{F0: 4294967295, F1: true}
                                _ = v72
                            case 1:
                                var v73 int32 = int32(v4)
                                fmt.Println(v5, v1)
                                switch fi37 {
                                case 582:
                                    v4 = v4
                                case 0:
                                    if ((fi39 | 128) >= (fi39 * (fi37 & (fi39 * 32767)))) {
                                        v71 -= v71
                                        v5 -= ((v5 + v5) - (-2.25 + v5))
                                        if false {
                                            v5 = v5
                                            if (!true) {
                                                if ((((v2 | v2) >= int8(fi37)) || (!false)) || false) {
                                                    switch "\n\t" {
                                                    case "":
                                                        if true {
                                                            v4 = 928
                                                            if true {
                                                                if true {
                                                                    var v74 int8 = int8(v38)
                                                                    switch fi37 {
                                                                    case 128:
                                                                        v73 = v73
                                                                    case -1:
                                                                        v71 -= uint16(v38)
                                                                    case 9223372036854775807:
                                                                        var v75 float64 = v3
                                                                        _ = v75
                                                                    }
                                                                    v5 = 0.0
                                                                    _ = v74
                                                                } else {
                                                                    if (!false) {
                                                                        var v76 int32 = int32(v38)
                                                                        _ = v76
                                                                    } else {
                                                                        var v77 int16 = int16(v2)
                                                                        _ = v77
                                                                    }
                                                                    _ = len(v6)
                                                                    if true {
                                                                        fmt.Println(v1, v1, fi37)
                                                                        var v78 uint = 402
                                                                        _ = v78
                                                                    }
                                                                }
                                                                if false {
                                                                    var v79 bool = true
                                                                    fmt.Println(v5, v38)
                                                                    _ = v79
                                                                }
                                                            }
                                                            v80 := GenS4{F0: uint64(v71), F1: 115}
                                                            _ = v80
                                                        } else {
                                                            v4 -= (v4 | v4)
                                                            var v81 int32 = (v73 ^ 42)
                                                            var v82 int32 = (v81 >> ((19) & 63))
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    case "foo bar":
                                                        var v83 int8 = v2
                                                        _ = v83
                                                    default:
                                                        var v84 int = -9223372036854775808
                                                        _ = v84
                                                    }
                                                    v6 = append(v6, (fi39 << ((65535) & 63)))
                                                    var v85 int32 = v73
                                                    _ = v85
                                                } else {
                                                    v86 := GenS7{F0: 42, F1: v38, F2: 520}
                                                    _ = v86
                                                }
                                            } else {
                                                if true {
                                                    switch ((v1 + v1) + ("" + "\n\t")) {
                                                    case "foo bar":
                                                        fmt.Println(fi37, fi37)
                                                        var v87 uint = 65535
                                                        _ = v87
                                                    default:
                                                        v73 += v73
                                                    }
                                                    if false {
                                                        var v88 float64 = (v3 + 3.14159)
                                                        v73 = (v73 % ((466) | 1))
                                                        v38 = (v38 ^ v38)
                                                        _ = v88
                                                    } else {
                                                        v89 := GenS4{F0: 1, F1: 0}
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        _ = v89
                                                    }
                                                    v5 += ((float32(v3) * (((v5 - (3.14159 * (v5 * 0.0))) - v5) * (3.14159 - v5))) + v5)
                                                    var v90 float32 = -2.25
                                                    _ = v90
                                                } else {
                                                    if len(v6) > 0 {
                                                        v6[0] = (-(fi39))
                                                    }
                                                }
                                                var v91 int8 = v2
                                                fmt.Println(v1, v3, v40)
                                                var v92 bool = ((v71 + (v71 & 42)) >= v71)
                                                _ = v91
                                                _ = v92
                                            }
                                        }
                                    }
                                case 476:
                                    fmt.Println(v3, v40)
                                default:
                                    fmt.Println(fi39)
                                    switch (fi37 << ((18446744073709551615) & 63)) {
                                    case 4294967295:
                                        var v93 int8 = -113
                                        _ = v93
                                    case -9223372036854775808:
                                        if true {
                                            var v94 float32 = v5
                                            _ = v94
                                        }
                                        v95 := GenS7{F0: 714, F1: v38, F2: 4294967295}
                                        v71 -= v71
                                        _ = v95
                                    default:
                                        if false {
                                            v96 := GenS8{F0: 0, F1: false}
                                            v71 = v71
                                            var v97 int64 = v4
                                            _ = v96
                                            _ = v97
                                        }
                                    }
                                    var v98 int = 32767
                                    v99 := GenS9{F0: (42 >= 255), F1: true, F2: v2}
                                    _ = v98
                                    _ = v99
                                }
                                _ = v73
                            }
                            v100 := GenS9{F0: true, F1: ((v71 & v71) != (v71 ^ (v71 / ((32767) | 1)))), F2: (v2 / (((v2 | 105)) | 1))}
                            v101 := GenS6{F0: (!false), F1: 42}
                            _ = v100
                            _ = v101
                            _ = v71
                        } else {
                            v3 -= ((0.1 * 3.14159) - (v3 * ((0.0 + (0.0 * v3)) / float64(((v5 * -2.25) - ((v5 / ((((1.5 - (0.0 / -2.25)) - float32((v3 / (v3 + (v3 + (float64(v5) * -2.25)))))) * (float32((v3 - v3)) + (v5 + (v5 * 1e10)))) / v5)) * v5))))))
                            var v102 int = (fi37 % ((-9223372036854775808) | 1))
                            var v103 int64 = v4
                            _ = v102
                            _ = v103
                        }
                        if (true || false) {
                            var v104 uint64 = 2
                            if false {
                                v4 -= 9223372036854775807
                                var v105 uint = 65535
                                v106 := GenS2{F0: (v2 ^ (v2 / ((v2) | 1))), F1: 2, F2: 2147483647}
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = v105
                                _ = v106
                            }
                            var v107 uint32 = 508
                            _ = v104
                            _ = v107
                        }
                        switch (v1 + v1) {
                        case "a":
                            v5 += v5
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    } else {
                        switch (fi39 | (fi37 % ((702) | 1))) {
                        case 0:
                            fmt.Println(v38, v2, fi39)
                        case 1:
                            v108 := GenS2{F0: -128, F1: 18446744073709551615, F2: -2147483648}
                            _ = v108
                        case -9223372036854775808:
                            if len(v6) > 0 {
                                v6[0] = (fi39 & (^(fi39)))
                            }
                        }
                        if (((1e10 / v5) <= float32((v3 * v3))) || true) {
                            var v109 float32 = (v5 - ((v5 / (v5 / v5)) * 3.14159))
                            _ = len(v6)
                            _ = v109
                        }
                        v110 := GenS7{F0: 255, F1: (^(v38)), F2: 1}
                        _ = v110
                    }
                case -49:
                    var v111 float32 = (v5 - 3.14159)
                    _ = v111
                }
                v112 := GenS4{F0: 1, F1: 491}
                _ = v112
            }
            _ = fi39
            _ = v40
        }
        for fi113 := 0; fi113 < 10; fi113++ {
            var v114 int8 = v2
            switch fi113 {
            case 127:
                v2 -= 100
            case 128:
                if (1 <= 2147483647) {
                    v2 -= v114
                    v115 := GenS10{F0: 413, F1: -1}
                    _ = v115
                } else {
                    if (!((169 <= 42) && (((!true) && (!((v5 * v5) == (((v5 * v5) + (((v5 + v5) / v5) * (v5 - (v5 - (v5 - v5))))) - v5)))) && (38 > 18446744073709551615)))) {
                        var v116 int8 = (v2 - v2)
                        _ = v116
                    }
                    fmt.Println(v114, v4, v1)
                }
            default:
                var v117 int8 = (v2 >> ((65535) & 63))
                _ = v117
            }
            v1 = v1
            if true {
                fmt.Println(v114, v2, v114)
            } else {
                fmt.Println(v1, fi113, fi113)
                var v118 float64 = v3
                if true {
                    v4 = (v4 % (((^(v4))) | 1))
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                switch (fi113 / (((^(fi37))) | 1)) {
                case 1:
                    v118 = 1e10
                }
                if len(v6) > 0 {
                    v6[0] = 127
                }
                _ = v118
            }
            _ = fi113
            _ = v114
        }
        _ = fi37
        _ = v38
    }
    var v119 int64 = -481
    if (!false) {
        fmt.Println(v119)
        v4 += v119
    }
    var v120 uint8 = 2
    v121 := GenS1{F0: v120, F1: -2147483648, F2: "hello"}
    v6 = append(v6, 2147483647)
    _ = len(v6)
    v122 := GenS4{F0: 18446744073709551615, F1: 2147483648}
    switch 0 {
    case 4294967295:
        var v123 int16 = -1
        _ = v123
    case 996:
        var v124 int64 = 0
        _ = v124
    case 32767:
        v5 -= v5
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v120 = v120
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
