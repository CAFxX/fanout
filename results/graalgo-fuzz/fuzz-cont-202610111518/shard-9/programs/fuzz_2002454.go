package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint; F2 int8 }
type GenS5 struct { F0 bool; F1 int64; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int; F2 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 uint64; F2 string }
func genFn0(p0 uint, p1 int) uint32 {
    p0 += p0
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 1, F1: p0}
        if false {
            switch (("a" + (("hello" + "hello") + "")) + "hello") {
            case "foo bar":
                var v10 int64 = 42
                _ = v10
            case "":
                if false {
                    p0 = 0
                    switch (((("hello" + "hello") + "") + ("\n\t" + ("a" + "\n\t"))) + "hello") {
                    case "foo bar":
                        var v11 uint = 18446744073709551615
                        _ = v11
                    default:
                        if false {
                            v12 := []int{582, 4294967295, p1}
                            v7 += p1
                            _ = v12
                        }
                    }
                    fmt.Println(p0, fi8)
                } else {
                    p1 = int(p0)
                }
            case "\n\t":
                v13 := GenS2{F0: int32(p0), F1: -2147483648, F2: 0}
                _ = v13
            }
            switch -660 {
            case -9223372036854775808:
                var v14 uint16 = 213
                _ = v14
            case 473:
                v15 := []int{(p1 ^ fi8), 2147483647, (-(p1))}
                _ = v15
            case 128:
                p0 -= (p0 ^ 0)
            }
            switch v7 {
            case 4294967295:
                switch "\n\t" {
                case "foo bar":
                    v16 := GenS1{F0: p0, F1: p0}
                    _ = v16
                case "a":
                    if false {
                        p1 = (p1 + -515)
                        if (!true) {
                            switch p1 {
                            case 9223372036854775807:
                                fmt.Println(p0, v7, p0)
                            case -305:
                                var v17 string = "foo bar"
                                _ = v17
                            }
                            var v18 float64 = 3.14159
                            p0 += (^(p0))
                            _ = v18
                        } else {
                            fmt.Println(fi8, p1, v7)
                            fmt.Println(p0, v7, p1)
                            var v19 int16 = -998
                            _ = v19
                        }
                    } else {
                        switch (("\n\t" + "\n\t") + "a") {
                        case "foo bar":
                            switch (p1 % (((v7 ^ v7)) | 1)) {
                            case 127:
                                v20 := GenS2{F0: 849, F1: -2147483648, F2: -267}
                                _ = v20
                            default:
                                v21 := []int{p1, 127, fi8}
                                _ = v21
                            }
                        case "":
                            var v22 int64 = -1
                            _ = v22
                        case "\n\t":
                            var v23 float64 = 3.14159
                            _ = v23
                        }
                        var v24 uint64 = 360
                        var v25 int16 = -1
                        _ = v24
                        _ = v25
                    }
                }
            case -136:
                if ("\n\t" != "foo bar") {
                    p1 -= (v7 + (v7 + 32767))
                    p0 -= (p0 + 65535)
                    v26 := GenS2{F0: 0, F1: 2147483647, F2: 1}
                    _ = v26
                } else {
                    fmt.Println(p0)
                }
            default:
                var v27 int = (fi8 + p1)
                _ = v27
            }
        } else {
            p0 = p0
            var v28 int = (fi8 - fi8)
            if ((((p1 >= (fi8 ^ 0)) && ((0.0 + (float64(-2.25) / (3.14159 * float64(((float32(1.5) / float32(float64(1e10))) / float32(1e10)))))) == 1e10)) && true) && (!(false && false))) {
                v29 := []int{}
                _ = v29
            } else {
                p1 += (p1 >> ((p0) & 63))
            }
            _ = v28
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := GenS3{F0: 1, F1: -128, F2: 655}
            v32 := GenS3{F0: 4294967295, F1: 10, F2: 988}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1(p0 int16, p1 uint32, p2 uint8) uint32 {
    var v33 int32 = 0
    p0 = -1
    p1 = p1
    p0 -= 42
    v34 := GenS1{F0: 18446744073709551615, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v34
    return (p1 / ((p1) | 1))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 float64 = -2.25
    var v5 int16 = 801
    var v6 uint64 = 18446744073709551615
    var v35 uint32 = 42
    v36 := GenS3{F0: 318, F1: 57, F2: 42}
    if (1 != 2) {
        fmt.Println(v4, v35, v35)
        v37 := GenS2{F0: -1, F1: -1, F2: 1}
        _ = v37
    }
    var v38 bool = ((^(v5)) == (v5 * (^(v5))))
    fmt.Println(v4, v4, v35)
    if v38 {
        v39 := []int{-2147483648}
        v40 := GenS1{F0: 440, F1: 2}
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        v42 := []int{(fi41 & (-(fi41))), (fi41 % (((fi41 << ((1) & 63))) | 1))}
        _ = fi41
        _ = v42
    }
    if (!v38) {
        v4 += v4
        v5 -= v5
    }
    v38 = (v38 && (1 == (v5 - 615)))
    switch "hello" {
    case "a":
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 float64 = (v4 - 0.1)
                _ = fi44
                _ = v45
            }
            var v46 uint64 = v6
            _ = fi43
            _ = v46
        }
    }
    if (float64((((0.1 - 0.1) * (((1.5 / 1.5) / 1.0) * 0.0)) / 3.14159)) > v4) {
        if true {
            switch ("foo bar" + ("foo bar" + ("" + "foo bar"))) {
            case "foo bar":
                var v47 int = 392
                _ = v47
            case "\n\t":
                var v48 uint16 = 42
                _ = v48
            case "":
                if (1e10 == float32((1.5 - v4))) {
                    _ = len(v1)
                    v35 = v35
                    for fi49 := 0; fi49 < 16; fi49++ {
                        if false {
                            if v38 {
                                v4 += v4
                                var v50 bool = (4294967295 > 19)
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 15; fi51++ {
                                _ = len(v1)
                                _ = fi51
                            }
                            var v52 int32 = -1
                            _ = v52
                        } else {
                            switch ("hello" + "") {
                            case "hello":
                                for fi53 := 0; fi53 < 20; fi53++ {
                                    var v54 uint8 = 0
                                    switch "hello" {
                                    case "hello":
                                        v55 := []int{64, fi49}
                                        _ = v55
                                    case "":
                                        v4 = ((v4 * 1e10) * float64((-2.25 - 1.5)))
                                    case "a":
                                        if (1 > 127) {
                                            var v56 float32 = 1.5
                                            v57 := GenS2{F0: 1, F1: -2147483648, F2: -1}
                                            _ = v56
                                            _ = v57
                                        } else {
                                            if v38 {
                                                v58 := GenS1{F0: 1, F1: 42}
                                                _ = v58
                                            }
                                        }
                                    default:
                                        delete(v3, "hello")
                                    }
                                    _ = fi53
                                    _ = v54
                                }
                            case "foo bar":
                                for fi59 := 0; fi59 < 20; fi59++ {
                                    if v38 {
                                        v35 -= (v35 / (((v35 - genFn1((v5 - v5), 2, 42))) | 1))
                                    } else {
                                        if ((^(v5)) < (v5 << ((1) & 63))) {
                                            var v60 int8 = -110
                                            _ = v60
                                        } else {
                                            v61 := GenS4{F0: ((v6 << ((65535) & 63)) > (v6 + 18446744073709551615)), F1: 0, F2: -1}
                                            if true {
                                                v62 := []int{}
                                                if ("" <= "\n\t") {
                                                    switch ((("hello" + ("hello" + (("" + "foo bar") + "hello"))) + "a") + (((("a" + "foo bar") + ("hello" + "")) + (("\n\t" + "") + "foo bar")) + "hello")) {
                                                    case "a":
                                                        var v63 int16 = -413
                                                        _ = v63
                                                    case "":
                                                        v3[("foo bar" + "hello")] = "a"
                                                    }
                                                    v62 = append(v62, (fi49 + 380))
                                                    switch fi59 {
                                                    case 308:
                                                        var v64 int = (-(fi59))
                                                        _ = v64
                                                    case -9223372036854775808:
                                                        fmt.Println(v35)
                                                    case -1:
                                                        v65 := GenS4{F0: false, F1: uint(v35), F2: 42}
                                                        _ = v65
                                                    default:
                                                        fmt.Println(v5, v6)
                                                    }
                                                } else {
                                                    if v38 {
                                                        _ = len(v62)
                                                        v4 -= (1.5 * (0.1 - (v4 + v4)))
                                                        if v38 {
                                                            if true {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    fmt.Println(v35, v6, fi59)
                                                                default:
                                                                    v6 = v6
                                                                }
                                                            } else {
                                                                switch (fi59 - (fi49 & fi49)) {
                                                                case 128:
                                                                    if len(v62) >= 2 {
                                                                        v62 = v62[1:len(v62)]
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch fi59 {
                                                                    case 127:
                                                                        var v66 uint8 = 127
                                                                        _ = v66
                                                                    case 2147483647:
                                                                        fmt.Println(v6)
                                                                    }
                                                                default:
                                                                    fmt.Println(fi59)
                                                                }
                                                                v67 := GenS4{F0: v38, F1: 255, F2: -128}
                                                                _ = v67
                                                            }
                                                            v38 = (!(0 == 37))
                                                            v35 -= (v35 ^ v35)
                                                            var v68 int32 = 1
                                                            _ = v68
                                                        } else {
                                                            v35 = v35
                                                            fmt.Println(v35, v38, v6)
                                                            v69 := GenS5{F0: v38, F1: 544, F2: 32767}
                                                            _ = v69
                                                        }
                                                    }
                                                    if len(v62) > 0 {
                                                        v62[0] = -9223372036854775808
                                                    }
                                                    v6 += (v6 & v6)
                                                }
                                                _ = v62
                                            }
                                            fmt.Println(v5)
                                            _ = v61
                                        }
                                        fmt.Println(v35, v4, fi49)
                                        v5 = (v5 & 561)
                                    }
                                    v70 := []int{}
                                    var v71 int32 = 42
                                    v72 := GenS4{F0: (v38 && (v71 != v71)), F1: 0, F2: 100}
                                    v71 = (v71 ^ (v71 | (v71 * -2147483648)))
                                    _ = fi59
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                            }
                        }
                        fmt.Println(v4, v4)
                        _ = fi49
                    }
                } else {
                    if true {
                        switch "a" {
                        case "\n\t":
                            for fi73 := 0; fi73 < 3; fi73++ {
                                for fi74 := 0; fi74 < 10; fi74++ {
                                    v75 := GenS3{F0: 0, F1: 45, F2: 595}
                                    v76 := GenS2{F0: -2147483648, F1: -2147483648, F2: -1}
                                    var v77 uint64 = (^(v6))
                                    var v78 int16 = (^(v5))
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                }
                                v79 := []int{2147483647}
                                _ = fi73
                                _ = v79
                            }
                        case "hello":
                            v80 := GenS6{F0: ((42 >= 65535) && (65535 >= 932))}
                            _ = v80
                        default:
                            v81 := GenS7{F0: 4294967295, F1: 9223372036854775807, F2: 18446744073709551615}
                            _ = v81
                        }
                    } else {
                        v82 := GenS8{F0: 42}
                        v83 := GenS9{F0: (v35 ^ v35)}
                        _ = v82
                        _ = v83
                    }
                }
            }
            switch -2147483648 {
            case 9223372036854775807:
                switch "a" {
                case "foo bar":
                    var v84 uint64 = (^(v6))
                    _ = v84
                case "a":
                    v5 -= v5
                }
            case 0:
                for fi85 := 0; fi85 < 7; fi85++ {
                    delete(v2, v35)
                    _ = fi85
                }
            }
            for fi86 := 0; fi86 < 17; fi86++ {
                if (!(18446744073709551615 < v6)) {
                    var v87 uint = 18446744073709551615
                    _ = v87
                } else {
                    if v38 {
                        var v88 int32 = -398
                        var v89 int32 = (^(v88))
                        var v90 uint64 = 18446744073709551615
                        _ = v88
                        _ = v89
                        _ = v90
                    } else {
                        v91 := GenS6{F0: v38}
                        v4 = (3.14159 + v4)
                        switch 743 {
                        case 4294967295:
                            v38 = false
                        case 307:
                            var v92 int32 = 1
                            _ = v92
                        }
                        _ = v91
                    }
                    var v93 int = fi86
                    _ = v93
                }
                switch (fi86 + fi86) {
                case 2147483647:
                    var v94 float64 = v4
                    _ = v94
                }
                fmt.Println(v5, v5, v4)
                _ = fi86
            }
        }
        fmt.Println(v5, v35, v5)
        if (836 <= 520) {
            v35 = v35
            for fi95 := 0; fi95 < 6; fi95++ {
                v96 := GenS10{F0: v5, F1: 258, F2: ("hello" + ("\n\t" + ("a" + ("" + ("" + "")))))}
                if v38 {
                    for fi97 := 0; fi97 < 20; fi97++ {
                        v98 := GenS2{F0: 42, F1: 0, F2: -1}
                        v99 := GenS4{F0: true, F1: 974, F2: -109}
                        _ = fi97
                        _ = v98
                        _ = v99
                    }
                    v100 := []int{}
                    _ = v100
                }
                if (!("foo bar" < "")) {
                    v38 = v38
                    v5 += v5
                    if (50 <= 1) {
                        v101 := []int{(fi95 >> ((65535) & 63)), fi95, fi95, (fi95 - fi95)}
                        var v102 int16 = (^(v5))
                        for fi103 := 0; fi103 < 13; fi103++ {
                            v5 -= v5
                            _ = fi103
                        }
                        _ = v101
                        _ = v102
                    } else {
                        for fi104 := 0; fi104 < 13; fi104++ {
                            _ = len(v2)
                            _ = fi104
                        }
                        for fi105 := 0; fi105 < 12; fi105++ {
                            v35 = genFn1((v5 << ((255) & 63)), 65535, 42)
                            if v38 {
                                var v106 int = fi105
                                v2[42] = (v35 << ((18446744073709551615) & 63))
                                v35 = 1
                                _ = v106
                            }
                            v5 = v5
                            _ = fi105
                        }
                        v6 = 18446744073709551615
                    }
                } else {
                    for fi107 := 0; fi107 < 7; fi107++ {
                        if v38 {
                            if (3.14159 <= (-2.25 * -2.25)) {
                                if v38 {
                                    v108 := []int{}
                                    _ = v108
                                }
                                v5 += (v5 / ((-204) | 1))
                                switch fi107 {
                                case 32768:
                                    var v109 uint64 = (v6 - v6)
                                    _ = v109
                                case 4294967295:
                                    v35 += 1
                                }
                                v5 += v5
                                fmt.Println(fi107, fi107)
                            } else {
                                fmt.Println(v6)
                            }
                            v5 = (v5 - 42)
                            v5 -= (v5 ^ -410)
                        }
                        _ = fi107
                    }
                    var v110 int16 = (^(v5))
                    fmt.Println(v38, v35, v6)
                    _ = v110
                }
                _ = fi95
                _ = v96
            }
        }
    } else {
        var v111 int64 = -9223372036854775808
        v112 := v1[-295]
        v113 := []int{146, -1, 128, 127}
        _ = v111
        _ = v112
        _ = v113
    }
    v114 := v2[4294967295]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
