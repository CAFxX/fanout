package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 int64; F2 int64 }
type GenS6 struct { F0 string; F1 bool; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 uint16; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS13 struct { F0 uint16; F1 int }

func main() {
    var v1 string = ("a" + "hello")
    v2 := []uint32{2147483647, 2}
    v3 := make(map[string]uint32)
    var v4 uint64 = 4294967295
    v4 -= v4
    v1 = "hello"
    var v5 int = 2147483647
    if false {
        var v6 bool = (!(255 == 15))
        _ = v6
    } else {
        if (!(0 == 100)) {
            v4 += (v4 | (v4 - (v4 ^ v4)))
        } else {
            var v7 int8 = 116
            _ = len(v2)
            var v8 int32 = 1
            _ = v7
            _ = v8
        }
        v3[(("a" + "a") + "\n\t")] = 0
    }
    switch (("hello" + (v1 + v1)) + "\n\t") {
    case "":
        switch (-(v5)) {
        case 0:
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := v3["foo bar"]
                fmt.Println(fi9, v5)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi9
                _ = v10
            }
        case -9223372036854775808:
            if (-2.25 <= 0.1) {
                switch -9223372036854775808 {
                case 0:
                    v2 = append(v2, 2)
                case 4294967295:
                    var v11 float32 = ((0.1 / 0.1) * 1e10)
                    _ = v11
                }
                v12 := GenS1{F0: (v5 % ((v5) | 1)), F1: 594}
                if (!(0 >= 310)) {
                    fmt.Println(v1, v5)
                    switch v1 {
                    case "a":
                        if false {
                            if (18446744073709551615 < 0) {
                                var v13 int16 = 32767
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    var v15 int64 = -9223372036854775808
                                    _ = len(v2)
                                    _ = fi14
                                    _ = v15
                                }
                                _ = v13
                            } else {
                                v1 = (((("foo bar" + v1) + v1) + "a") + "a")
                            }
                            var v16 int64 = int64(v5)
                            for fi17 := 0; fi17 < 15; fi17++ {
                                for fi18 := 0; fi18 < 9; fi18++ {
                                    switch (fi17 | (fi17 << ((752) & 63))) {
                                    case 127:
                                        v19 := v3[v1]
                                        _ = v19
                                    case 32767:
                                        var v20 int8 = 50
                                        _ = v20
                                    default:
                                        v21 := GenS2{F0: 42}
                                        _ = v21
                                    }
                                    v1 = v1
                                    fmt.Println(fi17, v4, fi18)
                                    _ = fi18
                                }
                                for fi22 := 0; fi22 < 9; fi22++ {
                                    if (v1 == v1) {
                                        v23 := GenS3{F0: v16, F1: v1, F2: 2147483647}
                                        _ = v23
                                    }
                                    _ = fi22
                                }
                                switch v1 {
                                case "foo bar":
                                    switch v1 {
                                    case "":
                                        if (float32(0.0) < 1e10) {
                                            v1 = "hello"
                                            if (false || (false || (2 > 255))) {
                                                var v24 uint = 823
                                                var v25 uint64 = 1
                                                var v26 bool = (!true)
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v27 int64 = (-(v16))
                                                v1 = ((v1 + "foo bar") + (v1 + v1))
                                                _ = v27
                                            }
                                            for fi28 := 0; fi28 < 12; fi28++ {
                                                v2 = append(v2, 529)
                                                _ = fi28
                                            }
                                        } else {
                                            for fi29 := 0; fi29 < 16; fi29++ {
                                                var v30 uint8 = 255
                                                _ = fi29
                                                _ = v30
                                            }
                                            fmt.Println(fi17, v4, v5)
                                        }
                                    case "foo bar":
                                        if false {
                                            v31 := GenS3{F0: v16, F1: "a", F2: 0}
                                            _ = v31
                                        }
                                    }
                                case "\n\t":
                                    var v32 uint8 = 113
                                    _ = v32
                                }
                                _ = fi17
                            }
                            _ = v16
                        }
                    default:
                        v3[""] = 65535
                    }
                    v4 += 923
                } else {
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v34 := GenS4{F0: -32768}
                        _ = fi33
                        _ = v34
                    }
                    if ((v4 >= 205) || true) {
                        for fi35 := 0; fi35 < 4; fi35++ {
                            for fi36 := 0; fi36 < 13; fi36++ {
                                _ = len(v3)
                                var v37 uint = 18446744073709551615
                                _ = fi36
                                _ = v37
                            }
                            _ = fi35
                        }
                        for fi38 := 0; fi38 < 12; fi38++ {
                            var v39 uint = 255
                            delete(v3, v1)
                            _ = fi38
                            _ = v39
                        }
                    } else {
                        var v40 float64 = (3.14159 * (0.0 * (0.1 * 3.14159)))
                        var v41 uint = 0
                        v42 := GenS1{F0: v5, F1: 42}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                _ = v12
            }
        case 4294967295:
            for fi43 := 0; fi43 < 7; fi43++ {
                if false {
                    for fi44 := 0; fi44 < 7; fi44++ {
                        _ = len(v2)
                        _ = fi44
                    }
                } else {
                    for fi45 := 0; fi45 < 18; fi45++ {
                        var v46 int32 = 0
                        switch (-(v5)) {
                        case -521:
                            v1 = "\n\t"
                            if (4294967295 < 2147483647) {
                                fmt.Println(v46)
                                fmt.Println(v46)
                                switch v1 {
                                case "a":
                                    v4 = (v4 - 2)
                                case "hello":
                                    delete(v3, "\n\t")
                                default:
                                    v47 := GenS2{F0: 261}
                                    _ = v47
                                }
                            } else {
                                v48 := v3[v1]
                                v3["hello"] = v48
                                fmt.Println(v5, v1, v5)
                                _ = v48
                            }
                        case 4294967295:
                            if ((3.14159 < -2.25) && (1 <= (v4 - 0))) {
                                switch "hello" {
                                case "":
                                    v46 = v46
                                case "foo bar":
                                    v49 := GenS5{F0: true, F1: 42, F2: -510}
                                    if (false && (61 <= 21)) {
                                        v4 += v4
                                    } else {
                                        var v50 int = (v5 >> ((65535) & 63))
                                        var v51 bool = (!true)
                                        var v52 uint8 = 255
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 uint16 = 2
                                    fmt.Println(v5, v5, v4)
                                    _ = v49
                                    _ = v53
                                default:
                                    if true {
                                        fmt.Println(fi43)
                                        v46 += (v46 + v46)
                                    } else {
                                        var v54 float32 = float32(3.14159)
                                        _ = v54
                                    }
                                }
                                switch v1 {
                                case "\n\t":
                                    v4 = 0
                                case "hello":
                                    v4 -= v4
                                default:
                                    var v55 float32 = ((3.14159 + (3.14159 * (0.1 * 0.1))) - float32((1.5 * -2.25)))
                                    _ = v55
                                }
                                var v56 bool = ((!(-39 == -32768)) && (100 > -80))
                                v5 -= (fi43 * (fi43 / (((v5 | 32768)) | 1)))
                                _ = v56
                            }
                        case -237:
                            v57 := GenS6{F0: "hello", F1: true, F2: 432}
                            _ = v57
                        default:
                            var v58 int8 = 42
                            _ = v58
                        }
                        _ = fi45
                        _ = v46
                    }
                }
                _ = fi43
            }
        }
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        v1 = "\n\t"
    default:
        if (!(240 != 127)) {
            switch (v5 * v5) {
            case 1:
                delete(v3, (v1 + (v1 + v1)))
            default:
                v4 = (v4 << ((18446744073709551615) & 63))
            }
            v1 = ("foo bar" + ((((v1 + v1) + "foo bar") + v1) + v1))
            v4 = 513
        } else {
            v5 = 127
            v59 := GenS7{F0: -1}
            switch "a" {
            case "\n\t":
                if (0.0 >= 3.14159) {
                    var v60 float32 = ((((1.5 + 1.5) - 1.5) * 3.14159) * 1.5)
                    var v61 uint8 = 255
                    _ = v60
                    _ = v61
                } else {
                    var v62 int8 = 127
                    v63 := GenS5{F0: true, F1: int64(v5), F2: 9223372036854775807}
                    for fi64 := 0; fi64 < 19; fi64++ {
                        switch ("a" + v1) {
                        case "foo bar":
                            for fi65 := 0; fi65 < 14; fi65++ {
                                v3[v1] = 292
                                switch (fi65 - 127) {
                                case 886:
                                    v4 -= (v4 << ((42) & 63))
                                    var v66 int = (fi65 / (((v5 >> ((255) & 63))) | 1))
                                    v67 := v3[v1]
                                    v68 := GenS3{F0: 9223372036854775807, F1: v1, F2: 2147483648}
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                case -1:
                                    if true {
                                        fmt.Println(fi65)
                                        v1 = v1
                                        switch fi64 {
                                        case 1:
                                            if (!(2 <= 557)) {
                                                v69 := GenS3{F0: 1, F1: "foo bar", F2: 516}
                                                var v70 int16 = 42
                                                v1 = v1
                                                _ = v69
                                                _ = v70
                                            } else {
                                                if false {
                                                    v1 = v1
                                                    if (!false) {
                                                        v5 = (fi65 / ((v5) | 1))
                                                    } else {
                                                        v62 -= v62
                                                        v62 += 0
                                                    }
                                                    if (18446744073709551615 < 898) {
                                                        if len(v2) > 0 {
                                                            v2[0] = 2
                                                        }
                                                    }
                                                    if ((true || (v1 == ("foo bar" + v1))) && false) {
                                                        var v71 float64 = 1.5
                                                        v1 = "foo bar"
                                                        v72 := GenS1{F0: 2147483647, F1: 512}
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        fmt.Println(v1)
                                                        v62 = 100
                                                        _ = len(v3)
                                                        if false {
                                                            v3[v1] = 2147483648
                                                            v3[(v1 + v1)] = 768
                                                        } else {
                                                            v73 := GenS7{F0: v62}
                                                            var v74 uint8 = 128
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        var v75 string = ""
                                                        var v76 uint16 = 0
                                                        var v77 int16 = 42
                                                        v78 := v3[""]
                                                        if false {
                                                            v79 := GenS2{F0: v76}
                                                            switch v5 {
                                                            case 4294967295:
                                                                v78 -= v78
                                                            case 1:
                                                                _ = len(v2)
                                                            case -501:
                                                                var v80 uint16 = (v76 + (v76 * v76))
                                                                _ = v80
                                                            }
                                                            v81 := GenS8{F0: -1, F1: (^(v76)), F2: (v1 + "foo bar")}
                                                            var v82 uint32 = 2147483647
                                                            switch 730 {
                                                            case 165:
                                                                fmt.Println(v62, v75, v78)
                                                            default:
                                                                v83 := GenS9{F0: 1}
                                                                _ = v83
                                                            }
                                                            v84 := GenS10{F0: (v76 + v76)}
                                                            delete(v3, ((v1 + v75) + (v1 + v75)))
                                                            _ = v79
                                                            _ = v81
                                                            _ = v82
                                                            _ = v84
                                                        } else {
                                                            v75 = v75
                                                            v85 := GenS11{F0: 2147483647, F1: (v77 ^ -32768)}
                                                            if true {
                                                                var v86 uint64 = 4294967295
                                                                switch v5 {
                                                                case 4294967295:
                                                                    var v87 float32 = float32(0.1)
                                                                    _ = v87
                                                                case 0:
                                                                    v88 := GenS2{F0: (^(v76))}
                                                                    if true {
                                                                        var v89 int32 = 0
                                                                        _ = v89
                                                                    } else {
                                                                        if (false && false) {
                                                                            v75 = "\n\t"
                                                                            v90 := GenS4{F0: (v77 ^ v77)}
                                                                            var v91 uint32 = (^(v78))
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                    _ = v88
                                                                case -714:
                                                                    delete(v3, "hello")
                                                                }
                                                                v78 = 65535
                                                                _ = v86
                                                            }
                                                            var v92 uint64 = v4
                                                            var v93 int8 = (v62 >> ((272) & 63))
                                                            _ = v85
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                        var v94 float32 = 1e10
                                                        var v95 uint = 472
                                                        if true {
                                                            v76 = v76
                                                            var v96 float64 = -2.25
                                                            _ = v96
                                                        }
                                                        switch "a" {
                                                        case "":
                                                            switch "hello" {
                                                            case "hello":
                                                                v1 = v1
                                                            case "":
                                                                if false {
                                                                    switch ("a" + v1) {
                                                                    case "hello":
                                                                        var v97 int16 = v77
                                                                        _ = v97
                                                                    case "":
                                                                        var v98 uint16 = v76
                                                                        _ = v98
                                                                    case "\n\t":
                                                                        switch (v5 | (-(v5))) {
                                                                        case -1:
                                                                            v99 := GenS1{F0: fi65, F1: 42}
                                                                            _ = v99
                                                                        }
                                                                    default:
                                                                        var v100 uint8 = 127
                                                                        _ = v100
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v101 uint64 = v4
                                                                _ = v101
                                                            default:
                                                                var v102 float64 = float64(0.1)
                                                                _ = v102
                                                            }
                                                            if ((v77 / ((v77) | 1)) > (v77 << ((v95) & 63))) {
                                                                var v103 uint = (v95 + 42)
                                                                _ = v103
                                                            }
                                                            v104 := GenS9{F0: 0}
                                                            if ((true || (748 == 1)) || (int8(v76) >= (v62 & v62))) {
                                                                var v105 uint = 1
                                                                _ = v105
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (v78 ^ 42)
                                                                }
                                                            }
                                                            var v106 int8 = (^(v62))
                                                            _ = v104
                                                            _ = v106
                                                        case "hello":
                                                            v3["foo bar"] = 34
                                                        case "a":
                                                            switch fi64 {
                                                            case -9223372036854775808:
                                                                var v107 uint32 = v78
                                                                _ = v107
                                                            case -733:
                                                                var v108 uint32 = (v78 | (^(v78)))
                                                                _ = v108
                                                            default:
                                                                if true {
                                                                    var v109 uint = v95
                                                                    _ = v109
                                                                }
                                                            }
                                                        }
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                        _ = v78
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    v4 += (v4 | 813)
                                                }
                                                if false {
                                                    delete(v3, "")
                                                    v62 = (v62 - v62)
                                                    switch ("hello" + "") {
                                                    case "\n\t":
                                                        var v110 int32 = 757
                                                        _ = v110
                                                    case "foo bar":
                                                        if (0 < 42) {
                                                            var v111 uint64 = (^(v4))
                                                            v111 -= 0
                                                            v4 -= 130
                                                            _ = v111
                                                        }
                                                    case "hello":
                                                        _ = len(v2)
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = 489
                                                    }
                                                }
                                                if (v62 >= (-(v62))) {
                                                    var v112 uint64 = (v4 * (v4 >> ((2) & 63)))
                                                    _ = v112
                                                } else {
                                                    var v113 int32 = -542
                                                    var v114 int = (^(fi64))
                                                    if len(v2) > 0 {
                                                        v2[0] = 65535
                                                    }
                                                    var v115 bool = (0.1 == -2.25)
                                                    var v116 uint64 = v4
                                                    v117 := GenS12{F0: 117, F1: 42, F2: 2}
                                                    switch v5 {
                                                    case 42:
                                                        v4 = v116
                                                    case -660:
                                                        var v118 uint = 4294967295
                                                        v119 := v3["hello"]
                                                        if ((!v115) && true) {
                                                            if (!false) {
                                                                var v120 int64 = 1
                                                                _ = v120
                                                            }
                                                            v121 := v3[v1]
                                                            _ = v121
                                                        }
                                                        _ = v118
                                                        _ = v119
                                                    case -2147483648:
                                                        var v122 string = v1
                                                        _ = v122
                                                    }
                                                    v115 = v115
                                                    _ = v113
                                                    _ = v114
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                }
                                            }
                                            v62 -= 1
                                            var v123 uint64 = (^(v4))
                                            var v124 uint = 0
                                            _ = v123
                                            _ = v124
                                        }
                                    }
                                case -9223372036854775808:
                                    v125 := GenS5{F0: (!(1 < 2147483647)), F1: 402, F2: 608}
                                    _ = v125
                                }
                                var v126 float32 = (3.14159 + -2.25)
                                _ = fi65
                                _ = v126
                            }
                        case "":
                            var v127 int32 = -398
                            _ = v127
                        case "a":
                            v5 += fi64
                        }
                        v128 := v3[v1]
                        _ = fi64
                        _ = v128
                    }
                    _ = v62
                    _ = v63
                }
            case "hello":
                v129 := GenS1{F0: (^(v5)), F1: -9223372036854775808}
                _ = v129
            default:
                var v130 float32 = (0.0 + float32((0.0 - 0.0)))
                _ = v130
            }
            _ = v59
        }
    }
    if true {
        switch v5 {
        case 9223372036854775807:
            _ = len(v3)
        case 126:
            switch (v5 >> ((1) & 63)) {
            case 0:
                fmt.Println(v4, v4)
            case 2147483647:
                v131 := GenS2{F0: 0}
                _ = v131
            case 128:
                if false {
                    v5 = 128
                    v4 -= 103
                } else {
                    v4 = (v4 ^ (v4 - v4))
                    var v132 float64 = (1.5 * 1.5)
                    _ = v132
                }
            default:
                v133 := GenS7{F0: 1}
                _ = v133
            }
        case -2147483648:
            fmt.Println(v1, v5)
        default:
            v134 := GenS4{F0: 32767}
            _ = v134
        }
        v4 -= v4
    } else {
        v4 = 621
        fmt.Println(v1, v1)
        var v135 int32 = 1
        _ = v135
    }
    var v136 uint8 = 1
    switch (v5 % ((v5) | 1)) {
    case -9223372036854775808:
        for fi137 := 0; fi137 < 14; fi137++ {
            _ = len(v2)
            v5 -= fi137
            _ = fi137
        }
    default:
        var v138 string = "hello"
        _ = v138
    }
    fmt.Println(v1, v1)
    v5 = (v5 | (v5 << ((255) & 63)))
    switch v1 {
    case "foo bar":
        v139 := GenS13{F0: 859, F1: (v5 + (-(v5)))}
        _ = v139
    case "hello":
        v3[v1] = 0
    default:
        v140 := v3[(v1 + v1)]
        _ = v140
    }
    fmt.Println("v1:", v1)
    fmt.Println("v136:", v136)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
