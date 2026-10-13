package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint; F1 uint16; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS8 struct { F0 uint16; F1 bool; F2 int }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 int64 }
type GenS12 struct { F0 int64; F1 int; F2 uint8 }
type GenS13 struct { F0 string; F1 string; F2 uint64 }
type GenS14 struct { F0 uint32 }
func genFn0(p0 int, p1 uint16, p2 uint) uint8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: 394, F1: 42}
        v6 := []int{fi4, (fi4 ^ p0), 4294967295}
        p1 += p1
        _ = fi4
        _ = v5
        _ = v6
    }
    p1 = uint16(p2)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{(p0 + 2147483647), (p0 & p0)}
        if ((1.5 * (1.5 - 0.0)) >= 0.1) {
            v8 = append(v8, (p0 & p0))
            if (0.0 == 0.1) {
                var v9 int64 = -1
                var v10 float32 = -2.25
                switch fi7 {
                case -586:
                    p1 = p1
                case 151:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        switch "hello" {
                        case "":
                            v12 := GenS1{F0: 1, F1: v9}
                            switch p0 {
                            case 66:
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                            _ = v12
                        case "hello":
                            if (-32768 < 626) {
                                v13 := GenS2{F0: -2147483648, F1: (^(fi11))}
                                _ = len(v8)
                                _ = v13
                            } else {
                                v8 = append(v8, (p0 - fi7))
                                var v14 uint16 = 42
                                _ = v14
                            }
                        }
                        var v15 int = (fi11 % ((42) | 1))
                        v16 := GenS3{F0: 2}
                        _ = fi11
                        _ = v15
                        _ = v16
                    }
                case 4294967295:
                    if (-1 == 2147483647) {
                        if true {
                            var v17 uint32 = 0
                            var v18 float32 = float32(1.5)
                            _ = v17
                            _ = v18
                        } else {
                            for fi19 := 0; fi19 < 19; fi19++ {
                                if false {
                                    if (!true) {
                                        switch "" {
                                        case "\n\t":
                                            v8 = append(v8, 0)
                                        case "":
                                            v8 = append(v8, fi19)
                                        case "a":
                                            var v20 string = ""
                                            _ = v20
                                        }
                                        v21 := GenS4{F0: -700, F1: 2147483647}
                                        _ = v21
                                    }
                                }
                                _ = fi19
                            }
                        }
                    } else {
                        p1 += p1
                        switch "hello" {
                        case "\n\t":
                            var v22 string = "\n\t"
                            _ = v22
                        case "hello":
                            for fi23 := 0; fi23 < 16; fi23++ {
                                v9 += 42
                                _ = fi23
                            }
                        case "":
                            v10 -= (v10 * float32(0.0))
                        }
                        var v24 string = ""
                        _ = v24
                    }
                default:
                    v25 := GenS1{F0: 0, F1: (v9 + (v9 / ((42) | 1)))}
                    _ = v25
                }
                _ = v10
                _ = v9
            }
        } else {
            _ = len(v8)
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int32 = 745
                if true {
                    var v28 uint16 = p1
                    _ = v28
                } else {
                    var v29 uint32 = 2147483647
                    v29 -= (v29 - v29)
                    _ = v29
                }
                p2 -= (p2 % ((p2) | 1))
                _ = fi26
                _ = v27
            }
        }
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 bool = true
            fmt.Println(p0, fi7, p2)
            var v32 int8 = 0
            var v33 int32 = -1
            v34 := GenS1{F0: int32(fi7), F1: 118}
            var v35 uint64 = 1
            if ("" <= "a") {
                fmt.Println(p1)
            } else {
                v36 := GenS5{F0: p2, F1: 255, F2: "foo bar"}
                _ = v36
            }
            v37 := GenS6{F0: (p2 & p2)}
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
            _ = v34
            _ = v35
            _ = v37
        }
        _ = fi7
        _ = v8
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        v39 := []int{(p0 << ((42) & 63)), 325, fi38}
        switch "\n\t" {
        case "\n\t":
            if true {
                p2 = 1
                var v40 int16 = -32768
                _ = v40
            }
        case "foo bar":
            p2 = (^(p2))
        default:
            p1 -= (p1 * (p1 % ((42) | 1)))
        }
        switch ("" + "\n\t") {
        case "a":
            fmt.Println(p2, p2)
        default:
            var v41 int16 = 1
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    var v42 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return 255
}

func main() {
    v1 := []int{4294967295}
    var v2 uint8 = 123
    var v3 float32 = -2.25
    v43 := GenS4{F0: -319, F1: 42}
    if true {
        var v44 float64 = 3.14159
        v3 -= 1.5
        _ = v44
    } else {
        v45 := GenS7{F0: 2, F1: 42, F2: 0}
        var v46 uint = 2
        _ = v45
        _ = v46
    }
    var v47 uint8 = 127
    var v48 int8 = -112
    var v49 int8 = v48
    v50 := GenS3{F0: v47}
    v49 -= v49
    v2 -= (v47 / (((v2 + 128)) | 1))
    _ = len(v1)
    var v51 uint16 = 167
    for fi52 := 0; fi52 < 20; fi52++ {
        v53 := GenS3{F0: (v47 | 19)}
        v54 := GenS7{F0: (v51 | 463), F1: 1, F2: uint32(v51)}
        switch -440 {
        case 9223372036854775807:
            switch fi52 {
            case -699:
                if (1000 >= 1) {
                    var v55 float32 = -2.25
                    var v56 int8 = (v49 - (v48 / ((100) | 1)))
                    _ = v55
                    _ = v56
                } else {
                    var v57 uint64 = 4294967295
                    _ = len(v1)
                    var v58 uint = 35
                    _ = v57
                    _ = v58
                }
            case 790:
                switch fi52 {
                case 1:
                    if true {
                        if true {
                            if (false && (!(2147483647 >= 2))) {
                                v47 -= v2
                                if false {
                                    v59 := GenS8{F0: v51, F1: (0 != uint(v2)), F2: -2147483648}
                                    switch (fi52 % (((fi52 & fi52)) | 1)) {
                                    case 974:
                                        if len(v1) > 0 {
                                            v1[0] = fi52
                                        }
                                    case 32767:
                                        var v60 uint32 = 87
                                        _ = v60
                                    case 1:
                                        switch "foo bar" {
                                        case "foo bar":
                                            for fi61 := 0; fi61 < 7; fi61++ {
                                                var v62 uint32 = 65535
                                                v47 += v2
                                                v3 = (1.5 * (v3 + -2.25))
                                                switch fi61 {
                                                case 32768:
                                                    var v63 uint64 = 2
                                                    _ = v63
                                                case 2147483647:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case -1:
                                                    v64 := GenS9{F0: "\n\t", F1: -1}
                                                    _ = v64
                                                default:
                                                    var v65 int = -9223372036854775808
                                                    if len(v1) > 0 {
                                                        v1[0] = (v65 - (fi52 / ((-190) | 1)))
                                                    }
                                                    v65 = fi52
                                                    _ = v65
                                                }
                                                fmt.Println(v48, v48, v62)
                                                var v66 bool = false
                                                v67 := GenS3{F0: 255}
                                                v62 -= (^(v62))
                                                _ = fi61
                                                _ = v62
                                                _ = v66
                                                _ = v67
                                            }
                                        case "a":
                                            for fi68 := 0; fi68 < 3; fi68++ {
                                                var v69 int8 = v48
                                                var v70 int = 2147483647
                                                var v71 uint64 = 379
                                                if len(v1) > 0 {
                                                    v1[0] = (v70 - (^(v70)))
                                                }
                                                _ = fi68
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            }
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = (^(fi52))
                                            }
                                        }
                                    }
                                    v48 = (v48 % ((v49) | 1))
                                    _ = v59
                                }
                            }
                            v72 := GenS3{F0: v47}
                            _ = v72
                        }
                        if (!((42 == -1) && (42 > -2147483648))) {
                            if false {
                                for fi73 := 0; fi73 < 2; fi73++ {
                                    v74 := GenS10{F0: -32768}
                                    var v75 uint8 = genFn0(fi52, 42, 65535)
                                    var v76 float32 = v3
                                    _ = fi73
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                }
                            }
                        } else {
                            v2 = (^(v2))
                            var v77 uint = 255
                            _ = v77
                        }
                        v1 = append(v1, fi52)
                    }
                case 2147483647:
                    if (((!true) && (!false)) || ((v49 | (v48 + (v49 % ((1) | 1)))) > v48)) {
                        for fi78 := 0; fi78 < 20; fi78++ {
                            fmt.Println(v51)
                            var v79 int16 = -580
                            _ = fi78
                            _ = v79
                        }
                    }
                case -1:
                    fmt.Println(v49, v51, v51)
                default:
                    var v80 string = ((("foo bar" + "a") + "hello") + "")
                    _ = v80
                }
            default:
                v81 := GenS10{F0: 1000}
                _ = v81
            }
        default:
            if false {
                v49 = v49
                v82 := GenS6{F0: 94}
                var v83 uint64 = 670
                _ = v82
                _ = v83
            } else {
                switch "a" {
                case "":
                    v3 += v3
                case "\n\t":
                    if false {
                        var v84 uint8 = genFn0(32768, 255, 1)
                        switch ("" + "\n\t") {
                        case "hello":
                            fmt.Println(v48, v3)
                        default:
                            var v85 uint64 = 828
                            _ = v85
                        }
                        if (1 != 1) {
                            if true {
                                var v86 float32 = v3
                                for fi87 := 0; fi87 < 10; fi87++ {
                                    var v88 uint64 = 0
                                    if false {
                                        if true {
                                            var v89 int32 = 233
                                            switch (^(fi52)) {
                                            case 626:
                                                if (!((v48 == 100) || (false || (0.1 != float64((((v3 - v86) / ((v86 + v86) * float32(0.0))) * float32((0.1 * (1.5 - 1e10))))))))) {
                                                    v88 -= 2
                                                    var v90 float32 = float32(float64(v86))
                                                    switch "hello" {
                                                    case "hello":
                                                        _ = len(v1)
                                                        fmt.Println(v86)
                                                        v47 = (v84 | (v2 & genFn0((fi87 - (fi87 >> ((42) & 63))), v51, 0)))
                                                    case "\n\t":
                                                        var v91 float32 = v3
                                                        _ = v91
                                                    }
                                                    _ = v90
                                                } else {
                                                    v92 := GenS5{F0: 255, F1: v51, F2: (("hello" + "\n\t") + "")}
                                                    var v93 bool = true
                                                    _ = v92
                                                    _ = v93
                                                }
                                            default:
                                                if (false || (!(true || (!(1000 > int16(v89)))))) {
                                                    if ((v88 % (((v88 / ((18446744073709551615) | 1))) | 1)) > (v88 - v88)) {
                                                        v94 := GenS9{F0: "hello", F1: 32767}
                                                        var v95 uint16 = 0
                                                        if true {
                                                            switch (fi87 % (((^(fi52))) | 1)) {
                                                            case 4294967295:
                                                                var v96 float64 = (1.5 - 1e10)
                                                                _ = v96
                                                            case 70:
                                                                v86 -= v86
                                                            default:
                                                                var v97 float64 = (1.5 - 0.1)
                                                                _ = v97
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v47 -= (v84 % ((genFn0(fi52, (v51 * v95), 923)) | 1))
                                                        }
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        v84 += (v47 % ((v2) | 1))
                                                        if (v47 < 1) {
                                                            _ = len(v1)
                                                            if len(v1) > 0 {
                                                                v1[0] = fi52
                                                            }
                                                            if (!true) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                fmt.Println(v51, fi87, v89)
                                                                var v98 float64 = 0.1
                                                                switch fi87 {
                                                                case -331:
                                                                    switch 0 {
                                                                    case 0:
                                                                        var v99 uint = 255
                                                                        _ = v99
                                                                    case 32768:
                                                                        v3 += v3
                                                                    case -906:
                                                                        if ((!true) || (127 == v48)) {
                                                                            if true {
                                                                                var v100 uint = 42
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = fi52
                                                                                }
                                                                                v1 = append(v1, fi52)
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                    default:
                                                                        v47 = v47
                                                                    }
                                                                case -1:
                                                                    v84 += 255
                                                                }
                                                                v1 = append(v1, fi52)
                                                                _ = v98
                                                            }
                                                        }
                                                    }
                                                    var v101 int64 = int64(v51)
                                                    var v102 uint64 = v88
                                                    _ = v101
                                                    _ = v102
                                                } else {
                                                    var v103 bool = (v51 > 0)
                                                    v2 = genFn0(int(v88), (v51 * v51), 47)
                                                    var v104 int8 = 1
                                                    _ = v103
                                                    _ = v104
                                                }
                                            }
                                            v3 -= (0.0 + v3)
                                            _ = v89
                                        } else {
                                            v3 = v86
                                            var v105 uint = 411
                                            var v106 int8 = (v49 - (-(v48)))
                                            _ = v105
                                            _ = v106
                                        }
                                        fmt.Println(v88, v2, fi87)
                                        var v107 float64 = (3.14159 / 1.5)
                                        _ = v107
                                    }
                                    v2 += (v2 % ((v84) | 1))
                                    var v108 int = 127
                                    _ = fi87
                                    _ = v108
                                    _ = v88
                                }
                                _ = v86
                            } else {
                                var v109 int8 = (v49 * (v49 << ((65535) & 63)))
                                if false {
                                    if (!(true || true)) {
                                        switch "a" {
                                        case "\n\t":
                                            _ = len(v1)
                                        case "a":
                                            fmt.Println(v47, v48, v49)
                                        case "hello":
                                            v2 = (v2 << ((424) & 63))
                                        default:
                                            if (0.1 < (-2.25 - float64(v3))) {
                                                fmt.Println(v48)
                                                for fi110 := 0; fi110 < 9; fi110++ {
                                                    var v111 string = ("a" + ("" + "hello"))
                                                    v112 := GenS11{F0: 1, F1: -289}
                                                    _ = fi110
                                                    _ = v111
                                                    _ = v112
                                                }
                                                v113 := GenS6{F0: 549}
                                                _ = v113
                                            } else {
                                                if (-821 < -1) {
                                                    var v114 int32 = -2147483648
                                                    _ = v114
                                                }
                                                v1 = append(v1, 0)
                                            }
                                        }
                                    } else {
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                for fi115 := 0; fi115 < 15; fi115++ {
                                                    _ = len(v1)
                                                    var v116 int16 = -246
                                                    _ = fi115
                                                    _ = v116
                                                }
                                            case "":
                                                var v117 int8 = 36
                                                _ = v117
                                            default:
                                                v109 -= (v109 / (((v49 & 57)) | 1))
                                            }
                                            if (42 < (-(v49))) {
                                                v118 := GenS11{F0: 32767, F1: -724}
                                                var v119 uint64 = 865
                                                _ = v118
                                                _ = v119
                                            }
                                            for fi120 := 0; fi120 < 13; fi120++ {
                                                var v121 uint64 = 18446744073709551615
                                                v51 = v51
                                                v122 := GenS12{F0: 567, F1: (fi52 + (fi120 + 42)), F2: v47}
                                                _ = fi120
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                    }
                                    for fi123 := 0; fi123 < 5; fi123++ {
                                        var v124 int8 = 1
                                        var v125 int64 = 873
                                        _ = len(v1)
                                        _ = fi123
                                        _ = v124
                                        _ = v125
                                    }
                                }
                                var v126 uint = 130
                                _ = v109
                                _ = v126
                            }
                            for fi127 := 0; fi127 < 1; fi127++ {
                                switch -188 {
                                case 847:
                                    if (66 == uint64(v2)) {
                                        fmt.Println(fi127, v49, v2)
                                        v84 += (v2 % ((v2) | 1))
                                        fmt.Println(v84, fi127, v2)
                                    } else {
                                        v128 := GenS4{F0: -412, F1: 619}
                                        _ = v128
                                    }
                                default:
                                    if true {
                                        v129 := GenS1{F0: -80, F1: -750}
                                        if ((0.1 >= -2.25) && (2147483648 > 602)) {
                                            if (true || (true || false)) {
                                                v1 = append(v1, fi52)
                                                var v130 string = "\n\t"
                                                _ = v130
                                            } else {
                                                fmt.Println(v3)
                                                switch fi127 {
                                                case 1:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if false {
                                                            if len(v1) > 0 {
                                                                v1[0] = -1
                                                            }
                                                        } else {
                                                            switch (fi52 & (fi52 * (fi52 - fi52))) {
                                                            case 2147483647:
                                                                v2 = genFn0((fi127 & 849), 441, 454)
                                                            case -11:
                                                                v84 += v84
                                                            case 42:
                                                                fmt.Println(v48)
                                                            }
                                                        }
                                                    default:
                                                        if (!true) {
                                                            var v131 int16 = 0
                                                            _ = v131
                                                        } else {
                                                            if (v47 <= (v84 * v47)) {
                                                                v132 := GenS7{F0: (v51 - 255), F1: 0, F2: 65535}
                                                                fmt.Println(fi127, v84)
                                                                fmt.Println(fi127, v3, fi127)
                                                                _ = v132
                                                            }
                                                        }
                                                    }
                                                case 42:
                                                    var v133 uint32 = 65535
                                                    _ = v133
                                                case 265:
                                                    v134 := GenS4{F0: int64(fi127), F1: 42}
                                                    _ = v134
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = fi127
                                                    }
                                                }
                                                v1 = append(v1, fi127)
                                            }
                                        }
                                        switch int(v2) {
                                        case 116:
                                            var v135 uint64 = 18446744073709551615
                                            _ = v135
                                        case 4294967295:
                                            v47 += v2
                                            v84 = v47
                                            v2 = v2
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        _ = v129
                                    }
                                }
                                fmt.Println(v47, v3, v49)
                                _ = fi127
                            }
                            switch fi52 {
                            case 2147483647:
                                switch "" {
                                case "foo bar":
                                    v3 = 0.1
                                case "\n\t":
                                    var v136 uint64 = 197
                                    _ = v136
                                case "hello":
                                    v137 := GenS13{F0: "\n\t", F1: "a", F2: 948}
                                    _ = v137
                                default:
                                    v138 := GenS2{F0: 649, F1: fi52}
                                    _ = v138
                                }
                            default:
                                v47 = v84
                            }
                        } else {
                            for fi139 := 0; fi139 < 14; fi139++ {
                                var v140 uint16 = (v51 * 32767)
                                v51 -= (v140 % (((v140 + v51)) | 1))
                                _ = fi139
                                _ = v140
                            }
                        }
                        _ = v84
                    }
                default:
                    v48 = -1
                }
                fmt.Println(v48)
                fmt.Println(v3)
            }
        }
        _ = fi52
        _ = v53
        _ = v54
    }
    if false {
        v141 := GenS14{F0: 65535}
        v3 = float32(0.1)
        fmt.Println(v48, v47)
        _ = v141
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v51
}
