package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 int16; F1 int; F2 string }
type GenS9 struct { F0 bool; F1 uint64; F2 uint }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint; F1 bool }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint64; F1 uint }
type GenS18 struct { F0 int64; F1 uint }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS21 struct { F0 string }

func main() {
    var v1 uint32 = 0
    v2 := []string{(("\n\t" + "a") + ("" + ("\n\t" + "\n\t"))), ("a" + (("hello" + ("\n\t" + ("a" + ((("foo bar" + "") + "hello") + ("a" + "foo bar"))))) + (("\n\t" + ("foo bar" + (("hello" + (("a" + "foo bar") + "\n\t")) + "hello"))) + ("\n\t" + "foo bar")))), ""}
    var v3 uint16 = 789
    var v4 int32 = -1
    v5 := []string{"", ("hello" + "a"), ("\n\t" + ("" + (("hello" + "foo bar") + "hello"))), "hello"}
    var v6 int16 = -1
    if (847 != (v1 - (v1 + v1))) {
        var v7 string = "foo bar"
        _ = v7
    } else {
        fmt.Println(v6)
        for fi8 := 0; fi8 < 10; fi8++ {
            if ((1.5 + -2.25) < 1e10) {
                for fi9 := 0; fi9 < 5; fi9++ {
                    switch "\n\t" {
                    case "":
                        if (!((!(true || (429 <= (-(v4))))) && true)) {
                            v3 = (v3 % (((^(v3))) | 1))
                            if false {
                                if true {
                                    if false {
                                        var v10 int = 127
                                        v11 := GenS1{F0: "foo bar", F1: (v3 >> ((817) & 63)), F2: 1}
                                        v5 = append(v5, "\n\t")
                                        v12 := GenS2{F0: v4, F1: 65535, F2: -130}
                                        var v13 bool = (-1 <= 42)
                                        var v14 int16 = (v6 | v6)
                                        var v15 int64 = -970
                                        _ = v10
                                        _ = v11
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                    }
                                    fmt.Println(fi8, v3, fi8)
                                }
                            }
                        }
                    case "foo bar":
                        var v16 float32 = 1.5
                        _ = v16
                    default:
                        if true {
                            v3 = 0
                        }
                        if ("hello" <= "hello") {
                            _ = len(v5)
                        } else {
                            switch (fi9 ^ (^(fi9))) {
                            case 32768:
                                v6 -= v6
                            case 9223372036854775807:
                                v1 += (v1 - (v1 % (((v1 | v1)) | 1)))
                            case 484:
                                var v17 uint32 = 65535
                                _ = v17
                            default:
                                v1 = (v1 / ((v1) | 1))
                            }
                            var v18 uint32 = v1
                            _ = v18
                        }
                    }
                    var v19 uint64 = 2
                    _ = fi9
                    _ = v19
                }
                for fi20 := 0; fi20 < 20; fi20++ {
                    v21 := GenS3{F0: (int32(v1) < 42), F1: 128}
                    var v22 int16 = v6
                    v22 -= v6
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                if false {
                    v1 = v1
                    if true {
                        v23 := GenS2{F0: v4, F1: 501, F2: v6}
                        if true {
                            v24 := GenS4{F0: (v6 ^ v6), F1: 4294967295, F2: int8(v6)}
                            _ = v24
                        } else {
                            v1 += 4294967295
                        }
                        for fi25 := 0; fi25 < 18; fi25++ {
                            fmt.Println(v3, v4)
                            var v26 uint8 = 252
                            var v27 string = "hello"
                            switch (fi8 * 32767) {
                            case 127:
                                switch -1 {
                                case -2147483648:
                                    var v28 uint = 65535
                                    _ = v28
                                case 9223372036854775807:
                                    v4 = v4
                                default:
                                    var v29 int = (fi8 % ((fi8) | 1))
                                    _ = v29
                                }
                            case -9223372036854775808:
                                var v30 uint = 692
                                _ = v30
                            case 9223372036854775807:
                                var v31 bool = (v6 <= int16(v4))
                                _ = v31
                            default:
                                v32 := GenS5{F0: fi25}
                                if ((v27 > "a") || false) {
                                    v33 := GenS6{F0: (v1 ^ (v1 << ((255) & 63)))}
                                    if len(v2) > 0 {
                                        v2[0] = (v27 + v27)
                                    }
                                    v27 = "hello"
                                    _ = v33
                                }
                                _ = v32
                            }
                            var v34 float64 = (3.14159 * 1.5)
                            v4 += v4
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v34
                        }
                        _ = v23
                    } else {
                        if (!true) {
                            for fi35 := 0; fi35 < 6; fi35++ {
                                switch "a" {
                                case "hello":
                                    v4 += v4
                                case "a":
                                    if true {
                                        var v36 int8 = -120
                                        _ = v36
                                    }
                                }
                                v1 = v1
                                _ = fi35
                            }
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + ("foo bar" + ("hello" + "hello")))
                            }
                        } else {
                            if false {
                                v4 = v4
                            } else {
                                v37 := GenS7{F0: (0.1 >= 0.0), F1: (!true)}
                                switch ("a" + "hello") {
                                case "a":
                                    v2 = append(v2, ("foo bar" + "hello"))
                                default:
                                    if false {
                                        for fi38 := 0; fi38 < 20; fi38++ {
                                            var v39 uint8 = 0
                                            if (v39 < 2) {
                                                var v40 uint32 = (v1 - 42)
                                                var v41 uint8 = 127
                                                if true {
                                                    v4 = (v4 << ((18446744073709551615) & 63))
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        switch fi8 {
                                                        case 510:
                                                            var v42 float32 = 1.5
                                                            v2 = append(v2, ("hello" + "\n\t"))
                                                            _ = v42
                                                        case 1:
                                                            v2 = append(v2, "hello")
                                                        case -1:
                                                            fmt.Println(v6, fi8, v3)
                                                        }
                                                    default:
                                                        v4 = -860
                                                    }
                                                    v4 = (v4 * (^(v4)))
                                                }
                                                _ = v40
                                                _ = v41
                                            } else {
                                                if false {
                                                    if true {
                                                        if ((v6 ^ v6) != (v6 / (((v6 >> ((1) & 63))) | 1))) {
                                                            v5 = append(v5, ("a" + "a"))
                                                        }
                                                        var v43 int32 = 2147483647
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v43
                                                    } else {
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(v39, v39)
                                                        default:
                                                            v44 := GenS7{F0: (v4 != (-(v4))), F1: ((-2.25 / 3.14159) != 0.1)}
                                                            _ = v44
                                                        }
                                                    }
                                                    v2 = append(v2, "foo bar")
                                                    switch "a" {
                                                    case "\n\t":
                                                        if false {
                                                            var v45 string = "\n\t"
                                                            _ = v45
                                                        } else {
                                                            var v46 bool = (!(false && true))
                                                            v47 := GenS7{F0: ((!(-9223372036854775808 <= 0)) || v46), F1: ((v1 * uint32(v3)) != (v1 << ((2) & 63)))}
                                                            v48 := GenS5{F0: (fi8 * fi38)}
                                                            fmt.Println(v1, fi38, fi8)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        var v49 string = "\n\t"
                                                        _ = v49
                                                    } else {
                                                        if (false || false) {
                                                            if (!false) {
                                                                v6 += (v6 & v6)
                                                            }
                                                            var v50 uint16 = v3
                                                            _ = v50
                                                        } else {
                                                            switch fi8 {
                                                            case -113:
                                                                _ = len(v5)
                                                            case 42:
                                                                v51 := GenS3{F0: false, F1: (v39 - (^(v39)))}
                                                                _ = v51
                                                            case -509:
                                                                fmt.Println(v3)
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        v1 = (v1 << ((370) & 63))
                                                        var v52 uint = 2
                                                        _ = v52
                                                    }
                                                    var v53 string = "a"
                                                    _ = v53
                                                }
                                                v39 = (v39 & 42)
                                            }
                                            _ = fi38
                                            _ = v39
                                        }
                                        v3 -= (v3 ^ (v3 | v3))
                                    } else {
                                        for fi54 := 0; fi54 < 11; fi54++ {
                                            v3 -= 2
                                            _ = fi54
                                        }
                                    }
                                }
                                _ = v37
                            }
                            fmt.Println(v3, v4)
                            switch ("" + (("hello" + "") + "\n\t")) {
                            case "hello":
                                if false {
                                    v1 += v1
                                }
                            case "a":
                                v55 := GenS8{F0: v6, F1: (fi8 >> ((255) & 63)), F2: ("" + "a")}
                                _ = v55
                            default:
                                v2 = append(v2, "")
                            }
                        }
                        if (502 > 343) {
                            v4 += -288
                            var v56 int64 = -9223372036854775808
                            v57 := GenS9{F0: (((810 > 1) && (true && false)) && (755 >= 42)), F1: 18446744073709551615, F2: 453}
                            _ = v56
                            _ = v57
                        } else {
                            for fi58 := 0; fi58 < 7; fi58++ {
                                var v59 uint16 = v3
                                _ = fi58
                                _ = v59
                            }
                        }
                        var v60 int = (fi8 & fi8)
                        _ = v60
                    }
                } else {
                    v61 := GenS5{F0: fi8}
                    _ = v61
                }
                if (fi8 <= fi8) {
                    var v62 uint = 752
                    if (((42 == (v3 | (^(v3)))) && false) || false) {
                        var v63 bool = false
                        v64 := GenS10{F0: v1}
                        _ = v63
                        _ = v64
                    } else {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            if false {
                                fmt.Println(v3, v1)
                            }
                        case "hello":
                            if ((fi8 != -9223372036854775808) || (false && (int(v3) >= int(v62)))) {
                                fmt.Println(fi8, v1)
                                v6 -= v6
                            } else {
                                v4 = 42
                            }
                        case "foo bar":
                            v65 := GenS5{F0: fi8}
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 8; fi66++ {
                            if true {
                                v67 := GenS1{F0: ("a" + "a"), F1: v3, F2: 2}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if false {
                                    v1 = v1
                                }
                                _ = v67
                            } else {
                                v6 += v6
                                if (10 <= -31) {
                                    v1 = v1
                                    v4 = v4
                                } else {
                                    if (!(1 == 4294967295)) {
                                        v5 = append(v5, "")
                                    }
                                    v3 = v3
                                    var v68 uint = v62
                                    _ = v68
                                }
                                fmt.Println(fi66, v1, v4)
                            }
                            var v69 float32 = (((1.5 * (0.0 + -2.25)) / (float32(0.1) * -2.25)) + 1.5)
                            _ = fi66
                            _ = v69
                        }
                    }
                    switch "" {
                    case "\n\t":
                        for fi70 := 0; fi70 < 13; fi70++ {
                            var v71 float32 = (float32(1.5) * 1.5)
                            _ = len(v2)
                            _ = len(v2)
                            _ = fi70
                            _ = v71
                        }
                    case "hello":
                        v4 -= -2147483648
                    default:
                        for fi72 := 0; fi72 < 15; fi72++ {
                            v2 = append(v2, "a")
                            var v73 int = fi72
                            v74 := GenS11{F0: 731, F1: (v3 >> ((v62) & 63)), F2: 37}
                            _ = fi72
                            _ = v73
                            _ = v74
                        }
                    }
                    _ = v62
                }
            }
            switch "hello" {
            case "hello":
                fmt.Println(v4, v3, fi8)
            case "":
                if (2 == v3) {
                    var v75 bool = false
                    _ = v75
                } else {
                    switch "" {
                    case "\n\t":
                        v4 = (v4 >> ((468) & 63))
                    default:
                        v4 = v4
                    }
                    if false {
                        for fi76 := 0; fi76 < 7; fi76++ {
                            if (true || (-9223372036854775808 >= 42)) {
                                var v77 int8 = -128
                                v78 := GenS12{F0: -590}
                                _ = v77
                                _ = v78
                            }
                            _ = fi76
                        }
                        var v79 uint64 = 473
                        _ = v79
                    } else {
                        var v80 uint8 = 0
                        v80 = (v80 & (v80 * 2))
                        var v81 int64 = -9223372036854775808
                        _ = v80
                        _ = v81
                    }
                    var v82 float64 = 1e10
                    _ = v82
                }
            }
            switch 9223372036854775807 {
            case -2147483648:
                for fi83 := 0; fi83 < 18; fi83++ {
                    var v84 uint64 = 0
                    var v85 uint8 = 42
                    _ = fi83
                    _ = v84
                    _ = v85
                }
            }
            _ = fi8
        }
        for fi86 := 0; fi86 < 1; fi86++ {
            v6 += (-(v6))
            v1 -= 2147483648
            v2 = append(v2, "a")
            _ = fi86
        }
    }
    v3 -= v3
    v87 := GenS13{F0: 27}
    for fi88 := 0; fi88 < 11; fi88++ {
        if (((1.5 + (1e10 * (-2.25 / float32(1.5)))) - 0.1) >= (3.14159 * (3.14159 + 3.14159))) {
            fmt.Println(v4, v3)
            var v89 int32 = -2147483648
            v90 := GenS2{F0: (v4 % ((v89) | 1)), F1: 975, F2: (v6 + (v6 % ((v6) | 1)))}
            _ = v89
            _ = v90
        }
        v91 := GenS1{F0: "hello", F1: (v3 | 32767), F2: 1}
        _ = fi88
        _ = v91
    }
    var v92 int16 = v6
    v93 := GenS14{F0: v4}
    v92 = v92
    if true {
        v6 = (v92 + v92)
    }
    if false {
        for fi94 := 0; fi94 < 14; fi94++ {
            for fi95 := 0; fi95 < 1; fi95++ {
                fmt.Println(fi95)
                fmt.Println(v6, v1, fi94)
                _ = fi95
            }
            v5 = append(v5, ((((("\n\t" + "\n\t") + "") + "foo bar") + "a") + "\n\t"))
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = fi94
        }
    }
    if (3.14159 < (float64(0.0) / (3.14159 / 1.5))) {
        if true {
            for fi96 := 0; fi96 < 2; fi96++ {
                if true {
                    v1 += (v1 ^ 42)
                    switch "" {
                    case "":
                        if false {
                            switch "\n\t" {
                            case "a":
                                if true {
                                    v3 -= 42
                                }
                            case "":
                                var v97 bool = ("foo bar" < "foo bar")
                                _ = v97
                            case "hello":
                                v98 := GenS15{F0: 689, F1: (32767 != 65535)}
                                _ = v98
                            default:
                                v99 := GenS16{F0: ""}
                                _ = v99
                            }
                        }
                    }
                } else {
                    if true {
                        var v100 uint64 = 162
                        _ = v100
                    } else {
                        v3 = v3
                    }
                }
                for fi101 := 0; fi101 < 6; fi101++ {
                    v92 = (^(v6))
                    switch ("foo bar" + (("\n\t" + "hello") + (("hello" + "hello") + "a"))) {
                    case "\n\t":
                        var v102 int = fi101
                        _ = v102
                    }
                    v6 += (v92 >> ((591) & 63))
                    var v103 uint8 = 101
                    v104 := GenS6{F0: 65535}
                    v2 = append(v2, ("a" + "foo bar"))
                    _ = fi101
                    _ = v103
                    _ = v104
                }
                _ = fi96
            }
        } else {
            if false {
                v105 := GenS17{F0: 685, F1: 279}
                v3 = (v3 & 90)
                var v106 int32 = (v4 | (v4 / ((-592) | 1)))
                _ = v105
                _ = v106
            }
            v1 = (v1 | uint32(v3))
        }
        if (true && (!true)) {
            v107 := GenS17{F0: 4294967295, F1: 255}
            v108 := GenS5{F0: 993}
            _ = v107
            _ = v108
        } else {
            v109 := GenS8{F0: v92, F1: 1, F2: "foo bar"}
            v1 = v1
            v110 := GenS5{F0: 166}
            _ = v109
            _ = v110
        }
        switch "a" {
        case "\n\t":
            var v111 int16 = 42
            _ = v111
        case "":
            switch ("\n\t" + ("" + "a")) {
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    v112 := GenS9{F0: false, F1: 1, F2: 65535}
                    _ = v112
                case "":
                    for fi113 := 0; fi113 < 12; fi113++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        for fi114 := 0; fi114 < 6; fi114++ {
                            if false {
                                var v115 float64 = 1e10
                                _ = v115
                            } else {
                                if ((((0.1 * (0.1 + ((1e10 - 1e10) + ((1.5 + 1.5) * (1.5 / 3.14159))))) >= -2.25) && false) || true) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v116 := GenS18{F0: 0, F1: 255}
                                        _ = v116
                                    case "":
                                        v5 = append(v5, "\n\t")
                                    case "foo bar":
                                        _ = len(v5)
                                    default:
                                        fmt.Println(v92, v6)
                                    }
                                    var v117 string = "a"
                                    var v118 bool = (((("hello" > "a") && true) || true) && (!false))
                                    v119 := GenS19{F0: fi113}
                                    var v120 int = 32768
                                    var v121 uint16 = v3
                                    _ = v117
                                    _ = v118
                                    _ = v119
                                    _ = v120
                                    _ = v121
                                } else {
                                    if false {
                                        var v122 uint64 = 567
                                        _ = v122
                                    } else {
                                        var v123 uint32 = (^(v1))
                                        _ = v123
                                    }
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        if true {
                                            if false {
                                                v3 = 592
                                            }
                                            fmt.Println(v3, fi114)
                                        }
                                    case "hello":
                                        v124 := GenS18{F0: -9223372036854775808, F1: 18446744073709551615}
                                        _ = v124
                                    case "":
                                        fmt.Println(v3)
                                    }
                                    v125 := GenS20{F0: 4294967295, F1: 151, F2: -1}
                                    if len(v2) > 0 {
                                        v2[0] = ("" + (("a" + "hello") + "hello"))
                                    }
                                    _ = v125
                                }
                                if true {
                                    if true {
                                        switch (-(fi113)) {
                                        case 408:
                                            v126 := GenS13{F0: 2}
                                            _ = v126
                                        }
                                        v127 := GenS16{F0: "a"}
                                        var v128 bool = ((v3 == 1) && false)
                                        _ = v127
                                        _ = v128
                                    }
                                    v129 := GenS20{F0: 18446744073709551615, F1: 42, F2: 1}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v129
                                } else {
                                    var v130 int16 = 450
                                    var v131 uint16 = 381
                                    if (int(v131) > -185) {
                                        fmt.Println(v131)
                                    }
                                    v3 = v131
                                    var v132 float32 = 1e10
                                    var v133 float32 = v132
                                    fmt.Println(v132, v3, fi114)
                                    _ = v130
                                    _ = v131
                                    _ = v132
                                    _ = v133
                                }
                                var v134 uint64 = 18446744073709551615
                                _ = v134
                            }
                            var v135 uint32 = (v1 << ((184) & 63))
                            _ = len(v2)
                            _ = fi114
                            _ = v135
                        }
                        _ = fi113
                    }
                case "a":
                    var v136 uint64 = 623
                    _ = v136
                }
            case "a":
                for fi137 := 0; fi137 < 18; fi137++ {
                    var v138 int64 = 9223372036854775807
                    switch (fi137 - int(v3)) {
                    case 420:
                        switch "foo bar" {
                        case "hello":
                            var v139 int64 = (-(v138))
                            _ = v139
                        case "":
                            v1 = v1
                        case "foo bar":
                            v140 := GenS21{F0: ("hello" + "")}
                            _ = v140
                        default:
                            _ = len(v2)
                        }
                    }
                    v138 += v138
                    _ = fi137
                    _ = v138
                }
            case "\n\t":
                v92 = int16(v4)
            }
        case "foo bar":
            var v141 int8 = -34
            _ = v141
        }
    } else {
        var v142 string = ""
        v5 = append(v5, v142)
        switch int(v1) {
        case 0:
            var v143 uint8 = 255
            _ = v143
        case -1:
            v3 += (v3 / ((v3) | 1))
        case 276:
            v142 = v142
        default:
            v144 := GenS1{F0: (v142 + (((v142 + v142) + (v142 + ((v142 + "hello") + (v142 + "hello")))) + v142)), F1: v3, F2: 0}
            _ = v144
        }
        _ = v142
    }
    for fi145 := 0; fi145 < 20; fi145++ {
        fmt.Println(v1)
        v146 := GenS7{F0: (!(v1 != v1)), F1: true}
        _ = fi145
        _ = v146
    }
    var v147 uint32 = (v1 / (((^(v1))) | 1))
    if (42 == 981) {
        if true {
            fmt.Println(v4, v3, v4)
            switch ("a" + "foo bar") {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = ""
                }
            case "\n\t":
                v148 := GenS19{F0: 4294967295}
                _ = v148
            case "hello":
                var v149 string = ("a" + "a")
                _ = v149
            }
        }
        var v150 uint64 = 42
        v150 = 1
        _ = v150
    } else {
        var v151 string = ("a" + (("hello" + ("hello" + "a")) + ("\n\t" + ("" + "foo bar"))))
        switch 128 {
        case -9223372036854775808:
            for fi152 := 0; fi152 < 19; fi152++ {
                v6 -= (v6 << ((1) & 63))
                for fi153 := 0; fi153 < 10; fi153++ {
                    if true {
                        v154 := GenS14{F0: (v4 >> ((770) & 63))}
                        v5 = append(v5, "a")
                        _ = v154
                    }
                    var v155 int8 = -1
                    fmt.Println(v1, v6)
                    _ = fi153
                    _ = v155
                }
                _ = fi152
            }
        }
        if ((v4 / (((v4 ^ (v4 / ((-189) | 1)))) | 1)) < v4) {
            fmt.Println(v1, v1, v92)
            for fi156 := 0; fi156 < 5; fi156++ {
                if ((1.5 * (1e10 + 0.0)) <= 0.1) {
                    v151 = v151
                    var v157 bool = true
                    _ = v157
                }
                var v158 uint16 = v3
                _ = fi156
                _ = v158
            }
        }
        _ = v151
    }
    var v159 int32 = v4
    for fi160 := 0; fi160 < 12; fi160++ {
        fmt.Println(v92, v4)
        _ = fi160
    }
    v147 = (v147 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v147:", v147)
    fmt.Println("v159:", v159)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v147
    _ = v159
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v87
    _ = v92
    _ = v93
}
