package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint16; F1 string; F2 uint16 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint32; F1 uint32; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int64; F1 int8; F2 uint }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 bool; F1 int32 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint8 }

func main() {
    var v1 int64 = 1
    var v2 uint = 983
    var v3 int16 = 1
    var v4 uint = v2
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 2
    var v7 string = (("a" + ("" + "hello")) + "a")
    var v8 uint16 = 32767
    v2 += (v4 & v2)
    v8 -= uint16(v6)
    if false {
        if false {
            v9 := GenS1{F0: 42, F1: 2147483647}
            v3 = v3
            _ = v9
        } else {
            if true {
                if ((v8 != (v8 - (v8 * (^(v8))))) && false) {
                    if (570 != 0) {
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v1 = 0
                            var v11 float32 = (float32(0.1) * (float32((1e10 - 3.14159)) - 1.5))
                            for fi12 := 0; fi12 < 20; fi12++ {
                                if ((fi10 >> (((^(v4))) & 63)) > 2147483647) {
                                    v13 := GenS2{F0: v3}
                                    fmt.Println(v8, v3)
                                    _ = v13
                                } else {
                                    v2 -= 255
                                    var v14 uint32 = 0
                                    var v15 uint16 = uint16(v4)
                                    v1 = v1
                                    v15 = v15
                                    switch v7 {
                                    case "a":
                                        if true {
                                            fmt.Println(v14, v8)
                                            v16 := []int{fi10}
                                            var v17 int32 = 203
                                            _ = v16
                                            _ = v17
                                        }
                                    case "":
                                        if true {
                                            v18 := GenS1{F0: 32767, F1: 42}
                                            v11 -= v11
                                            v15 = 32767
                                            switch v7 {
                                            case "a":
                                                v19 := GenS3{F0: v14, F1: 0}
                                                _ = v19
                                            }
                                            _ = v18
                                        } else {
                                            switch (fi10 + fi10) {
                                            case 504:
                                                switch (-(fi12)) {
                                                case 128:
                                                    if (-1 <= 127) {
                                                        var v20 uint16 = v8
                                                        v21 := []int{fi12}
                                                        v22 := GenS4{F0: int32(v6), F1: (v3 + (v3 & (v3 + v3)))}
                                                        _ = v20
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                case -2147483648:
                                                    delete(v5, v14)
                                                case 0:
                                                    if false {
                                                        v23 := GenS5{F0: v1, F1: -2147483648}
                                                        if (((v1 | v1) == 0) && false) {
                                                            switch "a" {
                                                            case "hello":
                                                                switch (("" + v7) + v7) {
                                                                case "a":
                                                                    if (42 < 255) {
                                                                        if (false && (0.0 >= (1.5 / -2.25))) {
                                                                            fmt.Println(v15)
                                                                        } else {
                                                                            v7 = "foo bar"
                                                                        }
                                                                    }
                                                                }
                                                            case "":
                                                                fmt.Println(v8, v11, fi12)
                                                            default:
                                                                if (v1 > (v1 ^ v1)) {
                                                                    v24 := v5[2147483647]
                                                                    var v25 float64 = 1.5
                                                                    switch fi12 {
                                                                    case -1:
                                                                        if (3.14159 > 3.14159) {
                                                                            if (false || (42 < 2)) {
                                                                                if (v1 != v1) {
                                                                                    switch v7 {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v1, v3, fi10)
                                                                                    case "hello":
                                                                                        if false {
                                                                                            v5[v24] = 2
                                                                                            if false {
                                                                                                v14 -= uint32(v2)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v26 float64 = -2.25
                                                                                        _ = v26
                                                                                    }
                                                                                    var v27 uint64 = (v6 | (v6 >> ((1) & 63)))
                                                                                    _ = v27
                                                                                }
                                                                                var v28 int16 = (v3 ^ (v3 + v3))
                                                                                _ = v28
                                                                            }
                                                                            switch fi10 {
                                                                            case 127:
                                                                                var v29 uint16 = (v8 * (^(v15)))
                                                                                _ = v29
                                                                            case 32767:
                                                                                fmt.Println(v24, v1, v3)
                                                                                switch (v7 + ("hello" + "a")) {
                                                                                case "":
                                                                                    if false {
                                                                                        v25 += (v25 + 0.1)
                                                                                    } else {
                                                                                        var v30 int64 = 9223372036854775807
                                                                                        if true {
                                                                                            v30 = (v30 | -839)
                                                                                        } else {
                                                                                            fmt.Println(v1)
                                                                                            v7 = (("\n\t" + v7) + ("foo bar" + (("" + v7) + "foo bar")))
                                                                                            var v31 uint = (v2 >> (((v4 - 255)) & 63))
                                                                                            _ = v31
                                                                                        }
                                                                                        var v32 float32 = v11
                                                                                        v33 := []int{int(v1), (fi10 + fi12), -2147483648, (fi10 % (((-(fi12))) | 1))}
                                                                                        _ = v30
                                                                                        _ = v32
                                                                                        _ = v33
                                                                                    }
                                                                                case "a":
                                                                                    v8 = v15
                                                                                case "hello":
                                                                                    v2 = (v4 * 4294967295)
                                                                                }
                                                                            case 32768:
                                                                                var v34 int16 = (v3 * (v3 ^ (v3 << ((65535) & 63))))
                                                                                _ = v34
                                                                            default:
                                                                                if ((v3 & v3) > (v3 * 42)) {
                                                                                    fmt.Println(v1)
                                                                                    v35 := []int{(fi12 * 32767), (fi12 % (((^(fi12))) | 1))}
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if (v4 >= uint(v24)) {
                                                                                v24 = (v24 << ((2) & 63))
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v36 float64 = v25
                                                                        _ = v36
                                                                    }
                                                                    _ = v24
                                                                    _ = v25
                                                                } else {
                                                                    v6 -= v6
                                                                    var v37 uint16 = v8
                                                                    v38 := GenS2{F0: -32768}
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                            }
                                                        }
                                                        var v39 bool = (v6 <= 2)
                                                        _ = v23
                                                        _ = v39
                                                    } else {
                                                        if (!false) {
                                                            v5[(^(v14))] = 65535
                                                            var v40 int8 = 42
                                                            v11 += v11
                                                            _ = v40
                                                        }
                                                        var v41 int16 = v3
                                                        _ = v41
                                                    }
                                                default:
                                                    v42 := GenS1{F0: 42, F1: 2147483647}
                                                    _ = v42
                                                }
                                            default:
                                                var v43 uint16 = 1
                                                _ = v43
                                            }
                                            var v44 uint = 2
                                            var v45 float32 = (v11 / 1e10)
                                            _ = v44
                                            _ = v45
                                        }
                                        v5[(v14 & v14)] = v14
                                        var v46 uint8 = 149
                                        _ = v46
                                    case "\n\t":
                                        v47 := v5[1]
                                        _ = v47
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                v5[13] = uint32(v2)
                                v2 = (^(v4))
                                _ = fi12
                            }
                            _ = fi10
                            _ = v11
                        }
                    }
                    v48 := GenS4{F0: int32(v3), F1: 1000}
                    for fi49 := 0; fi49 < 14; fi49++ {
                        fmt.Println(v2, v8)
                        _ = len(v5)
                        for fi50 := 0; fi50 < 3; fi50++ {
                            if ((v1 & v1) == v1) {
                                switch v7 {
                                case "a":
                                    var v51 uint64 = (^(v6))
                                    _ = v51
                                default:
                                    var v52 int32 = 1
                                    _ = v52
                                }
                                if true {
                                    if true {
                                        if true {
                                            var v53 uint32 = 2147483647
                                            v6 = (^(v6))
                                            _ = v53
                                        }
                                        fmt.Println(v4, v8, v6)
                                        if (1e10 != (float64(((-2.25 - (0.0 - (3.14159 * float32(0.1)))) * (float32(((1.5 * 1.5) * ((1.5 / 1e10) / 3.14159))) + (-2.25 / 1e10)))) / 1.0)) {
                                            if (0 <= v4) {
                                                var v54 int32 = 614
                                                v55 := []int{-189, (fi50 % ((fi50) | 1))}
                                                v6 = (v6 % (((^(v6))) | 1))
                                                _ = v54
                                                _ = v55
                                            }
                                            v1 = -1
                                            var v56 uint32 = 42
                                            _ = v56
                                        } else {
                                            var v57 int = (fi49 >> ((v2) & 63))
                                            switch fi50 {
                                            case -2147483648:
                                                fmt.Println(v1, fi50)
                                                if false {
                                                    switch ((((v7 + v7) + v7) + v7) + v7) {
                                                    case "\n\t":
                                                        v6 += v6
                                                    case "foo bar":
                                                        if true {
                                                            var v58 int32 = 42
                                                            delete(v5, 360)
                                                            if (int32(v3) != v58) {
                                                                switch 32767 {
                                                                case 523:
                                                                    v5[2147483648] = 2
                                                                case -599:
                                                                    v4 = 158
                                                                default:
                                                                    fmt.Println(v8)
                                                                    v59 := GenS3{F0: 493, F1: 672}
                                                                    v4 -= (v4 % (((^(v4))) | 1))
                                                                    _ = v59
                                                                }
                                                                var v60 string = v7
                                                                if (v6 >= 2) {
                                                                    if ((v3 | (v3 % ((v3) | 1))) == (^(v3))) {
                                                                        v61 := []int{(-(v57))}
                                                                        _ = v61
                                                                    }
                                                                } else {
                                                                    if (!true) {
                                                                        _ = len(v5)
                                                                    } else {
                                                                        if (1 >= -110) {
                                                                            if (v6 > (v6 / (((v6 / (((v6 % ((v6) | 1))) | 1))) | 1))) {
                                                                                var v62 float64 = 3.14159
                                                                                v5[0] = 301
                                                                                var v63 float64 = v62
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                        } else {
                                                                            v1 -= v1
                                                                            fmt.Println(v1, v60)
                                                                            v58 -= -1
                                                                            var v64 uint = (v4 >> ((v2) & 63))
                                                                            _ = v64
                                                                        }
                                                                        var v65 int64 = v1
                                                                        fmt.Println(v8, v65)
                                                                        _ = v65
                                                                    }
                                                                    v66 := GenS6{F0: (v8 / ((uint16(v6)) | 1)), F1: "\n\t", F2: 685}
                                                                    _ = v66
                                                                }
                                                                switch fi50 {
                                                                case 741:
                                                                    if false {
                                                                        switch fi50 {
                                                                        case 32768:
                                                                            fmt.Println(v4)
                                                                        case -207:
                                                                            v4 -= (v2 & (v4 % (((^(v4))) | 1)))
                                                                            fmt.Println(v6)
                                                                            v6 = (^(v6))
                                                                        default:
                                                                            v58 = (v58 / ((v58) | 1))
                                                                        }
                                                                        if false {
                                                                            v60 = "foo bar"
                                                                            v67 := GenS2{F0: v3}
                                                                            _ = v67
                                                                        }
                                                                        v58 += v58
                                                                        switch fi49 {
                                                                        case 32768:
                                                                            v3 = (v3 & (v3 % (((v3 - 1000)) | 1)))
                                                                        default:
                                                                            var v68 int32 = 841
                                                                            _ = v68
                                                                        }
                                                                        v2 -= v2
                                                                    } else {
                                                                        v69 := GenS5{F0: v1, F1: (^(v58))}
                                                                        _ = v69
                                                                    }
                                                                default:
                                                                    v70 := []int{v57, v57, (fi50 * 708), 2147483647}
                                                                    _ = v70
                                                                }
                                                                _ = v60
                                                            }
                                                            _ = v58
                                                        }
                                                    case "hello":
                                                        switch 343 {
                                                        case 128:
                                                            v71 := GenS7{F0: (v6 | v6), F1: v7}
                                                            _ = v71
                                                        default:
                                                            v72 := GenS8{F0: 513, F1: 0, F2: "\n\t"}
                                                            _ = v72
                                                        }
                                                        v4 -= v4
                                                    }
                                                }
                                            }
                                            _ = v57
                                        }
                                        switch v7 {
                                        case "\n\t":
                                            switch fi49 {
                                            case 128:
                                                fmt.Println(v4)
                                            case 127:
                                                v4 = 42
                                            }
                                        case "":
                                            switch fi49 {
                                            case 2147483647:
                                                v73 := []int{int(v6), (fi49 ^ 32767), (fi50 % (((-(fi49))) | 1)), fi50}
                                                _ = v73
                                            }
                                        case "hello":
                                            if ((fi49 & fi50) <= (fi49 * (-(fi49)))) {
                                                var v74 uint8 = 133
                                                var v75 int8 = -1
                                                v75 += v75
                                                _ = v74
                                                _ = v75
                                            } else {
                                                v4 = (^(v2))
                                                v3 -= v3
                                                v6 -= (^(v6))
                                                fmt.Println(v2, v6, v7)
                                            }
                                        default:
                                            v2 = uint(v6)
                                        }
                                    } else {
                                        if ((v1 & v1) != -9223372036854775808) {
                                            v76 := GenS9{F0: 1}
                                            fmt.Println(v1, v4)
                                            if true {
                                                v77 := GenS1{F0: v3, F1: 42}
                                                v78 := []int{fi50}
                                                var v79 uint = uint(v6)
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                            } else {
                                                _ = len(v5)
                                                fmt.Println(v7, v4, v3)
                                                v8 = v8
                                            }
                                            _ = v76
                                        } else {
                                            switch v7 {
                                            case "hello":
                                                v8 = (v8 >> (((^(v2))) & 63))
                                            case "":
                                                v2 += v4
                                            case "\n\t":
                                                v80 := GenS9{F0: v1}
                                                _ = v80
                                            }
                                            fmt.Println(v2)
                                            fmt.Println(v4, fi49, v2)
                                        }
                                    }
                                } else {
                                    v5[16] = 2147483648
                                    v81 := GenS10{F0: v7, F1: 127, F2: (^(v6))}
                                    var v82 float64 = -2.25
                                    _ = v81
                                    _ = v82
                                }
                                var v83 float64 = (0.1 - (1e10 / (float64(0.1) + -2.25)))
                                _ = v83
                            }
                            fmt.Println(v4)
                            _ = fi50
                        }
                        _ = fi49
                    }
                    _ = v48
                } else {
                    fmt.Println(v2)
                    v84 := GenS11{F0: 1}
                    _ = v84
                }
                for fi85 := 0; fi85 < 17; fi85++ {
                    if false {
                        for fi86 := 0; fi86 < 14; fi86++ {
                            if (!(-1 == 42)) {
                                switch "\n\t" {
                                case "hello":
                                    var v87 int32 = 2147483647
                                    _ = v87
                                case "\n\t":
                                    if true {
                                        var v88 int = (fi86 ^ (fi86 % (((fi85 * fi85)) | 1)))
                                        _ = v88
                                    }
                                default:
                                    v89 := GenS12{F0: (!false)}
                                    _ = v89
                                }
                                v7 = (v7 + (((v7 + ((("a" + v7) + (v7 + (v7 + (v7 + (v7 + (("a" + ((v7 + (v7 + "")) + "a")) + (v7 + v7))))))) + v7)) + (("" + "foo bar") + "foo bar")) + ((v7 + v7) + ("\n\t" + ((v7 + "\n\t") + v7)))))
                            }
                            var v90 string = v7
                            switch 1 {
                            case 582:
                                v91 := GenS9{F0: -1}
                                _ = v91
                            case 32767:
                                var v92 int64 = v1
                                _ = v92
                            default:
                                v1 += v1
                            }
                            _ = fi86
                            _ = v90
                        }
                        var v93 float32 = -2.25
                        _ = v93
                    }
                    _ = fi85
                }
            }
            fmt.Println(v8, v3, v8)
        }
        for fi94 := 0; fi94 < 15; fi94++ {
            v95 := GenS1{F0: v3, F1: -1}
            _ = len(v5)
            _ = fi94
            _ = v95
        }
        v96 := GenS8{F0: 2147483648, F1: 65535, F2: v7}
        _ = v96
    } else {
        v97 := GenS4{F0: 525, F1: v3}
        var v98 int64 = (v1 << (((v4 - 1)) & 63))
        _ = v97
        _ = v98
    }
    v8 += uint16(v4)
    for fi99 := 0; fi99 < 6; fi99++ {
        for fi100 := 0; fi100 < 3; fi100++ {
            var v101 uint16 = v8
            var v102 int8 = -128
            fmt.Println(v4, v101)
            _ = fi100
            _ = v101
            _ = v102
        }
        v6 = (v6 + (^(v6)))
        _ = fi99
    }
    var v103 int = -140
    v103 = v103
    for fi104 := 0; fi104 < 16; fi104++ {
        if (1.5 > 3.14159) {
            for fi105 := 0; fi105 < 12; fi105++ {
                v106 := GenS13{F0: 820, F1: v6}
                v107 := GenS13{F0: v2, F1: v6}
                _ = fi105
                _ = v106
                _ = v107
            }
            var v108 int8 = 127
            fmt.Println(v108, v3)
            _ = v108
        }
        var v109 int8 = 56
        for fi110 := 0; fi110 < 4; fi110++ {
            if (!(!true)) {
                var v111 float32 = ((1.5 - float32(0.1)) * 1e10)
                v112 := GenS14{F0: ""}
                _ = v111
                _ = v112
            }
            v113 := GenS5{F0: 9223372036854775807, F1: -596}
            var v114 uint = (^(v2))
            if false {
                v115 := GenS13{F0: v114, F1: (v6 & (v6 - (v6 << ((v2) & 63))))}
                v1 = (v1 * v1)
                v116 := []int{}
                delete(v5, 2147483647)
                _ = v115
                _ = v116
            }
            v1 += v1
            if (v114 < (v4 & v4)) {
                _ = len(v5)
                v117 := GenS2{F0: v3}
                v118 := GenS11{F0: 231}
                _ = v117
                _ = v118
            } else {
                switch "\n\t" {
                case "foo bar":
                    switch fi110 {
                    case -494:
                        v119 := GenS5{F0: -1, F1: 0}
                        var v120 uint8 = uint8(v2)
                        var v121 float64 = 1e10
                        _ = v119
                        _ = v120
                        _ = v121
                    case -1:
                        v122 := []int{2147483647, (fi104 * fi110), (v103 & -649), (^(fi110))}
                        _ = v122
                    case 2147483647:
                        v2 = (v4 * (v114 & v2))
                    default:
                        v7 = v7
                    }
                case "":
                    v123 := []int{(v103 ^ int(v1)), fi110}
                    _ = v123
                case "hello":
                    v8 = (v8 ^ v8)
                default:
                    fmt.Println(v109, fi110)
                    var v124 float64 = 1e10
                    if true {
                        var v125 int32 = -1
                        _ = v125
                    } else {
                        if false {
                            if false {
                                if false {
                                    v126 := GenS14{F0: "foo bar"}
                                    _ = v126
                                }
                                if true {
                                    var v127 bool = ((-2.25 + ((((0.1 + ((0.1 + (0.1 * 1.5)) * 1.5)) - 1e10) - (3.14159 * 1.5)) * float32(((1e10 - (3.14159 - 1.5)) + v124)))) > ((0.0 + (1.5 + 0.1)) / float32(v124)))
                                    _ = v127
                                }
                            }
                            if (!false) {
                                fmt.Println(v2, v1)
                                var v128 int8 = int8(v6)
                                v128 = (v109 + v109)
                                _ = v128
                            }
                        }
                    }
                    _ = v124
                }
                v129 := GenS10{F0: v7, F1: 36, F2: (v6 ^ v6)}
                fmt.Println(v6, v103)
                _ = v129
            }
            switch (fi110 % (((fi110 * v103)) | 1)) {
            case 32767:
                v4 += 2
            case 0:
                switch (("\n\t" + (v7 + (v7 + (((v7 + v7) + v7) + ("\n\t" + ("hello" + v7)))))) + "a") {
                case "":
                    if (0.1 >= 0.0) {
                        v130 := GenS9{F0: -101}
                        if (2 == (^(v6))) {
                            if true {
                                v131 := GenS15{F0: (v1 << ((v2) & 63)), F1: (v109 | (v109 + v109)), F2: (v114 << ((v4) & 63))}
                                _ = v131
                            }
                            if false {
                                var v132 float32 = float32(float64(-2.25))
                                var v133 int16 = v3
                                _ = v132
                                _ = v133
                            }
                        } else {
                            var v134 int32 = -2147483648
                            _ = v134
                        }
                        _ = v130
                    } else {
                        var v135 uint32 = 4294967295
                        var v136 uint8 = 127
                        if (543 < 255) {
                            var v137 uint64 = v6
                            var v138 int8 = 1
                            switch v7 {
                            case "hello":
                                v7 = ((((v7 + "hello") + v7) + "\n\t") + "hello")
                            case "foo bar":
                                v139 := GenS16{F0: uint32(v8)}
                                _ = v139
                            }
                            _ = v137
                            _ = v138
                        }
                        v1 = 402
                        _ = v135
                        _ = v136
                    }
                case "a":
                    v140 := GenS13{F0: (v4 & 65535), F1: v6}
                    _ = v140
                case "\n\t":
                    if false {
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            if false {
                                if true {
                                    v7 = (v7 + ("hello" + (v7 + v7)))
                                    var v141 uint32 = 970
                                    v142 := []int{v103, 2147483647}
                                    _ = v141
                                    _ = v142
                                }
                                v103 -= (fi110 & 32767)
                            } else {
                                v114 = v2
                                _ = len(v5)
                                v7 = (v7 + v7)
                            }
                        default:
                            v4 = 1
                            v2 = 1
                            switch ("hello" + "foo bar") {
                            case "a":
                                v7 = v7
                            case "\n\t":
                                _ = len(v5)
                            case "hello":
                                if (true && (v3 <= v3)) {
                                    v2 = (v4 % ((uint(v8)) | 1))
                                    v143 := GenS17{F0: (109 > -1), F1: 59}
                                    _ = v143
                                }
                            default:
                                if true {
                                    v144 := GenS18{F0: v3}
                                    var v145 uint32 = 872
                                    _ = v144
                                    _ = v145
                                }
                                v7 = v7
                                if (!((v109 ^ (v109 & v109)) != 100)) {
                                    v146 := GenS19{F0: 2}
                                    _ = v146
                                }
                            }
                        }
                    }
                }
            case 42:
                v147 := v5[884]
                _ = v147
            }
            var v148 int32 = -412
            v149 := GenS19{F0: 0}
            _ = fi110
            _ = v113
            _ = v114
            _ = v148
            _ = v149
        }
        _ = fi104
        _ = v109
    }
    for fi150 := 0; fi150 < 16; fi150++ {
        var v151 uint8 = 128
        var v152 int64 = (v1 & (v1 >> (((^(v4))) & 63)))
        if (!true) {
            fmt.Println(fi150)
        }
        _ = fi150
        _ = v151
        _ = v152
    }
    for fi153 := 0; fi153 < 11; fi153++ {
        v154 := []int{v103, (v103 - v103), 42}
        _ = fi153
        _ = v154
    }
    fmt.Println(v103, v103, v2)
    delete(v5, 798)
    v3 -= (^(v3))
    v155 := GenS9{F0: (v1 % ((int64(v3)) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v103
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
