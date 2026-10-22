package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint }
type GenS7 struct { F0 uint64; F1 int; F2 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16; F1 string; F2 int16 }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 string; F1 string }
func genFn0(p0 uint16) uint16 {
    p0 = (p0 % ((p0) | 1))
    var v4 float32 = float32(float64(float32(1e10)))
    _ = p0
    _ = v4
    return (p0 | p0)
}
func genFn1(p0 bool) uint32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int8 = -123
        if p0 {
            fmt.Println(v6, fi5, fi5)
            if p0 {
                var v7 uint = 255
                fmt.Println(v7, p0, p0)
                _ = v7
            } else {
                var v8 uint32 = 2147483647
                v9 := []int{(fi5 + fi5), fi5, (fi5 ^ 385)}
                _ = v8
                _ = v9
            }
            var v10 float32 = -2.25
            _ = v10
        } else {
            v11 := []int{}
            fmt.Println(v6, p0, p0)
            fmt.Println(fi5, fi5, fi5)
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    p0 = p0
    _ = p0
    return 1
}

func main() {
    v1 := []int{-2147483648, 6, 2147483647, -1}
    var v2 float64 = 1e10
    v3 := []uint32{}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 2, F1: -797}
        _ = fi12
        _ = v13
    }
    fmt.Println(v2)
    v14 := GenS2{F0: (("hello" + (("foo bar" + "foo bar") + "a")) + ("hello" + "foo bar"))}
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(fi15, v2)
        _ = fi15
    }
    if false {
        v2 = 3.14159
        for fi16 := 0; fi16 < 10; fi16++ {
            if false {
                v17 := GenS3{F0: 65535, F1: 1}
                switch "\n\t" {
                case "":
                    fmt.Println(fi16, v2)
                }
                _ = v17
            }
            var v18 bool = ((32767 >= 2) || (true && (true || true)))
            _ = fi16
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            v2 = v2
            var v20 int16 = -988
            _ = fi19
            _ = v20
        }
    } else {
        if (65535 == 544) {
            switch -2147483648 {
            case 127:
                v2 -= v2
            case 486:
                if (365 == 2) {
                    if len(v1) > 0 {
                        v1[0] = -9223372036854775808
                    }
                    v2 += ((v2 + (0.0 / float64(1.5))) + v2)
                }
            case 32768:
                var v21 uint32 = 737
                _ = v21
            default:
                if (true || (-1 <= 127)) {
                    if true {
                        v2 += (-2.25 + -2.25)
                        switch 128 {
                        case 1:
                            switch (("" + ("" + "hello")) + "foo bar") {
                            case "":
                                v2 = v2
                            case "foo bar":
                                v2 = (3.14159 - (v2 - v2))
                            case "hello":
                                v2 = 0.0
                            default:
                                var v22 string = ("foo bar" + "foo bar")
                                _ = v22
                            }
                        case 9223372036854775807:
                            if (551 == 32767) {
                                if ((!(65535 == 1)) && false) {
                                    var v23 float32 = float32((0.0 * v2))
                                    var v24 float64 = ((v2 / 1.0) / v2)
                                    if true {
                                        var v25 uint16 = 373
                                        _ = v25
                                    } else {
                                        v2 = v2
                                        v23 = (v23 + 1.5)
                                        if (!false) {
                                            _ = len(v3)
                                            v2 -= v2
                                            fmt.Println(v23, v2, v2)
                                        } else {
                                            fmt.Println(v24)
                                        }
                                    }
                                    _ = v23
                                    _ = v24
                                } else {
                                    if (((516 != 262) && (true || false)) && (!(42 <= 29))) {
                                        var v26 float64 = (v2 + v2)
                                        var v27 int8 = -1
                                        _ = v26
                                        _ = v27
                                    }
                                    v3 = append(v3, 4294967295)
                                }
                            } else {
                                v2 = v2
                            }
                        }
                        var v28 uint32 = 0
                        _ = v28
                    } else {
                        var v29 uint64 = 1
                        if (0.0 > 1.5) {
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v31 := GenS3{F0: 774, F1: -1}
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    if (((791 > 65535) || false) || (!(42 >= v29))) {
                                        var v33 uint32 = 42
                                        var v34 int64 = 9223372036854775807
                                        v2 = (1.5 + v2)
                                        var v35 bool = true
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    v2 = v2
                                    if len(v1) > 0 {
                                        v1[0] = fi30
                                    }
                                    _ = fi32
                                }
                                _ = fi30
                                _ = v31
                            }
                            v1 = append(v1, -9223372036854775808)
                        }
                        for fi36 := 0; fi36 < 12; fi36++ {
                            v37 := GenS4{F0: int(v29), F1: 2}
                            var v38 int32 = -751
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v29
                    }
                    v2 -= v2
                } else {
                    v2 = v2
                    if (!true) {
                        v2 += 0.1
                        if (-2.25 <= ((-2.25 * 1.5) * float32(((v2 / v2) + (v2 + v2))))) {
                            v1 = append(v1, 993)
                        } else {
                            for fi39 := 0; fi39 < 8; fi39++ {
                                if (42 >= 0) {
                                    v40 := GenS4{F0: (fi39 / (((fi39 - 0)) | 1)), F1: 65535}
                                    var v41 uint = 2
                                    if false {
                                        v42 := GenS5{F0: 4294967295}
                                        var v43 uint8 = 42
                                        switch fi39 {
                                        case 9223372036854775807:
                                            var v44 string = "a"
                                            _ = v44
                                        case 0:
                                            var v45 int32 = 1
                                            _ = v45
                                        case 128:
                                            for fi46 := 0; fi46 < 19; fi46++ {
                                                if true {
                                                    v47 := GenS3{F0: 42, F1: 0}
                                                    _ = v47
                                                }
                                                if false {
                                                    var v48 float32 = -2.25
                                                    if true {
                                                        v43 = 76
                                                        v48 += v48
                                                        switch "a" {
                                                        case "hello":
                                                            fmt.Println(v2, v2, v41)
                                                        case "foo bar":
                                                            v49 := GenS2{F0: "a"}
                                                            _ = v49
                                                        }
                                                    } else {
                                                        var v50 uint = 130
                                                        var v51 int8 = -128
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    _ = v48
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        v2 -= v2
                                                    case "foo bar":
                                                        switch 49 {
                                                        case 417:
                                                            fmt.Println(v43, v2)
                                                        case -533:
                                                            v41 = v41
                                                        default:
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v43 -= v43
                                                            case "\n\t":
                                                                v43 += v43
                                                            case "":
                                                                if (677 < 18446744073709551615) {
                                                                    var v52 int = fi46
                                                                    _ = v52
                                                                } else {
                                                                    var v53 int64 = -1
                                                                    var v54 uint32 = 0
                                                                    if false {
                                                                        fmt.Println(fi39, fi46)
                                                                        v41 = v41
                                                                        if false {
                                                                            v41 -= v41
                                                                            switch (^(fi39)) {
                                                                            case 4294967295:
                                                                                v2 = v2
                                                                            case -105:
                                                                                var v55 float32 = 3.14159
                                                                                _ = v55
                                                                            case -178:
                                                                                v56 := GenS4{F0: 32768, F1: 52}
                                                                                _ = v56
                                                                            }
                                                                            if true {
                                                                                var v57 int16 = 47
                                                                                v57 = (v57 - 1)
                                                                                var v58 uint16 = 845
                                                                                _ = v57
                                                                                _ = v58
                                                                            }
                                                                        }
                                                                    }
                                                                    switch fi39 {
                                                                    case -286:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v59 int = (fi39 / ((fi46) | 1))
                                                        _ = v59
                                                    default:
                                                        if false {
                                                            v41 = v41
                                                            var v60 int32 = 347
                                                            if false {
                                                                if (v41 < 0) {
                                                                    v60 = v60
                                                                }
                                                            } else {
                                                                v60 -= v60
                                                            }
                                                            _ = v60
                                                        } else {
                                                            fmt.Println(v41)
                                                            var v61 uint8 = v43
                                                            var v62 string = ("" + ("a" + "hello"))
                                                            v41 -= v41
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    }
                                                }
                                                switch fi46 {
                                                case -136:
                                                    var v63 int8 = -117
                                                    _ = v63
                                                default:
                                                    var v64 string = "hello"
                                                    _ = v64
                                                }
                                                _ = fi46
                                            }
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v40
                                    _ = v41
                                } else {
                                    var v65 uint32 = 65535
                                    if ((!true) && true) {
                                        if true {
                                            fmt.Println(v2, v2, v2)
                                        } else {
                                            v66 := GenS3{F0: 2, F1: 42}
                                            _ = v66
                                        }
                                        var v67 uint16 = 664
                                        _ = v67
                                    } else {
                                        var v68 float32 = 3.14159
                                        for fi69 := 0; fi69 < 5; fi69++ {
                                            switch (("\n\t" + "") + "foo bar") {
                                            case "":
                                                v70 := GenS3{F0: 4294967295, F1: 616}
                                                _ = v70
                                            case "\n\t":
                                                fmt.Println(v65, v65)
                                            default:
                                                var v71 int64 = 0
                                                _ = v71
                                            }
                                            if true {
                                                var v72 uint64 = 76
                                                var v73 int64 = 531
                                                var v74 uint16 = 255
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            } else {
                                                v75 := GenS6{F0: genFn1((0 <= 65535)), F1: 624, F2: 0}
                                                _ = v75
                                            }
                                            v68 += 0.0
                                            v76 := GenS4{F0: 27, F1: 0}
                                            v77 := GenS7{F0: 0, F1: fi39, F2: (v65 / (((v65 / (((v65 + (v65 % ((v65) | 1)))) | 1))) | 1))}
                                            v78 := GenS6{F0: v65, F1: uint64(v65), F2: 18446744073709551615}
                                            var v79 float32 = 1.5
                                            _ = fi69
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                            _ = v79
                                        }
                                        var v80 uint64 = 867
                                        _ = v68
                                        _ = v80
                                    }
                                    v2 -= ((v2 + (3.14159 + (v2 * v2))) / v2)
                                    _ = v65
                                }
                                var v81 float64 = float64(1.5)
                                _ = fi39
                                _ = v81
                            }
                        }
                        v2 = (1e10 * -2.25)
                    } else {
                        switch 32767 {
                        case 32768:
                            v2 -= -2.25
                        default:
                            if len(v1) > 0 {
                                v1[0] = 128
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = genFn1(true)
                        }
                        var v82 bool = ((!(!false)) && (!true))
                        _ = v82
                    }
                }
            }
            var v83 int32 = -1
            _ = v83
        } else {
            if true {
                switch 32768 {
                case 1:
                    switch "a" {
                    case "hello":
                        v3 = append(v3, 591)
                    default:
                        switch ("foo bar" + "foo bar") {
                        case "foo bar":
                            v2 = ((v2 - float64(1e10)) / (v2 * v2))
                        case "hello":
                            v1 = append(v1, 127)
                        case "a":
                            var v84 int = -9223372036854775808
                            _ = v84
                        default:
                            switch 32767 {
                            case 9223372036854775807:
                                v2 = float64((-2.25 / 1.0))
                            case 32768:
                                v3 = append(v3, 145)
                            case 127:
                                var v85 bool = false
                                _ = v85
                            default:
                                switch 128 {
                                case 285:
                                    v86 := GenS3{F0: 255, F1: 32767}
                                    _ = v86
                                default:
                                    for fi87 := 0; fi87 < 15; fi87++ {
                                        v88 := GenS5{F0: genFn1((0 != -1))}
                                        _ = fi87
                                        _ = v88
                                    }
                                }
                            }
                        }
                    }
                case 939:
                    fmt.Println(v2)
                case 32768:
                    var v89 float64 = float64((3.14159 + -2.25))
                    _ = v89
                default:
                    for fi90 := 0; fi90 < 15; fi90++ {
                        if true {
                            for fi91 := 0; fi91 < 2; fi91++ {
                                v92 := GenS1{F0: 42, F1: -9223372036854775808}
                                var v93 int8 = 127
                                _ = fi91
                                _ = v92
                                _ = v93
                            }
                            var v94 int = fi90
                            v94 = (v94 / (((^(fi90))) | 1))
                            _ = v94
                        } else {
                            switch (("\n\t" + "") + ((("\n\t" + "") + "") + (("foo bar" + "hello") + ("\n\t" + ("foo bar" + "foo bar"))))) {
                            case "foo bar":
                                v2 += v2
                            default:
                                var v95 uint = 18446744073709551615
                                _ = v95
                            }
                            var v96 int32 = -144
                            switch fi90 {
                            case -858:
                                if true {
                                    var v97 uint64 = 4294967295
                                    for fi98 := 0; fi98 < 7; fi98++ {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = fi98
                                    }
                                    _ = v97
                                }
                            case 128:
                                for fi99 := 0; fi99 < 7; fi99++ {
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        if (!true) {
                                            v3 = append(v3, 500)
                                            v2 -= v2
                                            v100 := GenS1{F0: 42, F1: 9223372036854775807}
                                            _ = v100
                                        } else {
                                            switch fi99 {
                                            case -694:
                                                _ = len(v1)
                                            default:
                                                if ((!(972 == 4294967295)) || false) {
                                                    var v101 bool = true
                                                    v96 -= 0
                                                    var v102 int = fi90
                                                    _ = v101
                                                    _ = v102
                                                }
                                            }
                                            fmt.Println(v96)
                                            v103 := GenS1{F0: 42, F1: 0}
                                            _ = v103
                                        }
                                    }
                                    var v104 int8 = 54
                                    switch fi90 {
                                    case 9223372036854775807:
                                        var v105 int16 = 1
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (!true) {
                                            if len(v1) > 0 {
                                                v1[0] = (fi90 - fi90)
                                            }
                                            var v106 float32 = 1.5
                                            v105 += (v105 - v105)
                                            _ = v106
                                        }
                                        _ = v105
                                    case -2147483648:
                                        switch "a" {
                                        case "foo bar":
                                            if false {
                                                v2 -= 1e10
                                                fmt.Println(v104, v96, v104)
                                                v107 := GenS6{F0: 1, F1: 1, F2: 777}
                                                _ = v107
                                            } else {
                                                var v108 uint32 = genFn1((!true))
                                                _ = len(v1)
                                                _ = v108
                                            }
                                        case "\n\t":
                                            if false {
                                                if (1 == v104) {
                                                    _ = len(v1)
                                                }
                                                if (false || false) {
                                                    var v109 int64 = -9223372036854775808
                                                    var v110 float64 = (0.0 * v2)
                                                    _ = v109
                                                    _ = v110
                                                }
                                                v111 := GenS3{F0: 255, F1: 1}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v111
                                            }
                                        case "a":
                                            if true {
                                                v96 -= (v96 * v96)
                                            }
                                        default:
                                            if false {
                                                v112 := GenS8{F0: 128}
                                                v104 -= v104
                                                _ = v112
                                            }
                                        }
                                    case -646:
                                        if ((!(!((!((int(v96) < (fi90 * 128)) || (424 < 239))) && false))) && false) {
                                            var v113 int16 = 0
                                            _ = v113
                                        }
                                    }
                                    _ = fi99
                                    _ = v104
                                }
                            }
                            _ = v96
                        }
                        if true {
                            v114 := GenS9{F0: 65535, F1: ("hello" + "foo bar"), F2: 0}
                            v3 = append(v3, 1)
                            v115 := GenS10{F0: "hello", F1: 100}
                            _ = v114
                            _ = v115
                        }
                        _ = fi90
                    }
                }
                _ = len(v1)
                switch -1 {
                case 946:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32767:
                    v2 += (((v2 - float64(float32(v2))) / v2) + 1.5)
                case 0:
                    fmt.Println(v2)
                default:
                    switch "a" {
                    case "foo bar":
                        if ((!false) || false) {
                            v2 = (v2 * v2)
                            v2 += (float64(1.5) / (3.14159 + (float64(((1.5 * 1e10) - (3.14159 - ((-2.25 / (float32(-2.25) + 1.5)) / -2.25)))) - v2)))
                        }
                    default:
                        if false {
                            if (180 > 22) {
                                var v116 uint32 = genFn1(true)
                                v117 := GenS9{F0: 209, F1: "\n\t", F2: 335}
                                _ = v116
                                _ = v117
                            } else {
                                var v118 int64 = -9223372036854775808
                                v119 := GenS4{F0: 1, F1: 65535}
                                _ = v118
                                _ = v119
                            }
                            var v120 int64 = 982
                            var v121 bool = false
                            _ = v120
                            _ = v121
                        } else {
                            for fi122 := 0; fi122 < 2; fi122++ {
                                var v123 int = (fi122 << ((1) & 63))
                                _ = fi122
                                _ = v123
                            }
                            fmt.Println(v2)
                        }
                    }
                }
            } else {
                var v124 uint8 = 123
                switch "hello" {
                case "hello":
                    for fi125 := 0; fi125 < 2; fi125++ {
                        v124 += v124
                        _ = fi125
                    }
                case "foo bar":
                    var v126 uint8 = (v124 | v124)
                    _ = v126
                default:
                    var v127 int64 = 0
                    _ = v127
                }
                _ = v124
            }
            v2 = v2
        }
        v128 := GenS11{F0: ("\n\t" + ""), F1: "a"}
        _ = v128
    }
    v2 += float64(float32(v2))
    fmt.Println(v2)
    for fi129 := 0; fi129 < 8; fi129++ {
        v2 = v2
        _ = fi129
    }
    switch -34 {
    case -2147483648:
        switch "foo bar" {
        case "a":
            if len(v1) > 0 {
                v1[0] = 128
            }
        case "\n\t":
            fmt.Println(v2)
        default:
            if (1e10 >= v2) {
                if true {
                    v2 += float64(-2.25)
                    var v130 string = ""
                    fmt.Println(v130, v2)
                    _ = v130
                }
                _ = len(v1)
                var v131 float64 = (float64((3.14159 * -2.25)) + v2)
                _ = v131
            }
        }
    case 32767:
        fmt.Println(v2)
    case 562:
        v2 += (3.14159 * 1e10)
    }
    v1 = append(v1, 9223372036854775807)
    fmt.Println(v2)
    for fi132 := 0; fi132 < 9; fi132++ {
        var v133 int64 = 42
        var v134 bool = ((!(!(true && (!false)))) || ((fi132 ^ fi132) < 128))
        _ = fi132
        _ = v133
        _ = v134
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
