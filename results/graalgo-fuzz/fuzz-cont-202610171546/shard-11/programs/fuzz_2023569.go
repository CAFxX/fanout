package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 uint; F1 int64; F2 bool }
type GenS13 struct { F0 int32; F1 uint }
type GenS14 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS15 struct { F0 string; F1 int16; F2 uint }
type GenS16 struct { F0 string; F1 int8; F2 uint8 }
type GenS17 struct { F0 int; F1 string }
type GenS18 struct { F0 string; F1 int16 }
type GenS19 struct { F0 uint32; F1 bool }
type GenS20 struct { F0 int64; F1 int }
type GenS21 struct { F0 int16; F1 string }
type GenS22 struct { F0 int64; F1 int; F2 uint32 }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 uint64; F1 int8 }
type GenS25 struct { F0 bool }
type GenS26 struct { F0 string; F1 uint16; F2 int8 }
func genFn0(p0 int16) int32 {
    p0 = (p0 | p0)
    var v7 int = 9223372036854775807
    switch "hello" {
    case "\n\t":
        var v8 uint32 = 4294967295
        _ = v8
    case "foo bar":
        v9 := []int{32768, (v7 >> ((65535) & 63))}
        _ = v9
    case "":
        v7 = (v7 | v7)
    }
    v10 := GenS1{F0: 65535, F1: 195}
    _ = p0
    _ = v10
    _ = v7
    return -2147483648
}

func main() {
    var v1 int16 = 1000
    var v2 int64 = 1
    var v3 int64 = v2
    v4 := make(map[string]bool)
    var v5 float32 = (0.1 - 0.1)
    var v6 uint32 = 2
    switch 0 {
    case -585:
        switch "a" {
        case "a":
            var v11 float64 = float64(v5)
            _ = v11
        }
    case -288:
        fmt.Println(v5)
    default:
        var v12 bool = true
        _ = v12
    }
    v13 := GenS1{F0: (v6 * v6), F1: 2}
    switch 32768 {
    case 1:
        var v14 int = 128
        _ = v14
    case 0:
        if true {
            v15 := []int{-1, 1, -2147483648}
            var v16 int16 = v1
            switch 128 {
            case 0:
                for fi17 := 0; fi17 < 2; fi17++ {
                    fmt.Println(v1, v16)
                    v18 := GenS1{F0: (v6 / (((v6 & 252)) | 1)), F1: 0}
                    _ = fi17
                    _ = v18
                }
            case 2147483647:
                if (123 <= 128) {
                    v19 := GenS1{F0: v6, F1: 1}
                    for fi20 := 0; fi20 < 13; fi20++ {
                        if false {
                            _ = len(v15)
                        } else {
                            v21 := GenS1{F0: (v6 & 270), F1: 1}
                            var v22 uint32 = (v6 * (v6 * 42))
                            for fi23 := 0; fi23 < 2; fi23++ {
                                var v24 int8 = 100
                                switch "a" {
                                case "":
                                    switch 1 {
                                    case -2147483648:
                                        if (true && (1 >= 829)) {
                                            if ((float64((v5 + float32(-2.25))) == float64((((v5 - v5) / 1.5) - float32(float64(float32(((((0.0 - (3.14159 + (3.14159 + 1.5))) + 0.1) + 0.0) - float64(v5)))))))) || true) {
                                                var v25 uint8 = 0
                                                var v26 int64 = v2
                                                _ = v25
                                                _ = v26
                                            }
                                            v27 := GenS2{F0: (fi23 % ((fi23) | 1)), F1: 183, F2: v22}
                                            _ = v27
                                        } else {
                                            v22 += (v6 / ((v6) | 1))
                                        }
                                    case 9223372036854775807:
                                        var v28 int8 = v24
                                        _ = v28
                                    case 32767:
                                        v29 := GenS2{F0: 9223372036854775807, F1: 128, F2: 2147483647}
                                        _ = v29
                                    }
                                default:
                                    switch "" {
                                    case "hello":
                                        if false {
                                            v30 := GenS3{F0: v3, F1: (fi20 - (fi23 >> ((4294967295) & 63))), F2: (^(v24))}
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                            _ = v30
                                        } else {
                                            var v31 int16 = v16
                                            _ = v31
                                        }
                                    }
                                }
                                _ = fi23
                                _ = v24
                            }
                            _ = v21
                            _ = v22
                        }
                        _ = fi20
                    }
                    _ = v19
                } else {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        fmt.Println(v3)
                        _ = fi32
                    }
                    for fi33 := 0; fi33 < 13; fi33++ {
                        v1 = v1
                        _ = fi33
                    }
                }
            }
            _ = v15
            _ = v16
        } else {
            var v34 int16 = 32767
            _ = v34
        }
    }
    v3 = 0
    if ((false || true) || false) {
        var v35 float32 = 1e10
        _ = v35
    } else {
        if ((((v6 - v6) < (^(v6))) && ((v3 + v3) < 9223372036854775807)) && true) {
            v36 := GenS3{F0: (v3 & v3), F1: 2147483647, F2: 100}
            switch 36 {
            case 1:
                var v37 uint16 = 65535
                _ = v37
            case 4294967295:
                switch "a" {
                case "a":
                    v5 = 1.5
                case "foo bar":
                    v2 += 0
                }
            }
            _ = v36
        }
    }
    fmt.Println(v6, v2, v2)
    v1 = v1
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 uint64 = 42
        _ = fi38
        _ = v39
    }
    v40 := GenS1{F0: (v6 << ((674) & 63)), F1: 31}
    v1 -= (v1 - -32768)
    v41 := []int{1, 9223372036854775807}
    var v42 string = ((("hello" + (((("foo bar" + "\n\t") + "") + "") + ("foo bar" + "\n\t"))) + (("hello" + (("a" + "hello") + ("\n\t" + "hello"))) + ("" + "a"))) + "")
    v43 := GenS4{F0: 42, F1: (v42 + v42), F2: (v2 + v3)}
    for fi44 := 0; fi44 < 14; fi44++ {
        for fi45 := 0; fi45 < 9; fi45++ {
            if true {
                fmt.Println(fi45)
                fmt.Println(v2)
                fmt.Println(fi45, fi45, v1)
            }
            if (false && false) {
                switch (fi45 & fi44) {
                case 0:
                    v46 := GenS2{F0: 32767, F1: 255, F2: (v6 ^ 2147483648)}
                    _ = v46
                case -9223372036854775808:
                    var v47 string = v42
                    _ = v47
                default:
                    var v48 string = v42
                    _ = v48
                }
                switch "a" {
                case "hello":
                    var v49 bool = true
                    _ = v49
                case "\n\t":
                    v1 += (v1 + v1)
                }
                var v50 int32 = -619
                var v51 int = fi45
                _ = v50
                _ = v51
            } else {
                var v52 uint = 65535
                fmt.Println(v1, fi44, v1)
                v41 = append(v41, (fi45 + 42))
                _ = v52
            }
            var v53 uint64 = 237
            _ = fi45
            _ = v53
        }
        fmt.Println(v5, v1)
        _ = fi44
    }
    switch 625 {
    case 553:
        var v54 uint8 = 1
        _ = v54
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        switch "foo bar" {
        case "\n\t":
            for fi56 := 0; fi56 < 18; fi56++ {
                v57 := GenS5{F0: (v1 << ((465) & 63)), F1: (^(v1))}
                _ = fi56
                _ = v57
            }
        case "a":
            var v58 uint = 38
            _ = v58
        case "hello":
            var v59 uint32 = v6
            _ = v59
        default:
            var v60 int64 = (v3 >> ((65535) & 63))
            _ = v60
        }
        if ((-(v2)) >= v3) {
            var v61 float32 = -2.25
            _ = v61
        }
        var v62 uint16 = uint16(v6)
        _ = fi55
        _ = v62
    }
    for fi63 := 0; fi63 < 14; fi63++ {
        for fi64 := 0; fi64 < 11; fi64++ {
            if true {
                switch (fi63 << ((255) & 63)) {
                case 0:
                    if ((fi64 << ((255) & 63)) <= fi63) {
                        v2 -= 11
                    }
                case 32767:
                    v65 := GenS1{F0: 65535, F1: 65}
                    _ = v65
                case 377:
                    if (!(2147483647 < (v6 - (v6 ^ (v6 / ((v6) | 1)))))) {
                        fmt.Println(fi64)
                        v4[((v42 + (v42 + "hello")) + "hello")] = ((((v42 + v42) + "a") + v42) <= "foo bar")
                        var v66 float64 = (1e10 / 3.14159)
                        var v67 int16 = -565
                        _ = v66
                        _ = v67
                    } else {
                        v6 = v6
                        if (false && (!(42 != 431))) {
                            var v68 int32 = 0
                            fmt.Println(v3, v1)
                            if ((503 >= 42) && (!(255 <= 42))) {
                                v68 = v68
                                v5 += (((v5 * v5) - v5) - float32(((3.14159 - 0.1) / 0.1)))
                                var v69 int8 = -1
                                delete(v4, "a")
                                if (((0 < v6) && (!true)) && false) {
                                    if (fi63 != 128) {
                                        var v70 bool = ((-(fi64)) >= (fi63 >> ((255) & 63)))
                                        _ = v70
                                    }
                                }
                                var v71 uint8 = 128
                                _ = v69
                                _ = v71
                            }
                            if false {
                                v72 := GenS1{F0: (^(v6)), F1: 2}
                                if false {
                                    _ = len(v4)
                                    _ = len(v41)
                                } else {
                                    if true {
                                        var v73 float32 = 0.1
                                        _ = v73
                                    } else {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v1 += v1
                                        case "hello":
                                            switch "" {
                                            case "foo bar":
                                                v6 += (v6 ^ v6)
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v74 uint8 = 42
                                                    _ = v74
                                                case "hello":
                                                    v1 = 42
                                                case "\n\t":
                                                    if (((fi63 >> ((42) & 63)) == fi64) || false) {
                                                        v75 := v4[v42]
                                                        v2 -= (v3 - int64(v68))
                                                        v6 = (^(v6))
                                                        var v76 int = fi63
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        if (7 >= 75) {
                                                            var v77 uint64 = 18446744073709551615
                                                            if (true || ((false || ((^(v77)) > v77)) || ((v68 < (v68 % (((v68 % ((-822) | 1))) | 1))) || true))) {
                                                                if len(v41) > 0 {
                                                                    v41[0] = int(v77)
                                                                }
                                                                v78 := v4[v42]
                                                                v5 = v5
                                                                v1 -= v1
                                                                fmt.Println(v1, v42, v6)
                                                                _ = v78
                                                            }
                                                            v79 := GenS6{F0: v68}
                                                            _ = v77
                                                            _ = v79
                                                        } else {
                                                            if false {
                                                                var v80 int8 = 46
                                                                fmt.Println(v42)
                                                                v81 := GenS1{F0: (v6 >> ((932) & 63)), F1: 128}
                                                                if false {
                                                                    if (!true) {
                                                                        var v82 int = (fi63 * (fi63 * fi63))
                                                                        var v83 uint8 = 199
                                                                        v84 := v4[""]
                                                                        _ = v82
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                    switch v42 {
                                                                    case "foo bar":
                                                                        v85 := GenS3{F0: v3, F1: 729, F2: (v80 & (v80 << ((uint(v6)) & 63)))}
                                                                        _ = v85
                                                                    default:
                                                                        v86 := GenS1{F0: 196, F1: 234}
                                                                        _ = v86
                                                                    }
                                                                    v2 = (v2 * v2)
                                                                }
                                                                switch (-(fi63)) {
                                                                case 598:
                                                                    var v87 float64 = 3.14159
                                                                    _ = v87
                                                                case -1:
                                                                    switch 32768 {
                                                                    case 9223372036854775807:
                                                                        var v88 int8 = -1
                                                                        _ = v88
                                                                    case 32767:
                                                                        v89 := GenS3{F0: 42, F1: (fi63 ^ 2147483647), F2: -1}
                                                                        _ = v89
                                                                    }
                                                                case -2147483648:
                                                                    switch fi63 {
                                                                    case 32768:
                                                                        var v90 int64 = v2
                                                                        _ = v90
                                                                    default:
                                                                        v68 = v68
                                                                    }
                                                                default:
                                                                    var v91 uint64 = 42
                                                                    _ = v91
                                                                }
                                                                _ = v80
                                                                _ = v81
                                                            } else {
                                                                var v92 string = v42
                                                                var v93 int64 = -9223372036854775808
                                                                var v94 uint8 = 128
                                                                _ = v92
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            var v95 int64 = (v2 | (v2 / ((1) | 1)))
                                                            _ = v95
                                                        }
                                                        switch fi64 {
                                                        case 128:
                                                            if ((float64(1.5) + (float64(v5) * 0.0)) <= 3.14159) {
                                                                v96 := v4[(v42 + v42)]
                                                                v2 = v2
                                                                v97 := GenS7{F0: 172}
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                        default:
                                                            v5 = v5
                                                        }
                                                        v3 += int64(v6)
                                                    }
                                                default:
                                                    if ("\n\t" != (v42 + v42)) {
                                                        v68 = (-(v68))
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v98 string = v42
                                                            _ = v98
                                                        case "hello":
                                                            if true {
                                                                fmt.Println(v1)
                                                                v1 += v1
                                                                v42 = "foo bar"
                                                                switch -1 {
                                                                case 616:
                                                                    if len(v41) > 0 {
                                                                        v41[0] = fi64
                                                                    }
                                                                case -9223372036854775808:
                                                                    if true {
                                                                        v99 := GenS8{F0: fi64, F1: -128}
                                                                        _ = v99
                                                                    } else {
                                                                        switch (-(fi63)) {
                                                                        case -683:
                                                                            var v100 int = -660
                                                                            _ = v100
                                                                        case 0:
                                                                            if len(v41) > 0 {
                                                                                v41[0] = 4294967295
                                                                            }
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                        switch -2147483648 {
                                                                        case 127:
                                                                            v5 = (v5 + float32(3.14159))
                                                                        default:
                                                                            var v101 uint32 = v6
                                                                            _ = v101
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v68 -= (-(v68))
                                                                        var v102 int8 = 0
                                                                        _ = v102
                                                                    }
                                                                case 128:
                                                                    var v103 float64 = float64(-2.25)
                                                                    _ = v103
                                                                }
                                                                switch (v42 + "foo bar") {
                                                                case "foo bar":
                                                                    v104 := GenS9{F0: 969}
                                                                    _ = v104
                                                                case "hello":
                                                                    fmt.Println(v3)
                                                                case "\n\t":
                                                                    fmt.Println(v3, v2)
                                                                }
                                                                var v105 uint32 = 361
                                                                _ = v105
                                                            } else {
                                                                fmt.Println(fi63, fi63, v3)
                                                                v6 -= 1
                                                                var v106 int16 = (v1 - v1)
                                                                _ = v106
                                                            }
                                                        case "a":
                                                            v41 = append(v41, fi63)
                                                        default:
                                                            var v107 float32 = v5
                                                            _ = v107
                                                        }
                                                        var v108 float32 = float32(0.0)
                                                        switch fi64 {
                                                        case 2147483647:
                                                            if true {
                                                                if false {
                                                                    var v109 uint32 = 2147483647
                                                                    delete(v4, "hello")
                                                                    v109 += v109
                                                                    _ = v109
                                                                }
                                                                v3 = v3
                                                                fmt.Println(v1, v3, v42)
                                                                v68 = int32(v3)
                                                                if (1.5 <= ((-2.25 - 1e10) - 0.0)) {
                                                                    if (18446744073709551615 < 291) {
                                                                        var v110 int64 = v2
                                                                        _ = v110
                                                                    } else {
                                                                        fmt.Println(fi63, v5)
                                                                    }
                                                                    switch fi63 {
                                                                    case 1:
                                                                        var v111 int = (^(fi63))
                                                                        _ = v111
                                                                    }
                                                                    v68 = (v68 | -1)
                                                                } else {
                                                                    if true {
                                                                        v3 -= (-(v2))
                                                                        if false {
                                                                            fmt.Println(v68)
                                                                            v2 = -883
                                                                            if (!true) {
                                                                                switch ("\n\t" + (v42 + ("foo bar" + v42))) {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        _ = len(v41)
                                                                                    } else {
                                                                                        if true {
                                                                                            var v112 uint8 = 255
                                                                                            switch v42 {
                                                                                            case "foo bar":
                                                                                                switch 9223372036854775807 {
                                                                                                case 482:
                                                                                                    v113 := GenS4{F0: 2, F1: v42, F2: (v2 | (^(v3)))}
                                                                                                    _ = v113
                                                                                                default:
                                                                                                    v114 := GenS10{F0: 32767}
                                                                                                    _ = v114
                                                                                                }
                                                                                            case "a":
                                                                                                var v115 uint64 = 0
                                                                                                _ = v115
                                                                                            case "hello":
                                                                                                v42 = v42
                                                                                            default:
                                                                                                if len(v41) > 0 {
                                                                                                    v41[0] = fi63
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v68, v108, v3)
                                                                                            if false {
                                                                                                delete(v4, "hello")
                                                                                            }
                                                                                            _ = v112
                                                                                        }
                                                                                        v5 -= v5
                                                                                        var v116 bool = false
                                                                                        _ = v116
                                                                                    }
                                                                                case "a":
                                                                                    switch fi64 {
                                                                                    case 32768:
                                                                                        v2 -= (v3 % ((v2) | 1))
                                                                                        _ = len(v41)
                                                                                        v42 = v42
                                                                                    case -1:
                                                                                        v68 = genFn0((v1 ^ v1))
                                                                                        fmt.Println(v1, v68, v3)
                                                                                    default:
                                                                                        v5 += v5
                                                                                    }
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        fmt.Println(v68, fi64, fi63)
                                                                                    } else {
                                                                                        var v117 int64 = (^(v2))
                                                                                        var v118 uint = 376
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                    }
                                                                                default:
                                                                                    var v119 uint32 = 2147483648
                                                                                    fmt.Println(v119, v3)
                                                                                    _ = v119
                                                                                }
                                                                            } else {
                                                                                v5 -= v108
                                                                            }
                                                                        } else {
                                                                            switch fi63 {
                                                                            case 42:
                                                                                if false {
                                                                                    if true {
                                                                                        var v120 uint8 = 0
                                                                                        v121 := GenS11{F0: (v3 | (v2 - v2)), F1: 0}
                                                                                        v68 -= v68
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                    }
                                                                                } else {
                                                                                    v2 -= v2
                                                                                }
                                                                            }
                                                                            v122 := GenS2{F0: fi64, F1: 255, F2: v6}
                                                                            v123 := GenS12{F0: 255, F1: (v3 & v3), F2: false}
                                                                            _ = v122
                                                                            _ = v123
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi63)
                                                                        var v124 int16 = 122
                                                                        var v125 float32 = float32(float64(1e10))
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                }
                                                            } else {
                                                                var v126 int32 = (-(v68))
                                                                _ = v126
                                                            }
                                                        case 42:
                                                            fmt.Println(fi64, v2, v42)
                                                        case -1:
                                                            var v127 uint8 = 53
                                                            _ = v127
                                                        }
                                                        _ = v108
                                                    }
                                                }
                                                switch "" {
                                                case "hello":
                                                    v4[(v42 + ("hello" + v42))] = (32767 <= 1)
                                                default:
                                                    var v128 float32 = (v5 + 3.14159)
                                                    _ = v128
                                                }
                                            }
                                        case "":
                                            fmt.Println(fi63, fi63, v42)
                                        default:
                                            v129 := GenS1{F0: v6, F1: 128}
                                            _ = v129
                                        }
                                        _ = len(v4)
                                        if (float64(-2.25) >= 1.5) {
                                            switch (-(fi64)) {
                                            case -1:
                                                var v130 uint = 0
                                                _ = v130
                                            default:
                                                fmt.Println(fi63)
                                            }
                                        } else {
                                            v4[((v42 + v42) + "\n\t")] = ((^(v68)) > v68)
                                            var v131 uint16 = 255
                                            var v132 int8 = 99
                                            fmt.Println(v5, v1, v3)
                                            _ = v131
                                            _ = v132
                                        }
                                    }
                                    if len(v41) >= 2 {
                                        v41 = v41[1:len(v41)]
                                    }
                                }
                                _ = v72
                            } else {
                                switch (fi64 ^ fi63) {
                                case 32768:
                                    var v133 uint64 = 296
                                    _ = v133
                                }
                            }
                            var v134 float32 = 0.0
                            _ = v134
                            _ = v68
                        }
                        v5 = v5
                    }
                default:
                    fmt.Println(fi63, v6)
                }
                if len(v41) > 0 {
                    v41[0] = (fi64 + 32767)
                }
                var v135 int16 = (v1 * (v1 - (-(v1))))
                if false {
                    v136 := GenS13{F0: -1, F1: 1}
                    if true {
                        var v137 int32 = 1
                        _ = v137
                    } else {
                        var v138 float32 = float32(1e10)
                        v139 := GenS14{F0: 1, F1: 2147483647, F2: 42}
                        switch 2147483647 {
                        case -833:
                            v4[v42] = (false && true)
                        default:
                            v135 += v1
                        }
                        _ = v138
                        _ = v139
                    }
                    var v140 uint32 = (v6 + v6)
                    _ = v136
                    _ = v140
                } else {
                    if true {
                        v135 = (v135 ^ v1)
                        v3 -= (v3 | (v2 % ((int64(v1)) | 1)))
                    } else {
                        v141 := GenS15{F0: v42, F1: v135, F2: 844}
                        _ = v141
                    }
                    switch ((v42 + "foo bar") + v42) {
                    case "":
                        switch (-(fi63)) {
                        case -216:
                            fmt.Println(v1, v3, v2)
                        default:
                            v142 := GenS16{F0: v42, F1: -97, F2: 127}
                            _ = v142
                        }
                    }
                    fmt.Println(v6, v3)
                }
                _ = v135
            }
            if true {
                if false {
                    var v143 int = int(v3)
                    _ = v143
                } else {
                    if (!true) {
                        if false {
                            switch v42 {
                            case "foo bar":
                                var v144 int16 = (v1 & v1)
                                v2 += v3
                                v5 -= float32((((((-2.25 * (-2.25 + (1e10 * (1e10 - 3.14159)))) / 1.5) / 0.1) - 1e10) * float64(v5)))
                                _ = v144
                            case "\n\t":
                                if (!(2147483647 < genFn0(-779))) {
                                    _ = len(v4)
                                    if false {
                                        var v145 bool = true
                                        _ = v145
                                    } else {
                                        v5 -= float32((1e10 + float64(v5)))
                                    }
                                }
                            case "a":
                                var v146 float32 = ((float32(3.14159) / v5) / v5)
                                _ = v146
                            }
                            _ = len(v4)
                            switch (^(fi64)) {
                            case -2147483648:
                                if (2 <= 2) {
                                    var v147 int64 = -277
                                    _ = v147
                                }
                            case 1:
                                switch (fi63 % ((fi63) | 1)) {
                                case 32767:
                                    v148 := GenS6{F0: 636}
                                    _ = v148
                                case 2147483647:
                                    v149 := GenS17{F0: fi63, F1: "\n\t"}
                                    _ = v149
                                }
                            }
                            var v150 uint = 65535
                            v3 = 861
                            v151 := GenS10{F0: 557}
                            _ = v150
                            _ = v151
                        }
                        if (true && (-2.25 < (1e10 / 1.0))) {
                            fmt.Println(fi64, v2)
                            var v152 uint64 = 0
                            _ = v152
                        } else {
                            v153 := GenS15{F0: "a", F1: v1, F2: 414}
                            _ = v153
                        }
                    } else {
                        if false {
                            fmt.Println(v3, v2)
                            v1 += 622
                            v1 -= (v1 & v1)
                        }
                        switch (fi63 - fi64) {
                        case 579:
                            var v154 float32 = v5
                            _ = v154
                        case 9223372036854775807:
                            v5 = v5
                        case 645:
                            v155 := GenS18{F0: v42, F1: v1}
                            _ = v155
                        default:
                            v5 -= v5
                        }
                        if true {
                            var v156 int = fi63
                            if (v2 <= v2) {
                                var v157 int8 = 1
                                switch fi63 {
                                case 966:
                                    v41 = append(v41, v156)
                                case 287:
                                    var v158 string = v42
                                    _ = v158
                                case 1:
                                    if false {
                                        v157 -= 1
                                    }
                                }
                                v156 = fi64
                                if (4294967295 == 925) {
                                    v156 = fi63
                                    fmt.Println(v3, fi64)
                                } else {
                                    var v159 uint = 159
                                    _ = v159
                                }
                                v160 := GenS1{F0: 89, F1: 255}
                                if (0 < 2147483647) {
                                    var v161 int64 = v3
                                    var v162 float32 = 3.14159
                                    v163 := GenS2{F0: (fi64 << ((859) & 63)), F1: 0, F2: v6}
                                    _ = v161
                                    _ = v162
                                    _ = v163
                                } else {
                                    switch v42 {
                                    case "\n\t":
                                        var v164 uint64 = 18446744073709551615
                                        _ = v164
                                    default:
                                        v165 := GenS19{F0: v6, F1: (false || true)}
                                        _ = v165
                                    }
                                    delete(v4, (((v42 + "") + "hello") + v42))
                                }
                                switch (-(v156)) {
                                case -1:
                                    v166 := v4[v42]
                                    _ = v166
                                }
                                switch 127 {
                                case -1:
                                    switch (fi63 ^ v156) {
                                    case 9223372036854775807:
                                        v41 = append(v41, (fi63 * fi64))
                                    case 0:
                                        v167 := GenS20{F0: 300, F1: (-(fi63))}
                                        _ = v167
                                    case 571:
                                        fmt.Println(v157, v6, v157)
                                    default:
                                        v6 += (v6 % ((1) | 1))
                                    }
                                default:
                                    _ = len(v41)
                                }
                                _ = v157
                                _ = v160
                            }
                            _ = v156
                        } else {
                            var v168 uint32 = v6
                            _ = v168
                        }
                    }
                    var v169 uint16 = 65535
                    var v170 int16 = -477
                    _ = v169
                    _ = v170
                }
                fmt.Println(v2, fi64)
                fmt.Println(fi64)
                v2 += (v3 * (v2 | v3))
                v6 += v6
            }
            _ = fi64
        }
        var v171 int = 0
        switch fi63 {
        case 2147483647:
            switch (v171 % ((752) | 1)) {
            case 128:
                v3 += v3
            case 4294967295:
                v172 := GenS10{F0: 2}
                _ = v172
            case 32767:
                v1 -= v1
            default:
                var v173 uint16 = 791
                _ = v173
            }
        case 32768:
            v174 := GenS5{F0: v1, F1: (^(v1))}
            _ = v174
        case -9223372036854775808:
            for fi175 := 0; fi175 < 19; fi175++ {
                var v176 float64 = float64((1e10 / (((-2.25 - ((((((((v5 / v5) / (v5 - v5)) - (v5 * (1e10 - (v5 / ((1.5 - (v5 * v5)) / (((1.5 / 1.5) * (v5 * v5)) + 3.14159)))))) + 0.0) / 1.0) + (0.0 + float32(0.1))) + float32((float64(v5) - ((1.5 * 1.5) / 1.0)))) + (v5 + v5))) * v5) / 3.14159)))
                _ = fi175
                _ = v176
            }
        default:
            switch "" {
            case "a":
                var v177 uint32 = 2
                _ = v177
            case "hello":
                if (!("hello" != v42)) {
                    var v178 float32 = float32((3.14159 / float64(1e10)))
                    for fi179 := 0; fi179 < 7; fi179++ {
                        var v180 uint8 = 42
                        switch (("a" + "foo bar") + ((v42 + ("hello" + v42)) + (v42 + "\n\t"))) {
                        case "foo bar":
                            if len(v41) > 0 {
                                v41[0] = fi63
                            }
                        case "hello":
                            if false {
                                var v181 uint32 = 1
                                v178 -= -2.25
                                if (v180 != v180) {
                                    var v182 float32 = (v178 * ((1e10 * v178) - -2.25))
                                    _ = v182
                                }
                                _ = v181
                            }
                        default:
                            var v183 int64 = v3
                            _ = v183
                        }
                        _ = fi179
                        _ = v180
                    }
                    _ = v178
                } else {
                    for fi184 := 0; fi184 < 16; fi184++ {
                        var v185 uint8 = 255
                        switch (v42 + v42) {
                        case "hello":
                            _ = len(v41)
                        case "a":
                            if true {
                                if false {
                                    if (0 <= 4294967295) {
                                        v185 = 127
                                        v185 -= v185
                                    } else {
                                        var v186 uint64 = 10
                                        var v187 float32 = 1.5
                                        _ = v186
                                        _ = v187
                                    }
                                    v185 += (v185 - 255)
                                    _ = len(v4)
                                    v188 := GenS21{F0: -86, F1: (v42 + (v42 + ((v42 + (v42 + (v42 + v42))) + v42)))}
                                    _ = v188
                                } else {
                                    switch (fi184 % ((-2147483648) | 1)) {
                                    case 4294967295:
                                        v189 := GenS14{F0: 18446744073709551615, F1: 2147483648, F2: 42}
                                        _ = v189
                                    case 1:
                                        v190 := GenS10{F0: 0}
                                        _ = v190
                                    default:
                                        if (!(207 == (v6 * v6))) {
                                            v1 = 42
                                            v191 := GenS22{F0: (v3 << ((0) & 63)), F1: (fi184 >> ((65535) & 63)), F2: uint32(v185)}
                                            switch (^(fi63)) {
                                            case 128:
                                                v5 = float32(3.14159)
                                            case -2147483648:
                                                v192 := GenS23{F0: v185}
                                                _ = v192
                                            }
                                            _ = v191
                                        }
                                    }
                                    if true {
                                        if true {
                                            v5 = (1.5 * float32(0.0))
                                            var v193 string = v42
                                            var v194 float64 = -2.25
                                            _ = v193
                                            _ = v194
                                        }
                                        v171 = v171
                                        if (int(v6) > 734) {
                                            v42 = (("a" + "a") + v42)
                                            fmt.Println(v1)
                                            v4[(v42 + "a")] = false
                                        }
                                    } else {
                                        switch v42 {
                                        case "\n\t":
                                            switch (v171 ^ 625) {
                                            case 2147483647:
                                                v6 -= 42
                                            case -929:
                                                fmt.Println(v1, v171, v5)
                                            default:
                                                v195 := GenS24{F0: 42, F1: int8(fi184)}
                                                _ = v195
                                            }
                                        case "":
                                            v196 := GenS17{F0: (fi63 >> ((2) & 63)), F1: v42}
                                            _ = v196
                                        case "hello":
                                            var v197 float32 = 1.5
                                            _ = v197
                                        default:
                                            if true {
                                                switch 464 {
                                                case -9223372036854775808:
                                                    v1 = (v1 / ((v1) | 1))
                                                case -1:
                                                    v2 = v3
                                                default:
                                                    fmt.Println(v2, v42, v5)
                                                    v1 += v1
                                                    if (!false) {
                                                        v2 -= -715
                                                    } else {
                                                        v185 += 253
                                                        if false {
                                                            v198 := GenS25{F0: ((!false) || false)}
                                                            var v199 uint32 = 1
                                                            _ = v198
                                                            _ = v199
                                                        } else {
                                                            v185 += 255
                                                            v200 := GenS12{F0: 2, F1: (v3 + v3), F2: (!false)}
                                                            _ = v200
                                                        }
                                                        v201 := GenS1{F0: 635, F1: v185}
                                                        _ = v201
                                                    }
                                                }
                                                var v202 uint = 515
                                                _ = v202
                                            }
                                        }
                                        v4[""] = (true || true)
                                        fmt.Println(fi63, v6)
                                    }
                                    if len(v41) > 0 {
                                        v41[0] = 143
                                    }
                                }
                            } else {
                                v203 := GenS26{F0: ("foo bar" + "foo bar"), F1: 470, F2: -49}
                                var v204 int32 = -2147483648
                                _ = v203
                                _ = v204
                            }
                            v1 -= (^(v1))
                            fmt.Println(v3)
                        case "":
                            v205 := GenS21{F0: (v1 - (v1 + 169)), F1: "foo bar"}
                            _ = v205
                        default:
                            v3 = v3
                        }
                        _ = fi184
                        _ = v185
                    }
                }
            case "foo bar":
                v2 += (v3 << ((42) & 63))
            }
        }
        _ = fi63
        _ = v171
    }
    var v206 int16 = (v1 / ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v206:", v206)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v206
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
