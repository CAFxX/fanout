package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int64; F1 string; F2 uint }
type GenS6 struct { F0 bool; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int; F1 bool; F2 bool }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint; F1 int16 }
type GenS13 struct { F0 bool; F1 uint; F2 uint }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 string; F1 int16 }
type GenS17 struct { F0 string; F1 uint32; F2 int16 }
type GenS18 struct { F0 int64; F1 bool; F2 string }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int; F1 bool; F2 string }
type GenS21 struct { F0 uint64; F1 int32 }
type GenS22 struct { F0 int8; F1 string; F2 uint64 }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 int32; F1 bool }
type GenS25 struct { F0 string; F1 uint16 }
type GenS26 struct { F0 int; F1 int }
type GenS27 struct { F0 int8 }
type GenS28 struct { F0 int64; F1 int32; F2 int16 }
type GenS29 struct { F0 int16; F1 int32; F2 uint8 }
type GenS30 struct { F0 uint16 }
type GenS31 struct { F0 int; F1 uint32; F2 int16 }
type GenS32 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS33 struct { F0 int }
type GenS34 struct { F0 uint64; F1 uint; F2 int }
type GenS35 struct { F0 int; F1 int64 }
type GenS36 struct { F0 uint; F1 uint32 }
type GenS37 struct { F0 uint; F1 uint64 }
type GenS38 struct { F0 uint32; F1 uint32 }
type GenS39 struct { F0 uint16; F1 int32; F2 int64 }
type GenS40 struct { F0 int8 }
type GenS41 struct { F0 uint8 }
type GenS42 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint16 = 1
    var v2 float64 = 0.1
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 int64 = 642
    v1 = (v1 - 777)
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int32 = -882
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v2, fi10)
        if (32767 <= 42) {
            var v11 uint16 = v1
            fmt.Println(v6, v6)
            var v12 int = (fi10 << ((206) & 63))
            _ = v11
            _ = v12
        } else {
            var v13 int8 = 0
            _ = v13
        }
        _ = fi10
    }
    v9 = (v9 % (((v9 / (((v9 ^ (^(v9)))) | 1))) | 1))
    for fi14 := 0; fi14 < 11; fi14++ {
        if (v9 != (v9 << ((65535) & 63))) {
            switch (-(fi14)) {
            case -558:
                if false {
                    var v15 int64 = 9223372036854775807
                    if ((1.5 + float64((float32(((v2 / v2) * 1.5)) / float32(1.5)))) <= 1e10) {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v1 = (v1 - v1)
                            _ = fi16
                        }
                        delete(v5, (fi14 ^ fi14))
                    }
                    v17 := []int{(fi14 + fi14), (^(fi14)), 0, fi14}
                    _ = v15
                    _ = v17
                } else {
                    if (-114 <= -53) {
                        var v18 float32 = ((float32(0.0) / 3.14159) + 1.5)
                        var v19 uint64 = (v9 + 0)
                        _ = v18
                        _ = v19
                    }
                    v4["foo bar"] = (float64(3.14159) != v2)
                }
            }
            var v20 int64 = (v6 - v6)
            v21 := GenS1{F0: "", F1: 370, F2: 1}
            _ = v20
            _ = v21
        }
        v2 = ((v2 + v2) * (v2 * -2.25))
        _ = fi14
    }
    var v22 int32 = -2147483648
    switch int(v9) {
    case 32767:
        v22 += v22
    case -905:
        delete(v5, 2147483647)
    }
    var v23 float32 = (float32(v2) + float32(v2))
    v1 += v1
    if (v22 != 21) {
        v22 += 0
        switch "a" {
        case "hello":
            switch ("\n\t" + "foo bar") {
            case "":
                switch "foo bar" {
                case "\n\t":
                    if ((v22 > v22) || (-1 >= -11)) {
                        if (!false) {
                            switch 4294967295 {
                            case 520:
                                var v24 uint16 = v1
                                _ = v24
                            case -2147483648:
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v26 uint16 = (v1 / (((v1 << ((255) & 63))) | 1))
                                        _ = v26
                                    case "a":
                                        var v27 uint = 0
                                        _ = v27
                                    }
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        fmt.Println(v2, v1, v22)
                                        var v29 uint32 = 65535
                                        v6 += int64(v29)
                                        _ = fi28
                                        _ = v29
                                    }
                                    v2 -= v2
                                    _ = fi25
                                }
                            case 0:
                                v30 := []int{9223372036854775807, -121}
                                _ = v30
                            }
                            fmt.Println(v1, v6)
                        }
                        if (1 == (v22 % (((v22 / ((v22) | 1))) | 1))) {
                            v1 -= (v1 ^ (v1 - v1))
                            var v31 uint32 = 4294967295
                            _ = v31
                        } else {
                            var v32 uint32 = 0
                            var v33 string = ("hello" + ("a" + ""))
                            fmt.Println(v1)
                            _ = v32
                            _ = v33
                        }
                        if (!((v1 - uint16(v9)) <= v1)) {
                            if (!(!true)) {
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    v23 -= 0.1
                                    _ = fi34
                                }
                                v35 := GenS2{F0: 255, F1: ((float64(v23) != v2) || false)}
                                v9 -= (v9 * 409)
                                _ = v35
                            } else {
                                v1 -= v1
                                var v36 int = -1
                                _ = v36
                            }
                        } else {
                            if true {
                                switch 2147483647 {
                                case 1:
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        fmt.Println(v9, v22)
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v38 int32 = (v22 >> ((255) & 63))
                                            _ = v38
                                        case "hello":
                                            v6 -= v6
                                        default:
                                            v39 := []int{}
                                            _ = v39
                                        }
                                        v22 -= -839
                                        _ = fi37
                                    }
                                default:
                                    fmt.Println(v2, v9, v9)
                                }
                                fmt.Println(v1, v22, v1)
                            }
                        }
                    }
                case "foo bar":
                    if ((v23 * (float32(float64(v23)) - 0.1)) <= (-2.25 * v23)) {
                        switch 42 {
                        case 603:
                            var v40 string = (("foo bar" + "a") + "foo bar")
                            _ = v40
                        case 42:
                            fmt.Println(v2)
                        case -900:
                            v41 := []int{32768, -2147483648}
                            _ = v41
                        }
                        var v42 int64 = -615
                        var v43 int32 = 40
                        _ = v42
                        _ = v43
                    }
                case "hello":
                    for fi44 := 0; fi44 < 5; fi44++ {
                        v6 = v6
                        _ = fi44
                    }
                }
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    var v45 uint32 = 2
                    _ = v45
                }
            case "foo bar":
                if (-1 < int16(v6)) {
                    v6 = (v6 / ((-9223372036854775808) | 1))
                    var v46 uint8 = 127
                    _ = v46
                }
            }
        case "foo bar":
            if true {
                var v47 string = ""
                if true {
                    var v48 int32 = v22
                    for fi49 := 0; fi49 < 11; fi49++ {
                        v50 := GenS2{F0: 0, F1: true}
                        if true {
                            v6 += (v6 >> ((42) & 63))
                        }
                        fmt.Println(v9, v9)
                        _ = fi49
                        _ = v50
                    }
                    if false {
                        for fi51 := 0; fi51 < 5; fi51++ {
                            _ = len(v5)
                            _ = fi51
                        }
                    }
                    _ = v48
                } else {
                    if (!(true && false)) {
                        if false {
                            if (true || ((!true) && true)) {
                                var v52 uint8 = 127
                                v22 += v22
                                var v53 uint16 = v1
                                _ = v52
                                _ = v53
                            } else {
                                if ((1 == (v22 << ((255) & 63))) && false) {
                                    var v54 int16 = 641
                                    if ((v6 & 409) > (v6 | int64(v54))) {
                                        v1 = (^(v1))
                                        if true {
                                            for fi55 := 0; fi55 < 12; fi55++ {
                                                switch v47 {
                                                case "":
                                                    v56 := GenS3{F0: ("\n\t" + v47)}
                                                    _ = v56
                                                }
                                                _ = fi55
                                            }
                                        } else {
                                            delete(v5, -897)
                                            v57 := []int{0, 2147483647}
                                            _ = v57
                                        }
                                        v58 := []int{-9223372036854775808, 4294967295, -776, -2147483648}
                                        _ = v58
                                    } else {
                                        switch v47 {
                                        case "":
                                            for fi59 := 0; fi59 < 19; fi59++ {
                                                delete(v3, (fi59 - (^(fi59))))
                                                switch fi59 {
                                                case 2147483647:
                                                    if true {
                                                        for fi60 := 0; fi60 < 10; fi60++ {
                                                            var v61 uint16 = (^(v1))
                                                            _ = fi60
                                                            _ = v61
                                                        }
                                                    } else {
                                                        if (!(((v54 >> ((2) & 63)) != v54) && (18446744073709551615 != 1))) {
                                                            var v62 bool = true
                                                            for fi63 := 0; fi63 < 13; fi63++ {
                                                                if ((v1 + (^(v1))) == (^(v1))) {
                                                                    if true {
                                                                        var v64 float32 = v23
                                                                        delete(v5, (fi63 & (fi59 % ((fi63) | 1))))
                                                                        _ = v64
                                                                    }
                                                                    fmt.Println(v22, v1)
                                                                } else {
                                                                    var v65 int16 = (-(v54))
                                                                    if (!(true || (((v6 ^ int64(v22)) > (v6 % ((812) | 1))) || (!false)))) {
                                                                        if v62 {
                                                                            var v66 int32 = v22
                                                                            _ = v66
                                                                        }
                                                                        v62 = (19 >= 100)
                                                                        v67 := GenS1{F0: v47, F1: 1000, F2: -1}
                                                                        v1 += v1
                                                                        _ = v67
                                                                    }
                                                                    _ = v65
                                                                }
                                                                if (674 > 4294967295) {
                                                                    var v68 int64 = (v6 + 77)
                                                                    switch ("hello" + (((v47 + v47) + "") + "hello")) {
                                                                    case "":
                                                                        if v62 {
                                                                            v69 := v4[v47]
                                                                            v54 = (v54 >> ((961) & 63))
                                                                            v6 -= -9223372036854775808
                                                                            _ = v69
                                                                        } else {
                                                                            v2 = ((v2 + float64(1e10)) / (0.0 + (float64(0.1) + 1.5)))
                                                                            v68 = (^(v6))
                                                                        }
                                                                    case "foo bar":
                                                                        var v70 int16 = (v54 - (v54 % (((v54 * (v54 + v54))) | 1)))
                                                                        v71 := []int{}
                                                                        if (((573 > 255) && (false || (true || v62))) && v62) {
                                                                            var v72 int = (fi63 / ((0) | 1))
                                                                            v3[int(v9)] = v47
                                                                            _ = v72
                                                                        }
                                                                        fmt.Println(fi59, v68)
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                    if v62 {
                                                                        v68 += v68
                                                                        var v73 uint = 4294967295
                                                                        _ = v73
                                                                    }
                                                                    _ = v68
                                                                } else {
                                                                    if v62 {
                                                                        var v74 uint64 = v9
                                                                        _ = v74
                                                                    }
                                                                    var v75 uint = 4294967295
                                                                    v76 := GenS1{F0: "", F1: (v54 ^ 306), F2: -76}
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                                v6 += 42
                                                                _ = fi63
                                                            }
                                                            fmt.Println(v23, v1, v23)
                                                            _ = v62
                                                        }
                                                        v23 -= v23
                                                    }
                                                case 1:
                                                    if (-8 != -128) {
                                                        v6 += int64(v9)
                                                    } else {
                                                        for fi77 := 0; fi77 < 10; fi77++ {
                                                            v78 := GenS4{F0: 0, F1: 1}
                                                            _ = len(v5)
                                                            delete(v4, "hello")
                                                            if false {
                                                                if (v1 == (v1 & 42)) {
                                                                    fmt.Println(v47, v47)
                                                                    v23 = 0.1
                                                                    v79 := GenS2{F0: 255, F1: false}
                                                                    _ = v79
                                                                }
                                                            } else {
                                                                var v80 uint16 = v1
                                                                _ = v80
                                                            }
                                                            if ((false && (((v23 + v23) * v23) > v23)) && (42 > 230)) {
                                                                switch fi59 {
                                                                case 128:
                                                                    if true {
                                                                        var v81 float64 = float64(v23)
                                                                        switch (-(fi59)) {
                                                                        case 0:
                                                                            fmt.Println(v54, v9)
                                                                        case 1:
                                                                            v82 := GenS5{F0: (v6 ^ 0), F1: "foo bar", F2: 883}
                                                                            _ = v82
                                                                        case -191:
                                                                            v83 := GenS6{F0: ((v1 * v1) != v1), F1: 42, F2: (v22 | v22)}
                                                                            _ = v83
                                                                        default:
                                                                            if false {
                                                                                switch -581 {
                                                                                case 32768:
                                                                                    if true {
                                                                                        switch 641 {
                                                                                        case 127:
                                                                                            var v84 uint16 = (v1 + (v1 / ((v1) | 1)))
                                                                                            _ = v84
                                                                                        case -2147483648:
                                                                                            v85 := GenS3{F0: v47}
                                                                                            _ = v85
                                                                                        case 4294967295:
                                                                                            v9 = (v9 & 507)
                                                                                            v86 := []int{fi77, int(v9), fi59}
                                                                                            var v87 string = (v47 + v47)
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        default:
                                                                                            v88 := GenS6{F0: (!true), F1: 0, F2: 2147483647}
                                                                                            _ = v88
                                                                                        }
                                                                                    } else {
                                                                                        var v89 uint32 = 65535
                                                                                        _ = v89
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch "" {
                                                                                case "hello":
                                                                                    var v90 string = "hello"
                                                                                    var v91 uint = 444
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                default:
                                                                                    v92 := GenS7{F0: 2, F1: (v6 | (v6 * v6))}
                                                                                    switch (v47 + "") {
                                                                                    case "hello":
                                                                                        switch v47 {
                                                                                        case "":
                                                                                            v93 := GenS8{F0: fi77, F1: (18446744073709551615 < 317), F2: ((v23 - ((float32(float64(v23)) + (v23 * 1.5)) / (1.5 / v23))) == v23)}
                                                                                            _ = v93
                                                                                        }
                                                                                    case "":
                                                                                        fmt.Println(v23, v2)
                                                                                    case "\n\t":
                                                                                        var v94 int32 = (v22 + (-(v22)))
                                                                                        if false {
                                                                                            if ((v94 - (-(v22))) >= v94) {
                                                                                                var v95 int8 = -102
                                                                                                var v96 uint32 = 42
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                            }
                                                                                        } else {
                                                                                            v97 := GenS1{F0: v47, F1: v54, F2: 1}
                                                                                            _ = len(v5)
                                                                                            _ = v97
                                                                                        }
                                                                                        if (true && (!((255 < 128) || (!(v1 < 670))))) {
                                                                                            v98 := v5[5]
                                                                                            var v99 float32 = (v23 - (v23 + 0.1))
                                                                                            switch v47 {
                                                                                            case "a":
                                                                                                var v100 string = v47
                                                                                                _ = v100
                                                                                            case "foo bar":
                                                                                                if v98 {
                                                                                                    switch ((v47 + v47) + (v47 + "a")) {
                                                                                                    case "\n\t":
                                                                                                        var v101 uint32 = 2147483648
                                                                                                        _ = v101
                                                                                                    default:
                                                                                                        v6 = 9223372036854775807
                                                                                                    }
                                                                                                    if (!(!v98)) {
                                                                                                        v102 := GenS2{F0: 127, F1: (0 != 0)}
                                                                                                        if ((!v98) && v98) {
                                                                                                            if ((0.1 - v81) == (v81 / 1e10)) {
                                                                                                                v98 = (42 != 4294967295)
                                                                                                                v103 := v3[127]
                                                                                                                v6 -= 42
                                                                                                                _ = v103
                                                                                                            } else {
                                                                                                                switch (-(fi59)) {
                                                                                                                case 4294967295:
                                                                                                                    var v104 int64 = (v6 * v6)
                                                                                                                    _ = v104
                                                                                                                case -2147483648:
                                                                                                                    v105 := GenS9{F0: 42, F1: 2}
                                                                                                                    _ = v105
                                                                                                                case 32768:
                                                                                                                    var v106 string = "\n\t"
                                                                                                                    _ = v106
                                                                                                                }
                                                                                                            }
                                                                                                            v107 := GenS1{F0: v47, F1: v54, F2: -1}
                                                                                                            if (v9 != (v9 * v9)) {
                                                                                                                if v98 {
                                                                                                                    var v108 int = (fi77 ^ (^(fi59)))
                                                                                                                    fmt.Println(v9)
                                                                                                                    _ = len(v3)
                                                                                                                    _ = v108
                                                                                                                } else {
                                                                                                                    v109 := GenS7{F0: v1, F1: (v6 + 42)}
                                                                                                                    if (v98 || ((!false) || (((-2.25 + 1.5) / v23) <= v99))) {
                                                                                                                        v5[-1] = v98
                                                                                                                        fmt.Println(v1, v81, fi77)
                                                                                                                        v110 := GenS7{F0: 2, F1: v6}
                                                                                                                        v111 := GenS5{F0: 42, F1: ((("\n\t" + v47) + (v47 + ((((("foo bar" + ("a" + "\n\t")) + "") + (("foo bar" + "hello") + "hello")) + v47) + v47))) + v47), F2: 183}
                                                                                                                        var v112 uint32 = 2147483648
                                                                                                                        _ = v110
                                                                                                                        _ = v111
                                                                                                                        _ = v112
                                                                                                                    } else {
                                                                                                                        v113 := []int{-1, (fi77 ^ fi77)}
                                                                                                                        _ = len(v113)
                                                                                                                        _ = v113
                                                                                                                    }
                                                                                                                    var v114 bool = (false || v98)
                                                                                                                    var v115 uint = 0
                                                                                                                    _ = v109
                                                                                                                    _ = v114
                                                                                                                    _ = v115
                                                                                                                }
                                                                                                                if (691 >= 844) {
                                                                                                                    fmt.Println(v23, v81, v22)
                                                                                                                    v116 := GenS9{F0: 9223372036854775807, F1: 42}
                                                                                                                    var v117 uint8 = 1
                                                                                                                    var v118 uint64 = v9
                                                                                                                    v119 := GenS10{F0: 0, F1: 0}
                                                                                                                    _ = v116
                                                                                                                    _ = v117
                                                                                                                    _ = v118
                                                                                                                    _ = v119
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v107
                                                                                                        } else {
                                                                                                            switch v47 {
                                                                                                            case "hello":
                                                                                                                var v120 bool = v98
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    if (!v120) {
                                                                                                                        v121 := GenS3{F0: "foo bar"}
                                                                                                                        _ = v121
                                                                                                                    } else {
                                                                                                                        if v120 {
                                                                                                                            _ = len(v5)
                                                                                                                            fmt.Println(v9, v54)
                                                                                                                            var v122 float32 = v23
                                                                                                                            _ = v122
                                                                                                                        } else {
                                                                                                                            var v123 uint = 2
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    if (!(v98 || true)) {
                                                                                                                        var v124 uint64 = (v9 + uint64(v1))
                                                                                                                        if false {
                                                                                                                            v6 = (^(v6))
                                                                                                                            var v125 int8 = -41
                                                                                                                            _ = v125
                                                                                                                        } else {
                                                                                                                            if (42 > uint(v9)) {
                                                                                                                                if ((v22 - (v22 & (v22 % ((42) | 1)))) > v94) {
                                                                                                                                    v126 := []int{(fi77 + fi59), (fi59 - (-(fi59)))}
                                                                                                                                    if v98 {
                                                                                                                                        if v120 {
                                                                                                                                            var v127 uint = 255
                                                                                                                                            v98 = v120
                                                                                                                                            var v128 string = (v47 + (v47 + v47))
                                                                                                                                            _ = v127
                                                                                                                                            _ = v128
                                                                                                                                        } else {
                                                                                                                                            if (1 <= 127) {
                                                                                                                                                v99 = 0.1
                                                                                                                                                v6 -= -105
                                                                                                                                                v22 -= 2147483647
                                                                                                                                            } else {
                                                                                                                                                if (169 > 42) {
                                                                                                                                                    v129 := GenS4{F0: 0, F1: 0}
                                                                                                                                                    if ((!(0 < v1)) && false) {
                                                                                                                                                        switch (^(fi59)) {
                                                                                                                                                        case 127:
                                                                                                                                                            var v130 float64 = v2
                                                                                                                                                            _ = v130
                                                                                                                                                        case -92:
                                                                                                                                                            var v131 int64 = v6
                                                                                                                                                            _ = v131
                                                                                                                                                        }
                                                                                                                                                        v47 = (v47 + (v47 + v47))
                                                                                                                                                        if v120 {
                                                                                                                                                            v132 := GenS6{F0: false, F1: 179, F2: -1}
                                                                                                                                                            _ = v132
                                                                                                                                                        } else {
                                                                                                                                                            var v133 int64 = 625
                                                                                                                                                            v9 -= (v124 / ((0) | 1))
                                                                                                                                                            delete(v5, fi59)
                                                                                                                                                            _ = v133
                                                                                                                                                        }
                                                                                                                                                        if (42 >= 10) {
                                                                                                                                                            var v134 uint32 = 0
                                                                                                                                                            if len(v126) > 0 {
                                                                                                                                                                v126[0] = fi59
                                                                                                                                                            }
                                                                                                                                                            if false {
                                                                                                                                                                _ = len(v126)
                                                                                                                                                                switch "" {
                                                                                                                                                                case "hello":
                                                                                                                                                                    v135 := GenS11{F0: 4294967295}
                                                                                                                                                                    _ = v135
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            switch fi77 {
                                                                                                                                                            case -790:
                                                                                                                                                                v98 = true
                                                                                                                                                            default:
                                                                                                                                                                if v120 {
                                                                                                                                                                    if v120 {
                                                                                                                                                                        v94 += v22
                                                                                                                                                                    } else {
                                                                                                                                                                        _ = len(v4)
                                                                                                                                                                    }
                                                                                                                                                                    v136 := GenS12{F0: 221, F1: 1}
                                                                                                                                                                    _ = v136
                                                                                                                                                                } else {
                                                                                                                                                                    var v137 uint64 = 510
                                                                                                                                                                    delete(v4, ((v47 + v47) + "foo bar"))
                                                                                                                                                                    v138 := GenS5{F0: 0, F1: v47, F2: 255}
                                                                                                                                                                    _ = v137
                                                                                                                                                                    _ = v138
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            v126 = append(v126, (^(fi77)))
                                                                                                                                                            _ = v134
                                                                                                                                                        }
                                                                                                                                                        if ((v1 != v1) && (false || ((443 >= (v94 ^ (v22 >> ((198) & 63)))) || (v6 == -87)))) {
                                                                                                                                                            v126 = append(v126, fi59)
                                                                                                                                                            v139 := GenS8{F0: fi77, F1: (!false), F2: false}
                                                                                                                                                            v54 = -32768
                                                                                                                                                            _ = v139
                                                                                                                                                        } else {
                                                                                                                                                            if v120 {
                                                                                                                                                                v140 := GenS13{F0: ((v47 != v47) && (v2 != 1.5)), F1: 118, F2: 92}
                                                                                                                                                                var v141 int = (fi59 ^ 140)
                                                                                                                                                                v124 = (^(v9))
                                                                                                                                                                if v98 {
                                                                                                                                                                    var v142 uint8 = 255
                                                                                                                                                                    _ = v142
                                                                                                                                                                } else {
                                                                                                                                                                    var v143 float64 = v81
                                                                                                                                                                    v144 := GenS1{F0: ((("\n\t" + v47) + "\n\t") + ""), F1: (^(v54)), F2: 42}
                                                                                                                                                                    fmt.Println(v98, fi77)
                                                                                                                                                                    _ = v143
                                                                                                                                                                    _ = v144
                                                                                                                                                                }
                                                                                                                                                                v54 += -624
                                                                                                                                                                var v145 bool = (int32(v54) >= -2147483648)
                                                                                                                                                                _ = v140
                                                                                                                                                                _ = v141
                                                                                                                                                                _ = v145
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v98, v94, v99)
                                                                                                                                                        fmt.Println(v124)
                                                                                                                                                        if len(v126) >= 2 {
                                                                                                                                                            v126 = v126[1:len(v126)]
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        switch ("hello" + (v47 + "\n\t")) {
                                                                                                                                                        case "hello":
                                                                                                                                                            fmt.Println(v23, v99, v9)
                                                                                                                                                        case "a":
                                                                                                                                                            if false {
                                                                                                                                                                switch "\n\t" {
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    if (127 > 1) {
                                                                                                                                                                        v146 := GenS12{F0: 42, F1: 1}
                                                                                                                                                                        if ((v47 + (v47 + ((v47 + "\n\t") + v47))) <= "") {
                                                                                                                                                                            var v147 uint8 = 1
                                                                                                                                                                            _ = v147
                                                                                                                                                                        }
                                                                                                                                                                        v5[(fi59 + (-(fi59)))] = ((v54 % ((0) | 1)) == v54)
                                                                                                                                                                        v94 -= int32(v6)
                                                                                                                                                                        _ = v146
                                                                                                                                                                    }
                                                                                                                                                                case "a":
                                                                                                                                                                    fmt.Println(v9, v124, v47)
                                                                                                                                                                case "":
                                                                                                                                                                    if ((v6 - 1) < v6) {
                                                                                                                                                                        fmt.Println(v124, v81, v9)
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                switch (v47 + "a") {
                                                                                                                                                                case "a":
                                                                                                                                                                    _ = len(v126)
                                                                                                                                                                default:
                                                                                                                                                                    if (v1 <= uint16(v9)) {
                                                                                                                                                                        if ("hello" < ((("" + ("\n\t" + "foo bar")) + v47) + v47)) {
                                                                                                                                                                            var v148 uint16 = (v1 / ((v1) | 1))
                                                                                                                                                                            _ = v148
                                                                                                                                                                        } else {
                                                                                                                                                                            if len(v126) > 0 {
                                                                                                                                                                                v126[0] = (fi59 + (fi59 * int(v54)))
                                                                                                                                                                            }
                                                                                                                                                                            switch fi59 {
                                                                                                                                                                            case 2147483647:
                                                                                                                                                                                var v149 int32 = v94
                                                                                                                                                                                _ = v149
                                                                                                                                                                            case 4294967295:
                                                                                                                                                                                switch 923 {
                                                                                                                                                                                case 941:
                                                                                                                                                                                    v54 += -87
                                                                                                                                                                                }
                                                                                                                                                                            case -132:
                                                                                                                                                                                if true {
                                                                                                                                                                                    if len(v126) >= 2 {
                                                                                                                                                                                        v126 = v126[1:len(v126)]
                                                                                                                                                                                    }
                                                                                                                                                                                } else {
                                                                                                                                                                                    if v120 {
                                                                                                                                                                                        if v120 {
                                                                                                                                                                                            v150 := GenS6{F0: (!v120), F1: 42, F2: v94}
                                                                                                                                                                                            v3[1] = (v47 + "")
                                                                                                                                                                                            _ = v150
                                                                                                                                                                                        } else {
                                                                                                                                                                                            if (uint(v9) < 4294967295) {
                                                                                                                                                                                                if (!v120) {
                                                                                                                                                                                                    if (true && (100 > int8(v6))) {
                                                                                                                                                                                                        v9 = 42
                                                                                                                                                                                                        if len(v126) > 0 {
                                                                                                                                                                                                            v126[0] = fi77
                                                                                                                                                                                                        }
                                                                                                                                                                                                    } else {
                                                                                                                                                                                                        v1 += (^(v1))
                                                                                                                                                                                                        v81 -= v2
                                                                                                                                                                                                        v99 = (((((v99 / 1.5) * (v23 - float32(v81))) - (float32(((-2.25 * v81) / ((1e10 - (v81 / float64(v23))) / 1e10))) - v99)) - (v23 + v99)) - v23)
                                                                                                                                                                                                    }
                                                                                                                                                                                                    var v151 uint64 = 2
                                                                                                                                                                                                    if true {
                                                                                                                                                                                                        v152 := GenS14{F0: v22}
                                                                                                                                                                                                        _ = v152
                                                                                                                                                                                                    }
                                                                                                                                                                                                    v153 := GenS8{F0: fi77, F1: v98, F2: true}
                                                                                                                                                                                                    var v154 uint = 65535
                                                                                                                                                                                                    _ = v151
                                                                                                                                                                                                    _ = v153
                                                                                                                                                                                                    _ = v154
                                                                                                                                                                                                }
                                                                                                                                                                                                v81 = float64((-2.25 * v99))
                                                                                                                                                                                            }
                                                                                                                                                                                        }
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v155 := v4[(v47 + v47)]
                                                                                                                                                                                        var v156 uint64 = v124
                                                                                                                                                                                        switch fi77 {
                                                                                                                                                                                        case 128:
                                                                                                                                                                                            v23 -= (v99 / (-2.25 / float32((1.5 + (float64(v23) / (v2 / (v2 - float64(float32(0.0)))))))))
                                                                                                                                                                                            switch (^(fi59)) {
                                                                                                                                                                                            case 2147483647:
                                                                                                                                                                                                if ((float32(v81) - (v99 / (v99 * v23))) > (v23 + v23)) {
                                                                                                                                                                                                    if v120 {
                                                                                                                                                                                                        switch "hello" {
                                                                                                                                                                                                        case "a":
                                                                                                                                                                                                            var v157 int32 = (-(v94))
                                                                                                                                                                                                            _ = v157
                                                                                                                                                                                                        }
                                                                                                                                                                                                    } else {
                                                                                                                                                                                                        var v158 uint16 = 412
                                                                                                                                                                                                        v9 += 1
                                                                                                                                                                                                        _ = v158
                                                                                                                                                                                                    }
                                                                                                                                                                                                    v159 := v5[fi59]
                                                                                                                                                                                                    if v155 {
                                                                                                                                                                                                        var v160 uint8 = 2
                                                                                                                                                                                                        _ = v160
                                                                                                                                                                                                    }
                                                                                                                                                                                                    v1 = v1
                                                                                                                                                                                                    var v161 uint = 42
                                                                                                                                                                                                    _ = v159
                                                                                                                                                                                                    _ = v161
                                                                                                                                                                                                }
                                                                                                                                                                                            case -401:
                                                                                                                                                                                                var v162 uint64 = (v124 / (((^(v156))) | 1))
                                                                                                                                                                                                _ = v162
                                                                                                                                                                                            default:
                                                                                                                                                                                                var v163 int32 = v22
                                                                                                                                                                                                _ = v163
                                                                                                                                                                                            }
                                                                                                                                                                                        case 174:
                                                                                                                                                                                            var v164 uint32 = uint32(v9)
                                                                                                                                                                                            _ = v164
                                                                                                                                                                                        case -9223372036854775808:
                                                                                                                                                                                            switch 2147483647 {
                                                                                                                                                                                            case 2147483647:
                                                                                                                                                                                                var v165 uint32 = 1
                                                                                                                                                                                                _ = v165
                                                                                                                                                                                            default:
                                                                                                                                                                                                v166 := GenS6{F0: (!((v124 & 4294967295) >= 4294967295)), F1: 127, F2: (v22 - v22)}
                                                                                                                                                                                                _ = v166
                                                                                                                                                                                            }
                                                                                                                                                                                        default:
                                                                                                                                                                                            v167 := GenS15{F0: v124}
                                                                                                                                                                                            v168 := GenS16{F0: "", F1: (v54 | v54)}
                                                                                                                                                                                            _ = v167
                                                                                                                                                                                            _ = v168
                                                                                                                                                                                        }
                                                                                                                                                                                        _ = v155
                                                                                                                                                                                        _ = v156
                                                                                                                                                                                    }
                                                                                                                                                                                    var v169 bool = v98
                                                                                                                                                                                    var v170 string = v47
                                                                                                                                                                                    var v171 int64 = 9223372036854775807
                                                                                                                                                                                    _ = v169
                                                                                                                                                                                    _ = v170
                                                                                                                                                                                    _ = v171
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                        if (!v120) {
                                                                                                                                                                            if len(v126) > 0 {
                                                                                                                                                                                v126[0] = (fi59 + fi59)
                                                                                                                                                                            }
                                                                                                                                                                            var v172 int8 = 0
                                                                                                                                                                            if (v124 < v124) {
                                                                                                                                                                                v1 += v1
                                                                                                                                                                                if false {
                                                                                                                                                                                    if v120 {
                                                                                                                                                                                        var v173 uint32 = 1
                                                                                                                                                                                        switch fi77 {
                                                                                                                                                                                        case 9223372036854775807:
                                                                                                                                                                                            var v174 uint = 65535
                                                                                                                                                                                            _ = v174
                                                                                                                                                                                        case 0:
                                                                                                                                                                                            var v175 int64 = 475
                                                                                                                                                                                            _ = v175
                                                                                                                                                                                        default:
                                                                                                                                                                                            var v176 uint = 18446744073709551615
                                                                                                                                                                                            _ = v176
                                                                                                                                                                                        }
                                                                                                                                                                                        _ = v173
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v172 = 1
                                                                                                                                                                                        v172 = 62
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                v23 -= (v23 / (v99 + 1e10))
                                                                                                                                                                                var v177 bool = (((-2.25 * v23) + v99) == v99)
                                                                                                                                                                                v54 = (v54 / ((v54) | 1))
                                                                                                                                                                                _ = v177
                                                                                                                                                                            } else {
                                                                                                                                                                                switch fi77 {
                                                                                                                                                                                case 453:
                                                                                                                                                                                    var v178 int = (fi77 >> ((42) & 63))
                                                                                                                                                                                    _ = v178
                                                                                                                                                                                default:
                                                                                                                                                                                    if len(v126) >= 2 {
                                                                                                                                                                                        v126 = v126[1:len(v126)]
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                var v179 uint16 = 65535
                                                                                                                                                                                _ = v179
                                                                                                                                                                            }
                                                                                                                                                                            _ = v172
                                                                                                                                                                        } else {
                                                                                                                                                                            fmt.Println(v94)
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                switch fi77 {
                                                                                                                                                                case 211:
                                                                                                                                                                    var v180 float32 = float32(1e10)
                                                                                                                                                                    _ = v180
                                                                                                                                                                case 1:
                                                                                                                                                                    fmt.Println(v94, v98, v124)
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        case "":
                                                                                                                                                            switch 32768 {
                                                                                                                                                            case 9223372036854775807:
                                                                                                                                                                if (!((!((0.0 - ((v81 * 1.5) - v81)) < 3.14159)) || true)) {
                                                                                                                                                                    v181 := GenS17{F0: v47, F1: 900, F2: -1}
                                                                                                                                                                    fmt.Println(v98, v9, fi59)
                                                                                                                                                                    _ = v181
                                                                                                                                                                } else {
                                                                                                                                                                    v23 = v99
                                                                                                                                                                }
                                                                                                                                                            default:
                                                                                                                                                                v182 := GenS18{F0: v6, F1: (18446744073709551615 >= (v9 & v9)), F2: ""}
                                                                                                                                                                _ = v182
                                                                                                                                                            }
                                                                                                                                                        default:
                                                                                                                                                            if (v1 > v1) {
                                                                                                                                                                fmt.Println(fi77, v98, v6)
                                                                                                                                                                if len(v126) > 0 {
                                                                                                                                                                    v126[0] = (fi77 & (fi77 + (-(fi59))))
                                                                                                                                                                }
                                                                                                                                                                var v183 uint16 = v1
                                                                                                                                                                var v184 string = "a"
                                                                                                                                                                v124 = (v9 & v124)
                                                                                                                                                                _ = v183
                                                                                                                                                                _ = v184
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        switch (^(fi77)) {
                                                                                                                                                        case -658:
                                                                                                                                                            v120 = v98
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    _ = v129
                                                                                                                                                }
                                                                                                                                                var v185 string = ("foo bar" + v47)
                                                                                                                                                _ = v185
                                                                                                                                            }
                                                                                                                                            delete(v5, int(v6))
                                                                                                                                            fmt.Println(v47)
                                                                                                                                        }
                                                                                                                                        _ = len(v126)
                                                                                                                                        var v186 uint64 = (v124 ^ (v124 & 42))
                                                                                                                                        _ = v186
                                                                                                                                    } else {
                                                                                                                                        switch ("a" + v47) {
                                                                                                                                        case "\n\t":
                                                                                                                                            var v187 uint64 = (v9 + v124)
                                                                                                                                            _ = v187
                                                                                                                                        default:
                                                                                                                                            switch int(v22) {
                                                                                                                                            case 42:
                                                                                                                                                v188 := v3[fi77]
                                                                                                                                                _ = v188
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    var v189 uint16 = 609
                                                                                                                                    _ = v126
                                                                                                                                    _ = v189
                                                                                                                                } else {
                                                                                                                                    v99 = (v23 + v99)
                                                                                                                                    v1 = v1
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if v98 {
                                                                                                                                    v81 += v2
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v190 uint16 = (v1 ^ v1)
                                                                                                                            if true {
                                                                                                                                v2 -= (3.14159 - (1.5 * (v2 - v2)))
                                                                                                                                v191 := GenS14{F0: (v94 ^ (v22 << ((uint(v190)) & 63)))}
                                                                                                                                _ = v191
                                                                                                                            } else {
                                                                                                                                v192 := []int{}
                                                                                                                                if (0 == v54) {
                                                                                                                                    v193 := GenS15{F0: v9}
                                                                                                                                    if (135 > 2) {
                                                                                                                                        switch fi77 {
                                                                                                                                        case 9223372036854775807:
                                                                                                                                            v1 = v1
                                                                                                                                        case 2147483647:
                                                                                                                                            v190 = 1
                                                                                                                                        case 32767:
                                                                                                                                            v190 -= 617
                                                                                                                                            var v194 int16 = (v54 + (v54 + (^(v54))))
                                                                                                                                            if (!v98) {
                                                                                                                                                if v98 {
                                                                                                                                                    v47 = ("" + v47)
                                                                                                                                                    switch 127 {
                                                                                                                                                    case -9223372036854775808:
                                                                                                                                                        fmt.Println(v54)
                                                                                                                                                    default:
                                                                                                                                                        var v195 float64 = v2
                                                                                                                                                        _ = v195
                                                                                                                                                    }
                                                                                                                                                    var v196 int16 = (v194 * (v54 + 817))
                                                                                                                                                    _ = v196
                                                                                                                                                } else {
                                                                                                                                                    v124 = 18446744073709551615
                                                                                                                                                }
                                                                                                                                                _ = len(v192)
                                                                                                                                                var v197 float64 = 1.5
                                                                                                                                                switch (fi59 * 9223372036854775807) {
                                                                                                                                                case 298:
                                                                                                                                                    var v198 int8 = 0
                                                                                                                                                    _ = v198
                                                                                                                                                case 0:
                                                                                                                                                    v98 = (4294967295 > 1)
                                                                                                                                                case 2147483647:
                                                                                                                                                    var v199 int = fi59
                                                                                                                                                    _ = v199
                                                                                                                                                }
                                                                                                                                                _ = v197
                                                                                                                                            }
                                                                                                                                            _ = v194
                                                                                                                                        }
                                                                                                                                        v6 = v6
                                                                                                                                        fmt.Println(v2)
                                                                                                                                        v1 += v190
                                                                                                                                        var v200 uint32 = 303
                                                                                                                                        _ = v200
                                                                                                                                    }
                                                                                                                                    _ = v193
                                                                                                                                }
                                                                                                                                v201 := GenS14{F0: v94}
                                                                                                                                _ = v192
                                                                                                                                _ = v201
                                                                                                                            }
                                                                                                                            _ = v190
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            switch 1 {
                                                                                                                            case -390:
                                                                                                                                v202 := GenS10{F0: 49, F1: 1}
                                                                                                                                _ = v202
                                                                                                                            case 4294967295:
                                                                                                                                var v203 string = v47
                                                                                                                                _ = v203
                                                                                                                            default:
                                                                                                                                fmt.Println(v54)
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if v120 {
                                                                                                                                var v204 int8 = 1
                                                                                                                                var v205 uint8 = 121
                                                                                                                                v206 := v4[(v47 + v47)]
                                                                                                                                _ = v204
                                                                                                                                _ = v205
                                                                                                                                _ = v206
                                                                                                                            } else {
                                                                                                                                var v207 string = v47
                                                                                                                                v208 := GenS10{F0: 0, F1: 0}
                                                                                                                                _ = v207
                                                                                                                                _ = v208
                                                                                                                            }
                                                                                                                            v99 -= 1e10
                                                                                                                            v209 := GenS1{F0: (v47 + v47), F1: -190, F2: 99}
                                                                                                                            _ = v209
                                                                                                                        }
                                                                                                                        _ = v124
                                                                                                                    } else {
                                                                                                                        v120 = (v22 > v22)
                                                                                                                        v210 := GenS19{F0: (v9 ^ (v9 | 262))}
                                                                                                                        _ = v210
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v120
                                                                                                            case "a":
                                                                                                                v211 := []int{fi59, (fi77 - (fi59 - -187)), (-(fi59)), 32768}
                                                                                                                _ = v211
                                                                                                            }
                                                                                                            var v212 int = (fi77 ^ (fi77 + (fi59 + -330)))
                                                                                                            switch int(v6) {
                                                                                                            case 0:
                                                                                                                v213 := GenS6{F0: false, F1: 128, F2: v94}
                                                                                                                _ = v213
                                                                                                            }
                                                                                                            var v214 uint8 = 128
                                                                                                            var v215 int64 = 9223372036854775807
                                                                                                            if ((v6 ^ (v6 | (v215 / ((int64(v214)) | 1)))) <= v6) {
                                                                                                                if false {
                                                                                                                    v81 += v81
                                                                                                                    v216 := []int{}
                                                                                                                    _ = len(v216)
                                                                                                                    v217 := GenS20{F0: (fi77 + (fi77 + (v212 | (fi59 % ((v212) | 1))))), F1: (!v98), F2: ((v47 + "a") + "hello")}
                                                                                                                    _ = v216
                                                                                                                    _ = v217
                                                                                                                } else {
                                                                                                                    var v218 uint32 = 449
                                                                                                                    v5[fi59] = v98
                                                                                                                    switch v212 {
                                                                                                                    case -2147483648:
                                                                                                                        if (!(!v98)) {
                                                                                                                            v219 := GenS21{F0: (v9 << ((2) & 63)), F1: 875}
                                                                                                                            v6 = -473
                                                                                                                            _ = v219
                                                                                                                        } else {
                                                                                                                            v94 += v94
                                                                                                                            if v98 {
                                                                                                                                v218 += v218
                                                                                                                                v220 := []int{32768, int(v9), fi59}
                                                                                                                                var v221 uint64 = 0
                                                                                                                                _ = v220
                                                                                                                                _ = v221
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case 127:
                                                                                                                        var v222 int32 = -2147483648
                                                                                                                        _ = v222
                                                                                                                    case 595:
                                                                                                                        v23 = v23
                                                                                                                        var v223 uint = 65535
                                                                                                                        var v224 uint16 = v1
                                                                                                                        _ = v223
                                                                                                                        _ = v224
                                                                                                                    default:
                                                                                                                        v81 = ((v81 + (float64(v99) + 0.0)) - 0.0)
                                                                                                                    }
                                                                                                                    if v98 {
                                                                                                                        if true {
                                                                                                                            v3[fi59] = v47
                                                                                                                            v225 := GenS2{F0: (v214 - v214), F1: (v98 && true)}
                                                                                                                            _ = v225
                                                                                                                        } else {
                                                                                                                            v226 := []int{v212, 127}
                                                                                                                            v218 -= 4294967295
                                                                                                                            if v98 {
                                                                                                                                v227 := GenS22{F0: 42, F1: ("hello" + v47), F2: 4294967295}
                                                                                                                                v1 -= (v1 >> ((1) & 63))
                                                                                                                                _ = v227
                                                                                                                            }
                                                                                                                            _ = v226
                                                                                                                        }
                                                                                                                        switch v212 {
                                                                                                                        case 32768:
                                                                                                                            fmt.Println(fi59)
                                                                                                                        case -9223372036854775808:
                                                                                                                            var v228 int = (fi59 - int(v1))
                                                                                                                            _ = v228
                                                                                                                        case 42:
                                                                                                                            var v229 uint = uint(v218)
                                                                                                                            _ = v229
                                                                                                                        default:
                                                                                                                            delete(v3, (fi77 * (fi77 * (fi77 | int(v1)))))
                                                                                                                        }
                                                                                                                        fmt.Println(v218, v9, v218)
                                                                                                                    } else {
                                                                                                                        var v230 float32 = 1.5
                                                                                                                        if (2 == uint(v218)) {
                                                                                                                            if true {
                                                                                                                                v218 -= (v218 ^ (v218 & (v218 % (((^(v218))) | 1))))
                                                                                                                                v218 += (v218 - 810)
                                                                                                                                if v98 {
                                                                                                                                    var v231 int32 = (v22 & 0)
                                                                                                                                    v232 := GenS8{F0: 939, F1: (v231 <= v94), F2: true}
                                                                                                                                    switch fi59 {
                                                                                                                                    case 32767:
                                                                                                                                        v233 := GenS23{F0: 219}
                                                                                                                                        _ = v233
                                                                                                                                    case 32768:
                                                                                                                                        v6 += (v215 & (v215 % ((-1) | 1)))
                                                                                                                                    case -607:
                                                                                                                                        v234 := []int{(v212 & (fi77 & -9223372036854775808))}
                                                                                                                                        v235 := GenS23{F0: (v54 & (v54 << ((0) & 63)))}
                                                                                                                                        v230 = float32(v81)
                                                                                                                                        _ = v234
                                                                                                                                        _ = v235
                                                                                                                                    }
                                                                                                                                    _ = v231
                                                                                                                                    _ = v232
                                                                                                                                } else {
                                                                                                                                    v236 := GenS24{F0: (^(v22)), F1: (((^(v218)) != v218) || v98)}
                                                                                                                                    var v237 int64 = v215
                                                                                                                                    _ = v236
                                                                                                                                    _ = v237
                                                                                                                                }
                                                                                                                                v22 = v22
                                                                                                                                v94 = int32(v9)
                                                                                                                                v81 += v2
                                                                                                                            } else {
                                                                                                                                fmt.Println(fi59, v23, fi59)
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v230
                                                                                                                    }
                                                                                                                    _ = v218
                                                                                                                }
                                                                                                                if v98 {
                                                                                                                    switch 125 {
                                                                                                                    case 128:
                                                                                                                        var v238 float32 = 1e10
                                                                                                                        _ = v238
                                                                                                                    case 4294967295:
                                                                                                                        if true {
                                                                                                                            if (v98 || (v47 <= ((v47 + ("a" + v47)) + v47))) {
                                                                                                                                v214 -= (v214 - 255)
                                                                                                                            } else {
                                                                                                                                _ = len(v3)
                                                                                                                                if ((true && v98) && v98) {
                                                                                                                                    var v239 uint8 = (v214 << ((0) & 63))
                                                                                                                                    switch (fi77 | (fi77 ^ (fi59 & fi59))) {
                                                                                                                                    case -894:
                                                                                                                                        v240 := GenS25{F0: (v47 + v47), F1: v1}
                                                                                                                                        _ = v240
                                                                                                                                    case -36:
                                                                                                                                        var v241 uint = 4294967295
                                                                                                                                        _ = v241
                                                                                                                                    case -579:
                                                                                                                                        fmt.Println(v23)
                                                                                                                                        var v242 int32 = 68
                                                                                                                                        v94 -= (v242 * (v242 / (((v94 >> ((18446744073709551615) & 63))) | 1)))
                                                                                                                                        _ = v242
                                                                                                                                    default:
                                                                                                                                        if (!v98) {
                                                                                                                                            var v243 uint32 = 897
                                                                                                                                            v244 := GenS25{F0: v47, F1: v1}
                                                                                                                                            _ = v243
                                                                                                                                            _ = v244
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v94, v54, v54)
                                                                                                                                            if true {
                                                                                                                                                var v245 float32 = 0.1
                                                                                                                                                var v246 uint64 = 4294967295
                                                                                                                                                _ = v245
                                                                                                                                                _ = v246
                                                                                                                                            } else {
                                                                                                                                                v247 := GenS26{F0: (fi77 % ((fi77) | 1)), F1: fi59}
                                                                                                                                                v215 = (-(v6))
                                                                                                                                                v239 -= (v239 | v214)
                                                                                                                                                _ = v247
                                                                                                                                            }
                                                                                                                                            if v98 {
                                                                                                                                                fmt.Println(v6)
                                                                                                                                                var v248 int16 = 42
                                                                                                                                                _ = v248
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v81 = (((float64(-2.25) - v81) / v81) / v2)
                                                                                                                                    _ = v239
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v249 := v3[474]
                                                                                                                            v98 = false
                                                                                                                            _ = v249
                                                                                                                        } else {
                                                                                                                            var v250 bool = false
                                                                                                                            if ((v9 / (((v9 % (((v9 | 42)) | 1))) | 1)) > v9) {
                                                                                                                                v54 -= (v54 % ((0) | 1))
                                                                                                                                v251 := v4[v47]
                                                                                                                                fmt.Println(v212)
                                                                                                                                if true {
                                                                                                                                    if v251 {
                                                                                                                                        if v251 {
                                                                                                                                            switch "a" {
                                                                                                                                            case "\n\t":
                                                                                                                                                var v252 float32 = v99
                                                                                                                                                _ = v252
                                                                                                                                            case "foo bar":
                                                                                                                                                v253 := []int{int(v54), fi77, fi77, (fi77 * v212)}
                                                                                                                                                _ = v253
                                                                                                                                            default:
                                                                                                                                                v6 -= 42
                                                                                                                                                var v254 uint = 2
                                                                                                                                                _ = v254
                                                                                                                                            }
                                                                                                                                            v94 -= (^(v94))
                                                                                                                                            v3[-327] = v47
                                                                                                                                            if ((^(v54)) < v54) {
                                                                                                                                                v255 := GenS23{F0: 337}
                                                                                                                                                v3[-115] = "a"
                                                                                                                                                if v251 {
                                                                                                                                                    var v256 bool = false
                                                                                                                                                    _ = v256
                                                                                                                                                }
                                                                                                                                                v251 = (127 < (v214 >> ((0) & 63)))
                                                                                                                                                v2 += ((3.14159 + v81) - -2.25)
                                                                                                                                                var v257 int32 = v94
                                                                                                                                                _ = v255
                                                                                                                                                _ = v257
                                                                                                                                            } else {
                                                                                                                                                delete(v4, v47)
                                                                                                                                            }
                                                                                                                                            var v258 uint64 = (v9 * (v9 & 0))
                                                                                                                                            v259 := GenS26{F0: fi77, F1: v212}
                                                                                                                                            if true {
                                                                                                                                                var v260 int32 = 0
                                                                                                                                                var v261 uint16 = v1
                                                                                                                                                _ = v260
                                                                                                                                                _ = v261
                                                                                                                                            } else {
                                                                                                                                                v262 := v5[(-(fi59))]
                                                                                                                                                if ((v6 >> ((2) & 63)) >= v6) {
                                                                                                                                                    v263 := GenS1{F0: (v47 + ("\n\t" + ("hello" + (("\n\t" + v47) + v47)))), F1: 42, F2: -1}
                                                                                                                                                    fmt.Println(fi59, v1)
                                                                                                                                                    _ = v263
                                                                                                                                                } else {
                                                                                                                                                    switch v47 {
                                                                                                                                                    case "":
                                                                                                                                                        v258 -= v9
                                                                                                                                                    case "a":
                                                                                                                                                        if (!(!(1 == 0))) {
                                                                                                                                                            v94 = (v94 << ((65535) & 63))
                                                                                                                                                            var v264 int16 = (v54 ^ v54)
                                                                                                                                                            v47 = (("a" + "") + v47)
                                                                                                                                                            var v265 bool = (!false)
                                                                                                                                                            v266 := GenS20{F0: 260, F1: v98, F2: "a"}
                                                                                                                                                            switch (v47 + v47) {
                                                                                                                                                            case "foo bar":
                                                                                                                                                                if ((42 != (^(fi77))) && v250) {
                                                                                                                                                                    var v267 int = (fi59 & -2147483648)
                                                                                                                                                                    v268 := GenS6{F0: v98, F1: v214, F2: -2147483648}
                                                                                                                                                                    v1 += (v1 ^ 952)
                                                                                                                                                                    v269 := v4[v47]
                                                                                                                                                                    var v270 int16 = (v264 / (((^(v264))) | 1))
                                                                                                                                                                    var v271 int64 = (v6 << ((804) & 63))
                                                                                                                                                                    _ = v267
                                                                                                                                                                    _ = v268
                                                                                                                                                                    _ = v269
                                                                                                                                                                    _ = v270
                                                                                                                                                                    _ = v271
                                                                                                                                                                }
                                                                                                                                                            case "\n\t":
                                                                                                                                                                switch ("" + v47) {
                                                                                                                                                                case "":
                                                                                                                                                                    var v272 float32 = v99
                                                                                                                                                                    _ = v272
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    v273 := []int{v212}
                                                                                                                                                                    _ = v273
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    var v274 string = ""
                                                                                                                                                                    _ = v274
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            _ = v264
                                                                                                                                                            _ = v265
                                                                                                                                                            _ = v266
                                                                                                                                                        }
                                                                                                                                                    case "\n\t":
                                                                                                                                                        switch v47 {
                                                                                                                                                        case "":
                                                                                                                                                            v275 := GenS27{F0: 1}
                                                                                                                                                            _ = v275
                                                                                                                                                        case "hello":
                                                                                                                                                            delete(v5, int(v215))
                                                                                                                                                        default:
                                                                                                                                                            switch fi59 {
                                                                                                                                                            case 789:
                                                                                                                                                                v5[(fi59 / (((v212 % ((-2147483648) | 1))) | 1))] = v262
                                                                                                                                                            case 9223372036854775807:
                                                                                                                                                                var v276 int32 = 42
                                                                                                                                                                _ = v276
                                                                                                                                                            case 869:
                                                                                                                                                                fmt.Println(v250, v98)
                                                                                                                                                            default:
                                                                                                                                                                v277 := GenS12{F0: 1, F1: 1000}
                                                                                                                                                                _ = v277
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    var v278 int16 = (v54 | (v54 & 940))
                                                                                                                                                    _ = v278
                                                                                                                                                }
                                                                                                                                                _ = v262
                                                                                                                                            }
                                                                                                                                            _ = v258
                                                                                                                                            _ = v259
                                                                                                                                        } else {
                                                                                                                                            v279 := GenS28{F0: -886, F1: 73, F2: v54}
                                                                                                                                            if ((v214 << ((42) & 63)) != 95) {
                                                                                                                                                var v280 int64 = v215
                                                                                                                                                fmt.Println(v94, v23, fi59)
                                                                                                                                                fmt.Println(v98)
                                                                                                                                                if v251 {
                                                                                                                                                    fmt.Println(v94, fi77)
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v47)
                                                                                                                                                    v281 := GenS29{F0: (v54 & (v54 + (v54 | 1000))), F1: (v22 >> ((0) & 63)), F2: uint8(v1)}
                                                                                                                                                    fmt.Println(v215)
                                                                                                                                                    _ = v281
                                                                                                                                                }
                                                                                                                                                var v282 string = "a"
                                                                                                                                                delete(v3, 42)
                                                                                                                                                _ = v280
                                                                                                                                                _ = v282
                                                                                                                                            }
                                                                                                                                            v283 := GenS9{F0: (fi59 - v212), F1: 171}
                                                                                                                                            _ = v279
                                                                                                                                            _ = v283
                                                                                                                                        }
                                                                                                                                        if ((!(208 < (v54 % ((v54) | 1)))) && (-599 <= v215)) {
                                                                                                                                            v5[fi59] = (v251 && (true || v251))
                                                                                                                                            switch (v212 << ((507) & 63)) {
                                                                                                                                            case -79:
                                                                                                                                                fmt.Println(v250, fi59, v6)
                                                                                                                                            case -169:
                                                                                                                                                switch fi59 {
                                                                                                                                                case -341:
                                                                                                                                                    v9 = (v9 ^ v9)
                                                                                                                                                default:
                                                                                                                                                    v250 = v250
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            v81 = v81
                                                                                                                                        } else {
                                                                                                                                            var v284 string = "\n\t"
                                                                                                                                            v214 = uint8(v1)
                                                                                                                                            var v285 uint8 = v214
                                                                                                                                            _ = v284
                                                                                                                                            _ = v285
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v3[(fi59 >> ((65535) & 63))] = v47
                                                                                                                                    v23 = 3.14159
                                                                                                                                } else {
                                                                                                                                    v286 := GenS30{F0: (v1 ^ v1)}
                                                                                                                                    v287 := GenS31{F0: (^(v212)), F1: 0, F2: (v54 << ((42) & 63))}
                                                                                                                                    v288 := GenS32{F0: 0, F1: v1, F2: 800}
                                                                                                                                    _ = v286
                                                                                                                                    _ = v287
                                                                                                                                    _ = v288
                                                                                                                                }
                                                                                                                                _ = v251
                                                                                                                            }
                                                                                                                            fmt.Println(v81, v214, v2)
                                                                                                                            _ = v250
                                                                                                                        }
                                                                                                                    case 32768:
                                                                                                                        v289 := []int{(fi77 % ((fi59) | 1)), 127}
                                                                                                                        _ = v289
                                                                                                                    }
                                                                                                                    switch v47 {
                                                                                                                    case "foo bar":
                                                                                                                        var v290 uint8 = 238
                                                                                                                        _ = v290
                                                                                                                    case "":
                                                                                                                        delete(v5, fi77)
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v212, v6)
                                                                                                                    }
                                                                                                                    var v291 uint16 = v1
                                                                                                                    _ = v291
                                                                                                                }
                                                                                                            } else {
                                                                                                                v99 = v99
                                                                                                            }
                                                                                                            _ = v212
                                                                                                            _ = v214
                                                                                                            _ = v215
                                                                                                        }
                                                                                                        _ = v102
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v2 += (0.0 / v81)
                                                                                            }
                                                                                            _ = v98
                                                                                            _ = v99
                                                                                        } else {
                                                                                            v1 += (v1 | (v1 % (((^(v1))) | 1)))
                                                                                            v292 := GenS21{F0: 479, F1: (v94 >> ((65535) & 63))}
                                                                                            v293 := []int{-9223372036854775808, -9223372036854775808}
                                                                                            v22 = v94
                                                                                            v294 := GenS33{F0: fi77}
                                                                                            if (true || false) {
                                                                                                var v295 uint32 = 345
                                                                                                _ = v295
                                                                                            } else {
                                                                                                v6 -= 801
                                                                                                v296 := GenS34{F0: (v9 ^ v9), F1: 65535, F2: fi77}
                                                                                                if len(v293) > 0 {
                                                                                                    v293[0] = fi77
                                                                                                }
                                                                                                _ = v296
                                                                                            }
                                                                                            _ = v292
                                                                                            _ = v293
                                                                                            _ = v294
                                                                                        }
                                                                                        switch v47 {
                                                                                        case "foo bar":
                                                                                            switch fi59 {
                                                                                            case -1:
                                                                                                v297 := GenS9{F0: 4294967295, F1: 2147483648}
                                                                                                if (!(0 != 67)) {
                                                                                                    if (("\n\t" + v47) == v47) {
                                                                                                        if false {
                                                                                                            var v298 uint = 65535
                                                                                                            _ = v298
                                                                                                        }
                                                                                                    } else {
                                                                                                        v299 := GenS2{F0: 55, F1: ((4294967295 < 65535) || (v23 > 0.0))}
                                                                                                        switch v47 {
                                                                                                        case "":
                                                                                                            v22 -= 42
                                                                                                        case "a":
                                                                                                            v5[32768] = ((v94 << ((0) & 63)) > v94)
                                                                                                            var v300 bool = (63 < -18)
                                                                                                            _ = v300
                                                                                                        }
                                                                                                        _ = v299
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    var v301 uint64 = v9
                                                                                                    fmt.Println(v94, v1, v9)
                                                                                                    _ = v301
                                                                                                }
                                                                                                _ = v297
                                                                                            case -613:
                                                                                                var v302 int8 = 0
                                                                                                _ = v302
                                                                                            default:
                                                                                                _ = len(v3)
                                                                                            }
                                                                                        }
                                                                                        _ = v94
                                                                                    }
                                                                                    _ = v92
                                                                                }
                                                                                var v303 int8 = 1
                                                                                _ = len(v3)
                                                                                _ = v303
                                                                            }
                                                                        }
                                                                        var v304 int16 = v54
                                                                        fmt.Println(v304)
                                                                        if true {
                                                                            v305 := GenS22{F0: 103, F1: ("hello" + v47), F2: (v9 + 18446744073709551615)}
                                                                            _ = v305
                                                                        }
                                                                        if true {
                                                                            v4[("" + "")] = (2 <= 4294967295)
                                                                        } else {
                                                                            var v306 float64 = 0.0
                                                                            if ((^(fi77)) != 42) {
                                                                                var v307 uint = 0
                                                                                _ = v307
                                                                            }
                                                                            _ = v306
                                                                        }
                                                                        _ = v304
                                                                        _ = v81
                                                                    } else {
                                                                        if true {
                                                                            switch (fi77 / (((fi59 + (fi77 ^ fi59))) | 1)) {
                                                                            case -9223372036854775808:
                                                                                v308 := GenS35{F0: (fi59 % ((fi77) | 1)), F1: -856}
                                                                                _ = v308
                                                                            case -90:
                                                                                if true {
                                                                                    if (!(((476 >= (v22 - (v22 | (^(v22))))) && (v2 < 1e10)) && true)) {
                                                                                        switch v47 {
                                                                                        case "":
                                                                                            v309 := GenS36{F0: 260, F1: 4294967295}
                                                                                            _ = v309
                                                                                        case "\n\t":
                                                                                            var v310 float64 = 1e10
                                                                                            switch fi59 {
                                                                                            case -446:
                                                                                                v311 := GenS16{F0: "foo bar", F1: 1000}
                                                                                                _ = v311
                                                                                            }
                                                                                            var v312 int64 = 9223372036854775807
                                                                                            _ = v310
                                                                                            _ = v312
                                                                                        }
                                                                                        _ = len(v4)
                                                                                    }
                                                                                    var v313 bool = ((v23 * v23) == 3.14159)
                                                                                    v6 += (v6 | int64(v1))
                                                                                    _ = v313
                                                                                }
                                                                                v1 -= (v1 % ((32767) | 1))
                                                                                v314 := GenS29{F0: (v54 + v54), F1: v22, F2: 128}
                                                                                if (387 == (v9 + (v9 & 2))) {
                                                                                    if (!false) {
                                                                                        var v315 int = fi77
                                                                                        var v316 int32 = v22
                                                                                        switch (v47 + "foo bar") {
                                                                                        case "hello":
                                                                                            var v317 float64 = (v2 / 1.0)
                                                                                            _ = v317
                                                                                        case "foo bar":
                                                                                            if (false || false) {
                                                                                                v318 := []int{}
                                                                                                v2 = ((v2 / v2) / -2.25)
                                                                                                _ = v318
                                                                                            } else {
                                                                                                var v319 bool = false
                                                                                                _ = v319
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(v315, v9)
                                                                                        default:
                                                                                            v320 := GenS37{F0: 0, F1: v9}
                                                                                            _ = v320
                                                                                        }
                                                                                        var v321 int64 = v6
                                                                                        _ = v315
                                                                                        _ = v316
                                                                                        _ = v321
                                                                                    } else {
                                                                                        v322 := GenS1{F0: "foo bar", F1: (v54 + -1), F2: 0}
                                                                                        fmt.Println(fi59, v47, v22)
                                                                                        v323 := GenS38{F0: 486, F1: uint32(v1)}
                                                                                        _ = v322
                                                                                        _ = v323
                                                                                    }
                                                                                    fmt.Println(v6, v9)
                                                                                    if (!(-765 <= int64(fi77))) {
                                                                                        if (v2 != ((0.1 - 1e10) * (-2.25 / (1.5 / 3.14159)))) {
                                                                                            v324 := GenS39{F0: 0, F1: v22, F2: v6}
                                                                                            fmt.Println(v9, v6)
                                                                                            v325 := GenS6{F0: (127 != 101), F1: 1, F2: (v22 + v22)}
                                                                                            var v326 int16 = (v54 / (((v54 >> ((2) & 63))) | 1))
                                                                                            _ = v324
                                                                                            _ = v325
                                                                                            _ = v326
                                                                                        }
                                                                                        v327 := []int{(fi59 | (fi59 + fi77)), (fi77 << ((544) & 63))}
                                                                                        var v328 float64 = 1e10
                                                                                        _ = v327
                                                                                        _ = v328
                                                                                    }
                                                                                }
                                                                                _ = v314
                                                                            case 9223372036854775807:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v6 = 9223372036854775807
                                                                                default:
                                                                                    var v329 string = "hello"
                                                                                    _ = v329
                                                                                }
                                                                            }
                                                                            fmt.Println(v47, v54, v54)
                                                                        } else {
                                                                            switch (v47 + v47) {
                                                                            case "":
                                                                                v3[int(v54)] = (v47 + (v47 + v47))
                                                                            case "a":
                                                                                v330 := v4[v47]
                                                                                _ = v330
                                                                            default:
                                                                                if (false || (1e10 < 0.0)) {
                                                                                    fmt.Println(v9)
                                                                                    switch 9223372036854775807 {
                                                                                    case 1:
                                                                                        _ = len(v4)
                                                                                    case 128:
                                                                                        v331 := GenS40{F0: 0}
                                                                                        _ = v331
                                                                                    case -9223372036854775808:
                                                                                        v332 := GenS5{F0: 62, F1: v47, F2: 18446744073709551615}
                                                                                        _ = v332
                                                                                    }
                                                                                    delete(v4, v47)
                                                                                }
                                                                            }
                                                                        }
                                                                        v1 += v1
                                                                    }
                                                                }
                                                                var v333 uint8 = 255
                                                                v334 := GenS7{F0: (v1 / ((42) | 1)), F1: (v6 + int64(v333))}
                                                                _ = v333
                                                                _ = v334
                                                            }
                                                            v23 = v23
                                                            _ = fi77
                                                            _ = v78
                                                        }
                                                    }
                                                }
                                                _ = fi59
                                            }
                                        default:
                                            fmt.Println(v47, v54, v23)
                                        }
                                        v22 -= (v22 / ((v22) | 1))
                                    }
                                    _ = v54
                                }
                                v2 -= ((1e10 / ((0.1 - (v2 * (((((v2 + v2) * (v2 * -2.25)) / v2) + v2) - v2))) / (v2 * (v2 * v2)))) * (v2 / v2))
                            }
                        } else {
                            v335 := []int{}
                            for fi336 := 0; fi336 < 7; fi336++ {
                                if true {
                                    v337 := GenS29{F0: -32768, F1: v22, F2: 1}
                                    if ((v6 * (v6 / (((v6 << ((2) & 63))) | 1))) > (-(v6))) {
                                        switch (-(fi336)) {
                                        case -189:
                                            v22 = v22
                                        default:
                                            var v338 uint16 = uint16(v9)
                                            _ = v338
                                        }
                                        var v339 int16 = 42
                                        _ = v339
                                    }
                                    fmt.Println(v22, v23, v22)
                                    _ = v337
                                } else {
                                    v9 -= v9
                                    v340 := GenS41{F0: 76}
                                    _ = v340
                                }
                                v2 = 0.0
                                _ = fi336
                            }
                            _ = v335
                        }
                        var v341 uint64 = (v9 | (v9 << ((547) & 63)))
                        _ = v341
                    }
                    var v342 int16 = -32768
                    _ = v342
                }
                v47 = ("" + v47)
                _ = v47
            }
        case "":
            var v343 string = ""
            _ = v343
        default:
            v344 := GenS24{F0: (v22 * -1), F1: (true || (true && (-9223372036854775808 <= (v6 * v6))))}
            _ = v344
        }
    } else {
        var v345 int = 128
        v1 -= 1
        _ = v345
    }
    var v346 uint16 = v1
    v347 := GenS42{F0: 79, F1: 1}
    v9 -= (^(v9))
    v348 := []int{-9223372036854775808}
    v6 = -411
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v346:", v346)
    fmt.Println("slicelen v348:", len(v348), v348)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v346
    _ = v347
    _ = v348
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
