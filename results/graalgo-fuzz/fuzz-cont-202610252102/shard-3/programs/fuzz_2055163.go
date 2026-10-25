package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int8; F1 int32; F2 uint16 }
type GenS6 struct { F0 int16; F1 int64; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16; F1 int64 }
type GenS14 struct { F0 string; F1 string; F2 string }
type GenS15 struct { F0 int64; F1 uint; F2 int64 }
type GenS16 struct { F0 bool; F1 int8; F2 int8 }
type GenS17 struct { F0 int64; F1 int16 }
type GenS18 struct { F0 uint32; F1 int8 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int32; F1 uint }
type GenS21 struct { F0 uint32; F1 int }
type GenS22 struct { F0 uint64 }
func genFn0() uint64 {
    var v4 uint8 = 191
    v5 := GenS1{F0: 255, F1: 245, F2: 884}
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int8 = 100
        if false {
            v8 := []int{fi6, fi6, (fi6 | (fi6 * fi6))}
            _ = v8
        } else {
            switch fi6 {
            case 9223372036854775807:
                var v9 int = fi6
                _ = v9
            default:
                switch int(v7) {
                case -1:
                    if true {
                        var v10 int64 = 42
                        v4 -= 127
                        _ = v10
                    } else {
                        v4 -= v4
                    }
                default:
                    switch fi6 {
                    case 132:
                        v11 := []int{fi6}
                        _ = v11
                    case -2147483648:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            v13 := GenS2{F0: 34, F1: 255, F2: 9223372036854775807}
                            _ = fi12
                            _ = v13
                        }
                    }
                }
            }
            if ("hello" > "\n\t") {
                fmt.Println(v7, v4)
            } else {
                v14 := GenS3{F0: 872, F1: 2}
                for fi15 := 0; fi15 < 13; fi15++ {
                    switch (^(fi6)) {
                    case 32768:
                        v7 -= 1
                    case 42:
                        v16 := GenS4{F0: -48, F1: (v4 ^ (v4 - 2))}
                        _ = v16
                    }
                    _ = fi15
                }
                fmt.Println(v4)
                _ = v14
            }
        }
        var v17 int64 = -9223372036854775808
        _ = fi6
        _ = v17
        _ = v7
    }
    _ = v4
    _ = v5
    return 42
}
func genFn1() int {
    v18 := []int{127, 0, 1}
    if true {
        if (0 >= -9223372036854775808) {
            if len(v18) > 0 {
                v18[0] = 4294967295
            }
            var v19 uint64 = 18446744073709551615
            v19 += v19
            _ = v19
        } else {
            switch "\n\t" {
            case "\n\t":
                if (0 < 69) {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v22 := GenS5{F0: 100, F1: 1, F2: 894}
                            _ = fi21
                            _ = v22
                        }
                        var v23 int32 = 0
                        var v24 uint32 = 4294967295
                        _ = fi20
                        _ = v23
                        _ = v24
                    }
                }
            case "foo bar":
                if false {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 int64 = 1
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        v27 := GenS5{F0: -128, F1: 2147483647, F2: 65535}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    var v28 int64 = -9223372036854775808
                    _ = v28
                } else {
                    var v29 int8 = 0
                    _ = v29
                }
            case "a":
                for fi30 := 0; fi30 < 3; fi30++ {
                    v18 = append(v18, fi30)
                    v31 := GenS2{F0: 0, F1: 2, F2: 0}
                    v32 := GenS6{F0: 42, F1: 1, F2: genFn0()}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            default:
                var v33 uint64 = 0
                _ = v33
            }
        }
        var v34 int16 = 32767
        fmt.Println(v34)
        _ = v34
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        var v35 bool = (float64(((float32(0.1) - -2.25) - (1.5 - -2.25))) == 3.14159)
        _ = v35
    }
    if len(v18) > 0 {
        v18[0] = -9223372036854775808
    }
    var v36 uint64 = 853
    _ = v18
    _ = v36
    return -1
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v37 := GenS7{F0: 0}
    switch 2147483647 {
    case -2147483648:
        if true {
            var v38 uint8 = 127
            _ = v38
        }
    case -602:
        var v39 bool = (!true)
        _ = v39
    case 42:
        fmt.Println(v1)
    }
    var v40 uint32 = 1
    for fi41 := 0; fi41 < 4; fi41++ {
        if false {
            var v42 string = ("foo bar" + "hello")
            var v43 string = "hello"
            _ = v42
            _ = v43
        }
        var v44 uint8 = 128
        _ = fi41
        _ = v44
    }
    v45 := []int{4294967295, int(v40)}
    if true {
        v46 := GenS8{F0: -957, F1: 9223372036854775807}
        if (2 > 304) {
            var v47 uint8 = 95
            _ = v47
        } else {
            switch 2147483647 {
            case 42:
                switch -69 {
                case 32768:
                    v48 := GenS9{F0: -982, F1: (0 >= 127)}
                    _ = v48
                case 4294967295:
                    v40 -= v40
                case 416:
                    v40 = v40
                default:
                    var v49 int32 = 1
                    _ = v49
                }
            case 4294967295:
                if ((0 == 9223372036854775807) && true) {
                    if false {
                        for fi50 := 0; fi50 < 4; fi50++ {
                            if len(v45) > 0 {
                                v45[0] = (fi50 - genFn1())
                            }
                            v40 += v40
                            fmt.Println(v40, fi50)
                            _ = fi50
                        }
                    } else {
                        var v51 float32 = 3.14159
                        if true {
                            if true {
                                if (0.0 <= (1.5 / 1e10)) {
                                    var v52 uint64 = 0
                                    for fi53 := 0; fi53 < 9; fi53++ {
                                        v45 = append(v45, (fi53 ^ fi53))
                                        v51 = 0.1
                                        _ = fi53
                                    }
                                    v40 = v40
                                    _ = v52
                                } else {
                                    fmt.Println(v51, v40, v51)
                                    fmt.Println(v40, v51, v1)
                                }
                                switch 127 {
                                case 4294967295:
                                    delete(v2, (v40 << ((304) & 63)))
                                default:
                                    var v54 int = -255
                                    _ = v54
                                }
                                v55 := GenS10{F0: -32768}
                                _ = v55
                            } else {
                                v1 = (v1 / ((172) | 1))
                            }
                            v56 := GenS11{F0: 2}
                            _ = v56
                        }
                        _ = v51
                    }
                    if false {
                        v40 -= v40
                        v1 = v1
                        var v57 float32 = 1e10
                        _ = v57
                    }
                    var v58 int8 = 42
                    _ = v58
                } else {
                    var v59 int = genFn1()
                    fmt.Println(v40)
                    fmt.Println(v1, v1, v59)
                    _ = v59
                }
            case 1:
                if false {
                    if true {
                        for fi60 := 0; fi60 < 9; fi60++ {
                            for fi61 := 0; fi61 < 3; fi61++ {
                                var v62 float32 = -2.25
                                v1 = v1
                                var v63 uint = (v1 + (v1 ^ (v1 >> (((v1 % (((v1 / (((^(v1))) | 1))) | 1))) & 63))))
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                            v64 := GenS12{F0: 47}
                            if len(v45) >= 2 {
                                v45 = v45[1:len(v45)]
                            }
                            _ = fi60
                            _ = v64
                        }
                        v65 := GenS8{F0: 1, F1: 42}
                        _ = v65
                    }
                    v40 = (v40 | v40)
                }
            default:
                var v66 float64 = ((0.0 - 0.1) + ((1e10 + 1.5) - float64(3.14159)))
                _ = v66
            }
        }
        _ = v46
    } else {
        if false {
            _ = len(v3)
            v3[4294967295] = false
            var v67 uint32 = 1
            _ = v67
        }
        if ("a" <= "") {
            for fi68 := 0; fi68 < 19; fi68++ {
                switch (fi68 << ((v1) & 63)) {
                case -2147483648:
                    if (false || false) {
                        _ = len(v45)
                        v40 += 2147483648
                        v40 -= v40
                    } else {
                        for fi69 := 0; fi69 < 5; fi69++ {
                            var v70 float32 = 1e10
                            var v71 uint32 = (^(v40))
                            if ("a" < "a") {
                                v72 := GenS9{F0: -2147483648, F1: (42 <= -1)}
                                _ = v72
                            }
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                    }
                case -488:
                    if ((uint32(v1) == 2147483647) && false) {
                        v1 = v1
                        if ((v1 - 255) <= v1) {
                            var v73 string = "hello"
                            _ = v73
                        } else {
                            switch "hello" {
                            case "hello":
                                v40 = (v40 << ((v1) & 63))
                            default:
                                fmt.Println(fi68, fi68, v1)
                            }
                            var v74 uint64 = genFn0()
                            v45 = append(v45, fi68)
                            _ = v74
                        }
                        var v75 int32 = -166
                        _ = v75
                    }
                }
                if true {
                    var v76 bool = (true && true)
                    var v77 uint8 = 128
                    switch "\n\t" {
                    case "hello":
                        if (127 >= v77) {
                            v77 += v77
                            var v78 float64 = 3.14159
                            fmt.Println(v77, v78, v40)
                            _ = v78
                        }
                    case "foo bar":
                        switch ("" + "hello") {
                        case "hello":
                            var v79 uint64 = 804
                            _ = v79
                        case "a":
                            var v80 int64 = 449
                            _ = v80
                        }
                    default:
                        fmt.Println(v77, v40)
                    }
                    _ = v76
                    _ = v77
                } else {
                    if (!(0 > uint8(v40))) {
                        v81 := GenS13{F0: 142, F1: 42}
                        switch (^(fi68)) {
                        case 2147483647:
                            for fi82 := 0; fi82 < 18; fi82++ {
                                var v83 float64 = 0.1
                                _ = fi82
                                _ = v83
                            }
                        case 32767:
                            v40 = (^(v40))
                        case 0:
                            var v84 uint8 = 255
                            _ = v84
                        default:
                            for fi85 := 0; fi85 < 20; fi85++ {
                                if true {
                                    var v86 int64 = -473
                                    v1 += 407
                                    _ = v86
                                }
                                var v87 uint64 = 0
                                _ = fi85
                                _ = v87
                            }
                        }
                        _ = v81
                    } else {
                        v1 = v1
                        v40 += 767
                        if len(v45) > 0 {
                            v45[0] = (fi68 - fi68)
                        }
                    }
                    if (127 <= 42) {
                        v40 += v40
                        switch 437 {
                        case 4294967295:
                            for fi88 := 0; fi88 < 16; fi88++ {
                                fmt.Println(fi68)
                                if false {
                                    _ = len(v2)
                                    var v89 uint32 = (v40 & 0)
                                    var v90 bool = (!(!(!true)))
                                    _ = v89
                                    _ = v90
                                } else {
                                    if true {
                                        var v91 bool = false
                                        _ = v91
                                    } else {
                                        switch (("a" + "") + "\n\t") {
                                        case "foo bar":
                                            v92 := GenS14{F0: (("foo bar" + "hello") + "foo bar"), F1: "foo bar", F2: "hello"}
                                            _ = v92
                                        case "a":
                                            v93 := GenS11{F0: 128}
                                            _ = v93
                                        }
                                        var v94 float64 = 3.14159
                                        _ = v94
                                    }
                                    v1 -= (v1 >> ((42) & 63))
                                    var v95 uint = 2
                                    _ = v95
                                }
                                v40 += (v40 ^ v40)
                                _ = fi88
                            }
                        case 869:
                            v40 = v40
                        case 2147483647:
                            if (genFn1() >= fi68) {
                                var v96 int64 = 1
                                _ = v96
                            }
                        default:
                            for fi97 := 0; fi97 < 16; fi97++ {
                                v40 = v40
                                v40 = v40
                                if false {
                                    switch ("\n\t" + "hello") {
                                    case "hello":
                                        if true {
                                            if len(v45) > 0 {
                                                v45[0] = fi97
                                            }
                                            v1 += (v1 % (((v1 * v1)) | 1))
                                        } else {
                                            v98 := GenS3{F0: 18446744073709551615, F1: (v1 / ((v1) | 1))}
                                            _ = v98
                                        }
                                    case "a":
                                        fmt.Println(fi97)
                                    default:
                                        switch (fi68 - fi68) {
                                        case 0:
                                            v40 += 0
                                        case 32768:
                                            v1 -= v1
                                        case 341:
                                            if (0.0 != (-2.25 * (float32(float64(1.5)) * 0.0))) {
                                                switch "a" {
                                                case "foo bar":
                                                    var v99 int64 = 9223372036854775807
                                                    _ = v99
                                                }
                                                v40 = (v40 >> (((^(v1))) & 63))
                                                _ = len(v3)
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    var v100 float32 = float32(0.0)
                                                    _ = v100
                                                case "":
                                                    var v101 int8 = 100
                                                    _ = v101
                                                default:
                                                    v40 += (v40 << ((v1) & 63))
                                                }
                                            }
                                        default:
                                            if false {
                                                if true {
                                                    var v102 string = "foo bar"
                                                    _ = v102
                                                } else {
                                                    switch ("foo bar" + "a") {
                                                    case "a":
                                                        v103 := v3[9223372036854775807]
                                                        _ = v103
                                                    case "\n\t":
                                                        var v104 uint = 42
                                                        _ = v104
                                                    case "foo bar":
                                                        v1 = 1
                                                    default:
                                                        var v105 uint = v1
                                                        _ = v105
                                                    }
                                                    var v106 int16 = 0
                                                    fmt.Println(v40, fi68)
                                                    if ((2147483647 != -1) || true) {
                                                        switch "hello" {
                                                        case "a":
                                                            var v107 float64 = 0.0
                                                            _ = v107
                                                        default:
                                                            v40 = (^(v40))
                                                        }
                                                        var v108 uint = 4294967295
                                                        if (32767 < 167) {
                                                            fmt.Println(v1, v106)
                                                        }
                                                        var v109 int8 = 0
                                                        if ((fi68 % ((fi68) | 1)) == (^(fi68))) {
                                                            v108 -= (v1 << (((v108 % (((v108 ^ 4294967295)) | 1))) & 63))
                                                            var v110 float32 = ((1.5 + float32((1e10 - 1.5))) / 3.14159)
                                                            _ = v110
                                                        } else {
                                                            switch ("foo bar" + ("a" + "\n\t")) {
                                                            case "hello":
                                                                v109 = 42
                                                            default:
                                                                _ = len(v45)
                                                            }
                                                            v111 := GenS15{F0: 1, F1: (v108 + (v108 - 255)), F2: -9223372036854775808}
                                                            _ = v111
                                                        }
                                                        v112 := GenS16{F0: true, F1: v109, F2: v109}
                                                        _ = v108
                                                        _ = v109
                                                        _ = v112
                                                    }
                                                    v1 += v1
                                                    v1 = 4294967295
                                                    var v113 uint16 = 286
                                                    _ = v106
                                                    _ = v113
                                                }
                                                fmt.Println(fi97)
                                                var v114 int32 = 610
                                                _ = v114
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    _ = len(v45)
                                                default:
                                                    v115 := GenS10{F0: 0}
                                                    _ = v115
                                                }
                                                var v116 int64 = 0
                                                v117 := GenS17{F0: v116, F1: -32768}
                                                _ = v116
                                                _ = v117
                                            }
                                        }
                                    }
                                    var v118 uint16 = 1
                                    _ = v118
                                }
                                _ = fi97
                            }
                        }
                    }
                    v2[(v40 ^ v40)] = "hello"
                }
                _ = fi68
            }
            for fi119 := 0; fi119 < 11; fi119++ {
                v40 += (v40 - 586)
                for fi120 := 0; fi120 < 18; fi120++ {
                    v121 := GenS1{F0: 969, F1: 251, F2: 528}
                    _ = fi120
                    _ = v121
                }
                if ((!true) && (127 == 255)) {
                    var v122 uint64 = 4294967295
                    var v123 uint32 = 634
                    for fi124 := 0; fi124 < 7; fi124++ {
                        v125 := GenS5{F0: 1, F1: 426, F2: 0}
                        _ = fi124
                        _ = v125
                    }
                    _ = v122
                    _ = v123
                } else {
                    v45 = append(v45, 1)
                }
                _ = fi119
            }
            var v126 int64 = -9223372036854775808
            _ = v126
        }
    }
    v127 := GenS16{F0: (false && (!false)), F1: 100, F2: 1}
    var v128 uint32 = (v40 + (v40 + (v40 - (v40 | (^(v40))))))
    if false {
        v129 := GenS18{F0: (v40 << (((v1 * (v1 - 42))) & 63)), F1: 127}
        var v130 float64 = 1.5
        _ = v129
        _ = v130
    }
    for fi131 := 0; fi131 < 19; fi131++ {
        if (!((((1e10 / (1e10 * 0.1)) * (0.0 / -2.25)) + 3.14159) == -2.25)) {
            v132 := v3[fi131]
            v133 := GenS19{F0: -128}
            _ = v132
            _ = v133
        } else {
            v40 += (^(v128))
            v134 := GenS20{F0: 1, F1: (v1 >> (((v1 | (v1 << (((v1 * (v1 - uint(v40)))) & 63)))) & 63))}
            if true {
                if false {
                    switch (fi131 << (((v1 * (v1 | 18446744073709551615))) & 63)) {
                    case -666:
                        v40 = (^(v128))
                    default:
                        var v135 uint8 = 0
                        _ = v135
                    }
                    v1 = (v1 % ((76) | 1))
                }
                for fi136 := 0; fi136 < 1; fi136++ {
                    var v137 float32 = -2.25
                    _ = fi136
                    _ = v137
                }
                v138 := GenS21{F0: 545, F1: (fi131 ^ fi131)}
                _ = v138
            } else {
                if true {
                    v3[365] = (-1 == -384)
                }
                fmt.Println(fi131)
            }
            _ = v134
        }
        var v139 uint64 = 42
        _ = fi131
        _ = v139
    }
    if true {
        var v140 int16 = -387
        _ = v140
    }
    var v141 uint8 = 127
    if false {
        v141 += 1
        fmt.Println(v40, v1)
    } else {
        v142 := GenS22{F0: 42}
        for fi143 := 0; fi143 < 12; fi143++ {
            v1 += (v1 << (((v1 >> (((v1 / ((v1) | 1))) & 63))) & 63))
            _ = fi143
        }
        switch 9223372036854775807 {
        case 127:
            fmt.Println(v128, v40, v1)
        default:
            for fi144 := 0; fi144 < 15; fi144++ {
                for fi145 := 0; fi145 < 4; fi145++ {
                    fmt.Println(v141, v141, fi144)
                    _ = fi145
                }
                var v146 int32 = 0
                v40 -= (^(v40))
                _ = fi144
                _ = v146
            }
        }
        _ = v142
    }
    if true {
        fmt.Println(v1)
        var v147 bool = false
        _ = v147
    } else {
        var v148 uint8 = (v141 / ((v141) | 1))
        v149 := GenS2{F0: uint32(v141), F1: 65535, F2: 185}
        fmt.Println(v1, v128)
        _ = v148
        _ = v149
    }
    for fi150 := 0; fi150 < 8; fi150++ {
        var v151 float32 = 1e10
        fmt.Println(v128, v151, v151)
        _ = fi150
        _ = v151
    }
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v141:", v141)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v127
    _ = v128
    _ = v141
    _ = v2
    _ = v3
    _ = v37
    _ = v40
    _ = v45
}
