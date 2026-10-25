package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int32; F1 string; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int; F2 string }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 uint32; F2 string }
type GenS12 struct { F0 int32; F1 int }
type GenS13 struct { F0 string; F1 int32 }
type GenS14 struct { F0 int8; F1 uint8 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int32; F1 int64 }
type GenS17 struct { F0 bool; F1 bool }
type GenS18 struct { F0 int8 }
func genFn0(p0 uint64) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            var v8 bool = (32767 >= 1000)
            _ = v8
        } else {
            v9 := GenS1{F0: 0, F1: -986}
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 string = "\n\t"
                _ = fi10
                _ = v11
            }
            fmt.Println(p0)
            _ = v9
        }
        if false {
            p0 += (^(p0))
        } else {
            for fi12 := 0; fi12 < 10; fi12++ {
                if true {
                    var v13 int32 = -1
                    p0 = 18446744073709551615
                    v13 = 645
                    _ = v13
                }
                v14 := GenS2{F0: 2, F1: -1}
                v15 := []int{0, -9223372036854775808, (fi7 | (fi12 << ((690) & 63)))}
                _ = fi12
                _ = v14
                _ = v15
            }
        }
        _ = fi7
    }
    v16 := []int{1, -287}
    switch 9223372036854775807 {
    case 32767:
        v17 := GenS3{F0: -1, F1: p0}
        _ = v17
    }
    var v18 int = 760
    var v19 float32 = ((-2.25 / 3.14159) / 0.1)
    _ = p0
    _ = v16
    _ = v18
    _ = v19
    return -1
}

func main() {
    var v1 int32 = -1
    var v2 uint64 = 124
    v3 := make(map[int]string)
    var v4 uint16 = 255
    var v5 float32 = 1.5
    var v6 string = ("foo bar" + "foo bar")
    fmt.Println(v5, v5)
    if (-23 == 100) {
        if ((!((v4 * v4) > v4)) || (v6 != v6)) {
            v5 += v5
            v20 := v3[0]
            _ = v20
        } else {
            switch v6 {
            case "":
                for fi21 := 0; fi21 < 1; fi21++ {
                    v1 = v1
                    _ = fi21
                }
            case "hello":
                v6 = "a"
            case "a":
                v22 := []int{}
                _ = v22
            default:
                v23 := GenS4{F0: 255, F1: v1}
                _ = v23
            }
        }
        if true {
            for fi24 := 0; fi24 < 13; fi24++ {
                switch (^(fi24)) {
                case 128:
                    var v25 uint64 = uint64(v4)
                    _ = v25
                case 825:
                    var v26 int8 = 42
                    _ = v26
                case 0:
                    v1 -= v1
                default:
                    v27 := []int{9223372036854775807}
                    _ = v27
                }
                fmt.Println(fi24)
                _ = fi24
            }
        } else {
            if false {
                switch (v6 + v6) {
                case "foo bar":
                    fmt.Println(v1, v6)
                case "hello":
                    fmt.Println(v6, v4, v6)
                }
                var v28 string = v6
                _ = v28
            } else {
                v29 := GenS5{F0: v1, F1: v6, F2: v2}
                var v30 bool = (true && true)
                v3[2147483647] = "hello"
                _ = v29
                _ = v30
            }
            fmt.Println(v5)
        }
    } else {
        v2 = (v2 / (((v2 | v2)) | 1))
        for fi31 := 0; fi31 < 13; fi31++ {
            for fi32 := 0; fi32 < 1; fi32++ {
                v33 := []int{(fi31 - (-(fi32))), fi31, (-(fi32)), fi31}
                v34 := GenS6{F0: (v4 + v4), F1: (fi31 - fi31), F2: v6}
                _ = fi32
                _ = v33
                _ = v34
            }
            if false {
                if (genFn0(4294967295) == genFn0(v2)) {
                    var v35 bool = false
                    _ = v35
                } else {
                    v36 := GenS2{F0: 2, F1: genFn0(0)}
                    var v37 uint = 1
                    for fi38 := 0; fi38 < 5; fi38++ {
                        v3[fi31] = v6
                        var v39 uint = (^(v37))
                        v6 = ((v6 + "hello") + (v6 + "a"))
                        _ = fi38
                        _ = v39
                    }
                    _ = v36
                    _ = v37
                }
                if false {
                    if ((207 > 112) || false) {
                        v3[(fi31 << ((2) & 63))] = ((v6 + "") + "foo bar")
                        switch (fi31 % ((fi31) | 1)) {
                        case 42:
                            switch ("\n\t" + v6) {
                            case "hello":
                                var v40 uint64 = (^(v2))
                                _ = v40
                            case "":
                                if (false || (!((!((!true) && true)) || false))) {
                                    if false {
                                        var v41 int8 = 100
                                        for fi42 := 0; fi42 < 9; fi42++ {
                                            delete(v3, 32768)
                                            var v43 float64 = 1.5
                                            var v44 uint = 2
                                            _ = fi42
                                            _ = v43
                                            _ = v44
                                        }
                                        _ = v41
                                    }
                                } else {
                                    for fi45 := 0; fi45 < 20; fi45++ {
                                        v3[fi45] = v6
                                        if ((v1 & (^(v1))) <= v1) {
                                            fmt.Println(fi31, v2)
                                        } else {
                                            if true {
                                                v5 -= v5
                                                v1 += -2147483648
                                            } else {
                                                var v46 int16 = -1
                                                _ = v46
                                            }
                                        }
                                        _ = fi45
                                    }
                                    v4 = (v4 * 0)
                                    v1 = 27
                                }
                            case "\n\t":
                                var v47 uint = 255
                                _ = v47
                            default:
                                v6 = v6
                            }
                        case 4294967295:
                            var v48 int64 = -1
                            _ = v48
                        case 9223372036854775807:
                            var v49 uint64 = (v2 * uint64(v4))
                            _ = v49
                        }
                        var v50 int64 = -9223372036854775808
                        _ = v50
                    } else {
                        if ((1 <= -1) || true) {
                            fmt.Println(fi31, v6)
                            switch (fi31 | fi31) {
                            case -534:
                                var v51 float32 = v5
                                _ = v51
                            case -601:
                                v52 := GenS1{F0: -128, F1: 9223372036854775807}
                                _ = v52
                            }
                            if (!true) {
                                switch v6 {
                                case "\n\t":
                                    v53 := []int{fi31}
                                    _ = v53
                                case "foo bar":
                                    v54 := GenS7{F0: (true || (v2 < v2)), F1: 2147483648}
                                    _ = v54
                                case "hello":
                                    v4 -= uint16(v2)
                                default:
                                    var v55 float64 = -2.25
                                    _ = v55
                                }
                                for fi56 := 0; fi56 < 8; fi56++ {
                                    var v57 uint16 = 76
                                    v58 := GenS8{F0: 2}
                                    switch v6 {
                                    case "hello":
                                        var v59 uint64 = v2
                                        _ = v59
                                    case "\n\t":
                                        v2 -= (v2 ^ (v2 / ((v2) | 1)))
                                    }
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                            } else {
                                switch ("a" + v6) {
                                case "a":
                                    _ = len(v3)
                                case "\n\t":
                                    var v60 string = v6
                                    _ = v60
                                default:
                                    delete(v3, fi31)
                                }
                                if (!false) {
                                    switch fi31 {
                                    case 9223372036854775807:
                                        for fi61 := 0; fi61 < 5; fi61++ {
                                            switch (fi31 - int(v1)) {
                                            case -502:
                                                switch "a" {
                                                case "":
                                                    v62 := GenS4{F0: 231, F1: 42}
                                                    _ = v62
                                                case "hello":
                                                    _ = len(v3)
                                                case "a":
                                                    v63 := GenS9{F0: v6, F1: v1}
                                                    _ = v63
                                                default:
                                                    var v64 float64 = ((float64(float32(((((1e10 - 1e10) / 3.14159) - float64(0.1)) / (1.5 + (-2.25 / -2.25))))) * 1e10) + (-2.25 + 0.1))
                                                    _ = v64
                                                }
                                            case 32767:
                                                if false {
                                                    fmt.Println(v6, fi61, v4)
                                                    var v65 string = (v6 + "foo bar")
                                                    _ = v65
                                                } else {
                                                    v2 -= (v2 >> ((0) & 63))
                                                }
                                            case 4294967295:
                                                if true {
                                                    v2 = v2
                                                    fmt.Println(v2, fi31, fi61)
                                                    v2 += v2
                                                    var v66 uint32 = 395
                                                    var v67 string = "\n\t"
                                                    if (0 > 0) {
                                                        v68 := GenS10{F0: -1}
                                                        var v69 int16 = -1
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    fmt.Println(fi61, v1)
                                                }
                                                if ((v4 / ((1) | 1)) <= v4) {
                                                    var v70 int = (^(fi31))
                                                    _ = v70
                                                }
                                            }
                                            _ = fi61
                                        }
                                    default:
                                        v71 := GenS11{F0: 4294967295, F1: 65535, F2: v6}
                                        _ = v71
                                    }
                                }
                                var v72 int16 = 342
                                _ = v72
                            }
                        }
                        v4 = (v4 - (v4 & (v4 | v4)))
                        switch fi31 {
                        case -2147483648:
                            v1 = -707
                        }
                    }
                    v73 := GenS12{F0: -1, F1: 732}
                    for fi74 := 0; fi74 < 11; fi74++ {
                        v75 := []int{}
                        _ = fi74
                        _ = v75
                    }
                    _ = v73
                }
            } else {
                v76 := []int{4294967295}
                _ = v76
            }
            _ = fi31
        }
    }
    v77 := GenS13{F0: ("" + (("a" + v6) + v6)), F1: v1}
    v78 := []int{32767, 2147483647, 4294967295, 2147483647}
    if (true || true) {
        _ = len(v78)
        if (((false && false) || true) && ((!(127 <= 104)) || false)) {
            if false {
                v79 := v3[32767]
                for fi80 := 0; fi80 < 3; fi80++ {
                    fmt.Println(v5)
                    _ = fi80
                }
                _ = v79
            } else {
                v81 := v3[1]
                _ = v81
            }
            v3[32768] = v6
            var v82 uint32 = 65535
            _ = v82
        } else {
            if (-378 != 1) {
                switch 648 {
                case -845:
                    var v83 string = "foo bar"
                    _ = v83
                }
                switch "a" {
                case "":
                    v1 = 832
                case "\n\t":
                    v1 = v1
                case "hello":
                    for fi84 := 0; fi84 < 19; fi84++ {
                        if (0.0 > (1.5 * 1.5)) {
                            var v85 int64 = 1
                            _ = v85
                        } else {
                            v2 = 18446744073709551615
                            v4 = v4
                        }
                        for fi86 := 0; fi86 < 16; fi86++ {
                            if len(v78) > 0 {
                                v78[0] = 32768
                            }
                            if ((-873 <= -1) || (!(0 >= 164))) {
                                switch (fi86 / ((fi84) | 1)) {
                                case 42:
                                    v5 -= (0.1 - 1e10)
                                }
                                v2 -= (v2 ^ 2)
                                var v87 uint64 = v2
                                if (v87 == (^(v87))) {
                                    v88 := GenS5{F0: int32(v87), F1: (v6 + v6), F2: uint64(v4)}
                                    if false {
                                        v89 := GenS2{F0: 99, F1: -128}
                                        var v90 int8 = 42
                                        v1 += v1
                                        _ = v89
                                        _ = v90
                                    }
                                    _ = v88
                                }
                                _ = len(v78)
                                var v91 int = (fi84 >> ((351) & 63))
                                _ = v87
                                _ = v91
                            } else {
                                var v92 int64 = 9223372036854775807
                                _ = len(v3)
                                fmt.Println(fi84, v2, fi84)
                                var v93 uint8 = 127
                                var v94 int8 = genFn0((v2 ^ (v2 & 2)))
                                var v95 uint64 = 155
                                _ = v92
                                _ = v93
                                _ = v94
                                _ = v95
                            }
                            _ = fi86
                        }
                        _ = fi84
                    }
                default:
                    if len(v78) > 0 {
                        v78[0] = -614
                    }
                }
                fmt.Println(v2, v6)
            }
        }
        switch v6 {
        case "":
            for fi96 := 0; fi96 < 5; fi96++ {
                if false {
                    for fi97 := 0; fi97 < 6; fi97++ {
                        v98 := GenS10{F0: 677}
                        var v99 uint16 = 741
                        if (((v5 - float32((((float64(0.0) - ((0.0 * (float64(float32((3.14159 + (((1.5 / (((0.0 - 1e10) / (-2.25 * 0.1)) - (1e10 * float64(float32((-2.25 / ((1e10 - 0.1) * (float64((1.5 + (0.0 * 1e10))) - float64((v5 / (-2.25 - 0.1))))))))))) + (-2.25 - -2.25)) - 1e10)))) - (0.1 - 3.14159))) - 1.5)) - (-2.25 / (0.1 / (-2.25 + 3.14159)))) - float64(v5)))) <= (0.0 * v5)) || (-2147483648 == v1)) {
                            v100 := GenS4{F0: 4294967295, F1: (v1 + (v1 + v1))}
                            var v101 int32 = v1
                            v2 -= (v2 | v2)
                            fmt.Println(v4, v101, v1)
                            _ = v100
                            _ = v101
                        }
                        _ = fi97
                        _ = v98
                        _ = v99
                    }
                    var v102 bool = false
                    var v103 int32 = (^(v1))
                    _ = v102
                    _ = v103
                }
                for fi104 := 0; fi104 < 1; fi104++ {
                    v2 = v2
                    v4 -= 66
                    _ = fi104
                }
                _ = fi96
            }
        case "hello":
            switch 4294967295 {
            case -828:
                if (false && false) {
                    fmt.Println(v4, v1, v5)
                    switch v6 {
                    case "a":
                        v5 = v5
                    case "\n\t":
                        if true {
                            v105 := GenS4{F0: 541, F1: 50}
                            _ = v105
                        } else {
                            v106 := GenS14{F0: -128, F1: 255}
                            if (2147483647 == 65535) {
                                _ = len(v78)
                            } else {
                                _ = len(v78)
                                v2 = (^(v2))
                                switch "a" {
                                case "hello":
                                    v6 = ""
                                case "\n\t":
                                    var v107 string = ("hello" + (v6 + v6))
                                    _ = v107
                                case "foo bar":
                                    var v108 uint32 = 2
                                    _ = v108
                                }
                            }
                            if (!true) {
                                v1 = 9
                                fmt.Println(v6, v5, v1)
                            } else {
                                var v109 float64 = 1e10
                                _ = v109
                            }
                            _ = v106
                        }
                    }
                } else {
                    var v110 uint64 = v2
                    _ = v110
                }
            case 0:
                v111 := v3[-314]
                _ = v111
            case -9223372036854775808:
                for fi112 := 0; fi112 < 12; fi112++ {
                    var v113 int8 = genFn0((v2 - (v2 % ((42) | 1))))
                    _ = fi112
                    _ = v113
                }
            default:
                v5 = v5
            }
        case "a":
            for fi114 := 0; fi114 < 14; fi114++ {
                if false {
                    for fi115 := 0; fi115 < 20; fi115++ {
                        v5 += 1.5
                        v78 = append(v78, (fi114 | fi114))
                        _ = fi115
                    }
                    v116 := GenS4{F0: 233, F1: (v1 * 0)}
                    _ = v116
                }
                for fi117 := 0; fi117 < 15; fi117++ {
                    var v118 int64 = -9223372036854775808
                    if (!(!true)) {
                        var v119 int8 = genFn0((v2 & v2))
                        switch (-(fi117)) {
                        case 899:
                            fmt.Println(v2, v4)
                        }
                        _ = v119
                    }
                    v5 += float32(0.1)
                    _ = fi117
                    _ = v118
                }
                var v120 int32 = (v1 * (v1 | (v1 ^ v1)))
                _ = fi114
                _ = v120
            }
        }
    } else {
        if len(v78) >= 2 {
            v78 = v78[1:len(v78)]
        }
        fmt.Println(v5, v5, v5)
    }
    var v121 float64 = (-2.25 * (((3.14159 * 0.1) * float64(v5)) * -2.25))
    v122 := GenS15{F0: -590}
    switch 606 {
    case 667:
        _ = len(v3)
    case 9223372036854775807:
        v1 += (v1 | v1)
    }
    fmt.Println(v6, v121, v4)
    fmt.Println(v121)
    for fi123 := 0; fi123 < 6; fi123++ {
        if true {
            var v124 uint = 4294967295
            if true {
                v125 := GenS16{F0: (^(v1)), F1: -1}
                if len(v78) > 0 {
                    v78[0] = int(v124)
                }
                switch (-(fi123)) {
                case 9223372036854775807:
                    var v126 bool = false
                    _ = v126
                case 3:
                    for fi127 := 0; fi127 < 16; fi127++ {
                        switch fi127 {
                        case -931:
                            switch (fi127 * fi123) {
                            case 128:
                                _ = len(v78)
                            case 882:
                                v128 := GenS17{F0: (!true), F1: true}
                                _ = v128
                            case 448:
                                var v129 float64 = (float64((-2.25 + ((1e10 - (-2.25 * v5)) - float32((float64(-2.25) - 3.14159))))) + v121)
                                _ = v129
                            }
                        case 1:
                            if false {
                                var v130 int64 = -174
                                _ = v130
                            } else {
                                if len(v78) > 0 {
                                    v78[0] = fi127
                                }
                                switch v6 {
                                case "a":
                                    var v131 uint64 = v2
                                    _ = v131
                                case "foo bar":
                                    v5 -= float32(float64(-2.25))
                                }
                                var v132 uint8 = 255
                                _ = v132
                            }
                        case 32767:
                            var v133 int16 = 1
                            _ = v133
                        }
                        if ((v2 % (((v2 << (((v124 >> ((v124) & 63))) & 63))) | 1)) > (v2 - (v2 + (v2 * v2)))) {
                            v121 -= (v121 * 0.1)
                        }
                        _ = fi127
                    }
                case 127:
                    if (true && (!((true || (uint8(v124) <= 200)) && (true && true)))) {
                        var v134 int16 = 571
                        _ = len(v3)
                        _ = len(v78)
                        _ = v134
                    }
                default:
                    for fi135 := 0; fi135 < 4; fi135++ {
                        if true {
                            v1 = -2147483648
                            var v136 float64 = 0.0
                            var v137 uint64 = 1
                            var v138 float32 = v5
                            _ = v136
                            _ = v137
                            _ = v138
                        } else {
                            fmt.Println(v5)
                            v139 := GenS9{F0: ((v6 + v6) + ("\n\t" + v6)), F1: (-(v1))}
                            fmt.Println(v6, fi135, v4)
                            _ = v139
                        }
                        fmt.Println(v2, v124)
                        delete(v3, (fi123 << ((v124) & 63)))
                        _ = fi135
                    }
                }
                _ = v125
            } else {
                fmt.Println(fi123)
                v1 += (v1 - v1)
                if false {
                    for fi140 := 0; fi140 < 12; fi140++ {
                        switch (fi123 & fi140) {
                        case 42:
                            v141 := GenS18{F0: genFn0(v2)}
                            _ = v141
                        }
                        var v142 uint32 = 433
                        _ = fi140
                        _ = v142
                    }
                } else {
                    var v143 int64 = 702
                    v5 += (-2.25 - (v5 * v5))
                    _ = v143
                }
            }
            _ = v124
        } else {
            if len(v78) > 0 {
                v78[0] = (fi123 >> ((456) & 63))
            }
            _ = len(v78)
            var v144 float32 = (v5 / (1.5 + v5))
            _ = v144
        }
        fmt.Println(fi123, fi123)
        for fi145 := 0; fi145 < 1; fi145++ {
            v3[42] = (v6 + v6)
            v1 = (v1 / (((v1 ^ (v1 + int32(v2)))) | 1))
            _ = fi145
        }
        _ = fi123
    }
    v5 = (v5 - v5)
    if true {
        v146 := GenS5{F0: v1, F1: "\n\t", F2: 4294967295}
        _ = len(v3)
        _ = v146
    }
    if ("foo bar" == "\n\t") {
        v5 += 1e10
    }
    for fi147 := 0; fi147 < 8; fi147++ {
        if len(v78) > 0 {
            v78[0] = 818
        }
        _ = fi147
    }
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
}
