package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 int64; F2 int8 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int16; F1 int64 }
type GenS16 struct { F0 int8; F1 int8 }
type GenS17 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS18 struct { F0 int8; F1 uint64 }
type GenS19 struct { F0 uint8; F1 uint32 }
func genFn0() uint {
    v6 := GenS1{F0: 764}
    if (0.1 <= (3.14159 + 3.14159)) {
        var v7 int32 = 1
        v7 = v7
        _ = v7
    } else {
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(fi8)
            if true {
                var v9 int64 = -1
                fmt.Println(fi8)
                _ = v9
            }
            _ = fi8
        }
        v10 := []int{32767, 835, -9223372036854775808, -9223372036854775808}
        _ = v10
    }
    _ = v6
    return 65535
}
func genFn1(p0 int, p1 uint16) bool {
    p0 = -2147483648
    fmt.Println(p1, p1)
    p1 = p1
    p0 -= p0
    _ = p0
    _ = p1
    return (-293 > 1000)
}

func main() {
    v1 := []int{215, -9223372036854775808}
    var v2 uint = 65535
    v3 := make(map[string]bool)
    var v4 uint8 = 40
    var v5 int64 = -590
    _ = len(v1)
    fmt.Println(v5)
    v11 := GenS2{F0: -1, F1: 71, F2: 160}
    v4 = v4
    if true {
        v1 = append(v1, 32768)
        var v12 float32 = 1e10
        _ = v12
    } else {
        v4 = (v4 + (v4 % ((v4) | 1)))
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 string = "\n\t"
            _ = fi13
            _ = v14
        }
    }
    var v15 uint8 = v4
    if false {
        switch -222 {
        case -532:
            if false {
                var v16 int = 32768
                switch (("a" + ("a" + "hello")) + "a") {
                case "a":
                    switch v16 {
                    case 0:
                        switch 128 {
                        case 1:
                            fmt.Println(v15)
                        default:
                            switch "hello" {
                            case "foo bar":
                                v17 := GenS3{F0: (^(v5))}
                                _ = v17
                            case "\n\t":
                                v16 = (v16 - 128)
                            case "hello":
                                var v18 string = "hello"
                                _ = v18
                            }
                        }
                    default:
                        v19 := GenS4{F0: (v16 | (v16 + v16))}
                        _ = v19
                    }
                case "\n\t":
                    v5 += v5
                case "foo bar":
                    v20 := GenS2{F0: -1, F1: 127, F2: 65535}
                    _ = v20
                }
                _ = v16
            } else {
                v21 := GenS5{F0: v4, F1: v2, F2: 4294967295}
                if true {
                    if true {
                        if (92 > 42) {
                            v22 := GenS6{F0: 42}
                            v15 = (v4 ^ (v15 * v4))
                            _ = v22
                        } else {
                            v2 -= (v2 / ((2) | 1))
                            if (128 < 4294967295) {
                                v5 += v5
                            }
                            var v23 bool = true
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 19; fi24++ {
                            v5 += (v5 + (v5 / ((9223372036854775807) | 1)))
                            for fi25 := 0; fi25 < 17; fi25++ {
                                var v26 int32 = 0
                                var v27 uint64 = 0
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            v28 := GenS5{F0: v4, F1: (v2 % ((v2) | 1)), F2: fi24}
                            _ = fi24
                            _ = v28
                        }
                        var v29 float64 = 0.1
                        _ = v29
                    } else {
                        v30 := GenS7{F0: false}
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v2 += v2
                        v32 := v3[("\n\t" + ("a" + (((((("\n\t" + ("foo bar" + "")) + "foo bar") + "\n\t") + "foo bar") + "a") + "hello")))]
                        for fi33 := 0; fi33 < 5; fi33++ {
                            switch (("\n\t" + ("a" + "\n\t")) + ("foo bar" + "foo bar")) {
                            case "":
                                fmt.Println(v2, fi31)
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v34 uint64 = 42
                            switch (fi33 >> (((v2 * genFn0())) & 63)) {
                            case -891:
                                delete(v3, "hello")
                            default:
                                v35 := v3["hello"]
                                _ = v35
                            }
                            _ = fi33
                            _ = v34
                        }
                        _ = fi31
                        _ = v32
                    }
                }
                switch "" {
                case "":
                    switch -1 {
                    case 1:
                        _ = len(v3)
                    case -1:
                        if true {
                            var v36 int = -856
                            var v37 float32 = ((0.0 + -2.25) - 0.0)
                            v38 := GenS8{F0: v36, F1: -128}
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            var v39 int16 = 42
                            if false {
                                var v40 float32 = 1e10
                                _ = v40
                            } else {
                                for fi41 := 0; fi41 < 8; fi41++ {
                                    if len(v1) > 0 {
                                        v1[0] = (fi41 & fi41)
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        v42 := GenS2{F0: 42, F1: 32, F2: 65535}
                                        _ = v42
                                    case "\n\t":
                                        v43 := GenS6{F0: 284}
                                        _ = v43
                                    case "a":
                                        if false {
                                            v44 := GenS6{F0: 2147483647}
                                            for fi45 := 0; fi45 < 2; fi45++ {
                                                v3["a"] = (2147483647 == 75)
                                                v46 := GenS9{F0: 4294967295, F1: 2, F2: 732}
                                                _ = fi45
                                                _ = v46
                                            }
                                            _ = v44
                                        }
                                    }
                                    _ = fi41
                                }
                            }
                            var v47 uint = v2
                            _ = v39
                            _ = v47
                        }
                    case -2147483648:
                        v3["a"] = (!(uint16(v2) <= 65535))
                    default:
                        fmt.Println(v2, v5)
                    }
                default:
                    if (18446744073709551615 < 299) {
                        var v48 float64 = -2.25
                        _ = v48
                    } else {
                        v15 = (v15 / ((197) | 1))
                        if (true && (0 != -885)) {
                            v49 := GenS10{F0: 65535, F1: 283}
                            _ = v49
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch 9223372036854775807 {
                            case 32767:
                                if (109 <= 448) {
                                    switch ("" + "a") {
                                    case "foo bar":
                                        var v50 bool = ((v2 >> ((v2) & 63)) >= (v2 + (v2 ^ uint(v4))))
                                        _ = v50
                                    case "hello":
                                        v2 += (v2 ^ v2)
                                    }
                                    if true {
                                        if (!(61 == uint16(v2))) {
                                            for fi51 := 0; fi51 < 7; fi51++ {
                                                for fi52 := 0; fi52 < 12; fi52++ {
                                                    var v53 int32 = -497
                                                    v54 := GenS4{F0: 2147483647}
                                                    var v55 int32 = (^(v53))
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                }
                                                _ = fi51
                                            }
                                            switch ("" + ("\n\t" + "a")) {
                                            case "hello":
                                                if (!((!false) && (!false))) {
                                                    if false {
                                                        var v56 uint64 = 0
                                                        fmt.Println(v15, v4)
                                                        var v57 string = "a"
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        var v58 uint8 = v15
                                                        if (2 == 464) {
                                                            switch "" {
                                                            case "foo bar":
                                                                for fi59 := 0; fi59 < 13; fi59++ {
                                                                    switch (fi59 ^ 0) {
                                                                    case 32767:
                                                                        v15 -= (v58 | (^(v15)))
                                                                    }
                                                                    _ = fi59
                                                                }
                                                            case "\n\t":
                                                                for fi60 := 0; fi60 < 17; fi60++ {
                                                                    var v61 int64 = (v5 ^ (v5 - (^(v5))))
                                                                    _ = fi60
                                                                    _ = v61
                                                                }
                                                            case "":
                                                                switch 2147483647 {
                                                                case 946:
                                                                    var v62 float32 = (0.0 * 0.0)
                                                                    _ = v62
                                                                }
                                                            }
                                                        }
                                                        v4 -= (^(v15))
                                                        _ = v58
                                                    }
                                                }
                                            case "foo bar":
                                                v4 -= v15
                                            default:
                                                v5 = (v5 + v5)
                                            }
                                        }
                                        _ = len(v3)
                                    } else {
                                        var v63 uint64 = 169
                                        for fi64 := 0; fi64 < 18; fi64++ {
                                            switch (^(fi64)) {
                                            case 42:
                                                switch (fi64 >> (((v2 & (v2 | (v2 ^ (v2 % (((v2 & (v2 / ((18446744073709551615) | 1)))) | 1)))))) & 63)) {
                                                case -1:
                                                    var v65 int32 = 0
                                                    _ = v65
                                                case 9223372036854775807:
                                                    if false {
                                                        for fi66 := 0; fi66 < 3; fi66++ {
                                                            v4 += (v4 - v15)
                                                            delete(v3, ("a" + "a"))
                                                            v67 := GenS11{F0: (^(v63))}
                                                            if false {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v68 bool = (!false)
                                                                v69 := GenS8{F0: fi66, F1: 96}
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = fi66
                                                            _ = v67
                                                        }
                                                        v5 = (v5 * (^(v5)))
                                                        switch (fi64 & fi64) {
                                                        case 32767:
                                                            fmt.Println(v2)
                                                        case -60:
                                                            if (0 < -894) {
                                                                var v70 uint32 = uint32(v63)
                                                                for fi71 := 0; fi71 < 7; fi71++ {
                                                                    v72 := GenS5{F0: (v15 * (v15 ^ (v4 & v15))), F1: (v2 + v2), F2: fi64}
                                                                    var v73 uint = (v2 >> ((4294967295) & 63))
                                                                    if true {
                                                                        v1 = append(v1, (fi64 % ((0) | 1)))
                                                                        v4 = (^(v15))
                                                                    }
                                                                    _ = fi71
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v70
                                                            } else {
                                                                _ = len(v1)
                                                                if false {
                                                                    v4 += v15
                                                                    if true {
                                                                        if (-1 > 0) {
                                                                            fmt.Println(v4)
                                                                            fmt.Println(v63, v15, fi64)
                                                                            var v74 uint32 = 2147483647
                                                                            _ = v74
                                                                        }
                                                                        fmt.Println(v5, v4)
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            if true {
                                                                                v2 += genFn0()
                                                                                for fi75 := 0; fi75 < 20; fi75++ {
                                                                                    v4 = (^(v15))
                                                                                    _ = fi75
                                                                                }
                                                                            }
                                                                        case "":
                                                                            var v76 float64 = 3.14159
                                                                            _ = v76
                                                                        default:
                                                                            if (18446744073709551615 <= uint(v15)) {
                                                                                var v77 float64 = 0.1
                                                                                v5 += v5
                                                                                _ = v77
                                                                            } else {
                                                                                for fi78 := 0; fi78 < 19; fi78++ {
                                                                                    v79 := GenS8{F0: fi78, F1: 1}
                                                                                    var v80 uint8 = 2
                                                                                    v4 = v80
                                                                                    _ = fi78
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                }
                                                                                for fi81 := 0; fi81 < 18; fi81++ {
                                                                                    switch ("a" + "\n\t") {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            switch (fi64 / (((-(fi64))) | 1)) {
                                                                                            case 4294967295:
                                                                                                v63 -= 775
                                                                                            case 702:
                                                                                                fmt.Println(fi64)
                                                                                                if false {
                                                                                                    v82 := GenS12{F0: fi64}
                                                                                                    var v83 uint64 = (v63 >> ((uint(v63)) & 63))
                                                                                                    _ = v82
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    v84 := GenS13{F0: 965, F1: v5, F2: -128}
                                                                                                    if true {
                                                                                                        v1 = append(v1, fi64)
                                                                                                    } else {
                                                                                                        v85 := GenS14{F0: v63}
                                                                                                        fmt.Println(v15, v4, v5)
                                                                                                        _ = len(v3)
                                                                                                        _ = len(v1)
                                                                                                        v4 = (^(v4))
                                                                                                        if (!false) {
                                                                                                            v86 := GenS12{F0: 32768}
                                                                                                            switch fi64 {
                                                                                                            case 1:
                                                                                                                v4 += 127
                                                                                                            case 2147483647:
                                                                                                                v15 -= (v15 + (v15 + uint8(v2)))
                                                                                                            }
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    if true {
                                                                                                        v5 = (v5 + (v5 >> (((^(v2))) & 63)))
                                                                                                    }
                                                                                                    _ = v84
                                                                                                }
                                                                                            case 2147483647:
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    v4 += v15
                                                                                                    var v87 uint32 = 902
                                                                                                    _ = v87
                                                                                                }
                                                                                            }
                                                                                            v3["a"] = false
                                                                                            var v88 uint32 = 4294967295
                                                                                            _ = v88
                                                                                        } else {
                                                                                            v63 = v63
                                                                                        }
                                                                                    }
                                                                                    var v89 int16 = 42
                                                                                    v90 := GenS9{F0: v63, F1: 2, F2: 1}
                                                                                    _ = fi81
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            if true {
                                                                                _ = len(v1)
                                                                                var v91 bool = false
                                                                                _ = v91
                                                                            } else {
                                                                                v92 := v3[("hello" + "\n\t")]
                                                                                v15 = v15
                                                                                _ = len(v3)
                                                                                _ = v92
                                                                            }
                                                                            for fi93 := 0; fi93 < 16; fi93++ {
                                                                                v15 = 92
                                                                                v1 = append(v1, int(v5))
                                                                                _ = fi93
                                                                            }
                                                                            var v94 float32 = 1e10
                                                                            _ = v94
                                                                        }
                                                                    }
                                                                    fmt.Println(v63, v15)
                                                                }
                                                                var v95 int8 = 5
                                                                _ = v95
                                                            }
                                                        default:
                                                            v5 = (v5 >> (((v2 + (^(v2)))) & 63))
                                                        }
                                                    }
                                                }
                                            }
                                            _ = fi64
                                        }
                                        var v96 int16 = 184
                                        _ = v63
                                        _ = v96
                                    }
                                    for fi97 := 0; fi97 < 5; fi97++ {
                                        var v98 uint = genFn0()
                                        if (false && (801 <= 65535)) {
                                            v2 = (v98 << ((v2) & 63))
                                        } else {
                                            v99 := GenS3{F0: v5}
                                            var v100 uint32 = 1
                                            _ = v100
                                            _ = v99
                                        }
                                        _ = fi97
                                        _ = v98
                                    }
                                }
                            }
                        }
                        if (!true) {
                            var v101 int32 = 42
                            _ = v101
                        } else {
                            fmt.Println(v15, v2, v15)
                            fmt.Println(v15)
                        }
                    }
                }
                _ = v21
            }
        case 128:
            var v102 uint16 = 726
            _ = v102
        case -931:
            if true {
                var v103 uint = (v2 / ((v2) | 1))
                var v104 int16 = 564
                var v105 uint16 = 42
                _ = v103
                _ = v104
                _ = v105
            } else {
                var v106 uint8 = v15
                _ = v106
            }
        }
        v3[((("foo bar" + ("hello" + ("\n\t" + ""))) + "hello") + "\n\t")] = ((false && false) || true)
    } else {
        if true {
            fmt.Println(v2)
            for fi107 := 0; fi107 < 5; fi107++ {
                v5 -= (v5 | 42)
                _ = fi107
            }
        } else {
            v5 += v5
            if (42 <= (v5 & (v5 << ((genFn0()) & 63)))) {
                v1 = append(v1, 1)
            }
            var v108 int64 = v5
            _ = v108
        }
        var v109 float64 = 0.0
        _ = len(v1)
        _ = v109
    }
    v110 := v3["hello"]
    fmt.Println(v15, v5, v110)
    if v110 {
        fmt.Println(v4)
        for fi111 := 0; fi111 < 14; fi111++ {
            v112 := GenS8{F0: 896, F1: 69}
            _ = fi111
            _ = v112
        }
    }
    for fi113 := 0; fi113 < 13; fi113++ {
        if (42 > v5) {
            if (v110 && v110) {
                var v114 int = (fi113 / (((fi113 - (fi113 >> ((v2) & 63)))) | 1))
                _ = v114
            } else {
                v115 := GenS15{F0: 642, F1: (-(v5))}
                _ = v115
            }
            switch int(v5) {
            case 672:
                v116 := GenS16{F0: 100, F1: 42}
                _ = v116
            case -904:
                switch ("hello" + "") {
                case "\n\t":
                    for fi117 := 0; fi117 < 7; fi117++ {
                        v118 := v3[("" + "hello")]
                        if (-1 >= 357) {
                            var v119 uint8 = (^(v4))
                            v118 = v118
                            _ = v119
                        } else {
                            if len(v1) > 0 {
                                v1[0] = (fi113 << (((v2 + 223)) & 63))
                            }
                            if true {
                                if (v118 && true) {
                                    v5 = -1
                                    v1 = append(v1, fi113)
                                } else {
                                    v4 = v15
                                    v120 := GenS1{F0: 255}
                                    var v121 uint64 = 422
                                    _ = v120
                                    _ = v121
                                }
                                v118 = (v118 && v110)
                            } else {
                                var v122 int64 = (v5 ^ v5)
                                v15 = v4
                                switch "hello" {
                                case "a":
                                    var v123 int32 = 42
                                    _ = v123
                                case "\n\t":
                                    var v124 uint16 = 643
                                    _ = v124
                                case "":
                                    delete(v3, "foo bar")
                                default:
                                    switch fi117 {
                                    case -2147483648:
                                        var v125 int16 = 380
                                        _ = v125
                                    }
                                }
                                _ = v122
                            }
                            if ((v110 || false) && (42 < -128)) {
                                if v118 {
                                    if false {
                                        v126 := GenS14{F0: 1}
                                        v15 = 127
                                        v127 := GenS6{F0: 424}
                                        v128 := GenS17{F0: 2, F1: 0, F2: 2147483647}
                                        _ = v126
                                        _ = v127
                                        _ = v128
                                    }
                                }
                            }
                        }
                        var v129 int8 = 100
                        _ = fi117
                        _ = v118
                        _ = v129
                    }
                default:
                    for fi130 := 0; fi130 < 12; fi130++ {
                        v131 := GenS18{F0: 42, F1: 2}
                        _ = fi130
                        _ = v131
                    }
                }
            }
            for fi132 := 0; fi132 < 15; fi132++ {
                var v133 uint64 = 38
                v110 = (1e10 > -2.25)
                _ = fi132
                _ = v133
            }
        } else {
            if true {
                delete(v3, "")
            } else {
                var v134 uint64 = 501
                var v135 uint8 = v15
                fmt.Println(v4, v135)
                _ = v134
                _ = v135
            }
            var v136 int = fi113
            for fi137 := 0; fi137 < 7; fi137++ {
                var v138 uint16 = 255
                _ = len(v3)
                fmt.Println(v2, v2)
                _ = fi137
                _ = v138
            }
            _ = v136
        }
        _ = fi113
    }
    if false {
        v139 := GenS19{F0: v15, F1: 406}
        fmt.Println(v4)
        _ = v139
    }
    var v140 int = 128
    if true {
        if len(v1) > 0 {
            v1[0] = (v140 | v140)
        }
    } else {
        if len(v1) > 0 {
            v1[0] = v140
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v140:", v140)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v110
    _ = v140
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
