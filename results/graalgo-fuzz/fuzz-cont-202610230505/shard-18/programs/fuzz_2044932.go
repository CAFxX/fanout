package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 bool; F2 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint16; F2 int }
type GenS7 struct { F0 int; F1 int8; F2 uint }
type GenS8 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 uint8; F1 int8 }
type GenS14 struct { F0 string; F1 uint }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int8; F1 uint32; F2 int8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int8; F1 uint64; F2 int64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint16; F1 int64; F2 int }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 int; F1 string; F2 int }
type GenS23 struct { F0 int64; F1 int; F2 int32 }
type GenS24 struct { F0 string }
type GenS25 struct { F0 int8; F1 string; F2 int8 }
type GenS26 struct { F0 uint64; F1 int64 }
type GenS27 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint32, p1 int32) uint8 {
    fmt.Println(p0, p0)
    var v6 uint32 = p0
    v6 += (^(p0))
    v7 := GenS1{F0: 893, F1: (!true)}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint8 = 1
    var v2 int16 = 32767
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 uint8 = v1
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v2, v5, v2)
        v2 -= (v2 + 364)
        var v9 int = (fi8 / ((int(v1)) | 1))
        _ = fi8
        _ = v9
    }
    v2 -= (-(v2))
    for fi10 := 0; fi10 < 20; fi10++ {
        if false {
            v2 = 42
            switch 323 {
            case -54:
                fmt.Println(v1, v1, v1)
            default:
                fmt.Println(v5, v1)
            }
        } else {
            var v11 uint16 = 2
            var v12 uint32 = 42
            _ = v11
            _ = v12
        }
        var v13 bool = (!(32767 == 32767))
        switch (fi10 | (fi10 % ((fi10) | 1))) {
        case 795:
            if v13 {
                v14 := GenS2{F0: v5, F1: v2}
                _ = v14
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    switch "hello" {
                    case "\n\t":
                        var v16 int16 = v2
                        _ = v16
                    case "":
                        v17 := GenS2{F0: (v1 ^ (v5 ^ (v1 + v5))), F1: v2}
                        _ = v17
                    default:
                        v18 := GenS3{F0: 581}
                        _ = v18
                    }
                    if v13 {
                        if (v1 < genFn0(0, -1)) {
                            v19 := []int{(fi15 & (fi10 + (fi15 / ((127) | 1)))), 0, (fi15 >> ((459) & 63)), (fi15 + fi10)}
                            _ = v19
                        } else {
                            v13 = (false || v13)
                        }
                        v1 = (v1 - 69)
                    } else {
                        v13 = ((!(!true)) && v13)
                        var v20 string = "\n\t"
                        var v21 string = ""
                        if ((0.0 / (1e10 * 1e10)) <= 0.1) {
                            if ((9223372036854775807 == -538) && v13) {
                                switch 128 {
                                case 127:
                                    var v22 uint8 = v1
                                    _ = v22
                                }
                            }
                            v13 = (!((v21 > v20) || v13))
                            v23 := []int{fi10, (fi10 / ((fi15) | 1)), (fi15 - fi15), (fi10 >> ((42) & 63))}
                            if ((-1 != -1) || (199 < (v1 % (((v5 >> ((uint(v1)) & 63))) | 1)))) {
                                var v24 bool = (!false)
                                fmt.Println(v20, fi10, v2)
                                _ = v24
                            } else {
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            }
                            _ = v23
                        } else {
                            var v25 string = v21
                            var v26 uint16 = 32767
                            var v27 uint64 = 42
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = v20
                        _ = v21
                    }
                    _ = fi15
                }
                v3[1] = (100 <= -87)
            }
        case 4294967295:
            var v28 int32 = int32(v2)
            _ = v28
        }
        _ = fi10
        _ = v13
    }
    if false {
        v1 = 127
        for fi29 := 0; fi29 < 1; fi29++ {
            v5 = v1
            var v30 int32 = -1
            _ = fi29
            _ = v30
        }
        var v31 int = 1
        _ = v31
    }
    v32 := GenS1{F0: 1, F1: (!true)}
    var v33 int = 507
    v34 := GenS3{F0: 4294967295}
    v33 = 4294967295
    switch (v33 - v33) {
    case -258:
        switch v33 {
        case -2147483648:
            if true {
                v35 := GenS3{F0: 0}
                _ = v35
            } else {
                for fi36 := 0; fi36 < 16; fi36++ {
                    v3[166] = true
                    fmt.Println(v1, v1, v5)
                    _ = fi36
                }
                fmt.Println(v5, v33, v5)
                var v37 int = (v33 - v33)
                _ = v37
            }
        }
    default:
        for fi38 := 0; fi38 < 14; fi38++ {
            switch fi38 {
            case -526:
                v1 = 1
            case -2147483648:
                v1 = (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))
            case 463:
                fmt.Println(fi38, v33)
            default:
                switch "foo bar" {
                case "foo bar":
                    if (0 < 2147483647) {
                        switch (fi38 ^ -9223372036854775808) {
                        case 537:
                            var v39 uint32 = 4294967295
                            _ = v39
                        default:
                            v1 += (v5 % ((genFn0(1, 1)) | 1))
                        }
                        var v40 uint32 = 817
                        _ = v40
                    } else {
                        var v41 float64 = float64(((float32(1.5) / (3.14159 - 0.0)) + (1.5 / 3.14159)))
                        var v42 bool = (true && (3.14159 == (0.0 / 1.5)))
                        _ = v41
                        _ = v42
                    }
                case "a":
                    var v43 bool = (false || false)
                    _ = v43
                }
            }
            if (("hello" + ("" + "")) == (("\n\t" + "\n\t") + "")) {
                for fi44 := 0; fi44 < 18; fi44++ {
                    var v45 bool = true
                    var v46 string = "\n\t"
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                for fi47 := 0; fi47 < 3; fi47++ {
                    v1 = 0
                    var v48 uint8 = v5
                    switch fi47 {
                    case 32767:
                        var v49 float64 = float64(1e10)
                        _ = v49
                    }
                    if (2147483647 == 423) {
                        var v50 int16 = v2
                        _ = v50
                    } else {
                        v51 := GenS2{F0: 143, F1: -1}
                        fmt.Println(v5)
                        _ = v51
                    }
                    _ = fi47
                    _ = v48
                }
            } else {
                if true {
                    var v52 int32 = 1
                    if false {
                        for fi53 := 0; fi53 < 7; fi53++ {
                            v52 -= v52
                            _ = fi53
                        }
                        var v54 string = "a"
                        _ = v54
                    }
                    v55 := []int{v33, v33, 1}
                    _ = v52
                    _ = v55
                } else {
                    var v56 float32 = (1.5 * (0.0 * float32(((-2.25 + -2.25) / float64(0.1)))))
                    var v57 int8 = 0
                    _ = v56
                    _ = v57
                }
                v58 := GenS3{F0: 2}
                _ = v58
            }
            v59 := GenS1{F0: 0, F1: false}
            _ = fi38
            _ = v59
        }
    }
    var v60 int16 = (v2 ^ (-(v2)))
    if true {
        v1 = 1
        v2 = 1
        fmt.Println(v5, v2)
    }
    if false {
        switch "" {
        case "\n\t":
            var v61 float32 = 0.1
            _ = v61
        default:
            for fi62 := 0; fi62 < 16; fi62++ {
                v63 := []int{(fi62 + (fi62 + fi62)), v33, (^(fi62))}
                var v64 uint64 = 18446744073709551615
                _ = fi62
                _ = v63
                _ = v64
            }
        }
    }
    switch (v33 % ((v33) | 1)) {
    case -408:
        v5 = v5
    case 991:
        v65 := GenS4{F0: -247, F1: true, F2: (v33 / ((-9223372036854775808) | 1))}
        _ = v65
    default:
        v33 -= v33
    }
    fmt.Println(v1)
    v66 := GenS4{F0: 42, F1: (v33 > 0), F2: v33}
    if true {
        for fi67 := 0; fi67 < 4; fi67++ {
            for fi68 := 0; fi68 < 2; fi68++ {
                v69 := GenS5{F0: 42}
                v70 := []int{(v33 - (v33 >> ((255) & 63)))}
                _ = fi68
                _ = v69
                _ = v70
            }
            var v71 uint = uint(v5)
            if (!false) {
                for fi72 := 0; fi72 < 16; fi72++ {
                    if (v2 > 1) {
                        if (-2147483648 != v33) {
                            v71 += 42
                            fmt.Println(v5)
                            v71 -= 899
                            var v73 float32 = -2.25
                            if true {
                                v73 = -2.25
                                if false {
                                    fmt.Println(fi72, v5, v60)
                                    switch ("" + ("\n\t" + "a")) {
                                    case "":
                                        if false {
                                            if (!(true || true)) {
                                                fmt.Println(v60, v73, fi72)
                                            } else {
                                                if ((v71 / ((125) | 1)) > v71) {
                                                    v2 += v60
                                                } else {
                                                    fmt.Println(v73)
                                                    v74 := []int{1, 516}
                                                    _ = v74
                                                }
                                                fmt.Println(fi72, fi72, v73)
                                            }
                                        }
                                    case "foo bar":
                                        v75 := GenS6{F0: v33, F1: 42, F2: fi67}
                                        _ = v75
                                    default:
                                        switch v33 {
                                        case 0:
                                            v33 = -2147483648
                                        case 42:
                                            if false {
                                                if ((float32(((((3.14159 * -2.25) * 1.5) - float64((3.14159 - 0.1))) / float64(float32(1e10)))) + float32(-2.25)) < v73) {
                                                    switch "a" {
                                                    case "\n\t":
                                                        fmt.Println(v2)
                                                    case "":
                                                        v76 := GenS7{F0: fi72, F1: 1, F2: v71}
                                                        _ = v76
                                                    }
                                                }
                                                v77 := GenS6{F0: fi72, F1: 255, F2: v33}
                                                var v78 int64 = -9223372036854775808
                                                _ = v77
                                                _ = v78
                                            }
                                        case 128:
                                            v71 -= (v71 * (v71 & 255))
                                        }
                                        if true {
                                            if (!(v71 < v71)) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v79 uint64 = 0
                                                    _ = v79
                                                case "\n\t":
                                                    fmt.Println(v2, v73, v1)
                                                case "":
                                                    fmt.Println(v2, v2, v1)
                                                }
                                                var v80 string = "a"
                                                switch v33 {
                                                case 2147483647:
                                                    var v81 float64 = (0.0 - ((-2.25 / 1.0) * 0.0))
                                                    _ = v81
                                                case -2147483648:
                                                    if (true && (-128 == 127)) {
                                                        fmt.Println(v2, fi72, v1)
                                                        if (!true) {
                                                            if (!(v71 >= 255)) {
                                                                if true {
                                                                    fmt.Println(v1)
                                                                } else {
                                                                    var v82 uint16 = 32767
                                                                    fmt.Println(v73)
                                                                    _ = v82
                                                                }
                                                                switch v80 {
                                                                case "\n\t":
                                                                    v5 = (v1 - (v1 * (v5 + v5)))
                                                                case "":
                                                                    v71 = 18446744073709551615
                                                                case "foo bar":
                                                                    delete(v3, 65535)
                                                                }
                                                                var v83 int64 = 0
                                                                _ = v83
                                                            }
                                                            if true {
                                                                fmt.Println(v1, fi67, fi72)
                                                                _ = len(v4)
                                                                switch v80 {
                                                                case "foo bar":
                                                                    if true {
                                                                        v84 := GenS1{F0: 18446744073709551615, F1: false}
                                                                        var v85 int8 = -1
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                case "hello":
                                                                    v86 := GenS6{F0: v33, F1: 865, F2: (fi72 / ((fi72) | 1))}
                                                                    _ = v86
                                                                case "a":
                                                                    v1 = 30
                                                                }
                                                            } else {
                                                                v1 += v1
                                                            }
                                                            v1 = v1
                                                        } else {
                                                            var v87 int64 = -9223372036854775808
                                                            v87 -= (v87 | v87)
                                                            v88 := v4[fi72]
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        switch "" {
                                                        case "":
                                                            fmt.Println(v1, v1)
                                                        case "foo bar":
                                                            v33 += v33
                                                        case "\n\t":
                                                            delete(v3, 0)
                                                        default:
                                                            if ((-42 < 1) && false) {
                                                                var v89 uint32 = 2147483647
                                                                var v90 int16 = (v2 << (((v71 & v71)) & 63))
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                if (false && false) {
                                                                    v5 = v5
                                                                    var v91 uint32 = 2147483647
                                                                    _ = v91
                                                                } else {
                                                                    v92 := GenS8{F0: false, F1: 0, F2: 65535}
                                                                    _ = v92
                                                                }
                                                                var v93 int = 4294967295
                                                                v94 := GenS9{F0: v71}
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                        }
                                                    } else {
                                                        v80 = "\n\t"
                                                        v60 = (v2 ^ v60)
                                                        var v95 uint8 = v5
                                                        _ = v95
                                                    }
                                                }
                                                _ = v80
                                            } else {
                                                fmt.Println(v33, v60, v71)
                                            }
                                            var v96 uint64 = 913
                                            _ = v96
                                        } else {
                                            fmt.Println(v60, v71, v5)
                                            if (((v1 <= (^(v1))) && (v60 <= (v60 + 0))) && (18446744073709551615 >= 18446744073709551615)) {
                                                var v97 uint64 = 2
                                                _ = v97
                                            } else {
                                                v98 := []int{}
                                                if false {
                                                    if false {
                                                        var v99 int = v33
                                                        v100 := GenS2{F0: (v5 ^ (v5 + (v1 << ((2) & 63)))), F1: v60}
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        v60 -= 918
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        switch 9223372036854775807 {
                                                        case -1:
                                                            v33 -= fi67
                                                        case 127:
                                                            v101 := GenS4{F0: -1, F1: false, F2: 42}
                                                            _ = v101
                                                        case 638:
                                                            v33 -= fi67
                                                        }
                                                        if false {
                                                            var v102 bool = false
                                                            if len(v98) > 0 {
                                                                v98[0] = (fi67 - v33)
                                                            }
                                                            _ = v102
                                                        }
                                                        v103 := GenS7{F0: -678, F1: 0, F2: v71}
                                                        _ = v103
                                                    case "":
                                                        v2 = (v2 * 1000)
                                                    case "hello":
                                                        var v104 float32 = ((v73 * (v73 - (v73 - 0.1))) / v73)
                                                        _ = v104
                                                    }
                                                    if (385 > 555) {
                                                        switch (fi72 - fi72) {
                                                        case 490:
                                                            switch "a" {
                                                            case "\n\t":
                                                                if false {
                                                                    v105 := GenS9{F0: (v71 % ((0) | 1))}
                                                                    switch v33 {
                                                                    case 2147483647:
                                                                        var v106 uint8 = v1
                                                                        _ = v106
                                                                    case 860:
                                                                        v4[v33] = 203
                                                                    case 9223372036854775807:
                                                                        var v107 int64 = -1
                                                                        _ = v107
                                                                    default:
                                                                        var v108 int64 = -745
                                                                        _ = v108
                                                                    }
                                                                    if (true || false) {
                                                                        if false {
                                                                            if (-1 == 0) {
                                                                                if false {
                                                                                    if ((9223372036854775807 >= 1) && (true && ((true || (true && true)) || (42 > 1)))) {
                                                                                        if len(v98) > 0 {
                                                                                            v98[0] = (fi72 / ((v33) | 1))
                                                                                        }
                                                                                        var v109 int64 = 0
                                                                                        if (true && false) {
                                                                                            switch -442 {
                                                                                            case 1:
                                                                                                switch v33 {
                                                                                                case -420:
                                                                                                    var v110 uint16 = 2
                                                                                                    _ = v110
                                                                                                case 4294967295:
                                                                                                    var v111 string = ("a" + "\n\t")
                                                                                                    _ = v111
                                                                                                case -2147483648:
                                                                                                    if (float64(v73) >= ((0.0 * float64(3.14159)) - (3.14159 * (3.14159 - (1.5 + (0.0 / ((3.14159 / float64(-2.25)) + (1.5 / (0.1 / 1.0))))))))) {
                                                                                                        var v112 float32 = 0.0
                                                                                                        switch ("a" + "hello") {
                                                                                                        case "foo bar":
                                                                                                            switch "foo bar" {
                                                                                                            case "hello":
                                                                                                                var v113 uint32 = 1
                                                                                                                _ = v113
                                                                                                            case "a":
                                                                                                                v114 := GenS10{F0: 2147483647}
                                                                                                                _ = v114
                                                                                                            case "":
                                                                                                                if (!(("hello" + "foo bar") != "hello")) {
                                                                                                                    v33 += 128
                                                                                                                    v115 := GenS8{F0: (false && false), F1: 1, F2: 65535}
                                                                                                                    _ = v115
                                                                                                                }
                                                                                                            default:
                                                                                                                v112 = (0.1 + v112)
                                                                                                            }
                                                                                                            var v116 float64 = 0.1
                                                                                                            _ = len(v4)
                                                                                                            v117 := GenS11{F0: 252}
                                                                                                            _ = v116
                                                                                                            _ = v117
                                                                                                        default:
                                                                                                            fmt.Println(v1, v112)
                                                                                                        }
                                                                                                        _ = v112
                                                                                                    }
                                                                                                    v118 := GenS1{F0: 172, F1: (366 != 567)}
                                                                                                    _ = v118
                                                                                                default:
                                                                                                    var v119 uint64 = 2
                                                                                                    _ = v119
                                                                                                }
                                                                                            case 0:
                                                                                                if (uint32(v5) < 324) {
                                                                                                    var v120 float64 = 1.5
                                                                                                    switch fi67 {
                                                                                                    case -1:
                                                                                                        fmt.Println(v120)
                                                                                                    case -858:
                                                                                                        var v121 uint8 = v5
                                                                                                        _ = v121
                                                                                                    case 592:
                                                                                                        switch (^(fi67)) {
                                                                                                        case 128:
                                                                                                            v2 -= 1000
                                                                                                        case 32768:
                                                                                                            var v122 int = (fi67 << ((uint(v1)) & 63))
                                                                                                            _ = v122
                                                                                                        }
                                                                                                    default:
                                                                                                        var v123 bool = (v1 > v1)
                                                                                                        _ = v123
                                                                                                    }
                                                                                                    v120 = 0.0
                                                                                                    _ = v120
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v124 float64 = 0.0
                                                                                            _ = v124
                                                                                        }
                                                                                        _ = v109
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            v1 = genFn0(671, 1)
                                                                                        }
                                                                                        fmt.Println(fi72, v60, v60)
                                                                                    }
                                                                                    v1 = 1
                                                                                    if len(v98) > 0 {
                                                                                        v98[0] = (fi67 | fi72)
                                                                                    }
                                                                                    var v125 int64 = 45
                                                                                    _ = v125
                                                                                } else {
                                                                                    var v126 int = 32767
                                                                                    _ = v126
                                                                                }
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v71 += (v71 << ((102) & 63))
                                                                                case "":
                                                                                    var v127 float64 = (float64((v73 + (v73 + v73))) - ((3.14159 - (1e10 - (0.1 + 1e10))) * 0.1))
                                                                                    _ = v127
                                                                                case "foo bar":
                                                                                    v33 = int(v1)
                                                                                default:
                                                                                    var v128 uint32 = 0
                                                                                    _ = v128
                                                                                }
                                                                            } else {
                                                                                v2 = (v2 >> ((v71) & 63))
                                                                                switch fi67 {
                                                                                case 128:
                                                                                    var v129 int16 = 42
                                                                                    _ = v129
                                                                                case 2147483647:
                                                                                    var v130 int32 = 42
                                                                                    _ = v130
                                                                                }
                                                                                var v131 int16 = (v60 / ((v60) | 1))
                                                                                _ = v131
                                                                            }
                                                                            var v132 uint32 = 12
                                                                            _ = v132
                                                                        } else {
                                                                            v71 -= (^(v71))
                                                                            v133 := GenS8{F0: true, F1: 23, F2: 42}
                                                                            fmt.Println(v2, v2)
                                                                            _ = v133
                                                                        }
                                                                        switch (fi72 << (((v71 ^ v71)) & 63)) {
                                                                        case 4294967295:
                                                                            fmt.Println(v5)
                                                                        case -807:
                                                                            if ((1.5 * ((3.14159 + 1.5) + float32((0.0 - 0.0)))) < 3.14159) {
                                                                                v1 -= (v1 / ((v1) | 1))
                                                                                v134 := GenS8{F0: (-166 >= int32(v71)), F1: 2, F2: 42}
                                                                                if (((-2.25 - 0.1) != float32(3.14159)) || true) {
                                                                                    v73 += v73
                                                                                } else {
                                                                                    v135 := GenS2{F0: v5, F1: v60}
                                                                                    _ = v135
                                                                                }
                                                                                _ = v134
                                                                            }
                                                                            switch 0 {
                                                                            case 4294967295:
                                                                                v5 = (v5 & v5)
                                                                            case 1:
                                                                                v3[65535] = (v71 != (v71 & v71))
                                                                            case 0:
                                                                                switch (fi72 << ((v71) & 63)) {
                                                                                case 128:
                                                                                    if len(v98) > 0 {
                                                                                        v98[0] = -1
                                                                                    }
                                                                                case 117:
                                                                                    fmt.Println(v2, v2)
                                                                                }
                                                                            default:
                                                                                var v136 float32 = -2.25
                                                                                delete(v4, fi67)
                                                                                if true {
                                                                                    v60 = v2
                                                                                    v60 = (v60 / ((v60) | 1))
                                                                                    v1 += 127
                                                                                } else {
                                                                                    var v137 int32 = 42
                                                                                    switch ("foo bar" + (("foo bar" + (("hello" + ("hello" + ("" + ("a" + ("\n\t" + ""))))) + "hello")) + ("foo bar" + "foo bar"))) {
                                                                                    case "":
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            var v138 uint = 919
                                                                                            _ = v138
                                                                                        case "":
                                                                                            _ = len(v4)
                                                                                            v73 -= 0.0
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v2, fi72, v73)
                                                                                            case "":
                                                                                                var v139 int = (fi67 - fi67)
                                                                                                if true {
                                                                                                    var v140 float64 = (1.5 * (-2.25 + 0.0))
                                                                                                    _ = v140
                                                                                                }
                                                                                                v139 -= (-(v139))
                                                                                                _ = v139
                                                                                            default:
                                                                                                _ = len(v4)
                                                                                            }
                                                                                            v136 = (float32((0.0 - 0.0)) - (v136 - (v136 / (v136 / ((-2.25 / (v136 / v136)) * (v73 * v73))))))
                                                                                            var v141 int = v33
                                                                                            if false {
                                                                                                if len(v98) > 0 {
                                                                                                    v98[0] = v141
                                                                                                }
                                                                                            } else {
                                                                                                switch 2147483647 {
                                                                                                case 311:
                                                                                                    var v142 float32 = (v73 - v73)
                                                                                                    _ = v142
                                                                                                case -2147483648:
                                                                                                    var v143 int32 = (^(v137))
                                                                                                    _ = v143
                                                                                                case 32768:
                                                                                                    v71 = (v71 ^ (v71 - 2))
                                                                                                }
                                                                                                if len(v98) > 0 {
                                                                                                    v98[0] = v33
                                                                                                }
                                                                                            }
                                                                                            _ = v141
                                                                                        case "hello":
                                                                                            delete(v3, 390)
                                                                                        default:
                                                                                            v144 := GenS2{F0: genFn0(65535, v137), F1: (v60 - v2)}
                                                                                            _ = v144
                                                                                        }
                                                                                    case "a":
                                                                                        var v145 int64 = 0
                                                                                        _ = v145
                                                                                    default:
                                                                                        var v146 uint32 = 360
                                                                                        _ = v146
                                                                                    }
                                                                                    _ = len(v3)
                                                                                    _ = v137
                                                                                }
                                                                                _ = v136
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v147 := GenS11{F0: 42}
                                                                        _ = v147
                                                                    }
                                                                    _ = v105
                                                                } else {
                                                                    v1 += v1
                                                                    v5 = v1
                                                                    switch "" {
                                                                    case "hello":
                                                                        if false {
                                                                            v148 := GenS5{F0: 1}
                                                                            fmt.Println(v73, v73, v5)
                                                                            v149 := GenS11{F0: -339}
                                                                            _ = v148
                                                                            _ = v149
                                                                        }
                                                                    case "foo bar":
                                                                        if (-427 > (v33 & fi67)) {
                                                                            if ("" <= ("\n\t" + "hello")) {
                                                                                if (false || (-2.25 != -2.25)) {
                                                                                    v5 = (v5 << ((18446744073709551615) & 63))
                                                                                }
                                                                                var v150 uint8 = genFn0(uint32(v71), 1)
                                                                                _ = v150
                                                                            } else {
                                                                                if (true || (true && (true || (835 != 0)))) {
                                                                                    v151 := GenS1{F0: 2, F1: (-1 > int(v60))}
                                                                                    _ = v151
                                                                                } else {
                                                                                    var v152 int16 = v60
                                                                                    if (!("\n\t" >= "")) {
                                                                                        var v153 float32 = (v73 * (v73 - 1e10))
                                                                                        v60 = v2
                                                                                        v154 := GenS12{F0: 65535, F1: 48}
                                                                                        var v155 uint32 = 2147483647
                                                                                        _ = v153
                                                                                        _ = v154
                                                                                        _ = v155
                                                                                    }
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        var v156 string = "a"
                                                                                        _ = v156
                                                                                    case "":
                                                                                        var v157 bool = false
                                                                                        _ = v157
                                                                                    default:
                                                                                        v33 += (fi67 << (((v71 - (v71 + (v71 & v71)))) & 63))
                                                                                    }
                                                                                    v158 := GenS13{F0: v1, F1: int8(fi67)}
                                                                                    _ = v152
                                                                                    _ = v158
                                                                                }
                                                                                v159 := GenS14{F0: "foo bar", F1: (v71 / ((v71) | 1))}
                                                                                _ = v159
                                                                            }
                                                                            fmt.Println(v60, v71)
                                                                        } else {
                                                                            if len(v98) >= 2 {
                                                                                v98 = v98[1:len(v98)]
                                                                            }
                                                                            v160 := GenS13{F0: (v5 * v1), F1: 100}
                                                                            _ = v160
                                                                        }
                                                                    case "":
                                                                        _ = len(v3)
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v161 := GenS8{F0: (778 > v2), F1: 65535, F2: 250}
                                                                _ = v161
                                                            case "a":
                                                                v162 := GenS15{F0: 2147483647}
                                                                _ = v162
                                                            }
                                                        case 404:
                                                            fmt.Println(v1, v33, v5)
                                                        default:
                                                            if (((true && ((v33 | fi67) <= fi72)) && false) || false) {
                                                                v163 := GenS4{F0: -1, F1: false, F2: -9223372036854775808}
                                                                if (!(!(((1e10 + (float64((float32(1.5) + 1.5)) + 0.0)) + (-2.25 / 3.14159)) <= 3.14159))) {
                                                                    v2 = v2
                                                                    v164 := v3[42]
                                                                    _ = v164
                                                                }
                                                                v98 = append(v98, 740)
                                                                _ = v163
                                                            }
                                                        }
                                                        v2 -= v2
                                                        v73 -= float32((1.5 * (-2.25 * 0.1)))
                                                    } else {
                                                        var v165 int = fi67
                                                        switch fi67 {
                                                        case 4294967295:
                                                            if false {
                                                                v1 += v5
                                                            }
                                                        case 32767:
                                                            _ = len(v98)
                                                        case 128:
                                                            if false {
                                                                v98 = append(v98, -1)
                                                                var v166 uint = v71
                                                                _ = v166
                                                            }
                                                        default:
                                                            v167 := GenS3{F0: 967}
                                                            _ = v167
                                                        }
                                                        fmt.Println(fi72, fi72)
                                                        _ = v165
                                                    }
                                                }
                                                v73 = (0.0 * v73)
                                                _ = v98
                                            }
                                            fmt.Println(fi67, v73)
                                        }
                                        fmt.Println(v2)
                                        v168 := GenS4{F0: -2147483648, F1: true, F2: int(v5)}
                                        _ = v168
                                    }
                                } else {
                                    v169 := GenS12{F0: v71, F1: 2147483648}
                                    _ = len(v4)
                                    _ = v169
                                }
                                var v170 float32 = 1.5
                                _ = v170
                            }
                            _ = v73
                        }
                        v171 := v3[2147483647]
                        if v171 {
                            v172 := GenS13{F0: (v5 % ((genFn0(2, 345)) | 1)), F1: 68}
                            if v171 {
                                if v171 {
                                    fmt.Println(v71, v60)
                                }
                            } else {
                                v173 := []int{(fi72 << (((v71 * (v71 / ((v71) | 1)))) & 63)), fi67, 2147483647}
                                _ = v173
                            }
                            var v174 float64 = float64((-2.25 - 0.1))
                            _ = v172
                            _ = v174
                        }
                        _ = v171
                    } else {
                        var v175 int64 = -649
                        v176 := []int{-9223372036854775808}
                        _ = v175
                        _ = v176
                    }
                    _ = fi72
                }
                switch (-(v33)) {
                case 32768:
                    switch "" {
                    case "\n\t":
                        switch (v33 - (fi67 % (((^(v33))) | 1))) {
                        case 127:
                            if false {
                                v177 := GenS11{F0: 1}
                                _ = v177
                            } else {
                                if (genFn0(4294967295, 1) >= 208) {
                                    v178 := GenS10{F0: 65535}
                                    var v179 float32 = ((float32((-2.25 / 3.14159)) / 1e10) + -2.25)
                                    _ = v178
                                    _ = v179
                                }
                                v1 = v5
                            }
                        case 430:
                            for fi180 := 0; fi180 < 2; fi180++ {
                                var v181 int64 = -1
                                switch ("" + ("a" + ("\n\t" + "a"))) {
                                case "foo bar":
                                    switch fi180 {
                                    case -764:
                                        var v182 int = fi180
                                        _ = v182
                                    case -1:
                                        v183 := GenS16{F0: 0, F1: 476, F2: -128}
                                        _ = v183
                                    default:
                                        v184 := GenS6{F0: v33, F1: uint16(v71), F2: (fi180 << (((^(v71))) & 63))}
                                        if true {
                                            fmt.Println(v1, v5)
                                            v185 := GenS15{F0: 65535}
                                            v186 := GenS17{F0: 42}
                                            _ = v185
                                            _ = v186
                                        } else {
                                            v181 += v181
                                            if true {
                                                v60 += -817
                                                v5 = (v1 - (v1 ^ 15))
                                            } else {
                                                v187 := []int{-749, v33}
                                                if false {
                                                    var v188 int16 = (-(v60))
                                                    if (42 != 548) {
                                                        fmt.Println(fi67, v60, v60)
                                                        var v189 uint = (v71 ^ 675)
                                                        v187 = append(v187, fi67)
                                                        if (true && (false && false)) {
                                                            switch ("" + "") {
                                                            case "":
                                                                if (4294967295 > 156) {
                                                                    v190 := GenS18{F0: -17, F1: 0, F2: v181}
                                                                    v191 := GenS19{F0: "hello"}
                                                                    v192 := v4[(v33 % (((fi180 | int(v181))) | 1))]
                                                                    var v193 uint32 = v192
                                                                    _ = v190
                                                                    _ = v191
                                                                    _ = v192
                                                                    _ = v193
                                                                } else {
                                                                    v194 := GenS8{F0: true, F1: 1, F2: uint16(v1)}
                                                                    v195 := GenS20{F0: 65535, F1: v181, F2: v33}
                                                                    v5 -= 0
                                                                    _ = v194
                                                                    _ = v195
                                                                }
                                                                v3[1] = false
                                                                if (true && false) {
                                                                    v196 := v4[(fi180 / (((fi180 / (((fi67 ^ fi180)) | 1))) | 1))]
                                                                    if len(v187) >= 2 {
                                                                        v187 = v187[1:len(v187)]
                                                                    }
                                                                    delete(v4, (^(fi180)))
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v3[(^(v196))] = (823 != 0)
                                                                    default:
                                                                        if (v33 < (fi67 & (fi67 << (((^(v189))) & 63)))) {
                                                                            if true {
                                                                                _ = len(v4)
                                                                            } else {
                                                                                v196 = (v196 >> ((v71) & 63))
                                                                                switch ("" + "hello") {
                                                                                case "hello":
                                                                                    v197 := GenS5{F0: 42}
                                                                                    _ = v197
                                                                                case "\n\t":
                                                                                    if len(v187) > 0 {
                                                                                        v187[0] = (fi67 | (fi180 / (((v33 >> (((^(v71))) & 63))) | 1)))
                                                                                    }
                                                                                default:
                                                                                    if (((1 <= 373) && (0.0 > 0.1)) || (float32(1.5) < 3.14159)) {
                                                                                        v188 = v60
                                                                                        var v198 int = (fi67 | -9223372036854775808)
                                                                                        if false {
                                                                                            v199 := GenS16{F0: 100, F1: 1, F2: 42}
                                                                                            var v200 uint32 = (v196 | 65535)
                                                                                            _ = v199
                                                                                            _ = v200
                                                                                        }
                                                                                        if (v1 <= (v1 * (^(v1)))) {
                                                                                            v198 -= (-(fi180))
                                                                                            delete(v4, (v198 - (^(fi67))))
                                                                                        } else {
                                                                                            switch 127 {
                                                                                            case 85:
                                                                                                if ((v71 < v189) && (!false)) {
                                                                                                    fmt.Println(v1)
                                                                                                    v201 := v4[fi67]
                                                                                                    _ = v201
                                                                                                } else {
                                                                                                    v202 := GenS21{F0: 18446744073709551615}
                                                                                                    _ = v202
                                                                                                }
                                                                                            case -1:
                                                                                                var v203 int8 = 42
                                                                                                _ = v203
                                                                                            default:
                                                                                                fmt.Println(v60)
                                                                                            }
                                                                                        }
                                                                                        var v204 string = "\n\t"
                                                                                        _ = v198
                                                                                        _ = v204
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v205 float64 = (0.0 + -2.25)
                                                                            v188 -= v2
                                                                            _ = v205
                                                                        }
                                                                    }
                                                                    v5 += (^(v1))
                                                                    _ = len(v187)
                                                                    _ = v196
                                                                }
                                                            case "\n\t":
                                                                if (100 != -1) {
                                                                    v4[fi180] = 289
                                                                } else {
                                                                    if (false || (-2.25 <= (1e10 - -2.25))) {
                                                                        switch fi67 {
                                                                        case 42:
                                                                            v4[127] = 2147483647
                                                                        case 333:
                                                                            var v206 int32 = 1
                                                                            if false {
                                                                                if (false && (!(42 <= 84))) {
                                                                                    if (!true) {
                                                                                        if false {
                                                                                            var v207 float64 = (float64((((((float32(((3.14159 + 0.0) / (3.14159 + (0.1 - (float64(0.1) * 0.1))))) - 0.0) * (1.5 * (-2.25 - float32((3.14159 + 0.1))))) / (float32(3.14159) / 0.1)) + 0.1) - float32(1.5))) / 1.0)
                                                                                            _ = v207
                                                                                        } else {
                                                                                            v2 -= (-(v2))
                                                                                        }
                                                                                    } else {
                                                                                        v208 := GenS21{F0: 2}
                                                                                        _ = v208
                                                                                    }
                                                                                    v60 = -1
                                                                                    if false {
                                                                                        var v209 uint = 1
                                                                                        if ("\n\t" <= "hello") {
                                                                                            var v210 uint32 = uint32(v1)
                                                                                            v211 := GenS17{F0: 0}
                                                                                            v4[129] = 65535
                                                                                            _ = v210
                                                                                            _ = v211
                                                                                        }
                                                                                        _ = v209
                                                                                    } else {
                                                                                        v33 = (-(fi180))
                                                                                    }
                                                                                }
                                                                            }
                                                                            v188 -= v2
                                                                            _ = v206
                                                                        default:
                                                                            v60 = (v2 / ((v60) | 1))
                                                                        }
                                                                    } else {
                                                                        v2 = v60
                                                                        var v212 bool = false
                                                                        switch ("a" + "\n\t") {
                                                                        case "foo bar":
                                                                            var v213 float32 = ((0.1 - (0.0 - (1.5 * float32(((3.14159 - 1e10) - 0.1))))) / -2.25)
                                                                            _ = v213
                                                                        case "a":
                                                                            v33 -= 42
                                                                        case "":
                                                                            v188 -= (v2 * v60)
                                                                        default:
                                                                            fmt.Println(v1)
                                                                        }
                                                                        v214 := GenS22{F0: 127, F1: "\n\t", F2: int(v189)}
                                                                        var v215 uint64 = 42
                                                                        var v216 int = (^(v33))
                                                                        _ = v212
                                                                        _ = v214
                                                                        _ = v215
                                                                        _ = v216
                                                                    }
                                                                    _ = len(v187)
                                                                }
                                                            case "foo bar":
                                                                var v217 uint32 = 2
                                                                _ = v217
                                                            }
                                                        }
                                                        _ = v189
                                                    }
                                                    var v218 int64 = (v181 - (^(v181)))
                                                    v218 = v181
                                                    _ = v188
                                                    _ = v218
                                                } else {
                                                    v2 -= -855
                                                    var v219 string = "a"
                                                    switch fi180 {
                                                    case 127:
                                                        v219 = "\n\t"
                                                    case 1:
                                                        if ((v2 / ((v2) | 1)) != v2) {
                                                            v60 -= 0
                                                            var v220 float64 = (float64((0.0 + 3.14159)) + 0.0)
                                                            switch -2147483648 {
                                                            case -112:
                                                                if false {
                                                                    v220 += (float64(0.1) * float64((1e10 + 0.0)))
                                                                    var v221 string = "a"
                                                                    fmt.Println(v220)
                                                                    var v222 uint8 = (v5 | v1)
                                                                    _ = v221
                                                                    _ = v222
                                                                }
                                                            case 0:
                                                                var v223 uint32 = 831
                                                                _ = v223
                                                            }
                                                            _ = len(v3)
                                                            _ = len(v187)
                                                            v224 := GenS23{F0: v181, F1: fi67, F2: 1}
                                                            _ = v220
                                                            _ = v224
                                                        }
                                                    case 385:
                                                        v225 := GenS24{F0: ""}
                                                        _ = v225
                                                    default:
                                                        if false {
                                                            var v226 uint8 = genFn0(4294967295, 42)
                                                            if len(v187) > 0 {
                                                                v187[0] = int(v2)
                                                            }
                                                            v1 = (v226 << ((255) & 63))
                                                            var v227 int8 = int8(v71)
                                                            _ = len(v3)
                                                            _ = v226
                                                            _ = v227
                                                        } else {
                                                            v228 := GenS25{F0: -128, F1: v219, F2: 127}
                                                            v229 := GenS22{F0: v33, F1: v219, F2: (fi180 | v33)}
                                                            v230 := GenS25{F0: -1, F1: v219, F2: 1}
                                                            _ = v228
                                                            _ = v229
                                                            _ = v230
                                                        }
                                                    }
                                                    _ = v219
                                                }
                                                var v231 uint8 = (^(v5))
                                                _ = v187
                                                _ = v231
                                            }
                                        }
                                        _ = v184
                                    }
                                default:
                                    v60 += 1
                                }
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v232 := []int{2147483647}
                                        _ = v232
                                    case "hello":
                                        var v233 bool = ((v181 >> (((v71 & (v71 * 913))) & 63)) > (v181 + v181))
                                        _ = v233
                                    }
                                    v33 = int(v181)
                                    var v234 int64 = (v181 % (((v181 * -222)) | 1))
                                    _ = v234
                                }
                                _ = fi180
                                _ = v181
                            }
                        default:
                            for fi235 := 0; fi235 < 7; fi235++ {
                                v236 := GenS26{F0: 4294967295, F1: -9223372036854775808}
                                v60 = 1000
                                fmt.Println(v33, v60)
                                _ = fi235
                                _ = v236
                            }
                        }
                    case "hello":
                        var v237 bool = false
                        _ = v237
                    default:
                        var v238 int16 = v2
                        _ = v238
                    }
                case -104:
                    if false {
                        if false {
                            var v239 uint16 = 2
                            if true {
                                var v240 uint = v71
                                var v241 int8 = -44
                                v33 += (v33 ^ (fi67 | (fi67 - (v33 / (((fi67 << ((18446744073709551615) & 63))) | 1)))))
                                _ = v240
                                _ = v241
                            } else {
                                v242 := []int{fi67, fi67}
                                _ = v242
                            }
                            delete(v4, fi67)
                            _ = v239
                        }
                    }
                case -9223372036854775808:
                    v1 = genFn0(2147483647, 0)
                default:
                    switch (v33 << (((v71 / ((v71) | 1))) & 63)) {
                    case -20:
                        for fi243 := 0; fi243 < 5; fi243++ {
                            v71 -= v71
                            _ = fi243
                        }
                    case -2147483648:
                        switch v33 {
                        case 1:
                            v2 = 32767
                        case 32767:
                            if ((!(127 <= -29)) || (0 == 536)) {
                                fmt.Println(v1)
                            }
                        default:
                            fmt.Println(v1, v1)
                        }
                    case 2147483647:
                        v244 := GenS6{F0: fi67, F1: 32767, F2: (v33 >> ((18446744073709551615) & 63))}
                        _ = v244
                    }
                }
                v2 = v2
            }
            _ = fi67
            _ = v71
        }
        v245 := GenS27{F0: 1, F1: 42}
        _ = v245
    } else {
        var v246 int32 = 2147483647
        v247 := GenS16{F0: -11, F1: 151, F2: 1}
        for fi248 := 0; fi248 < 18; fi248++ {
            v3[uint32(v5)] = false
            v60 = v60
            for fi249 := 0; fi249 < 16; fi249++ {
                if false {
                    var v250 int32 = v246
                    _ = v250
                }
                v246 = v246
                _ = fi249
            }
            _ = fi248
        }
        _ = v246
        _ = v247
    }
    var v251 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v251:", v251)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v251
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v60
    _ = v66
}
