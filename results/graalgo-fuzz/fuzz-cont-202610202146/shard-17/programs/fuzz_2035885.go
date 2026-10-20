package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS5 struct { F0 string; F1 uint16; F2 bool }
type GenS6 struct { F0 string; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int32; F1 int32; F2 uint32 }
type GenS11 struct { F0 bool }
func genFn0(p0 int8) int32 {
    if true {
        p0 += p0
        for fi6 := 0; fi6 < 19; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                switch (fi7 - fi6) {
                case 968:
                    var v8 uint64 = 212
                    _ = v8
                case -430:
                    var v9 bool = (-32768 < 1000)
                    _ = v9
                case -9223372036854775808:
                    var v10 float64 = 0.1
                    _ = v10
                default:
                    var v11 float64 = ((((0.1 / 1e10) / (-2.25 * 3.14159)) * 0.1) / float64(1e10))
                    if (false || (fi7 <= fi6)) {
                        v11 = v11
                        var v12 int32 = 2147483647
                        v13 := GenS1{F0: v12, F1: (p0 | p0)}
                        var v14 uint = 1
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 uint32 = 2147483647
                    _ = v11
                    _ = v15
                }
                _ = fi7
            }
            _ = fi6
        }
    } else {
        switch 128 {
        case 144:
            if true {
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 int = -1
                    v17 = v17
                    _ = fi16
                    _ = v17
                }
                if true {
                    fmt.Println(p0)
                    v18 := GenS1{F0: 2147483647, F1: p0}
                    switch -384 {
                    case -2147483648:
                        v19 := GenS2{F0: "foo bar"}
                        _ = v19
                    case 42:
                        v20 := []int{479, 2147483647}
                        _ = v20
                    default:
                        for fi21 := 0; fi21 < 13; fi21++ {
                            fmt.Println(fi21, p0)
                            v22 := GenS1{F0: -39, F1: -128}
                            _ = fi21
                            _ = v22
                        }
                    }
                    _ = v18
                } else {
                    var v23 float32 = 3.14159
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 string = "\n\t"
                        v26 := []int{(-(fi24)), -717}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    if (false || false) {
                        v27 := []int{-856, 9223372036854775807}
                        var v28 int32 = -103
                        _ = v27
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            var v30 bool = (32767 <= 902)
                            _ = fi29
                            _ = v30
                        }
                        var v31 string = "\n\t"
                        if (187 > 18446744073709551615) {
                            p0 -= 82
                        }
                        _ = v31
                    }
                    _ = v23
                }
            }
        case 447:
            v32 := []int{32768, 9223372036854775807, -1}
            _ = v32
        default:
            var v33 uint = 389
            _ = v33
        }
        v34 := GenS1{F0: -1, F1: (p0 + (p0 - 0))}
        if true {
            switch "" {
            case "":
                for fi35 := 0; fi35 < 11; fi35++ {
                    p0 = p0
                    _ = fi35
                }
            case "a":
                v36 := GenS1{F0: 203, F1: p0}
                _ = v36
            }
        } else {
            var v37 int16 = 480
            if (false && false) {
                fmt.Println(v37, v37)
                if ((-(v37)) != (v37 & 42)) {
                    var v38 float64 = 0.0
                    _ = v38
                }
                var v39 uint8 = 0
                _ = v39
            } else {
                if (914 < 4294967295) {
                    p0 = -74
                    p0 = -31
                } else {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            v42 := []int{(fi41 + (fi40 << ((2) & 63)))}
                            _ = fi41
                            _ = v42
                        }
                        v43 := GenS3{F0: -775, F1: 1}
                        _ = fi40
                        _ = v43
                    }
                    if true {
                        v44 := GenS2{F0: ("a" + "foo bar")}
                        p0 -= p0
                        _ = v44
                    }
                }
            }
            if (("hello" == "foo bar") && false) {
                var v45 uint64 = 308
                _ = v45
            }
            _ = v37
        }
        _ = v34
    }
    v46 := GenS4{F0: 38, F1: 1, F2: 181}
    if (9223372036854775807 < -949) {
        v47 := []int{346, 4294967295, 20}
        var v48 float64 = 1.5
        fmt.Println(v48, p0)
        _ = v47
        _ = v48
    } else {
        var v49 int64 = 500
        v50 := GenS1{F0: 42, F1: p0}
        _ = v49
        _ = v50
    }
    switch 4294967295 {
    case 128:
        v51 := []int{-328}
        _ = v51
    default:
        p0 -= (p0 * (p0 & p0))
    }
    _ = p0
    _ = v46
    return -1
}

func main() {
    var v1 int64 = 0
    var v2 uint32 = 4294967295
    v3 := make(map[string]uint32)
    var v4 uint8 = 255
    var v5 uint = 119
    var v52 int16 = 1
    for fi53 := 0; fi53 < 15; fi53++ {
        var v54 float64 = 1e10
        switch "a" {
        case "":
            v1 -= (-(v1))
        default:
            v55 := v3[""]
            _ = v55
        }
        _ = fi53
        _ = v54
    }
    if false {
        fmt.Println(v1)
        _ = len(v3)
    }
    if true {
        if (0 < v2) {
            if ((v4 - 42) <= v4) {
                fmt.Println(v5, v5)
            } else {
                if (!(true && false)) {
                    var v56 uint64 = 2
                    v57 := []int{128}
                    var v58 uint = (v5 | v5)
                    _ = v56
                    _ = v57
                    _ = v58
                } else {
                    var v59 uint8 = v4
                    _ = v59
                }
            }
            var v60 uint8 = 0
            switch 9223372036854775807 {
            case -573:
                v61 := GenS1{F0: genFn0(-128), F1: -128}
                _ = v61
            case -824:
                var v62 uint32 = (v2 ^ v2)
                _ = v62
            case 32767:
                var v63 uint = v5
                _ = v63
            default:
                v5 = v5
            }
            _ = v60
        } else {
            var v64 int16 = v52
            _ = v64
        }
        fmt.Println(v52, v1, v52)
    } else {
        v52 = (v52 % ((v52) | 1))
    }
    for fi65 := 0; fi65 < 9; fi65++ {
        var v66 float32 = -2.25
        v67 := []int{128, fi65, (fi65 % ((228) | 1))}
        v2 -= 2147483648
        _ = fi65
        _ = v66
        _ = v67
    }
    switch ((("foo bar" + "foo bar") + "a") + "hello") {
    case "hello":
        fmt.Println(v5, v4, v1)
    }
    delete(v3, "")
    v5 = 18446744073709551615
    if (!false) {
        var v68 float64 = 0.0
        _ = v68
    } else {
        v69 := v3["\n\t"]
        switch -443 {
        case -345:
            if (0 == -80) {
                v70 := GenS5{F0: ("a" + "hello"), F1: 32767, F2: false}
                if (305 > 42) {
                    var v71 int32 = 761
                    var v72 uint8 = 255
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        for fi73 := 0; fi73 < 17; fi73++ {
                            var v74 int8 = 127
                            fmt.Println(v2)
                            _ = fi73
                            _ = v74
                        }
                    case "a":
                        for fi75 := 0; fi75 < 4; fi75++ {
                            var v76 int8 = 17
                            _ = fi75
                            _ = v76
                        }
                    case "foo bar":
                        v3["\n\t"] = v69
                    }
                    _ = v71
                    _ = v72
                } else {
                    v52 = (^(v52))
                    fmt.Println(v69, v69)
                    v4 = (^(v4))
                }
                _ = v70
            } else {
                for fi77 := 0; fi77 < 17; fi77++ {
                    v78 := []int{fi77}
                    if (((-(v1)) != v1) && (true && (!true))) {
                        v79 := GenS6{F0: "foo bar", F1: 0, F2: v1}
                        if (false && false) {
                            fmt.Println(v5, v1)
                            var v80 float32 = float32(3.14159)
                            _ = v80
                        } else {
                            v81 := GenS5{F0: "", F1: 1, F2: (!true)}
                            for fi82 := 0; fi82 < 17; fi82++ {
                                if false {
                                    v83 := GenS4{F0: v69, F1: 127, F2: 4294967295}
                                    var v84 uint8 = v4
                                    _ = v83
                                    _ = v84
                                }
                                v1 = v1
                                v69 = v69
                                _ = fi82
                            }
                            _ = v81
                        }
                        v2 = uint32(v5)
                        _ = v79
                    }
                    _ = fi77
                    _ = v78
                }
                switch "\n\t" {
                case "":
                    if (v5 >= 42) {
                        switch -451 {
                        case 9223372036854775807:
                            v69 = (v69 | 0)
                        case -9223372036854775808:
                            v4 = (^(v4))
                        case -841:
                            v85 := GenS1{F0: int32(v1), F1: 1}
                            _ = v85
                        default:
                            if false {
                                v52 -= v52
                            }
                        }
                    } else {
                        v86 := []int{218, 1, 953}
                        _ = v86
                    }
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        var v87 uint16 = 32767
                        _ = v87
                    default:
                        switch 0 {
                        case 128:
                            for fi88 := 0; fi88 < 11; fi88++ {
                                if (1.5 == -2.25) {
                                    var v89 int32 = 42
                                    fmt.Println(v52, v5, fi88)
                                    _ = v89
                                } else {
                                    v90 := GenS7{F0: 32767, F1: 2}
                                    v2 = (v69 / ((612) | 1))
                                    _ = v90
                                }
                                _ = fi88
                            }
                        }
                    }
                case "hello":
                    var v91 float64 = 1e10
                    _ = v91
                default:
                    if false {
                        switch "a" {
                        case "hello":
                            for fi92 := 0; fi92 < 2; fi92++ {
                                for fi93 := 0; fi93 < 7; fi93++ {
                                    if (true && false) {
                                        v94 := []int{fi93, fi93, (fi92 | fi92), fi93}
                                        var v95 int8 = -1
                                        _ = v94
                                        _ = v95
                                    }
                                    if (v5 > (v5 - 954)) {
                                        var v96 int = (-(fi93))
                                        fmt.Println(v4, v52)
                                        _ = v96
                                    }
                                    _ = fi93
                                }
                                v5 = 2
                                _ = fi92
                            }
                        }
                        v97 := []int{32768, 128, -2147483648, 301}
                        for fi98 := 0; fi98 < 4; fi98++ {
                            switch "a" {
                            case "":
                                for fi99 := 0; fi99 < 3; fi99++ {
                                    var v100 uint32 = (v69 | v2)
                                    switch ("hello" + "hello") {
                                    case "":
                                        if (int(v1) != (fi98 * (fi99 >> ((612) & 63)))) {
                                            v1 -= (v1 % ((v1) | 1))
                                        }
                                    case "hello":
                                        if true {
                                            v101 := GenS8{F0: (^(v52)), F1: (v69 << ((v5) & 63)), F2: 0}
                                            v100 += v2
                                            v100 = 0
                                            _ = v101
                                        }
                                    }
                                    _ = fi99
                                    _ = v100
                                }
                            default:
                                var v102 int = 948
                                _ = v102
                            }
                            v52 += (v52 ^ 599)
                            v69 -= v2
                            _ = fi98
                        }
                        _ = v97
                    }
                }
                for fi103 := 0; fi103 < 12; fi103++ {
                    delete(v3, "hello")
                    for fi104 := 0; fi104 < 16; fi104++ {
                        switch (fi103 % (((fi104 * fi104)) | 1)) {
                        case -2147483648:
                            v105 := GenS4{F0: 2, F1: 1, F2: 1}
                            _ = v105
                        case 4294967295:
                            v106 := GenS9{F0: (v52 - v52)}
                            _ = v106
                        case 32767:
                            fmt.Println(v2)
                        }
                        if true {
                            v52 += v52
                            var v107 int16 = v52
                            var v108 uint = uint(v69)
                            var v109 int32 = 0
                            _ = v107
                            _ = v108
                            _ = v109
                        } else {
                            var v110 string = ("hello" + "foo bar")
                            _ = v110
                        }
                        _ = len(v3)
                        _ = fi104
                    }
                    _ = fi103
                }
            }
        case 0:
            v5 = 4294967295
        }
        switch "a" {
        case "foo bar":
            v2 -= v2
        case "hello":
            v5 += (v5 ^ (^(v5)))
        case "":
            v111 := GenS10{F0: 42, F1: -1, F2: (v2 + 1)}
            _ = v111
        }
        _ = v69
    }
    var v112 bool = true
    v3[(("" + "") + "a")] = 902
    if ((v1 + -9223372036854775808) <= (v1 | (v1 >> ((v5) & 63)))) {
        switch "hello" {
        case "a":
            var v113 string = ""
            _ = v113
        case "\n\t":
            if v112 {
                var v114 float32 = ((float32(-2.25) * 1e10) - 0.1)
                var v115 uint = v5
                _ = v114
                _ = v115
            }
        default:
            v4 = 175
        }
        v4 -= (v4 * 178)
        v116 := []int{}
        _ = v116
    } else {
        v117 := GenS11{F0: false}
        switch -502 {
        case -1:
            v3[("a" + ("a" + "a"))] = (v2 / ((42) | 1))
        }
        v4 += (^(v4))
        _ = v117
    }
    v118 := []int{-9223372036854775808, 2147483647, 9, 4294967295}
    v119 := GenS11{F0: (true && true)}
    v52 -= (v52 << (((^(v5))) & 63))
    var v120 uint8 = (v4 & v4)
    var v121 uint32 = 889
    var v122 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v118:", len(v118), v118)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v112
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
