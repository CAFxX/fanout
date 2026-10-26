package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint; F2 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 int16; F1 uint; F2 int8 }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint32; F1 int8 }
type GenS14 struct { F0 string; F1 uint }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint8; F1 int16; F2 bool }
type GenS18 struct { F0 uint16; F1 int16; F2 bool }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint; F1 uint; F2 int16 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 int16; F1 int32 }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 string; F1 int64 }
type GenS27 struct { F0 int16; F1 int8 }
type GenS28 struct { F0 uint16 }
type GenS29 struct { F0 bool }
type GenS30 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS31 struct { F0 bool }
type GenS32 struct { F0 uint32 }
func genFn0(p0 int, p1 bool, p2 int32) int8 {
    var v5 bool = false
    p1 = false
    v6 := GenS1{F0: 18}
    var v7 int32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 86
}
func genFn1(p0 int64) int16 {
    var v8 int8 = 42
    v9 := GenS1{F0: 105}
    v10 := GenS2{F0: 451, F1: (p0 + p0), F2: (449 != 610)}
    var v11 int64 = p0
    if false {
        var v12 float64 = 0.1
        var v13 int16 = 1
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -685
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 284
    v3 := make(map[uint32]uint32)
    var v4 int64 = 0
    v2 = v2
    switch -1 {
    case 2147483647:
        v14 := []int{32767, -643, 543}
        _ = v14
    default:
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{32768, -2147483648, fi15, -746}
            for fi17 := 0; fi17 < 17; fi17++ {
                _ = len(v1)
                v2 -= v2
                _ = fi17
            }
            var v18 string = ("a" + ("\n\t" + "a"))
            _ = fi15
            _ = v16
            _ = v18
        }
    }
    v2 = (v2 | (v2 / (((v2 & 1)) | 1)))
    v19 := GenS2{F0: v4, F1: 1, F2: (1e10 > (0.1 * -2.25))}
    fmt.Println(v2)
    switch -393 {
    case -1:
        if (1 >= 1000) {
            var v20 int32 = v2
            var v21 bool = (0.1 > (0.0 / (0.0 - 1e10)))
            _ = v20
            _ = v21
        } else {
            switch -987 {
            case 128:
                v4 = v4
            }
        }
    case 9223372036854775807:
        switch 0 {
        case 77:
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 uint8 = 58
                _ = fi22
                _ = v23
            }
        }
    case 32767:
        if false {
            v24 := GenS2{F0: v4, F1: (v4 % ((42) | 1)), F2: true}
            fmt.Println(v2, v2)
            v25 := GenS2{F0: v4, F1: v4, F2: false}
            _ = v24
            _ = v25
        } else {
            if true {
                v4 += (v4 >> ((65535) & 63))
                if (18446744073709551615 > 18446744073709551615) {
                    if true {
                        _ = len(v1)
                        for fi26 := 0; fi26 < 18; fi26++ {
                            v4 = v4
                            if (true || ((42 > -1) || false)) {
                                if false {
                                    for fi27 := 0; fi27 < 13; fi27++ {
                                        if (((0.1 / 1.0) != ((1.5 / -2.25) * 0.1)) || (1.5 == (0.1 + 1e10))) {
                                            v28 := GenS1{F0: 128}
                                            _ = v28
                                        }
                                        if (!true) {
                                            fmt.Println(fi26, fi26)
                                            if ("\n\t" < "") {
                                                if true {
                                                    var v29 int = -103
                                                    switch "" {
                                                    case "foo bar":
                                                        var v30 int8 = 100
                                                        _ = v30
                                                    case "\n\t":
                                                        fmt.Println(fi27)
                                                    case "a":
                                                        v31 := GenS3{F0: 1000, F1: 97}
                                                        _ = v31
                                                    }
                                                    v1[2] = "foo bar"
                                                    _ = v29
                                                }
                                                _ = len(v3)
                                                fmt.Println(fi27, v4)
                                            } else {
                                                if true {
                                                    v32 := v3[2147483648]
                                                    var v33 uint8 = 127
                                                    switch (fi27 * (fi26 & fi26)) {
                                                    case 784:
                                                        if false {
                                                            if (v4 >= (v4 * 42)) {
                                                                if (!false) {
                                                                    v2 -= v2
                                                                    if false {
                                                                        if (false || true) {
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v34 uint = 4294967295
                                                                                _ = v34
                                                                            case "\n\t":
                                                                                v35 := v3[2]
                                                                                _ = v35
                                                                            }
                                                                            if (1000 > 1000) {
                                                                                v36 := GenS4{F0: v32}
                                                                                var v37 int8 = 80
                                                                                if (float32(3.14159) <= (((-2.25 + 3.14159) / 0.1) + 3.14159)) {
                                                                                    v38 := v1[(v32 * v32)]
                                                                                    v3[v32] = v32
                                                                                    v39 := []int{int(v4), (fi27 << ((65535) & 63))}
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                }
                                                                                _ = v36
                                                                                _ = v37
                                                                            } else {
                                                                                var v40 int32 = (v2 - v2)
                                                                                v4 -= v4
                                                                                v32 -= v32
                                                                                _ = v40
                                                                            }
                                                                            v41 := []int{42, fi27, (fi26 * fi26), fi26}
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        v1[(v32 >> ((42) & 63))] = ("" + "hello")
                                                                        switch "" {
                                                                        case "hello":
                                                                            v42 := v1[v32]
                                                                            _ = v42
                                                                        case "foo bar":
                                                                            v43 := GenS5{F0: 569, F1: 554, F2: ((("foo bar" + "a") + ("hello" + ("" + "\n\t"))) + "foo bar")}
                                                                            _ = v43
                                                                        }
                                                                        v33 -= 94
                                                                    }
                                                                } else {
                                                                    v44 := GenS1{F0: 1}
                                                                    _ = v44
                                                                }
                                                                v45 := GenS6{F0: (fi27 % ((fi26) | 1))}
                                                                switch (fi27 & -377) {
                                                                case 4294967295:
                                                                    switch fi26 {
                                                                    case 42:
                                                                        fmt.Println(v33, v33)
                                                                    case 343:
                                                                        v2 -= v2
                                                                    case -2147483648:
                                                                        v46 := []int{(fi26 - 0)}
                                                                        _ = v46
                                                                    default:
                                                                        fmt.Println(v33, fi26, fi27)
                                                                    }
                                                                    fmt.Println(fi27, v33, fi26)
                                                                    v4 -= int64(v2)
                                                                case -2147483648:
                                                                    v47 := GenS3{F0: 718, F1: (^(v33))}
                                                                    _ = v47
                                                                case -237:
                                                                    switch (-(fi26)) {
                                                                    case 32768:
                                                                        var v48 string = "foo bar"
                                                                        _ = v48
                                                                    default:
                                                                        if (384 <= 845) {
                                                                            if false {
                                                                                fmt.Println(fi26, v4, fi27)
                                                                                var v49 uint64 = 92
                                                                                v2 = v2
                                                                                switch (-(fi26)) {
                                                                                case 9223372036854775807:
                                                                                    var v50 int = fi27
                                                                                    _ = v50
                                                                                case 2147483647:
                                                                                    v51 := []int{fi26, fi27, 496}
                                                                                    _ = v51
                                                                                }
                                                                                _ = v49
                                                                            }
                                                                            fmt.Println(v32, fi27)
                                                                            var v52 uint64 = 2
                                                                            _ = v52
                                                                        } else {
                                                                            var v53 uint = 1
                                                                            v4 += (v4 % (((v4 >> ((v53) & 63))) | 1))
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                default:
                                                                    var v54 string = ("" + "")
                                                                    _ = v54
                                                                }
                                                                _ = v45
                                                            } else {
                                                                v55 := GenS1{F0: v33}
                                                                _ = v55
                                                            }
                                                        } else {
                                                            v4 = v4
                                                            v56 := GenS1{F0: 255}
                                                            v32 -= v32
                                                            _ = v56
                                                        }
                                                    default:
                                                        v57 := []int{(fi26 & fi26), fi26}
                                                        _ = v57
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    switch ("a" + "hello") {
                                                    case "foo bar":
                                                        var v58 bool = (85 <= 128)
                                                        _ = v58
                                                    default:
                                                        v59 := GenS5{F0: 42, F1: 4294967295, F2: "\n\t"}
                                                        _ = v59
                                                    }
                                                    if true {
                                                        fmt.Println(fi26, fi27)
                                                        var v60 int16 = genFn1(v4)
                                                        _ = v60
                                                    }
                                                }
                                                v61 := GenS7{F0: 127}
                                                _ = v61
                                            }
                                            v62 := GenS8{F0: 4294967295, F1: 42}
                                            v4 = 0
                                            var v63 uint16 = 198
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = fi27
                                    }
                                    v2 += -605
                                }
                                var v64 string = "hello"
                                v65 := GenS7{F0: 2}
                                _ = v64
                                _ = v65
                            } else {
                                v4 -= (v4 % ((332) | 1))
                                var v66 uint = 42
                                _ = v66
                            }
                            _ = fi26
                        }
                    }
                } else {
                    v67 := GenS6{F0: 127}
                    if true {
                        v4 -= v4
                        if (0.0 > 1e10) {
                            if true {
                                v2 = v2
                                var v68 bool = true
                                v2 = 0
                                _ = v68
                            } else {
                                var v69 uint16 = 2
                                switch 0 {
                                case 444:
                                    var v70 float64 = 3.14159
                                    _ = v70
                                default:
                                    var v71 uint16 = 0
                                    _ = v71
                                }
                                var v72 int16 = -32768
                                _ = v69
                                _ = v72
                            }
                        } else {
                            fmt.Println(v2, v2)
                            v2 = (v2 ^ (-(v2)))
                        }
                    } else {
                        if (true || true) {
                            var v73 int = -2147483648
                            var v74 float64 = 1.5
                            v4 = 214
                            _ = v73
                            _ = v74
                        } else {
                            switch "a" {
                            case "":
                                if ((-2.25 + float32((float64(3.14159) - (1e10 + -2.25)))) > -2.25) {
                                    var v75 int64 = 42
                                    var v76 bool = (!true)
                                    if false {
                                        for fi77 := 0; fi77 < 10; fi77++ {
                                            var v78 int8 = 100
                                            var v79 int64 = (-(v75))
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                        switch "" {
                                        case "\n\t":
                                            v80 := []int{}
                                            _ = v80
                                        case "hello":
                                            v3[2] = 42
                                        case "foo bar":
                                            v81 := GenS6{F0: -260}
                                            _ = v81
                                        default:
                                            v75 += v4
                                        }
                                        if (v76 && v76) {
                                            fmt.Println(v4, v2, v4)
                                            switch 127 {
                                            case -1:
                                                if true {
                                                    v82 := GenS3{F0: 492, F1: 2}
                                                    _ = v82
                                                }
                                            case -316:
                                                for fi83 := 0; fi83 < 17; fi83++ {
                                                    v76 = v76
                                                    var v84 uint = 536
                                                    _ = fi83
                                                    _ = v84
                                                }
                                            case 0:
                                                if false {
                                                    if false {
                                                        if ((873 == 18446744073709551615) || v76) {
                                                            v4 = v4
                                                            var v85 string = "hello"
                                                            _ = v85
                                                        } else {
                                                            if v76 {
                                                                v86 := GenS1{F0: 0}
                                                                _ = v86
                                                            }
                                                            v76 = true
                                                            if true {
                                                                switch ("" + "a") {
                                                                case "foo bar":
                                                                    var v87 int16 = genFn1(v75)
                                                                    _ = v87
                                                                case "hello":
                                                                    _ = len(v1)
                                                                default:
                                                                    for fi88 := 0; fi88 < 9; fi88++ {
                                                                        v76 = false
                                                                        switch int(v75) {
                                                                        case -1:
                                                                            if true {
                                                                                var v89 int64 = (^(v4))
                                                                                _ = v89
                                                                            } else {
                                                                                for fi90 := 0; fi90 < 3; fi90++ {
                                                                                    v91 := []int{221, fi88}
                                                                                    _ = fi90
                                                                                    _ = v91
                                                                                }
                                                                                for fi92 := 0; fi92 < 19; fi92++ {
                                                                                    if v76 {
                                                                                        delete(v3, 455)
                                                                                        var v93 int64 = (v4 * v75)
                                                                                        _ = v93
                                                                                    }
                                                                                    fmt.Println(fi92, v2)
                                                                                    _ = fi92
                                                                                }
                                                                                v2 += v2
                                                                            }
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v75, v75)
                                                                        default:
                                                                            switch fi88 {
                                                                            case 429:
                                                                                v94 := []int{}
                                                                                _ = v94
                                                                            case -1:
                                                                                if v76 {
                                                                                    v75 -= v4
                                                                                    for fi95 := 0; fi95 < 13; fi95++ {
                                                                                        v2 += (-(v2))
                                                                                        var v96 int32 = (v2 / (((v2 + -2147483648)) | 1))
                                                                                        _ = fi95
                                                                                        _ = v96
                                                                                    }
                                                                                    for fi97 := 0; fi97 < 10; fi97++ {
                                                                                        if false {
                                                                                            var v98 uint32 = 2147483648
                                                                                            switch ("\n\t" + "hello") {
                                                                                            case "\n\t":
                                                                                                if (v75 > v4) {
                                                                                                    var v99 uint32 = (v98 / ((v98) | 1))
                                                                                                    v100 := v3[v99]
                                                                                                    v2 += (v2 ^ 2147483647)
                                                                                                    _ = v100
                                                                                                    _ = v99
                                                                                                }
                                                                                            case "hello":
                                                                                                var v101 bool = true
                                                                                                _ = v101
                                                                                            }
                                                                                            var v102 int = fi97
                                                                                            _ = v102
                                                                                            _ = v98
                                                                                        }
                                                                                        v103 := GenS2{F0: (-(v75)), F1: v4, F2: (6 >= 513)}
                                                                                        _ = fi97
                                                                                        _ = v103
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                if v76 {
                                                                                    if (65535 <= 841) {
                                                                                        v104 := v1[2]
                                                                                        v105 := GenS8{F0: 902, F1: 1}
                                                                                        _ = v104
                                                                                        _ = v105
                                                                                    } else {
                                                                                        switch (-(fi88)) {
                                                                                        case 32767:
                                                                                            fmt.Println(v75, v75, v4)
                                                                                        case -2147483648:
                                                                                            if ((0.0 - 0.1) < (1.5 * 0.1)) {
                                                                                                v76 = v76
                                                                                                if v76 {
                                                                                                    v106 := GenS4{F0: 1}
                                                                                                    if false {
                                                                                                        fmt.Println(v75)
                                                                                                        var v107 uint16 = 42
                                                                                                        var v108 string = (("" + "") + "a")
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    } else {
                                                                                                        fmt.Println(v76, v2, v2)
                                                                                                        switch ("foo bar" + ("foo bar" + ("a" + "a"))) {
                                                                                                        case "\n\t":
                                                                                                            for fi109 := 0; fi109 < 10; fi109++ {
                                                                                                                v2 -= (-(v2))
                                                                                                                var v110 float64 = ((((0.0 / 0.1) * (0.0 * -2.25)) + (1e10 - 1.5)) * (-2.25 + ((1.5 - (3.14159 - 1e10)) - -2.25)))
                                                                                                                if (85 > 2) {
                                                                                                                    v111 := []int{(fi88 ^ fi109), (fi109 / ((fi88) | 1)), (fi109 + fi109), fi109}
                                                                                                                    _ = v111
                                                                                                                }
                                                                                                                _ = fi109
                                                                                                                _ = v110
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v112 int16 = genFn1((v4 & (v75 * -9223372036854775808)))
                                                                                                            _ = v112
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(fi88, v2, v76)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v106
                                                                                                }
                                                                                                for fi113 := 0; fi113 < 2; fi113++ {
                                                                                                    if v76 {
                                                                                                        var v114 int64 = v4
                                                                                                        v115 := GenS8{F0: 1, F1: 692}
                                                                                                        if v76 {
                                                                                                            var v116 bool = (0.1 <= 3.14159)
                                                                                                            var v117 int8 = -1
                                                                                                            if false {
                                                                                                                if v76 {
                                                                                                                    v1[152] = "hello"
                                                                                                                } else {
                                                                                                                    v117 += v117
                                                                                                                    var v118 int16 = -180
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                                switch (((("hello" + "a") + "a") + ("" + "foo bar")) + ("" + "foo bar")) {
                                                                                                                case "foo bar":
                                                                                                                    var v119 int8 = v117
                                                                                                                    _ = v119
                                                                                                                case "":
                                                                                                                    var v120 uint8 = 0
                                                                                                                    _ = v120
                                                                                                                case "a":
                                                                                                                    var v121 uint32 = 1
                                                                                                                    _ = v121
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v75, v114)
                                                                                                                v117 += v117
                                                                                                                if (-2.25 == 3.14159) {
                                                                                                                    v122 := v1[42]
                                                                                                                    var v123 uint16 = 0
                                                                                                                    var v124 uint64 = 350
                                                                                                                    var v125 float64 = ((3.14159 + (-2.25 - 0.0)) - ((-2.25 + -2.25) * (-2.25 * (3.14159 + (((1.5 - 1e10) + float64(1.5)) * 1e10)))))
                                                                                                                    if false {
                                                                                                                        v1[1] = ("foo bar" + (v122 + ("foo bar" + "\n\t")))
                                                                                                                        var v126 string = "hello"
                                                                                                                        _ = v126
                                                                                                                    }
                                                                                                                    _ = v122
                                                                                                                    _ = v123
                                                                                                                    _ = v124
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v116
                                                                                                            _ = v117
                                                                                                        }
                                                                                                        _ = v114
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    v4 = v75
                                                                                                    fmt.Println(v4)
                                                                                                    _ = fi113
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v76, v2, v2)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v127 uint32 = 543
                                                                        _ = fi88
                                                                        _ = v127
                                                                    }
                                                                }
                                                                fmt.Println(v2)
                                                            }
                                                        }
                                                    }
                                                    var v128 bool = false
                                                    _ = v128
                                                }
                                            }
                                            var v129 uint32 = 4294967295
                                            _ = v129
                                        }
                                    } else {
                                        var v130 int64 = 42
                                        v76 = (1e10 <= 1.5)
                                        _ = v130
                                    }
                                    _ = v75
                                    _ = v76
                                }
                            case "foo bar":
                                v2 -= int32(v4)
                            case "\n\t":
                                switch "hello" {
                                case "":
                                    var v131 float32 = (3.14159 * (((0.0 * (1e10 * (3.14159 + 1e10))) - 0.1) * 1e10))
                                    _ = v131
                                case "a":
                                    v4 -= (-(v4))
                                case "hello":
                                    if true {
                                        if false {
                                            var v132 string = "hello"
                                            if (float64(((3.14159 + 3.14159) + -2.25)) > float64(-2.25)) {
                                                v132 = "hello"
                                                if (1 == genFn0(128, false, int32(v4))) {
                                                    for fi133 := 0; fi133 < 8; fi133++ {
                                                        var v134 int16 = 830
                                                        v135 := GenS9{F0: false, F1: 221}
                                                        delete(v1, 42)
                                                        _ = fi133
                                                        _ = v134
                                                        _ = v135
                                                    }
                                                    fmt.Println(v4)
                                                } else {
                                                    for fi136 := 0; fi136 < 13; fi136++ {
                                                        v137 := []int{}
                                                        var v138 int32 = 0
                                                        _ = len(v3)
                                                        _ = fi136
                                                        _ = v137
                                                        _ = v138
                                                    }
                                                }
                                            }
                                            _ = v132
                                        } else {
                                            var v139 uint = 0
                                            v140 := GenS6{F0: 2147483647}
                                            for fi141 := 0; fi141 < 6; fi141++ {
                                                var v142 int16 = 42
                                                _ = fi141
                                                _ = v142
                                            }
                                            _ = v139
                                            _ = v140
                                        }
                                        if false {
                                            v2 += v2
                                            for fi143 := 0; fi143 < 20; fi143++ {
                                                for fi144 := 0; fi144 < 16; fi144++ {
                                                    if false {
                                                        v2 += (v2 * (v2 << ((42) & 63)))
                                                        if (!(0 == 4294967295)) {
                                                            if (!(!false)) {
                                                                fmt.Println(fi144)
                                                                var v145 uint = 18446744073709551615
                                                                _ = v145
                                                            }
                                                            if true {
                                                                var v146 int16 = genFn1(v4)
                                                                _ = v146
                                                            }
                                                        }
                                                    }
                                                    v2 = (v2 | 0)
                                                    _ = fi144
                                                }
                                                _ = fi143
                                            }
                                            if false {
                                                if true {
                                                    fmt.Println(v2, v2)
                                                }
                                                if true {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v147 int8 = 100
                                                        _ = v147
                                                    default:
                                                        v148 := GenS10{F0: 273, F1: 255, F2: -128}
                                                        _ = v148
                                                    }
                                                    v149 := GenS11{F0: 42, F1: ("foo bar" + "")}
                                                    _ = v149
                                                }
                                            }
                                        }
                                    } else {
                                        v4 = v4
                                        v2 -= 42
                                        var v150 uint32 = 1
                                        _ = v150
                                    }
                                default:
                                    if (!(0 < 4294967295)) {
                                        v151 := []int{}
                                        v4 += 484
                                        if (!true) {
                                            fmt.Println(v2, v2)
                                            v2 = (v2 - v2)
                                        } else {
                                            v2 = v2
                                        }
                                        _ = v151
                                    }
                                }
                            }
                            for fi152 := 0; fi152 < 12; fi152++ {
                                var v153 int16 = 1000
                                if true {
                                    v4 += v4
                                    v154 := []int{32768, 9223372036854775807}
                                    _ = v154
                                }
                                delete(v3, 42)
                                _ = fi152
                                _ = v153
                            }
                        }
                        v155 := GenS6{F0: 9223372036854775807}
                        var v156 int32 = v2
                        _ = v155
                        _ = v156
                    }
                    _ = v67
                }
                v4 += 42
            } else {
                var v157 int16 = 10
                fmt.Println(v157, v157, v157)
                v2 += (-(v2))
                _ = v157
            }
            v2 += v2
        }
    }
    var v158 string = "a"
    var v159 uint16 = 255
    v2 = (^(v2))
    var v160 uint64 = 0
    if true {
        _ = len(v3)
        switch "\n\t" {
        case "":
            v4 -= 1
        }
    }
    for fi161 := 0; fi161 < 13; fi161++ {
        for fi162 := 0; fi162 < 7; fi162++ {
            if (int32(v159) >= -258) {
                switch ("\n\t" + v158) {
                case "\n\t":
                    var v163 float32 = -2.25
                    _ = v163
                case "hello":
                    fmt.Println(fi161)
                    switch "a" {
                    case "hello":
                        if true {
                            if (genFn0((fi162 >> ((4294967295) & 63)), (!false), (v2 & v2)) == 23) {
                                var v164 float32 = (0.1 / 1.0)
                                var v165 int64 = v4
                                var v166 int64 = (v4 >> ((523) & 63))
                                v167 := []int{}
                                v1[42] = "hello"
                                delete(v1, 2147483647)
                                var v168 float32 = 1.5
                                _ = v164
                                _ = v165
                                _ = v166
                                _ = v167
                                _ = v168
                            } else {
                                fmt.Println(v159)
                                var v169 float64 = 0.1
                                var v170 float32 = 3.14159
                                v170 += v170
                                _ = len(v1)
                                _ = v169
                                _ = v170
                            }
                            var v171 bool = true
                            v172 := []int{}
                            _ = v171
                            _ = v172
                        } else {
                            if true {
                                if true {
                                    if (!((v2 >> ((18446744073709551615) & 63)) != v2)) {
                                        v173 := GenS9{F0: false, F1: 127}
                                        fmt.Println(v159)
                                        v159 -= (v159 >> ((950) & 63))
                                        _ = v173
                                    }
                                }
                            }
                            var v174 string = (v158 + "foo bar")
                            _ = v174
                        }
                    case "foo bar":
                        var v175 string = ("foo bar" + v158)
                        _ = v175
                    case "":
                        var v176 uint32 = 644
                        _ = v176
                    }
                    v177 := []int{-2147483648}
                    if (!(!(42 > 0))) {
                        switch ((v158 + (((v158 + (v158 + "\n\t")) + v158) + "foo bar")) + "foo bar") {
                        case "hello":
                            if false {
                                v178 := GenS3{F0: 1, F1: 2}
                                var v179 uint8 = 128
                                _ = v178
                                _ = v179
                            } else {
                                if (true || true) {
                                    v180 := GenS12{F0: int32(v160)}
                                    var v181 uint8 = 127
                                    if (!false) {
                                        if (!false) {
                                            v158 = "hello"
                                            var v182 int = (-(fi161))
                                            switch fi161 {
                                            case 42:
                                                var v183 uint = 65535
                                                _ = v183
                                            case 128:
                                                var v184 uint8 = (v181 % ((0) | 1))
                                                _ = v184
                                            case 110:
                                                v185 := GenS13{F0: 804, F1: 49}
                                                _ = v185
                                            default:
                                                fmt.Println(v160, v4)
                                            }
                                            var v186 int64 = (v4 >> ((42) & 63))
                                            _ = v182
                                            _ = v186
                                        }
                                        switch (v158 + v158) {
                                        case "":
                                            var v187 uint16 = v159
                                            _ = v187
                                        case "\n\t":
                                            var v188 float32 = (((3.14159 * 3.14159) * 0.1) + (1e10 * 0.0))
                                            _ = v188
                                        case "a":
                                            if false {
                                                if len(v177) > 0 {
                                                    v177[0] = fi162
                                                }
                                                var v189 bool = (true && (true || (true || true)))
                                                if (665 != (fi162 & fi162)) {
                                                    v190 := GenS14{F0: v158, F1: 948}
                                                    _ = v190
                                                } else {
                                                    if (!false) {
                                                        switch (^(fi161)) {
                                                        case 32768:
                                                            if (!(((v181 ^ v181) != v181) && false)) {
                                                                _ = len(v177)
                                                            }
                                                        case 128:
                                                            fmt.Println(fi161, v4, v158)
                                                        case 42:
                                                            var v191 uint16 = v159
                                                            fmt.Println(v191)
                                                            switch fi162 {
                                                            case 32768:
                                                                if (v189 || (!(v189 || (v189 || v189)))) {
                                                                    v192 := GenS3{F0: 32767, F1: 128}
                                                                    if v189 {
                                                                        if v189 {
                                                                            switch v158 {
                                                                            case "hello":
                                                                                v193 := v1[118]
                                                                                _ = v193
                                                                            case "foo bar":
                                                                                v194 := GenS6{F0: fi162}
                                                                                _ = v194
                                                                            }
                                                                        }
                                                                    }
                                                                    var v195 int8 = 127
                                                                    var v196 uint64 = (v160 * (v160 % ((v160) | 1)))
                                                                    _ = v192
                                                                    _ = v195
                                                                    _ = v196
                                                                } else {
                                                                    if len(v177) >= 2 {
                                                                        v177 = v177[1:len(v177)]
                                                                    }
                                                                    v160 = v160
                                                                    if ((650 != 137) || false) {
                                                                        v191 -= 916
                                                                        fmt.Println(v159)
                                                                        var v197 int = 42
                                                                        _ = v197
                                                                    } else {
                                                                        v198 := GenS4{F0: 1}
                                                                        switch -9223372036854775808 {
                                                                        case 42:
                                                                            switch v158 {
                                                                            case "a":
                                                                                var v199 uint32 = 42
                                                                                _ = v199
                                                                            default:
                                                                                var v200 uint32 = 4294967295
                                                                                _ = v200
                                                                            }
                                                                        default:
                                                                            var v201 uint = 0
                                                                            _ = v201
                                                                        }
                                                                        if ((v160 << ((123) & 63)) == 407) {
                                                                            if ((42 >= (fi162 + 9223372036854775807)) && v189) {
                                                                                var v202 uint16 = (v191 & (v191 >> ((556) & 63)))
                                                                                _ = v202
                                                                            }
                                                                            v159 = v191
                                                                            v203 := GenS7{F0: (v181 ^ v181)}
                                                                            _ = v203
                                                                        } else {
                                                                            switch v158 {
                                                                            case "a":
                                                                                var v204 int64 = (v4 << ((295) & 63))
                                                                                _ = v204
                                                                            }
                                                                        }
                                                                        v205 := GenS3{F0: 42, F1: 255}
                                                                        _ = v198
                                                                        _ = v205
                                                                    }
                                                                    v1[2] = (v158 + (v158 + ""))
                                                                    var v206 float32 = 0.1
                                                                    v189 = (!v189)
                                                                    var v207 int32 = -389
                                                                    _ = v206
                                                                    _ = v207
                                                                }
                                                            default:
                                                                v208 := GenS13{F0: 42, F1: -1}
                                                                _ = v208
                                                            }
                                                            _ = v191
                                                        default:
                                                            switch (v158 + v158) {
                                                            case "hello":
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        delete(v1, 0)
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v4 -= (v4 >> ((18446744073709551615) & 63))
                                                                v1[2] = "foo bar"
                                                                if len(v177) > 0 {
                                                                    v177[0] = 32768
                                                                }
                                                                v189 = v189
                                                                var v209 int = fi162
                                                                v2 -= v2
                                                                _ = v209
                                                            case "\n\t":
                                                                if true {
                                                                    var v210 uint8 = (v181 % ((v181) | 1))
                                                                    v189 = (!false)
                                                                    _ = v210
                                                                } else {
                                                                    var v211 int16 = 125
                                                                    fmt.Println(v160)
                                                                    _ = v211
                                                                }
                                                            }
                                                        }
                                                        v181 -= v181
                                                        switch v158 {
                                                        case "foo bar":
                                                            var v212 uint32 = 375
                                                            _ = v212
                                                        case "":
                                                            if (!v189) {
                                                                v213 := GenS5{F0: (v160 - v160), F1: 0, F2: v158}
                                                                if v189 {
                                                                    switch (^(fi161)) {
                                                                    case -9223372036854775808:
                                                                        switch 2147483647 {
                                                                        case 127:
                                                                            fmt.Println(v189, fi161)
                                                                        case -588:
                                                                            v159 = v159
                                                                        }
                                                                    case 9223372036854775807:
                                                                        v160 = (^(v160))
                                                                    default:
                                                                        if (2 <= (v160 << ((0) & 63))) {
                                                                            fmt.Println(fi162, v159)
                                                                            var v214 float64 = 3.14159
                                                                            if v189 {
                                                                                if (1 > 32767) {
                                                                                    if (1 < (v181 + (v181 % ((v181) | 1)))) {
                                                                                        var v215 int = -584
                                                                                        _ = v215
                                                                                    } else {
                                                                                        v216 := GenS7{F0: v181}
                                                                                        _ = v216
                                                                                    }
                                                                                    fmt.Println(fi162, v214, v214)
                                                                                }
                                                                                delete(v3, 2147483648)
                                                                                switch fi161 {
                                                                                case 1:
                                                                                    if len(v177) > 0 {
                                                                                        v177[0] = (fi161 | (-(fi161)))
                                                                                    }
                                                                                default:
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v217 int32 = 1
                                                                                        _ = v217
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if ((fi161 | fi162) <= fi161) {
                                                                                    var v218 string = v158
                                                                                    _ = v218
                                                                                } else {
                                                                                    fmt.Println(v160, v158, v189)
                                                                                    switch (v158 + v158) {
                                                                                    case "":
                                                                                        if (((-32768 == 0) || (!v189)) || v189) {
                                                                                            var v219 int16 = 0
                                                                                            delete(v1, 2)
                                                                                            v4 -= -953
                                                                                            _ = v219
                                                                                        }
                                                                                    default:
                                                                                        switch ("a" + v158) {
                                                                                        case "hello":
                                                                                            v1[2147483647] = "\n\t"
                                                                                        case "foo bar":
                                                                                            v3[2] = 2
                                                                                        }
                                                                                    }
                                                                                    v159 -= v159
                                                                                }
                                                                            }
                                                                            _ = v214
                                                                        }
                                                                    }
                                                                    v158 = v158
                                                                }
                                                                if (3.14159 > 0.0) {
                                                                    var v220 int64 = v4
                                                                    _ = v220
                                                                } else {
                                                                    v221 := GenS9{F0: v189, F1: 252}
                                                                    _ = v221
                                                                }
                                                                _ = v213
                                                            }
                                                        case "\n\t":
                                                            v3[2147483648] = 4294967295
                                                        default:
                                                            if ((fi161 & fi162) == fi162) {
                                                                var v222 int = (fi161 >> ((18446744073709551615) & 63))
                                                                var v223 string = ("a" + ((v158 + (v158 + (v158 + v158))) + ("\n\t" + v158)))
                                                                _ = v222
                                                                _ = v223
                                                            }
                                                        }
                                                    } else {
                                                        var v224 int32 = v2
                                                        var v225 string = v158
                                                        v226 := GenS11{F0: 698, F1: "a"}
                                                        var v227 string = (v225 + "\n\t")
                                                        v189 = false
                                                        _ = v224
                                                        _ = v225
                                                        _ = v226
                                                        _ = v227
                                                    }
                                                    if (uint(v160) <= 665) {
                                                        if v189 {
                                                            var v228 int64 = v4
                                                            _ = v228
                                                        } else {
                                                            if v189 {
                                                                switch (fi161 & 128) {
                                                                case -9223372036854775808:
                                                                    if v189 {
                                                                        var v229 float64 = 3.14159
                                                                        v4 -= -1
                                                                        var v230 uint64 = v160
                                                                        _ = v229
                                                                        _ = v230
                                                                    } else {
                                                                        if (true || (-15 == int8(v160))) {
                                                                            v189 = (false && (637 >= 0))
                                                                        } else {
                                                                            v160 = (v160 ^ uint64(v181))
                                                                            switch fi162 {
                                                                            case -9223372036854775808:
                                                                                var v231 string = v158
                                                                                v189 = (66 <= -68)
                                                                                v181 += v181
                                                                                if (-80 <= -38) {
                                                                                    v181 = v181
                                                                                    if true {
                                                                                        switch v231 {
                                                                                        case "hello":
                                                                                            var v232 bool = true
                                                                                            _ = v232
                                                                                        case "foo bar":
                                                                                            var v233 uint16 = v159
                                                                                            _ = v233
                                                                                        case "":
                                                                                            _ = len(v177)
                                                                                        }
                                                                                        var v234 float32 = (0.1 - 3.14159)
                                                                                        var v235 int16 = 0
                                                                                        _ = v234
                                                                                        _ = v235
                                                                                    }
                                                                                    var v236 uint32 = uint32(v160)
                                                                                    _ = v236
                                                                                } else {
                                                                                    if v189 {
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            v237 := GenS3{F0: 950, F1: v181}
                                                                                            if (!v189) {
                                                                                                fmt.Println(v189, v4, fi161)
                                                                                                v238 := v3[1]
                                                                                                v231 = "hello"
                                                                                                _ = v238
                                                                                            } else {
                                                                                                var v239 int16 = 1
                                                                                                v240 := GenS15{F0: (v239 - genFn1(104))}
                                                                                                _ = v239
                                                                                                _ = v240
                                                                                            }
                                                                                            v158 = (v158 + v158)
                                                                                            _ = v237
                                                                                        }
                                                                                        delete(v3, 2147483647)
                                                                                        if (!(1.5 >= 1.5)) {
                                                                                            var v241 uint32 = 2147483648
                                                                                            if (!(-1 != -58)) {
                                                                                                _ = len(v1)
                                                                                                v242 := v1[(^(v241))]
                                                                                                fmt.Println(v189, fi161)
                                                                                                _ = v242
                                                                                            }
                                                                                            v189 = true
                                                                                            var v243 int16 = -32768
                                                                                            v244 := GenS16{F0: ((v160 >> ((255) & 63)) >= (^(v160)))}
                                                                                            _ = v241
                                                                                            _ = v243
                                                                                            _ = v244
                                                                                        } else {
                                                                                            var v245 int16 = 1000
                                                                                            switch (fi162 + fi161) {
                                                                                            case 4294967295:
                                                                                                var v246 float32 = 0.1
                                                                                                var v247 uint8 = 127
                                                                                                _ = v246
                                                                                                _ = v247
                                                                                            case 127:
                                                                                                v2 = (v2 + (-(v2)))
                                                                                            }
                                                                                            switch (fi162 / (((fi162 - (fi161 % ((fi161) | 1)))) | 1)) {
                                                                                            case -440:
                                                                                                switch v158 {
                                                                                                case "hello":
                                                                                                    if len(v177) > 0 {
                                                                                                        v177[0] = (fi161 * (fi162 << ((0) & 63)))
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v159 += (v159 | v159)
                                                                                                case "\n\t":
                                                                                                    var v248 int = fi161
                                                                                                    _ = v248
                                                                                                default:
                                                                                                    v160 = (v160 * (v160 & v160))
                                                                                                }
                                                                                                _ = len(v177)
                                                                                                if v189 {
                                                                                                    v159 = v159
                                                                                                    v249 := GenS4{F0: 0}
                                                                                                    if true {
                                                                                                        if (!v189) {
                                                                                                            if v189 {
                                                                                                                v250 := GenS17{F0: (^(v181)), F1: (v245 * v245), F2: v189}
                                                                                                                v2 = (v2 * (-(v2)))
                                                                                                                _ = v250
                                                                                                            } else {
                                                                                                                var v251 int64 = v4
                                                                                                                v245 = (v245 ^ 1000)
                                                                                                                _ = v251
                                                                                                            }
                                                                                                            v158 = v231
                                                                                                            v252 := GenS16{F0: (!(!((^(fi162)) == (fi161 << ((65535) & 63)))))}
                                                                                                            _ = v252
                                                                                                        } else {
                                                                                                            if len(v177) >= 2 {
                                                                                                                v177 = v177[1:len(v177)]
                                                                                                            }
                                                                                                            var v253 uint64 = 18446744073709551615
                                                                                                            v254 := GenS9{F0: (32768 >= (fi162 / ((fi162) | 1))), F1: v181}
                                                                                                            _ = v253
                                                                                                            _ = v254
                                                                                                        }
                                                                                                        var v255 int64 = (v4 % ((v4) | 1))
                                                                                                        v256 := GenS18{F0: (^(v159)), F1: -601, F2: v189}
                                                                                                        _ = v255
                                                                                                        _ = v256
                                                                                                    }
                                                                                                    _ = v249
                                                                                                } else {
                                                                                                    v245 = v245
                                                                                                }
                                                                                                _ = len(v1)
                                                                                                v158 = "hello"
                                                                                            default:
                                                                                                v257 := GenS7{F0: (v181 % (((v181 & v181)) | 1))}
                                                                                                _ = v257
                                                                                            }
                                                                                            _ = v245
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v231, v2, v160)
                                                                                        v159 = (v159 << ((18446744073709551615) & 63))
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    v2 += int32(v181)
                                                                                    if v189 {
                                                                                        var v258 float32 = (1.5 + 0.0)
                                                                                        v189 = false
                                                                                        _ = v258
                                                                                    }
                                                                                }
                                                                                switch (fi161 << ((2) & 63)) {
                                                                                case -1:
                                                                                    if (((18446744073709551615 < (v160 & (^(v160)))) || (1 != 4294967295)) && true) {
                                                                                        v259 := GenS9{F0: ((v189 || true) && ((1 > 825) || ((fi161 | 298) != (-(fi162))))), F1: 128}
                                                                                        v1[830] = v231
                                                                                        _ = v259
                                                                                    } else {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            if false {
                                                                                                switch v158 {
                                                                                                case "":
                                                                                                    _ = len(v3)
                                                                                                    if ((42 < 1) || (((((-2.25 / 1.0) - float32((1e10 * float64(3.14159)))) - 0.1) - 3.14159) != 3.14159)) {
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                v260 := GenS2{F0: 9223372036854775807, F1: 142, F2: false}
                                                                                                                _ = v260
                                                                                                            } else {
                                                                                                                v261 := GenS19{F0: ((v158 + (v231 + v231)) + v231)}
                                                                                                                if v189 {
                                                                                                                    if v189 {
                                                                                                                        v2 = int32(fi162)
                                                                                                                        v262 := v1[0]
                                                                                                                        switch fi162 {
                                                                                                                        case 0:
                                                                                                                            fmt.Println(v160)
                                                                                                                        case 454:
                                                                                                                            v160 = 1
                                                                                                                        case 4294967295:
                                                                                                                            v263 := GenS20{F0: 498, F1: 18446744073709551615, F2: int16(fi162)}
                                                                                                                            _ = v263
                                                                                                                        }
                                                                                                                        _ = v262
                                                                                                                    }
                                                                                                                }
                                                                                                                v264 := GenS19{F0: (((v158 + ("hello" + (v231 + v158))) + v231) + v158)}
                                                                                                                _ = v261
                                                                                                                _ = v264
                                                                                                            }
                                                                                                            var v265 uint8 = (v181 % ((v181) | 1))
                                                                                                            var v266 float64 = 0.1
                                                                                                            _ = v265
                                                                                                            _ = v266
                                                                                                        } else {
                                                                                                            v2 = v2
                                                                                                            v189 = ((!true) || v189)
                                                                                                        }
                                                                                                        v4 += v4
                                                                                                        var v267 uint8 = (v181 << ((904) & 63))
                                                                                                        _ = v267
                                                                                                    }
                                                                                                    _ = len(v177)
                                                                                                    v160 += v160
                                                                                                case "foo bar":
                                                                                                    var v268 string = ("\n\t" + ((v158 + v231) + v231))
                                                                                                    _ = v268
                                                                                                case "\n\t":
                                                                                                    v269 := GenS21{F0: "\n\t"}
                                                                                                    _ = v269
                                                                                                default:
                                                                                                    var v270 uint16 = v159
                                                                                                    _ = v270
                                                                                                }
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v271 int64 = 9223372036854775807
                                                                                                    v158 = ("foo bar" + "foo bar")
                                                                                                    v272 := GenS22{F0: genFn1(v271), F1: v2}
                                                                                                    _ = v271
                                                                                                    _ = v272
                                                                                                }
                                                                                                var v273 int16 = 42
                                                                                                if (1 <= fi161) {
                                                                                                    var v274 int16 = v273
                                                                                                    v275 := v3[192]
                                                                                                    _ = v274
                                                                                                    _ = v275
                                                                                                } else {
                                                                                                    var v276 string = v231
                                                                                                    _ = v276
                                                                                                }
                                                                                                _ = v273
                                                                                            }
                                                                                        case "\n\t":
                                                                                            if true {
                                                                                                var v277 float64 = 1e10
                                                                                                v2 -= v2
                                                                                                _ = v277
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 2147483647:
                                                                                    if v189 {
                                                                                        v231 = v231
                                                                                    }
                                                                                case 42:
                                                                                    switch (v158 + v158) {
                                                                                    case "\n\t":
                                                                                        v278 := GenS23{F0: 42}
                                                                                        _ = v278
                                                                                    case "a":
                                                                                        var v279 float32 = 0.1
                                                                                        _ = v279
                                                                                    }
                                                                                default:
                                                                                    v1[uint32(v181)] = "foo bar"
                                                                                }
                                                                                _ = v231
                                                                            case 42:
                                                                                if v189 {
                                                                                    fmt.Println(v189, v189)
                                                                                    if (float64(float32(1e10)) != 0.1) {
                                                                                        v189 = v189
                                                                                    }
                                                                                }
                                                                            default:
                                                                                switch fi161 {
                                                                                case 2147483647:
                                                                                    switch fi162 {
                                                                                    case -9223372036854775808:
                                                                                        if true {
                                                                                            switch v158 {
                                                                                            case "a":
                                                                                                var v280 uint16 = v159
                                                                                                _ = v280
                                                                                            case "foo bar":
                                                                                                _ = len(v1)
                                                                                            case "":
                                                                                                var v281 int64 = (v4 | (v4 & v4))
                                                                                                _ = v281
                                                                                            }
                                                                                            v4 -= (-(v4))
                                                                                        } else {
                                                                                            var v282 uint = uint(v160)
                                                                                            _ = v282
                                                                                        }
                                                                                    case -2147483648:
                                                                                        fmt.Println(fi161, v160, v181)
                                                                                    default:
                                                                                        if len(v177) >= 2 {
                                                                                            v177 = v177[1:len(v177)]
                                                                                        }
                                                                                    }
                                                                                case 32768:
                                                                                    v283 := GenS24{F0: (v160 & v160)}
                                                                                    _ = v283
                                                                                case 128:
                                                                                    var v284 uint16 = (v159 ^ v159)
                                                                                    _ = v284
                                                                                }
                                                                                v3[2147483648] = 42
                                                                                if (!(fi162 <= 300)) {
                                                                                    _ = len(v177)
                                                                                } else {
                                                                                    var v285 uint = 0
                                                                                    v189 = v189
                                                                                    v159 += (^(v159))
                                                                                    v286 := GenS20{F0: 2, F1: v285, F2: -994}
                                                                                    _ = v285
                                                                                    _ = v286
                                                                                }
                                                                                var v287 string = (v158 + ("" + v158))
                                                                                _ = v287
                                                                            }
                                                                            var v288 int64 = 1
                                                                            _ = v288
                                                                        }
                                                                        v160 = (v160 | v160)
                                                                        var v289 string = v158
                                                                        _ = v289
                                                                    }
                                                                default:
                                                                    var v290 uint32 = 2147483647
                                                                    _ = v290
                                                                }
                                                                switch fi161 {
                                                                case 9223372036854775807:
                                                                    if ((((0 < -1) || v189) || true) || v189) {
                                                                        if false {
                                                                            if (v4 == v4) {
                                                                                v291 := GenS25{F0: 0}
                                                                                v292 := GenS26{F0: v158, F1: (v4 * 0)}
                                                                                switch ((v158 + v158) + v158) {
                                                                                case "foo bar":
                                                                                    v4 -= v4
                                                                                case "\n\t":
                                                                                    v160 -= (v160 + (^(v160)))
                                                                                case "hello":
                                                                                    _ = len(v177)
                                                                                }
                                                                                _ = v291
                                                                                _ = v292
                                                                            } else {
                                                                                var v293 int = fi162
                                                                                if len(v177) > 0 {
                                                                                    v177[0] = (v293 % (((v293 >> ((292) & 63))) | 1))
                                                                                }
                                                                                _ = v293
                                                                            }
                                                                            v160 -= v160
                                                                        } else {
                                                                            if false {
                                                                                switch fi162 {
                                                                                case 2147483647:
                                                                                    v294 := v1[66]
                                                                                    _ = v294
                                                                                case 1:
                                                                                    _ = len(v3)
                                                                                default:
                                                                                    switch (fi162 << ((611) & 63)) {
                                                                                    case -378:
                                                                                        fmt.Println(v159, v189)
                                                                                    }
                                                                                    v295 := GenS13{F0: 65535, F1: 42}
                                                                                    fmt.Println(fi161, v4)
                                                                                    _ = v295
                                                                                }
                                                                                _ = len(v177)
                                                                                v296 := v1[2147483648]
                                                                                _ = v296
                                                                            }
                                                                            v297 := GenS8{F0: (v160 ^ (v160 ^ v160)), F1: uint16(v160)}
                                                                            v298 := GenS10{F0: genFn1(v4), F1: 520, F2: 127}
                                                                            _ = v297
                                                                            _ = v298
                                                                        }
                                                                        v299 := v3[42]
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            fmt.Println(v181, v159)
                                                                        }
                                                                        _ = v299
                                                                    } else {
                                                                        var v300 int16 = 0
                                                                        _ = v300
                                                                    }
                                                                case 42:
                                                                    v301 := GenS3{F0: 881, F1: v181}
                                                                    _ = v301
                                                                case -1:
                                                                    v302 := GenS27{F0: -1, F1: 6}
                                                                    _ = v302
                                                                }
                                                                switch "hello" {
                                                                case "hello":
                                                                    if v189 {
                                                                        v189 = v189
                                                                        v158 = ("" + ((v158 + v158) + ("foo bar" + v158)))
                                                                    } else {
                                                                        switch 42 {
                                                                        case 128:
                                                                            fmt.Println(v2, fi161, v189)
                                                                        default:
                                                                            if ((v4 & v4) == v4) {
                                                                                v303 := GenS28{F0: v159}
                                                                                _ = v303
                                                                            }
                                                                        }
                                                                    }
                                                                    if v189 {
                                                                        var v304 int = 472
                                                                        _ = v304
                                                                    } else {
                                                                        v2 += v2
                                                                        v4 += v4
                                                                        switch fi162 {
                                                                        case 1:
                                                                            switch fi162 {
                                                                            case 761:
                                                                                v305 := GenS29{F0: (((1 == 42) && false) && ((v159 ^ (^(v159))) != (v159 << ((2) & 63))))}
                                                                                _ = v305
                                                                            case 9223372036854775807:
                                                                                var v306 uint8 = (v181 >> ((484) & 63))
                                                                                _ = v306
                                                                            default:
                                                                                v307 := GenS25{F0: 785}
                                                                                var v308 int8 = 61
                                                                                _ = v307
                                                                                _ = v308
                                                                            }
                                                                        case -2147483648:
                                                                            v309 := GenS30{F0: v181, F1: v181, F2: (v181 | v181)}
                                                                            _ = v309
                                                                        case -1:
                                                                            v160 = (v160 / (((v160 + uint64(v159))) | 1))
                                                                        default:
                                                                            delete(v3, 738)
                                                                        }
                                                                    }
                                                                    v310 := GenS31{F0: true}
                                                                    _ = v310
                                                                default:
                                                                    if len(v177) > 0 {
                                                                        v177[0] = (fi161 << ((42) & 63))
                                                                    }
                                                                }
                                                                if ((v160 % ((v160) | 1)) >= 19) {
                                                                    var v311 float32 = 3.14159
                                                                    var v312 int16 = 0
                                                                    _ = v311
                                                                    _ = v312
                                                                } else {
                                                                    v181 -= (^(v181))
                                                                    v158 = v158
                                                                }
                                                            } else {
                                                                fmt.Println(v2, v159)
                                                                v313 := GenS1{F0: 128}
                                                                if true {
                                                                    switch fi162 {
                                                                    case 132:
                                                                        var v314 uint64 = v160
                                                                        _ = v314
                                                                    case 4294967295:
                                                                        v158 = (v158 + "")
                                                                    case 0:
                                                                        var v315 int32 = -2147483648
                                                                        _ = v315
                                                                    }
                                                                    v159 += (v159 % (((v159 * v159)) | 1))
                                                                }
                                                                _ = v313
                                                            }
                                                            v4 = 1
                                                        }
                                                    } else {
                                                        v316 := GenS32{F0: 0}
                                                        fmt.Println(fi162)
                                                        _ = v316
                                                    }
                                                }
                                                _ = v189
                                            }
                                        default:
                                            fmt.Println(v4, v4)
                                        }
                                    } else {
                                        var v317 int8 = -94
                                        if false {
                                            v158 = ("a" + v158)
                                            var v318 float32 = (3.14159 - 1e10)
                                            var v319 uint16 = (v159 / (((^(v159))) | 1))
                                            _ = v318
                                            _ = v319
                                        } else {
                                            var v320 uint64 = v160
                                            _ = v320
                                        }
                                        _ = v317
                                    }
                                    _ = v180
                                    _ = v181
                                } else {
                                    switch v158 {
                                    case "a":
                                        var v321 int32 = v2
                                        _ = v321
                                    default:
                                        var v322 int64 = v4
                                        _ = v322
                                    }
                                    v159 = 1
                                    v160 -= 18446744073709551615
                                }
                            }
                        default:
                            var v323 int16 = -1
                            _ = v323
                        }
                    }
                    v1[2] = v158
                    v324 := GenS25{F0: genFn1(-1)}
                    fmt.Println(v2, fi161)
                    _ = v177
                    _ = v324
                }
                if ((v2 & -2147483648) < -1) {
                    var v325 uint8 = 187
                    fmt.Println(v158, v159)
                    v326 := v1[uint32(v159)]
                    _ = v325
                    _ = v326
                }
            }
            _ = fi162
        }
        var v327 uint = 255
        _ = fi161
        _ = v327
    }
    v160 += (v160 + v160)
    var v328 int32 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v328:", v328)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v158
    _ = v159
    _ = v160
    _ = v19
    _ = v2
    _ = v3
    _ = v328
    _ = v4
}
