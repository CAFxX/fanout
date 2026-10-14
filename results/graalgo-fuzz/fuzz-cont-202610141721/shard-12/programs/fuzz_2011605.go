package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int16; F2 uint16 }
type GenS6 struct { F0 string; F1 uint8; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 string; F1 int; F2 uint64 }
type GenS10 struct { F0 int16; F1 string; F2 int16 }
type GenS11 struct { F0 int32; F1 int32; F2 int8 }
type GenS12 struct { F0 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint32 = 2147483648
    v3 := make(map[int]uint32)
    var v4 uint = 1
    var v5 string = "hello"
    var v6 int = 2147483647
    var v7 int32 = -1
    v7 -= (v7 - int32(v1))
    var v8 string = v5
    v9 := GenS1{F0: 2, F1: 0, F2: 774}
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint16 = 478
        if (!(v1 <= (v1 >> ((576) & 63)))) {
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := []int{584, int(v11), (fi10 << (((^(v4))) & 63)), (fi10 / ((int(v1)) | 1))}
                var v14 float32 = (0.0 * ((0.1 / 1.5) * 1e10))
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 int64 = (v1 - int64(v4))
            for fi16 := 0; fi16 < 11; fi16++ {
                switch (-(fi10)) {
                case -2147483648:
                    if (42 >= -128) {
                        v15 += v15
                        if (!((v11 != 761) || (1 == 1000))) {
                            v17 := GenS1{F0: 127, F1: 4294967295, F2: (v11 | 2)}
                            _ = v17
                        }
                    } else {
                        switch "a" {
                        case "hello":
                            var v18 bool = (!(127 > 42))
                            _ = v18
                        case "foo bar":
                            switch fi16 {
                            case -447:
                                v3[(fi16 / ((127) | 1))] = (^(v2))
                            case 127:
                                v19 := v3[fi10]
                                _ = v19
                            case 1:
                                var v20 int16 = 1
                                _ = v20
                            default:
                                if false {
                                    v1 += (v15 % ((v1) | 1))
                                }
                            }
                        case "a":
                            v3[(fi16 | 854)] = (v2 - v2)
                        }
                        if false {
                            switch 128 {
                            case 59:
                                switch v8 {
                                case "":
                                    v1 += (v1 | 42)
                                    var v21 int8 = 69
                                    _ = v21
                                default:
                                    var v22 uint32 = v2
                                    _ = v22
                                }
                            case -394:
                                v5 = v8
                            case 2147483647:
                                v11 = 65535
                            }
                            switch int(v11) {
                            case -272:
                                v23 := v3[(fi10 & -598)]
                                _ = v23
                            case 1:
                                var v24 string = v5
                                _ = v24
                            case -2147483648:
                                fmt.Println(v11, v15, v6)
                            default:
                                v4 += v4
                            }
                        }
                    }
                case -1:
                    var v25 uint16 = (v11 % (((v11 | v11)) | 1))
                    _ = v25
                case 0:
                    switch v6 {
                    case 9223372036854775807:
                        var v26 int16 = 42
                        _ = v26
                    case -1:
                        if true {
                            v8 = v5
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                var v27 float32 = 1e10
                                _ = v27
                            }
                            var v28 int16 = -32768
                            if ((v4 % (((v4 - (^(v4)))) | 1)) > 18446744073709551615) {
                                var v29 bool = (false && true)
                                _ = v29
                            } else {
                                var v30 bool = (2 >= (v2 - 676))
                                _ = v30
                            }
                            _ = v28
                        }
                    case 32767:
                        v31 := v3[(fi10 * (fi16 & v6))]
                        _ = v31
                    }
                }
                var v32 bool = true
                if (fi10 <= (fi16 % ((fi10) | 1))) {
                    switch (v6 ^ (^(fi16))) {
                    case 32767:
                        v33 := GenS2{F0: -1, F1: 42, F2: (v11 | v11)}
                        var v34 bool = (-117 < 127)
                        var v35 float32 = float32(1.5)
                        _ = v33
                        _ = v34
                        _ = v35
                    case -9223372036854775808:
                        fmt.Println(v1, v1, v5)
                    default:
                        fmt.Println(v6, v8, fi16)
                    }
                    var v36 uint64 = 420
                    if ((v32 || v32) && (fi10 <= fi16)) {
                        var v37 int = (fi16 >> ((2) & 63))
                        fmt.Println(v4)
                        _ = v37
                    } else {
                        fmt.Println(v4, v7)
                        fmt.Println(v11)
                    }
                    _ = v36
                }
                _ = fi16
                _ = v32
            }
            _ = v15
        } else {
            _ = len(v3)
        }
        _ = fi10
        _ = v11
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        v5 = (v5 + "foo bar")
        if (-128 < 0) {
            v3[v6] = 42
        } else {
            for fi39 := 0; fi39 < 20; fi39++ {
                switch (v8 + "\n\t") {
                case "\n\t":
                    var v40 uint32 = (v2 >> ((v4) & 63))
                    _ = v40
                case "a":
                    v5 = v5
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        if false {
                            if false {
                                fmt.Println(fi39, fi38)
                                switch fi38 {
                                case -469:
                                    if (1 == 185) {
                                        var v41 int32 = (v7 + v7)
                                        fmt.Println(v2)
                                        _ = v41
                                    } else {
                                        switch (fi39 - 42) {
                                        case 128:
                                            var v42 uint16 = 956
                                            _ = v42
                                        default:
                                            if false {
                                                if (!false) {
                                                    v43 := GenS2{F0: -32768, F1: 0, F2: uint16(v4)}
                                                    v44 := []int{2147483647, (v6 * int(v2)), (fi39 % ((fi38) | 1)), (v6 >> ((1) & 63))}
                                                    v6 += (v6 - (v6 - fi38))
                                                    switch (fi38 / ((v6) | 1)) {
                                                    case 32768:
                                                        var v45 float32 = 0.1
                                                        _ = v45
                                                    default:
                                                        _ = len(v44)
                                                    }
                                                    fmt.Println(v5, v2, fi39)
                                                    if false {
                                                        delete(v3, fi38)
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    v2 += 468
                                                    if false {
                                                        fmt.Println(fi39)
                                                        v2 += 2
                                                        if false {
                                                            v5 = v8
                                                            v3[(-(v6))] = 2147483647
                                                            if (!(!(0 == 2))) {
                                                                var v46 int32 = v7
                                                                v47 := GenS2{F0: -124, F1: 128, F2: uint16(v4)}
                                                                _ = v46
                                                                _ = v47
                                                            } else {
                                                                if false {
                                                                    if false {
                                                                        if (!true) {
                                                                            if true {
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    if (0 > (v2 ^ (v2 << ((v4) & 63)))) {
                                                                                        if (68 == 295) {
                                                                                            var v48 int8 = 65
                                                                                            var v49 uint64 = 0
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                        }
                                                                                        v5 = v5
                                                                                    }
                                                                                case "hello":
                                                                                    if (!((369 <= 2) || true)) {
                                                                                        delete(v3, v6)
                                                                                        v50 := []int{(v6 / (((fi39 / (((fi39 / (((fi38 & fi38)) | 1))) | 1))) | 1)), (v6 & (fi39 << (((v4 % (((v4 & (v4 ^ 1))) | 1))) & 63))), (fi39 >> ((v4) & 63))}
                                                                                        _ = v50
                                                                                    }
                                                                                case "a":
                                                                                    if (0.1 > 0.0) {
                                                                                        if true {
                                                                                            switch v5 {
                                                                                            case "\n\t":
                                                                                                v51 := []int{fi39, (fi39 + (fi39 << (((v4 | (v4 + 899))) & 63)))}
                                                                                                _ = v51
                                                                                            case "":
                                                                                                switch v5 {
                                                                                                case "":
                                                                                                    v52 := GenS2{F0: -32768, F1: 255, F2: 42}
                                                                                                    _ = v52
                                                                                                case "a":
                                                                                                    fmt.Println(v8, v5)
                                                                                                default:
                                                                                                    switch "" {
                                                                                                    case "foo bar":
                                                                                                        v5 = (("hello" + v5) + ((((v5 + "\n\t") + "") + (v5 + v5)) + "hello"))
                                                                                                    case "":
                                                                                                        v7 += v7
                                                                                                    default:
                                                                                                        switch fi38 {
                                                                                                        case 2147483647:
                                                                                                            v53 := GenS2{F0: 662, F1: 2, F2: 645}
                                                                                                            _ = v53
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v54 := []int{}
                                                                                                _ = v54
                                                                                            }
                                                                                            if (false && (true && (((v6 >> (((^(v4))) & 63)) == fi38) || (!(false && false))))) {
                                                                                                v5 = ("foo bar" + v8)
                                                                                            }
                                                                                            fmt.Println(v4)
                                                                                            var v55 float64 = (1.5 + float64(1e10))
                                                                                            if true {
                                                                                                v1 -= v1
                                                                                                v56 := GenS3{F0: int16(v1), F1: (v7 / (((v7 - 2147483647)) | 1))}
                                                                                                v57 := v3[v6]
                                                                                                v58 := v3[fi38]
                                                                                                if false {
                                                                                                    v59 := GenS1{F0: 1, F1: 42, F2: 1}
                                                                                                    var v60 int8 = 127
                                                                                                    if (true && (fi38 != fi38)) {
                                                                                                        if (!false) {
                                                                                                            var v61 uint = v4
                                                                                                            _ = v61
                                                                                                        }
                                                                                                        var v62 int8 = (v60 & 42)
                                                                                                        v63 := v3[(fi39 - (v6 / (((fi38 & (v6 * 43))) | 1)))]
                                                                                                        _ = v62
                                                                                                        _ = v63
                                                                                                    }
                                                                                                    _ = v59
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    if ((v55 != 0.1) || true) {
                                                                                                        var v64 int16 = -231
                                                                                                        var v65 string = v8
                                                                                                        if true {
                                                                                                            v66 := GenS2{F0: (v64 << (((^(v4))) & 63)), F1: 126, F2: 159}
                                                                                                            _ = v66
                                                                                                        }
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                    } else {
                                                                                                        v67 := []int{}
                                                                                                        _ = v67
                                                                                                    }
                                                                                                }
                                                                                                _ = v56
                                                                                                _ = v57
                                                                                                _ = v58
                                                                                            }
                                                                                            _ = v55
                                                                                        }
                                                                                        fmt.Println(v2)
                                                                                        v6 -= (fi39 | (fi38 | (fi39 + fi39)))
                                                                                    } else {
                                                                                        v1 -= (v1 * 693)
                                                                                    }
                                                                                default:
                                                                                    v68 := GenS4{F0: 0}
                                                                                    _ = v68
                                                                                }
                                                                            } else {
                                                                                v7 += (v7 + (v7 & 0))
                                                                            }
                                                                            var v69 float32 = -2.25
                                                                            _ = v69
                                                                        }
                                                                        if true {
                                                                            if (0 == 354) {
                                                                                _ = len(v3)
                                                                                v8 = v5
                                                                                v8 = ("\n\t" + "a")
                                                                                v7 -= v7
                                                                                delete(v3, fi39)
                                                                                var v70 float32 = 0.1
                                                                                _ = v70
                                                                            }
                                                                            var v71 bool = true
                                                                            var v72 int16 = 1
                                                                            _ = v71
                                                                            _ = v72
                                                                        } else {
                                                                            var v73 int8 = 8
                                                                            _ = v73
                                                                        }
                                                                        var v74 uint32 = v2
                                                                        _ = v74
                                                                    } else {
                                                                        v4 = 815
                                                                        v6 = -7
                                                                    }
                                                                } else {
                                                                    v75 := GenS2{F0: -603, F1: 128, F2: 1}
                                                                    v76 := GenS1{F0: 206, F1: 0, F2: 42}
                                                                    v77 := v3[(fi38 << (((^(v4))) & 63))]
                                                                    _ = v75
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                                var v78 uint16 = 43
                                                                fmt.Println(v5, v6, v2)
                                                                switch v8 {
                                                                case "a":
                                                                    var v79 uint32 = (^(v2))
                                                                    _ = v79
                                                                case "\n\t":
                                                                    v7 = v7
                                                                case "":
                                                                    v80 := []int{fi39, v6}
                                                                    _ = v80
                                                                }
                                                                _ = v78
                                                            }
                                                        } else {
                                                            fmt.Println(fi39, v1, v8)
                                                            var v81 string = ((v5 + v5) + "\n\t")
                                                            var v82 string = v5
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    } else {
                                                        v83 := GenS4{F0: 42}
                                                        _ = v83
                                                    }
                                                    var v84 int16 = 1
                                                    v85 := []int{0, v6, v6, (fi39 ^ (fi39 * 407))}
                                                    _ = v84
                                                    _ = v85
                                                }
                                                switch 128 {
                                                case 457:
                                                    if true {
                                                        fmt.Println(fi39, v1, v7)
                                                        v2 = 799
                                                        v86 := GenS5{F0: 653, F1: 0, F2: 115}
                                                        _ = v86
                                                    }
                                                    v87 := GenS6{F0: v5, F1: 253, F2: (2 < 18446744073709551615)}
                                                    delete(v3, fi38)
                                                    _ = v87
                                                default:
                                                    v6 = 166
                                                }
                                                var v88 float64 = ((-2.25 - 3.14159) * 1e10)
                                                _ = v88
                                            } else {
                                                if false {
                                                    var v89 int = v6
                                                    switch (fi39 - v89) {
                                                    case 1:
                                                        v5 = "hello"
                                                    case 2147483647:
                                                        v90 := GenS7{F0: 65535}
                                                        _ = v90
                                                    case -9223372036854775808:
                                                        v91 := GenS7{F0: 1}
                                                        _ = v91
                                                    default:
                                                        var v92 uint64 = 1
                                                        _ = v92
                                                    }
                                                    _ = v89
                                                } else {
                                                    fmt.Println(v1, v4)
                                                    v93 := []int{128, fi38, 9223372036854775807}
                                                    if len(v93) >= 2 {
                                                        v93 = v93[1:len(v93)]
                                                    }
                                                    _ = v93
                                                }
                                                var v94 uint = (v4 << (((^(v4))) & 63))
                                                _ = v94
                                            }
                                        }
                                    }
                                case 32767:
                                    v4 -= v4
                                default:
                                    fmt.Println(v5, fi39, v5)
                                }
                                if (255 > 33) {
                                    if true {
                                        switch 42 {
                                        case -2147483648:
                                            v95 := []int{(fi38 >> ((v4) & 63)), v6, v6}
                                            _ = v95
                                        case 127:
                                            v4 += 1
                                        case 380:
                                            v96 := []int{fi38}
                                            _ = v96
                                        }
                                    } else {
                                        var v97 uint64 = 1
                                        _ = v97
                                    }
                                }
                                var v98 uint = (v4 << (((v4 % ((42) | 1))) & 63))
                                _ = v98
                            }
                            fmt.Println(v5, v5, v8)
                            if (true && (0.0 < 0.0)) {
                                var v99 int16 = int16(v7)
                                _ = v99
                            }
                        }
                    case "\n\t":
                        var v100 int = -919
                        _ = v100
                    }
                }
                switch (v8 + "\n\t") {
                case "foo bar":
                    v5 = (v5 + "")
                case "\n\t":
                    var v101 string = v5
                    _ = v101
                }
                if true {
                    if false {
                        var v102 float32 = 0.1
                        _ = v102
                    }
                    if false {
                        var v103 string = v8
                        v104 := GenS8{F0: 2, F1: v7}
                        _ = v103
                        _ = v104
                    }
                    fmt.Println(v6)
                }
                if false {
                    var v105 uint64 = 1
                    var v106 bool = true
                    if (!(42 <= (v2 - v2))) {
                        v107 := GenS2{F0: 1, F1: 85, F2: 0}
                        v108 := GenS1{F0: 2, F1: v105, F2: 32767}
                        var v109 int32 = v7
                        _ = v107
                        _ = v108
                        _ = v109
                    }
                    _ = v105
                    _ = v106
                }
                fmt.Println(fi38, v2)
                _ = fi39
            }
            v110 := GenS2{F0: 42, F1: 42, F2: 839}
            v4 += 2
            _ = v110
        }
        _ = fi38
    }
    if false {
        var v111 int32 = v7
        for fi112 := 0; fi112 < 1; fi112++ {
            var v113 uint8 = 42
            var v114 int16 = -509
            _ = fi112
            _ = v113
            _ = v114
        }
        _ = v111
    } else {
        v115 := GenS7{F0: 1}
        fmt.Println(v2, v6, v4)
        _ = v115
    }
    v116 := []int{(v6 | -9223372036854775808), (v6 >> ((v4) & 63))}
    if true {
        v117 := GenS7{F0: 2}
        v4 = 962
        for fi118 := 0; fi118 < 5; fi118++ {
            for fi119 := 0; fi119 < 1; fi119++ {
                v6 -= (^(v6))
                _ = fi119
            }
            _ = fi118
        }
        _ = v117
    } else {
        if false {
            var v120 uint16 = 255
            v121 := GenS9{F0: v5, F1: (v6 * (v6 / (((v6 | 42)) | 1))), F2: 42}
            v122 := v3[v6]
            _ = v120
            _ = v121
            _ = v122
        }
    }
    for fi123 := 0; fi123 < 10; fi123++ {
        switch 0 {
        case -46:
            var v124 int64 = 586
            _ = v124
        case 570:
            var v125 bool = false
            _ = v125
        }
        v126 := GenS9{F0: ("a" + v8), F1: (-(fi123)), F2: 18446744073709551615}
        var v127 float32 = (((-2.25 * 3.14159) * 1e10) / 1.5)
        _ = fi123
        _ = v126
        _ = v127
    }
    for fi128 := 0; fi128 < 15; fi128++ {
        v2 += v2
        _ = fi128
    }
    switch "" {
    case "a":
        for fi129 := 0; fi129 < 20; fi129++ {
            v1 -= v1
            v130 := GenS10{F0: 258, F1: "hello", F2: -588}
            if true {
                v8 = ((("foo bar" + v8) + (("hello" + (v8 + ("a" + (v5 + "hello")))) + "foo bar")) + ("foo bar" + "hello"))
                var v131 uint32 = 2
                _ = v131
            } else {
                v8 = ((v8 + (("" + v8) + "hello")) + "a")
                v132 := GenS11{F0: (v7 / ((v7) | 1)), F1: v7, F2: -67}
                v133 := GenS11{F0: (v7 << ((v4) & 63)), F1: 0, F2: 42}
                _ = v132
                _ = v133
            }
            _ = fi129
            _ = v130
        }
    case "":
        var v134 uint64 = 42
        _ = v134
    case "\n\t":
        switch v6 {
        case 0:
            for fi135 := 0; fi135 < 7; fi135++ {
                v5 = ("a" + "foo bar")
                var v136 float64 = 0.0
                _ = fi135
                _ = v136
            }
        case 1:
            v116 = append(v116, v6)
        case -9223372036854775808:
            var v137 int64 = v1
            _ = v137
        default:
            v138 := GenS12{F0: 42}
            _ = v138
        }
    }
    v4 += v4
    var v139 uint16 = 255
    var v140 uint64 = 4294967295
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v116:", len(v116), v116)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v116
    _ = v139
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
