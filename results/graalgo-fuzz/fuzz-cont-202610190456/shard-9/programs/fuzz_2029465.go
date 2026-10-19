package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int32 }
type GenS4 struct { F0 int64; F1 int16; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 bool }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int32; F1 string; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 int32; F2 int16 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 uint; F2 uint8 }
type GenS13 struct { F0 int16; F1 string }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 uint32 }
type GenS18 struct { F0 uint; F1 bool; F2 int8 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 string }
type GenS21 struct { F0 string; F1 int16 }
type GenS22 struct { F0 uint; F1 int }
type GenS23 struct { F0 uint; F1 uint }
type GenS24 struct { F0 uint16 }
func genFn0() string {
    if false {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int64 = 0
            _ = fi7
            _ = v8
        }
    } else {
        var v9 float64 = float64(float32(0.1))
        _ = v9
    }
    var v10 int = 32768
    _ = v10
    return "hello"
}

func main() {
    var v1 int16 = 908
    var v2 uint16 = 2
    var v3 string = "\n\t"
    v4 := make(map[uint32]bool)
    var v5 int16 = (v1 | v1)
    var v6 bool = (1.5 >= ((0.1 + ((-2.25 - (0.0 / 3.14159)) - ((float64(1.5) / (1.5 / (3.14159 / -2.25))) + 0.0))) - 0.0))
    var v11 int64 = int64(v1)
    v2 -= (v2 & 792)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint32 = 4294967295
        fmt.Println(v5)
        switch (("\n\t" + v3) + "\n\t") {
        case "":
            var v14 int = (-(fi12))
            _ = v14
        case "a":
            if (1 > v13) {
                v2 = 479
                v2 -= (v2 / ((962) | 1))
                v15 := GenS1{F0: 2147483647}
                _ = v15
            }
        case "\n\t":
            v2 = (v2 >> ((2) & 63))
        }
        _ = fi12
        _ = v13
    }
    var v16 uint = 179
    for fi17 := 0; fi17 < 8; fi17++ {
        switch ("" + (v3 + ((("\n\t" + v3) + "foo bar") + v3))) {
        case "hello":
            fmt.Println(v3, v3)
        default:
            v18 := []int{(fi17 + 128), fi17, fi17}
            _ = v18
        }
        _ = fi17
    }
    v19 := []int{127, 94, -2147483648}
    if ((((0.1 + 0.1) * (1.5 + (0.0 + 1.5))) - 0.1) > (3.14159 * (((0.1 / 1.0) + float64((-2.25 - float32((float64(0.0) - (0.0 * float64(1.5))))))) / 1.5))) {
        switch v3 {
        case "hello":
            v20 := GenS1{F0: 635}
            _ = v20
        case "a":
            var v21 string = v3
            _ = v21
        }
        var v22 uint64 = 0
        v16 -= v16
        _ = v22
    }
    fmt.Println(v1, v6, v16)
    for fi23 := 0; fi23 < 6; fi23++ {
        if len(v19) > 0 {
            v19[0] = -35
        }
        switch v3 {
        case "foo bar":
            switch v3 {
            case "a":
                if v6 {
                    v24 := GenS1{F0: 4294967295}
                    _ = v24
                }
            default:
                v25 := GenS1{F0: 293}
                _ = v25
            }
        case "\n\t":
            var v26 uint64 = 84
            _ = v26
        default:
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 uint32 = 0
                _ = fi27
                _ = v28
            }
        }
        switch (v3 + v3) {
        case "hello":
            for fi29 := 0; fi29 < 9; fi29++ {
                if v6 {
                    if ((v11 == v11) || (!false)) {
                        v11 = -9223372036854775808
                        fmt.Println(v6, v6, v6)
                        if v6 {
                            var v30 uint32 = 133
                            switch "a" {
                            case "hello":
                                if (!(("foo bar" + (v3 + v3)) != v3)) {
                                    if len(v19) > 0 {
                                        v19[0] = fi23
                                    }
                                    if (!true) {
                                        _ = len(v19)
                                        v31 := GenS2{F0: (v2 / ((986) | 1)), F1: 127, F2: v1}
                                        v32 := GenS3{F0: 942, F1: 127, F2: -2147483648}
                                        _ = v31
                                        _ = v32
                                    } else {
                                        switch fi29 {
                                        case 1:
                                            v1 += v1
                                        case -2147483648:
                                            v30 -= v30
                                        }
                                        var v33 int8 = -89
                                        switch int(v30) {
                                        case 1:
                                            fmt.Println(v5, v33)
                                        default:
                                            v30 -= 1
                                        }
                                        v33 = (-(v33))
                                        if v6 {
                                            v5 += 0
                                            var v34 float32 = (-2.25 / (-2.25 / 3.14159))
                                            _ = v34
                                        }
                                        _ = v33
                                    }
                                    var v35 int32 = -968
                                    _ = v35
                                } else {
                                    switch 42 {
                                    case -273:
                                        v36 := GenS1{F0: v30}
                                        _ = v36
                                    default:
                                        switch v3 {
                                        case "":
                                            v37 := GenS2{F0: v2, F1: 42, F2: v1}
                                            _ = v37
                                        case "hello":
                                            v5 = v5
                                        case "a":
                                            v19 = append(v19, fi23)
                                            if (uint32(v16) == v30) {
                                                if (v3 != v3) {
                                                    var v38 int64 = (-(v11))
                                                    var v39 float32 = -2.25
                                                    _ = v38
                                                    _ = v39
                                                }
                                                if ((810 <= 18446744073709551615) || (v6 && (-1 < 907))) {
                                                    var v40 int = int(v2)
                                                    _ = v40
                                                } else {
                                                    switch fi23 {
                                                    case 931:
                                                        var v41 uint16 = v2
                                                        _ = v41
                                                    default:
                                                        v42 := GenS2{F0: v2, F1: 1, F2: v5}
                                                        _ = v42
                                                    }
                                                    v6 = (v6 && (float32(1.5) >= 1e10))
                                                    var v43 int16 = 1
                                                    _ = v43
                                                }
                                                delete(v4, (v30 / (((^(v30))) | 1)))
                                                v44 := GenS1{F0: (v30 / (((^(v30))) | 1))}
                                                v16 -= v16
                                                if (100 != -1) {
                                                    if ((((!v6) || v6) && (!v6)) && (!(v16 == (^(v16))))) {
                                                        var v45 int8 = -68
                                                        if v6 {
                                                            fmt.Println(v30)
                                                            _ = len(v19)
                                                        } else {
                                                            v46 := GenS2{F0: v2, F1: uint8(v2), F2: (^(v1))}
                                                            v19 = append(v19, (-(fi29)))
                                                            v47 := GenS2{F0: v2, F1: 128, F2: v1}
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        _ = v45
                                                    } else {
                                                        var v48 float32 = 3.14159
                                                        _ = v48
                                                    }
                                                }
                                                _ = len(v4)
                                                fmt.Println(v11, v5)
                                                _ = v44
                                            } else {
                                                if v6 {
                                                    _ = len(v19)
                                                    var v49 uint16 = (^(v2))
                                                    _ = v49
                                                }
                                                var v50 uint = 2
                                                _ = v50
                                            }
                                            var v51 float64 = 1.5
                                            if false {
                                                v30 -= (v30 | (v30 << ((255) & 63)))
                                            } else {
                                                v11 = (v11 * v11)
                                            }
                                            var v52 int64 = (v11 / ((v11) | 1))
                                            var v53 int64 = v11
                                            var v54 uint32 = 4294967295
                                            var v55 int16 = (v5 << (((^(v16))) & 63))
                                            switch (fi29 - 127) {
                                            case 0:
                                                var v56 int32 = 2147483647
                                                _ = v56
                                            }
                                            if (!true) {
                                                switch -377 {
                                                case -376:
                                                    var v57 string = v3
                                                    _ = v57
                                                case 32768:
                                                    v58 := GenS1{F0: 65535}
                                                    _ = v58
                                                default:
                                                    fmt.Println(v5)
                                                }
                                                _ = len(v4)
                                                v59 := GenS4{F0: -688, F1: (v55 >> (((v16 << (((^(v16))) & 63))) & 63)), F2: 42}
                                                _ = v59
                                            } else {
                                                if (!(((^(v2)) < v2) || v6)) {
                                                    v60 := GenS3{F0: -218, F1: uint8(v2), F2: 0}
                                                    switch ("a" + "\n\t") {
                                                    case "hello":
                                                        _ = len(v4)
                                                    case "\n\t":
                                                        var v61 int8 = 93
                                                        _ = v61
                                                    }
                                                    fmt.Println(fi29, v53, v53)
                                                    _ = v60
                                                } else {
                                                    if v6 {
                                                        v62 := GenS5{F0: (v30 - (v54 - (v30 << (((v16 & (v16 ^ v16))) & 63)))), F1: 65535, F2: ((!v6) && false)}
                                                        if v6 {
                                                            fmt.Println(v16, v5, v3)
                                                            _ = len(v19)
                                                            v63 := GenS6{F0: v6, F1: ("foo bar" + v3)}
                                                            switch 33 {
                                                            case 130:
                                                                if (972 != 4294967295) {
                                                                    if true {
                                                                        var v64 int64 = -9223372036854775808
                                                                        _ = v64
                                                                    } else {
                                                                        var v65 int16 = (v5 << (((v16 * 65535)) & 63))
                                                                        v66 := GenS3{F0: 1, F1: 0, F2: 2147483647}
                                                                        v55 += v1
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        v52 = (-(v52))
                                                                    case "":
                                                                        var v67 uint64 = 4294967295
                                                                        _ = v67
                                                                    case "foo bar":
                                                                        var v68 int64 = v52
                                                                        _ = v68
                                                                    default:
                                                                        v69 := GenS7{F0: -2147483648, F1: v3, F2: (v1 | 42)}
                                                                        _ = v69
                                                                    }
                                                                }
                                                            case 32768:
                                                                v70 := GenS7{F0: 74, F1: ("hello" + ("foo bar" + ("a" + v3))), F2: int16(v16)}
                                                                var v71 uint8 = 255
                                                                delete(v4, (v30 - 800))
                                                                _ = v70
                                                                _ = v71
                                                            case 808:
                                                                v52 = (v52 + (v52 << (((v16 + v16)) & 63)))
                                                            default:
                                                                _ = len(v19)
                                                                if false {
                                                                    v5 = v5
                                                                    var v72 bool = ((true || (((v51 + (v51 - (((-2.25 + (v51 / v51)) + (1.5 + v51)) - (float64(float32(v51)) / float64(1.5))))) + v51) < v51)) || (v3 < v3))
                                                                    v73 := v4[65535]
                                                                    if ((v53 >> (((^(v16))) & 63)) <= v53) {
                                                                        var v74 string = ""
                                                                        v51 += (v51 * ((v51 * -2.25) + (0.1 / (float64(0.1) * (float64(1e10) * (v51 / (float64(1.5) / v51)))))))
                                                                        fmt.Println(fi23, v72, v16)
                                                                        v55 -= 1000
                                                                        _ = v74
                                                                    } else {
                                                                        var v75 uint = v16
                                                                        _ = v75
                                                                    }
                                                                    _ = v72
                                                                    _ = v73
                                                                } else {
                                                                    var v76 uint16 = v2
                                                                    _ = v76
                                                                }
                                                            }
                                                            _ = v63
                                                        } else {
                                                            fmt.Println(fi29)
                                                            if (v51 <= (1e10 * v51)) {
                                                                var v77 int32 = 0
                                                                if v6 {
                                                                    v2 = v2
                                                                } else {
                                                                    switch fi29 {
                                                                    case -939:
                                                                        var v78 uint8 = 1
                                                                        _ = v78
                                                                    case -1:
                                                                        _ = len(v19)
                                                                    case 2147483647:
                                                                        switch (^(fi29)) {
                                                                        case -41:
                                                                            v53 -= v11
                                                                        default:
                                                                            switch fi23 {
                                                                            case -210:
                                                                                if (v6 || v6) {
                                                                                    var v79 bool = false
                                                                                    v52 += v53
                                                                                    v52 = (^(v11))
                                                                                    _ = v79
                                                                                }
                                                                            default:
                                                                                v80 := GenS8{F0: 2}
                                                                                _ = v80
                                                                            }
                                                                            v81 := GenS6{F0: (v16 < 4294967295), F1: v3}
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                    v82 := GenS9{F0: (v16 + (v16 & 1)), F1: (-(v77)), F2: v55}
                                                                    if (v6 && (42 < 100)) {
                                                                        if ((v1 >> ((v16) & 63)) <= int16(fi29)) {
                                                                            var v83 float64 = v51
                                                                            switch (fi29 % ((fi29) | 1)) {
                                                                            case 32768:
                                                                                fmt.Println(v54)
                                                                            default:
                                                                                fmt.Println(v1, v5)
                                                                            }
                                                                            if (uint8(v16) > uint8(v2)) {
                                                                                var v84 int32 = (v77 >> ((v16) & 63))
                                                                                _ = v84
                                                                            } else {
                                                                                switch (("foo bar" + (v3 + "")) + "") {
                                                                                case "foo bar":
                                                                                    fmt.Println(v83, v53)
                                                                                case "":
                                                                                    var v85 string = v3
                                                                                    v5 = -628
                                                                                    var v86 uint = (v16 | uint(v54))
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                var v87 uint = (v16 >> ((183) & 63))
                                                                                _ = v87
                                                                            }
                                                                            if (2 > 255) {
                                                                                switch (v3 + v3) {
                                                                                case "a":
                                                                                    v88 := v4[v30]
                                                                                    _ = v88
                                                                                case "":
                                                                                    v89 := GenS10{F0: 167, F1: v77}
                                                                                    _ = v89
                                                                                default:
                                                                                    var v90 float32 = -2.25
                                                                                    _ = v90
                                                                                }
                                                                            } else {
                                                                                var v91 int64 = v53
                                                                                _ = v91
                                                                            }
                                                                            v5 -= -1
                                                                            if true {
                                                                                if v6 {
                                                                                    var v92 int8 = 1
                                                                                    v93 := GenS11{F0: int(v92)}
                                                                                    v94 := GenS1{F0: v30}
                                                                                    _ = v92
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                                v1 -= (-(v5))
                                                                                v19 = append(v19, fi29)
                                                                                fmt.Println(v5)
                                                                            } else {
                                                                                v95 := GenS11{F0: fi23}
                                                                                _ = v95
                                                                            }
                                                                            var v96 int16 = (v5 + int16(v2))
                                                                            _ = v83
                                                                            _ = v96
                                                                        }
                                                                        v30 = (v30 / ((65535) | 1))
                                                                        var v97 float32 = (float32(float64(float32(v51))) - 0.0)
                                                                        _ = v97
                                                                    }
                                                                    _ = v82
                                                                }
                                                                fmt.Println(v2, v51, v51)
                                                                var v98 uint64 = 42
                                                                _ = v77
                                                                _ = v98
                                                            } else {
                                                                switch ("a" + (v3 + ("\n\t" + v3))) {
                                                                case "foo bar":
                                                                    v52 += int64(v16)
                                                                case "a":
                                                                    v1 = (^(v5))
                                                                default:
                                                                    var v99 uint32 = (v54 >> ((507) & 63))
                                                                    if v6 {
                                                                        fmt.Println(v16, v30)
                                                                        if true {
                                                                            v100 := v4[2147483647]
                                                                            v101 := GenS10{F0: 1, F1: 1}
                                                                            _ = v100
                                                                            _ = v101
                                                                        }
                                                                        if v6 {
                                                                            var v102 int = 1
                                                                            var v103 string = "\n\t"
                                                                            delete(v4, 1)
                                                                            _ = v102
                                                                            _ = v103
                                                                        } else {
                                                                            if (!((v99 | 790) < 292)) {
                                                                                v3 = (v3 + v3)
                                                                            } else {
                                                                                if len(v19) > 0 {
                                                                                    v19[0] = -1
                                                                                }
                                                                                v52 -= (-(v53))
                                                                            }
                                                                            v104 := GenS12{F0: fi29, F1: 431, F2: 42}
                                                                            v105 := GenS13{F0: (^(v55)), F1: "\n\t"}
                                                                            _ = v104
                                                                            _ = v105
                                                                        }
                                                                    } else {
                                                                        if v6 {
                                                                            v30 = v54
                                                                            v106 := GenS14{F0: ("" + v3)}
                                                                            if v6 {
                                                                                v19 = append(v19, (fi29 % ((fi23) | 1)))
                                                                                v107 := GenS15{F0: 1}
                                                                                if len(v19) > 0 {
                                                                                    v19[0] = (-(fi23))
                                                                                }
                                                                                switch -991 {
                                                                                case 171:
                                                                                    v52 += (v11 / ((int64(v2)) | 1))
                                                                                case -184:
                                                                                    switch (("hello" + v3) + v3) {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v3)
                                                                                    case "a":
                                                                                        var v108 string = ((v3 + v3) + "foo bar")
                                                                                        _ = v108
                                                                                    default:
                                                                                        if len(v19) > 0 {
                                                                                            v19[0] = fi23
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v109 float32 = -2.25
                                                                                _ = v107
                                                                                _ = v109
                                                                            }
                                                                            _ = v106
                                                                        }
                                                                        v110 := GenS7{F0: int32(v2), F1: "", F2: (v5 | (v1 ^ (v1 << ((v16) & 63))))}
                                                                        switch (fi23 << (((v16 - (v16 & (v16 % (((^(v16))) | 1))))) & 63)) {
                                                                        case 127:
                                                                            v16 += v16
                                                                        }
                                                                        var v111 uint64 = uint64(v2)
                                                                        _ = v110
                                                                        _ = v111
                                                                    }
                                                                    v112 := v4[v30]
                                                                    _ = v112
                                                                    _ = v99
                                                                }
                                                                v113 := GenS16{F0: fi23}
                                                                _ = v113
                                                            }
                                                        }
                                                        _ = v62
                                                    } else {
                                                        fmt.Println(v11, v5, v51)
                                                        if (int(v16) < (fi29 | -44)) {
                                                            var v114 uint32 = uint32(v16)
                                                            v55 = -1
                                                            v115 := GenS17{F0: 2, F1: uint32(v16)}
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                        v51 -= v51
                                                    }
                                                    v30 -= 4294967295
                                                }
                                                v16 += v16
                                                switch v3 {
                                                case "":
                                                    fmt.Println(v5)
                                                case "a":
                                                    delete(v4, v30)
                                                case "foo bar":
                                                    switch v3 {
                                                    case "a":
                                                        var v116 int32 = 855
                                                        _ = v116
                                                    case "foo bar":
                                                        var v117 uint8 = 10
                                                        _ = v117
                                                    case "":
                                                        fmt.Println(v52)
                                                    default:
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                    }
                                                    v118 := v4[65535]
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if v6 {
                                                            fmt.Println(fi23, v55)
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                _ = len(v4)
                                                            default:
                                                                var v119 uint64 = 42
                                                                _ = v119
                                                            }
                                                            v3 = v3
                                                            switch v3 {
                                                            case "\n\t":
                                                                if (v6 && (fi29 > (fi23 | (fi23 * -416)))) {
                                                                    v120 := GenS18{F0: (v16 * v16), F1: ((!true) || v118), F2: -97}
                                                                    var v121 uint32 = v30
                                                                    var v122 uint16 = v2
                                                                    _ = v120
                                                                    _ = v121
                                                                    _ = v122
                                                                } else {
                                                                    v5 = (v1 >> ((v16) & 63))
                                                                }
                                                            case "a":
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v123 := GenS19{F0: (fi23 / ((int(v16)) | 1))}
                                                                        _ = v123
                                                                    case "hello":
                                                                        v51 += float64(1e10)
                                                                    case "":
                                                                        fmt.Println(v11, v5, v16)
                                                                        var v124 int32 = -2147483648
                                                                        var v125 int8 = 42
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                    if (v6 && (v3 != "a")) {
                                                                        v126 := GenS3{F0: -869, F1: 0, F2: 2147483647}
                                                                        _ = v126
                                                                    } else {
                                                                        var v127 string = v3
                                                                        v127 = v3
                                                                        var v128 uint64 = 0
                                                                        _ = v127
                                                                        _ = v128
                                                                    }
                                                                } else {
                                                                    var v129 int32 = -615
                                                                    switch (fi29 << (((v16 - (v16 | (v16 + v16)))) & 63)) {
                                                                    case -9223372036854775808:
                                                                        if (0 < 2) {
                                                                            fmt.Println(v2)
                                                                        }
                                                                    }
                                                                    _ = v129
                                                                }
                                                            default:
                                                                v130 := v4[v30]
                                                                _ = v130
                                                            }
                                                            _ = len(v19)
                                                        } else {
                                                            v131 := GenS1{F0: (v54 | 265)}
                                                            var v132 int8 = -79
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                    case "\n\t":
                                                        switch (v3 + v3) {
                                                        case "a":
                                                            if ((v16 & 642) == (v16 & (v16 % ((v16) | 1)))) {
                                                                v133 := GenS20{F0: (v3 + (v3 + ("\n\t" + ((v3 + (v3 + (v3 + (("" + "a") + ((("a" + "") + "a") + ("hello" + "hello")))))) + v3))))}
                                                                _ = v133
                                                            } else {
                                                                v30 -= v30
                                                                v6 = true
                                                            }
                                                        case "hello":
                                                            v134 := GenS8{F0: uint8(v2)}
                                                            _ = v134
                                                        case "foo bar":
                                                            v135 := GenS21{F0: "\n\t", F1: 1000}
                                                            _ = v135
                                                        }
                                                    case "hello":
                                                        v6 = v6
                                                    }
                                                    _ = v118
                                                }
                                            }
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                }
                            case "a":
                                v11 = v11
                            default:
                                v1 += 183
                            }
                            _ = v30
                        }
                        v136 := GenS7{F0: 347, F1: "foo bar", F2: v5}
                        _ = v136
                    }
                } else {
                    var v137 float64 = float64(0.0)
                    v138 := v4[473]
                    _ = v137
                    _ = v138
                }
                _ = fi29
            }
        case "a":
            if ((0.0 - ((1e10 + float32(-2.25)) + 1.5)) == (((1.5 + 1e10) + 0.1) - -2.25)) {
                switch fi23 {
                case 1:
                    if len(v19) > 0 {
                        v19[0] = 127
                    }
                case -416:
                    for fi139 := 0; fi139 < 7; fi139++ {
                        if (!false) {
                            v16 = 159
                            var v140 int32 = 2147483647
                            v3 = v3
                            _ = v140
                        }
                        _ = fi139
                    }
                }
                var v141 int64 = v11
                switch ("foo bar" + (((v3 + "\n\t") + v3) + ("foo bar" + ((v3 + (v3 + v3)) + v3)))) {
                case "\n\t":
                    v3 = ((("\n\t" + v3) + v3) + "foo bar")
                case "a":
                    v2 += v2
                }
                _ = v141
            } else {
                var v142 int64 = v11
                var v143 bool = false
                for fi144 := 0; fi144 < 14; fi144++ {
                    if v6 {
                        var v145 int16 = int16(fi144)
                        if ((fi23 != (fi23 | -2147483648)) || v6) {
                            if ((-2.25 * 0.0) > 0.1) {
                                switch v3 {
                                case "foo bar":
                                    var v146 uint32 = 65535
                                    _ = v146
                                case "\n\t":
                                    fmt.Println(v3, v11)
                                    switch fi144 {
                                    case -861:
                                        switch v3 {
                                        case "a":
                                            if (!(v3 == v3)) {
                                                var v147 int8 = 42
                                                fmt.Println(v6, fi23)
                                                if len(v19) > 0 {
                                                    v19[0] = fi23
                                                }
                                                _ = v147
                                            } else {
                                                var v148 bool = v143
                                                switch (v3 + "\n\t") {
                                                case "foo bar":
                                                    v148 = v148
                                                    if (v5 == (v5 - (v1 | v1))) {
                                                        switch -9223372036854775808 {
                                                        case -1:
                                                            if ((v143 || (!(v16 < v16))) || v143) {
                                                                var v149 string = ("" + ((v3 + v3) + "foo bar"))
                                                                var v150 uint64 = 574
                                                                v151 := GenS1{F0: 2}
                                                                v152 := GenS22{F0: (v16 & v16), F1: (-(fi23))}
                                                                _ = v149
                                                                _ = v150
                                                                _ = v151
                                                                _ = v152
                                                            } else {
                                                                if (((v142 / ((-9223372036854775808) | 1)) <= (v11 - (v142 * 1))) || ((fi144 ^ -191) <= -568)) {
                                                                    v11 -= (v11 * -634)
                                                                    fmt.Println(v145)
                                                                } else {
                                                                    if v6 {
                                                                        var v153 string = v3
                                                                        fmt.Println(v145, v153, v145)
                                                                        var v154 int = -171
                                                                        v4[2147483647] = v6
                                                                        _ = v153
                                                                        _ = v154
                                                                    }
                                                                }
                                                                var v155 bool = (((0.1 * (1e10 + 0.0)) + 0.1) >= 3.14159)
                                                                v156 := GenS8{F0: 2}
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                        default:
                                                            if ((float32(0.0) / 1.0) != (3.14159 * (((0.0 * -2.25) * 0.1) - 0.0))) {
                                                                v16 += (v16 ^ v16)
                                                            } else {
                                                                if false {
                                                                    var v157 int16 = (v1 * (^(v1)))
                                                                    var v158 float32 = 0.0
                                                                    _ = v157
                                                                    _ = v158
                                                                }
                                                                switch fi23 {
                                                                case 128:
                                                                    v148 = ((v142 >> ((0) & 63)) != (-(v11)))
                                                                case -588:
                                                                    fmt.Println(v3, fi23)
                                                                default:
                                                                    if v143 {
                                                                        var v159 int8 = 0
                                                                        v16 += v16
                                                                        v19 = append(v19, 4294967295)
                                                                        v1 = (v145 * int16(v159))
                                                                        v159 -= (v159 / (((v159 >> (((^(v16))) & 63))) | 1))
                                                                        fmt.Println(v159, v6, v1)
                                                                        _ = v159
                                                                    }
                                                                }
                                                                v160 := GenS15{F0: 127}
                                                                _ = v160
                                                            }
                                                        }
                                                    } else {
                                                        var v161 uint8 = 255
                                                        _ = v161
                                                    }
                                                    v143 = ((^(fi144)) < (fi144 << (((v16 | 2)) & 63)))
                                                case "":
                                                    fmt.Println(v142, fi144)
                                                }
                                                fmt.Println(v11, v11, v3)
                                                var v162 uint = 4294967295
                                                var v163 float64 = 3.14159
                                                _ = v148
                                                _ = v162
                                                _ = v163
                                            }
                                        case "foo bar":
                                            if v6 {
                                                var v164 string = v3
                                                v19 = append(v19, (fi23 << (((v16 << ((v16) & 63))) & 63)))
                                                var v165 bool = ((((1e10 / (0.1 - 1.5)) / 1.0) - 0.0) == 1.5)
                                                switch (v164 + "") {
                                                case "":
                                                    v166 := v4[0]
                                                    _ = v166
                                                default:
                                                    v167 := v4[241]
                                                    _ = v167
                                                }
                                                if v165 {
                                                    if (v1 == v145) {
                                                        if (-280 > -2147483648) {
                                                            var v168 float64 = float64(((((((-2.25 * 1.5) + -2.25) / 1.0) + 3.14159) / 1.0) * float32((3.14159 / float64(-2.25)))))
                                                            switch 0 {
                                                            case 2147483647:
                                                                v169 := GenS7{F0: 42, F1: "foo bar", F2: v1}
                                                                _ = v169
                                                            }
                                                            _ = v168
                                                        }
                                                        var v170 uint32 = 2
                                                        _ = v170
                                                    }
                                                    var v171 uint32 = 0
                                                    v3 = v3
                                                    v5 = v5
                                                    if v6 {
                                                        _ = len(v19)
                                                    }
                                                    switch -9223372036854775808 {
                                                    case -1:
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if len(v19) > 0 {
                                                                v19[0] = (^(fi144))
                                                            }
                                                        case "\n\t":
                                                            var v172 float32 = 1e10
                                                            _ = v172
                                                        }
                                                    case 127:
                                                        v16 += 18446744073709551615
                                                        if (v171 != 1) {
                                                            switch "\n\t" {
                                                            case "":
                                                                var v173 uint16 = (v2 - (v2 ^ (^(v2))))
                                                                _ = v173
                                                            }
                                                        } else {
                                                            v2 += 928
                                                            if v6 {
                                                                v174 := GenS8{F0: 127}
                                                                v145 -= (v145 << ((v16) & 63))
                                                                var v175 string = (v164 + (v3 + v164))
                                                                _ = v174
                                                                _ = v175
                                                            }
                                                        }
                                                        fmt.Println(v145, v2, v3)
                                                    }
                                                    v145 = v1
                                                    v145 = v145
                                                    _ = v171
                                                } else {
                                                    v176 := GenS4{F0: (-(v142)), F1: 1, F2: v11}
                                                    var v177 uint16 = (v2 * (v2 % (((v2 >> (((v16 & v16)) & 63))) | 1)))
                                                    _ = v176
                                                    _ = v177
                                                }
                                                if len(v19) > 0 {
                                                    v19[0] = 0
                                                }
                                                _ = v164
                                                _ = v165
                                            }
                                        default:
                                            v178 := GenS23{F0: v16, F1: v16}
                                            v3 = ((v3 + v3) + "")
                                            _ = v178
                                        }
                                    case -58:
                                        v1 -= (v1 >> ((42) & 63))
                                    case -1:
                                        _ = len(v4)
                                    default:
                                        if (-128 >= 31) {
                                            v1 -= (v5 + v145)
                                        }
                                        if (v143 || v143) {
                                            v6 = v143
                                            v142 = v142
                                            v4[603] = (!v143)
                                        } else {
                                            var v179 int8 = -128
                                            fmt.Println(v2, v142, fi144)
                                            _ = v179
                                        }
                                        _ = len(v19)
                                    }
                                }
                            } else {
                                switch ("foo bar" + "") {
                                case "hello":
                                    _ = len(v4)
                                default:
                                    if false {
                                        var v180 int16 = v1
                                        var v181 int8 = -1
                                        var v182 int = (fi144 + (fi144 - int(v181)))
                                        _ = v180
                                        _ = v181
                                        _ = v182
                                    } else {
                                        var v183 int16 = (^(v1))
                                        var v184 int16 = (v1 + 32767)
                                        if v6 {
                                            if (3.14159 >= 3.14159) {
                                                v16 = v16
                                                v143 = v143
                                            } else {
                                                v4[323] = v6
                                                if v143 {
                                                    v2 += v2
                                                    var v185 int16 = (v184 >> ((v16) & 63))
                                                    _ = v185
                                                } else {
                                                    v186 := GenS24{F0: (v2 >> ((570) & 63))}
                                                    _ = len(v4)
                                                    _ = v186
                                                }
                                            }
                                            v187 := GenS15{F0: 100}
                                            v16 = (v16 + (v16 >> (((v16 | 1)) & 63)))
                                            _ = v187
                                        } else {
                                            v145 = v145
                                            v1 = 42
                                            fmt.Println(v16, v184)
                                        }
                                        _ = v183
                                        _ = v184
                                    }
                                }
                                v188 := GenS10{F0: 127, F1: -682}
                                if v6 {
                                    v6 = true
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    switch -9223372036854775808 {
                                    case -9223372036854775808:
                                        if (4294967295 != 42) {
                                            v5 = v5
                                            v11 = v11
                                            var v189 int16 = v1
                                            _ = v189
                                        } else {
                                            v190 := GenS23{F0: 18446744073709551615, F1: (v16 - (^(v16)))}
                                            delete(v4, 1)
                                            if (0.1 == (3.14159 + 3.14159)) {
                                                if true {
                                                    v191 := GenS24{F0: v2}
                                                    _ = v191
                                                }
                                                v11 += -9223372036854775808
                                                var v192 uint32 = 791
                                                switch "" {
                                                case "\n\t":
                                                    var v193 bool = (1 >= 0)
                                                    _ = v193
                                                case "":
                                                    v6 = (!(v143 && v143))
                                                }
                                                v19 = append(v19, (fi23 * 0))
                                                var v194 int16 = v145
                                                _ = v192
                                                _ = v194
                                            } else {
                                                fmt.Println(fi23, fi23, v6)
                                                if (13 != 0) {
                                                    v145 = (v5 / (((v5 << (((v16 + v16)) & 63))) | 1))
                                                    if (100 > 123) {
                                                        var v195 uint32 = 42
                                                        var v196 uint8 = 1
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                        _ = len(v4)
                                                        _ = v195
                                                        _ = v196
                                                    }
                                                } else {
                                                    v6 = v6
                                                    fmt.Println(v3, v2)
                                                }
                                            }
                                            _ = v190
                                        }
                                    }
                                    var v197 uint16 = v2
                                    _ = v197
                                }
                                _ = v188
                            }
                            fmt.Println(v6)
                            var v198 uint64 = 2
                            _ = v198
                        } else {
                            var v199 int8 = 100
                            v143 = (65535 != v16)
                            var v200 int64 = 558
                            _ = v199
                            _ = v200
                        }
                        _ = v145
                    }
                    _ = fi144
                }
                _ = v142
                _ = v143
            }
        default:
            v6 = v6
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
