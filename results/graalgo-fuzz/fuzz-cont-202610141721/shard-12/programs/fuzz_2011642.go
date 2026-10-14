package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int32; F1 int32; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 int64; F2 uint }
type GenS9 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS10 struct { F0 bool; F1 uint64; F2 bool }
type GenS11 struct { F0 int; F1 int16; F2 string }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int; F1 int16; F2 int16 }
type GenS14 struct { F0 uint16; F1 uint; F2 int32 }
type GenS15 struct { F0 uint32; F1 bool }
type GenS16 struct { F0 int8; F1 int8 }
type GenS17 struct { F0 int32; F1 int16 }
func genFn0(p0 uint, p1 uint16) int64 {
    p1 -= (p1 + p1)
    var v6 float64 = (float64(-2.25) + -2.25)
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 float32 = (1.5 * 1.5)
    var v2 int16 = -32768
    var v3 float32 = v1
    var v4 string = "hello"
    v5 := make(map[int]int)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{(fi7 * fi7)}
        _ = fi7
        _ = v8
    }
    v9 := []int{4294967295}
    for fi10 := 0; fi10 < 1; fi10++ {
        switch fi10 {
        case 127:
            v1 -= v1
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (true || true) {
            v12 := GenS1{F0: (true || true), F1: 874, F2: -1}
            if true {
                var v13 int16 = (v2 & v2)
                v14 := GenS1{F0: (((float64((1.5 + v3)) + 3.14159) <= (0.0 / 1e10)) || false), F1: 784, F2: -2147483648}
                switch "foo bar" {
                case "a":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        fmt.Println(fi11, fi15)
                        var v16 int16 = v13
                        switch "\n\t" {
                        case "a":
                            var v17 float64 = ((1e10 / ((float64(0.1) / 1.0) + 3.14159)) + -2.25)
                            _ = v17
                        case "hello":
                            var v18 string = (v4 + v4)
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                default:
                    fmt.Println(v1, v3, fi11)
                }
                _ = v13
                _ = v14
            } else {
                if (!(!false)) {
                    switch ("a" + "foo bar") {
                    case "":
                        var v19 uint = 1
                        _ = v19
                    default:
                        var v20 float64 = (1.5 * float64(0.0))
                        _ = v20
                    }
                } else {
                    if (-1 <= -1) {
                        if ((v1 - 0.1) != (3.14159 + v3)) {
                            fmt.Println(v2)
                            v21 := GenS2{F0: 2}
                            for fi22 := 0; fi22 < 14; fi22++ {
                                v23 := GenS2{F0: 37}
                                _ = fi22
                                _ = v23
                            }
                            _ = v21
                        }
                    }
                }
                switch v4 {
                case "":
                    for fi24 := 0; fi24 < 9; fi24++ {
                        var v25 int32 = -1
                        _ = fi24
                        _ = v25
                    }
                default:
                    if true {
                        v9 = append(v9, (fi11 | fi11))
                    } else {
                        switch fi11 {
                        case 127:
                            for fi26 := 0; fi26 < 13; fi26++ {
                                if true {
                                    v1 = (v3 * ((v3 * v3) + v3))
                                    v27 := GenS1{F0: true, F1: 621, F2: 42}
                                    _ = v27
                                }
                                if false {
                                    v1 += -2.25
                                    v28 := GenS3{F0: 42, F1: false, F2: ((v4 + (v4 + v4)) + v4)}
                                    _ = v28
                                }
                                fmt.Println(v1, fi11)
                                _ = fi26
                            }
                        default:
                            for fi29 := 0; fi29 < 17; fi29++ {
                                delete(v5, fi29)
                                if (65535 >= 32767) {
                                    v5[(fi29 >> ((42) & 63))] = 9223372036854775807
                                    _ = len(v9)
                                }
                                _ = fi29
                            }
                        }
                        if (1.5 > 1.5) {
                            var v30 int64 = 0
                            v30 = (v30 << ((1) & 63))
                            if false {
                                if (1 < 1) {
                                    v31 := GenS3{F0: 42, F1: false, F2: (v4 + ((v4 + "hello") + v4))}
                                    _ = v31
                                }
                                v30 = v30
                            } else {
                                var v32 uint16 = 255
                                if (fi11 != 128) {
                                    if ((fi11 == fi11) || (!true)) {
                                        v30 += v30
                                    } else {
                                        if true {
                                            var v33 int32 = 42
                                            _ = v33
                                        }
                                        switch v4 {
                                        case "hello":
                                            var v34 float64 = (((3.14159 * 0.0) * float64(v3)) - ((3.14159 - -2.25) - 1.5))
                                            _ = v34
                                        case "":
                                            v35 := GenS2{F0: 0}
                                            _ = v35
                                        case "a":
                                            if false {
                                                for fi36 := 0; fi36 < 2; fi36++ {
                                                    if false {
                                                        switch "" {
                                                        case "\n\t":
                                                            v37 := GenS2{F0: 42}
                                                            _ = v37
                                                        }
                                                        if (1 >= 42) {
                                                            if true {
                                                                var v38 int16 = 1
                                                                v2 -= v38
                                                                _ = v38
                                                            } else {
                                                                if false {
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                    v39 := GenS3{F0: 65535, F1: true, F2: v4}
                                                                    _ = v39
                                                                }
                                                                v1 = (v3 + 1.5)
                                                                var v40 uint32 = 2147483648
                                                                _ = v40
                                                            }
                                                            var v41 uint32 = 487
                                                            if (v30 >= (^(v30))) {
                                                                switch fi11 {
                                                                case 128:
                                                                    switch fi36 {
                                                                    case 2147483647:
                                                                        var v42 bool = true
                                                                        _ = v42
                                                                    case 128:
                                                                        v43 := GenS2{F0: 127}
                                                                        _ = v43
                                                                    }
                                                                case 0:
                                                                    var v44 int8 = 100
                                                                    v41 -= (^(v41))
                                                                    switch fi36 {
                                                                    case 598:
                                                                        var v45 float32 = v3
                                                                        _ = v45
                                                                    default:
                                                                        v32 -= (v32 / ((v32) | 1))
                                                                    }
                                                                    if ((-(v44)) == v44) {
                                                                        v46 := GenS4{F0: (v2 % (((v2 >> ((195) & 63))) | 1)), F1: 4294967295, F2: (v41 & v41)}
                                                                        _ = v46
                                                                    } else {
                                                                        var v47 int16 = v2
                                                                        v30 += 0
                                                                        v48 := GenS5{F0: -34, F1: 806, F2: 599}
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    _ = v44
                                                                case 2147483647:
                                                                    fmt.Println(v1)
                                                                }
                                                                var v49 string = v4
                                                                v5[fi36] = (fi36 / ((fi36) | 1))
                                                                _ = v49
                                                            }
                                                            _ = v41
                                                        }
                                                        v50 := GenS5{F0: 454, F1: int32(v30), F2: 65535}
                                                        var v51 uint8 = 2
                                                        fmt.Println(fi11, fi36, v3)
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        switch ("\n\t" + (("a" + "hello") + "")) {
                                                        case "hello":
                                                            var v52 uint16 = 2
                                                            _ = v52
                                                        case "a":
                                                            switch v4 {
                                                            case "a":
                                                                var v53 uint16 = (v32 & (^(v32)))
                                                                _ = v53
                                                            }
                                                        default:
                                                            if false {
                                                                if true {
                                                                    if (0 != (v2 | 1)) {
                                                                        var v54 bool = ((v32 % ((v32) | 1)) <= (v32 - v32))
                                                                        var v55 float32 = (((((((-2.25 * ((v1 - (v3 / (((v1 - v3) / (v1 - (v1 / v1))) - v3))) / 1e10)) * 1.5) + v1) / v3) / -2.25) - v3) / -2.25)
                                                                        _ = v54
                                                                        _ = v55
                                                                    } else {
                                                                        if (!true) {
                                                                            var v56 int32 = 1
                                                                            var v57 uint64 = 2
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        switch int(v32) {
                                                                        case 127:
                                                                            v58 := GenS4{F0: -1, F1: 146, F2: 578}
                                                                            _ = v58
                                                                        case 42:
                                                                            var v59 uint32 = 2147483648
                                                                            _ = v59
                                                                        case 32767:
                                                                            if false {
                                                                                delete(v5, 42)
                                                                            }
                                                                            v60 := GenS6{F0: -128}
                                                                            fmt.Println(v30, v32)
                                                                            _ = v60
                                                                        }
                                                                        v61 := GenS6{F0: -1}
                                                                        var v62 int64 = -1
                                                                        v3 = v1
                                                                        switch ("a" + v4) {
                                                                        case "a":
                                                                            v2 += v2
                                                                        case "hello":
                                                                            v2 -= v2
                                                                        case "foo bar":
                                                                            var v63 int = (-(fi36))
                                                                            _ = v63
                                                                        default:
                                                                            if (0 >= 42) {
                                                                                switch 91 {
                                                                                case -9223372036854775808:
                                                                                    v64 := GenS7{F0: 127}
                                                                                    _ = v64
                                                                                case 32767:
                                                                                    if true {
                                                                                        var v65 int8 = 127
                                                                                        _ = v65
                                                                                    }
                                                                                    var v66 uint8 = 128
                                                                                    fmt.Println(fi11)
                                                                                    _ = v66
                                                                                case 1:
                                                                                    fmt.Println(v4, fi11)
                                                                                }
                                                                            }
                                                                        }
                                                                        var v67 bool = (12 > 103)
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v67
                                                                    }
                                                                    if false {
                                                                        var v68 uint = 42
                                                                        var v69 string = v4
                                                                        _ = v68
                                                                        _ = v69
                                                                    } else {
                                                                        v30 = (v30 - (v30 | 1))
                                                                        v32 -= v32
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v70 int = (fi36 % ((0) | 1))
                                                                        fmt.Println(fi36)
                                                                        v9 = append(v9, fi11)
                                                                        v71 := GenS8{F0: 2, F1: 1, F2: 65535}
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        v4 = (((((v4 + "foo bar") + v4) + "foo bar") + "foo bar") + v4)
                                                                    }
                                                                    fmt.Println(fi11, v1)
                                                                }
                                                                if (!false) {
                                                                    fmt.Println(v32, v32, fi11)
                                                                    var v72 bool = true
                                                                    if true {
                                                                        v32 += (v32 & (v32 - 202))
                                                                        v2 = -681
                                                                        var v73 float64 = ((-2.25 * 3.14159) + float64(v1))
                                                                        var v74 float64 = (-2.25 + v73)
                                                                        if (255 == 42) {
                                                                            fmt.Println(v30)
                                                                            v30 += 613
                                                                        }
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                    v30 = v30
                                                                    v2 = (-(v2))
                                                                    _ = v72
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if (!((v32 * (v32 ^ (v32 + (v32 << ((685) & 63))))) == 2)) {
                                                        v75 := v5[728]
                                                        var v76 uint8 = 42
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        fmt.Println(fi36, v1, v30)
                                                        fmt.Println(v4, v3, v30)
                                                    }
                                                    v77 := GenS7{F0: 0}
                                                    _ = fi36
                                                    _ = v77
                                                }
                                                v78 := GenS3{F0: 2147483647, F1: false, F2: v4}
                                                fmt.Println(v1, v30, fi11)
                                                _ = v78
                                            }
                                        default:
                                            switch 354 {
                                            case 843:
                                                v32 = 1
                                            }
                                        }
                                        v79 := v5[1]
                                        _ = v79
                                    }
                                    var v80 uint = 2
                                    _ = v80
                                }
                                _ = v32
                            }
                            _ = v30
                        }
                    }
                }
                switch (fi11 + int(v2)) {
                case -557:
                    switch v4 {
                    case "a":
                        v3 += v1
                    case "hello":
                        var v81 uint64 = 18446744073709551615
                        _ = v81
                    }
                case -2147483648:
                    var v82 int64 = 1
                    _ = v82
                }
            }
            _ = v12
        } else {
            if (42 == 4294967295) {
                v83 := GenS1{F0: ((genFn0(65535, 255) < int64(v2)) || true), F1: 2147483647, F2: -2147483648}
                v84 := GenS8{F0: 42, F1: -1, F2: 65535}
                v85 := GenS3{F0: 553, F1: (v2 >= v2), F2: ((v4 + v4) + (v4 + ((("foo bar" + v4) + "hello") + ((("" + v4) + v4) + v4))))}
                _ = v83
                _ = v84
                _ = v85
            }
        }
        fmt.Println(v3)
        _ = fi11
    }
    if false {
        v3 += v3
    } else {
        fmt.Println(v2, v4)
        if true {
            if (-9223372036854775808 <= genFn0(18446744073709551615, 2)) {
                if false {
                    v86 := GenS6{F0: 100}
                    switch ("hello" + v4) {
                    case "a":
                        for fi87 := 0; fi87 < 13; fi87++ {
                            var v88 string = v4
                            _ = fi87
                            _ = v88
                        }
                    default:
                        v89 := GenS9{F0: 255, F1: (1.5 == float32(((-2.25 / 1.0) / 1.5))), F2: 483}
                        _ = v89
                    }
                    _ = v86
                }
                for fi90 := 0; fi90 < 11; fi90++ {
                    for fi91 := 0; fi91 < 5; fi91++ {
                        switch v4 {
                        case "\n\t":
                            v92 := GenS10{F0: (((291 == 0) || false) && true), F1: 427, F2: true}
                            _ = v92
                        }
                        v4 = v4
                        var v93 int32 = -1
                        switch (fi91 ^ (fi91 ^ 32767)) {
                        case 32767:
                            v4 = v4
                        case -1:
                            var v94 int8 = -111
                            _ = v94
                        case 0:
                            v95 := GenS4{F0: v2, F1: 35, F2: 0}
                            _ = v95
                        default:
                            if (false && (((((((((-2.25 - 1e10) * (3.14159 - -2.25)) + (1.5 * (3.14159 * (1e10 * 1e10)))) + 3.14159) - ((3.14159 - (0.1 + (3.14159 - 1e10))) - 3.14159)) - (float64(float32((0.0 - float64(((1e10 - v1) + float32(3.14159)))))) - 3.14159)) + 0.0) != 0.1) || (815 == 969))) {
                                var v96 float64 = 1.5
                                fmt.Println(fi91, v2, v3)
                                switch fi90 {
                                case 0:
                                    fmt.Println(v4)
                                case 646:
                                    _ = len(v5)
                                case 128:
                                    fmt.Println(v2)
                                }
                                _ = v96
                            } else {
                                v2 += v2
                                v2 -= -32768
                                v3 = v3
                            }
                            var v97 int64 = -932
                            _ = v97
                        }
                        _ = fi91
                        _ = v93
                    }
                    var v98 int8 = 100
                    switch fi90 {
                    case -9223372036854775808:
                        if false {
                            for fi99 := 0; fi99 < 20; fi99++ {
                                v100 := GenS11{F0: fi90, F1: (v2 % ((v2) | 1)), F2: (v4 + v4)}
                                _ = fi99
                                _ = v100
                            }
                            var v101 int16 = (v2 | -817)
                            v102 := GenS1{F0: true, F1: 0, F2: 0}
                            _ = v101
                            _ = v102
                        } else {
                            fmt.Println(v98)
                            switch (v4 + v4) {
                            case "":
                                fmt.Println(v4)
                            }
                            if (false && (!true)) {
                                v103 := GenS12{F0: genFn0(277, 2)}
                                if (!true) {
                                    v2 = 1000
                                    if true {
                                        var v104 uint64 = 570
                                        _ = v104
                                    }
                                    fmt.Println(v2, v3)
                                } else {
                                    if (255 > 42) {
                                        if true {
                                            v2 = (-(v2))
                                        }
                                        v98 += v98
                                        for fi105 := 0; fi105 < 2; fi105++ {
                                            var v106 int64 = 9223372036854775807
                                            switch -312 {
                                            case 997:
                                                var v107 int32 = int32(v98)
                                                _ = v107
                                            default:
                                                v4 = ("foo bar" + v4)
                                            }
                                            _ = fi105
                                            _ = v106
                                        }
                                    }
                                    v2 = v2
                                    v108 := v5[42]
                                    _ = v108
                                }
                                _ = v103
                            } else {
                                var v109 int = 127
                                if (v109 <= v109) {
                                    var v110 string = v4
                                    delete(v5, fi90)
                                    var v111 uint32 = 710
                                    _ = v110
                                    _ = v111
                                } else {
                                    switch v4 {
                                    case "\n\t":
                                        var v112 float64 = (((3.14159 * 1e10) * (1.5 + 1.5)) * 1e10)
                                        _ = v112
                                    }
                                    var v113 string = ((v4 + (v4 + "hello")) + "")
                                    fmt.Println(fi90)
                                    _ = v113
                                }
                                _ = v109
                            }
                        }
                    case 32767:
                        _ = len(v5)
                    }
                    _ = fi90
                    _ = v98
                }
            }
        }
    }
    for fi114 := 0; fi114 < 8; fi114++ {
        switch "\n\t" {
        case "":
            if len(v9) > 0 {
                v9[0] = int(v2)
            }
        case "\n\t":
            if (127 > 127) {
                var v115 uint64 = 1
                if true {
                    for fi116 := 0; fi116 < 9; fi116++ {
                        var v117 string = v4
                        v1 = float32(float64(v1))
                        _ = fi116
                        _ = v117
                    }
                    v4 = "hello"
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v118 uint8 = 255
                        _ = v118
                    case "foo bar":
                        v4 = ((((v4 + "hello") + "a") + "a") + ("foo bar" + "a"))
                    }
                    var v119 int = 1
                    _ = v119
                }
                _ = v115
            } else {
                for fi120 := 0; fi120 < 9; fi120++ {
                    v2 -= v2
                    var v121 string = "foo bar"
                    v122 := GenS13{F0: (fi114 % ((-253) | 1)), F1: v2, F2: v2}
                    v1 = v1
                    _ = fi120
                    _ = v121
                    _ = v122
                }
                for fi123 := 0; fi123 < 8; fi123++ {
                    v3 += ((((float32((0.1 * 3.14159)) * v1) + float32((float64(3.14159) + 1.5))) + float32(1e10)) * v1)
                    v124 := GenS14{F0: 65535, F1: 456, F2: 781}
                    if true {
                        v125 := GenS10{F0: (!true), F1: 18, F2: false}
                        if (446 == -462) {
                            if ((v3 / float32(3.14159)) < (v3 - (v3 * 0.1))) {
                                v4 = v4
                                switch v4 {
                                case "a":
                                    switch fi123 {
                                    case -1:
                                        var v126 float32 = 1e10
                                        _ = v126
                                    case -2147483648:
                                        if (fi114 < (fi123 + (fi114 + (^(fi123))))) {
                                            v3 = (v1 / (1.5 / float32(1.5)))
                                            v3 = (float32(((1.5 - (((1e10 - 3.14159) / 1.0) + (1.5 * (float64(1e10) / ((0.1 + 1e10) / 1.5))))) * 1.5)) * (3.14159 * 0.0))
                                            delete(v5, (fi114 * fi123))
                                            v1 = v1
                                            var v127 uint = 65535
                                            v1 = 0.1
                                            var v128 int64 = 42
                                            v129 := GenS15{F0: 2, F1: false}
                                            _ = v127
                                            _ = v128
                                            _ = v129
                                        } else {
                                            var v130 float32 = 0.1
                                            _ = v130
                                        }
                                    }
                                default:
                                    v4 = (v4 + v4)
                                }
                            } else {
                                v131 := v5[fi114]
                                v3 -= (float32(1.5) + v3)
                                _ = len(v9)
                                var v132 float32 = v3
                                fmt.Println(v131, v1, v131)
                                v133 := GenS16{F0: 93, F1: -1}
                                _ = v131
                                _ = v132
                                _ = v133
                            }
                            v2 += v2
                            v134 := GenS7{F0: -52}
                            _ = v134
                        } else {
                            _ = len(v5)
                        }
                        _ = v125
                    } else {
                        var v135 uint8 = 61
                        v135 -= v135
                        v2 = (v2 ^ 0)
                        var v136 float32 = (1.5 * (v1 * v1))
                        var v137 int16 = v2
                        _ = v135
                        _ = v136
                        _ = v137
                    }
                    if false {
                        var v138 bool = true
                        if false {
                            if v138 {
                                v9 = append(v9, fi123)
                                v139 := GenS16{F0: 127, F1: -89}
                                _ = v139
                            } else {
                                if (2 < 0) {
                                    switch "\n\t" {
                                    case "a":
                                        var v140 uint = 18446744073709551615
                                        _ = v140
                                    default:
                                        if (!v138) {
                                            v141 := GenS2{F0: 255}
                                            fmt.Println(fi114, fi114, fi114)
                                            if (42 > 679) {
                                                var v142 uint64 = 4294967295
                                                if (((v4 + "a") + v4) > "\n\t") {
                                                    v143 := GenS9{F0: 42, F1: v138, F2: 281}
                                                    var v144 int64 = -1
                                                    _ = v143
                                                    _ = v144
                                                } else {
                                                    v2 = (v2 & (v2 | v2))
                                                    var v145 int = (^(fi123))
                                                    _ = v145
                                                }
                                                _ = v142
                                            } else {
                                                switch v4 {
                                                case "hello":
                                                    v3 = 3.14159
                                                case "foo bar":
                                                    v4 = v4
                                                case "":
                                                    if v138 {
                                                        if (255 != 27) {
                                                            if false {
                                                                var v146 int64 = -879
                                                                _ = v146
                                                            } else {
                                                                delete(v5, fi123)
                                                            }
                                                            if false {
                                                                var v147 int16 = (-(v2))
                                                                switch "" {
                                                                case "\n\t":
                                                                    v148 := GenS13{F0: fi114, F1: 32767, F2: v147}
                                                                    _ = v148
                                                                case "foo bar":
                                                                    if (-307 != 66) {
                                                                        v1 += float32(0.1)
                                                                    } else {
                                                                        v1 += 3.14159
                                                                        var v149 int32 = 2147483647
                                                                        _ = v149
                                                                    }
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v150 int8 = 0
                                                                        _ = v150
                                                                    }
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                case "":
                                                                    v151 := GenS12{F0: genFn0(42, 1)}
                                                                    _ = v151
                                                                default:
                                                                    if (v138 || (v4 >= "hello")) {
                                                                        delete(v5, (fi114 - fi123))
                                                                    } else {
                                                                        v2 = (v2 - (v2 & (v147 >> ((2) & 63))))
                                                                        if v138 {
                                                                            fmt.Println(v138, v138)
                                                                            var v152 int8 = -1
                                                                            _ = v152
                                                                        } else {
                                                                            var v153 uint8 = 2
                                                                            _ = v153
                                                                        }
                                                                    }
                                                                }
                                                                v1 -= (float32(0.0) * (v3 * (v3 + v1)))
                                                                _ = v147
                                                            } else {
                                                                v1 -= 3.14159
                                                            }
                                                        } else {
                                                            switch "foo bar" {
                                                            case "a":
                                                                v138 = (((true || false) || false) || v138)
                                                            }
                                                            v9 = append(v9, (fi114 / ((fi114) | 1)))
                                                            v4 = (v4 + (v4 + "foo bar"))
                                                        }
                                                        v4 = "a"
                                                    }
                                                default:
                                                    v154 := GenS12{F0: 647}
                                                    _ = v154
                                                }
                                            }
                                            _ = v141
                                        } else {
                                            var v155 uint = 1
                                            _ = v155
                                        }
                                        v4 = v4
                                    }
                                }
                                v5[0] = (fi123 + (fi114 | fi114))
                            }
                            if ((fi114 >> ((933) & 63)) == (fi114 * 2147483647)) {
                                var v156 int32 = 2147483647
                                var v157 int8 = -1
                                var v158 int16 = v2
                                var v159 float64 = 1.5
                                var v160 int32 = v156
                                _ = v156
                                _ = v157
                                _ = v158
                                _ = v159
                                _ = v160
                            } else {
                                var v161 int8 = -128
                                _ = v161
                            }
                            v162 := GenS3{F0: 500, F1: v138, F2: v4}
                            _ = v162
                        }
                        _ = v138
                    }
                    _ = fi123
                    _ = v124
                }
            }
        case "a":
            switch (v4 + ((((("a" + v4) + (((v4 + ("foo bar" + v4)) + (v4 + ((v4 + v4) + v4))) + v4)) + (("foo bar" + ("\n\t" + v4)) + (v4 + v4))) + (((v4 + v4) + (v4 + (v4 + "\n\t"))) + "")) + v4)) {
            case "hello":
                v163 := GenS17{F0: -264, F1: v2}
                _ = v163
            }
        }
        _ = fi114
    }
    v164 := GenS11{F0: 127, F1: (v2 * -32768), F2: "\n\t"}
    var v165 float32 = 0.0
    v166 := GenS13{F0: 4294967295, F1: (v2 >> ((255) & 63)), F2: (-(v2))}
    v1 += -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v164
    _ = v165
    _ = v166
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
