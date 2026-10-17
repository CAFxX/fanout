package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int8; F2 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint; F1 int32; F2 bool }
type GenS15 struct { F0 int16; F1 bool }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint64; F1 uint; F2 string }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint32; F1 bool; F2 uint }
func genFn0() uint8 {
    var v6 uint = 112
    if (0 > 42) {
        if false {
            v6 += 868
            if true {
                for fi7 := 0; fi7 < 4; fi7++ {
                    v8 := []int{fi7}
                    var v9 int16 = 42
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            } else {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := GenS1{F0: 2}
                    _ = fi10
                    _ = v11
                }
                v6 = v6
            }
            if false {
                if ((2 != 1) || true) {
                    if ((!true) || false) {
                        fmt.Println(v6)
                    }
                    v12 := GenS1{F0: 2}
                    _ = v12
                }
                v13 := GenS2{F0: "a", F1: 44, F2: 32767}
                switch "foo bar" {
                case "":
                    switch ("hello" + "hello") {
                    case "\n\t":
                        v14 := []int{}
                        _ = v14
                    case "":
                        v15 := GenS2{F0: "hello", F1: 13, F2: 50}
                        _ = v15
                    case "a":
                        v16 := GenS1{F0: 65535}
                        _ = v16
                    }
                }
                _ = v13
            }
        }
        var v17 string = "a"
        _ = v17
    } else {
        if (62 != 0) {
            if (true && ((v6 | 4294967295) <= v6)) {
                switch (("a" + "") + "\n\t") {
                case "hello":
                    var v18 int32 = 1
                    _ = v18
                case "\n\t":
                    var v19 uint = 255
                    _ = v19
                }
                var v20 float64 = 0.1
                v21 := GenS2{F0: "foo bar", F1: -1, F2: 255}
                _ = v20
                _ = v21
            }
            fmt.Println(v6)
            switch "\n\t" {
            case "\n\t":
                v22 := []int{-987, 1}
                _ = v22
            case "":
                fmt.Println(v6)
            default:
                var v23 int8 = 127
                _ = v23
            }
        } else {
            var v24 float64 = 1.5
            fmt.Println(v6, v24)
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 int8 = 108
                v24 = v24
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        v27 := []int{860, -35, -363}
        _ = v27
    }
    v6 -= v6
    switch "a" {
    case "hello":
        v6 = (v6 << ((903) & 63))
    case "foo bar":
        v6 = v6
    }
    _ = v6
    return 0
}

func main() {
    var v1 int64 = 179
    v2 := []uint32{}
    v3 := make(map[int]bool)
    var v4 int64 = (v1 % ((v1) | 1))
    var v5 float64 = (-2.25 * (0.0 + -2.25))
    var v28 uint16 = 2
    var v29 int64 = (v4 + (v4 & 0))
    for fi30 := 0; fi30 < 6; fi30++ {
        fmt.Println(v28, v5)
        _ = fi30
    }
    v31 := GenS1{F0: 65535}
    switch 128 {
    case -197:
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 float32 = -2.25
            v34 := GenS3{F0: ("hello" + "\n\t")}
            _ = fi32
            _ = v33
            _ = v34
        }
    case -2147483648:
        if ((18446744073709551615 >= 4294967295) || (2 <= 217)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch "" {
            case "":
                switch 0 {
                case -9223372036854775808:
                    var v35 int32 = -1
                    _ = v35
                case 2147483647:
                    if (65535 <= 1) {
                        v4 += 1
                        var v36 float64 = 3.14159
                        _ = v36
                    }
                case 217:
                    for fi37 := 0; fi37 < 10; fi37++ {
                        var v38 bool = (1 > 2147483647)
                        v39 := v3[-688]
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
            case "foo bar":
                v5 += (v5 / 1e10)
            case "\n\t":
                var v40 int32 = int32(v4)
                _ = v40
            }
            v41 := GenS4{F0: 361, F1: 18446744073709551615}
            _ = v41
        } else {
            v42 := GenS1{F0: (v28 & (v28 ^ (v28 - v28)))}
            _ = v42
        }
    default:
        if false {
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
    }
    if (false || true) {
        for fi43 := 0; fi43 < 3; fi43++ {
            v28 += v28
            var v44 uint64 = 0
            v45 := GenS3{F0: ("foo bar" + "foo bar")}
            _ = fi43
            _ = v44
            _ = v45
        }
    } else {
        v5 += v5
        if (("\n\t" + "") <= "foo bar") {
            fmt.Println(v28, v29)
            if true {
                var v46 int16 = 1
                if ((v1 == v4) || true) {
                    v47 := GenS5{F0: (!false), F1: 4294967295}
                    v48 := GenS5{F0: (("" + "a") == ("a" + "\n\t")), F1: 42}
                    _ = v47
                    _ = v48
                } else {
                    v3[42] = false
                    v29 = (-(v4))
                    var v49 int8 = 100
                    _ = v49
                }
                delete(v3, -1)
                _ = v46
            } else {
                var v50 uint16 = v28
                fmt.Println(v1, v29, v4)
                _ = v50
            }
        } else {
            v4 += (v29 - (v29 % ((v4) | 1)))
            var v51 int8 = -128
            if (false && true) {
                for fi52 := 0; fi52 < 9; fi52++ {
                    switch "foo bar" {
                    case "hello":
                        for fi53 := 0; fi53 < 12; fi53++ {
                            var v54 uint16 = (v28 << ((705) & 63))
                            if false {
                                var v55 uint8 = genFn0()
                                v56 := GenS6{F0: 1, F1: v51}
                                v1 = 42
                                _ = v55
                                _ = v56
                            } else {
                                v28 += (v28 & (^(v54)))
                                fmt.Println(v54, v5)
                            }
                            var v57 uint64 = 1
                            _ = fi53
                            _ = v54
                            _ = v57
                        }
                    case "a":
                        if (42 < 1) {
                            if false {
                                switch 900 {
                                case 32767:
                                    for fi58 := 0; fi58 < 14; fi58++ {
                                        var v59 int16 = 1
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                                switch "hello" {
                                case "foo bar":
                                    for fi60 := 0; fi60 < 4; fi60++ {
                                        v4 -= v4
                                        v29 += v29
                                        v29 -= -1
                                        _ = fi60
                                    }
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                case "\n\t":
                                    switch -9223372036854775808 {
                                    case 127:
                                        v2 = append(v2, 2147483647)
                                    }
                                default:
                                    if (-2.25 != 3.14159) {
                                        var v61 int8 = v51
                                        if len(v2) > 0 {
                                            v2[0] = 511
                                        }
                                        _ = v61
                                    } else {
                                        var v62 string = "\n\t"
                                        switch (fi52 << ((255) & 63)) {
                                        case 32767:
                                            v63 := GenS3{F0: "a"}
                                            _ = v63
                                        case 9223372036854775807:
                                            v29 = v4
                                        default:
                                            var v64 uint16 = (v28 % ((v28) | 1))
                                            _ = v64
                                        }
                                        _ = v62
                                    }
                                }
                                v65 := GenS6{F0: 4294967295, F1: v51}
                                _ = v65
                            }
                            if true {
                                delete(v3, (fi52 * -7))
                                if (!(1 <= 18446744073709551615)) {
                                    var v66 uint = 861
                                    for fi67 := 0; fi67 < 19; fi67++ {
                                        v68 := GenS1{F0: v28}
                                        v28 = 0
                                        var v69 uint8 = genFn0()
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                    }
                                    if false {
                                        for fi70 := 0; fi70 < 18; fi70++ {
                                            var v71 uint64 = 42
                                            _ = fi70
                                            _ = v71
                                        }
                                        var v72 int8 = 100
                                        v73 := v3[(fi52 + fi52)]
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v5 += 0.0
                                        v5 = v5
                                        _ = len(v2)
                                    }
                                    _ = v66
                                } else {
                                    fmt.Println(v51, v4, fi52)
                                    v29 += (v29 & v29)
                                }
                                var v74 uint8 = 128
                                _ = v74
                            } else {
                                fmt.Println(fi52, v29)
                            }
                        }
                    case "foo bar":
                        for fi75 := 0; fi75 < 1; fi75++ {
                            var v76 int8 = (v51 % (((v51 * 100)) | 1))
                            v76 -= -128
                            _ = fi75
                            _ = v76
                        }
                    default:
                        var v77 bool = (360 == 18446744073709551615)
                        _ = v77
                    }
                    var v78 int64 = (v29 - (v1 * v4))
                    _ = fi52
                    _ = v78
                }
            }
            _ = v51
        }
        v1 = 0
    }
    v79 := GenS2{F0: "hello", F1: 100, F2: (v28 ^ v28)}
    fmt.Println(v1)
    var v80 int16 = -32768
    if ((v1 >> ((18446744073709551615) & 63)) == (v29 * int64(v80))) {
        if true {
            switch 32768 {
            case 4294967295:
                v81 := GenS7{F0: v1}
                _ = v81
            case 1:
                switch -565 {
                case 9223372036854775807:
                    v28 = (^(v28))
                case 1:
                    if false {
                        if false {
                            v82 := GenS3{F0: "foo bar"}
                            v5 = v5
                            _ = v82
                        } else {
                            if true {
                                fmt.Println(v5, v28)
                                for fi83 := 0; fi83 < 13; fi83++ {
                                    v1 = 1
                                    switch fi83 {
                                    case -920:
                                        if (-1 >= -127) {
                                            var v84 int64 = v4
                                            _ = v84
                                        } else {
                                            if false {
                                                v85 := GenS6{F0: 798, F1: -128}
                                                v86 := GenS3{F0: "hello"}
                                                if true {
                                                    switch 9223372036854775807 {
                                                    case 127:
                                                        if false {
                                                            if (4294967295 < 42) {
                                                                v2 = append(v2, 0)
                                                            }
                                                        }
                                                    default:
                                                        if (281 == (v80 ^ (v80 / ((int16(v28)) | 1)))) {
                                                            for fi87 := 0; fi87 < 2; fi87++ {
                                                                var v88 int16 = (v80 / (((v80 & (v80 / ((42) | 1)))) | 1))
                                                                _ = fi87
                                                                _ = v88
                                                            }
                                                            for fi89 := 0; fi89 < 15; fi89++ {
                                                                v80 = v80
                                                                _ = fi89
                                                            }
                                                            var v90 uint = 65535
                                                            _ = v90
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v4, fi83, fi83)
                                                }
                                                _ = v85
                                                _ = v86
                                            } else {
                                                for fi91 := 0; fi91 < 7; fi91++ {
                                                    v92 := GenS8{F0: 2147483647, F1: 117, F2: 0}
                                                    _ = fi91
                                                    _ = v92
                                                }
                                                delete(v3, fi83)
                                            }
                                            v5 -= float64((3.14159 + float32(-2.25)))
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                    if ((false && ((1.5 - 1.5) != v5)) || ((^(v28)) <= v28)) {
                                        if (2 > 4294967295) {
                                            for fi93 := 0; fi93 < 20; fi93++ {
                                                if false {
                                                    v3[(fi83 | fi83)] = true
                                                    var v94 float32 = -2.25
                                                    fmt.Println(v5, v29)
                                                    _ = v94
                                                } else {
                                                    v95 := GenS3{F0: ((("" + "a") + "foo bar") + (("a" + "foo bar") + (("\n\t" + ("\n\t" + "")) + "\n\t")))}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v96 string = ""
                                                        _ = v96
                                                    case "hello":
                                                        var v97 string = ""
                                                        _ = v97
                                                    }
                                                    _ = v95
                                                }
                                                v98 := GenS9{F0: "hello"}
                                                _ = fi93
                                                _ = v98
                                            }
                                            switch fi83 {
                                            case 9223372036854775807:
                                                v28 += v28
                                            case 127:
                                                if false {
                                                    var v99 int16 = v80
                                                    var v100 bool = false
                                                    _ = v100
                                                    _ = v99
                                                }
                                            }
                                        } else {
                                            var v101 uint16 = (v28 ^ (v28 * v28))
                                            v1 -= (v4 * (^(v1)))
                                            v5 += float64(-2.25)
                                            _ = v101
                                        }
                                    }
                                    _ = fi83
                                }
                            }
                            v80 = (v80 / (((v80 & (^(v80)))) | 1))
                        }
                        var v102 uint8 = 255
                        _ = v102
                    } else {
                        _ = len(v3)
                        for fi103 := 0; fi103 < 1; fi103++ {
                            var v104 float64 = v5
                            delete(v3, fi103)
                            _ = fi103
                            _ = v104
                        }
                        v80 = (v80 + v80)
                    }
                case -9223372036854775808:
                    if false {
                        for fi105 := 0; fi105 < 3; fi105++ {
                            fmt.Println(fi105)
                            v106 := GenS10{F0: v80, F1: fi105}
                            if (758 < 42) {
                                if (216 <= 18446744073709551615) {
                                    switch ((("hello" + "a") + "hello") + (("\n\t" + "a") + ("a" + ("a" + "\n\t")))) {
                                    case "":
                                        v107 := GenS11{F0: uint32(v28), F1: 102}
                                        _ = v107
                                    }
                                    v28 = v28
                                    if ((fi105 / ((fi105) | 1)) != (fi105 * (fi105 % ((fi105) | 1)))) {
                                        v108 := GenS12{F0: -9223372036854775808}
                                        var v109 float64 = v5
                                        _ = v108
                                        _ = v109
                                    } else {
                                        fmt.Println(v29, v5, fi105)
                                    }
                                } else {
                                    if true {
                                        var v110 int32 = -1
                                        switch "\n\t" {
                                        case "hello":
                                            var v111 uint32 = 1
                                            _ = v111
                                        case "\n\t":
                                            switch "hello" {
                                            case "\n\t":
                                                var v112 int16 = (v80 | (-(v80)))
                                                _ = v112
                                            default:
                                                var v113 int8 = 127
                                                _ = v113
                                            }
                                        case "":
                                            fmt.Println(v29)
                                        }
                                        if ((!(128 != 1)) || true) {
                                            var v114 int8 = 0
                                            switch fi105 {
                                            case 9223372036854775807:
                                                v28 = v28
                                            case 42:
                                                var v115 bool = false
                                                _ = v115
                                            }
                                            delete(v3, (fi105 * (^(fi105))))
                                            _ = v114
                                        } else {
                                            for fi116 := 0; fi116 < 16; fi116++ {
                                                var v117 string = "\n\t"
                                                v1 = (v1 ^ v29)
                                                v118 := GenS13{F0: (v110 - v110)}
                                                v4 += int64(v28)
                                                _ = fi116
                                                _ = v117
                                                _ = v118
                                            }
                                            var v119 int16 = (v80 - v80)
                                            _ = v119
                                        }
                                        _ = v110
                                    }
                                    v120 := GenS12{F0: int64(v28)}
                                    v121 := v3[fi105]
                                    _ = v120
                                    _ = v121
                                }
                            }
                            _ = fi105
                            _ = v106
                        }
                        if (568 < -1) {
                            v5 -= 1e10
                        } else {
                            fmt.Println(v5, v1, v1)
                            var v122 float32 = 3.14159
                            delete(v3, -1)
                            _ = v122
                        }
                        v123 := GenS14{F0: 0, F1: 1, F2: false}
                        _ = v123
                    }
                default:
                    if (!true) {
                        if (true && (((!true) || (-72 >= -1)) || false)) {
                            _ = len(v2)
                        } else {
                            v29 = (^(v4))
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            switch "" {
                            case "foo bar":
                                if (((0.1 != 0.0) && false) && false) {
                                    v124 := GenS10{F0: -32768, F1: -191}
                                    _ = v124
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if false {
                                        var v125 int = 1
                                        var v126 float32 = 0.1
                                        _ = v125
                                        _ = v126
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            var v127 int16 = (v80 + -1)
                                            _ = v127
                                        case "":
                                            v28 = v28
                                        }
                                    }
                                    for fi128 := 0; fi128 < 11; fi128++ {
                                        var v129 int32 = 1
                                        _ = fi128
                                        _ = v129
                                    }
                                }
                            }
                        }
                    } else {
                        v1 += (v29 ^ (v1 ^ v29))
                    }
                }
            case 9223372036854775807:
                v130 := GenS3{F0: ("a" + ("hello" + "\n\t"))}
                _ = v130
            }
            if (!(!(4294967295 != 1))) {
                v80 = v80
            } else {
                v5 -= v5
                v1 += v4
                for fi131 := 0; fi131 < 10; fi131++ {
                    for fi132 := 0; fi132 < 18; fi132++ {
                        switch "foo bar" {
                        case "a":
                            v133 := GenS11{F0: uint32(v28), F1: 18}
                            _ = v133
                        case "\n\t":
                            v134 := GenS15{F0: 274, F1: (v28 == (^(v28)))}
                            _ = v134
                        default:
                            if true {
                                v80 -= v80
                                if true {
                                    fmt.Println(v4)
                                    v135 := GenS16{F0: 1}
                                    var v136 int32 = 1
                                    v4 = (v29 << ((105) & 63))
                                    _ = v135
                                    _ = v136
                                } else {
                                    fmt.Println(v28, fi131)
                                    if false {
                                        fmt.Println(v28)
                                        var v137 string = "\n\t"
                                        _ = v137
                                    } else {
                                        delete(v3, fi131)
                                        v138 := GenS5{F0: (842 <= 611), F1: 0}
                                        _ = v138
                                    }
                                }
                                var v139 uint = 42
                                _ = v139
                            }
                        }
                        var v140 float64 = (float64(3.14159) + (1.5 + (v5 / 1.5)))
                        _ = fi132
                        _ = v140
                    }
                    for fi141 := 0; fi141 < 2; fi141++ {
                        if false {
                            v142 := GenS7{F0: 42}
                            if ((v1 / ((v1) | 1)) <= 0) {
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                case "a":
                                    var v143 int32 = -89
                                    _ = v143
                                }
                            }
                            _ = v142
                        }
                        switch (fi131 % (((fi131 / ((fi131) | 1))) | 1)) {
                        case -782:
                            v144 := GenS15{F0: (v80 & v80), F1: true}
                            v1 += 42
                            v29 = (v1 ^ int64(v80))
                            _ = v144
                        case 32767:
                            v145 := GenS17{F0: 42, F1: 0, F2: "\n\t"}
                            _ = v145
                        case 371:
                            v1 -= (v29 - (v4 & 0))
                        default:
                            v146 := GenS18{F0: true}
                            var v147 float32 = -2.25
                            _ = v146
                            _ = v147
                        }
                        _ = fi141
                    }
                    _ = fi131
                }
            }
        }
    }
    _ = len(v2)
    delete(v3, 694)
    for fi148 := 0; fi148 < 5; fi148++ {
        fmt.Println(v4, fi148, v5)
        v149 := GenS19{F0: 65535, F1: true, F2: 0}
        _ = fi148
        _ = v149
    }
    if true {
        v150 := GenS18{F0: false}
        v2 = append(v2, 9)
        var v151 uint16 = v28
        _ = v150
        _ = v151
    }
    v3[9223372036854775807] = (((1.5 - ((0.1 - 0.1) - ((1e10 - ((1.5 - (0.0 + 3.14159)) / 1e10)) / -2.25))) - 0.1) <= 1e10)
    var v152 int16 = v80
    var v153 uint32 = 507
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v152
    _ = v153
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v79
    _ = v80
}
