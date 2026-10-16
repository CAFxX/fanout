package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 int16; F1 int8; F2 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 int8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int8; F1 string }
type GenS15 struct { F0 bool; F1 uint8 }
type GenS16 struct { F0 uint64; F1 string; F2 uint32 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS19 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint8, p1 uint64, p2 uint) uint8 {
    fmt.Println(p0, p0)
    p1 = p1
    var v5 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]string)
    v3 := []uint32{85, 4294967295}
    var v4 float32 = -2.25
    if false {
        v4 = ((((v4 * v4) - (0.1 * (v4 * v4))) / 1.5) * -2.25)
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 0, F1: (("\n\t" + "a") + "hello"), F2: genFn0(0, 740, (^(v1)))}
        _ = fi6
        _ = v7
    }
    var v8 int8 = 1
    fmt.Println(v1)
    for fi9 := 0; fi9 < 3; fi9++ {
        if ((!(0 < 1)) || false) {
            v1 = (v1 << ((v1) & 63))
            v10 := GenS1{F0: 65535, F1: ("\n\t" + ""), F2: 127}
            switch "\n\t" {
            case "\n\t":
                var v11 float64 = 1.5
                _ = v11
            case "":
                for fi12 := 0; fi12 < 7; fi12++ {
                    fmt.Println(v8, v1)
                    fmt.Println(fi12)
                    _ = fi12
                }
            default:
                var v13 uint64 = 42
                _ = v13
            }
            _ = v10
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                v8 = 1
                _ = fi14
            }
            delete(v2, (("hello" + "hello") + "a"))
        }
        _ = fi9
    }
    switch 128 {
    case -2147483648:
        fmt.Println(v4)
    case 6:
        var v15 int8 = 0
        _ = v15
    case 664:
        fmt.Println(v8)
    }
    var v16 uint16 = 255
    switch "\n\t" {
    case "a":
        switch ("a" + ("hello" + "\n\t")) {
        case "foo bar":
            if (uint(v16) > (^(v1))) {
                switch "foo bar" {
                case "\n\t":
                    if false {
                        if (0 != 2147483647) {
                            var v17 int = 127
                            fmt.Println(v8, v8)
                            fmt.Println(v1)
                            _ = v17
                        } else {
                            v1 += 0
                            for fi18 := 0; fi18 < 11; fi18++ {
                                var v19 uint64 = 4294967295
                                switch fi18 {
                                case 127:
                                    var v20 int64 = int64(fi18)
                                    _ = v20
                                default:
                                    v21 := GenS2{F0: v19, F1: int32(v1), F2: (fi18 % ((301) | 1))}
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    _ = len(v2)
                                    var v23 uint16 = uint16(v19)
                                    _ = fi22
                                    _ = v23
                                }
                                _ = fi18
                                _ = v19
                            }
                            fmt.Println(v8)
                        }
                        var v24 uint32 = 824
                        switch ("hello" + "") {
                        case "foo bar":
                            if true {
                                if true {
                                    fmt.Println(v8, v1)
                                    switch -89 {
                                    case 9223372036854775807:
                                        v3 = append(v3, (v24 + 2147483648))
                                    case 2147483647:
                                        var v25 int64 = 0
                                        _ = v25
                                    }
                                } else {
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        var v27 int8 = (-(v8))
                                        v24 = v24
                                        v28 := GenS3{F0: 0, F1: v8, F2: -569}
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                    if false {
                                        var v29 float64 = 0.0
                                        _ = v29
                                    }
                                }
                            } else {
                                var v30 int8 = v8
                                _ = v30
                            }
                        case "a":
                            v4 -= 3.14159
                        case "hello":
                            v1 -= v1
                        }
                        _ = v24
                    }
                case "a":
                    var v31 uint16 = (v16 ^ v16)
                    _ = v31
                }
            }
        }
    }
    if false {
        if false {
            var v32 float64 = (1e10 + (3.14159 - (-2.25 + float64(3.14159))))
            _ = v32
        }
        switch "hello" {
        case "hello":
            var v33 uint64 = 2
            _ = v33
        case "\n\t":
            if (0 >= 0) {
                v2["\n\t"] = ""
                if (!false) {
                    v16 = 42
                    for fi34 := 0; fi34 < 7; fi34++ {
                        fmt.Println(v16)
                        fmt.Println(v16)
                        _ = fi34
                    }
                    v4 += (v4 * v4)
                }
            } else {
                if false {
                    v4 = v4
                    v1 = 0
                    if false {
                        for fi35 := 0; fi35 < 19; fi35++ {
                            v4 = v4
                            var v36 int = fi35
                            _ = fi35
                            _ = v36
                        }
                    } else {
                        if true {
                            if (true && ((true && (false && false)) || false)) {
                                v8 -= v8
                                var v37 float64 = (1.5 + (1e10 + (((1e10 / (((float64(v4) + 1.5) + 1.5) + 0.0)) + 1e10) / ((-2.25 / 0.1) - 0.0))))
                                for fi38 := 0; fi38 < 12; fi38++ {
                                    v16 = v16
                                    _ = fi38
                                }
                                _ = v37
                            }
                        } else {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                for fi40 := 0; fi40 < 8; fi40++ {
                                    if false {
                                        var v41 uint64 = 1
                                        _ = v41
                                    } else {
                                        var v42 uint = v1
                                        v43 := v2[""]
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = fi40
                                }
                                if true {
                                    fmt.Println(v16)
                                } else {
                                    if (v16 == v16) {
                                        v4 = v4
                                    } else {
                                        var v44 uint64 = 1
                                        switch (fi39 ^ fi39) {
                                        case 42:
                                            var v45 uint64 = v44
                                            _ = v45
                                        }
                                        for fi46 := 0; fi46 < 8; fi46++ {
                                            var v47 int8 = (v8 * v8)
                                            _ = fi46
                                            _ = v47
                                        }
                                        _ = v44
                                    }
                                    v48 := GenS1{F0: (v16 * (v16 << ((v1) & 63))), F1: "a", F2: uint8(v16)}
                                    v49 := GenS2{F0: 106, F1: int32(fi39), F2: (fi39 >> ((624) & 63))}
                                    _ = v48
                                    _ = v49
                                }
                                v8 = v8
                                _ = fi39
                            }
                            var v50 int16 = -716
                            v8 -= 42
                            _ = v50
                        }
                        fmt.Println(v16, v16, v16)
                        var v51 int = 32767
                        _ = v51
                    }
                } else {
                    for fi52 := 0; fi52 < 12; fi52++ {
                        switch "a" {
                        case "hello":
                            v53 := GenS3{F0: -32768, F1: v8, F2: 1000}
                            _ = v53
                        }
                        var v54 bool = true
                        var v55 int16 = 1000
                        _ = fi52
                        _ = v54
                        _ = v55
                    }
                    if false {
                        switch 0 {
                        case -549:
                            v1 += 4294967295
                        }
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v56 float64 = ((1.5 - 0.0) - (float64(v4) / (1.5 + 0.0)))
                            if (127 == int(v1)) {
                                for fi57 := 0; fi57 < 9; fi57++ {
                                    if true {
                                        v1 -= 42
                                        switch ("foo bar" + "hello") {
                                        case "hello":
                                            switch (("\n\t" + "foo bar") + "\n\t") {
                                            case "foo bar":
                                                delete(v2, ("hello" + "hello"))
                                            }
                                        }
                                        v56 = v56
                                    } else {
                                        v4 = v4
                                    }
                                    var v58 int32 = 42
                                    switch ("a" + "hello") {
                                    case "\n\t":
                                        var v59 uint16 = 169
                                        _ = v59
                                    case "hello":
                                        for fi60 := 0; fi60 < 18; fi60++ {
                                            v56 = v56
                                            switch "\n\t" {
                                            case "\n\t":
                                                switch (fi57 % ((fi60) | 1)) {
                                                case -1:
                                                    switch fi60 {
                                                    case 2147483647:
                                                        var v61 uint = v1
                                                        _ = v61
                                                    case -111:
                                                        var v62 bool = true
                                                        _ = v62
                                                    case 843:
                                                        if false {
                                                            var v63 uint64 = 271
                                                            _ = v63
                                                        } else {
                                                            var v64 uint = v1
                                                            v65 := GenS3{F0: 1, F1: (^(v8)), F2: 1000}
                                                            if true {
                                                                var v66 uint64 = 400
                                                                _ = v66
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    default:
                                                        var v67 int32 = -599
                                                        _ = v67
                                                    }
                                                case 565:
                                                    var v68 int64 = 263
                                                    _ = v68
                                                default:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v1 = 0
                                                    case "hello":
                                                        if true {
                                                            var v69 int8 = (v8 * (-(v8)))
                                                            v70 := GenS4{F0: v16, F1: 32767, F2: "hello"}
                                                            v71 := GenS1{F0: v16, F1: "\n\t", F2: uint8(v16)}
                                                            var v72 int16 = -944
                                                            switch "foo bar" {
                                                            case "a":
                                                                v72 += v72
                                                            case "foo bar":
                                                                v73 := GenS5{F0: (^(v1))}
                                                                _ = v73
                                                            }
                                                            fmt.Println(v56, v16, v4)
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    default:
                                                        if true {
                                                            v74 := GenS6{F0: 689, F1: 2147483647, F2: (^(v1))}
                                                            v58 -= int32(v8)
                                                            _ = v74
                                                        }
                                                    }
                                                }
                                            case "a":
                                                var v75 bool = (v8 < (v8 / ((v8) | 1)))
                                                _ = v75
                                            default:
                                                v8 = v8
                                            }
                                            var v76 uint = 4294967295
                                            _ = fi60
                                            _ = v76
                                        }
                                    default:
                                        switch fi57 {
                                        case 523:
                                            v4 += ((v4 - 3.14159) + 0.0)
                                        case 128:
                                            v77 := GenS7{F0: "\n\t", F1: 1}
                                            _ = v77
                                        case 837:
                                            _ = len(v2)
                                        default:
                                            _ = len(v3)
                                        }
                                    }
                                    _ = fi57
                                    _ = v58
                                }
                                v78 := GenS3{F0: 549, F1: v8, F2: -1}
                                v4 -= ((float32(((0.0 + v56) / v56)) - float32(0.1)) / v4)
                                _ = v78
                            } else {
                                fmt.Println(v1, v56, v56)
                            }
                            _ = v56
                        }
                        if false {
                            v79 := v2[(("\n\t" + "foo bar") + ("hello" + ("a" + "hello")))]
                            v80 := GenS5{F0: (v1 << ((v1) & 63))}
                            _ = v79
                            _ = v80
                        } else {
                            var v81 uint8 = 193
                            v82 := v2["foo bar"]
                            _ = v81
                            _ = v82
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if true {
                    var v83 uint64 = 18446744073709551615
                    for fi84 := 0; fi84 < 19; fi84++ {
                        v83 += 0
                        v4 = v4
                        _ = fi84
                    }
                    _ = v83
                }
                if (true && (true && true)) {
                    switch 128 {
                    case 32768:
                        v85 := GenS8{F0: (^(v16)), F1: ("hello" + "\n\t")}
                        _ = v85
                    case 127:
                        fmt.Println(v16)
                    case 32767:
                        v1 -= (v1 | v1)
                    }
                } else {
                    for fi86 := 0; fi86 < 15; fi86++ {
                        var v87 string = "hello"
                        _ = fi86
                        _ = v87
                    }
                    fmt.Println(v8, v8, v16)
                }
            }
        default:
            if false {
                switch 1 {
                case 127:
                    if ((v1 * v1) != 452) {
                        v88 := GenS7{F0: ((("a" + "") + "") + "\n\t"), F1: (v16 & (v16 % ((v16) | 1)))}
                        if true {
                            v1 = 42
                            if true {
                                if (870 == 65535) {
                                    v3 = append(v3, 879)
                                    if (0.1 >= (v4 - 1.5)) {
                                        var v89 bool = true
                                        v90 := GenS3{F0: 1000, F1: int8(v16), F2: 789}
                                        _ = v89
                                        _ = v90
                                    }
                                    fmt.Println(v8, v4)
                                } else {
                                    for fi91 := 0; fi91 < 19; fi91++ {
                                        for fi92 := 0; fi92 < 1; fi92++ {
                                            var v93 int32 = 1
                                            _ = fi92
                                            _ = v93
                                        }
                                        _ = fi91
                                    }
                                    var v94 uint32 = 2147483648
                                    fmt.Println(v94, v1)
                                    _ = v94
                                }
                                for fi95 := 0; fi95 < 18; fi95++ {
                                    var v96 int16 = -32768
                                    if false {
                                        v1 -= v1
                                        if true {
                                            if (true || false) {
                                                if false {
                                                    v97 := GenS6{F0: 0, F1: 1, F2: (v1 >> (((v1 + v1)) & 63))}
                                                    var v98 string = "\n\t"
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    v96 -= v96
                                                }
                                                v2[""] = "\n\t"
                                                v4 += 1e10
                                            } else {
                                                if ((0 > -1) && true) {
                                                    for fi99 := 0; fi99 < 2; fi99++ {
                                                        fmt.Println(v4, fi99, v1)
                                                        _ = fi99
                                                    }
                                                    v16 = v16
                                                } else {
                                                    for fi100 := 0; fi100 < 8; fi100++ {
                                                        v16 += v16
                                                        var v101 float64 = 0.0
                                                        v102 := GenS2{F0: 790, F1: -2147483648, F2: fi95}
                                                        _ = fi100
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                    fmt.Println(v8, fi95, v1)
                                                }
                                                var v103 uint64 = 959
                                                fmt.Println(v4)
                                                _ = v103
                                            }
                                            var v104 string = "\n\t"
                                            _ = v104
                                        } else {
                                            var v105 int8 = (v8 ^ 100)
                                            switch 2147483647 {
                                            case -543:
                                                v96 -= v96
                                            }
                                            v8 = (v105 - -1)
                                            _ = v105
                                        }
                                    }
                                    if ((0.1 * 1.5) == 0.0) {
                                        v106 := GenS1{F0: (v16 * 255), F1: "\n\t", F2: 127}
                                        if (-7 == v8) {
                                            if (65535 == 65535) {
                                                if false {
                                                    fmt.Println(v8, v96)
                                                } else {
                                                    if (true || ((v8 == v8) && (!false))) {
                                                        v96 += 1
                                                    }
                                                    if false {
                                                        v107 := GenS7{F0: "", F1: uint16(v1)}
                                                        v3 = append(v3, 4294967295)
                                                        _ = v107
                                                    }
                                                    v16 = v16
                                                }
                                                v3 = append(v3, 2147483647)
                                                if (!false) {
                                                    if ((false || ((true || (((!true) && ((v96 << ((v1) & 63)) == (v96 * (v96 - (^(v96)))))) && false)) || ((18446744073709551615 < 950) || ((-9223372036854775808 == -9223372036854775808) || ((^(v8)) < (v8 ^ (v8 % ((-25) | 1)))))))) && (((("hello" + "") > "hello") && true) || (v4 < 1e10))) {
                                                        for fi108 := 0; fi108 < 8; fi108++ {
                                                            v2[("\n\t" + ((("a" + ("\n\t" + "")) + ("foo bar" + "a")) + "\n\t"))] = (("a" + ("" + "hello")) + ("a" + "foo bar"))
                                                            v109 := GenS9{F0: 42, F1: v8}
                                                            var v110 uint64 = uint64(v1)
                                                            _ = fi108
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                        fmt.Println(v16, v16, v96)
                                                    } else {
                                                        switch (fi95 % ((9223372036854775807) | 1)) {
                                                        case -9223372036854775808:
                                                            var v111 string = "hello"
                                                            _ = v111
                                                        case 32768:
                                                            v112 := v2[(("foo bar" + ("foo bar" + "foo bar")) + ("a" + ((("\n\t" + "") + "a") + "a")))]
                                                            _ = v112
                                                        default:
                                                            for fi113 := 0; fi113 < 18; fi113++ {
                                                                if false {
                                                                    v4 = v4
                                                                    var v114 float64 = ((((1e10 / 1.5) + (((((0.0 * (0.1 + 1.5)) - 1.5) - 0.1) + 1.5) - 0.0)) * 0.1) * 1e10)
                                                                    v115 := GenS10{F0: -20}
                                                                    _ = v114
                                                                    _ = v115
                                                                }
                                                                if true {
                                                                    var v116 int8 = v8
                                                                    _ = v116
                                                                }
                                                                _ = fi113
                                                            }
                                                        }
                                                        if ("a" < (("hello" + ("foo bar" + "foo bar")) + "hello")) {
                                                            if false {
                                                                if true {
                                                                    v96 = v96
                                                                    switch fi95 {
                                                                    case -9223372036854775808:
                                                                        v117 := GenS11{F0: "a", F1: (v8 / ((v8) | 1))}
                                                                        _ = v117
                                                                    case 128:
                                                                        v96 += v96
                                                                    }
                                                                    switch ("hello" + "") {
                                                                    case "hello":
                                                                        if true {
                                                                            v8 = 5
                                                                        }
                                                                    case "\n\t":
                                                                        var v118 float64 = -2.25
                                                                        _ = v118
                                                                    case "foo bar":
                                                                        switch fi95 {
                                                                        case 354:
                                                                            if (57 >= 4294967295) {
                                                                                switch (fi95 * 296) {
                                                                                case 4294967295:
                                                                                    for fi119 := 0; fi119 < 7; fi119++ {
                                                                                        fmt.Println(fi119, fi119, v96)
                                                                                        v96 -= 1
                                                                                        fmt.Println(v4, v1, v1)
                                                                                        _ = fi119
                                                                                    }
                                                                                case 32768:
                                                                                    if true {
                                                                                        delete(v2, (("foo bar" + "\n\t") + "\n\t"))
                                                                                    } else {
                                                                                        if (1 == 2147483647) {
                                                                                            if (false && false) {
                                                                                                fmt.Println(fi95)
                                                                                                var v120 bool = true
                                                                                                _ = v120
                                                                                            }
                                                                                            if true {
                                                                                                fmt.Println(v1, v4, v96)
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v121 uint32 = 2147483647
                                                                                                    if (true || true) {
                                                                                                        v8 += -128
                                                                                                        if (false || ((65535 != 1) && false)) {
                                                                                                            var v122 float32 = ((1.5 - v4) - ((float32(-2.25) + ((v4 / v4) / (float32(3.14159) / (v4 + 3.14159)))) / v4))
                                                                                                            _ = v122
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                for fi123 := 0; fi123 < 9; fi123++ {
                                                                                                                    _ = len(v2)
                                                                                                                    v1 -= v1
                                                                                                                    _ = fi123
                                                                                                                }
                                                                                                                var v124 float64 = (1e10 * ((0.0 - (float64(v4) / -2.25)) * -2.25))
                                                                                                                _ = v124
                                                                                                            } else {
                                                                                                                var v125 float32 = (0.0 / v4)
                                                                                                                _ = v125
                                                                                                            }
                                                                                                            var v126 float32 = v4
                                                                                                            _ = v126
                                                                                                        }
                                                                                                        if false {
                                                                                                            v121 += (v121 * 428)
                                                                                                            for fi127 := 0; fi127 < 17; fi127++ {
                                                                                                                var v128 uint8 = 135
                                                                                                                var v129 float64 = 1.5
                                                                                                                fmt.Println(v8, v16, v16)
                                                                                                                _ = fi127
                                                                                                                _ = v128
                                                                                                                _ = v129
                                                                                                            }
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                switch (fi95 ^ (fi95 - (fi95 | (-(fi95))))) {
                                                                                                                case 113:
                                                                                                                    fmt.Println(v1, v4, v8)
                                                                                                                case -9223372036854775808:
                                                                                                                    v130 := GenS12{F0: 97}
                                                                                                                    _ = v130
                                                                                                                case 319:
                                                                                                                    v1 = (v1 - 42)
                                                                                                                default:
                                                                                                                    if true {
                                                                                                                        for fi131 := 0; fi131 < 12; fi131++ {
                                                                                                                            if false {
                                                                                                                                var v132 float64 = 3.14159
                                                                                                                                v4 -= v4
                                                                                                                                _ = v132
                                                                                                                            } else {
                                                                                                                                v133 := GenS13{F0: 135}
                                                                                                                                _ = v133
                                                                                                                            }
                                                                                                                            _ = fi131
                                                                                                                        }
                                                                                                                        v96 = 42
                                                                                                                        if (v96 <= v96) {
                                                                                                                            v4 -= (v4 * float32(0.0))
                                                                                                                        } else {
                                                                                                                            for fi134 := 0; fi134 < 1; fi134++ {
                                                                                                                                fmt.Println(v4, v4, v1)
                                                                                                                                switch -1 {
                                                                                                                                case 2147483647:
                                                                                                                                    v4 = v4
                                                                                                                                case 219:
                                                                                                                                    v16 = v16
                                                                                                                                case 32768:
                                                                                                                                    v3 = append(v3, v121)
                                                                                                                                }
                                                                                                                                _ = fi134
                                                                                                                            }
                                                                                                                            var v135 int64 = -807
                                                                                                                            _ = v135
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                var v136 uint16 = 901
                                                                                                                _ = v136
                                                                                                            } else {
                                                                                                                if (((float32((float64(float32((float64((v4 + 0.0)) - 1.5))) / 0.1)) / v4) > 3.14159) || (v121 < (v121 / ((v121) | 1)))) {
                                                                                                                    if (("foo bar" + ("hello" + ("" + "hello"))) != "hello") {
                                                                                                                        switch (("" + "foo bar") + ("" + "\n\t")) {
                                                                                                                        case "":
                                                                                                                            v16 = v16
                                                                                                                        default:
                                                                                                                            for fi137 := 0; fi137 < 18; fi137++ {
                                                                                                                                fmt.Println(v16, v8, v96)
                                                                                                                                if len(v3) >= 2 {
                                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    v138 := GenS12{F0: 0}
                                                                                                                                    _ = v138
                                                                                                                                } else {
                                                                                                                                    v139 := GenS14{F0: int8(v1), F1: ""}
                                                                                                                                    fmt.Println(v1, v121, v121)
                                                                                                                                    v121 += v121
                                                                                                                                    var v140 uint8 = 1
                                                                                                                                    v140 -= (v140 * v140)
                                                                                                                                    _ = v139
                                                                                                                                    _ = v140
                                                                                                                                }
                                                                                                                                _ = fi137
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            for fi141 := 0; fi141 < 12; fi141++ {
                                                                                                                                if (("a" + "foo bar") >= "foo bar") {
                                                                                                                                    v1 += (v1 | (v1 >> ((v1) & 63)))
                                                                                                                                    var v142 int8 = v8
                                                                                                                                    if false {
                                                                                                                                        switch "" {
                                                                                                                                        case "hello":
                                                                                                                                            var v143 float32 = (v4 - 0.1)
                                                                                                                                            _ = v143
                                                                                                                                        case "foo bar":
                                                                                                                                            v3 = append(v3, v121)
                                                                                                                                        default:
                                                                                                                                            v144 := GenS15{F0: false, F1: 0}
                                                                                                                                            _ = v144
                                                                                                                                        }
                                                                                                                                        if (false && false) {
                                                                                                                                            fmt.Println(v142, fi141, v4)
                                                                                                                                            v4 = (float32((0.0 / 1e10)) - v4)
                                                                                                                                        } else {
                                                                                                                                            var v145 int = (fi141 >> ((1) & 63))
                                                                                                                                            switch "" {
                                                                                                                                            case "\n\t":
                                                                                                                                                v145 = fi95
                                                                                                                                            }
                                                                                                                                            _ = v145
                                                                                                                                        }
                                                                                                                                        v3 = append(v3, 2147483648)
                                                                                                                                    }
                                                                                                                                    _ = v142
                                                                                                                                }
                                                                                                                                v8 = (v8 / ((v8) | 1))
                                                                                                                                var v146 string = "foo bar"
                                                                                                                                _ = fi141
                                                                                                                                _ = v146
                                                                                                                            }
                                                                                                                            var v147 float64 = -2.25
                                                                                                                            if ((!(32767 <= 2)) && (128 != 2)) {
                                                                                                                                fmt.Println(v8)
                                                                                                                            } else {
                                                                                                                                v1 -= (v1 ^ v1)
                                                                                                                                for fi148 := 0; fi148 < 11; fi148++ {
                                                                                                                                    var v149 int64 = 9223372036854775807
                                                                                                                                    var v150 int32 = -509
                                                                                                                                    switch "\n\t" {
                                                                                                                                    case "\n\t":
                                                                                                                                        var v151 int32 = v150
                                                                                                                                        _ = v151
                                                                                                                                    case "a":
                                                                                                                                        v152 := GenS11{F0: "", F1: 2}
                                                                                                                                        _ = v152
                                                                                                                                    }
                                                                                                                                    if ((^(fi95)) > (-(fi95))) {
                                                                                                                                        v1 = 312
                                                                                                                                        var v153 int64 = 1
                                                                                                                                        v147 = 0.1
                                                                                                                                        _ = v153
                                                                                                                                    }
                                                                                                                                    v154 := GenS16{F0: 2, F1: ((("hello" + "a") + "hello") + "foo bar"), F2: 4294967295}
                                                                                                                                    _ = fi148
                                                                                                                                    _ = v149
                                                                                                                                    _ = v150
                                                                                                                                    _ = v154
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v147
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(v96, v16)
                                                                                                                    v1 = 42
                                                                                                                }
                                                                                                                v155 := v2["foo bar"]
                                                                                                                _ = len(v3)
                                                                                                                _ = v155
                                                                                                            }
                                                                                                            var v156 uint = 42
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                var v157 bool = ((uint8(v1) != 1) && true)
                                                                                                                _ = v157
                                                                                                            case "\n\t":
                                                                                                                switch "\n\t" {
                                                                                                                case "foo bar":
                                                                                                                    delete(v2, "a")
                                                                                                                case "":
                                                                                                                    if true {
                                                                                                                        v4 = (float32(1e10) + float32((0.0 * 0.0)))
                                                                                                                    } else {
                                                                                                                        if (true || true) {
                                                                                                                            var v158 float64 = 3.14159
                                                                                                                            _ = v158
                                                                                                                        }
                                                                                                                        fmt.Println(v16, fi95)
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    if (0.0 == 1.5) {
                                                                                                                        for fi159 := 0; fi159 < 18; fi159++ {
                                                                                                                            v96 = (v96 % ((v96) | 1))
                                                                                                                            var v160 int64 = 0
                                                                                                                            if len(v3) > 0 {
                                                                                                                                v3[0] = (v121 + 42)
                                                                                                                            }
                                                                                                                            _ = fi159
                                                                                                                            _ = v160
                                                                                                                        }
                                                                                                                        switch 294 {
                                                                                                                        case 18:
                                                                                                                            var v161 int64 = 50
                                                                                                                            _ = v161
                                                                                                                        case 245:
                                                                                                                            v1 += (v1 ^ v156)
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v162 := GenS17{F0: v8}
                                                                                                                    _ = v162
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v163 := GenS14{F0: (v8 & (v8 | (v8 * (v8 ^ v8)))), F1: ("" + "a")}
                                                                                                                _ = v163
                                                                                                            default:
                                                                                                                v96 = v96
                                                                                                            }
                                                                                                            _ = v156
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v121
                                                                                                }
                                                                                                for fi164 := 0; fi164 < 1; fi164++ {
                                                                                                    v165 := GenS18{F0: ((100 == v8) && false), F1: 0, F2: 73}
                                                                                                    var v166 uint64 = 462
                                                                                                    _ = fi164
                                                                                                    _ = v165
                                                                                                    _ = v166
                                                                                                }
                                                                                            }
                                                                                            v4 = 3.14159
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if (!(0 == 999)) {
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        v96 = v96
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi95, v1, v96)
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            var v167 bool = false
                                                                                            var v168 int = (fi95 / ((128) | 1))
                                                                                            _ = v167
                                                                                            _ = v168
                                                                                        } else {
                                                                                            fmt.Println(fi95, v8, v96)
                                                                                            var v169 int16 = (v96 << ((v1) & 63))
                                                                                            var v170 uint32 = 0
                                                                                            _ = v169
                                                                                            _ = v170
                                                                                        }
                                                                                    default:
                                                                                        v171 := GenS19{F0: 18446744073709551615, F1: -801}
                                                                                        _ = v171
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v1, v4, v1)
                                                                                    for fi172 := 0; fi172 < 5; fi172++ {
                                                                                        v16 += (v16 >> ((65535) & 63))
                                                                                        _ = fi172
                                                                                    }
                                                                                }
                                                                            }
                                                                        case -1:
                                                                            switch (fi95 * fi95) {
                                                                            case 623:
                                                                                var v173 uint8 = uint8(v16)
                                                                                _ = v173
                                                                            case 969:
                                                                                v3 = append(v3, 4294967295)
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (127 == 128) {
                                                                            var v174 uint = v1
                                                                            _ = v174
                                                                        } else {
                                                                            _ = len(v3)
                                                                            fmt.Println(v16, v1)
                                                                        }
                                                                    }
                                                                } else {
                                                                    v1 -= (v1 | v1)
                                                                    fmt.Println(v16, fi95)
                                                                }
                                                            } else {
                                                                var v175 int8 = v8
                                                                v1 -= (v1 | v1)
                                                                _ = v175
                                                            }
                                                            var v176 int64 = -1
                                                            _ = v176
                                                        }
                                                        fmt.Println(v96, v1)
                                                    }
                                                }
                                            } else {
                                                v4 -= v4
                                            }
                                            delete(v2, (("" + "foo bar") + "hello"))
                                            v177 := GenS8{F0: (v16 - (v16 | (v16 >> ((25) & 63)))), F1: "a"}
                                            _ = v177
                                        } else {
                                            var v178 uint16 = v16
                                            _ = v178
                                        }
                                        _ = v106
                                    }
                                    _ = fi95
                                    _ = v96
                                }
                            } else {
                                var v179 string = ("foo bar" + ("hello" + (("" + ("a" + "")) + "hello")))
                                v8 = int8(v16)
                                v180 := v2[v179]
                                _ = v179
                                _ = v180
                            }
                        } else {
                            var v181 uint32 = 781
                            v181 = v181
                            _ = v181
                        }
                        _ = v88
                    }
                case 4294967295:
                    v1 = v1
                case 222:
                    var v182 int32 = -227
                    _ = v182
                }
                var v183 float32 = ((v4 - (3.14159 + v4)) + v4)
                fmt.Println(v183)
                _ = v183
            } else {
                for fi184 := 0; fi184 < 18; fi184++ {
                    if false {
                        v16 -= (^(v16))
                    }
                    v3 = append(v3, 122)
                    var v185 string = ("a" + (("" + ("" + "hello")) + "foo bar"))
                    _ = fi184
                    _ = v185
                }
            }
        }
        v2["a"] = ("\n\t" + "foo bar")
    }
    for fi186 := 0; fi186 < 20; fi186++ {
        v4 = v4
        var v187 int32 = 1
        v16 -= (^(v16))
        _ = fi186
        _ = v187
    }
    v1 -= (v1 ^ v1)
    fmt.Println(v1)
    v4 += 0.0
    var v188 string = "\n\t"
    var v189 int64 = 1
    var v190 float64 = (0.0 / -2.25)
    for fi191 := 0; fi191 < 20; fi191++ {
        fmt.Println(v4)
        fmt.Println(v189, v189, v188)
        v192 := GenS4{F0: v16, F1: -1, F2: ""}
        _ = fi191
        _ = v192
    }
    var v193 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v188:", v188)
    fmt.Println("v189:", v189)
    fmt.Println("v190:", v190)
    fmt.Println("v193:", v193)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v188
    _ = v189
    _ = v190
    _ = v193
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
