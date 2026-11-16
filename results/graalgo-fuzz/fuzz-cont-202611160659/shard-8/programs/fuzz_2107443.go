package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int64; F1 string; F2 string }
type GenS8 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    v1 := []int{2147483647}
    var v2 float32 = (1e10 + 1.5)
    var v3 int16 = 1
    for fi4 := 0; fi4 < 18; fi4++ {
        if false {
            v2 += -2.25
        } else {
            if (((0.1 - 0.1) >= 0.1) || false) {
                v5 := GenS1{F0: ("\n\t" + "hello"), F1: false}
                _ = v5
            } else {
                var v6 uint64 = 430
                _ = v6
            }
            var v7 int = fi4
            _ = v7
        }
        if (1 < 139) {
            var v8 bool = (134 <= 128)
            var v9 float64 = 3.14159
            switch "hello" {
            case "\n\t":
                v9 += v9
            case "":
                for fi10 := 0; fi10 < 2; fi10++ {
                    if v8 {
                        var v11 float64 = v9
                        var v12 int64 = 0
                        var v13 float64 = -2.25
                        if (!(-847 == -437)) {
                            v2 -= ((3.14159 * ((v2 * 1.5) - v2)) / (v2 / float32(-2.25)))
                        }
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v8, v11, v12)
                        case "\n\t":
                            fmt.Println(fi4)
                        default:
                            var v14 uint8 = 0
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (!(4294967295 != 869)) {
                        if (!((2 < 42) && (!v8))) {
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            v15 := GenS1{F0: "\n\t", F1: false}
                            if (4294967295 == 332) {
                                fmt.Println(fi4)
                                if (v8 || (5 != -128)) {
                                    switch "" {
                                    case "a":
                                        fmt.Println(v8)
                                    default:
                                        v16 := GenS2{F0: 255, F1: 1, F2: ("hello" + "a")}
                                        _ = v16
                                    }
                                    fmt.Println(v3)
                                    v9 = (v9 - v9)
                                } else {
                                    _ = len(v1)
                                }
                                v3 = v3
                                var v17 uint32 = 810
                                switch "foo bar" {
                                case "foo bar":
                                    var v18 int = (fi10 - fi4)
                                    _ = v18
                                case "a":
                                    fmt.Println(v3, v3, v3)
                                case "hello":
                                    var v19 uint16 = 65535
                                    _ = v19
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v17
                            } else {
                                var v20 int32 = 0
                                v2 = v2
                                v9 = v9
                                _ = v20
                            }
                            _ = v15
                        }
                        var v21 uint64 = 2
                        _ = v21
                    } else {
                        if v8 {
                            v2 += 1e10
                        } else {
                            var v22 uint = 255
                            v1 = append(v1, (fi4 >> ((765) & 63)))
                            var v23 string = ("foo bar" + "hello")
                            var v24 uint64 = 42
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        if (-477 != -1) {
                            v3 = v3
                            var v25 uint = 0
                            _ = v25
                        }
                        v8 = (4 >= 128)
                    }
                    _ = fi10
                }
            }
            _ = v8
            _ = v9
        } else {
            v26 := GenS2{F0: 128, F1: 1, F2: ("\n\t" + "\n\t")}
            v27 := GenS2{F0: 80, F1: 42, F2: "foo bar"}
            _ = v26
            _ = v27
        }
        switch "" {
        case "hello":
            var v28 int16 = (v3 / ((v3) | 1))
            _ = v28
        case "a":
            fmt.Println(v2, v2)
        case "foo bar":
            var v29 bool = (65535 == 65535)
            _ = v29
        }
        _ = fi4
    }
    fmt.Println(v3, v3)
    v30 := GenS3{F0: 451, F1: 174}
    var v31 uint8 = 42
    for fi32 := 0; fi32 < 1; fi32++ {
        v2 -= -2.25
        fmt.Println(v31, v3)
        if false {
            switch "hello" {
            case "a":
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 uint32 = 42
                    v34 -= (v34 * uint32(v31))
                    _ = fi33
                    _ = v34
                }
            case "foo bar":
                if (v31 != 0) {
                    v35 := GenS4{F0: 69, F1: v31}
                    if false {
                        var v36 int32 = 378
                        _ = v36
                    }
                    v37 := GenS2{F0: (^(v31)), F1: -1, F2: "hello"}
                    _ = v35
                    _ = v37
                } else {
                    v38 := GenS5{F0: 152}
                    _ = v38
                }
            case "":
                switch "\n\t" {
                case "":
                    for fi39 := 0; fi39 < 19; fi39++ {
                        v3 += 32767
                        v1 = append(v1, 4294967295)
                        v40 := GenS3{F0: 1, F1: (fi39 * (fi32 - (fi32 & fi39)))}
                        _ = fi39
                        _ = v40
                    }
                case "hello":
                    for fi41 := 0; fi41 < 1; fi41++ {
                        if len(v1) > 0 {
                            v1[0] = fi32
                        }
                        _ = fi41
                    }
                default:
                    var v42 bool = true
                    _ = v42
                }
            default:
                v3 -= 678
            }
            var v43 uint16 = 42
            v44 := GenS6{F0: 2147483647, F1: 8}
            _ = v43
            _ = v44
        }
        _ = fi32
    }
    if true {
        for fi45 := 0; fi45 < 11; fi45++ {
            var v46 uint32 = 0
            if (456 != 1) {
                for fi47 := 0; fi47 < 13; fi47++ {
                    switch fi47 {
                    case 9223372036854775807:
                        switch "a" {
                        case "foo bar":
                            v48 := GenS2{F0: (v31 % ((250) | 1)), F1: 42, F2: ("a" + "")}
                            _ = v48
                        case "a":
                            v46 -= (^(v46))
                        default:
                            var v49 int16 = 207
                            _ = v49
                        }
                    case 286:
                        switch fi47 {
                        case 0:
                            switch "\n\t" {
                            case "a":
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = (fi47 / ((-2147483648) | 1))
                                    }
                                    if false {
                                        var v50 int16 = v3
                                        if true {
                                            fmt.Println(v31, v2, fi47)
                                        } else {
                                            var v51 uint64 = 0
                                            if len(v1) > 0 {
                                                v1[0] = 380
                                            }
                                            var v52 int16 = 1000
                                            switch fi45 {
                                            case 0:
                                                v46 -= v46
                                            case 246:
                                                fmt.Println(v46)
                                            default:
                                                var v53 uint32 = v46
                                                _ = v53
                                            }
                                            var v54 uint64 = v51
                                            _ = v51
                                            _ = v52
                                            _ = v54
                                        }
                                        _ = v50
                                    } else {
                                        v55 := GenS7{F0: 9223372036854775807, F1: ("foo bar" + "hello"), F2: (("a" + "\n\t") + "a")}
                                        if (false || false) {
                                            switch ("foo bar" + "\n\t") {
                                            case "":
                                                v31 -= 42
                                            case "hello":
                                                _ = len(v1)
                                            case "\n\t":
                                                v56 := GenS5{F0: v46}
                                                _ = v56
                                            }
                                            v1 = append(v1, (fi47 / ((-240) | 1)))
                                            if true {
                                                fmt.Println(fi47, v46, v31)
                                                if (!false) {
                                                    v2 = -2.25
                                                    if len(v1) > 0 {
                                                        v1[0] = fi45
                                                    }
                                                    var v57 string = ("hello" + ("foo bar" + (("foo bar" + "") + "")))
                                                    _ = v57
                                                }
                                                fmt.Println(v3, v31, v46)
                                            } else {
                                                if (((1 == 2) && true) || false) {
                                                    var v58 int = (fi45 - 4294967295)
                                                    fmt.Println(v3)
                                                    _ = v58
                                                }
                                                fmt.Println(v3)
                                            }
                                        }
                                        if (18446744073709551615 <= 4294967295) {
                                            if (!true) {
                                                switch "hello" {
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if true {
                                                            v1 = append(v1, (fi47 << ((938) & 63)))
                                                            v2 = 0.0
                                                            var v59 uint = 2
                                                            _ = v59
                                                        }
                                                    case "hello":
                                                        v2 += (v2 + (0.1 / 3.14159))
                                                    }
                                                case "":
                                                    if (!((v31 & v31) == v31)) {
                                                        var v60 uint8 = (v31 / (((v31 & v31)) | 1))
                                                        var v61 uint = 42
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        switch "\n\t" {
                                                        case "a":
                                                            v1 = append(v1, (fi47 * (fi47 | (fi47 * 4294967295))))
                                                        case "foo bar":
                                                            if false {
                                                                if false {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (-(fi45))
                                                                    }
                                                                    v46 = (v46 % ((v46) | 1))
                                                                } else {
                                                                    var v62 bool = true
                                                                    _ = v62
                                                                }
                                                                var v63 uint64 = 596
                                                                v64 := GenS7{F0: -1, F1: (("a" + "\n\t") + "a"), F2: "hello"}
                                                                v63 += v63
                                                                fmt.Println(v31, v2, v2)
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                        case "\n\t":
                                                            v65 := GenS3{F0: 871, F1: (fi45 | fi45)}
                                                            _ = v65
                                                        default:
                                                            v1 = append(v1, 32767)
                                                        }
                                                        var v66 uint64 = uint64(v31)
                                                        fmt.Println(v66)
                                                        _ = v66
                                                    }
                                                }
                                                switch (fi47 + (^(fi45))) {
                                                case -9223372036854775808:
                                                    var v67 uint8 = (v31 + (v31 - (v31 & (v31 & v31))))
                                                    _ = v67
                                                case 127:
                                                    v31 = (^(v31))
                                                case -157:
                                                    v68 := GenS6{F0: 648, F1: 1}
                                                    _ = v68
                                                default:
                                                    switch "" {
                                                    case "\n\t":
                                                        v69 := GenS6{F0: 191, F1: 604}
                                                        _ = v69
                                                    case "hello":
                                                        var v70 uint8 = (v31 * (v31 & 0))
                                                        _ = v70
                                                    case "a":
                                                        var v71 int64 = 816
                                                        switch fi47 {
                                                        case 1:
                                                            v3 += (v3 % ((v3) | 1))
                                                        }
                                                        var v72 int64 = v71
                                                        var v73 int32 = 42
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    default:
                                                        if len(v1) > 0 {
                                                            v1[0] = 2147483647
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi47)
                                            }
                                        } else {
                                            var v74 float32 = 1e10
                                            v46 += (v46 * (v46 + v46))
                                            var v75 int64 = -989
                                            _ = v74
                                            _ = v75
                                        }
                                        _ = v55
                                    }
                                    v2 = v2
                                } else {
                                    if (892 <= 280) {
                                        var v76 uint8 = 42
                                        _ = len(v1)
                                        var v77 float32 = float32(-2.25)
                                        _ = v76
                                        _ = v77
                                    } else {
                                        v78 := GenS5{F0: 4294967295}
                                        var v79 float64 = 3.14159
                                        _ = v78
                                        _ = v79
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (fi47 * fi45)
                                    }
                                }
                            case "hello":
                                var v80 int32 = 558
                                _ = v80
                            }
                            v1 = append(v1, (fi45 ^ (^(fi47))))
                            if false {
                                var v81 uint = 42
                                if ((((!(v2 > v2)) || true) && (false || true)) && ((true && (true || (v3 == v3))) && (367 != int32(fi45)))) {
                                    v2 = 1.5
                                }
                                _ = len(v1)
                                if len(v1) > 0 {
                                    v1[0] = (-(fi45))
                                }
                                if len(v1) > 0 {
                                    v1[0] = 32768
                                }
                                var v82 uint = v81
                                switch ("\n\t" + "a") {
                                case "a":
                                    v1 = append(v1, 32768)
                                    v83 := GenS2{F0: v31, F1: 42, F2: "foo bar"}
                                    fmt.Println(v2, v46, v82)
                                    v84 := GenS8{F0: uint8(v46), F1: 18446744073709551615, F2: (-(v3))}
                                    switch (fi47 & (fi47 >> ((v81) & 63))) {
                                    case 399:
                                        v85 := GenS9{F0: v82}
                                        _ = v85
                                    case -9223372036854775808:
                                        switch (fi47 / ((fi45) | 1)) {
                                        case 42:
                                            v82 += v82
                                        }
                                    case 42:
                                        switch 234 {
                                        case 32768:
                                            if (!false) {
                                                v1 = append(v1, -2147483648)
                                                v31 -= v31
                                            }
                                            v46 += (v46 | (v46 - (v46 >> (((v81 * 1)) & 63))))
                                            var v86 float32 = float32(1.5)
                                            _ = v86
                                        case 42:
                                            v31 -= v31
                                        default:
                                            var v87 string = "\n\t"
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v87
                                        }
                                    default:
                                        v88 := GenS4{F0: 2147483647, F1: 66}
                                        _ = v88
                                    }
                                    _ = v83
                                    _ = v84
                                default:
                                    v2 = -2.25
                                    fmt.Println(fi45, v2, v31)
                                    switch "hello" {
                                    case "a":
                                        if (v31 == v31) {
                                            v82 = 962
                                        } else {
                                            v89 := GenS10{F0: 430}
                                            if true {
                                                v90 := GenS11{F0: (v31 >> ((1) & 63)), F1: true}
                                                fmt.Println(fi45, v46)
                                                v91 := GenS7{F0: 42, F1: ("a" + ("hello" + "\n\t")), F2: ("foo bar" + ("\n\t" + ("foo bar" + "foo bar")))}
                                                _ = v90
                                                _ = v91
                                            } else {
                                                var v92 int64 = 661
                                                v93 := GenS8{F0: 127, F1: 624, F2: (v3 + (v3 - 0))}
                                                var v94 string = "hello"
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                            _ = v89
                                        }
                                    }
                                }
                                _ = v81
                                _ = v82
                            } else {
                                fmt.Println(v2)
                                if false {
                                    v1 = append(v1, fi45)
                                    fmt.Println(v31)
                                    if (!false) {
                                        v95 := GenS9{F0: 81}
                                        if false {
                                            v96 := GenS4{F0: 42, F1: v31}
                                            v1 = append(v1, (fi47 << ((255) & 63)))
                                            v97 := GenS3{F0: 65535, F1: fi47}
                                            _ = v96
                                            _ = v97
                                        }
                                        v98 := GenS6{F0: 42, F1: 1}
                                        fmt.Println(fi45, v2, v3)
                                        v3 = (v3 >> ((255) & 63))
                                        _ = v95
                                        _ = v98
                                    }
                                }
                                v2 = float32(1.5)
                                v1 = append(v1, 127)
                            }
                            var v99 int8 = 127
                            v3 -= v3
                            _ = v99
                        case -2147483648:
                            var v100 uint64 = 4294967295
                            _ = v100
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        fmt.Println(fi45, fi45)
                    }
                    v2 -= 0.1
                    v101 := GenS8{F0: (v31 / ((128) | 1)), F1: 4294967295, F2: (v3 - -32768)}
                    _ = fi47
                    _ = v101
                }
                fmt.Println(v3)
                fmt.Println(v46, v46)
            }
            _ = fi45
            _ = v46
        }
        v3 = -760
        _ = len(v1)
    } else {
        var v102 string = (("foo bar" + "\n\t") + "a")
        _ = v102
    }
    v103 := GenS12{F0: 4294967295, F1: int8(v3), F2: v31}
    v2 = 3.14159
    var v104 uint8 = 127
    v3 = v3
    v3 += v3
    if false {
        for fi105 := 0; fi105 < 3; fi105++ {
            switch (fi105 % ((fi105) | 1)) {
            case -427:
                v3 -= v3
            case 32767:
                if true {
                    v104 = (v104 ^ v104)
                }
            }
            var v106 int64 = 1
            if (false || (v2 == float32(1.5))) {
                switch 127 {
                case -858:
                    for fi107 := 0; fi107 < 10; fi107++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi107
                    }
                case -88:
                    fmt.Println(v31, v2, fi105)
                default:
                    if false {
                        _ = len(v1)
                        var v108 float64 = 1.5
                        v109 := GenS9{F0: 18446744073709551615}
                        _ = v108
                        _ = v109
                    } else {
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v110 uint32 = 1
                        _ = v110
                    }
                }
            } else {
                v31 += 128
            }
            _ = fi105
            _ = v106
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
