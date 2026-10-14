package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int64; F1 int8; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint32) uint {
    v7 := GenS1{F0: -1, F1: -32768, F2: (^(p0))}
    var v8 string = ((("\n\t" + "foo bar") + "a") + "hello")
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int8, p1 uint32, p2 uint64) int16 {
    switch 0 {
    case 0:
        for fi9 := 0; fi9 < 8; fi9++ {
            p2 = 42
            if true {
                var v10 float32 = float32((0.1 * (((0.0 * 1.5) / (0.1 + 1e10)) - 3.14159)))
                for fi11 := 0; fi11 < 12; fi11++ {
                    if true {
                        var v12 int8 = p0
                        var v13 float64 = 1e10
                        _ = v12
                        _ = v13
                    }
                    _ = fi11
                }
                switch "\n\t" {
                case "a":
                    if ((p2 % (((p2 % ((p2) | 1))) | 1)) == (p2 ^ 0)) {
                        p0 -= 0
                        var v14 uint32 = p1
                        var v15 uint16 = 0
                        _ = v14
                        _ = v15
                    }
                case "":
                    fmt.Println(p1, p0, p2)
                }
                _ = v10
            } else {
                p2 -= 2
            }
            _ = fi9
        }
    }
    v16 := []int{-2147483648, -2147483648, 0}
    p1 = p1
    v17 := GenS2{F0: "a", F1: 545, F2: p1}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v5 bool = (1 == 368)
    var v6 uint16 = 0
    delete(v3, 0)
    v5 = (0.1 == 1e10)
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            v5 = ((!(((("foo bar" + (((("" + ("\n\t" + "")) + "\n\t") + ("hello" + "foo bar")) + ("a" + "foo bar"))) + "foo bar") + "foo bar") != (("a" + "") + "a"))) && v5)
            _ = fi19
        }
        v20 := []int{4294967295, -2147483648, fi18, fi18}
        _ = fi18
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            switch "hello" {
            case "a":
                var v23 bool = (656 != 1)
                _ = v23
            case "foo bar":
                switch ("a" + "a") {
                case "\n\t":
                    v24 := GenS3{F0: 42, F1: (v6 >> ((65535) & 63))}
                    _ = v24
                case "":
                    switch 4294967295 {
                    case 127:
                        v5 = (1e10 <= 0.0)
                    case -732:
                        v25 := GenS3{F0: 456, F1: 514}
                        _ = v25
                    case -9223372036854775808:
                        v6 = v6
                    default:
                        var v26 bool = true
                        fmt.Println(v5)
                        var v27 string = ((("hello" + "foo bar") + "foo bar") + "a")
                        _ = v26
                        _ = v27
                    }
                case "hello":
                    var v28 uint32 = 534
                    _ = v28
                }
            case "\n\t":
                v29 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 2}
                _ = v29
            default:
                fmt.Println(fi22, fi22)
            }
            _ = fi22
        }
        var v30 bool = false
        _ = fi21
        _ = v30
    }
    var v31 int16 = 1
    if v5 {
        var v32 bool = (v31 <= genFn1(42, 882, 884))
        v1[1] = (0 < 1)
        if v32 {
            for fi33 := 0; fi33 < 7; fi33++ {
                v34 := []int{}
                for fi35 := 0; fi35 < 2; fi35++ {
                    if (v32 || (v5 || (v5 || v32))) {
                        if (v32 && (!false)) {
                            _ = len(v2)
                            switch ("" + "hello") {
                            case "hello":
                                v31 = -648
                            case "foo bar":
                                v32 = v5
                            default:
                                var v36 uint = 1
                                _ = v36
                            }
                        }
                        var v37 int = 247
                        v38 := GenS2{F0: "\n\t", F1: 654, F2: 2147483648}
                        var v39 bool = (-9223372036854775808 >= 0)
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    var v40 string = "hello"
                    _ = fi35
                    _ = v40
                }
                var v41 string = ("foo bar" + "hello")
                _ = fi33
                _ = v34
                _ = v41
            }
            for fi42 := 0; fi42 < 2; fi42++ {
                v43 := []int{(fi42 + 32767), (fi42 | (fi42 & fi42)), (fi42 + (fi42 + (fi42 % ((fi42) | 1))))}
                if v32 {
                    if true {
                        for fi44 := 0; fi44 < 15; fi44++ {
                            delete(v3, 297)
                            v43 = append(v43, (^(fi42)))
                            var v45 uint32 = 2
                            _ = fi44
                            _ = v45
                        }
                        var v46 int8 = -128
                        _ = v46
                    } else {
                        v47 := GenS1{F0: 0, F1: int16(fi42), F2: 745}
                        _ = v47
                    }
                }
                _ = fi42
                _ = v43
            }
        } else {
            if (v6 == (v6 - 32767)) {
                v48 := GenS2{F0: "\n\t", F1: 0, F2: 2147483648}
                v5 = (!(((v6 >> ((genFn0(2147483648)) & 63)) > 111) && (9223372036854775807 > 9223372036854775807)))
                _ = v48
            } else {
                v49 := []int{-608, 241}
                _ = v49
            }
            delete(v4, 9223372036854775807)
            for fi50 := 0; fi50 < 20; fi50++ {
                if (-105 <= 1) {
                    v31 += (-(v31))
                    for fi51 := 0; fi51 < 15; fi51++ {
                        v4[fi51] = (("hello" + "") + "hello")
                        _ = fi51
                    }
                    var v52 float64 = float64(float32(3.14159))
                    _ = v52
                } else {
                    var v53 float64 = -2.25
                    switch "\n\t" {
                    case "foo bar":
                        v5 = (2 != 42)
                    case "\n\t":
                        if (!v32) {
                            v31 = v31
                            var v54 float64 = ((v53 / (1e10 / ((1.5 - 1e10) - 3.14159))) * v53)
                            _ = v54
                        } else {
                            fmt.Println(v6, v5)
                        }
                    case "":
                        v55 := GenS4{F0: 255, F1: (^(fi50))}
                        _ = v55
                    }
                    _ = v53
                }
                _ = fi50
            }
        }
        _ = v32
    }
    fmt.Println(v5)
    var v56 uint = 1
    delete(v2, uint32(v56))
    switch "a" {
    case "a":
        for fi57 := 0; fi57 < 8; fi57++ {
            var v58 int16 = v31
            for fi59 := 0; fi59 < 5; fi59++ {
                if v5 {
                    var v60 uint32 = 508
                    _ = v60
                }
                if (v6 != (v6 & (v6 + v6))) {
                    v58 -= (^(v31))
                    if false {
                        if v5 {
                            var v61 int = 1
                            v31 = (v58 ^ int16(fi59))
                            _ = v61
                        } else {
                            fmt.Println(fi57, v6)
                            if v5 {
                                var v62 uint32 = uint32(v6)
                                _ = v62
                            } else {
                                var v63 uint16 = (v6 | (v6 >> ((623) & 63)))
                                if v5 {
                                    v64 := []int{fi59, (fi57 / ((fi59) | 1)), fi57, fi59}
                                    v5 = false
                                    fmt.Println(v31, v31)
                                    _ = v64
                                }
                                var v65 int = 4294967295
                                _ = v63
                                _ = v65
                            }
                            v6 += v6
                        }
                        switch (fi57 / (((-(fi57))) | 1)) {
                        case 32768:
                            if (124 > 204) {
                                fmt.Println(v56)
                            } else {
                                if v5 {
                                    v66 := GenS3{F0: 18446744073709551615, F1: (v6 ^ v6)}
                                    _ = len(v2)
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(fi59, v6, v31)
                                    }
                                    _ = v66
                                } else {
                                    switch fi59 {
                                    case -9223372036854775808:
                                        v31 = (v58 + v58)
                                    case 0:
                                        var v67 int = fi57
                                        _ = v67
                                    case -1:
                                        switch (^(fi57)) {
                                        case 9223372036854775807:
                                            if (genFn0(2147483647) != v56) {
                                                v68 := GenS2{F0: "a", F1: 4294967295, F2: 2147483647}
                                                v56 -= (v56 % (((^(v56))) | 1))
                                                _ = v68
                                            } else {
                                                fmt.Println(v58, v56)
                                                fmt.Println(v58, fi57, v58)
                                                v69 := GenS1{F0: 9223372036854775807, F1: int16(fi57), F2: 2147483647}
                                                if (!false) {
                                                    if v5 {
                                                        v5 = true
                                                        v5 = false
                                                        v70 := GenS2{F0: ((((("" + ("a" + "foo bar")) + (("a" + "\n\t") + "hello")) + "\n\t") + "hello") + "\n\t"), F1: 2, F2: 42}
                                                        v71 := GenS2{F0: ("foo bar" + (("foo bar" + "foo bar") + ("foo bar" + "foo bar"))), F1: 407, F2: 4294967295}
                                                        _ = v70
                                                        _ = v71
                                                    } else {
                                                        if v5 {
                                                            if (394 == uint32(v56)) {
                                                                fmt.Println(v31, fi57, fi59)
                                                                fmt.Println(v6, v5)
                                                            } else {
                                                                v5 = ((1 < 0) || false)
                                                            }
                                                            v72 := []int{fi59, 0, (fi59 >> ((42) & 63))}
                                                            v72 = append(v72, -2147483648)
                                                            _ = v72
                                                        } else {
                                                            switch fi57 {
                                                            case 0:
                                                                var v73 uint32 = 42
                                                                _ = v73
                                                            case -1:
                                                                if ((1e10 * float32((1.5 * (-2.25 + -2.25)))) >= 0.1) {
                                                                    fmt.Println(v6)
                                                                    switch 32768 {
                                                                    case 127:
                                                                        var v74 bool = (v58 <= 344)
                                                                        _ = v74
                                                                    case 32768:
                                                                        v75 := []int{313, (fi59 ^ fi57)}
                                                                        _ = v75
                                                                    case -9223372036854775808:
                                                                        var v76 float32 = (1e10 - 3.14159)
                                                                        _ = v76
                                                                    default:
                                                                        v6 -= v6
                                                                    }
                                                                    var v77 int64 = 9223372036854775807
                                                                    _ = v77
                                                                }
                                                            case 42:
                                                                v6 -= (v6 * (v6 % ((805) | 1)))
                                                            }
                                                            var v78 bool = ((true || (!(615 <= 47))) && ((438 < uint32(v6)) && (!(611 > (v56 >> ((genFn0(794)) & 63))))))
                                                            v79 := []int{-1, fi57, (fi59 ^ fi57)}
                                                            v58 = 344
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        if (!v5) {
                                                            v6 -= v6
                                                            v31 = -221
                                                        } else {
                                                            if (((v6 * (v6 * (v6 >> (((v56 ^ v56)) & 63)))) <= 1) || (((v6 - v6) >= (v6 | (v6 & (v6 >> ((v56) & 63))))) && false)) {
                                                                v80 := []int{}
                                                                switch (("\n\t" + "a") + "foo bar") {
                                                                case "":
                                                                    switch 2147483647 {
                                                                    case 0:
                                                                        var v81 uint = 18446744073709551615
                                                                        _ = v81
                                                                    case 9223372036854775807:
                                                                        var v82 string = ("\n\t" + "foo bar")
                                                                        _ = v82
                                                                    }
                                                                case "foo bar":
                                                                    var v83 float64 = 0.1
                                                                    _ = v83
                                                                case "hello":
                                                                    if true {
                                                                        v5 = v5
                                                                        var v84 bool = ((((0.0 - float32((0.0 / (1.5 + -2.25)))) / 1e10) * 0.1) == (1.5 * ((0.1 - 0.0) + 1.5)))
                                                                        if (v5 || v5) {
                                                                            var v85 uint8 = 42
                                                                            _ = v85
                                                                        } else {
                                                                            v84 = (fi59 <= (fi57 | fi59))
                                                                        }
                                                                        switch 128 {
                                                                        case 42:
                                                                            _ = len(v80)
                                                                        default:
                                                                            v80 = append(v80, fi57)
                                                                        }
                                                                        _ = v84
                                                                    }
                                                                default:
                                                                    if false {
                                                                        if (!v5) {
                                                                            _ = len(v80)
                                                                            fmt.Println(v56, v56, v6)
                                                                            switch (fi59 % ((fi57) | 1)) {
                                                                            case 32768:
                                                                                switch (^(fi57)) {
                                                                                case 129:
                                                                                    var v86 uint32 = 2147483647
                                                                                    _ = v86
                                                                                }
                                                                            case 9:
                                                                                fmt.Println(v56, v56, fi57)
                                                                            }
                                                                            v87 := GenS4{F0: 1, F1: fi57}
                                                                            var v88 uint64 = uint64(v56)
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        v5 = (1.5 >= 3.14159)
                                                                        fmt.Println(v58)
                                                                    }
                                                                }
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    if (2147483647 == 0) {
                                                                        fmt.Println(v6, v56)
                                                                        v89 := GenS5{F0: 42, F1: 1, F2: 0}
                                                                        _ = v89
                                                                    }
                                                                default:
                                                                    v6 += (v6 ^ (v6 | v6))
                                                                }
                                                                _ = v80
                                                            }
                                                        }
                                                        var v90 bool = v5
                                                        _ = v90
                                                    }
                                                    v91 := []int{0, (fi59 % (((fi57 + 2147483647)) | 1)), fi59, fi57}
                                                    fmt.Println(fi57, v56, v5)
                                                    _ = len(v1)
                                                    _ = v91
                                                } else {
                                                    switch (("a" + "\n\t") + "foo bar") {
                                                    case "":
                                                        var v92 bool = v5
                                                        _ = v92
                                                    case "hello":
                                                        if (("a" + "\n\t") == "") {
                                                            if (v5 || (int32(fi59) <= 42)) {
                                                                v31 += 42
                                                                v93 := GenS4{F0: 2, F1: (fi59 ^ (fi57 * 127))}
                                                                _ = v93
                                                            }
                                                        } else {
                                                            var v94 uint32 = 627
                                                            _ = v94
                                                        }
                                                    }
                                                }
                                                _ = v69
                                            }
                                        case 104:
                                            v58 = v31
                                        case 861:
                                            var v95 int32 = 42
                                            _ = v95
                                        }
                                    }
                                }
                                v96 := GenS1{F0: 1, F1: (v58 ^ (v58 >> (((v56 % ((genFn0(2147483648)) | 1))) & 63))), F2: 4294967295}
                                _ = v96
                            }
                        case 127:
                            var v97 string = (("foo bar" + "foo bar") + "\n\t")
                            _ = v97
                        case 922:
                            var v98 int16 = v58
                            _ = v98
                        default:
                            v58 = v58
                        }
                    }
                    var v99 uint32 = 624
                    _ = v99
                } else {
                    v100 := GenS2{F0: "\n\t", F1: 2147483648, F2: 2147483648}
                    v101 := GenS6{F0: 42}
                    switch (fi57 & fi57) {
                    case 479:
                        v31 += v31
                    default:
                        v102 := GenS7{F0: -128, F1: 2147483648}
                        _ = v102
                    }
                    v103 := []int{fi59, (^(fi59)), 127}
                    v31 = v58
                    _ = v100
                    _ = v101
                    _ = v103
                }
                v31 = (v58 | 32767)
                _ = fi59
            }
            for fi104 := 0; fi104 < 16; fi104++ {
                fmt.Println(v56, v6, fi104)
                fmt.Println(v31, v6, fi57)
                var v105 float32 = (1.5 - 3.14159)
                _ = fi104
                _ = v105
            }
            _ = fi57
            _ = v58
        }
    default:
        if true {
            switch "" {
            case "":
                for fi106 := 0; fi106 < 15; fi106++ {
                    fmt.Println(v56, fi106)
                    var v107 uint64 = 1
                    v108 := []int{fi106, fi106, 1}
                    _ = fi106
                    _ = v107
                    _ = v108
                }
            case "hello":
                v56 -= (v56 | v56)
            default:
                var v109 int = -2147483648
                _ = v109
            }
        } else {
            v6 -= (v6 / ((v6) | 1))
        }
    }
    var v110 string = "hello"
    v56 = v56
    switch -9223372036854775808 {
    case -815:
        var v111 uint16 = v6
        _ = v111
    case 0:
        v6 += v6
    default:
        var v112 int16 = (v31 >> ((genFn0(2147483647)) & 63))
        _ = v112
    }
    var v113 float32 = float32(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v110:", v110)
    fmt.Println("v113:", v113)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v110
    _ = v113
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
