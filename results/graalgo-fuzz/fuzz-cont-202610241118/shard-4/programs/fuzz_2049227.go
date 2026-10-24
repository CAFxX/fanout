package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 int64; F2 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 bool; F1 int32; F2 uint64 }
type GenS13 struct { F0 int16; F1 int32; F2 int16 }
type GenS14 struct { F0 int64; F1 uint8 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint8 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            if (!((("hello" + "") != "hello") || true)) {
                if true {
                    var v8 uint = 305
                    fmt.Println(fi6)
                    _ = v8
                } else {
                    if true {
                        var v9 uint = 290
                        _ = v9
                    }
                    if true {
                        var v10 uint32 = 2147483648
                        var v11 int = (fi6 >> ((4294967295) & 63))
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(fi7, fi7)
                        var v12 uint = 4294967295
                        _ = v12
                    }
                }
            }
            var v13 uint = 42
            v14 := GenS1{F0: -32768, F1: 18446744073709551615, F2: 1000}
            _ = fi7
            _ = v13
            _ = v14
        }
        var v15 int8 = 80
        _ = fi6
        _ = v15
    }
    var v16 int = -291
    v17 := []int{v16}
    if false {
        v16 = 9223372036854775807
        fmt.Println(v16)
        v16 += 128
    }
    _ = v16
    _ = v17
    return 53
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 uint16 = 255
    var v4 uint64 = 117
    var v5 float32 = 3.14159
    var v18 uint64 = (^(v4))
    if false {
        v3 = v3
        switch ("hello" + "") {
        case "\n\t":
            var v19 float64 = -2.25
            _ = v19
        case "foo bar":
            if true {
                v20 := GenS1{F0: 42, F1: v18, F2: -100}
                _ = v20
            } else {
                v18 = v4
            }
        case "hello":
            switch -9223372036854775808 {
            case -360:
                switch "a" {
                case "foo bar":
                    if true {
                        v3 -= 2
                    }
                case "hello":
                    if false {
                        var v21 int16 = 1
                        v22 := GenS2{F0: 0}
                        if true {
                            if true {
                                var v23 uint8 = 69
                                switch "foo bar" {
                                case "hello":
                                    var v24 int64 = 42
                                    _ = v24
                                case "\n\t":
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        if true {
                                            switch fi25 {
                                            case 1:
                                                switch (fi25 >> ((2) & 63)) {
                                                case 0:
                                                    v26 := GenS1{F0: v21, F1: v4, F2: (v21 / ((42) | 1))}
                                                    _ = v26
                                                case -9223372036854775808:
                                                    v23 = v23
                                                case 229:
                                                    var v27 float64 = (3.14159 / 1.0)
                                                    _ = v27
                                                }
                                            case -450:
                                                var v28 uint16 = v3
                                                _ = v28
                                            case 42:
                                                fmt.Println(v18, fi25, v18)
                                            }
                                        } else {
                                            switch fi25 {
                                            case 32767:
                                                v29 := GenS3{F0: 100, F1: genFn0()}
                                                _ = v29
                                            default:
                                                for fi30 := 0; fi30 < 10; fi30++ {
                                                    v31 := GenS2{F0: 2147483647}
                                                    _ = fi30
                                                    _ = v31
                                                }
                                            }
                                            if true {
                                                for fi32 := 0; fi32 < 5; fi32++ {
                                                    var v33 uint64 = v4
                                                    fmt.Println(v21, v3)
                                                    v34 := GenS4{F0: (^(fi25)), F1: ("hello" + ((("\n\t" + "hello") + (("\n\t" + "hello") + (("a" + "\n\t") + ""))) + "foo bar"))}
                                                    _ = fi32
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                            if false {
                                                var v35 int = 127
                                                for fi36 := 0; fi36 < 15; fi36++ {
                                                    var v37 uint = uint(v4)
                                                    var v38 uint16 = v3
                                                    _ = fi36
                                                    _ = v37
                                                    _ = v38
                                                }
                                                for fi39 := 0; fi39 < 9; fi39++ {
                                                    var v40 uint16 = (v3 * 0)
                                                    var v41 int64 = 0
                                                    v42 := []int{fi25, fi39, (v35 % ((v35) | 1)), (v35 << ((42) & 63))}
                                                    var v43 int16 = 32767
                                                    var v44 int32 = -2147483648
                                                    fmt.Println(v41, v43, v41)
                                                    var v45 int64 = (v41 % ((v41) | 1))
                                                    _ = fi39
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                }
                                                _ = v35
                                            } else {
                                                for fi46 := 0; fi46 < 3; fi46++ {
                                                    v23 = v23
                                                    v23 = v23
                                                    var v47 uint32 = 0
                                                    _ = fi46
                                                    _ = v47
                                                }
                                                for fi48 := 0; fi48 < 7; fi48++ {
                                                    var v49 int = fi25
                                                    var v50 string = "hello"
                                                    _ = fi48
                                                    _ = v49
                                                    _ = v50
                                                }
                                                v51 := []int{(fi25 & (fi25 - fi25)), 821, fi25, fi25}
                                                _ = v51
                                            }
                                        }
                                        _ = fi25
                                    }
                                }
                                _ = v23
                            }
                            switch "\n\t" {
                            case "foo bar":
                                var v52 uint8 = 255
                                _ = v52
                            case "":
                                v53 := []int{int(v3), 0}
                                _ = v53
                            default:
                                v54 := []int{127, 9223372036854775807, 0, 9223372036854775807}
                                _ = v54
                            }
                            v55 := v1[42]
                            _ = v55
                        } else {
                            switch "a" {
                            case "hello":
                                switch 128 {
                                case 9223372036854775807:
                                    var v56 bool = false
                                    _ = v56
                                default:
                                    v57 := v2[("a" + "a")]
                                    _ = v57
                                }
                            default:
                                v58 := []int{-2147483648, -2147483648, 1}
                                _ = v58
                            }
                        }
                        _ = v21
                        _ = v22
                    } else {
                        v59 := []int{32768, 706, -785}
                        _ = v59
                    }
                }
            case -589:
                switch 9223372036854775807 {
                case 32767:
                    for fi60 := 0; fi60 < 5; fi60++ {
                        if false {
                            v61 := GenS5{F0: "\n\t"}
                            if true {
                                for fi62 := 0; fi62 < 8; fi62++ {
                                    fmt.Println(fi60)
                                    var v63 uint32 = 2147483647
                                    _ = fi62
                                    _ = v63
                                }
                            }
                            if (1 < -755) {
                                fmt.Println(v4, v3, fi60)
                                switch "" {
                                case "a":
                                    fmt.Println(v5, v18)
                                case "\n\t":
                                    fmt.Println(fi60)
                                default:
                                    v64 := GenS5{F0: "hello"}
                                    _ = v64
                                }
                            }
                            _ = v61
                        } else {
                            var v65 uint8 = 2
                            fmt.Println(v3, v4, v3)
                            v3 -= uint16(v65)
                            _ = v65
                        }
                        v18 = v4
                        _ = fi60
                    }
                case 317:
                    if true {
                        var v66 int64 = -867
                        _ = v66
                    } else {
                        if false {
                            switch "" {
                            case "\n\t":
                                var v67 uint32 = 219
                                _ = v67
                            }
                            switch ("foo bar" + ("a" + "")) {
                            case "foo bar":
                                for fi68 := 0; fi68 < 19; fi68++ {
                                    var v69 uint8 = 1
                                    for fi70 := 0; fi70 < 11; fi70++ {
                                        var v71 float32 = (((v5 / v5) + 3.14159) * v5)
                                        _ = fi70
                                        _ = v71
                                    }
                                    switch (("\n\t" + "foo bar") + "\n\t") {
                                    case "a":
                                        var v72 uint32 = 2147483647
                                        _ = v72
                                    case "\n\t":
                                        if (!true) {
                                            fmt.Println(v4, v4, fi68)
                                        } else {
                                            _ = len(v1)
                                        }
                                    case "hello":
                                        if false {
                                            delete(v2, ("hello" + ("foo bar" + "a")))
                                            if ((-71 == 42) && (127 <= 127)) {
                                                if (127 > 76) {
                                                    var v73 float32 = v5
                                                    if false {
                                                        v18 += (v18 >> ((0) & 63))
                                                        for fi74 := 0; fi74 < 6; fi74++ {
                                                            var v75 int64 = 246
                                                            var v76 int32 = 42
                                                            var v77 int64 = (v75 & 42)
                                                            _ = fi74
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        v78 := []int{(fi68 * int(v18)), 1}
                                                        _ = v78
                                                    } else {
                                                        for fi79 := 0; fi79 < 15; fi79++ {
                                                            fmt.Println(v5, v73)
                                                            _ = fi79
                                                        }
                                                    }
                                                    switch fi68 {
                                                    case 232:
                                                        if false {
                                                            if ((^(fi68)) != (fi68 >> ((65535) & 63))) {
                                                                var v80 uint16 = 0
                                                                v4 = 476
                                                                v81 := []int{}
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            switch fi68 {
                                                            case 476:
                                                                v5 -= float32(((3.14159 + 3.14159) - -2.25))
                                                            case 32768:
                                                                fmt.Println(v3)
                                                            case 127:
                                                                v5 = 3.14159
                                                            default:
                                                                v82 := GenS6{F0: v69, F1: 42, F2: -23}
                                                                _ = v82
                                                            }
                                                            for fi83 := 0; fi83 < 17; fi83++ {
                                                                v4 = (v4 % ((42) | 1))
                                                                v84 := GenS3{F0: 127, F1: -128}
                                                                _ = fi83
                                                                _ = v84
                                                            }
                                                        } else {
                                                            v85 := []int{(fi68 | 32768)}
                                                            for fi86 := 0; fi86 < 13; fi86++ {
                                                                var v87 string = ("" + "a")
                                                                _ = fi86
                                                                _ = v87
                                                            }
                                                            for fi88 := 0; fi88 < 20; fi88++ {
                                                                if ((v3 * uint16(v18)) == v3) {
                                                                    var v89 int64 = 9223372036854775807
                                                                    var v90 int64 = v89
                                                                    v5 += v73
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                                if false {
                                                                    var v91 int32 = -2147483648
                                                                    _ = len(v85)
                                                                    var v92 uint32 = 4294967295
                                                                    if (41 <= 0) {
                                                                        var v93 int64 = -9223372036854775808
                                                                        _ = v93
                                                                    } else {
                                                                        var v94 bool = (!true)
                                                                        _ = v94
                                                                    }
                                                                    var v95 string = "hello"
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v95
                                                                } else {
                                                                    v5 += 3.14159
                                                                    delete(v2, "foo bar")
                                                                    v96 := GenS7{F0: -1}
                                                                    _ = v96
                                                                }
                                                                _ = fi88
                                                            }
                                                            _ = v85
                                                        }
                                                    case 0:
                                                        v97 := []int{(fi68 % (((fi68 | fi68)) | 1)), 127, -661, fi68}
                                                        _ = v97
                                                    case 4294967295:
                                                        delete(v1, 2147483647)
                                                    default:
                                                        v98 := GenS1{F0: 42, F1: (v18 % ((457) | 1)), F2: 286}
                                                        _ = v98
                                                    }
                                                    _ = v73
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            v1[fi68] = false
                                                            switch ("hello" + "a") {
                                                            case "\n\t":
                                                                v99 := GenS8{F0: (("" + (("a" + "a") + "\n\t")) + "a"), F1: (("foo bar" + ("" + "\n\t")) + "hello")}
                                                                _ = v99
                                                            }
                                                        } else {
                                                            fmt.Println(v69, fi68, fi68)
                                                            v100 := GenS9{F0: 988}
                                                            _ = v100
                                                        }
                                                    case "a":
                                                        for fi101 := 0; fi101 < 20; fi101++ {
                                                            fmt.Println(v18)
                                                            if ("a" != "a") {
                                                                v69 += (v69 * 127)
                                                            } else {
                                                                var v102 uint32 = 4294967295
                                                                if true {
                                                                    v103 := GenS6{F0: v69, F1: 1, F2: -47}
                                                                    _ = v103
                                                                }
                                                                if (!(("" == ((("a" + "") + "a") + ("hello" + "\n\t"))) && ((fi68 * (fi101 ^ 0)) > int(v18)))) {
                                                                    fmt.Println(fi101, v18)
                                                                }
                                                                fmt.Println(v4)
                                                                _ = v102
                                                            }
                                                            _ = fi101
                                                        }
                                                    case "hello":
                                                        for fi104 := 0; fi104 < 15; fi104++ {
                                                            delete(v1, fi104)
                                                            v18 += v18
                                                            _ = fi104
                                                        }
                                                    }
                                                    switch (("a" + "foo bar") + "a") {
                                                    case "a":
                                                        fmt.Println(v18, fi68, v18)
                                                    }
                                                }
                                            } else {
                                                var v105 int = -2147483648
                                                v69 = 128
                                                var v106 int32 = -36
                                                _ = v105
                                                _ = v106
                                            }
                                        } else {
                                            var v107 float32 = float32((1.5 - 3.14159))
                                            _ = v107
                                        }
                                    }
                                    _ = fi68
                                    _ = v69
                                }
                            case "hello":
                                var v108 uint32 = 42
                                _ = v108
                            case "":
                                for fi109 := 0; fi109 < 6; fi109++ {
                                    if ((v3 & v3) > (v3 * v3)) {
                                        var v110 int = fi109
                                        if (false || true) {
                                            v5 += v5
                                            v111 := []int{(fi109 | -369)}
                                            delete(v2, ("a" + "foo bar"))
                                            _ = v111
                                        }
                                        switch ("foo bar" + ("\n\t" + ("foo bar" + "\n\t"))) {
                                        case "foo bar":
                                            switch ("\n\t" + "") {
                                            case "a":
                                                for fi112 := 0; fi112 < 13; fi112++ {
                                                    var v113 bool = true
                                                    _ = fi112
                                                    _ = v113
                                                }
                                            case "":
                                                switch fi109 {
                                                case -1:
                                                    var v114 uint16 = 42
                                                    _ = v114
                                                case -461:
                                                    if (!(v18 <= (^(v4)))) {
                                                        v4 += v18
                                                        var v115 uint16 = (v3 << ((1) & 63))
                                                        _ = v115
                                                    } else {
                                                        var v116 int32 = 42
                                                        _ = v116
                                                    }
                                                }
                                            default:
                                                fmt.Println(fi109)
                                            }
                                        case "\n\t":
                                            var v117 int64 = 42
                                            _ = v117
                                        case "":
                                            for fi118 := 0; fi118 < 13; fi118++ {
                                                var v119 uint8 = 127
                                                v120 := GenS10{F0: v4}
                                                v121 := []int{-1, 0, (fi118 + 42), v110}
                                                _ = fi118
                                                _ = v119
                                                _ = v120
                                                _ = v121
                                            }
                                        }
                                        _ = v110
                                    } else {
                                        var v122 uint16 = 65535
                                        if false {
                                            if true {
                                                var v123 int8 = -25
                                                for fi124 := 0; fi124 < 20; fi124++ {
                                                    v125 := []int{(^(fi109))}
                                                    if (0.0 >= float32(0.1)) {
                                                        v122 = v3
                                                    }
                                                    _ = fi124
                                                    _ = v125
                                                }
                                                v126 := GenS7{F0: (^(v123))}
                                                _ = v123
                                                _ = v126
                                            } else {
                                                switch ("" + ("\n\t" + "\n\t")) {
                                                case "\n\t":
                                                    for fi127 := 0; fi127 < 12; fi127++ {
                                                        var v128 uint32 = 530
                                                        _ = fi127
                                                        _ = v128
                                                    }
                                                default:
                                                    v129 := GenS11{F0: 42}
                                                    _ = v129
                                                }
                                                for fi130 := 0; fi130 < 2; fi130++ {
                                                    v122 += (^(v3))
                                                    if (false && false) {
                                                        if (1 > -1) {
                                                            switch (^(fi130)) {
                                                            case 910:
                                                                switch fi130 {
                                                                case 814:
                                                                    v5 = (v5 - -2.25)
                                                                case 4294967295:
                                                                    v131 := GenS12{F0: (false && ((true || true) && true)), F1: 1, F2: v18}
                                                                    _ = v131
                                                                case 945:
                                                                    v132 := []int{fi109, fi109}
                                                                    delete(v2, ("\n\t" + "a"))
                                                                    var v133 uint16 = (v122 | 255)
                                                                    switch ("a" + "\n\t") {
                                                                    case "a":
                                                                        switch ("a" + ("\n\t" + "foo bar")) {
                                                                        case "hello":
                                                                            v18 = 945
                                                                        case "":
                                                                            v134 := v1[(fi109 & 4294967295)]
                                                                            _ = v134
                                                                        case "\n\t":
                                                                            v122 -= (v133 << ((561) & 63))
                                                                        }
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v135 uint16 = (v122 + v122)
                                                                            _ = v135
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            v136 := v2[""]
                                                                            if len(v132) >= 2 {
                                                                                v132 = v132[1:len(v132)]
                                                                            }
                                                                            v137 := v1[(fi109 - (fi130 % ((fi130) | 1)))]
                                                                            if false {
                                                                                v4 = (v4 / (((v4 & 85)) | 1))
                                                                                v1[(fi109 >> ((362) & 63))] = v137
                                                                            } else {
                                                                                v138 := GenS13{F0: 1000, F1: 42, F2: -691}
                                                                                fmt.Println(v137)
                                                                                _ = v138
                                                                            }
                                                                            _ = v136
                                                                            _ = v137
                                                                        }
                                                                    }
                                                                    v3 = v122
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                            case 664:
                                                                v18 += (v18 / (((v18 - (v18 << ((2) & 63)))) | 1))
                                                            case -9223372036854775808:
                                                                if true {
                                                                    if false {
                                                                        var v139 int64 = 0
                                                                        _ = v139
                                                                    } else {
                                                                        v5 -= (0.0 - v5)
                                                                        fmt.Println(v122)
                                                                    }
                                                                    fmt.Println(v3, fi109)
                                                                }
                                                            }
                                                            v3 = 756
                                                        } else {
                                                            fmt.Println(fi109, fi109)
                                                            switch (-(fi130)) {
                                                            case 4294967295:
                                                                var v140 bool = (("" + "\n\t") > (("hello" + "hello") + "hello"))
                                                                _ = v140
                                                            default:
                                                                v3 -= (v3 + v3)
                                                            }
                                                            if true {
                                                                if false {
                                                                    delete(v1, (fi109 ^ fi109))
                                                                }
                                                            }
                                                        }
                                                        v5 += v5
                                                        var v141 bool = (!false)
                                                        _ = v141
                                                    }
                                                    _ = fi130
                                                }
                                            }
                                        } else {
                                            var v142 int16 = 0
                                            if false {
                                                switch (fi109 - (fi109 >> ((uint(v18)) & 63))) {
                                                case -1:
                                                    switch fi109 {
                                                    case -236:
                                                        for fi143 := 0; fi143 < 9; fi143++ {
                                                            v144 := GenS4{F0: (fi109 & fi143), F1: ("foo bar" + "\n\t")}
                                                            delete(v1, (fi143 >> ((125) & 63)))
                                                            if (((v142 * (v142 ^ v142)) <= -47) && (!((!false) && true))) {
                                                                var v145 string = "foo bar"
                                                                var v146 uint8 = 240
                                                                var v147 bool = true
                                                                _ = v145
                                                                _ = v146
                                                                _ = v147
                                                            } else {
                                                                v18 = v18
                                                                v18 = (v18 | 18446744073709551615)
                                                            }
                                                            if false {
                                                                v148 := GenS14{F0: 42, F1: 255}
                                                                _ = v148
                                                            }
                                                            _ = fi143
                                                            _ = v144
                                                        }
                                                    }
                                                }
                                                v5 += (v5 - (v5 + 1e10))
                                                switch -9223372036854775808 {
                                                case 127:
                                                    var v149 float32 = v5
                                                    _ = v149
                                                case -1:
                                                    var v150 uint64 = (^(v4))
                                                    _ = v150
                                                case 0:
                                                    var v151 int64 = 0
                                                    _ = v151
                                                default:
                                                    _ = len(v2)
                                                }
                                            } else {
                                                var v152 int32 = 1
                                                _ = v152
                                            }
                                            v153 := GenS9{F0: v4}
                                            _ = v142
                                            _ = v153
                                        }
                                        var v154 uint8 = 127
                                        _ = v122
                                        _ = v154
                                    }
                                    v1[128] = (!(!false))
                                    v5 -= (float32(1e10) - (0.0 - (v5 + v5)))
                                    _ = fi109
                                }
                            default:
                                v155 := GenS15{F0: 2}
                                _ = v155
                            }
                        }
                        var v156 int32 = 1
                        _ = v156
                    }
                case 9223372036854775807:
                    v2["\n\t"] = "foo bar"
                }
            case -2147483648:
                var v157 bool = (true && (!true))
                _ = v157
            default:
                v3 = (v3 ^ (v3 / (((v3 ^ 586)) | 1)))
            }
        }
        switch 9223372036854775807 {
        case 1:
            if (!(177 < 0)) {
                v3 = v3
                v5 += ((v5 + (v5 * 0.1)) / v5)
                v158 := []int{750}
                _ = v158
            }
        }
    }
    v18 += (v18 << ((957) & 63))
    var v159 int16 = 32767
    var v160 float32 = (0.1 - v5)
    if false {
        fmt.Println(v5)
        var v161 int = 128
        var v162 int8 = -1
        _ = v161
        _ = v162
    }
    fmt.Println(v3, v4)
    v4 += (v18 >> ((0) & 63))
    v163 := GenS6{F0: 42, F1: int64(v4), F2: -82}
    v18 += 582
    var v164 float32 = v160
    v165 := GenS16{F0: v159}
    var v166 int16 = 42
    v3 -= (v3 * v3)
    var v167 int16 = v166
    v168 := GenS17{F0: 42}
    v169 := GenS3{F0: 14, F1: 100}
    v170 := v1[2147483647]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v164:", v164)
    fmt.Println("v166:", v166)
    fmt.Println("v167:", v167)
    fmt.Println("v170:", v170)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v159
    _ = v160
    _ = v163
    _ = v164
    _ = v165
    _ = v166
    _ = v167
    _ = v168
    _ = v169
    _ = v170
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
