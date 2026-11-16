package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint }
type GenS7 struct { F0 bool; F1 string; F2 uint32 }
type GenS8 struct { F0 uint }
func genFn0(p0 uint16, p1 int) uint32 {
    p0 += (^(p0))
    switch (p1 % (((-(p1))) | 1)) {
    case 32768:
        p1 = -112
    }
    if (399 > -1) {
        v5 := GenS1{F0: 974, F1: p0}
        v6 := []int{p1}
        _ = v5
        _ = v6
    } else {
        switch -1 {
        case 32767:
            v7 := GenS1{F0: 1, F1: p0}
            _ = v7
        case 9223372036854775807:
            if false {
                if false {
                    v8 := GenS2{F0: 18446744073709551615, F1: (^(p0))}
                    _ = v8
                } else {
                    var v9 uint64 = 685
                    _ = v9
                }
                if false {
                    v10 := GenS2{F0: 408, F1: p0}
                    var v11 float32 = 0.1
                    _ = v10
                    _ = v11
                }
                var v12 uint32 = 42
                _ = v12
            }
        case -1:
            switch ("hello" + (("foo bar" + "foo bar") + "a")) {
            case "\n\t":
                if (42 <= -32768) {
                    var v13 int8 = 127
                    _ = v13
                }
            }
        }
        var v14 int16 = -32768
        _ = v14
    }
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1(p0 uint32) uint64 {
    p0 = genFn0(42, 127)
    fmt.Println(p0)
    switch 1 {
    case 221:
        fmt.Println(p0)
    }
    p0 = (p0 * (^(p0)))
    var v15 bool = ((!false) && true)
    _ = p0
    _ = v15
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 uint = 86
    v4 := make(map[uint32]bool)
    var v16 uint32 = genFn0(255, -451)
    v3 += v3
    v3 = v3
    if false {
        _ = len(v2)
    }
    v16 -= genFn0(255, -849)
    v17 := v2[1]
    v16 = (v17 % (((v16 - 2147483648)) | 1))
    if true {
        fmt.Println(v17, v16, v17)
    } else {
        v18 := GenS3{F0: 119}
        v17 = (v17 + 84)
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v3 -= v3
        _ = fi19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS4{F0: 31}
        fmt.Println(v16, v3)
        _ = fi20
        _ = v21
    }
    v3 += v3
    v2[480] = v16
    if false {
        if true {
            var v22 int = int(v3)
            switch 9223372036854775807 {
            case -1:
                if false {
                    var v23 int32 = -1
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if (-1 == -371) {
                            v22 = int(v17)
                            v25 := GenS1{F0: uint(v17), F1: 255}
                            _ = v25
                        }
                        _ = fi24
                    }
                }
            case 128:
                v26 := []int{(v22 + 2147483647), v22, v22, (v22 - (v22 & v22))}
                _ = v26
            default:
                v27 := []int{-217, (v22 + (v22 / (((v22 | (v22 % (((v22 | v22)) | 1)))) | 1))), v22, v22}
                _ = v27
            }
            _ = v22
        }
        v28 := GenS1{F0: 872, F1: 729}
        if (((!false) || false) && (((0.0 + -2.25) > 1.5) || ((0.1 + (((((float32((0.0 + 1.5)) * float32(1.5)) + 1e10) + 1.5) + 0.0) * ((1e10 * 0.0) * -2.25))) >= ((0.1 + 0.0) + (1e10 + 3.14159))))) {
            if false {
                v29 := GenS5{F0: 0}
                v30 := []int{-1, -765, 32767, 0}
                _ = v29
                _ = v30
            } else {
                v31 := []int{249, 0, 9223372036854775807}
                fmt.Println(v17, v17, v17)
                v17 += 4294967295
                _ = v31
            }
            if true {
                if (("hello" + "\n\t") >= (("a" + "\n\t") + (("foo bar" + "foo bar") + "hello"))) {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 uint8 = 2
                        _ = fi32
                        _ = v33
                    }
                    var v34 float32 = float32(3.14159)
                    fmt.Println(v34, v16)
                    _ = v34
                }
                v1[32768] = (v17 ^ 2147483648)
            }
        } else {
            if ((0.1 == 1.5) || (42 >= 1)) {
                v17 += genFn0(739, 1)
                var v35 int8 = 127
                _ = v35
            } else {
                fmt.Println(v17)
                var v36 int16 = -35
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 int32 = -2147483648
                    var v39 uint = 18446744073709551615
                    v40 := GenS4{F0: 18446744073709551615}
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = v36
            }
        }
        _ = v28
    }
    if (true || true) {
        if true {
            for fi41 := 0; fi41 < 4; fi41++ {
                v42 := []int{}
                for fi43 := 0; fi43 < 13; fi43++ {
                    v3 = (v3 ^ 65535)
                    _ = fi43
                }
                _ = fi41
                _ = v42
            }
            switch -2147483648 {
            case 4294967295:
                v17 = 2147483648
            default:
                switch "\n\t" {
                case "foo bar":
                    for fi44 := 0; fi44 < 17; fi44++ {
                        v17 = (v16 % ((v16) | 1))
                        switch "a" {
                        case "a":
                            if ((^(v16)) == (v17 * v17)) {
                                switch -561 {
                                case -1:
                                    for fi45 := 0; fi45 < 9; fi45++ {
                                        v46 := GenS2{F0: v3, F1: 570}
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                                switch (fi44 | (fi44 & (fi44 * (fi44 + (fi44 >> (((^(v3))) & 63)))))) {
                                case -1:
                                    if (v17 != v16) {
                                        if ("hello" <= ("a" + "\n\t")) {
                                            v16 -= v16
                                            for fi47 := 0; fi47 < 20; fi47++ {
                                                if (!(86 < fi47)) {
                                                    var v48 float64 = 3.14159
                                                    switch "" {
                                                    case "":
                                                        v48 = v48
                                                    default:
                                                        v16 += (^(v17))
                                                    }
                                                    _ = v48
                                                }
                                                delete(v1, fi44)
                                                _ = fi47
                                            }
                                        }
                                        var v49 int32 = -2147483648
                                        _ = v49
                                    }
                                case -193:
                                    if ((!false) || (((2 <= 2) && (!(false && false))) || false)) {
                                        if (89 >= 32767) {
                                            fmt.Println(v16, v3)
                                        } else {
                                            for fi50 := 0; fi50 < 16; fi50++ {
                                                if false {
                                                    if true {
                                                        if (true && true) {
                                                            var v51 string = ("hello" + ("\n\t" + "\n\t"))
                                                            _ = v51
                                                        } else {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v52 int8 = 100
                                                                _ = v52
                                                            case "foo bar":
                                                                v2[fi44] = (v16 << ((v3) & 63))
                                                            case "\n\t":
                                                                v53 := GenS5{F0: 23}
                                                                _ = v53
                                                            }
                                                            fmt.Println(fi44)
                                                            if (true && (-1 >= -652)) {
                                                                v17 = v17
                                                                var v54 int16 = 0
                                                                _ = v54
                                                            }
                                                        }
                                                        v17 += v17
                                                    }
                                                }
                                                if false {
                                                    v3 += 255
                                                }
                                                switch 42 {
                                                case 324:
                                                    v16 = (v16 | v17)
                                                case 42:
                                                    v55 := GenS2{F0: v3, F1: 42}
                                                    _ = v55
                                                default:
                                                    if false {
                                                        if true {
                                                            v3 = v3
                                                        }
                                                    } else {
                                                        switch ("hello" + "hello") {
                                                        case "":
                                                            var v56 float32 = (((((1e10 * (0.0 * 1.5)) * 1e10) - 3.14159) * 1.5) * 0.1)
                                                            _ = v56
                                                        default:
                                                            if false {
                                                                if false {
                                                                    delete(v4, 2)
                                                                } else {
                                                                    v17 = v17
                                                                    v57 := GenS5{F0: 40}
                                                                    _ = v57
                                                                }
                                                                v3 = (^(v3))
                                                                v58 := []int{}
                                                                _ = v58
                                                            }
                                                        }
                                                        v59 := []int{9223372036854775807}
                                                        _ = v59
                                                    }
                                                }
                                                var v60 uint64 = 2
                                                var v61 float32 = 1e10
                                                if ((!(786 > (fi50 * 42))) && (0 < 189)) {
                                                    var v62 int64 = 0
                                                    _ = v62
                                                } else {
                                                    fmt.Println(v16, v17, v61)
                                                    if false {
                                                        v16 += genFn0(255, fi50)
                                                        var v63 int8 = 127
                                                        _ = v63
                                                    } else {
                                                        if true {
                                                            var v64 string = "hello"
                                                            _ = v64
                                                        } else {
                                                            var v65 uint32 = v17
                                                            if (!(!(v3 <= (v3 / ((502) | 1))))) {
                                                                switch fi50 {
                                                                case -983:
                                                                    var v66 int16 = -1
                                                                    _ = v66
                                                                case 127:
                                                                    v67 := []int{-922, fi50, fi44, (^(fi44))}
                                                                    switch -486 {
                                                                    case 0:
                                                                        v67 = append(v67, (fi50 - -805))
                                                                    case -182:
                                                                        v65 = 279
                                                                    }
                                                                    _ = v67
                                                                }
                                                            } else {
                                                                var v68 string = "a"
                                                                var v69 string = (("" + v68) + "\n\t")
                                                                var v70 uint64 = uint64(v17)
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            _ = v65
                                                        }
                                                        if true {
                                                            var v71 int16 = -1
                                                            v3 = (^(v3))
                                                            v60 = 18446744073709551615
                                                            _ = v71
                                                        } else {
                                                            if ((v60 >= 0) && true) {
                                                                v72 := []int{(fi50 ^ fi50)}
                                                                v61 += 0.0
                                                                _ = v72
                                                            }
                                                            switch (fi50 % (((fi50 % ((fi44) | 1))) | 1)) {
                                                            case 990:
                                                                if false {
                                                                    var v73 uint = v3
                                                                    _ = len(v2)
                                                                    v61 = v61
                                                                    v16 -= (v17 - 109)
                                                                    if (("foo bar" + "foo bar") < ("a" + (("hello" + "a") + (("" + ("\n\t" + ((("a" + "a") + "hello") + ("a" + ("\n\t" + "foo bar"))))) + (("hello" + "a") + ("\n\t" + ("\n\t" + ((("\n\t" + "hello") + ("" + "foo bar")) + (("foo bar" + "") + "a"))))))))) {
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            v16 += 2147483648
                                                                        case "\n\t":
                                                                            var v74 uint8 = 0
                                                                            _ = v74
                                                                        default:
                                                                            v73 -= v73
                                                                        }
                                                                        if true {
                                                                            fmt.Println(fi44)
                                                                            var v75 uint = uint(v60)
                                                                            var v76 int32 = -354
                                                                            _ = v75
                                                                            _ = v76
                                                                        } else {
                                                                            if false {
                                                                                v16 = v17
                                                                                delete(v4, v17)
                                                                            } else {
                                                                                fmt.Println(fi50, v61)
                                                                                if ((-113 >= -23) && false) {
                                                                                    var v77 uint64 = v60
                                                                                    v60 = genFn1(764)
                                                                                    _ = v77
                                                                                }
                                                                                var v78 float32 = v61
                                                                                _ = v78
                                                                            }
                                                                            fmt.Println(fi50, v17)
                                                                            v79 := []int{fi50, 624, fi44, (fi50 - -37)}
                                                                            var v80 int16 = 0
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        switch ("foo bar" + "a") {
                                                                        case "\n\t":
                                                                            v17 = v17
                                                                        case "":
                                                                            if (v17 > (v16 ^ v16)) {
                                                                                switch "" {
                                                                                case "a":
                                                                                    var v81 int16 = 1000
                                                                                    _ = v81
                                                                                default:
                                                                                    switch (fi50 * fi50) {
                                                                                    case -2147483648:
                                                                                        var v82 uint8 = 1
                                                                                        _ = v82
                                                                                    }
                                                                                }
                                                                                v83 := GenS2{F0: (v73 >> (((v3 / (((v3 * (^(v73)))) | 1))) & 63)), F1: 390}
                                                                                if (true && (fi50 <= (fi44 ^ (fi44 % (((fi44 * fi44)) | 1))))) {
                                                                                    v17 += genFn0(42, fi44)
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        switch fi44 {
                                                                                        case 128:
                                                                                            var v84 int32 = 0
                                                                                            _ = v84
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v85 := []int{(fi50 | fi50), (fi44 % ((-475) | 1)), fi50, fi44}
                                                                                        _ = v85
                                                                                    case "":
                                                                                        var v86 float32 = float32(0.0)
                                                                                        _ = v86
                                                                                    }
                                                                                    var v87 uint32 = v16
                                                                                    _ = v87
                                                                                } else {
                                                                                    var v88 uint32 = 0
                                                                                    v89 := GenS4{F0: v60}
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                }
                                                                                v16 = v16
                                                                                v60 = v60
                                                                                _ = v83
                                                                            }
                                                                        default:
                                                                            v90 := GenS2{F0: 1, F1: 49}
                                                                            _ = v90
                                                                        }
                                                                        v3 = (v73 | (v73 << ((v73) & 63)))
                                                                        switch (^(fi44)) {
                                                                        case 0:
                                                                            var v91 string = "hello"
                                                                            _ = v91
                                                                        case 32767:
                                                                            if false {
                                                                                v92 := []int{-9223372036854775808, 1}
                                                                                _ = v92
                                                                            } else {
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    var v93 uint8 = 55
                                                                                    _ = v93
                                                                                case "hello":
                                                                                    if false {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            if ((!true) && true) {
                                                                                                v61 = -2.25
                                                                                                v73 -= v3
                                                                                                v61 += (1e10 - v61)
                                                                                                var v94 bool = true
                                                                                                _ = v94
                                                                                            }
                                                                                        case "":
                                                                                            switch ("a" + "hello") {
                                                                                            case "a":
                                                                                                v95 := GenS2{F0: v73, F1: 796}
                                                                                                _ = v95
                                                                                            case "":
                                                                                                v17 = v16
                                                                                            }
                                                                                            v17 = genFn0(255, 2147483647)
                                                                                            v96 := GenS6{F0: 2147483647, F1: 0, F2: (^(v3))}
                                                                                            _ = v96
                                                                                        case "foo bar":
                                                                                            v97 := v4[uint32(v60)]
                                                                                            _ = v97
                                                                                        default:
                                                                                            var v98 int16 = 1
                                                                                            _ = v98
                                                                                        }
                                                                                        v99 := []int{fi50, fi44, fi50, (-(fi50))}
                                                                                        fmt.Println(v16, v17)
                                                                                        _ = v99
                                                                                    } else {
                                                                                        if (v60 != (v60 % ((genFn1(v17)) | 1))) {
                                                                                            v17 = v17
                                                                                            v100 := GenS7{F0: false, F1: (("hello" + "") + ("" + "\n\t")), F2: (v17 % ((genFn0(32767, 42)) | 1))}
                                                                                            var v101 uint64 = genFn1(genFn0(2, fi44))
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                        }
                                                                                        fmt.Println(v61)
                                                                                    }
                                                                                case "a":
                                                                                    var v102 uint16 = 769
                                                                                    _ = v102
                                                                                default:
                                                                                    v103 := []int{531, (fi44 * (fi44 - (fi44 / ((-219) | 1)))), (fi50 * 4294967295), (^(fi50))}
                                                                                    _ = v103
                                                                                }
                                                                                v61 = -2.25
                                                                                v61 = 3.14159
                                                                                v104 := GenS2{F0: 18446744073709551615, F1: 42}
                                                                                _ = v104
                                                                            }
                                                                        case -181:
                                                                            var v105 int32 = 495
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    _ = v73
                                                                }
                                                            }
                                                        }
                                                        v61 += (3.14159 + v61)
                                                    }
                                                    var v106 uint8 = 203
                                                    if true {
                                                        v60 += 2
                                                    }
                                                    var v107 uint8 = v106
                                                    _ = v106
                                                    _ = v107
                                                }
                                                _ = fi50
                                                _ = v60
                                                _ = v61
                                            }
                                            v16 = v16
                                            v108 := []int{(fi44 >> (((v3 * (v3 ^ 42))) & 63)), (fi44 + fi44)}
                                            _ = v108
                                        }
                                    } else {
                                        if false {
                                            v109 := v4[(^(v16))]
                                            var v110 int = (fi44 & (fi44 / (((fi44 + fi44)) | 1)))
                                            _ = v109
                                            _ = v110
                                        }
                                        var v111 uint64 = 42
                                        _ = v111
                                    }
                                default:
                                    if false {
                                        v17 -= uint32(v3)
                                        var v112 bool = (!true)
                                        _ = v112
                                    }
                                }
                                switch "hello" {
                                case "hello":
                                    v16 -= genFn0(42, 42)
                                case "":
                                    v17 += 42
                                case "a":
                                    var v113 int8 = -128
                                    _ = v113
                                default:
                                    v3 += (v3 << ((v3) & 63))
                                }
                            } else {
                                v114 := []int{(fi44 >> ((v3) & 63)), (fi44 & (fi44 >> (((v3 * v3)) & 63))), -473, 9223372036854775807}
                                _ = v114
                            }
                        default:
                            if false {
                                v115 := GenS1{F0: (v3 ^ v3), F1: 18}
                                _ = v115
                            } else {
                                var v116 uint16 = uint16(v16)
                                v17 = (^(v16))
                                _ = v116
                            }
                        }
                        _ = fi44
                    }
                case "hello":
                    fmt.Println(v16, v16)
                default:
                    var v117 int64 = 9223372036854775807
                    _ = v117
                }
            }
        } else {
            _ = len(v2)
            fmt.Println(v17, v17, v3)
        }
    } else {
        v118 := GenS2{F0: 2, F1: 186}
        switch 563 {
        case -1:
            var v119 bool = false
            _ = v119
        case -239:
            v120 := v2[0]
            _ = v120
        case 32768:
            v17 -= v17
        default:
            switch "foo bar" {
            case "hello":
                fmt.Println(v3, v16, v3)
            default:
                var v121 uint64 = 42
                _ = v121
            }
        }
        _ = v118
    }
    if true {
        v122 := GenS8{F0: v3}
        var v123 int16 = 1000
        _ = v122
        _ = v123
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
