package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint64; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 uint8; F1 uint }
type GenS12 struct { F0 uint8; F1 int16 }
type GenS13 struct { F0 int8; F1 uint8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int16; F1 uint64; F2 int8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int32; F1 int; F2 int }
type GenS18 struct { F0 uint32; F1 int64; F2 uint8 }

func main() {
    v1 := []uint32{42, 806, 2147483647}
    var v2 bool = true
    v3 := []string{""}
    var v4 int16 = 1000
    v5 := GenS1{F0: 4294967295, F1: 50, F2: -9223372036854775808}
    v2 = v2
    v6 := GenS1{F0: 960, F1: 42, F2: 791}
    v1 = append(v1, 4294967295)
    fmt.Println(v2, v4)
    if v2 {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    } else {
        v2 = (!(-2.25 < -2.25))
        var v7 float32 = 0.0
        for fi8 := 0; fi8 < 5; fi8++ {
            switch 127 {
            case -9223372036854775808:
                v4 += (v4 + v4)
            case 128:
                for fi9 := 0; fi9 < 18; fi9++ {
                    fmt.Println(v4, fi8, v2)
                    var v10 int32 = 423
                    var v11 uint = 1
                    v7 = 0.1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                v2 = false
            }
            _ = fi8
        }
        _ = v7
    }
    if (1 >= 4294967295) {
        _ = len(v1)
        if ((("a" + "foo bar") <= "") && v2) {
            v12 := GenS2{F0: -520}
            var v13 uint64 = 52
            _ = v12
            _ = v13
        } else {
            var v14 uint64 = 767
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    v3 = append(v3, "foo bar")
                    v17 := GenS3{F0: v2, F1: 4294967295}
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            switch "hello" {
            case "":
                var v18 uint8 = 42
                _ = v18
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v4 = -323
            }
            _ = v14
        }
        v4 -= v4
    } else {
        if true {
            if v2 {
                for fi19 := 0; fi19 < 1; fi19++ {
                    v3 = append(v3, "foo bar")
                    _ = fi19
                }
                v20 := GenS4{F0: 32767}
                _ = v20
            } else {
                if v2 {
                    if v2 {
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v21 string = "a"
                        _ = v21
                    } else {
                        fmt.Println(v4, v4)
                        v2 = (true && (((v4 >> ((257) & 63)) == (v4 * 1)) && false))
                    }
                    if v2 {
                        var v22 int16 = v4
                        var v23 int64 = 1
                        if ((-2.25 - float32(1.5)) == ((-2.25 / 0.1) / ((0.0 - 1e10) / 1e10))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (127 != 195) {
                                v23 -= v23
                                v22 -= (v4 << ((656) & 63))
                                for fi24 := 0; fi24 < 15; fi24++ {
                                    for fi25 := 0; fi25 < 16; fi25++ {
                                        v22 = v4
                                        _ = fi25
                                    }
                                    v3 = append(v3, ("a" + "a"))
                                    _ = fi24
                                }
                            }
                            for fi26 := 0; fi26 < 14; fi26++ {
                                if (-1 <= 0) {
                                    fmt.Println(v22, fi26, v4)
                                }
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    var v28 uint = 1
                                    _ = fi27
                                    _ = v28
                                }
                                _ = fi26
                            }
                        } else {
                            for fi29 := 0; fi29 < 15; fi29++ {
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v2, v4, v2)
                                }
                                fmt.Println(fi29)
                                _ = fi29
                            }
                            v23 -= v23
                            var v30 bool = ((^(v22)) >= v4)
                            _ = v30
                        }
                        _ = v22
                        _ = v23
                    }
                }
            }
            if (255 >= 65535) {
                if (151 <= 1) {
                    var v31 float64 = 1.5
                    fmt.Println(v2, v2)
                    _ = v31
                } else {
                    v32 := GenS5{F0: -1, F1: 453, F2: 41}
                    switch ("" + ((("hello" + "a") + "") + "\n\t")) {
                    case "":
                        if (507 > 7) {
                            for fi33 := 0; fi33 < 14; fi33++ {
                                v34 := GenS4{F0: 127}
                                _ = fi33
                                _ = v34
                            }
                        }
                    case "foo bar":
                        var v35 int32 = 42
                        _ = v35
                    case "a":
                        v36 := GenS1{F0: 42, F1: 782, F2: 9223372036854775807}
                        _ = v36
                    default:
                        var v37 int16 = v4
                        _ = v37
                    }
                    var v38 int8 = 0
                    _ = v32
                    _ = v38
                }
            } else {
                fmt.Println(v2, v2)
                var v39 int = -1
                _ = v39
            }
        }
        if false {
            if (65535 < 4294967295) {
                var v40 bool = (!(4294967295 != 1))
                v41 := GenS3{F0: v40, F1: 65535}
                var v42 uint8 = 128
                _ = v40
                _ = v41
                _ = v42
            } else {
                var v43 int64 = 9223372036854775807
                var v44 uint8 = 42
                if len(v1) > 0 {
                    v1[0] = 480
                }
                _ = v43
                _ = v44
            }
            var v45 uint = 255
            fmt.Println(v45, v4, v45)
            _ = v45
        } else {
            switch "hello" {
            case "a":
                if (-692 > 0) {
                    v4 = v4
                    var v46 int16 = 0
                    _ = v46
                }
            case "":
                for fi47 := 0; fi47 < 9; fi47++ {
                    for fi48 := 0; fi48 < 8; fi48++ {
                        var v49 int = (fi48 << ((4294967295) & 63))
                        var v50 uint32 = 856
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    var v51 uint = 632
                    var v52 uint = (^(v51))
                    _ = fi47
                    _ = v51
                    _ = v52
                }
            case "foo bar":
                var v53 uint = 254
                _ = v53
            default:
                for fi54 := 0; fi54 < 15; fi54++ {
                    switch "hello" {
                    case "":
                        var v55 int32 = -2147483648
                        _ = v55
                    case "hello":
                        fmt.Println(v2)
                    case "\n\t":
                        v4 -= (v4 << ((4294967295) & 63))
                    default:
                        fmt.Println(fi54)
                    }
                    _ = fi54
                }
            }
            for fi56 := 0; fi56 < 16; fi56++ {
                v57 := GenS6{F0: 251}
                if ((1 <= 255) && true) {
                    var v58 bool = (!true)
                    switch ("a" + ("a" + ("a" + "foo bar"))) {
                    case "\n\t":
                        switch (^(fi56)) {
                        case -505:
                            v4 = -1
                        case 42:
                            v59 := GenS5{F0: (fi56 << ((255) & 63)), F1: 18446744073709551615, F2: 507}
                            _ = v59
                        }
                    case "foo bar":
                        v2 = (("hello" + (("a" + ("\n\t" + ("" + ((("foo bar" + "hello") + ("hello" + "")) + "")))) + (("a" + "\n\t") + "\n\t"))) == ("a" + "a"))
                    }
                    for fi60 := 0; fi60 < 3; fi60++ {
                        if (-9223372036854775808 < 0) {
                            v2 = v58
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v4, v2, v4)
                            if v2 {
                                v58 = false
                            }
                        } else {
                            if (!v58) {
                                switch (("a" + "\n\t") + "\n\t") {
                                case "a":
                                    if v2 {
                                        v58 = v2
                                        v61 := GenS6{F0: 953}
                                        _ = v61
                                    } else {
                                        fmt.Println(v58)
                                        switch (-(fi56)) {
                                        case 466:
                                            fmt.Println(v58)
                                        case 42:
                                            v62 := GenS7{F0: 1, F1: 255, F2: -401}
                                            _ = v62
                                        case -2147483648:
                                            var v63 uint8 = 0
                                            _ = v63
                                        }
                                        switch fi56 {
                                        case 127:
                                            var v64 string = "\n\t"
                                            _ = v64
                                        case 42:
                                            switch (("a" + "") + "hello") {
                                            case "":
                                                v65 := GenS3{F0: v2, F1: 4294967295}
                                                _ = v65
                                            default:
                                                var v66 int64 = -972
                                                _ = v66
                                            }
                                        case 958:
                                            if v2 {
                                                v58 = (!false)
                                                fmt.Println(v2, fi56)
                                                switch 1 {
                                                case 1:
                                                    var v67 uint64 = 2
                                                    _ = v67
                                                case -922:
                                                    v58 = (1.5 == (1e10 * 3.14159))
                                                default:
                                                    if (!(v4 <= v4)) {
                                                        var v68 int8 = -1
                                                        _ = v68
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(fi56)
                                                    var v69 int32 = 8
                                                    fmt.Println(fi60, v2, fi56)
                                                    _ = v69
                                                } else {
                                                    if v2 {
                                                        var v70 int32 = 1
                                                        v71 := GenS4{F0: 2147483647}
                                                        v58 = (((!v2) || ((!v2) && (!v2))) || v2)
                                                        if len(v3) > 0 {
                                                            v3[0] = "\n\t"
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = ""
                                                        }
                                                        _ = v70
                                                        _ = v71
                                                    } else {
                                                        v72 := GenS8{F0: (-(v4)), F1: "a"}
                                                        v73 := GenS3{F0: true, F1: 65535}
                                                        if v2 {
                                                            v58 = true
                                                            fmt.Println(fi60, v58)
                                                            var v74 uint16 = 64
                                                            _ = v74
                                                        }
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    switch (("a" + (((("hello" + "foo bar") + ((("\n\t" + "a") + "a") + ("a" + "\n\t"))) + "hello") + "foo bar")) + "a") {
                                                    case "":
                                                        fmt.Println(v2)
                                                    default:
                                                        if false {
                                                            v75 := GenS9{F0: 4294967295, F1: (394 != 883)}
                                                            v58 = v2
                                                            _ = v75
                                                        }
                                                    }
                                                }
                                                if true {
                                                    switch "hello" {
                                                    case "a":
                                                        v58 = true
                                                    }
                                                    v76 := GenS3{F0: (0 != 0), F1: 314}
                                                    var v77 uint = 42
                                                    _ = v76
                                                    _ = v77
                                                }
                                                switch ((("hello" + ("hello" + "a")) + "a") + "a") {
                                                case "a":
                                                    var v78 int64 = 9223372036854775807
                                                    _ = v78
                                                }
                                                if true {
                                                    var v79 string = ("\n\t" + ((("" + "\n\t") + "a") + "foo bar"))
                                                    _ = v79
                                                } else {
                                                    v58 = v2
                                                    v4 = v4
                                                    v2 = v58
                                                }
                                                if true {
                                                    if (!v2) {
                                                        if ((!((fi60 * fi60) != (-(fi56)))) || (0.1 > 0.1)) {
                                                            var v80 float32 = 0.0
                                                            v81 := GenS10{F0: (fi56 ^ 128), F1: 585}
                                                            v82 := GenS11{F0: 187, F1: 255}
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        } else {
                                                            var v83 bool = true
                                                            switch 83 {
                                                            case 2147483647:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            case -9223372036854775808:
                                                                var v84 uint64 = 0
                                                                _ = v84
                                                            case 951:
                                                                if (((3.14159 - float64((0.1 + (float32(float64(0.1)) + (float32(1e10) * ((float32(1e10) / 0.1) * float32((float64(1e10) - (-2.25 + 1e10))))))))) == float64(1e10)) || (v83 || (fi56 != -731))) {
                                                                    if (!v83) {
                                                                        var v85 int16 = (v4 & v4)
                                                                        v4 -= (v4 << ((1) & 63))
                                                                        _ = v85
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 2147483648
                                                                    }
                                                                }
                                                            default:
                                                                if true {
                                                                    v86 := GenS2{F0: 241}
                                                                    var v87 uint32 = 2147483647
                                                                    v4 = (v4 ^ 1)
                                                                    _ = v86
                                                                    _ = v87
                                                                } else {
                                                                    v2 = v2
                                                                }
                                                            }
                                                            var v88 int8 = -92
                                                            v88 -= (v88 | v88)
                                                            _ = v83
                                                            _ = v88
                                                        }
                                                    }
                                                }
                                                fmt.Println(v2, v4, fi60)
                                            }
                                        }
                                        v1 = append(v1, 42)
                                        var v89 float64 = 1e10
                                        _ = v89
                                    }
                                    var v90 string = ("hello" + "hello")
                                    v4 += 702
                                    _ = v90
                                case "":
                                    switch (fi56 & fi56) {
                                    case 32768:
                                        if v2 {
                                            if v2 {
                                                v4 += (v4 & (v4 - (v4 ^ (v4 >> ((765) & 63)))))
                                                v4 = -1
                                                var v91 float64 = 1.5
                                                var v92 uint = 2
                                                fmt.Println(v4)
                                                _ = v91
                                                _ = v92
                                            } else {
                                                v2 = v2
                                                v93 := GenS12{F0: 142, F1: (v4 - v4)}
                                                _ = v93
                                            }
                                        } else {
                                            v4 = v4
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    case -9223372036854775808:
                                        v58 = (fi60 == (fi56 | (^(fi60))))
                                    }
                                default:
                                    if (v4 < v4) {
                                        v1 = append(v1, 4294967295)
                                        switch ((("\n\t" + "") + "\n\t") + "\n\t") {
                                        case "hello":
                                            switch "foo bar" {
                                            case "a":
                                                if v2 {
                                                    v4 = v4
                                                    v94 := GenS4{F0: (fi60 >> ((18446744073709551615) & 63))}
                                                    if false {
                                                        var v95 int32 = 0
                                                        _ = v95
                                                    }
                                                    _ = v94
                                                }
                                            case "foo bar":
                                                var v96 uint16 = 42
                                                _ = v96
                                            }
                                        default:
                                            var v97 int8 = 0
                                            _ = v97
                                        }
                                    } else {
                                        var v98 uint32 = 91
                                        var v99 int8 = -98
                                        _ = v98
                                        _ = v99
                                    }
                                }
                                if v2 {
                                    if (float64(-2.25) > 3.14159) {
                                        v4 -= (v4 & v4)
                                        if false {
                                            switch (fi56 * fi60) {
                                            case 127:
                                                v100 := GenS13{F0: -128, F1: 162}
                                                _ = v100
                                            case 42:
                                                if v2 {
                                                    var v101 int32 = 843
                                                    var v102 string = ("foo bar" + "foo bar")
                                                    _ = v101
                                                    _ = v102
                                                } else {
                                                    var v103 bool = ((!false) || false)
                                                    v104 := GenS14{F0: 100}
                                                    v4 += v4
                                                    _ = v103
                                                    _ = v104
                                                }
                                            case 239:
                                                switch "hello" {
                                                case "hello":
                                                    v4 += v4
                                                default:
                                                    v2 = v58
                                                }
                                            default:
                                                if ("a" <= "foo bar") {
                                                    switch "" {
                                                    case "foo bar":
                                                        if (-336 <= 0) {
                                                            if v58 {
                                                                fmt.Println(fi56)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v4 -= v4
                                                                var v105 uint64 = 929
                                                                if false {
                                                                    if (-1 > 0) {
                                                                        var v106 int32 = -1
                                                                        fmt.Println(v105, v2, v4)
                                                                        var v107 int64 = 1
                                                                        _ = v106
                                                                        _ = v107
                                                                    } else {
                                                                        var v108 int32 = 2147483647
                                                                        v109 := GenS1{F0: (v105 + 18446744073709551615), F1: 65535, F2: 0}
                                                                        var v110 int32 = -1
                                                                        if true {
                                                                            var v111 float32 = ((1e10 * 0.1) + 1.5)
                                                                            v2 = true
                                                                            v108 = (v110 * 42)
                                                                            _ = v111
                                                                        }
                                                                        _ = v108
                                                                        _ = v109
                                                                        _ = v110
                                                                    }
                                                                    if (3.14159 == ((((((((-2.25 / 1.0) * (3.14159 - -2.25)) + (-2.25 + 1.5)) / 3.14159) - 1e10) + float32((3.14159 + ((0.1 * 1e10) - 0.0)))) / float32(0.1)) + float32(0.1))) {
                                                                        v2 = ((^(v105)) <= (v105 / ((v105) | 1)))
                                                                        switch fi56 {
                                                                        case 32768:
                                                                            if v2 {
                                                                                v112 := GenS15{F0: v4, F1: v105, F2: 1}
                                                                                _ = v112
                                                                            }
                                                                        case -8:
                                                                            v113 := GenS8{F0: 0, F1: ("" + "\n\t")}
                                                                            _ = v113
                                                                        }
                                                                        var v114 int64 = 9223372036854775807
                                                                        _ = v114
                                                                    } else {
                                                                        var v115 uint = 255
                                                                        if v2 {
                                                                            switch (("\n\t" + "\n\t") + ("" + (("foo bar" + ("foo bar" + "\n\t")) + "hello"))) {
                                                                            case "foo bar":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            case "":
                                                                                var v116 uint8 = 0
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = "foo bar"
                                                                                }
                                                                                _ = v116
                                                                            case "hello":
                                                                                if (0 != 1) {
                                                                                    v115 += (v115 % (((^(v115))) | 1))
                                                                                    if v2 {
                                                                                        var v117 bool = true
                                                                                        v4 -= (v4 % (((-(v4))) | 1))
                                                                                        _ = v117
                                                                                    } else {
                                                                                        if (fi60 >= 2147483647) {
                                                                                            fmt.Println(fi60, v4)
                                                                                            v118 := GenS15{F0: (v4 | (-(v4))), F1: 1, F2: 0}
                                                                                            _ = v118
                                                                                        } else {
                                                                                            v119 := GenS13{F0: 100, F1: 42}
                                                                                            _ = v119
                                                                                        }
                                                                                        var v120 int64 = 0
                                                                                        _ = v120
                                                                                    }
                                                                                } else {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 65535
                                                                                    }
                                                                                    var v121 uint8 = 235
                                                                                    _ = v121
                                                                                }
                                                                            default:
                                                                                v122 := GenS3{F0: (-1 >= v4), F1: 1}
                                                                                _ = v122
                                                                            }
                                                                            v115 -= v115
                                                                            v3 = append(v3, ("foo bar" + (("hello" + "a") + "")))
                                                                        }
                                                                        _ = v115
                                                                    }
                                                                    switch ("foo bar" + "\n\t") {
                                                                    case "a":
                                                                        v2 = v58
                                                                    case "foo bar":
                                                                        fmt.Println(fi56, v2, v105)
                                                                    case "":
                                                                        var v123 uint16 = 2
                                                                        _ = v123
                                                                    default:
                                                                        v124 := GenS16{F0: "foo bar"}
                                                                        _ = v124
                                                                    }
                                                                }
                                                                _ = v105
                                                            } else {
                                                                var v125 uint8 = 50
                                                                v125 = v125
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v125
                                                            }
                                                            if ((-2147483648 < 0) && v58) {
                                                                v2 = true
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v126 int16 = v4
                                                                    _ = v126
                                                                case "foo bar":
                                                                    if v2 {
                                                                        if ((!(!(v4 != v4))) || ((!false) || false)) {
                                                                            var v127 int16 = int16(fi56)
                                                                            _ = len(v1)
                                                                            if (fi60 <= (-(fi60))) {
                                                                                _ = len(v3)
                                                                                if (!v58) {
                                                                                    v2 = v2
                                                                                    var v128 uint8 = 174
                                                                                    var v129 int = (fi60 & (fi56 / (((fi60 + (fi56 + fi60))) | 1)))
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    var v130 int64 = 9223372036854775807
                                                                                    v58 = true
                                                                                    if (662 == 42) {
                                                                                        var v131 uint32 = 4294967295
                                                                                        _ = v131
                                                                                    }
                                                                                    _ = v130
                                                                                }
                                                                                v2 = ((!true) || v2)
                                                                            }
                                                                            _ = v127
                                                                        }
                                                                        fmt.Println(v58)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    var v132 int64 = -9223372036854775808
                                                                    _ = v132
                                                                case "":
                                                                    v3 = append(v3, "hello")
                                                                    v133 := GenS13{F0: -1, F1: 196}
                                                                    v2 = true
                                                                    _ = v133
                                                                }
                                                            } else {
                                                                if v2 {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        fmt.Println(fi56)
                                                                        var v134 int64 = -9223372036854775808
                                                                        _ = v134
                                                                    case "a":
                                                                        switch (("a" + "a") + "\n\t") {
                                                                        case "foo bar":
                                                                            switch ("foo bar" + "a") {
                                                                            case "foo bar":
                                                                                fmt.Println(v4, v2, fi56)
                                                                            default:
                                                                                v2 = (-1 >= -259)
                                                                            }
                                                                        default:
                                                                            v58 = v2
                                                                        }
                                                                    }
                                                                    v3 = append(v3, "foo bar")
                                                                    switch fi56 {
                                                                    case 32767:
                                                                        if (42 == 783) {
                                                                            var v135 bool = ((591 < 308) || true)
                                                                            var v136 float64 = (-2.25 - 3.14159)
                                                                            _ = v135
                                                                            _ = v136
                                                                        } else {
                                                                            var v137 int32 = 0
                                                                            v138 := GenS17{F0: 2147483647, F1: (fi60 & (fi60 * fi56)), F2: fi60}
                                                                            v58 = v58
                                                                            v139 := GenS16{F0: ""}
                                                                            v58 = (v2 || (65535 <= 2147483647))
                                                                            v140 := GenS13{F0: -75, F1: 211}
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 4294967295
                                                                            }
                                                                            _ = v137
                                                                            _ = v138
                                                                            _ = v139
                                                                            _ = v140
                                                                        }
                                                                    case -9223372036854775808:
                                                                        if v2 {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "\n\t"
                                                                            }
                                                                            if v2 {
                                                                                var v141 float64 = (float64(float32((0.0 * float64((-2.25 - 3.14159))))) / 1.0)
                                                                                _ = v141
                                                                            } else {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v2 = v2
                                                                                }
                                                                                var v142 int64 = -565
                                                                                fmt.Println(fi56)
                                                                                _ = v142
                                                                            }
                                                                            v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
                                                                        } else {
                                                                            var v143 int32 = 2147483647
                                                                            if (!(((-2.25 / 1.5) - 1e10) != 0.0)) {
                                                                                switch (fi60 + (fi56 >> ((65535) & 63))) {
                                                                                case -9223372036854775808:
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "foo bar"
                                                                                    }
                                                                                default:
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "hello"
                                                                                    }
                                                                                }
                                                                                var v144 int64 = 1
                                                                                _ = v144
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            if ((v143 >> ((65535) & 63)) <= v143) {
                                                                                fmt.Println(v58)
                                                                            }
                                                                            _ = v143
                                                                        }
                                                                    }
                                                                    v4 = -335
                                                                }
                                                                v58 = (9223372036854775807 <= -9223372036854775808)
                                                                var v145 uint16 = 42
                                                                _ = v145
                                                            }
                                                            var v146 uint64 = 1
                                                            _ = v146
                                                        } else {
                                                            var v147 uint64 = 15
                                                            var v148 uint = 65535
                                                            _ = v147
                                                            _ = v148
                                                        }
                                                    case "\n\t":
                                                        v149 := GenS1{F0: 2, F1: 94, F2: -1}
                                                        _ = v149
                                                    }
                                                    _ = len(v1)
                                                    v1 = append(v1, 4294967295)
                                                } else {
                                                    v58 = v58
                                                    v150 := GenS18{F0: 1, F1: 575, F2: 127}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v150
                                                }
                                            }
                                            var v151 uint32 = 63
                                            _ = v151
                                        }
                                    } else {
                                        var v152 int16 = v4
                                        _ = v152
                                    }
                                } else {
                                    v2 = (!false)
                                    if (9223372036854775807 == 42) {
                                        var v153 string = ("a" + "\n\t")
                                        var v154 int8 = 1
                                        fmt.Println(v4, fi56, fi56)
                                        var v155 bool = (v58 || (!(false || (542 > 65535))))
                                        v58 = (821 < 143)
                                        _ = v153
                                        _ = v154
                                        _ = v155
                                    }
                                    switch "a" {
                                    case "hello":
                                        v1 = append(v1, 561)
                                    case "foo bar":
                                        if (-1 < -9223372036854775808) {
                                            _ = len(v3)
                                        }
                                    case "a":
                                        v2 = (!(!(!(358 >= 191))))
                                    }
                                    var v156 uint = 4294967295
                                    var v157 uint32 = 65535
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v156 -= 320
                                    _ = v156
                                    _ = v157
                                }
                            }
                        }
                        _ = fi60
                    }
                    _ = v58
                } else {
                    var v158 bool = v2
                    _ = v158
                }
                var v159 int32 = 882
                _ = fi56
                _ = v159
                _ = v57
            }
        }
    }
    v2 = v2
    var v160 uint16 = 65535
    fmt.Println(v4, v2, v160)
    switch "foo bar" {
    case "hello":
        v160 = v160
    case "":
        switch 128 {
        case 32768:
            v160 = v160
        default:
            v161 := GenS15{F0: 0, F1: 18446744073709551615, F2: -1}
            _ = v161
        }
    default:
        v160 = v160
    }
    fmt.Println(v4, v4)
    switch ((("\n\t" + ("hello" + ((("" + ("a" + "a")) + "foo bar") + (("a" + ("\n\t" + "\n\t")) + "hello")))) + "") + "\n\t") {
    case "hello":
        for fi162 := 0; fi162 < 15; fi162++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v2)
            case "a":
                v163 := GenS16{F0: "hello"}
                _ = v163
            case "hello":
                fmt.Println(v2)
            }
            _ = fi162
        }
    }
    fmt.Println(v160, v160)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
