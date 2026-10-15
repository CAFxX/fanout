package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
type GenS3 struct { F0 uint; F1 string; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int16; F2 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 bool; F1 string; F2 uint16 }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 uint32; F1 string; F2 int64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint32; F1 uint; F2 int32 }
type GenS15 struct { F0 int32; F1 int16 }
type GenS16 struct { F0 int64; F1 uint32; F2 string }
type GenS17 struct { F0 int8; F1 uint16; F2 int8 }
type GenS18 struct { F0 int32; F1 string }
func genFn0(p0 int64) uint8 {
    v6 := GenS1{F0: 127, F1: 1}
    var v7 int8 = -113
    _ = p0
    _ = v6
    _ = v7
    return 227
}
func genFn1() int16 {
    v8 := []int{-711, 127, 1}
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint = 577
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS2{F0: ("foo bar" + "hello"), F1: -99, F2: 32767}
            _ = fi13
            _ = v14
        }
        if false {
            switch fi12 {
            case 995:
                _ = len(v8)
            }
            v11 += v11
        }
        _ = fi12
    }
    if len(v8) > 0 {
        v8[0] = 819
    }
    _ = v11
    _ = v8
    return 1000
}

func main() {
    v1 := []string{(("hello" + "\n\t") + "foo bar")}
    v2 := make(map[uint32]int)
    v3 := []uint32{1, 4294967295}
    v4 := []string{"\n\t", "a"}
    v5 := make(map[string]bool)
    _ = len(v1)
    var v15 string = ""
    if (1 < -2147483648) {
        var v16 uint8 = 1
        _ = v16
    }
    switch (v15 + v15) {
    case "hello":
        for fi17 := 0; fi17 < 5; fi17++ {
            switch fi17 {
            case 0:
                v15 = (v15 + ("\n\t" + v15))
            case -9223372036854775808:
                if (false && (255 >= 195)) {
                    v18 := GenS2{F0: (v15 + "hello"), F1: 0, F2: 1}
                    if (!true) {
                        if true {
                            v19 := GenS3{F0: 255, F1: v15, F2: 0}
                            if len(v1) > 0 {
                                v1[0] = (v15 + v15)
                            }
                            v2[296] = (fi17 % (((fi17 + fi17)) | 1))
                            _ = v19
                        }
                        v15 = v15
                        v15 = (("" + v15) + v15)
                    }
                    if (!true) {
                        if true {
                            var v20 bool = (4294967295 != 2)
                            fmt.Println(fi17)
                            var v21 uint16 = 745
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = v18
                } else {
                    var v22 uint8 = 1
                    v15 = ("" + "\n\t")
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v15 = ((v15 + "hello") + (v15 + (("\n\t" + (v15 + ("foo bar" + v15))) + v15)))
                        v22 = (v22 + 2)
                        if false {
                            v15 = v15
                            var v24 uint32 = 2147483648
                            _ = v24
                        }
                        if (196 <= 255) {
                            var v25 int32 = 0
                            v25 = v25
                            _ = v25
                        }
                        _ = fi23
                    }
                    _ = v22
                }
            case 128:
                fmt.Println(v15)
            }
            _ = fi17
        }
    case "a":
        var v26 int8 = 0
        _ = v26
    case "foo bar":
        if (!(true && (-67 == 127))) {
            var v27 uint16 = 234
            v28 := GenS4{F0: (^(v27))}
            for fi29 := 0; fi29 < 13; fi29++ {
                fmt.Println(v15, v15, v27)
                switch "a" {
                case "foo bar":
                    v30 := GenS5{F0: 42, F1: -1}
                    _ = v30
                case "":
                    v31 := GenS4{F0: 42}
                    _ = v31
                case "\n\t":
                    if (!true) {
                        v27 -= (v27 << ((0) & 63))
                        if true {
                            v27 = v27
                            v27 = (v27 << ((1) & 63))
                            fmt.Println(fi29, v15, fi29)
                        }
                        _ = len(v1)
                    }
                default:
                    v3 = append(v3, 2)
                }
                delete(v5, v15)
                _ = fi29
            }
            _ = v27
            _ = v28
        } else {
            if ((18446744073709551615 != 0) || true) {
                fmt.Println(v15)
            }
            switch (v15 + "") {
            case "\n\t":
                v32 := GenS6{F0: -128}
                _ = v32
            }
        }
    }
    v33 := v2[0]
    for fi34 := 0; fi34 < 2; fi34++ {
        switch 32767 {
        case -843:
            v33 += v33
        case 32767:
            if (255 == 42) {
                if ((255 >= 0) || (!true)) {
                    v15 = v15
                    var v35 int32 = 1
                    _ = v35
                } else {
                    var v36 int8 = -128
                    v37 := GenS1{F0: 0, F1: 42}
                    for fi38 := 0; fi38 < 12; fi38++ {
                        v15 = ("a" + ("" + v15))
                        v3 = append(v3, 0)
                        var v39 uint = 42
                        var v40 float64 = 0.0
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = v36
                    _ = v37
                }
                fmt.Println(v15, fi34, fi34)
                if ((3.14159 == float32(0.1)) && false) {
                    var v41 uint = 1
                    var v42 bool = false
                    v33 = v33
                    _ = v41
                    _ = v42
                }
            } else {
                switch (fi34 * -213) {
                case 0:
                    var v43 float32 = (1.5 - (3.14159 - -2.25))
                    _ = v43
                case -746:
                    for fi44 := 0; fi44 < 7; fi44++ {
                        v33 -= (fi44 << ((999) & 63))
                        _ = fi44
                    }
                case 32767:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    v46 := v2[892]
                    switch (v15 + (v15 + v15)) {
                    case "foo bar":
                        v47 := GenS1{F0: 132, F1: 628}
                        _ = v47
                    default:
                        fmt.Println(fi45)
                    }
                    _ = len(v5)
                    _ = fi45
                    _ = v46
                }
            }
        case 2147483647:
            switch "foo bar" {
            case "":
                if (false || (1.5 == ((0.0 / ((1.5 - ((1e10 - 3.14159) - (1.5 / 1.5))) + 1.5)) / 1e10))) {
                    for fi48 := 0; fi48 < 12; fi48++ {
                        switch v33 {
                        case -23:
                            var v49 bool = ((true || ("hello" == "\n\t")) || false)
                            _ = v49
                        case -1:
                            var v50 bool = (((128 != 42) || false) || false)
                            _ = v50
                        default:
                            var v51 int16 = -582
                            _ = v51
                        }
                        fmt.Println(v33, v15)
                        if true {
                            switch (v33 - fi48) {
                            case 42:
                                v52 := GenS7{F0: 1, F1: 835, F2: genFn1()}
                                switch "" {
                                case "a":
                                    v15 = ("foo bar" + "")
                                case "\n\t":
                                    v33 -= (v33 + (fi34 ^ (fi34 & (v33 % (((fi48 - v33)) | 1)))))
                                default:
                                    v15 = ("\n\t" + v15)
                                }
                                v53 := GenS2{F0: ("a" + (v15 + (("hello" + "\n\t") + v15))), F1: 0, F2: 814}
                                _ = v52
                                _ = v53
                            case -9223372036854775808:
                                fmt.Println(fi48, v15)
                            }
                            v54 := v2[65535]
                            if true {
                                switch (((v15 + "a") + v15) + v15) {
                                case "hello":
                                    v5[v15] = (2147483647 >= 292)
                                    var v55 uint32 = 42
                                    var v56 float32 = 3.14159
                                    _ = v55
                                    _ = v56
                                case "":
                                    fmt.Println(v54, v33, v54)
                                case "a":
                                    if (34 != -1) {
                                        switch v15 {
                                        case "a":
                                            var v57 uint32 = 2147483648
                                            _ = v57
                                        case "":
                                            fmt.Println(v15, v33)
                                        case "foo bar":
                                            if false {
                                                v5[v15] = (!(1e10 <= 1e10))
                                                v58 := GenS8{F0: 493}
                                                _ = v58
                                            }
                                        }
                                        var v59 uint32 = 1
                                        switch 32768 {
                                        case 32768:
                                            var v60 float32 = 3.14159
                                            _ = v60
                                        case -2147483648:
                                            switch (fi48 / ((v54) | 1)) {
                                            case 0:
                                                var v61 uint8 = 255
                                                _ = v61
                                            default:
                                                switch (fi48 + fi34) {
                                                case 2147483647:
                                                    switch (v15 + v15) {
                                                    case "foo bar":
                                                        v3 = append(v3, 442)
                                                    case "a":
                                                        _ = len(v2)
                                                    case "\n\t":
                                                        if true {
                                                            v5[v15] = true
                                                            if (!(((false || (!true)) && false) && true)) {
                                                                switch -9223372036854775808 {
                                                                case -1:
                                                                    if false {
                                                                        if false {
                                                                            fmt.Println(v33, fi48, fi48)
                                                                            fmt.Println(fi48, v54, fi48)
                                                                            v62 := GenS2{F0: v15, F1: -4, F2: 1}
                                                                            _ = v62
                                                                        } else {
                                                                            v63 := GenS8{F0: 32767}
                                                                            if true {
                                                                                switch fi48 {
                                                                                case 32767:
                                                                                    v33 += (fi48 - (fi34 % ((fi34) | 1)))
                                                                                    var v64 string = v15
                                                                                    v64 = ("\n\t" + v15)
                                                                                    _ = v64
                                                                                default:
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        _ = len(v2)
                                                                                    default:
                                                                                        var v65 uint16 = 789
                                                                                        _ = v65
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v66 := GenS2{F0: v15, F1: -50, F2: -32768}
                                                                                _ = v66
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v54)
                                                                            } else {
                                                                                switch 128 {
                                                                                case 127:
                                                                                    fmt.Println(fi48, fi48, v15)
                                                                                case 32768:
                                                                                    v15 = v15
                                                                                }
                                                                                v67 := v2[v59]
                                                                                _ = v67
                                                                            }
                                                                            _ = v63
                                                                        }
                                                                        if (("foo bar" + (v15 + "")) < v15) {
                                                                            if (!((v15 + "foo bar") < v15)) {
                                                                                v33 -= 0
                                                                                var v68 int32 = -24
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v59
                                                                                }
                                                                                _ = v68
                                                                            } else {
                                                                                if false {
                                                                                    var v69 float64 = -2.25
                                                                                    var v70 int16 = genFn1()
                                                                                    v15 = "hello"
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                            if false {
                                                                                v71 := GenS8{F0: -32768}
                                                                                _ = v71
                                                                            } else {
                                                                                fmt.Println(v15, v33, v54)
                                                                                v72 := GenS9{F0: 683, F1: -651}
                                                                                var v73 uint8 = 2
                                                                                v74 := GenS10{F0: true, F1: v15, F2: 2}
                                                                                _ = v72
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                            if ((2 == 42) && false) {
                                                                                var v75 int8 = 0
                                                                                v76 := GenS6{F0: (v75 * v75)}
                                                                                switch (v54 & (v54 * 282)) {
                                                                                case -2147483648:
                                                                                    v77 := GenS3{F0: 4294967295, F1: ("\n\t" + v15), F2: v33}
                                                                                    _ = v77
                                                                                }
                                                                                _ = v75
                                                                                _ = v76
                                                                            } else {
                                                                                var v78 uint32 = (^(v59))
                                                                                _ = v78
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v79 := GenS11{F0: ((1.5 / 1.0) == (0.0 + ((((-2.25 - -2.25) / 3.14159) - (1.5 + float64(-2.25))) * 3.14159))), F1: 182}
                                                                        _ = v79
                                                                    }
                                                                case -64:
                                                                    var v80 uint64 = 1
                                                                    _ = v80
                                                                default:
                                                                    if ((!false) || true) {
                                                                        v59 -= 51
                                                                        v54 -= (v33 ^ fi48)
                                                                        v3 = append(v3, 2)
                                                                        var v81 float32 = 0.0
                                                                        switch (-(fi34)) {
                                                                        case 127:
                                                                            v33 += fi48
                                                                            if (65535 > 855) {
                                                                                fmt.Println(v33, v54, fi48)
                                                                                v15 = v15
                                                                                v33 += v54
                                                                            }
                                                                        }
                                                                        var v82 uint32 = (^(v59))
                                                                        _ = len(v5)
                                                                        if false {
                                                                            fmt.Println(v81)
                                                                            v59 -= v59
                                                                            fmt.Println(v54, fi48)
                                                                            var v83 uint = 2
                                                                            v84 := GenS12{F0: (^(v59)), F1: v15, F2: 380}
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            if false {
                                                                                v5[(v15 + "hello")] = ((v82 & 337) > (v59 & 2147483647))
                                                                                var v85 uint8 = 0
                                                                                if (v15 <= v15) {
                                                                                    fmt.Println(v81)
                                                                                    v86 := GenS13{F0: 42}
                                                                                    v87 := GenS14{F0: v82, F1: 18446744073709551615, F2: -2147483648}
                                                                                    var v88 string = ""
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                } else {
                                                                                    v89 := GenS15{F0: 0, F1: -32768}
                                                                                    var v90 int16 = 472
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                }
                                                                                v4 = append(v4, v15)
                                                                                v91 := v2[(v82 * 512)]
                                                                                _ = v85
                                                                                _ = v91
                                                                            }
                                                                            var v92 uint64 = 1
                                                                            v3 = append(v3, (^(v82)))
                                                                            _ = v92
                                                                        }
                                                                        v93 := GenS11{F0: (42 != -901), F1: 759}
                                                                        fmt.Println(v81, fi48)
                                                                        _ = v81
                                                                        _ = v82
                                                                        _ = v93
                                                                    }
                                                                }
                                                                _ = len(v1)
                                                            } else {
                                                                v33 -= (v33 & 42)
                                                            }
                                                            v94 := GenS15{F0: -2147483648, F1: 1000}
                                                            _ = v94
                                                        }
                                                    default:
                                                        var v95 bool = true
                                                        _ = v95
                                                    }
                                                case 4294967295:
                                                    v96 := GenS1{F0: genFn0(9223372036854775807), F1: genFn1()}
                                                    _ = v96
                                                default:
                                                    v97 := GenS16{F0: -1, F1: v59, F2: ""}
                                                    _ = v97
                                                }
                                            }
                                        case 128:
                                            v15 = v15
                                        }
                                        _ = v59
                                    } else {
                                        v54 = v33
                                        switch (fi48 | v33) {
                                        case -2147483648:
                                            if false {
                                                if (-2147483648 == -2147483648) {
                                                    switch v15 {
                                                    case "\n\t":
                                                        v33 = fi34
                                                    case "a":
                                                        fmt.Println(v15, fi34, fi48)
                                                    case "hello":
                                                        v98 := GenS10{F0: (2 <= 523), F1: (v15 + "\n\t"), F2: 42}
                                                        _ = v98
                                                    default:
                                                        var v99 uint64 = 0
                                                        _ = v99
                                                    }
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v100 uint8 = 42
                                                    _ = v100
                                                }
                                            } else {
                                                v101 := GenS8{F0: 32767}
                                                _ = v101
                                            }
                                        case -8:
                                            v102 := GenS17{F0: -126, F1: 1, F2: 45}
                                            _ = v102
                                        default:
                                            var v103 float64 = -2.25
                                            _ = v103
                                        }
                                        var v104 bool = true
                                        v33 -= (fi48 - fi48)
                                        _ = v104
                                    }
                                }
                                if true {
                                    var v105 int64 = 42
                                    switch (v15 + v15) {
                                    case "foo bar":
                                        v4 = append(v4, "foo bar")
                                    case "a":
                                        var v106 uint16 = 66
                                        _ = v106
                                    case "":
                                        if true {
                                            if false {
                                                var v107 int = 127
                                                _ = v107
                                            } else {
                                                fmt.Println(v15, v33)
                                            }
                                        }
                                    default:
                                        v108 := GenS18{F0: -151, F1: "hello"}
                                        _ = v108
                                    }
                                    _ = v105
                                } else {
                                    var v109 int = (fi48 | fi48)
                                    v33 += (v33 & (-(fi34)))
                                    _ = v109
                                }
                                var v110 int = 32767
                                _ = v110
                            }
                            _ = v54
                        }
                        _ = fi48
                    }
                    if (true || true) {
                        for fi111 := 0; fi111 < 3; fi111++ {
                            if len(v1) > 0 {
                                v1[0] = ((v15 + "a") + ((((v15 + "hello") + "\n\t") + ("\n\t" + ("foo bar" + (("hello" + v15) + v15)))) + v15))
                            }
                            _ = fi111
                        }
                    }
                } else {
                    for fi112 := 0; fi112 < 11; fi112++ {
                        var v113 int32 = 219
                        if true {
                            var v114 string = ("\n\t" + v15)
                            _ = v114
                        }
                        _ = fi112
                        _ = v113
                    }
                }
            }
        default:
            if false {
                v33 = 148
            }
        }
        v115 := v2[147]
        var v116 uint32 = 2147483647
        _ = fi34
        _ = v115
        _ = v116
    }
    v33 += v33
    for fi117 := 0; fi117 < 20; fi117++ {
        var v118 uint = 412
        _ = fi117
        _ = v118
    }
    v15 = v15
    var v119 int8 = 45
    v33 = -1
    v33 -= v33
    switch v33 {
    case -149:
        v15 = ("\n\t" + (v15 + "a"))
    case 614:
        if (v15 != v15) {
            _ = len(v3)
        }
    case -9:
        _ = len(v3)
    }
    for fi120 := 0; fi120 < 19; fi120++ {
        v119 = (v119 ^ 100)
        v119 -= v119
        _ = fi120
    }
    v119 += v119
    v119 = v119
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v119
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
