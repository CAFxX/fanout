package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 uint8; F1 uint }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint32; F1 uint8; F2 string }
type GenS16 struct { F0 int64; F1 bool; F2 int64 }
type GenS17 struct { F0 int8; F1 int8 }
func genFn0(p0 int32) int8 {
    v7 := []int{9223372036854775807, -1, 923, 980}
    var v8 int16 = 32767
    _ = p0
    _ = v7
    _ = v8
    return 100
}
func genFn1() bool {
    v9 := []int{32767}
    _ = len(v9)
    var v10 int16 = 100
    _ = v10
    _ = v9
    return (1 < v10)
}

func main() {
    var v1 int8 = 77
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 int64 = -306
    v5 := make(map[uint32]bool)
    v6 := make(map[string]uint32)
    if (42 > 989) {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint16 = 65535
            v1 += (v1 % (((v1 / ((v1) | 1))) | 1))
            _ = fi11
            _ = v12
        }
        var v13 string = "hello"
        v13 = "a"
        _ = v13
    }
    fmt.Println(v1, v1)
    fmt.Println(v4, v4)
    v4 += v4
    v14 := GenS1{F0: -98}
    if (((v4 & (v4 * 572)) != -9223372036854775808) && false) {
        v15 := v5[2147483647]
        v4 += -9223372036854775808
        if (v1 >= (v1 & genFn0(-673))) {
            var v16 int16 = 32767
            _ = v16
        }
        _ = v15
    } else {
        v4 -= 35
        v1 = genFn0(2147483647)
        var v17 uint64 = 4294967295
        _ = v17
    }
    var v18 uint32 = 730
    if true {
        v1 = genFn0(1)
    } else {
        v18 -= (v18 << ((1) & 63))
        switch ("" + "\n\t") {
        case "":
            var v19 uint8 = 223
            _ = v19
        case "foo bar":
            if (-54 > -1) {
                v4 += 9223372036854775807
                v20 := []int{30}
                if (false || true) {
                    var v21 int8 = v1
                    for fi22 := 0; fi22 < 5; fi22++ {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            switch "\n\t" {
                            case "\n\t":
                                switch -9223372036854775808 {
                                case 4294967295:
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                }
                            case "a":
                                if true {
                                    var v24 float64 = 1e10
                                    if (3.14159 == 1e10) {
                                        switch "" {
                                        case "hello":
                                            switch "hello" {
                                            case "a":
                                                v25 := GenS1{F0: fi23}
                                                _ = v25
                                            case "":
                                                v21 = v21
                                            case "\n\t":
                                                fmt.Println(v4, v18, v4)
                                            default:
                                                if ((1.5 + -2.25) < 1e10) {
                                                    if ((^(v18)) != (v18 & (v18 | v18))) {
                                                        v26 := GenS2{F0: 255, F1: v4, F2: 522}
                                                        var v27 int = fi22
                                                        if true {
                                                            v28 := GenS2{F0: 0, F1: 1, F2: 532}
                                                            if true {
                                                                v29 := GenS3{F0: 1000}
                                                                _ = v29
                                                            }
                                                            var v30 uint32 = 4294967295
                                                            _ = v28
                                                            _ = v30
                                                        }
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        v31 := GenS3{F0: 623}
                                                        var v32 float64 = v24
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    if false {
                                                        v24 = float64(0.0)
                                                        var v33 int64 = v4
                                                        _ = v33
                                                    }
                                                }
                                            }
                                            var v34 float32 = ((0.0 + 0.1) + 3.14159)
                                            if true {
                                                fmt.Println(v24, v18)
                                                v35 := GenS2{F0: 1, F1: v4, F2: 963}
                                                _ = v35
                                            } else {
                                                switch fi23 {
                                                case 32767:
                                                    fmt.Println(v4, v21)
                                                case 2147483647:
                                                    v20 = append(v20, fi23)
                                                }
                                                var v36 int = (fi22 - -760)
                                                var v37 int32 = -2147483648
                                                _ = v36
                                                _ = v37
                                            }
                                            _ = v34
                                        case "a":
                                            v38 := GenS3{F0: 482}
                                            _ = v38
                                        case "":
                                            if (2 == 42) {
                                                switch fi23 {
                                                case 9223372036854775807:
                                                    _ = len(v20)
                                                case 128:
                                                    delete(v3, 32767)
                                                case -720:
                                                    v24 += v24
                                                default:
                                                    var v39 float32 = 1e10
                                                    _ = v39
                                                }
                                            } else {
                                                v24 = float64(0.0)
                                            }
                                        }
                                        if (true || (42 <= 541)) {
                                            if false {
                                                v4 += (v4 << ((42) & 63))
                                                v40 := GenS2{F0: 128, F1: 628, F2: 63}
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if false {
                                                        switch fi22 {
                                                        case 0:
                                                            switch fi22 {
                                                            case 1:
                                                                v41 := GenS1{F0: (fi22 - fi23)}
                                                                _ = v41
                                                            }
                                                        case 142:
                                                            var v42 string = ""
                                                            _ = v42
                                                        }
                                                        var v43 uint8 = 128
                                                        var v44 uint16 = 255
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        if false {
                                                            v18 = (v18 + 2)
                                                            v18 -= 2147483648
                                                            v45 := GenS4{F0: v21, F1: 2}
                                                            _ = v45
                                                        }
                                                        v4 -= v4
                                                    }
                                                case "hello":
                                                    v21 -= (-(v21))
                                                    v46 := GenS4{F0: genFn0(292), F1: 4294967295}
                                                    var v47 int = fi23
                                                    _ = v46
                                                    _ = v47
                                                default:
                                                    v48 := GenS3{F0: 752}
                                                    _ = v48
                                                }
                                                _ = v40
                                            }
                                            if false {
                                                v49 := v2[2147483648]
                                                if len(v20) > 0 {
                                                    v20[0] = (fi23 * 531)
                                                }
                                                if (!((v1 * genFn0(-1)) == 1)) {
                                                    if (-1 >= 42) {
                                                        v18 -= v18
                                                        var v50 int16 = 1
                                                        if (!true) {
                                                            var v51 uint = 65535
                                                            if (v51 >= 2) {
                                                                if false {
                                                                    v51 -= 2
                                                                    v24 = v24
                                                                    _ = len(v20)
                                                                    fmt.Println(v24, fi22)
                                                                } else {
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                    v21 += (v21 / ((100) | 1))
                                                                    switch fi22 {
                                                                    case -1:
                                                                        if (((false || false) && (-2.25 > (1e10 - 1e10))) && true) {
                                                                            v52 := GenS4{F0: int8(v51), F1: 255}
                                                                            var v53 int64 = v4
                                                                            v21 += (v21 & genFn0(-1))
                                                                            _ = v52
                                                                            _ = v53
                                                                        } else {
                                                                            var v54 uint32 = (v18 | (^(v18)))
                                                                            var v55 uint = v51
                                                                            switch fi23 {
                                                                            case -353:
                                                                                _ = len(v6)
                                                                            case 32767:
                                                                                fmt.Println(v51)
                                                                            case 604:
                                                                                var v56 uint8 = 128
                                                                                _ = v56
                                                                            }
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                    case -2147483648:
                                                                        v4 = (v4 & v4)
                                                                    default:
                                                                        v18 = 1
                                                                    }
                                                                }
                                                            }
                                                            if (42 != (v18 - v18)) {
                                                                if (2 >= 127) {
                                                                    v4 = (v4 / (((v4 - (v4 * v4))) | 1))
                                                                    v4 = v4
                                                                    if ((v1 != v21) && true) {
                                                                        fmt.Println(v21, v49)
                                                                        switch fi23 {
                                                                        case 4294967295:
                                                                            v57 := GenS5{F0: "", F1: true}
                                                                            _ = v57
                                                                        case 535:
                                                                            v4 += int64(v1)
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if false {
                                                                            v1 = (-(v21))
                                                                        }
                                                                        v58 := GenS6{F0: 1, F1: ""}
                                                                        _ = v58
                                                                    }
                                                                    v18 += (v49 / ((639) | 1))
                                                                }
                                                                if len(v20) > 0 {
                                                                    v20[0] = fi23
                                                                }
                                                            } else {
                                                                var v59 string = ("a" + "")
                                                                v60 := GenS7{F0: v50}
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            _ = v51
                                                        }
                                                        _ = v50
                                                    }
                                                }
                                                fmt.Println(v21, v1)
                                                _ = v49
                                            } else {
                                                if false {
                                                    var v61 uint32 = (v18 | (v18 >> ((911) & 63)))
                                                    _ = v61
                                                }
                                                var v62 string = "foo bar"
                                                fmt.Println(v24)
                                                _ = v62
                                            }
                                        }
                                        v63 := GenS7{F0: -1}
                                        _ = v63
                                    }
                                    _ = v24
                                }
                            case "foo bar":
                                if (2 == 1) {
                                    if true {
                                        var v64 float64 = ((-2.25 * 1e10) * (3.14159 * -2.25))
                                        var v65 uint8 = uint8(v18)
                                        _ = v64
                                        _ = v65
                                    } else {
                                        if (fi22 == (fi23 % (((^(fi22))) | 1))) {
                                            var v66 string = "a"
                                            switch "\n\t" {
                                            case "":
                                                if ((!false) || false) {
                                                    if (!(!true)) {
                                                        fmt.Println(fi22, v4)
                                                        var v67 float32 = (-2.25 * 3.14159)
                                                        var v68 int8 = v21
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    if ((-394 < 42) || true) {
                                                        v69 := v2[v18]
                                                        var v70 int16 = 1
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    var v71 uint32 = v18
                                                    _ = v71
                                                }
                                            case "\n\t":
                                                v72 := GenS4{F0: 69, F1: 4294967295}
                                                _ = v72
                                            case "hello":
                                                if len(v20) > 0 {
                                                    v20[0] = fi22
                                                }
                                            default:
                                                switch (fi23 >> ((0) & 63)) {
                                                case 2147483647:
                                                    var v73 string = "\n\t"
                                                    _ = v73
                                                case 1:
                                                    var v74 uint8 = uint8(v18)
                                                    _ = v74
                                                case 128:
                                                    v20 = append(v20, (-(fi23)))
                                                }
                                            }
                                            _ = v66
                                        } else {
                                            v3[(fi23 * 44)] = 65535
                                        }
                                        v21 -= genFn0(-190)
                                    }
                                } else {
                                    _ = len(v6)
                                    if ((float64(0.0) >= 0.0) && false) {
                                        var v75 uint = 18446744073709551615
                                        _ = v75
                                    } else {
                                        var v76 int8 = -71
                                        v77 := GenS8{F0: (^(v21)), F1: "a"}
                                        _ = v76
                                        _ = v77
                                    }
                                }
                            }
                            v78 := GenS3{F0: -1}
                            _ = fi23
                            _ = v78
                        }
                        for fi79 := 0; fi79 < 4; fi79++ {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = fi79
                        }
                        v3[fi22] = (v18 << ((0) & 63))
                        _ = fi22
                    }
                    var v80 int = -2147483648
                    _ = v21
                    _ = v80
                }
                _ = v20
            }
        default:
            v18 += (v18 | 585)
        }
    }
    v81 := GenS9{F0: 32767}
    var v82 uint64 = 0
    var v83 uint32 = (^(v18))
    v84 := GenS2{F0: 82, F1: -1, F2: v82}
    if (316 < v82) {
        v85 := GenS1{F0: 291}
        v18 += v83
        _ = v85
    } else {
        for fi86 := 0; fi86 < 8; fi86++ {
            if (uint(v18) > 23) {
                switch "" {
                case "foo bar":
                    v2[(v18 | v18)] = v83
                case "hello":
                    v87 := GenS5{F0: "\n\t", F1: (!true)}
                    _ = v87
                case "a":
                    for fi88 := 0; fi88 < 7; fi88++ {
                        var v89 int32 = -809
                        v90 := []int{(fi88 / ((fi86) | 1)), (fi86 + (-(fi88))), 127, (fi88 / (((fi88 | -54)) | 1))}
                        if (-619 != -1) {
                            v82 = 4294967295
                        } else {
                            v1 -= -1
                        }
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                }
            }
            var v91 uint = 65535
            switch "hello" {
            case "":
                switch (^(fi86)) {
                case -1:
                    _ = len(v5)
                case 32767:
                    if (!(v82 >= (v82 % (((v82 - v82)) | 1)))) {
                        v92 := []int{(fi86 * int(v83)), 991}
                        var v93 int = (fi86 % ((fi86) | 1))
                        switch v93 {
                        case 4294967295:
                            for fi94 := 0; fi94 < 2; fi94++ {
                                if false {
                                    v91 = (v91 + (v91 + v91))
                                    var v95 bool = (v83 < (^(v83)))
                                    _ = v95
                                } else {
                                    switch "hello" {
                                    case "":
                                        if (v1 >= v1) {
                                            if (v91 > 54) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v96 uint8 = 0
                                                    _ = v96
                                                }
                                            } else {
                                                if true {
                                                    var v97 uint64 = (v82 & 398)
                                                    v98 := GenS6{F0: 991, F1: ""}
                                                    fmt.Println(v4)
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    v93 = 0
                                                    v2[464] = 4294967295
                                                    var v99 uint16 = 816
                                                    v93 = fi86
                                                    switch "hello" {
                                                    case "":
                                                        v82 += v82
                                                    case "foo bar":
                                                        var v100 float32 = 3.14159
                                                        _ = v100
                                                    }
                                                    switch fi86 {
                                                    case 4294967295:
                                                        var v101 float32 = 0.0
                                                        _ = v101
                                                    case -742:
                                                        var v102 int = (^(fi86))
                                                        _ = v102
                                                    }
                                                    var v103 int = (fi86 / ((v93) | 1))
                                                    _ = v103
                                                    _ = v99
                                                }
                                                if false {
                                                    if ((42 > 306) || false) {
                                                        v104 := GenS3{F0: -32768}
                                                        v82 -= v82
                                                        _ = v104
                                                    } else {
                                                        var v105 int = 32768
                                                        v105 -= 42
                                                        if (!true) {
                                                            var v106 int = 4294967295
                                                            v107 := GenS2{F0: 241, F1: v4, F2: 42}
                                                            _ = v106
                                                            _ = v107
                                                        } else {
                                                            if false {
                                                                var v108 int8 = (v1 & (v1 - v1))
                                                                _ = v108
                                                            }
                                                        }
                                                        _ = v105
                                                    }
                                                    switch 0 {
                                                    case -2147483648:
                                                        switch int(v1) {
                                                        case -1:
                                                            v91 = (v91 % ((v91) | 1))
                                                        case 42:
                                                            var v109 int64 = v4
                                                            _ = v109
                                                        }
                                                    case 446:
                                                        fmt.Println(v91)
                                                    }
                                                    switch v93 {
                                                    case -140:
                                                        if false {
                                                            v110 := GenS4{F0: -77, F1: (v91 % ((4294967295) | 1))}
                                                            v93 -= v93
                                                            _ = v110
                                                        } else {
                                                            v111 := GenS4{F0: 1, F1: (^(v91))}
                                                            _ = v111
                                                        }
                                                    case 1:
                                                        var v112 uint16 = 32767
                                                        v113 := GenS3{F0: 1000}
                                                        _ = v112
                                                        _ = v113
                                                    case 601:
                                                        if len(v92) > 0 {
                                                            v92[0] = (fi86 - v93)
                                                        }
                                                    }
                                                } else {
                                                    v92 = append(v92, (fi94 ^ 9223372036854775807))
                                                }
                                                v114 := GenS10{F0: (v1 - v1), F1: "hello"}
                                                var v115 int8 = v1
                                                _ = v114
                                                _ = v115
                                            }
                                            v92 = append(v92, (fi86 & (v93 & int(v82))))
                                            var v116 float32 = ((1.5 + 1.5) - (0.0 + 1.5))
                                            _ = v116
                                        }
                                    default:
                                        v4 -= v4
                                    }
                                }
                                if false {
                                    fmt.Println(v18)
                                    var v117 uint16 = 105
                                    _ = v117
                                }
                                switch "a" {
                                case "foo bar":
                                    var v118 int64 = (v4 + int64(v83))
                                    _ = v118
                                case "":
                                    var v119 int32 = 54
                                    _ = v119
                                }
                                _ = fi94
                            }
                        case 1:
                            v82 = v82
                        case 42:
                            fmt.Println(v93, v91)
                        }
                        _ = v92
                        _ = v93
                    } else {
                        v83 = v83
                        v120 := []int{2147483647, (fi86 | (^(fi86)))}
                        _ = v120
                    }
                case 2147483647:
                    delete(v6, "")
                }
            case "a":
                v4 = v4
            default:
                v18 += 1
            }
            _ = fi86
            _ = v91
        }
        if false {
            switch ("hello" + "foo bar") {
            case "a":
                if false {
                    var v121 int64 = 0
                    _ = v121
                }
            }
        } else {
            fmt.Println(v4, v83)
            for fi122 := 0; fi122 < 20; fi122++ {
                if ((127 < 128) && (true && false)) {
                    v123 := GenS7{F0: int16(fi122)}
                    var v124 float64 = 1.5
                    for fi125 := 0; fi125 < 19; fi125++ {
                        if true {
                            var v126 int8 = -1
                            v82 = (^(v82))
                            v127 := []int{(fi122 >> ((uint(v83)) & 63)), (fi125 | fi125), (^(fi122)), fi125}
                            _ = v126
                            _ = v127
                        } else {
                            v128 := []int{}
                            var v129 int = 42
                            if (!true) {
                                switch (("\n\t" + "\n\t") + "foo bar") {
                                case "hello":
                                    var v130 uint32 = v83
                                    _ = v130
                                case "":
                                    switch "" {
                                    case "\n\t":
                                        v131 := GenS11{F0: "foo bar", F1: 671}
                                        _ = v131
                                    case "":
                                        if true {
                                            fmt.Println(v83, v18)
                                        } else {
                                            v124 = float64(0.1)
                                            var v132 int16 = -32768
                                            _ = v132
                                        }
                                    default:
                                        v133 := GenS5{F0: "hello", F1: false}
                                        _ = v133
                                    }
                                case "\n\t":
                                    switch "hello" {
                                    case "hello":
                                        switch "" {
                                        case "a":
                                            v124 -= 1e10
                                        case "":
                                            if true {
                                                if (0 < 139) {
                                                    v124 += (0.1 - -2.25)
                                                    if (false && true) {
                                                        v1 = v1
                                                    }
                                                    v134 := GenS12{F0: 187, F1: 84}
                                                    _ = v134
                                                } else {
                                                    v135 := GenS1{F0: int(v18)}
                                                    fmt.Println(fi122, v129)
                                                    var v136 bool = (266 == 2)
                                                    _ = v135
                                                    _ = v136
                                                }
                                                v4 = 9223372036854775807
                                                v124 = v124
                                            }
                                        }
                                    default:
                                        v137 := GenS13{F0: 95, F1: 890}
                                        _ = v137
                                    }
                                default:
                                    fmt.Println(fi125)
                                }
                                var v138 uint64 = v82
                                fmt.Println(v18)
                                _ = v138
                            } else {
                                v139 := v5[1]
                                v4 = -1
                                if (!(!v139)) {
                                    if len(v128) >= 2 {
                                        v128 = v128[1:len(v128)]
                                    }
                                    fmt.Println(v4, v139, v1)
                                    fmt.Println(v124, v83, v82)
                                }
                                switch "foo bar" {
                                case "hello":
                                    v82 -= (v82 << ((20) & 63))
                                default:
                                    switch (v129 | fi122) {
                                    case 32768:
                                        switch (fi125 ^ (^(v129))) {
                                        case -473:
                                            v140 := v6["a"]
                                            _ = v140
                                        case 1:
                                            var v141 int8 = 100
                                            _ = v141
                                        case 390:
                                            if (!(false && v139)) {
                                                v18 += 4294967295
                                            }
                                        }
                                    case 128:
                                        if v139 {
                                            if false {
                                                v124 = 1.5
                                            } else {
                                                v83 = (v18 << ((0) & 63))
                                                v128 = append(v128, (fi122 - 32767))
                                            }
                                        }
                                    case 127:
                                        if v139 {
                                            if len(v128) > 0 {
                                                v128[0] = (v129 ^ fi122)
                                            }
                                        } else {
                                            v83 -= (v18 + v83)
                                            if ((!false) && (!v139)) {
                                                if false {
                                                    if (v139 && (float64(0.0) <= -2.25)) {
                                                        var v142 int64 = v4
                                                        var v143 int8 = (v1 / ((78) | 1))
                                                        _ = v142
                                                        _ = v143
                                                    } else {
                                                        v144 := GenS5{F0: "foo bar", F1: (1 != 42)}
                                                        fmt.Println(v129, v82, v124)
                                                        v145 := v3[-9223372036854775808]
                                                        if false {
                                                            var v146 int8 = genFn0(0)
                                                            v82 = (v82 | 813)
                                                            if v139 {
                                                                var v147 uint64 = (v82 * v82)
                                                                if ((0.1 - (1e10 - 1.5)) == 1e10) {
                                                                    v139 = (6 < v83)
                                                                    switch (((("foo bar" + "hello") + (("hello" + "a") + ("foo bar" + "hello"))) + ("foo bar" + ("" + ""))) + "\n\t") {
                                                                    case "":
                                                                        if len(v128) > 0 {
                                                                            v128[0] = 127
                                                                        }
                                                                    case "a":
                                                                        v82 += uint64(v145)
                                                                    default:
                                                                        if (("\n\t" == "hello") && (!(1e10 != (0.1 - 1e10)))) {
                                                                            if (562 != 0) {
                                                                                if (!(v124 >= (v124 - 0.1))) {
                                                                                    var v148 int16 = 49
                                                                                    fmt.Println(v18, v139, v146)
                                                                                    _ = v148
                                                                                }
                                                                            }
                                                                            var v149 uint = 18446744073709551615
                                                                            switch fi125 {
                                                                            case -9223372036854775808:
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    var v150 uint = (v149 << ((v149) & 63))
                                                                                    _ = v150
                                                                                case "":
                                                                                    var v151 float32 = (-2.25 * 0.0)
                                                                                    _ = v151
                                                                                case "a":
                                                                                    v139 = (!v139)
                                                                                default:
                                                                                    _ = len(v128)
                                                                                }
                                                                            case -1:
                                                                                var v152 bool = (!((v139 && (true && v139)) || ((35 == 128) || v139)))
                                                                                _ = v152
                                                                            case -850:
                                                                                v153 := GenS12{F0: 692, F1: int32(v145)}
                                                                                _ = v153
                                                                            default:
                                                                                switch "hello" {
                                                                                case "":
                                                                                    if len(v128) >= 2 {
                                                                                        v128 = v128[1:len(v128)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            v3[v129] = 958
                                                                            _ = v149
                                                                        }
                                                                    }
                                                                }
                                                                var v154 uint32 = v83
                                                                if ((v4 >= v4) && (!((false || ((("foo bar" + (("\n\t" + "a") + "")) + "foo bar") <= "foo bar")) && false))) {
                                                                    var v155 int64 = (v4 * (v4 << ((88) & 63)))
                                                                    switch int(v4) {
                                                                    case 0:
                                                                        var v156 int16 = 1
                                                                        _ = v156
                                                                    case 871:
                                                                        switch ("foo bar" + "") {
                                                                        case "foo bar":
                                                                            var v157 uint64 = (^(v147))
                                                                            _ = v157
                                                                        case "hello":
                                                                            switch (v129 >> ((2) & 63)) {
                                                                            case 1:
                                                                                var v158 uint16 = 2
                                                                                _ = v158
                                                                            case 0:
                                                                                if (("a" > ("a" + ("hello" + "a"))) || v139) {
                                                                                    v154 += 2
                                                                                    var v159 uint8 = 174
                                                                                    if len(v128) >= 2 {
                                                                                        v128 = v128[1:len(v128)]
                                                                                    }
                                                                                    _ = v159
                                                                                } else {
                                                                                    v4 = -284
                                                                                }
                                                                                var v160 uint = 1
                                                                                var v161 uint32 = 1
                                                                                _ = v160
                                                                                _ = v161
                                                                            default:
                                                                                v162 := GenS14{F0: -1}
                                                                                var v163 uint8 = 127
                                                                                var v164 float32 = 3.14159
                                                                                _ = v162
                                                                                _ = v163
                                                                                _ = v164
                                                                            }
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                if true {
                                                                                    v165 := GenS1{F0: 127}
                                                                                    v145 = (v83 - 92)
                                                                                    _ = v165
                                                                                } else {
                                                                                    v129 = fi122
                                                                                    var v166 int32 = int32(v82)
                                                                                    _ = v166
                                                                                }
                                                                            }
                                                                        default:
                                                                            v167 := v6[("hello" + "foo bar")]
                                                                            v154 = (^(v83))
                                                                            var v168 uint64 = (v82 << ((0) & 63))
                                                                            _ = v167
                                                                            _ = v168
                                                                        }
                                                                    default:
                                                                        if (!v139) {
                                                                            v155 += (^(v155))
                                                                            v139 = true
                                                                            if (!(v154 > (v145 - (v154 >> ((65535) & 63))))) {
                                                                                v124 = v124
                                                                            }
                                                                        }
                                                                        var v169 bool = v139
                                                                        var v170 int64 = -713
                                                                        _ = v169
                                                                        _ = v170
                                                                    }
                                                                    _ = v155
                                                                }
                                                                _ = v147
                                                                _ = v154
                                                            }
                                                            _ = v146
                                                        }
                                                        v171 := GenS14{F0: 540}
                                                        v129 = (^(v129))
                                                        _ = v144
                                                        _ = v145
                                                        _ = v171
                                                    }
                                                }
                                                var v172 int = (fi122 + (v129 % ((v129) | 1)))
                                                _ = v172
                                            } else {
                                                v173 := GenS15{F0: 578, F1: 100, F2: ("hello" + "hello")}
                                                if len(v128) >= 2 {
                                                    v128 = v128[1:len(v128)]
                                                }
                                                _ = v173
                                            }
                                            switch (-(v129)) {
                                            case 483:
                                                v174 := GenS6{F0: 283, F1: "a"}
                                                var v175 uint16 = 464
                                                var v176 float32 = float32(v124)
                                                v18 -= (v83 / ((v83) | 1))
                                                _ = v174
                                                _ = v175
                                                _ = v176
                                            }
                                        }
                                    default:
                                        var v177 float64 = (v124 / v124)
                                        _ = v177
                                    }
                                }
                                _ = v139
                            }
                            _ = v128
                            _ = v129
                        }
                        if true {
                            var v178 uint64 = (v82 / ((v82) | 1))
                            var v179 bool = (false && (float64(-2.25) >= (v124 - (-2.25 + (v124 * float64(float32((v124 / v124))))))))
                            v1 = -128
                            v180 := []int{}
                            fmt.Println(v83, fi122)
                            _ = v178
                            _ = v179
                            _ = v180
                        }
                        v181 := []int{fi122, fi125, fi122, fi125}
                        fmt.Println(fi125, fi122)
                        switch (fi125 * fi122) {
                        case 9223372036854775807:
                            v1 -= (v1 % (((^(v1))) | 1))
                        case 636:
                            if (true && false) {
                                fmt.Println(v124, v83, v83)
                            } else {
                                var v182 float64 = float64(1e10)
                                _ = v182
                            }
                        default:
                            v183 := v5[(v18 + (v18 + (v83 ^ (v83 | 873))))]
                            _ = v183
                        }
                        _ = fi125
                        _ = v181
                    }
                    _ = v123
                    _ = v124
                } else {
                    if true {
                        if (true || ((v82 != 18446744073709551615) || false)) {
                            v184 := GenS13{F0: 8, F1: 0}
                            _ = v184
                        } else {
                            v185 := []int{1, fi122}
                            if (true && true) {
                                v186 := GenS6{F0: -1, F1: ""}
                                switch ("\n\t" + ("hello" + ("foo bar" + "hello"))) {
                                case "foo bar":
                                    v187 := GenS8{F0: genFn0(-318), F1: "\n\t"}
                                    _ = v187
                                default:
                                    switch (fi122 << ((18446744073709551615) & 63)) {
                                    case -2147483648:
                                        v188 := GenS16{F0: -853, F1: (128 >= 52), F2: v4}
                                        _ = v188
                                    }
                                }
                                var v189 uint32 = (^(v83))
                                _ = v186
                                _ = v189
                            }
                            switch "\n\t" {
                            case "a":
                                if (!((fi122 >> ((18446744073709551615) & 63)) >= -2147483648)) {
                                    v4 = v4
                                    v4 = (v4 & -9223372036854775808)
                                } else {
                                    delete(v6, (("foo bar" + ("" + "a")) + "hello"))
                                    v190 := GenS12{F0: 255, F1: 265}
                                    v191 := GenS14{F0: 40}
                                    _ = v190
                                    _ = v191
                                }
                            case "":
                                if true {
                                    fmt.Println(v83)
                                    var v192 uint = 65535
                                    _ = v192
                                }
                            case "\n\t":
                                v193 := GenS2{F0: 1, F1: v4, F2: v82}
                                _ = v193
                            default:
                                var v194 uint32 = 0
                                _ = v194
                            }
                            _ = v185
                        }
                    }
                    var v195 int32 = 901
                    v196 := GenS17{F0: v1, F1: -26}
                    _ = v195
                    _ = v196
                }
                var v197 uint = 42
                v4 = 0
                _ = fi122
                _ = v197
            }
        }
    }
    v4 = -1
    for fi198 := 0; fi198 < 12; fi198++ {
        v199 := []int{9223372036854775807}
        switch ("foo bar" + (("a" + "") + "\n\t")) {
        case "foo bar":
            v4 -= (v4 ^ v4)
        }
        delete(v2, 1)
        _ = fi198
        _ = v199
    }
    var v200 uint64 = (v82 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v200:", v200)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
