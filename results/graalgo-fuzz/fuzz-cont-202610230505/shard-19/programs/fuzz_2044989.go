package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 string; F1 bool; F2 uint8 }
type GenS12 struct { F0 uint8; F1 int; F2 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int; F1 uint64; F2 uint64 }
type GenS15 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 uint16 = 32767
    var v4 int16 = 42
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(v3, v3)
        _ = fi5
    }
    fmt.Println(v3)
    var v6 float32 = 1e10
    switch 765 {
    case -834:
        v7 := []int{0, 12}
        _ = v7
    case 128:
        fmt.Println(v4, v6, v4)
    case -378:
        var v8 uint8 = 128
        _ = v8
    }
    var v9 uint = 255
    v10 := GenS1{F0: (v4 >> (((v9 % (((v9 >> ((65535) & 63))) | 1))) & 63))}
    switch "\n\t" {
    case "a":
        var v11 uint16 = 539
        _ = v11
    case "\n\t":
        switch -107 {
        case 128:
            var v12 int = -921
            _ = v12
        case 32767:
            v13 := GenS1{F0: (v4 >> ((v9) & 63))}
            _ = v13
        case -77:
            v14 := GenS2{F0: (^(v3))}
            _ = v14
        default:
            for fi15 := 0; fi15 < 20; fi15++ {
                v3 = v3
                _ = fi15
            }
        }
    default:
        switch ("\n\t" + "") {
        case "foo bar":
            v4 = int16(v9)
        case "\n\t":
            v4 += (-(v4))
        case "a":
            v16 := GenS1{F0: -1}
            _ = v16
        }
    }
    if false {
        v9 = (v9 ^ (v9 / (((v9 & v9)) | 1)))
        var v17 int = 42
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        if (!true) {
            v9 = v9
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                if ((v4 & (v4 >> ((v9) & 63))) >= -1) {
                    v9 += v9
                    v6 += float32(((1.5 + 1e10) + (1.5 - -2.25)))
                } else {
                    if (!false) {
                        var v20 uint = uint(v3)
                        if (105 <= 1) {
                            v6 = float32(float64((v6 * v6)))
                            var v21 int64 = 9223372036854775807
                            v21 -= (v21 % (((v21 ^ v21)) | 1))
                            _ = v21
                        } else {
                            v22 := GenS3{F0: 70}
                            var v23 int16 = (^(v4))
                            _ = v22
                            _ = v23
                        }
                        v4 += v4
                        v3 = 2
                        _ = v20
                    }
                    v1[2] = "hello"
                    var v24 int8 = 127
                    _ = v24
                }
                _ = fi19
            }
            if false {
                if (9223372036854775807 != -9223372036854775808) {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v4 = v4
                        _ = fi25
                    }
                    if true {
                        v6 = v6
                        if false {
                            for fi26 := 0; fi26 < 3; fi26++ {
                                var v27 int32 = 0
                                v28 := GenS3{F0: uint(v3)}
                                if ((v4 % ((-32768) | 1)) == v4) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if false {
                                            var v29 float32 = -2.25
                                            switch (((("hello" + "") + "a") + "hello") + ("hello" + "a")) {
                                            case "hello":
                                                if false {
                                                    switch (fi18 | (fi18 % ((fi18) | 1))) {
                                                    case 0:
                                                        fmt.Println(v9)
                                                    case 361:
                                                        v3 -= 0
                                                    default:
                                                        v30 := []int{(fi18 + (-(fi26))), (fi26 / ((fi18) | 1))}
                                                        _ = v30
                                                    }
                                                } else {
                                                    if (70 == 100) {
                                                        fmt.Println(v6, fi18)
                                                        var v31 int64 = 10
                                                        var v32 uint8 = uint8(v9)
                                                        if false {
                                                            v33 := v1[0]
                                                            v34 := []int{(fi26 / ((-742) | 1)), (^(fi26)), (fi26 + fi26), (fi18 + 42)}
                                                            switch "hello" {
                                                            case "":
                                                                switch v33 {
                                                                case "foo bar":
                                                                    if false {
                                                                        var v35 float32 = v29
                                                                        v36 := GenS1{F0: (v4 + (-(v4)))}
                                                                        switch v33 {
                                                                        case "":
                                                                            fmt.Println(v4)
                                                                        }
                                                                        _ = v35
                                                                        _ = v36
                                                                    } else {
                                                                        v37 := GenS2{F0: 255}
                                                                        var v38 int8 = 76
                                                                        _ = v37
                                                                        _ = v38
                                                                    }
                                                                    if ((v32 - (v32 & v32)) <= (v32 - v32)) {
                                                                        if true {
                                                                            var v39 uint8 = 134
                                                                            var v40 int = -2147483648
                                                                            _ = v39
                                                                            _ = v40
                                                                        }
                                                                        if (true && (87 > 127)) {
                                                                            var v41 bool = false
                                                                            fmt.Println(fi18, v31, v31)
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        if (v4 >= v4) {
                                                                            v42 := GenS1{F0: int16(v31)}
                                                                            var v43 uint32 = 19
                                                                            v31 -= v31
                                                                            if (false || (true && (true || true))) {
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v31 += 0
                                                                                case "foo bar":
                                                                                    v29 = float32(0.1)
                                                                                case "":
                                                                                    var v44 float64 = (1.5 - -2.25)
                                                                                    _ = v44
                                                                                }
                                                                                if true {
                                                                                    v29 = (float32(0.1) + v29)
                                                                                } else {
                                                                                    if false {
                                                                                        var v45 int = fi18
                                                                                        v46 := GenS4{F0: uint8(v9), F1: uint(v3)}
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                    } else {
                                                                                        var v47 int8 = -37
                                                                                        _ = v47
                                                                                    }
                                                                                    if false {
                                                                                        var v48 float64 = float64(v6)
                                                                                        if (v31 == 9223372036854775807) {
                                                                                            var v49 uint = v9
                                                                                            _ = v49
                                                                                        }
                                                                                        switch fi18 {
                                                                                        case 2147483647:
                                                                                            fmt.Println(v33, v6, v4)
                                                                                        case 69:
                                                                                            var v50 uint = v9
                                                                                            _ = v50
                                                                                        }
                                                                                        var v51 float32 = v6
                                                                                        v52 := GenS5{F0: fi26}
                                                                                        _ = v48
                                                                                        _ = v51
                                                                                        _ = v52
                                                                                    } else {
                                                                                        v53 := GenS2{F0: 255}
                                                                                        v43 += 1
                                                                                        _ = v53
                                                                                    }
                                                                                }
                                                                                if (!(((v3 >> ((v9) & 63)) <= 255) || (v4 >= 594))) {
                                                                                    if (v31 <= v31) {
                                                                                        v54 := GenS5{F0: fi26}
                                                                                        delete(v2, (v43 | (v43 / ((uint32(v9)) | 1))))
                                                                                        var v55 float64 = ((((0.0 * -2.25) / (3.14159 + 3.14159)) * 3.14159) / 3.14159)
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    }
                                                                                    switch (v33 + "foo bar") {
                                                                                    case "\n\t":
                                                                                        v56 := GenS6{F0: v3, F1: v27}
                                                                                        _ = v56
                                                                                    case "":
                                                                                        v57 := GenS2{F0: (v3 ^ 64)}
                                                                                        _ = v57
                                                                                    }
                                                                                    if false {
                                                                                        _ = len(v34)
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    v29 = float32(3.14159)
                                                                                    var v58 bool = true
                                                                                    v59 := GenS4{F0: (v32 - uint8(v43)), F1: v9}
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                } else {
                                                                                    fmt.Println(v31)
                                                                                    var v60 float32 = v6
                                                                                    var v61 int32 = v27
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                                fmt.Println(v27, v9, v33)
                                                                            }
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                        var v62 int64 = v31
                                                                        _ = v62
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v63 := GenS7{F0: ((fi18 * fi26) != (-(fi26)))}
                                                                        _ = v63
                                                                    case "a":
                                                                        var v64 int32 = v27
                                                                        _ = v64
                                                                    }
                                                                default:
                                                                    var v65 string = ((v33 + v33) + "foo bar")
                                                                    _ = v65
                                                                }
                                                            case "a":
                                                                var v66 uint64 = 42
                                                                _ = v66
                                                            case "foo bar":
                                                                v2[2147483648] = ((v33 + ("foo bar" + ("\n\t" + (v33 + v33)))) + v33)
                                                            default:
                                                                var v67 int32 = (v27 - v27)
                                                                _ = v67
                                                            }
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        v68 := GenS8{F0: 63, F1: v27}
                                                        _ = v31
                                                        _ = v32
                                                        _ = v68
                                                    } else {
                                                        _ = len(v2)
                                                        var v69 uint8 = 42
                                                        if ((float32(3.14159) + float32((0.0 - 1e10))) <= v6) {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if false {
                                                                    if ((1 < -1) || ((v4 - v4) <= (v4 - (v4 + (v4 / ((v4) | 1)))))) {
                                                                        var v70 int32 = v27
                                                                        var v71 bool = true
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        switch (fi18 << ((uint(v69)) & 63)) {
                                                                        case 127:
                                                                            v69 += v69
                                                                        case -778:
                                                                            v72 := []int{-390}
                                                                            _ = v72
                                                                        case -1:
                                                                            v73 := v1[65535]
                                                                            _ = v73
                                                                        }
                                                                        var v74 int16 = 1
                                                                        v9 -= v9
                                                                        _ = v74
                                                                    }
                                                                } else {
                                                                    v75 := GenS9{F0: 0, F1: 2147483647}
                                                                    fmt.Println(v9, fi18, v3)
                                                                    if true {
                                                                        var v76 float64 = (((((float64((v6 * 0.1)) / ((-2.25 - -2.25) + float64((1e10 + v29)))) * 0.1) - ((0.0 * -2.25) - 1.5)) - 1.5) - 0.0)
                                                                        if false {
                                                                            v9 = (v9 + v9)
                                                                            v2[603] = "\n\t"
                                                                            v77 := GenS2{F0: (v3 - 820)}
                                                                            _ = v77
                                                                        } else {
                                                                            if false {
                                                                                var v78 float32 = (-2.25 * (0.0 / 0.1))
                                                                                _ = v78
                                                                            } else {
                                                                                switch (fi26 - fi26) {
                                                                                case 32768:
                                                                                    var v79 int32 = (^(v27))
                                                                                    _ = v79
                                                                                default:
                                                                                    var v80 uint64 = 1
                                                                                    _ = v80
                                                                                }
                                                                                var v81 uint = v9
                                                                                v29 += v6
                                                                                var v82 uint8 = (v69 * 255)
                                                                                fmt.Println(v76, v27)
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        _ = v76
                                                                    } else {
                                                                        var v83 int8 = -128
                                                                        _ = v83
                                                                    }
                                                                    _ = v75
                                                                }
                                                            case "\n\t":
                                                                v6 += v29
                                                            case "a":
                                                                var v84 float32 = v29
                                                                _ = v84
                                                            default:
                                                                var v85 int32 = v27
                                                                _ = v85
                                                            }
                                                            fmt.Println(v29, v3)
                                                            switch fi26 {
                                                            case 128:
                                                                v86 := GenS6{F0: v3, F1: v27}
                                                                _ = v86
                                                            case -707:
                                                                v87 := []int{860, 127}
                                                                _ = v87
                                                            default:
                                                                if (-74 > 11) {
                                                                    v27 += (v27 - v27)
                                                                }
                                                            }
                                                            v4 = v4
                                                        } else {
                                                            if false {
                                                                v88 := []int{}
                                                                var v89 int32 = v27
                                                                var v90 uint32 = 1
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                v91 := GenS3{F0: v9}
                                                                _ = v91
                                                            }
                                                        }
                                                        _ = v69
                                                    }
                                                    switch "a" {
                                                    case "foo bar":
                                                        v27 = v27
                                                    }
                                                    v9 += (v9 << (((v9 << ((4294967295) & 63))) & 63))
                                                }
                                            case "a":
                                                switch fi26 {
                                                case 127:
                                                    if (127 >= (fi26 * (fi18 - fi18))) {
                                                        v92 := GenS9{F0: -128, F1: 4294967295}
                                                        if ("hello" == ("" + "")) {
                                                            v93 := []int{int(v9), int(v4), 4294967295, (-(fi18))}
                                                            fmt.Println(fi18)
                                                            v94 := GenS10{F0: 2147483647, F1: true}
                                                            var v95 int64 = 0
                                                            _ = v93
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            v4 += v4
                                                        }
                                                        _ = v92
                                                    }
                                                default:
                                                    var v96 int32 = (v27 & v27)
                                                    _ = v96
                                                }
                                            case "\n\t":
                                                var v97 int32 = v27
                                                _ = v97
                                            default:
                                                v27 = 2147483647
                                            }
                                            _ = len(v1)
                                            var v98 uint8 = 128
                                            _ = v29
                                            _ = v98
                                        } else {
                                            var v99 int32 = v27
                                            v100 := v1[2147483648]
                                            _ = v100
                                            _ = v99
                                        }
                                    case "hello":
                                        if (!(0 <= 9223372036854775807)) {
                                            switch fi18 {
                                            case -608:
                                                var v101 string = "a"
                                                _ = v101
                                            case 9223372036854775807:
                                                switch "a" {
                                                case "a":
                                                    v102 := []int{fi26, fi26, (fi26 | fi26)}
                                                    _ = v102
                                                case "\n\t":
                                                    v103 := []int{(fi26 % ((0) | 1))}
                                                    _ = v103
                                                default:
                                                    v4 += (^(v4))
                                                }
                                            default:
                                                var v104 int64 = 937
                                                _ = v104
                                            }
                                            if (false || true) {
                                                v27 = (v27 & 1)
                                                fmt.Println(v27, v9, v6)
                                            }
                                            var v105 uint16 = 997
                                            _ = v105
                                        }
                                    case "":
                                        v9 += (v9 * (v9 | 65535))
                                    }
                                    v27 = v27
                                    v106 := GenS11{F0: "a", F1: (false && (!(143 >= 396))), F2: 152}
                                    _ = v106
                                }
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            for fi107 := 0; fi107 < 5; fi107++ {
                                delete(v2, 1)
                                _ = fi107
                            }
                        }
                    } else {
                        var v108 int8 = int8(v9)
                        _ = v108
                    }
                } else {
                    v1[2147483647] = "hello"
                    for fi109 := 0; fi109 < 4; fi109++ {
                        v110 := GenS3{F0: (v9 & v9)}
                        v111 := []int{-72}
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                }
            }
        }
        _ = fi18
    }
    switch 51 {
    case 42:
        var v112 int8 = 42
        _ = v112
    case 32767:
        fmt.Println(v6, v3)
    default:
        if (-2.25 >= v6) {
            if false {
                switch ("foo bar" + ("a" + "foo bar")) {
                case "\n\t":
                    for fi113 := 0; fi113 < 2; fi113++ {
                        var v114 int = fi113
                        switch (fi113 >> ((255) & 63)) {
                        case -799:
                            v3 = (v3 >> ((v9) & 63))
                        case -1:
                            if ((true || false) || ((!false) && (true && (false || (((-128 == 62) && (!false)) || (2 >= 42)))))) {
                                for fi115 := 0; fi115 < 14; fi115++ {
                                    v4 += (v4 | 42)
                                    v6 = 0.1
                                    _ = fi115
                                }
                            } else {
                                for fi116 := 0; fi116 < 16; fi116++ {
                                    switch "foo bar" {
                                    case "hello":
                                        if (false || false) {
                                            v3 = v3
                                        }
                                    case "\n\t":
                                        fmt.Println(fi116, fi116)
                                    case "a":
                                        switch fi113 {
                                        case 32767:
                                            v117 := GenS12{F0: 127, F1: (fi116 << (((v9 * (v9 * (^(v9))))) & 63)), F2: 42}
                                            _ = v117
                                        case 4294967295:
                                            var v118 int8 = 42
                                            _ = v118
                                        }
                                    }
                                    _ = fi116
                                }
                                var v119 int8 = 12
                                _ = v119
                            }
                        }
                        _ = fi113
                        _ = v114
                    }
                case "":
                    var v120 float32 = 1.5
                    _ = v120
                case "foo bar":
                    v6 -= 1.5
                default:
                    if ((-32768 <= v4) && false) {
                        fmt.Println(v9, v6, v6)
                    }
                }
                switch -9223372036854775808 {
                case 0:
                    var v121 uint = (v9 * 532)
                    _ = v121
                }
                v3 = v3
            } else {
                fmt.Println(v9)
                v4 = (v4 + (v4 & 1))
                for fi122 := 0; fi122 < 9; fi122++ {
                    v4 -= (v4 - 32767)
                    if false {
                        v123 := []int{(fi122 + (^(fi122))), 127, fi122}
                        fmt.Println(v9)
                        _ = v123
                    }
                    _ = fi122
                }
            }
            var v124 int32 = 0
            _ = v124
        }
    }
    if (2147483648 < 2) {
        v125 := GenS8{F0: 0, F1: 1}
        v126 := v2[42]
        _ = v125
        _ = v126
    } else {
        var v127 uint64 = 0
        v128 := []int{-1, 1}
        if (true || (!true)) {
            var v129 int32 = 1
            var v130 uint64 = v127
            _ = v129
            _ = v130
        }
        _ = v127
        _ = v128
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            var v131 uint32 = 2147483648
            _ = v131
        case "foo bar":
            for fi132 := 0; fi132 < 8; fi132++ {
                fmt.Println(v6, fi132)
                for fi133 := 0; fi133 < 8; fi133++ {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v9)
                    default:
                        fmt.Println(v9)
                    }
                    v4 = (-(v4))
                    var v134 float64 = ((((3.14159 - (float64(1.5) / -2.25)) * ((0.0 / 1.5) - 0.0)) * (0.1 * 0.1)) / (0.0 - (1.5 - ((1.5 - (((1.5 + float64(float32((((3.14159 * 0.0) - -2.25) + 1.5)))) + (1.5 - ((1.5 / 0.1) + -2.25))) - 1.5)) + 0.0))))
                    var v135 int = fi133
                    _ = fi133
                    _ = v134
                    _ = v135
                }
                _ = fi132
            }
        case "hello":
            if true {
                fmt.Println(v9, v9)
                v6 = v6
                var v136 int64 = -1
                _ = v136
            }
        default:
            switch ("a" + "") {
            case "\n\t":
                v4 -= (-(v4))
            case "foo bar":
                for fi137 := 0; fi137 < 7; fi137++ {
                    for fi138 := 0; fi138 < 17; fi138++ {
                        fmt.Println(fi137)
                        v139 := []int{fi137, fi138}
                        _ = fi138
                        _ = v139
                    }
                    _ = fi137
                }
            case "hello":
                if true {
                    var v140 bool = true
                    v6 = (0.1 / (-2.25 / v6))
                    _ = v140
                } else {
                    if true {
                        switch 42 {
                        case 4294967295:
                            for fi141 := 0; fi141 < 12; fi141++ {
                                for fi142 := 0; fi142 < 4; fi142++ {
                                    if (3.14159 > 1.5) {
                                        var v143 bool = (!true)
                                        _ = v143
                                    }
                                    v6 = v6
                                    v6 = (float32((0.0 - float64(((1.5 / v6) * v6)))) / 3.14159)
                                    _ = fi142
                                }
                                _ = fi141
                            }
                        case -1:
                            switch "hello" {
                            case "foo bar":
                                v4 = v4
                            case "hello":
                                var v144 bool = (32768 < 902)
                                _ = v144
                            default:
                                v9 -= 255
                            }
                        default:
                            if false {
                                if false {
                                    for fi145 := 0; fi145 < 6; fi145++ {
                                        v146 := v2[737]
                                        var v147 float64 = 3.14159
                                        if true {
                                            for fi148 := 0; fi148 < 8; fi148++ {
                                                v6 -= float32(3.14159)
                                                v149 := []int{9223372036854775807, (fi145 / ((747) | 1)), fi148, (fi145 | fi145)}
                                                if len(v149) >= 2 {
                                                    v149 = v149[1:len(v149)]
                                                }
                                                _ = fi148
                                                _ = v149
                                            }
                                            v147 += v147
                                            v2[2] = v146
                                        } else {
                                            if false {
                                                if true {
                                                    for fi150 := 0; fi150 < 11; fi150++ {
                                                        var v151 int16 = 0
                                                        v147 += (v147 * v147)
                                                        var v152 uint8 = 2
                                                        v153 := GenS1{F0: (-(v151))}
                                                        _ = fi150
                                                        _ = v151
                                                        _ = v152
                                                        _ = v153
                                                    }
                                                    if false {
                                                        var v154 int64 = 0
                                                        _ = v154
                                                    }
                                                }
                                                var v155 uint8 = 128
                                                _ = v155
                                            }
                                        }
                                        _ = fi145
                                        _ = v146
                                        _ = v147
                                    }
                                } else {
                                    v6 += v6
                                    var v156 int = -2147483648
                                    _ = v156
                                }
                                v157 := GenS13{F0: 833}
                                _ = v157
                            }
                        }
                        var v158 int = -273
                        _ = v158
                    } else {
                        var v159 float64 = 3.14159
                        for fi160 := 0; fi160 < 2; fi160++ {
                            var v161 uint16 = v3
                            v161 -= v161
                            _ = fi160
                            _ = v161
                        }
                        v4 += v4
                        _ = v159
                    }
                    fmt.Println(v3, v4)
                }
            default:
                switch "" {
                case "hello":
                    for fi162 := 0; fi162 < 19; fi162++ {
                        v163 := GenS14{F0: fi162, F1: 18446744073709551615, F2: 1}
                        for fi164 := 0; fi164 < 1; fi164++ {
                            v165 := GenS11{F0: "\n\t", F1: ((fi162 << ((917) & 63)) <= (fi162 >> (((v9 & (v9 & (v9 - 741)))) & 63))), F2: 2}
                            _ = fi164
                            _ = v165
                        }
                        _ = fi162
                        _ = v163
                    }
                case "foo bar":
                    var v166 uint16 = v3
                    _ = v166
                default:
                    var v167 int = 1
                    _ = v167
                }
            }
        }
    } else {
        v168 := []int{}
        v169 := GenS15{F0: 42, F1: 42}
        _ = v168
        _ = v169
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
