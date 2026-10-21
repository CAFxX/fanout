package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int64; F2 int8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 bool; F1 uint16; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 bool; F2 uint16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 bool; F2 uint32 }
type GenS12 struct { F0 uint32; F1 int8 }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint64 = 4294967295
    fmt.Println(v3)
    for fi4 := 0; fi4 < 7; fi4++ {
        v3 -= v3
        _ = fi4
    }
    switch "hello" {
    case "foo bar":
        var v5 uint8 = 224
        _ = v5
    case "hello":
        switch -181 {
        case 32768:
            if false {
                fmt.Println(v3)
            } else {
                fmt.Println(v3)
                switch -868 {
                case -879:
                    switch 2147483647 {
                    case -675:
                        fmt.Println(v3)
                    case 9223372036854775807:
                        var v6 string = ""
                        _ = v6
                    }
                case 42:
                    var v7 uint32 = 2147483647
                    _ = v7
                case 128:
                    for fi8 := 0; fi8 < 1; fi8++ {
                        if (984 >= 0) {
                            v9 := GenS1{F0: (fi8 ^ (fi8 - -911)), F1: fi8, F2: 791}
                            switch -495 {
                            case 42:
                                if true {
                                    if (!(!true)) {
                                        var v10 uint8 = 87
                                        _ = v10
                                    } else {
                                        if true {
                                            v3 = (v3 >> ((18446744073709551615) & 63))
                                            v3 -= v3
                                            var v11 int32 = 2147483647
                                            _ = v11
                                        }
                                        for fi12 := 0; fi12 < 17; fi12++ {
                                            if true {
                                                v3 = v3
                                                v3 = 4294967295
                                                switch ("foo bar" + "hello") {
                                                case "hello":
                                                    v3 = v3
                                                case "\n\t":
                                                    v13 := []int{(fi8 | (fi8 * fi12)), (fi12 | (fi8 + (fi8 % (((fi12 << ((255) & 63))) | 1)))), (-(fi8))}
                                                    _ = v13
                                                default:
                                                    fmt.Println(fi12, fi12, fi12)
                                                }
                                            }
                                            if false {
                                                var v14 int8 = 127
                                                v15 := []int{(fi12 ^ fi12), (fi12 / (((^(fi8))) | 1)), fi8}
                                                var v16 int8 = (v14 % (((v14 * v14)) | 1))
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                            } else {
                                                if false {
                                                    if true {
                                                        switch fi8 {
                                                        case 0:
                                                            v17 := v2[799]
                                                            _ = v17
                                                        case 9223372036854775807:
                                                            v3 = v3
                                                        default:
                                                            if true {
                                                                v3 = 1
                                                            } else {
                                                                var v18 int8 = 1
                                                                v3 = (^(v3))
                                                                v18 = (v18 >> ((18446744073709551615) & 63))
                                                                v19 := []int{fi12, fi8, 0}
                                                                if (127 <= 255) {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if len(v19) >= 2 {
                                                                            v19 = v19[1:len(v19)]
                                                                        }
                                                                    case "\n\t":
                                                                        v3 += (v3 << ((4294967295) & 63))
                                                                    case "":
                                                                        v18 -= v18
                                                                    }
                                                                }
                                                                v20 := GenS1{F0: fi8, F1: (fi12 >> ((59) & 63)), F2: 674}
                                                                if (1 != 32767) {
                                                                    var v21 uint32 = 2
                                                                    v18 = v18
                                                                    delete(v2, 2147483648)
                                                                    switch (fi12 - (fi12 / (((fi8 + fi8)) | 1))) {
                                                                    case -621:
                                                                        v22 := GenS2{F0: false}
                                                                        _ = v22
                                                                    case 121:
                                                                        var v23 float64 = 1e10
                                                                        _ = v23
                                                                    case -498:
                                                                        switch (fi12 | fi8) {
                                                                        case 2147483647:
                                                                            var v24 string = (("a" + "\n\t") + ("foo bar" + ("\n\t" + (("" + "hello") + "hello"))))
                                                                            _ = v24
                                                                        case 4294967295:
                                                                            if (614 < 42) {
                                                                                v25 := GenS3{F0: (^(v3)), F1: -747, F2: v18}
                                                                                _ = v25
                                                                            } else {
                                                                                var v26 bool = (42 == 2)
                                                                                fmt.Println(v3, fi8, fi8)
                                                                                _ = v26
                                                                            }
                                                                        case 0:
                                                                            var v27 string = ("" + "hello")
                                                                            _ = v27
                                                                        }
                                                                        v28 := v2[v21]
                                                                        var v29 uint32 = 345
                                                                        var v30 uint32 = v21
                                                                        fmt.Println(v28, v30)
                                                                        _ = v28
                                                                        _ = v29
                                                                        _ = v30
                                                                    }
                                                                    _ = v21
                                                                }
                                                                _ = v18
                                                                _ = v19
                                                                _ = v20
                                                            }
                                                        }
                                                    }
                                                    if (true && ((true && true) || false)) {
                                                        v31 := v1[979]
                                                        var v32 bool = true
                                                        v33 := GenS2{F0: (!(!v32))}
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v3 += 0
                                                        v34 := GenS1{F0: fi12, F1: 9223372036854775807, F2: 194}
                                                        _ = v34
                                                    }
                                                    var v35 int32 = 1
                                                    _ = v35
                                                } else {
                                                    v3 -= (v3 & v3)
                                                    if ((((440 != 138) || false) && ((false && ((fi8 & (^(fi12))) > fi12)) || true)) && false) {
                                                        if false {
                                                            v3 -= (^(v3))
                                                        } else {
                                                            v36 := GenS4{F0: -128, F1: 200}
                                                            fmt.Println(fi12, fi8)
                                                            _ = v36
                                                        }
                                                        if true {
                                                            v37 := []int{fi12, (fi8 * 9223372036854775807)}
                                                            v38 := GenS2{F0: (2 == 127)}
                                                            _ = v37
                                                            _ = v38
                                                        } else {
                                                            v3 += v3
                                                            v39 := []int{(fi12 & 628), (fi12 / (((^(fi12))) | 1)), 996, fi8}
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v40 int32 = 2147483647
                                                                var v41 uint16 = 0
                                                                var v42 uint8 = uint8(v41)
                                                                v3 += v3
                                                                if false {
                                                                    fmt.Println(fi8)
                                                                    v43 := GenS3{F0: 468, F1: 1, F2: 100}
                                                                    _ = v43
                                                                }
                                                                fmt.Println(fi8, v42, v42)
                                                                _ = v40
                                                                _ = v41
                                                                _ = v42
                                                            case "hello":
                                                                v3 += v3
                                                            }
                                                            _ = v39
                                                        }
                                                        v44 := []int{(fi8 >> ((685) & 63))}
                                                        v45 := GenS5{F0: false, F1: 912, F2: 1000}
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        v46 := GenS1{F0: fi8, F1: fi8, F2: 2147483648}
                                                        switch (fi8 * (fi8 * fi12)) {
                                                        case 32768:
                                                            var v47 int8 = 100
                                                            _ = v47
                                                        default:
                                                            var v48 int = fi12
                                                            _ = v48
                                                        }
                                                        v3 += (^(v3))
                                                        _ = v46
                                                    }
                                                }
                                            }
                                            _ = fi12
                                        }
                                    }
                                } else {
                                    v3 -= 1
                                    switch "" {
                                    case "\n\t":
                                        v49 := []int{}
                                        _ = v49
                                    case "foo bar":
                                        var v50 bool = ((32767 == 1) && (!(1.5 == (float32(3.14159) + -2.25))))
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 6; fi51++ {
                                        var v52 uint = 1
                                        v3 += (^(v3))
                                        _ = fi51
                                        _ = v52
                                    }
                                }
                            default:
                                v53 := v1[4294967295]
                                _ = v53
                            }
                            _ = v9
                        }
                        _ = fi8
                    }
                default:
                    for fi54 := 0; fi54 < 11; fi54++ {
                        if false {
                            v55 := GenS3{F0: (v3 / ((2) | 1)), F1: -1, F2: 0}
                            _ = v55
                        } else {
                            var v56 bool = (false && true)
                            fmt.Println(v56)
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 8; fi57++ {
                            v3 = (^(v3))
                            v58 := GenS3{F0: 189, F1: 811, F2: 91}
                            _ = fi57
                            _ = v58
                        }
                        _ = fi54
                    }
                }
                v3 = 310
            }
        case 32767:
            var v59 uint16 = 63
            _ = v59
        case 128:
            v60 := []int{}
            _ = v60
        }
    }
    var v61 uint16 = 635
    switch 32767 {
    case 32768:
        var v62 int64 = 1
        _ = v62
    case -9223372036854775808:
        v63 := GenS4{F0: 1, F1: 42}
        _ = v63
    default:
        if (-2.25 == (1e10 + float64(float32((((-2.25 / 1.0) * (0.1 * 0.1)) - 1e10))))) {
            v61 += v61
            var v64 float64 = float64((float32((0.0 * 0.1)) + 1.5))
            v2[42] = 1
            _ = v64
        } else {
            fmt.Println(v61)
            v65 := GenS5{F0: (!true), F1: (^(v61)), F2: -32768}
            v3 = v3
            _ = v65
        }
    }
    switch 127 {
    case 32767:
        v3 -= 0
    case -2147483648:
        switch 2147483647 {
        case 9223372036854775807:
            if false {
                for fi66 := 0; fi66 < 7; fi66++ {
                    var v67 uint8 = 0
                    fmt.Println(v3, v61, v67)
                    if true {
                        var v68 uint = 18446744073709551615
                        var v69 int8 = 45
                        _ = v68
                        _ = v69
                    }
                    _ = fi66
                    _ = v67
                }
            } else {
                var v70 int = -9223372036854775808
                for fi71 := 0; fi71 < 18; fi71++ {
                    var v72 float32 = -2.25
                    if false {
                        var v73 string = "foo bar"
                        _ = v73
                    } else {
                        switch ("a" + ("\n\t" + ("" + "hello"))) {
                        case "hello":
                            for fi74 := 0; fi74 < 13; fi74++ {
                                _ = len(v2)
                                v75 := v2[uint32(v61)]
                                if true {
                                    switch ("foo bar" + "a") {
                                    case "":
                                        fmt.Println(v61)
                                    case "hello":
                                        v76 := GenS6{F0: -1}
                                        _ = v76
                                    }
                                    fmt.Println(v72)
                                }
                                _ = fi74
                                _ = v75
                            }
                        case "foo bar":
                            var v77 int32 = -59
                            _ = v77
                        case "\n\t":
                            var v78 int32 = -723
                            _ = v78
                        }
                    }
                    _ = fi71
                    _ = v72
                }
                v3 = (^(v3))
                _ = v70
            }
        case -356:
            v79 := GenS5{F0: true, F1: (^(v61)), F2: -552}
            _ = v79
        case 32768:
            v61 = 2
        default:
            var v80 uint = 513
            _ = v80
        }
    case 608:
        v81 := GenS7{F0: 128, F1: 2}
        _ = v81
    default:
        for fi82 := 0; fi82 < 16; fi82++ {
            for fi83 := 0; fi83 < 17; fi83++ {
                if true {
                    v3 -= (v3 * v3)
                    switch fi83 {
                    case -2147483648:
                        var v84 uint32 = 2147483648
                        _ = v84
                    case 0:
                        v61 -= (^(v61))
                    }
                } else {
                    if false {
                        if (v3 >= 971) {
                            v1[190] = (fi83 << ((18446744073709551615) & 63))
                            switch ("a" + "hello") {
                            case "hello":
                                var v85 int64 = 0
                                v86 := GenS4{F0: -97, F1: 899}
                                _ = v85
                                _ = v86
                            case "":
                                switch 127 {
                                case -1:
                                    v61 = v61
                                case 0:
                                    if (208 > 255) {
                                        if (("\n\t" + ("\n\t" + "hello")) > ((("hello" + "hello") + "\n\t") + ("\n\t" + ((("hello" + "hello") + "foo bar") + "a")))) {
                                            v3 += 673
                                        } else {
                                            v61 = v61
                                            var v87 string = "hello"
                                            v88 := GenS8{F0: 296}
                                            v89 := []int{fi83, fi82}
                                            if false {
                                                var v90 bool = true
                                                switch (fi82 * (fi82 - fi82)) {
                                                case -9223372036854775808:
                                                    v91 := GenS9{F0: 1, F1: v90, F2: (v61 % ((65535) | 1))}
                                                    _ = v91
                                                case 2147483647:
                                                    if (!true) {
                                                        switch "" {
                                                        case "hello":
                                                            v87 = "hello"
                                                        case "":
                                                            v90 = v90
                                                        default:
                                                            if v90 {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v92 uint64 = (v3 << ((0) & 63))
                                                                    _ = v92
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                            } else {
                                                                if v90 {
                                                                    var v93 string = v87
                                                                    v94 := GenS3{F0: 42, F1: 9223372036854775807, F2: -72}
                                                                    var v95 int32 = 460
                                                                    v61 += v61
                                                                    _ = v93
                                                                    _ = v94
                                                                    _ = v95
                                                                } else {
                                                                    switch v87 {
                                                                    case "foo bar":
                                                                        fmt.Println(v90, v61)
                                                                    }
                                                                    switch 4294967295 {
                                                                    case 127:
                                                                        var v96 float32 = 0.1
                                                                        _ = v96
                                                                    }
                                                                    fmt.Println(v3)
                                                                }
                                                                if (1 != -1) {
                                                                    var v97 float64 = 0.1
                                                                    _ = v97
                                                                } else {
                                                                    v3 -= v3
                                                                    if len(v89) >= 2 {
                                                                        v89 = v89[1:len(v89)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v98 float64 = ((3.14159 + ((1.5 + 1e10) + float64(-2.25))) - 0.0)
                                                        _ = v98
                                                    } else {
                                                        var v99 float32 = 0.1
                                                        _ = v99
                                                    }
                                                case 9223372036854775807:
                                                    v3 = v3
                                                    if (true && (((!(((1.5 / 1.0) - 3.14159) > 1e10)) || (v90 && v90)) && (0 >= 0))) {
                                                        v3 += v3
                                                    }
                                                default:
                                                    v100 := GenS2{F0: ((128 <= 50) && v90)}
                                                    v61 += v61
                                                    _ = v100
                                                }
                                                if true {
                                                    v101 := GenS2{F0: v90}
                                                    switch v87 {
                                                    case "\n\t":
                                                        var v102 uint32 = 2147483647
                                                        _ = v102
                                                    case "foo bar":
                                                        var v103 bool = true
                                                        _ = v103
                                                    case "hello":
                                                        v90 = v90
                                                    default:
                                                        v61 = (v61 + 2)
                                                    }
                                                    if v90 {
                                                        switch v87 {
                                                        case "foo bar":
                                                            var v104 string = "hello"
                                                            _ = v104
                                                        case "a":
                                                            if len(v89) > 0 {
                                                                v89[0] = fi82
                                                            }
                                                        case "hello":
                                                            v3 -= 0
                                                        default:
                                                            if ((v61 | 32767) < 430) {
                                                                var v105 string = ("\n\t" + v87)
                                                                if v90 {
                                                                    fmt.Println(v87)
                                                                } else {
                                                                    var v106 uint32 = 0
                                                                    v89 = append(v89, fi83)
                                                                    var v107 int = (fi83 % ((fi83) | 1))
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                                var v108 int32 = -1
                                                                _ = v105
                                                                _ = v108
                                                            } else {
                                                                if v90 {
                                                                    if false {
                                                                        fmt.Println(fi83, v90, v61)
                                                                        v87 = (v87 + (v87 + ""))
                                                                        v3 = v3
                                                                    }
                                                                }
                                                                if len(v89) > 0 {
                                                                    v89[0] = (fi82 << ((42) & 63))
                                                                }
                                                            }
                                                        }
                                                        v109 := GenS2{F0: v90}
                                                        _ = v109
                                                    }
                                                    fmt.Println(v61, fi82)
                                                    _ = v101
                                                } else {
                                                    var v110 int32 = 1
                                                    _ = v110
                                                }
                                                fmt.Println(fi83, v90)
                                                _ = v90
                                            }
                                            var v111 string = v87
                                            if (!(1.5 > -2.25)) {
                                                fmt.Println(v87, v3)
                                                v112 := GenS8{F0: 2}
                                                switch (v87 + v111) {
                                                case "\n\t":
                                                    switch fi82 {
                                                    case 447:
                                                        v113 := GenS10{F0: (v61 & 996)}
                                                        _ = v113
                                                    }
                                                default:
                                                    if (fi82 != fi82) {
                                                        v87 = v87
                                                        if ((!((0.1 + (float64(0.0) / 1.0)) != 0.1)) && true) {
                                                            fmt.Println(v87, fi83, v111)
                                                        }
                                                    } else {
                                                        var v114 uint64 = (v3 - 866)
                                                        v115 := GenS8{F0: v61}
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                }
                                                _ = v112
                                            } else {
                                                v87 = v87
                                                v87 = ("" + (v111 + (v111 + "")))
                                                if ((v61 << ((770) & 63)) > 1) {
                                                    if (!false) {
                                                        var v116 uint = 366
                                                        v61 = v61
                                                        fmt.Println(v116, v111, v111)
                                                        _ = v116
                                                    } else {
                                                        if len(v89) >= 2 {
                                                            v89 = v89[1:len(v89)]
                                                        }
                                                        var v117 string = "hello"
                                                        v111 = v87
                                                        _ = v117
                                                    }
                                                } else {
                                                    v118 := GenS3{F0: 4294967295, F1: 42, F2: 127}
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(v111)
                                                    case "foo bar":
                                                        fmt.Println(v61, v87, v3)
                                                    case "":
                                                        v87 = v111
                                                    default:
                                                        if true {
                                                            v61 += v61
                                                            switch fi82 {
                                                            case 724:
                                                                v89 = append(v89, fi82)
                                                            default:
                                                                v87 = "a"
                                                            }
                                                        } else {
                                                            switch v87 {
                                                            case "a":
                                                                v3 = uint64(v61)
                                                            default:
                                                                v87 = v111
                                                            }
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    if true {
                                                        v119 := GenS4{F0: 40, F1: 2147483647}
                                                        var v120 float64 = 1.5
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if (-1 == -9223372036854775808) {
                                                                v3 = (v3 & v3)
                                                                if false {
                                                                    v121 := GenS7{F0: fi83, F1: v61}
                                                                    var v122 float32 = 0.1
                                                                    var v123 float32 = (3.14159 / 0.1)
                                                                    _ = v121
                                                                    _ = v122
                                                                    _ = v123
                                                                } else {
                                                                    v120 = v120
                                                                    if (!(true && false)) {
                                                                        fmt.Println(v61, fi83, v3)
                                                                        if ((("" + v87) + v87) == ((v111 + v111) + (v87 + v87))) {
                                                                            if false {
                                                                                v89 = append(v89, 1)
                                                                                var v124 int32 = 1
                                                                                if (255 == 218) {
                                                                                    var v125 float64 = v120
                                                                                    var v126 string = v87
                                                                                    var v127 uint32 = 0
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                } else {
                                                                                    if ((0 <= -747) || (-2147483648 <= 2147483647)) {
                                                                                        var v128 int64 = int64(v61)
                                                                                        _ = v128
                                                                                    }
                                                                                    var v129 int = 4294967295
                                                                                    fmt.Println(v120)
                                                                                    _ = v129
                                                                                }
                                                                                _ = v124
                                                                            }
                                                                            var v130 float64 = 0.0
                                                                            v131 := GenS11{F0: 487, F1: (481 >= 80), F2: 1}
                                                                            _ = v130
                                                                            _ = v131
                                                                        } else {
                                                                            fmt.Println(v111)
                                                                        }
                                                                        var v132 uint = 1
                                                                        _ = v132
                                                                    } else {
                                                                        v133 := GenS12{F0: 2, F1: 1}
                                                                        v134 := GenS13{F0: v87, F1: 2147483648}
                                                                        _ = v133
                                                                        _ = v134
                                                                    }
                                                                    fmt.Println(v61, fi82)
                                                                }
                                                                fmt.Println(v3)
                                                            }
                                                        default:
                                                            if false {
                                                                if len(v89) > 0 {
                                                                    v89[0] = fi83
                                                                }
                                                            }
                                                        }
                                                        var v135 uint = 1
                                                        _ = v119
                                                        _ = v120
                                                        _ = v135
                                                    } else {
                                                        v136 := GenS9{F0: 2147483647, F1: (1 != 42), F2: v61}
                                                        _ = v136
                                                    }
                                                    var v137 uint64 = v3
                                                    _ = v118
                                                    _ = v137
                                                }
                                                switch (fi83 | fi83) {
                                                case 731:
                                                    v138 := GenS3{F0: (v3 % (((v3 | v3)) | 1)), F1: 245, F2: int8(fi83)}
                                                    if (((0.1 * 1e10) > -2.25) && (-2.25 < 1.5)) {
                                                        v139 := GenS6{F0: 42}
                                                        _ = v139
                                                    }
                                                    if true {
                                                        var v140 int32 = -1
                                                        _ = v140
                                                    } else {
                                                        v89 = append(v89, fi82)
                                                        v141 := v1[0]
                                                        _ = v141
                                                    }
                                                    _ = v138
                                                case 42:
                                                    v142 := GenS4{F0: 4, F1: 42}
                                                    _ = v142
                                                default:
                                                    var v143 int64 = -474
                                                    _ = v143
                                                }
                                                v2[42] = fi83
                                                switch (fi83 | (fi82 % ((2147483647) | 1))) {
                                                case 32768:
                                                    var v144 uint32 = 0
                                                    if false {
                                                        var v145 uint64 = 226
                                                        v146 := GenS6{F0: -953}
                                                        switch -29 {
                                                        case 371:
                                                            if len(v89) > 0 {
                                                                v89[0] = fi83
                                                            }
                                                            v147 := GenS5{F0: (!(255 <= (v61 + v61))), F1: 32767, F2: -1}
                                                            if len(v89) > 0 {
                                                                v89[0] = fi82
                                                            }
                                                            switch fi83 {
                                                            case -464:
                                                                v87 = (v87 + v87)
                                                            case 127:
                                                                delete(v2, v144)
                                                            default:
                                                                v148 := GenS14{F0: uint32(v61)}
                                                                _ = v148
                                                            }
                                                            var v149 float32 = 1e10
                                                            _ = v147
                                                            _ = v149
                                                        case 977:
                                                            fmt.Println(v87)
                                                        }
                                                        _ = v145
                                                        _ = v146
                                                    } else {
                                                        if (42 <= 65535) {
                                                            fmt.Println(v87)
                                                        } else {
                                                            v3 = 374
                                                        }
                                                        v89 = append(v89, fi82)
                                                        v61 = v61
                                                        fmt.Println(fi83, fi82)
                                                        var v150 uint = 4294967295
                                                        _ = v150
                                                    }
                                                    v151 := GenS8{F0: v61}
                                                    _ = v144
                                                    _ = v151
                                                default:
                                                    v87 = (v87 + (v111 + v111))
                                                }
                                            }
                                            var v152 int = fi83
                                            if len(v89) >= 2 {
                                                v89 = v89[1:len(v89)]
                                            }
                                            _ = v111
                                            _ = v152
                                            _ = v87
                                            _ = v88
                                            _ = v89
                                        }
                                        v153 := GenS8{F0: 829}
                                        _ = v153
                                    }
                                    fmt.Println(v61, fi83, v3)
                                case 127:
                                    switch fi83 {
                                    case 2147483647:
                                        v154 := GenS15{F0: v3, F1: -125}
                                        _ = v154
                                    }
                                }
                            default:
                                var v155 int32 = -1
                                _ = v155
                            }
                        } else {
                            var v156 float32 = (3.14159 - 0.0)
                            _ = v156
                        }
                        v61 = (v61 % ((42) | 1))
                    }
                    fmt.Println(v3)
                    v61 -= (v61 ^ 175)
                }
                v3 = (v3 | v3)
                v3 -= 416
                var v157 string = ("\n\t" + "a")
                _ = fi83
                _ = v157
            }
            var v158 float32 = 3.14159
            _ = fi82
            _ = v158
        }
    }
    v159 := GenS1{F0: -754, F1: 9223372036854775807, F2: 42}
    var v160 uint64 = uint64(v61)
    v160 = (v160 % ((1) | 1))
    v61 += (v61 - v61)
    v61 += 373
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v160:", v160)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v159
    _ = v160
    _ = v2
    _ = v3
    _ = v61
}
