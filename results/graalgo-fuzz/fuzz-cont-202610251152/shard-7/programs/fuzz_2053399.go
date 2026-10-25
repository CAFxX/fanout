package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int32; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint64; F1 int64; F2 bool }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint64; F1 uint8 }
type GenS13 struct { F0 int16; F1 int; F2 uint16 }
type GenS14 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int64; F1 uint16 }
func genFn0() int8 {
    var v5 int = 580
    if false {
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                switch v5 {
                case -618:
                    v8 := []int{(^(v5))}
                    _ = v8
                case 42:
                    fmt.Println(v5, fi7, fi6)
                    fmt.Println(fi6)
                    fmt.Println(v5, fi7)
                    v5 -= (fi7 - v5)
                case -395:
                    v5 = v5
                }
                if (false || (((!false) && ((fi6 & (v5 / (((^(fi7))) | 1))) > (fi6 >> ((42) & 63)))) && false)) {
                    var v9 uint16 = 0
                    v10 := GenS1{F0: 255, F1: 1, F2: 0}
                    _ = v10
                    _ = v9
                } else {
                    var v11 uint16 = 771
                    _ = v11
                }
                var v12 uint8 = 255
                _ = fi7
                _ = v12
            }
            v13 := []int{fi6}
            v5 = v5
            _ = fi6
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(v5, fi14, fi15)
                v16 := GenS1{F0: 1, F1: -812, F2: 2}
                v17 := GenS1{F0: 128, F1: 42, F2: 158}
                v18 := GenS2{F0: 971, F1: 450}
                v19 := GenS3{F0: 255}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                switch "a" {
                case "\n\t":
                    var v21 bool = false
                    _ = v21
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi14)
                    case "foo bar":
                        var v22 string = "foo bar"
                        _ = v22
                    case "a":
                        v23 := GenS3{F0: 1}
                        _ = v23
                    default:
                        v5 += fi14
                    }
                    v24 := []int{(v5 & fi14), 0, fi20}
                    v25 := GenS4{F0: 0, F1: 2, F2: 533}
                    var v26 int16 = -389
                    _ = v24
                    _ = v25
                    _ = v26
                case "":
                    if ((100 != 1) && (false && true)) {
                        if (794 < 867) {
                            var v27 uint16 = 674
                            var v28 float64 = 1.5
                            var v29 uint64 = 21
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            if (1 <= 422) {
                                switch "foo bar" {
                                case "\n\t":
                                    v30 := []int{(v5 + 32768), -559, 2147483647}
                                    var v31 int32 = 1
                                    _ = v30
                                    _ = v31
                                case "foo bar":
                                    if (((-(v5)) < v5) || (!true)) {
                                        if (true || true) {
                                            var v32 string = "a"
                                            if false {
                                                switch (^(fi14)) {
                                                case 917:
                                                    v33 := GenS5{F0: -1, F1: 2147483648}
                                                    fmt.Println(v5, v32, fi14)
                                                    var v34 uint32 = 42
                                                    _ = v33
                                                    _ = v34
                                                case 128:
                                                    var v35 uint16 = 186
                                                    _ = v35
                                                default:
                                                    v32 = v32
                                                }
                                                v36 := []int{(fi20 | (fi20 * -368))}
                                                _ = v36
                                            } else {
                                                v37 := GenS5{F0: 42, F1: 2}
                                                var v38 bool = (!(18446744073709551615 == 1))
                                                var v39 float32 = 1.5
                                                switch ((((v32 + v32) + "") + (v32 + "hello")) + v32) {
                                                case "\n\t":
                                                    if false {
                                                        var v40 int = fi20
                                                        if false {
                                                            var v41 uint = 42
                                                            v42 := []int{v40, 32767, (fi14 >> ((v41) & 63)), (v40 * fi14)}
                                                            v41 = 625
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        _ = v40
                                                    } else {
                                                        if (4294967295 >= 0) {
                                                            v39 = 3.14159
                                                            v39 = (v39 * v39)
                                                            var v43 int16 = -1
                                                            _ = v43
                                                        } else {
                                                            var v44 bool = v38
                                                            v45 := GenS6{F0: 1}
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        v46 := GenS2{F0: 2, F1: 1}
                                                        _ = v46
                                                    }
                                                case "":
                                                    if (1 < 0) {
                                                        v39 -= 3.14159
                                                        fmt.Println(v5, v38)
                                                    }
                                                default:
                                                    if true {
                                                        v38 = (v38 && v38)
                                                        v47 := GenS7{F0: 925, F1: v5}
                                                        v5 += 2147483647
                                                        _ = v47
                                                    } else {
                                                        v48 := []int{(fi14 + (v5 << ((0) & 63)))}
                                                        _ = v48
                                                    }
                                                }
                                                if (1000 == 0) {
                                                    switch ("" + (v32 + "hello")) {
                                                    case "a":
                                                        var v49 int32 = -2147483648
                                                        _ = v49
                                                    case "\n\t":
                                                        switch (fi20 - (v5 & -2147483648)) {
                                                        case 2147483647:
                                                            if (0.1 > (0.0 - 3.14159)) {
                                                                switch fi14 {
                                                                case 965:
                                                                    var v50 uint = 18446744073709551615
                                                                    _ = v50
                                                                }
                                                                v51 := GenS8{F0: 1, F1: 171, F2: 1}
                                                                fmt.Println(fi20, fi20)
                                                                _ = v51
                                                            } else {
                                                                var v52 float64 = (float64(v39) * -2.25)
                                                                fmt.Println(fi14)
                                                                _ = v52
                                                            }
                                                        case 9223372036854775807:
                                                            if (100 >= -128) {
                                                                var v53 float64 = -2.25
                                                                v38 = v38
                                                                v39 = 3.14159
                                                                v39 += 1.5
                                                                fmt.Println(v39, v5)
                                                                _ = v53
                                                            } else {
                                                                if (1.5 < 0.1) {
                                                                    v5 = 2147483647
                                                                    v32 = ""
                                                                }
                                                            }
                                                        default:
                                                            v54 := []int{v5}
                                                            _ = v54
                                                        }
                                                    case "":
                                                        v55 := []int{(fi14 + (fi14 + (fi14 % ((9223372036854775807) | 1))))}
                                                        _ = v55
                                                    }
                                                    switch -598 {
                                                    case -9223372036854775808:
                                                        v5 = (fi14 ^ (v5 * (v5 % ((-9223372036854775808) | 1))))
                                                    }
                                                    switch (fi14 | (fi20 | (fi20 - (fi14 / (((fi14 + (v5 % (((^(fi14))) | 1)))) | 1))))) {
                                                    case 127:
                                                        if (!(-947 <= 1)) {
                                                            if false {
                                                                v38 = (v38 || v38)
                                                            } else {
                                                                var v56 bool = (72 != -128)
                                                                _ = v56
                                                            }
                                                            v32 = "a"
                                                        }
                                                        v57 := GenS1{F0: 1, F1: -381, F2: 255}
                                                        v39 += 0.0
                                                        if false {
                                                            v5 = v5
                                                            var v58 float32 = v39
                                                            v59 := []int{(fi20 & fi20), v5}
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v39 += v39
                                                            var v60 uint64 = 2
                                                            _ = v60
                                                        }
                                                        v61 := []int{0, v5, (fi20 * v5), (fi14 >> ((642) & 63))}
                                                        fmt.Println(fi14, fi20, v38)
                                                        var v62 uint64 = 176
                                                        var v63 float32 = v39
                                                        var v64 int = 1
                                                        _ = v57
                                                        _ = v61
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    case -869:
                                                        v39 = v39
                                                    }
                                                }
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v32
                                        } else {
                                            switch (("\n\t" + "foo bar") + "foo bar") {
                                            case "foo bar":
                                                var v65 uint = 255
                                                _ = v65
                                            case "\n\t":
                                                switch v5 {
                                                case 4294967295:
                                                    switch 801 {
                                                    case 4294967295:
                                                        v5 = fi14
                                                    case -1:
                                                        var v66 string = ((("\n\t" + "hello") + "\n\t") + "\n\t")
                                                        _ = v66
                                                    case -2147483648:
                                                        var v67 string = "foo bar"
                                                        _ = v67
                                                    default:
                                                        v5 += fi14
                                                    }
                                                case 32768:
                                                    if ((100 < 42) || true) {
                                                        switch "a" {
                                                        case "":
                                                            v5 -= v5
                                                        case "a":
                                                            v68 := []int{(v5 << ((255) & 63))}
                                                            _ = v68
                                                        case "hello":
                                                            switch fi20 {
                                                            case 916:
                                                                var v69 bool = false
                                                                _ = v69
                                                            }
                                                        default:
                                                            if false {
                                                                if (!(-9223372036854775808 != 42)) {
                                                                    if ((("\n\t" == ("foo bar" + "hello")) && true) && true) {
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            v70 := GenS6{F0: 127}
                                                                            _ = v70
                                                                        case "hello":
                                                                            var v71 int8 = -128
                                                                            _ = v71
                                                                        case "a":
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                var v72 uint16 = 32767
                                                                                _ = v72
                                                                            case "":
                                                                                v5 -= fi14
                                                                            default:
                                                                                v73 := []int{(^(fi20)), (-(fi20)), fi20}
                                                                                _ = v73
                                                                            }
                                                                        default:
                                                                            fmt.Println(v5)
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi14, fi14)
                                                                        var v74 string = "a"
                                                                        var v75 uint = 732
                                                                        if false {
                                                                            if (!false) {
                                                                                if true {
                                                                                    var v76 float64 = float64((((0.1 + 0.1) / 0.1) + 1.5))
                                                                                    if true {
                                                                                        var v77 float32 = 3.14159
                                                                                        if ((2 != 2) && true) {
                                                                                            var v78 bool = ((true || true) && true)
                                                                                            v74 = "a"
                                                                                            _ = v78
                                                                                        } else {
                                                                                            v79 := GenS9{F0: 629, F1: 1, F2: (42 > -9223372036854775808)}
                                                                                            _ = v79
                                                                                        }
                                                                                        _ = v77
                                                                                    } else {
                                                                                        v80 := []int{-65, 128}
                                                                                        if len(v80) > 0 {
                                                                                            v80[0] = 32768
                                                                                        }
                                                                                        _ = v80
                                                                                    }
                                                                                    v81 := GenS3{F0: 255}
                                                                                    _ = v76
                                                                                    _ = v81
                                                                                }
                                                                                v82 := []int{-409, fi14, -2147483648, (fi14 & (fi20 * (-(fi14))))}
                                                                                switch (fi20 * fi14) {
                                                                                case -1:
                                                                                    v75 -= (v75 ^ (^(v75)))
                                                                                default:
                                                                                    v83 := GenS10{F0: v75, F1: 0}
                                                                                    _ = v83
                                                                                }
                                                                                _ = len(v82)
                                                                                _ = v82
                                                                            } else {
                                                                                v84 := GenS4{F0: 0, F1: 1, F2: 795}
                                                                                v5 -= (v5 & 4294967295)
                                                                                if (92 < 1) {
                                                                                    fmt.Println(v75, v74)
                                                                                    switch -2147483648 {
                                                                                    case 2147483647:
                                                                                        v85 := []int{(fi14 % (((fi14 & 586)) | 1)), -1, v5, 127}
                                                                                        _ = v85
                                                                                    case 32767:
                                                                                        if true {
                                                                                            v86 := GenS5{F0: int16(v75), F1: 0}
                                                                                            _ = v86
                                                                                        } else {
                                                                                            v75 += 485
                                                                                            var v87 uint32 = 4294967295
                                                                                            switch ("\n\t" + v74) {
                                                                                            case "":
                                                                                                var v88 uint64 = 838
                                                                                                _ = v88
                                                                                            }
                                                                                            _ = v87
                                                                                        }
                                                                                    case -2147483648:
                                                                                        var v89 int8 = -14
                                                                                        v90 := GenS6{F0: 0}
                                                                                        v91 := GenS11{F0: -1}
                                                                                        v5 -= fi20
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    default:
                                                                                        switch v74 {
                                                                                        case "":
                                                                                            var v92 int16 = 32767
                                                                                            _ = v92
                                                                                        case "\n\t":
                                                                                            v93 := []int{(fi14 ^ (v5 + fi14)), (fi14 << ((65535) & 63)), 128, 9223372036854775807}
                                                                                            _ = v93
                                                                                        case "foo bar":
                                                                                            if ((((!(fi14 != v5)) && false) || (-20 >= -98)) && true) {
                                                                                                v94 := GenS1{F0: 255, F1: 0, F2: 2}
                                                                                                _ = v94
                                                                                            } else {
                                                                                                v95 := []int{(v5 >> ((v75) & 63)), (v5 % ((-558) | 1)), (v5 ^ (v5 * (v5 / ((-2147483648) | 1)))), 1}
                                                                                                _ = v95
                                                                                            }
                                                                                        default:
                                                                                            v96 := GenS12{F0: 2, F1: 236}
                                                                                            _ = v96
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v97 := GenS13{F0: -804, F1: (fi14 | fi20), F2: 2}
                                                                                    v98 := []int{2147483647, (-(fi20)), (fi20 >> (((v75 >> (((v75 & 42)) & 63))) & 63)), fi20}
                                                                                    var v99 float64 = 0.0
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                                _ = v84
                                                                            }
                                                                            var v100 int16 = -1
                                                                            _ = v100
                                                                        } else {
                                                                            v5 = (v5 | v5)
                                                                            var v101 uint64 = 311
                                                                            _ = v101
                                                                        }
                                                                        v75 += v75
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                    v102 := GenS5{F0: -1, F1: 2147483648}
                                                                    _ = v102
                                                                } else {
                                                                    v103 := GenS11{F0: 126}
                                                                    v5 += (-(fi20))
                                                                    _ = v103
                                                                }
                                                            }
                                                        }
                                                        var v104 float32 = 1.5
                                                        var v105 int64 = -9223372036854775808
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        switch v5 {
                                        case 127:
                                            v5 = fi14
                                        case -959:
                                            fmt.Println(fi14, v5, fi14)
                                        case 129:
                                            v106 := []int{fi14, (fi20 | (^(fi20)))}
                                            _ = v106
                                        default:
                                            if (!false) {
                                                v5 += (fi20 | fi14)
                                            } else {
                                                v5 = 0
                                                var v107 uint8 = 115
                                                fmt.Println(v5, v107)
                                                _ = v107
                                            }
                                        }
                                        switch fi20 {
                                        case 32767:
                                            fmt.Println(v5)
                                        case 928:
                                            if (!false) {
                                                var v108 int32 = 1
                                                _ = v108
                                            } else {
                                                var v109 uint = 577
                                                _ = v109
                                            }
                                        case -334:
                                            v110 := GenS13{F0: 1, F1: (fi20 >> ((42) & 63)), F2: 2}
                                            _ = v110
                                        }
                                        if (1e10 >= 0.1) {
                                            var v111 int32 = -566
                                            v112 := []int{fi20, v5, (fi20 / ((2147483647) | 1)), fi20}
                                            _ = v111
                                            _ = v112
                                        } else {
                                            var v113 uint64 = 18446744073709551615
                                            v113 = v113
                                            _ = v113
                                        }
                                        var v114 int16 = 1
                                        _ = v114
                                    }
                                }
                                v5 = v5
                            } else {
                                var v115 uint = 327
                                v116 := []int{}
                                _ = v115
                                _ = v116
                            }
                        }
                        v117 := GenS5{F0: 1, F1: 375}
                        var v118 uint8 = 120
                        _ = v117
                        _ = v118
                    } else {
                        v5 -= fi14
                        var v119 uint = 42
                        if true {
                            fmt.Println(v5, fi14)
                            v120 := GenS6{F0: 255}
                            _ = v120
                        }
                        _ = v119
                    }
                }
                v121 := GenS4{F0: 1000, F1: 1, F2: 42}
                switch "foo bar" {
                case "hello":
                    if (0 >= 0) {
                        v5 = fi14
                        v5 += (fi14 * (fi20 - 42))
                        fmt.Println(fi20, v5, fi14)
                    }
                }
                var v122 string = ""
                switch v122 {
                case "hello":
                    var v123 uint32 = 42
                    _ = v123
                }
                _ = fi20
                _ = v121
                _ = v122
            }
            switch ((((("\n\t" + "") + "hello") + "foo bar") + "foo bar") + "foo bar") {
            case "hello":
                var v124 string = ""
                _ = v124
            default:
                var v125 uint = 18446744073709551615
                _ = v125
            }
            _ = fi14
        }
    } else {
        var v126 int16 = 1000
        _ = v126
    }
    v5 += v5
    _ = v5
    return 0
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 2
    v4 -= (^(v4))
    v127 := GenS14{F0: 142, F1: 1, F2: 2147483647}
    for fi128 := 0; fi128 < 5; fi128++ {
        fmt.Println(v1, fi128, v1)
        fmt.Println(v1, fi128)
        _ = fi128
    }
    var v129 uint16 = 526
    if (!false) {
        var v130 int8 = -128
        for fi131 := 0; fi131 < 9; fi131++ {
            _ = len(v3)
            _ = fi131
        }
        fmt.Println(v4, v129)
        _ = v130
    }
    if (!(1000 <= -1)) {
        v129 += v129
        for fi132 := 0; fi132 < 7; fi132++ {
            var v133 uint32 = v4
            if (-2.25 <= float64(1.5)) {
                var v134 int16 = -37
                var v135 bool = (-9223372036854775808 <= 42)
                if v135 {
                    v1 = v1
                    v136 := GenS15{F0: 127}
                    _ = v136
                }
                _ = v134
                _ = v135
            } else {
                for fi137 := 0; fi137 < 18; fi137++ {
                    v1 += v1
                    _ = fi137
                }
                var v138 uint32 = (v4 / ((v4) | 1))
                v4 -= (v138 * (v138 / ((v133) | 1)))
                _ = v138
            }
            _ = fi132
            _ = v133
        }
    }
    for fi139 := 0; fi139 < 5; fi139++ {
        v140 := GenS16{F0: 1, F1: v129}
        fmt.Println(v4, v4, v1)
        var v141 float64 = 1e10
        _ = fi139
        _ = v140
        _ = v141
    }
    v142 := GenS6{F0: 127}
    var v143 int64 = 9223372036854775807
    var v144 uint16 = (^(v129))
    v144 -= (v144 & (v129 >> ((18446744073709551615) & 63)))
    var v145 int64 = (v143 & v143)
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v127
    _ = v129
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v2
    _ = v3
    _ = v4
}
