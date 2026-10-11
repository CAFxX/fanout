package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32; F1 bool; F2 uint16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16; F1 uint }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint64; F1 int; F2 uint }
type GenS16 struct { F0 int; F1 bool }
func genFn0() int32 {
    v5 := GenS1{F0: 2, F1: true, F2: 713}
    if (1 > 80) {
        var v6 uint64 = 477
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 1, F1: true, F2: 42}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 0, F1: ((false && false) || (247 <= 127)), F2: uint16(v6)}
        _ = v6
        _ = v9
    } else {
        if false {
            v10 := GenS2{F0: true, F1: 128}
            if (!false) {
                var v11 int = -2147483648
                switch "foo bar" {
                case "foo bar":
                    var v12 uint16 = 1
                    _ = v12
                case "hello":
                    v11 -= (v11 + 127)
                case "":
                    var v13 uint32 = 2147483648
                    _ = v13
                }
                v14 := []int{}
                _ = v11
                _ = v14
            } else {
                v15 := GenS3{F0: "hello", F1: "a"}
                var v16 int64 = 1
                fmt.Println(v16)
                _ = v15
                _ = v16
            }
            _ = v10
        } else {
            v17 := GenS1{F0: 142, F1: true, F2: 2}
            if true {
                var v18 float64 = 1e10
                _ = v18
            } else {
                var v19 string = (("hello" + "hello") + "\n\t")
                _ = v19
            }
            _ = v17
        }
        var v20 float32 = ((0.0 + -2.25) - 1.5)
        v21 := []int{0}
        _ = v20
        _ = v21
    }
    _ = v5
    return -1
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    var v3 float64 = (((-2.25 / float64((1.5 * 0.1))) / (float64((-2.25 + (1.5 + 3.14159))) * 1.5)) * (((-2.25 - 1e10) / 3.14159) - float64(-2.25)))
    v4 := make(map[string]bool)
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            switch "a" {
            case "hello":
                switch "a" {
                case "\n\t":
                    v23 := []int{-9223372036854775808}
                    _ = v23
                case "":
                    switch fi22 {
                    case -231:
                        v3 = v3
                    case -108:
                        if false {
                            v24 := []int{}
                            _ = v24
                        }
                    case 232:
                        var v25 uint64 = 42
                        _ = v25
                    default:
                        fmt.Println(v1)
                    }
                case "hello":
                    if true {
                        if false {
                            for fi26 := 0; fi26 < 14; fi26++ {
                                _ = len(v2)
                                _ = fi26
                            }
                        } else {
                            v3 = v3
                            var v27 uint64 = 666
                            v1 = v1
                            _ = v27
                        }
                        v1 = v1
                    }
                }
            case "foo bar":
                var v28 int64 = -1
                _ = v28
            default:
                if false {
                    v29 := v2[(v1 | 588)]
                    _ = v29
                } else {
                    v30 := GenS3{F0: "hello", F1: "hello"}
                    _ = v30
                }
            }
            var v31 uint8 = 42
            _ = v31
        }
        fmt.Println(v3)
        _ = fi22
    }
    fmt.Println(v3)
    var v32 uint = 1
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 float32 = 1e10
        fmt.Println(v3, v1)
        _ = fi33
        _ = v34
    }
    if (-627 >= 898) {
        v1 = v1
        delete(v2, v1)
        switch ("hello" + "foo bar") {
        case "\n\t":
            if (!false) {
                if false {
                    var v35 uint64 = 14
                    for fi36 := 0; fi36 < 13; fi36++ {
                        for fi37 := 0; fi37 < 19; fi37++ {
                            if (!(v32 <= (v32 % (((v32 << (((v32 << (((v32 ^ (v32 % (((v32 - (v32 & v32))) | 1)))) & 63))) & 63))) | 1)))) {
                                if false {
                                    switch -2147483648 {
                                    case -1:
                                        v32 = 687
                                    }
                                } else {
                                    v38 := []int{32767, (fi37 + fi37)}
                                    v1 = (v1 >> (((v32 - (v32 + v32))) & 63))
                                    _ = v38
                                }
                                if (255 >= 42) {
                                    v3 = (float64(-2.25) * v3)
                                    var v39 int8 = -1
                                    _ = v39
                                } else {
                                    v40 := GenS2{F0: true, F1: 127}
                                    v41 := GenS4{F0: (v1 - (v1 % ((v1) | 1))), F1: ""}
                                    if ((!true) || false) {
                                        v42 := GenS4{F0: v1, F1: ""}
                                        v35 = 1
                                        if (42 >= 761) {
                                            v35 = (v35 & 0)
                                            var v43 int = int(v35)
                                            v44 := GenS5{F0: false}
                                            v45 := GenS6{F0: v1, F1: 1, F2: uint16(v32)}
                                            if (!(561 == 0)) {
                                                v3 = ((((v3 / (0.0 - 1e10)) + (3.14159 / float64(0.1))) / v3) + v3)
                                            }
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v42
                                    } else {
                                        if false {
                                            v46 := []int{}
                                            _ = v46
                                        } else {
                                            v47 := GenS6{F0: uint32(v35), F1: 261, F2: 65535}
                                            var v48 uint8 = 167
                                            v49 := GenS1{F0: (^(v1)), F1: (-1 != -2147483648), F2: 65535}
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                    var v50 int16 = -1
                                    _ = v40
                                    _ = v41
                                    _ = v50
                                }
                            } else {
                                switch "foo bar" {
                                case "":
                                    switch (fi36 & (fi36 % (((fi36 >> ((v32) & 63))) | 1))) {
                                    case -660:
                                        fmt.Println(fi36, fi37, fi36)
                                    default:
                                        if false {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v51 string = ("foo bar" + ("a" + "a"))
                                                _ = v51
                                            case "a":
                                                var v52 uint32 = 42
                                                _ = v52
                                            default:
                                                v53 := GenS1{F0: (v1 | (v1 / (((v1 - (v1 | uint32(v35)))) | 1))), F1: ((-2.25 - ((0.1 * float32((0.1 * ((v3 + v3) - 1.5)))) - 3.14159)) < float32((((v3 * v3) + (1.5 * v3)) * v3))), F2: 2}
                                                _ = v53
                                            }
                                            fmt.Println(v35)
                                            v54 := []int{0, fi37, fi37, fi37}
                                            v1 = v1
                                            if true {
                                                var v55 uint = v32
                                                _ = len(v4)
                                                _ = v55
                                            }
                                            _ = v54
                                        }
                                    }
                                    var v56 int8 = 1
                                    _ = v56
                                }
                                switch "" {
                                case "a":
                                    var v57 int8 = int8(v32)
                                    _ = v57
                                case "foo bar":
                                    if ((127 <= 127) || true) {
                                        var v58 int32 = genFn0()
                                        v59 := GenS6{F0: (v1 / ((811) | 1)), F1: v32, F2: 42}
                                        _ = v58
                                        _ = v59
                                    }
                                case "hello":
                                    v3 = v3
                                }
                            }
                            _ = fi37
                        }
                        v32 = v32
                        _ = fi36
                    }
                    var v60 bool = false
                    _ = v35
                    _ = v60
                }
                var v61 uint8 = 2
                if true {
                    var v62 int64 = -1
                    if false {
                        v61 = (v61 ^ (v61 >> (((v32 - (v32 / ((1) | 1)))) & 63)))
                        v1 = (v1 & uint32(v61))
                        var v63 float32 = 0.1
                        _ = v63
                    }
                    var v64 bool = true
                    _ = v62
                    _ = v64
                }
                _ = v61
            }
        case "hello":
            _ = len(v2)
        case "a":
            switch -2147483648 {
            case 68:
                v32 -= 2
            case -1:
                for fi65 := 0; fi65 < 7; fi65++ {
                    v66 := GenS4{F0: v1, F1: "\n\t"}
                    var v67 uint = 65535
                    for fi68 := 0; fi68 < 2; fi68++ {
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            v69 := []int{-731}
                            _ = v69
                        case "\n\t":
                            var v70 int8 = -100
                            _ = v70
                        case "a":
                            var v71 uint8 = 0
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                var v72 string = ("" + "\n\t")
                                _ = v72
                            case "a":
                                var v73 int = (fi65 | fi68)
                                _ = v73
                            }
                            v67 = v67
                            _ = v71
                        default:
                            v74 := GenS5{F0: false}
                            _ = v74
                        }
                        _ = fi68
                    }
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            case 0:
                _ = len(v2)
            default:
                var v75 int8 = 127
                _ = v75
            }
        default:
            if false {
                if (true || true) {
                    if false {
                        switch 9223372036854775807 {
                        case 4294967295:
                            for fi76 := 0; fi76 < 12; fi76++ {
                                switch fi76 {
                                case -9223372036854775808:
                                    fmt.Println(v3, v32)
                                case 127:
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(fi76)
                                    case "a":
                                        v1 = uint32(v32)
                                    default:
                                        v3 += 0.1
                                    }
                                case 208:
                                    for fi77 := 0; fi77 < 16; fi77++ {
                                        v2[(v1 * v1)] = "\n\t"
                                        _ = fi77
                                    }
                                default:
                                    switch ("foo bar" + (("a" + (("" + ("" + ("\n\t" + "hello"))) + "")) + ("foo bar" + "\n\t"))) {
                                    case "":
                                        v32 = 18446744073709551615
                                    case "\n\t":
                                        v78 := []int{(fi76 % ((fi76) | 1)), fi76, -321, 9223372036854775807}
                                        _ = v78
                                    case "hello":
                                        if false {
                                            var v79 bool = (!(-85 >= -128))
                                            v79 = ((!(0.0 != (0.0 + 1.5))) || (true || v79))
                                            _ = v79
                                        } else {
                                            fmt.Println(v3, v1)
                                            var v80 float64 = v3
                                            for fi81 := 0; fi81 < 13; fi81++ {
                                                fmt.Println(fi76)
                                                switch (fi76 / (((-(fi81))) | 1)) {
                                                case -589:
                                                    v82 := v2[(^(v1))]
                                                    _ = v82
                                                case 4294967295:
                                                    var v83 bool = (!false)
                                                    _ = v83
                                                default:
                                                    if (42 == 42) {
                                                        if false {
                                                            var v84 float32 = (1e10 - float32(float64(((0.1 - 0.0) - ((0.0 + 1.5) - (1e10 * 0.1))))))
                                                            _ = v84
                                                        }
                                                    } else {
                                                        fmt.Println(v3)
                                                        v85 := GenS7{F0: 477}
                                                        fmt.Println(v32, fi76)
                                                        _ = v85
                                                    }
                                                }
                                                if false {
                                                    var v86 float32 = 0.1
                                                    _ = v86
                                                }
                                                _ = fi81
                                            }
                                            _ = v80
                                        }
                                    default:
                                        for fi87 := 0; fi87 < 10; fi87++ {
                                            var v88 uint64 = 42
                                            v32 -= (^(v32))
                                            _ = fi87
                                            _ = v88
                                        }
                                    }
                                }
                                v89 := []int{}
                                _ = fi76
                                _ = v89
                            }
                        case 32767:
                            if (!true) {
                                if true {
                                    var v90 int32 = -488
                                    if false {
                                        var v91 uint32 = v1
                                        if (!("hello" == "foo bar")) {
                                            switch "\n\t" {
                                            case "a":
                                                fmt.Println(v3, v1)
                                            case "\n\t":
                                                fmt.Println(v1, v32, v90)
                                            case "":
                                                for fi92 := 0; fi92 < 5; fi92++ {
                                                    v3 = 1e10
                                                    switch 508 {
                                                    case -9223372036854775808:
                                                        fmt.Println(fi92)
                                                    }
                                                    switch "a" {
                                                    case "":
                                                        var v93 uint16 = 0
                                                        _ = v93
                                                    }
                                                    _ = fi92
                                                }
                                            }
                                            v32 -= 903
                                            v94 := GenS2{F0: (32767 != -1), F1: 127}
                                            _ = v94
                                        }
                                        _ = v91
                                    }
                                    _ = v90
                                }
                                v1 = 4294967295
                            }
                        default:
                            var v95 int32 = 2147483647
                            _ = v95
                        }
                        v96 := GenS8{F0: genFn0(), F1: true, F2: 42}
                        _ = v96
                    } else {
                        for fi97 := 0; fi97 < 13; fi97++ {
                            var v98 int32 = genFn0()
                            _ = fi97
                            _ = v98
                        }
                    }
                    var v99 int64 = 42
                    v1 = 4294967295
                    _ = v99
                } else {
                    v32 = (^(v32))
                    var v100 float32 = 0.0
                    _ = v100
                }
            } else {
                if (128 != -2147483648) {
                    v32 = (^(v32))
                    if (true || ((v1 + (v1 * (v1 << (((v32 << ((65535) & 63))) & 63)))) != (v1 << ((v32) & 63)))) {
                        v1 = (^(v1))
                        if (2 > 127) {
                            fmt.Println(v1)
                            for fi101 := 0; fi101 < 4; fi101++ {
                                switch fi101 {
                                case 376:
                                    v102 := []int{}
                                    _ = v102
                                case -958:
                                    v103 := GenS9{F0: 2147483647}
                                    _ = v103
                                default:
                                    var v104 uint32 = (v1 / (((v1 / (((v1 << (((^(v32))) & 63))) | 1))) | 1))
                                    _ = v104
                                }
                                var v105 uint32 = 1
                                _ = fi101
                                _ = v105
                            }
                            v106 := GenS10{F0: 1, F1: 6}
                            _ = v106
                        }
                        v107 := []int{199, 0, 948}
                        _ = v107
                    } else {
                        var v108 int64 = -922
                        for fi109 := 0; fi109 < 16; fi109++ {
                            switch fi109 {
                            case -1:
                                if (uint32(v32) >= 2147483648) {
                                    fmt.Println(v1, v1, fi109)
                                    var v110 int16 = 42
                                    _ = v110
                                } else {
                                    for fi111 := 0; fi111 < 7; fi111++ {
                                        var v112 float64 = -2.25
                                        var v113 int64 = -9223372036854775808
                                        v114 := []int{}
                                        if ((fi109 & -786) < (^(fi111))) {
                                            v32 += (v32 & v32)
                                            v1 = (^(v1))
                                            switch (^(fi109)) {
                                            case -9223372036854775808:
                                                var v115 int = int(v108)
                                                _ = v115
                                            default:
                                                v116 := GenS4{F0: (^(v1)), F1: "\n\t"}
                                                _ = v116
                                            }
                                        }
                                        v32 -= (v32 | v32)
                                        var v117 int64 = (v113 - v108)
                                        v2[(v1 % ((0) | 1))] = ("a" + ("foo bar" + "hello"))
                                        _ = fi111
                                        _ = v112
                                        _ = v113
                                        _ = v114
                                        _ = v117
                                    }
                                    switch 150 {
                                    case 32767:
                                        v108 += 0
                                    default:
                                        for fi118 := 0; fi118 < 5; fi118++ {
                                            if true {
                                                fmt.Println(v1)
                                                var v119 int32 = 1
                                                fmt.Println(v119)
                                                _ = v119
                                            }
                                            var v120 uint = (v32 ^ 2)
                                            _ = fi118
                                            _ = v120
                                        }
                                    }
                                    for fi121 := 0; fi121 < 17; fi121++ {
                                        fmt.Println(fi109)
                                        var v122 string = ""
                                        _ = fi121
                                        _ = v122
                                    }
                                }
                            default:
                                v1 = (v1 >> ((0) & 63))
                            }
                            _ = fi109
                        }
                        var v123 int64 = (v108 + v108)
                        _ = v108
                        _ = v123
                    }
                }
            }
        }
    }
    v32 += v32
    v124 := v4["\n\t"]
    switch 899 {
    case -724:
        var v125 uint32 = v1
        _ = v125
    case -1:
        for fi126 := 0; fi126 < 7; fi126++ {
            v124 = (2147483647 >= -2147483648)
            _ = fi126
        }
    }
    var v127 float32 = 0.0
    var v128 int64 = -444
    fmt.Println(v1, v1, v1)
    v129 := []int{}
    if false {
        fmt.Println(v3)
        for fi130 := 0; fi130 < 10; fi130++ {
            var v131 int16 = 573
            _ = fi130
            _ = v131
        }
        for fi132 := 0; fi132 < 9; fi132++ {
            if (0 >= -90) {
                var v133 bool = (!((!v124) && (false || false)))
                var v134 int8 = -128
                _ = len(v2)
                _ = v133
                _ = v134
            }
            switch "" {
            case "a":
                v135 := GenS7{F0: int32(v32)}
                _ = v135
            default:
                if len(v129) >= 2 {
                    v129 = v129[1:len(v129)]
                }
            }
            _ = fi132
        }
    }
    if (-2147483648 <= 1) {
        v129 = append(v129, 32768)
    } else {
        switch int(v32) {
        case 32767:
            v128 = (^(v128))
        default:
            if v124 {
                switch 32767 {
                case 9223372036854775807:
                    v136 := GenS6{F0: (^(v1)), F1: 42, F2: 368}
                    _ = v136
                case 0:
                    for fi137 := 0; fi137 < 20; fi137++ {
                        if true {
                            var v138 uint = v32
                            if (866 != 568) {
                                var v139 uint32 = (v1 % ((v1) | 1))
                                switch ("" + ("" + "hello")) {
                                case "":
                                    if v124 {
                                        var v140 bool = false
                                        v141 := GenS11{F0: v138, F1: (^(v138))}
                                        if true {
                                            for fi142 := 0; fi142 < 3; fi142++ {
                                                switch "" {
                                                case "hello":
                                                    v143 := GenS12{F0: ("" + "foo bar")}
                                                    _ = v143
                                                case "a":
                                                    if (((v1 & (v139 / ((v1) | 1))) <= (v139 << (((v32 * uint(v139))) & 63))) && false) {
                                                        fmt.Println(fi142, v1, v127)
                                                        v1 -= v1
                                                    } else {
                                                        var v144 int8 = 100
                                                        if (float32((v3 + (v3 / v3))) >= 1e10) {
                                                            var v145 int64 = -9223372036854775808
                                                            v139 = (^(v139))
                                                            _ = v145
                                                        } else {
                                                            v124 = v124
                                                            v139 -= v1
                                                        }
                                                        fmt.Println(v144)
                                                        _ = v144
                                                    }
                                                case "\n\t":
                                                    var v146 string = ("a" + "a")
                                                    if false {
                                                        if len(v129) >= 2 {
                                                            v129 = v129[1:len(v129)]
                                                        }
                                                        var v147 string = ("" + v146)
                                                        var v148 int8 = -94
                                                        _ = v147
                                                        _ = v148
                                                    } else {
                                                        v127 += 0.0
                                                        if len(v129) >= 2 {
                                                            v129 = v129[1:len(v129)]
                                                        }
                                                        switch (fi137 % (((fi142 % ((-20) | 1))) | 1)) {
                                                        case -2147483648:
                                                            var v149 uint16 = 2
                                                            _ = v149
                                                        case 32768:
                                                            if v124 {
                                                                fmt.Println(fi142, fi137)
                                                                v139 += 42
                                                                var v150 bool = false
                                                                _ = v150
                                                            }
                                                        case 4294967295:
                                                            var v151 int64 = v128
                                                            _ = v151
                                                        }
                                                        v152 := GenS13{F0: 1, F1: v138}
                                                        var v153 uint8 = 128
                                                        var v154 int16 = 33
                                                        var v155 float32 = (v127 / float32((v3 * ((v3 - -2.25) / (3.14159 + v3)))))
                                                        _ = len(v4)
                                                        if v140 {
                                                            var v156 uint = (v138 % ((257) | 1))
                                                            var v157 int64 = 9223372036854775807
                                                            fmt.Println(v124, v139)
                                                            v158 := GenS4{F0: (^(v139)), F1: v146}
                                                            _ = v156
                                                            _ = v157
                                                            _ = v158
                                                        }
                                                        _ = len(v2)
                                                        _ = v152
                                                        _ = v153
                                                        _ = v154
                                                        _ = v155
                                                    }
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v32, v140)
                                                    case "a":
                                                        var v159 uint64 = 1
                                                        _ = v159
                                                    default:
                                                        v160 := v2[(v1 + v1)]
                                                        _ = v160
                                                    }
                                                    v139 += (v139 ^ (v1 | (v139 % ((v139) | 1))))
                                                    _ = v146
                                                }
                                                delete(v2, 0)
                                                _ = fi142
                                            }
                                        }
                                        _ = v140
                                        _ = v141
                                    }
                                case "foo bar":
                                    fmt.Println(v127, fi137)
                                default:
                                    v161 := GenS14{F0: 127}
                                    _ = v161
                                }
                                _ = v139
                            } else {
                                v129 = append(v129, fi137)
                                var v162 uint16 = 42
                                _ = v162
                            }
                            _ = v138
                        } else {
                            _ = len(v129)
                            var v163 int16 = -32768
                            _ = v163
                        }
                        _ = fi137
                    }
                case -1:
                    v1 = 4294967295
                }
            } else {
                fmt.Println(v32)
                if (100 == 42) {
                    v164 := GenS15{F0: 4294967295, F1: 127, F2: (^(v32))}
                    _ = v164
                } else {
                    var v165 int64 = v128
                    v166 := GenS16{F0: 0, F1: (1 <= 1)}
                    v32 -= (v32 / ((v32) | 1))
                    _ = v165
                    _ = v166
                }
            }
        }
        fmt.Println(v1)
        var v167 uint = (v32 & (^(v32)))
        _ = v167
    }
    v32 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("slicelen v129:", len(v129), v129)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v124
    _ = v127
    _ = v128
    _ = v129
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
