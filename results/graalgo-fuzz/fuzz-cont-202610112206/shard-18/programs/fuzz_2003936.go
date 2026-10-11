package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 uint; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16; F1 int8 }
type GenS14 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS15 struct { F0 int; F1 uint; F2 int16 }
type GenS16 struct { F0 int; F1 int16 }
type GenS17 struct { F0 uint64; F1 int }
type GenS18 struct { F0 uint32; F1 string }
type GenS19 struct { F0 uint; F1 int16; F2 int }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 string = ("\n\t" + "\n\t")
    var v4 bool = false
    if false {
        fmt.Println(v3, v3)
    }
    v4 = (2 == 1)
    switch "\n\t" {
    case "a":
        if (32767 == -175) {
            fmt.Println(v3, v4)
            v3 = v3
        } else {
            if (173 <= 22) {
                for fi5 := 0; fi5 < 1; fi5++ {
                    if v4 {
                        v4 = true
                    } else {
                        v6 := GenS1{F0: 477, F1: 1, F2: fi5}
                        v4 = true
                        for fi7 := 0; fi7 < 7; fi7++ {
                            v8 := GenS1{F0: -1, F1: 39, F2: (fi7 % ((fi5) | 1))}
                            fmt.Println(fi5, fi5)
                            v3 = ((v3 + (v3 + "a")) + (((((v3 + ("hello" + (((v3 + v3) + v3) + "hello"))) + ((v3 + ((v3 + ("" + (((("\n\t" + ((v3 + "\n\t") + v3)) + "") + "a") + "foo bar"))) + v3)) + v3)) + (v3 + ((("a" + ("" + ((v3 + (v3 + "foo bar")) + ("" + "\n\t")))) + "") + (((v3 + v3) + "hello") + v3)))) + v3) + ((v3 + v3) + "a")))
                            _ = fi7
                            _ = v8
                        }
                        _ = v6
                    }
                    var v9 float64 = 3.14159
                    switch fi5 {
                    case 32767:
                        v3 = "\n\t"
                    default:
                        v9 = v9
                    }
                    _ = fi5
                    _ = v9
                }
            } else {
                _ = len(v1)
                if v4 {
                    v10 := []int{486, 724, 32767}
                    _ = v10
                }
                if (!(128 != 127)) {
                    switch 0 {
                    case 1:
                        v11 := []int{}
                        _ = v11
                    case 0:
                        v12 := GenS2{F0: 565, F1: v3}
                        _ = v12
                    default:
                        if true {
                            v4 = false
                            fmt.Println(v4)
                        } else {
                            var v13 float32 = (1.5 / -2.25)
                            if v4 {
                                v3 = v3
                                v4 = (42 <= 1000)
                            }
                            _ = v13
                        }
                    }
                    switch 32767 {
                    case 42:
                        v4 = false
                    }
                    var v14 int16 = 42
                    _ = v14
                } else {
                    v3 = v3
                }
            }
            var v15 int = 42
            var v16 uint64 = 4294967295
            _ = v15
            _ = v16
        }
    case "hello":
        delete(v1, 1)
    case "foo bar":
        v4 = (-1 > 42)
    default:
        _ = len(v1)
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int = (fi17 << ((18446744073709551615) & 63))
            switch fi18 {
            case -9223372036854775808:
                fmt.Println(v4, v3)
            default:
                var v20 uint8 = 255
                _ = v20
            }
            v19 -= (v19 | (v19 << ((0) & 63)))
            _ = fi18
            _ = v19
        }
        if v4 {
            v21 := []int{(fi17 | -1), -1}
            var v22 int = (fi17 - fi17)
            if (2 == 2) {
                v3 = v3
                v3 = "foo bar"
            } else {
                v23 := v2[1]
                var v24 uint64 = 1
                _ = v23
                _ = v24
            }
            _ = v21
            _ = v22
        }
        _ = fi17
    }
    if (2147483647 != 0) {
        v25 := GenS2{F0: 63, F1: v3}
        _ = v25
    }
    v4 = v4
    var v26 uint64 = 0
    if v4 {
        v26 += v26
        v27 := GenS1{F0: -261, F1: 0, F2: -584}
        _ = v27
    }
    switch 127 {
    case 32767:
        fmt.Println(v3, v4, v4)
    case 202:
        v28 := GenS3{F0: 779, F1: true}
        _ = v28
    case 833:
        if (!true) {
            for fi29 := 0; fi29 < 4; fi29++ {
                switch (fi29 / ((fi29) | 1)) {
                case 4294967295:
                    var v30 int32 = 42
                    _ = v30
                default:
                    var v31 uint16 = 855
                    _ = v31
                }
                v32 := v2[42]
                _ = fi29
                _ = v32
            }
        } else {
            v33 := GenS3{F0: 75, F1: ((v4 || v4) && v4)}
            switch 9223372036854775807 {
            case 127:
                v34 := []int{1}
                _ = v34
            case 9223372036854775807:
                var v35 float32 = 3.14159
                _ = v35
            case 42:
                for fi36 := 0; fi36 < 15; fi36++ {
                    v3 = (v3 + ((v3 + "") + ""))
                    for fi37 := 0; fi37 < 9; fi37++ {
                        switch "\n\t" {
                        case "hello":
                            v38 := GenS1{F0: 1000, F1: -57, F2: fi36}
                            _ = v38
                        case "":
                            v3 = v3
                        }
                        var v39 uint8 = 127
                        v40 := GenS1{F0: 121, F1: 42, F2: (fi36 << ((2) & 63))}
                        _ = fi37
                        _ = v39
                        _ = v40
                    }
                    _ = fi36
                }
            }
            _ = v33
        }
    default:
        var v41 int8 = -128
        _ = v41
    }
    var v42 uint8 = 1
    v4 = true
    fmt.Println(v42, v26, v3)
    if (70 != 0) {
        _ = len(v1)
        var v43 int64 = 0
        for fi44 := 0; fi44 < 13; fi44++ {
            if (42 > 995) {
                var v45 int32 = -1
                if ((v4 || false) || (uint(v42) != 4294967295)) {
                    delete(v1, 2)
                    switch (("a" + v3) + v3) {
                    case "\n\t":
                        switch 511 {
                        case -1:
                            v43 = 0
                        case 9223372036854775807:
                            var v46 int32 = v45
                            _ = v46
                        }
                    default:
                        if v4 {
                            for fi47 := 0; fi47 < 6; fi47++ {
                                if (v3 < ((v3 + "a") + "hello")) {
                                    v48 := []int{fi47}
                                    v49 := GenS3{F0: 2, F1: ((v26 | (^(v26))) != v26)}
                                    v50 := GenS1{F0: 42, F1: 127, F2: (^(fi47))}
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                } else {
                                    fmt.Println(v43)
                                }
                                v51 := GenS3{F0: 47, F1: (32767 != 341)}
                                _ = fi47
                                _ = v51
                            }
                            v52 := []int{fi44, (fi44 - (fi44 / (((fi44 | (fi44 + (fi44 ^ fi44)))) | 1))), (-(fi44))}
                            var v53 float32 = -2.25
                            _ = v52
                            _ = v53
                        } else {
                            var v54 float64 = 0.0
                            var v55 int = 604
                            _ = v54
                            _ = v55
                        }
                    }
                }
                _ = v45
            } else {
                var v56 int32 = 31
                v56 -= -505
                if (((!true) || (-370 < 32767)) && (int(v56) > fi44)) {
                    for fi57 := 0; fi57 < 7; fi57++ {
                        if (!(v4 && v4)) {
                            if v4 {
                                var v58 int16 = 1
                                if v4 {
                                    fmt.Println(v56)
                                    if (138 < 0) {
                                        fmt.Println(v4, fi44)
                                        fmt.Println(v43, v56)
                                        v59 := []int{}
                                        _ = v59
                                    }
                                    if v4 {
                                        fmt.Println(fi57, v43)
                                        v43 -= v43
                                    } else {
                                        if false {
                                            v58 = v58
                                            switch fi57 {
                                            case -9:
                                                switch ((((("\n\t" + "a") + "a") + "hello") + v3) + "a") {
                                                case "a":
                                                    v60 := []int{(fi44 | 42)}
                                                    _ = v60
                                                case "foo bar":
                                                    switch (fi57 * (fi44 / (((-(fi44))) | 1))) {
                                                    case 732:
                                                        var v61 float64 = 1e10
                                                        _ = v61
                                                    case -2147483648:
                                                        v62 := []int{(fi57 >> ((18446744073709551615) & 63)), -9223372036854775808, (-(fi57))}
                                                        if (!v4) {
                                                            v58 -= v58
                                                        }
                                                        fmt.Println(v43, v3, fi44)
                                                        _ = v62
                                                    case -1:
                                                        fmt.Println(v4)
                                                    }
                                                default:
                                                    if (v43 != (v43 >> ((65535) & 63))) {
                                                        var v63 float64 = 1.5
                                                        fmt.Println(v43, v58)
                                                        fmt.Println(v26)
                                                        _ = v63
                                                    } else {
                                                        if v4 {
                                                            var v64 uint64 = v26
                                                            _ = v64
                                                        } else {
                                                            v42 += 0
                                                        }
                                                        var v65 uint16 = 2
                                                        _ = v65
                                                    }
                                                }
                                            case -2147483648:
                                                var v66 float32 = ((1.5 - 0.1) - (float32(0.1) - 1e10))
                                                _ = v66
                                            case 9223372036854775807:
                                                if (false || (65535 != 255)) {
                                                    v3 = (("" + "") + v3)
                                                    switch int(v42) {
                                                    case 32768:
                                                        if (311 >= 287) {
                                                            v67 := GenS4{F0: 1, F1: 271}
                                                            _ = v67
                                                        } else {
                                                            if v4 {
                                                                if true {
                                                                    v58 -= int16(v56)
                                                                } else {
                                                                    if (!v4) {
                                                                        v56 = v56
                                                                    }
                                                                    var v68 uint32 = 269
                                                                    v69 := []int{(^(fi57)), (^(fi57)), fi57}
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                var v70 uint8 = v42
                                                                _ = v70
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        var v71 int16 = v58
                                                        _ = v71
                                                    default:
                                                        v72 := []int{(fi44 | fi57), -9223372036854775808, fi57}
                                                        v73 := GenS5{F0: (v58 ^ (v58 + 42))}
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                }
                                            }
                                            v3 = "\n\t"
                                        } else {
                                            if ((!(true && (v43 == (^(v43))))) || ((!v4) && false)) {
                                                _ = len(v1)
                                                delete(v2, 2)
                                            } else {
                                                v74 := v1[167]
                                                v56 -= v56
                                                _ = v74
                                            }
                                            if (v26 >= (v26 << ((uint(v26)) & 63))) {
                                                v1[1] = fi57
                                                var v75 float64 = 1.5
                                                switch 32767 {
                                                case 4294967295:
                                                    delete(v2, 2)
                                                    var v76 float64 = (1e10 - v75)
                                                    var v77 uint = 813
                                                    v78 := GenS1{F0: v58, F1: 0, F2: fi44}
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                case -274:
                                                    v56 -= v56
                                                default:
                                                    fmt.Println(v3, fi44, v75)
                                                }
                                                _ = v75
                                            } else {
                                                var v79 int64 = 935
                                                v80 := GenS5{F0: v58}
                                                v81 := GenS6{F0: 795}
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            }
                                        }
                                    }
                                    switch v3 {
                                    case "\n\t":
                                        switch v3 {
                                        case "a":
                                            switch ("" + v3) {
                                            case "\n\t":
                                                var v82 int16 = -582
                                                _ = v82
                                            case "hello":
                                                if (v4 || (3.14159 < (0.1 + (((1e10 / (0.1 - (((1e10 + (3.14159 / (3.14159 + 0.0))) * ((((0.1 * -2.25) - 3.14159) / (((0.0 / 3.14159) / 1.0) + (1e10 + 0.0))) - 0.1)) + (0.0 * 3.14159)))) / 1.0) - ((3.14159 + 3.14159) / 0.1))))) {
                                                    v42 += v42
                                                    v56 = v56
                                                }
                                            }
                                        case "hello":
                                            switch fi44 {
                                            case 180:
                                                var v83 int64 = -9223372036854775808
                                                _ = v83
                                            }
                                        case "":
                                            v42 -= 255
                                            v3 = ""
                                            v84 := GenS4{F0: v26, F1: 787}
                                            _ = v84
                                        default:
                                            var v85 int = -9223372036854775808
                                            _ = v85
                                        }
                                    case "foo bar":
                                        fmt.Println(v42, fi57, v56)
                                    case "hello":
                                        var v86 string = v3
                                        _ = v86
                                    }
                                    v87 := GenS1{F0: (v58 ^ v58), F1: -58, F2: (-(fi44))}
                                    _ = v87
                                } else {
                                    switch v3 {
                                    case "foo bar":
                                        fmt.Println(fi57, fi57)
                                    case "hello":
                                        if (!v4) {
                                            if (224 >= 18446744073709551615) {
                                                if ((v58 / (((v58 >> ((4294967295) & 63))) | 1)) == v58) {
                                                    v26 = 378
                                                    var v88 uint16 = 1
                                                    var v89 float64 = float64(0.0)
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    v90 := GenS4{F0: 2, F1: 4294967295}
                                                    _ = v90
                                                }
                                                var v91 uint32 = 817
                                                if true {
                                                    var v92 uint8 = (v42 - uint8(v91))
                                                    if ((32767 > v58) || ((42 > int(v42)) && false)) {
                                                        v92 -= (v92 << ((18446744073709551615) & 63))
                                                        var v93 int32 = (-(v56))
                                                        _ = v93
                                                    } else {
                                                        v3 = v3
                                                        var v94 bool = v4
                                                        _ = v94
                                                    }
                                                    v4 = false
                                                    _ = v92
                                                }
                                                _ = v91
                                            } else {
                                                var v95 float32 = -2.25
                                                v96 := v2[65535]
                                                fmt.Println(fi57, v56, v56)
                                                _ = v95
                                                _ = v96
                                            }
                                            if (!v4) {
                                                v97 := []int{9223372036854775807, fi44, int(v58)}
                                                _ = v97
                                            }
                                        } else {
                                            v98 := GenS3{F0: 255, F1: (0 <= 0)}
                                            fmt.Println(v43)
                                            fmt.Println(v43)
                                            _ = v98
                                        }
                                    case "a":
                                        switch (fi44 * (fi44 + fi44)) {
                                        case -2147483648:
                                            if false {
                                                if ((65535 < 667) && (0 != v26)) {
                                                    fmt.Println(v3)
                                                    if (!(v26 != v26)) {
                                                        var v99 float64 = 1e10
                                                        v100 := GenS7{F0: v3, F1: 926}
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        if ((v3 != "a") && (!false)) {
                                                            var v101 int16 = -662
                                                            v102 := GenS8{F0: 0, F1: 2, F2: v43}
                                                            var v103 uint64 = (v26 / (((v26 - (^(v26)))) | 1))
                                                            _ = v101
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        switch (fi44 ^ 42) {
                                                        case -2147483648:
                                                            v2[42] = fi44
                                                        case 725:
                                                            if ((v26 % ((1) | 1)) < v26) {
                                                                v2[65535] = fi44
                                                            }
                                                        case 32767:
                                                            fmt.Println(fi57, fi44, v42)
                                                        }
                                                        v56 = (v56 % ((v56) | 1))
                                                    }
                                                    var v104 float64 = (-2.25 + (((3.14159 - 1e10) - 0.0) - 0.0))
                                                    _ = v104
                                                }
                                                v105 := GenS5{F0: int16(fi57)}
                                                v42 = 128
                                                _ = v105
                                            } else {
                                                var v106 int8 = 1
                                                var v107 string = "a"
                                                var v108 float32 = 0.0
                                                _ = v106
                                                _ = v107
                                                _ = v108
                                            }
                                        case 0:
                                            v1[4294967295] = (fi44 ^ 9223372036854775807)
                                        case 32768:
                                            fmt.Println(v56, v4)
                                        default:
                                            v56 = v56
                                        }
                                        if (430 <= 65535) {
                                            switch (v3 + v3) {
                                            case "\n\t":
                                                fmt.Println(v42)
                                            }
                                            v58 += v58
                                        } else {
                                            switch v3 {
                                            case "":
                                                var v109 float32 = 0.0
                                                _ = v109
                                            case "a":
                                                v1[708] = 2147483647
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                        v58 += v58
                                        var v110 uint8 = v42
                                        _ = v110
                                    default:
                                        v111 := GenS9{F0: 42, F1: 0}
                                        _ = v111
                                    }
                                }
                                v58 = (v58 + v58)
                                _ = v58
                            } else {
                                var v112 string = "a"
                                v113 := GenS5{F0: 551}
                                v114 := GenS10{F0: (v42 - 1), F1: (v56 + (v56 + (v56 >> ((18446744073709551615) & 63))))}
                                v115 := GenS10{F0: 128, F1: (v56 | (v56 << ((173) & 63)))}
                                fmt.Println(v56, v26, v4)
                                switch (fi44 / ((402) | 1)) {
                                case 128:
                                    if v4 {
                                        var v116 uint64 = v26
                                        if (42 > (v56 << ((1) & 63))) {
                                            if (true && true) {
                                                fmt.Println(v26)
                                                if (!v4) {
                                                    fmt.Println(v43, v112, v116)
                                                    var v117 float64 = (3.14159 / float64((3.14159 - -2.25)))
                                                    if (!false) {
                                                        if (!(-2.25 > (1e10 / (((1.5 * (1.5 + ((3.14159 * (-2.25 * (1e10 / 1.5))) * 1.5))) * ((0.0 + (0.1 + 1e10)) * float32(v117))) * 0.0)))) {
                                                            switch fi57 {
                                                            case 4294967295:
                                                                v4 = (0 > (v42 << ((65535) & 63)))
                                                            case 2147483647:
                                                                var v118 int = -906
                                                                _ = v118
                                                            case 1:
                                                                v119 := GenS3{F0: uint16(v26), F1: true}
                                                                _ = v119
                                                            }
                                                        } else {
                                                            if (int16(v26) < -32768) {
                                                                fmt.Println(v116, v3)
                                                                v3 = "a"
                                                                v120 := []int{}
                                                                _ = v120
                                                            }
                                                            switch fi57 {
                                                            case 32767:
                                                                switch fi44 {
                                                                case 128:
                                                                    v121 := GenS11{F0: v42, F1: 42, F2: 2}
                                                                    _ = v121
                                                                case 517:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        v42 -= v42
                                                                    default:
                                                                        switch (^(fi57)) {
                                                                        case 32768:
                                                                            v122 := GenS7{F0: (v112 + v3), F1: 1}
                                                                            _ = v122
                                                                        case 0:
                                                                            v3 = (v112 + v112)
                                                                        case 1:
                                                                            v42 -= v42
                                                                            v123 := GenS1{F0: 730, F1: 127, F2: (fi44 & -2147483648)}
                                                                            v117 += v117
                                                                            var v124 uint = 65535
                                                                            var v125 uint16 = 32767
                                                                            v126 := GenS4{F0: v26, F1: (v124 & v124)}
                                                                            _ = v123
                                                                            _ = v124
                                                                            _ = v125
                                                                            _ = v126
                                                                        default:
                                                                            switch (v112 + ((v3 + "hello") + "hello")) {
                                                                            case "hello":
                                                                                var v127 int32 = v56
                                                                                _ = v127
                                                                            case "":
                                                                                fmt.Println(v56)
                                                                            case "\n\t":
                                                                                v117 += v117
                                                                                if v4 {
                                                                                    v128 := GenS10{F0: v42, F1: (v56 ^ v56)}
                                                                                    var v129 int = fi57
                                                                                    if (!v4) {
                                                                                        var v130 uint8 = v42
                                                                                        var v131 string = (v3 + v112)
                                                                                        _ = v130
                                                                                        _ = v131
                                                                                    }
                                                                                    v132 := []int{(fi44 - -1), v129}
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                    _ = v132
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                case -9223372036854775808:
                                                                    v133 := GenS4{F0: 0, F1: 817}
                                                                    _ = v133
                                                                default:
                                                                    var v134 int32 = -44
                                                                    var v135 uint16 = 42
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            case -2147483648:
                                                                var v136 uint64 = (v116 % ((332) | 1))
                                                                _ = v136
                                                            case 830:
                                                                v137 := GenS12{F0: 127}
                                                                _ = v137
                                                            default:
                                                                var v138 int16 = 1000
                                                                _ = v138
                                                            }
                                                        }
                                                    }
                                                    _ = v117
                                                }
                                            } else {
                                                v112 = (v112 + ("\n\t" + ((v112 + ("\n\t" + "foo bar")) + (("" + (v3 + "a")) + v3))))
                                                v139 := GenS13{F0: -32768, F1: 1}
                                                _ = v139
                                            }
                                            v112 = v112
                                            if (v4 || false) {
                                                switch (fi44 ^ (fi57 & (fi44 ^ 933))) {
                                                case 128:
                                                    if v4 {
                                                        v140 := []int{}
                                                        _ = v140
                                                    }
                                                    var v141 bool = true
                                                    fmt.Println(fi57, fi44, v26)
                                                    _ = v141
                                                default:
                                                    switch -996 {
                                                    case 42:
                                                        v142 := GenS14{F0: 366, F1: 2147483647, F2: 60}
                                                        _ = v142
                                                    case 9223372036854775807:
                                                        v143 := v1[2147483648]
                                                        if (1.5 <= ((-2.25 * (((0.1 + (1.5 * 1.5)) - 0.0) - 1e10)) - (3.14159 / 1.0))) {
                                                            switch fi57 {
                                                            case -1:
                                                                fmt.Println(v116, v116)
                                                                v3 = (v3 + "\n\t")
                                                                v4 = true
                                                                var v144 bool = (true && (!(0 > -102)))
                                                                _ = v144
                                                            case 32768:
                                                                v145 := GenS15{F0: v143, F1: 445, F2: -285}
                                                                _ = v145
                                                            case 42:
                                                                var v146 float32 = 3.14159
                                                                _ = v146
                                                            }
                                                            var v147 int = (^(fi57))
                                                            switch -569 {
                                                            case 4294967295:
                                                                v42 = (v42 ^ v42)
                                                            case -2147483648:
                                                                v42 -= 2
                                                                v148 := GenS8{F0: 771, F1: 32767, F2: (v43 % (((^(v43))) | 1))}
                                                                fmt.Println(fi57)
                                                                _ = v148
                                                            default:
                                                                v149 := GenS14{F0: 465, F1: 4294967295, F2: (v42 & 42)}
                                                                _ = v149
                                                            }
                                                            _ = v147
                                                        }
                                                        _ = v143
                                                    case 412:
                                                        v150 := GenS15{F0: 9223372036854775807, F1: 562, F2: -32768}
                                                        _ = v150
                                                    default:
                                                        v151 := v1[1]
                                                        _ = v151
                                                    }
                                                }
                                                var v152 uint16 = 55
                                                fmt.Println(fi57, v4, v152)
                                                _ = v152
                                            } else {
                                                v42 += (v42 + (v42 * 138))
                                                if true {
                                                    v153 := GenS10{F0: (v42 << ((4294967295) & 63)), F1: (v56 << ((65535) & 63))}
                                                    switch fi57 {
                                                    case 128:
                                                        var v154 uint32 = 65535
                                                        _ = v154
                                                    case 32768:
                                                        fmt.Println(fi57)
                                                    case 0:
                                                        v155 := v2[2147483647]
                                                        _ = v155
                                                    }
                                                    v43 = v43
                                                    if false {
                                                        var v156 uint = 4294967295
                                                        v157 := v1[65535]
                                                        var v158 uint16 = 0
                                                        fmt.Println(v26, v112, fi57)
                                                        switch ((("foo bar" + v112) + (v3 + v3)) + v3) {
                                                        case "hello":
                                                            if true {
                                                                if (v4 || (!(!((v42 * v42) <= (v42 - 0))))) {
                                                                    var v159 string = v112
                                                                    v160 := []int{}
                                                                    switch ("" + v3) {
                                                                    case "\n\t":
                                                                        v161 := GenS12{F0: (v42 % ((42) | 1))}
                                                                        _ = v161
                                                                    }
                                                                    v42 += v42
                                                                    var v162 int32 = (v56 / (((v56 + v56)) | 1))
                                                                    _ = v159
                                                                    _ = v160
                                                                    _ = v162
                                                                }
                                                                v156 = uint(v158)
                                                            }
                                                        default:
                                                            v163 := v2[42]
                                                            _ = v163
                                                        }
                                                        _ = v156
                                                        _ = v157
                                                        _ = v158
                                                    }
                                                    if ((^(v56)) <= (v56 * 1)) {
                                                        if (!(2147483647 >= 2147483648)) {
                                                            v56 -= (v56 >> ((255) & 63))
                                                        }
                                                        var v164 int32 = v56
                                                        _ = v164
                                                    } else {
                                                        var v165 float32 = 3.14159
                                                        fmt.Println(v4, v43, fi44)
                                                        _ = v165
                                                    }
                                                    v166 := GenS16{F0: -432, F1: 32767}
                                                    _ = v153
                                                    _ = v166
                                                } else {
                                                    _ = len(v1)
                                                    var v167 uint = 18446744073709551615
                                                    _ = v167
                                                }
                                            }
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(v42)
                                            case "":
                                                var v168 uint64 = (^(v26))
                                                _ = v168
                                            case "foo bar":
                                                if true {
                                                    if (!(v56 == 0)) {
                                                        var v169 int32 = 769
                                                        _ = v169
                                                    } else {
                                                        switch (fi57 | (^(fi44))) {
                                                        case 9223372036854775807:
                                                            if (v4 || (0 != 0)) {
                                                                v170 := GenS4{F0: (v116 - (v26 << ((255) & 63))), F1: 0}
                                                                if v4 {
                                                                    switch v112 {
                                                                    case "\n\t":
                                                                        switch fi44 {
                                                                        case 4294967295:
                                                                            fmt.Println(v42)
                                                                        default:
                                                                            v171 := GenS11{F0: uint8(v116), F1: 1, F2: v42}
                                                                            _ = v171
                                                                        }
                                                                    case "hello":
                                                                        var v172 int64 = 370
                                                                        _ = v172
                                                                    default:
                                                                        fmt.Println(v56)
                                                                    }
                                                                } else {
                                                                    v173 := []int{(fi44 | 4294967295)}
                                                                    fmt.Println(v26)
                                                                    v173 = append(v173, (fi44 | 4294967295))
                                                                    if (-91 >= 42) {
                                                                        var v174 uint8 = v42
                                                                        v175 := GenS6{F0: 65535}
                                                                        var v176 int16 = 545
                                                                        _ = v174
                                                                        _ = v175
                                                                        _ = v176
                                                                    } else {
                                                                        if false {
                                                                            if v4 {
                                                                                var v177 uint16 = 0
                                                                                var v178 int32 = (v56 >> ((255) & 63))
                                                                                _ = v177
                                                                                _ = v178
                                                                            } else {
                                                                                var v179 float32 = 3.14159
                                                                                _ = v179
                                                                            }
                                                                        }
                                                                        v1[2] = (fi44 * (^(fi44)))
                                                                        var v180 int = fi57
                                                                        fmt.Println(v43)
                                                                        v181 := v1[0]
                                                                        if (((!(0 != 42)) && (!(911 != 195))) || v4) {
                                                                            v182 := GenS17{F0: (v26 + 18446744073709551615), F1: (v180 % ((1) | 1))}
                                                                            _ = v182
                                                                        }
                                                                        v183 := GenS7{F0: v112, F1: 255}
                                                                        _ = v180
                                                                        _ = v181
                                                                        _ = v183
                                                                    }
                                                                    _ = v173
                                                                }
                                                                v56 -= (^(v56))
                                                                v184 := GenS16{F0: fi57, F1: 0}
                                                                var v185 uint64 = v26
                                                                v43 -= (^(v43))
                                                                _ = v170
                                                                _ = v184
                                                                _ = v185
                                                            }
                                                        case 127:
                                                            var v186 int32 = (v56 | v56)
                                                            _ = v186
                                                        case 32768:
                                                            switch v3 {
                                                            case "a":
                                                                switch -9223372036854775808 {
                                                                case -446:
                                                                    var v187 int32 = (v56 & (^(v56)))
                                                                    _ = v187
                                                                case 42:
                                                                    var v188 uint64 = v116
                                                                    _ = v188
                                                                case -9223372036854775808:
                                                                    if v4 {
                                                                        switch -26 {
                                                                        case 128:
                                                                            fmt.Println(fi44)
                                                                        }
                                                                        var v189 float32 = 1.5
                                                                        if ((v43 ^ v43) < v43) {
                                                                            v116 = v26
                                                                            switch v3 {
                                                                            case "\n\t":
                                                                                var v190 bool = v4
                                                                                _ = v190
                                                                            case "":
                                                                                fmt.Println(fi44, v43, v112)
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v43, v4, v42)
                                                                            v56 = v56
                                                                            var v191 float64 = 0.1
                                                                            var v192 int64 = v43
                                                                            _ = v191
                                                                            _ = v192
                                                                        }
                                                                        v43 -= 52
                                                                        var v193 int64 = v43
                                                                        var v194 uint8 = (v42 ^ (^(v42)))
                                                                        fmt.Println(fi44)
                                                                        _ = v189
                                                                        _ = v193
                                                                        _ = v194
                                                                    }
                                                                }
                                                                v56 = v56
                                                                var v195 int32 = v56
                                                                _ = v195
                                                            case "":
                                                                var v196 int8 = -1
                                                                _ = v196
                                                            case "\n\t":
                                                                var v197 uint8 = v42
                                                                _ = v197
                                                            default:
                                                                switch v3 {
                                                                case "a":
                                                                    if (-75 <= -64) {
                                                                        v3 = "a"
                                                                    }
                                                                case "\n\t":
                                                                    var v198 uint8 = 1
                                                                    _ = v198
                                                                default:
                                                                    var v199 string = ""
                                                                    _ = v199
                                                                }
                                                            }
                                                        }
                                                        v112 = "foo bar"
                                                        if (42 <= 1) {
                                                            if (1.5 > 3.14159) {
                                                                if (v3 <= "hello") {
                                                                    if false {
                                                                        switch fi44 {
                                                                        case 42:
                                                                            if (v4 && v4) {
                                                                                v200 := []int{fi44}
                                                                                if v4 {
                                                                                    var v201 uint = 416
                                                                                    v202 := GenS18{F0: 751, F1: "\n\t"}
                                                                                    v201 = 18446744073709551615
                                                                                    _ = v201
                                                                                    _ = v202
                                                                                }
                                                                                v3 = ((v3 + "") + v112)
                                                                                _ = v200
                                                                            } else {
                                                                                v203 := []int{(fi57 ^ fi44), fi44, 42}
                                                                                _ = len(v203)
                                                                                fmt.Println(v43)
                                                                                _ = v203
                                                                            }
                                                                        case 1:
                                                                            v3 = "\n\t"
                                                                        case -2147483648:
                                                                            switch fi57 {
                                                                            case -756:
                                                                                v116 += (v116 % ((v116) | 1))
                                                                            default:
                                                                                if ("foo bar" == "foo bar") {
                                                                                    switch ((v112 + v112) + (v112 + (v3 + v112))) {
                                                                                    case "hello":
                                                                                        var v204 int16 = 0
                                                                                        _ = v204
                                                                                    default:
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    v3 = "\n\t"
                                                                                }
                                                                            }
                                                                        }
                                                                        var v205 float64 = 0.1
                                                                        fmt.Println(v205)
                                                                        v206 := []int{(fi57 - 32767), fi57, (fi44 - -2147483648), fi44}
                                                                        fmt.Println(v116, v116, v4)
                                                                        switch v112 {
                                                                        case "hello":
                                                                            if false {
                                                                                switch fi44 {
                                                                                case 2147483647:
                                                                                    _ = len(v206)
                                                                                case -1:
                                                                                    var v207 float32 = 0.0
                                                                                    _ = v207
                                                                                }
                                                                                v116 += 418
                                                                                v43 = int64(v116)
                                                                            } else {
                                                                                var v208 int32 = v56
                                                                                switch -2147483648 {
                                                                                case -9223372036854775808:
                                                                                    v205 += ((v205 / v205) * v205)
                                                                                case 9223372036854775807:
                                                                                    v3 = ((("\n\t" + v3) + "a") + ("a" + v112))
                                                                                default:
                                                                                    v209 := GenS10{F0: (v42 << ((255) & 63)), F1: (v56 % (((v208 / (((v208 << ((65535) & 63))) | 1))) | 1))}
                                                                                    _ = v209
                                                                                }
                                                                                switch v112 {
                                                                                case "foo bar":
                                                                                    v112 = ((v3 + "a") + v112)
                                                                                    v26 = v26
                                                                                case "\n\t":
                                                                                    v42 -= 1
                                                                                }
                                                                                _ = v208
                                                                            }
                                                                            fmt.Println(v42, v56)
                                                                        }
                                                                        _ = v205
                                                                        _ = v206
                                                                    }
                                                                    v3 = "a"
                                                                } else {
                                                                    fmt.Println(v112, v43, fi44)
                                                                }
                                                                v43 -= -9223372036854775808
                                                            }
                                                        }
                                                        v4 = true
                                                        v43 = v43
                                                        switch ((v3 + "foo bar") + ("a" + "")) {
                                                        case "\n\t":
                                                            if v4 {
                                                                fmt.Println(v43)
                                                                v2[2147483648] = fi44
                                                                v26 = (v26 + 0)
                                                            } else {
                                                                v4 = (false || (fi44 < 127))
                                                                v2[1] = fi44
                                                            }
                                                            var v210 int = (fi57 | fi44)
                                                            var v211 float64 = (float64(0.1) / 1.0)
                                                            _ = v210
                                                            _ = v211
                                                        }
                                                        v212 := []int{-949, (fi57 + (fi44 + 32767)), (fi57 ^ 0)}
                                                        fmt.Println(v4, v4, v42)
                                                        if (1 <= 0) {
                                                            var v213 uint8 = 243
                                                            _ = v213
                                                        } else {
                                                            v43 += -417
                                                            if false {
                                                                switch v112 {
                                                                case "hello":
                                                                    var v214 uint16 = 2
                                                                    _ = v214
                                                                default:
                                                                    if v4 {
                                                                        if (((((1.5 / 3.14159) - 1.5) - (-2.25 / 1.5)) + (0.0 * -2.25)) != (1.5 - ((float32(float64(((0.0 / 3.14159) / 1.0))) - ((1.5 / 1e10) * float32(((0.0 / -2.25) - (1.5 * -2.25))))) - 1e10))) {
                                                                            fmt.Println(v26, fi57)
                                                                            var v215 string = "foo bar"
                                                                            var v216 uint8 = 176
                                                                            fmt.Println(v116, fi57, v42)
                                                                            _ = v215
                                                                            _ = v216
                                                                        } else {
                                                                            fmt.Println(v26, v116)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v212
                                                    }
                                                    v43 -= 9223372036854775807
                                                } else {
                                                    v217 := GenS14{F0: 65535, F1: 65535, F2: (^(v42))}
                                                    switch fi57 {
                                                    case 1:
                                                        var v218 uint64 = v116
                                                        _ = v218
                                                    }
                                                    _ = v217
                                                }
                                            default:
                                                var v219 int = (fi44 % ((fi57) | 1))
                                                _ = v219
                                            }
                                            v220 := v2[2147483648]
                                            _ = v220
                                        }
                                        fmt.Println(v112, v4, v3)
                                        v42 += (v42 + (v42 * v42))
                                        _ = v116
                                    }
                                case -2147483648:
                                    var v221 uint8 = (v42 - (v42 + (^(v42))))
                                    _ = v221
                                }
                                _ = v112
                                _ = v113
                                _ = v114
                                _ = v115
                            }
                            v42 -= v42
                        }
                        fmt.Println(v26, v26, v4)
                        var v222 string = v3
                        if (-66 != -64) {
                            var v223 int16 = -1
                            _ = v223
                        } else {
                            v224 := GenS15{F0: (fi44 >> ((65535) & 63)), F1: 765, F2: -266}
                            _ = v224
                        }
                        switch (fi44 - (-(fi44))) {
                        case 127:
                            _ = len(v2)
                        default:
                            switch (fi57 | (fi44 - (fi57 - (fi44 % ((fi44) | 1))))) {
                            case 2147483647:
                                var v225 int32 = v56
                                _ = v225
                            case 882:
                                var v226 int64 = (v43 >> ((65535) & 63))
                                _ = v226
                            case -436:
                                v4 = v4
                            default:
                                v56 -= v56
                            }
                        }
                        _ = fi57
                        _ = v222
                    }
                } else {
                    switch (v3 + "") {
                    case "foo bar":
                        var v227 float32 = 1e10
                        _ = v227
                    }
                    delete(v2, 2147483648)
                }
                _ = v56
            }
            _ = fi44
        }
        _ = v43
    }
    var v228 uint64 = v26
    if v4 {
        fmt.Println(v4, v42, v3)
        delete(v2, 4294967295)
    } else {
        v228 = v26
    }
    v229 := GenS19{F0: 255, F1: -106, F2: -2147483648}
    switch 2147483647 {
    case 273:
        v230 := GenS20{F0: 100}
        _ = v230
    case 1:
        var v231 string = "foo bar"
        _ = v231
    case -1:
        for fi232 := 0; fi232 < 5; fi232++ {
            v228 += (^(v228))
            fmt.Println(fi232)
            _ = fi232
        }
    default:
        v233 := GenS21{F0: 969}
        _ = v233
    }
    v234 := GenS6{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v228:", v228)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v228
    _ = v229
    _ = v234
    _ = v26
    _ = v3
    _ = v4
    _ = v42
}
