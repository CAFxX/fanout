package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 bool; F2 int32 }
type GenS7 struct { F0 int64; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 bool }
type GenS9 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 int64; F1 int16; F2 int16 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 int8 }
type GenS15 struct { F0 int8; F1 uint }
type GenS16 struct { F0 uint16; F1 int64 }
type GenS17 struct { F0 int8; F1 uint8; F2 uint }
type GenS18 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint64) int {
    v7 := []int{846, 0}
    if true {
        var v8 int = 127
        v7 = append(v7, -9223372036854775808)
        if len(v7) > 0 {
            v7[0] = (v8 << ((2) & 63))
        }
        _ = v8
    } else {
        p0 = 255
        p1 -= (p1 / ((0) | 1))
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 32768
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 42
    var v3 int32 = v2
    var v4 float64 = (((1e10 - (((((3.14159 + 1.5) - 0.1) * (1e10 * 0.1)) - 0.0) * 1.5)) / (1e10 - float64(1.5))) / 3.14159)
    v5 := make(map[int]uint32)
    var v6 int8 = 100
    v9 := v5[int(v3)]
    v10 := GenS1{F0: 46}
    var v11 uint32 = (v9 + 43)
    v12 := GenS1{F0: 87}
    var v13 uint64 = 986
    if true {
        switch "hello" {
        case "foo bar":
            switch "a" {
            case "":
                fmt.Println(v13, v11)
            case "hello":
                _ = len(v1)
            case "a":
                v14 := []int{-2147483648, int(v9)}
                _ = v14
            default:
                v15 := GenS1{F0: 255}
                _ = v15
            }
        case "a":
            v4 -= (3.14159 * float64((3.14159 / 1e10)))
        case "hello":
            if (-581 == 42) {
                switch -245 {
                case 128:
                    for fi16 := 0; fi16 < 4; fi16++ {
                        fmt.Println(v3)
                        v17 := []int{fi16, (fi16 / (((fi16 + genFn0(255, v13))) | 1))}
                        _ = fi16
                        _ = v17
                    }
                case -51:
                    v4 = -2.25
                }
                v18 := []int{-1, 1}
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                v11 = v9
                v6 = 100
                var v20 uint64 = (v13 + 42)
                _ = fi19
                _ = v20
            }
        }
        fmt.Println(v3)
    }
    var v21 uint64 = (^(v13))
    v22 := GenS1{F0: 1}
    v13 = 18446744073709551615
    for fi23 := 0; fi23 < 9; fi23++ {
        if (float64(-2.25) != (0.1 - ((-2.25 - float64(3.14159)) * (v4 + (v4 - 1e10))))) {
            v11 += (v11 ^ (v11 << ((466) & 63)))
        } else {
            if (true || (0.1 == 3.14159)) {
                if true {
                    if true {
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v25 := []int{fi24, (fi24 % ((fi23) | 1)), genFn0(757, v21), fi24}
                            _ = len(v1)
                            if len(v25) > 0 {
                                v25[0] = fi23
                            }
                            _ = fi24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 7; fi26++ {
                            delete(v5, -9223372036854775808)
                            v27 := GenS2{F0: 42, F1: 1}
                            _ = fi26
                            _ = v27
                        }
                        switch fi23 {
                        case 2147483647:
                            var v28 int16 = -1
                            _ = v28
                        case 9223372036854775807:
                            v4 = v4
                        default:
                            if false {
                                v11 = v9
                                v2 += 433
                                v29 := GenS2{F0: genFn0(65535, v21), F1: -810}
                                _ = v29
                            } else {
                                if true {
                                    for fi30 := 0; fi30 < 5; fi30++ {
                                        v11 = 2147483647
                                        v31 := []int{(fi30 % (((fi30 << ((18446744073709551615) & 63))) | 1))}
                                        var v32 float64 = (v4 * -2.25)
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                    delete(v5, (^(fi23)))
                                } else {
                                    var v33 uint16 = 65535
                                    if true {
                                        var v34 uint16 = v33
                                        for fi35 := 0; fi35 < 6; fi35++ {
                                            var v36 uint32 = v9
                                            v37 := []int{}
                                            if true {
                                                v6 += v6
                                            } else {
                                                fmt.Println(v4, v13)
                                                var v38 int = (fi23 >> ((0) & 63))
                                                _ = v38
                                            }
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                        }
                                        for fi39 := 0; fi39 < 11; fi39++ {
                                            v9 = 710
                                            switch (-(fi23)) {
                                            case 1:
                                                v5[(^(fi23))] = 677
                                            case 884:
                                                fmt.Println(v4)
                                            case 4294967295:
                                                fmt.Println(fi39)
                                            }
                                            _ = fi39
                                        }
                                        _ = v34
                                    }
                                    fmt.Println(fi23, v9, v21)
                                    _ = v33
                                }
                                v13 -= 1
                                v5[fi23] = v11
                            }
                        }
                    }
                    var v40 uint16 = 361
                    v9 = v11
                    _ = v40
                }
                switch ("\n\t" + (("hello" + "") + "foo bar")) {
                case "\n\t":
                    switch fi23 {
                    case 128:
                        var v41 uint64 = v13
                        _ = v41
                    case 0:
                        fmt.Println(v11)
                    case 32767:
                        if ("hello" >= "a") {
                            if false {
                                v4 += -2.25
                                v11 -= (v9 & 626)
                                if false {
                                    delete(v1, v11)
                                } else {
                                    for fi42 := 0; fi42 < 14; fi42++ {
                                        fmt.Println(v2, fi23)
                                        v43 := []int{fi42}
                                        _ = fi42
                                        _ = v43
                                    }
                                }
                            }
                            v44 := GenS2{F0: (fi23 & (^(fi23))), F1: -1}
                            var v45 int32 = v3
                            _ = v44
                            _ = v45
                        }
                    default:
                        v46 := GenS3{F0: 4294967295, F1: 1, F2: 128}
                        _ = v46
                    }
                default:
                    v47 := v1[(v9 / ((1) | 1))]
                    _ = v47
                }
                if true {
                    v48 := []int{(fi23 | genFn0(2, (^(v13)))), 42}
                    if false {
                        var v49 uint16 = 0
                        switch (("a" + "a") + ("" + (("a" + "a") + ""))) {
                        case "":
                            for fi50 := 0; fi50 < 6; fi50++ {
                                var v51 uint = 65535
                                if ((65535 >= uint16(v21)) || (-1 < int16(v3))) {
                                    v52 := v1[v11]
                                    _ = v52
                                } else {
                                    fmt.Println(v21, v11, fi50)
                                    v53 := GenS4{F0: "", F1: (v11 ^ 748)}
                                    if true {
                                        var v54 int32 = v2
                                        _ = v54
                                    } else {
                                        v55 := GenS5{F0: 32767}
                                        v4 = -2.25
                                        v56 := GenS6{F0: -137, F1: false, F2: (v2 << ((v51) & 63))}
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v53
                                }
                                v4 = ((1e10 - 0.0) / float64(((0.1 + 3.14159) / float32((1e10 / v4)))))
                                var v57 string = ""
                                var v58 int64 = -9223372036854775808
                                _ = fi50
                                _ = v51
                                _ = v57
                                _ = v58
                            }
                        default:
                            v59 := GenS7{F0: int64(fi23), F1: 42, F2: -1}
                            _ = v59
                        }
                        v13 = (v13 >> ((42) & 63))
                        _ = v49
                    }
                    if true {
                        if true {
                            if false {
                                if true {
                                    switch "a" {
                                    case "":
                                        var v60 uint8 = 123
                                        _ = v60
                                    case "hello":
                                        v61 := v5[fi23]
                                        _ = v61
                                    }
                                    v6 += (v6 >> ((65535) & 63))
                                }
                            }
                            switch "foo bar" {
                            case "a":
                                var v62 int32 = v3
                                _ = v62
                            }
                            for fi63 := 0; fi63 < 17; fi63++ {
                                switch genFn0(65535, v21) {
                                case -474:
                                    var v64 int32 = (v3 % (((v2 % (((v3 ^ (v3 - (v3 ^ v2)))) | 1))) | 1))
                                    _ = v64
                                }
                                v6 = (v6 * v6)
                                var v65 int16 = int16(v6)
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(fi63)
                                default:
                                    v66 := GenS8{F0: 217, F1: (v13 & v21), F2: true}
                                    _ = v66
                                }
                                fmt.Println(v6, v6, v9)
                                var v67 int64 = -9223372036854775808
                                _ = fi63
                                _ = v65
                                _ = v67
                            }
                        }
                    } else {
                        v2 -= v2
                        v3 -= (-(v3))
                        if true {
                            fmt.Println(v13, v6, v11)
                        }
                    }
                    _ = v48
                } else {
                    if ((v9 >= (v11 * v9)) && false) {
                        var v68 int16 = -1
                        switch (fi23 << ((1) & 63)) {
                        case 510:
                            if false {
                                if (489 < int64(fi23)) {
                                    fmt.Println(fi23)
                                    fmt.Println(v6)
                                    if (false && (false && ((42 <= 1) || false))) {
                                        v2 = 0
                                        for fi69 := 0; fi69 < 2; fi69++ {
                                            if (65535 >= 42) {
                                                v3 = v3
                                            }
                                            if true {
                                                fmt.Println(fi69, v9, fi69)
                                                if (((1.5 + 1.5) != ((1e10 - 0.0) + 0.1)) && false) {
                                                    var v70 uint16 = 0
                                                    _ = v70
                                                } else {
                                                    if (255 < 18446744073709551615) {
                                                        var v71 int16 = v68
                                                        switch fi23 {
                                                        case 542:
                                                            v2 = (v2 ^ 1)
                                                        case 237:
                                                            if ((float64((1e10 - float32((0.1 / v4)))) > (v4 + v4)) || (9223372036854775807 < -1)) {
                                                                delete(v5, int(v68))
                                                                v72 := GenS1{F0: 2}
                                                                _ = v72
                                                            } else {
                                                                v73 := GenS9{F0: 0, F1: 924, F2: (v13 & (v13 % (((v21 ^ v13)) | 1)))}
                                                                v9 = 2147483647
                                                                v68 = v68
                                                                _ = v73
                                                            }
                                                        }
                                                        v68 = v68
                                                        _ = v71
                                                    } else {
                                                        if true {
                                                            var v74 uint8 = 127
                                                            switch "" {
                                                            case "foo bar":
                                                                _ = len(v5)
                                                            }
                                                            _ = v74
                                                        } else {
                                                            var v75 int8 = v6
                                                            v76 := []int{fi23}
                                                            v3 -= 578
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                    }
                                                    fmt.Println(v2)
                                                }
                                                switch "hello" {
                                                case "a":
                                                    _ = len(v5)
                                                case "":
                                                    if ((fi23 + (fi23 / (((fi23 / ((genFn0(32767, v21)) | 1))) | 1))) != (-(fi69))) {
                                                        v77 := GenS10{F0: (v3 ^ v2)}
                                                        if false {
                                                            fmt.Println(v9, v13)
                                                            if ((32 >= 127) && (false && true)) {
                                                                switch ("" + "hello") {
                                                                case "":
                                                                    var v78 uint32 = (v9 * 2)
                                                                    if true {
                                                                        if true {
                                                                            var v79 int16 = -690
                                                                            var v80 uint = 255
                                                                            _ = v79
                                                                            _ = v80
                                                                        } else {
                                                                            var v81 uint64 = 1
                                                                            v2 = v3
                                                                            var v82 float32 = 1e10
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                    }
                                                                    _ = v78
                                                                case "\n\t":
                                                                    switch "a" {
                                                                    case "a":
                                                                        switch fi23 {
                                                                        case -2147483648:
                                                                            switch fi23 {
                                                                            case 32768:
                                                                                v83 := GenS11{F0: 779, F1: (v6 % ((int8(v68)) | 1))}
                                                                                _ = v83
                                                                            default:
                                                                                v13 += 774
                                                                            }
                                                                        case -9223372036854775808:
                                                                            v84 := GenS9{F0: -9223372036854775808, F1: 493, F2: (v13 - (v13 & 1))}
                                                                            _ = v84
                                                                        case -888:
                                                                            if (255 > 255) {
                                                                                v85 := GenS12{F0: int64(v9), F1: v68, F2: 0}
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v68, v68)
                                                                    case "foo bar":
                                                                        var v86 bool = (1 < 426)
                                                                        _ = v86
                                                                    default:
                                                                        switch -278 {
                                                                        case -2147483648:
                                                                            var v87 string = (("\n\t" + "foo bar") + "")
                                                                            _ = v87
                                                                        }
                                                                    }
                                                                default:
                                                                    if ("\n\t" != "foo bar") {
                                                                        v88 := GenS13{F0: (-(v3))}
                                                                        _ = v88
                                                                    }
                                                                }
                                                                v11 += (v9 - v9)
                                                                if (!((-2.25 + ((0.1 / 1.0) / float32(1.5))) == (float32((v4 + 1.5)) * -2.25))) {
                                                                    var v89 string = "hello"
                                                                    switch ((v89 + "hello") + v89) {
                                                                    case "a":
                                                                        delete(v1, 2147483647)
                                                                        v9 += (v9 & (^(v11)))
                                                                    case "\n\t":
                                                                        v13 -= (v21 & (v13 & v13))
                                                                    default:
                                                                        v90 := []int{fi23, (fi23 - -1), (fi23 & fi69)}
                                                                        _ = v90
                                                                    }
                                                                    if true {
                                                                        var v91 float32 = 1e10
                                                                        fmt.Println(v9)
                                                                        _ = v91
                                                                    }
                                                                    _ = v89
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        fmt.Println(v13, v11)
                                                                    case "\n\t":
                                                                        fmt.Println(v11, v68)
                                                                    default:
                                                                        fmt.Println(v21, fi69, fi69)
                                                                    }
                                                                    delete(v1, v9)
                                                                    if false {
                                                                        v92 := []int{(fi23 - -841)}
                                                                        if ((((!true) || (!false)) && false) && true) {
                                                                            v2 -= 390
                                                                            v68 = (v68 - v68)
                                                                        } else {
                                                                            if len(v92) > 0 {
                                                                                v92[0] = (^(fi69))
                                                                            }
                                                                            v13 -= (v21 & v21)
                                                                            fmt.Println(v3, v2)
                                                                        }
                                                                        _ = v92
                                                                    } else {
                                                                        v93 := []int{fi69}
                                                                        var v94 int64 = 9223372036854775807
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            v9 -= (v9 % (((v11 + (^(v9)))) | 1))
                                                            var v95 uint32 = v11
                                                            switch ("hello" + "\n\t") {
                                                            case "a":
                                                                v21 += (v21 + v13)
                                                            default:
                                                                var v96 uint64 = uint64(v9)
                                                                _ = v96
                                                            }
                                                            _ = v95
                                                        }
                                                        _ = v77
                                                    } else {
                                                        v9 -= (v11 - v11)
                                                    }
                                                }
                                            } else {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if true {
                                                        v97 := []int{(fi23 << ((255) & 63)), fi69}
                                                        var v98 uint32 = (v11 & (v11 >> ((680) & 63)))
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        var v99 uint16 = 153
                                                        _ = v99
                                                    }
                                                default:
                                                    var v100 uint = 794
                                                    _ = v100
                                                }
                                                v101 := v5[fi69]
                                                var v102 float64 = v4
                                                _ = v101
                                                _ = v102
                                            }
                                            if ((43 < 94) && false) {
                                                if false {
                                                    var v103 uint64 = 1
                                                    _ = v103
                                                }
                                                v104 := v1[v9]
                                                _ = v104
                                            } else {
                                                switch (((("\n\t" + "\n\t") + "\n\t") + "\n\t") + ((("foo bar" + ((("\n\t" + (((("\n\t" + "foo bar") + (("hello" + ("a" + "foo bar")) + "a")) + "a") + "foo bar")) + "") + "foo bar")) + "foo bar") + ("" + "hello"))) {
                                                case "hello":
                                                    v5[9223372036854775807] = v11
                                                case "foo bar":
                                                    if (!((!false) && true)) {
                                                        fmt.Println(v2, v11, v9)
                                                    } else {
                                                        v2 += 2147483647
                                                        switch fi23 {
                                                        case 1:
                                                            if (false && true) {
                                                                switch "a" {
                                                                case "":
                                                                    switch fi69 {
                                                                    case 673:
                                                                        v105 := GenS11{F0: (v9 + (v11 * v9)), F1: (v6 / ((v6) | 1))}
                                                                        _ = v105
                                                                    default:
                                                                        var v106 uint64 = (v13 >> ((160) & 63))
                                                                        var v107 float64 = (float64(-2.25) / -2.25)
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                case "hello":
                                                                    var v108 float64 = (v4 + v4)
                                                                    _ = v108
                                                                case "\n\t":
                                                                    var v109 uint8 = 255
                                                                    _ = v109
                                                                default:
                                                                    v2 = v2
                                                                    switch ("" + ("a" + "")) {
                                                                    case "hello":
                                                                        switch "a" {
                                                                        case "a":
                                                                            fmt.Println(v2, v6, v11)
                                                                        }
                                                                    default:
                                                                        v9 = v9
                                                                    }
                                                                }
                                                                v11 -= v9
                                                            }
                                                        case 2147483647:
                                                            if (!(true && (!false))) {
                                                                var v110 float32 = (1e10 * 1e10)
                                                                _ = v110
                                                            } else {
                                                                v111 := []int{127, (^(fi69)), (fi69 >> ((18446744073709551615) & 63))}
                                                                var v112 int64 = 0
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            var v113 string = (("\n\t" + "a") + "a")
                                                            _ = v113
                                                        default:
                                                            fmt.Println(v68, v13, v21)
                                                            v3 -= (v3 & -1)
                                                        }
                                                    }
                                                }
                                            }
                                            _ = fi69
                                        }
                                    } else {
                                        fmt.Println(v21, v4)
                                        for fi114 := 0; fi114 < 9; fi114++ {
                                            if true {
                                                v21 = 540
                                                var v115 int32 = (v3 * 42)
                                                _ = v115
                                            }
                                            v116 := GenS14{F0: 1, F1: (v6 / (((v6 - -1)) | 1))}
                                            if (!true) {
                                                if true {
                                                    v117 := GenS15{F0: int8(v3), F1: 289}
                                                    _ = len(v1)
                                                    _ = v117
                                                }
                                                v5[(fi114 >> ((uint(v11)) & 63))] = 0
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    fmt.Println(v6)
                                                case "hello":
                                                    var v118 float32 = 0.1
                                                    _ = v118
                                                default:
                                                    fmt.Println(v13)
                                                }
                                            } else {
                                                v3 = (v3 >> ((65535) & 63))
                                                if false {
                                                    v119 := GenS14{F0: 255, F1: v6}
                                                    delete(v5, (-(fi23)))
                                                    fmt.Println(v13, v2, v3)
                                                    if true {
                                                        var v120 int16 = (v68 >> ((661) & 63))
                                                        _ = v120
                                                    }
                                                    v4 = 0.1
                                                    v121 := GenS2{F0: (fi23 >> ((0) & 63)), F1: 420}
                                                    _ = v119
                                                    _ = v121
                                                } else {
                                                    var v122 uint = 60
                                                    v21 += v13
                                                    var v123 int = genFn0(2, (v21 & (v21 / ((0) | 1))))
                                                    _ = v122
                                                    _ = v123
                                                }
                                            }
                                            _ = fi114
                                            _ = v116
                                        }
                                    }
                                } else {
                                    v11 = (v9 & v9)
                                    if false {
                                        switch "\n\t" {
                                        case "hello":
                                            if false {
                                                _ = len(v1)
                                                for fi124 := 0; fi124 < 10; fi124++ {
                                                    var v125 uint8 = 0
                                                    v126 := GenS14{F0: 1, F1: v6}
                                                    var v127 uint32 = v9
                                                    _ = fi124
                                                    _ = v125
                                                    _ = v126
                                                    _ = v127
                                                }
                                            } else {
                                                for fi128 := 0; fi128 < 2; fi128++ {
                                                    v21 = 2
                                                    _ = fi128
                                                }
                                                for fi129 := 0; fi129 < 7; fi129++ {
                                                    var v130 uint16 = 65535
                                                    if (255 >= 255) {
                                                        v21 = v13
                                                        v3 -= -2147483648
                                                    } else {
                                                        var v131 int8 = v6
                                                        fmt.Println(v4, v11, fi23)
                                                        _ = v131
                                                    }
                                                    var v132 string = "a"
                                                    _ = fi129
                                                    _ = v130
                                                    _ = v132
                                                }
                                            }
                                        case "a":
                                            var v133 bool = false
                                            _ = v133
                                        default:
                                            for fi134 := 0; fi134 < 13; fi134++ {
                                                if (255 < 65535) {
                                                    v11 = 65535
                                                }
                                                var v135 int64 = 380
                                                _ = fi134
                                                _ = v135
                                            }
                                        }
                                        if true {
                                            fmt.Println(v11, fi23, v11)
                                            if (("hello" > "hello") && true) {
                                                var v136 int8 = 100
                                                v5[(fi23 % (((fi23 >> ((2) & 63))) | 1))] = v11
                                                _ = v136
                                            } else {
                                                var v137 float32 = -2.25
                                                v6 -= v6
                                                _ = v137
                                            }
                                            if ((v13 & 0) <= 1) {
                                                for fi138 := 0; fi138 < 19; fi138++ {
                                                    var v139 float32 = float32((-2.25 + (v4 + float64(((3.14159 / (0.0 + 1e10)) - -2.25)))))
                                                    v140 := []int{-2147483648, fi23, fi23}
                                                    v11 -= 2
                                                    _ = fi138
                                                    _ = v139
                                                    _ = v140
                                                }
                                                v141 := GenS6{F0: int64(v2), F1: (!false), F2: v2}
                                                _ = v141
                                            } else {
                                                fmt.Println(v11)
                                                var v142 int16 = v68
                                                _ = v142
                                            }
                                        } else {
                                            v4 += v4
                                        }
                                        switch fi23 {
                                        case 122:
                                            v4 = v4
                                        }
                                    } else {
                                        v1[v9] = ((false || true) && (false && false))
                                        fmt.Println(v21, v6)
                                    }
                                    if true {
                                        switch (("foo bar" + "") + (("hello" + ("" + "hello")) + "foo bar")) {
                                        case "foo bar":
                                            fmt.Println(v11, v11)
                                        case "hello":
                                            for fi143 := 0; fi143 < 3; fi143++ {
                                                var v144 int8 = 42
                                                _ = fi143
                                                _ = v144
                                            }
                                        }
                                    } else {
                                        v145 := GenS16{F0: 0, F1: 42}
                                        _ = v145
                                    }
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v146 := GenS17{F0: (^(v6)), F1: 196, F2: 65535}
                                    _ = v146
                                case "":
                                    var v147 int8 = int8(v11)
                                    _ = v147
                                default:
                                    for fi148 := 0; fi148 < 1; fi148++ {
                                        v13 = 0
                                        if false {
                                            var v149 string = ("hello" + "a")
                                            var v150 float64 = 0.0
                                            v151 := GenS12{F0: 0, F1: (v68 / (((v68 ^ (v68 << ((1) & 63)))) | 1)), F2: (v68 % ((v68) | 1))}
                                            _ = v149
                                            _ = v150
                                            _ = v151
                                        }
                                        var v152 uint32 = (^(v11))
                                        _ = fi148
                                        _ = v152
                                    }
                                }
                                v13 = (^(v21))
                            } else {
                                if false {
                                    switch (fi23 | fi23) {
                                    case 127:
                                        v21 = (v13 + (v21 << ((65535) & 63)))
                                    case 42:
                                        if true {
                                            v153 := v5[(fi23 % ((fi23) | 1))]
                                            _ = v153
                                        } else {
                                            var v154 float64 = float64((1.5 + float32(v4)))
                                            var v155 uint16 = uint16(v9)
                                            _ = v154
                                            _ = v155
                                        }
                                    case 764:
                                        var v156 int = fi23
                                        _ = v156
                                    }
                                } else {
                                    for fi157 := 0; fi157 < 11; fi157++ {
                                        fmt.Println(v21, v21, fi157)
                                        v158 := []int{32767, genFn0(2, v21), fi157}
                                        _ = fi157
                                        _ = v158
                                    }
                                    v159 := GenS18{F0: (v68 % ((v68) | 1))}
                                    _ = v159
                                }
                                switch (("" + "\n\t") + "a") {
                                case "":
                                    var v160 float64 = v4
                                    _ = v160
                                }
                                v6 = v6
                            }
                        }
                        v161 := GenS8{F0: 127, F1: v13, F2: false}
                        _ = v161
                        _ = v68
                    }
                }
            }
        }
        for fi162 := 0; fi162 < 2; fi162++ {
            fmt.Println(fi162, v13)
            v163 := []int{genFn0(255, (v13 | uint64(v11))), genFn0(42, v13)}
            _ = fi162
            _ = v163
        }
        _ = fi23
    }
    fmt.Println(v6)
    v13 = (v21 >> ((4294967295) & 63))
    var v164 uint32 = uint32(v13)
    switch genFn0(255, 18446744073709551615) {
    case -557:
        for fi165 := 0; fi165 < 15; fi165++ {
            var v166 int32 = (v3 >> ((2) & 63))
            if false {
                switch (fi165 / (((fi165 + (^(fi165)))) | 1)) {
                case 0:
                    v13 += v13
                case 500:
                    for fi167 := 0; fi167 < 5; fi167++ {
                        if true {
                            var v168 int32 = 42
                            if true {
                                fmt.Println(v2, v4)
                            } else {
                                switch fi165 {
                                case 1:
                                    if true {
                                        v169 := GenS2{F0: (fi165 ^ genFn0(656, (v21 + (v13 ^ v13)))), F1: 9223372036854775807}
                                        _ = v169
                                    }
                                default:
                                    if true {
                                        v170 := []int{fi165, fi167, 362}
                                        var v171 uint8 = 127
                                        switch (fi165 | (fi167 * (fi165 & (^(fi165))))) {
                                        case -9223372036854775808:
                                            if false {
                                                var v172 bool = (637 < v164)
                                                if ((v172 && v172) || false) {
                                                    v166 = v2
                                                }
                                                v171 += v171
                                                switch (((("" + "hello") + ("hello" + ("\n\t" + ("hello" + (("hello" + (("a" + "") + "foo bar")) + ("a" + (("" + (((("\n\t" + "foo bar") + "") + "\n\t") + ("a" + ""))) + ("\n\t" + "foo bar")))))))) + "foo bar") + "a") {
                                                case "\n\t":
                                                    if false {
                                                        if ((32767 < 227) || false) {
                                                            if (!v172) {
                                                                v173 := GenS1{F0: uint16(v13)}
                                                                var v174 uint32 = (v11 + uint32(v13))
                                                                _ = v173
                                                                _ = v174
                                                            }
                                                            v175 := GenS5{F0: -1}
                                                            var v176 int64 = 9223372036854775807
                                                            _ = v175
                                                            _ = v176
                                                        } else {
                                                            fmt.Println(v166, v4)
                                                        }
                                                        v177 := GenS2{F0: fi165, F1: 0}
                                                        _ = v177
                                                    }
                                                default:
                                                    if (v3 > v3) {
                                                        delete(v5, genFn0(995, uint64(v11)))
                                                        v178 := GenS7{F0: -1, F1: 1, F2: 713}
                                                        _ = v178
                                                    }
                                                }
                                                v6 += (v6 >> ((0) & 63))
                                                _ = v172
                                            }
                                        case -1:
                                            switch "" {
                                            case "\n\t":
                                                switch (fi167 >> ((392) & 63)) {
                                                case 0:
                                                    var v179 uint16 = 537
                                                    _ = v179
                                                case -615:
                                                    var v180 string = "a"
                                                    _ = v180
                                                }
                                            default:
                                                switch 4294967295 {
                                                case 127:
                                                    v11 += (v11 >> ((879) & 63))
                                                }
                                            }
                                        case -915:
                                            var v181 uint32 = (v9 / (((v11 + 42)) | 1))
                                            _ = v181
                                        }
                                        _ = v170
                                        _ = v171
                                    }
                                }
                            }
                            switch -2147483648 {
                            case 32767:
                                var v182 uint8 = 89
                                _ = v182
                            case 128:
                                var v183 float64 = v4
                                _ = v183
                            case 1:
                                fmt.Println(v9, v4)
                            }
                            _ = v168
                        } else {
                            var v184 float64 = v4
                            var v185 int32 = v3
                            var v186 string = "a"
                            _ = v184
                            _ = v185
                            _ = v186
                        }
                        var v187 int8 = v6
                        var v188 int32 = v166
                        _ = fi167
                        _ = v187
                        _ = v188
                    }
                case -567:
                    v189 := []int{}
                    _ = v189
                default:
                    if true {
                        var v190 float64 = (0.0 / 1e10)
                        v21 = v21
                        for fi191 := 0; fi191 < 12; fi191++ {
                            v192 := GenS3{F0: 935, F1: 255, F2: 32}
                            _ = fi191
                            _ = v192
                        }
                        _ = v190
                    } else {
                        if true {
                            var v193 int16 = 42
                            _ = v193
                        } else {
                            v13 = v13
                        }
                    }
                }
                v194 := []int{(^(fi165))}
                _ = v194
            }
            v195 := []int{629, 1}
            _ = fi165
            _ = v166
            _ = v195
        }
    case 32768:
        fmt.Println(v3)
    default:
        var v196 uint16 = 414
        _ = v196
    }
    var v197 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v164:", v164)
    fmt.Println("v197:", v197)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v164
    _ = v197
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
