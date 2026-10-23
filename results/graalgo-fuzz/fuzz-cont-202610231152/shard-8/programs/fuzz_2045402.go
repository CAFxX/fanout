package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 int64; F2 int8 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 bool; F1 uint64; F2 int }
type GenS13 struct { F0 int16; F1 int; F2 uint }
type GenS14 struct { F0 int8; F1 int16 }
type GenS15 struct { F0 int64; F1 int32; F2 uint64 }
type GenS16 struct { F0 int8; F1 uint32; F2 int8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int16; F1 int; F2 uint32 }
type GenS19 struct { F0 uint16; F1 bool }
type GenS20 struct { F0 uint16; F1 int; F2 uint64 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int; F1 uint8 }
type GenS24 struct { F0 int }
type GenS25 struct { F0 uint8 }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint8 = 255
        v6 = v6
        for fi7 := 0; fi7 < 6; fi7++ {
            switch (fi7 << ((2) & 63)) {
            case -2147483648:
                var v8 uint64 = 18446744073709551615
                _ = v8
            default:
                var v9 int32 = 1
                _ = v9
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    switch "a" {
    case "\n\t":
        var v10 int = 324
        _ = v10
    default:
        if true {
            switch "foo bar" {
            case "\n\t":
                var v11 uint8 = 42
                _ = v11
            case "foo bar":
                var v12 uint32 = 4294967295
                _ = v12
            default:
                switch 32768 {
                case 9223372036854775807:
                    var v13 int16 = -32768
                    _ = v13
                case 721:
                    var v14 string = ""
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        if (false && (!true)) {
                            if false {
                                var v16 uint64 = 261
                                v17 := []int{42, -2147483648}
                                _ = v16
                                _ = v17
                            }
                        }
                        if (65535 <= 185) {
                            if false {
                                v18 := GenS1{F0: 1}
                                fmt.Println(fi15)
                                switch ("\n\t" + "foo bar") {
                                case "hello":
                                    var v19 float32 = -2.25
                                    _ = v19
                                case "foo bar":
                                    var v20 uint32 = 2
                                    _ = v20
                                case "":
                                    var v21 uint32 = 891
                                    _ = v21
                                }
                                _ = v18
                            }
                            switch (("\n\t" + ("hello" + (("" + "\n\t") + "hello"))) + "foo bar") {
                            case "a":
                                fmt.Println(fi15)
                            case "foo bar":
                                if (0.0 != (1e10 + (float32(0.1) - 0.1))) {
                                    var v22 uint32 = 776
                                    switch fi15 {
                                    case 4294967295:
                                        v23 := GenS1{F0: -2147483648}
                                        _ = v23
                                    case 1:
                                        v24 := []int{(fi15 + fi15), fi15, (fi15 / (((^(fi15))) | 1))}
                                        _ = v24
                                    case -2147483648:
                                        switch fi15 {
                                        case 9223372036854775807:
                                            var v25 int16 = -740
                                            _ = v25
                                        }
                                    default:
                                        var v26 uint = 65535
                                        _ = v26
                                    }
                                    switch fi15 {
                                    case -1:
                                        v22 -= (^(v22))
                                    default:
                                        v22 += (v22 >> ((65535) & 63))
                                    }
                                    _ = v22
                                } else {
                                    for fi27 := 0; fi27 < 8; fi27++ {
                                        var v28 int32 = 2147483647
                                        var v29 uint = 50
                                        switch (fi15 << (((v29 + 255)) & 63)) {
                                        case -745:
                                            switch "hello" {
                                            case "":
                                                v30 := []int{-2147483648, 694, (-(fi27)), fi15}
                                                _ = v30
                                            case "hello":
                                                var v31 uint64 = 1
                                                _ = v31
                                            case "a":
                                                fmt.Println(fi15, fi15)
                                            }
                                            switch fi27 {
                                            case -768:
                                                if true {
                                                    v29 = v29
                                                    v32 := []int{(fi27 + fi15), (fi27 ^ (fi27 / (((fi15 >> (((^(v29))) & 63))) | 1))), -39, (fi15 - fi15)}
                                                    _ = v32
                                                } else {
                                                    v29 += v29
                                                    v28 += -1
                                                }
                                            case 128:
                                                v28 = v28
                                            case -271:
                                                switch fi15 {
                                                case 127:
                                                    v33 := GenS1{F0: (v28 / ((-2147483648) | 1))}
                                                    _ = v33
                                                case 0:
                                                    var v34 bool = true
                                                    _ = v34
                                                case 1:
                                                    v35 := GenS1{F0: (v28 * (v28 | (v28 / (((^(v28))) | 1))))}
                                                    _ = v35
                                                default:
                                                    v28 = v28
                                                }
                                            default:
                                                v36 := GenS1{F0: (v28 ^ v28)}
                                                _ = v36
                                            }
                                            var v37 string = (("hello" + "") + ("" + "foo bar"))
                                            var v38 int32 = v28
                                            v39 := GenS2{F0: "foo bar", F1: v38, F2: true}
                                            v29 = 531
                                            v28 -= (v38 + v28)
                                            switch fi27 {
                                            case 32767:
                                                v37 = (v37 + v37)
                                            case -9223372036854775808:
                                                fmt.Println(fi27, v38)
                                            case 42:
                                                v29 = 4294967295
                                            }
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        default:
                                            var v40 uint32 = 546
                                            _ = v40
                                        }
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                    fmt.Println(fi15)
                                    switch "\n\t" {
                                    case "hello":
                                        var v41 uint = 4294967295
                                        _ = v41
                                    }
                                }
                            }
                        }
                        for fi42 := 0; fi42 < 13; fi42++ {
                            v43 := GenS1{F0: 1}
                            var v44 int8 = 53
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        _ = fi15
                    }
                }
            }
            switch 1 {
            case 32767:
                var v45 bool = (2147483648 < 2)
                _ = v45
            }
            v46 := GenS2{F0: "a", F1: 1, F2: false}
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        v48 := GenS2{F0: ((("hello" + "a") + "foo bar") + ("hello" + ("\n\t" + ""))), F1: -366, F2: true}
        var v49 string = "hello"
        _ = fi47
        _ = v48
        _ = v49
    }
    return 4294967295
}
func genFn1(p0 uint16) uint64 {
    for fi50 := 0; fi50 < 3; fi50++ {
        for fi51 := 0; fi51 < 16; fi51++ {
            if (!(true || (656 > 1))) {
                fmt.Println(fi50)
            } else {
                v52 := []int{(fi50 << ((255) & 63)), (fi51 >> ((42) & 63)), 846}
                fmt.Println(p0, fi50)
                _ = v52
            }
            v53 := GenS2{F0: ("\n\t" + ((("" + ("hello" + "a")) + "a") + "")), F1: -126, F2: ("" == (((("foo bar" + "foo bar") + "a") + "a") + ("foo bar" + "foo bar")))}
            if (!false) {
                fmt.Println(fi51, fi50, p0)
                p0 = (p0 / ((p0) | 1))
            } else {
                var v54 int = (fi51 * (fi50 & fi50))
                var v55 uint16 = (^(p0))
                _ = v54
                _ = v55
            }
            _ = fi51
            _ = v53
        }
        var v56 int64 = -516
        _ = fi50
        _ = v56
    }
    var v57 uint64 = genFn0()
    _ = p0
    _ = v57
    return v57
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint16 = 0
    for fi58 := 0; fi58 < 13; fi58++ {
        var v59 uint64 = 18446744073709551615
        var v60 uint = 255
        _ = fi58
        _ = v59
        _ = v60
    }
    var v61 uint = 348
    if true {
        v62 := []int{869, 1}
        for fi63 := 0; fi63 < 10; fi63++ {
            switch (fi63 - fi63) {
            case 4294967295:
                var v64 int = 42
                _ = v64
            case 9223372036854775807:
                switch 32768 {
                case -993:
                    if (2 < 0) {
                        switch ("a" + ("a" + "foo bar")) {
                        case "foo bar":
                            if false {
                                if (127 >= 42) {
                                    _ = len(v62)
                                    switch "" {
                                    case "foo bar":
                                        if false {
                                            v1 += 1000
                                            if ((42 > 42) && (!(42 > 713))) {
                                                var v65 bool = (!false)
                                                for fi66 := 0; fi66 < 4; fi66++ {
                                                    var v67 uint16 = 65535
                                                    v68 := GenS3{F0: "a", F1: 2147483647}
                                                    _ = fi66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                if len(v62) >= 2 {
                                                    v62 = v62[1:len(v62)]
                                                }
                                                _ = v65
                                            } else {
                                                v1 -= -201
                                            }
                                        }
                                    case "a":
                                        var v69 uint16 = (^(v4))
                                        _ = v69
                                    }
                                } else {
                                    v61 = v61
                                }
                            } else {
                                v70 := GenS4{F0: 0}
                                v1 -= 0
                                _ = v70
                            }
                        case "\n\t":
                            v4 = (v4 & (v4 % ((v4) | 1)))
                        case "hello":
                            switch 128 {
                            case -843:
                                var v71 uint8 = 163
                                _ = v71
                            default:
                                var v72 bool = (884 != 0)
                                _ = v72
                            }
                        }
                        if ((!(!false)) || false) {
                            v62 = append(v62, (-(fi63)))
                            v73 := GenS1{F0: 42}
                            _ = v73
                        } else {
                            if true {
                                v1 += (v1 % ((755) | 1))
                            } else {
                                if (!true) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if true {
                                            if (!(!false)) {
                                                fmt.Println(v61, v4)
                                                v1 = (v1 % ((1000) | 1))
                                            }
                                            v74 := GenS2{F0: "", F1: -1, F2: ((^(fi63)) > fi63)}
                                            for fi75 := 0; fi75 < 1; fi75++ {
                                                if len(v62) > 0 {
                                                    v62[0] = (-(fi63))
                                                }
                                                v1 += -32768
                                                _ = len(v2)
                                                var v76 float32 = float32((0.1 / 0.1))
                                                _ = fi75
                                                _ = v76
                                            }
                                            _ = v74
                                        }
                                    case "hello":
                                        var v77 uint64 = 2
                                        _ = v77
                                    }
                                    var v78 int = (fi63 - fi63)
                                    switch 128 {
                                    case 42:
                                        var v79 int32 = -2147483648
                                        _ = v79
                                    case 4294967295:
                                        if (!(!false)) {
                                            if len(v62) >= 2 {
                                                v62 = v62[1:len(v62)]
                                            }
                                        } else {
                                            if (!false) {
                                                v62 = append(v62, 9223372036854775807)
                                            } else {
                                                switch v78 {
                                                case 42:
                                                    v80 := GenS3{F0: (("" + "hello") + "foo bar"), F1: 42}
                                                    _ = v80
                                                case 127:
                                                    for fi81 := 0; fi81 < 13; fi81++ {
                                                        if true {
                                                            if (false || true) {
                                                                _ = len(v2)
                                                            }
                                                            var v82 uint8 = 144
                                                            _ = v82
                                                        } else {
                                                            switch (^(fi81)) {
                                                            case -809:
                                                                fmt.Println(v4, v1, fi81)
                                                            default:
                                                                if false {
                                                                    if false {
                                                                        if (-9223372036854775808 <= 0) {
                                                                            v1 = v1
                                                                            if len(v62) > 0 {
                                                                                v62[0] = -244
                                                                            }
                                                                        }
                                                                    }
                                                                    v78 += (fi81 / ((v78) | 1))
                                                                    var v83 int64 = 42
                                                                    _ = v83
                                                                }
                                                                v4 += 255
                                                                if true {
                                                                    v62 = append(v62, 127)
                                                                    v61 = (v61 >> ((246) & 63))
                                                                    v84 := GenS5{F0: (119 <= 141), F1: 3}
                                                                    _ = v84
                                                                } else {
                                                                    v4 -= (v4 / (((v4 - (v4 % ((v4) | 1)))) | 1))
                                                                    if true {
                                                                        var v85 int = 0
                                                                        fmt.Println(v85)
                                                                        _ = v85
                                                                    } else {
                                                                        if false {
                                                                            switch "a" {
                                                                            case "a":
                                                                                switch "" {
                                                                                case "hello":
                                                                                    fmt.Println(v61, v78, v61)
                                                                                }
                                                                            }
                                                                            if false {
                                                                                v86 := GenS5{F0: (872 == 293), F1: 255}
                                                                                if ((1e10 != 0.1) || false) {
                                                                                    v78 = (v78 >> (((v61 * (v61 / ((v61) | 1)))) & 63))
                                                                                    var v87 float32 = (3.14159 * float32(0.1))
                                                                                    if false {
                                                                                        v62 = append(v62, (v78 | (fi63 | v78)))
                                                                                        v1 = (v1 << ((v61) & 63))
                                                                                    }
                                                                                    var v88 uint16 = v4
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                } else {
                                                                                    if false {
                                                                                        if true {
                                                                                            var v89 int16 = v1
                                                                                            var v90 float32 = (float32(3.14159) / 1.0)
                                                                                            _ = v89
                                                                                            _ = v90
                                                                                        }
                                                                                        var v91 uint32 = 204
                                                                                        v4 += v4
                                                                                        _ = v91
                                                                                    } else {
                                                                                        var v92 uint16 = 2
                                                                                        v93 := GenS2{F0: "\n\t", F1: 405, F2: true}
                                                                                        v78 -= v78
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                    v4 = v4
                                                                                }
                                                                                _ = v86
                                                                            } else {
                                                                                v1 += (v1 << ((255) & 63))
                                                                                var v94 int8 = -64
                                                                                v3[(^(fi63))] = 1
                                                                                _ = v94
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v3[fi63] = 157
                                                        }
                                                        v61 = (v61 + v61)
                                                        switch (^(v78)) {
                                                        case -9223372036854775808:
                                                            switch "a" {
                                                            case "":
                                                                fmt.Println(v4, fi63, v78)
                                                            case "hello":
                                                                v1 = 42
                                                            case "foo bar":
                                                                v95 := GenS5{F0: (!true), F1: 0}
                                                                _ = v95
                                                            default:
                                                                v4 = (v4 + (v4 + v4))
                                                            }
                                                        case 1:
                                                            v96 := GenS3{F0: "foo bar", F1: -2147483648}
                                                            _ = v96
                                                        case -905:
                                                            var v97 uint = 627
                                                            _ = v97
                                                        default:
                                                            v78 += (v78 & fi63)
                                                        }
                                                        _ = fi81
                                                    }
                                                default:
                                                    switch v78 {
                                                    case 634:
                                                        var v98 int16 = 1000
                                                        _ = v98
                                                    case 128:
                                                        if (483 <= 634) {
                                                            var v99 int64 = 0
                                                            fmt.Println(v1, v99)
                                                            fmt.Println(v1, v61, v4)
                                                            _ = v99
                                                        } else {
                                                            if len(v62) > 0 {
                                                                v62[0] = fi63
                                                            }
                                                        }
                                                    case 32768:
                                                        v4 = (v4 * (^(v4)))
                                                    default:
                                                        v1 -= v1
                                                    }
                                                }
                                                if false {
                                                    for fi100 := 0; fi100 < 10; fi100++ {
                                                        switch ("a" + "\n\t") {
                                                        case "":
                                                            v101 := GenS2{F0: "foo bar", F1: 2147483647, F2: (true && true)}
                                                            _ = v101
                                                        case "a":
                                                            var v102 float64 = 1.5
                                                            _ = v102
                                                        case "foo bar":
                                                            fmt.Println(fi100, v61, fi63)
                                                            v103 := GenS6{F0: 127, F1: "foo bar"}
                                                            fmt.Println(fi63)
                                                            _ = v103
                                                        }
                                                        v62 = append(v62, 405)
                                                        v78 = fi63
                                                        if true {
                                                            v104 := GenS5{F0: (!(0 > 1)), F1: 128}
                                                            var v105 float32 = 0.1
                                                            _ = v104
                                                            _ = v105
                                                        } else {
                                                            var v106 uint = v61
                                                            _ = v106
                                                        }
                                                        _ = fi100
                                                    }
                                                    v107 := GenS6{F0: 69, F1: ("foo bar" + ("a" + ""))}
                                                    _ = v107
                                                }
                                                v78 = 213
                                            }
                                        }
                                    default:
                                        v108 := v3[fi63]
                                        _ = v108
                                    }
                                    _ = v78
                                }
                                for fi109 := 0; fi109 < 16; fi109++ {
                                    var v110 uint8 = 2
                                    if true {
                                        var v111 uint32 = 2147483647
                                        v112 := GenS7{F0: -1}
                                        _ = v111
                                        _ = v112
                                    } else {
                                        fmt.Println(v61, v110)
                                    }
                                    _ = fi109
                                    _ = v110
                                }
                                if false {
                                    switch ("\n\t" + "hello") {
                                    case "a":
                                        v4 += (v4 + (v4 - (v4 & (v4 | 90))))
                                    case "hello":
                                        if false {
                                            fmt.Println(v4)
                                            var v113 uint = v61
                                            for fi114 := 0; fi114 < 15; fi114++ {
                                                var v115 int8 = -69
                                                _ = fi114
                                                _ = v115
                                            }
                                            _ = v113
                                        }
                                    case "":
                                        v62 = append(v62, (fi63 >> ((v61) & 63)))
                                    }
                                }
                            }
                        }
                    } else {
                        v1 = -831
                        fmt.Println(v61, v4, fi63)
                        switch ("foo bar" + "a") {
                        case "hello":
                            _ = len(v3)
                        case "foo bar":
                            var v116 int = 4294967295
                            _ = v116
                        case "":
                            v1 = v1
                        default:
                            for fi117 := 0; fi117 < 7; fi117++ {
                                v118 := GenS8{F0: (9223372036854775807 > int64(v1))}
                                fmt.Println(v61, fi117)
                                v119 := GenS9{F0: 128, F1: 9223372036854775807, F2: -128}
                                _ = fi117
                                _ = v118
                                _ = v119
                            }
                        }
                    }
                case -17:
                    switch "a" {
                    case "":
                        v61 = 4294967295
                    default:
                        v120 := GenS5{F0: (false && ((fi63 | fi63) >= -733)), F1: uint8(v61)}
                        _ = v120
                    }
                case 32768:
                    _ = len(v62)
                }
            default:
                for fi121 := 0; fi121 < 18; fi121++ {
                    switch ("" + "") {
                    case "":
                        if false {
                            v122 := GenS10{F0: -1, F1: 27}
                            if (!(int64(v4) == 0)) {
                                var v123 uint16 = v4
                                _ = v123
                            } else {
                                v4 -= (v4 % (((^(v4))) | 1))
                            }
                            if (!((v1 > v1) || (-1 <= -1))) {
                                var v124 uint16 = 700
                                var v125 int32 = int32(v124)
                                v126 := GenS11{F0: (v125 & v125), F1: 42}
                                fmt.Println(v4)
                                v61 = (v61 << ((v61) & 63))
                                _ = v124
                                _ = v125
                                _ = v126
                            }
                            _ = v122
                        } else {
                            fmt.Println(fi63)
                            fmt.Println(v4)
                            switch (("foo bar" + ("" + ("hello" + "hello"))) + "foo bar") {
                            case "\n\t":
                                var v127 int8 = -1
                                _ = v127
                            case "hello":
                                if ((v61 * 0) >= 447) {
                                    delete(v3, fi63)
                                    var v128 int8 = 10
                                    v129 := GenS12{F0: true, F1: 659, F2: 128}
                                    _ = v128
                                    _ = v129
                                }
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    v4 = (v4 & v4)
                                case "":
                                    fmt.Println(v1, fi63, fi63)
                                    if false {
                                        v1 += 980
                                        var v130 bool = true
                                        _ = v130
                                    } else {
                                        if (667 != 0) {
                                            v4 = (v4 & (^(v4)))
                                            v131 := GenS13{F0: 32767, F1: -9223372036854775808, F2: 2}
                                            _ = v131
                                        }
                                        v4 = v4
                                        v132 := GenS10{F0: -1, F1: (fi63 | (^(fi63)))}
                                        _ = v132
                                    }
                                case "\n\t":
                                    _ = len(v3)
                                default:
                                    v133 := GenS3{F0: "a", F1: 492}
                                    _ = v133
                                }
                            }
                        }
                    case "\n\t":
                        v134 := GenS14{F0: 42, F1: (v1 & (-(v1)))}
                        _ = v134
                    }
                    if true {
                        v61 = 1
                        var v135 uint8 = 1
                        _ = v135
                    } else {
                        var v136 uint32 = 2
                        _ = v136
                    }
                    _ = fi121
                }
            }
            if (("\n\t" == "") || (2 == 166)) {
                v61 += v61
                v3[fi63] = 4294967295
                v4 = v4
            }
            _ = fi63
        }
        _ = v62
    }
    if (!false) {
        v4 = 32767
        v1 -= 32767
        v137 := []int{4294967295}
        _ = v137
    } else {
        v1 = -312
        for fi138 := 0; fi138 < 2; fi138++ {
            v4 -= uint16(v61)
            fmt.Println(fi138, fi138)
            fmt.Println(fi138, v4)
            _ = fi138
        }
        if true {
            v1 = v1
        } else {
            v139 := GenS3{F0: "foo bar", F1: 42}
            if true {
                _ = len(v3)
            }
            _ = v139
        }
    }
    if ("hello" > (("\n\t" + "") + ("hello" + "hello"))) {
        var v140 int32 = -987
        switch "\n\t" {
        case "a":
            var v141 uint32 = 2147483648
            _ = v141
        default:
            var v142 int64 = 0
            _ = v142
        }
        if false {
            if false {
                switch "hello" {
                case "foo bar":
                    v61 += v61
                case "hello":
                    if (-871 == v1) {
                        fmt.Println(v1)
                        v143 := []int{}
                        for fi144 := 0; fi144 < 11; fi144++ {
                            v4 = (v4 * (v4 / (((^(v4))) | 1)))
                            _ = fi144
                        }
                        _ = v143
                    }
                case "a":
                    if (!true) {
                        var v145 int = 4294967295
                        _ = v145
                    }
                default:
                    if (0 <= v1) {
                        fmt.Println(v1, v1)
                        v140 += v140
                    }
                }
                var v146 uint8 = 127
                var v147 uint16 = (v4 * (^(v4)))
                _ = v146
                _ = v147
            }
        }
        _ = v140
    } else {
        v148 := []int{465, -221}
        switch "a" {
        case "hello":
            var v149 string = ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + ("\n\t" + ((("" + "hello") + (("hello" + "hello") + "")) + "")))
            _ = v149
        case "":
            v61 = 4294967295
        default:
            switch "" {
            case "a":
                for fi150 := 0; fi150 < 8; fi150++ {
                    if (!true) {
                        if true {
                            if true {
                                if (1.5 < 1.5) {
                                    if (9223372036854775807 != 1) {
                                        v151 := GenS9{F0: 1, F1: 872, F2: -1}
                                        _ = v151
                                    }
                                    var v152 uint32 = 42
                                    _ = v152
                                } else {
                                    v153 := GenS7{F0: 42}
                                    v61 += 65535
                                    _ = v153
                                }
                                var v154 int64 = 265
                                var v155 uint8 = 2
                                _ = v154
                                _ = v155
                            }
                            for fi156 := 0; fi156 < 9; fi156++ {
                                v157 := GenS7{F0: 1}
                                fmt.Println(v4, fi150)
                                if (1.5 < (0.0 / -2.25)) {
                                    v61 += 761
                                    if (-1 != 35) {
                                        v4 = v4
                                        if len(v148) > 0 {
                                            v148[0] = fi150
                                        }
                                        v158 := GenS3{F0: "foo bar", F1: 133}
                                        if (51 >= -128) {
                                            switch ("a" + "a") {
                                            case "hello":
                                                fmt.Println(v4, fi150, v61)
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    if true {
                                                        if ((v61 / (((v61 % ((v61) | 1))) | 1)) == v61) {
                                                            v2[fi156] = fi156
                                                        }
                                                        v1 = -32768
                                                        if len(v148) > 0 {
                                                            v148[0] = fi150
                                                        }
                                                    }
                                                case "foo bar":
                                                    _ = len(v3)
                                                case "\n\t":
                                                    switch "a" {
                                                    case "":
                                                        switch "hello" {
                                                        case "hello":
                                                            var v159 uint16 = v4
                                                            _ = v159
                                                        case "a":
                                                            var v160 uint = v61
                                                            _ = v160
                                                        }
                                                    default:
                                                        v148 = append(v148, fi156)
                                                        v161 := GenS14{F0: -128, F1: -32768}
                                                        switch "" {
                                                        case "":
                                                            v4 -= 993
                                                        case "hello":
                                                            if false {
                                                                var v162 bool = false
                                                                v148 = append(v148, (fi150 >> ((v61) & 63)))
                                                                _ = v162
                                                            }
                                                        }
                                                        _ = v161
                                                    }
                                                }
                                            }
                                            fmt.Println(fi150)
                                            if len(v148) > 0 {
                                                v148[0] = 1
                                            }
                                        }
                                        if len(v148) >= 2 {
                                            v148 = v148[1:len(v148)]
                                        }
                                        _ = v158
                                    }
                                }
                                _ = fi156
                                _ = v157
                            }
                            fmt.Println(v61)
                        }
                        var v163 string = "hello"
                        _ = v163
                    } else {
                        for fi164 := 0; fi164 < 8; fi164++ {
                            if false {
                                if (int64(fi164) != 0) {
                                    v165 := GenS15{F0: -171, F1: 329, F2: 2}
                                    v166 := v2[fi164]
                                    _ = v165
                                    _ = v166
                                }
                                if ((!false) || (!false)) {
                                    switch (("\n\t" + "hello") + ("a" + ("\n\t" + (((("\n\t" + "foo bar") + ("" + "foo bar")) + ("hello" + ("" + "\n\t"))) + "")))) {
                                    case "\n\t":
                                        fmt.Println(v1)
                                        if true {
                                            switch -2147483648 {
                                            case 2147483647:
                                                delete(v3, (fi150 % (((fi164 << ((v61) & 63))) | 1)))
                                            case 326:
                                                v167 := v3[(fi150 ^ (-(fi164)))]
                                                v61 = v61
                                                _ = v167
                                            }
                                            _ = len(v148)
                                            v4 = 30
                                            v168 := GenS16{F0: -86, F1: 65535, F2: 102}
                                            delete(v3, 42)
                                            _ = v168
                                        }
                                    }
                                    var v169 uint = v61
                                    _ = v169
                                }
                            }
                            _ = fi164
                        }
                    }
                    v170 := GenS12{F0: false, F1: 42, F2: fi150}
                    _ = fi150
                    _ = v170
                }
            }
        }
        _ = v148
    }
    fmt.Println(v61, v1)
    fmt.Println(v4, v61, v4)
    var v171 int = 0
    var v172 uint16 = 255
    var v173 bool = false
    for fi174 := 0; fi174 < 9; fi174++ {
        delete(v2, (fi174 << ((594) & 63)))
        if v173 {
            fmt.Println(v172)
            fmt.Println(v1, fi174)
        } else {
            v1 += (v1 & 32767)
            v1 -= (v1 + (v1 / (((-(v1))) | 1)))
        }
        _ = fi174
    }
    switch (v171 << ((v61) & 63)) {
    case 4294967295:
        if v173 {
            if v173 {
                switch "hello" {
                case "hello":
                    for fi175 := 0; fi175 < 4; fi175++ {
                        var v176 uint32 = uint32(v4)
                        _ = fi175
                        _ = v176
                    }
                case "\n\t":
                    fmt.Println(v171, v173)
                }
                var v177 string = "\n\t"
                var v178 uint8 = 102
                _ = v177
                _ = v178
            } else {
                var v179 uint8 = 255
                _ = v179
            }
            if false {
                if v173 {
                    for fi180 := 0; fi180 < 14; fi180++ {
                        switch "\n\t" {
                        case "\n\t":
                            v173 = false
                        case "a":
                            if v173 {
                                delete(v2, v171)
                                if (1e10 <= (((3.14159 / 1.0) * (3.14159 - 0.0)) * ((float32(0.1) * 3.14159) * 1e10))) {
                                    fmt.Println(v173, fi180)
                                    var v181 string = "foo bar"
                                    _ = v181
                                }
                                switch "hello" {
                                case "":
                                    var v182 int64 = -396
                                    _ = v182
                                case "foo bar":
                                    for fi183 := 0; fi183 < 18; fi183++ {
                                        var v184 uint32 = 65535
                                        var v185 uint64 = 56
                                        _ = fi183
                                        _ = v184
                                        _ = v185
                                    }
                                default:
                                    fmt.Println(v171)
                                }
                            }
                        case "foo bar":
                            if ((v171 ^ (fi180 << (((^(v61))) & 63))) == 32768) {
                                var v186 int64 = -921
                                var v187 uint64 = 0
                                _ = v186
                                _ = v187
                            }
                        }
                        _ = fi180
                    }
                }
                v171 = (v171 - (v171 + -9223372036854775808))
                v4 -= v172
            } else {
                v171 = v171
                fmt.Println(v171, v61)
                if false {
                    if (766 != -1) {
                        v171 += (v171 - 2147483647)
                    } else {
                        if false {
                            v171 = v171
                        } else {
                            for fi188 := 0; fi188 < 4; fi188++ {
                                v61 = (v61 % (((v61 % (((v61 + 18446744073709551615)) | 1))) | 1))
                                if v173 {
                                    v173 = v173
                                }
                                _ = fi188
                            }
                        }
                        v189 := GenS13{F0: v1, F1: -812, F2: (^(v61))}
                        _ = v189
                    }
                }
            }
            if (4294967295 < 22) {
                if false {
                    v190 := GenS13{F0: v1, F1: v171, F2: (^(v61))}
                    v191 := GenS16{F0: -1, F1: 175, F2: -62}
                    v61 = (v61 >> ((v61) & 63))
                    _ = v190
                    _ = v191
                } else {
                    v192 := []int{(v171 / ((v171) | 1)), (-(v171)), v171}
                    _ = v192
                }
                v193 := GenS16{F0: -121, F1: 42, F2: -128}
                switch "foo bar" {
                case "\n\t":
                    for fi194 := 0; fi194 < 17; fi194++ {
                        var v195 uint64 = genFn1(v172)
                        for fi196 := 0; fi196 < 8; fi196++ {
                            if ((v172 / (((v4 >> ((uint(v4)) & 63))) | 1)) <= (v4 / ((0) | 1))) {
                                if ((932 < (v61 / (((v61 / (((v61 ^ (v61 ^ (v61 & 4294967295)))) | 1))) | 1))) && (v173 && ((v171 - (fi196 + fi194)) < fi194))) {
                                    var v197 float64 = 0.1
                                    v195 += v195
                                    switch (fi194 & v171) {
                                    case 32768:
                                        var v198 int = fi196
                                        _ = v198
                                    case -708:
                                        v61 = 1
                                    default:
                                        v173 = ("a" >= "\n\t")
                                    }
                                    v171 = -19
                                    if (9223372036854775807 > int64(v61)) {
                                        v1 += (^(v1))
                                        var v199 bool = (18446744073709551615 <= v195)
                                        _ = v199
                                    }
                                    _ = v197
                                } else {
                                    if true {
                                        v173 = (42 != 42)
                                    }
                                }
                                if (!((v173 || (!v173)) || true)) {
                                    v172 += (v4 << ((255) & 63))
                                    v200 := GenS17{F0: 532}
                                    _ = v200
                                }
                            }
                            var v201 uint = 18446744073709551615
                            if true {
                                v202 := GenS18{F0: (-(v1)), F1: (fi196 << ((v61) & 63)), F2: 1}
                                v172 = v4
                                _ = v202
                            }
                            _ = fi196
                            _ = v201
                        }
                        for fi203 := 0; fi203 < 11; fi203++ {
                            var v204 float32 = -2.25
                            var v205 int32 = 2147483647
                            v195 = (v195 | (^(v195)))
                            fmt.Println(v173, v173, v1)
                            _ = fi203
                            _ = v204
                            _ = v205
                        }
                        _ = fi194
                        _ = v195
                    }
                }
                _ = v193
            }
        } else {
            var v206 uint16 = 2
            if (v171 <= (v171 ^ 4294967295)) {
                v207 := v2[(^(v171))]
                for fi208 := 0; fi208 < 19; fi208++ {
                    for fi209 := 0; fi209 < 14; fi209++ {
                        switch (-(v207)) {
                        case 128:
                            fmt.Println(v171, v1)
                        case 127:
                            fmt.Println(v171, v207, v61)
                        case 4294967295:
                            v207 = 9223372036854775807
                        default:
                            v210 := v2[(fi209 / (((fi209 % (((fi208 / (((fi209 >> (((v61 % (((v61 % (((v61 % ((v61) | 1))) | 1))) | 1))) & 63))) | 1))) | 1))) | 1))]
                            _ = v210
                        }
                        switch "hello" {
                        case "":
                            _ = len(v2)
                        default:
                            var v211 uint64 = 42
                            v212 := v2[128]
                            v213 := GenS5{F0: ((int32(v171) == -655) || v173), F1: 128}
                            switch "hello" {
                            case "a":
                                switch (v212 * -9223372036854775808) {
                                case 2147483647:
                                    fmt.Println(v171, v1, v211)
                                }
                            case "foo bar":
                                switch "foo bar" {
                                case "":
                                    if v173 {
                                        if true {
                                            var v214 bool = (false && (v173 && (false && v173)))
                                            _ = v214
                                        }
                                        var v215 uint = 151
                                        var v216 int32 = -1
                                        _ = v215
                                        _ = v216
                                    }
                                case "foo bar":
                                    if false {
                                        v61 = (v61 * v61)
                                        if ((!(v1 < (v1 - 32767))) || (((^(v1)) > -32768) && (986 == -2147483648))) {
                                            var v217 int = 4294967295
                                            v218 := GenS18{F0: (v1 % ((v1) | 1)), F1: v212, F2: 4294967295}
                                            var v219 uint16 = 255
                                            _ = v217
                                            _ = v218
                                            _ = v219
                                        } else {
                                            v212 = fi209
                                            var v220 float64 = 0.1
                                            _ = v220
                                        }
                                    }
                                default:
                                    if true {
                                        v221 := GenS11{F0: -9, F1: v1}
                                        _ = v221
                                    }
                                }
                            case "\n\t":
                                _ = len(v2)
                                v4 = 32767
                                if true {
                                    v222 := GenS1{F0: -94}
                                    _ = v222
                                } else {
                                    if (!(!(0 < 802))) {
                                        if v173 {
                                            v4 -= v4
                                            if v173 {
                                                var v223 bool = v173
                                                v171 -= v212
                                                v224 := []int{v207, (-(v212)), fi209, v207}
                                                _ = v223
                                                _ = v224
                                            } else {
                                                switch ("\n\t" + ("hello" + "hello")) {
                                                case "a":
                                                    v225 := []int{9223372036854775807, (v171 << ((v61) & 63))}
                                                    _ = v225
                                                case "":
                                                    var v226 uint32 = 2147483647
                                                    _ = v226
                                                case "hello":
                                                    _ = len(v2)
                                                }
                                                v227 := GenS19{F0: (v172 | (v4 >> ((v61) & 63))), F1: (!v173)}
                                                _ = v227
                                            }
                                            if (255 == 128) {
                                                var v228 float64 = 1e10
                                                var v229 uint32 = 4294967295
                                                _ = v228
                                                _ = v229
                                            } else {
                                                var v230 bool = v173
                                                _ = v230
                                            }
                                        } else {
                                            v206 += 2
                                            fmt.Println(v211)
                                            var v231 int64 = -1
                                            _ = v231
                                        }
                                    } else {
                                        fmt.Println(v4)
                                        v232 := []int{}
                                        _ = v232
                                    }
                                }
                                if true {
                                    v233 := v2[v207]
                                    var v234 float64 = 0.1
                                    v235 := GenS13{F0: v1, F1: (fi209 << ((42) & 63)), F2: v61}
                                    _ = v233
                                    _ = v234
                                    _ = v235
                                } else {
                                    delete(v3, 128)
                                    v3[(^(v212))] = 2
                                    var v236 uint64 = v211
                                    _ = v236
                                }
                                if false {
                                    var v237 float64 = 0.1
                                    _ = v237
                                } else {
                                    if false {
                                        v211 = (v211 >> (((v61 - (^(v61)))) & 63))
                                        v173 = v173
                                        v238 := GenS20{F0: v172, F1: fi209, F2: (v211 + v211)}
                                        _ = v238
                                    }
                                    var v239 int16 = 416
                                    v240 := []int{}
                                    _ = v239
                                    _ = v240
                                }
                            }
                            var v241 uint8 = 128
                            _ = v211
                            _ = v212
                            _ = v213
                            _ = v241
                        }
                        v242 := GenS21{F0: v61}
                        _ = fi209
                        _ = v242
                    }
                    _ = fi208
                }
                _ = v207
            }
            var v243 uint32 = 65535
            _ = v206
            _ = v243
        }
    default:
        var v244 uint = (v61 ^ 18446744073709551615)
        _ = v244
    }
    var v245 int16 = (v1 % ((0) | 1))
    v246 := []int{4294967295, v171, v171}
    v1 = (v1 >> ((4294967295) & 63))
    if (42 >= 0) {
        switch ("\n\t" + "hello") {
        case "hello":
            for fi247 := 0; fi247 < 12; fi247++ {
                if false {
                    for fi248 := 0; fi248 < 6; fi248++ {
                        var v249 int16 = (v1 & (-(v1)))
                        var v250 int32 = 2147483647
                        v4 = (v172 ^ v4)
                        _ = fi248
                        _ = v249
                        _ = v250
                    }
                }
                switch v171 {
                case 2147483647:
                    if (0 > 965) {
                        if ((-(v245)) >= v1) {
                            fmt.Println(v1, v171)
                            if len(v246) > 0 {
                                v246[0] = v171
                            }
                            var v251 string = "\n\t"
                            _ = v251
                        } else {
                            var v252 string = "foo bar"
                            _ = len(v246)
                            v253 := GenS14{F0: -1, F1: (v1 | v1)}
                            _ = v252
                            _ = v253
                        }
                        if v173 {
                            switch (fi247 + v171) {
                            case 0:
                                var v254 float64 = (((0.1 + 1.5) + 1.5) / 3.14159)
                                _ = v254
                            case 9223372036854775807:
                                v255 := GenS3{F0: "a", F1: -1}
                                _ = v255
                            default:
                                switch "hello" {
                                case "a":
                                    for fi256 := 0; fi256 < 6; fi256++ {
                                        var v257 int = (fi247 << ((167) & 63))
                                        _ = fi256
                                        _ = v257
                                    }
                                case "\n\t":
                                    var v258 int64 = -9223372036854775808
                                    _ = v258
                                default:
                                    var v259 int32 = -2147483648
                                    _ = v259
                                }
                            }
                            v246 = append(v246, (fi247 >> (((v61 | v61)) & 63)))
                        } else {
                            var v260 string = "foo bar"
                            _ = v260
                        }
                        for fi261 := 0; fi261 < 7; fi261++ {
                            var v262 uint8 = 2
                            v61 += (^(v61))
                            v263 := GenS22{F0: "hello"}
                            var v264 uint16 = v172
                            if (!v173) {
                                var v265 bool = (v61 <= (v61 & (v61 - 482)))
                                v266 := GenS20{F0: v264, F1: fi247, F2: 837}
                                var v267 int = fi247
                                _ = v265
                                _ = v266
                                _ = v267
                            } else {
                                v268 := GenS23{F0: 42, F1: v262}
                                var v269 int8 = -128
                                if len(v246) >= 2 {
                                    v246 = v246[1:len(v246)]
                                }
                                if (!false) {
                                    var v270 int32 = 1
                                    v1 += v1
                                    v270 = v270
                                    _ = v270
                                } else {
                                    if true {
                                        if (!v173) {
                                            v171 -= -2147483648
                                            v3[(fi247 & v171)] = 42
                                            v245 -= (-(v1))
                                            switch fi261 {
                                            case 2147483647:
                                                _ = len(v246)
                                            case 245:
                                                var v271 float64 = 0.0
                                                _ = v271
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    if len(v246) > 0 {
                                                        v246[0] = (fi247 * fi247)
                                                    }
                                                case "hello":
                                                    if true {
                                                        v272 := GenS6{F0: (v269 & -18), F1: "a"}
                                                        v273 := GenS13{F0: v245, F1: -9223372036854775808, F2: 256}
                                                        _ = v272
                                                        _ = v273
                                                    }
                                                case "":
                                                    v246 = append(v246, 164)
                                                }
                                            }
                                        }
                                        v4 = v4
                                        switch "foo bar" {
                                        case "":
                                            var v274 bool = (!v173)
                                            _ = v274
                                        case "foo bar":
                                            var v275 int8 = v269
                                            _ = v275
                                        }
                                        if v173 {
                                            v172 = (^(v172))
                                            if len(v246) >= 2 {
                                                v246 = v246[1:len(v246)]
                                            }
                                            v276 := GenS9{F0: (^(v262)), F1: -9223372036854775808, F2: v269}
                                            _ = v276
                                        } else {
                                            v277 := GenS24{F0: (fi261 / (((v171 | v171)) | 1))}
                                            v245 = v1
                                            _ = v277
                                        }
                                    }
                                    v172 = (v264 ^ (v264 | uint16(v61)))
                                    v1 += v1
                                }
                                var v278 uint16 = (v4 / ((1) | 1))
                                _ = v268
                                _ = v269
                                _ = v278
                            }
                            _ = fi261
                            _ = v262
                            _ = v263
                            _ = v264
                        }
                    } else {
                        var v279 int8 = -69
                        var v280 int16 = v1
                        v281 := GenS25{F0: 2}
                        _ = v279
                        _ = v280
                        _ = v281
                    }
                case 128:
                    var v282 uint16 = v172
                    _ = v282
                case 42:
                    v283 := GenS17{F0: 165}
                    _ = v283
                default:
                    for fi284 := 0; fi284 < 1; fi284++ {
                        v4 = v4
                        _ = fi284
                    }
                }
                _ = fi247
            }
        case "\n\t":
            v285 := GenS16{F0: 127, F1: 4294967295, F2: 0}
            _ = v285
        }
    } else {
        v245 -= v245
        if v173 {
            v4 = v4
        }
        v245 -= v1
    }
    var v286 uint8 = 76
    for fi287 := 0; fi287 < 2; fi287++ {
        v288 := v3[127]
        v172 = v172
        _ = fi287
        _ = v288
    }
    fmt.Println("v1:", v1)
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v245:", v245)
    fmt.Println("slicelen v246:", len(v246), v246)
    fmt.Println("v286:", v286)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v171
    _ = v172
    _ = v173
    _ = v2
    _ = v245
    _ = v246
    _ = v286
    _ = v3
    _ = v4
    _ = v61
}
