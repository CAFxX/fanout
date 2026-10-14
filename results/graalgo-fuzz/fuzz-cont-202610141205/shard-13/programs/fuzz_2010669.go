package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint8; F1 int32; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 uint8; F1 string; F2 uint16 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 uint64; F1 uint }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint8; F1 uint16 }
type GenS16 struct { F0 int; F1 uint16 }
type GenS17 struct { F0 int; F1 int16; F2 int64 }
type GenS18 struct { F0 uint32; F1 int8; F2 bool }
type GenS19 struct { F0 int; F1 uint64 }
type GenS20 struct { F0 int64; F1 int8; F2 int }
type GenS21 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS22 struct { F0 uint64; F1 int8; F2 int64 }
type GenS23 struct { F0 uint64; F1 int32 }
type GenS24 struct { F0 uint64; F1 bool; F2 bool }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 uint32; F1 uint; F2 int64 }
type GenS28 struct { F0 uint }
type GenS29 struct { F0 uint32; F1 int8 }
type GenS30 struct { F0 int; F1 int64; F2 uint64 }
type GenS31 struct { F0 bool }
type GenS32 struct { F0 uint64 }
type GenS33 struct { F0 uint8; F1 uint32 }
type GenS34 struct { F0 uint }
type GenS35 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS36 struct { F0 bool; F1 int; F2 int8 }
type GenS37 struct { F0 uint; F1 int32 }
type GenS38 struct { F0 int64; F1 uint32; F2 int64 }
type GenS39 struct { F0 uint8 }
type GenS40 struct { F0 int16; F1 bool; F2 int }
type GenS41 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 int8, p2 uint8) int32 {
    v4 := GenS1{F0: -128, F1: 472}
    fmt.Println(p0)
    v5 := GenS2{F0: false}
    var v6 uint32 = 1
    v7 := GenS3{F0: -1, F1: "", F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int8 = -1
    v8 := GenS4{F0: 0, F1: -9223372036854775808}
    if (42 < 42) {
        if (false && true) {
            v9 := []int{0}
            if (808 == 65535) {
                switch "a" {
                case "hello":
                    fmt.Println(v3)
                case "a":
                    if false {
                        var v10 string = "foo bar"
                        v10 = "foo bar"
                        _ = v10
                    } else {
                        v11 := GenS5{F0: -9223372036854775808, F1: 0}
                        if len(v9) > 0 {
                            v9[0] = 9223372036854775807
                        }
                        _ = v11
                    }
                default:
                    v9 = append(v9, 42)
                }
                v12 := GenS6{F0: 0, F1: false}
                _ = v12
            } else {
                var v13 int64 = 383
                _ = v13
            }
            _ = v9
        } else {
            if (42 == -2147483648) {
                v3 -= v3
                if true {
                    v14 := GenS1{F0: 100, F1: 0}
                    _ = v14
                } else {
                    if false {
                        v15 := []int{9223372036854775807, -9223372036854775808}
                        _ = v15
                    } else {
                        fmt.Println(v3)
                        var v16 int = 2147483647
                        _ = v16
                    }
                    v17 := []int{-20, -2147483648, 127, 128}
                    _ = v17
                }
            }
        }
        if true {
            for fi18 := 0; fi18 < 18; fi18++ {
                v3 = (v3 - v3)
                for fi19 := 0; fi19 < 6; fi19++ {
                    v3 -= 127
                    v3 = v3
                    _ = fi19
                }
                _ = fi18
            }
        } else {
            var v20 uint8 = 0
            if true {
                v20 -= (v20 * 1)
                var v21 bool = false
                v22 := GenS5{F0: -1, F1: -1}
                _ = v21
                _ = v22
            }
            _ = v20
        }
        v23 := []int{-728}
        _ = v23
    }
    v3 = (^(v3))
    v24 := []int{1, 1, 128}
    switch 0 {
    case 1:
        fmt.Println(v3)
    case 128:
        for fi25 := 0; fi25 < 17; fi25++ {
            switch 127 {
            case -314:
                var v26 int = fi25
                _ = v26
            case 32767:
                v27 := GenS7{F0: 180, F1: -1, F2: (v3 | 42)}
                _ = v27
            case 32768:
                for fi28 := 0; fi28 < 17; fi28++ {
                    if (("a" + ("\n\t" + ("" + "a"))) != (("hello" + "foo bar") + "\n\t")) {
                        v29 := GenS8{F0: 4294967295, F1: 248}
                        var v30 uint32 = 42
                        _ = v29
                        _ = v30
                    } else {
                        v31 := v1[-9223372036854775808]
                        switch "hello" {
                        case "":
                            v32 := GenS5{F0: -900, F1: 42}
                            _ = v32
                        case "\n\t":
                            _ = len(v1)
                        case "foo bar":
                            var v33 int16 = 32767
                            _ = v33
                        default:
                            var v34 int64 = 0
                            _ = v34
                        }
                        _ = v31
                    }
                    v35 := GenS9{F0: 18, F1: (("" + ("" + ("a" + "\n\t"))) + "a"), F2: 32767}
                    fmt.Println(fi25, v3)
                    _ = fi28
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 bool = false
                var v38 string = ("" + "foo bar")
                var v39 uint8 = 156
                v40 := GenS2{F0: (!v37)}
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi25
        }
    case -9223372036854775808:
        for fi41 := 0; fi41 < 17; fi41++ {
            v42 := GenS2{F0: (32768 == fi41)}
            if len(v24) > 0 {
                v24[0] = fi41
            }
            var v43 float32 = (-2.25 * (1e10 + (float32((3.14159 - 1e10)) / 1.5)))
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        switch (fi44 >> ((2) & 63)) {
        case -735:
            v3 -= v3
        case 0:
            v45 := GenS10{F0: false, F1: false}
            _ = v45
        case -9223372036854775808:
            if true {
                v46 := GenS7{F0: 255, F1: 0, F2: -61}
                for fi47 := 0; fi47 < 4; fi47++ {
                    if true {
                        var v48 uint32 = 2147483648
                        switch (-(fi47)) {
                        case 32767:
                            v48 = v48
                        case 127:
                            v49 := GenS11{F0: 123, F1: 713}
                            v3 += 87
                            var v50 int8 = (v3 / ((-128) | 1))
                            _ = v49
                            _ = v50
                        case 42:
                            fmt.Println(v3, fi47)
                        }
                        switch -587 {
                        case -9223372036854775808:
                            if true {
                                if true {
                                    if false {
                                        v48 -= (v48 >> ((296) & 63))
                                        var v51 float64 = 0.1
                                        _ = v51
                                    } else {
                                        fmt.Println(fi47, fi47)
                                        var v52 uint8 = 14
                                        var v53 string = ""
                                        _ = v52
                                        _ = v53
                                    }
                                } else {
                                    if (true && (v48 <= (v48 & (v48 - v48)))) {
                                        fmt.Println(v48)
                                        var v54 string = ("hello" + "a")
                                        switch ((v54 + ((v54 + (v54 + v54)) + "\n\t")) + ("" + "hello")) {
                                        case "hello":
                                            v55 := GenS12{F0: 1, F1: 1}
                                            _ = v55
                                        case "":
                                            switch fi44 {
                                            case -9223372036854775808:
                                                fmt.Println(fi44, v3, fi44)
                                            default:
                                                v24 = append(v24, fi47)
                                            }
                                        default:
                                            v3 = (v3 >> ((0) & 63))
                                            v48 = 2
                                            if true {
                                                switch fi47 {
                                                case -363:
                                                    switch fi47 {
                                                    case -1:
                                                        v54 = ((v54 + (v54 + v54)) + ((v54 + v54) + (("hello" + v54) + ("a" + v54))))
                                                    case 32767:
                                                        var v56 int16 = 32767
                                                        _ = v56
                                                    case 9223372036854775807:
                                                        v48 += (v48 << ((2) & 63))
                                                    default:
                                                        var v57 uint = 18446744073709551615
                                                        _ = v57
                                                    }
                                                case 852:
                                                    if true {
                                                        var v58 uint32 = (v48 % (((v48 * v48)) | 1))
                                                        v58 -= 42
                                                        if true {
                                                            v59 := GenS6{F0: v58, F1: (fi44 <= (fi47 - -1))}
                                                            _ = v59
                                                        } else {
                                                            v60 := GenS2{F0: (float32(((0.0 / 3.14159) / 1.0)) == float32(float64(3.14159)))}
                                                            v61 := GenS6{F0: (v48 / (((v58 | (^(v48)))) | 1)), F1: ((32767 != -267) || true)}
                                                            delete(v1, -2147483648)
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        _ = v58
                                                    }
                                                case 2147483647:
                                                    v62 := GenS7{F0: 42, F1: -2147483648, F2: (v3 ^ 9)}
                                                    v54 = "hello"
                                                    v63 := GenS13{F0: -1}
                                                    _ = v62
                                                    _ = v63
                                                }
                                                if true {
                                                    if false {
                                                        if len(v24) > 0 {
                                                            v24[0] = (^(fi47))
                                                        }
                                                        v48 += (v48 ^ v48)
                                                    } else {
                                                        if (65535 >= 2) {
                                                            v64 := GenS14{F0: -307}
                                                            var v65 int32 = genFn0(1, v3, 2)
                                                            v66 := GenS13{F0: 32767}
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        } else {
                                                            switch "" {
                                                            case "a":
                                                                v67 := GenS15{F0: 121, F1: 156}
                                                                fmt.Println(v3, fi44)
                                                                _ = v67
                                                            }
                                                            fmt.Println(v54, v3)
                                                        }
                                                    }
                                                    v68 := GenS16{F0: fi47, F1: 32767}
                                                    _ = v68
                                                } else {
                                                    var v69 int8 = (v3 - v3)
                                                    switch ((v54 + ("foo bar" + "a")) + v54) {
                                                    case "hello":
                                                        var v70 string = "hello"
                                                        _ = v70
                                                    case "foo bar":
                                                        if true {
                                                            if true {
                                                                if true {
                                                                    if true {
                                                                        switch ("hello" + "a") {
                                                                        case "a":
                                                                            fmt.Println(v54, v54)
                                                                        }
                                                                    } else {
                                                                        if (!((true && (true || false)) && false)) {
                                                                            v71 := GenS17{F0: (-(fi44)), F1: 0, F2: 9223372036854775807}
                                                                            var v72 uint16 = 0
                                                                            var v73 int8 = (v3 % (((^(v3))) | 1))
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        switch -532 {
                                                                        case -2147483648:
                                                                            var v74 uint32 = 785
                                                                            fmt.Println(fi47, v69)
                                                                            var v75 int8 = v69
                                                                            _ = v74
                                                                            _ = v75
                                                                        default:
                                                                            fmt.Println(fi47, v69)
                                                                            var v76 uint = 550
                                                                            switch fi47 {
                                                                            case 1:
                                                                                v24 = append(v24, (-(fi47)))
                                                                            case 2147483647:
                                                                                v77 := GenS8{F0: uint64(v48), F1: 154}
                                                                                _ = v77
                                                                            case 32768:
                                                                                v78 := v2[65535]
                                                                                _ = v78
                                                                            default:
                                                                                v79 := GenS2{F0: (!(42 >= -32768))}
                                                                                _ = v79
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                        var v80 string = (v54 + v54)
                                                                        _ = v80
                                                                    }
                                                                    v69 -= v3
                                                                }
                                                                if (42 != 127) {
                                                                    v24 = append(v24, 127)
                                                                } else {
                                                                    var v81 float64 = (((-2.25 / 1.0) / 0.1) + float64(float32(3.14159)))
                                                                    fmt.Println(v48, v81)
                                                                    _ = v81
                                                                }
                                                            } else {
                                                                var v82 int32 = 1
                                                                v48 += 4294967295
                                                                v69 -= v3
                                                                _ = v82
                                                            }
                                                            var v83 int64 = -670
                                                            if (0 == 0) {
                                                                if (32767 >= 42) {
                                                                    v69 = (^(v3))
                                                                    v84 := GenS12{F0: 9223372036854775807, F1: -1}
                                                                    if true {
                                                                        v48 += 294
                                                                        if false {
                                                                            var v85 float64 = 3.14159
                                                                            v69 -= -128
                                                                            v69 = v69
                                                                            _ = v85
                                                                        }
                                                                        if (!false) {
                                                                            fmt.Println(v3, v48, fi44)
                                                                        }
                                                                        switch 32768 {
                                                                        case -451:
                                                                            v83 = (v83 / ((v83) | 1))
                                                                        case 32767:
                                                                            var v86 uint16 = 32767
                                                                            _ = v86
                                                                        }
                                                                        if (false && true) {
                                                                            v3 = int8(v48)
                                                                            v48 = (v48 >> ((18446744073709551615) & 63))
                                                                            var v87 int8 = v3
                                                                            switch v54 {
                                                                            case "\n\t":
                                                                                v83 += (^(v83))
                                                                            case "foo bar":
                                                                                if ((!false) && false) {
                                                                                    var v88 int = 0
                                                                                    delete(v1, (fi44 + v88))
                                                                                    var v89 int8 = 127
                                                                                    var v90 uint = uint(v48)
                                                                                    v69 = -70
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                } else {
                                                                                    fmt.Println(v69, v54)
                                                                                    v91 := GenS7{F0: 255, F1: 353, F2: (-(v3))}
                                                                                    v83 = -9223372036854775808
                                                                                    fmt.Println(fi47, v48)
                                                                                    _ = v91
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v83, v87, v3)
                                                                            }
                                                                            _ = v87
                                                                        } else {
                                                                            switch v54 {
                                                                            case "a":
                                                                                v1[int(v69)] = fi44
                                                                            default:
                                                                                if false {
                                                                                    _ = len(v1)
                                                                                    if true {
                                                                                        var v92 int32 = 0
                                                                                        _ = v92
                                                                                    }
                                                                                }
                                                                            }
                                                                            v3 += 76
                                                                        }
                                                                    }
                                                                    _ = v84
                                                                } else {
                                                                    fmt.Println(fi44, fi47, v83)
                                                                    v93 := GenS13{F0: 0}
                                                                    if (42 < (fi44 * fi47)) {
                                                                        switch (("hello" + "hello") + "\n\t") {
                                                                        case "hello":
                                                                            v3 += (^(v3))
                                                                        case "":
                                                                            v94 := GenS10{F0: false, F1: (!true)}
                                                                            _ = v94
                                                                        case "a":
                                                                            v95 := GenS5{F0: 42, F1: (fi44 & (^(fi47)))}
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                    _ = v93
                                                                }
                                                                var v96 string = (v54 + v54)
                                                                v97 := GenS10{F0: true, F1: (78 <= 884)}
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            if false {
                                                                if ("foo bar" != "a") {
                                                                    var v98 uint64 = 1
                                                                    if (genFn0(-9223372036854775808, v3, 81) != 894) {
                                                                        switch v54 {
                                                                        case "\n\t":
                                                                            if false {
                                                                                if true {
                                                                                    switch v54 {
                                                                                    case "foo bar":
                                                                                        v48 += v48
                                                                                    case "a":
                                                                                        var v99 int8 = (-(v69))
                                                                                        _ = len(v1)
                                                                                        var v100 int64 = v83
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    default:
                                                                                        var v101 string = "hello"
                                                                                        v83 = v83
                                                                                        _ = v101
                                                                                    }
                                                                                    switch v54 {
                                                                                    case "hello":
                                                                                        var v102 uint32 = 0
                                                                                        _ = v102
                                                                                    case "":
                                                                                        fmt.Println(v83)
                                                                                    case "a":
                                                                                        var v103 uint8 = 1
                                                                                        _ = v103
                                                                                    }
                                                                                    v104 := GenS3{F0: (v69 ^ v69), F1: v54, F2: v98}
                                                                                    if len(v24) > 0 {
                                                                                        v24[0] = fi44
                                                                                    }
                                                                                    _ = v104
                                                                                } else {
                                                                                    if true {
                                                                                        var v105 int = -299
                                                                                        _ = v105
                                                                                    } else {
                                                                                        _ = len(v24)
                                                                                        if false {
                                                                                            v106 := GenS16{F0: 32767, F1: 576}
                                                                                            if (!(!true)) {
                                                                                                var v107 int16 = -1
                                                                                                _ = v107
                                                                                            }
                                                                                            _ = v106
                                                                                        }
                                                                                        var v108 int32 = 2147483647
                                                                                        _ = v108
                                                                                    }
                                                                                    v3 -= (v3 >> ((255) & 63))
                                                                                }
                                                                                if len(v24) > 0 {
                                                                                    v24[0] = 32768
                                                                                }
                                                                                var v109 int8 = (v69 - (v3 + -128))
                                                                                _ = v109
                                                                            }
                                                                            var v110 uint16 = 0
                                                                            _ = len(v1)
                                                                            _ = v110
                                                                        }
                                                                        var v111 float32 = float32((0.0 + (-2.25 * 0.0)))
                                                                        _ = v111
                                                                    } else {
                                                                        v83 = (-(v83))
                                                                        switch v54 {
                                                                        case "hello":
                                                                            if len(v24) >= 2 {
                                                                                v24 = v24[1:len(v24)]
                                                                            }
                                                                        case "foo bar":
                                                                            v112 := GenS8{F0: (v98 % ((0) | 1)), F1: 2}
                                                                            _ = v112
                                                                        case "\n\t":
                                                                            v113 := GenS18{F0: (v48 - (v48 + 561)), F1: (^(v3)), F2: true}
                                                                            _ = v113
                                                                        default:
                                                                            var v114 uint = uint(v98)
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                    v3 += v3
                                                                    _ = v98
                                                                } else {
                                                                    v54 = ""
                                                                    v54 = ((v54 + v54) + (v54 + ""))
                                                                    v69 -= (v69 + v69)
                                                                }
                                                            }
                                                            _ = v83
                                                        } else {
                                                            v115 := GenS19{F0: (fi47 - 32768), F1: 42}
                                                            _ = v115
                                                        }
                                                    case "":
                                                        var v116 string = (("a" + "") + v54)
                                                        _ = v116
                                                    }
                                                    fmt.Println(v48, v48, fi47)
                                                    _ = v69
                                                }
                                            }
                                        }
                                        _ = v54
                                    }
                                }
                                fmt.Println(v48)
                                var v117 uint = 715
                                var v118 uint32 = (v48 | 2147483648)
                                _ = v117
                                _ = v118
                            }
                        case 0:
                            v119 := GenS20{F0: 1, F1: 1, F2: (fi47 << ((288) & 63))}
                            _ = v119
                        }
                        v3 = (v3 & (-(v3)))
                        if true {
                            if (-647 >= 844) {
                                if false {
                                    fmt.Println(fi47, v3, v48)
                                    var v120 uint64 = 303
                                    _ = len(v24)
                                    var v121 uint16 = 255
                                    var v122 uint64 = (v120 - v120)
                                    v48 = (^(v48))
                                    _ = v120
                                    _ = v121
                                    _ = v122
                                } else {
                                    delete(v1, (fi47 & fi44))
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        var v123 uint16 = 32767
                                        _ = v123
                                    case "hello":
                                        var v124 string = "\n\t"
                                        if len(v24) > 0 {
                                            v24[0] = 128
                                        }
                                        switch 99 {
                                        case 32768:
                                            v48 += (^(v48))
                                        }
                                        _ = v124
                                    case "foo bar":
                                        var v125 bool = (float32((0.0 - float64(3.14159))) == 1.5)
                                        _ = v125
                                    default:
                                        v126 := GenS6{F0: (v48 | (v48 - v48)), F1: false}
                                        _ = v126
                                    }
                                }
                                fmt.Println(fi47)
                            }
                            switch "\n\t" {
                            case "a":
                                v127 := GenS21{F0: uint8(v48), F1: 65535, F2: 305}
                                if false {
                                    fmt.Println(v3)
                                    v48 = 532
                                }
                                _ = v127
                            case "foo bar":
                                if ((fi44 & fi44) != (fi44 | -110)) {
                                    v128 := GenS19{F0: (fi47 - fi47), F1: 630}
                                    if true {
                                        v129 := GenS22{F0: 434, F1: (v3 | 100), F2: -9223372036854775808}
                                        v130 := GenS17{F0: (fi47 % (((fi44 << ((988) & 63))) | 1)), F1: 32767, F2: 42}
                                        _ = v129
                                        _ = v130
                                    } else {
                                        var v131 float32 = 1e10
                                        _ = v131
                                    }
                                    _ = v128
                                }
                            case "hello":
                                v132 := GenS19{F0: int(v48), F1: 2}
                                _ = v132
                            default:
                                v3 += v3
                            }
                        } else {
                            if true {
                                fmt.Println(v3, fi44, v3)
                                if (884 >= 4294967295) {
                                    v48 = v48
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    v48 += 705
                                }
                                switch (fi44 | (fi44 >> ((746) & 63))) {
                                case -39:
                                    switch "a" {
                                    case "\n\t":
                                        switch ("hello" + "a") {
                                        case "":
                                            v24 = append(v24, (^(fi47)))
                                        case "a":
                                            v24 = append(v24, (fi44 >> ((4294967295) & 63)))
                                        case "hello":
                                            switch "\n\t" {
                                            case "hello":
                                                v3 += 127
                                            case "":
                                                if false {
                                                    v133 := GenS23{F0: 2, F1: -1}
                                                    _ = v133
                                                }
                                            }
                                        }
                                    case "hello":
                                        if (0 >= 2) {
                                            var v134 float32 = (((1.5 / (1.5 + -2.25)) * -2.25) / -2.25)
                                            switch (fi47 >> ((4294967295) & 63)) {
                                            case 0:
                                                var v135 float64 = 0.1
                                                _ = v135
                                            case -2147483648:
                                                v136 := v1[(^(fi47))]
                                                _ = v136
                                            }
                                            var v137 float64 = (3.14159 + 1e10)
                                            _ = v134
                                            _ = v137
                                        } else {
                                            v3 += v3
                                            var v138 uint = 18446744073709551615
                                            _ = v138
                                        }
                                    default:
                                        if false {
                                            var v139 float64 = 0.0
                                            if false {
                                                v140 := GenS20{F0: 42, F1: v3, F2: (fi47 & fi44)}
                                                var v141 int = 1
                                                if true {
                                                    switch -2147483648 {
                                                    case -2147483648:
                                                        fmt.Println(v48, v3, v141)
                                                    case 32767:
                                                        fmt.Println(fi44)
                                                    case 128:
                                                        v3 = (v3 * 0)
                                                    default:
                                                        if true {
                                                            v3 = -128
                                                            if len(v24) > 0 {
                                                                v24[0] = (fi47 >> ((18446744073709551615) & 63))
                                                            }
                                                            _ = len(v1)
                                                        } else {
                                                            var v142 int = fi47
                                                            if len(v24) >= 2 {
                                                                v24 = v24[1:len(v24)]
                                                            }
                                                            v141 = v141
                                                            _ = v142
                                                        }
                                                    }
                                                }
                                                if (2 <= 140) {
                                                    v139 = float64(0.0)
                                                }
                                                v3 += 0
                                                fmt.Println(fi47, fi44, v48)
                                                _ = v140
                                                _ = v141
                                            }
                                            _ = v139
                                        }
                                    }
                                    switch fi44 {
                                    case -47:
                                        delete(v1, (fi44 ^ -2147483648))
                                    case 32767:
                                        v48 = v48
                                    }
                                    v48 -= v48
                                default:
                                    if (v48 == v48) {
                                        v143 := GenS13{F0: 0}
                                        switch 9223372036854775807 {
                                        case -9223372036854775808:
                                            v2[v48] = ("foo bar" + ("\n\t" + "hello"))
                                        default:
                                            v144 := GenS24{F0: uint64(v48), F1: false, F2: (v3 < v3)}
                                            _ = v144
                                        }
                                        _ = v143
                                    } else {
                                        v1[fi44] = 127
                                        var v145 float32 = (float32(((-2.25 * float64(((3.14159 + (3.14159 + 0.0)) / (0.0 + -2.25)))) - 3.14159)) / 1.5)
                                        _ = v145
                                    }
                                }
                            }
                            var v146 uint32 = v48
                            if (!true) {
                                var v147 int16 = 42
                                v148 := GenS12{F0: 1, F1: v147}
                                var v149 uint16 = 0
                                _ = v147
                                _ = v148
                                _ = v149
                            } else {
                                v150 := GenS15{F0: 255, F1: 42}
                                fmt.Println(fi47, v3, fi47)
                                _ = v150
                            }
                            _ = v146
                        }
                        v151 := GenS24{F0: 1, F1: false, F2: (float64(-2.25) >= 0.0)}
                        _ = v151
                        _ = v48
                    }
                    _ = fi47
                }
                _ = v46
            } else {
                v152 := GenS21{F0: 82, F1: 1, F2: 0}
                for fi153 := 0; fi153 < 1; fi153++ {
                    fmt.Println(v3, fi44, v3)
                    v3 = (v3 ^ (v3 << ((918) & 63)))
                    _ = fi153
                }
                _ = v152
            }
        }
        v154 := GenS25{F0: (v3 & (v3 * 42))}
        for fi155 := 0; fi155 < 1; fi155++ {
            if (true || true) {
                v156 := GenS12{F0: -9223372036854775808, F1: -644}
                switch "\n\t" {
                case "\n\t":
                    var v157 uint8 = 127
                    _ = v157
                case "a":
                    v158 := GenS26{F0: true}
                    _ = v158
                }
                v159 := v2[65535]
                v1[fi155] = (fi155 - fi155)
                if ((v159 + (v159 + v159)) == v159) {
                    v160 := GenS9{F0: 48, F1: v159, F2: 32767}
                    v24 = append(v24, (fi44 / (((fi155 - fi155)) | 1)))
                    if true {
                        var v161 uint64 = 2
                        if false {
                            v3 = (-(v3))
                        }
                        v159 = v159
                        switch (v159 + (v159 + ("" + ("hello" + (v159 + "hello"))))) {
                        case "":
                            if false {
                                var v162 float64 = 0.1
                                var v163 uint8 = 0
                                if (42 <= 692) {
                                    var v164 int = (fi44 * (fi44 | int(v163)))
                                    var v165 int16 = int16(v161)
                                    _ = v164
                                    _ = v165
                                } else {
                                    delete(v1, fi44)
                                }
                                _ = v162
                                _ = v163
                            }
                        case "a":
                            var v166 int8 = 100
                            _ = v166
                        default:
                            var v167 int16 = -889
                            _ = v167
                        }
                        _ = v161
                    } else {
                        v3 = v3
                        var v168 uint8 = 2
                        _ = v168
                    }
                    _ = v160
                }
                _ = v156
                _ = v159
            }
            if false {
                v3 = v3
                v3 += v3
                if (2 == 4294967295) {
                    v3 = (v3 ^ (-(v3)))
                    if true {
                        v2[0] = ""
                        var v169 bool = (!(32767 == 42))
                        var v170 float64 = -2.25
                        _ = v169
                        _ = v170
                    } else {
                        if true {
                            v171 := GenS23{F0: 40, F1: -2147483648}
                            _ = v171
                        } else {
                            fmt.Println(v3, v3)
                            var v172 uint = 2
                            if (false || ((false || false) && true)) {
                                var v173 int8 = v3
                                _ = len(v1)
                                _ = v173
                            }
                            _ = v172
                        }
                        if false {
                            if true {
                                v24 = append(v24, fi44)
                                v174 := GenS27{F0: 1, F1: 65535, F2: 663}
                                v3 -= v3
                                var v175 int8 = (-(v3))
                                v3 += v175
                                var v176 float32 = 0.0
                                fmt.Println(v3)
                                _ = v174
                                _ = v175
                                _ = v176
                            } else {
                                var v177 uint16 = 737
                                _ = v177
                            }
                            switch "hello" {
                            case "foo bar":
                                switch ("foo bar" + "a") {
                                case "hello":
                                    var v178 int64 = 9223372036854775807
                                    _ = v178
                                case "a":
                                    var v179 uint32 = 546
                                    _ = v179
                                case "foo bar":
                                    fmt.Println(v3)
                                }
                            case "a":
                                v3 = (v3 * (^(v3)))
                            }
                            v3 += v3
                            switch (fi44 + fi155) {
                            case 4294967295:
                                if true {
                                    v180 := GenS14{F0: 677}
                                    if false {
                                        _ = len(v1)
                                        v181 := GenS7{F0: 101, F1: 0, F2: 0}
                                        if len(v24) > 0 {
                                            v24[0] = fi155
                                        }
                                        _ = v181
                                    } else {
                                        v182 := GenS28{F0: 252}
                                        v3 = -1
                                        v3 = 127
                                        _ = v182
                                    }
                                    if false {
                                        fmt.Println(fi155, fi155, fi155)
                                        if false {
                                            var v183 float32 = float32(0.1)
                                            switch (fi44 & 1) {
                                            case 32767:
                                                if true {
                                                    fmt.Println(fi44, v183)
                                                    v183 = (v183 + v183)
                                                    v3 = -63
                                                }
                                            case -532:
                                                v184 := GenS29{F0: 65535, F1: (v3 + -61)}
                                                _ = v184
                                            case 1:
                                                var v185 uint8 = 122
                                                _ = v185
                                            default:
                                                if (473 < 125) {
                                                    fmt.Println(fi44, fi44, fi155)
                                                }
                                            }
                                            var v186 uint64 = 42
                                            _ = v183
                                            _ = v186
                                        }
                                        switch 945 {
                                        case -2147483648:
                                            v3 = v3
                                        case 9223372036854775807:
                                            v3 = 1
                                        case 2147483647:
                                            fmt.Println(fi44)
                                        default:
                                            v3 = (v3 % ((v3) | 1))
                                        }
                                    } else {
                                        var v187 int = (fi44 & fi44)
                                        v188 := GenS18{F0: 342, F1: (v3 + (v3 << ((144) & 63))), F2: false}
                                        switch ("\n\t" + "hello") {
                                        case "foo bar":
                                            v24 = append(v24, fi155)
                                        case "":
                                            if (false || (-1 == -9223372036854775808)) {
                                                v189 := GenS30{F0: (^(v187)), F1: 42, F2: 4294967295}
                                                switch v187 {
                                                case 2147483647:
                                                    v187 += v187
                                                case 1:
                                                    fmt.Println(v3, fi155)
                                                default:
                                                    v190 := GenS12{F0: -1, F1: 317}
                                                    v191 := GenS2{F0: (false && (v3 > (v3 | (v3 - (-(v3))))))}
                                                    switch (("\n\t" + "\n\t") + "") {
                                                    case "":
                                                        v187 += fi44
                                                    case "foo bar":
                                                        delete(v2, 757)
                                                    default:
                                                        var v192 uint = 4294967295
                                                        _ = v192
                                                    }
                                                    _ = v190
                                                    _ = v191
                                                }
                                                if true {
                                                    switch fi155 {
                                                    case 4294967295:
                                                        v193 := GenS31{F0: true}
                                                        _ = v193
                                                    case 127:
                                                        var v194 int8 = (v3 - -1)
                                                        _ = v194
                                                    default:
                                                        v195 := GenS21{F0: 127, F1: 65535, F2: 0}
                                                        _ = v195
                                                    }
                                                    var v196 int = fi155
                                                    v197 := GenS32{F0: 1}
                                                    _ = v196
                                                    _ = v197
                                                } else {
                                                    v198 := GenS23{F0: 67, F1: 767}
                                                    _ = v198
                                                }
                                                if true {
                                                    switch v187 {
                                                    case 32768:
                                                        v3 = -1
                                                    case 94:
                                                        v199 := v1[fi155]
                                                        _ = v199
                                                    }
                                                    if false {
                                                        if (139 >= 0) {
                                                            v200 := GenS4{F0: 42, F1: -267}
                                                            _ = v200
                                                        }
                                                        v3 -= (v3 % ((v3) | 1))
                                                        if len(v24) > 0 {
                                                            v24[0] = v187
                                                        }
                                                    }
                                                } else {
                                                    v201 := GenS15{F0: 127, F1: 613}
                                                    var v202 float64 = (-2.25 * 1.5)
                                                    _ = v201
                                                    _ = v202
                                                }
                                                _ = v189
                                            } else {
                                                if false {
                                                    var v203 float32 = 3.14159
                                                    var v204 uint64 = 29
                                                    v204 = v204
                                                    _ = v203
                                                    _ = v204
                                                } else {
                                                    fmt.Println(fi44, v187)
                                                    v187 -= fi44
                                                }
                                                v3 -= (-(v3))
                                            }
                                            var v205 uint = 210
                                            v206 := v2[538]
                                            _ = v205
                                            _ = v206
                                        case "a":
                                            if (1 >= 65535) {
                                                if len(v24) > 0 {
                                                    v24[0] = (fi44 | (fi44 * (^(fi155))))
                                                }
                                                var v207 uint64 = 42
                                                _ = v207
                                            }
                                        }
                                        _ = v187
                                        _ = v188
                                    }
                                    _ = v180
                                }
                            }
                        } else {
                            var v208 uint32 = 2147483648
                            _ = v208
                        }
                    }
                } else {
                    v209 := GenS15{F0: 2, F1: 37}
                    switch (fi44 % ((fi155) | 1)) {
                    case -1:
                        var v210 float32 = 0.0
                        _ = v210
                    case 42:
                        if false {
                            v211 := GenS33{F0: 89, F1: 42}
                            if true {
                                fmt.Println(fi44, fi155)
                                var v212 int32 = -1
                                if (459 == 2) {
                                    delete(v2, 490)
                                }
                                _ = v212
                            }
                            _ = v211
                        } else {
                            v3 = v3
                            v213 := GenS34{F0: 2}
                            var v214 float64 = -2.25
                            _ = v213
                            _ = v214
                        }
                    }
                    _ = v209
                }
            }
            switch -2147483648 {
            case 4294967295:
                var v215 float32 = 0.1
                _ = v215
            default:
                fmt.Println(fi155)
            }
            v3 = v3
            v3 -= v3
            switch ("hello" + (("\n\t" + ("foo bar" + "\n\t")) + "")) {
            case "":
                v3 = -92
            case "\n\t":
                v216 := GenS10{F0: false, F1: (!(!(255 == 18446744073709551615)))}
                _ = v216
            }
            var v217 string = ""
            _ = fi155
            _ = v217
        }
        _ = fi44
        _ = v154
    }
    if true {
        var v218 bool = (!false)
        var v219 bool = (v218 && (-32768 == -32768))
        if v218 {
            if (!(float64(1e10) <= (1.5 + 1e10))) {
                fmt.Println(v3)
            }
            for fi220 := 0; fi220 < 3; fi220++ {
                for fi221 := 0; fi221 < 7; fi221++ {
                    if true {
                        v2[402] = ("a" + ("\n\t" + "a"))
                        switch "\n\t" {
                        case "a":
                            v3 = (v3 * -1)
                        case "":
                            var v222 uint8 = 60
                            _ = v222
                        }
                        var v223 uint64 = 117
                        v219 = true
                        v224 := GenS18{F0: 42, F1: (v3 * 5), F2: (9223372036854775807 < -1)}
                        v219 = v218
                        var v225 float64 = (3.14159 / 3.14159)
                        _ = v223
                        _ = v224
                        _ = v225
                    }
                    _ = fi221
                }
                _ = fi220
            }
            switch "a" {
            case "\n\t":
                var v226 int32 = 1
                _ = v226
            case "hello":
                for fi227 := 0; fi227 < 13; fi227++ {
                    v219 = false
                    fmt.Println(v219, fi227)
                    for fi228 := 0; fi228 < 17; fi228++ {
                        v218 = false
                        var v229 bool = false
                        _ = fi228
                        _ = v229
                    }
                    _ = fi227
                }
            case "foo bar":
                if v219 {
                    _ = len(v24)
                    for fi230 := 0; fi230 < 5; fi230++ {
                        switch "hello" {
                        case "\n\t":
                            v231 := GenS35{F0: 2, F1: 1, F2: 0}
                            _ = v231
                        case "foo bar":
                            for fi232 := 0; fi232 < 20; fi232++ {
                                fmt.Println(v218)
                                _ = fi232
                            }
                        case "hello":
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                        default:
                            if (true || (v218 && ((false || (true && v218)) && v219))) {
                                v3 = -76
                            }
                        }
                        _ = fi230
                    }
                    fmt.Println(v218, v219, v3)
                }
            }
        } else {
            var v233 bool = (253 < 42)
            var v234 uint64 = 42
            _ = v233
            _ = v234
        }
        _ = v218
        _ = v219
    } else {
        fmt.Println(v3)
        delete(v1, 35)
    }
    switch "" {
    case "\n\t":
        if (-766 >= 32767) {
            for fi235 := 0; fi235 < 2; fi235++ {
                v236 := GenS27{F0: 795, F1: 1, F2: -9223372036854775808}
                _ = len(v1)
                _ = fi235
                _ = v236
            }
            if true {
                for fi237 := 0; fi237 < 3; fi237++ {
                    v3 = -1
                    var v238 bool = (!(genFn0(-512, v3, 234) > genFn0(-1, (v3 << ((1) & 63)), 255)))
                    _ = fi237
                    _ = v238
                }
            } else {
                switch "hello" {
                case "foo bar":
                    var v239 int32 = 42
                    _ = v239
                case "hello":
                    for fi240 := 0; fi240 < 4; fi240++ {
                        for fi241 := 0; fi241 < 6; fi241++ {
                            var v242 uint8 = 255
                            var v243 uint = 18446744073709551615
                            var v244 bool = true
                            _ = fi241
                            _ = v242
                            _ = v243
                            _ = v244
                        }
                        _ = fi240
                    }
                case "a":
                    if false {
                        v3 -= (v3 | (-(v3)))
                    }
                }
                v245 := GenS6{F0: 2, F1: (171 > 1)}
                _ = v245
            }
            var v246 string = ("hello" + "hello")
            _ = v246
        } else {
            if true {
                for fi247 := 0; fi247 < 3; fi247++ {
                    v3 += (v3 / ((1) | 1))
                    v3 = 123
                    _ = fi247
                }
                for fi248 := 0; fi248 < 3; fi248++ {
                    switch (fi248 % ((fi248) | 1)) {
                    case 0:
                        v3 = (v3 << ((18446744073709551615) & 63))
                    default:
                        if true {
                            switch fi248 {
                            case -828:
                                var v249 uint32 = 938
                                _ = v249
                            default:
                                for fi250 := 0; fi250 < 14; fi250++ {
                                    fmt.Println(v3)
                                    v3 += (v3 | (v3 | v3))
                                    _ = fi250
                                }
                            }
                            if (255 == 42) {
                                _ = len(v2)
                                for fi251 := 0; fi251 < 19; fi251++ {
                                    if false {
                                        v252 := GenS36{F0: (true && ((true && true) || true)), F1: fi248, F2: (v3 * (v3 ^ v3))}
                                        _ = v252
                                    } else {
                                        var v253 uint32 = 2147483647
                                        _ = v253
                                    }
                                    _ = fi251
                                }
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                            } else {
                                v3 += (v3 << ((1) & 63))
                                if true {
                                    switch fi248 {
                                    case 924:
                                        switch 32768 {
                                        case 42:
                                            v24 = append(v24, fi248)
                                        case -214:
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        }
                                    case 32768:
                                        if (fi248 > fi248) {
                                            var v254 uint64 = 4294967295
                                            v3 -= v3
                                            var v255 int32 = -2147483648
                                            _ = v254
                                            _ = v255
                                        }
                                    case -329:
                                        switch (fi248 & (^(fi248))) {
                                        case 2147483647:
                                            var v256 float32 = -2.25
                                            _ = v256
                                        case 741:
                                            for fi257 := 0; fi257 < 4; fi257++ {
                                                fmt.Println(fi257)
                                                var v258 int = fi248
                                                v258 += v258
                                                _ = fi257
                                                _ = v258
                                            }
                                        }
                                    default:
                                        switch 32767 {
                                        case 1:
                                            fmt.Println(fi248)
                                        default:
                                            _ = len(v2)
                                        }
                                    }
                                    for fi259 := 0; fi259 < 11; fi259++ {
                                        if (fi259 < (fi259 >> ((1) & 63))) {
                                            v3 -= v3
                                            if false {
                                                var v260 uint16 = 65535
                                                _ = v260
                                            } else {
                                                v3 = v3
                                            }
                                            v261 := GenS19{F0: 679, F1: 42}
                                            _ = v261
                                        } else {
                                            var v262 float64 = 3.14159
                                            v263 := GenS37{F0: 617, F1: -2147483648}
                                            _ = v262
                                            _ = v263
                                        }
                                        var v264 uint32 = 2147483647
                                        v265 := v1[fi259]
                                        _ = fi259
                                        _ = v264
                                        _ = v265
                                    }
                                } else {
                                    delete(v1, fi248)
                                }
                            }
                            switch "hello" {
                            case "a":
                                v3 = -128
                            case "":
                                if false {
                                    switch ("a" + ("" + ("hello" + ""))) {
                                    case "a":
                                        fmt.Println(v3)
                                    }
                                    v266 := GenS19{F0: fi248, F1: 4294967295}
                                    _ = v266
                                }
                            }
                        } else {
                            v3 += v3
                            fmt.Println(fi248)
                        }
                    }
                    _ = fi248
                }
            }
        }
    case "foo bar":
        v267 := GenS38{F0: 0, F1: 919, F2: -1}
        _ = v267
    default:
        v268 := GenS39{F0: 2}
        _ = v268
    }
    switch "" {
    case "\n\t":
        switch 394 {
        case -9223372036854775808:
            switch -9223372036854775808 {
            case 32768:
                _ = len(v24)
            case -451:
                if (0 == 4) {
                    var v269 int64 = -725
                    fmt.Println(v269, v3)
                    switch "" {
                    case "a":
                        switch 32767 {
                        case 605:
                            v24 = append(v24, 42)
                        case 1:
                            var v270 uint32 = 65535
                            _ = v270
                        default:
                            switch 42 {
                            case -861:
                                if false {
                                    for fi271 := 0; fi271 < 15; fi271++ {
                                        v272 := v1[fi271]
                                        _ = fi271
                                        _ = v272
                                    }
                                    v3 += (v3 / ((-105) | 1))
                                    v273 := GenS10{F0: true, F1: (386 > 754)}
                                    _ = v273
                                }
                            case -685:
                                for fi274 := 0; fi274 < 18; fi274++ {
                                    var v275 float32 = (float32(0.1) - 1.5)
                                    v3 += (v3 + -53)
                                    _ = fi274
                                    _ = v275
                                }
                            default:
                                var v276 uint = 42
                                _ = v276
                            }
                        }
                    case "hello":
                        v277 := GenS40{F0: -32768, F1: (("\n\t" + "hello") >= ("\n\t" + ("foo bar" + ("foo bar" + "\n\t")))), F2: 2147483647}
                        _ = v277
                    }
                    _ = v269
                } else {
                    for fi278 := 0; fi278 < 17; fi278++ {
                        delete(v1, (^(fi278)))
                        for fi279 := 0; fi279 < 20; fi279++ {
                            v3 = v3
                            fmt.Println(fi278, v3)
                            v280 := GenS40{F0: 42, F1: (!(((!((fi279 + (fi279 ^ (-(fi278)))) > fi278)) && (false && (9223372036854775807 == 9223372036854775807))) || ((v3 | v3) == v3))), F2: (fi278 + fi279)}
                            if (2 <= 124) {
                                v281 := GenS41{F0: 0, F1: -32768}
                                _ = v281
                            }
                            _ = fi279
                            _ = v280
                        }
                        _ = fi278
                    }
                    v282 := GenS5{F0: 4294967295, F1: 90}
                    var v283 int8 = (v3 << ((470) & 63))
                    _ = v282
                    _ = v283
                }
            }
        case 32768:
            v3 -= 0
        default:
            v2[65535] = ("\n\t" + "foo bar")
        }
    case "foo bar":
        v3 += v3
    default:
        for fi284 := 0; fi284 < 3; fi284++ {
            for fi285 := 0; fi285 < 12; fi285++ {
                v1[(fi285 * 4294967295)] = 9223372036854775807
                var v286 uint32 = 469
                switch "foo bar" {
                case "hello":
                    v287 := GenS37{F0: 1, F1: -636}
                    _ = v287
                case "a":
                    v286 += (v286 & (v286 & (v286 + (v286 ^ (^(v286))))))
                default:
                    v3 += v3
                }
                _ = fi285
                _ = v286
            }
            if true {
                var v288 uint32 = 322
                if true {
                    var v289 uint8 = 2
                    fmt.Println(v289)
                    _ = v289
                }
                _ = len(v1)
                _ = v288
            }
            fmt.Println(fi284)
            _ = fi284
        }
    }
    fmt.Println(v3)
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v8
}
