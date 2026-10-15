package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint; F1 uint32; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint; F1 int; F2 bool }
type GenS12 struct { F0 uint64; F1 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 string }
type GenS15 struct { F0 int64; F1 int8; F2 uint32 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int16; F1 uint64; F2 int32 }
type GenS18 struct { F0 int8; F1 uint64; F2 bool }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int16; F1 uint16; F2 int }
type GenS21 struct { F0 int8; F1 int64 }
type GenS22 struct { F0 int32; F1 string }
type GenS23 struct { F0 uint8; F1 uint }
type GenS24 struct { F0 uint; F1 string }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 uint; F1 int8 }
type GenS27 struct { F0 int8; F1 uint8 }
type GenS28 struct { F0 bool; F1 int16 }
type GenS29 struct { F0 int32; F1 uint; F2 string }
type GenS30 struct { F0 int64; F1 int; F2 uint }
type GenS31 struct { F0 int32; F1 string }
type GenS32 struct { F0 uint; F1 string }
type GenS33 struct { F0 int; F1 int64; F2 int }
type GenS34 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS35 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS36 struct { F0 uint32; F1 int8 }
type GenS37 struct { F0 int; F1 int8 }
type GenS38 struct { F0 bool }

func main() {
    var v1 uint16 = 362
    var v2 float64 = (3.14159 + float64(1.5))
    v3 := make(map[string]bool)
    switch 32768 {
    case 0:
        v2 += (float64(1.5) / v2)
    case 127:
        var v4 int32 = 0
        _ = v4
    case 2147483647:
        v5 := GenS1{F0: 9, F1: 216, F2: -490}
        _ = v5
    }
    v6 := GenS2{F0: (!true)}
    var v7 int32 = 1
    v8 := []int{-439}
    if true {
        switch 886 {
        case 0:
            v9 := GenS1{F0: 128, F1: 0, F2: 601}
            _ = v9
        case 942:
            fmt.Println(v2, v1, v2)
        case 42:
            if false {
                var v10 uint16 = v1
                switch 127 {
                case 2147483647:
                    v1 = (v10 & (v1 * v1))
                case 32767:
                    v11 := GenS1{F0: uint8(v10), F1: 42, F2: -903}
                    _ = v11
                case 32768:
                    if (!true) {
                        if true {
                            fmt.Println(v1, v2, v2)
                            switch -583 {
                            case 32767:
                                if (107 == 1) {
                                    _ = len(v3)
                                    v12 := v3[("\n\t" + "hello")]
                                    var v13 int64 = 9223372036854775807
                                    _ = v12
                                    _ = v13
                                } else {
                                    var v14 uint16 = (v1 - v1)
                                    switch "" {
                                    case "a":
                                        v15 := GenS2{F0: (0 >= 1)}
                                        _ = v15
                                    case "\n\t":
                                        for fi16 := 0; fi16 < 4; fi16++ {
                                            for fi17 := 0; fi17 < 17; fi17++ {
                                                v18 := GenS1{F0: 127, F1: 42, F2: 9223372036854775807}
                                                _ = len(v8)
                                                v8 = append(v8, fi17)
                                                switch fi16 {
                                                case 4294967295:
                                                    switch "a" {
                                                    case "a":
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v19 float32 = 1.5
                                                            _ = v19
                                                        case "\n\t":
                                                            v20 := v3["a"]
                                                            _ = v20
                                                        case "hello":
                                                            if true {
                                                                if len(v8) > 0 {
                                                                    v8[0] = fi17
                                                                }
                                                                v3["foo bar"] = false
                                                                if false {
                                                                    var v21 uint16 = v14
                                                                    v7 = v7
                                                                    v1 += 65535
                                                                    _ = v21
                                                                }
                                                            } else {
                                                                var v22 string = "foo bar"
                                                                v2 = ((float64((float32(v2) + 1.5)) / v2) - ((v2 + float64(0.1)) - v2))
                                                                if ((-(v7)) >= v7) {
                                                                    var v23 uint8 = 1
                                                                    if (!(284 < (^(v7)))) {
                                                                        if (42 >= 818) {
                                                                            v1 = 0
                                                                            var v24 int8 = 100
                                                                            var v25 uint = 0
                                                                            _ = v24
                                                                            _ = v25
                                                                        }
                                                                    }
                                                                    _ = len(v8)
                                                                    var v26 string = v22
                                                                    var v27 uint8 = (^(v23))
                                                                    _ = v23
                                                                    _ = v26
                                                                    _ = v27
                                                                }
                                                                _ = v22
                                                            }
                                                        default:
                                                            var v28 uint = 4294967295
                                                            _ = v28
                                                        }
                                                    case "hello":
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v8 = append(v8, 127)
                                                        default:
                                                            var v29 string = "foo bar"
                                                            _ = v29
                                                        }
                                                    default:
                                                        var v30 int = 2147483647
                                                        _ = v30
                                                    }
                                                case -2147483648:
                                                    v2 = 1.5
                                                }
                                                v14 = (v10 << ((2) & 63))
                                                _ = fi17
                                                _ = v18
                                            }
                                            _ = fi16
                                        }
                                    case "hello":
                                        v7 = v7
                                    default:
                                        if false {
                                            v31 := GenS1{F0: 2, F1: 205, F2: -250}
                                            _ = v31
                                        } else {
                                            if false {
                                                if len(v8) > 0 {
                                                    v8[0] = -150
                                                }
                                                fmt.Println(v7, v2, v14)
                                            } else {
                                                v14 += 65535
                                                switch ("hello" + "\n\t") {
                                                case "a":
                                                    v14 -= 2
                                                case "hello":
                                                    var v32 int8 = 42
                                                    _ = v32
                                                case "\n\t":
                                                    v33 := GenS2{F0: ((2147483647 != 545) && true)}
                                                    _ = v33
                                                default:
                                                    v3["foo bar"] = true
                                                }
                                            }
                                            if true {
                                                v34 := GenS3{F0: 609, F1: -1, F2: 1}
                                                v35 := GenS2{F0: false}
                                                if (-2.25 > 1.5) {
                                                    _ = len(v8)
                                                    switch -2147483648 {
                                                    case 2147483647:
                                                        if (1e10 > 1.5) {
                                                            var v36 uint16 = v14
                                                            _ = v36
                                                        } else {
                                                            fmt.Println(v7, v2)
                                                            for fi37 := 0; fi37 < 2; fi37++ {
                                                                if (0 == 32767) {
                                                                    v2 = v2
                                                                    var v38 int64 = -539
                                                                    v10 += v14
                                                                    _ = v38
                                                                }
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                _ = len(v3)
                                                                _ = fi37
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    for fi39 := 0; fi39 < 19; fi39++ {
                                                        switch "\n\t" {
                                                        case "":
                                                            v40 := GenS4{F0: 384, F1: (("" + "") + ("\n\t" + "foo bar")), F2: 9223372036854775807}
                                                            _ = v40
                                                        case "a":
                                                            _ = len(v8)
                                                        case "\n\t":
                                                            for fi41 := 0; fi41 < 19; fi41++ {
                                                                var v42 int8 = -128
                                                                if false {
                                                                    var v43 string = "a"
                                                                    delete(v3, "\n\t")
                                                                    fmt.Println(v2)
                                                                    v44 := GenS2{F0: false}
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                _ = fi41
                                                                _ = v42
                                                            }
                                                        }
                                                        if (false || ((0.0 != (v2 * (((v2 + (v2 - (v2 / 1.5))) + 0.1) / v2))) && (!(18446744073709551615 < 2)))) {
                                                            if (((v10 & v10) < v10) || true) {
                                                                var v45 string = "\n\t"
                                                                if false {
                                                                    var v46 uint8 = 255
                                                                    v46 -= v46
                                                                    _ = v46
                                                                }
                                                                for fi47 := 0; fi47 < 19; fi47++ {
                                                                    v45 = v45
                                                                    _ = fi47
                                                                }
                                                                _ = v45
                                                            }
                                                            for fi48 := 0; fi48 < 19; fi48++ {
                                                                if false {
                                                                    switch fi48 {
                                                                    case 127:
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v1 += v14
                                                                        case "a":
                                                                            var v49 float64 = v2
                                                                            _ = v49
                                                                        default:
                                                                            fmt.Println(fi48)
                                                                        }
                                                                        if (!(!(!(!(!false))))) {
                                                                            switch fi48 {
                                                                            case 1:
                                                                                v10 += (v10 ^ (v1 * v1))
                                                                            case 9223372036854775807:
                                                                                if (-128 <= 1) {
                                                                                    var v50 int16 = -32768
                                                                                    _ = v50
                                                                                } else {
                                                                                    fmt.Println(v7)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v10 += (^(v14))
                                                                            var v51 uint32 = 2
                                                                            _ = v51
                                                                        }
                                                                        fmt.Println(v1)
                                                                    case 9223372036854775807:
                                                                        v52 := GenS5{F0: 163}
                                                                        _ = v52
                                                                    }
                                                                    var v53 int8 = 127
                                                                    if (132 == 128) {
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            var v54 int64 = -47
                                                                            _ = v54
                                                                        case "hello":
                                                                            if true {
                                                                                _ = len(v8)
                                                                                var v55 string = "foo bar"
                                                                                if false {
                                                                                    var v56 int = 32767
                                                                                    var v57 float32 = -2.25
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                } else {
                                                                                    fmt.Println(v2)
                                                                                    if true {
                                                                                        switch 993 {
                                                                                        case -9223372036854775808:
                                                                                            if len(v8) > 0 {
                                                                                                v8[0] = 822
                                                                                            }
                                                                                        case -1:
                                                                                            _ = len(v3)
                                                                                        case 32768:
                                                                                            v58 := v3["foo bar"]
                                                                                            _ = v58
                                                                                        default:
                                                                                            if false {
                                                                                                var v59 uint32 = 2
                                                                                                _ = v59
                                                                                            }
                                                                                        }
                                                                                        var v60 int64 = -1
                                                                                        _ = v60
                                                                                    }
                                                                                    v14 -= v14
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        v61 := GenS2{F0: (1e10 == (float32((float64(((1e10 / 1e10) / ((((1e10 + 1e10) - 3.14159) - (0.0 - (3.14159 * 3.14159))) + float32(1.5)))) / v2)) * float32(v2)))}
                                                                                        _ = v61
                                                                                    }
                                                                                }
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                        v62 := GenS4{F0: 65535, F1: "a", F2: 241}
                                                                        v63 := v3["a"]
                                                                        _ = v62
                                                                        _ = v63
                                                                    } else {
                                                                        var v64 int16 = -1
                                                                        _ = v64
                                                                    }
                                                                    switch (fi48 / (((^(fi39))) | 1)) {
                                                                    case 873:
                                                                        v8 = append(v8, (-(fi39)))
                                                                    default:
                                                                        if false {
                                                                            switch -2147483648 {
                                                                            case 32767:
                                                                                v65 := GenS6{F0: int64(v7), F1: 889}
                                                                                fmt.Println(v2, v2, v14)
                                                                                _ = v65
                                                                            case 42:
                                                                                if true {
                                                                                    v7 += (-(v7))
                                                                                    v66 := GenS7{F0: 1, F1: 65535, F2: -1}
                                                                                    switch fi39 {
                                                                                    case 921:
                                                                                        var v67 int64 = -9223372036854775808
                                                                                        _ = v67
                                                                                    case 0:
                                                                                        v7 -= (-(v7))
                                                                                    }
                                                                                    _ = v66
                                                                                } else {
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v68 := GenS8{F0: 0, F1: (v14 * v10)}
                                                                                        _ = v68
                                                                                    case "hello":
                                                                                        v1 += v10
                                                                                        var v69 uint = 0
                                                                                        _ = v69
                                                                                    case "":
                                                                                        v70 := GenS8{F0: 474, F1: 42}
                                                                                        _ = v70
                                                                                    default:
                                                                                        var v71 uint = 65535
                                                                                        _ = v71
                                                                                    }
                                                                                    var v72 int32 = v7
                                                                                    if (v2 <= -2.25) {
                                                                                        v73 := GenS1{F0: 234, F1: 2, F2: -1}
                                                                                        v74 := GenS9{F0: 902, F1: ("" + ("\n\t" + ("a" + (("\n\t" + "") + "\n\t"))))}
                                                                                        v75 := GenS10{F0: "hello"}
                                                                                        v76 := GenS11{F0: 315, F1: 127, F2: (-1 == 502)}
                                                                                        if true {
                                                                                            var v77 int64 = 89
                                                                                            _ = v77
                                                                                        }
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    } else {
                                                                                        var v78 float64 = -2.25
                                                                                        _ = v78
                                                                                    }
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                            v8 = append(v8, (fi39 ^ (^(fi48))))
                                                                        } else {
                                                                            if false {
                                                                                v79 := GenS12{F0: 818, F1: 35}
                                                                                var v80 uint8 = 127
                                                                                v81 := GenS7{F0: 18446744073709551615, F1: 4294967295, F2: int8(v80)}
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                v1 += (v14 - 1)
                                                                                var v82 uint8 = 149
                                                                                _ = v82
                                                                            }
                                                                            if true {
                                                                                var v83 int32 = 42
                                                                                _ = v83
                                                                            }
                                                                        }
                                                                    }
                                                                    var v84 bool = true
                                                                    _ = v53
                                                                    _ = v84
                                                                }
                                                                _ = fi48
                                                            }
                                                        }
                                                        v2 = (1.5 + v2)
                                                        _ = fi39
                                                    }
                                                }
                                                _ = v34
                                                _ = v35
                                            }
                                            fmt.Println(v1, v10, v1)
                                        }
                                    }
                                    fmt.Println(v14)
                                    _ = v14
                                }
                            case -2147483648:
                                v2 = v2
                            case 42:
                                for fi85 := 0; fi85 < 7; fi85++ {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = fi85
                                }
                            }
                        }
                        v86 := v3[("a" + "\n\t")]
                        v2 = ((v2 * 1e10) * v2)
                        _ = v86
                    }
                }
                switch 4294967295 {
                case -9223372036854775808:
                    var v87 float32 = 3.14159
                    _ = v87
                case -219:
                    v88 := GenS8{F0: 1, F1: (v1 ^ 42)}
                    _ = v88
                case 2147483647:
                    _ = len(v3)
                }
                _ = v10
            } else {
                if true {
                    var v89 int32 = -1
                    fmt.Println(v89, v1)
                    v89 += (v89 * 2147483647)
                    _ = v89
                } else {
                    var v90 int32 = -277
                    _ = v90
                }
                for fi91 := 0; fi91 < 17; fi91++ {
                    v2 -= -2.25
                    v7 = (v7 * (-(v7)))
                    fmt.Println(v1, fi91)
                    _ = fi91
                }
            }
        }
    } else {
        for fi92 := 0; fi92 < 4; fi92++ {
            for fi93 := 0; fi93 < 18; fi93++ {
                v94 := GenS6{F0: 9223372036854775807, F1: 65535}
                v3["foo bar"] = (42 < 42)
                _ = fi93
                _ = v94
            }
            var v95 uint8 = 42
            _ = fi92
            _ = v95
        }
        if true {
            if (false || (!((!(!false)) && false))) {
                v96 := GenS1{F0: 1, F1: 42, F2: 0}
                _ = v96
            } else {
                var v97 uint32 = 2147483647
                if false {
                    var v98 int64 = 1
                    v99 := GenS13{F0: 1}
                    var v100 uint8 = 42
                    _ = v100
                    _ = v98
                    _ = v99
                }
                _ = v97
            }
            for fi101 := 0; fi101 < 8; fi101++ {
                var v102 uint = 255
                if true {
                    v7 = v7
                    if false {
                        for fi103 := 0; fi103 < 14; fi103++ {
                            v104 := GenS14{F0: v7, F1: "\n\t"}
                            v102 += 0
                            var v105 uint8 = 2
                            if true {
                                v102 = v102
                                var v106 int32 = 0
                                _ = v106
                            }
                            if (-624 == 9223372036854775807) {
                                fmt.Println(v102, fi103, v7)
                            }
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                    } else {
                        fmt.Println(fi101, v7, v1)
                        if ((v102 >> ((v102) & 63)) == v102) {
                            for fi107 := 0; fi107 < 19; fi107++ {
                                v3[("foo bar" + ("\n\t" + ""))] = (-76 <= 0)
                                _ = fi107
                            }
                            var v108 uint32 = 4294967295
                            for fi109 := 0; fi109 < 12; fi109++ {
                                var v110 int16 = 1000
                                if (v1 < (v1 / (((v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))) | 1))) {
                                    if (((true && true) && (!false)) || (false || (!true))) {
                                        fmt.Println(v108)
                                    } else {
                                        v111 := v3["a"]
                                        if len(v8) > 0 {
                                            v8[0] = 32768
                                        }
                                        switch (fi109 / ((32767) | 1)) {
                                        case 2147483647:
                                            var v112 string = "\n\t"
                                            _ = v112
                                        case -9223372036854775808:
                                            var v113 int16 = -37
                                            _ = v113
                                        case 127:
                                            fmt.Println(v108, v1, v7)
                                        default:
                                            var v114 uint16 = v1
                                            _ = v114
                                        }
                                        if ("a" <= "a") {
                                            switch "" {
                                            case "hello":
                                                if (!(v110 <= (v110 & v110))) {
                                                    if v111 {
                                                        var v115 int64 = 42
                                                        var v116 bool = (v110 == 1000)
                                                        var v117 uint = v102
                                                        _ = v115
                                                        _ = v116
                                                        _ = v117
                                                    } else {
                                                        var v118 int16 = -32768
                                                        if len(v8) > 0 {
                                                            v8[0] = (-(fi101))
                                                        }
                                                        v2 = float64((0.0 + (1e10 - float32(float64((((1e10 * 3.14159) + 0.1) + 3.14159))))))
                                                        _ = v118
                                                    }
                                                    var v119 int64 = 42
                                                    v111 = v111
                                                    if (2 > uint8(v1)) {
                                                        var v120 int32 = (v7 ^ (v7 + v7))
                                                        v110 += v110
                                                        var v121 int8 = 23
                                                        _ = v120
                                                        _ = v121
                                                    } else {
                                                        var v122 int64 = 350
                                                        v111 = false
                                                        v119 += 42
                                                        _ = v122
                                                    }
                                                    fmt.Println(v2, v111)
                                                    _ = v119
                                                } else {
                                                    v123 := GenS1{F0: 0, F1: v102, F2: 42}
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    var v124 uint8 = 0
                                                    v110 = (v110 ^ int16(v1))
                                                    if true {
                                                        v125 := GenS7{F0: (v102 / ((4294967295) | 1)), F1: v108, F2: 79}
                                                        v126 := GenS2{F0: v111}
                                                        _ = v125
                                                        _ = v126
                                                    } else {
                                                        var v127 int32 = (v7 + 2147483647)
                                                        _ = v127
                                                    }
                                                    _ = v123
                                                    _ = v124
                                                }
                                            case "a":
                                                var v128 uint32 = 974
                                                _ = v128
                                            default:
                                                _ = len(v8)
                                            }
                                            switch fi101 {
                                            case -2147483648:
                                                var v129 uint64 = 278
                                                _ = v129
                                            case 1:
                                                v7 = -1
                                            default:
                                                if (((0 > 0) || false) || v111) {
                                                    var v130 float64 = v2
                                                    if (0 == -990) {
                                                        v102 += (^(v102))
                                                        v3["foo bar"] = (1e10 >= (0.1 + v130))
                                                        v3[("" + "a")] = false
                                                    }
                                                    _ = v130
                                                } else {
                                                    v131 := GenS15{F0: 0, F1: 100, F2: (v108 << (((v102 ^ v102)) & 63))}
                                                    v132 := GenS16{F0: (v110 & (^(v110)))}
                                                    _ = v131
                                                    _ = v132
                                                }
                                            }
                                            fmt.Println(v110, v102)
                                            v7 = int32(v102)
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        } else {
                                            v110 += v110
                                            var v133 uint64 = 42
                                            delete(v3, "foo bar")
                                            v2 -= v2
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            _ = v133
                                        }
                                        _ = v111
                                    }
                                    if (fi109 >= fi109) {
                                        v7 += (-(v7))
                                        fmt.Println(v102, v7)
                                        switch "\n\t" {
                                        case "\n\t":
                                            v1 = (^(v1))
                                        case "":
                                            var v134 bool = true
                                            _ = v134
                                        default:
                                            v135 := GenS1{F0: 1, F1: v102, F2: 9223372036854775807}
                                            _ = v135
                                        }
                                    } else {
                                        v108 += (v108 % (((^(v108))) | 1))
                                        v136 := GenS17{F0: (v110 ^ (-(v110))), F1: 2, F2: v7}
                                        if true {
                                            var v137 int16 = v110
                                            _ = v137
                                        } else {
                                            if ((0 >= (v102 << (((v102 % ((2) | 1))) & 63))) && ((!(!true)) || ((v7 / (((v7 + (v7 ^ v7))) | 1)) >= v7))) {
                                                switch fi109 {
                                                case 127:
                                                    var v138 uint16 = v1
                                                    _ = v138
                                                }
                                                var v139 int64 = 1
                                                fmt.Println(v110)
                                                _ = v139
                                            }
                                        }
                                        v102 -= v102
                                        v110 = -1
                                        _ = v136
                                    }
                                    v140 := GenS8{F0: 42, F1: 765}
                                    _ = v140
                                } else {
                                    if false {
                                        fmt.Println(v110)
                                        v7 = -516
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    } else {
                                        fmt.Println(v110)
                                        v2 = v2
                                    }
                                }
                                _ = fi109
                                _ = v110
                            }
                            _ = v108
                        } else {
                            switch (fi101 * fi101) {
                            case 128:
                                if (388 <= 4294967295) {
                                    switch "" {
                                    case "a":
                                        switch ("" + "hello") {
                                        case "\n\t":
                                            var v141 int8 = 0
                                            _ = v141
                                        case "foo bar":
                                            if len(v8) > 0 {
                                                v8[0] = fi101
                                            }
                                        case "a":
                                            fmt.Println(fi101)
                                        }
                                    default:
                                        if (346 >= 0) {
                                            v2 = v2
                                        } else {
                                            delete(v3, ("a" + "a"))
                                        }
                                    }
                                }
                            case 250:
                                var v142 int64 = 12
                                _ = v142
                            case 32768:
                                var v143 int32 = 1
                                _ = v143
                            default:
                                if false {
                                    var v144 int16 = -949
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi101, v2)
                                    case "a":
                                        for fi145 := 0; fi145 < 5; fi145++ {
                                            v7 -= v7
                                            fmt.Println(v1, v102, v1)
                                            var v146 float64 = v2
                                            var v147 int64 = 1
                                            v102 -= v102
                                            _ = fi145
                                            _ = v146
                                            _ = v147
                                        }
                                    case "foo bar":
                                        v144 -= v144
                                    }
                                    v148 := GenS7{F0: (v102 / ((v102) | 1)), F1: 0, F2: int8(v1)}
                                    _ = v144
                                    _ = v148
                                } else {
                                    if true {
                                        v2 = 1.5
                                        var v149 uint8 = 9
                                        _ = v149
                                    }
                                }
                            }
                        }
                        var v150 float64 = v2
                        _ = v150
                    }
                } else {
                    var v151 uint32 = 0
                    for fi152 := 0; fi152 < 4; fi152++ {
                        if (!(82 == 42)) {
                            v153 := v3["hello"]
                            switch -1 {
                            case -1:
                                var v154 int8 = -1
                                _ = v154
                            case -175:
                                var v155 int64 = int64(v102)
                                var v156 uint8 = 42
                                var v157 int16 = 32767
                                _ = v155
                                _ = v156
                                _ = v157
                            }
                            v153 = true
                            _ = v153
                        }
                        _ = fi152
                    }
                    _ = v151
                }
                v7 -= v7
                _ = fi101
                _ = v102
            }
        } else {
            if ((v7 | 10) > v7) {
                fmt.Println(v1, v2)
            }
        }
    }
    if (0 > -1) {
        v158 := GenS2{F0: false}
        v2 += float64(-2.25)
        v159 := GenS18{F0: 127, F1: 18446744073709551615, F2: false}
        _ = v158
        _ = v159
    } else {
        v160 := v3[(("hello" + "foo bar") + "hello")]
        var v161 uint64 = 708
        _ = v160
        _ = v161
    }
    for fi162 := 0; fi162 < 4; fi162++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi162
    }
    v163 := GenS18{F0: int8(v7), F1: 927, F2: true}
    var v164 string = "\n\t"
    fmt.Println(v7, v2, v1)
    v1 -= (v1 / ((v1) | 1))
    v165 := v3[""]
    v166 := GenS19{F0: 4294967295}
    for fi167 := 0; fi167 < 12; fi167++ {
        _ = len(v8)
        switch (fi167 + 660) {
        case 668:
            if v165 {
                for fi168 := 0; fi168 < 15; fi168++ {
                    v169 := GenS9{F0: 461, F1: v164}
                    var v170 int16 = 1
                    var v171 float64 = float64(0.1)
                    if (-26 > 31) {
                        var v172 float32 = (1.5 - (((0.0 + 0.1) / 1.5) / 1.5))
                        var v173 uint64 = 42
                        var v174 int16 = v170
                        _ = v172
                        _ = v173
                        _ = v174
                    }
                    _ = fi168
                    _ = v169
                    _ = v170
                    _ = v171
                }
                var v175 int = (^(fi167))
                for fi176 := 0; fi176 < 8; fi176++ {
                    v177 := v3[(v164 + v164)]
                    v178 := GenS1{F0: 1, F1: 18446744073709551615, F2: 42}
                    _ = fi176
                    _ = v177
                    _ = v178
                }
                _ = v175
            }
        case 32767:
            for fi179 := 0; fi179 < 15; fi179++ {
                delete(v3, v164)
                if false {
                    fmt.Println(v165, v1, v165)
                    v180 := GenS13{F0: 2}
                    fmt.Println(v1)
                    v181 := GenS18{F0: 42, F1: 216, F2: v165}
                    v182 := GenS7{F0: 255, F1: 2, F2: 1}
                    switch ((((v164 + v164) + (v164 + "a")) + (v164 + "")) + ((((v164 + (v164 + v164)) + ("" + v164)) + v164) + "")) {
                    case "a":
                        var v183 int8 = 0
                        _ = v183
                    case "\n\t":
                        switch "hello" {
                        case "":
                            var v184 int8 = 1
                            _ = v184
                        case "a":
                            var v185 int8 = -1
                            _ = v185
                        case "foo bar":
                            var v186 float32 = (0.0 + 1e10)
                            _ = v186
                        default:
                            switch v164 {
                            case "a":
                                v164 = (v164 + "a")
                            case "foo bar":
                                switch fi179 {
                                case -2147483648:
                                    v1 -= v1
                                case -1:
                                    var v187 uint = 18446744073709551615
                                    _ = v187
                                }
                            case "hello":
                                var v188 uint = 42
                                _ = v188
                            }
                        }
                    case "hello":
                        if (!(("" + v164) >= "a")) {
                            var v189 int32 = (v7 & v7)
                            _ = v189
                        } else {
                            var v190 uint8 = 1
                            _ = v190
                        }
                    default:
                        v191 := GenS20{F0: 0, F1: v1, F2: (fi179 >> ((18446744073709551615) & 63))}
                        _ = v191
                    }
                    switch v164 {
                    case "a":
                        if (false || v165) {
                            if (v165 && (!v165)) {
                                var v192 uint8 = 175
                                v193 := GenS2{F0: v165}
                                if v165 {
                                    var v194 float64 = -2.25
                                    if false {
                                        if ((!(42 != -1)) || false) {
                                            delete(v3, v164)
                                        } else {
                                            if (v165 || false) {
                                                var v195 uint16 = (v1 ^ 32767)
                                                switch (v164 + v164) {
                                                case "foo bar":
                                                    v2 -= 0.1
                                                default:
                                                    if ((-2.25 + 0.1) >= (0.0 * 0.1)) {
                                                        v3[(("a" + v164) + "")] = (v164 > ("a" + v164))
                                                    } else {
                                                        var v196 uint64 = 134
                                                        _ = v196
                                                    }
                                                }
                                                if v165 {
                                                    if (-86 > -1) {
                                                        _ = len(v3)
                                                    } else {
                                                        v194 = v194
                                                        var v197 uint32 = 1
                                                        _ = v197
                                                    }
                                                    v165 = v165
                                                    v198 := GenS5{F0: v192}
                                                    _ = v198
                                                } else {
                                                    fmt.Println(v195, v194)
                                                    v199 := GenS21{F0: 127, F1: 0}
                                                    _ = len(v8)
                                                    _ = v199
                                                }
                                                v200 := GenS3{F0: 0, F1: -1, F2: 605}
                                                var v201 uint8 = (v192 - 1)
                                                _ = v195
                                                _ = v200
                                                _ = v201
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                if true {
                                                    var v202 int16 = 159
                                                    v1 = 752
                                                    switch v164 {
                                                    case "hello":
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        fmt.Println(v7)
                                                        v203 := GenS9{F0: uint32(v192), F1: ("a" + v164)}
                                                        v204 := GenS20{F0: v202, F1: v1, F2: fi167}
                                                        _ = v203
                                                        _ = v204
                                                    default:
                                                        v8 = append(v8, (^(fi167)))
                                                    }
                                                    var v205 int8 = -122
                                                    _ = v202
                                                    _ = v205
                                                } else {
                                                    var v206 uint8 = (v192 << ((293) & 63))
                                                    fmt.Println(fi167)
                                                    fmt.Println(fi167)
                                                    _ = v206
                                                }
                                            case "hello":
                                                if false {
                                                    if len(v8) > 0 {
                                                        v8[0] = (fi179 + -205)
                                                    }
                                                    fmt.Println(v194, v1)
                                                } else {
                                                    v207 := GenS22{F0: v7, F1: ((v164 + "foo bar") + ((((v164 + (v164 + ("foo bar" + ((v164 + v164) + v164)))) + v164) + v164) + v164))}
                                                    v7 += -2147483648
                                                    var v208 float32 = 1e10
                                                    _ = v207
                                                    _ = v208
                                                }
                                            default:
                                                v2 = (v2 - 1.5)
                                            }
                                        }
                                        if true {
                                            fmt.Println(v192, v194, v194)
                                        }
                                    }
                                    if (((!(-1 > -32768)) || v165) || ((1e10 + (0.1 + 3.14159)) == 0.0)) {
                                        if v165 {
                                            fmt.Println(v192, v2)
                                            if false {
                                                if true {
                                                    var v209 uint64 = 394
                                                    fmt.Println(v209)
                                                    _ = v209
                                                }
                                                v194 -= ((1e10 * v2) - ((v2 - (0.0 + 0.1)) + (v194 * (float64(0.1) + float64(3.14159)))))
                                                var v210 bool = false
                                                _ = v210
                                            }
                                        }
                                        switch (fi179 - (fi179 & (-(fi179)))) {
                                        case 127:
                                            var v211 uint16 = (v1 % ((v1) | 1))
                                            _ = v211
                                        case -2147483648:
                                            v1 = 0
                                        case -9223372036854775808:
                                            if (v164 < v164) {
                                                switch fi179 {
                                                case 912:
                                                    switch v164 {
                                                    case "hello":
                                                        if true {
                                                            var v212 int64 = 113
                                                            _ = v212
                                                        }
                                                    case "\n\t":
                                                        var v213 int = (fi179 >> ((857) & 63))
                                                        _ = v213
                                                    case "":
                                                        var v214 bool = ((v192 - (v192 & (^(v192)))) >= v192)
                                                        _ = v214
                                                    default:
                                                        var v215 string = ((v164 + "") + "hello")
                                                        _ = v215
                                                    }
                                                case 2147483647:
                                                    if v165 {
                                                        v216 := GenS23{F0: v192, F1: 0}
                                                        v217 := GenS11{F0: 65535, F1: (fi179 * fi179), F2: (!(37 > 42))}
                                                        _ = v216
                                                        _ = v217
                                                    } else {
                                                        v194 += (float64(1.5) + v194)
                                                        switch 32767 {
                                                        case 2147483647:
                                                            if len(v8) >= 2 {
                                                                v8 = v8[1:len(v8)]
                                                            }
                                                        case -43:
                                                            v218 := v3["hello"]
                                                            _ = v218
                                                        }
                                                        v219 := GenS5{F0: v192}
                                                        _ = v219
                                                    }
                                                }
                                                v194 = float64(-2.25)
                                                _ = len(v8)
                                            } else {
                                                if false {
                                                    if v165 {
                                                        v194 -= v194
                                                        v220 := GenS12{F0: 1, F1: 631}
                                                        if ((608 != 42) || v165) {
                                                            v194 += -2.25
                                                            fmt.Println(fi167)
                                                        }
                                                        _ = v220
                                                    } else {
                                                        v221 := GenS15{F0: -111, F1: -128, F2: 2147483647}
                                                        switch v164 {
                                                        case "hello":
                                                            v2 -= float64(1.5)
                                                        case "\n\t":
                                                            if (!(((-9223372036854775808 > -906) && (!(v192 >= v192))) && (-1 < 127))) {
                                                                if ((!(115 >= 825)) && v165) {
                                                                    fmt.Println(v192, fi179, v1)
                                                                    v164 = ("a" + "a")
                                                                } else {
                                                                    fmt.Println(v194, v192, fi167)
                                                                    v222 := GenS24{F0: 1, F1: v164}
                                                                    _ = v222
                                                                }
                                                                var v223 int8 = 42
                                                                v223 -= (v223 % ((v223) | 1))
                                                                _ = v223
                                                            } else {
                                                                fmt.Println(fi167, fi167, v194)
                                                                v7 = (^(v7))
                                                                v224 := v3["a"]
                                                                _ = v224
                                                            }
                                                        default:
                                                            var v225 int64 = -1
                                                            _ = v225
                                                        }
                                                        v226 := GenS18{F0: 0, F1: 1, F2: v165}
                                                        _ = v221
                                                        _ = v226
                                                    }
                                                    v164 = ""
                                                }
                                                var v227 uint8 = 127
                                                _ = v227
                                            }
                                        }
                                        var v228 uint = 0
                                        _ = v228
                                    } else {
                                        switch ("hello" + v164) {
                                        case "foo bar":
                                            fmt.Println(v7, v2, fi179)
                                        case "\n\t":
                                            switch (-(fi179)) {
                                            case 1:
                                                delete(v3, ("hello" + v164))
                                            case 127:
                                                if ((^(v1)) == (^(v1))) {
                                                    switch fi167 {
                                                    case 1:
                                                        if (v165 && ((((0.0 / 0.1) / ((float32(v194) / -2.25) / 1.0)) / 0.1) == 3.14159)) {
                                                            switch -1 {
                                                            case 0:
                                                                fmt.Println(fi179)
                                                            case 42:
                                                                fmt.Println(v2)
                                                            default:
                                                                fmt.Println(fi179)
                                                            }
                                                            fmt.Println(fi179, v1, v192)
                                                            var v229 uint8 = v192
                                                            switch "a" {
                                                            case "":
                                                                var v230 uint8 = v229
                                                                _ = v230
                                                            default:
                                                                v231 := GenS25{F0: 0}
                                                                _ = v231
                                                            }
                                                            _ = v229
                                                        }
                                                    case 9223372036854775807:
                                                        var v232 int8 = -109
                                                        _ = v232
                                                    case 32768:
                                                        v233 := GenS3{F0: -286, F1: -109, F2: uint32(v1)}
                                                        _ = v233
                                                    }
                                                } else {
                                                    v234 := GenS20{F0: 1000, F1: 65535, F2: 42}
                                                    var v235 int32 = (v7 >> ((255) & 63))
                                                    var v236 uint32 = 65535
                                                    _ = v234
                                                    _ = v235
                                                    _ = v236
                                                }
                                            default:
                                                v192 = v192
                                                if false {
                                                    v164 = (((v164 + (((("hello" + v164) + (v164 + v164)) + v164) + v164)) + "a") + v164)
                                                    v237 := GenS26{F0: 18446744073709551615, F1: 42}
                                                    _ = v237
                                                } else {
                                                    fmt.Println(v164, v7, v2)
                                                    var v238 int16 = 0
                                                    _ = v238
                                                }
                                            }
                                        default:
                                            v239 := GenS5{F0: 42}
                                            _ = v239
                                        }
                                        v164 = (v164 + "hello")
                                        if v165 {
                                            v240 := GenS17{F0: 1, F1: 207, F2: (v7 * v7)}
                                            v7 -= 2147483647
                                            if (int16(v1) >= -1) {
                                                v241 := v3[v164]
                                                switch v164 {
                                                case "a":
                                                    if v165 {
                                                        var v242 int = (fi167 >> ((4294967295) & 63))
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        if v165 {
                                                            v8 = append(v8, -774)
                                                            if true {
                                                                v243 := GenS27{F0: 126, F1: (v192 / ((v192) | 1))}
                                                                _ = v243
                                                            } else {
                                                                v192 = (v192 | (v192 & 1))
                                                                if v241 {
                                                                    if (float64(0.0) == 1e10) {
                                                                        fmt.Println(v2, v242, v242)
                                                                        v1 += (v1 - 255)
                                                                    } else {
                                                                        var v244 int16 = 0
                                                                        _ = v244
                                                                    }
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v245 int32 = (v7 << ((255) & 63))
                                                                        _ = v245
                                                                    default:
                                                                        var v246 bool = ((!false) && (false && v241))
                                                                        _ = v246
                                                                    }
                                                                }
                                                            }
                                                            if (!((v1 | (v1 << ((585) & 63))) <= v1)) {
                                                                var v247 int = (fi167 * (v242 - v242))
                                                                v7 += (v7 | v7)
                                                                var v248 int16 = 146
                                                                _ = v247
                                                                _ = v248
                                                            }
                                                        } else {
                                                            delete(v3, "foo bar")
                                                        }
                                                        var v249 bool = (0 != -32768)
                                                        v250 := GenS25{F0: 1}
                                                        _ = v242
                                                        _ = v249
                                                        _ = v250
                                                    }
                                                case "\n\t":
                                                    if len(v8) > 0 {
                                                        v8[0] = 458
                                                    }
                                                default:
                                                    if false {
                                                        v251 := GenS8{F0: 335, F1: (v1 ^ (v1 - v1))}
                                                        _ = v251
                                                    } else {
                                                        v252 := GenS14{F0: (-(v7)), F1: v164}
                                                        fmt.Println(v164, fi167)
                                                        if len(v8) > 0 {
                                                            v8[0] = int(v7)
                                                        }
                                                        _ = v252
                                                    }
                                                }
                                                switch fi179 {
                                                case 1:
                                                    switch v164 {
                                                    case "foo bar":
                                                        v253 := GenS28{F0: v165, F1: 201}
                                                        _ = v253
                                                    case "hello":
                                                        v254 := GenS22{F0: v7, F1: "foo bar"}
                                                        _ = v254
                                                    default:
                                                        v3[("" + (("a" + v164) + "a"))] = (-455 != -1)
                                                    }
                                                }
                                                _ = v241
                                            }
                                            _ = v240
                                        }
                                        var v255 int = fi167
                                        _ = v255
                                    }
                                    _ = v194
                                }
                                _ = v192
                                _ = v193
                            }
                            v165 = v165
                        } else {
                            delete(v3, (v164 + "\n\t"))
                            var v256 int64 = 992
                            var v257 uint32 = 936
                            _ = v256
                            _ = v257
                        }
                    default:
                        switch v164 {
                        case "\n\t":
                            var v258 int8 = -14
                            _ = v258
                        default:
                            v259 := GenS4{F0: 359, F1: v164, F2: -1}
                            _ = v259
                        }
                    }
                    delete(v3, v164)
                    if ((true || (-1 <= 1)) || v165) {
                        switch 2147483647 {
                        case 32768:
                            if v165 {
                                fmt.Println(fi179, v7)
                                if (4294967295 <= 1) {
                                    if (!(!true)) {
                                        if v165 {
                                            v165 = v165
                                            fmt.Println(v164, v7, v2)
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            v3[("hello" + v164)] = (false && (((v165 && v165) || (227 == 168)) || true))
                                            v260 := GenS12{F0: 18446744073709551615, F1: 632}
                                            _ = v260
                                        } else {
                                            switch -156 {
                                            case 128:
                                                if (3.14159 > (-2.25 * 0.0)) {
                                                    if ((v7 + v7) < (v7 - (v7 & (v7 << ((643) & 63))))) {
                                                        v1 -= v1
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        v7 += -1
                                                    } else {
                                                        if false {
                                                            if len(v8) >= 2 {
                                                                v8 = v8[1:len(v8)]
                                                            }
                                                        }
                                                        var v261 int32 = v7
                                                        _ = len(v3)
                                                        _ = v261
                                                    }
                                                    v7 = v7
                                                    if (!(2 > 255)) {
                                                        v262 := GenS7{F0: 255, F1: 65535, F2: -1}
                                                        var v263 uint16 = v1
                                                        if ((v165 || ((3.14159 + 0.1) != 1e10)) && (v165 && v165)) {
                                                            v3[(v164 + "hello")] = (true && (!((-436 != v7) && false)))
                                                            v7 -= v7
                                                        }
                                                        _ = v262
                                                        _ = v263
                                                    } else {
                                                        fmt.Println(v165)
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    }
                                                }
                                            case 32768:
                                                switch fi179 {
                                                case -1:
                                                    var v264 uint = 1
                                                    _ = v264
                                                case 500:
                                                    if v165 {
                                                        var v265 bool = true
                                                        var v266 uint64 = 18446744073709551615
                                                        _ = v265
                                                        _ = v266
                                                    }
                                                default:
                                                    var v267 uint16 = v1
                                                    _ = v267
                                                }
                                            default:
                                                fmt.Println(v7, v164, v165)
                                            }
                                        }
                                        if (true && v165) {
                                            var v268 int = 262
                                            var v269 bool = v165
                                            if (1 != 2147483648) {
                                                switch v268 {
                                                case 9223372036854775807:
                                                    var v270 float32 = 1e10
                                                    _ = v270
                                                }
                                                v271 := GenS5{F0: 2}
                                                _ = v271
                                            } else {
                                                if (!(v165 || v269)) {
                                                    v272 := GenS29{F0: (v7 / (((v7 >> ((65535) & 63))) | 1)), F1: 803, F2: v164}
                                                    _ = v272
                                                } else {
                                                    v269 = v269
                                                    fmt.Println(v269, v269, v268)
                                                }
                                                v268 = (v268 << ((42) & 63))
                                                if v269 {
                                                    fmt.Println(v268, v165)
                                                    var v273 int32 = v7
                                                    if v165 {
                                                        v274 := GenS30{F0: -9223372036854775808, F1: v268, F2: 18446744073709551615}
                                                        switch "" {
                                                        case "foo bar":
                                                            v275 := GenS31{F0: (v273 + (^(v7))), F1: v164}
                                                            _ = v275
                                                        case "hello":
                                                            var v276 float64 = float64(3.14159)
                                                            _ = v276
                                                        }
                                                        _ = v274
                                                    } else {
                                                        switch (fi167 << ((0) & 63)) {
                                                        case 0:
                                                            v273 = v7
                                                        case 128:
                                                            v269 = ((v164 + v164) >= v164)
                                                        case 32767:
                                                            if v269 {
                                                                if ("a" == ("\n\t" + (v164 + "foo bar"))) {
                                                                    v277 := GenS10{F0: v164}
                                                                    v165 = v269
                                                                    fmt.Println(fi167, v165, v273)
                                                                    _ = v277
                                                                } else {
                                                                    switch 32768 {
                                                                    case 2147483647:
                                                                        var v278 int32 = v7
                                                                        _ = v278
                                                                    case 32767:
                                                                        if v269 {
                                                                            v279 := GenS32{F0: 2, F1: v164}
                                                                            _ = v279
                                                                        } else {
                                                                            var v280 string = ("hello" + ((v164 + v164) + ((v164 + v164) + "")))
                                                                            switch (fi179 >> ((65535) & 63)) {
                                                                            case -937:
                                                                                v280 = "hello"
                                                                            case 128:
                                                                                fmt.Println(v269)
                                                                            case -283:
                                                                                var v281 int64 = 42
                                                                                _ = v281
                                                                            default:
                                                                                fmt.Println(fi179, v280)
                                                                                v1 = (v1 ^ (v1 & 65535))
                                                                            }
                                                                            v282 := GenS17{F0: -1, F1: 0, F2: (v273 * 366)}
                                                                            _ = v280
                                                                            _ = v282
                                                                        }
                                                                    }
                                                                    fmt.Println(v164)
                                                                    v8 = append(v8, fi167)
                                                                    v1 = (v1 & v1)
                                                                }
                                                            } else {
                                                                v7 -= v7
                                                                v8 = append(v8, (fi167 | (fi179 / (((v268 * (fi167 % ((fi167) | 1)))) | 1))))
                                                            }
                                                        }
                                                        if false {
                                                            if false {
                                                                var v283 uint64 = 158
                                                                var v284 uint = 255
                                                                v7 -= v7
                                                                _ = v283
                                                                _ = v284
                                                            }
                                                            _ = len(v8)
                                                            v269 = ((v165 && (v165 || v165)) && (!(v164 <= (((v164 + "") + ((v164 + v164) + "hello")) + v164))))
                                                        }
                                                    }
                                                    _ = v273
                                                }
                                            }
                                            _ = v268
                                            _ = v269
                                        } else {
                                            if true {
                                                v165 = (-9223372036854775808 < 0)
                                                v2 -= v2
                                                if true {
                                                    switch (^(fi179)) {
                                                    case -9223372036854775808:
                                                        v1 = v1
                                                    case 42:
                                                        fmt.Println(v2)
                                                        var v285 int32 = v7
                                                        _ = v285
                                                    case -872:
                                                        var v286 uint64 = 246
                                                        _ = v286
                                                    }
                                                    var v287 int8 = 42
                                                    _ = v287
                                                }
                                            } else {
                                                v3[v164] = (2 != 42)
                                                v288 := GenS33{F0: 132, F1: -1, F2: (fi167 % ((fi167) | 1))}
                                                v7 += (v7 * 42)
                                                var v289 uint32 = 2147483648
                                                v290 := GenS34{F0: v1, F1: 2147483648, F2: 255}
                                                _ = v288
                                                _ = v289
                                                _ = v290
                                            }
                                            v291 := GenS18{F0: 127, F1: 42, F2: false}
                                            _ = v291
                                        }
                                    } else {
                                        v1 -= v1
                                    }
                                    fmt.Println(fi179, fi167, v2)
                                } else {
                                    var v292 uint64 = 0
                                    v2 = v2
                                    v2 += ((v2 / v2) + -2.25)
                                    _ = v292
                                }
                                fmt.Println(v7, fi179)
                                switch v164 {
                                case "hello":
                                    v1 = v1
                                default:
                                    if v165 {
                                        switch 1 {
                                        case -1:
                                            var v293 float64 = v2
                                            _ = v293
                                        }
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v294 int = -2147483648
                                            _ = v294
                                        case "":
                                            fmt.Println(v7)
                                        case "\n\t":
                                            v2 += ((v2 - 0.1) * -2.25)
                                        }
                                        if ("\n\t" >= v164) {
                                            v295 := GenS24{F0: 955, F1: v164}
                                            var v296 uint32 = 42
                                            v297 := GenS27{F0: 100, F1: 128}
                                            _ = v295
                                            _ = v296
                                            _ = v297
                                        }
                                        if (v7 == -766) {
                                            v1 += v1
                                        }
                                    } else {
                                        fmt.Println(v2)
                                    }
                                }
                                v7 = 42
                            } else {
                                switch (fi179 / (((fi167 ^ -44)) | 1)) {
                                case -381:
                                    v164 = v164
                                case -1:
                                    var v298 bool = (414 < 0)
                                    _ = v298
                                case -980:
                                    v299 := GenS17{F0: 42, F1: 888, F2: (v7 % (((v7 & 621)) | 1))}
                                    _ = v299
                                default:
                                    if v165 {
                                        v300 := GenS35{F0: 4294967295, F1: 2, F2: 188}
                                        v1 = v1
                                        v1 += v1
                                        v301 := GenS1{F0: 255, F1: 255, F2: 9223372036854775807}
                                        _ = v300
                                        _ = v301
                                    }
                                }
                                if ((v7 <= (-(v7))) || (-9223372036854775808 > -52)) {
                                    var v302 int16 = -1
                                    v303 := GenS36{F0: 65535, F1: 0}
                                    _ = v302
                                    _ = v303
                                } else {
                                    switch ((v164 + ("hello" + ("" + v164))) + "") {
                                    case "\n\t":
                                        _ = len(v3)
                                        v8 = append(v8, (fi167 ^ -1))
                                        var v304 int64 = 42
                                        _ = v304
                                    case "hello":
                                        switch v164 {
                                        case "foo bar":
                                            switch (((v164 + (v164 + v164)) + v164) + (v164 + v164)) {
                                            case "":
                                                var v305 uint32 = 2147483647
                                                _ = v305
                                            case "\n\t":
                                                if ((float64((1.5 + (float32(v2) - float32(0.0)))) + float64(1.5)) > 1e10) {
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    var v306 int64 = int64(v1)
                                                    switch fi179 {
                                                    case -1:
                                                        var v307 uint16 = v1
                                                        _ = v307
                                                    case 0:
                                                        switch v164 {
                                                        case "hello":
                                                            switch v164 {
                                                            case "foo bar":
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                v3[v164] = true
                                                            case "\n\t":
                                                                var v308 float32 = ((0.1 - (((-2.25 * 1e10) + (0.1 / (1.5 - 3.14159))) + ((1.5 + (1.5 / float32(1e10))) - 1.5))) * (-2.25 * -2.25))
                                                                var v309 int8 = -1
                                                                _ = v308
                                                                _ = v309
                                                            default:
                                                                switch v164 {
                                                                case "\n\t":
                                                                    v164 = ""
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v310 := GenS16{F0: -636}
                                                            _ = v310
                                                        default:
                                                            v311 := GenS11{F0: 948, F1: (fi179 ^ (-(fi179))), F2: (v2 == (((0.0 - 1e10) + float64(((-2.25 * 0.0) - (((0.1 + 1.5) - float32(v2)) - float32(v2))))) * ((1e10 / (v2 / v2)) / v2)))}
                                                            _ = v311
                                                        }
                                                    }
                                                    _ = v306
                                                } else {
                                                    var v312 bool = (!((18446744073709551615 == 65535) && ((0.1 * 1e10) <= float64(-2.25))))
                                                    _ = v312
                                                }
                                            case "hello":
                                                switch v164 {
                                                case "\n\t":
                                                    delete(v3, v164)
                                                default:
                                                    v1 -= (v1 / ((v1) | 1))
                                                }
                                            default:
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                            }
                                        }
                                    case "a":
                                        v165 = (v165 && (-128 == -106))
                                    }
                                    var v313 float32 = (float32((v2 * v2)) + (0.0 / float32(v2)))
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v313
                                }
                            }
                        default:
                            var v314 float32 = 0.0
                            _ = v314
                        }
                    } else {
                        if v165 {
                            if true {
                                fmt.Println(fi179, fi167)
                                var v315 uint64 = 4294967295
                                fmt.Println(v165)
                                _ = v315
                            }
                            fmt.Println(v165)
                        } else {
                            v7 = int32(fi179)
                            delete(v3, "hello")
                            v316 := GenS22{F0: (v7 + v7), F1: ""}
                            _ = v316
                        }
                        v165 = (!(v165 || (2 != 65535)))
                        v164 = v164
                    }
                    var v317 float64 = float64((3.14159 / float32((v2 + (float64(((1e10 - (1e10 / ((((3.14159 / float32((1e10 / v2))) / (0.0 - 0.1)) / (0.1 - 1e10)) + (float32(v2) * (float32(float64(0.0)) + -2.25))))) - 1e10)) + v2)))))
                    v317 -= 1.5
                    _ = v180
                    _ = v181
                    _ = v182
                    _ = v317
                }
                v318 := GenS37{F0: (fi167 % ((589) | 1)), F1: 112}
                _ = fi179
                _ = v318
            }
        case 1:
            fmt.Println(v164, v7)
        default:
            v319 := GenS38{F0: (!(0 != -853))}
            _ = v319
        }
        fmt.Println(v7)
        _ = fi167
    }
    var v320 int16 = 1
    v321 := GenS12{F0: 42, F1: 155}
    fmt.Println("v1:", v1)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v320:", v320)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v163
    _ = v164
    _ = v165
    _ = v166
    _ = v2
    _ = v3
    _ = v320
    _ = v321
    _ = v6
    _ = v7
    _ = v8
}
