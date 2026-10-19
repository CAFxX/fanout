package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 uint64; F1 uint8; F2 bool }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 bool = (4294967295 == -9223372036854775808)
    var v2 uint16 = 32767
    v3 := []uint32{4294967295, 0}
    var v4 int16 = 1
    v5 := []int{32767, 9223372036854775807, -928, 743}
    v6 := make(map[uint32]uint32)
    _ = len(v6)
    _ = len(v3)
    switch 32768 {
    case 974:
        var v7 uint16 = 3
        _ = v7
    case -950:
        _ = len(v6)
    default:
        v5 = append(v5, -340)
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 441
                }
            }
            _ = fi9
        }
        _ = fi8
    }
    v10 := GenS1{F0: -124, F1: 1}
    if v1 {
        if false {
            for fi11 := 0; fi11 < 10; fi11++ {
                switch (^(fi11)) {
                case -9223372036854775808:
                    if v1 {
                        var v12 uint = 1
                        v4 += (^(v4))
                        v4 = (v4 % ((1000) | 1))
                        _ = v12
                    }
                }
                var v13 int64 = 0
                _ = fi11
                _ = v13
            }
            if len(v3) > 0 {
                v3[0] = 271
            }
            var v14 string = ""
            _ = v14
        } else {
            switch "a" {
            case "hello":
                if v1 {
                    v4 = -848
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v16 := GenS1{F0: -128, F1: -436}
                        _ = fi15
                        _ = v16
                    }
                    if true {
                        var v17 int64 = 0
                        v17 += v17
                        var v18 int = 0
                        _ = v17
                        _ = v18
                    }
                }
            }
            fmt.Println(v2, v4)
            if v1 {
                fmt.Println(v4)
                switch "foo bar" {
                case "\n\t":
                    var v19 int16 = (v4 / (((v4 >> ((65535) & 63))) | 1))
                    _ = v19
                case "a":
                    var v20 bool = (!v1)
                    _ = v20
                }
                v4 += (v4 / ((v4) | 1))
            } else {
                v5 = append(v5, 42)
                var v21 uint64 = 887
                v6[42] = 65535
                _ = v21
            }
        }
        var v22 uint8 = 42
        for fi23 := 0; fi23 < 17; fi23++ {
            switch ("hello" + "hello") {
            case "hello":
                if (881 <= 2) {
                    var v24 int32 = -1
                    _ = v24
                }
            case "foo bar":
                switch "a" {
                case "hello":
                    v2 += 790
                case "a":
                    for fi25 := 0; fi25 < 18; fi25++ {
                        switch fi25 {
                        case -2147483648:
                            v26 := GenS1{F0: 0, F1: -1}
                            _ = v26
                        case 438:
                            fmt.Println(v1)
                        default:
                            switch (fi23 << ((4294967295) & 63)) {
                            case 194:
                                v1 = false
                            case 729:
                                var v27 string = "hello"
                                _ = v27
                            case -704:
                                var v28 string = "hello"
                                _ = v28
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v2, v4)
                        default:
                            fmt.Println(fi23, v2)
                        }
                        if (65535 == 2) {
                            var v29 int = (fi25 - -9223372036854775808)
                            switch v29 {
                            case -21:
                                switch "hello" {
                                case "hello":
                                    delete(v6, 2)
                                case "a":
                                    v1 = v1
                                }
                            case 2147483647:
                                var v30 float32 = float32(3.14159)
                                var v31 bool = false
                                _ = v30
                                _ = v31
                            case 183:
                                var v32 float64 = 0.1
                                _ = v32
                            default:
                                v6[4294967295] = 65535
                            }
                            var v33 uint = uint(v22)
                            _ = v29
                            _ = v33
                        }
                        _ = fi25
                    }
                case "":
                    v2 += v2
                }
            case "":
                var v34 uint = 65535
                _ = v34
            }
            v4 -= 1
            _ = fi23
        }
        _ = v22
    }
    if v1 {
        switch "a" {
        case "a":
            for fi35 := 0; fi35 < 5; fi35++ {
                v1 = (455 <= 65535)
                var v36 string = "a"
                for fi37 := 0; fi37 < 3; fi37++ {
                    v38 := GenS2{F0: 753, F1: "", F2: 188}
                    if (0.0 == 1.5) {
                        v5 = append(v5, fi37)
                        switch -116 {
                        case -2147483648:
                            var v39 string = v36
                            _ = v39
                        case 962:
                            fmt.Println(fi35, v2)
                        case 2147483647:
                            v40 := v6[65535]
                            _ = v40
                        default:
                            var v41 float32 = float32(1.5)
                            _ = v41
                        }
                        v42 := GenS1{F0: 0, F1: 9223372036854775807}
                        _ = v42
                    }
                    _ = fi37
                    _ = v38
                }
                _ = fi35
                _ = v36
            }
        case "\n\t":
            if (("" + ("" + "")) > "") {
                if ((float64((0.0 + 0.1)) / 0.1) != (1.5 - ((-2.25 * (((((-2.25 + (((0.1 / 1.0) - (float64(1e10) + 1e10)) / (3.14159 * 0.1))) + 0.0) * 1e10) - 0.0) * (-2.25 / 1.0))) / float64(-2.25)))) {
                    var v43 float64 = (1.5 * (-2.25 / 1e10))
                    var v44 int32 = 0
                    if (((4294967295 > 42) && ("a" < (("a" + ("foo bar" + "foo bar")) + "foo bar"))) && false) {
                        fmt.Println(v44, v2, v43)
                        v45 := GenS3{F0: v4, F1: 127}
                        if (128 >= 657) {
                            for fi46 := 0; fi46 < 2; fi46++ {
                                for fi47 := 0; fi47 < 18; fi47++ {
                                    if len(v5) > 0 {
                                        v5[0] = (-(fi47))
                                    }
                                    v48 := GenS4{F0: 0, F1: int8(v4), F2: 4294967295}
                                    v49 := GenS3{F0: 42, F1: 1}
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                                var v50 int8 = 42
                                v50 = (v50 % ((-44) | 1))
                                _ = fi46
                                _ = v50
                            }
                        }
                        _ = v45
                    } else {
                        if (v2 < (v2 ^ (v2 >> ((255) & 63)))) {
                            var v51 uint64 = 1
                            v52 := GenS4{F0: 1, F1: 100, F2: 58}
                            _ = v51
                            _ = v52
                        }
                        if false {
                            v53 := GenS2{F0: 0, F1: "a", F2: 87}
                            v54 := GenS5{F0: (v2 & 32767)}
                            var v55 int32 = (v44 + v44)
                            _ = v53
                            _ = v54
                            _ = v55
                        } else {
                            v56 := v6[2147483647]
                            for fi57 := 0; fi57 < 16; fi57++ {
                                v1 = ((fi57 * (-(fi57))) < int(v2))
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    if v1 {
                                        for fi58 := 0; fi58 < 9; fi58++ {
                                            var v59 int64 = 42
                                            var v60 float64 = (((float64(float32(v43)) / float64((0.1 / float32((((1e10 * float64(0.0)) + (0.1 * v43)) * 0.0))))) + (1e10 / v43)) / 3.14159)
                                            v59 += v59
                                            _ = fi58
                                            _ = v59
                                            _ = v60
                                        }
                                        var v61 uint64 = 42
                                        _ = v61
                                    } else {
                                        if v1 {
                                            v62 := GenS6{F0: v1, F1: 65535}
                                            _ = v62
                                        } else {
                                            var v63 int64 = -262
                                            v4 -= (-(v4))
                                            _ = v63
                                        }
                                        for fi64 := 0; fi64 < 1; fi64++ {
                                            fmt.Println(v44, v44, v1)
                                            _ = fi64
                                        }
                                    }
                                case "hello":
                                    v56 = 2147483647
                                case "a":
                                    var v65 int64 = -1
                                    _ = v65
                                }
                                for fi66 := 0; fi66 < 19; fi66++ {
                                    v44 -= (-(v44))
                                    v67 := GenS4{F0: v44, F1: 42, F2: (v56 << ((0) & 63))}
                                    var v68 uint32 = (v56 / ((v56) | 1))
                                    _ = fi66
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi57
                            }
                            _ = v56
                        }
                    }
                    _ = v43
                    _ = v44
                } else {
                    switch 4294967295 {
                    case 128:
                        _ = len(v6)
                    }
                    if (v1 && true) {
                        v69 := GenS2{F0: 2147483647, F1: "a", F2: 127}
                        fmt.Println(v4, v4)
                        _ = v69
                    } else {
                        if (100 != -118) {
                            switch "a" {
                            case "hello":
                                v70 := GenS7{F0: 28, F1: 4, F2: 42}
                                _ = v70
                            case "":
                                var v71 int = 1
                                _ = v71
                            case "\n\t":
                                v2 = 0
                            }
                            if (v1 && true) {
                                switch ("" + ("foo bar" + "hello")) {
                                case "":
                                    var v72 uint = 65535
                                    _ = v72
                                case "hello":
                                    if (float64((-2.25 + 1e10)) != (3.14159 / float64((float32((0.1 - -2.25)) - -2.25)))) {
                                        for fi73 := 0; fi73 < 3; fi73++ {
                                            switch "hello" {
                                            case "\n\t":
                                                var v74 bool = v1
                                                _ = v74
                                            case "":
                                                var v75 int16 = 32767
                                                _ = v75
                                            case "foo bar":
                                                v1 = (false && (false && (4294967295 >= 318)))
                                            default:
                                                if (int32(fi73) < 42) {
                                                    v6[0] = 42
                                                } else {
                                                    v2 += 0
                                                    var v76 int8 = 42
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v76
                                                }
                                            }
                                            v4 += (v4 & v4)
                                            var v77 int32 = -1
                                            _ = fi73
                                            _ = v77
                                        }
                                        var v78 int16 = (v4 << ((1) & 63))
                                        var v79 int16 = v78
                                        _ = v78
                                        _ = v79
                                    }
                                default:
                                    switch 4294967295 {
                                    case -184:
                                        var v80 uint32 = 42
                                        _ = v80
                                    case 42:
                                        var v81 int16 = -1
                                        _ = v81
                                    default:
                                        switch -1 {
                                        case 2147483647:
                                            fmt.Println(v2, v2, v4)
                                        case 127:
                                            for fi82 := 0; fi82 < 18; fi82++ {
                                                v83 := GenS3{F0: v4, F1: -128}
                                                _ = fi82
                                                _ = v83
                                            }
                                        default:
                                            _ = len(v6)
                                        }
                                    }
                                }
                                var v84 uint = 18446744073709551615
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    v4 = v4
                                case "a":
                                    switch 77 {
                                    case 9223372036854775807:
                                        _ = len(v6)
                                    case 1:
                                        v85 := GenS8{F0: 2}
                                        _ = v85
                                    case -9223372036854775808:
                                        if false {
                                            for fi86 := 0; fi86 < 13; fi86++ {
                                                v1 = v1
                                                _ = fi86
                                            }
                                            v2 = v2
                                            fmt.Println(v84)
                                        }
                                    default:
                                        switch ((("a" + "") + ("a" + (("foo bar" + "foo bar") + "hello"))) + ((("foo bar" + "a") + "hello") + "a")) {
                                        case "foo bar":
                                            v87 := GenS9{F0: 127, F1: 4294967295}
                                            _ = v87
                                        case "\n\t":
                                            switch 4294967295 {
                                            case 0:
                                                if false {
                                                    var v88 int16 = 1
                                                    if (!v1) {
                                                        var v89 uint32 = 1
                                                        v6[v89] = v89
                                                        _ = v89
                                                    } else {
                                                        var v90 uint32 = 2147483647
                                                        fmt.Println(v84)
                                                        if (true && (!((v84 <= (v84 - (^(v84)))) && (v84 == v84)))) {
                                                            var v91 float64 = 0.1
                                                            if (true && true) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 328
                                                                }
                                                                var v92 float64 = 0.0
                                                                _ = v92
                                                            } else {
                                                                fmt.Println(v84)
                                                            }
                                                            _ = v91
                                                        } else {
                                                            v93 := GenS2{F0: (v90 % ((2147483648) | 1)), F1: "\n\t", F2: 152}
                                                            _ = v93
                                                        }
                                                        _ = v90
                                                    }
                                                    if v1 {
                                                        if len(v3) > 0 {
                                                            v3[0] = 221
                                                        }
                                                        fmt.Println(v88, v84)
                                                    } else {
                                                        if v1 {
                                                            if (-84 != -128) {
                                                                if len(v5) > 0 {
                                                                    v5[0] = -9223372036854775808
                                                                }
                                                            }
                                                            v94 := GenS3{F0: 32767, F1: 1}
                                                            _ = v94
                                                        }
                                                        var v95 uint64 = 3
                                                        _ = v95
                                                    }
                                                    _ = v88
                                                }
                                            case 532:
                                                v96 := GenS10{F0: 4294967295, F1: 200, F2: true}
                                                _ = v96
                                            }
                                        }
                                    }
                                case "":
                                    for fi97 := 0; fi97 < 4; fi97++ {
                                        if v1 {
                                            if v1 {
                                                v98 := GenS11{F0: ("" + "\n\t"), F1: uint8(v84)}
                                                _ = v98
                                            }
                                        }
                                        for fi99 := 0; fi99 < 14; fi99++ {
                                            v100 := GenS10{F0: 2, F1: 2, F2: (42 >= -403)}
                                            var v101 uint = uint(v2)
                                            var v102 int16 = (v4 | -1)
                                            _ = fi99
                                            _ = v100
                                            _ = v101
                                            _ = v102
                                        }
                                        _ = fi97
                                    }
                                }
                                _ = v84
                            } else {
                                v6[2147483648] = 781
                            }
                            var v103 uint64 = 153
                            _ = v103
                        }
                    }
                    if (true || v1) {
                        switch 2147483647 {
                        case 4294967295:
                            v104 := GenS6{F0: ("hello" != "a"), F1: 4294967295}
                            _ = v104
                        case 0:
                            var v105 uint8 = 2
                            _ = v105
                        default:
                            if v1 {
                                if len(v3) > 0 {
                                    v3[0] = 531
                                }
                                v2 += 338
                            } else {
                                fmt.Println(v4, v1, v2)
                            }
                        }
                        var v106 int8 = 0
                        _ = v106
                    }
                }
                for fi107 := 0; fi107 < 20; fi107++ {
                    if (0 == v2) {
                        v4 -= -73
                    }
                    _ = fi107
                }
            }
        default:
            var v108 uint16 = v2
            _ = v108
        }
        switch -768 {
        case 297:
            _ = len(v5)
        default:
            for fi109 := 0; fi109 < 2; fi109++ {
                switch ("a" + ("" + "foo bar")) {
                case "":
                    fmt.Println(v2)
                }
                v4 = v4
                var v110 bool = v1
                _ = fi109
                _ = v110
            }
        }
    } else {
        switch 9223372036854775807 {
        case 4294967295:
            var v111 int8 = 0
            _ = v111
        default:
            var v112 int16 = (v4 << ((42) & 63))
            _ = v112
        }
    }
    fmt.Println(v2)
    var v113 int8 = -4
    for fi114 := 0; fi114 < 5; fi114++ {
        if ((1e10 * (1.5 - 1.5)) != 0.0) {
            v1 = ((!v1) || ("" != ("\n\t" + "")))
        }
        v115 := GenS3{F0: (-(v4)), F1: int8(fi114)}
        v116 := GenS12{F0: 2147483647}
        _ = fi114
        _ = v115
        _ = v116
    }
    if (291 != -896) {
        var v117 uint32 = 42
        var v118 uint = 0
        var v119 float32 = (0.0 + 3.14159)
        _ = v117
        _ = v118
        _ = v119
    } else {
        v120 := GenS1{F0: 127, F1: 218}
        v2 = (v2 * 255)
        v1 = v1
        _ = v120
    }
    v113 = (v113 * v113)
    v121 := v6[2]
    var v122 int16 = 32767
    var v123 string = ("hello" + "foo bar")
    switch 139 {
    case 32767:
        if true {
            for fi124 := 0; fi124 < 3; fi124++ {
                v2 = v2
                _ = fi124
            }
            var v125 int32 = 42
            _ = v125
        }
    case 883:
        v4 += -828
    default:
        var v126 int8 = (v113 + (v113 * v113))
        _ = v126
    }
    fmt.Println(v123, v123, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v113
    _ = v121
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
