package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 bool; F1 int32; F2 uint }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 int8; F1 uint8; F2 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int16 }
func genFn0() int64 {
    var v4 uint16 = 2
    switch "" {
    case "hello":
        switch 723 {
        case 32768:
            v5 := []int{}
            _ = v5
        }
    default:
        for fi6 := 0; fi6 < 17; fi6++ {
            switch (fi6 & 32767) {
            case -9223372036854775808:
                v7 := GenS1{F0: -218, F1: 342, F2: (fi6 | 1)}
                _ = v7
            case -147:
                var v8 string = ("" + "\n\t")
                _ = v8
            default:
                var v9 float64 = 0.1
                _ = v9
            }
            _ = fi6
        }
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        switch (fi10 - (fi10 << ((255) & 63))) {
        case -1:
            v4 = 42
        case -986:
            switch "\n\t" {
            case "hello":
                if false {
                    v11 := []int{}
                    var v12 uint = 1
                    fmt.Println(v12, fi10, v12)
                    _ = v11
                    _ = v12
                }
            case "foo bar":
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 int = (fi13 * (fi10 % (((fi13 ^ fi13)) | 1)))
                    v15 := GenS1{F0: 1, F1: -54, F2: (fi10 + 32768)}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "\n\t":
                for fi16 := 0; fi16 < 20; fi16++ {
                    v4 -= (v4 | v4)
                    if false {
                        if true {
                            var v17 int32 = -1
                            v18 := []int{fi16, (fi16 % ((-643) | 1))}
                            v17 = 1
                            _ = v17
                            _ = v18
                        } else {
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    fmt.Println(fi10)
                                case "a":
                                    v19 := []int{(fi16 * (fi10 ^ fi16)), (fi16 - 0), fi16}
                                    _ = v19
                                }
                                v4 = (v4 >> ((255) & 63))
                            }
                            var v20 string = ("foo bar" + "hello")
                            _ = v20
                        }
                    } else {
                        var v21 float32 = 1e10
                        v22 := GenS1{F0: 42, F1: 2147483647, F2: 706}
                        switch "" {
                        case "\n\t":
                            fmt.Println(fi16, fi16, v4)
                        }
                        _ = v21
                        _ = v22
                    }
                    var v23 bool = false
                    _ = fi16
                    _ = v23
                }
            }
        case 32768:
            for fi24 := 0; fi24 < 8; fi24++ {
                if true {
                    switch 9223372036854775807 {
                    case -724:
                        v4 += v4
                    }
                    v4 = v4
                    if (42 != 42) {
                        if ((("\n\t" + "foo bar") <= ("\n\t" + "foo bar")) && (true && (false && (true && true)))) {
                            switch fi10 {
                            case 32767:
                                if true {
                                    fmt.Println(fi10, fi24)
                                    var v25 uint16 = (v4 / ((1) | 1))
                                    _ = v25
                                }
                            case -1:
                                v26 := []int{fi24, 1}
                                _ = v26
                            case 2147483647:
                                if false {
                                    fmt.Println(fi24)
                                    v27 := GenS1{F0: 42, F1: -1, F2: -1}
                                    if (953 != -1) {
                                        v28 := []int{0, fi10}
                                        switch -1 {
                                        case 2147483647:
                                            if ((1e10 + 1e10) < 1.5) {
                                                v4 += v4
                                                switch (fi10 ^ (fi10 * fi24)) {
                                                case 204:
                                                    switch "a" {
                                                    case "":
                                                        if true {
                                                            var v29 string = ("" + "")
                                                            var v30 int8 = 0
                                                            _ = v29
                                                            _ = v30
                                                        } else {
                                                            v4 += v4
                                                            v31 := GenS1{F0: -32768, F1: 0, F2: 567}
                                                            v4 -= v4
                                                            _ = v31
                                                        }
                                                    case "a":
                                                        v28 = append(v28, (fi10 + 32768))
                                                    case "foo bar":
                                                        if (!(true || false)) {
                                                            v28 = append(v28, (fi10 ^ fi10))
                                                        }
                                                    default:
                                                        var v32 int8 = 49
                                                        _ = v32
                                                    }
                                                case -377:
                                                    v4 = 65535
                                                case 127:
                                                    switch "hello" {
                                                    case "":
                                                        if true {
                                                            if len(v28) >= 2 {
                                                                v28 = v28[1:len(v28)]
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v4 += (v4 & (v4 ^ v4))
                                                    default:
                                                        var v33 int8 = -128
                                                        _ = v33
                                                    }
                                                default:
                                                    if false {
                                                        var v34 uint16 = (v4 >> ((0) & 63))
                                                        var v35 uint16 = 255
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        if len(v28) > 0 {
                                                            v28[0] = fi10
                                                        }
                                                        switch ("a" + ("\n\t" + "\n\t")) {
                                                        case "a":
                                                            if (!(((0 != 170) && (!(42 == 32767))) && (42 != 0))) {
                                                                var v36 int16 = 1
                                                                if (772 > 18446744073709551615) {
                                                                    var v37 string = "\n\t"
                                                                    v36 += (v36 / ((-1) | 1))
                                                                    _ = v37
                                                                }
                                                                _ = v36
                                                            } else {
                                                                v38 := GenS1{F0: -32768, F1: -2147483648, F2: 32768}
                                                                if true {
                                                                    var v39 int32 = -1
                                                                    _ = v39
                                                                } else {
                                                                    v4 -= (v4 - v4)
                                                                }
                                                                v40 := GenS1{F0: 42, F1: 31, F2: fi10}
                                                                _ = v38
                                                                _ = v40
                                                            }
                                                        }
                                                        if (!(!(1 < -1))) {
                                                            if true {
                                                                v4 -= 1
                                                            } else {
                                                                var v41 uint64 = uint64(v4)
                                                                _ = v41
                                                            }
                                                            v42 := GenS1{F0: 1000, F1: 0, F2: 32768}
                                                            if (4294967295 >= 1) {
                                                                v4 = v4
                                                                var v43 int8 = 0
                                                                _ = v43
                                                            }
                                                            _ = v42
                                                        } else {
                                                            var v44 int32 = 0
                                                            var v45 int = 42
                                                            v45 = v45
                                                            if false {
                                                                fmt.Println(v44)
                                                                fmt.Println(fi24, fi24, fi10)
                                                            }
                                                            var v46 int32 = v44
                                                            var v47 int16 = 970
                                                            _ = v44
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    }
                                                }
                                                v48 := GenS1{F0: -696, F1: 0, F2: (^(fi24))}
                                                _ = v48
                                            } else {
                                                if (0.0 != 1e10) {
                                                    v4 = (v4 & (v4 | (v4 + 32767)))
                                                    fmt.Println(fi24)
                                                    if false {
                                                        if (0 == 127) {
                                                            fmt.Println(fi24)
                                                            switch (fi10 ^ fi10) {
                                                            case 429:
                                                                var v49 bool = (1 >= 42)
                                                                _ = v49
                                                            case 32768:
                                                                v4 = (v4 / (((v4 << ((1) & 63))) | 1))
                                                            case -333:
                                                                if (!false) {
                                                                    if true {
                                                                        var v50 uint16 = (v4 - (^(v4)))
                                                                        var v51 int8 = -1
                                                                        _ = v50
                                                                        _ = v51
                                                                    }
                                                                    v52 := GenS1{F0: -32768, F1: -525, F2: fi10}
                                                                    if (!false) {
                                                                        var v53 int64 = -9223372036854775808
                                                                        v54 := GenS2{F0: v4}
                                                                        _ = v53
                                                                        _ = v54
                                                                    } else {
                                                                        var v55 int = (fi24 ^ 4294967295)
                                                                        var v56 uint = 42
                                                                        v55 = (-(fi10))
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    _ = v52
                                                                } else {
                                                                    if true {
                                                                        v57 := GenS1{F0: 155, F1: -2147483648, F2: (-(fi24))}
                                                                        if len(v28) > 0 {
                                                                            v28[0] = 0
                                                                        }
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 int8 = 0
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            switch fi10 {
                                                                            case 0:
                                                                                v4 = 236
                                                                            case 32768:
                                                                                switch fi24 {
                                                                                case 127:
                                                                                    v59 := GenS2{F0: v4}
                                                                                    _ = v59
                                                                                case 42:
                                                                                    if false {
                                                                                        v58 += v58
                                                                                        v4 += 645
                                                                                        var v60 uint = 255
                                                                                        _ = v60
                                                                                    }
                                                                                default:
                                                                                    v61 := GenS3{F0: 0}
                                                                                    _ = v61
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v58)
                                                                        case "":
                                                                            var v62 bool = ((42 == 42) || (!false))
                                                                            _ = v62
                                                                        default:
                                                                            var v63 int32 = 1
                                                                            _ = v63
                                                                        }
                                                                        v4 += v4
                                                                        _ = v58
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            var v64 int8 = -18
                                                            _ = len(v28)
                                                            v64 = -89
                                                            _ = v64
                                                        } else {
                                                            fmt.Println(fi24, v4)
                                                            var v65 int32 = -786
                                                            v4 = (^(v4))
                                                            _ = v65
                                                        }
                                                    }
                                                }
                                                var v66 float32 = float32((3.14159 + ((3.14159 * 1.5) / float64(1e10))))
                                                var v67 uint = 4294967295
                                                _ = v66
                                                _ = v67
                                            }
                                        case -35:
                                            fmt.Println(fi24, fi24, fi10)
                                        }
                                        fmt.Println(v4, fi24, fi10)
                                        _ = v28
                                    }
                                    _ = v27
                                } else {
                                    switch (fi10 + -2147483648) {
                                    case -403:
                                        var v68 int32 = -1
                                        _ = v68
                                    case 1:
                                        v4 = (v4 - 2)
                                    case -288:
                                        v69 := []int{0}
                                        _ = v69
                                    default:
                                        var v70 int16 = -32768
                                        _ = v70
                                    }
                                }
                            default:
                                v71 := GenS4{F0: 4294967295}
                                _ = v71
                            }
                            v72 := GenS3{F0: -1}
                            _ = v72
                        } else {
                            if false {
                                v73 := GenS5{F0: -255}
                                var v74 float32 = (1e10 * 0.1)
                                v74 = 1e10
                                var v75 uint64 = 64
                                var v76 uint32 = 1
                                v77 := []int{(fi24 >> ((2) & 63)), int(v76), -731}
                                v75 = 42
                                _ = v73
                                _ = v74
                                _ = v75
                                _ = v76
                                _ = v77
                            } else {
                                if false {
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v78 bool = (-71 > 127)
                                            _ = v78
                                        case "a":
                                            var v79 int16 = 1
                                            _ = v79
                                        case "foo bar":
                                            v80 := GenS2{F0: (v4 - (v4 + v4))}
                                            _ = v80
                                        default:
                                            v4 = v4
                                            var v81 uint16 = (v4 + 255)
                                            v82 := []int{fi10}
                                            _ = v81
                                            _ = v82
                                        }
                                        v4 = v4
                                        v83 := []int{}
                                        v4 += v4
                                        _ = v83
                                    } else {
                                        v4 -= v4
                                    }
                                }
                                if (("foo bar" + "\n\t") >= "foo bar") {
                                    if false {
                                        v84 := GenS6{F0: 2147483648, F1: fi24}
                                        fmt.Println(fi24, fi10, fi24)
                                        _ = v84
                                    } else {
                                        if (!false) {
                                            var v85 string = ("hello" + "")
                                            var v86 uint32 = 65535
                                            var v87 uint = 4294967295
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        }
                                        switch fi24 {
                                        case 2147483647:
                                            var v88 float32 = 0.0
                                            _ = v88
                                        default:
                                            fmt.Println(fi10, fi10, fi24)
                                        }
                                        v4 = v4
                                    }
                                    fmt.Println(fi24, fi24)
                                    if (!(true && (866 != 809))) {
                                        v89 := []int{(fi24 / (((fi10 >> ((229) & 63))) | 1)), (^(fi24)), fi10, (fi10 * fi24)}
                                        _ = v89
                                    }
                                } else {
                                    var v90 uint16 = 2
                                    _ = v90
                                }
                            }
                        }
                    }
                } else {
                    v91 := GenS4{F0: 18446744073709551615}
                    var v92 int16 = 42
                    switch 16 {
                    case -580:
                        switch fi10 {
                        case -525:
                            if (!true) {
                                v92 -= (v92 % ((-1) | 1))
                                v4 += v4
                                v4 += v4
                                v93 := []int{fi24, fi10, int(v4), 0}
                                _ = v93
                            } else {
                                var v94 int = fi24
                                _ = v94
                            }
                            if false {
                                v95 := []int{0, fi10}
                                _ = v95
                            } else {
                                if (false || (-1 != -128)) {
                                    if false {
                                        var v96 uint32 = 65535
                                        v97 := GenS3{F0: -1}
                                        _ = v96
                                        _ = v97
                                    } else {
                                        fmt.Println(v4)
                                        var v98 int32 = -2147483648
                                        _ = v98
                                    }
                                    v4 += (v4 | 2)
                                    var v99 uint32 = 42
                                    _ = v99
                                } else {
                                    v100 := []int{-1, fi10}
                                    v101 := GenS7{F0: false, F1: 1, F2: 255}
                                    _ = v100
                                    _ = v101
                                }
                                v4 = (v4 & 223)
                            }
                        case 531:
                            fmt.Println(v4, fi24)
                        default:
                            v4 = (v4 | 660)
                            var v102 int = (fi10 * fi10)
                            v103 := []int{32768, fi10, -115}
                            _ = v102
                            _ = v103
                        }
                    }
                    _ = v91
                    _ = v92
                }
                v104 := GenS8{F0: 213, F1: 100}
                v4 += (v4 >> ((4294967295) & 63))
                _ = fi24
                _ = v104
            }
        default:
            v105 := []int{(fi10 | fi10), fi10, 9223372036854775807, (fi10 - (fi10 % ((int(v4)) | 1)))}
            _ = v105
        }
        v4 = v4
        _ = fi10
    }
    if (true || true) {
        var v106 uint = 2
        for fi107 := 0; fi107 < 19; fi107++ {
            v108 := GenS2{F0: (v4 >> ((357) & 63))}
            if (false || true) {
                v109 := []int{(fi107 >> ((786) & 63)), fi107, 42, fi107}
                var v110 int32 = 431
                v4 = (v4 - (v4 << ((v106) & 63)))
                _ = v109
                _ = v110
            } else {
                if (true || (!false)) {
                    var v111 int = fi107
                    v4 += v4
                    _ = v111
                }
                var v112 uint = 150
                _ = v112
            }
            fmt.Println(v4)
            _ = fi107
            _ = v108
        }
        _ = v106
    } else {
        v113 := GenS9{F0: 42, F1: 128, F2: 255}
        v4 = v4
        v114 := []int{-323, 0, 42, 32767}
        _ = v113
        _ = v114
    }
    _ = v4
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 bool = (32768 != 42)
    v115 := GenS10{F0: "\n\t"}
    v3 = v3
    v3 = v3
    if v3 {
        v3 = v3
        fmt.Println(v3)
    }
    v116 := []int{1, -477, 306, -444}
    for fi117 := 0; fi117 < 3; fi117++ {
        var v118 int32 = 2147483647
        _ = fi117
        _ = v118
    }
    var v119 int = -2147483648
    v116 = append(v116, (v119 / ((-347) | 1)))
    if false {
        for fi120 := 0; fi120 < 12; fi120++ {
            if v3 {
                fmt.Println(fi120)
            } else {
                v119 = (fi120 & (v119 | 9223372036854775807))
                var v121 uint32 = 65535
                var v122 int64 = genFn0()
                _ = v121
                _ = v122
            }
            v123 := GenS9{F0: -104, F1: 164, F2: 65535}
            fmt.Println(v3, v119)
            _ = fi120
            _ = v123
        }
        delete(v1, (v119 / (((v119 + 36)) | 1)))
        var v124 bool = (((1e10 * 1e10) * 1e10) >= (3.14159 - 0.0))
        _ = v124
    } else {
        if (42 > 331) {
            var v125 uint = 344
            v126 := GenS6{F0: 65535, F1: (^(v119))}
            v127 := GenS4{F0: 0}
            _ = v125
            _ = v126
            _ = v127
        }
    }
    for fi128 := 0; fi128 < 19; fi128++ {
        if v3 {
            if (true && (1000 > -32768)) {
                fmt.Println(v119, v3)
            } else {
                switch (fi128 % (((fi128 * v119)) | 1)) {
                case -2147483648:
                    if (!(265 <= 0)) {
                        var v129 int8 = 42
                        for fi130 := 0; fi130 < 19; fi130++ {
                            var v131 uint16 = 255
                            _ = fi130
                            _ = v131
                        }
                        var v132 uint16 = 2
                        _ = v129
                        _ = v132
                    } else {
                        if v3 {
                            var v133 uint16 = 0
                            fmt.Println(v119, fi128)
                            v2[v119] = fi128
                            _ = v133
                        }
                    }
                case 127:
                    v3 = v3
                case 42:
                    for fi134 := 0; fi134 < 3; fi134++ {
                        v135 := GenS7{F0: (!(564 < 42)), F1: 884, F2: 2}
                        v136 := GenS8{F0: 18446744073709551615, F1: 1}
                        _ = fi134
                        _ = v135
                        _ = v136
                    }
                }
                var v137 int16 = -129
                _ = v137
            }
        }
        v119 = fi128
        _ = fi128
    }
    v1[v119] = 408
    v138 := GenS10{F0: "foo bar"}
    switch v119 {
    case 127:
        if v3 {
            for fi139 := 0; fi139 < 8; fi139++ {
                var v140 int16 = 1
                for fi141 := 0; fi141 < 1; fi141++ {
                    v1[(v119 + (v119 / ((fi139) | 1)))] = 771
                    _ = len(v116)
                    v119 -= fi139
                    if (100 == 1) {
                        v142 := GenS11{F0: (v140 ^ v140)}
                        _ = len(v2)
                        if len(v116) >= 2 {
                            v116 = v116[1:len(v116)]
                        }
                        _ = v142
                    }
                    _ = fi141
                }
                _ = fi139
                _ = v140
            }
            fmt.Println(v3, v119)
        } else {
            var v143 bool = v3
            v144 := GenS12{F0: 582}
            v145 := GenS10{F0: ((("a" + "\n\t") + ("foo bar" + "foo bar")) + "a")}
            _ = v143
            _ = v144
            _ = v145
        }
    default:
        v119 += (v119 * -627)
    }
    for fi146 := 0; fi146 < 11; fi146++ {
        v119 = v119
        v147 := GenS10{F0: "foo bar"}
        _ = fi146
        _ = v147
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v116:", len(v116), v116)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v115
    _ = v116
    _ = v119
    _ = v138
    _ = v2
    _ = v3
}
