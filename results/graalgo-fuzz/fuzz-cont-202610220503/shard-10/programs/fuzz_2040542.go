package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 uint64; F1 string; F2 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int64; F1 int; F2 int16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint64; F1 int64; F2 int8 }
type GenS17 struct { F0 int8; F1 int16; F2 int32 }
type GenS18 struct { F0 uint16; F1 int }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint) uint {
    fmt.Println(p0)
    v4 := []int{}
    var v5 uint8 = 128
    p1 += (p1 | 65535)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 uint16, p1 string) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p1 = ("" + (p1 + p1))
        for fi7 := 0; fi7 < 2; fi7++ {
            p1 = p1
            var v8 bool = (int64(p0) == 165)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    switch p1 {
    case "":
        switch -1 {
        case 4294967295:
            var v9 uint = 413
            _ = v9
        default:
            v10 := GenS1{F0: 127, F1: 2147483647}
            _ = v10
        }
    case "foo bar":
        v11 := []int{4294967295, 4294967295, -2147483648, 9223372036854775807}
        _ = v11
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{2}
    v2 := []uint32{282}
    var v3 bool = false
    var v12 int64 = 0
    if false {
        v12 += (v12 >> ((42) & 63))
        fmt.Println(v3, v3)
        var v13 string = ""
        _ = v13
    } else {
        v14 := GenS1{F0: -1, F1: 1}
        var v15 uint64 = 18446744073709551615
        var v16 uint16 = 32767
        _ = v14
        _ = v15
        _ = v16
    }
    v12 -= (^(v12))
    v3 = ((((((1e10 + 1.5) != (0.1 * 3.14159)) && (((1e10 + float64(-2.25)) == ((3.14159 - 3.14159) - 1e10)) && v3)) && (42 <= 248)) || (!v3)) || v3)
    v3 = (42 != 131)
    switch -1 {
    case -933:
        v12 -= -1
    case 9223372036854775807:
        v1 = append(v1, 42)
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int64 = v12
        _ = fi17
        _ = v18
    }
    if v3 {
        v19 := GenS2{F0: 0, F1: -1}
        var v20 int64 = (v12 << ((18446744073709551615) & 63))
        _ = v19
        _ = v20
    } else {
        v3 = true
        var v21 float32 = float32(1e10)
        if (false && v3) {
            for fi22 := 0; fi22 < 4; fi22++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v23 uint64 = 974
                if ((fi22 & 0) != (fi22 - 127)) {
                    fmt.Println(v21)
                    v24 := GenS2{F0: (v12 / ((v12) | 1)), F1: v12}
                    v23 -= (v23 - (v23 + v23))
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            v25 := GenS2{F0: -520, F1: (v12 ^ 0)}
            fmt.Println(v3)
            _ = v25
        } else {
            if (0 >= 65535) {
                switch -9223372036854775808 {
                case 127:
                    v26 := GenS3{F0: -1, F1: 42}
                    _ = v26
                case -9223372036854775808:
                    v21 = float32(-2.25)
                case 100:
                    v27 := GenS4{F0: -408, F1: 2}
                    _ = v27
                }
                v28 := GenS2{F0: v12, F1: v12}
                if v3 {
                    if ((42 < -32768) || (float64((float32(1.5) - float32(1.5))) > -2.25)) {
                        var v29 int64 = v12
                        for fi30 := 0; fi30 < 1; fi30++ {
                            switch fi30 {
                            case 9223372036854775807:
                                switch fi30 {
                                case -974:
                                    for fi31 := 0; fi31 < 18; fi31++ {
                                        var v32 uint8 = 2
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v33 bool = v3
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                case -757:
                                    var v34 int64 = (v12 % (((v12 - (^(v29)))) | 1))
                                    _ = v34
                                case -9223372036854775808:
                                    var v35 int = fi30
                                    _ = v35
                                }
                            }
                            _ = fi30
                        }
                        v29 -= (v12 % ((0) | 1))
                        _ = v29
                    } else {
                        fmt.Println(v12, v12)
                    }
                    if true {
                        var v36 uint8 = 42
                        v37 := GenS5{F0: 18446744073709551615, F1: -32768, F2: 2147483647}
                        switch 0 {
                        case 2147483647:
                            var v38 uint = genFn0((1 > 122), 2)
                            _ = v38
                        }
                        _ = v36
                        _ = v37
                    } else {
                        v12 += v12
                        _ = len(v2)
                    }
                } else {
                    for fi39 := 0; fi39 < 10; fi39++ {
                        var v40 uint = genFn0(true, 2)
                        if v3 {
                            var v41 int32 = -553
                            if ((v21 - (v21 / 1.0)) <= v21) {
                                if (32767 <= 908) {
                                    if v3 {
                                        v2 = append(v2, 2147483647)
                                        v42 := GenS6{F0: 118}
                                        _ = v42
                                    } else {
                                        if ((!v3) || (468 >= 2)) {
                                            v12 += v12
                                            v3 = (-790 < 42)
                                        } else {
                                            v3 = v3
                                        }
                                        if (v3 || (genFn0((v40 != 2), (v40 % ((0) | 1))) >= genFn0((255 < 1), v40))) {
                                            v43 := GenS3{F0: (^(v41)), F1: (v40 * (v40 % ((v40) | 1)))}
                                            _ = v43
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v44 int64 = v12
                                            v45 := GenS5{F0: 42, F1: 32767, F2: (v41 + -2147483648)}
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    fmt.Println(fi39, v40, v3)
                                    v3 = (!true)
                                }
                            } else {
                                v12 += v12
                            }
                            if v3 {
                                _ = len(v1)
                                if false {
                                    for fi46 := 0; fi46 < 7; fi46++ {
                                        if false {
                                            v3 = v3
                                            if ((1000 < 0) && (42 != 255)) {
                                                fmt.Println(v12)
                                                var v47 int16 = int16(v40)
                                                if true {
                                                    v21 += v21
                                                } else {
                                                    if v3 {
                                                        var v48 int8 = int8(fi46)
                                                        var v49 float64 = 1e10
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    v1 = append(v1, 1)
                                                    _ = len(v2)
                                                }
                                                _ = v47
                                            }
                                        } else {
                                            v12 = v12
                                        }
                                        _ = fi46
                                    }
                                    if (!true) {
                                        switch ("" + ("hello" + "foo bar")) {
                                        case "hello":
                                            for fi50 := 0; fi50 < 14; fi50++ {
                                                if v3 {
                                                    fmt.Println(v21)
                                                    v21 += v21
                                                    var v51 int8 = 0
                                                    _ = v51
                                                } else {
                                                    if true {
                                                        var v52 uint = v40
                                                        if (67 != 128) {
                                                            fmt.Println(fi50)
                                                            var v53 int8 = 100
                                                            fmt.Println(v40)
                                                            _ = v53
                                                        } else {
                                                            v21 += -2.25
                                                            var v54 float64 = 1e10
                                                            fmt.Println(fi50, v41, v12)
                                                            var v55 float64 = (((v54 / float64(v21)) / (v54 / v54)) + (float64(v21) + (v54 + v54)))
                                                            v56 := GenS7{F0: (v12 * (v12 + v12)), F1: 325, F2: (v41 % ((1) | 1))}
                                                            v40 -= v40
                                                            _ = len(v1)
                                                            if v3 {
                                                                v21 -= v21
                                                            } else {
                                                                v55 = -2.25
                                                                if len(v2) > 0 {
                                                                    v2[0] = 65535
                                                                }
                                                                v12 = -9223372036854775808
                                                            }
                                                            v2 = append(v2, 2147483647)
                                                            switch 0 {
                                                            case 42:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 127:
                                                                _ = len(v2)
                                                            default:
                                                                if (1 >= genFn1(65535, "foo bar")) {
                                                                    v21 += v21
                                                                    v21 = 0.1
                                                                }
                                                            }
                                                            v57 := GenS8{F0: 42}
                                                            var v58 int64 = v12
                                                            var v59 int8 = -22
                                                            if (true || v3) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                if false {
                                                                    if ("" <= "\n\t") {
                                                                        fmt.Println(v59, v12, v52)
                                                                    } else {
                                                                        v60 := GenS9{F0: 65535, F1: (v41 % ((-508) | 1))}
                                                                        if (((("\n\t" + (((("foo bar" + (("foo bar" + ("" + "\n\t")) + "foo bar")) + "") + (("a" + "a") + ("hello" + ((("\n\t" + "\n\t") + (("" + ("\n\t" + "hello")) + (("" + "foo bar") + ""))) + ("hello" + "hello"))))) + "a")) + "\n\t") + ("\n\t" + "\n\t")) >= "a") {
                                                                            fmt.Println(v12, v52)
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v21 -= (v21 + float32(v55))
                                                                            default:
                                                                                v58 = v12
                                                                            }
                                                                            switch (fi50 >> ((v52) & 63)) {
                                                                            case 1:
                                                                                var v61 uint32 = 0
                                                                                _ = v61
                                                                            default:
                                                                                var v62 uint64 = 4294967295
                                                                                _ = v62
                                                                            }
                                                                        }
                                                                        _ = v60
                                                                    }
                                                                    if (v3 && v3) {
                                                                        var v63 uint8 = 42
                                                                        _ = v63
                                                                    } else {
                                                                        v1 = append(v1, 677)
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 787
                                                                        }
                                                                        var v64 uint32 = 348
                                                                        _ = v64
                                                                    }
                                                                    switch (fi39 % ((fi39) | 1)) {
                                                                    case 461:
                                                                        v41 += (v41 >> (((v52 | 351)) & 63))
                                                                    case -420:
                                                                        v55 += v55
                                                                    case 2147483647:
                                                                        v1 = append(v1, 114)
                                                                    default:
                                                                        if false {
                                                                            fmt.Println(v55, v52)
                                                                        } else {
                                                                            fmt.Println(v59, v52, v41)
                                                                            switch ("foo bar" + "") {
                                                                            case "":
                                                                                fmt.Println(v12, v12)
                                                                            case "hello":
                                                                                var v65 int64 = v12
                                                                                v41 = -993
                                                                                _ = v65
                                                                            case "foo bar":
                                                                                fmt.Println(v59, v52, v58)
                                                                                v66 := GenS9{F0: 499, F1: v41}
                                                                                v67 := GenS3{F0: (v41 | v41), F1: genFn0(v3, (v40 / ((genFn0((!false), (v40 | v40))) | 1)))}
                                                                                _ = v66
                                                                                _ = v67
                                                                            default:
                                                                                v3 = (false || v3)
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v68 uint32 = 115
                                                                    var v69 float32 = 1e10
                                                                    v70 := GenS10{F0: 417, F1: (int64(v68) == 665)}
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = 1
                                                                }
                                                            } else {
                                                                switch "a" {
                                                                case "hello":
                                                                    v1 = append(v1, 2)
                                                                    var v71 int = 32767
                                                                    var v72 string = "hello"
                                                                    v1 = append(v1, 506)
                                                                    v12 = v58
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            }
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        _ = v52
                                                    } else {
                                                        v12 += (v12 + v12)
                                                        switch "" {
                                                        case "\n\t":
                                                            if (834 >= 42) {
                                                                if v3 {
                                                                    var v73 uint8 = 181
                                                                    var v74 int64 = v12
                                                                    if v3 {
                                                                        if v3 {
                                                                            v75 := GenS4{F0: (fi50 >> ((v40) & 63)), F1: (v73 % ((93) | 1))}
                                                                            if (232 > int16(fi39)) {
                                                                                v12 = v12
                                                                                if false {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if (0.0 > v21) {
                                                                                    var v76 int16 = 1000
                                                                                    var v77 uint16 = 1
                                                                                    switch fi50 {
                                                                                    case 429:
                                                                                        var v78 int = (^(fi50))
                                                                                        _ = v78
                                                                                    case -1:
                                                                                        var v79 uint32 = 910
                                                                                        _ = v79
                                                                                    default:
                                                                                        if v3 {
                                                                                            var v80 uint = uint(v77)
                                                                                            if (false && (!(v77 == uint16(v73)))) {
                                                                                                switch "a" {
                                                                                                case "foo bar":
                                                                                                    v3 = ((v41 >> ((4294967295) & 63)) >= (v41 + -113))
                                                                                                }
                                                                                            }
                                                                                            _ = v80
                                                                                        }
                                                                                    }
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                fmt.Println(v12, v41)
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                        v73 += (v73 - (v73 | v73))
                                                                    }
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                                var v81 float64 = 3.14159
                                                                _ = v81
                                                            }
                                                        default:
                                                            v21 = -2.25
                                                        }
                                                        if (!true) {
                                                            v82 := GenS11{F0: 0, F1: "", F2: 32767}
                                                            v3 = (46 <= genFn1(32767, "a"))
                                                            var v83 uint64 = 42
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    }
                                                    _ = len(v2)
                                                }
                                                fmt.Println(v12, v21)
                                                var v84 uint32 = 935
                                                _ = fi50
                                                _ = v84
                                            }
                                        }
                                        fmt.Println(v21, v3, v12)
                                    } else {
                                        for fi85 := 0; fi85 < 18; fi85++ {
                                            if v3 {
                                                switch "foo bar" {
                                                case "":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v86 int16 = 1000
                                                        _ = v86
                                                    case "a":
                                                        v3 = ((false && (!(!true))) && v3)
                                                    }
                                                case "foo bar":
                                                    if (255 > 255) {
                                                        if (v3 || (2147483647 == 65535)) {
                                                            v12 += v12
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        } else {
                                                            fmt.Println(v21)
                                                        }
                                                        var v87 int16 = 1000
                                                        var v88 bool = (!v3)
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        var v89 float64 = 1e10
                                                        v90 := GenS12{F0: v40}
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                case "\n\t":
                                                    var v91 int64 = v12
                                                    _ = v91
                                                default:
                                                    if v3 {
                                                        v2 = append(v2, 2147483648)
                                                        var v92 float64 = 0.1
                                                        _ = v92
                                                    } else {
                                                        v40 += v40
                                                    }
                                                }
                                                v3 = v3
                                                switch "a" {
                                                case "foo bar":
                                                    switch "" {
                                                    case "foo bar":
                                                        v93 := GenS7{F0: v12, F1: 32767, F2: 585}
                                                        _ = v93
                                                    case "hello":
                                                        fmt.Println(fi85, v3, fi39)
                                                    }
                                                case "a":
                                                    fmt.Println(fi39, v21)
                                                    var v94 int = (fi39 / (((fi85 << (((v40 + genFn0((false && (1 >= 1000)), v40))) & 63))) | 1))
                                                    _ = v94
                                                case "":
                                                    v95 := GenS2{F0: -9223372036854775808, F1: -1}
                                                    _ = v95
                                                }
                                                var v96 float64 = 1.5
                                                _ = v96
                                            } else {
                                                v97 := GenS2{F0: -1, F1: v12}
                                                _ = v97
                                            }
                                            if v3 {
                                                fmt.Println(v41, fi85)
                                                v98 := GenS13{F0: (-(fi39))}
                                                var v99 float32 = (float32(3.14159) + 1.5)
                                                _ = v98
                                                _ = v99
                                            } else {
                                                v3 = (v3 && (float32(0.0) <= v21))
                                                var v100 uint64 = 2
                                                _ = v100
                                            }
                                            fmt.Println(v21, v21)
                                            _ = fi85
                                        }
                                    }
                                }
                            } else {
                                _ = len(v1)
                                var v101 uint32 = 1
                                _ = v101
                            }
                            _ = v41
                        }
                        v21 = v21
                        _ = fi39
                        _ = v40
                    }
                }
                _ = v28
            } else {
                if (((!(((v12 | (v12 + v12)) < (-(v12))) && (((v3 || true) || (!v3)) || (((!v3) || false) || v3)))) && v3) && (0 >= 42)) {
                    if v3 {
                        v102 := GenS7{F0: 20, F1: 65535, F2: -321}
                        switch -1 {
                        case -2147483648:
                            switch 32768 {
                            case -4:
                                v21 += v21
                            case -28:
                                var v103 int64 = v12
                                _ = v103
                            case 545:
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                            default:
                                if (!(false && (!(2 <= 96)))) {
                                    var v104 int64 = v12
                                    v105 := GenS14{F0: -257, F1: -9223372036854775808, F2: -1}
                                    for fi106 := 0; fi106 < 15; fi106++ {
                                        _ = len(v1)
                                        for fi107 := 0; fi107 < 11; fi107++ {
                                            switch fi106 {
                                            case -2147483648:
                                                if false {
                                                    var v108 string = ""
                                                    if false {
                                                        v109 := GenS15{F0: 19}
                                                        switch (v108 + v108) {
                                                        case "foo bar":
                                                            switch "\n\t" {
                                                            case "a":
                                                                if v3 {
                                                                    fmt.Println(fi107)
                                                                    var v110 float32 = v21
                                                                    _ = v110
                                                                }
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v111 uint64 = 42
                                                            var v112 float64 = float64(v21)
                                                            var v113 string = v108
                                                            var v114 string = (((v108 + "foo bar") + "") + v113)
                                                            var v115 uint64 = (^(v111))
                                                            v115 = genFn1(42, "\n\t")
                                                            _ = v111
                                                            _ = v112
                                                            _ = v113
                                                            _ = v114
                                                            _ = v115
                                                        case "\n\t":
                                                            if v3 {
                                                                _ = len(v1)
                                                                var v116 float32 = v21
                                                                v108 = "a"
                                                                _ = v116
                                                            } else {
                                                                if (840 <= 32767) {
                                                                    if true {
                                                                        switch 2147483647 {
                                                                        case -2147483648:
                                                                            switch ((v108 + "a") + ("" + v108)) {
                                                                            case "a":
                                                                                v108 = v108
                                                                            case "\n\t":
                                                                                fmt.Println(v104)
                                                                            }
                                                                        case 42:
                                                                            switch (fi107 / (((fi107 ^ fi106)) | 1)) {
                                                                            case 127:
                                                                                v12 -= v104
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                        default:
                                                                            switch 9223372036854775807 {
                                                                            case 42:
                                                                                if v3 {
                                                                                    fmt.Println(v108, v104)
                                                                                } else {
                                                                                    var v117 uint64 = 2
                                                                                    _ = v117
                                                                                }
                                                                                v12 = v12
                                                                                if false {
                                                                                    v118 := GenS16{F0: 0, F1: v104, F2: int8(fi107)}
                                                                                    switch ("" + v108) {
                                                                                    case "":
                                                                                        if v3 {
                                                                                            var v119 string = (v108 + v108)
                                                                                            if v3 {
                                                                                                if (2147483647 >= -415) {
                                                                                                    if (!v3) {
                                                                                                        switch (fi107 | fi107) {
                                                                                                        case 4294967295:
                                                                                                            _ = len(v2)
                                                                                                        case -1:
                                                                                                            v120 := GenS5{F0: 0, F1: -32768, F2: 0}
                                                                                                            _ = v120
                                                                                                        case -729:
                                                                                                            v121 := GenS9{F0: 18, F1: 1}
                                                                                                            _ = v121
                                                                                                        default:
                                                                                                            fmt.Println(v119)
                                                                                                            fmt.Println(v21)
                                                                                                            v12 = -849
                                                                                                        }
                                                                                                        v122 := GenS14{F0: -1, F1: (fi106 << ((4294967295) & 63)), F2: 32767}
                                                                                                        fmt.Println(v119, v21)
                                                                                                        _ = v122
                                                                                                    } else {
                                                                                                        if ((v104 - v12) >= v12) {
                                                                                                            v12 = v104
                                                                                                        } else {
                                                                                                            fmt.Println(fi107)
                                                                                                        }
                                                                                                        v119 = (v108 + v119)
                                                                                                        v123 := GenS17{F0: 127, F1: -522, F2: -380}
                                                                                                        _ = v123
                                                                                                    }
                                                                                                }
                                                                                                switch v108 {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v119, v3)
                                                                                                case "foo bar":
                                                                                                    switch (fi106 / ((fi106) | 1)) {
                                                                                                    case 42:
                                                                                                        v21 = v21
                                                                                                    default:
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    v119 = "\n\t"
                                                                                                }
                                                                                                var v124 float32 = float32(3.14159)
                                                                                                _ = v124
                                                                                            }
                                                                                            v119 = v108
                                                                                            _ = v119
                                                                                        } else {
                                                                                            v125 := GenS18{F0: 1, F1: fi106}
                                                                                            v126 := GenS6{F0: 127}
                                                                                            _ = v125
                                                                                            _ = v126
                                                                                        }
                                                                                    case "hello":
                                                                                        if (false || (0.0 >= 3.14159)) {
                                                                                            switch -226 {
                                                                                            case 4294967295:
                                                                                                v127 := GenS19{F0: 42}
                                                                                                var v128 float32 = float32(1.5)
                                                                                                v108 = (v108 + "hello")
                                                                                                switch (fi107 % (((fi106 % (((fi106 << ((2) & 63))) | 1))) | 1)) {
                                                                                                case 32767:
                                                                                                    var v129 uint8 = 0
                                                                                                    v130 := GenS20{F0: 128}
                                                                                                    var v131 string = "a"
                                                                                                    _ = v129
                                                                                                    _ = v130
                                                                                                    _ = v131
                                                                                                case 32768:
                                                                                                    var v132 float64 = 0.0
                                                                                                    _ = v132
                                                                                                default:
                                                                                                    fmt.Println(v3, v108)
                                                                                                }
                                                                                                fmt.Println(fi106, v21)
                                                                                                _ = v127
                                                                                                _ = v128
                                                                                            case 226:
                                                                                                v133 := GenS14{F0: -44, F1: fi106, F2: -1}
                                                                                                _ = v133
                                                                                            case 127:
                                                                                                v12 -= (v104 + v12)
                                                                                            }
                                                                                            if v3 {
                                                                                                v3 = (!false)
                                                                                                v21 += float32(-2.25)
                                                                                                v134 := GenS11{F0: 4294967295, F1: v108, F2: -416}
                                                                                                _ = v134
                                                                                            } else {
                                                                                                fmt.Println(v108, v108)
                                                                                            }
                                                                                            v135 := GenS21{F0: 53}
                                                                                            _ = v135
                                                                                        }
                                                                                    }
                                                                                    var v136 int = fi107
                                                                                    _ = v118
                                                                                    _ = v136
                                                                                } else {
                                                                                    var v137 int = (fi107 & (fi106 % (((fi107 ^ -2147483648)) | 1)))
                                                                                    v12 = 375
                                                                                    _ = v137
                                                                                }
                                                                            case -379:
                                                                                v138 := GenS15{F0: 127}
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                    }
                                                                    if v3 {
                                                                        if (127 <= 49) {
                                                                            v104 = v104
                                                                            fmt.Println(v21)
                                                                            v108 = v108
                                                                            fmt.Println(v12, v21)
                                                                        } else {
                                                                            v2 = append(v2, 871)
                                                                            var v139 int = 2147483647
                                                                            _ = v139
                                                                        }
                                                                        var v140 bool = (!v3)
                                                                        var v141 uint64 = 404
                                                                        _ = v140
                                                                        _ = v141
                                                                    } else {
                                                                        var v142 int32 = -2147483648
                                                                        _ = v142
                                                                    }
                                                                } else {
                                                                    var v143 uint16 = 221
                                                                    var v144 int8 = 1
                                                                    _ = v143
                                                                    _ = v144
                                                                }
                                                                fmt.Println(v108, fi107)
                                                                if (fi106 == 4294967295) {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 65535
                                                                    }
                                                                    fmt.Println(v108, v21)
                                                                    if (4294967295 <= 4294967295) {
                                                                        var v145 uint16 = 2
                                                                        v145 += (v145 ^ v145)
                                                                        var v146 uint = uint(v145)
                                                                        _ = v145
                                                                        _ = v146
                                                                    }
                                                                    v21 = v21
                                                                    v3 = (4294967295 != 865)
                                                                } else {
                                                                    if v3 {
                                                                        var v147 float64 = 1.5
                                                                        _ = v147
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v1 = append(v1, 777)
                                                        _ = v109
                                                    } else {
                                                        if (true && true) {
                                                            v104 = v12
                                                            var v148 bool = ((("a" + ((((("hello" + (v108 + v108)) + ((v108 + v108) + v108)) + v108) + v108) + v108)) + "hello") >= v108)
                                                            var v149 uint8 = 42
                                                            fmt.Println(v12, v104, v148)
                                                            _ = v148
                                                            _ = v149
                                                        }
                                                        if false {
                                                            _ = len(v1)
                                                            var v150 uint = genFn0(v3, 65535)
                                                            _ = v150
                                                        }
                                                        if (!(0 == 0)) {
                                                            var v151 int64 = 911
                                                            _ = v151
                                                        } else {
                                                            if v3 {
                                                                fmt.Println(v104, v104)
                                                            }
                                                            v21 -= (0.0 + v21)
                                                            v3 = ((!false) && (!((v108 < "") && ((!((^(fi107)) >= fi106)) || true))))
                                                            v3 = (17 == 42)
                                                            v3 = (128 == 1)
                                                            v152 := GenS7{F0: v104, F1: 42, F2: -1}
                                                            _ = v152
                                                        }
                                                        v3 = (2 > 255)
                                                        var v153 uint8 = 1
                                                        _ = v153
                                                    }
                                                    _ = v108
                                                } else {
                                                    if (-713 != (fi107 * (fi106 | fi107))) {
                                                        var v154 int32 = 42
                                                        if len(v2) > 0 {
                                                            v2[0] = 818
                                                        }
                                                        v104 += (v104 ^ (v12 & (-(v12))))
                                                        v21 -= (float32((3.14159 * (0.0 / (1.5 - 3.14159)))) / (((v21 - v21) * (float32(1.5) + v21)) * 1e10))
                                                        v21 = v21
                                                        _ = v154
                                                    } else {
                                                        _ = len(v2)
                                                        if v3 {
                                                            var v155 uint64 = 18446744073709551615
                                                            _ = v155
                                                        } else {
                                                            var v156 int16 = -591
                                                            var v157 int8 = 0
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                        v2 = append(v2, 65535)
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v12 = v104
                                                }
                                                fmt.Println(fi107, v21, v104)
                                                if v3 {
                                                    v158 := GenS1{F0: -100, F1: -425}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v158
                                                }
                                            case 0:
                                                var v159 int32 = -400
                                                var v160 uint8 = 255
                                                _ = v159
                                                _ = v160
                                            case 9223372036854775807:
                                                var v161 int16 = -32768
                                                _ = v161
                                            }
                                            switch fi106 {
                                            case 32767:
                                                _ = len(v1)
                                            case -2147483648:
                                                if v3 {
                                                    switch fi106 {
                                                    case 807:
                                                        switch "a" {
                                                        case "foo bar":
                                                            if v3 {
                                                                var v162 int64 = (v12 & v12)
                                                                _ = v162
                                                            } else {
                                                                v1 = append(v1, 2147483648)
                                                                var v163 uint = 267
                                                                _ = v163
                                                            }
                                                        case "hello":
                                                            v3 = v3
                                                        }
                                                    }
                                                }
                                            case 500:
                                                v104 += -1
                                            default:
                                                fmt.Println(v12)
                                            }
                                            fmt.Println(v3, v104)
                                            _ = fi107
                                        }
                                        if true {
                                            var v164 int = 530
                                            if (42 == -52) {
                                                _ = len(v2)
                                                v165 := GenS22{F0: 0}
                                                for fi166 := 0; fi166 < 2; fi166++ {
                                                    if (!(805 == 4294967295)) {
                                                        var v167 int8 = 42
                                                        var v168 int = 2147483647
                                                        _ = v167
                                                        _ = v168
                                                    }
                                                    _ = fi166
                                                }
                                                _ = v165
                                            }
                                            _ = v164
                                        }
                                        _ = fi106
                                    }
                                    _ = v104
                                    _ = v105
                                } else {
                                    if ((-1 > 32767) && v3) {
                                        var v169 uint8 = 141
                                        fmt.Println(v169)
                                        _ = v169
                                    }
                                }
                            }
                        case 444:
                            v21 -= (1e10 + v21)
                        case 42:
                            v170 := GenS22{F0: 400}
                            _ = v170
                        default:
                            fmt.Println(v12)
                        }
                        _ = v102
                    } else {
                        for fi171 := 0; fi171 < 20; fi171++ {
                            v3 = (100 != 0)
                            switch -497 {
                            case 9223372036854775807:
                                var v172 uint64 = 4294967295
                                _ = v172
                            case 42:
                                v173 := GenS23{F0: 65535}
                                _ = v173
                            case -838:
                                if ("a" == (((("a" + "a") + "a") + "\n\t") + "")) {
                                    if v3 {
                                        var v174 float64 = float64((v21 - 1.5))
                                        var v175 uint8 = 42
                                        _ = v174
                                        _ = v175
                                    }
                                } else {
                                    v3 = (!v3)
                                }
                            }
                            _ = fi171
                        }
                    }
                    var v176 int16 = -1
                    _ = v176
                } else {
                    if v3 {
                        v21 -= float32(1.5)
                        v3 = (((((float64(0.1) - 1.5) * float64(((1e10 * v21) / (v21 + v21)))) > float64(v21)) || (!(42 >= 128))) || (155 < 796))
                        v3 = (!false)
                    }
                }
            }
            fmt.Println(v12, v21, v3)
        }
        _ = v21
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!v3) {
        switch ("hello" + "\n\t") {
        case "hello":
            var v177 int16 = 1000
            _ = v177
        }
    } else {
        for fi178 := 0; fi178 < 19; fi178++ {
            for fi179 := 0; fi179 < 19; fi179++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v180 int32 = 2147483647
                _ = fi179
                _ = v180
            }
            _ = fi178
        }
    }
    switch "" {
    case "foo bar":
        var v181 string = "\n\t"
        _ = v181
    }
    v12 -= v12
    for fi182 := 0; fi182 < 20; fi182++ {
        for fi183 := 0; fi183 < 14; fi183++ {
            fmt.Println(fi183, v12, fi182)
            v12 += -9223372036854775808
            _ = fi183
        }
        _ = fi182
    }
    var v184 int = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v184:", v184)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v184
    _ = v2
    _ = v3
}
