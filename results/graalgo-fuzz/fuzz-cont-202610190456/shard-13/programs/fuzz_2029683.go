package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint8; F1 bool; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int; F2 int16 }
type GenS8 struct { F0 uint; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 uint16; F1 string; F2 string }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32; F1 uint32; F2 uint }
type GenS15 struct { F0 int64; F1 int8; F2 uint64 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 4; fi6++ {
        switch fi6 {
        case -9223372036854775808:
            v7 := []int{}
            _ = v7
        default:
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := []int{(fi6 << ((65535) & 63)), fi6, 32768}
                if false {
                    fmt.Println(fi8, fi6)
                    v10 := GenS1{F0: 255, F1: (fi6 >> ((18446744073709551615) & 63))}
                    var v11 uint = 42
                    if true {
                        switch ((("a" + (("foo bar" + "foo bar") + "foo bar")) + (("a" + ("\n\t" + ("a" + ("a" + "\n\t")))) + "")) + "") {
                        case "":
                            var v12 bool = (51 > 127)
                            _ = v12
                        case "foo bar":
                            if ((true && true) || true) {
                                if false {
                                    v11 -= (v11 << (((v11 << (((v11 | 392)) & 63))) & 63))
                                    var v13 float64 = float64((1.5 * (-2.25 * (1.5 + (0.0 * 3.14159)))))
                                    _ = v13
                                } else {
                                    var v14 float64 = ((-2.25 / 1.0) - (float64(3.14159) / (1.5 + 0.1)))
                                    _ = v14
                                }
                                fmt.Println(fi6, v11, fi6)
                                switch "a" {
                                case "\n\t":
                                    var v15 int = (fi8 << ((18446744073709551615) & 63))
                                    _ = v15
                                }
                            } else {
                                var v16 uint32 = 124
                                var v17 string = ""
                                var v18 uint16 = 48
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        case "\n\t":
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        default:
                            v11 += (v11 * v11)
                        }
                        var v19 float32 = float32(1.5)
                        v19 = ((v19 + ((3.14159 - v19) - 0.0)) / float32(((float64(v19) + -2.25) * 0.0)))
                        _ = v19
                    }
                    switch (("" + "a") + "") {
                    case "\n\t":
                        v11 += v11
                    case "hello":
                        v11 = 255
                    case "":
                        switch ("\n\t" + "a") {
                        case "":
                            var v20 int32 = 1
                            _ = v20
                        default:
                            var v21 int = (fi8 & fi6)
                            _ = v21
                        }
                    default:
                        var v22 bool = true
                        _ = v22
                    }
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        }
        if (!true) {
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS2{F0: 0, F1: 127, F2: 1}
                v25 := GenS3{F0: 386}
                switch fi23 {
                case -82:
                    var v26 int8 = 42
                    _ = v26
                case 0:
                    if false {
                        var v27 uint32 = 91
                        var v28 uint16 = 0
                        switch "\n\t" {
                        case "":
                            v29 := GenS4{F0: "hello", F1: (v28 << ((0) & 63))}
                            _ = v29
                        default:
                            var v30 int8 = -1
                            _ = v30
                        }
                        _ = v27
                        _ = v28
                    }
                case 32767:
                    switch -2147483648 {
                    case -2147483648:
                        var v31 string = ("a" + ("hello" + "a"))
                        _ = v31
                    default:
                        var v32 int = (fi23 + fi23)
                        _ = v32
                    }
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                if true {
                    v34 := GenS5{F0: 2, F1: (1.5 >= 1.5), F2: "a"}
                    if (!(-2.25 < (((0.0 + 1e10) - 0.1) - -2.25))) {
                        if false {
                            fmt.Println(fi6, fi33)
                            v35 := []int{(fi33 - fi6), (fi6 >> ((255) & 63)), fi6}
                            _ = v35
                        }
                        v36 := GenS4{F0: ("" + "foo bar"), F1: 42}
                        v37 := []int{fi6}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int16 = -847
                        v39 := []int{(fi33 & (fi6 % ((fi6) | 1))), 9223372036854775807, fi6}
                        _ = v38
                        _ = v39
                    }
                    _ = v34
                } else {
                    var v40 uint = 980
                    _ = v40
                }
                fmt.Println(fi33)
                v41 := GenS6{F0: 42}
                _ = fi33
                _ = v41
            }
            if (!(741 < 754)) {
                v42 := GenS7{F0: "foo bar", F1: 4294967295, F2: 542}
                switch ("a" + "foo bar") {
                case "hello":
                    var v43 float64 = ((1e10 * 0.1) - 0.1)
                    _ = v43
                }
                _ = v42
            } else {
                if false {
                    var v44 int16 = -646
                    _ = v44
                }
                for fi45 := 0; fi45 < 6; fi45++ {
                    if true {
                        var v46 int = (-(fi6))
                        v46 = 0
                        fmt.Println(v46)
                        _ = v46
                    } else {
                        if false {
                            if true {
                                var v47 uint16 = 32767
                                fmt.Println(fi45)
                                _ = v47
                            } else {
                                v48 := []int{fi6, fi6}
                                var v49 float64 = (((3.14159 / 1.5) - (1.5 * 3.14159)) / 1.0)
                                _ = v48
                                _ = v49
                            }
                        } else {
                            var v50 uint = 649
                            var v51 int = (fi6 | (fi45 << ((v50) & 63)))
                            v50 = (v50 - (v50 * v50))
                            v51 -= (^(fi6))
                            _ = v50
                            _ = v51
                        }
                    }
                    _ = fi45
                }
                v52 := []int{fi6, (fi6 * (fi6 & (fi6 << ((4294967295) & 63)))), (fi6 << ((894) & 63)), (fi6 & (fi6 % ((-756) | 1)))}
                _ = v52
            }
        } else {
            if false {
                var v53 uint32 = 387
                _ = v53
            }
            switch (fi6 - 42) {
            case 9223372036854775807:
                v54 := []int{2147483647}
                _ = v54
            case 32767:
                v55 := GenS3{F0: 0}
                _ = v55
            }
        }
        _ = fi6
    }
    var v56 uint64 = 18446744073709551615
    v56 = (v56 ^ v56)
    var v57 int64 = 42
    _ = v56
    _ = v57
    return 0
}

func main() {
    var v1 int = 32768
    v2 := []uint32{42}
    v3 := []uint32{2, 147}
    var v4 uint64 = 4294967295
    var v5 uint16 = 65535
    if (int(v5) >= v1) {
        var v58 float32 = 3.14159
        var v59 int = (^(v1))
        _ = v58
        _ = v59
    }
    var v60 int64 = 1
    fmt.Println(v1, v5)
    for fi61 := 0; fi61 < 14; fi61++ {
        var v62 int8 = 4
        v5 = (v5 >> ((709) & 63))
        for fi63 := 0; fi63 < 13; fi63++ {
            if (409 < 1) {
                v64 := GenS8{F0: 42, F1: -1, F2: 187}
                v2 = append(v2, 42)
                v62 -= v62
                _ = v64
            } else {
                if ((!true) && ((fi61 & 128) <= int(v5))) {
                    if ((true || true) && false) {
                        v65 := GenS8{F0: 255, F1: v60, F2: 128}
                        if false {
                            var v66 uint32 = 398
                            _ = v66
                        } else {
                            var v67 uint16 = v5
                            if false {
                                v1 = fi63
                            } else {
                                var v68 uint16 = v5
                                var v69 uint8 = 87
                                if ((v62 & v62) > 1) {
                                    fmt.Println(v60, v60)
                                }
                                _ = v68
                                _ = v69
                            }
                            _ = v67
                        }
                        var v70 int8 = (v62 - (v62 & (v62 ^ int8(fi61))))
                        _ = v65
                        _ = v70
                    }
                    if (fi61 != (v1 | (v1 & 53))) {
                        switch ("foo bar" + "a") {
                        case "":
                            switch "\n\t" {
                            case "hello":
                                _ = len(v2)
                            }
                        case "foo bar":
                            v4 = v4
                        case "hello":
                            if (-218 > 435) {
                                switch "a" {
                                case "a":
                                    v62 = -128
                                default:
                                    if (0.0 <= 0.1) {
                                        var v71 string = ""
                                        var v72 bool = (-915 <= (fi63 * (fi61 ^ fi61)))
                                        var v73 uint32 = 2147483648
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v1 += 32767
                                        fmt.Println(v1, v62)
                                        v1 -= (fi61 - 447)
                                    }
                                }
                                v62 = v62
                            } else {
                                var v74 uint = 1
                                var v75 uint32 = 24
                                switch 32768 {
                                case 1:
                                    var v76 string = "foo bar"
                                    _ = v76
                                case 42:
                                    if false {
                                        v4 -= (v4 | 4294967295)
                                        v77 := GenS9{F0: (v75 >> (((^(v74))) & 63)), F1: (v75 << ((v74) & 63)), F2: (v5 + v5)}
                                        var v78 float32 = (1.5 + 1.5)
                                        var v79 int32 = -1
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    } else {
                                        v62 += v62
                                        var v80 uint64 = 161
                                        var v81 uint8 = 2
                                        var v82 float32 = 1.5
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                default:
                                    var v83 int = v1
                                    _ = v83
                                }
                                if len(v3) > 0 {
                                    v3[0] = (v75 ^ v75)
                                }
                                fmt.Println(v5)
                                fmt.Println(fi61, v4)
                                _ = len(v2)
                                _ = v74
                                _ = v75
                            }
                        default:
                            fmt.Println(v60, v60)
                        }
                        v5 += 42
                        switch "a" {
                        case "a":
                            _ = len(v2)
                        }
                        if false {
                            v1 -= (fi61 - (v1 ^ fi63))
                            var v84 int8 = v62
                            v85 := GenS8{F0: 4294967295, F1: v60, F2: 2}
                            _ = v84
                            _ = v85
                        }
                    } else {
                        _ = len(v3)
                        var v86 uint8 = 128
                        _ = v86
                    }
                }
                fmt.Println(v60)
            }
            if ("hello" <= "a") {
                fmt.Println(v1, v4)
                v60 = (^(v60))
                var v87 int8 = (-(v62))
                _ = v87
            } else {
                fmt.Println(v62, v1)
                var v88 float32 = ((float32(0.1) / (((((1e10 / -2.25) / -2.25) + (float32((1.5 - 0.0)) - 0.0)) * float32(1e10)) - 0.1)) / -2.25)
                _ = v88
            }
            _ = fi63
        }
        _ = fi61
        _ = v62
    }
    if false {
        fmt.Println(v4, v60, v5)
        for fi89 := 0; fi89 < 17; fi89++ {
            for fi90 := 0; fi90 < 9; fi90++ {
                var v91 uint16 = (^(v5))
                var v92 bool = ((v91 != (v91 - 262)) || true)
                if len(v3) > 0 {
                    v3[0] = uint32(v4)
                }
                _ = fi90
                _ = v91
                _ = v92
            }
            v93 := GenS2{F0: 7, F1: -128, F2: 0}
            _ = fi89
            _ = v93
        }
        fmt.Println(v60, v60, v1)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi94 := 0; fi94 < 17; fi94++ {
            for fi95 := 0; fi95 < 4; fi95++ {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                var v96 uint = 0
                _ = fi95
                _ = v96
            }
            var v97 uint32 = 42
            fmt.Println(v97)
            _ = fi94
            _ = v97
        }
        _ = len(v3)
    }
    fmt.Println(v60, v1)
    if (0 <= 80) {
        v4 = v4
        v5 += v5
    } else {
        fmt.Println(v5, v1)
        switch ("\n\t" + "a") {
        case "":
            if false {
                var v98 int8 = 1
                for fi99 := 0; fi99 < 10; fi99++ {
                    for fi100 := 0; fi100 < 4; fi100++ {
                        fmt.Println(v5)
                        var v101 int64 = -1
                        v102 := GenS5{F0: 255, F1: (!true), F2: (("" + ("foo bar" + "hello")) + "hello")}
                        _ = fi100
                        _ = v101
                        _ = v102
                    }
                    switch "hello" {
                    case "hello":
                        for fi103 := 0; fi103 < 5; fi103++ {
                            v98 = 1
                            var v104 int8 = v98
                            _ = fi103
                            _ = v104
                        }
                    case "":
                        var v105 float32 = -2.25
                        _ = v105
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi99
                }
                for fi106 := 0; fi106 < 15; fi106++ {
                    if false {
                        fmt.Println(v98)
                    }
                    _ = fi106
                }
                _ = v98
            }
        case "hello":
            v107 := GenS10{F0: 2147483648, F1: ("a" + "")}
            _ = v107
        case "foo bar":
            fmt.Println(v4, v60, v5)
        }
    }
    var v108 uint = 891
    v109 := GenS6{F0: v4}
    _ = len(v3)
    v60 = (^(v60))
    switch v1 {
    case 128:
        if (false || (!(v60 == v60))) {
            if len(v2) > 0 {
                v2[0] = 31
            }
            if false {
                v110 := GenS11{F0: v5, F1: ("\n\t" + "a"), F2: "\n\t"}
                _ = v110
            } else {
                v111 := GenS3{F0: v5}
                var v112 uint64 = (v4 & (v4 + 78))
                _ = v111
                _ = v112
            }
            var v113 bool = false
            _ = v113
        } else {
            _ = len(v3)
        }
    case 836:
        if true {
            fmt.Println(v108, v108)
            v3 = append(v3, uint32(v5))
            v1 += (v1 << (((v108 / (((v108 << ((v108) & 63))) | 1))) & 63))
        }
    case 42:
        for fi114 := 0; fi114 < 5; fi114++ {
            v1 += (fi114 / ((fi114) | 1))
            fmt.Println(v108)
            v1 += (v1 ^ 127)
            _ = fi114
        }
    }
    v1 += (-(v1))
    if true {
        if (false || true) {
            if (100 != 42) {
                v115 := GenS7{F0: ("a" + "foo bar"), F1: -949, F2: 1}
                if (!((true && (!(255 != v108))) && (-27 != 42))) {
                    v5 = v5
                    if true {
                        _ = len(v2)
                        v116 := GenS10{F0: 1, F1: (("\n\t" + "foo bar") + "hello")}
                        _ = v116
                    } else {
                        fmt.Println(v4)
                        v117 := GenS11{F0: v5, F1: (("" + (("hello" + "foo bar") + "\n\t")) + "\n\t"), F2: ("a" + "")}
                        for fi118 := 0; fi118 < 5; fi118++ {
                            if true {
                                for fi119 := 0; fi119 < 8; fi119++ {
                                    if true {
                                        var v120 int8 = 1
                                        v121 := GenS12{F0: "foo bar", F1: 0}
                                        _ = v120
                                        _ = v121
                                    } else {
                                        if true {
                                            v122 := GenS13{F0: (v108 % (((v108 << (((v108 % (((v108 ^ v108)) | 1))) & 63))) | 1))}
                                            switch -906 {
                                            case -551:
                                                v60 -= (v60 | (v60 * v60))
                                            case -1:
                                                v4 = (^(v4))
                                            }
                                            v5 -= 42
                                            _ = v122
                                        } else {
                                            v5 = v5
                                            v60 -= v60
                                            v1 += -186
                                            var v123 float64 = 1.5
                                            if len(v3) > 0 {
                                                v3[0] = uint32(v108)
                                            }
                                            v124 := GenS12{F0: "hello", F1: 2}
                                            _ = v123
                                            _ = v124
                                        }
                                    }
                                    _ = fi119
                                }
                                v125 := GenS14{F0: 192, F1: 65535, F2: uint(v5)}
                                v4 = (v4 - v4)
                                _ = v125
                            } else {
                                var v126 int8 = 8
                                _ = len(v2)
                                switch int(v60) {
                                case 324:
                                    if (26 <= 2) {
                                        v127 := GenS4{F0: "a", F1: 2}
                                        switch (v1 + (v1 ^ fi118)) {
                                        case -723:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            v128 := GenS9{F0: 2147483648, F1: 42, F2: (^(v5))}
                                            _ = v128
                                        }
                                        _ = v127
                                    } else {
                                        v60 -= (v60 / ((9223372036854775807) | 1))
                                        switch (v1 & (fi118 * (v1 % (((^(fi118))) | 1)))) {
                                        case 127:
                                            v129 := GenS8{F0: 2, F1: v60, F2: 2}
                                            _ = v129
                                        }
                                        var v130 int8 = v126
                                        _ = v130
                                    }
                                default:
                                    var v131 uint32 = 65535
                                    _ = v131
                                }
                                _ = v126
                            }
                            for fi132 := 0; fi132 < 11; fi132++ {
                                v60 -= 9223372036854775807
                                v3 = append(v3, 4294967295)
                                v60 = (v60 % (((v60 | v60)) | 1))
                                _ = fi132
                            }
                            _ = fi118
                        }
                        _ = v117
                    }
                }
                v1 = (v1 & (v1 << ((255) & 63)))
                _ = v115
            } else {
                var v133 uint = 65535
                if (false && true) {
                    var v134 int32 = 0
                    _ = v134
                }
                _ = v133
            }
        }
        if true {
            if false {
                for fi135 := 0; fi135 < 13; fi135++ {
                    for fi136 := 0; fi136 < 2; fi136++ {
                        _ = len(v2)
                        var v137 int32 = 42
                        if ((((v5 > (v5 + (v5 - v5))) || true) && (v60 > (-(v60)))) || (false && false)) {
                            v108 += v108
                            if false {
                                v138 := GenS15{F0: int64(v5), F1: 100, F2: (^(v4))}
                                if (v4 <= 18446744073709551615) {
                                    var v139 float64 = 3.14159
                                    v108 += v108
                                    _ = v139
                                } else {
                                    v60 += 42
                                }
                                _ = v138
                            } else {
                                if false {
                                    if (-494 >= v1) {
                                        fmt.Println(v1, fi136, v4)
                                        fmt.Println(v5)
                                    }
                                }
                            }
                        } else {
                            if ((3.14159 * 1.5) >= 0.0) {
                                v137 -= 42
                                v108 = (v108 / ((v108) | 1))
                            }
                            if (!(9223372036854775807 <= 1)) {
                                v4 -= v4
                                if ((-(fi135)) == (fi135 + 9223372036854775807)) {
                                    var v140 int64 = (v60 - (v60 >> ((v108) & 63)))
                                    _ = v140
                                } else {
                                    switch -9223372036854775808 {
                                    case 32767:
                                        v141 := GenS10{F0: 860, F1: "hello"}
                                        _ = v141
                                    case 127:
                                        v1 = fi136
                                    case 2147483647:
                                        _ = len(v3)
                                    }
                                }
                                if true {
                                    var v142 uint16 = v5
                                    _ = v142
                                }
                                var v143 bool = (false && false)
                                _ = len(v2)
                                _ = v143
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(fi136, v108, fi136)
                            v2 = append(v2, 0)
                        }
                        var v144 int8 = -1
                        _ = fi136
                        _ = v137
                        _ = v144
                    }
                    _ = fi135
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = v1
        }
    }
    switch ("" + "a") {
    case "foo bar":
        v145 := GenS10{F0: 0, F1: "a"}
        _ = v145
    case "\n\t":
        var v146 int32 = 612
        _ = v146
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
