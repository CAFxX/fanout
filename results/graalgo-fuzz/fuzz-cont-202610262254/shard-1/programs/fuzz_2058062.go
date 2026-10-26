package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int64; F2 bool }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int32; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 string; F1 uint32 }
type GenS13 struct { F0 uint64; F1 uint64; F2 int }
type GenS14 struct { F0 int; F1 uint8 }
type GenS15 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS16 struct { F0 int64; F1 uint64 }
type GenS17 struct { F0 bool; F1 string; F2 uint16 }
type GenS18 struct { F0 int8; F1 int64 }

func main() {
    var v1 int = -1
    v2 := make(map[int]bool)
    var v3 string = ("a" + "foo bar")
    var v4 uint64 = 770
    if (true || true) {
        var v5 int32 = 0
        if (127 != -128) {
            if true {
                v1 += v1
            } else {
                v6 := []int{v1, v1, v1, v1}
                v3 = v3
                v7 := GenS1{F0: 4294967295, F1: 637}
                _ = v6
                _ = v7
            }
        } else {
            switch -1 {
            case 4294967295:
                if false {
                    if false {
                        v8 := GenS1{F0: 65535, F1: 18446744073709551615}
                        var v9 float32 = 0.0
                        v3 = v3
                        _ = v8
                        _ = v9
                    }
                }
            case 128:
                var v10 int8 = 42
                _ = v10
            case -158:
                v11 := GenS1{F0: 975, F1: 65535}
                _ = v11
            }
        }
        _ = v5
    }
    var v12 uint = 0
    switch v3 {
    case "foo bar":
        if (false && ((v4 >> ((v12) & 63)) < v4)) {
            switch "a" {
            case "a":
                v3 = "foo bar"
            }
            v4 = v4
        } else {
            v13 := GenS2{F0: -1, F1: true, F2: -128}
            if false {
                var v14 uint16 = 42
                for fi15 := 0; fi15 < 14; fi15++ {
                    fmt.Println(v12, v1)
                    v16 := []int{}
                    v17 := GenS2{F0: 42, F1: false, F2: 1}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch v3 {
                    case "\n\t":
                        if (255 > 255) {
                            var v19 string = (v3 + "foo bar")
                            var v20 int32 = 2147483647
                            if false {
                                v21 := []int{(fi18 + int(v12)), (^(v1)), v1, -9223372036854775808}
                                _ = v21
                            } else {
                                _ = len(v2)
                            }
                            _ = v19
                            _ = v20
                        } else {
                            var v22 uint = 65535
                            for fi23 := 0; fi23 < 3; fi23++ {
                                v24 := GenS3{F0: -32768, F1: (v14 & 617), F2: v14}
                                if false {
                                    v12 -= (v12 & (v22 | (v12 / ((v12) | 1))))
                                } else {
                                    if false {
                                        var v25 int32 = -750
                                        _ = v25
                                    } else {
                                        v14 = 1
                                    }
                                }
                                _ = fi23
                                _ = v24
                            }
                            if (1e10 <= 1.5) {
                                v26 := GenS4{F0: (v14 & v14), F1: 970, F2: (!false)}
                                v27 := GenS5{F0: (57 < -9223372036854775808), F1: 42}
                                fmt.Println(v12, v3)
                                _ = v26
                                _ = v27
                            }
                            _ = v22
                        }
                    case "":
                        switch v3 {
                        case "a":
                            v1 += -183
                        case "foo bar":
                            var v28 uint8 = uint8(v4)
                            _ = v28
                        case "":
                            if false {
                                var v29 int16 = -32768
                                var v30 string = v3
                                _ = v29
                                _ = v30
                            } else {
                                _ = len(v2)
                            }
                        }
                    case "hello":
                        v3 = v3
                    }
                    _ = fi18
                }
                _ = v14
            }
            _ = v13
        }
    case "hello":
        if true {
            if (65535 < 2147483648) {
                for fi31 := 0; fi31 < 6; fi31++ {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        if (!false) {
                            v33 := GenS4{F0: 392, F1: 122, F2: true}
                            _ = v33
                        } else {
                            v34 := []int{fi31, 127, 128, (v1 % (((v1 | v1)) | 1))}
                            var v35 int = fi31
                            switch v3 {
                            case "a":
                                fmt.Println(v4, v1, v12)
                            default:
                                _ = len(v2)
                            }
                            if false {
                                v36 := GenS4{F0: 255, F1: -1, F2: (1 != 1)}
                                if (v3 == v3) {
                                    v37 := GenS4{F0: uint16(v4), F1: 42, F2: false}
                                    _ = v37
                                }
                                if (!("hello" > "")) {
                                    var v38 uint16 = 0
                                    _ = v38
                                } else {
                                    if len(v34) >= 2 {
                                        v34 = v34[1:len(v34)]
                                    }
                                }
                                _ = v36
                            } else {
                                var v39 int8 = 42
                                switch fi31 {
                                case -2147483648:
                                    var v40 int = fi32
                                    _ = v40
                                case 42:
                                    v39 = 0
                                case 290:
                                    if false {
                                        if (1.5 < 1.5) {
                                            v41 := GenS1{F0: 0, F1: (v12 * v12)}
                                            _ = v41
                                        } else {
                                            v42 := GenS4{F0: 1, F1: 161, F2: ((^(v12)) <= 968)}
                                            var v43 uint32 = 0
                                            _ = v42
                                            _ = v43
                                        }
                                        if (true || false) {
                                            var v44 uint32 = 1
                                            v1 = fi32
                                            v45 := GenS6{F0: -1, F1: (v4 | v4), F2: 676}
                                            _ = v44
                                            _ = v45
                                        }
                                        if (0.0 != 0.1) {
                                            v46 := GenS7{F0: (v12 ^ (v12 & (^(v12))))}
                                            v12 -= (v12 & v12)
                                            _ = v46
                                        }
                                    }
                                }
                                if ((-2.25 / 3.14159) <= (((-2.25 / 1.5) - (-2.25 * 1.5)) - ((1.5 + 1.5) / 3.14159))) {
                                    if true {
                                        var v47 float64 = 0.0
                                        var v48 bool = (v3 == ("\n\t" + (v3 + v3)))
                                        delete(v2, (fi32 - 32767))
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                _ = v39
                            }
                            _ = v34
                            _ = v35
                        }
                        switch "hello" {
                        case "":
                            switch v3 {
                            case "foo bar":
                                var v49 uint = 65535
                                _ = v49
                            case "":
                                var v50 string = (v3 + (v3 + ""))
                                _ = v50
                            case "hello":
                                var v51 int64 = -1
                                _ = v51
                            }
                        case "hello":
                            var v52 float64 = (0.0 + (1.5 - 1.5))
                            _ = v52
                        }
                        if (889 >= 42) {
                            if false {
                                var v53 uint8 = 255
                                fmt.Println(fi31, v12)
                                switch v3 {
                                case "":
                                    var v54 int16 = -32768
                                    _ = v54
                                case "\n\t":
                                    v53 = v53
                                }
                                _ = v53
                            }
                        }
                        _ = fi32
                    }
                    var v55 uint8 = 42
                    _ = fi31
                    _ = v55
                }
                if false {
                    v56 := GenS5{F0: false, F1: int64(v12)}
                    v1 = 484
                    v57 := []int{v1, (v1 | (v1 + v1))}
                    _ = v56
                    _ = v57
                } else {
                    var v58 int8 = 113
                    var v59 int8 = (v58 + 100)
                    var v60 int16 = -1
                    _ = v58
                    _ = v59
                    _ = v60
                }
            }
        } else {
            var v61 float32 = 0.0
            _ = v61
        }
    case "\n\t":
        v12 = v12
    default:
        for fi62 := 0; fi62 < 12; fi62++ {
            var v63 int = v1
            var v64 uint16 = 32767
            if ((v4 & (v4 & v4)) < v4) {
                v12 -= 0
                v3 = v3
                var v65 uint64 = 461
                _ = v65
            } else {
                switch v3 {
                case "\n\t":
                    v66 := GenS8{F0: (v4 / (((v4 | 949)) | 1)), F1: -2147483648}
                    _ = v66
                }
                var v67 uint = v12
                var v68 int16 = 1000
                _ = v67
                _ = v68
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    var v69 float64 = 1e10
    var v70 float64 = 1.5
    fmt.Println(v1)
    v71 := []int{0}
    switch v1 {
    case 32767:
        v4 = 42
    case 549:
        if false {
            if (-12 >= 0) {
                var v72 int32 = 42
                _ = v72
            } else {
                delete(v2, -2147483648)
                if len(v71) > 0 {
                    v71[0] = v1
                }
                if true {
                    if false {
                        fmt.Println(v70, v3)
                        v73 := GenS9{F0: ("hello" + v3), F1: 1}
                        v74 := GenS10{F0: 42}
                        _ = v73
                        _ = v74
                    } else {
                        if true {
                            v1 += (v1 - (v1 / ((v1) | 1)))
                            v75 := GenS4{F0: 42, F1: -1, F2: true}
                            v76 := GenS9{F0: v3, F1: 802}
                            _ = v75
                            _ = v76
                        } else {
                            _ = len(v2)
                            for fi77 := 0; fi77 < 8; fi77++ {
                                fmt.Println(v3, fi77)
                                v70 = (float64(1.5) + v70)
                                _ = fi77
                            }
                            v78 := GenS4{F0: 1, F1: 768, F2: true}
                            _ = v78
                        }
                    }
                } else {
                    switch v1 {
                    case -373:
                        v3 = (v3 + v3)
                    default:
                        if ((245 <= v12) && (v12 < v12)) {
                            fmt.Println(v1)
                        } else {
                            if (!false) {
                                v79 := v2[32767]
                                for fi80 := 0; fi80 < 2; fi80++ {
                                    switch "hello" {
                                    case "":
                                        var v81 string = v3
                                        _ = v81
                                    case "\n\t":
                                        var v82 float64 = v70
                                        _ = v82
                                    }
                                    for fi83 := 0; fi83 < 19; fi83++ {
                                        var v84 uint64 = 4294967295
                                        _ = fi83
                                        _ = v84
                                    }
                                    _ = fi80
                                }
                                _ = v79
                            } else {
                                var v85 bool = false
                                switch (v3 + (v3 + v3)) {
                                case "":
                                    v1 = (v1 & 42)
                                case "a":
                                    v86 := GenS8{F0: uint64(v12), F1: 975}
                                    _ = v86
                                case "foo bar":
                                    for fi87 := 0; fi87 < 7; fi87++ {
                                        fmt.Println(v69, v4)
                                        _ = fi87
                                    }
                                default:
                                    fmt.Println(v69, v12, v70)
                                }
                                v69 = v69
                                _ = v85
                            }
                            var v88 bool = true
                            _ = v88
                        }
                    }
                    v12 = 601
                    var v89 int = (v1 & v1)
                    _ = v89
                }
            }
            fmt.Println(v69, v69)
            if true {
                if (-273 >= 590) {
                    var v90 bool = (0 >= 11)
                    delete(v2, (v1 ^ v1))
                    _ = v90
                } else {
                    if true {
                        var v91 int16 = 1000
                        if false {
                            fmt.Println(v70, v69)
                        } else {
                            var v92 uint8 = 0
                            v3 = v3
                            _ = v92
                        }
                        v93 := GenS11{F0: -1}
                        _ = v91
                        _ = v93
                    } else {
                        for fi94 := 0; fi94 < 1; fi94++ {
                            _ = len(v71)
                            if false {
                                switch v3 {
                                case "foo bar":
                                    var v95 uint = v12
                                    _ = v95
                                case "hello":
                                    v2[(fi94 * (^(v1)))] = (42 != -1)
                                case "a":
                                    v3 = ("\n\t" + "foo bar")
                                default:
                                    v70 -= 1e10
                                }
                                if ("foo bar" < ((v3 + v3) + (v3 + ""))) {
                                    v96 := GenS12{F0: v3, F1: 0}
                                    _ = v96
                                } else {
                                    var v97 uint = v12
                                    v4 = v4
                                    var v98 bool = true
                                    _ = v97
                                    _ = v98
                                }
                                switch fi94 {
                                case 762:
                                    v70 = float64(1e10)
                                case 128:
                                    v99 := GenS13{F0: 42, F1: (v4 - v4), F2: int(v4)}
                                    _ = v99
                                default:
                                    var v100 uint16 = 2
                                    _ = v100
                                }
                            } else {
                                v69 = v70
                                v12 = (v12 & v12)
                            }
                            _ = fi94
                        }
                    }
                }
            }
        }
    case 127:
        var v101 int = (v1 * (v1 | 4294967295))
        _ = v101
    default:
        switch 2147483647 {
        case 42:
            var v102 int8 = 64
            _ = v102
        case 9223372036854775807:
            v103 := GenS14{F0: (v1 * (v1 % ((v1) | 1))), F1: 128}
            _ = v103
        case -1:
            if true {
                if (623 >= 1) {
                    v1 = (v1 - v1)
                    for fi104 := 0; fi104 < 4; fi104++ {
                        v70 = v69
                        _ = fi104
                    }
                } else {
                    v4 = (v4 & v4)
                    var v105 uint = v12
                    for fi106 := 0; fi106 < 17; fi106++ {
                        for fi107 := 0; fi107 < 7; fi107++ {
                            _ = len(v2)
                            var v108 int32 = 0
                            v2[-2147483648] = ((false || false) || (2 > 2))
                            _ = fi107
                            _ = v108
                        }
                        _ = fi106
                    }
                    _ = v105
                }
            } else {
                var v109 uint8 = 0
                var v110 int8 = 88
                _ = v109
                _ = v110
            }
        }
    }
    if (false && false) {
        v12 -= v12
        v12 = (v12 ^ v12)
        switch v1 {
        case 128:
            v111 := GenS10{F0: 21}
            _ = v111
        case 127:
            _ = len(v2)
        }
    }
    if ((false && true) || true) {
        var v112 float64 = v69
        _ = v112
    }
    var v113 uint8 = 110
    if true {
        if (!(42 <= 32767)) {
            if (false || false) {
                switch v3 {
                case "a":
                    if false {
                        var v114 uint8 = v113
                        switch v3 {
                        case "a":
                            if (0 < 127) {
                                v115 := v2[v1]
                                if (v114 != (v114 - (v114 * 2))) {
                                    if v115 {
                                        var v116 uint = (^(v12))
                                        _ = v116
                                    } else {
                                        v117 := GenS1{F0: (v12 << (((v12 >> ((4294967295) & 63))) & 63)), F1: v12}
                                        if v115 {
                                            var v118 int64 = -699
                                            var v119 int32 = 128
                                            switch ("a" + v3) {
                                            case "hello":
                                                var v120 uint16 = 42
                                                _ = v120
                                            default:
                                                for fi121 := 0; fi121 < 10; fi121++ {
                                                    if ((9223372036854775807 == 128) || (-2.25 < v69)) {
                                                        for fi122 := 0; fi122 < 2; fi122++ {
                                                            fmt.Println(v69, v114, v113)
                                                            if (4294967295 < 569) {
                                                                v115 = (v115 || (!v115))
                                                                v123 := v2[(-(v1))]
                                                                v3 = v3
                                                                v115 = (28 == 2)
                                                                _ = v123
                                                            } else {
                                                                v118 = (v118 >> ((v12) & 63))
                                                                var v124 int32 = v119
                                                                var v125 bool = (fi121 != (-(fi122)))
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                            var v126 int16 = -676
                                                            var v127 uint = 2
                                                            _ = fi122
                                                            _ = v126
                                                            _ = v127
                                                        }
                                                        v128 := GenS4{F0: uint16(v4), F1: (v118 - v118), F2: v115}
                                                        _ = v128
                                                    }
                                                    var v129 int16 = int16(v119)
                                                    _ = fi121
                                                    _ = v129
                                                }
                                            }
                                            _ = v118
                                            _ = v119
                                        }
                                        _ = v117
                                    }
                                    if true {
                                        for fi130 := 0; fi130 < 15; fi130++ {
                                            var v131 uint = v12
                                            _ = fi130
                                            _ = v131
                                        }
                                    } else {
                                        v69 = -2.25
                                    }
                                    for fi132 := 0; fi132 < 19; fi132++ {
                                        v113 -= v113
                                        var v133 string = v3
                                        var v134 uint8 = 255
                                        _ = fi132
                                        _ = v133
                                        _ = v134
                                    }
                                } else {
                                    v12 = v12
                                }
                                delete(v2, (^(v1)))
                                _ = v115
                            } else {
                                if false {
                                    for fi135 := 0; fi135 < 1; fi135++ {
                                        var v136 bool = (true && (((3.14159 == -2.25) || (!false)) || (1 <= 0)))
                                        v137 := GenS15{F0: v4, F1: uint16(v12), F2: 123}
                                        _ = fi135
                                        _ = v136
                                        _ = v137
                                    }
                                    fmt.Println(v3)
                                } else {
                                    if true {
                                        for fi138 := 0; fi138 < 9; fi138++ {
                                            v139 := GenS16{F0: -431, F1: (v4 >> ((v12) & 63))}
                                            _ = fi138
                                            _ = v139
                                        }
                                        v71 = append(v71, (v1 ^ v1))
                                    } else {
                                        var v140 int64 = -9223372036854775808
                                        delete(v2, (^(v1)))
                                        fmt.Println(v70)
                                        _ = v140
                                    }
                                    if len(v71) >= 2 {
                                        v71 = v71[1:len(v71)]
                                    }
                                    for fi141 := 0; fi141 < 5; fi141++ {
                                        v113 -= uint8(v12)
                                        v142 := GenS17{F0: (float64(1.5) >= (v70 / (v70 - float64((((float32(((3.14159 * v70) / 1.0)) + (float32(1.5) * (((float32((v70 - 3.14159)) * ((3.14159 / 3.14159) - float32(v69))) - 1e10) - -2.25))) - 1.5) - 0.0))))), F1: v3, F2: 42}
                                        _ = fi141
                                        _ = v142
                                    }
                                }
                                switch "a" {
                                case "":
                                    fmt.Println(v12, v12, v4)
                                default:
                                    for fi143 := 0; fi143 < 20; fi143++ {
                                        var v144 bool = false
                                        fmt.Println(v114, v144)
                                        _ = fi143
                                        _ = v144
                                    }
                                }
                                v145 := v2[(v1 >> (((v12 / ((1) | 1))) & 63))]
                                _ = v145
                            }
                        }
                        _ = v114
                    } else {
                        v69 += 3.14159
                        v146 := GenS18{F0: 100, F1: 1}
                        _ = v146
                    }
                case "foo bar":
                    var v147 float64 = (float64((float32(1e10) - float32(v70))) - 0.0)
                    _ = v147
                }
                if len(v71) > 0 {
                    v71[0] = (v1 & v1)
                }
                _ = len(v2)
            }
            if ((1e10 * (float32(v70) / -2.25)) >= (((float32((((v70 - v70) + v70) - (0.1 - v70))) * -2.25) / (float32((1e10 - ((1.5 * 0.0) * float64((0.0 - 1.5))))) - (3.14159 + 0.0))) / 1.0)) {
                _ = len(v71)
                for fi148 := 0; fi148 < 4; fi148++ {
                    if false {
                        switch v3 {
                        case "\n\t":
                            var v149 int32 = -2147483648
                            _ = v149
                        case "foo bar":
                            fmt.Println(v69)
                        case "hello":
                            for fi150 := 0; fi150 < 11; fi150++ {
                                var v151 string = v3
                                v3 = v3
                                _ = fi150
                                _ = v151
                            }
                        }
                        var v152 int32 = 42
                        var v153 int8 = 42
                        _ = v152
                        _ = v153
                    }
                    _ = fi148
                }
            }
        }
    }
    v71 = append(v71, (v1 / ((1) | 1)))
    var v154 int16 = 1
    fmt.Println(v113, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v12:", v12)
    fmt.Println("v154:", v154)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v113
    _ = v12
    _ = v154
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v71
}
