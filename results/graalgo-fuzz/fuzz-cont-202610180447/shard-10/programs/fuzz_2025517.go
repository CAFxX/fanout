package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool; F1 int32; F2 int16 }
type GenS10 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string; F1 bool }
type GenS14 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0(p0 uint, p1 int32, p2 int8) uint {
    v7 := GenS1{F0: "foo bar"}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int = (fi8 * fi8)
        fmt.Println(p2, p2, p1)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1(p0 uint16, p1 int8) int64 {
    v10 := []int{-1, 2147483647, 0}
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    return -9223372036854775808
}

func main() {
    var v1 int8 = 59
    var v2 uint16 = 212
    v3 := make(map[string]uint32)
    var v4 uint32 = 42
    v5 := []int{}
    v6 := make(map[int]int)
    _ = len(v3)
    if (!false) {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 float32 = (1.5 * -2.25)
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(v2, v12, fi13)
                if true {
                    if true {
                        switch "a" {
                        case "hello":
                            if ((0 != 1) || true) {
                                delete(v6, fi11)
                                v14 := GenS1{F0: "a"}
                                if (42 != 65535) {
                                    delete(v3, (("hello" + "") + "foo bar"))
                                    v4 += (v4 + v4)
                                    fmt.Println(v1)
                                } else {
                                    v6[(^(fi11))] = (fi13 << ((uint(v2)) & 63))
                                }
                                _ = v14
                            }
                        case "":
                            v1 -= v1
                            v15 := GenS1{F0: "foo bar"}
                            v12 -= 0.1
                            var v16 bool = true
                            v17 := GenS1{F0: "foo bar"}
                            _ = v15
                            _ = v16
                            _ = v17
                        default:
                            v5 = append(v5, fi13)
                        }
                        var v18 string = "\n\t"
                        if (v18 >= "foo bar") {
                            v19 := GenS2{F0: 113}
                            delete(v3, "foo bar")
                            v18 = v18
                            _ = v19
                        }
                        var v20 float64 = -2.25
                        _ = v18
                        _ = v20
                    } else {
                        v4 -= v4
                    }
                    v21 := GenS3{F0: 0, F1: 42}
                    var v22 bool = true
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS3{F0: 1, F1: 2}
                    v2 -= (v2 & 255)
                    _ = v23
                }
                var v24 int = fi11
                _ = fi13
                _ = v24
            }
            switch (^(fi11)) {
            case 0:
                if (!true) {
                    switch ("hello" + ("foo bar" + "hello")) {
                    case "a":
                        switch "hello" {
                        case "a":
                            v4 -= (v4 & 2147483647)
                        case "foo bar":
                            var v25 float32 = 0.1
                            _ = v25
                        case "\n\t":
                            v26 := GenS2{F0: 255}
                            _ = v26
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    default:
                        switch (fi11 >> ((2) & 63)) {
                        case -1:
                            if true {
                                if (!(fi11 != fi11)) {
                                    v2 = 255
                                    if ((66 > 42) && true) {
                                        v1 = v1
                                        if true {
                                            var v27 int16 = -1
                                            v28 := GenS4{F0: 2147483648, F1: "a"}
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    if (false && ((3.14159 - (((float64(((-2.25 * v12) - (v12 + float32((float64(3.14159) * 0.0))))) * 3.14159) - 3.14159) - 1.5)) <= 0.0)) {
                                        switch ("" + ("a" + "hello")) {
                                        case "a":
                                            v29 := GenS5{F0: 880, F1: 4294967295}
                                            _ = v29
                                        case "":
                                            v4 = 2147483647
                                        }
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v30 int16 = -967
                                        if (false || (2 != 159)) {
                                            var v31 uint8 = 90
                                            v32 := v6[0]
                                            _ = v31
                                            _ = v32
                                        } else {
                                            for fi33 := 0; fi33 < 14; fi33++ {
                                                var v34 int16 = (v30 + v30)
                                                _ = fi33
                                                _ = v34
                                            }
                                            var v35 int16 = (v30 ^ (v30 | (v30 / ((v30) | 1))))
                                            _ = v35
                                        }
                                        _ = v30
                                    }
                                } else {
                                    v36 := GenS6{F0: v2, F1: ""}
                                    if (false && ((v4 + v4) > (v4 - 2147483648))) {
                                        var v37 int64 = -765
                                        switch fi11 {
                                        case 4294967295:
                                            if true {
                                                v4 += (^(v4))
                                            } else {
                                                var v38 int16 = -1
                                                var v39 bool = true
                                                v39 = v39
                                                _ = v38
                                                _ = v39
                                            }
                                        case -2147483648:
                                            var v40 int64 = (^(v37))
                                            _ = v40
                                        case 127:
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(v12, v37)
                                            case "\n\t":
                                                fmt.Println(fi11, fi11, v4)
                                            case "a":
                                                v4 = v4
                                            default:
                                                if ((("" + "\n\t") + "hello") <= ((("\n\t" + "hello") + ((("hello" + "") + (((("a" + "a") + ("foo bar" + "")) + "foo bar") + "foo bar")) + (("\n\t" + "hello") + ("foo bar" + ("foo bar" + "hello"))))) + (("foo bar" + "\n\t") + ("" + (("\n\t" + ("" + "a")) + "foo bar"))))) {
                                                    var v41 float32 = 0.0
                                                    var v42 int64 = (v37 & (v37 * int64(v4)))
                                                    var v43 int64 = v42
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            }
                                        }
                                        fmt.Println(v2, v12, v37)
                                        _ = v37
                                    } else {
                                        var v44 int32 = -1
                                        switch fi11 {
                                        case -9223372036854775808:
                                            v2 -= v2
                                        case 42:
                                            v45 := GenS7{F0: 18446744073709551615, F1: 128}
                                            _ = v45
                                        case 4294967295:
                                            v4 -= 42
                                        }
                                        if false {
                                            for fi46 := 0; fi46 < 12; fi46++ {
                                                v47 := GenS8{F0: 140}
                                                _ = fi46
                                                _ = v47
                                            }
                                            v48 := GenS1{F0: "\n\t"}
                                            v1 = (-(v1))
                                            _ = v48
                                        } else {
                                            for fi49 := 0; fi49 < 3; fi49++ {
                                                _ = len(v6)
                                                v50 := GenS6{F0: (v2 >> ((2) & 63)), F1: ("foo bar" + "\n\t")}
                                                switch ("hello" + "a") {
                                                case "foo bar":
                                                    v51 := GenS7{F0: 51, F1: 128}
                                                    _ = v51
                                                case "a":
                                                    v52 := GenS3{F0: 128, F1: 2}
                                                    _ = v52
                                                }
                                                _ = fi49
                                                _ = v50
                                            }
                                        }
                                        _ = v44
                                    }
                                    var v53 uint = genFn0(174, 2147483647, int8(v4))
                                    _ = v36
                                    _ = v53
                                }
                                v1 += v1
                            }
                        case 127:
                            if true {
                                for fi54 := 0; fi54 < 6; fi54++ {
                                    var v55 int = (fi11 >> ((42) & 63))
                                    v2 += v2
                                    if (uint16(v4) < (v2 << ((genFn0(1, -1, v1)) & 63))) {
                                        delete(v6, fi54)
                                        fmt.Println(fi11, v1)
                                        switch (fi11 % ((fi11) | 1)) {
                                        case -110:
                                            v56 := GenS6{F0: 0, F1: "hello"}
                                            _ = v56
                                        }
                                    }
                                    if (413 == -32768) {
                                        var v57 uint8 = 128
                                        v1 += (v1 + (v1 >> ((65535) & 63)))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v57
                                    }
                                    v1 += v1
                                    _ = fi54
                                    _ = v55
                                }
                                if (v4 < v4) {
                                    if len(v5) > 0 {
                                        v5[0] = -1
                                    }
                                }
                                fmt.Println(v4, v1)
                            }
                        case -9223372036854775808:
                            for fi58 := 0; fi58 < 11; fi58++ {
                                fmt.Println(fi11)
                                var v59 string = "hello"
                                switch v59 {
                                case "":
                                    fmt.Println(v4, fi11)
                                case "foo bar":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(fi11, fi11, v59)
                                default:
                                    if (52 == 631) {
                                        if (128 >= 0) {
                                            if true {
                                                v1 = v1
                                                if true {
                                                    if false {
                                                        switch (^(fi11)) {
                                                        case 127:
                                                            v4 = 65535
                                                        case 1:
                                                            v3["foo bar"] = (^(v4))
                                                        default:
                                                            var v60 uint32 = v4
                                                            _ = v60
                                                        }
                                                    }
                                                    v61 := GenS4{F0: v4, F1: v59}
                                                    _ = v61
                                                } else {
                                                    var v62 uint64 = 18446744073709551615
                                                    _ = v62
                                                }
                                                v63 := GenS9{F0: (v12 > ((1.5 - v12) - 1e10)), F1: 2147483647, F2: 32767}
                                                _ = v63
                                            }
                                            v4 += (v4 - v4)
                                        }
                                        _ = len(v6)
                                        v4 += 369
                                    }
                                }
                                _ = fi58
                                _ = v59
                            }
                        default:
                            if true {
                                var v64 uint8 = 180
                                _ = v64
                            } else {
                                if true {
                                    var v65 bool = true
                                    var v66 int = fi11
                                    _ = v65
                                    _ = v66
                                }
                                for fi67 := 0; fi67 < 1; fi67++ {
                                    v4 = v4
                                    v2 -= v2
                                    var v68 float32 = 1e10
                                    v68 = 3.14159
                                    _ = fi67
                                    _ = v68
                                }
                                var v69 bool = (v12 <= v12)
                                _ = v69
                            }
                        }
                    }
                    var v70 uint32 = (v4 << ((65535) & 63))
                    _ = v70
                }
            case 917:
                v71 := GenS6{F0: (^(v2)), F1: "\n\t"}
                _ = v71
            }
            _ = fi11
            _ = v12
        }
        v2 += 602
        v72 := GenS8{F0: 42}
        _ = v72
    } else {
        if (247 >= uint8(v2)) {
            fmt.Println(v4, v2, v4)
            v73 := v6[127]
            if true {
                var v74 uint8 = 89
                _ = len(v3)
                for fi75 := 0; fi75 < 3; fi75++ {
                    for fi76 := 0; fi76 < 4; fi76++ {
                        if len(v5) > 0 {
                            v5[0] = (-(fi76))
                        }
                        var v77 uint16 = (v2 ^ (v2 >> ((2) & 63)))
                        _ = fi76
                        _ = v77
                    }
                    var v78 uint32 = (v4 - v4)
                    _ = fi75
                    _ = v78
                }
                _ = v74
            } else {
                v79 := GenS10{F0: 0, F1: 1, F2: 4294967295}
                var v80 uint16 = (v2 - (v2 - v2))
                _ = v79
                _ = v80
            }
            _ = v73
        } else {
            for fi81 := 0; fi81 < 6; fi81++ {
                fmt.Println(fi81, fi81)
                _ = fi81
            }
            for fi82 := 0; fi82 < 19; fi82++ {
                _ = len(v5)
                var v83 float32 = 1.5
                v2 += 65535
                _ = fi82
                _ = v83
            }
        }
        v2 += 530
        var v84 int8 = (v1 * -128)
        _ = v84
    }
    for fi85 := 0; fi85 < 3; fi85++ {
        switch "hello" {
        case "":
            v1 = v1
        default:
            v3["\n\t"] = 2
        }
        v2 -= v2
        _ = fi85
    }
    for fi86 := 0; fi86 < 1; fi86++ {
        var v87 int32 = 42
        v88 := GenS8{F0: -9223372036854775808}
        _ = fi86
        _ = v87
        _ = v88
    }
    if (1 > (v1 - v1)) {
        fmt.Println(v1, v2, v4)
    }
    v89 := GenS11{F0: 42}
    for fi90 := 0; fi90 < 18; fi90++ {
        for fi91 := 0; fi91 < 8; fi91++ {
            if (1.5 == ((float64((1e10 - 0.0)) / 0.1) * 1e10)) {
                if false {
                    v1 = (v1 >> ((1) & 63))
                    _ = len(v6)
                }
                switch "hello" {
                case "\n\t":
                    var v92 bool = true
                    _ = v92
                default:
                    v3[""] = 4294967295
                }
                if ("\n\t" == ("a" + ("\n\t" + "\n\t"))) {
                    var v93 int = 128
                    v94 := GenS3{F0: 102, F1: 42}
                    var v95 bool = false
                    var v96 uint8 = 128
                    var v97 string = "\n\t"
                    _ = v93
                    _ = v94
                    _ = v95
                    _ = v96
                    _ = v97
                } else {
                    v4 = v4
                }
            } else {
                v4 += (^(v4))
            }
            _ = fi91
        }
        for fi98 := 0; fi98 < 16; fi98++ {
            if (-1 == -60) {
                if (0 >= 0) {
                    var v99 int = 32767
                    var v100 float64 = (1.5 + -2.25)
                    _ = v100
                    _ = v99
                }
                var v101 int16 = 1
                v102 := GenS12{F0: 42}
                _ = v101
                _ = v102
            } else {
                v5 = append(v5, 1)
            }
            _ = fi98
        }
        for fi103 := 0; fi103 < 2; fi103++ {
            var v104 float64 = 0.0
            _ = fi103
            _ = v104
        }
        _ = fi90
    }
    _ = len(v3)
    if ((v4 >> ((1) & 63)) > 2147483648) {
        fmt.Println(v4)
    }
    if (true && true) {
        switch 4294967295 {
        case 32768:
            if ((v2 * v2) != (v2 ^ (v2 * 1))) {
                for fi105 := 0; fi105 < 4; fi105++ {
                    v2 = 0
                    _ = fi105
                }
            } else {
                v106 := GenS12{F0: 0}
                _ = v106
            }
        default:
            _ = len(v5)
        }
    }
    v4 += (^(v4))
    fmt.Println(v2, v1, v2)
    fmt.Println(v4)
    v4 = 2147483647
    v4 += 0
    for fi107 := 0; fi107 < 6; fi107++ {
        var v108 bool = ((v2 > v2) && false)
        delete(v6, (fi107 - (fi107 * (fi107 / ((fi107) | 1)))))
        _ = fi107
        _ = v108
    }
    switch "a" {
    case "a":
        if false {
            v1 += v1
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        } else {
            if false {
                if (1e10 <= 0.0) {
                    v2 -= v2
                } else {
                    if (94 < -1) {
                        switch ("hello" + ("hello" + "foo bar")) {
                        case "a":
                            var v109 int16 = int16(v2)
                            _ = v109
                        case "":
                            fmt.Println(v4, v1)
                        case "hello":
                            var v110 int = 716
                            _ = v110
                        default:
                            for fi111 := 0; fi111 < 8; fi111++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                for fi112 := 0; fi112 < 13; fi112++ {
                                    var v113 bool = true
                                    v114 := GenS12{F0: 1000}
                                    var v115 bool = v113
                                    if (!v113) {
                                        var v116 int32 = -2147483648
                                        if v113 {
                                            var v117 int8 = (v1 >> ((2) & 63))
                                            _ = v117
                                        } else {
                                            var v118 uint32 = 42
                                            if (973 < 178) {
                                                if (196 < -1) {
                                                    if v113 {
                                                        v113 = (v4 == v4)
                                                        v4 += (v4 << ((genFn0(85, 2147483647, 42)) & 63))
                                                    } else {
                                                        v119 := v6[766]
                                                        var v120 int64 = 841
                                                        var v121 uint64 = 713
                                                        _ = v119
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                    v115 = v113
                                                } else {
                                                    switch (fi112 | 1) {
                                                    case 920:
                                                        if v113 {
                                                            switch fi111 {
                                                            case 4294967295:
                                                                v2 = (^(v2))
                                                            }
                                                            var v122 uint8 = 42
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch "" {
                                                                case "foo bar":
                                                                    if (!((v2 < (v2 ^ 61)) && v113)) {
                                                                        fmt.Println(v1, v113, v113)
                                                                    } else {
                                                                        var v123 int8 = int8(fi112)
                                                                        _ = v123
                                                                    }
                                                                default:
                                                                    var v124 int8 = v1
                                                                    _ = v124
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(fi111)
                                                            case "hello":
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v125 uint8 = 1
                                                                    _ = v125
                                                                case "a":
                                                                    var v126 uint = 255
                                                                    _ = v126
                                                                default:
                                                                    var v127 uint16 = uint16(v122)
                                                                    _ = v127
                                                                }
                                                            default:
                                                                v2 = 255
                                                            }
                                                            _ = v122
                                                        } else {
                                                            fmt.Println(v113, v2, fi111)
                                                            v4 = v118
                                                        }
                                                    case -67:
                                                        if (((0.1 - 1.5) / float64(1e10)) <= (3.14159 - -2.25)) {
                                                            if v113 {
                                                                delete(v6, (fi111 % ((2147483647) | 1)))
                                                            }
                                                            var v128 uint64 = 1
                                                            _ = v128
                                                        }
                                                    default:
                                                        v115 = (float32(-2.25) == 1e10)
                                                    }
                                                }
                                                v118 = (v118 ^ 391)
                                            }
                                            _ = v118
                                        }
                                        _ = v116
                                    }
                                    _ = fi112
                                    _ = v113
                                    _ = v114
                                    _ = v115
                                }
                                _ = fi111
                            }
                        }
                        var v129 uint8 = 1
                        _ = v129
                    } else {
                        v1 -= v1
                        v130 := GenS13{F0: "hello", F1: true}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v130
                    }
                    var v131 uint = 1
                    for fi132 := 0; fi132 < 1; fi132++ {
                        delete(v3, "")
                        _ = fi132
                    }
                    _ = v131
                }
            }
            var v133 int16 = -32768
            _ = v133
        }
    case "hello":
        for fi134 := 0; fi134 < 18; fi134++ {
            var v135 uint = 519
            if true {
                if (true || true) {
                    _ = len(v5)
                    if (false || ((fi134 * fi134) != (fi134 / ((fi134) | 1)))) {
                        if false {
                            switch fi134 {
                            case 1:
                                v6[1] = fi134
                            case 42:
                                fmt.Println(v135, v1, v2)
                            }
                            for fi136 := 0; fi136 < 20; fi136++ {
                                fmt.Println(fi136, v2, v4)
                                v137 := GenS14{F0: 1, F1: v1, F2: 0}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi136
                                _ = v137
                            }
                        } else {
                            v2 += (v2 - 32767)
                            if false {
                                _ = len(v5)
                                delete(v3, (("\n\t" + ("\n\t" + ("\n\t" + ("" + "")))) + "a"))
                            }
                        }
                        switch (fi134 - 9223372036854775807) {
                        case 127:
                            var v138 float32 = float32((-2.25 + 1.5))
                            _ = v138
                        case 0:
                            v139 := GenS8{F0: 1}
                            _ = v139
                        case -2147483648:
                            for fi140 := 0; fi140 < 3; fi140++ {
                                switch "a" {
                                case "foo bar":
                                    v4 = (v4 % ((v4) | 1))
                                    v1 -= v1
                                default:
                                    v2 = (^(v2))
                                }
                                switch (("foo bar" + "") + "\n\t") {
                                case "hello":
                                    v141 := GenS5{F0: 4294967295, F1: 0}
                                    _ = v141
                                case "\n\t":
                                    switch "a" {
                                    case "hello":
                                        var v142 int16 = -233
                                        _ = v142
                                    case "\n\t":
                                        v1 -= v1
                                    }
                                case "":
                                    v135 -= v135
                                }
                                _ = fi140
                            }
                        }
                        var v143 int32 = 2147483647
                        _ = v143
                    }
                }
                var v144 uint = 65535
                switch "" {
                case "foo bar":
                    var v145 int8 = v1
                    _ = v145
                case "":
                    if ((v2 * (^(v2))) != (v2 / ((v2) | 1))) {
                        v146 := GenS4{F0: v4, F1: ("a" + ("a" + "hello"))}
                        _ = v146
                    }
                }
                _ = v144
            } else {
                v1 = v1
                var v147 uint16 = 2
                v2 = v147
                _ = v147
            }
            if true {
                var v148 int64 = -297
                switch (fi134 / ((int(v135)) | 1)) {
                case 762:
                    var v149 int8 = (v1 | v1)
                    _ = v149
                case 42:
                    if false {
                        var v150 uint16 = (v2 * 65535)
                        _ = v150
                    }
                case 4294967295:
                    v2 = v2
                }
                v135 += (v135 & 2)
                _ = v148
            } else {
                for fi151 := 0; fi151 < 6; fi151++ {
                    v135 -= v135
                    var v152 int32 = 2147483647
                    var v153 int64 = 797
                    _ = fi151
                    _ = v152
                    _ = v153
                }
                switch -816 {
                case 787:
                    for fi154 := 0; fi154 < 2; fi154++ {
                        if (255 <= 236) {
                            fmt.Println(v4, fi134, v1)
                            v135 += 0
                            v1 -= (^(v1))
                        } else {
                            var v155 int = (fi154 & (^(fi154)))
                            switch 42 {
                            case 0:
                                if true {
                                    var v156 float32 = 3.14159
                                    _ = v156
                                } else {
                                    var v157 int32 = 0
                                    v158 := GenS2{F0: 0}
                                    _ = v157
                                    _ = v158
                                }
                            default:
                                var v159 int = v155
                                _ = v159
                            }
                            var v160 float32 = 3.14159
                            _ = v155
                            _ = v160
                        }
                        _ = fi154
                    }
                default:
                    v4 = 1
                }
            }
            _ = fi134
            _ = v135
        }
    case "\n\t":
        var v161 float32 = 3.14159
        _ = v161
    default:
        var v162 uint32 = (v4 - 2)
        _ = v162
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
}
