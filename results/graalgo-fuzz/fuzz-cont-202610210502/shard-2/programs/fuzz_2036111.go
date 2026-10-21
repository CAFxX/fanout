package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int32; F1 int64; F2 int32 }
type GenS6 struct { F0 int32; F1 int; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 int8; F2 int16 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    var v5 int8 = 11
    switch "a" {
    case "foo bar":
        v6 := v2[127]
        _ = v6
    }
    fmt.Println(v3, v3)
    v1 = v1
    v7 := GenS1{F0: -769}
    _ = len(v4)
    if (((3.14159 / 3.14159) / float64(1.5)) != float64(1e10)) {
        v5 += 100
        var v8 int32 = -2147483648
        v9 := GenS1{F0: (^(v8))}
        _ = v8
        _ = v9
    } else {
        v1 = 127
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            v4[931] = false
        } else {
            v1 -= v1
            for fi11 := 0; fi11 < 2; fi11++ {
                if false {
                    if true {
                        if true {
                            delete(v2, fi10)
                            if false {
                                v1 -= 42
                                v12 := GenS1{F0: -1}
                                if false {
                                    var v13 bool = (843 <= 42)
                                    if v13 {
                                        _ = len(v2)
                                    } else {
                                        var v14 int16 = -131
                                        v15 := []int{fi10}
                                        _ = v14
                                        _ = v15
                                    }
                                    switch fi10 {
                                    case -9223372036854775808:
                                        v16 := []int{fi10, -2147483648}
                                        _ = v16
                                    case 127:
                                        switch (fi11 + (^(fi10))) {
                                        case 668:
                                            var v17 int64 = -9223372036854775808
                                            _ = v17
                                        case 42:
                                            v3 = 42
                                        case 1:
                                            if (("a" + "\n\t") > "") {
                                                v1 = (v3 + (^(v3)))
                                                var v18 uint16 = 2
                                                if (v13 && v13) {
                                                    v19 := GenS2{F0: v18, F1: 2}
                                                    v20 := GenS3{F0: v18, F1: 42, F2: 42}
                                                    _ = v19
                                                    _ = v20
                                                }
                                                _ = v18
                                            } else {
                                                switch "hello" {
                                                case "a":
                                                    if true {
                                                        v2[(fi10 >> ((551) & 63))] = 65535
                                                        if v13 {
                                                            v13 = v13
                                                            v4[1] = ((!(((^(v3)) >= (v3 << ((4294967295) & 63))) && v13)) || ((v3 * v3) == 42))
                                                            var v21 int8 = 100
                                                            _ = v21
                                                        }
                                                    } else {
                                                        v1 -= (v3 ^ v3)
                                                        if v13 {
                                                            v13 = (v13 && true)
                                                        } else {
                                                            var v22 int8 = (v5 * (v5 >> ((4294967295) & 63)))
                                                            if true {
                                                                var v23 float32 = (1.5 * 0.0)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v5 += v5
                                                                case "":
                                                                    var v24 int = -1
                                                                    _ = v24
                                                                case "hello":
                                                                    if v13 {
                                                                        switch (fi11 << ((344) & 63)) {
                                                                        case 0:
                                                                            var v25 int8 = (v22 & (v22 ^ (v22 & -92)))
                                                                            _ = v25
                                                                        }
                                                                        v26 := GenS1{F0: 209}
                                                                        v27 := GenS1{F0: 2147483647}
                                                                        _ = v26
                                                                        _ = v27
                                                                    } else {
                                                                        var v28 int64 = 42
                                                                        _ = v28
                                                                    }
                                                                default:
                                                                    v23 = (((float32(1e10) + v23) / 0.1) / v23)
                                                                }
                                                                switch fi11 {
                                                                case 32767:
                                                                    v5 -= 1
                                                                case 42:
                                                                    if true {
                                                                        v29 := []int{fi10, fi10}
                                                                        _ = v29
                                                                    } else {
                                                                        v30 := GenS1{F0: -792}
                                                                        v22 = (v22 / ((v5) | 1))
                                                                        _ = v30
                                                                    }
                                                                case -2147483648:
                                                                    var v31 uint8 = v3
                                                                    _ = v31
                                                                }
                                                                _ = v23
                                                            }
                                                            var v32 uint8 = v1
                                                            _ = v22
                                                            _ = v32
                                                        }
                                                    }
                                                }
                                                if v13 {
                                                    switch -1 {
                                                    case 9223372036854775807:
                                                        if false {
                                                            v33 := v2[-615]
                                                            _ = v33
                                                        } else {
                                                            if (!v13) {
                                                                v5 = v5
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v34 float32 = ((float32((-2.25 * float64((3.14159 + (-2.25 + 0.1))))) * (1e10 / (((-2.25 * 0.0) - 3.14159) * -2.25))) / -2.25)
                                                                    _ = v34
                                                                }
                                                            } else {
                                                                v35 := GenS4{F0: "a", F1: 868}
                                                                if ("hello" >= "foo bar") {
                                                                    v13 = ((!v13) && v13)
                                                                    var v36 int = fi11
                                                                    _ = v36
                                                                }
                                                                _ = v35
                                                            }
                                                        }
                                                    case 0:
                                                        v13 = (v13 && v13)
                                                    default:
                                                        var v37 uint64 = 18446744073709551615
                                                        if ((1.5 > -2.25) || v13) {
                                                            fmt.Println(fi10, v13)
                                                            var v38 uint64 = 965
                                                            v39 := []int{(fi11 % ((-2147483648) | 1)), fi11}
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            v37 = v37
                                                            v37 = (v37 & (v37 & v37))
                                                            v5 += 127
                                                        }
                                                        _ = v37
                                                    }
                                                    var v40 int = (fi10 / ((fi10) | 1))
                                                    _ = v40
                                                }
                                                if (940 < 0) {
                                                    var v41 string = "\n\t"
                                                    var v42 uint32 = 1
                                                    if (v13 || v13) {
                                                        if v13 {
                                                            switch v41 {
                                                            case "foo bar":
                                                                v5 = v5
                                                            case "hello":
                                                                v43 := GenS3{F0: 0, F1: 2147483647, F2: 18446744073709551615}
                                                                _ = v43
                                                            }
                                                            var v44 int16 = -1
                                                            v45 := GenS5{F0: 756, F1: 42, F2: 0}
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            if (!(2 > 969)) {
                                                                switch ("foo bar" + v41) {
                                                                case "\n\t":
                                                                    v42 = 42
                                                                }
                                                                if (v42 != v42) {
                                                                    fmt.Println(v42, fi10, fi10)
                                                                } else {
                                                                    delete(v2, (^(fi11)))
                                                                }
                                                                fmt.Println(v41)
                                                                if (243 < 42) {
                                                                    if (42 >= 292) {
                                                                        if (v41 >= "a") {
                                                                            v3 = (^(v1))
                                                                        } else {
                                                                            if ((((("\n\t" + "\n\t") + "hello") > "hello") || (!v13)) || v13) {
                                                                                var v46 int64 = -9223372036854775808
                                                                                _ = v46
                                                                            } else {
                                                                                var v47 uint8 = v1
                                                                                _ = v47
                                                                            }
                                                                            var v48 float64 = 0.0
                                                                            if false {
                                                                                var v49 uint32 = 634
                                                                                v3 += v3
                                                                                v41 = "hello"
                                                                                v50 := GenS6{F0: 0, F1: (fi10 % (((-(fi11))) | 1)), F2: 426}
                                                                                v51 := []int{(fi10 | -2147483648), (fi11 / ((32767) | 1))}
                                                                                v52 := GenS7{F0: -32768}
                                                                                _ = v49
                                                                                _ = v50
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                            _ = v48
                                                                        }
                                                                        v13 = (0 == -500)
                                                                    } else {
                                                                        v13 = (!v13)
                                                                        v53 := []int{(fi10 - fi10), fi10}
                                                                        _ = v53
                                                                    }
                                                                }
                                                            } else {
                                                                var v54 uint16 = 65535
                                                                switch fi10 {
                                                                case 9223372036854775807:
                                                                    if v13 {
                                                                        switch fi10 {
                                                                        case 128:
                                                                            var v55 uint64 = 1
                                                                            _ = v55
                                                                        case 1:
                                                                            var v56 uint16 = 255
                                                                            _ = v56
                                                                        case 2147483647:
                                                                            delete(v4, 1)
                                                                        default:
                                                                            _ = len(v2)
                                                                        }
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v57 int8 = 42
                                                                    _ = v57
                                                                case -110:
                                                                    var v58 int32 = -1
                                                                    _ = v58
                                                                }
                                                                _ = v54
                                                            }
                                                            v42 = v42
                                                        }
                                                        fmt.Println(v13)
                                                        v59 := []int{}
                                                        var v60 int64 = 0
                                                        var v61 int = fi11
                                                        _ = v59
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    v2[(fi11 * fi11)] = 1
                                                    if (v13 || v13) {
                                                        var v62 uint64 = 213
                                                        _ = v62
                                                    }
                                                    _ = v41
                                                    _ = v42
                                                }
                                                fmt.Println(v5)
                                            }
                                        default:
                                            if v13 {
                                                fmt.Println(v5, fi10, v5)
                                                fmt.Println(v3, fi10, v1)
                                                fmt.Println(v5, v1)
                                            } else {
                                                v63 := []int{}
                                                v3 = (v1 % ((175) | 1))
                                                if (false && (true || v13)) {
                                                    v64 := GenS8{F0: ("" + "hello"), F1: 65535}
                                                    _ = v64
                                                } else {
                                                    if v13 {
                                                        v5 = v5
                                                        fmt.Println(v3, v3, fi10)
                                                    }
                                                }
                                                v63 = append(v63, fi10)
                                                _ = v63
                                            }
                                        }
                                    case -1:
                                        v65 := GenS9{F0: (v5 >> ((764) & 63))}
                                        _ = v65
                                    }
                                    v66 := GenS8{F0: ("\n\t" + ("" + "hello")), F1: 32767}
                                    switch "hello" {
                                    case "a":
                                        if v13 {
                                            v67 := GenS2{F0: 65535, F1: 1}
                                            if v13 {
                                                if (!(255 == (v3 - (v3 % (((v3 << ((42) & 63))) | 1))))) {
                                                    v68 := GenS4{F0: "a", F1: 42}
                                                    v69 := GenS10{F0: -375}
                                                    var v70 float32 = ((0.1 / (float32(1e10) + 1.5)) - ((1e10 + (float32(1.5) + 1e10)) + 0.1))
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    fmt.Println(v13)
                                                    v71 := []int{fi10, 573, -2147483648, fi11}
                                                    _ = v71
                                                }
                                            }
                                            _ = v67
                                        }
                                    }
                                    v72 := GenS2{F0: 191, F1: 505}
                                    _ = v13
                                    _ = v66
                                    _ = v72
                                } else {
                                    _ = len(v2)
                                    var v73 int32 = -1
                                    v4[2147483648] = (true && true)
                                    _ = v73
                                }
                                v74 := GenS2{F0: 65535, F1: 2}
                                var v75 uint32 = 1
                                _ = v12
                                _ = v74
                                _ = v75
                            }
                            if true {
                                switch fi11 {
                                case -1:
                                    var v76 float64 = (((1.5 * -2.25) - 1.5) / ((float64((1e10 / 0.1)) + 3.14159) + (1e10 * (1.5 * 3.14159))))
                                    _ = v76
                                }
                            } else {
                                if (!(!(2147483647 <= 938))) {
                                    v5 = v5
                                } else {
                                    fmt.Println(v5, v1)
                                    if true {
                                        var v77 int8 = (v5 >> ((654) & 63))
                                        if (42 >= 18446744073709551615) {
                                            v78 := []int{(-(fi10)), fi10}
                                            _ = v78
                                        } else {
                                            v79 := GenS8{F0: "a", F1: 795}
                                            if true {
                                                switch (-(fi11)) {
                                                case 128:
                                                    delete(v2, fi10)
                                                case 32767:
                                                    var v80 int = fi11
                                                    _ = v80
                                                default:
                                                    v81 := GenS2{F0: 1, F1: 478}
                                                    _ = v81
                                                }
                                                var v82 uint = 18446744073709551615
                                                _ = v82
                                            } else {
                                                if false {
                                                    v77 -= (v5 ^ v5)
                                                } else {
                                                    var v83 uint64 = 2
                                                    v84 := v4[550]
                                                    v85 := []int{fi11, (-(fi11)), -2147483648}
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                }
                                            }
                                            _ = v79
                                        }
                                        _ = v77
                                    }
                                    var v86 bool = true
                                    _ = v86
                                }
                                v87 := GenS10{F0: 0}
                                fmt.Println(fi11, v3, fi10)
                                _ = v87
                            }
                        }
                        if false {
                            fmt.Println(v1, v5)
                            var v88 string = ("hello" + "\n\t")
                            _ = v88
                        } else {
                            var v89 uint = 44
                            v1 += v1
                            var v90 uint64 = 4294967295
                            _ = v89
                            _ = v90
                        }
                    } else {
                        if (3.14159 < float64(1.5)) {
                            if true {
                                var v91 uint64 = 1
                                fmt.Println(v91)
                                fmt.Println(v5)
                                v92 := GenS3{F0: 32767, F1: -2147483648, F2: 18446744073709551615}
                                _ = v91
                                _ = v92
                            } else {
                                if true {
                                    v2[9223372036854775807] = 4294967295
                                } else {
                                    if true {
                                        var v93 uint32 = 4294967295
                                        switch 187 {
                                        case 9223372036854775807:
                                            var v94 int = 32768
                                            _ = v94
                                        case 0:
                                            var v95 int64 = 1
                                            _ = v95
                                        case 990:
                                            switch "hello" {
                                            case "a":
                                                v96 := []int{(fi10 + fi11)}
                                                _ = v96
                                            case "foo bar":
                                                v97 := []int{fi11, fi10, fi10}
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch fi11 {
                                                    case 128:
                                                        v5 -= v5
                                                    case -2147483648:
                                                        if (0 == 2) {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v1 = 255
                                                            case "a":
                                                                fmt.Println(fi11)
                                                                fmt.Println(v3, v1)
                                                                v98 := GenS11{F0: (v3 | 0)}
                                                                _ = v98
                                                            default:
                                                                v3 -= v3
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "a":
                                                            switch "foo bar" {
                                                            case "":
                                                                v99 := v4[(v93 / (((v93 & 42)) | 1))]
                                                                _ = v99
                                                            case "\n\t":
                                                                var v100 uint = 746
                                                                _ = v100
                                                            case "foo bar":
                                                                delete(v2, -2147483648)
                                                            }
                                                        case "\n\t":
                                                            if (-200 >= 2147483647) {
                                                                var v101 uint16 = 42
                                                                v102 := GenS8{F0: ("foo bar" + "foo bar"), F1: uint16(v1)}
                                                                _ = v101
                                                                _ = v102
                                                            }
                                                        }
                                                        v3 = (v1 + (v1 + (v1 - (v3 - (^(v1))))))
                                                    case 2147483647:
                                                        v103 := GenS5{F0: -2147483648, F1: 42, F2: 373}
                                                        _ = v103
                                                    default:
                                                        v104 := GenS6{F0: -602, F1: (fi11 - -444), F2: 1000}
                                                        _ = v104
                                                    }
                                                case "hello":
                                                    v3 += v3
                                                case "":
                                                    fmt.Println(fi11, v93)
                                                }
                                                v3 += 1
                                                var v105 uint64 = 18446744073709551615
                                                _ = v105
                                                _ = v97
                                            case "":
                                                if (int8(fi10) > 92) {
                                                    v106 := GenS8{F0: "hello", F1: 0}
                                                    if (false || (1 >= 4294967295)) {
                                                        var v107 uint8 = 85
                                                        fmt.Println(v1, v5, fi11)
                                                        v93 = v93
                                                        _ = v107
                                                    } else {
                                                        if ((-(fi10)) != fi10) {
                                                            var v108 float64 = (1e10 / (0.1 / -2.25))
                                                            _ = v108
                                                        } else {
                                                            var v109 float64 = -2.25
                                                            if (-9223372036854775808 < 42) {
                                                                var v110 uint8 = 28
                                                                fmt.Println(v109, fi10)
                                                                _ = v110
                                                            }
                                                            _ = v109
                                                        }
                                                        v3 -= v3
                                                        if (("a" >= "hello") || true) {
                                                            switch ("\n\t" + "") {
                                                            case "a":
                                                                var v111 int16 = -110
                                                                _ = v111
                                                            case "hello":
                                                                v93 = v93
                                                            case "":
                                                                if (!true) {
                                                                    var v112 uint = 18446744073709551615
                                                                    _ = v112
                                                                }
                                                            }
                                                        } else {
                                                            v113 := GenS1{F0: 42}
                                                            _ = v113
                                                        }
                                                    }
                                                    _ = v106
                                                }
                                            default:
                                                delete(v4, v93)
                                            }
                                        default:
                                            delete(v4, v93)
                                        }
                                        _ = v93
                                    } else {
                                        fmt.Println(fi11, v5)
                                    }
                                    var v114 string = "foo bar"
                                    var v115 uint32 = 42
                                    _ = v114
                                    _ = v115
                                }
                                v1 += v1
                            }
                            var v116 uint32 = 6
                            _ = v116
                        } else {
                            v1 += v1
                            var v117 string = ""
                            switch v117 {
                            case "a":
                                v1 -= v3
                            default:
                                var v118 int = fi11
                                _ = v118
                            }
                            _ = v117
                        }
                        _ = len(v4)
                    }
                    var v119 int = fi10
                    var v120 float32 = ((1.5 + (0.0 + (-2.25 / ((0.0 * 1.5) - (-2.25 + -2.25))))) + float32(1e10))
                    v119 = 2147483647
                    switch (fi10 % ((fi10) | 1)) {
                    case 32767:
                        var v121 uint = 723
                        _ = v121
                    case 1:
                        switch ((("hello" + ("" + ("\n\t" + ("" + (("foo bar" + "") + "hello"))))) + ("foo bar" + "")) + "") {
                        case "\n\t":
                            v3 += (^(v1))
                        default:
                            var v122 int32 = -2147483648
                            _ = v122
                        }
                    case 128:
                        fmt.Println(v3, fi11, v1)
                    default:
                        var v123 int16 = -416
                        _ = v123
                    }
                    _ = v119
                    _ = v120
                }
                _ = fi11
            }
            v124 := GenS12{F0: true, F1: -49, F2: 516}
            _ = v124
        }
        v125 := []int{(-(fi10)), (fi10 & fi10)}
        _ = fi10
        _ = v125
    }
    var v126 int = 128
    switch 4294967295 {
    case -9223372036854775808:
        for fi127 := 0; fi127 < 6; fi127++ {
            v128 := []int{(^(fi127)), 32767}
            switch ("a" + "foo bar") {
            case "hello":
                if ((v126 + 1) < 128) {
                    for fi129 := 0; fi129 < 19; fi129++ {
                        if true {
                            delete(v2, -539)
                        }
                        _ = fi129
                    }
                    if false {
                        var v130 int32 = -632
                        for fi131 := 0; fi131 < 4; fi131++ {
                            var v132 int8 = v5
                            v5 = -128
                            var v133 bool = true
                            _ = fi131
                            _ = v132
                            _ = v133
                        }
                        switch "" {
                        case "hello":
                            var v134 int32 = (^(v130))
                            _ = v134
                        default:
                            _ = len(v128)
                        }
                        _ = v130
                    }
                    v135 := GenS4{F0: ("foo bar" + ("hello" + "foo bar")), F1: uint(v3)}
                    _ = v135
                } else {
                    var v136 float32 = 3.14159
                    v137 := GenS7{F0: 0}
                    _ = v136
                    _ = v137
                }
            case "":
                _ = len(v2)
            case "a":
                v3 = v1
            }
            v2[(v126 & v126)] = 531
            _ = fi127
            _ = v128
        }
    case 2147483647:
        var v138 string = "foo bar"
        _ = v138
    case -2147483648:
        if false {
            var v139 int16 = -310
            var v140 string = ("hello" + "foo bar")
            switch v140 {
            case "":
                v141 := []int{}
                _ = v141
            case "a":
                switch v140 {
                case "a":
                    var v142 uint8 = v1
                    _ = v142
                case "foo bar":
                    fmt.Println(v139, v126, v126)
                case "\n\t":
                    switch 9223372036854775807 {
                    case 328:
                        var v143 string = ("" + (v140 + v140))
                        _ = v143
                    }
                }
            case "foo bar":
                v1 += (v3 ^ v3)
            }
            _ = v139
            _ = v140
        } else {
            var v144 uint = 4294967295
            v144 += v144
            _ = v144
        }
    default:
        if false {
            delete(v4, 677)
        }
    }
    v4[118] = true
    var v145 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v126
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
