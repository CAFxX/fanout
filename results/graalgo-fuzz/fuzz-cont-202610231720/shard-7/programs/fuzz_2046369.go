package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint32; F2 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32; F1 int32 }
type GenS11 struct { F0 uint64; F1 int8 }

func main() {
    var v1 bool = false
    var v2 float64 = (0.0 + 1e10)
    v3 := make(map[string]bool)
    var v4 int8 = 100
    v5 := make(map[uint32]bool)
    var v6 uint8 = 42
    if false {
        switch "\n\t" {
        case "a":
            var v7 float64 = v2
            _ = v7
        case "\n\t":
            if (65535 <= 4294967295) {
                v8 := []int{4294967295, 127, 490}
                _ = v8
            } else {
                if (v1 || (!false)) {
                    v9 := GenS1{F0: 42}
                    _ = v9
                }
            }
        case "hello":
            v6 += (v6 ^ v6)
        }
        v10 := GenS1{F0: -2147483648}
        if false {
            v11 := GenS2{F0: v4, F1: 65535}
            switch "hello" {
            case "":
                var v12 uint64 = 2
                _ = v12
            case "\n\t":
                var v13 bool = false
                _ = v13
            case "hello":
                if false {
                    fmt.Println(v2)
                    v4 = (^(v4))
                    v1 = v1
                } else {
                    if (("a" + "a") == ("" + (("\n\t" + "") + (("\n\t" + "a") + "hello")))) {
                        switch "\n\t" {
                        case "foo bar":
                            v14 := GenS2{F0: 0, F1: 9}
                            _ = v14
                        case "hello":
                            for fi15 := 0; fi15 < 17; fi15++ {
                                if false {
                                    var v16 float32 = (((3.14159 - float32(((0.0 - (v2 * v2)) + float64(0.1)))) + (float32(1.5) + float32(float64(-2.25)))) * (3.14159 - (((float32(v2) * 1e10) * (1e10 / (float32(v2) - (1.5 + -2.25)))) - 0.1)))
                                    for fi17 := 0; fi17 < 7; fi17++ {
                                        v18 := GenS2{F0: v4, F1: 273}
                                        v2 = 1.5
                                        _ = fi17
                                        _ = v18
                                    }
                                    if (v1 && v1) {
                                        var v19 uint = 18446744073709551615
                                        if (!v1) {
                                            var v20 int32 = -2147483648
                                            var v21 uint16 = 65535
                                            _ = v20
                                            _ = v21
                                        }
                                        v22 := GenS1{F0: -1}
                                        _ = v19
                                        _ = v22
                                    } else {
                                        v23 := GenS2{F0: (^(v4)), F1: 445}
                                        v24 := []int{}
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v16
                                } else {
                                    var v25 int8 = 127
                                    fmt.Println(v25)
                                    _ = v25
                                }
                                _ = fi15
                            }
                        case "":
                            v26 := GenS1{F0: 0}
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 8; fi27++ {
                                if v1 {
                                    for fi28 := 0; fi28 < 2; fi28++ {
                                        switch fi27 {
                                        case 4294967295:
                                            if (4294967295 > 846) {
                                                v2 -= float64(0.0)
                                            }
                                            var v29 int32 = 846
                                            var v30 int8 = -6
                                            _ = v29
                                            _ = v30
                                        default:
                                            _ = len(v5)
                                        }
                                        var v31 uint16 = 2
                                        if (821 != 462) {
                                            v1 = false
                                            v32 := []int{fi27, (fi27 + (-(fi28)))}
                                            _ = v32
                                        } else {
                                            switch 92 {
                                            case 0:
                                                if ((v31 + v31) != (v31 + (v31 - v31))) {
                                                    var v33 uint16 = v31
                                                    fmt.Println(fi28)
                                                    v1 = true
                                                    _ = v33
                                                } else {
                                                    v34 := GenS2{F0: (v4 * (-(v4))), F1: 2147483647}
                                                    v35 := GenS1{F0: 1}
                                                    _ = v34
                                                    _ = v35
                                                }
                                            case 127:
                                                v36 := v3["hello"]
                                                _ = v36
                                            }
                                        }
                                        _ = fi28
                                        _ = v31
                                    }
                                    var v37 int16 = 32767
                                    v6 = 128
                                    _ = v37
                                }
                                _ = fi27
                            }
                        }
                        var v38 uint64 = 18446744073709551615
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 20; fi39++ {
                        fmt.Println(v6)
                        if false {
                            v40 := GenS1{F0: 125}
                            fmt.Println(v6, v4)
                            v41 := GenS3{F0: 0, F1: (^(v6)), F2: "hello"}
                            _ = v40
                            _ = v41
                        }
                        v42 := []int{}
                        _ = fi39
                        _ = v42
                    }
                }
            default:
                v5[0] = (false || ((!(42 > -537)) && (967 > 2)))
            }
            _ = v11
        }
        _ = v10
    }
    v43 := GenS1{F0: 2147483647}
    if (v1 || v1) {
        v44 := GenS4{F0: v6}
        var v45 int64 = -1
        for fi46 := 0; fi46 < 18; fi46++ {
            for fi47 := 0; fi47 < 2; fi47++ {
                var v48 uint16 = 255
                var v49 uint16 = (v48 ^ v48)
                _ = fi47
                _ = v48
                _ = v49
            }
            v1 = v1
            v4 = (v4 & v4)
            _ = fi46
        }
        _ = v44
        _ = v45
    } else {
        v50 := GenS1{F0: -469}
        v51 := GenS1{F0: 27}
        var v52 int8 = v4
        _ = v50
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        delete(v5, 111)
        if (float32(v2) > ((-2.25 - 0.1) - float32(v2))) {
            if (!v1) {
                v4 = v4
            }
            if v1 {
                for fi54 := 0; fi54 < 15; fi54++ {
                    if v1 {
                        v55 := []int{(fi54 & 221), (fi53 >> ((255) & 63))}
                        if true {
                            if (!(!(1 >= 725))) {
                                fmt.Println(v1)
                            } else {
                                switch (("\n\t" + ("hello" + ("a" + "\n\t"))) + "foo bar") {
                                case "a":
                                    v2 -= (v2 * 1e10)
                                case "\n\t":
                                    switch "hello" {
                                    case "a":
                                        v1 = (fi53 > (fi54 << ((42) & 63)))
                                    default:
                                        v56 := GenS5{F0: uint(v6)}
                                        var v57 uint8 = 1
                                        var v58 bool = (65535 < 681)
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                }
                            }
                            var v59 int32 = 44
                            v60 := v3[(("a" + (("\n\t" + "a") + (("foo bar" + ("" + "")) + "a"))) + "foo bar")]
                            v1 = ((1000 < 32767) || v60)
                            _ = v59
                            _ = v60
                        } else {
                            if (false && (1 > -1)) {
                                fmt.Println(fi54)
                                var v61 int = (fi54 % ((fi54) | 1))
                                var v62 uint = 255
                                fmt.Println(v1, fi54, v6)
                                _ = v61
                                _ = v62
                            } else {
                                fmt.Println(v1, v2)
                            }
                            var v63 int64 = 1
                            var v64 int32 = -2147483648
                            _ = v63
                            _ = v64
                        }
                        _ = v55
                    } else {
                        v65 := []int{fi53}
                        _ = v65
                    }
                    v2 = (v2 / (v2 - (v2 / v2)))
                    switch "a" {
                    case "":
                        var v66 int32 = 0
                        _ = v66
                    default:
                        switch fi54 {
                        case -9223372036854775808:
                            var v67 uint = 159
                            _ = v67
                        case -228:
                            switch fi53 {
                            case 0:
                                v4 -= v4
                            default:
                                var v68 uint8 = (v6 % ((179) | 1))
                                _ = v68
                            }
                        case 9223372036854775807:
                            v69 := GenS2{F0: (v4 % (((-(v4))) | 1)), F1: 2147483648}
                            _ = v69
                        default:
                            v70 := GenS1{F0: 0}
                            _ = v70
                        }
                    }
                    v71 := GenS1{F0: 0}
                    _ = fi54
                    _ = v71
                }
            }
            v72 := v3["hello"]
            _ = v72
        } else {
            v6 = 0
            var v73 uint64 = 18446744073709551615
            var v74 int16 = -673
            _ = v73
            _ = v74
        }
        _ = fi53
    }
    v75 := []int{}
    v3["\n\t"] = false
    var v76 uint = 18446744073709551615
    if (!(((0.0 / (v2 + (1.5 - (((v2 + -2.25) / 1e10) + float64(float32(-2.25)))))) * (0.1 + float64((-2.25 / float32(((1.5 + v2) - -2.25)))))) >= v2)) {
        switch 42 {
        case -1:
            for fi77 := 0; fi77 < 8; fi77++ {
                if false {
                    var v78 int = (fi77 % ((fi77) | 1))
                    v1 = v1
                    delete(v5, 4294967295)
                    _ = v78
                } else {
                    for fi79 := 0; fi79 < 8; fi79++ {
                        if len(v75) > 0 {
                            v75[0] = fi77
                        }
                        v80 := GenS6{F0: 1}
                        var v81 bool = (-724 <= 42)
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                }
                if ((v6 << (((v76 * 2)) & 63)) > 128) {
                    if false {
                        v4 = -64
                        var v82 uint = (^(v76))
                        if v1 {
                            v3[(("\n\t" + "hello") + "a")] = v1
                            switch "foo bar" {
                            case "":
                                fmt.Println(v6, v4)
                            case "hello":
                                v83 := GenS7{F0: v1, F1: 1, F2: (v4 | v4)}
                                _ = v83
                            }
                        }
                        _ = v82
                    } else {
                        var v84 uint = v76
                        if (((v6 + v6) > v6) || (!true)) {
                            var v85 uint32 = 0
                            var v86 uint16 = 255
                            v3["\n\t"] = v1
                            _ = v85
                            _ = v86
                        }
                        var v87 int16 = 1
                        _ = v84
                        _ = v87
                    }
                    fmt.Println(fi77)
                }
                if true {
                    fmt.Println(fi77, fi77, v2)
                    for fi88 := 0; fi88 < 5; fi88++ {
                        var v89 int64 = 42
                        _ = fi88
                        _ = v89
                    }
                    if v1 {
                        var v90 int64 = 1
                        if (236 == 92) {
                            v76 = (^(v76))
                            for fi91 := 0; fi91 < 18; fi91++ {
                                _ = len(v75)
                                switch (fi91 >> ((42) & 63)) {
                                case -225:
                                    v4 += 1
                                }
                                _ = fi91
                            }
                        }
                        _ = v90
                    }
                }
                _ = fi77
            }
        }
        var v92 int16 = -32768
        _ = v92
    } else {
        if false {
            if v1 {
                if ("\n\t" != "") {
                    if (!(!v1)) {
                        switch (((("hello" + "foo bar") + ("" + "\n\t")) + ("" + ("foo bar" + "foo bar"))) + "a") {
                        case "a":
                            v76 = (v76 * v76)
                        case "":
                            if (-916 <= 42) {
                                switch (("hello" + ("" + "foo bar")) + "") {
                                case "foo bar":
                                    if true {
                                        switch ("hello" + (("hello" + ("hello" + ("" + ("\n\t" + ("foo bar" + ""))))) + "")) {
                                        case "a":
                                            for fi93 := 0; fi93 < 14; fi93++ {
                                                switch fi93 {
                                                case 1:
                                                    var v94 uint8 = (v6 / (((v6 * (v6 >> (((v76 | (v76 + v76))) & 63)))) | 1))
                                                    _ = v94
                                                case 4294967295:
                                                    v6 = v6
                                                }
                                                v95 := v5[838]
                                                if v95 {
                                                    var v96 uint64 = 852
                                                    _ = v96
                                                }
                                                _ = fi93
                                                _ = v95
                                            }
                                        default:
                                            v4 = (v4 * 1)
                                        }
                                    } else {
                                        v97 := GenS6{F0: 0}
                                        v75 = append(v75, 2147483647)
                                        for fi98 := 0; fi98 < 12; fi98++ {
                                            for fi99 := 0; fi99 < 3; fi99++ {
                                                if len(v75) >= 2 {
                                                    v75 = v75[1:len(v75)]
                                                }
                                                var v100 string = ((("\n\t" + "") + "a") + "hello")
                                                var v101 uint8 = v6
                                                _ = fi99
                                                _ = v100
                                                _ = v101
                                            }
                                            v1 = true
                                            v102 := GenS8{F0: 2147483647}
                                            _ = fi98
                                            _ = v102
                                        }
                                        _ = v97
                                    }
                                case "a":
                                    switch -785 {
                                    case 4294967295:
                                        var v103 uint32 = 2147483648
                                        _ = v103
                                    }
                                case "hello":
                                    v104 := GenS6{F0: -796}
                                    _ = v104
                                default:
                                    if v1 {
                                        for fi105 := 0; fi105 < 11; fi105++ {
                                            var v106 float32 = 1.5
                                            _ = fi105
                                            _ = v106
                                        }
                                    } else {
                                        fmt.Println(v76, v76)
                                    }
                                }
                                fmt.Println(v1)
                            } else {
                                v4 -= 77
                                var v107 int64 = -1
                                _ = v107
                            }
                        case "foo bar":
                            fmt.Println(v76)
                        default:
                            var v108 uint16 = 65535
                            _ = v108
                        }
                        switch "a" {
                        case "hello":
                            for fi109 := 0; fi109 < 8; fi109++ {
                                if (18446744073709551615 >= 18446744073709551615) {
                                    var v110 uint32 = 282
                                    switch (-(fi109)) {
                                    case 120:
                                        if (1000 != 42) {
                                            v2 += float64(0.0)
                                            fmt.Println(v76, v1, v1)
                                            fmt.Println(v6)
                                        }
                                    case 42:
                                        v111 := GenS4{F0: 127}
                                        _ = v111
                                    }
                                    v1 = v1
                                    _ = v110
                                }
                                fmt.Println(v76, v2, v4)
                                _ = fi109
                            }
                        case "a":
                            var v112 float64 = float64(3.14159)
                            _ = v112
                        case "\n\t":
                            fmt.Println(v2)
                        }
                        var v113 uint16 = 32767
                        _ = v113
                    }
                } else {
                    var v114 bool = (0 == 65535)
                    var v115 int32 = -336
                    if ("hello" >= ("" + "hello")) {
                        v116 := GenS9{F0: v114}
                        if ((v6 << ((0) & 63)) != v6) {
                            v76 = (v76 >> (((v76 + 42)) & 63))
                            v2 = v2
                            fmt.Println(v2, v2, v76)
                        } else {
                            for fi117 := 0; fi117 < 11; fi117++ {
                                v114 = (v2 > (3.14159 - float64(0.1)))
                                v3[""] = ((v2 / float64(-2.25)) >= float64(1e10))
                                var v118 float64 = float64(1.5)
                                _ = fi117
                                _ = v118
                            }
                            var v119 int8 = -31
                            _ = v119
                        }
                        for fi120 := 0; fi120 < 11; fi120++ {
                            for fi121 := 0; fi121 < 18; fi121++ {
                                if (!(((!(true && v114)) || v114) || true)) {
                                    v122 := GenS10{F0: 739, F1: v115}
                                    var v123 uint16 = 0
                                    _ = v122
                                    _ = v123
                                } else {
                                    v115 += v115
                                    if (32767 >= 32767) {
                                        switch (-(fi120)) {
                                        case 32767:
                                            var v124 uint8 = v6
                                            _ = v124
                                        case -2147483648:
                                            var v125 float64 = v2
                                            _ = v125
                                        case 128:
                                            v76 -= v76
                                        }
                                        if len(v75) >= 2 {
                                            v75 = v75[1:len(v75)]
                                        }
                                    }
                                }
                                v114 = true
                                _ = fi121
                            }
                            _ = len(v75)
                            _ = fi120
                        }
                        _ = v116
                    }
                    _ = v114
                    _ = v115
                }
                for fi126 := 0; fi126 < 17; fi126++ {
                    fmt.Println(fi126, fi126, v6)
                    v75 = append(v75, fi126)
                    _ = fi126
                }
            }
        } else {
            for fi127 := 0; fi127 < 2; fi127++ {
                v128 := v5[1]
                var v129 int32 = 1
                _ = fi127
                _ = v128
                _ = v129
            }
            v2 += 3.14159
            for fi130 := 0; fi130 < 9; fi130++ {
                var v131 int64 = 582
                _ = fi130
                _ = v131
            }
        }
        v4 = (v4 - (^(v4)))
    }
    if len(v75) >= 2 {
        v75 = v75[1:len(v75)]
    }
    if ((!v1) && false) {
        v76 = (v76 | 0)
    } else {
        var v132 int16 = 0
        _ = v132
    }
    var v133 int64 = -620
    if ((v4 >= v4) || (!true)) {
        if v1 {
            var v134 float64 = (-2.25 / 1.0)
            v135 := v5[448]
            _ = v134
            _ = v135
        }
        v136 := GenS8{F0: 744}
        _ = v136
    } else {
        fmt.Println(v2)
        v4 -= v4
        fmt.Println(v6)
    }
    for fi137 := 0; fi137 < 14; fi137++ {
        v138 := GenS8{F0: 1}
        _ = fi137
        _ = v138
    }
    for fi139 := 0; fi139 < 13; fi139++ {
        if ((!true) || (v1 || ("\n\t" > ("a" + ("" + "a"))))) {
            var v140 int64 = (-(v133))
            _ = v140
        } else {
            v75 = append(v75, 127)
        }
        for fi141 := 0; fi141 < 17; fi141++ {
            if false {
                var v142 float64 = (0.0 - 1e10)
                if len(v75) > 0 {
                    v75[0] = fi141
                }
                _ = v142
            }
            if false {
                fmt.Println(v6)
                var v143 float32 = float32(((v2 / (float64((0.1 * float32(v2))) + float64(0.1))) * ((v2 - 0.0) * 1.5)))
                if (-271 <= int16(v6)) {
                    var v144 int = (fi141 ^ (fi139 * 127))
                    if (v2 < -2.25) {
                        if v1 {
                            var v145 int = (^(v144))
                            v76 = (v76 | (v76 ^ (v76 - (v76 | (v76 | v76)))))
                            if (!true) {
                                switch (fi141 * v145) {
                                case -1:
                                    var v146 uint16 = 255
                                    if (!v1) {
                                        var v147 int64 = 42
                                        _ = v147
                                    } else {
                                        if v1 {
                                            var v148 int8 = (^(v4))
                                            var v149 uint32 = 84
                                            if (fi139 != v144) {
                                                var v150 int8 = (v4 * 4)
                                                _ = v150
                                            } else {
                                                var v151 uint32 = 337
                                                v152 := GenS3{F0: 255, F1: (v6 - v6), F2: "hello"}
                                                _ = v151
                                                _ = v152
                                            }
                                            _ = v148
                                            _ = v149
                                        } else {
                                            v5[2] = (42 == -916)
                                            if true {
                                                v143 = (v143 + (v143 / v143))
                                                fmt.Println(v144, v4)
                                            }
                                        }
                                    }
                                    _ = v146
                                default:
                                    switch (^(v144)) {
                                    case 0:
                                        v75 = append(v75, int(v6))
                                        _ = len(v5)
                                    default:
                                        v76 = 0
                                    }
                                }
                            }
                            _ = v145
                        } else {
                            if ((v133 * (v133 % (((^(v133))) | 1))) != (v133 & (v133 | (v133 - (v133 % ((1) | 1)))))) {
                                if v1 {
                                    var v153 int16 = 1000
                                    v153 = 464
                                    _ = v153
                                } else {
                                    v76 = (v76 - (v76 % ((255) | 1)))
                                    delete(v5, 22)
                                }
                                var v154 int8 = v4
                                delete(v3, "foo bar")
                                _ = v154
                            } else {
                                v1 = v1
                            }
                        }
                    }
                    _ = v144
                } else {
                    var v155 uint8 = v6
                    delete(v3, "a")
                    fmt.Println(v1)
                    _ = v155
                }
                _ = v143
            } else {
                if v1 {
                    v1 = v1
                    v133 = v133
                    switch (fi141 + (fi139 & fi141)) {
                    case -449:
                        var v156 string = "hello"
                        _ = v156
                    }
                    var v157 uint64 = 0
                    var v158 uint64 = v157
                    var v159 uint64 = uint64(v76)
                    _ = v157
                    _ = v158
                    _ = v159
                } else {
                    v160 := GenS11{F0: 714, F1: 0}
                    if len(v75) >= 2 {
                        v75 = v75[1:len(v75)]
                    }
                    _ = v160
                }
                var v161 uint64 = 1
                _ = v161
            }
            _ = fi141
        }
        _ = fi139
    }
    if false {
        var v162 uint = v76
        _ = v162
    }
    fmt.Println("v1:", v1)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
