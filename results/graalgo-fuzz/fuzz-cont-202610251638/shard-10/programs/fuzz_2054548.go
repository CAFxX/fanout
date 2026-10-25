package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 uint8; F1 uint8; F2 bool }
type GenS11 struct { F0 uint; F1 uint64 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 int32; F1 int32; F2 uint }
type GenS14 struct { F0 uint16; F1 uint; F2 int16 }
func genFn0() bool {
    v6 := []int{-9223372036854775808}
    var v7 uint8 = 2
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 uint8, p1 int64, p2 uint64) uint32 {
    var v8 string = "foo bar"
    switch "hello" {
    case "foo bar":
        p2 = p2
    case "\n\t":
        var v9 uint = 65535
        _ = v9
    }
    var v10 float32 = 0.1
    switch "" {
    case "hello":
        if true {
            v8 = (v8 + "a")
            p0 = 255
            var v11 bool = (32767 != -137)
            _ = v11
        } else {
            var v12 uint64 = 1
            for fi13 := 0; fi13 < 13; fi13++ {
                if (2 < 18446744073709551615) {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v8 = ("" + "\n\t")
                        var v15 string = ((v8 + "") + v8)
                        var v16 int64 = p1
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if false {
                        var v17 uint64 = 2
                        v18 := []int{0}
                        _ = v17
                        _ = v18
                    } else {
                        v19 := []int{}
                        p0 += 42
                        _ = v19
                    }
                }
                v20 := []int{fi13, (fi13 - 4294967295), (fi13 / (((fi13 & fi13)) | 1)), (^(fi13))}
                _ = fi13
                _ = v20
            }
            _ = v12
        }
    }
    if (!false) {
        var v21 string = "a"
        if ((p1 >> ((0) & 63)) != (p1 | (p1 & p1))) {
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 bool = false
                var v24 uint32 = 2
                switch v8 {
                case "a":
                    if false {
                        switch 1 {
                        case -120:
                            if ((32767 >= 65535) || (1 >= -1)) {
                                v25 := GenS1{F0: 32767, F1: 127}
                                p1 = (p1 >> ((506) & 63))
                                v26 := []int{fi22, fi22, (fi22 - (fi22 / ((-959) | 1))), int(p1)}
                                _ = v25
                                _ = v26
                            }
                        case -413:
                            for fi27 := 0; fi27 < 16; fi27++ {
                                v28 := GenS2{F0: 311, F1: 255}
                                if v23 {
                                    var v29 uint64 = 589
                                    if v23 {
                                        var v30 int8 = -55
                                        _ = v30
                                    } else {
                                        var v31 int32 = 559
                                        if v23 {
                                            if false {
                                                v29 = 0
                                                if (42 <= 32767) {
                                                    v32 := GenS2{F0: (v31 / ((-771) | 1)), F1: 65535}
                                                    _ = v32
                                                } else {
                                                    var v33 int8 = -1
                                                    v34 := []int{(fi27 + fi22), (fi22 << ((501) & 63)), (fi27 / ((fi27) | 1))}
                                                    var v35 int64 = -437
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                                if (v23 || v23) {
                                                    v36 := []int{fi22}
                                                    v10 = (1e10 + (0.1 - (v10 / (v10 - v10))))
                                                    _ = v36
                                                } else {
                                                    var v37 int = fi22
                                                    _ = v37
                                                }
                                            }
                                            fmt.Println(p1)
                                        } else {
                                            if (v23 || (("foo bar" + v21) != "\n\t")) {
                                                v31 = v31
                                            } else {
                                                var v38 int32 = (v31 & 0)
                                                _ = v38
                                            }
                                            switch (fi22 / ((957) | 1)) {
                                            case -9223372036854775808:
                                                var v39 uint16 = 65535
                                                v10 += v10
                                                v40 := GenS1{F0: 1, F1: -1}
                                                _ = v39
                                                _ = v40
                                            }
                                            switch (v21 + "a") {
                                            case "\n\t":
                                                v41 := []int{}
                                                _ = v41
                                            case "hello":
                                                switch 2147483647 {
                                                case -319:
                                                    var v42 int32 = (v31 - (^(v31)))
                                                    _ = v42
                                                }
                                            case "foo bar":
                                                v43 := []int{0}
                                                _ = v43
                                            default:
                                                if ((v21 + "hello") < (v21 + (((v8 + v8) + (("foo bar" + v21) + v8)) + v8))) {
                                                    if v23 {
                                                        fmt.Println(p2, p1, v8)
                                                        p0 += (p0 - p0)
                                                        var v44 int8 = -57
                                                        _ = v44
                                                    } else {
                                                        v24 += v24
                                                        v45 := []int{fi27, 127}
                                                        _ = v45
                                                    }
                                                }
                                                v46 := []int{}
                                                _ = v46
                                            }
                                        }
                                        var v47 string = v21
                                        _ = v31
                                        _ = v47
                                    }
                                    v48 := GenS1{F0: int16(p0), F1: -128}
                                    _ = v29
                                    _ = v48
                                } else {
                                    v49 := GenS3{F0: 65535}
                                    _ = v49
                                }
                                _ = fi27
                                _ = v28
                            }
                        case 32767:
                            v50 := GenS4{F0: ((v10 >= v10) && (0 >= -1))}
                            _ = v50
                        }
                        if (0.1 != (3.14159 + -2.25)) {
                            var v51 float64 = float64((v10 - 3.14159))
                            fmt.Println(v23, v10, fi22)
                            _ = v51
                        } else {
                            p2 = (p2 + p2)
                            var v52 uint = 42
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 6; fi53++ {
                            if v23 {
                                var v54 bool = false
                                fmt.Println(p2, p2)
                                _ = v54
                            }
                            fmt.Println(v24, v8, fi53)
                            _ = fi53
                        }
                    }
                case "hello":
                    switch 32767 {
                    case 2147483647:
                        for fi55 := 0; fi55 < 12; fi55++ {
                            p0 += 2
                            switch ("" + v8) {
                            case "\n\t":
                                var v56 uint8 = (p0 * 42)
                                _ = v56
                            default:
                                if v23 {
                                    switch (fi55 - fi55) {
                                    case 9223372036854775807:
                                        v8 = "a"
                                    }
                                    v57 := GenS5{F0: v21}
                                    v58 := GenS3{F0: 65535}
                                    if v23 {
                                        switch v8 {
                                        case "hello":
                                            var v59 string = ((v21 + ((("\n\t" + v21) + "foo bar") + v21)) + (v21 + (v21 + (((v8 + "\n\t") + (v21 + "\n\t")) + "a"))))
                                            _ = v59
                                        case "":
                                            v60 := []int{fi22, (-(fi55)), (fi55 / ((-1) | 1)), (fi55 - 127)}
                                            _ = v60
                                        }
                                    }
                                    _ = v57
                                    _ = v58
                                } else {
                                    v61 := GenS6{F0: 255, F1: 1, F2: (p2 >> ((255) & 63))}
                                    var v62 bool = v23
                                    _ = v61
                                    _ = v62
                                }
                            }
                            _ = fi55
                        }
                    case -935:
                        v63 := []int{}
                        _ = v63
                    case 127:
                        var v64 int16 = -1
                        _ = v64
                    default:
                        var v65 uint64 = 42
                        _ = v65
                    }
                case "foo bar":
                    var v66 uint = 1
                    _ = v66
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            for fi67 := 0; fi67 < 14; fi67++ {
                v21 = (v21 + v21)
                for fi68 := 0; fi68 < 2; fi68++ {
                    if true {
                        if (!true) {
                            if ((v10 / (v10 - (v10 + 0.0))) > ((float32((((-2.25 - 0.1) * 0.1) + 3.14159)) / 0.1) / ((((0.0 + (float32(((1.5 + float64(float32(float64(v10)))) + float64(float32(1e10)))) + (0.0 / (v10 * -2.25)))) / v10) + (-2.25 * v10)) / float32((3.14159 * -2.25))))) {
                                var v69 float32 = v10
                                if ((-(fi67)) != 127) {
                                    if true {
                                        var v70 float32 = (0.1 / (0.1 * v10))
                                        _ = v70
                                    } else {
                                        v10 -= v10
                                        var v71 float32 = 0.1
                                        _ = v71
                                    }
                                    if ((float64(v10) / ((3.14159 - (((0.1 * 3.14159) + 1.5) * float64(v10))) / float64(((0.0 / v10) / float32(3.14159))))) <= (0.0 - -2.25)) {
                                        fmt.Println(fi67, v8)
                                        if (!((p2 > p2) && false)) {
                                            v8 = (v21 + ("\n\t" + v8))
                                            v72 := []int{fi68, (fi68 & int(p1)), fi68, (-(fi68))}
                                            _ = v72
                                        } else {
                                            v73 := []int{}
                                            v69 = ((v69 - v69) - (float32(-2.25) + 0.0))
                                            _ = v73
                                        }
                                        v8 = (("a" + "\n\t") + "a")
                                    }
                                } else {
                                    v74 := GenS3{F0: 140}
                                    v8 = v21
                                    _ = v74
                                }
                                _ = v69
                            } else {
                                v75 := []int{fi68, fi67, 127, -27}
                                _ = v75
                            }
                        }
                        p2 = (p2 % ((2) | 1))
                    }
                    fmt.Println(p0, fi67)
                    _ = fi68
                }
                _ = fi67
            }
            switch ("" + "foo bar") {
            case "foo bar":
                v21 = "a"
            default:
                fmt.Println(p2)
            }
        }
        if true {
            var v76 uint64 = 1
            _ = v76
        } else {
            if ((((v21 + "") + v21) >= "a") || (v10 != -2.25)) {
                var v77 float32 = v10
                var v78 bool = false
                _ = v77
                _ = v78
            } else {
                var v79 uint16 = 585
                var v80 float64 = 1.5
                _ = v79
                _ = v80
            }
            p2 = (p2 >> ((1) & 63))
            fmt.Println(p0, p1)
        }
        _ = v21
    } else {
        switch -528 {
        case 42:
            v81 := []int{-1}
            _ = v81
        case -497:
            v8 = "foo bar"
        case 32768:
            for fi82 := 0; fi82 < 1; fi82++ {
                var v83 float32 = float32(1e10)
                _ = fi82
                _ = v83
            }
        }
        v84 := GenS7{F0: 255, F1: 503}
        p0 -= (p0 ^ (^(p0)))
        _ = v84
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    return 2
}

func main() {
    var v1 int = -2147483648
    var v2 float64 = (-2.25 * (1.5 - 1e10))
    var v3 int32 = -363
    var v4 uint8 = 127
    var v5 int16 = -32768
    v85 := GenS1{F0: -1, F1: 42}
    switch int(v5) {
    case -415:
        v86 := []int{-525, 771, (v1 * -2147483648)}
        _ = v86
    }
    v2 = (1.5 - (1.5 * v2))
    switch v1 {
    case 164:
        fmt.Println(v3, v4)
    case 2147483647:
        if ((true && false) && true) {
            for fi87 := 0; fi87 < 14; fi87++ {
                var v88 int = (fi87 / ((-113) | 1))
                var v89 uint16 = 852
                switch "foo bar" {
                case "foo bar":
                    v90 := []int{(v88 % ((127) | 1)), v1, (v88 % ((int(v3)) | 1))}
                    _ = v90
                case "a":
                    var v91 bool = false
                    _ = v91
                case "":
                    if ((1e10 - (3.14159 * (1.5 + (float32(3.14159) * (3.14159 * -2.25))))) < float32((v2 * v2))) {
                        var v92 int8 = 47
                        v4 = (^(v4))
                        v93 := []int{fi87, (fi87 ^ (fi87 << ((65535) & 63))), 141}
                        _ = v92
                        _ = v93
                    } else {
                        if false {
                            if false {
                                v4 -= v4
                            } else {
                                for fi94 := 0; fi94 < 16; fi94++ {
                                    v5 = v5
                                    v3 -= 0
                                    v4 = (^(v4))
                                    _ = fi94
                                }
                                switch v88 {
                                case 32767:
                                    if true {
                                        v5 = v5
                                        var v95 uint = 15
                                        _ = v95
                                    } else {
                                        v1 -= v88
                                    }
                                case -397:
                                    for fi96 := 0; fi96 < 16; fi96++ {
                                        v1 += (v1 | (v1 * (v88 % ((0) | 1))))
                                        v97 := []int{v1, 9223372036854775807, (v1 / ((fi87) | 1))}
                                        if true {
                                            if (!("" < ("hello" + ((("\n\t" + "a") + ("foo bar" + "a")) + "\n\t")))) {
                                                switch -1 {
                                                case 4294967295:
                                                    fmt.Println(v4)
                                                }
                                                switch v1 {
                                                case 42:
                                                    if true {
                                                        v89 = v89
                                                    }
                                                case 4294967295:
                                                    v1 = (v1 / ((v88) | 1))
                                                }
                                            } else {
                                                if (v4 != (v4 << ((163) & 63))) {
                                                    if false {
                                                        if (-2.25 >= 3.14159) {
                                                            if true {
                                                                var v98 int32 = -2147483648
                                                                var v99 uint = 764
                                                                var v100 uint64 = 42
                                                                _ = v100
                                                                _ = v98
                                                                _ = v99
                                                            } else {
                                                                v5 += (v5 * (^(v5)))
                                                                v101 := GenS2{F0: v3, F1: (^(v89))}
                                                                v4 = 178
                                                                _ = v101
                                                            }
                                                            var v102 string = ((("" + ("a" + "")) + "\n\t") + "")
                                                            _ = v102
                                                        }
                                                    } else {
                                                        v3 = -844
                                                        fmt.Println(v3, fi87)
                                                        var v103 int8 = -120
                                                        _ = v103
                                                    }
                                                    var v104 uint = 1
                                                    var v105 float32 = float32((v2 - 1.5))
                                                    _ = v104
                                                    _ = v105
                                                }
                                                v88 = (v1 - (v88 % ((1) | 1)))
                                                v2 -= (v2 / v2)
                                            }
                                            v106 := GenS1{F0: (v5 % ((0) | 1)), F1: -124}
                                            v89 = 255
                                            _ = v106
                                        }
                                        _ = fi96
                                        _ = v97
                                    }
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    var v107 uint8 = (v4 - v4)
                                    _ = v107
                                }
                            }
                            var v108 int = (fi87 ^ (fi87 % ((1) | 1)))
                            switch 128 {
                            case 9223372036854775807:
                                v109 := []int{(v1 - -848)}
                                _ = v109
                            case -108:
                                v110 := []int{}
                                _ = v110
                            default:
                                fmt.Println(v89, v88, v3)
                            }
                            _ = v108
                        }
                    }
                default:
                    switch ("a" + "a") {
                    case "a":
                        v111 := []int{fi87, int(v4), (v1 % (((-(v1))) | 1)), -1}
                        _ = v111
                    case "":
                        var v112 uint64 = 1
                        _ = v112
                    default:
                        v113 := []int{32767}
                        _ = v113
                    }
                }
                _ = fi87
                _ = v88
                _ = v89
            }
            for fi114 := 0; fi114 < 6; fi114++ {
                v115 := GenS8{F0: genFn1(v4, 42, 149)}
                for fi116 := 0; fi116 < 14; fi116++ {
                    if true {
                        v117 := []int{fi114, fi116}
                        _ = v117
                    } else {
                        v4 = v4
                        v2 += v2
                    }
                    _ = fi116
                }
                fmt.Println(v1)
                _ = fi114
                _ = v115
            }
        }
    }
    for fi118 := 0; fi118 < 4; fi118++ {
        v2 = v2
        v119 := GenS2{F0: v3, F1: 781}
        v4 = 238
        _ = fi118
        _ = v119
    }
    v1 += (^(v1))
    v120 := []int{(v1 + (v1 ^ v1)), (v1 & v1)}
    if true {
        for fi121 := 0; fi121 < 1; fi121++ {
            if (2147483648 == 575) {
                for fi122 := 0; fi122 < 18; fi122++ {
                    if (0 > 0) {
                        var v123 uint32 = genFn1((v4 >> ((2) & 63)), -9223372036854775808, 839)
                        if true {
                            switch (("\n\t" + "\n\t") + "\n\t") {
                            case "":
                                if (!false) {
                                    fmt.Println(v2, fi122)
                                    if true {
                                        switch (("hello" + "foo bar") + ("foo bar" + "foo bar")) {
                                        case "a":
                                            var v124 float64 = 0.1
                                            _ = v124
                                        case "":
                                            fmt.Println(v123, v4)
                                        case "hello":
                                            v1 = (fi122 ^ v1)
                                            switch (fi122 | v1) {
                                            case 109:
                                                switch ((((("hello" + "foo bar") + ("a" + "\n\t")) + "hello") + "") + "foo bar") {
                                                case "a":
                                                    var v125 int8 = -105
                                                    _ = v125
                                                case "":
                                                    switch v1 {
                                                    case 0:
                                                        fmt.Println(v2, fi121)
                                                    case 1:
                                                        var v126 int = fi121
                                                        _ = v126
                                                    case 917:
                                                        var v127 string = "foo bar"
                                                        _ = v127
                                                    default:
                                                        var v128 int8 = 100
                                                        _ = v128
                                                    }
                                                case "foo bar":
                                                    v129 := GenS9{F0: (v4 | v4), F1: 127}
                                                    _ = v129
                                                default:
                                                    var v130 uint = 255
                                                    _ = v130
                                                }
                                            default:
                                                if len(v120) >= 2 {
                                                    v120 = v120[1:len(v120)]
                                                }
                                            }
                                            v123 -= (v123 + v123)
                                        }
                                        if (0 >= 100) {
                                            if false {
                                                if true {
                                                    var v131 int8 = 100
                                                    _ = v131
                                                }
                                                if (!(-9223372036854775808 != 9223372036854775807)) {
                                                    var v132 uint32 = (^(v123))
                                                    var v133 uint32 = 798
                                                    _ = v132
                                                    _ = v133
                                                } else {
                                                    if (false || true) {
                                                        if (true || (("hello" != "a") || true)) {
                                                            if false {
                                                                var v134 uint = 65535
                                                                _ = v134
                                                            }
                                                            var v135 string = ""
                                                            _ = v135
                                                        }
                                                        var v136 float32 = 1e10
                                                        v123 = genFn1((v4 ^ 0), 42, 42)
                                                        v123 = (^(v123))
                                                        v2 += 1.5
                                                        _ = v136
                                                    }
                                                    fmt.Println(fi122, v1, fi122)
                                                }
                                            }
                                            v120 = append(v120, fi121)
                                        } else {
                                            v4 -= (^(v4))
                                            if len(v120) > 0 {
                                                v120[0] = v1
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi122, v123, fi121)
                                        var v137 uint16 = 1
                                        fmt.Println(v5, v4)
                                        var v138 string = ("hello" + "")
                                        fmt.Println(v137)
                                        v139 := GenS2{F0: (v3 >> ((609) & 63)), F1: v137}
                                        _ = v137
                                        _ = v138
                                        _ = v139
                                    }
                                    v1 = 128
                                }
                                fmt.Println(fi122)
                            case "foo bar":
                                v3 -= (v3 / ((v3) | 1))
                            }
                            if false {
                                v3 -= -2147483648
                                if len(v120) > 0 {
                                    v120[0] = (fi121 + v1)
                                }
                                if len(v120) >= 2 {
                                    v120 = v120[1:len(v120)]
                                }
                            } else {
                                if false {
                                    var v140 float64 = (v2 - ((3.14159 + ((1e10 * (float64((3.14159 * (1.5 - 0.0))) - (float64(float32((v2 - v2))) + v2))) / v2)) / 1.0))
                                    fmt.Println(v140)
                                    _ = v140
                                }
                                v5 = v5
                                v1 = 42
                                if true {
                                    v120 = append(v120, int(v123))
                                    if true {
                                        v3 = v3
                                    } else {
                                        switch "" {
                                        case "a":
                                            if (8 < 127) {
                                                if ((genFn1(42, int64(v123), 42) != (v123 + v123)) || (v4 < v4)) {
                                                    if false {
                                                        if false {
                                                            var v141 float64 = (v2 + float64((1.5 * 0.1)))
                                                            var v142 uint64 = 906
                                                            _ = v141
                                                            _ = v142
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v143 bool = ((!true) && (!true))
                                                            _ = v143
                                                        case "a":
                                                            fmt.Println(fi122)
                                                        }
                                                    }
                                                    fmt.Println(fi121, fi122)
                                                    var v144 int = -89
                                                    _ = v144
                                                }
                                            } else {
                                                var v145 int8 = -128
                                                if (fi122 < (fi122 << ((1) & 63))) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v146 int = (v1 % ((fi122) | 1))
                                                        _ = v146
                                                    case "a":
                                                        v120 = append(v120, 2147483647)
                                                    default:
                                                        v123 = v123
                                                    }
                                                    var v147 string = ""
                                                    var v148 string = ""
                                                    _ = v147
                                                    _ = v148
                                                } else {
                                                    v3 -= (v3 ^ v3)
                                                    v4 = (v4 + v4)
                                                    if false {
                                                        fmt.Println(v2, v145, v4)
                                                    }
                                                }
                                                _ = v145
                                            }
                                        default:
                                            fmt.Println(v4, v3)
                                        }
                                    }
                                    v3 = v3
                                } else {
                                    fmt.Println(v3, v5, v4)
                                    var v149 int = (v1 >> ((748) & 63))
                                    var v150 uint = 4294967295
                                    _ = v149
                                    _ = v150
                                }
                            }
                        } else {
                            v151 := GenS10{F0: uint8(v123), F1: (v4 >> ((74) & 63)), F2: true}
                            v5 = 1000
                            _ = v151
                        }
                        var v152 uint32 = (v123 & (v123 % ((2147483648) | 1)))
                        _ = v123
                        _ = v152
                    } else {
                        v153 := GenS11{F0: 1, F1: 55}
                        _ = v153
                    }
                    _ = fi122
                }
                var v154 int32 = v3
                _ = v154
            } else {
                v155 := GenS12{F0: (!(true || ((0 < 4294967295) && true))), F1: (v4 == (v4 << ((2) & 63)))}
                v4 = v4
                _ = v155
            }
            for fi156 := 0; fi156 < 11; fi156++ {
                v157 := GenS13{F0: (v3 ^ v3), F1: v3, F2: 255}
                _ = fi156
                _ = v157
            }
            var v158 uint16 = 581
            _ = fi121
            _ = v158
        }
        var v159 int32 = 2147483647
        fmt.Println(v5, v2, v1)
        _ = v159
    }
    var v160 uint32 = 4294967295
    var v161 int = (v1 % ((v1) | 1))
    fmt.Println(v3, v5)
    v162 := GenS3{F0: 65535}
    fmt.Println(v2)
    v163 := GenS12{F0: (!((v3 % ((int32(v160)) | 1)) == 0)), F1: true}
    if true {
        var v164 int8 = 0
        _ = v164
    }
    _ = len(v120)
    if true {
        if len(v120) > 0 {
            v120[0] = int(v5)
        }
    } else {
        switch ("" + (("a" + "foo bar") + "foo bar")) {
        case "foo bar":
            switch "a" {
            case "foo bar":
                fmt.Println(v2)
            case "\n\t":
                v165 := GenS11{F0: uint(v160), F1: 964}
                _ = v165
            default:
                var v166 bool = false
                _ = v166
            }
        case "":
            if (false && ((-(v161)) <= (v1 * (v161 | (v1 * (v1 | 9223372036854775807)))))) {
                v167 := GenS14{F0: 255, F1: 375, F2: 0}
                v168 := GenS4{F0: (!("hello" != "a"))}
                var v169 bool = true
                _ = v167
                _ = v168
                _ = v169
            }
        case "hello":
            var v170 int16 = int16(v4)
            _ = v170
        }
    }
    var v171 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v120:", len(v120), v120)
    fmt.Println("v160:", v160)
    fmt.Println("v161:", v161)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v120
    _ = v160
    _ = v161
    _ = v162
    _ = v163
    _ = v171
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
