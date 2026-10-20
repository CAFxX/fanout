package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint16; F1 int8; F2 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS13 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS14 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 string; F1 int; F2 uint32 }
func genFn0() int16 {
    if (-1 >= -2147483648) {
        for fi6 := 0; fi6 < 2; fi6++ {
            if false {
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 int32 = 2147483647
                    var v9 uint64 = 819
                    var v10 uint16 = 1
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            }
            v11 := []int{}
            _ = fi6
            _ = v11
        }
        var v12 int = 1
        _ = v12
    } else {
        var v13 uint64 = 1
        v13 = v13
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 string = (("foo bar" + (("" + "") + ("" + "hello"))) + ("foo bar" + "a"))
        _ = fi14
        _ = v15
    }
    var v16 int8 = 31
    _ = v16
    return 544
}
func genFn1(p0 int8) int32 {
    p0 = 42
    switch -9223372036854775808 {
    case 1:
        var v17 uint64 = 526
        _ = v17
    case 32767:
        if (!(!false)) {
            if true {
                fmt.Println(p0)
                p0 = -38
            }
        }
    case 514:
        fmt.Println(p0)
    default:
        p0 = (-(p0))
    }
    p0 = p0
    p0 -= 0
    _ = p0
    return 0
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{4294967295, 0}
    v3 := make(map[string]uint32)
    v4 := []int{-9223372036854775808, 32767, 297}
    v5 := []uint32{2147483647, 42, 4294967295}
    if true {
        if (false && false) {
            switch (v1 + "\n\t") {
            case "a":
                fmt.Println(v1)
            case "":
                fmt.Println(v1)
            }
        } else {
            v1 = ("hello" + ("a" + "foo bar"))
        }
        var v18 uint8 = 184
        _ = v18
    } else {
        switch 4294967295 {
        case 690:
            var v19 string = (v1 + v1)
            _ = v19
        }
        switch 736 {
        case 373:
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint64 = 135
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if true {
                    switch v1 {
                    case "a":
                        v1 = "a"
                    case "":
                        fmt.Println(fi20)
                    default:
                        for fi22 := 0; fi22 < 13; fi22++ {
                            if (255 > 255) {
                                var v23 int32 = -1
                                var v24 float64 = 1.5
                                var v25 string = v1
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                if false {
                                    if (255 >= 32767) {
                                        v4 = append(v4, fi20)
                                        if true {
                                            if true {
                                                v26 := GenS1{F0: 1, F1: 660, F2: 42}
                                                var v27 uint64 = 42
                                                var v28 uint8 = 2
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                            if (!(255 <= 994)) {
                                                var v29 uint64 = v21
                                                switch v1 {
                                                case "":
                                                    switch 942 {
                                                    case 2147483647:
                                                        var v30 uint32 = 653
                                                        _ = len(v3)
                                                        _ = v30
                                                    case 42:
                                                        v31 := GenS1{F0: 189, F1: 65535, F2: 65535}
                                                        _ = v31
                                                    case 895:
                                                        switch "hello" {
                                                        case "":
                                                            fmt.Println(fi20, fi22, fi22)
                                                            v32 := GenS2{F0: 834, F1: 32767, F2: (fi22 - fi20)}
                                                            var v33 int = -9223372036854775808
                                                            _ = v32
                                                            _ = v33
                                                        case "hello":
                                                            v1 = "a"
                                                        case "\n\t":
                                                            var v34 bool = (!true)
                                                            v5 = append(v5, 0)
                                                            _ = v34
                                                        default:
                                                            var v35 bool = true
                                                            _ = v35
                                                        }
                                                        if (false || false) {
                                                            v29 -= (^(v21))
                                                            v1 = v1
                                                            v1 = v1
                                                        }
                                                        v21 += v21
                                                    default:
                                                        v36 := GenS3{F0: ((-1 > 973) && true), F1: 42}
                                                        _ = v36
                                                    }
                                                default:
                                                    v37 := GenS2{F0: 255, F1: fi22, F2: -2147483648}
                                                    _ = v37
                                                }
                                                fmt.Println(fi20)
                                                _ = v29
                                            }
                                            var v38 int16 = -171
                                            switch "foo bar" {
                                            case "\n\t":
                                                if len(v5) > 0 {
                                                    v5[0] = 1
                                                }
                                            case "foo bar":
                                                var v39 string = v1
                                                _ = v39
                                            }
                                            switch "" {
                                            case "a":
                                                v5 = append(v5, 42)
                                            default:
                                                fmt.Println(v21, v1, fi22)
                                            }
                                            var v40 float32 = 3.14159
                                            _ = v38
                                            _ = v40
                                        } else {
                                            if ((1.5 - -2.25) > (1.5 / 1.0)) {
                                                switch fi22 {
                                                case 128:
                                                    var v41 uint16 = 2
                                                    _ = v41
                                                case 2147483647:
                                                    switch fi22 {
                                                    case -332:
                                                        var v42 int64 = 9223372036854775807
                                                        _ = v42
                                                    case 32768:
                                                        var v43 int = fi22
                                                        _ = v43
                                                    case 42:
                                                        var v44 string = "hello"
                                                        _ = v44
                                                    }
                                                default:
                                                    v45 := GenS4{F0: "\n\t", F1: 1, F2: (v21 * 169)}
                                                    _ = v45
                                                }
                                                var v46 uint16 = 959
                                                v21 += v21
                                                v47 := GenS1{F0: 1, F1: v46, F2: (v46 % ((970) | 1))}
                                                var v48 uint32 = 0
                                                var v49 string = v1
                                                v50 := GenS1{F0: 2147483647, F1: v46, F2: (^(v46))}
                                                v21 = uint64(v48)
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            } else {
                                                v21 = 2
                                            }
                                            if false {
                                                v1 = (v1 + "foo bar")
                                                v1 = v1
                                                var v51 int32 = int32(fi22)
                                                _ = v51
                                            } else {
                                                if (!true) {
                                                    v52 := GenS1{F0: 0, F1: 855, F2: 963}
                                                    v3[v1] = 42
                                                    _ = v52
                                                }
                                                v53 := v3["foo bar"]
                                                _ = v53
                                            }
                                            var v54 int64 = -325
                                            _ = v54
                                        }
                                    } else {
                                        var v55 int64 = 42
                                        v1 = v1
                                        if len(v4) > 0 {
                                            v4[0] = (fi20 | fi22)
                                        }
                                        switch fi22 {
                                        case -2147483648:
                                            v3["\n\t"] = 695
                                        case 128:
                                            v1 = v1
                                        case -1:
                                            v55 += 9223372036854775807
                                        default:
                                            if true {
                                                v21 = v21
                                            }
                                            v55 = (v55 >> ((1) & 63))
                                        }
                                        _ = v55
                                    }
                                    v1 = ""
                                    v56 := GenS1{F0: 706, F1: 2, F2: 42}
                                    _ = v56
                                } else {
                                    v57 := GenS5{F0: 65535, F1: 1, F2: 65535}
                                    _ = v57
                                }
                                fmt.Println(v1, fi20, v21)
                                v21 -= 18446744073709551615
                                if (true && ((genFn1(-1) == -313) || ((fi20 % ((fi22) | 1)) >= (fi20 * (fi22 + fi20))))) {
                                    switch (fi20 * -605) {
                                    case -1:
                                        v21 = v21
                                        fmt.Println(fi20, v21)
                                    case 32768:
                                        v21 = (v21 - 2)
                                    default:
                                        v58 := GenS4{F0: "foo bar", F1: v21, F2: (v21 + 4294967295)}
                                        _ = v58
                                    }
                                }
                            }
                            if (-32768 > -354) {
                                var v59 uint8 = 197
                                if (!(true && (v59 < (v59 % ((v59) | 1))))) {
                                    var v60 int32 = -2147483648
                                    v61 := GenS6{F0: 127, F1: 422}
                                    v1 = v1
                                    _ = v60
                                    _ = v61
                                }
                                v62 := GenS7{F0: 531, F1: 91, F2: v1}
                                _ = v59
                                _ = v62
                            } else {
                                if true {
                                    var v63 uint8 = 237
                                    _ = v63
                                } else {
                                    switch (fi22 & fi22) {
                                    case 1:
                                        var v64 uint16 = 247
                                        _ = v64
                                    case 0:
                                        v65 := GenS2{F0: 952, F1: fi22, F2: (fi22 | fi22)}
                                        _ = v65
                                    case 127:
                                        var v66 uint64 = 4294967295
                                        _ = v66
                                    }
                                }
                                switch -2147483648 {
                                case -108:
                                    fmt.Println(v1, v1, fi20)
                                case 127:
                                    if false {
                                        switch ("a" + (("foo bar" + (("\n\t" + v1) + v1)) + "hello")) {
                                        case "foo bar":
                                            v3[((v1 + v1) + "a")] = 42
                                        }
                                    }
                                case 4294967295:
                                    v67 := GenS8{F0: ((("hello" + v1) + "") + ("\n\t" + (v1 + (v1 + ""))))}
                                    _ = v67
                                }
                                var v68 uint8 = 2
                                _ = v68
                            }
                            _ = fi22
                        }
                    }
                    var v69 uint16 = 948
                    for fi70 := 0; fi70 < 2; fi70++ {
                        switch fi20 {
                        case 9223372036854775807:
                            var v71 int8 = 42
                            _ = v71
                        default:
                            v69 = (v69 + v69)
                            v69 -= v69
                            v1 = "\n\t"
                            fmt.Println(v69, fi20)
                            var v72 bool = true
                            _ = v72
                        }
                        v21 += (v21 ^ 161)
                        _ = fi70
                    }
                    _ = v69
                } else {
                    v1 = v1
                    var v73 float64 = (3.14159 - 0.1)
                    for fi74 := 0; fi74 < 11; fi74++ {
                        if true {
                            switch v1 {
                            case "foo bar":
                                v1 = v1
                            }
                            if (fi74 >= 127) {
                                fmt.Println(fi74, v73)
                                var v75 float32 = float32(((float64(-2.25) - v73) - 0.1))
                                v76 := GenS9{F0: 127}
                                _ = v75
                                _ = v76
                            } else {
                                var v77 uint = 1
                                if (-371 > 0) {
                                    fmt.Println(fi74, fi20)
                                } else {
                                    _ = len(v3)
                                    var v78 int8 = -128
                                    v79 := GenS2{F0: (v77 - (v77 ^ uint(v21))), F1: int(v78), F2: 0}
                                    _ = v78
                                    _ = v79
                                }
                                _ = v77
                            }
                            switch fi74 {
                            case 1:
                                fmt.Println(fi74, v73)
                            }
                        }
                        var v80 bool = false
                        fmt.Println(fi20, fi74)
                        _ = fi74
                        _ = v80
                    }
                    _ = v73
                }
                _ = fi20
                _ = v21
            }
        default:
            if true {
                switch (v1 + ("" + ("a" + ("hello" + v1)))) {
                case "":
                    v1 = "\n\t"
                case "a":
                    for fi81 := 0; fi81 < 6; fi81++ {
                        v82 := GenS6{F0: 42, F1: 1}
                        _ = fi81
                        _ = v82
                    }
                case "hello":
                    for fi83 := 0; fi83 < 1; fi83++ {
                        var v84 uint32 = 833
                        v84 -= 0
                        _ = fi83
                        _ = v84
                    }
                default:
                    if true {
                        if false {
                            v1 = (v1 + v1)
                            v85 := GenS10{F0: 1}
                            for fi86 := 0; fi86 < 7; fi86++ {
                                v1 = (("\n\t" + "foo bar") + ((v1 + v1) + v1))
                                if (3.14159 == -2.25) {
                                    v87 := GenS8{F0: ""}
                                    if (727 > 2147483648) {
                                        v1 = v1
                                    }
                                    switch v1 {
                                    case "":
                                        var v88 uint64 = 4294967295
                                        _ = v88
                                    case "hello":
                                        switch (v1 + "a") {
                                        case "foo bar":
                                            _ = len(v3)
                                        default:
                                            switch v1 {
                                            case "hello":
                                                if (!true) {
                                                    v1 = v1
                                                    for fi89 := 0; fi89 < 19; fi89++ {
                                                        v1 = v1
                                                        _ = fi89
                                                    }
                                                    var v90 float64 = (-2.25 + 3.14159)
                                                    _ = v90
                                                } else {
                                                    for fi91 := 0; fi91 < 11; fi91++ {
                                                        v92 := GenS11{F0: true, F1: 33}
                                                        _ = fi91
                                                        _ = v92
                                                    }
                                                }
                                            case "":
                                                fmt.Println(v1, v1)
                                            case "\n\t":
                                                if (2 >= 65535) {
                                                    if ((-1 <= -1) || false) {
                                                        v93 := GenS11{F0: false, F1: 7}
                                                        if false {
                                                            v1 = "\n\t"
                                                            if (278 == 1) {
                                                                for fi94 := 0; fi94 < 16; fi94++ {
                                                                    v95 := GenS7{F0: 42, F1: -1, F2: v1}
                                                                    if (true && true) {
                                                                        var v96 int16 = 1000
                                                                        v96 = (v96 >> ((2) & 63))
                                                                        v96 -= v96
                                                                        _ = v96
                                                                    }
                                                                    v1 = (((v1 + "a") + v1) + "")
                                                                    _ = fi94
                                                                    _ = v95
                                                                }
                                                                if (1e10 >= 1.5) {
                                                                    v1 = v1
                                                                    for fi97 := 0; fi97 < 16; fi97++ {
                                                                        fmt.Println(v1, v1, fi97)
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 65535
                                                                        }
                                                                        var v98 uint16 = 255
                                                                        if ((v98 - 255) >= (v98 / ((v98) | 1))) {
                                                                            var v99 bool = false
                                                                            if v99 {
                                                                                switch (fi97 | int(v98)) {
                                                                                case 42:
                                                                                    fmt.Println(fi97)
                                                                                }
                                                                            } else {
                                                                                if v99 {
                                                                                    if (!v99) {
                                                                                        if v99 {
                                                                                            v99 = (-128 != 127)
                                                                                            var v100 int = 1
                                                                                            _ = v100
                                                                                        } else {
                                                                                            v101 := GenS12{F0: 18446744073709551615, F1: (!(2147483648 == 0)), F2: 0}
                                                                                            _ = len(v5)
                                                                                            _ = v101
                                                                                        }
                                                                                        if (!false) {
                                                                                            var v102 uint8 = 30
                                                                                            if false {
                                                                                                v103 := GenS13{F0: 475, F1: 4294967295, F2: (v98 ^ uint16(v102))}
                                                                                                _ = len(v3)
                                                                                                var v104 int8 = 127
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            }
                                                                                            _ = v102
                                                                                        } else {
                                                                                            _ = len(v4)
                                                                                        }
                                                                                    } else {
                                                                                        var v105 string = (v1 + ((("a" + "") + v1) + "hello"))
                                                                                        if v99 {
                                                                                            v106 := GenS7{F0: v98, F1: -9, F2: "foo bar"}
                                                                                            _ = v106
                                                                                        } else {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (fi86 * fi86)
                                                                                            }
                                                                                            v107 := GenS11{F0: (!false), F1: 2}
                                                                                            _ = len(v2)
                                                                                            _ = v107
                                                                                        }
                                                                                        var v108 int32 = 42
                                                                                        if (!v99) {
                                                                                            var v109 int = fi86
                                                                                            _ = v109
                                                                                        } else {
                                                                                            v110 := GenS14{F0: 1, F1: 90, F2: 4294967295}
                                                                                            var v111 int32 = 2147483647
                                                                                            _ = v110
                                                                                            _ = v111
                                                                                        }
                                                                                        _ = v105
                                                                                        _ = v108
                                                                                    }
                                                                                    var v112 string = ((v1 + v1) + "\n\t")
                                                                                    _ = v112
                                                                                }
                                                                                fmt.Println(v99, v99, fi97)
                                                                            }
                                                                            _ = v99
                                                                        }
                                                                        var v113 float64 = 3.14159
                                                                        if false {
                                                                            fmt.Println(fi97, v98)
                                                                        } else {
                                                                            var v114 uint16 = (^(v98))
                                                                            if false {
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    v98 += v98
                                                                                case "a":
                                                                                    var v115 uint8 = 2
                                                                                    _ = v115
                                                                                }
                                                                                v116 := GenS7{F0: (^(v114)), F1: 0, F2: v1}
                                                                                _ = v116
                                                                            }
                                                                            fmt.Println(fi97, v1, fi97)
                                                                            _ = v114
                                                                        }
                                                                        var v117 int8 = 127
                                                                        if (v98 > (v98 >> ((2) & 63))) {
                                                                            var v118 uint16 = v98
                                                                            _ = len(v4)
                                                                            _ = v118
                                                                        } else {
                                                                            var v119 uint16 = v98
                                                                            switch ("a" + ("\n\t" + v1)) {
                                                                            case "a":
                                                                                if (!(!(((0.0 - 1e10) - 0.0) > float32((3.14159 + v113))))) {
                                                                                    var v120 uint = 199
                                                                                    if false {
                                                                                        var v121 float64 = 0.1
                                                                                        _ = v121
                                                                                    } else {
                                                                                        var v122 int32 = int32(fi86)
                                                                                        _ = v122
                                                                                    }
                                                                                    if true {
                                                                                        v3[(v1 + v1)] = 1
                                                                                        var v123 uint8 = 1
                                                                                        if false {
                                                                                            v124 := GenS15{F0: 0}
                                                                                            _ = v124
                                                                                        }
                                                                                        _ = v123
                                                                                    } else {
                                                                                        v117 = (v117 * -1)
                                                                                    }
                                                                                    _ = v120
                                                                                }
                                                                            case "hello":
                                                                                v98 += (v119 % ((v119) | 1))
                                                                            }
                                                                            _ = v119
                                                                        }
                                                                        _ = fi97
                                                                        _ = v113
                                                                        _ = v117
                                                                        _ = v98
                                                                    }
                                                                } else {
                                                                    v1 = ""
                                                                    v1 = v1
                                                                }
                                                            }
                                                        } else {
                                                            delete(v3, v1)
                                                            var v125 float32 = 1e10
                                                            _ = v125
                                                        }
                                                        for fi126 := 0; fi126 < 16; fi126++ {
                                                            v127 := GenS13{F0: -2147483648, F1: 1, F2: 42}
                                                            if (!false) {
                                                                var v128 bool = false
                                                                fmt.Println(v1)
                                                                _ = v128
                                                            }
                                                            _ = fi126
                                                            _ = v127
                                                        }
                                                        _ = v93
                                                    } else {
                                                        fmt.Println(v1, v1)
                                                    }
                                                    var v129 uint32 = 1
                                                    if (((fi86 | (fi86 - (fi86 >> ((380) & 63)))) < (fi86 & (fi86 / ((4294967295) | 1)))) || (false && false)) {
                                                        var v130 int32 = -660
                                                        _ = v130
                                                    } else {
                                                        var v131 int64 = -970
                                                        delete(v3, v1)
                                                        v132 := v3[v1]
                                                        _ = v131
                                                        _ = v132
                                                    }
                                                    _ = v129
                                                } else {
                                                    fmt.Println(v1)
                                                    _ = len(v2)
                                                }
                                            default:
                                                if (65535 >= 0) {
                                                    var v133 float32 = 0.0
                                                    fmt.Println(fi86, fi86)
                                                    var v134 uint32 = 42
                                                    _ = v133
                                                    _ = v134
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        v2 = append(v2, 2)
                                    default:
                                        v1 = (v1 + v1)
                                    }
                                    _ = v87
                                }
                                var v135 float64 = 1.5
                                _ = fi86
                                _ = v135
                            }
                            _ = v85
                        }
                        v1 = "hello"
                    } else {
                        fmt.Println(v1)
                    }
                }
                v1 = (((("" + ((((v1 + "a") + "a") + "hello") + v1)) + ("hello" + "hello")) + v1) + v1)
            }
        }
    }
    v1 = (v1 + ((v1 + ("foo bar" + ((v1 + (v1 + "\n\t")) + v1))) + ("foo bar" + ((v1 + "hello") + ("hello" + ("hello" + v1))))))
    var v136 uint8 = 2
    var v137 int32 = 42
    var v138 int8 = 0
    fmt.Println(v136, v136)
    fmt.Println(v1)
    var v139 uint64 = 18446744073709551615
    v140 := GenS16{F0: (v1 + ""), F1: 2147483647, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v136:", v136)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v139:", v139)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v136
    _ = v137
    _ = v138
    _ = v139
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
