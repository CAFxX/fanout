package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 uint; F1 int; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int; F1 int64; F2 bool }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS11 struct { F0 int64 }
func genFn0() uint64 {
    switch 151 {
    case 2147483647:
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 int16 = 0
            switch "\n\t" {
            case "hello":
                var v7 uint8 = 42
                _ = v7
            case "":
                switch "foo bar" {
                case "\n\t":
                    v8 := GenS1{F0: 2}
                    _ = v8
                }
            case "foo bar":
                var v9 int32 = -429
                _ = v9
            default:
                var v10 int = fi5
                _ = v10
            }
            _ = fi5
            _ = v6
        }
    case 9223372036854775807:
        var v11 uint32 = 2147483647
        _ = v11
    }
    var v12 uint = 255
    v13 := GenS1{F0: 92}
    v14 := GenS2{F0: -541, F1: 65535, F2: ((!(115 < -9223372036854775808)) && false)}
    var v15 int64 = 552
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 2
}
func genFn1() int64 {
    if (!true) {
        v16 := GenS2{F0: 4294967295, F1: 42, F2: (true && (!(0 <= 42)))}
        for fi17 := 0; fi17 < 10; fi17++ {
            if (!false) {
                var v18 uint64 = 673
                _ = v18
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(fi17, fi17)
                    if true {
                        var v20 uint8 = 255
                        if ((2 <= v20) && false) {
                            var v21 int8 = -128
                            fmt.Println(fi17)
                            _ = v21
                        }
                        var v22 int8 = 5
                        _ = v20
                        _ = v22
                    }
                    v23 := GenS3{F0: -1, F1: 371, F2: 65535}
                    if false {
                        var v24 string = ("" + ((("" + "a") + "a") + ("a" + "")))
                        var v25 string = ("hello" + v24)
                        _ = v24
                        _ = v25
                    }
                    var v26 float32 = 1e10
                    v27 := GenS1{F0: 4294967295}
                    _ = fi19
                    _ = v23
                    _ = v26
                    _ = v27
                }
                var v28 uint16 = 0
                v28 -= v28
                _ = v28
            }
            _ = fi17
        }
        _ = v16
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        if false {
            var v30 uint = 42
            for fi31 := 0; fi31 < 16; fi31++ {
                v30 = 255
                if ((!(fi29 == fi29)) && (18446744073709551615 > 971)) {
                    if (!true) {
                        v32 := []int{fi29, fi29, fi29, (fi31 & fi29)}
                        _ = v32
                    }
                    v30 += 255
                } else {
                    v33 := []int{(fi31 / ((32767) | 1)), 0, 0, 4294967295}
                    _ = v33
                }
                fmt.Println(v30, fi31, v30)
                _ = fi31
            }
            var v34 float64 = 1.5
            _ = v30
            _ = v34
        } else {
            v35 := GenS1{F0: 2147483648}
            for fi36 := 0; fi36 < 9; fi36++ {
                switch "\n\t" {
                case "foo bar":
                    var v37 string = "foo bar"
                    _ = v37
                case "":
                    v38 := []int{(fi36 << ((2) & 63))}
                    _ = v38
                default:
                    var v39 int64 = 42
                    _ = v39
                }
                switch "foo bar" {
                case "hello":
                    var v40 uint64 = 546
                    _ = v40
                case "a":
                    var v41 string = "hello"
                    _ = v41
                case "\n\t":
                    var v42 bool = ((1e10 / (1.5 / (1e10 - 1.5))) > float32((1.5 * float64((0.0 + 1.5)))))
                    _ = v42
                default:
                    if (fi36 != fi36) {
                        if false {
                            var v43 uint8 = 42
                            _ = v43
                        } else {
                            fmt.Println(fi36, fi29)
                            var v44 int16 = -1
                            _ = v44
                        }
                        var v45 int16 = 1000
                        _ = v45
                    }
                }
                var v46 uint8 = 0
                _ = fi36
                _ = v46
            }
            fmt.Println(fi29)
            _ = v35
        }
        switch fi29 {
        case 32768:
            if true {
                var v47 float32 = (0.0 + 0.1)
                _ = v47
            } else {
                if true {
                    var v48 uint32 = 499
                    var v49 int64 = -450
                    v48 -= v48
                    _ = v48
                    _ = v49
                } else {
                    v50 := []int{}
                    var v51 uint8 = 255
                    _ = v50
                    _ = v51
                }
                v52 := []int{}
                _ = v52
            }
        case -805:
            v53 := GenS3{F0: 0, F1: 1, F2: 2}
            _ = v53
        default:
            fmt.Println(fi29)
        }
        var v54 string = "hello"
        _ = fi29
        _ = v54
    }
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    v4 := []int{-837}
    delete(v2, v3)
    var v55 bool = true
    _ = len(v2)
    v56 := GenS1{F0: 65535}
    fmt.Println(v3)
    if (1 > 42) {
        switch v3 {
        case "hello":
            switch v3 {
            case "foo bar":
                _ = len(v2)
            case "":
                var v57 int = 32768
                _ = v57
            default:
                for fi58 := 0; fi58 < 6; fi58++ {
                    if (42 < 42) {
                        v55 = v55
                        v59 := GenS4{F0: 4294967295, F1: fi58, F2: 684}
                        for fi60 := 0; fi60 < 2; fi60++ {
                            fmt.Println(v55)
                            if v55 {
                                fmt.Println(fi60, v55, v3)
                                if v55 {
                                    var v61 uint8 = 127
                                    _ = v61
                                }
                                var v62 bool = v55
                                v3 = (v3 + ("a" + ""))
                                _ = v62
                            } else {
                                var v63 int = fi60
                                v2[("a" + v3)] = v55
                                v64 := GenS4{F0: 2, F1: fi58, F2: 32767}
                                _ = v63
                                _ = v64
                            }
                            _ = fi60
                        }
                        _ = v59
                    }
                    switch -1 {
                    case 872:
                        v65 := GenS5{F0: 65535, F1: -691, F2: genFn0()}
                        _ = v65
                    case 9223372036854775807:
                        if ((0.1 - 1e10) < (0.1 * 1e10)) {
                            v66 := GenS4{F0: 352, F1: 32768, F2: 116}
                            var v67 uint64 = 4294967295
                            v67 = v67
                            _ = v66
                            _ = v67
                        } else {
                            fmt.Println(v55, fi58, v3)
                            _ = len(v4)
                        }
                    }
                    switch v3 {
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = (fi58 % ((fi58) | 1))
                        }
                    default:
                        var v68 float64 = 3.14159
                        _ = v68
                    }
                    _ = fi58
                }
            }
        case "a":
            v2[("foo bar" + v3)] = (!v55)
        }
        for fi69 := 0; fi69 < 17; fi69++ {
            for fi70 := 0; fi70 < 14; fi70++ {
                var v71 uint16 = 65535
                _ = fi70
                _ = v71
            }
            if (2147483647 <= 2147483647) {
                var v72 string = "hello"
                _ = v72
            }
            _ = fi69
        }
        if ((0 >= 0) && true) {
            if v55 {
                for fi73 := 0; fi73 < 8; fi73++ {
                    v55 = v55
                    if v55 {
                        for fi74 := 0; fi74 < 4; fi74++ {
                            v3 = v3
                            v3 = v3
                            if (!false) {
                                var v75 string = ""
                                _ = v75
                            }
                            var v76 int8 = -110
                            switch (fi74 / ((42) | 1)) {
                            case 4294967295:
                                var v77 float32 = (1e10 / 0.1)
                                _ = v77
                            case -9223372036854775808:
                                v55 = ((1 >= 255) || (false || (9223372036854775807 >= -403)))
                            case 32767:
                                fmt.Println(fi74, v3)
                            default:
                                _ = len(v1)
                            }
                            var v78 float64 = 1e10
                            var v79 uint32 = 42
                            if v55 {
                                var v80 int64 = -339
                                var v81 int16 = -1
                                _ = v80
                                _ = v81
                            } else {
                                var v82 int64 = -9223372036854775808
                                var v83 uint16 = 401
                                _ = v82
                                _ = v83
                            }
                            _ = fi74
                            _ = v76
                            _ = v78
                            _ = v79
                        }
                        v84 := GenS6{F0: 4294967295, F1: 174, F2: 4294967295}
                        switch 1 {
                        case -538:
                            v85 := GenS3{F0: 412, F1: 0, F2: 624}
                            _ = v85
                        case 32768:
                            switch 2147483647 {
                            case 127:
                                var v86 int64 = -642
                                _ = v86
                            case -9223372036854775808:
                                if ((1000 != 599) && (((3.14159 + ((((0.0 - 1e10) / ((-2.25 - (float32((0.1 * (0.0 / (-2.25 * float64((0.1 + 0.1)))))) + 1e10)) * ((-2.25 - 0.1) + 1e10))) - 3.14159) + 1.5)) / 3.14159) > ((0.1 + -2.25) * (((((0.0 + (((1e10 / 3.14159) - ((0.0 * 3.14159) - 1.5)) - (3.14159 / (float32(0.0) + 0.1)))) - ((3.14159 / -2.25) - 3.14159)) * 3.14159) - float32(float64(1.5))) / float32(1e10))))) {
                                    v87 := GenS7{F0: fi73, F1: -1, F2: (0.0 > (1.5 * float32(float64(0.1))))}
                                    _ = v87
                                } else {
                                    fmt.Println(v55, fi73)
                                }
                            default:
                                var v88 string = "\n\t"
                                _ = v88
                            }
                        case 472:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = v84
                    } else {
                        v55 = (4294967295 != 1)
                        var v89 float64 = -2.25
                        for fi90 := 0; fi90 < 17; fi90++ {
                            var v91 float32 = (float32(v89) / 1.5)
                            var v92 int = (fi90 - (fi73 / (((fi73 >> ((293) & 63))) | 1)))
                            fmt.Println(v91)
                            _ = fi90
                            _ = v91
                            _ = v92
                        }
                        _ = v89
                    }
                    _ = fi73
                }
            }
            var v93 uint16 = 42
            if v55 {
                var v94 bool = false
                v95 := GenS8{F0: v3, F1: 758}
                v55 = (true && (42 > 2147483647))
                _ = v94
                _ = v95
            } else {
                var v96 uint32 = 2147483648
                _ = v96
            }
            _ = v93
        }
    } else {
        v97 := GenS3{F0: 1000, F1: 2, F2: 2}
        v98 := GenS8{F0: "\n\t", F1: 1}
        _ = v97
        _ = v98
    }
    if (-530 != 1) {
        if ((true || ((((-2.25 / 3.14159) - 1e10) + 0.1) < ((0.0 - ((1.5 - 1e10) / (1.5 / 1.0))) * 3.14159))) || v55) {
            v2[v3] = v55
        }
        v55 = v55
    }
    for fi99 := 0; fi99 < 4; fi99++ {
        for fi100 := 0; fi100 < 18; fi100++ {
            var v101 uint32 = 4294967295
            v102 := GenS2{F0: fi99, F1: 736, F2: (true && v55)}
            var v103 uint16 = 2
            v3 = (v3 + v3)
            switch ("\n\t" + v3) {
            case "hello":
                v103 -= 65535
                var v104 string = "a"
                _ = v104
            case "\n\t":
                var v105 int64 = 178
                _ = v105
            default:
                fmt.Println(v55)
            }
            _ = fi100
            _ = v101
            _ = v102
            _ = v103
        }
        var v106 uint16 = 255
        if (!("\n\t" < (v3 + ((((v3 + v3) + "a") + "a") + v3)))) {
            if (-128 == 42) {
                for fi107 := 0; fi107 < 4; fi107++ {
                    var v108 float64 = (3.14159 + 3.14159)
                    var v109 uint32 = 26
                    _ = fi107
                    _ = v108
                    _ = v109
                }
                for fi110 := 0; fi110 < 2; fi110++ {
                    fmt.Println(fi99, v106)
                    v3 = ("" + ((v3 + "hello") + (v3 + v3)))
                    _ = fi110
                }
                var v111 int64 = -9223372036854775808
                _ = v111
            }
            for fi112 := 0; fi112 < 16; fi112++ {
                if (1 <= -1) {
                    if (v55 && v55) {
                        fmt.Println(v106, fi112, v55)
                        v113 := v2[v3]
                        switch fi112 {
                        case 32767:
                            v2[""] = v55
                        default:
                            switch v3 {
                            case "":
                                v113 = v55
                            case "a":
                                v113 = ((v55 || v55) && (2 != 4294967295))
                            default:
                                fmt.Println(v113)
                            }
                        }
                        fmt.Println(v55)
                        v2[("hello" + "a")] = v55
                        _ = v113
                    }
                    if v55 {
                        v55 = false
                    } else {
                        switch (fi112 / ((fi112) | 1)) {
                        case 746:
                            var v114 bool = ((1000 == -1) || (v3 == "hello"))
                            _ = v114
                        case 951:
                            if v55 {
                                _ = len(v2)
                                v106 += v106
                                _ = len(v1)
                            } else {
                                v115 := GenS9{F0: v3, F1: 0}
                                var v116 string = "a"
                                if true {
                                    v106 = v106
                                    v116 = (v116 + ("" + "\n\t"))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v115
                                _ = v116
                            }
                        case -9223372036854775808:
                            switch v3 {
                            case "":
                                switch v3 {
                                case "a":
                                    var v117 uint16 = v106
                                    _ = v117
                                case "foo bar":
                                    var v118 int32 = 887
                                    _ = v118
                                default:
                                    switch fi99 {
                                    case -1:
                                        if v55 {
                                            if len(v4) > 0 {
                                                v4[0] = fi99
                                            }
                                        } else {
                                            delete(v2, ("" + ""))
                                            switch ((v3 + v3) + "a") {
                                            case "a":
                                                v119 := v2[v3]
                                                _ = v119
                                            case "foo bar":
                                                var v120 bool = ((!v55) || (((1e10 == 3.14159) || ((1e10 / 0.1) < 1.5)) || false))
                                                _ = v120
                                            }
                                            fmt.Println(fi99, v55, v55)
                                        }
                                    case 32768:
                                        if v55 {
                                            _ = len(v2)
                                            switch (fi99 & fi99) {
                                            case 1:
                                                switch v3 {
                                                case "\n\t":
                                                    fmt.Println(v106, fi112, v55)
                                                default:
                                                    var v121 int64 = 0
                                                    _ = v121
                                                }
                                            }
                                            v122 := GenS4{F0: genFn0(), F1: 128, F2: -383}
                                            _ = v122
                                        }
                                    }
                                }
                            case "hello":
                                if v55 {
                                    var v123 int = fi99
                                    v55 = (1.5 != -2.25)
                                    v124 := v2[v3]
                                    _ = v123
                                    _ = v124
                                }
                            default:
                                fmt.Println(v3, v55)
                            }
                        }
                        delete(v2, "hello")
                    }
                }
                _ = fi112
            }
        } else {
            v125 := v2["hello"]
            _ = v125
        }
        _ = fi99
        _ = v106
    }
    if v55 {
        switch -402 {
        case 9223372036854775807:
            v3 = ("" + v3)
        case 241:
            v55 = v55
        case 0:
            _ = len(v2)
        default:
            v126 := GenS8{F0: v3, F1: 1000}
            _ = v126
        }
        fmt.Println(v3)
        for fi127 := 0; fi127 < 6; fi127++ {
            var v128 float64 = -2.25
            var v129 bool = (497 <= 18446744073709551615)
            _ = fi127
            _ = v128
            _ = v129
        }
    } else {
        var v130 uint16 = 65535
        for fi131 := 0; fi131 < 17; fi131++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = ""
            for fi132 := 0; fi132 < 10; fi132++ {
                switch fi132 {
                case 32768:
                    fmt.Println(fi132, v55, v55)
                case 42:
                    _ = len(v1)
                case -313:
                    v55 = (v55 || false)
                default:
                    var v133 bool = v55
                    _ = v133
                }
                var v134 uint = 659
                v134 = v134
                _ = fi132
                _ = v134
            }
            _ = fi131
        }
        _ = v130
    }
    for fi135 := 0; fi135 < 2; fi135++ {
        for fi136 := 0; fi136 < 15; fi136++ {
            v137 := GenS5{F0: 18446744073709551615, F1: -2147483648, F2: 18446744073709551615}
            if v55 {
                if (!true) {
                    v55 = (!(-1 >= 9223372036854775807))
                    if true {
                        if true {
                            fmt.Println(fi135, fi135, fi136)
                        }
                    }
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v138 := GenS8{F0: v3, F1: -1}
                    if (v55 && v55) {
                        switch -308 {
                        case 1:
                            if (32767 >= 1) {
                                v3 = (v3 + ("" + "hello"))
                                v139 := GenS10{F0: 2, F1: -286, F2: 591}
                                v3 = (v3 + v3)
                                _ = v139
                            } else {
                                var v140 uint8 = 127
                                var v141 bool = false
                                var v142 int32 = -2147483648
                                _ = v140
                                _ = v141
                                _ = v142
                            }
                        case 2147483647:
                            v55 = ((fi136 * fi136) != (fi136 ^ (fi136 << ((195) & 63))))
                        case -60:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v143 int16 = 0
                        fmt.Println(v55, v3)
                        _ = v143
                    } else {
                        if v55 {
                            if true {
                                switch (v3 + (("a" + "") + v3)) {
                                case "hello":
                                    var v144 int = fi136
                                    v55 = v55
                                    if v55 {
                                        var v145 uint8 = 255
                                        var v146 uint8 = 53
                                        _ = v145
                                        _ = v146
                                    }
                                    _ = v144
                                case "a":
                                    v3 = v3
                                }
                            } else {
                                switch v3 {
                                case "foo bar":
                                    if v55 {
                                        fmt.Println(fi136, v55, fi136)
                                    }
                                case "hello":
                                    switch 939 {
                                    case -2147483648:
                                        fmt.Println(v55, fi135, fi136)
                                    case 127:
                                        if v55 {
                                            v3 = v3
                                            v147 := v2["foo bar"]
                                            _ = v147
                                        }
                                    }
                                }
                            }
                            if true {
                                fmt.Println(fi135, v55, fi135)
                                var v148 float32 = 3.14159
                                switch fi135 {
                                case -105:
                                    if (!v55) {
                                        switch v3 {
                                        case "foo bar":
                                            var v149 int64 = -1
                                            _ = v149
                                        case "\n\t":
                                            fmt.Println(fi135, v55, fi135)
                                        case "":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v150 string = (v3 + v3)
                                            _ = v150
                                        }
                                        v55 = v55
                                        fmt.Println(fi135, v148)
                                    }
                                }
                                _ = v148
                            } else {
                                var v151 float32 = ((3.14159 + (float32(1.5) * float32(0.1))) - 1.5)
                                _ = v151
                            }
                        } else {
                            v3 = v3
                        }
                        var v152 bool = (!true)
                        if (255 != 74) {
                            v153 := v2["hello"]
                            v3 = v3
                            _ = v153
                        }
                        delete(v2, (("a" + ((((v3 + ((("" + v3) + (v3 + v3)) + (v3 + ""))) + v3) + "a") + v3)) + "hello"))
                        var v154 int32 = 32
                        _ = v152
                        _ = v154
                    }
                    _ = v138
                }
                v55 = false
                if v55 {
                    if v55 {
                        fmt.Println(fi136, v55, fi135)
                    } else {
                        fmt.Println(v55, fi136, fi136)
                        var v155 uint64 = genFn0()
                        v156 := GenS5{F0: 255, F1: -506, F2: v155}
                        _ = v155
                        _ = v156
                    }
                    var v157 float32 = 1.5
                    switch "foo bar" {
                    case "foo bar":
                        v55 = (1 < 42)
                    }
                    _ = v157
                }
            } else {
                v3 = v3
            }
            if (!false) {
                switch v3 {
                case "foo bar":
                    v55 = false
                case "hello":
                    v158 := GenS10{F0: 42, F1: -1, F2: 896}
                    _ = v158
                }
                fmt.Println(v3, fi135)
            }
            _ = fi136
            _ = v137
        }
        if false {
            v3 = "foo bar"
            for fi159 := 0; fi159 < 1; fi159++ {
                switch -270 {
                case -1:
                    var v160 uint8 = 1
                    _ = v160
                case 128:
                    var v161 int16 = -1
                    _ = v161
                }
                if v55 {
                    var v162 float64 = -2.25
                    v55 = (-2147483648 >= -1)
                    _ = v162
                } else {
                    v3 = (v3 + (v3 + ("" + "foo bar")))
                    v163 := GenS10{F0: 4294967295, F1: 42, F2: genFn0()}
                    v55 = (v55 && (genFn1() < -583))
                    _ = v163
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi159
            }
        }
        _ = fi135
    }
    v3 = (v3 + v3)
    for fi164 := 0; fi164 < 20; fi164++ {
        for fi165 := 0; fi165 < 1; fi165++ {
            v3 = ("foo bar" + v3)
            _ = fi165
        }
        switch (v3 + (((v3 + (v3 + v3)) + v3) + (((v3 + v3) + (v3 + "a")) + (v3 + "foo bar")))) {
        case "foo bar":
            var v166 int = (fi164 << ((1) & 63))
            _ = v166
        }
        _ = fi164
    }
    v55 = (-19 <= 0)
    v55 = v55
    switch ((v3 + v3) + v3) {
    case "\n\t":
        switch "hello" {
        case "":
            v167 := GenS11{F0: 9223372036854775807}
            _ = v167
        case "a":
            var v168 uint16 = 255
            _ = v168
        case "\n\t":
            v55 = (v55 && true)
        default:
            v55 = v55
        }
    case "a":
        var v169 float32 = 1e10
        _ = v169
    case "":
        for fi170 := 0; fi170 < 13; fi170++ {
            fmt.Println(fi170, v3, v3)
            _ = fi170
        }
    default:
        var v171 float32 = float32(0.1)
        _ = v171
    }
    var v172 bool = true
    var v173 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v172
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
}
