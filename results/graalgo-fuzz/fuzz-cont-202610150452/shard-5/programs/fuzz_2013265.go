package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 uint32; F2 bool }
type GenS9 struct { F0 int8; F1 uint8; F2 string }
type GenS10 struct { F0 string; F1 uint16; F2 uint64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8; F1 int16; F2 uint64 }
type GenS13 struct { F0 uint; F1 uint64; F2 uint16 }
func genFn0(p0 int16, p1 uint32, p2 int64) int32 {
    if true {
        if (127 == 29) {
            fmt.Println(p1, p1)
            for fi6 := 0; fi6 < 14; fi6++ {
                switch (^(fi6)) {
                case 965:
                    var v7 int64 = p2
                    _ = v7
                }
                _ = fi6
            }
        } else {
            v8 := GenS1{F0: 32768, F1: -103, F2: 65535}
            v9 := []int{0, 42}
            _ = v8
            _ = v9
        }
    }
    p0 = (p0 ^ -1)
    p2 = p2
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return -775
}
func genFn1(p0 int32) uint {
    if (!(0 != -218)) {
        p0 -= (^(p0))
        switch -1 {
        case 9223372036854775807:
            switch "a" {
            case "\n\t":
                fmt.Println(p0)
            case "hello":
                switch -9223372036854775808 {
                case 32768:
                    var v10 uint8 = 1
                    _ = v10
                case 4294967295:
                    for fi11 := 0; fi11 < 7; fi11++ {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            p0 -= p0
                            var v13 int8 = -95
                            _ = fi12
                            _ = v13
                        }
                        _ = fi11
                    }
                default:
                    p0 = (p0 ^ (-(p0)))
                }
            }
        default:
            if true {
                p0 = p0
                fmt.Println(p0)
                v14 := GenS2{F0: (("" + ("a" + "hello")) + ("foo bar" + "\n\t")), F1: 0}
                _ = v14
            }
        }
        v15 := GenS1{F0: 32768, F1: 4294967295, F2: 255}
        _ = v15
    }
    var v16 uint32 = 65535
    v17 := GenS2{F0: "hello", F1: 42}
    _ = p0
    _ = v16
    _ = v17
    return 65535
}

func main() {
    var v1 int16 = -32768
    var v2 bool = false
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    var v5 string = "\n\t"
    for fi18 := 0; fi18 < 5; fi18++ {
        switch "foo bar" {
        case "hello":
            if v2 {
                v19 := GenS2{F0: v5, F1: (v4 >> ((65535) & 63))}
                _ = v19
            } else {
                v4 = 42
                v2 = true
            }
        case "\n\t":
            v1 += 1
        case "foo bar":
            v20 := GenS3{F0: 30}
            _ = v20
        default:
            if (false && (int32(fi18) < 0)) {
                fmt.Println(v5, v5, v5)
                v2 = (2 == 4294967295)
                if (-2147483648 > 604) {
                    switch "hello" {
                    case "":
                        fmt.Println(fi18)
                    case "hello":
                        v4 = 1
                    }
                } else {
                    fmt.Println(fi18, v4)
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 bool = false
                        v23 := GenS4{F0: 1, F1: -1, F2: 2147483647}
                        var v24 int64 = 9223372036854775807
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v5 = "foo bar"
                        if false {
                            fmt.Println(fi25, fi25, v4)
                        }
                        v3[42] = v2
                        var v26 uint32 = 938
                        if true {
                            var v27 uint8 = 2
                            if v2 {
                                v28 := []int{(-(fi25)), -601, -2147483648, fi25}
                                switch fi18 {
                                case -175:
                                    fmt.Println(v4, v26)
                                case 0:
                                    if v2 {
                                        fmt.Println(v1, v1)
                                        switch (fi25 & fi25) {
                                        case 42:
                                            v27 += v27
                                        }
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    } else {
                                        fmt.Println(v1, v27)
                                        v26 -= (v26 ^ v26)
                                    }
                                case 202:
                                    v5 = "a"
                                }
                                v29 := v3[(v26 | 50)]
                                _ = v28
                                _ = v29
                            } else {
                                v4 += v4
                                _ = len(v3)
                            }
                            var v30 string = "\n\t"
                            _ = v27
                            _ = v30
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            }
        }
        var v31 uint64 = 42
        _ = fi18
        _ = v31
    }
    var v32 string = v5
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 string = (v32 + v5)
        _ = fi33
        _ = v34
    }
    switch ("foo bar" + (v32 + v5)) {
    case "":
        v35 := v3[42]
        _ = v35
    default:
        var v36 int64 = (v4 % (((v4 << ((2) & 63))) | 1))
        _ = v36
    }
    if ((-1 == v1) || v2) {
        var v37 uint32 = 2
        if false {
            v37 += v37
            for fi38 := 0; fi38 < 12; fi38++ {
                if (fi38 < (fi38 << ((65535) & 63))) {
                    v1 += v1
                    if (1.5 <= -2.25) {
                        v39 := GenS4{F0: (^(v37)), F1: 246, F2: fi38}
                        _ = v39
                    }
                    if true {
                        v40 := GenS3{F0: 4}
                        _ = v40
                    } else {
                        var v41 int8 = 42
                        if (-128 <= (-(v41))) {
                            for fi42 := 0; fi42 < 20; fi42++ {
                                fmt.Println(v41, v5)
                                switch ((v32 + ("" + (v5 + v32))) + "hello") {
                                case "foo bar":
                                    if (false || v2) {
                                        v5 = (v32 + v32)
                                        var v43 uint16 = 922
                                        switch ("\n\t" + "hello") {
                                        case "foo bar":
                                            v44 := []int{fi38, (fi38 | (fi42 | fi38)), fi38, 0}
                                            _ = v44
                                        case "":
                                            switch 462 {
                                            case -1:
                                                switch (v5 + v5) {
                                                case "a":
                                                    var v45 uint32 = v37
                                                    _ = v45
                                                case "foo bar":
                                                    if v2 {
                                                        var v46 string = (v32 + (v5 + ((v32 + "foo bar") + "a")))
                                                        v37 -= (^(v37))
                                                        _ = v46
                                                    } else {
                                                        var v47 string = v32
                                                        v48 := GenS5{F0: (v2 && true)}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                case "hello":
                                                    if ((((1.5 + 0.1) >= 3.14159) && true) && v2) {
                                                        var v49 uint32 = (v37 ^ (v37 >> ((966) & 63)))
                                                        _ = v49
                                                    } else {
                                                        v37 -= (^(v37))
                                                        fmt.Println(v1)
                                                        var v50 float32 = -2.25
                                                        _ = v50
                                                    }
                                                }
                                            case 486:
                                                if v2 {
                                                    fmt.Println(v43, v43)
                                                    v37 += (v37 << ((464) & 63))
                                                    fmt.Println(fi38, fi38)
                                                    fmt.Println(v1, fi38)
                                                    v1 += v1
                                                } else {
                                                    fmt.Println(v1, v1, v4)
                                                }
                                            }
                                        }
                                        _ = v43
                                    } else {
                                        v51 := GenS6{F0: 255}
                                        v41 -= (v41 << ((uint(v37)) & 63))
                                        v4 -= (v4 ^ (v4 >> ((669) & 63)))
                                        if ((!(-2.25 == 1e10)) || false) {
                                            var v52 uint8 = 127
                                            _ = v52
                                        } else {
                                            v37 = v37
                                        }
                                        _ = v51
                                    }
                                case "hello":
                                    fmt.Println(v32)
                                case "\n\t":
                                    v4 -= v4
                                default:
                                    var v53 float64 = 0.1
                                    _ = v53
                                }
                                v5 = v32
                                fmt.Println(v32)
                                v5 = (v32 + v32)
                                v54 := []int{fi38, fi42, (fi42 & fi42)}
                                _ = fi42
                                _ = v54
                            }
                            fmt.Println(v2)
                        } else {
                            var v55 uint16 = 556
                            var v56 int16 = v1
                            switch "a" {
                            case "foo bar":
                                _ = len(v3)
                            case "":
                                v57 := v3[65535]
                                _ = v57
                            case "a":
                                v58 := GenS7{F0: 205}
                                _ = v58
                            default:
                                if ((v37 * v37) < (v37 ^ v37)) {
                                    var v59 string = v32
                                    _ = v59
                                } else {
                                    var v60 float64 = 1e10
                                    for fi61 := 0; fi61 < 19; fi61++ {
                                        var v62 int16 = int16(v55)
                                        _ = fi61
                                        _ = v62
                                    }
                                    _ = v60
                                }
                            }
                            _ = v55
                            _ = v56
                        }
                        _ = v41
                    }
                } else {
                    fmt.Println(v4, v2)
                }
                _ = fi38
            }
        } else {
            v63 := GenS8{F0: (v4 % ((-526) | 1)), F1: 2147483647, F2: (32767 >= 825)}
            fmt.Println(v37)
            _ = v63
        }
        v4 += v4
        _ = v37
    }
    v64 := []int{1, -842}
    var v65 int16 = (v1 * v1)
    if false {
        for fi66 := 0; fi66 < 20; fi66++ {
            fmt.Println(v65, v5)
            var v67 uint8 = 2
            _ = fi66
            _ = v67
        }
        if (0 > 886) {
            v32 = v32
            var v68 uint16 = 1
            fmt.Println(v2)
            _ = v68
        }
        for fi69 := 0; fi69 < 2; fi69++ {
            v70 := GenS1{F0: (fi69 % ((0) | 1)), F1: (fi69 / ((fi69) | 1)), F2: 2}
            v71 := GenS3{F0: 54}
            _ = fi69
            _ = v70
            _ = v71
        }
    } else {
        if true {
            var v72 uint64 = 540
            if false {
                v73 := GenS6{F0: 65535}
                if false {
                    v32 = "hello"
                    v65 -= (v1 | (v1 << ((93) & 63)))
                    v65 = (v65 ^ 1000)
                } else {
                    var v74 int = 292
                    _ = v74
                }
                v75 := GenS9{F0: -126, F1: 128, F2: ("" + (("hello" + v5) + (((("" + "hello") + v5) + v5) + ((v32 + v32) + v32))))}
                _ = v73
                _ = v75
            }
            _ = v72
        } else {
            if false {
                var v76 int32 = -609
                _ = len(v64)
                _ = v76
            } else {
                v77 := GenS1{F0: 673, F1: 32767, F2: 0}
                fmt.Println(v1)
                _ = v77
            }
        }
        v5 = ("foo bar" + v32)
        for fi78 := 0; fi78 < 11; fi78++ {
            v5 = ("" + ((v32 + (v32 + "\n\t")) + (v32 + "a")))
            _ = fi78
        }
    }
    var v79 int32 = 0
    v2 = ((!v2) && true)
    v3[858] = (v65 >= v65)
    v80 := GenS10{F0: "\n\t", F1: 508, F2: 733}
    var v81 uint16 = 635
    fmt.Println(v1, v79, v32)
    if ((v5 + (("hello" + v32) + "hello")) != v32) {
        var v82 int32 = v79
        if len(v64) >= 2 {
            v64 = v64[1:len(v64)]
        }
        _ = v82
    } else {
        if (0 > -27) {
            switch ((v32 + "hello") + ((v32 + ("a" + v5)) + v5)) {
            case "foo bar":
                for fi83 := 0; fi83 < 9; fi83++ {
                    switch ("foo bar" + v5) {
                    case "\n\t":
                        v2 = false
                    case "foo bar":
                        if v2 {
                            var v84 int64 = (v4 - (v4 * 9223372036854775807))
                            fmt.Println(v5, v65)
                            if v2 {
                                switch fi83 {
                                case -9223372036854775808:
                                    for fi85 := 0; fi85 < 12; fi85++ {
                                        var v86 uint32 = 0
                                        _ = fi85
                                        _ = v86
                                    }
                                }
                                var v87 int16 = (v1 ^ 0)
                                _ = v87
                            } else {
                                v88 := GenS11{F0: -8}
                                switch v5 {
                                case "":
                                    var v89 int32 = v79
                                    _ = v89
                                case "a":
                                    v1 = v1
                                case "foo bar":
                                    if v2 {
                                        v90 := GenS1{F0: fi83, F1: (fi83 + (^(fi83))), F2: v81}
                                        v91 := GenS12{F0: 127, F1: (v1 ^ (-(v65))), F2: 927}
                                        _ = v90
                                        _ = v91
                                    }
                                }
                                _ = v88
                            }
                            _ = v84
                        }
                    default:
                        v2 = (v2 || v2)
                    }
                    fmt.Println(v81, v65)
                    for fi92 := 0; fi92 < 7; fi92++ {
                        var v93 uint8 = 251
                        if len(v64) >= 2 {
                            v64 = v64[1:len(v64)]
                        }
                        delete(v3, 380)
                        _ = fi92
                        _ = v93
                    }
                    _ = fi83
                }
            default:
                for fi94 := 0; fi94 < 2; fi94++ {
                    for fi95 := 0; fi95 < 8; fi95++ {
                        v65 = int16(fi94)
                        delete(v3, 554)
                        _ = len(v64)
                        _ = fi95
                    }
                    if v2 {
                        v3[587] = true
                    } else {
                        v96 := GenS1{F0: (fi94 | int(v81)), F1: (fi94 >> ((77) & 63)), F2: v81}
                        if (v2 || v2) {
                            switch 42 {
                            case 4294967295:
                                var v97 uint8 = 46
                                _ = v97
                            case 774:
                                v98 := GenS7{F0: 42}
                                _ = v98
                            default:
                                var v99 uint8 = 2
                                _ = v99
                            }
                        } else {
                            switch "a" {
                            case "":
                                fmt.Println(v32)
                            }
                            var v100 int8 = 42
                            var v101 uint8 = 128
                            _ = v100
                            _ = v101
                        }
                        _ = v96
                    }
                    _ = fi94
                }
            }
            v102 := GenS7{F0: 712}
            _ = v102
        }
        if v2 {
            v103 := GenS9{F0: 100, F1: 0, F2: (v5 + ("\n\t" + "hello"))}
            switch "a" {
            case "":
                for fi104 := 0; fi104 < 17; fi104++ {
                    var v105 uint = 187
                    if ((1.5 * (0.1 * 1e10)) != float64(((-2.25 + 3.14159) / (1e10 + ((float32(1.5) + 0.1) / 1e10))))) {
                        var v106 float32 = (-2.25 - float32(0.1))
                        v1 += (-(v1))
                        fmt.Println(v4, v2)
                        _ = v106
                    } else {
                        for fi107 := 0; fi107 < 17; fi107++ {
                            if len(v64) > 0 {
                                v64[0] = (fi104 & fi107)
                            }
                            _ = fi107
                        }
                        for fi108 := 0; fi108 < 5; fi108++ {
                            var v109 uint8 = 255
                            _ = fi108
                            _ = v109
                        }
                    }
                    if v2 {
                        switch v32 {
                        case "a":
                            var v110 uint16 = (v81 >> (((v105 * (v105 - (^(v105))))) & 63))
                            _ = v110
                        case "":
                            v64 = append(v64, (fi104 - fi104))
                        case "foo bar":
                            v79 += int32(v81)
                        default:
                            v111 := GenS8{F0: 1, F1: 65535, F2: (-2 != (v4 << ((2) & 63)))}
                            _ = v111
                        }
                    }
                    _ = fi104
                    _ = v105
                }
            case "hello":
                v65 -= 42
            case "foo bar":
                v112 := v3[2147483648]
                _ = v112
            default:
                if (-159 < 654) {
                    var v113 float64 = -2.25
                    _ = v113
                } else {
                    v65 += 42
                }
            }
            if (v2 || (10 > 128)) {
                switch 42 {
                case 128:
                    if true {
                        fmt.Println(v5, v65, v1)
                    }
                case -1:
                    fmt.Println(v81)
                default:
                    switch 32768 {
                    case 128:
                        if v2 {
                            v4 = v4
                            switch 127 {
                            case 32768:
                                var v114 int8 = -1
                                _ = v114
                            case 4294967295:
                                var v115 uint32 = 0
                                _ = v115
                            default:
                                if (v2 && v2) {
                                    var v116 string = (v5 + (v32 + v5))
                                    _ = v116
                                } else {
                                    for fi117 := 0; fi117 < 14; fi117++ {
                                        var v118 string = "\n\t"
                                        _ = fi117
                                        _ = v118
                                    }
                                }
                            }
                        } else {
                            for fi119 := 0; fi119 < 10; fi119++ {
                                v64 = append(v64, 4294967295)
                                v120 := GenS12{F0: 0, F1: (v1 - v1), F2: 4294967295}
                                for fi121 := 0; fi121 < 6; fi121++ {
                                    var v122 uint32 = 2
                                    v123 := GenS12{F0: 99, F1: (v1 & v1), F2: 4294967295}
                                    _ = fi121
                                    _ = v122
                                    _ = v123
                                }
                                _ = fi119
                                _ = v120
                            }
                            var v124 int = 32
                            if (!v2) {
                                v81 = (^(v81))
                            } else {
                                v124 = v124
                                v81 = (v81 & 0)
                            }
                            _ = v124
                        }
                    case 0:
                        var v125 float64 = 1e10
                        _ = v125
                    case -1:
                        if true {
                            v126 := GenS1{F0: 127, F1: -100, F2: (^(v81))}
                            _ = v126
                        } else {
                            var v127 int = 397
                            for fi128 := 0; fi128 < 13; fi128++ {
                                var v129 uint = 4294967295
                                if (1 == 2147483647) {
                                    v130 := GenS5{F0: v2}
                                    v131 := GenS13{F0: (v129 & 1), F1: 16, F2: (v81 / ((v81) | 1))}
                                    var v132 uint16 = 32767
                                    _ = v130
                                    _ = v131
                                    _ = v132
                                }
                                switch "hello" {
                                case "a":
                                    if true {
                                        for fi133 := 0; fi133 < 19; fi133++ {
                                            if len(v64) >= 2 {
                                                v64 = v64[1:len(v64)]
                                            }
                                            _ = fi133
                                        }
                                    } else {
                                        switch v127 {
                                        case -1:
                                            v3[1] = true
                                        case 680:
                                            if (float64(3.14159) > (((float64(1e10) + -2.25) / float64((float32(float64(-2.25)) / 1.0))) - 0.0)) {
                                                v81 = 255
                                                switch v32 {
                                                case "":
                                                    var v134 uint16 = 42
                                                    _ = v134
                                                case "hello":
                                                    if len(v64) > 0 {
                                                        v64[0] = (fi128 << ((65535) & 63))
                                                    }
                                                default:
                                                    var v135 uint32 = 2147483648
                                                    _ = v135
                                                }
                                                fmt.Println(v32)
                                            }
                                        case -497:
                                            if (!(!v2)) {
                                                v81 = 675
                                                if ((((v32 + ("\n\t" + v32)) + ((v32 + (v32 + "foo bar")) + v32)) + ("" + v5)) >= "hello") {
                                                    for fi136 := 0; fi136 < 8; fi136++ {
                                                        v137 := GenS10{F0: "foo bar", F1: (v81 ^ v81), F2: 2}
                                                        _ = fi136
                                                        _ = v137
                                                    }
                                                    var v138 int32 = (^(v79))
                                                    _ = v138
                                                } else {
                                                    fmt.Println(v81, v1, v5)
                                                }
                                                var v139 uint8 = 0
                                                _ = v139
                                            } else {
                                                v65 -= 32767
                                                switch v127 {
                                                case 127:
                                                    v1 += v1
                                                case -2147483648:
                                                    if (v1 == v65) {
                                                        v4 -= v4
                                                        var v140 float64 = ((1.5 / (1.5 - (1.5 * float64(0.1)))) / ((-2.25 / float64(((0.1 * 0.1) * 0.1))) * 1.5))
                                                        _ = v140
                                                    } else {
                                                        if (v2 && v2) {
                                                            v5 = v5
                                                            switch (v32 + "hello") {
                                                            case "foo bar":
                                                                v2 = ((^(v79)) > v79)
                                                            default:
                                                                fmt.Println(v129, v129, v4)
                                                            }
                                                        }
                                                    }
                                                case 42:
                                                    var v141 uint32 = 4294967295
                                                    _ = v141
                                                default:
                                                    if len(v64) > 0 {
                                                        v64[0] = v127
                                                    }
                                                }
                                            }
                                        default:
                                            for fi142 := 0; fi142 < 13; fi142++ {
                                                var v143 uint = v129
                                                var v144 float32 = ((1e10 / 1.0) - 3.14159)
                                                var v145 uint16 = (v81 >> ((v129) & 63))
                                                delete(v3, 42)
                                                v146 := v3[748]
                                                _ = fi142
                                                _ = v143
                                                _ = v144
                                                _ = v145
                                                _ = v146
                                            }
                                        }
                                    }
                                case "":
                                    v129 += v129
                                default:
                                    var v147 int64 = (v4 * 9223372036854775807)
                                    _ = v147
                                }
                                _ = fi128
                                _ = v129
                            }
                            v81 -= 42
                            _ = v127
                        }
                    }
                }
            }
            _ = v103
        }
    }
    switch -368 {
    case 32767:
        var v148 int32 = v79
        _ = v148
    }
    var v149 int = -166
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v65:", v65)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v149
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v79
    _ = v80
    _ = v81
}
