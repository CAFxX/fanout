package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 int64; F1 int64; F2 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 int16 }
type GenS11 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS12 struct { F0 int8; F1 uint8; F2 int16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64; F1 int; F2 int }
type GenS16 struct { F0 uint64; F1 int }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint; F1 uint32 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint8; F1 uint }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 uint; F1 int64 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 int32; F1 uint8; F2 int64 }
type GenS27 struct { F0 int64; F1 int32; F2 uint64 }
type GenS28 struct { F0 uint; F1 uint64 }
type GenS29 struct { F0 int16; F1 uint64 }
type GenS30 struct { F0 uint }
type GenS31 struct { F0 string; F1 int }

func main() {
    var v1 int16 = 1000
    var v2 string = "foo bar"
    var v3 uint64 = 1
    v4 := make(map[uint32]int)
    var v5 uint8 = 128
    for fi6 := 0; fi6 < 19; fi6++ {
        if ((v3 ^ (v3 >> ((4294967295) & 63))) < v3) {
            delete(v4, 4294967295)
            var v7 string = v2
            var v8 float32 = float32(1.5)
            _ = v7
            _ = v8
        }
        fmt.Println(v5)
        _ = fi6
    }
    v2 = v2
    var v9 uint32 = 1
    switch ((v2 + v2) + v2) {
    case "":
        for fi10 := 0; fi10 < 18; fi10++ {
            v9 -= v9
            fmt.Println(v5, v5, v2)
            switch v2 {
            case "hello":
                switch (v2 + "a") {
                case "a":
                    fmt.Println(v9)
                }
            case "":
                v1 += (v1 ^ (v1 ^ v1))
            }
            _ = fi10
        }
    case "a":
        var v11 uint16 = 65535
        _ = v11
    default:
        var v12 int8 = 12
        _ = v12
    }
    v13 := GenS1{F0: (v5 & v5), F1: (false || false)}
    var v14 uint = 18446744073709551615
    for fi15 := 0; fi15 < 8; fi15++ {
        v4[(v9 & v9)] = (fi15 | fi15)
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(v2, v9, v3)
            v9 -= v9
            if ((2147483647 >= -103) && true) {
                if (449 == fi16) {
                    var v17 float32 = -2.25
                    var v18 int16 = v1
                    _ = v17
                    _ = v18
                }
                if ((v9 ^ (v9 - (v9 * (v9 / ((v9) | 1))))) >= 65535) {
                    if false {
                        if (v14 == v14) {
                            v19 := GenS1{F0: (v5 + (v5 / ((255) | 1))), F1: true}
                            var v20 int = 746
                            _ = v19
                            _ = v20
                        } else {
                            v3 = 1
                            v1 -= (^(v1))
                            v3 += (^(v3))
                        }
                        if (("foo bar" + "hello") > "foo bar") {
                            fmt.Println(v5, v2)
                            var v21 int32 = 42
                            _ = v21
                        } else {
                            var v22 float64 = (1e10 * 0.1)
                            _ = v22
                        }
                    }
                    v23 := GenS2{F0: (^(v5)), F1: (^(v1))}
                    var v24 uint16 = 42
                    _ = v23
                    _ = v24
                } else {
                    switch v2 {
                    case "":
                        var v25 int64 = 1
                        _ = v25
                    case "a":
                        fmt.Println(fi16, v5, v3)
                    }
                    v9 = (v9 % ((v9) | 1))
                    v9 += v9
                    switch fi16 {
                    case -202:
                        var v26 uint = v14
                        _ = v26
                    case 32767:
                        var v27 bool = (1.5 > (float32(3.14159) * (-2.25 + float32(0.0))))
                        _ = v27
                    default:
                        v2 = (v2 + "")
                    }
                    v28 := []int{(-(fi16)), -546, 0, fi15}
                    var v29 float64 = 0.0
                    _ = v28
                    _ = v29
                }
            } else {
                v30 := GenS3{F0: v5}
                _ = v30
            }
            _ = fi16
        }
        _ = fi15
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS1{F0: 85, F1: true}
        var v33 uint = v14
        v34 := GenS4{F0: v1, F1: (v5 % ((238) | 1)), F2: 1}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    v14 -= uint(v9)
    v35 := []int{4294967295, 4294967295, 4294967295, 277}
    if false {
        var v36 int8 = 121
        for fi37 := 0; fi37 < 14; fi37++ {
            v38 := GenS1{F0: v5, F1: true}
            if true {
                fmt.Println(v14)
                v9 += v9
                switch v2 {
                case "":
                    for fi39 := 0; fi39 < 12; fi39++ {
                        v40 := GenS1{F0: (v5 ^ v5), F1: false}
                        if ((!false) || ((((3.14159 / 0.1) + 0.1) <= 3.14159) || (false || false))) {
                            if (true || (!((fi37 + (-(fi39))) <= (^(fi37))))) {
                                var v41 int32 = int32(v3)
                                _ = v41
                            }
                            switch fi37 {
                            case 4294967295:
                                switch v2 {
                                case "\n\t":
                                    v14 = v14
                                }
                            case -562:
                                fmt.Println(v36, fi39)
                            default:
                                fmt.Println(v1, v14, v2)
                                var v42 uint64 = (v3 - (^(v3)))
                                switch (fi37 & fi37) {
                                case 1:
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                case 9223372036854775807:
                                    v36 += (v36 / (((-(v36))) | 1))
                                }
                                _ = v42
                            }
                            var v43 uint16 = 1
                            v44 := GenS5{F0: 1, F1: -1, F2: 494}
                            _ = v43
                            _ = v44
                        } else {
                            if true {
                                var v45 uint64 = v3
                                if false {
                                    switch (v2 + v2) {
                                    case "foo bar":
                                        v46 := GenS6{F0: false}
                                        _ = v46
                                    case "":
                                        switch (v2 + (v2 + v2)) {
                                        case "\n\t":
                                            switch v2 {
                                            case "hello":
                                                v45 += v3
                                            default:
                                                var v47 string = "hello"
                                                _ = v47
                                            }
                                            var v48 int = 4294967295
                                            _ = v48
                                        case "hello":
                                            v49 := GenS4{F0: (v1 - -1), F1: (v5 | (v5 >> (((v14 - uint(v9))) & 63))), F2: int8(v3)}
                                            _ = v49
                                        }
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                    case "hello":
                                        v2 = (v2 + v2)
                                    }
                                    var v50 int32 = 42
                                    _ = v50
                                } else {
                                    if (true || false) {
                                        if (!false) {
                                            var v51 uint = 0
                                            v36 -= (v36 + v36)
                                            v52 := GenS7{F0: (v36 << ((v14) & 63))}
                                            if (((v2 + ("" + (v2 + v2))) + v2) >= (v2 + v2)) {
                                                _ = len(v35)
                                            } else {
                                                v3 = (^(v3))
                                                var v53 uint16 = 42
                                                _ = v53
                                            }
                                            _ = v51
                                            _ = v52
                                        } else {
                                            v54 := GenS1{F0: v5, F1: (uint64(v5) == (^(v3)))}
                                            var v55 int32 = -1
                                            if (((((v2 + (v2 + (v2 + ((v2 + v2) + v2)))) + "a") < v2) || true) || (1 < 781)) {
                                                if false {
                                                    if false {
                                                        var v56 uint16 = 255
                                                        var v57 bool = (float32(-2.25) != ((-2.25 * -2.25) / 0.1))
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                } else {
                                                    if false {
                                                        v58 := GenS1{F0: v5, F1: true}
                                                        var v59 uint16 = 65535
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        v60 := GenS8{F0: 255, F1: ""}
                                                        _ = v60
                                                    }
                                                }
                                                if (0 != 816) {
                                                    fmt.Println(v5, v1)
                                                    _ = len(v35)
                                                } else {
                                                    fmt.Println(v55)
                                                    v1 = -32768
                                                }
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    v61 := GenS9{F0: v5}
                                                    _ = v61
                                                case "\n\t":
                                                    v62 := GenS9{F0: 1}
                                                    _ = v62
                                                case "hello":
                                                    switch fi39 {
                                                    case 32767:
                                                        var v63 float32 = 1.5
                                                        v55 += (v55 % ((int32(v45)) | 1))
                                                        _ = v63
                                                    case 587:
                                                        v36 = (v36 << (((v14 % ((v14) | 1))) & 63))
                                                    }
                                                }
                                                switch v2 {
                                                case "hello":
                                                    switch fi39 {
                                                    case -133:
                                                        delete(v4, v9)
                                                    }
                                                case "a":
                                                    var v64 uint16 = 929
                                                    _ = v64
                                                default:
                                                    v3 = v3
                                                }
                                                if false {
                                                    v36 = 100
                                                    switch -504 {
                                                    case 32767:
                                                        switch v2 {
                                                        case "foo bar":
                                                            var v65 int32 = v55
                                                            _ = v65
                                                        case "\n\t":
                                                            fmt.Println(v36, v9, v1)
                                                        case "":
                                                            if false {
                                                                v1 += v1
                                                                _ = len(v4)
                                                            } else {
                                                                var v66 uint16 = 696
                                                                v67 := v4[v9]
                                                                var v68 int32 = v55
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        default:
                                                            fmt.Println(v1)
                                                        }
                                                    default:
                                                        v69 := GenS6{F0: true}
                                                        _ = v69
                                                    }
                                                    var v70 float32 = -2.25
                                                    v4[(v9 << ((v14) & 63))] = (^(fi37))
                                                    _ = v70
                                                }
                                            }
                                            _ = v54
                                            _ = v55
                                        }
                                        v1 = (v1 / ((v1) | 1))
                                    } else {
                                        var v71 string = ("foo bar" + "\n\t")
                                        v5 = (v5 - 0)
                                        switch int(v3) {
                                        case 1:
                                            if (!false) {
                                                if false {
                                                    var v72 uint64 = v45
                                                    var v73 uint32 = v9
                                                    _ = len(v35)
                                                    v74 := GenS10{F0: -248, F1: v1}
                                                    var v75 uint8 = v5
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                                v3 = 0
                                                v5 += 128
                                            }
                                        case 673:
                                            switch fi37 {
                                            case 1:
                                                _ = len(v4)
                                                var v76 int8 = (v36 % (((v36 | (v36 & (v36 / (((v36 * int8(fi37))) | 1))))) | 1))
                                                _ = v76
                                            case -1:
                                                var v77 uint = 1
                                                _ = v77
                                            default:
                                                var v78 int64 = 411
                                                _ = v78
                                            }
                                        case -2147483648:
                                            v71 = "hello"
                                        default:
                                            v36 = int8(v9)
                                        }
                                        _ = v71
                                    }
                                    v79 := GenS1{F0: 247, F1: true}
                                    v80 := GenS11{F0: (v45 ^ v45), F1: (v1 | (v1 - (v1 * (v1 % ((v1) | 1))))), F2: 0}
                                    _ = v79
                                    _ = v80
                                }
                                v14 += 2
                                v2 = v2
                                _ = v45
                            }
                            var v81 int16 = (v1 | 32767)
                            fmt.Println(v5, v14)
                            switch (v2 + "\n\t") {
                            case "":
                                if true {
                                    if true {
                                        var v82 int16 = (v81 & (v81 % (((v1 & 0)) | 1)))
                                        var v83 float64 = 3.14159
                                        var v84 int8 = (v36 / ((v36) | 1))
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                    var v85 uint8 = v5
                                    v86 := v4[(v9 + (v9 ^ v9))]
                                    _ = v85
                                    _ = v86
                                } else {
                                    v36 = -128
                                    v87 := GenS12{F0: v36, F1: (v5 ^ (v5 | 0)), F2: 0}
                                    if false {
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                    }
                                    _ = v87
                                }
                            case "a":
                                var v88 int = fi39
                                _ = v88
                            default:
                                if false {
                                    v36 = v36
                                    switch fi37 {
                                    case -2147483648:
                                        delete(v4, (v9 << ((1) & 63)))
                                    case 644:
                                        if (!true) {
                                            var v89 float32 = 3.14159
                                            v14 = (v14 / ((v14) | 1))
                                            v90 := GenS13{F0: 490}
                                            _ = v89
                                            _ = v90
                                        } else {
                                            if (687 != 42) {
                                                v1 = v81
                                                v91 := GenS9{F0: v5}
                                                _ = v91
                                            } else {
                                                if false {
                                                    var v92 int8 = v36
                                                    _ = v92
                                                }
                                                if false {
                                                    v5 += v5
                                                } else {
                                                    if (((v2 + v2) < (v2 + v2)) && (!false)) {
                                                        var v93 bool = true
                                                        switch (v2 + v2) {
                                                        case "a":
                                                            v9 = 65535
                                                        case "hello":
                                                            v1 += (v81 % (((v81 >> ((v14) & 63))) | 1))
                                                        case "":
                                                            if (false || v93) {
                                                                var v94 int16 = (-(v1))
                                                                switch v2 {
                                                                case "foo bar":
                                                                    if v93 {
                                                                        var v95 float32 = -2.25
                                                                        delete(v4, v9)
                                                                        _ = v95
                                                                    }
                                                                    v96 := GenS5{F0: -702, F1: 42, F2: 463}
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        var v97 uint8 = v5
                                                                        _ = v97
                                                                    case "foo bar":
                                                                        if v93 {
                                                                            if ((v94 & v1) == (v94 >> ((v14) & 63))) {
                                                                                var v98 uint16 = 832
                                                                                _ = v98
                                                                            } else {
                                                                                var v99 int64 = 0
                                                                                fmt.Println(v99, fi37)
                                                                                _ = v99
                                                                            }
                                                                            v4[0] = (fi37 & (-(fi39)))
                                                                        }
                                                                    case "":
                                                                        v9 = v9
                                                                    default:
                                                                        fmt.Println(v36, v5, v1)
                                                                    }
                                                                    _ = v96
                                                                case "\n\t":
                                                                    var v100 float64 = 1e10
                                                                    _ = v100
                                                                default:
                                                                    if len(v35) > 0 {
                                                                        v35[0] = int(v9)
                                                                    }
                                                                    var v101 uint8 = 0
                                                                    _ = v101
                                                                }
                                                                _ = v94
                                                            } else {
                                                                v102 := v4[v9]
                                                                _ = v102
                                                            }
                                                        default:
                                                            v103 := GenS12{F0: (v36 >> ((uint(v5)) & 63)), F1: 2, F2: v1}
                                                            v14 -= (v14 >> (((v14 | 1)) & 63))
                                                            v9 += v9
                                                            _ = v103
                                                        }
                                                        if v93 {
                                                            v81 -= v1
                                                            fmt.Println(v5)
                                                        } else {
                                                            if (!(!false)) {
                                                                if ((^(v14)) <= 0) {
                                                                    var v104 int8 = 100
                                                                    _ = v104
                                                                } else {
                                                                    v14 += v14
                                                                    v105 := GenS14{F0: (v5 - v5)}
                                                                    _ = v105
                                                                }
                                                                var v106 int = 4294967295
                                                                _ = v106
                                                            }
                                                            var v107 int = (fi37 * (fi37 | int(v14)))
                                                            if (0.0 < ((0.1 / -2.25) - 0.0)) {
                                                                if ((v2 < ((v2 + v2) + (((("hello" + "") + (v2 + v2)) + v2) + v2))) || (728 != 32767)) {
                                                                    switch v107 {
                                                                    case -2147483648:
                                                                        fmt.Println(v14)
                                                                    }
                                                                    if v93 {
                                                                        var v108 int = (fi37 * (fi37 << (((^(v14))) & 63)))
                                                                        _ = v108
                                                                    }
                                                                    var v109 uint32 = uint32(v5)
                                                                    _ = v109
                                                                }
                                                            }
                                                            _ = v107
                                                        }
                                                        _ = v93
                                                    }
                                                }
                                                v4[uint32(v5)] = fi37
                                            }
                                            _ = len(v4)
                                        }
                                    case 2147483647:
                                        var v110 uint8 = (v5 % ((v5) | 1))
                                        _ = v110
                                    }
                                    switch v2 {
                                    case "foo bar":
                                        if true {
                                            v9 = (v9 << ((450) & 63))
                                            fmt.Println(fi37)
                                            var v111 float64 = -2.25
                                            var v112 uint32 = 2
                                            _ = v111
                                            _ = v112
                                        }
                                    }
                                } else {
                                    v4[uint32(v5)] = fi37
                                    if false {
                                        var v113 float32 = 3.14159
                                        var v114 int32 = 1
                                        _ = v113
                                        _ = v114
                                    }
                                }
                            }
                            if true {
                                if ((v36 / ((int8(v9)) | 1)) == 85) {
                                    delete(v4, v9)
                                } else {
                                    if (true && (0 > -1)) {
                                        v115 := GenS6{F0: false}
                                        switch 995 {
                                        case 127:
                                            if (v14 == 4294967295) {
                                                if true {
                                                    var v116 int = fi37
                                                    var v117 int16 = v1
                                                    v5 -= uint8(v3)
                                                    _ = v116
                                                    _ = v117
                                                }
                                                v36 = 100
                                            } else {
                                                switch (fi39 * (fi39 + fi39)) {
                                                case 42:
                                                    switch (fi37 * (fi37 - fi39)) {
                                                    case 128:
                                                        if (0 >= 9223372036854775807) {
                                                            delete(v4, v9)
                                                            switch 32767 {
                                                            case 329:
                                                                v118 := v4[v9]
                                                                _ = v118
                                                            default:
                                                                fmt.Println(v1, fi37)
                                                            }
                                                            v3 = v3
                                                            fmt.Println(v2, v5)
                                                            v119 := GenS12{F0: -128, F1: v5, F2: v1}
                                                            _ = v119
                                                        } else {
                                                            if (float32(-2.25) == (0.1 - 0.1)) {
                                                                switch v2 {
                                                                case "":
                                                                    var v120 uint16 = 0
                                                                    _ = v120
                                                                case "hello":
                                                                    v14 = 1
                                                                case "foo bar":
                                                                    var v121 int16 = (v81 + (v1 * 42))
                                                                    _ = v121
                                                                }
                                                                v122 := GenS7{F0: v36}
                                                                _ = v122
                                                            }
                                                            var v123 int64 = 361
                                                            if true {
                                                                var v124 bool = true
                                                                fmt.Println(fi39)
                                                                switch "" {
                                                                case "a":
                                                                    v125 := GenS15{F0: (v123 * (v123 >> ((v14) & 63))), F1: (fi37 - int(v123)), F2: (^(fi39))}
                                                                    _ = v125
                                                                case "":
                                                                    if (!v124) {
                                                                        v126 := v4[987]
                                                                        v127 := v4[107]
                                                                        _ = v126
                                                                        _ = v127
                                                                    } else {
                                                                        v128 := v4[v9]
                                                                        fmt.Println(v5)
                                                                        v5 = v5
                                                                        switch (-(fi39)) {
                                                                        case 2147483647:
                                                                            var v129 int32 = 250
                                                                            _ = v129
                                                                        case 426:
                                                                            v9 -= (^(v9))
                                                                        }
                                                                        _ = v128
                                                                    }
                                                                case "hello":
                                                                    switch -9223372036854775808 {
                                                                    case -535:
                                                                        fmt.Println(v36, v36, v9)
                                                                    case 127:
                                                                        var v130 uint64 = (v3 - v3)
                                                                        _ = v130
                                                                    default:
                                                                        v81 = v81
                                                                    }
                                                                default:
                                                                    if false {
                                                                        var v131 int = 1
                                                                        v81 = (v1 / (((^(v1))) | 1))
                                                                        _ = v131
                                                                    } else {
                                                                        var v132 bool = false
                                                                        _ = v132
                                                                    }
                                                                }
                                                                var v133 uint16 = uint16(v9)
                                                                _ = len(v4)
                                                                if v124 {
                                                                    if len(v35) > 0 {
                                                                        v35[0] = fi37
                                                                    }
                                                                    if false {
                                                                        _ = len(v4)
                                                                        var v134 string = "hello"
                                                                        _ = v134
                                                                    } else {
                                                                        if len(v35) >= 2 {
                                                                            v35 = v35[1:len(v35)]
                                                                        }
                                                                        _ = len(v4)
                                                                        if len(v35) > 0 {
                                                                            v35[0] = (fi37 + fi37)
                                                                        }
                                                                    }
                                                                }
                                                                _ = v124
                                                                _ = v133
                                                            } else {
                                                                var v135 uint8 = uint8(v14)
                                                                _ = len(v35)
                                                                v14 -= 4294967295
                                                                _ = v135
                                                            }
                                                            _ = v123
                                                        }
                                                    default:
                                                        if (int16(v5) <= -180) {
                                                            v136 := GenS7{F0: v36}
                                                            v81 = v1
                                                            v137 := GenS16{F0: v3, F1: fi37}
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            if ((v9 ^ v9) != v9) {
                                                v138 := GenS17{F0: (fi37 & -9223372036854775808)}
                                                var v139 int = fi39
                                                switch ("\n\t" + (v2 + v2)) {
                                                case "":
                                                    v140 := GenS18{F0: (v14 * v14), F1: uint32(v5)}
                                                    _ = v140
                                                case "a":
                                                    if false {
                                                        delete(v4, 688)
                                                    } else {
                                                        v14 -= uint(v5)
                                                        fmt.Println(v5, v9)
                                                        v14 = (v14 % (((v14 >> (((v14 / ((v14) | 1))) & 63))) | 1))
                                                    }
                                                case "hello":
                                                    v141 := GenS19{F0: -9223372036854775808}
                                                    _ = v141
                                                }
                                                _ = v138
                                                _ = v139
                                            }
                                        }
                                        v3 -= v3
                                        _ = v115
                                    }
                                }
                                v3 = 2
                            } else {
                                var v142 float32 = (-2.25 + 3.14159)
                                switch ("hello" + ("a" + (v2 + "foo bar"))) {
                                case "":
                                    v143 := GenS15{F0: 9223372036854775807, F1: fi39, F2: (fi37 | fi37)}
                                    _ = v143
                                case "hello":
                                    if false {
                                        var v144 float64 = float64(v142)
                                        _ = v144
                                    }
                                }
                                if (((v2 + v2) + v2) >= (v2 + v2)) {
                                    v145 := GenS20{F0: v36}
                                    v146 := GenS11{F0: uint64(v5), F1: (v81 | 266), F2: v9}
                                    _ = v145
                                    _ = v146
                                } else {
                                    fmt.Println(v142, v14)
                                    switch fi39 {
                                    case 32767:
                                        switch ((v2 + v2) + "a") {
                                        case "a":
                                            var v147 int32 = 1
                                            _ = v147
                                        }
                                    case 2147483647:
                                        v148 := GenS21{F0: 1}
                                        _ = v148
                                    }
                                    var v149 string = (("foo bar" + "") + (v2 + "foo bar"))
                                    _ = v149
                                }
                                var v150 uint32 = (v9 << ((377) & 63))
                                _ = v142
                                _ = v150
                            }
                            var v151 int64 = 1
                            v35 = append(v35, 32768)
                            _ = v151
                            _ = v81
                        }
                        if (!(!true)) {
                            if (true || false) {
                                v35 = append(v35, fi39)
                            }
                            if (false && false) {
                                v5 = v5
                            } else {
                                if false {
                                    v5 -= 1
                                    var v152 int16 = -32768
                                    switch -552 {
                                    case 32768:
                                        v153 := GenS21{F0: 9223372036854775807}
                                        _ = v153
                                    case -9223372036854775808:
                                        delete(v4, v9)
                                    }
                                    var v154 uint16 = 42
                                    v152 = (^(v1))
                                    v155 := GenS8{F0: (v5 ^ 2), F1: (v2 + v2)}
                                    _ = v152
                                    _ = v154
                                    _ = v155
                                } else {
                                    v156 := GenS15{F0: -9223372036854775808, F1: fi39, F2: fi37}
                                    switch (fi39 + (fi37 | (-(fi39)))) {
                                    case 160:
                                        var v157 string = v2
                                        _ = v157
                                    case 938:
                                        fmt.Println(v3, fi39)
                                        switch (fi39 >> ((v14) & 63)) {
                                        case 4294967295:
                                            v3 = uint64(v9)
                                            v35 = append(v35, 301)
                                        case 17:
                                            if len(v35) > 0 {
                                                v35[0] = 0
                                            }
                                        case 0:
                                            switch v2 {
                                            case "":
                                                switch -128 {
                                                case -1:
                                                    v14 -= uint(v3)
                                                case 9223372036854775807:
                                                    v158 := GenS22{F0: (v5 * 2), F1: v14}
                                                    _ = v158
                                                case 32768:
                                                    v1 = (v1 % (((^(v1))) | 1))
                                                }
                                            }
                                        default:
                                            switch fi39 {
                                            case -2147483648:
                                                _ = len(v35)
                                            case 128:
                                                v159 := GenS23{F0: 2147483648}
                                                _ = v159
                                            case 421:
                                                switch ((((((v2 + (("hello" + v2) + "hello")) + "hello") + (v2 + v2)) + ("\n\t" + v2)) + "\n\t") + v2) {
                                                case "hello":
                                                    var v160 int16 = v1
                                                    _ = v160
                                                default:
                                                    switch 0 {
                                                    case 141:
                                                        v9 = v9
                                                    case 1:
                                                        var v161 int = (fi37 % ((fi39) | 1))
                                                        _ = v161
                                                    default:
                                                        if (!true) {
                                                            switch fi37 {
                                                            case 32767:
                                                                var v162 uint = v14
                                                                _ = v162
                                                            case 128:
                                                                _ = len(v35)
                                                            case 861:
                                                                fmt.Println(fi39, v1)
                                                            default:
                                                                v14 = 65535
                                                            }
                                                            var v163 uint64 = v3
                                                            switch (fi39 * fi37) {
                                                            case 1:
                                                                v4[0] = (-(fi39))
                                                            }
                                                            _ = v163
                                                        } else {
                                                            var v164 int8 = (-(v36))
                                                            _ = v164
                                                        }
                                                    }
                                                    var v165 int32 = int32(fi39)
                                                    if (-9223372036854775808 >= (fi39 / ((fi39) | 1))) {
                                                        v166 := GenS24{F0: 905, F1: -9223372036854775808}
                                                        if false {
                                                            if false {
                                                                if ((float32(0.0) - -2.25) >= 0.1) {
                                                                    if len(v35) >= 2 {
                                                                        v35 = v35[1:len(v35)]
                                                                    }
                                                                } else {
                                                                    var v167 string = (v2 + (v2 + "\n\t"))
                                                                    _ = v167
                                                                }
                                                            } else {
                                                                v36 = (v36 | (v36 / (((v36 ^ (v36 & v36))) | 1)))
                                                                switch (fi39 >> (((v14 & v14)) & 63)) {
                                                                case 128:
                                                                    var v168 uint32 = v9
                                                                    _ = v168
                                                                case -9223372036854775808:
                                                                    fmt.Println(v5, fi39)
                                                                case -2147483648:
                                                                    if ((v5 ^ (v5 << ((965) & 63))) <= v5) {
                                                                        if ((("" + "hello") + "hello") < v2) {
                                                                            var v169 int32 = (v165 | int32(v1))
                                                                            _ = v169
                                                                        } else {
                                                                            var v170 string = ((v2 + "\n\t") + v2)
                                                                            _ = v170
                                                                        }
                                                                    } else {
                                                                        var v171 uint = uint(v9)
                                                                        _ = v171
                                                                    }
                                                                }
                                                                fmt.Println(fi39)
                                                            }
                                                            v9 += v9
                                                            if false {
                                                                v2 = (v2 + v2)
                                                                if ((!((v14 * v14) <= uint(v9))) || (0 < int64(v9))) {
                                                                    fmt.Println(v14, v36)
                                                                    v172 := v4[(v9 * v9)]
                                                                    switch int(v9) {
                                                                    case 982:
                                                                        v173 := GenS7{F0: v36}
                                                                        _ = v173
                                                                    default:
                                                                        switch fi37 {
                                                                        case 32767:
                                                                            switch (-(fi37)) {
                                                                            case -2147483648:
                                                                                if len(v35) > 0 {
                                                                                    v35[0] = (fi37 / ((2147483647) | 1))
                                                                                }
                                                                            case 128:
                                                                                v174 := GenS2{F0: (v5 * (^(v5))), F1: (v1 | (v1 * (v1 ^ v1)))}
                                                                                switch (v2 + v2) {
                                                                                case "hello":
                                                                                    if (!false) {
                                                                                        v4[(v9 | v9)] = 32768
                                                                                    } else {
                                                                                        switch -2147483648 {
                                                                                        case -9223372036854775808:
                                                                                            v175 := GenS8{F0: (v5 >> ((v14) & 63)), F1: v2}
                                                                                            _ = v175
                                                                                        case 879:
                                                                                            var v176 uint = v14
                                                                                            _ = v176
                                                                                        default:
                                                                                            var v177 uint = 2
                                                                                            _ = v177
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v178 int64 = 902
                                                                                _ = v174
                                                                                _ = v178
                                                                            case -687:
                                                                                var v179 int32 = (v165 ^ (v165 | v165))
                                                                                if false {
                                                                                    var v180 string = (v2 + v2)
                                                                                    fmt.Println(v3)
                                                                                    _ = v180
                                                                                } else {
                                                                                    if (!true) {
                                                                                        if true {
                                                                                            var v181 int64 = 644
                                                                                            var v182 float64 = 1.5
                                                                                            _ = v181
                                                                                            _ = v182
                                                                                        } else {
                                                                                            var v183 bool = ((v165 << (((v14 / ((18446744073709551615) | 1))) & 63)) <= v165)
                                                                                            _ = len(v35)
                                                                                            _ = v183
                                                                                        }
                                                                                        fmt.Println(v36)
                                                                                        var v184 bool = false
                                                                                        _ = v184
                                                                                    }
                                                                                    var v185 string = "\n\t"
                                                                                    switch (v185 + v2) {
                                                                                    case "foo bar":
                                                                                        switch (fi37 & int(v1)) {
                                                                                        case -1:
                                                                                            v186 := GenS13{F0: 2}
                                                                                            _ = v186
                                                                                        }
                                                                                    case "a":
                                                                                        if ((v14 / ((uint(v3)) | 1)) != 320) {
                                                                                            fmt.Println(fi37, v5)
                                                                                            delete(v4, v9)
                                                                                            v5 = 42
                                                                                        }
                                                                                    default:
                                                                                        var v187 uint8 = v5
                                                                                        _ = v187
                                                                                    }
                                                                                    _ = v185
                                                                                }
                                                                                var v188 int32 = -1
                                                                                switch int(v14) {
                                                                                case 127:
                                                                                    v14 += (v14 | (v14 - (v14 % ((v14) | 1))))
                                                                                case -1:
                                                                                    v189 := GenS23{F0: v9}
                                                                                    _ = v189
                                                                                case 525:
                                                                                    v1 = (v1 | (v1 / ((1) | 1)))
                                                                                default:
                                                                                    v190 := GenS25{F0: v1}
                                                                                    _ = v190
                                                                                }
                                                                                _ = v179
                                                                                _ = v188
                                                                            default:
                                                                                fmt.Println(v5, v36)
                                                                            }
                                                                        case 32768:
                                                                            v165 -= v165
                                                                        default:
                                                                            var v191 float64 = (-2.25 / float64(3.14159))
                                                                            _ = v191
                                                                        }
                                                                    }
                                                                    _ = v172
                                                                }
                                                                var v192 uint64 = 4294967295
                                                                _ = v192
                                                            } else {
                                                                v193 := GenS23{F0: (v9 % (((v9 + (v9 >> (((v14 / ((v14) | 1))) & 63)))) | 1))}
                                                                _ = v193
                                                            }
                                                        }
                                                        var v194 uint = v14
                                                        _ = v166
                                                        _ = v194
                                                    } else {
                                                        fmt.Println(v1, v1)
                                                        v36 -= int8(fi37)
                                                    }
                                                    if true {
                                                        fmt.Println(v36)
                                                        v165 -= v165
                                                        v195 := GenS13{F0: 32767}
                                                        _ = v195
                                                    } else {
                                                        v196 := GenS15{F0: 42, F1: fi37, F2: (fi39 | fi39)}
                                                        if true {
                                                            v197 := GenS26{F0: v165, F1: v5, F2: -489}
                                                            var v198 bool = ("foo bar" != v2)
                                                            fmt.Println(v1)
                                                            v165 = v165
                                                            _ = v197
                                                            _ = v198
                                                        } else {
                                                            v35 = append(v35, int(v3))
                                                            v1 += (-(v1))
                                                            _ = len(v4)
                                                        }
                                                        var v199 int32 = v165
                                                        _ = v196
                                                        _ = v199
                                                    }
                                                    v14 = 741
                                                    _ = v165
                                                }
                                            }
                                        }
                                    }
                                    _ = v156
                                }
                                if true {
                                    _ = len(v35)
                                    var v200 float32 = 1.5
                                    _ = v200
                                } else {
                                    v4[v9] = fi39
                                    var v201 int16 = -785
                                    _ = v201
                                }
                            }
                        } else {
                            if true {
                                _ = len(v35)
                            }
                            switch v2 {
                            case "foo bar":
                                fmt.Println(v14, v14)
                                v202 := GenS27{F0: 0, F1: -2147483648, F2: 18446744073709551615}
                                if (1 >= 255) {
                                    v203 := GenS12{F0: (v36 ^ v36), F1: (v5 & (v5 * 42)), F2: (v1 * (v1 >> (((v14 & uint(v3))) & 63)))}
                                    fmt.Println(v9, fi37)
                                    _ = len(v4)
                                    _ = v203
                                }
                                fmt.Println(v2)
                                if len(v35) > 0 {
                                    v35[0] = (fi39 % ((-935) | 1))
                                }
                                _ = v202
                            }
                            v5 -= (v5 % (((v5 / ((v5) | 1))) | 1))
                            v9 = v9
                            var v204 int = fi37
                            v3 -= v3
                            v4[(v9 << ((v14) & 63))] = (fi37 & int(v3))
                            _ = v204
                        }
                        _ = fi39
                        _ = v40
                    }
                case "foo bar":
                    var v205 int16 = (v1 + 1000)
                    _ = v205
                case "a":
                    if true {
                        if (((1.5 - (1e10 + 3.14159)) + (1.5 + 1e10)) > 1e10) {
                            fmt.Println(v3, v3, fi37)
                            delete(v4, 481)
                        }
                    } else {
                        v206 := GenS25{F0: (v1 + (^(v1)))}
                        delete(v4, v9)
                        _ = v206
                    }
                }
            } else {
                switch (fi37 * 580) {
                case 127:
                    v207 := GenS28{F0: uint(v5), F1: (^(v3))}
                    _ = v207
                case -2147483648:
                    v208 := GenS26{F0: 2147483647, F1: (v5 & uint8(v3)), F2: -9223372036854775808}
                    _ = v208
                default:
                    v3 += 2
                }
                var v209 int16 = v1
                v210 := GenS29{F0: 1000, F1: v3}
                _ = v209
                _ = v210
            }
            _ = fi37
            _ = v38
        }
        if (-2.25 >= (-2.25 * 1e10)) {
            v4[v9] = 2147483647
            if true {
                var v211 uint32 = 2147483647
                v36 = (v36 << ((v14) & 63))
                switch ("a" + "\n\t") {
                case "a":
                    v9 -= (v211 & v9)
                case "hello":
                    v212 := GenS13{F0: 2}
                    _ = v212
                case "foo bar":
                    var v213 int64 = 234
                    _ = v213
                default:
                    v14 = 2
                }
                _ = v211
            }
        } else {
            if true {
                if false {
                    _ = len(v4)
                }
                var v214 int64 = 52
                var v215 int16 = v1
                _ = v214
                _ = v215
            } else {
                switch ("foo bar" + ((("foo bar" + v2) + v2) + ((v2 + (("\n\t" + (((v2 + (("" + ("" + "foo bar")) + v2)) + (("a" + ("\n\t" + v2)) + v2)) + v2)) + (v2 + v2))) + "hello"))) {
                case "hello":
                    v216 := GenS30{F0: uint(v5)}
                    _ = v216
                case "a":
                    v5 = 2
                default:
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                }
            }
            switch 384 {
            case 4294967295:
                v217 := GenS21{F0: 9223372036854775807}
                _ = v217
            case 9223372036854775807:
                for fi218 := 0; fi218 < 19; fi218++ {
                    v219 := GenS16{F0: (^(v3)), F1: fi218}
                    var v220 string = ("a" + "hello")
                    _ = fi218
                    _ = v219
                    _ = v220
                }
            default:
                v221 := GenS31{F0: ("" + ((v2 + "\n\t") + "hello")), F1: 127}
                _ = v221
            }
            v9 -= (v9 - v9)
        }
        _ = v36
    }
    for fi222 := 0; fi222 < 10; fi222++ {
        for fi223 := 0; fi223 < 20; fi223++ {
            if (42 == 512) {
                var v224 uint32 = (v9 - (v9 - (v9 / (((v9 / (((^(v9))) | 1))) | 1))))
                switch v2 {
                case "hello":
                    delete(v4, 206)
                }
                var v225 uint = (v14 * (v14 * (v14 | (v14 >> ((v14) & 63)))))
                v5 = uint8(v225)
                _ = v224
                _ = v225
            } else {
                v226 := v4[v9]
                var v227 int8 = 0
                _ = v226
                _ = v227
            }
            _ = fi223
        }
        _ = fi222
    }
    v228 := GenS3{F0: (v5 / ((127) | 1))}
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    v9 = 796
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v228
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
