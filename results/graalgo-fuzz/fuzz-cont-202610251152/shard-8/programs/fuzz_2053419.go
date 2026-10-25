package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS5 struct { F0 string; F1 int64; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string; F1 int16; F2 uint8 }
type GenS12 struct { F0 bool; F1 bool; F2 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 uint64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 uint) uint32 {
    v7 := []int{-957, 32767, 2147483647}
    var v8 uint = p0
    _ = p0
    _ = v7
    _ = v8
    return 65535
}
func genFn1() int8 {
    switch -671 {
    case 0:
        var v9 string = ("foo bar" + "foo bar")
        _ = v9
    case 708:
        if (42 < 100) {
            var v10 float64 = 1.5
            v11 := GenS1{F0: 0, F1: 18446744073709551615, F2: 127}
            _ = v10
            _ = v11
        } else {
            switch "" {
            case "\n\t":
                v12 := GenS1{F0: 775, F1: 2, F2: 78}
                _ = v12
            case "hello":
                var v13 float32 = (float32(float64((1.5 - -2.25))) / ((-2.25 + float32(0.0)) / (1.5 + -2.25)))
                _ = v13
            }
        }
    case -2147483648:
        var v14 int8 = -103
        _ = v14
    default:
        var v15 uint8 = 3
        _ = v15
    }
    var v16 bool = (true && true)
    for fi17 := 0; fi17 < 3; fi17++ {
        if true {
            if true {
                v18 := GenS2{F0: 32768}
                var v19 int = (-(fi17))
                _ = v18
                _ = v19
            } else {
                v20 := GenS3{F0: 2, F1: v16}
                switch "foo bar" {
                case "hello":
                    v16 = (v16 && v16)
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS3{F0: 18446744073709551615, F1: (133 <= 478)}
                    v16 = v16
                    v23 := []int{fi17, (^(fi17)), (fi21 | fi21), (-(fi17))}
                    v16 = ((1.5 + 1.5) != 1.5)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                if v16 {
                    v16 = ("foo bar" != "\n\t")
                    v25 := GenS4{F0: 1, F1: 255, F2: -128}
                    _ = v25
                }
                _ = fi24
            }
            if false {
                if (!true) {
                    if true {
                        for fi26 := 0; fi26 < 20; fi26++ {
                            v27 := GenS5{F0: "a", F1: -215, F2: 2}
                            var v28 uint = 255
                            v29 := []int{fi17, fi26, (fi26 - fi26)}
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        v16 = (("a" + "a") <= "hello")
                    }
                }
                var v30 uint16 = 61
                if v16 {
                    var v31 int16 = 1000
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v16 = (true || true)
                        v33 := GenS1{F0: -2147483648, F1: 410, F2: 2}
                        _ = fi32
                        _ = v33
                    }
                    _ = v31
                }
                _ = v30
            } else {
                v16 = v16
                var v34 int8 = 1
                var v35 int = (fi17 << ((32) & 63))
                _ = v34
                _ = v35
            }
        }
        fmt.Println(v16, v16)
        v16 = v16
        _ = fi17
    }
    v16 = v16
    var v36 int64 = -1
    _ = v16
    _ = v36
    return -48
}

func main() {
    var v1 bool = (-1 != 42)
    v2 := []int{395, 127}
    var v3 uint = 4294967295
    v4 := make(map[int]int)
    v5 := []int{2147483647, 37}
    var v6 int16 = -1
    var v37 int64 = 1
    v38 := GenS5{F0: "a", F1: v37, F2: v3}
    if v1 {
        var v39 uint32 = 786
        switch -1 {
        case 2147483647:
            v40 := GenS6{F0: 193}
            _ = v40
        case -2147483648:
            for fi41 := 0; fi41 < 9; fi41++ {
                var v42 float32 = 3.14159
                _ = fi41
                _ = v42
            }
        case 127:
            switch 331 {
            case -404:
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
            default:
                if (2 == 42) {
                    switch ("foo bar" + ("" + "hello")) {
                    case "":
                        switch 858 {
                        case 128:
                            _ = len(v4)
                        default:
                            switch ("\n\t" + "\n\t") {
                            case "":
                                v4[4294967295] = 4294967295
                            case "\n\t":
                                var v43 uint32 = v39
                                _ = v43
                            case "a":
                                v1 = (113 >= 414)
                            }
                        }
                    default:
                        for fi44 := 0; fi44 < 6; fi44++ {
                            if v1 {
                                if false {
                                    v39 += v39
                                }
                                if v1 {
                                    v37 += (v37 ^ -1)
                                    fmt.Println(v39)
                                    if (float64(((float32(1e10) * float32(0.0)) - -2.25)) > 3.14159) {
                                        switch (("\n\t" + ((((("a" + ("\n\t" + (("hello" + ("\n\t" + "foo bar")) + (("foo bar" + "foo bar") + "")))) + "") + "hello") + ("\n\t" + "a")) + "\n\t")) + "a") {
                                        case "hello":
                                            v4[int(v37)] = (fi44 | int(v6))
                                        case "":
                                            for fi45 := 0; fi45 < 14; fi45++ {
                                                var v46 uint8 = 218
                                                _ = fi45
                                                _ = v46
                                            }
                                        case "foo bar":
                                            if (v1 && false) {
                                                if (114 != 42) {
                                                    v47 := v4[(fi44 * (fi44 * (fi44 - (fi44 % (((fi44 ^ 42)) | 1)))))]
                                                    v48 := GenS3{F0: 1, F1: ((18446744073709551615 == uint64(v39)) || (!v1))}
                                                    _ = v47
                                                    _ = v48
                                                }
                                                v1 = true
                                            }
                                        default:
                                            var v49 uint = uint(v39)
                                            _ = v49
                                        }
                                        for fi50 := 0; fi50 < 8; fi50++ {
                                            var v51 uint64 = 0
                                            var v52 int16 = v6
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                } else {
                                    v53 := GenS7{F0: v37}
                                    _ = v53
                                }
                                if false {
                                    v1 = ((v1 && ((v37 - v37) > (v37 - (v37 >> ((369) & 63))))) || (("hello" + "a") <= (((("" + "hello") + "") + "hello") + ((("a" + ("" + "")) + "foo bar") + "hello"))))
                                }
                            } else {
                                v37 -= v37
                                switch ("\n\t" + "foo bar") {
                                case "a":
                                    switch "" {
                                    case "\n\t":
                                        v4[(fi44 / ((fi44) | 1))] = fi44
                                    case "a":
                                        v3 += uint(v39)
                                    case "":
                                        _ = len(v5)
                                    }
                                default:
                                    if false {
                                        fmt.Println(fi44, v39)
                                        switch "a" {
                                        case "a":
                                            switch 4294967295 {
                                            case -2147483648:
                                                v39 = 42
                                            case 1:
                                                if (255 < 255) {
                                                    var v54 uint64 = 1
                                                    switch (fi44 & (fi44 / ((-1) | 1))) {
                                                    case -9223372036854775808:
                                                        if true {
                                                            if ((1.5 / (0.1 / 3.14159)) < 0.0) {
                                                                if (!(!(true && true))) {
                                                                    v55 := GenS2{F0: fi44}
                                                                    _ = v55
                                                                } else {
                                                                    var v56 uint32 = (v39 / ((1) | 1))
                                                                    v57 := GenS2{F0: (-(fi44))}
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                            }
                                                            var v58 int64 = (v37 / ((-335) | 1))
                                                            _ = v58
                                                        }
                                                    }
                                                    _ = v54
                                                } else {
                                                    for fi59 := 0; fi59 < 10; fi59++ {
                                                        v3 = v3
                                                        v60 := GenS8{F0: fi44, F1: -207}
                                                        _ = fi59
                                                        _ = v60
                                                    }
                                                    for fi61 := 0; fi61 < 19; fi61++ {
                                                        if false {
                                                            if v1 {
                                                                var v62 uint8 = 87
                                                                _ = v62
                                                            } else {
                                                                if len(v5) > 0 {
                                                                    v5[0] = 564
                                                                }
                                                                switch (-(fi44)) {
                                                                case -26:
                                                                    fmt.Println(v1, v1)
                                                                default:
                                                                    v63 := GenS9{F0: v37, F1: 1, F2: 255}
                                                                    _ = v63
                                                                }
                                                                v64 := GenS10{F0: v1}
                                                                _ = v64
                                                            }
                                                            v65 := GenS1{F0: 42, F1: 0, F2: 128}
                                                            _ = v65
                                                        } else {
                                                            var v66 uint32 = genFn0((v3 / ((v3) | 1)))
                                                            var v67 uint16 = 32767
                                                            var v68 float32 = 3.14159
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        v69 := GenS11{F0: "", F1: (v6 >> (((v3 / ((v3) | 1))) & 63)), F2: 2}
                                                        _ = fi61
                                                        _ = v69
                                                    }
                                                    var v70 int32 = -819
                                                    _ = v70
                                                }
                                            case 2147483647:
                                                var v71 int16 = v6
                                                _ = v71
                                            default:
                                                var v72 float32 = 1.5
                                                _ = v72
                                            }
                                        case "":
                                            v1 = (0.1 > 1.5)
                                        default:
                                            fmt.Println(fi44, v3)
                                        }
                                        var v73 int = fi44
                                        _ = v73
                                    } else {
                                        if v1 {
                                            v74 := GenS8{F0: 127, F1: -2147483648}
                                            _ = v74
                                        }
                                    }
                                }
                            }
                            _ = fi44
                        }
                    }
                    for fi75 := 0; fi75 < 10; fi75++ {
                        var v76 int16 = v6
                        var v77 int64 = (-(v37))
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                } else {
                    switch -2147483648 {
                    case -702:
                        if (!true) {
                            var v78 bool = v1
                            if (65535 <= v39) {
                                fmt.Println(v1, v3)
                                var v79 uint64 = 1
                                var v80 uint16 = 862
                                _ = v79
                                _ = v80
                            }
                            _ = v78
                        }
                    case -843:
                        var v81 bool = false
                        _ = v81
                    case -220:
                        if true {
                            if len(v5) > 0 {
                                v5[0] = -2147483648
                            }
                            v82 := GenS4{F0: 29, F1: 127, F2: 127}
                            _ = v82
                        }
                    default:
                        var v83 int8 = 1
                        _ = v83
                    }
                }
            }
        default:
            if (((0 == 42) && (int16(v39) <= v6)) && v1) {
                switch "" {
                case "a":
                    v84 := GenS12{F0: (v6 != 748), F1: false, F2: 42}
                    _ = v84
                case "foo bar":
                    v85 := GenS2{F0: 128}
                    _ = v85
                case "":
                    var v86 uint = (v3 & v3)
                    _ = v86
                default:
                    v37 -= (v37 % (((v37 | (v37 / ((int64(v6)) | 1)))) | 1))
                }
                fmt.Println(v39)
            }
        }
        fmt.Println(v6)
        _ = v39
    } else {
        for fi87 := 0; fi87 < 7; fi87++ {
            v37 -= (v37 ^ (v37 + (v37 << ((v3) & 63))))
            v88 := GenS13{F0: 2}
            for fi89 := 0; fi89 < 15; fi89++ {
                v1 = v1
                if len(v5) > 0 {
                    v5[0] = fi87
                }
                _ = fi89
            }
            _ = fi87
            _ = v88
        }
        fmt.Println(v6, v6)
        var v90 int64 = v37
        _ = v90
    }
    v5 = append(v5, 42)
    if true {
        v2 = append(v2, 2147483647)
        for fi91 := 0; fi91 < 12; fi91++ {
            v6 += 425
            switch "foo bar" {
            case "\n\t":
                v92 := GenS14{F0: (fi91 & (fi91 & 1)), F1: 18446744073709551615}
                _ = v92
            case "foo bar":
                switch (fi91 + (fi91 - 128)) {
                case -1:
                    var v93 string = "a"
                    _ = v93
                case 558:
                    switch ("foo bar" + ((("a" + "foo bar") + ("\n\t" + "hello")) + "")) {
                    case "hello":
                        var v94 string = "\n\t"
                        _ = v94
                    }
                default:
                    v3 += 18446744073709551615
                }
            case "hello":
                fmt.Println(fi91, v3)
            }
            _ = fi91
        }
        fmt.Println(v6, v3)
    } else {
        if false {
            switch 127 {
            case 32768:
                for fi95 := 0; fi95 < 6; fi95++ {
                    if (!false) {
                        var v96 uint16 = 1
                        _ = v96
                    }
                    _ = fi95
                }
            case 42:
                for fi97 := 0; fi97 < 9; fi97++ {
                    v37 = (^(v37))
                    for fi98 := 0; fi98 < 6; fi98++ {
                        v99 := GenS14{F0: (fi97 << (((v3 + (v3 >> ((v3) & 63)))) & 63)), F1: 4294967295}
                        v3 -= v3
                        v100 := v4[(fi98 + (fi97 | fi97))]
                        _ = fi98
                        _ = v100
                        _ = v99
                    }
                    _ = fi97
                }
            default:
                var v101 int16 = v6
                _ = v101
            }
            for fi102 := 0; fi102 < 7; fi102++ {
                var v103 int32 = -479
                if (1e10 < (0.1 + -2.25)) {
                    if ("foo bar" < ("hello" + ("\n\t" + ((("" + "a") + ("a" + "\n\t")) + ((("" + "a") + "foo bar") + "foo bar"))))) {
                        if false {
                            v104 := GenS15{F0: v6}
                            v1 = ((fi102 + fi102) == fi102)
                            for fi105 := 0; fi105 < 9; fi105++ {
                                v106 := GenS9{F0: v37, F1: 0, F2: 255}
                                _ = fi105
                                _ = v106
                            }
                            _ = v104
                        } else {
                            fmt.Println(fi102)
                            var v107 int32 = v103
                            _ = v107
                        }
                    }
                    var v108 float64 = 1e10
                    switch (fi102 / (((^(fi102))) | 1)) {
                    case 128:
                        for fi109 := 0; fi109 < 10; fi109++ {
                            switch (fi109 >> ((v3) & 63)) {
                            case 136:
                                v3 = 893
                            case 0:
                                var v110 float32 = 0.1
                                _ = v110
                            default:
                                v6 = (v6 % (((v6 << ((116) & 63))) | 1))
                            }
                            v6 += (^(v6))
                            switch ("\n\t" + "foo bar") {
                            case "":
                                if v1 {
                                    var v111 int32 = v103
                                    _ = v111
                                }
                            case "hello":
                                fmt.Println(v6, v37, v103)
                            case "foo bar":
                                fmt.Println(v103, fi102, v108)
                            default:
                                v112 := GenS16{F0: (v37 / ((v37) | 1)), F1: 18446744073709551615, F2: 42}
                                _ = v112
                            }
                            _ = fi109
                        }
                    case 32768:
                        v108 -= 3.14159
                    case 516:
                        if true {
                            v108 += (v108 / v108)
                        }
                    }
                    _ = v108
                }
                _ = fi102
                _ = v103
            }
        }
    }
    var v113 uint8 = 1
    for fi114 := 0; fi114 < 7; fi114++ {
        if v1 {
            v115 := GenS7{F0: (v37 % (((v37 << (((v3 * 4294967295)) & 63))) | 1))}
            _ = v115
        }
        fmt.Println(fi114)
        var v116 int = fi114
        _ = fi114
        _ = v116
    }
    v113 -= (v113 << (((v3 / (((^(v3))) | 1))) & 63))
    v4[2147483647] = 0
    if true {
        var v117 int32 = 42
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
