package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 uint; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint16; F1 int }
type GenS12 struct { F0 uint32; F1 int16; F2 bool }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint; F1 uint; F2 string }
type GenS16 struct { F0 int; F1 uint64; F2 int64 }
func genFn0(p0 uint, p1 int32, p2 uint32) uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            switch (fi5 % ((fi5) | 1)) {
            case 4294967295:
                p0 -= 65535
            }
            var v7 uint64 = 0
            if (3.14159 < 1e10) {
                var v8 string = ""
                _ = v8
            } else {
                switch fi6 {
                case 9223372036854775807:
                    switch (fi6 % ((fi6) | 1)) {
                    case -9223372036854775808:
                        var v9 uint8 = 20
                        _ = v9
                    case 0:
                        v10 := GenS1{F0: 127, F1: (fi5 | fi5), F2: (p0 << ((p0) & 63))}
                        _ = v10
                    default:
                        fmt.Println(fi6)
                    }
                case 128:
                    var v11 string = "foo bar"
                    _ = v11
                case -393:
                    p1 -= (p1 + p1)
                }
                v7 = v7
            }
            _ = fi6
            _ = v7
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            p2 = (^(p2))
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        var v14 string = ((("a" + "\n\t") + "hello") + ("foo bar" + "foo bar"))
        _ = fi5
        _ = v14
    }
    if (false && true) {
        if (2 <= 82) {
            var v15 int = 2147483647
            _ = v15
        } else {
            v16 := GenS1{F0: 220, F1: 32768, F2: (p0 << (((p0 + p0)) & 63))}
            _ = v16
        }
        var v17 int32 = p1
        var v18 float64 = float64(0.1)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        switch 4294967295 {
        case 32768:
            if false {
                v20 := GenS2{F0: 112, F1: 4294967295, F2: 127}
                p2 = (^(p2))
                v21 := GenS2{F0: 2, F1: (fi19 * fi19), F2: 1}
                _ = v20
                _ = v21
            }
        case 0:
            if true {
                if true {
                    p2 = p2
                    fmt.Println(p0)
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 int = 32767
                        if (!false) {
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        } else {
                            v23 = fi22
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                v25 := []int{467, fi19, fi19, fi19}
                if true {
                    if true {
                        v25 = append(v25, (fi19 | int(p0)))
                    }
                }
                _ = v25
            } else {
                p0 = p0
                if true {
                    fmt.Println(fi19, p1)
                    p2 += p2
                }
            }
        }
        _ = fi19
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int64 = 42
        _ = fi26
        _ = v27
    }
    v28 := GenS2{F0: 127, F1: 1, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return p0
}
func genFn1(p0 uint) int16 {
    switch 285 {
    case 127:
        p0 += 65535
    case -21:
        v29 := GenS1{F0: 0, F1: -9223372036854775808, F2: 21}
        _ = v29
    }
    var v30 uint = 0
    _ = p0
    _ = v30
    return 0
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := []string{}
    var v3 bool = (!false)
    v4 := make(map[int]string)
    _ = len(v4)
    if (float32(float64((float32(1e10) * 0.0))) <= 3.14159) {
        fmt.Println(v3)
        var v31 float32 = 1e10
        _ = v31
    }
    var v32 string = "\n\t"
    switch v32 {
    case "\n\t":
        var v33 int8 = 100
        _ = v33
    case "a":
        if v3 {
            if len(v2) > 0 {
                v2[0] = (v32 + ("hello" + v32))
            }
        } else {
            if v3 {
                v32 = v32
                for fi34 := 0; fi34 < 16; fi34++ {
                    v35 := GenS3{F0: genFn0(605, -1, 2)}
                    v3 = v3
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 20; fi36++ {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        v3 = false
                        var v38 uint8 = 127
                        _ = fi37
                        _ = v38
                    }
                    _ = fi36
                }
            }
            v39 := GenS3{F0: 867}
            switch v32 {
            case "":
                v32 = "hello"
            default:
                var v40 int8 = -128
                _ = v40
            }
            _ = v39
        }
    }
    v41 := v4[-9223372036854775808]
    switch "hello" {
    case "foo bar":
        v42 := GenS4{F0: 99}
        _ = v42
    case "\n\t":
        v41 = v32
    case "hello":
        var v43 float64 = (-2.25 + (-2.25 + 1.5))
        _ = v43
    }
    var v44 int = -599
    v45 := v4[v44]
    for fi46 := 0; fi46 < 13; fi46++ {
        v47 := GenS3{F0: 2}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi46
        _ = v47
    }
    v48 := GenS2{F0: 0, F1: (-(v44)), F2: 145}
    v32 = "hello"
    switch 0 {
    case -9223372036854775808:
        v32 = v45
    case 42:
        for fi49 := 0; fi49 < 3; fi49++ {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = len(v2)
            _ = fi49
        }
    case 93:
        var v50 uint16 = 40
        _ = v50
    default:
        if (false || v3) {
            v41 = v45
            switch "foo bar" {
            case "a":
                var v51 uint64 = 4294967295
                _ = v51
            case "\n\t":
                var v52 uint = 0
                _ = v52
            }
            v32 = v32
        } else {
            var v53 int64 = 0
            v1 = append(v1, v44)
            for fi54 := 0; fi54 < 9; fi54++ {
                fmt.Println(v32, v45)
                var v55 uint64 = 873
                _ = fi54
                _ = v55
            }
            _ = v53
        }
    }
    if false {
        switch (v41 + "foo bar") {
        case "foo bar":
            if true {
                v45 = "\n\t"
            }
        case "hello":
            v2 = append(v2, "")
        }
    } else {
        for fi56 := 0; fi56 < 14; fi56++ {
            if false {
                v3 = v3
                v41 = v32
                switch v32 {
                case "a":
                    if v3 {
                        if true {
                            var v57 float64 = 0.0
                            var v58 int8 = -104
                            _ = v57
                            _ = v58
                        } else {
                            for fi59 := 0; fi59 < 4; fi59++ {
                                fmt.Println(v3, v45, v44)
                                _ = fi59
                            }
                        }
                    } else {
                        v60 := GenS2{F0: 22, F1: fi56, F2: 127}
                        var v61 int32 = 1
                        _ = v60
                        _ = v61
                    }
                case "\n\t":
                    v62 := GenS4{F0: 1}
                    _ = v62
                case "foo bar":
                    v44 = fi56
                default:
                    v45 = "foo bar"
                }
            }
            var v63 float32 = ((1e10 * (-2.25 - 0.0)) * 0.1)
            if (-9223372036854775808 > -887) {
                switch (v44 + v44) {
                case -961:
                    v4[-11] = ""
                case 9223372036854775807:
                    var v64 uint32 = 2
                    _ = v64
                case -482:
                    v65 := GenS4{F0: 1}
                    _ = v65
                default:
                    for fi66 := 0; fi66 < 12; fi66++ {
                        if (!v3) {
                            if ((fi66 / ((fi66) | 1)) >= v44) {
                                if len(v1) > 0 {
                                    v1[0] = (fi56 / (((fi66 >> ((214) & 63))) | 1))
                                }
                            } else {
                                switch v32 {
                                case "foo bar":
                                    v67 := GenS4{F0: -1}
                                    _ = v67
                                case "a":
                                    fmt.Println(v63, v44)
                                case "hello":
                                    v63 = v63
                                default:
                                    fmt.Println(v63, fi56, v63)
                                    if (110 >= -78) {
                                        v68 := GenS5{F0: v3, F1: 873}
                                        v69 := GenS5{F0: ((0.0 > 0.0) || (42 == 1)), F1: 42}
                                        switch (v45 + "hello") {
                                        case "a":
                                            v3 = v3
                                        case "hello":
                                            v70 := GenS6{F0: (false && (1 != -1))}
                                            v41 = v32
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v70
                                        case "foo bar":
                                            v44 += (fi56 << ((587) & 63))
                                        }
                                        if (!v3) {
                                            if v3 {
                                                var v71 float64 = 0.0
                                                _ = len(v2)
                                                if true {
                                                    if (v3 || false) {
                                                        fmt.Println(v41, fi56)
                                                        v72 := GenS4{F0: 2147483647}
                                                        v32 = v45
                                                        _ = v72
                                                    } else {
                                                        delete(v4, (fi56 + fi56))
                                                        v32 = (v45 + v45)
                                                        var v73 uint16 = 16
                                                        switch v44 {
                                                        case 1:
                                                            v74 := GenS5{F0: (v71 != -2.25), F1: 425}
                                                            _ = v74
                                                        case 252:
                                                            if (!v3) {
                                                                v71 += v71
                                                            } else {
                                                                if false {
                                                                    v75 := GenS7{F0: 127, F1: 2, F2: v45}
                                                                    _ = v75
                                                                }
                                                                switch fi56 {
                                                                case 32767:
                                                                    v44 -= (v44 >> ((42) & 63))
                                                                case 4294967295:
                                                                    switch "a" {
                                                                    case "":
                                                                        if ((-1 != 1) && (42 >= 1)) {
                                                                            v63 -= (v63 * v63)
                                                                            v76 := v4[(fi56 + (fi56 ^ (v44 | (fi56 << ((584) & 63)))))]
                                                                            if ((!(v63 >= (v63 / v63))) && (1 == int64(v73))) {
                                                                                fmt.Println(fi56)
                                                                                switch (v44 / (((v44 | fi66)) | 1)) {
                                                                                case -1:
                                                                                    var v77 uint64 = 2
                                                                                    _ = v77
                                                                                }
                                                                                _ = len(v4)
                                                                                if ((float32(v71) - v63) != v63) {
                                                                                    var v78 int64 = -9223372036854775808
                                                                                    _ = v78
                                                                                } else {
                                                                                    var v79 float32 = (v63 + v63)
                                                                                    if (v3 || v3) {
                                                                                        if (3.14159 <= 3.14159) {
                                                                                            v71 += ((float64(v63) / -2.25) + (v71 + ((1.5 * 1.5) - v71)))
                                                                                        } else {
                                                                                            v45 = v45
                                                                                            var v80 int64 = -606
                                                                                            _ = v80
                                                                                        }
                                                                                        var v81 int8 = 127
                                                                                        _ = v81
                                                                                    }
                                                                                    switch v41 {
                                                                                    case "hello":
                                                                                        v63 = ((float32(((v71 / v71) + v71)) + (-2.25 / v63)) - (float32(v71) + float32(v71)))
                                                                                    case "a":
                                                                                        var v82 int64 = -1
                                                                                        _ = v82
                                                                                    }
                                                                                    _ = v79
                                                                                }
                                                                                v41 = v32
                                                                            } else {
                                                                                fmt.Println(fi56, v73)
                                                                                var v83 float32 = v63
                                                                                _ = v83
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                    case "foo bar":
                                                                        var v84 float32 = (v63 - ((float32((1e10 / v71)) + v63) / (0.1 + (v63 / v63))))
                                                                        _ = v84
                                                                    default:
                                                                        fmt.Println(v44, v41)
                                                                    }
                                                                default:
                                                                    switch (fi66 % (((fi56 | (v44 << ((255) & 63)))) | 1)) {
                                                                    case -1:
                                                                        v85 := GenS5{F0: false, F1: 65535}
                                                                        _ = v85
                                                                    case 42:
                                                                        if (-465 == -1) {
                                                                            var v86 uint = 798
                                                                            _ = v86
                                                                        }
                                                                    case 32767:
                                                                        v45 = "foo bar"
                                                                    }
                                                                }
                                                                v63 -= float32(v71)
                                                            }
                                                        case -9223372036854775808:
                                                            if v3 {
                                                                fmt.Println(v45, v45)
                                                                v4[-9223372036854775808] = (v41 + v41)
                                                            } else {
                                                                var v87 int64 = 9223372036854775807
                                                                switch fi56 {
                                                                case 839:
                                                                    v45 = ("\n\t" + ((("" + v41) + "a") + "hello"))
                                                                case 127:
                                                                    var v88 float64 = 0.0
                                                                    _ = v88
                                                                default:
                                                                    var v89 float32 = v63
                                                                    _ = v89
                                                                }
                                                                v90 := v4[(v44 - (fi66 ^ (v44 / (((fi66 + (fi56 & (fi56 & v44)))) | 1))))]
                                                                _ = v87
                                                                _ = v90
                                                            }
                                                        }
                                                        switch (v44 ^ (v44 / (((fi56 / ((127) | 1))) | 1))) {
                                                        case 294:
                                                            if v3 {
                                                                var v91 uint8 = 208
                                                                if len(v2) > 0 {
                                                                    v2[0] = "foo bar"
                                                                }
                                                                if false {
                                                                    fmt.Println(v73, v3)
                                                                    v3 = true
                                                                    v92 := v4[42]
                                                                    _ = v92
                                                                }
                                                                _ = v91
                                                            } else {
                                                                v63 -= (((v63 * v63) + 0.0) * v63)
                                                                if (!(v73 < (v73 >> ((18446744073709551615) & 63)))) {
                                                                    v41 = "hello"
                                                                    _ = len(v1)
                                                                    if ((-1 < 1) && (810 != 0)) {
                                                                        switch (-(fi66)) {
                                                                        case 868:
                                                                            v93 := GenS5{F0: v3, F1: 2}
                                                                            _ = v93
                                                                        case -2147483648:
                                                                            v41 = v45
                                                                        case 32767:
                                                                            if (-52 > 127) {
                                                                                v94 := GenS8{F0: 0}
                                                                                _ = v94
                                                                            } else {
                                                                                if v3 {
                                                                                    v73 = (v73 * (v73 ^ v73))
                                                                                    if false {
                                                                                        var v95 int16 = -1
                                                                                        v95 += v95
                                                                                        if true {
                                                                                            v44 += v44
                                                                                            v63 += v63
                                                                                            if (-2.25 >= v71) {
                                                                                                var v96 float32 = (float32(0.1) * v63)
                                                                                                var v97 bool = true
                                                                                                if false {
                                                                                                    var v98 int16 = (v95 ^ v95)
                                                                                                    v1 = append(v1, fi66)
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = (fi56 / ((42) | 1))
                                                                                                    }
                                                                                                    _ = v98
                                                                                                }
                                                                                                var v99 uint64 = 1
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                                _ = v99
                                                                                            } else {
                                                                                                switch int(v95) {
                                                                                                case 128:
                                                                                                    if (251 == -178) {
                                                                                                        v100 := GenS7{F0: -923, F1: genFn0(1, 0, 4294967295), F2: ("hello" + (((v32 + "\n\t") + (v41 + v41)) + (v45 + v32)))}
                                                                                                        var v101 uint64 = 205
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    }
                                                                                                case 32768:
                                                                                                    v71 = v71
                                                                                                default:
                                                                                                    if (v3 || v3) {
                                                                                                        v32 = v41
                                                                                                        if (42 >= 61) {
                                                                                                            var v102 int16 = genFn1(2)
                                                                                                            var v103 int64 = -9223372036854775808
                                                                                                            _ = v102
                                                                                                            _ = v103
                                                                                                        } else {
                                                                                                            var v104 uint8 = 255
                                                                                                            v105 := GenS1{F0: v104, F1: fi56, F2: 18446744073709551615}
                                                                                                            _ = v104
                                                                                                            _ = v105
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if false {
                                                                                                fmt.Println(v44)
                                                                                            }
                                                                                            fmt.Println(v95, v44, v73)
                                                                                        }
                                                                                        _ = v95
                                                                                    } else {
                                                                                        var v106 int8 = 127
                                                                                        var v107 uint = 1
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                    }
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                } else {
                                                                                    v41 = "\n\t"
                                                                                    v108 := GenS9{F0: 1, F1: 70, F2: 65535}
                                                                                    _ = v108
                                                                                }
                                                                            }
                                                                        }
                                                                        var v109 int32 = -2147483648
                                                                        _ = v109
                                                                    }
                                                                }
                                                            }
                                                        case 128:
                                                            v45 = v41
                                                        case -9223372036854775808:
                                                            if (false && true) {
                                                                if ((42 <= 88) || (132 != (^(v73)))) {
                                                                    var v110 uint8 = 127
                                                                    _ = v110
                                                                }
                                                                var v111 uint = 2
                                                                if len(v2) > 0 {
                                                                    v2[0] = "foo bar"
                                                                }
                                                                _ = v111
                                                            } else {
                                                                v1 = append(v1, v44)
                                                            }
                                                        default:
                                                            switch ("hello" + ("" + v41)) {
                                                            case "":
                                                                v112 := GenS1{F0: 0, F1: v44, F2: 42}
                                                                _ = v112
                                                            default:
                                                                fmt.Println(v73, v45)
                                                            }
                                                        }
                                                        _ = v73
                                                    }
                                                }
                                                _ = v71
                                            }
                                            var v113 string = v41
                                            _ = v113
                                        } else {
                                            fmt.Println(v45, fi56, fi66)
                                            v114 := GenS5{F0: v3, F1: 361}
                                            if v3 {
                                                if ((!(-9223372036854775808 != -9223372036854775808)) && false) {
                                                    fmt.Println(v45, v32, v32)
                                                } else {
                                                    switch 574 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v41, v41)
                                                    case -545:
                                                        var v115 uint16 = 1
                                                        _ = v115
                                                    default:
                                                        switch fi56 {
                                                        case 127:
                                                            var v116 string = (("\n\t" + v32) + v32)
                                                            _ = v116
                                                        case 9223372036854775807:
                                                            switch "foo bar" {
                                                            case "a":
                                                                if (36 != 0) {
                                                                    v117 := GenS10{F0: -1, F1: 2}
                                                                    var v118 uint64 = 1
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                            case "foo bar":
                                                                if (128 > 0) {
                                                                    var v119 bool = true
                                                                    switch (fi56 / (((v44 & v44)) | 1)) {
                                                                    case -1:
                                                                        var v120 int16 = 42
                                                                        _ = v120
                                                                    case 565:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case 773:
                                                                        if true {
                                                                            fmt.Println(fi66)
                                                                            if false {
                                                                                var v121 int16 = -1
                                                                                if (v3 || v119) {
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        var v122 float64 = 3.14159
                                                                                        _ = v122
                                                                                    case "\n\t":
                                                                                        if (!v3) {
                                                                                            v4[(-(fi66))] = v32
                                                                                            switch -9223372036854775808 {
                                                                                            case 32768:
                                                                                                if (348 != -596) {
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                            case 32767:
                                                                                                var v123 float32 = v63
                                                                                                _ = v123
                                                                                            case 1:
                                                                                                var v124 uint16 = 0
                                                                                                _ = v124
                                                                                            default:
                                                                                                _ = len(v4)
                                                                                            }
                                                                                            switch (v44 - (v44 + v44)) {
                                                                                            case 127:
                                                                                                var v125 uint = 65535
                                                                                                _ = v125
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(fi56, v63)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v126 uint16 = 2
                                                                                        _ = v126
                                                                                    }
                                                                                } else {
                                                                                    v44 += (fi56 - 4294967295)
                                                                                    switch v32 {
                                                                                    case "hello":
                                                                                        var v127 uint64 = 1
                                                                                        _ = v127
                                                                                    case "a":
                                                                                        var v128 int16 = v121
                                                                                        _ = v128
                                                                                    default:
                                                                                        if v119 {
                                                                                            v3 = (v3 && ((!(95 >= 100)) && (342 == 0)))
                                                                                            if (-327 != 0) {
                                                                                                v41 = ""
                                                                                                switch -233 {
                                                                                                case 42:
                                                                                                    var v129 bool = v3
                                                                                                    _ = v129
                                                                                                case -2147483648:
                                                                                                    var v130 string = (v41 + (v41 + v32))
                                                                                                    _ = v130
                                                                                                case 128:
                                                                                                    fmt.Println(v32)
                                                                                                    _ = len(v2)
                                                                                                    v63 = v63
                                                                                                }
                                                                                                v131 := GenS2{F0: 70, F1: (fi66 ^ -822), F2: 3}
                                                                                                _ = v131
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v4[v44] = ("hello" + ("\n\t" + v41))
                                                                                }
                                                                                v119 = true
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v132 bool = (2147483647 < 901)
                                                                                if true {
                                                                                    var v133 uint64 = 42
                                                                                    _ = v133
                                                                                }
                                                                                _ = v121
                                                                                _ = v132
                                                                            } else {
                                                                                var v134 int = v44
                                                                                v134 = (v44 ^ (fi66 * 128))
                                                                                _ = v134
                                                                            }
                                                                            if v119 {
                                                                                v135 := GenS11{F0: 65535, F1: (v44 << ((118) & 63))}
                                                                                if (!((v63 - (v63 + v63)) == v63)) {
                                                                                    var v136 float64 = 0.0
                                                                                    fmt.Println(v136, v41, v32)
                                                                                    _ = v136
                                                                                } else {
                                                                                    v137 := v4[v44]
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        v3 = false
                                                                                    case "":
                                                                                        v4[(fi66 >> ((0) & 63))] = ""
                                                                                    case "foo bar":
                                                                                        v138 := GenS12{F0: 469, F1: genFn1(genFn0(65535, 1, 481)), F2: true}
                                                                                        _ = v138
                                                                                    default:
                                                                                        v45 = v45
                                                                                    }
                                                                                    _ = v137
                                                                                }
                                                                                _ = v135
                                                                            } else {
                                                                                v32 = "hello"
                                                                                v44 -= v44
                                                                                v139 := GenS3{F0: 18446744073709551615}
                                                                                _ = v139
                                                                            }
                                                                        }
                                                                    default:
                                                                        v41 = "hello"
                                                                    }
                                                                    if true {
                                                                        if false {
                                                                            var v140 int8 = -1
                                                                            v141 := GenS7{F0: -9223372036854775808, F1: 822, F2: v32}
                                                                            v45 = v45
                                                                            _ = v140
                                                                            _ = v141
                                                                        }
                                                                    }
                                                                    _ = v119
                                                                } else {
                                                                    v2 = append(v2, v41)
                                                                    var v142 int64 = 1
                                                                    _ = v142
                                                                }
                                                            case "\n\t":
                                                                _ = len(v4)
                                                            }
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                    }
                                                    if (!v3) {
                                                        var v143 int64 = 1
                                                        _ = v143
                                                    } else {
                                                        if v3 {
                                                            v144 := GenS9{F0: 120, F1: 255, F2: 955}
                                                            fmt.Println(v45, v3)
                                                            var v145 int64 = 9223372036854775807
                                                            _ = v144
                                                            _ = v145
                                                        }
                                                    }
                                                }
                                                var v146 int8 = -56
                                                _ = v146
                                            } else {
                                                var v147 bool = v3
                                                fmt.Println(v63, v63)
                                                _ = v147
                                            }
                                            _ = len(v4)
                                            _ = v114
                                        }
                                        _ = v68
                                        _ = v69
                                    } else {
                                        delete(v4, -138)
                                        switch (v44 >> ((1) & 63)) {
                                        case -2147483648:
                                            fmt.Println(v41)
                                        case 942:
                                            v32 = (v41 + v32)
                                        case 127:
                                            if v3 {
                                                var v148 float64 = (float64(v63) / (0.1 * 3.14159))
                                                var v149 uint32 = 1
                                                var v150 uint64 = 0
                                                var v151 float32 = 1e10
                                                _ = v148
                                                _ = v149
                                                _ = v150
                                                _ = v151
                                            }
                                        default:
                                            v152 := GenS13{F0: true}
                                            _ = v152
                                        }
                                    }
                                    v153 := GenS14{F0: 2}
                                    _ = v153
                                }
                            }
                            if true {
                                if false {
                                    if (51 != 1) {
                                        switch v45 {
                                        case "\n\t":
                                            var v154 int16 = -914
                                            _ = v154
                                        case "foo bar":
                                            if v3 {
                                                var v155 uint64 = 0
                                                var v156 int8 = 0
                                                _ = v155
                                                _ = v156
                                            }
                                            fmt.Println(v32)
                                        case "a":
                                            if v3 {
                                                v157 := GenS6{F0: v3}
                                                v63 = v63
                                                _ = v157
                                            }
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = 739
                                            }
                                            fmt.Println(fi56)
                                        }
                                    }
                                } else {
                                    if ((v41 + ("\n\t" + "hello")) >= v41) {
                                        v63 -= -2.25
                                        fmt.Println(fi56, fi56)
                                        v63 = v63
                                    } else {
                                        fmt.Println(v44, v63, v63)
                                        v41 = v41
                                        if false {
                                            _ = len(v2)
                                            var v158 int16 = -1
                                            _ = v158
                                        } else {
                                            if (!(!((((1e10 + ((3.14159 * (-2.25 * float64(float32(float64(float32(3.14159)))))) * 0.1)) + 1e10) / (((1e10 + ((1e10 / 0.1) / 3.14159)) + 1.5) * 0.1)) > (1.5 / 1.0)))) {
                                                v45 = v41
                                            }
                                        }
                                    }
                                    var v159 int16 = 1
                                    _ = v159
                                }
                                v63 += 3.14159
                            } else {
                                v160 := GenS12{F0: 0, F1: 430, F2: true}
                                _ = v160
                            }
                        } else {
                            var v161 int16 = 503
                            var v162 int16 = 32767
                            if (128 < 42) {
                                v41 = v41
                                switch 32767 {
                                case -812:
                                    var v163 uint = 255
                                    _ = v163
                                case -1:
                                    var v164 int8 = 45
                                    _ = v164
                                case 42:
                                    var v165 int = (fi66 - fi66)
                                    _ = v165
                                }
                                var v166 string = ((v32 + "") + (v41 + "foo bar"))
                                _ = v166
                            } else {
                                v167 := v4[fi66]
                                _ = len(v1)
                                if false {
                                    v168 := GenS10{F0: fi56, F1: 65535}
                                    v169 := GenS15{F0: 0, F1: 264, F2: ("" + (v41 + "\n\t"))}
                                    _ = v168
                                    _ = v169
                                } else {
                                    var v170 float32 = v63
                                    switch "hello" {
                                    case "hello":
                                        var v171 uint8 = 255
                                        _ = v171
                                    case "\n\t":
                                        v32 = "a"
                                    default:
                                        v44 = (v44 << ((4294967295) & 63))
                                    }
                                    v170 -= v63
                                    _ = v170
                                }
                                v172 := v4[(fi56 + fi66)]
                                v44 = v44
                                _ = v167
                                _ = v172
                            }
                            v63 = (float32(0.1) / (v63 + v63))
                            _ = v161
                            _ = v162
                        }
                        _ = fi66
                    }
                }
                for fi173 := 0; fi173 < 1; fi173++ {
                    v32 = (("hello" + (v45 + (v41 + v41))) + ("a" + ("" + v45)))
                    if (463 < 1) {
                        v174 := GenS10{F0: (v44 ^ (fi173 ^ fi56)), F1: 213}
                        _ = v174
                    } else {
                        switch -899 {
                        case 599:
                            if (127 != 66) {
                                var v175 int32 = -286
                                var v176 int = v44
                                var v177 uint32 = 0
                                _ = v175
                                _ = v176
                                _ = v177
                            }
                        case 794:
                            v32 = v41
                        case -2147483648:
                            fmt.Println(v44)
                        default:
                            v44 = -1
                        }
                        var v178 uint = genFn0(233, 42, 2147483647)
                        v179 := GenS16{F0: (v44 + 180), F1: 42, F2: 9223372036854775807}
                        _ = v178
                        _ = v179
                    }
                    _ = fi173
                }
            }
            _ = fi56
            _ = v63
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v48
}
