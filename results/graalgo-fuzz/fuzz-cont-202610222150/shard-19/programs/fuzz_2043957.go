package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int64; F1 string; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 bool; F2 int }
type GenS9 struct { F0 int16; F1 string; F2 uint }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int16; F1 uint32; F2 int64 }
type GenS14 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS15 struct { F0 int32; F1 bool }
func genFn0(p0 int, p1 uint, p2 int8) int8 {
    var v7 int16 = 271
    p1 = (^(p1))
    v8 := []int{int(p1), p0, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 string = ""
    var v2 uint16 = 1
    var v3 bool = true
    var v4 string = (v1 + (v1 + v1))
    var v5 uint = 238
    var v6 uint32 = 700
    v9 := []int{32767, 113, 32767, 344}
    if v3 {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if ("a" == (v4 + v4)) {
            v10 := GenS1{F0: 0, F1: 714}
            v3 = v3
            _ = v10
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint32 = (v6 & (v6 | v6))
                var v13 int8 = genFn0((fi11 | fi11), v5, -107)
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 float32 = float32((-2.25 - ((-2.25 / float64(0.1)) - 0.1)))
                    fmt.Println(v16, v2, fi14)
                    _ = fi15
                    _ = v16
                }
                if v3 {
                    v17 := GenS2{F0: 876, F1: 0, F2: uint32(v5)}
                    _ = v17
                }
                var v18 uint = (v5 | (v5 % ((v5) | 1)))
                _ = fi14
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int16 = -524
                _ = fi20
                _ = v21
            }
            switch (fi19 % (((fi19 + -851)) | 1)) {
            case 630:
                switch (("" + "foo bar") + "\n\t") {
                case "":
                    var v22 int16 = -1
                    _ = v22
                case "a":
                    if (1 != -1) {
                        var v23 int = 854
                        v6 = (v6 & (v6 * v6))
                        _ = v23
                    }
                default:
                    switch (fi19 - 128) {
                    case 32767:
                        switch ("\n\t" + v1) {
                        case "hello":
                            if (0.0 > -2.25) {
                                v24 := GenS3{F0: 42}
                                switch fi19 {
                                case -9223372036854775808:
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        var v26 float32 = -2.25
                                        _ = fi25
                                        _ = v26
                                    }
                                case -2147483648:
                                    if (!(v6 < v6)) {
                                        if v3 {
                                            for fi27 := 0; fi27 < 16; fi27++ {
                                                var v28 bool = (0 < int64(v6))
                                                switch fi19 {
                                                case -9223372036854775808:
                                                    v29 := GenS1{F0: fi19, F1: 4294967295}
                                                    _ = v29
                                                case -735:
                                                    v2 = v2
                                                }
                                                var v30 uint8 = 0
                                                _ = fi27
                                                _ = v28
                                                _ = v30
                                            }
                                            var v31 int16 = 751
                                            _ = v31
                                        }
                                    } else {
                                        switch (v1 + v1) {
                                        case "a":
                                            if ((0 <= 2) && v3) {
                                                if ((fi19 % ((fi19) | 1)) == fi19) {
                                                    switch (fi19 & 9223372036854775807) {
                                                    case -1:
                                                        if len(v9) > 0 {
                                                            v9[0] = fi19
                                                        }
                                                    case 42:
                                                        switch (fi19 - (fi19 / ((fi19) | 1))) {
                                                        case 4294967295:
                                                            var v32 string = ""
                                                            _ = v32
                                                        }
                                                    case -529:
                                                        switch (v4 + ((v4 + v4) + v4)) {
                                                        case "foo bar":
                                                            var v33 float32 = -2.25
                                                            _ = v33
                                                        case "hello":
                                                            _ = len(v9)
                                                        case "":
                                                            var v34 bool = (v1 >= "")
                                                            _ = v34
                                                        default:
                                                            if v3 {
                                                                if v3 {
                                                                    for fi35 := 0; fi35 < 2; fi35++ {
                                                                        v1 = (v1 + (v4 + v4))
                                                                        v3 = (!v3)
                                                                        _ = fi35
                                                                    }
                                                                    v9 = append(v9, fi19)
                                                                    v1 = v1
                                                                } else {
                                                                    fmt.Println(v3, v1)
                                                                    _ = len(v9)
                                                                }
                                                                var v36 int32 = 42
                                                                if v3 {
                                                                    switch v4 {
                                                                    case "a":
                                                                        v3 = (v36 <= (v36 ^ v36))
                                                                    }
                                                                    var v37 string = ""
                                                                    _ = v37
                                                                } else {
                                                                    if (v3 && v3) {
                                                                        v6 = (v6 / ((v6) | 1))
                                                                        var v38 uint16 = (v2 - v2)
                                                                        _ = v38
                                                                    } else {
                                                                        v1 = (v1 + (v4 + v1))
                                                                        v6 = v6
                                                                        for fi39 := 0; fi39 < 12; fi39++ {
                                                                            if v3 {
                                                                                fmt.Println(v5)
                                                                            }
                                                                            v4 = v1
                                                                            _ = fi39
                                                                        }
                                                                    }
                                                                    if v3 {
                                                                        for fi40 := 0; fi40 < 15; fi40++ {
                                                                            v9 = append(v9, fi19)
                                                                            _ = fi40
                                                                        }
                                                                        v41 := GenS3{F0: int32(v2)}
                                                                        _ = len(v9)
                                                                        _ = v41
                                                                    } else {
                                                                        v2 = 65535
                                                                        v42 := GenS1{F0: (fi19 << ((v5) & 63)), F1: 1}
                                                                        _ = v42
                                                                    }
                                                                }
                                                                _ = v36
                                                            } else {
                                                                switch "" {
                                                                case "":
                                                                    v1 = v1
                                                                case "foo bar":
                                                                    if ((v2 | (v2 / ((2) | 1))) <= (v2 & (v2 | 32767))) {
                                                                        if v3 {
                                                                            if len(v9) > 0 {
                                                                                v9[0] = 32768
                                                                            }
                                                                        } else {
                                                                            var v43 float64 = 3.14159
                                                                            v44 := GenS4{F0: "", F1: v3}
                                                                            _ = v43
                                                                            _ = v44
                                                                        }
                                                                        for fi45 := 0; fi45 < 19; fi45++ {
                                                                            v2 -= v2
                                                                            _ = fi45
                                                                        }
                                                                    } else {
                                                                        for fi46 := 0; fi46 < 12; fi46++ {
                                                                            v47 := GenS1{F0: (fi19 * (fi19 - (-(fi19)))), F1: 251}
                                                                            switch 6 {
                                                                            case 0:
                                                                                fmt.Println(v6, v2, fi19)
                                                                            case -9223372036854775808:
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        var v48 int64 = 9223372036854775807
                                                                                        _ = v48
                                                                                    }
                                                                                    var v49 uint8 = 226
                                                                                    var v50 uint8 = 2
                                                                                    _ = v49
                                                                                    _ = v50
                                                                                case "\n\t":
                                                                                    if len(v9) >= 2 {
                                                                                        v9 = v9[1:len(v9)]
                                                                                    }
                                                                                }
                                                                            case 128:
                                                                                if (!((-2147483648 == 960) && (-584 >= -595))) {
                                                                                    if false {
                                                                                        if ((((v4 + "foo bar") + v4) + "") >= "\n\t") {
                                                                                            v51 := GenS4{F0: "\n\t", F1: ((v3 && (-128 < 54)) && (!true))}
                                                                                            v52 := GenS5{F0: 139, F1: v2}
                                                                                            if (2 <= 667) {
                                                                                                fmt.Println(v1, v4)
                                                                                                var v53 int32 = 0
                                                                                                v9 = append(v9, 32767)
                                                                                                v9 = append(v9, fi19)
                                                                                                _ = v53
                                                                                            } else {
                                                                                                var v54 int64 = 1
                                                                                                _ = v54
                                                                                            }
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                        } else {
                                                                                            fmt.Println(v2, v3)
                                                                                            v3 = (v3 || (!true))
                                                                                        }
                                                                                        if (((v3 || (0.0 > 1.5)) && (-554 > -32768)) || v3) {
                                                                                            switch (fi19 & -9223372036854775808) {
                                                                                            case 32767:
                                                                                                var v55 uint16 = 65535
                                                                                                _ = v55
                                                                                            case 728:
                                                                                                _ = len(v9)
                                                                                            default:
                                                                                                if true {
                                                                                                    var v56 int8 = 100
                                                                                                    _ = v56
                                                                                                } else {
                                                                                                    if v3 {
                                                                                                        v3 = (0 > -664)
                                                                                                        v57 := GenS5{F0: 98, F1: (^(v2))}
                                                                                                        _ = v57
                                                                                                    }
                                                                                                    var v58 uint8 = 128
                                                                                                    fmt.Println(v3, v6)
                                                                                                    _ = v58
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if v3 {
                                                                                                var v59 int8 = 0
                                                                                                var v60 uint64 = 2
                                                                                                _ = v59
                                                                                                _ = v60
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v9 = append(v9, -146)
                                                                                        if v3 {
                                                                                            var v61 bool = (!(102 >= 127))
                                                                                            _ = v61
                                                                                        }
                                                                                    }
                                                                                    v4 = ("hello" + (v1 + v4))
                                                                                    v6 = v6
                                                                                } else {
                                                                                    if v3 {
                                                                                        fmt.Println(v1, fi19, v4)
                                                                                    }
                                                                                    v62 := GenS1{F0: (fi19 & fi46), F1: 1}
                                                                                    v3 = v3
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            _ = fi46
                                                                            _ = v47
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v63 := GenS6{F0: 0, F1: "foo bar", F2: v3}
                                                                    _ = v63
                                                                default:
                                                                    switch 0 {
                                                                    case 42:
                                                                        var v64 int32 = 825
                                                                        _ = v64
                                                                    case 9223372036854775807:
                                                                        v65 := GenS3{F0: -545}
                                                                        _ = v65
                                                                    case 4294967295:
                                                                        v6 = 0
                                                                    default:
                                                                        for fi66 := 0; fi66 < 14; fi66++ {
                                                                            if false {
                                                                                v67 := GenS7{F0: 32767}
                                                                                _ = v67
                                                                            }
                                                                            _ = fi66
                                                                        }
                                                                    }
                                                                }
                                                                if len(v9) > 0 {
                                                                    v9[0] = fi19
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v68 := GenS8{F0: uint16(v5), F1: false, F2: fi19}
                                                    _ = v68
                                                }
                                            }
                                        case "hello":
                                            if (false && v3) {
                                                v2 = (v2 ^ (v2 + (v2 / (((v2 - v2)) | 1))))
                                            } else {
                                                var v69 uint = (^(v5))
                                                _ = v69
                                            }
                                        }
                                        for fi70 := 0; fi70 < 17; fi70++ {
                                            if v3 {
                                                fmt.Println(v1)
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                if (!(425 == 4294967295)) {
                                                    if false {
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        v9 = append(v9, fi70)
                                                    } else {
                                                        _ = len(v9)
                                                        v3 = (1 > 128)
                                                        v71 := GenS9{F0: 42, F1: "foo bar", F2: (v5 + 0)}
                                                        _ = v71
                                                    }
                                                }
                                            } else {
                                                v6 = 2
                                            }
                                            var v72 uint8 = 0
                                            fmt.Println(v2, v72)
                                            _ = fi70
                                            _ = v72
                                        }
                                    }
                                default:
                                    if v3 {
                                        var v73 float32 = 1.5
                                        _ = v73
                                    } else {
                                        var v74 bool = true
                                        v1 = v1
                                        var v75 uint32 = (v6 * v6)
                                        _ = v74
                                        _ = v75
                                    }
                                }
                                if (int8(v5) != int8(v5)) {
                                    for fi76 := 0; fi76 < 18; fi76++ {
                                        var v77 uint = v5
                                        v78 := GenS10{F0: (^(v5)), F1: int8(v2)}
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                    }
                                    switch fi19 {
                                    case 32768:
                                        if (false || (32767 < 42)) {
                                            var v79 float64 = 0.0
                                            _ = v79
                                        } else {
                                            fmt.Println(v5, v4, v3)
                                        }
                                    }
                                }
                                _ = v24
                            }
                        case "\n\t":
                            if true {
                                var v80 uint8 = 255
                                if (0.0 != float32(1e10)) {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    switch (v1 + "foo bar") {
                                    case "":
                                        if (1 < fi19) {
                                            if len(v9) > 0 {
                                                v9[0] = -9223372036854775808
                                            }
                                        } else {
                                            v1 = v4
                                            v81 := GenS11{F0: 42}
                                            _ = v81
                                        }
                                    case "foo bar":
                                        var v82 uint8 = v80
                                        _ = v82
                                    default:
                                        var v83 uint64 = 164
                                        _ = v83
                                    }
                                }
                                _ = v80
                            }
                        default:
                            if v3 {
                                if (42 == uint8(v6)) {
                                    var v84 uint16 = (v2 ^ 65535)
                                    for fi85 := 0; fi85 < 15; fi85++ {
                                        v84 += (v84 + v84)
                                        fmt.Println(fi19, v2, v3)
                                        var v86 float32 = 0.0
                                        _ = fi85
                                        _ = v86
                                    }
                                    _ = v84
                                }
                            }
                        }
                    }
                }
            case 42:
                if (v3 && ((1e10 + 0.0) == (3.14159 / 3.14159))) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    switch v4 {
                    case "\n\t":
                        for fi87 := 0; fi87 < 17; fi87++ {
                            _ = len(v9)
                            if (42 < 0) {
                                if v3 {
                                    var v88 float64 = 1.5
                                    _ = v88
                                }
                                v89 := GenS5{F0: 2, F1: v2}
                                var v90 string = "hello"
                                _ = v89
                                _ = v90
                            }
                            v91 := GenS12{F0: 42}
                            _ = fi87
                            _ = v91
                        }
                    case "hello":
                        var v92 float64 = -2.25
                        _ = v92
                    case "":
                        fmt.Println(v1, v4, v5)
                    }
                    for fi93 := 0; fi93 < 12; fi93++ {
                        var v94 string = "hello"
                        var v95 uint16 = (^(v2))
                        _ = fi93
                        _ = v94
                        _ = v95
                    }
                }
            default:
                for fi96 := 0; fi96 < 17; fi96++ {
                    v97 := GenS13{F0: 1, F1: (^(v6)), F2: 1}
                    _ = fi96
                    _ = v97
                }
            }
            _ = fi19
        }
    } else {
        var v98 int32 = 0
        var v99 float32 = 0.0
        _ = v98
        _ = v99
    }
    if len(v9) > 0 {
        v9[0] = -195
    }
    v100 := GenS14{F0: (v4 != v4), F1: (^(v2)), F2: 97}
    var v101 int32 = -2147483648
    var v102 int16 = 1000
    var v103 uint64 = 966
    var v104 uint16 = (v2 + 42)
    switch "hello" {
    case "foo bar":
        for fi105 := 0; fi105 < 19; fi105++ {
            v106 := GenS10{F0: (v5 * v5), F1: genFn0((-(fi105)), v5, 42)}
            _ = fi105
            _ = v106
        }
    case "":
        if (true && (!(42 > 17))) {
            for fi107 := 0; fi107 < 8; fi107++ {
                v103 -= 342
                switch "foo bar" {
                case "foo bar":
                    var v108 int32 = v101
                    _ = v108
                case "hello":
                    fmt.Println(v1)
                }
                v109 := GenS3{F0: (v101 * (v101 & 0))}
                _ = fi107
                _ = v109
            }
            var v110 int16 = v102
            v111 := GenS4{F0: v4, F1: v3}
            _ = v110
            _ = v111
        } else {
            v104 = (v2 + uint16(v6))
            fmt.Println(v103)
            var v112 int64 = 822
            _ = v112
        }
    case "\n\t":
        for fi113 := 0; fi113 < 3; fi113++ {
            v114 := GenS10{F0: 1, F1: 127}
            if v3 {
                for fi115 := 0; fi115 < 19; fi115++ {
                    fmt.Println(v6)
                    var v116 uint32 = 225
                    v117 := GenS7{F0: v104}
                    _ = fi115
                    _ = v116
                    _ = v117
                }
                fmt.Println(v2, v102)
            }
            _ = fi113
            _ = v114
        }
    default:
        if v3 {
            var v118 float32 = 0.1
            for fi119 := 0; fi119 < 8; fi119++ {
                var v120 uint64 = v103
                if len(v9) > 0 {
                    v9[0] = fi119
                }
                if true {
                    v121 := GenS15{F0: (v101 | (v101 % ((0) | 1))), F1: ((^(v6)) < (v6 % (((^(v6))) | 1)))}
                    _ = v121
                }
                _ = fi119
                _ = v120
            }
            _ = v118
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
