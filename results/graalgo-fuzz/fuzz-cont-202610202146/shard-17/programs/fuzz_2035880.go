package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int8; F2 string }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 string; F2 uint64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint64; F1 uint }
type GenS14 struct { F0 int32; F1 int8; F2 string }
type GenS15 struct { F0 uint8; F1 uint16; F2 uint }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int64; F1 string }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int8; F1 uint8 }
type GenS21 struct { F0 uint; F1 string; F2 bool }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 bool; F1 int32; F2 int64 }
type GenS24 struct { F0 uint16; F1 int32; F2 uint }
type GenS25 struct { F0 uint; F1 int16 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 bool; F1 int32; F2 bool }
type GenS28 struct { F0 int8; F1 int8 }
type GenS29 struct { F0 int; F1 uint }
type GenS30 struct { F0 int8; F1 uint16 }
type GenS31 struct { F0 int32 }
type GenS32 struct { F0 int16; F1 uint8 }
type GenS33 struct { F0 int32; F1 uint8; F2 int }
type GenS34 struct { F0 int16; F1 int16; F2 int8 }
type GenS35 struct { F0 int8; F1 uint16; F2 string }
type GenS36 struct { F0 int32; F1 int }
type GenS37 struct { F0 int; F1 uint16 }
type GenS38 struct { F0 int8 }
type GenS39 struct { F0 int8; F1 uint32; F2 int16 }
type GenS40 struct { F0 uint8 }
type GenS41 struct { F0 uint32 }
func genFn0(p0 uint, p1 int64) uint {
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            for fi7 := 0; fi7 < 14; fi7++ {
                if (29 != 0) {
                    v8 := GenS1{F0: "foo bar", F1: 255, F2: 2147483647}
                    p1 += 0
                    _ = v8
                } else {
                    v9 := GenS2{F0: -1, F1: (p0 + p0)}
                    if true {
                        p1 = (p1 & (p1 + 665))
                        var v10 uint = (p0 | p0)
                        p1 += (p1 >> ((1) & 63))
                        _ = v10
                    } else {
                        v11 := []int{(fi7 / ((fi7) | 1)), fi7, fi6}
                        _ = v11
                    }
                    if true {
                        var v12 float32 = 3.14159
                        v13 := GenS3{F0: 255}
                        v14 := GenS2{F0: -9223372036854775808, F1: p0}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = v9
                }
                var v15 int = -2147483648
                _ = fi7
                _ = v15
            }
        }
        if (((0.1 / (1.5 * 0.1)) + float32(3.14159)) != 0.0) {
            var v16 uint32 = 924
            var v17 uint8 = 103
            v18 := []int{(fi6 >> (((p0 + 18446744073709551615)) & 63)), fi6, (fi6 + (fi6 << ((p0) & 63))), (fi6 & (fi6 + 0))}
            _ = v16
            _ = v17
            _ = v18
        }
        v19 := []int{32767, (^(fi6)), 705}
        _ = fi6
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int16 = -1
        switch fi20 {
        case -51:
            p0 += (p0 & 42)
        }
        _ = fi20
        _ = v21
    }
    var v22 int16 = 1000
    switch "foo bar" {
    case "hello":
        switch -998 {
        case 1:
            for fi23 := 0; fi23 < 5; fi23++ {
                if (true && (2 <= 32767)) {
                    v24 := []int{0, fi23, 32768, (fi23 % ((128) | 1))}
                    p1 = (p1 % (((p1 >> ((p0) & 63))) | 1))
                    _ = v24
                } else {
                    fmt.Println(p1)
                    p1 = (p1 << ((p0) & 63))
                }
                _ = fi23
            }
        case 9223372036854775807:
            var v25 uint32 = 1
            _ = v25
        case 0:
            fmt.Println(p1)
        default:
            var v26 uint8 = 141
            _ = v26
        }
    case "\n\t":
        var v27 float64 = float64((0.1 / ((3.14159 - -2.25) - float32((3.14159 + float64(0.1))))))
        _ = v27
    case "":
        fmt.Println(v22, p1, v22)
    default:
        switch ("foo bar" + "") {
        case "a":
            p1 -= p1
        case "hello":
            for fi28 := 0; fi28 < 12; fi28++ {
                switch fi28 {
                case -1:
                    switch "" {
                    case "a":
                        var v29 uint16 = 0
                        _ = v29
                    case "hello":
                        v30 := GenS1{F0: ("a" + "foo bar"), F1: 576, F2: 2147483647}
                        _ = v30
                    default:
                        var v31 uint8 = 2
                        _ = v31
                    }
                case 2147483647:
                    v32 := []int{fi28, (fi28 ^ (fi28 | 1)), 444}
                    _ = v32
                case 4294967295:
                    if (650 >= 1) {
                        for fi33 := 0; fi33 < 10; fi33++ {
                            v22 = (v22 & 646)
                            var v34 uint64 = uint64(p0)
                            switch (fi28 * fi33) {
                            case 210:
                                p1 -= p1
                            case -9223372036854775808:
                                var v35 int64 = (p1 + p1)
                                _ = v35
                            case 42:
                                fmt.Println(p1, fi28)
                            default:
                                if true {
                                    if true {
                                        if (1.5 <= float32(float64(0.1))) {
                                            var v36 float64 = 1.5
                                            _ = v36
                                        } else {
                                            var v37 int32 = 42
                                            p1 = p1
                                            _ = v37
                                        }
                                    }
                                    v38 := GenS4{F0: 127}
                                    _ = v38
                                } else {
                                    if true {
                                        fmt.Println(fi28)
                                    } else {
                                        var v39 string = "hello"
                                        _ = v39
                                    }
                                    switch fi33 {
                                    case 32767:
                                        if ((((1.5 * -2.25) + 1.5) + (3.14159 + (1e10 * 0.0))) > 1e10) {
                                            fmt.Println(fi28, v22)
                                            var v40 string = "\n\t"
                                            _ = v40
                                        } else {
                                            var v41 bool = true
                                            var v42 int8 = -90
                                            _ = v41
                                            _ = v42
                                        }
                                    case 0:
                                        v43 := []int{(fi33 >> ((0) & 63))}
                                        _ = v43
                                    default:
                                        if false {
                                            var v44 float32 = float32(1e10)
                                            _ = v44
                                        } else {
                                            switch fi28 {
                                            case 0:
                                                var v45 int8 = 94
                                                _ = v45
                                            }
                                            fmt.Println(p0)
                                        }
                                    }
                                }
                            }
                            v46 := []int{(fi33 + fi33), fi33, -1}
                            var v47 int8 = 127
                            _ = fi33
                            _ = v34
                            _ = v46
                            _ = v47
                        }
                    } else {
                        v48 := GenS5{F0: 2}
                        _ = v48
                    }
                default:
                    v49 := GenS6{F0: ((-(p1)) == p1), F1: (-(p1)), F2: 342}
                    _ = v49
                }
                v50 := GenS5{F0: 1}
                _ = fi28
                _ = v50
            }
        case "":
            if true {
                for fi51 := 0; fi51 < 7; fi51++ {
                    fmt.Println(fi51, fi51, v22)
                    v52 := GenS6{F0: (!(!(p1 > int64(p0)))), F1: (p1 % (((-(p1))) | 1)), F2: 0}
                    var v53 float64 = 0.1
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                if ((0 != 9223372036854775807) && (158 == 42)) {
                    fmt.Println(p1, p1, v22)
                    var v54 int32 = 42
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 2; fi55++ {
                        for fi56 := 0; fi56 < 15; fi56++ {
                            fmt.Println(fi56)
                            p1 = p1
                            v57 := []int{fi55, (fi55 & 339), (-(fi55))}
                            _ = fi56
                            _ = v57
                        }
                        for fi58 := 0; fi58 < 10; fi58++ {
                            var v59 int64 = p1
                            _ = fi58
                            _ = v59
                        }
                        var v60 int64 = 9223372036854775807
                        _ = fi55
                        _ = v60
                    }
                    switch "\n\t" {
                    case "foo bar":
                        for fi61 := 0; fi61 < 8; fi61++ {
                            v22 += (v22 >> ((p0) & 63))
                            if ((((!(true || false)) || (p1 == (p1 & (^(p1))))) && true) || true) {
                                v62 := []int{331}
                                _ = v62
                            }
                            if (-2.25 <= 3.14159) {
                                if true {
                                    v22 = -1
                                    p1 = (p1 * 1)
                                } else {
                                    if (1 != uint32(p0)) {
                                        p1 = (-(p1))
                                        for fi63 := 0; fi63 < 7; fi63++ {
                                            v22 -= 556
                                            var v64 int16 = 172
                                            _ = fi63
                                            _ = v64
                                        }
                                        for fi65 := 0; fi65 < 9; fi65++ {
                                            switch ("" + (("a" + "") + ("" + "a"))) {
                                            case "hello":
                                                v66 := []int{(fi61 | fi65)}
                                                _ = v66
                                            case "foo bar":
                                                v67 := GenS7{F0: 127, F1: 42, F2: "foo bar"}
                                                _ = v67
                                            case "\n\t":
                                                v68 := []int{}
                                                _ = v68
                                            default:
                                                p1 += p1
                                            }
                                            p1 = (p1 % ((p1) | 1))
                                            switch "" {
                                            case "hello":
                                                if true {
                                                    switch 32768 {
                                                    case 32768:
                                                        var v69 float32 = (float32(3.14159) * 0.0)
                                                        _ = v69
                                                    default:
                                                        var v70 uint32 = 4294967295
                                                        _ = v70
                                                    }
                                                } else {
                                                    if true {
                                                        switch (fi65 % ((496) | 1)) {
                                                        case 9223372036854775807:
                                                            v71 := []int{fi61, fi61, (fi65 + (fi65 - (fi61 ^ 127)))}
                                                            _ = v71
                                                        case 32767:
                                                            if true {
                                                                v72 := []int{2147483647, fi61, (-(fi65))}
                                                                p1 -= (p1 + p1)
                                                                fmt.Println(p0, v22)
                                                                _ = v72
                                                            } else {
                                                                if false {
                                                                    v73 := []int{fi65, (fi61 >> ((255) & 63))}
                                                                    if len(v73) >= 2 {
                                                                        v73 = v73[1:len(v73)]
                                                                    }
                                                                    _ = v73
                                                                } else {
                                                                    var v74 uint16 = 2
                                                                    v75 := []int{470, (fi65 >> ((p0) & 63)), 32768}
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                var v76 int8 = 100
                                                                _ = v76
                                                            }
                                                        default:
                                                            v77 := GenS8{F0: 255, F1: 42}
                                                            _ = v77
                                                        }
                                                        v78 := GenS9{F0: 392}
                                                        if (!(false || false)) {
                                                            p0 = (p0 * p0)
                                                            v79 := GenS10{F0: "\n\t"}
                                                            if (!true) {
                                                                if (p0 > (p0 * (^(p0)))) {
                                                                    var v80 uint64 = 970
                                                                    p0 = 255
                                                                    var v81 float32 = 1.5
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    var v82 int16 = (v22 + v22)
                                                                    v83 := []int{fi61, (-(fi61)), fi61}
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            } else {
                                                                v84 := GenS11{F0: (v22 | v22), F1: "", F2: 18446744073709551615}
                                                                v85 := []int{(fi65 ^ (fi65 >> ((255) & 63)))}
                                                                var v86 int8 = 88
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            if false {
                                                                fmt.Println(v22, fi65)
                                                                v87 := GenS12{F0: 1}
                                                                v88 := GenS3{F0: 1}
                                                                _ = v87
                                                                _ = v88
                                                            } else {
                                                                var v89 string = ("\n\t" + "a")
                                                                v90 := []int{(^(fi65)), 2147483647, (fi65 + fi65)}
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            p1 -= (p1 - (p1 % ((p1) | 1)))
                                                            var v91 uint32 = 174
                                                            p0 = (p0 + (p0 + p0))
                                                            if (!((p0 << ((255) & 63)) >= 138)) {
                                                                var v92 uint64 = 638
                                                                v93 := GenS13{F0: (v92 / (((v92 - v92)) | 1)), F1: (p0 / ((p0) | 1))}
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                            fmt.Println(fi65, fi61, v22)
                                                            switch (fi61 & -447) {
                                                            case 127:
                                                                if ((p0 <= uint(v91)) && (true || true)) {
                                                                    fmt.Println(v91, fi65, p1)
                                                                    switch fi65 {
                                                                    case 2147483647:
                                                                        v91 -= (v91 >> ((707) & 63))
                                                                    default:
                                                                        v91 += 2147483647
                                                                    }
                                                                    p1 = (p1 ^ p1)
                                                                }
                                                                p0 -= 1
                                                                v94 := []int{(fi65 << ((p0) & 63)), 4294967295, fi61, int(v22)}
                                                                _ = v94
                                                            case 0:
                                                                v91 = 65535
                                                            default:
                                                                switch fi61 {
                                                                case 1:
                                                                    var v95 uint16 = 2
                                                                    _ = v95
                                                                default:
                                                                    var v96 int8 = -128
                                                                    var v97 float32 = float32(-2.25)
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                            }
                                                            switch (^(fi65)) {
                                                            case -9223372036854775808:
                                                                fmt.Println(fi65)
                                                                v98 := GenS10{F0: "\n\t"}
                                                                _ = v98
                                                            }
                                                            if false {
                                                                fmt.Println(fi65)
                                                                v99 := []int{(^(fi65))}
                                                                v99 = append(v99, fi61)
                                                                _ = v99
                                                            }
                                                            _ = v79
                                                            _ = v91
                                                        }
                                                        _ = v78
                                                    } else {
                                                        p1 = (p1 | p1)
                                                        v100 := GenS14{F0: 798, F1: -75, F2: (("" + ("foo bar" + "foo bar")) + ("foo bar" + "hello"))}
                                                        _ = v100
                                                    }
                                                }
                                            case "\n\t":
                                                v101 := []int{fi65, fi65, (^(fi65)), fi61}
                                                _ = v101
                                            case "":
                                                fmt.Println(v22)
                                            }
                                            fmt.Println(fi65)
                                            _ = fi65
                                        }
                                    }
                                    fmt.Println(p1, v22, v22)
                                    if false {
                                        var v102 float64 = 1.5
                                        switch ("hello" + ("hello" + ("" + "a"))) {
                                        case "":
                                            v103 := GenS15{F0: 127, F1: 398, F2: p0}
                                            _ = v103
                                        default:
                                            v104 := GenS16{F0: p0}
                                            _ = v104
                                        }
                                        v22 -= (v22 + 42)
                                        _ = v102
                                    } else {
                                        var v105 int8 = 127
                                        _ = v105
                                    }
                                }
                            } else {
                                var v106 uint16 = 228
                                fmt.Println(p0, v106, fi61)
                                v107 := GenS12{F0: 133}
                                _ = v106
                                _ = v107
                            }
                            _ = fi61
                        }
                    default:
                        p0 -= p0
                    }
                    for fi108 := 0; fi108 < 14; fi108++ {
                        for fi109 := 0; fi109 < 20; fi109++ {
                            v110 := GenS8{F0: 2, F1: 2147483647}
                            _ = fi109
                            _ = v110
                        }
                        _ = fi108
                    }
                }
                p1 += (p1 & p1)
            }
        default:
            fmt.Println(v22, v22)
        }
    }
    v22 -= 1
    _ = p0
    _ = p1
    _ = v22
    return (p0 % ((914) | 1))
}
func genFn1(p0 int16, p1 int, p2 string) int {
    var v111 int8 = 63
    var v112 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v111
    _ = v112
    return -218
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    var v3 float64 = 1.5
    var v4 uint64 = 473
    var v5 int16 = -1
    v4 += (v4 ^ v4)
    fmt.Println(v3, v4)
    if true {
        if (((-2.25 * (0.1 / 0.1)) + 1.5) <= 0.1) {
            if len(v2) > 0 {
                v2[0] = 1
            }
            v113 := GenS9{F0: 183}
            var v114 uint64 = v4
            _ = v113
            _ = v114
        } else {
            v4 = v4
            var v115 int8 = 127
            _ = v115
        }
        v3 -= 1.5
        v116 := GenS9{F0: 2147483648}
        _ = v116
    } else {
        v3 = 0.1
        if true {
            if false {
                if false {
                    for fi117 := 0; fi117 < 7; fi117++ {
                        v2 = append(v2, 4294967295)
                        if true {
                            _ = len(v2)
                            v5 = (^(v5))
                        } else {
                            fmt.Println(fi117, fi117)
                            var v118 uint64 = (v4 % (((^(v4))) | 1))
                            v4 = (v4 >> ((2) & 63))
                            _ = v118
                        }
                        _ = fi117
                    }
                } else {
                    var v119 int = -9223372036854775808
                    switch "foo bar" {
                    case "":
                        v119 = v119
                    case "a":
                        v120 := GenS17{F0: -203, F1: ("" + "a")}
                        _ = v120
                    case "\n\t":
                        var v121 int8 = -23
                        _ = v121
                    default:
                        switch "\n\t" {
                        case "":
                            v122 := GenS18{F0: v5}
                            _ = v122
                        case "foo bar":
                            var v123 bool = (true || false)
                            _ = v123
                        case "hello":
                            switch "\n\t" {
                            case "a":
                                var v124 uint = 65535
                                _ = v124
                            case "":
                                for fi125 := 0; fi125 < 7; fi125++ {
                                    var v126 bool = true
                                    var v127 int8 = 127
                                    v128 := GenS15{F0: 255, F1: 0, F2: 4294967295}
                                    _ = fi125
                                    _ = v126
                                    _ = v127
                                    _ = v128
                                }
                            }
                        default:
                            var v129 float64 = ((float64(-2.25) - float64((-2.25 / 1.5))) + (3.14159 - v3))
                            _ = v129
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v119
                }
            }
        } else {
            for fi130 := 0; fi130 < 15; fi130++ {
                var v131 int64 = -9223372036854775808
                _ = fi130
                _ = v131
            }
        }
        if true {
            if false {
                if false {
                    v5 = v5
                } else {
                    v132 := GenS11{F0: -1, F1: "a", F2: (v4 % (((v4 | v4)) | 1))}
                    fmt.Println(v5, v5, v3)
                    _ = v132
                }
                if true {
                    for fi133 := 0; fi133 < 3; fi133++ {
                        fmt.Println(v3, v4, v4)
                        var v134 int64 = 366
                        if true {
                            v4 = v4
                            for fi135 := 0; fi135 < 17; fi135++ {
                                v5 = 42
                                var v136 uint64 = v4
                                _ = fi135
                                _ = v136
                            }
                            v3 += 3.14159
                        } else {
                            v3 = ((v3 * v3) + v3)
                            v3 = v3
                        }
                        _ = fi133
                        _ = v134
                    }
                    v5 = 0
                    var v137 uint8 = 255
                    _ = v137
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                fmt.Println(v4, v3, v5)
            } else {
                v3 = v3
            }
        } else {
            switch -218 {
            case -620:
                v138 := GenS19{F0: 4294967295}
                _ = v138
            case 4294967295:
                delete(v1, "")
            }
        }
    }
    v139 := GenS16{F0: 255}
    var v140 float32 = 0.0
    var v141 int64 = -9223372036854775808
    fmt.Println(v140, v141)
    v4 = 258
    v4 -= v4
    if true {
        for fi142 := 0; fi142 < 18; fi142++ {
            fmt.Println(v141)
            _ = fi142
        }
        v3 = (0.1 / v3)
        for fi143 := 0; fi143 < 18; fi143++ {
            v144 := GenS20{F0: 100, F1: 42}
            for fi145 := 0; fi145 < 5; fi145++ {
                v146 := GenS21{F0: 4294967295, F1: "hello", F2: (255 >= 32767)}
                v3 = float64(v140)
                _ = fi145
                _ = v146
            }
            _ = fi143
            _ = v144
        }
    } else {
        for fi147 := 0; fi147 < 11; fi147++ {
            switch (("foo bar" + "") + "a") {
            case "a":
                if true {
                    delete(v1, ("hello" + "\n\t"))
                    for fi148 := 0; fi148 < 4; fi148++ {
                        var v149 uint64 = (v4 << ((262) & 63))
                        v140 += v140
                        v4 = (^(v149))
                        _ = fi148
                        _ = v149
                    }
                } else {
                    delete(v1, "hello")
                    if (v3 <= (0.1 - 1.5)) {
                        switch "" {
                        case "a":
                            fmt.Println(v140)
                        case "\n\t":
                            v150 := GenS12{F0: 127}
                            _ = v150
                        }
                        var v151 int64 = (v141 >> ((genFn0(42, (-(v141)))) & 63))
                        _ = v151
                    }
                }
            }
            for fi152 := 0; fi152 < 14; fi152++ {
                v141 = (^(v141))
                fmt.Println(v3)
                if ((1 <= 875) || (!true)) {
                    var v153 float32 = (v140 * v140)
                    if (147 <= 0) {
                        var v154 int64 = v141
                        if (128 == 127) {
                            delete(v1, "hello")
                        } else {
                            v153 -= (1e10 - float32(v3))
                            if false {
                                var v155 uint64 = v4
                                v156 := GenS22{F0: true}
                                v4 = (v155 / ((v4) | 1))
                                var v157 uint32 = 1
                                v158 := GenS17{F0: (v141 % ((v141) | 1)), F1: "hello"}
                                if ((float64(v140) / ((float64(((v153 / v140) * (v153 * 0.0))) / (0.1 * 0.1)) * v3)) != v3) {
                                    v159 := GenS3{F0: 248}
                                    if false {
                                        var v160 int32 = -2147483648
                                        fmt.Println(v155, v157, v155)
                                        v161 := v1["foo bar"]
                                        switch ("" + "") {
                                        case "a":
                                            v162 := v1["a"]
                                            _ = v162
                                        case "hello":
                                            var v163 int16 = (v5 & (-(v5)))
                                            _ = v163
                                        }
                                        if false {
                                            switch (^(fi147)) {
                                            case 32767:
                                                v3 += (float64((((v153 + v153) + (v140 * v140)) * (v153 - float32(1.5)))) + (v3 - v3))
                                            case -2147483648:
                                                var v164 float32 = v140
                                                switch "hello" {
                                                case "\n\t":
                                                    v161 = (v161 | (v161 ^ v161))
                                                case "a":
                                                    if true {
                                                        var v165 int8 = -128
                                                        v161 -= (v161 | (^(v161)))
                                                        switch "" {
                                                        case "\n\t":
                                                            var v166 int = (v161 ^ int(v157))
                                                            _ = v166
                                                        case "":
                                                            var v167 uint8 = 65
                                                            _ = v167
                                                        case "foo bar":
                                                            switch ("foo bar" + (("a" + "hello") + "foo bar")) {
                                                            case "":
                                                                v154 -= (v141 ^ (v141 % (((v141 | v141)) | 1)))
                                                            case "foo bar":
                                                                v164 += v140
                                                            case "\n\t":
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v168 uint = genFn0(4294967295, -1)
                                                                    _ = v168
                                                                case "foo bar":
                                                                    switch "a" {
                                                                    case "a":
                                                                        v140 = 3.14159
                                                                    case "foo bar":
                                                                        if (25 >= 0) {
                                                                            v1["a"] = (fi147 / (((fi152 & v161)) | 1))
                                                                            if ((v154 & v141) <= v141) {
                                                                                v169 := GenS23{F0: (true || false), F1: v160, F2: 42}
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                fmt.Println(v160)
                                                                                _ = v169
                                                                            } else {
                                                                                v155 -= v155
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        v160 = (v160 + v160)
                                                                    }
                                                                }
                                                            default:
                                                                var v170 int8 = (v165 << ((1) & 63))
                                                                _ = v170
                                                            }
                                                        default:
                                                            fmt.Println(v157, v3, v164)
                                                        }
                                                        _ = v165
                                                    }
                                                    v1[("a" + "a")] = 32767
                                                case "":
                                                    if (v157 >= (v157 & 65535)) {
                                                        switch ("" + "hello") {
                                                        case "foo bar":
                                                            v171 := GenS24{F0: 2, F1: int32(v154), F2: genFn0(962, v141)}
                                                            _ = v171
                                                        case "\n\t":
                                                            v172 := GenS25{F0: 18446744073709551615, F1: v5}
                                                            _ = v172
                                                        }
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v161 -= fi147
                                                        _ = len(v2)
                                                    }
                                                }
                                                v173 := GenS26{F0: -70}
                                                _ = v164
                                                _ = v173
                                            }
                                            switch (fi147 & fi147) {
                                            case 384:
                                                switch "hello" {
                                                case "a":
                                                    switch v161 {
                                                    case 42:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case -2147483648:
                                                        v4 = v4
                                                    default:
                                                        if ((v5 & (v5 & v5)) <= -1) {
                                                            fmt.Println(fi152)
                                                            v4 = v4
                                                        }
                                                    }
                                                case "hello":
                                                    if (false || true) {
                                                        var v174 int8 = 1
                                                        v3 = v3
                                                        _ = v174
                                                    } else {
                                                        v4 -= 2
                                                        v175 := GenS26{F0: -128}
                                                        if true {
                                                            v155 = (v4 << ((65535) & 63))
                                                            delete(v1, "")
                                                        } else {
                                                            fmt.Println(fi147, fi152)
                                                            var v176 int8 = -1
                                                            _ = v176
                                                        }
                                                        v177 := GenS3{F0: 44}
                                                        _ = v175
                                                        _ = v177
                                                    }
                                                }
                                            case 26:
                                                _ = len(v1)
                                            case 32767:
                                                switch "a" {
                                                case "\n\t":
                                                    switch v161 {
                                                    case -523:
                                                        if false {
                                                            if (false && false) {
                                                                if (((!true) || true) || false) {
                                                                    v178 := GenS1{F0: "\n\t", F1: 255, F2: (v157 & (v157 & (^(v157))))}
                                                                    if (809 >= 292) {
                                                                        var v179 int64 = -617
                                                                        v154 -= 1
                                                                        _ = v179
                                                                    }
                                                                    v180 := GenS27{F0: (v5 < -1), F1: (-(v160)), F2: false}
                                                                    fmt.Println(v157)
                                                                    fmt.Println(v153, v3, v155)
                                                                    _ = v178
                                                                    _ = v180
                                                                }
                                                            }
                                                            var v181 uint8 = 127
                                                            if ((v154 % (((v154 & (v154 * v154))) | 1)) <= v141) {
                                                                var v182 float32 = float32(1e10)
                                                                _ = v182
                                                            }
                                                            v140 -= float32(v3)
                                                            _ = v181
                                                        }
                                                    default:
                                                        var v183 float32 = v140
                                                        _ = v183
                                                    }
                                                }
                                                fmt.Println(v141, v140)
                                                fmt.Println(v153)
                                            default:
                                                var v184 float32 = v153
                                                _ = v184
                                            }
                                        } else {
                                            var v185 int16 = v5
                                            fmt.Println(v157)
                                            _ = len(v2)
                                            switch "\n\t" {
                                            case "a":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "foo bar":
                                                switch 42 {
                                                case 2147483647:
                                                    v4 -= (v155 << ((4294967295) & 63))
                                                    var v186 uint32 = (v157 >> ((0) & 63))
                                                    _ = v186
                                                }
                                            case "hello":
                                                switch genFn1(1000, (v161 | (fi147 % ((fi152) | 1))), "a") {
                                                case 1:
                                                    v161 += (fi152 % ((877) | 1))
                                                case -1:
                                                    v187 := GenS28{F0: int8(v157), F1: 100}
                                                    _ = v187
                                                }
                                            default:
                                                v2 = append(v2, v157)
                                            }
                                            if (v157 <= v157) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch v161 {
                                                case 0:
                                                    v3 -= 3.14159
                                                case 32767:
                                                    v188 := GenS29{F0: (fi152 % ((-433) | 1)), F1: genFn0(463, v154)}
                                                    _ = v188
                                                }
                                                v189 := GenS30{F0: 109, F1: 939}
                                                _ = v189
                                            }
                                            _ = v185
                                        }
                                        var v190 uint16 = uint16(v157)
                                        fmt.Println(v140)
                                        _ = v160
                                        _ = v161
                                        _ = v190
                                    } else {
                                        if true {
                                            fmt.Println(v3)
                                        } else {
                                            switch (fi147 | -2147483648) {
                                            case 1:
                                                switch "" {
                                                case "":
                                                    v191 := GenS22{F0: false}
                                                    _ = v191
                                                case "a":
                                                    var v192 float64 = ((1e10 / v3) - (float64((((v140 * v153) / v140) * v140)) * ((v3 * (float64(v140) / v3)) - float64(v153))))
                                                    _ = v192
                                                }
                                            }
                                            var v193 int = 128
                                            if (!false) {
                                                v194 := GenS31{F0: 0}
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(v141)
                                                case "a":
                                                    var v195 uint64 = 1
                                                    _ = v195
                                                }
                                                fmt.Println(v4, v141, v4)
                                                _ = v194
                                            }
                                            _ = v193
                                        }
                                        if false {
                                            if false {
                                                v196 := GenS23{F0: (true && false), F1: 42, F2: v141}
                                                if false {
                                                    v154 = 42
                                                    if false {
                                                        v155 = v4
                                                        switch fi152 {
                                                        case 42:
                                                            var v197 float64 = (v3 / 1.5)
                                                            _ = v197
                                                        case 2147483647:
                                                            var v198 int64 = v154
                                                            _ = v198
                                                        case -458:
                                                            var v199 uint = 2
                                                            _ = v199
                                                        }
                                                    }
                                                    v200 := GenS10{F0: "\n\t"}
                                                    _ = len(v1)
                                                    v201 := GenS32{F0: 1, F1: 2}
                                                    _ = v200
                                                    _ = v201
                                                } else {
                                                    if ((false || (float32(float64(v140)) <= (3.14159 - 0.0))) && false) {
                                                        v4 = (v155 - (v155 * 42))
                                                        var v202 int32 = 0
                                                        _ = v202
                                                    } else {
                                                        v155 -= 0
                                                        v3 += float64(v153)
                                                    }
                                                    if false {
                                                        v203 := GenS32{F0: 1, F1: 127}
                                                        v154 += 9223372036854775807
                                                        var v204 string = ("foo bar" + ("foo bar" + "\n\t"))
                                                        _ = v203
                                                        _ = v204
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = 65535
                                                        }
                                                    }
                                                }
                                                fmt.Println(v141, v157)
                                                v155 = 226
                                                v140 = float32((v3 - ((v3 / 0.1) / (float64((float32(v3) + 1.5)) * (1.5 * 0.1)))))
                                                if false {
                                                    fmt.Println(v153, v3, v155)
                                                    if true {
                                                        var v205 uint16 = 32767
                                                        _ = v205
                                                    } else {
                                                        var v206 int32 = 2147483647
                                                        _ = v206
                                                    }
                                                } else {
                                                    var v207 uint32 = (v157 ^ v157)
                                                    v154 = int64(fi147)
                                                    _ = v207
                                                }
                                                _ = v196
                                            }
                                            if ((v155 / ((1) | 1)) < 362) {
                                                v155 += v4
                                                var v208 float32 = v140
                                                _ = v208
                                            } else {
                                                v209 := GenS2{F0: (^(v154)), F1: 42}
                                                if (2 == 128) {
                                                    if false {
                                                        v210 := GenS33{F0: -2147483648, F1: 0, F2: fi147}
                                                        _ = v210
                                                    } else {
                                                        fmt.Println(v3, v157)
                                                    }
                                                    if false {
                                                        if (((-(fi147)) != fi152) && (!(!false))) {
                                                            v153 -= float32(((float64((0.1 * 0.0)) + v3) - v3))
                                                            switch fi152 {
                                                            case -549:
                                                                var v211 float32 = (v153 + v140)
                                                                _ = v211
                                                            default:
                                                                v212 := GenS34{F0: 666, F1: 1, F2: 2}
                                                                _ = v212
                                                            }
                                                        }
                                                        if true {
                                                            v213 := GenS35{F0: 42, F1: 2, F2: ("a" + "")}
                                                            _ = v213
                                                        } else {
                                                            _ = len(v1)
                                                            v214 := GenS36{F0: 1, F1: genFn1((v5 >> ((559) & 63)), genFn1(v5, fi152, "foo bar"), "\n\t")}
                                                            v153 += ((v153 * v140) / v153)
                                                            _ = v214
                                                        }
                                                        var v215 float32 = (((1e10 - 0.1) * ((1e10 * (v153 / (1e10 - v140))) - (3.14159 / (float32((1.5 + v3)) + v140)))) * 3.14159)
                                                        _ = v215
                                                    }
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        v216 := GenS27{F0: true, F1: 2147483647, F2: (v4 > 0)}
                                                        _ = v216
                                                    case "foo bar":
                                                        fmt.Println(v155)
                                                    default:
                                                        v217 := GenS15{F0: 22, F1: uint16(v157), F2: 687}
                                                        _ = v217
                                                    }
                                                    var v218 bool = (42 <= 65535)
                                                    if true {
                                                        if v218 {
                                                            var v219 uint16 = 65535
                                                            _ = v219
                                                        } else {
                                                            var v220 uint32 = (v157 % ((v157) | 1))
                                                            if (2 < 0) {
                                                                if v218 {
                                                                    fmt.Println(v155)
                                                                    var v221 int8 = 42
                                                                    _ = v221
                                                                } else {
                                                                    if ("" <= "foo bar") {
                                                                        var v222 int = (fi147 + genFn1((v5 - -1), fi152, "\n\t"))
                                                                        var v223 float32 = float32(1e10)
                                                                        _ = v222
                                                                        _ = v223
                                                                    }
                                                                    v2 = append(v2, v220)
                                                                }
                                                                var v224 uint64 = (v4 | v155)
                                                                _ = v224
                                                            }
                                                            _ = len(v2)
                                                            _ = v220
                                                        }
                                                        var v225 float32 = (1e10 * ((0.1 - v140) + ((-2.25 - (3.14159 - v153)) * (v153 * 0.1))))
                                                        v4 += (v4 + (v4 ^ (v155 * v4)))
                                                        _ = v225
                                                    }
                                                    _ = v218
                                                }
                                                var v226 bool = true
                                                _ = v209
                                                _ = v226
                                            }
                                            v3 += 0.1
                                        } else {
                                            v154 = v141
                                        }
                                        v227 := GenS36{F0: 303, F1: 42}
                                        _ = v227
                                    }
                                    v5 = -32768
                                    _ = v159
                                }
                                _ = v155
                                _ = v156
                                _ = v157
                                _ = v158
                            } else {
                                v4 = (v4 - (v4 << ((1) & 63)))
                                _ = len(v2)
                            }
                        }
                        _ = v154
                    } else {
                        if (-128 == 42) {
                            fmt.Println(v5, v3)
                            v140 = (v153 * (((v153 * v153) + v140) - 3.14159))
                            if true {
                                v228 := GenS13{F0: (v4 >> ((42) & 63)), F1: 773}
                                fmt.Println(v5)
                                v229 := GenS37{F0: (fi147 / ((fi147) | 1)), F1: 0}
                                v230 := GenS22{F0: true}
                                _ = v228
                                _ = v229
                                _ = v230
                            } else {
                                if true {
                                    v153 = 0.1
                                    var v231 uint8 = 128
                                    fmt.Println(v4, v4)
                                    _ = v231
                                } else {
                                    v232 := GenS25{F0: 571, F1: (v5 * (v5 - (v5 % ((v5) | 1))))}
                                    v2 = append(v2, 4294967295)
                                    switch ("hello" + ("\n\t" + "")) {
                                    case "foo bar":
                                        v233 := GenS38{F0: 1}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if (((("\n\t" + "") + (("" + (("" + ("a" + ("\n\t" + "foo bar"))) + "")) + ("\n\t" + ""))) + (("foo bar" + "foo bar") + "foo bar")) == "\n\t") {
                                            fmt.Println(v153)
                                            var v234 int16 = v5
                                            _ = v234
                                        } else {
                                            v5 += (v5 | v5)
                                            if (("hello" + (("a" + "") + ("" + "a"))) > "") {
                                                var v235 int64 = v141
                                                _ = v235
                                            }
                                        }
                                        if true {
                                            var v236 string = "a"
                                            _ = v236
                                        }
                                        v4 = 577
                                        _ = v233
                                    case "a":
                                        fmt.Println(v4, v5, v4)
                                    }
                                    v237 := GenS39{F0: 100, F1: 646, F2: (v5 | (-(v5)))}
                                    _ = v232
                                    _ = v237
                                }
                            }
                        }
                        switch 128 {
                        case 1:
                            if ((4294967295 == 506) || true) {
                                v238 := GenS40{F0: 0}
                                _ = v238
                            }
                        default:
                            v239 := GenS41{F0: 4294967295}
                            _ = v239
                        }
                        if (((false || false) || ("foo bar" > ("\n\t" + ("\n\t" + "\n\t")))) && (((32767 != 255) && (0 != 2)) || false)) {
                            _ = len(v1)
                            var v240 float64 = 3.14159
                            _ = v240
                        }
                    }
                    _ = v153
                } else {
                    var v241 float64 = 0.0
                    v242 := v1[(("\n\t" + "") + "\n\t")]
                    if true {
                        var v243 int16 = v5
                        _ = v243
                    } else {
                        var v244 uint16 = 2
                        var v245 int32 = 541
                        _ = v244
                        _ = v245
                    }
                    _ = v241
                    _ = v242
                }
                _ = fi152
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi147
        }
        for fi246 := 0; fi246 < 3; fi246++ {
            for fi247 := 0; fi247 < 8; fi247++ {
                var v248 int = genFn1(v5, (fi246 << ((636) & 63)), "a")
                v3 = float64(3.14159)
                _ = fi247
                _ = v248
            }
            v5 -= 0
            _ = fi246
        }
    }
    var v249 float32 = (v140 - v140)
    var v250 uint = 0
    var v251 uint8 = 255
    if (9223372036854775807 > -2147483648) {
        v140 += (v140 * (((3.14159 * (v249 - v140)) - 0.0) * -2.25))
    }
    v249 -= (v140 * (3.14159 - v140))
    for fi252 := 0; fi252 < 17; fi252++ {
        switch fi252 {
        case 2147483647:
            var v253 int32 = 0
            _ = v253
        case -111:
            v250 = (v250 & 1)
        default:
            v3 = 1e10
        }
        _ = fi252
    }
    v1["\n\t"] = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v249:", v249)
    fmt.Println("v250:", v250)
    fmt.Println("v251:", v251)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v139
    _ = v140
    _ = v141
    _ = v2
    _ = v249
    _ = v250
    _ = v251
    _ = v3
    _ = v4
    _ = v5
}
