package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 uint64; F2 int }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint64; F1 uint }
type GenS13 struct { F0 int16; F1 string; F2 uint16 }
type GenS14 struct { F0 uint32; F1 uint; F2 int }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int; F1 bool; F2 int8 }
type GenS17 struct { F0 int; F1 int8 }
type GenS18 struct { F0 int32; F1 bool }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 351
    var v3 int16 = 0
    v4 := v1[530]
    fmt.Println(v3, v3)
    fmt.Println(v2, v4)
    v3 = int16(v2)
    v5 := GenS1{F0: -38}
    var v6 int16 = (v3 >> ((42) & 63))
    var v7 uint64 = 525
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 0
        fmt.Println(v9, v4)
        _ = fi8
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    switch -659 {
    case 128:
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 float64 = (3.14159 * -2.25)
                fmt.Println(v6)
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    if false {
        var v14 int64 = (v10 - 852)
        _ = v14
    } else {
        if true {
            v10 += -664
            v2 -= (v2 + v2)
            if (!(((-2.25 == 3.14159) || (v10 != 42)) || true)) {
                if true {
                    switch v4 {
                    case "hello":
                        v4 = v4
                    case "a":
                        for fi15 := 0; fi15 < 5; fi15++ {
                            switch int(v3) {
                            case 9223372036854775807:
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    var v17 int = 128
                                    _ = fi16
                                    _ = v17
                                }
                            case -9223372036854775808:
                                v18 := GenS2{F0: (((("a" + (v4 + v4)) + "") + ("foo bar" + v4)) + ("a" + v4)), F1: 100}
                                _ = v18
                            case 1:
                                v7 -= v7
                            }
                            for fi19 := 0; fi19 < 15; fi19++ {
                                if true {
                                    if false {
                                        if false {
                                            fmt.Println(v10, v10)
                                            v7 += 0
                                            var v20 int32 = 0
                                            _ = v20
                                        } else {
                                            if false {
                                                v4 = v4
                                                delete(v1, fi15)
                                                var v21 int32 = 42
                                                _ = v21
                                            } else {
                                                if true {
                                                    v1[(fi15 - fi19)] = v4
                                                }
                                                v6 -= v6
                                            }
                                        }
                                        if false {
                                            v22 := GenS3{F0: -128, F1: 0}
                                            v7 -= 42
                                            _ = v22
                                        } else {
                                            if ((fi15 & fi15) != fi19) {
                                                v23 := GenS1{F0: 100}
                                                switch ((v4 + "a") + "foo bar") {
                                                case "\n\t":
                                                    v24 := []int{(fi15 - (fi19 % (((^(fi19))) | 1))), 32768, (fi19 >> ((255) & 63))}
                                                    var v25 int16 = int16(v7)
                                                    _ = v24
                                                    _ = v25
                                                case "":
                                                    var v26 uint8 = 2
                                                    _ = v26
                                                case "foo bar":
                                                    if true {
                                                        var v27 uint32 = 2
                                                        _ = v27
                                                    }
                                                }
                                                v7 = (v7 >> ((527) & 63))
                                                _ = v23
                                            }
                                            _ = len(v1)
                                        }
                                    }
                                    var v28 bool = true
                                    _ = v28
                                } else {
                                    fmt.Println(v4)
                                    if (false || true) {
                                        var v29 string = v4
                                        _ = v29
                                    }
                                }
                                if true {
                                    v30 := GenS4{F0: v10, F1: (4294967295 > 2147483648), F2: 23}
                                    switch fi19 {
                                    case 2147483647:
                                        v2 += v2
                                    case 0:
                                        var v31 bool = ((1.5 * (-2.25 - (3.14159 * 0.0))) <= 3.14159)
                                        _ = v31
                                    case 146:
                                        if true {
                                            v3 -= 42
                                        } else {
                                            v32 := []int{-822, (fi15 / ((fi19) | 1)), int(v10), fi19}
                                            v33 := GenS4{F0: v10, F1: false, F2: 128}
                                            v6 += (v3 + (v6 + v6))
                                            v34 := GenS5{F0: -1}
                                            v35 := v1[fi15]
                                            v2 += v2
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                    default:
                                        v36 := []int{(fi15 / (((fi19 >> ((255) & 63))) | 1)), fi19, (^(fi19)), 128}
                                        _ = v36
                                    }
                                    _ = v30
                                }
                                _ = fi19
                            }
                            _ = fi15
                        }
                    case "foo bar":
                        v37 := GenS6{F0: (v7 | (v7 | v7)), F1: 2147483648}
                        _ = v37
                    }
                    fmt.Println(v2, v7, v6)
                    var v38 uint32 = 4294967295
                    _ = v38
                }
                switch 32768 {
                case -947:
                    v39 := []int{127, 32767, 4294967295}
                    _ = v39
                case -2147483648:
                    if (!true) {
                        if (true && true) {
                            var v40 float64 = (-2.25 * float64((1e10 - 0.0)))
                            _ = v40
                        } else {
                            switch -594 {
                            case 42:
                                for fi41 := 0; fi41 < 13; fi41++ {
                                    fmt.Println(v7, v4)
                                    v42 := GenS7{F0: 1, F1: 148}
                                    _ = fi41
                                    _ = v42
                                }
                            case -228:
                                for fi43 := 0; fi43 < 11; fi43++ {
                                    v44 := GenS8{F0: 779}
                                    v2 = 953
                                    if false {
                                        v10 += v10
                                        if false {
                                            for fi45 := 0; fi45 < 2; fi45++ {
                                                if (255 <= 28) {
                                                    if false {
                                                        v4 = (((v4 + "") + "a") + "foo bar")
                                                        if (!true) {
                                                            switch "a" {
                                                            case "":
                                                                var v46 bool = false
                                                                _ = v46
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v10)
                                                    var v47 float32 = (0.0 - float32(0.0))
                                                    v2 -= 42
                                                    _ = v47
                                                } else {
                                                    v48 := GenS8{F0: v7}
                                                    if (false && true) {
                                                        switch v4 {
                                                        case "":
                                                            v10 = -1
                                                        case "hello":
                                                            v49 := GenS4{F0: v10, F1: false, F2: 127}
                                                            _ = v49
                                                        case "\n\t":
                                                            v50 := []int{}
                                                            _ = v50
                                                        }
                                                        fmt.Println(v10)
                                                        if (false && true) {
                                                            var v51 uint16 = 0
                                                            switch -48 {
                                                            case 742:
                                                                v1[fi45] = ((("\n\t" + ("a" + v4)) + (((v4 + v4) + (((v4 + v4) + v4) + "a")) + v4)) + "foo bar")
                                                            case -9223372036854775808:
                                                                var v52 int64 = v10
                                                                if ((0.0 != -2.25) && (-22 <= -1)) {
                                                                    switch v4 {
                                                                    case "\n\t":
                                                                        v6 -= -194
                                                                    case "":
                                                                        v51 -= v51
                                                                    }
                                                                    var v53 float64 = 3.14159
                                                                    _ = v53
                                                                } else {
                                                                    fmt.Println(v7)
                                                                    v54 := GenS2{F0: "a", F1: 100}
                                                                    var v55 uint = 65535
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                if (1 > 0) {
                                                                    v56 := []int{(fi43 + 4294967295), (fi43 - fi45), (fi43 / (((fi45 ^ 42)) | 1)), fi45}
                                                                    switch ("\n\t" + v4) {
                                                                    case "foo bar":
                                                                        var v57 uint = 1
                                                                        _ = v57
                                                                    case "\n\t":
                                                                        var v58 float32 = ((0.1 / 1.0) - 0.1)
                                                                        _ = v58
                                                                    case "":
                                                                        var v59 uint16 = v51
                                                                        _ = v59
                                                                    default:
                                                                        v4 = "foo bar"
                                                                    }
                                                                    _ = v56
                                                                } else {
                                                                    fmt.Println(fi45, v4, v7)
                                                                    v60 := GenS9{F0: -5, F1: v7, F2: (^(fi45))}
                                                                    switch v4 {
                                                                    case "hello":
                                                                        switch ("foo bar" + v4) {
                                                                        case "a":
                                                                            fmt.Println(v52, fi43, v52)
                                                                        case "foo bar":
                                                                            v61 := []int{(fi45 | 32767), (fi45 << ((255) & 63)), fi43}
                                                                            _ = v61
                                                                        }
                                                                    default:
                                                                        var v62 uint8 = 128
                                                                        _ = v62
                                                                    }
                                                                    _ = v60
                                                                }
                                                                _ = v52
                                                            }
                                                            _ = v51
                                                        }
                                                    }
                                                    var v63 uint = 1
                                                    v64 := GenS10{F0: 4294967295, F1: 128}
                                                    if false {
                                                        var v65 uint64 = v7
                                                        var v66 uint = (v63 - (v63 + 4294967295))
                                                        switch (v4 + (v4 + (v4 + ("foo bar" + "foo bar")))) {
                                                        case "":
                                                            fmt.Println(v65, v3, v7)
                                                        case "a":
                                                            var v67 float32 = 0.0
                                                            _ = v67
                                                        }
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    v68 := []int{-9223372036854775808, fi45, int(v7), fi45}
                                                    var v69 uint32 = 899
                                                    if true {
                                                        var v70 bool = false
                                                        _ = v70
                                                    } else {
                                                        v6 -= 331
                                                    }
                                                    _ = v48
                                                    _ = v63
                                                    _ = v64
                                                    _ = v68
                                                    _ = v69
                                                }
                                                v3 = 657
                                                _ = fi45
                                            }
                                        } else {
                                            var v71 uint = 65535
                                            _ = v71
                                        }
                                        switch (("foo bar" + v4) + (v4 + (v4 + v4))) {
                                        case "foo bar":
                                            for fi72 := 0; fi72 < 15; fi72++ {
                                                v6 = 1
                                                if (-802 >= (v6 & v6)) {
                                                    v6 += (v3 | v3)
                                                    v73 := []int{}
                                                    _ = v73
                                                } else {
                                                    var v74 uint64 = 2
                                                    _ = v74
                                                }
                                                _ = fi72
                                            }
                                        case "\n\t":
                                            if true {
                                                v6 = -125
                                                fmt.Println(v4, v10, v7)
                                            }
                                        case "hello":
                                            switch fi43 {
                                            case 0:
                                                v75 := GenS10{F0: 2, F1: 255}
                                                _ = v75
                                            case -2147483648:
                                                if ((((((^(fi43)) <= fi43) && (!((0 > 0) && true))) || true) || true) && (true && false)) {
                                                    if (v2 < (v2 | v2)) {
                                                        if true {
                                                            for fi76 := 0; fi76 < 16; fi76++ {
                                                                var v77 int16 = (v6 + v6)
                                                                v10 -= v10
                                                                v78 := GenS11{F0: v4}
                                                                _ = fi76
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                        } else {
                                                            v79 := GenS12{F0: v7, F1: 1}
                                                            var v80 int32 = 2147483647
                                                            var v81 int16 = (^(v6))
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        switch 42 {
                                                        case -156:
                                                            if (((v6 / ((-227) | 1)) >= (v6 ^ v3)) || (false || ((fi43 - (fi43 << ((117) & 63))) < fi43))) {
                                                                for fi82 := 0; fi82 < 11; fi82++ {
                                                                    if false {
                                                                        var v83 uint16 = (v2 | (v2 << ((42) & 63)))
                                                                        v2 = (v83 - (v2 >> ((18446744073709551615) & 63)))
                                                                        if (!((0.1 != 0.0) || (false && false))) {
                                                                            if (!true) {
                                                                                var v84 uint64 = v7
                                                                                v4 = (v4 + ("" + "a"))
                                                                                var v85 int8 = 100
                                                                                var v86 uint64 = v84
                                                                                _ = v84
                                                                                _ = v85
                                                                                _ = v86
                                                                            } else {
                                                                                var v87 uint32 = 784
                                                                                _ = v87
                                                                            }
                                                                            var v88 uint64 = (^(v7))
                                                                            _ = v88
                                                                        } else {
                                                                            fmt.Println(v83, v2)
                                                                            v4 = ((("hello" + v4) + "hello") + "")
                                                                        }
                                                                        _ = v83
                                                                    }
                                                                    v7 -= v7
                                                                    fmt.Println(v6, v2)
                                                                    _ = fi82
                                                                }
                                                            } else {
                                                                var v89 uint32 = uint32(v7)
                                                                if true {
                                                                    if false {
                                                                        var v90 float32 = -2.25
                                                                        _ = v90
                                                                    }
                                                                    v6 -= v6
                                                                }
                                                                v91 := []int{fi43, fi43}
                                                                _ = v89
                                                                _ = v91
                                                            }
                                                        case -877:
                                                            for fi92 := 0; fi92 < 14; fi92++ {
                                                                v2 += v2
                                                                _ = fi92
                                                            }
                                                        }
                                                        for fi93 := 0; fi93 < 4; fi93++ {
                                                            if (!true) {
                                                                switch 4294967295 {
                                                                case -826:
                                                                    switch "" {
                                                                    case "hello":
                                                                        v1[462] = (v4 + v4)
                                                                    default:
                                                                        switch ("foo bar" + "a") {
                                                                        case "\n\t":
                                                                            v1[fi43] = (v4 + (v4 + ("foo bar" + ("\n\t" + (v4 + ((v4 + v4) + "hello"))))))
                                                                            v94 := GenS13{F0: (v6 - 0), F1: (("foo bar" + (((v4 + v4) + (v4 + ("" + v4))) + (v4 + (v4 + v4)))) + (v4 + ("\n\t" + (v4 + (((v4 + "\n\t") + ("hello" + v4)) + v4))))), F2: v2}
                                                                            switch (v4 + v4) {
                                                                            case "":
                                                                                fmt.Println(fi93)
                                                                            }
                                                                            _ = v94
                                                                        case "":
                                                                            var v95 uint8 = 245
                                                                            _ = v95
                                                                        default:
                                                                            v96 := GenS2{F0: v4, F1: -128}
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                case 9223372036854775807:
                                                                    fmt.Println(v6)
                                                                    if true {
                                                                        v6 = v3
                                                                        v97 := v1[-309]
                                                                        fmt.Println(v6, fi93, v4)
                                                                        _ = v97
                                                                    } else {
                                                                        var v98 uint32 = 1
                                                                        var v99 int64 = (v10 << ((uint(v7)) & 63))
                                                                        v100 := GenS5{F0: 127}
                                                                        _ = v100
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    var v101 int16 = (v3 | v6)
                                                                    v102 := GenS13{F0: -32768, F1: (("foo bar" + "a") + v4), F2: v2}
                                                                    _ = v101
                                                                    _ = v102
                                                                case -781:
                                                                    v4 = ((v4 + ((v4 + "a") + v4)) + (("\n\t" + v4) + v4))
                                                                }
                                                                if (true || (false && true)) {
                                                                    v3 -= (v6 + (v3 % ((v6) | 1)))
                                                                } else {
                                                                    v103 := []int{}
                                                                    fmt.Println(v3, v4)
                                                                    v2 += (v2 << ((949) & 63))
                                                                    _ = v103
                                                                }
                                                            } else {
                                                                v7 -= 878
                                                            }
                                                            switch 0 {
                                                            case -1:
                                                                fmt.Println(v4, fi93)
                                                            default:
                                                                var v104 int64 = v10
                                                                _ = v104
                                                            }
                                                            _ = fi93
                                                        }
                                                    } else {
                                                        switch (v4 + (("hello" + "\n\t") + ((v4 + (v4 + v4)) + "\n\t"))) {
                                                        case "a":
                                                            for fi105 := 0; fi105 < 11; fi105++ {
                                                                var v106 bool = (-344 >= v10)
                                                                _ = fi105
                                                                _ = v106
                                                            }
                                                        case "":
                                                            var v107 float64 = ((3.14159 * float64(3.14159)) - ((((0.0 * -2.25) + -2.25) + (0.0 * 0.0)) - 1.5))
                                                            _ = v107
                                                        }
                                                    }
                                                    if false {
                                                        v108 := GenS14{F0: 2, F1: 65535, F2: (fi43 + (fi43 | int(v10)))}
                                                        v109 := []int{-2147483648, fi43}
                                                        switch int(v2) {
                                                        case -2147483648:
                                                            for fi110 := 0; fi110 < 1; fi110++ {
                                                                if false {
                                                                    v4 = v4
                                                                    v111 := GenS15{F0: 2147483647}
                                                                    _ = v111
                                                                }
                                                                if false {
                                                                    var v112 int64 = (^(v10))
                                                                    switch fi43 {
                                                                    case -2147483648:
                                                                        v3 = -1
                                                                    case 9223372036854775807:
                                                                        var v113 string = ("a" + v4)
                                                                        _ = v113
                                                                    default:
                                                                        v114 := v1[(fi43 | 127)]
                                                                        _ = v114
                                                                    }
                                                                    v115 := GenS1{F0: -128}
                                                                    _ = v112
                                                                    _ = v115
                                                                } else {
                                                                    var v116 int8 = 42
                                                                    var v117 int32 = 2147483647
                                                                    fmt.Println(fi43, fi43)
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                                _ = fi110
                                                            }
                                                        case -9223372036854775808:
                                                            for fi118 := 0; fi118 < 9; fi118++ {
                                                                var v119 uint = 2
                                                                if (42 >= 121) {
                                                                    fmt.Println(v119, fi118)
                                                                } else {
                                                                    if (false || false) {
                                                                        var v120 uint64 = 220
                                                                        v3 += int16(fi43)
                                                                        v119 += 73
                                                                        _ = v120
                                                                    }
                                                                    _ = len(v109)
                                                                    var v121 int32 = -1
                                                                    _ = v121
                                                                }
                                                                v122 := GenS16{F0: 127, F1: true, F2: 127}
                                                                _ = fi118
                                                                _ = v119
                                                                _ = v122
                                                            }
                                                        default:
                                                            for fi123 := 0; fi123 < 19; fi123++ {
                                                                var v124 int = (fi123 ^ -571)
                                                                var v125 float64 = ((1e10 + 0.1) * 1.5)
                                                                v7 = (v7 % ((v7) | 1))
                                                                _ = fi123
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                        }
                                                        _ = v108
                                                        _ = v109
                                                    } else {
                                                        v126 := []int{fi43, (fi43 - (fi43 & 380)), fi43, (-(fi43))}
                                                        _ = v126
                                                    }
                                                    var v127 float32 = (0.1 - (float32(-2.25) - 0.1))
                                                    _ = v127
                                                }
                                            default:
                                                if true {
                                                    switch fi43 {
                                                    case 9223372036854775807:
                                                        if true {
                                                            switch v4 {
                                                            case "hello":
                                                                v6 += (v6 & (-(v3)))
                                                            case "foo bar":
                                                                var v128 int64 = (v10 ^ v10)
                                                                _ = v128
                                                            default:
                                                                if false {
                                                                    v1[(-(fi43))] = (v4 + "\n\t")
                                                                    for fi129 := 0; fi129 < 19; fi129++ {
                                                                        var v130 bool = (0 != 42)
                                                                        v131 := GenS12{F0: v7, F1: 1}
                                                                        v4 = v4
                                                                        _ = fi129
                                                                        _ = v130
                                                                        _ = v131
                                                                    }
                                                                    var v132 bool = (1.5 != -2.25)
                                                                    _ = v132
                                                                } else {
                                                                    v133 := GenS17{F0: 1, F1: -1}
                                                                    v134 := GenS18{F0: 1, F1: true}
                                                                    fmt.Println(fi43, v6, v2)
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                            }
                                                        } else {
                                                            var v135 uint8 = 51
                                                            _ = v135
                                                        }
                                                    case 128:
                                                        var v136 int16 = 1
                                                        _ = v136
                                                    case 32768:
                                                        v137 := []int{}
                                                        _ = v137
                                                    default:
                                                        v6 += v3
                                                    }
                                                    var v138 uint64 = v7
                                                    for fi139 := 0; fi139 < 17; fi139++ {
                                                        if false {
                                                            delete(v1, fi43)
                                                            v140 := GenS15{F0: -1}
                                                            v141 := GenS4{F0: (v10 >> ((135) & 63)), F1: true, F2: 42}
                                                            _ = v140
                                                            _ = v141
                                                        }
                                                        fmt.Println(v7, v138, v10)
                                                        _ = fi139
                                                    }
                                                    _ = v138
                                                }
                                            }
                                        }
                                    } else {
                                        var v142 int = fi43
                                        _ = v142
                                    }
                                    _ = fi43
                                    _ = v44
                                }
                            case 559:
                                if ((v2 & v2) != 584) {
                                    var v143 uint16 = 32767
                                    for fi144 := 0; fi144 < 3; fi144++ {
                                        v3 = 321
                                        switch ("a" + ("a" + "foo bar")) {
                                        case "a":
                                            for fi145 := 0; fi145 < 3; fi145++ {
                                                var v146 uint8 = 164
                                                switch "\n\t" {
                                                case "a":
                                                    v7 = (v7 >> ((1) & 63))
                                                }
                                                _ = fi145
                                                _ = v146
                                            }
                                        case "":
                                            if (v2 > 32767) {
                                                for fi147 := 0; fi147 < 18; fi147++ {
                                                    var v148 bool = false
                                                    v149 := GenS18{F0: -2147483648, F1: v148}
                                                    _ = fi147
                                                    _ = v148
                                                    _ = v149
                                                }
                                                var v150 uint = uint(v7)
                                                _ = v150
                                            } else {
                                                fmt.Println(v2, v7, v10)
                                                v151 := GenS3{F0: 0, F1: (v6 - -1)}
                                                v152 := GenS2{F0: v4, F1: 127}
                                                _ = v151
                                                _ = v152
                                            }
                                        default:
                                            v153 := GenS7{F0: 1, F1: 255}
                                            _ = v153
                                        }
                                        v6 = 0
                                        _ = fi144
                                    }
                                    _ = v143
                                } else {
                                    _ = len(v1)
                                    v3 = v6
                                    v154 := []int{}
                                    _ = v154
                                }
                            }
                            v155 := GenS17{F0: 97, F1: 0}
                            v6 -= (v3 ^ 484)
                            _ = v155
                        }
                    }
                default:
                    v156 := GenS19{F0: "hello"}
                    _ = v156
                }
                for fi157 := 0; fi157 < 11; fi157++ {
                    v6 = (v6 / ((1) | 1))
                    if false {
                        if true {
                            v158 := GenS6{F0: v7, F1: 582}
                            v2 -= v2
                            _ = v158
                        } else {
                            v7 = (v7 + v7)
                            var v159 float64 = 0.0
                            _ = v159
                        }
                        v1[(^(fi157))] = ((v4 + v4) + "\n\t")
                    } else {
                        v160 := GenS20{F0: (v7 % (((v7 & (v7 << ((uint(v7)) & 63)))) | 1)), F1: -472}
                        v161 := []int{fi157}
                        _ = v160
                        _ = v161
                    }
                    if (100 >= 1) {
                        var v162 int8 = -128
                        v4 = (v4 + v4)
                        v10 = (v10 | v10)
                        _ = v162
                    } else {
                        v163 := []int{(fi157 >> ((332) & 63)), 933, 127, fi157}
                        var v164 int64 = v10
                        _ = v163
                        _ = v164
                    }
                    _ = fi157
                }
            } else {
                v10 -= (v10 & (^(v10)))
            }
        }
        v165 := GenS20{F0: (^(v7)), F1: -420}
        fmt.Println(v2, v10, v6)
        _ = v165
    }
    v166 := v1[0]
    if (v7 == 4294967295) {
        fmt.Println(v6, v7)
        var v167 float64 = 0.1
        v168 := GenS9{F0: 100, F1: (v7 * (v7 & (v7 + v7))), F2: -9223372036854775808}
        _ = v167
        _ = v168
    }
    switch 32768 {
    case -9223372036854775808:
        v166 = (("\n\t" + ((v4 + (("" + ("\n\t" + v166)) + "foo bar")) + (("a" + "a") + v166))) + "\n\t")
    case 42:
        v169 := GenS12{F0: 18446744073709551615, F1: 1}
        _ = v169
    }
    v170 := []int{32768, 1, 828, -965}
    v171 := GenS15{F0: 0}
    v10 -= (v10 ^ (v10 - 9223372036854775807))
    fmt.Println(v7, v6, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v166:", v166)
    fmt.Println("slicelen v170:", len(v170), v170)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v166
    _ = v170
    _ = v171
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
