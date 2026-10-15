package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int; F2 bool }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 int8 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS8 struct { F0 bool; F1 int16; F2 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(fi4)
        var v5 float64 = float64(0.1)
        v6 := GenS1{F0: 42, F1: 127}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int32 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        if ("\n\t" < ("" + "a")) {
            v7 = 2147483647
            v7 -= (^(v7))
        }
        v7 = (v7 | v7)
        var v9 int64 = 817
        _ = fi8
        _ = v9
    }
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 bool) uint32 {
    if p1 {
        var v10 uint16 = 135
        switch "" {
        case "\n\t":
            if false {
                var v11 float64 = 1e10
                for fi12 := 0; fi12 < 14; fi12++ {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v14 := []int{(fi13 >> ((4294967295) & 63)), (fi13 * fi13)}
                        var v15 string = "foo bar"
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v11 += v11
                    _ = fi12
                }
                _ = v11
            }
        case "":
            if p1 {
                if (v10 <= (v10 - (^(v10)))) {
                    if p1 {
                        var v16 int16 = p0
                        _ = v16
                    } else {
                        v17 := GenS2{F0: "foo bar"}
                        v18 := GenS3{F0: p0, F1: -573, F2: false}
                        v19 := GenS1{F0: 100, F1: 28}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    var v20 uint64 = 18446744073709551615
                    _ = v20
                } else {
                    fmt.Println(v10, p0, v10)
                    var v21 string = "foo bar"
                    var v22 float32 = 1.5
                    _ = v21
                    _ = v22
                }
                if (2 < 0) {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        if true {
                            v24 := []int{(fi23 ^ (^(fi23))), -1, (fi23 & 128)}
                            _ = v24
                        }
                        _ = fi23
                    }
                }
            }
        case "hello":
            p1 = true
        default:
            switch -2147483648 {
            case 128:
                v10 = (v10 ^ 32767)
            case 127:
                switch 127 {
                case -868:
                    var v25 uint32 = 2147483648
                    _ = v25
                case -691:
                    p1 = (v10 != (v10 % ((v10) | 1)))
                case 0:
                    var v26 bool = p1
                    _ = v26
                }
            case 32768:
                switch (("\n\t" + ("foo bar" + "hello")) + "hello") {
                case "\n\t":
                    v10 += (v10 >> ((1) & 63))
                }
            }
        }
        _ = v10
    } else {
        if (42 < 4294967295) {
            p0 += (-(p0))
            switch "\n\t" {
            case "":
                var v27 int8 = 11
                _ = v27
            }
            var v28 bool = ("" < "")
            _ = v28
        }
    }
    var v29 uint64 = 0
    _ = p0
    _ = p1
    _ = v29
    return 2147483647
}

func main() {
    var v1 bool = ((!true) && (!true))
    var v2 int = 127
    var v3 uint8 = 1
    fmt.Println(v3, v3)
    if (("hello" + (("" + ("foo bar" + "\n\t")) + (("foo bar" + (("foo bar" + "\n\t") + "foo bar")) + ("" + "hello")))) == "hello") {
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 bool = v1
            for fi32 := 0; fi32 < 7; fi32++ {
                v1 = (((0.1 + 3.14159) >= 1.5) || ((v3 << ((255) & 63)) <= v3))
                _ = fi32
            }
            v33 := GenS2{F0: "a"}
            _ = fi30
            _ = v31
            _ = v33
        }
        var v34 int = 32768
        _ = v34
    } else {
        v2 -= v2
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            v37 := []int{(v2 % (((v2 * 128)) | 1)), int(v3)}
            if v1 {
                fmt.Println(v2, v3, v2)
                v38 := GenS4{F0: (^(v2)), F1: v1, F2: ((((float32(-2.25) * 1e10) * 1.5) / 1.0) < 0.0)}
                _ = v38
            }
            var v39 int = -9223372036854775808
            var v40 uint64 = 995
            var v41 int64 = 0
            if false {
                v40 -= v40
            } else {
                switch "foo bar" {
                case "\n\t":
                    v1 = true
                default:
                    var v42 bool = v1
                    _ = v42
                }
            }
            v43 := GenS4{F0: (v39 * fi36), F1: v1, F2: v1}
            _ = fi36
            _ = v37
            _ = v39
            _ = v40
            _ = v41
            _ = v43
        }
        v44 := GenS5{F0: 2147483648, F1: 42, F2: -18}
        _ = fi35
        _ = v44
    }
    switch "foo bar" {
    case "foo bar":
        for fi45 := 0; fi45 < 14; fi45++ {
            v3 = (v3 / ((v3) | 1))
            for fi46 := 0; fi46 < 6; fi46++ {
                v1 = ("" >= "hello")
                v47 := GenS6{F0: -27, F1: -99}
                _ = fi46
                _ = v47
            }
            _ = fi45
        }
    default:
        if v1 {
            v48 := []int{(v2 * v2), v2}
            v3 += v3
            var v49 string = "hello"
            _ = v48
            _ = v49
        }
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 int8 = 0
        _ = fi50
        _ = v51
    }
    var v52 int32 = 255
    switch v2 {
    case 32768:
        var v53 int8 = 0
        _ = v53
    case -383:
        v54 := []int{(v2 / (((v2 & v2)) | 1)), (v2 << ((18446744073709551615) & 63)), 4294967295, v2}
        _ = v54
    case 0:
        fmt.Println(v52)
    }
    if v1 {
        for fi55 := 0; fi55 < 16; fi55++ {
            v56 := []int{(^(fi55)), (v2 << ((4294967295) & 63)), (fi55 ^ (^(fi55)))}
            if len(v56) > 0 {
                v56[0] = (fi55 << ((251) & 63))
            }
            _ = fi55
            _ = v56
        }
        for fi57 := 0; fi57 < 15; fi57++ {
            fmt.Println(fi57, v52, v52)
            switch "hello" {
            case "hello":
                for fi58 := 0; fi58 < 14; fi58++ {
                    switch "\n\t" {
                    case "a":
                        var v59 string = ("\n\t" + "foo bar")
                        _ = v59
                    }
                    _ = fi58
                }
            default:
                if v1 {
                    switch ((("" + "hello") + "") + ("\n\t" + ("\n\t" + ""))) {
                    case "a":
                        var v60 int8 = 127
                        _ = v60
                    case "\n\t":
                        if (v52 >= (v52 << ((0) & 63))) {
                            v61 := GenS5{F0: 1, F1: 127, F2: -97}
                            _ = v61
                        } else {
                            fmt.Println(v2)
                            v62 := GenS1{F0: 100, F1: -1}
                            switch "foo bar" {
                            case "\n\t":
                                v63 := []int{(v2 / ((fi57) | 1)), (fi57 | v2), (v2 ^ 601)}
                                _ = v63
                            case "a":
                                v64 := GenS6{F0: 1, F1: 127}
                                _ = v64
                            }
                            _ = v62
                        }
                    case "":
                        var v65 uint32 = 0
                        _ = v65
                    default:
                        v3 += v3
                    }
                    for fi66 := 0; fi66 < 19; fi66++ {
                        v2 += (-(fi57))
                        v52 += 42
                        var v67 int64 = 211
                        _ = fi66
                        _ = v67
                    }
                }
            }
            switch (fi57 | (v2 << ((255) & 63))) {
            case 904:
                for fi68 := 0; fi68 < 19; fi68++ {
                    v69 := GenS4{F0: (fi57 ^ 32768), F1: (1.5 >= -2.25), F2: (0 > v3)}
                    switch fi68 {
                    case -269:
                        v70 := GenS5{F0: 417, F1: 127, F2: -1}
                        _ = v70
                    case -888:
                        v2 = (^(v2))
                    default:
                        v3 = (^(v3))
                    }
                    fmt.Println(v3, v52)
                    if (2 != 356) {
                        v71 := []int{(v2 & 127), 975, 663}
                        v2 -= fi68
                        _ = v71
                    }
                    _ = fi68
                    _ = v69
                }
            case -2147483648:
                for fi72 := 0; fi72 < 12; fi72++ {
                    if false {
                        if true {
                            v1 = false
                        } else {
                            fmt.Println(fi72, v52)
                            if (int16(v3) > 32767) {
                                var v73 int64 = -396
                                if v1 {
                                    v3 = (^(v3))
                                    if ((v2 % (((-(fi57))) | 1)) >= 765) {
                                        v2 = (fi72 << ((4294967295) & 63))
                                    }
                                    v74 := []int{}
                                    _ = v74
                                }
                                v52 += (v52 + v52)
                                _ = v73
                            }
                            switch -1 {
                            case -842:
                                var v75 string = "foo bar"
                                _ = v75
                            case -2147483648:
                                v76 := GenS4{F0: fi57, F1: ("hello" <= "foo bar"), F2: v1}
                                _ = v76
                            }
                            var v77 uint64 = 42
                            fmt.Println(v52)
                            _ = v77
                        }
                        v3 = (v3 ^ v3)
                        if (!v1) {
                            var v78 int16 = -32768
                            v79 := []int{fi57, 127, (fi57 * 707)}
                            _ = v78
                            _ = v79
                        } else {
                            v1 = v1
                            switch 214 {
                            case 996:
                                switch (fi57 / (((fi57 & 9223372036854775807)) | 1)) {
                                case -699:
                                    v80 := GenS7{F0: 541, F1: 0, F2: 65535}
                                    _ = v80
                                default:
                                    var v81 int32 = (v52 & v52)
                                    _ = v81
                                }
                            default:
                                fmt.Println(v52)
                            }
                            v82 := GenS4{F0: 1, F1: (((!((((v3 % (((v3 * v3)) | 1)) == (^(v3))) || (true || (1 != 2))) || (2 != 657))) && v1) || (!((v3 + v3) >= (v3 << ((4294967295) & 63))))), F2: v1}
                            _ = v82
                        }
                    } else {
                        var v83 uint8 = (v3 % (((v3 / (((^(v3))) | 1))) | 1))
                        fmt.Println(v3)
                        _ = v83
                    }
                    var v84 uint32 = 975
                    var v85 uint64 = 1
                    _ = fi72
                    _ = v84
                    _ = v85
                }
            case -9223372036854775808:
                if (584 < 1) {
                    v1 = (0 == v52)
                    v86 := []int{(fi57 << ((65535) & 63))}
                    _ = v86
                } else {
                    v87 := []int{(fi57 ^ 127), -631}
                    _ = v87
                }
            }
            _ = fi57
        }
    }
    for fi88 := 0; fi88 < 7; fi88++ {
        v89 := []int{(fi88 * fi88), 2147483647}
        _ = fi88
        _ = v89
    }
    for fi90 := 0; fi90 < 5; fi90++ {
        v91 := GenS8{F0: (v1 || false), F1: int16(v52), F2: 42}
        v92 := []int{v2}
        if false {
            if false {
                v52 = (v52 ^ (^(v52)))
                for fi93 := 0; fi93 < 15; fi93++ {
                    var v94 float32 = 3.14159
                    fmt.Println(v3, v3)
                    _ = fi93
                    _ = v94
                }
                var v95 bool = ((42 < 2) && true)
                _ = v95
            } else {
                var v96 uint8 = (v3 ^ 2)
                if ("\n\t" < "\n\t") {
                    var v97 float32 = (1e10 - ((0.0 / 1e10) * 1.5))
                    var v98 int8 = -19
                    _ = v97
                    _ = v98
                }
                _ = v96
            }
        } else {
            for fi99 := 0; fi99 < 16; fi99++ {
                var v100 uint8 = v3
                if len(v92) >= 2 {
                    v92 = v92[1:len(v92)]
                }
                v101 := GenS1{F0: int8(v52), F1: 100}
                _ = len(v92)
                _ = fi99
                _ = v100
                _ = v101
            }
            if true {
                if v1 {
                    for fi102 := 0; fi102 < 13; fi102++ {
                        if len(v92) >= 2 {
                            v92 = v92[1:len(v92)]
                        }
                        if ((v1 && v1) && (57 == 1)) {
                            if (v1 && v1) {
                                v103 := GenS9{F0: -490}
                                _ = v103
                            } else {
                                var v104 int64 = 415
                                if len(v92) > 0 {
                                    v92[0] = (fi90 ^ -448)
                                }
                                v105 := GenS10{F0: 1000}
                                _ = v104
                                _ = v105
                            }
                            fmt.Println(v1, v3)
                        }
                        v106 := GenS3{F0: int16(fi90), F1: (^(fi90)), F2: (372 > genFn1(1000, ((9223372036854775807 != 1) || (!v1))))}
                        _ = fi102
                        _ = v106
                    }
                    v107 := GenS8{F0: (false && false), F1: 0, F2: 2147483647}
                    _ = v107
                }
                var v108 int64 = 9223372036854775807
                for fi109 := 0; fi109 < 10; fi109++ {
                    v1 = v1
                    v110 := GenS3{F0: int16(v108), F1: (-(v2)), F2: v1}
                    _ = fi109
                    _ = v110
                }
                _ = v108
            } else {
                if (-1 >= 0) {
                    v52 -= -2147483648
                    if (1 != 854) {
                        fmt.Println(v52, v2, fi90)
                        if (!(v1 && (!v1))) {
                            var v111 float32 = 0.0
                            _ = v111
                        }
                    } else {
                        for fi112 := 0; fi112 < 15; fi112++ {
                            var v113 uint64 = 418
                            v114 := GenS8{F0: v1, F1: 1, F2: 2147483648}
                            _ = fi112
                            _ = v113
                            _ = v114
                        }
                    }
                    v1 = (!true)
                } else {
                    if (42 <= 182) {
                        var v115 int32 = (v52 ^ v52)
                        _ = v115
                    }
                    switch v2 {
                    case 9223372036854775807:
                        for fi116 := 0; fi116 < 13; fi116++ {
                            if (v1 || ((^(v52)) >= (v52 & 42))) {
                                switch (fi116 / (((-(v2))) | 1)) {
                                case -1:
                                    v92 = append(v92, fi116)
                                    var v117 uint = 65535
                                    var v118 int32 = v52
                                    _ = v117
                                    _ = v118
                                case 286:
                                    v52 = v52
                                case 127:
                                    fmt.Println(v2)
                                }
                                _ = len(v92)
                            }
                            v3 += v3
                            _ = fi116
                        }
                    }
                }
                v119 := GenS3{F0: 566, F1: (fi90 | fi90), F2: ((int8(v52) == 127) && ("a" > "\n\t"))}
                var v120 int64 = genFn0()
                _ = v119
                _ = v120
            }
        }
        _ = fi90
        _ = v91
        _ = v92
    }
    for fi121 := 0; fi121 < 13; fi121++ {
        var v122 uint = 18446744073709551615
        switch v2 {
        case 32768:
            v123 := []int{}
            _ = v123
        case 2147483647:
            var v124 float32 = -2.25
            _ = v124
        }
        v125 := GenS6{F0: int8(v3), F1: -128}
        _ = fi121
        _ = v122
        _ = v125
    }
    v126 := GenS7{F0: 155, F1: 2, F2: 65535}
    v127 := []int{v2, (v2 | v2)}
    switch v2 {
    case -470:
        v128 := GenS8{F0: (1000 <= 1), F1: 1, F2: 638}
        _ = v128
    default:
        fmt.Println(v52)
    }
    var v129 int8 = 42
    if true {
        fmt.Println(v3, v1)
        v129 = -87
        v3 += v3
    }
    fmt.Println(v2, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v127:", len(v127), v127)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v126
    _ = v127
    _ = v129
    _ = v2
    _ = v3
    _ = v52
}
