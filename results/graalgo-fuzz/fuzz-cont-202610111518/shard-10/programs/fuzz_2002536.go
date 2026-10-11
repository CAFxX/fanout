package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int8; F1 uint64 }
type GenS14 struct { F0 int64; F1 int }
type GenS15 struct { F0 int32; F1 uint32 }
type GenS16 struct { F0 int64; F1 uint }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int16; F1 uint }
type GenS19 struct { F0 int8; F1 uint8; F2 string }

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]int)
    var v3 float64 = 0.0
    v1 = 1
    var v4 int = 9223372036854775807
    if (!true) {
        fmt.Println(v3, v3)
        fmt.Println(v4, v3)
    } else {
        var v5 int64 = -701
        v6 := []int{v4, 237, v4}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 987, F1: -405}
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 bool = true
        if v9 {
            for fi10 := 0; fi10 < 1; fi10++ {
                switch ("\n\t" + "hello") {
                case "hello":
                    switch v4 {
                    case -1:
                        if false {
                            if true {
                                if v9 {
                                    v11 := []int{}
                                    _ = v11
                                } else {
                                    if true {
                                        if (fi10 >= 42) {
                                            v12 := GenS1{F0: 0, F1: 1}
                                            var v13 int16 = 32767
                                            _ = v12
                                            _ = v13
                                        }
                                        var v14 string = "foo bar"
                                        _ = v14
                                    } else {
                                        if v9 {
                                            v15 := []int{(v4 + 42), (fi8 & 802), (fi10 - v4)}
                                            _ = v15
                                        }
                                        switch "a" {
                                        case "a":
                                            v16 := GenS2{F0: 158, F1: v4}
                                            _ = v16
                                        case "\n\t":
                                            var v17 uint64 = 18446744073709551615
                                            _ = v17
                                        default:
                                            var v18 string = (("foo bar" + "hello") + "\n\t")
                                            _ = v18
                                        }
                                    }
                                    fmt.Println(fi8, v4, v9)
                                }
                                v19 := []int{(fi10 ^ (fi10 * fi10))}
                                _ = v19
                            } else {
                                if (-889 == 1) {
                                    v20 := GenS1{F0: 0, F1: 0}
                                    if (64 < 42) {
                                        if (1 == 0) {
                                            v21 := GenS2{F0: 255, F1: (^(fi8))}
                                            var v22 int16 = 1000
                                            if (("" < "hello") || v9) {
                                                var v23 string = ("" + ("hello" + "\n\t"))
                                                _ = v23
                                            } else {
                                                var v24 float32 = (1e10 / -2.25)
                                                var v25 int = (^(fi8))
                                                var v26 uint64 = 0
                                                v4 += fi8
                                                v27 := GenS2{F0: 2, F1: (fi8 & v4)}
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                            v28 := GenS3{F0: (!((v22 * (^(v22))) == -48))}
                                            v1 = (v1 | (v1 / (((v1 & (v1 + v1))) | 1)))
                                            _ = v21
                                            _ = v22
                                            _ = v28
                                        } else {
                                            if (((!false) || (0 < 2147483647)) || (18446744073709551615 == 0)) {
                                                var v29 bool = (!false)
                                                var v30 uint16 = (v1 | (v1 + 154))
                                                fmt.Println(v4, v4, v4)
                                                _ = v29
                                                _ = v30
                                            } else {
                                                v2[(("a" + "foo bar") + "a")] = fi8
                                            }
                                            if (!false) {
                                                v4 -= 42
                                                var v31 int = -40
                                                switch (v4 ^ v31) {
                                                case -34:
                                                    fmt.Println(v1)
                                                case 128:
                                                    fmt.Println(fi8, fi10)
                                                }
                                                _ = len(v2)
                                                _ = v31
                                            }
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            v32 := GenS2{F0: 217, F1: -2147483648}
                                            _ = v32
                                        case "\n\t":
                                            if (v9 && v9) {
                                                v9 = (!(("hello" + "a") < ("foo bar" + "")))
                                                v3 = 1.5
                                            } else {
                                                v33 := GenS2{F0: 2, F1: v4}
                                                if true {
                                                    switch v4 {
                                                    case 211:
                                                        var v34 uint16 = (v1 * (v1 * (v1 * v1)))
                                                        _ = v34
                                                    case 32768:
                                                        if (v9 && (true || (594 != 18446744073709551615))) {
                                                            if (!v9) {
                                                                var v35 float64 = v3
                                                                var v36 int = (v4 + (v4 - 9223372036854775807))
                                                                v37 := []int{v36, (fi10 - (fi10 - v36)), (fi10 ^ fi10)}
                                                                _ = v35
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                v1 = 65535
                                                                if v9 {
                                                                    if (!(408 != (v1 >> ((439) & 63)))) {
                                                                        if v9 {
                                                                            v38 := GenS3{F0: v9}
                                                                            _ = v38
                                                                        } else {
                                                                            _ = len(v2)
                                                                        }
                                                                    } else {
                                                                        var v39 float64 = ((v3 * ((1e10 + v3) - (v3 + -2.25))) + (((3.14159 + 0.1) * v3) + v3))
                                                                        switch "hello" {
                                                                        case "a":
                                                                            v40 := GenS1{F0: 65535, F1: -1}
                                                                            _ = v40
                                                                        default:
                                                                            v1 += v1
                                                                        }
                                                                        _ = v39
                                                                    }
                                                                    v4 = fi8
                                                                } else {
                                                                    var v41 uint8 = 7
                                                                    v3 = 0.0
                                                                    _ = v41
                                                                }
                                                            }
                                                            fmt.Println(fi10, v9)
                                                        } else {
                                                            switch ("" + "hello") {
                                                            case "\n\t":
                                                                var v42 int = 32768
                                                                _ = v42
                                                            case "foo bar":
                                                                v43 := GenS4{F0: 0, F1: 127}
                                                                _ = v43
                                                            default:
                                                                var v44 int = (fi8 & (fi8 >> ((0) & 63)))
                                                                v45 := GenS2{F0: 49, F1: (v4 | (fi8 ^ v4))}
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                            v1 += v1
                                                            v46 := GenS5{F0: 1, F1: 205}
                                                            _ = v46
                                                        }
                                                        v47 := []int{}
                                                        v9 = true
                                                        _ = v47
                                                    case 657:
                                                        v9 = (v9 || (((v9 || (1000 > -27)) || v9) && (9223372036854775807 != 42)))
                                                    }
                                                    v48 := []int{(fi8 * fi10), 1}
                                                    _ = v48
                                                }
                                                _ = v33
                                            }
                                        case "a":
                                            fmt.Println(v9, v4, v1)
                                            fmt.Println(v9, v4)
                                        }
                                    } else {
                                        v49 := GenS6{F0: 600, F1: v9}
                                        var v50 uint32 = 2
                                        _ = v49
                                        _ = v50
                                    }
                                    var v51 int64 = -9223372036854775808
                                    _ = v20
                                    _ = v51
                                } else {
                                    var v52 uint16 = 1
                                    _ = v52
                                }
                            }
                        }
                    default:
                        var v53 bool = (v9 && true)
                        _ = v53
                    }
                case "foo bar":
                    var v54 uint64 = 1
                    _ = v54
                default:
                    switch (("" + "\n\t") + "\n\t") {
                    case "\n\t":
                        _ = len(v2)
                    case "":
                        v3 += 0.0
                    case "hello":
                        var v55 float64 = ((((v3 + -2.25) / float64(float32((((float64(1.5) - v3) * (3.14159 - (v3 * ((3.14159 - ((((3.14159 - float64((3.14159 * (0.0 + (0.1 / 1.0))))) / v3) / v3) + v3)) + float64(-2.25))))) + (float64(1e10) - ((v3 / v3) + ((((v3 - 0.1) - (v3 - 1.5)) - 0.0) / float64((0.1 / -2.25))))))))) + float64(float32(v3))) - v3)
                        _ = v55
                    }
                }
                var v56 uint8 = 45
                _ = fi10
                _ = v56
            }
            for fi57 := 0; fi57 < 15; fi57++ {
                if false {
                    fmt.Println(fi8, v1, fi8)
                    v58 := GenS3{F0: v9}
                    if v9 {
                        if v9 {
                            fmt.Println(v9, fi8)
                        }
                    } else {
                        v1 -= v1
                        var v59 bool = v9
                        _ = v59
                    }
                    _ = v58
                } else {
                    if v9 {
                        v60 := GenS7{F0: -1}
                        _ = v60
                    } else {
                        switch ("" + ("" + ("foo bar" + ("foo bar" + "")))) {
                        case "foo bar":
                            v61 := GenS8{F0: -2147483648}
                            _ = v61
                        default:
                            v62 := GenS1{F0: 42, F1: 348}
                            var v63 float32 = 1e10
                            _ = v62
                            _ = v63
                        }
                    }
                }
                switch ("\n\t" + ("hello" + "foo bar")) {
                case "\n\t":
                    v64 := GenS8{F0: 1}
                    _ = v64
                case "hello":
                    switch 4294967295 {
                    case 32767:
                        v4 -= -2147483648
                    case -9223372036854775808:
                        fmt.Println(v9, fi57, v3)
                    case 2147483647:
                        if (true || v9) {
                            fmt.Println(v3)
                            if v9 {
                                if v9 {
                                    switch fi57 {
                                    case 2147483647:
                                        switch "" {
                                        case "a":
                                            v65 := GenS9{F0: (v1 - v1)}
                                            v3 += 1.5
                                            v4 += fi57
                                            switch ("foo bar" + "foo bar") {
                                            case "foo bar":
                                                switch "a" {
                                                case "foo bar":
                                                    v1 = (v1 * (v1 >> ((1) & 63)))
                                                    v66 := v2["hello"]
                                                    if (int32(v1) <= 42) {
                                                        v67 := []int{(fi8 - (-(fi57))), (v4 * 2147483647), (-(fi57)), fi57}
                                                        _ = v67
                                                    }
                                                    _ = v66
                                                case "hello":
                                                    fmt.Println(fi57, v3)
                                                    v68 := GenS2{F0: 255, F1: fi57}
                                                    _ = v68
                                                case "":
                                                    fmt.Println(v3, fi8, v9)
                                                default:
                                                    var v69 int64 = -9223372036854775808
                                                    v4 = (v4 + -9223372036854775808)
                                                    _ = v69
                                                }
                                            case "hello":
                                                if true {
                                                    v70 := v2["\n\t"]
                                                    if false {
                                                        switch (v70 & v70) {
                                                        case 479:
                                                            var v71 uint32 = 1
                                                            _ = v71
                                                        case 131:
                                                            v72 := []int{v70}
                                                            _ = v72
                                                        default:
                                                            v3 += v3
                                                        }
                                                        var v73 int8 = 0
                                                        if (2147483647 < 0) {
                                                            var v74 int = (-(fi8))
                                                            var v75 int16 = int16(v73)
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        _ = v73
                                                    } else {
                                                        if v9 {
                                                            switch v4 {
                                                            case -945:
                                                                v3 = ((1.5 * (-2.25 - v3)) + float64(float32(-2.25)))
                                                            }
                                                            var v76 int64 = 9223372036854775807
                                                            v9 = false
                                                            _ = v76
                                                        }
                                                        switch fi8 {
                                                        case 127:
                                                            if v9 {
                                                                var v77 int64 = 9223372036854775807
                                                                var v78 uint16 = (v1 >> ((4294967295) & 63))
                                                                v70 -= (fi57 << ((0) & 63))
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                _ = len(v2)
                                                                var v79 uint8 = 2
                                                                _ = v79
                                                            }
                                                        default:
                                                            v3 += 1.5
                                                        }
                                                        var v80 int64 = -9223372036854775808
                                                        _ = v80
                                                    }
                                                    _ = v70
                                                } else {
                                                    var v81 int = (fi8 ^ (v4 % ((v4) | 1)))
                                                    v82 := []int{}
                                                    _ = v81
                                                    _ = v82
                                                }
                                            case "\n\t":
                                                v83 := GenS10{F0: "", F1: 206}
                                                _ = v83
                                            }
                                            _ = v65
                                        case "":
                                            switch (fi57 * fi8) {
                                            case 128:
                                                v1 += v1
                                            default:
                                                if ((float32(float64(0.0)) + (-2.25 + 1e10)) > (-2.25 - (1e10 / (0.0 + 0.1)))) {
                                                    if v9 {
                                                        switch fi8 {
                                                        case 32767:
                                                            fmt.Println(v4, v4, fi8)
                                                        case 42:
                                                            switch (fi8 * (fi8 * (v4 >> ((80) & 63)))) {
                                                            case -9223372036854775808:
                                                                v84 := []int{(v4 >> ((4294967295) & 63)), 128, 781, (fi57 + -233)}
                                                                _ = v84
                                                            case -205:
                                                                delete(v2, "hello")
                                                            }
                                                        case -22:
                                                            if v9 {
                                                                var v85 uint = 18446744073709551615
                                                                switch v4 {
                                                                case 428:
                                                                    var v86 uint16 = v1
                                                                    _ = v86
                                                                case -138:
                                                                    if v9 {
                                                                        v87 := GenS11{F0: 255, F1: v1, F2: 62}
                                                                        _ = v87
                                                                    } else {
                                                                        v88 := GenS5{F0: 42, F1: 893}
                                                                        fmt.Println(v85, v85, v1)
                                                                        fmt.Println(v85)
                                                                        _ = v88
                                                                    }
                                                                    var v89 int8 = 100
                                                                    _ = v89
                                                                default:
                                                                    v3 = (v3 * -2.25)
                                                                    fmt.Println(v9, v3)
                                                                }
                                                                var v90 int64 = -251
                                                                _ = v85
                                                                _ = v90
                                                            }
                                                        default:
                                                            v91 := []int{}
                                                            _ = v91
                                                        }
                                                        var v92 float32 = -2.25
                                                        _ = v92
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v1, v1, v1)
                                                        }
                                                        switch (v4 ^ fi57) {
                                                        case 694:
                                                            v1 -= (v1 - 62)
                                                        }
                                                        v93 := GenS4{F0: -1, F1: 127}
                                                        _ = v93
                                                    }
                                                    v94 := GenS12{F0: uint(v1)}
                                                    _ = v94
                                                } else {
                                                    var v95 uint64 = 1
                                                    var v96 int8 = 0
                                                    _ = v95
                                                    _ = v96
                                                }
                                            }
                                        case "\n\t":
                                            var v97 bool = false
                                            _ = v97
                                        default:
                                            if v9 {
                                                v98 := []int{(v4 - -1)}
                                                _ = v98
                                            }
                                        }
                                    default:
                                        if v9 {
                                            var v99 float64 = (1.5 / v3)
                                            _ = v99
                                        } else {
                                            v100 := GenS13{F0: 42, F1: 18446744073709551615}
                                            var v101 uint64 = 4294967295
                                            v102 := []int{}
                                            _ = v100
                                            _ = v101
                                            _ = v102
                                        }
                                    }
                                    v103 := GenS14{F0: 0, F1: fi8}
                                    _ = v103
                                } else {
                                    if (!v9) {
                                        v9 = v9
                                    }
                                    var v104 uint = 1
                                    v105 := GenS15{F0: 0, F1: 0}
                                    _ = v104
                                    _ = v105
                                }
                                v2["\n\t"] = (fi57 >> ((0) & 63))
                                if v9 {
                                    switch v4 {
                                    case -1:
                                        if false {
                                            var v106 uint64 = 42
                                            _ = v106
                                        } else {
                                            if (!v9) {
                                                var v107 int = (^(fi57))
                                                _ = v107
                                            } else {
                                                var v108 int = fi57
                                                if false {
                                                    var v109 uint16 = 0
                                                    _ = v109
                                                } else {
                                                    var v110 string = "a"
                                                    v2["a"] = (fi8 & (fi57 << ((2) & 63)))
                                                    v111 := []int{(fi57 / ((v108) | 1)), (fi8 << ((626) & 63))}
                                                    var v112 uint16 = 1
                                                    if (!v9) {
                                                        fmt.Println(v9, fi57, v108)
                                                    } else {
                                                        if len(v111) > 0 {
                                                            v111[0] = v4
                                                        }
                                                    }
                                                    _ = v110
                                                    _ = v111
                                                    _ = v112
                                                }
                                                _ = len(v2)
                                                v113 := GenS16{F0: 1, F1: 4294967295}
                                                _ = v108
                                                _ = v113
                                            }
                                            if v9 {
                                                v114 := GenS9{F0: 445}
                                                _ = v114
                                            } else {
                                                switch (fi8 - (fi57 + (-(fi57)))) {
                                                case 9223372036854775807:
                                                    v115 := GenS17{F0: v1}
                                                    _ = v115
                                                case 913:
                                                    var v116 uint16 = (v1 % ((783) | 1))
                                                    _ = v116
                                                default:
                                                    v117 := GenS10{F0: ((("a" + "foo bar") + "hello") + "a"), F1: v1}
                                                    _ = v117
                                                }
                                            }
                                        }
                                    case 32767:
                                        fmt.Println(v3)
                                    default:
                                        v118 := GenS5{F0: -1, F1: 0}
                                        _ = v118
                                    }
                                    v9 = (!true)
                                    v3 += (1.5 - (3.14159 / ((v3 * 3.14159) / 0.1)))
                                }
                            } else {
                                if true {
                                    fmt.Println(fi8, v3, v3)
                                } else {
                                    var v119 int64 = -1
                                    var v120 uint32 = 2147483648
                                    _ = v119
                                    _ = v120
                                }
                                fmt.Println(v9, v3, v4)
                                fmt.Println(v3)
                            }
                        }
                    default:
                        var v121 uint = 65535
                        _ = v121
                    }
                }
                fmt.Println(v9)
                _ = fi57
            }
            if v9 {
                var v122 float32 = -2.25
                for fi123 := 0; fi123 < 6; fi123++ {
                    switch fi8 {
                    case 127:
                        v124 := GenS18{F0: 42, F1: 561}
                        _ = v124
                    case 25:
                        var v125 int32 = -2147483648
                        _ = v125
                    case 2147483647:
                        fmt.Println(v122, fi8, fi8)
                    }
                    v3 = (((v3 * v3) / v3) / (v3 - (0.0 * v3)))
                    _ = fi123
                }
                v126 := []int{}
                _ = v122
                _ = v126
            } else {
                var v127 float64 = 1e10
                _ = v127
            }
        }
        v3 = v3
        _ = fi8
        _ = v9
    }
    for fi128 := 0; fi128 < 14; fi128++ {
        if (!false) {
            var v129 bool = (-1 <= -1)
            v130 := []int{(fi128 & (fi128 - v4)), (-(fi128))}
            _ = v129
            _ = v130
        }
        var v131 uint = 65535
        _ = fi128
        _ = v131
    }
    v1 += v1
    for fi132 := 0; fi132 < 3; fi132++ {
        var v133 float32 = 3.14159
        for fi134 := 0; fi134 < 1; fi134++ {
            var v135 uint64 = 1
            if true {
                fmt.Println(fi132, v135)
                var v136 float64 = 0.0
                _ = v136
            } else {
                fmt.Println(fi132)
                v137 := GenS6{F0: (v135 - (v135 & 2)), F1: (!(896 >= fi132))}
                var v138 int = 32768
                _ = v137
                _ = v138
            }
            _ = fi134
            _ = v135
        }
        v139 := []int{v4}
        _ = fi132
        _ = v133
        _ = v139
    }
    switch "a" {
    case "foo bar":
        if false {
            v140 := GenS19{F0: -1, F1: 230, F2: "hello"}
            v141 := GenS3{F0: (v1 >= v1)}
            for fi142 := 0; fi142 < 6; fi142++ {
                v3 -= 0.1
                _ = fi142
            }
            _ = v140
            _ = v141
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
