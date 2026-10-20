package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int16; F1 bool; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 int8; F2 int32 }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string }
func genFn0() uint64 {
    if (true && (!(!true))) {
        v6 := GenS1{F0: -363, F1: 361, F2: 1}
        if false {
            for fi7 := 0; fi7 < 19; fi7++ {
                fmt.Println(fi7)
                v8 := []int{9223372036854775807, (fi7 % ((847) | 1)), (-(fi7)), fi7}
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int8 = 42
                _ = fi9
                _ = v10
            }
            if (false || false) {
                var v11 bool = ((!true) && (32768 != 711))
                var v12 float64 = (1e10 / 1e10)
                if ((-2.25 != ((-2.25 - v12) + v12)) || (!v11)) {
                    fmt.Println(v11, v12)
                    v12 = v12
                    if true {
                        v13 := GenS2{F0: 42, F1: (4294967295 <= 42), F2: 42}
                        var v14 int16 = -1
                        _ = v13
                        _ = v14
                    }
                }
                _ = v11
                _ = v12
            }
        }
        _ = v6
    }
    var v15 int64 = 0
    var v16 float64 = 1.5
    v17 := []int{9223372036854775807}
    _ = v15
    _ = v16
    _ = v17
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (0.0 + 0.0)
    v3 := []string{}
    v4 := []uint32{42, 42, 4294967295}
    var v5 bool = true
    var v18 int64 = 397
    fmt.Println(v2, v2)
    var v19 float32 = (0.0 * ((((float32((v2 - ((0.0 * (float64(((-2.25 + -2.25) / 1e10)) - 0.1)) - v2))) * -2.25) * (((-2.25 + ((0.0 * 0.0) - 3.14159)) / 3.14159) - 0.0)) - float32(v2)) / ((0.0 + 0.1) - 1.5)))
    if (-1 > 9223372036854775807) {
        for fi20 := 0; fi20 < 7; fi20++ {
            v18 = 0
            v21 := GenS3{F0: 1}
            _ = fi20
            _ = v21
        }
    } else {
        v19 = v19
        switch -667 {
        case -9223372036854775808:
            if v5 {
                fmt.Println(v19, v2, v5)
                fmt.Println(v5)
            } else {
                v22 := GenS1{F0: 147, F1: 42, F2: 0}
                _ = v22
            }
        case -560:
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS4{F0: (v18 - (v18 >> ((2) & 63))), F1: 2}
                var v25 int8 = 0
                _ = fi23
                _ = v24
                _ = v25
            }
        case 42:
            var v26 int16 = 1
            _ = v26
        }
    }
    if (0 < 800) {
        v2 = v2
        for fi27 := 0; fi27 < 13; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                v18 = (v18 & (v18 & (v18 / ((v18) | 1))))
                var v29 string = ("" + "foo bar")
                switch 32768 {
                case 9223372036854775807:
                    if len(v3) > 0 {
                        v3[0] = v29
                    }
                }
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 uint = 18446744073709551615
            _ = fi30
            _ = v31
        }
    }
    var v32 int32 = -801
    if (v5 || (v5 || v5)) {
        switch 0 {
        case 42:
            v2 -= v2
        }
        if v5 {
            v18 = v18
        } else {
            var v33 uint16 = 1
            _ = v33
        }
    }
    var v34 uint = 42
    var v35 uint16 = 383
    v36 := GenS1{F0: 127, F1: genFn0(), F2: 0}
    v37 := GenS5{F0: (494 >= 65535), F1: 0}
    fmt.Println(v5, v2)
    if true {
        fmt.Println(v18, v34)
    }
    if (v5 || v5) {
        v32 = v32
        _ = len(v3)
    }
    if v5 {
        if false {
            for fi38 := 0; fi38 < 2; fi38++ {
                if true {
                    delete(v1, (fi38 & 981))
                    if false {
                        fmt.Println(v19, v35, v5)
                        fmt.Println(fi38)
                        var v39 uint64 = 179
                        _ = v39
                    }
                    v35 = 32767
                } else {
                    v40 := GenS6{F0: v5, F1: false}
                    _ = v40
                }
                v41 := GenS6{F0: v5, F1: true}
                v35 = 0
                _ = fi38
                _ = v41
            }
        } else {
            v42 := GenS1{F0: 9223372036854775807, F1: 1, F2: -7}
            var v43 uint8 = 1
            _ = v42
            _ = v43
        }
    }
    switch "a" {
    case "hello":
        var v44 float32 = v19
        _ = v44
    }
    if v5 {
        if v5 {
            switch "hello" {
            case "":
                if v5 {
                    switch "a" {
                    case "a":
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    case "":
                        switch ("" + "a") {
                        case "a":
                            if false {
                                switch 128 {
                                case -2147483648:
                                    v45 := GenS2{F0: 42, F1: (-128 != -107), F2: (^(v35))}
                                    _ = v45
                                case 42:
                                    var v46 uint64 = genFn0()
                                    _ = v46
                                case 0:
                                    fmt.Println(v5, v35, v18)
                                }
                                for fi47 := 0; fi47 < 8; fi47++ {
                                    fmt.Println(v18, v34)
                                    _ = fi47
                                }
                                switch "a" {
                                case "\n\t":
                                    for fi48 := 0; fi48 < 16; fi48++ {
                                        v49 := GenS7{F0: (fi48 ^ (fi48 - (fi48 - fi48))), F1: v32}
                                        switch "\n\t" {
                                        case "\n\t":
                                            _ = len(v1)
                                        }
                                        for fi50 := 0; fi50 < 7; fi50++ {
                                            var v51 uint = v34
                                            _ = fi50
                                            _ = v51
                                        }
                                        _ = fi48
                                        _ = v49
                                    }
                                case "a":
                                    var v52 uint16 = (v35 - 196)
                                    _ = v52
                                default:
                                    var v53 string = ("" + "")
                                    _ = v53
                                }
                            }
                        }
                    default:
                        if (-452 == 1) {
                            if (v35 > v35) {
                                v54 := GenS8{F0: 1000, F1: (!true), F2: v34}
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    for fi56 := 0; fi56 < 4; fi56++ {
                                        var v57 uint = v34
                                        v58 := GenS3{F0: v35}
                                        switch (fi56 % (((fi56 - (-(fi56)))) | 1)) {
                                        case 127:
                                            if (true && v5) {
                                                var v59 int64 = (v18 + v18)
                                                var v60 uint16 = 42
                                                var v61 uint64 = 708
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            } else {
                                                var v62 uint32 = 42
                                                _ = v62
                                            }
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = 960
                                            }
                                            if v5 {
                                                fmt.Println(v5, v32)
                                            }
                                            if true {
                                                if v5 {
                                                    var v63 int = fi55
                                                    if true {
                                                        v19 = 0.1
                                                        if true {
                                                            if (-1 == 100) {
                                                                if (248 != 1) {
                                                                    v5 = v5
                                                                }
                                                                var v64 int32 = (v32 << ((v57) & 63))
                                                                _ = len(v1)
                                                                _ = v64
                                                            }
                                                        } else {
                                                            var v65 uint8 = 72
                                                            _ = v65
                                                        }
                                                        if (("hello" != ("foo bar" + "\n\t")) && (("a" + ("" + ("" + ("foo bar" + "a")))) < ("foo bar" + ""))) {
                                                            v1[(fi56 * v63)] = fi56
                                                            var v66 int64 = v18
                                                            _ = v66
                                                        } else {
                                                            if v5 {
                                                                var v67 uint64 = 18446744073709551615
                                                                v18 -= (v18 * (-(v18)))
                                                                _ = v67
                                                            } else {
                                                                var v68 int32 = (^(v32))
                                                                if v5 {
                                                                    v68 = (v68 - v32)
                                                                    var v69 int = fi55
                                                                    switch (("a" + "foo bar") + "hello") {
                                                                    case "":
                                                                        v2 = -2.25
                                                                        var v70 uint32 = 519
                                                                        if (!true) {
                                                                            v1[int(v18)] = (-(fi56))
                                                                            if v5 {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ((("foo bar" + "\n\t") + "\n\t") + "hello")
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v35 -= v35
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (("hello" + ("a" + "a")) + "a")
                                                                        }
                                                                        v32 -= (v32 - (v68 << ((v34) & 63)))
                                                                        _ = v70
                                                                    }
                                                                    _ = v69
                                                                }
                                                                _ = v68
                                                            }
                                                            fmt.Println(v63, fi55)
                                                            v71 := GenS5{F0: (v5 && v5), F1: 42}
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            fmt.Println(v19, fi56)
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v72 uint8 = 127
                                                                _ = v72
                                                            case "foo bar":
                                                                switch ("foo bar" + ("a" + (("foo bar" + "a") + ""))) {
                                                                case "foo bar":
                                                                    var v73 bool = (!((v18 & (v18 * v18)) <= v18))
                                                                    _ = v73
                                                                }
                                                            }
                                                            _ = v71
                                                        }
                                                    }
                                                    v74 := GenS3{F0: v35}
                                                    _ = v63
                                                    _ = v74
                                                } else {
                                                    v32 -= (^(v32))
                                                }
                                            }
                                        }
                                        if v5 {
                                            if true {
                                                switch "\n\t" {
                                                case "a":
                                                    var v75 bool = (2 == 42)
                                                    _ = v75
                                                case "hello":
                                                    var v76 int32 = -1
                                                    _ = v76
                                                case "\n\t":
                                                    var v77 uint = 1
                                                    _ = v77
                                                default:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v78 := GenS2{F0: 0, F1: v5, F2: (v35 - 65535)}
                                                        _ = v78
                                                    case "":
                                                        var v79 uint = (v57 + v34)
                                                        delete(v1, fi56)
                                                        _ = v79
                                                    case "foo bar":
                                                        var v80 uint8 = 214
                                                        _ = v80
                                                    default:
                                                        if len(v3) > 0 {
                                                            v3[0] = "foo bar"
                                                        }
                                                    }
                                                }
                                                v81 := GenS9{F0: genFn0()}
                                                v32 = -2147483648
                                                var v82 int32 = v32
                                                v19 += float32((float64((v19 + (v19 * 3.14159))) / 0.1))
                                                _ = v81
                                                _ = v82
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = 65535
                                            }
                                        } else {
                                            var v83 string = ("hello" + "\n\t")
                                            _ = v83
                                        }
                                        switch ("foo bar" + "hello") {
                                        case "a":
                                            switch "foo bar" {
                                            case "hello":
                                                var v84 float32 = 1e10
                                                _ = v84
                                            default:
                                                v18 += (^(v18))
                                            }
                                        case "foo bar":
                                            fmt.Println(v32, fi56, fi56)
                                        }
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                    var v85 int16 = 681
                                    v5 = v5
                                    _ = fi55
                                    _ = v85
                                }
                                if (4294967295 != 888) {
                                    var v86 float64 = v2
                                    switch -2147483648 {
                                    case 0:
                                        if false {
                                            v87 := GenS10{F0: 127, F1: -128, F2: -2147483648}
                                            var v88 uint64 = genFn0()
                                            _ = v87
                                            _ = v88
                                        } else {
                                            if (42 <= v35) {
                                                fmt.Println(v2, v86, v35)
                                            } else {
                                                if true {
                                                    var v89 float64 = ((0.1 * v2) + -2.25)
                                                    fmt.Println(v35)
                                                    _ = v89
                                                } else {
                                                    var v90 int64 = v18
                                                    var v91 float32 = v19
                                                    var v92 int32 = v32
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                }
                                            }
                                            v18 = (^(v18))
                                            v93 := v1[-127]
                                            _ = v93
                                        }
                                    default:
                                        if v5 {
                                            var v94 int = 32767
                                            v34 -= (^(v34))
                                            var v95 int8 = -1
                                            _ = v94
                                            _ = v95
                                        } else {
                                            var v96 int32 = (v32 | -2147483648)
                                            var v97 uint = 679
                                            _ = v96
                                            _ = v97
                                        }
                                    }
                                    switch -2147483648 {
                                    case 2147483647:
                                        var v98 string = ("a" + ("a" + "\n\t"))
                                        _ = v98
                                    case 901:
                                        if len(v3) > 0 {
                                            v3[0] = ("foo bar" + ("" + "hello"))
                                        }
                                    case -974:
                                        v5 = ((!v5) && v5)
                                    default:
                                        fmt.Println(v2)
                                    }
                                    _ = v86
                                }
                                _ = v54
                            }
                            if v5 {
                                v35 -= (v35 & v35)
                                var v99 int32 = v32
                                if (42 > uint32(v35)) {
                                    switch "" {
                                    case "foo bar":
                                        if false {
                                            v35 = 42
                                            switch 128 {
                                            case 128:
                                                var v100 int = 32767
                                                _ = v100
                                            case 32767:
                                                _ = len(v3)
                                            case 9223372036854775807:
                                                _ = len(v4)
                                            }
                                        } else {
                                            v5 = (v5 && true)
                                        }
                                    }
                                    var v101 uint8 = 128
                                    _ = v101
                                } else {
                                    switch 0 {
                                    case 9223372036854775807:
                                        if (388 > genFn0()) {
                                            v102 := GenS9{F0: 1}
                                            switch "a" {
                                            case "":
                                                var v103 int = 4294967295
                                                _ = v103
                                            case "hello":
                                                v35 -= v35
                                            case "\n\t":
                                                v104 := GenS11{F0: v99, F1: 42}
                                                _ = v104
                                            default:
                                                var v105 float64 = 3.14159
                                                _ = v105
                                            }
                                            _ = v102
                                        }
                                    case -2147483648:
                                        v106 := GenS12{F0: 512}
                                        _ = v106
                                    default:
                                        switch "\n\t" {
                                        case "":
                                            if (!(1 < 42)) {
                                                switch -1 {
                                                case -783:
                                                    var v107 int8 = -128
                                                    _ = v107
                                                default:
                                                    var v108 int8 = 0
                                                    _ = v108
                                                }
                                                switch -499 {
                                                case 9223372036854775807:
                                                    var v109 float64 = 1e10
                                                    _ = v109
                                                case 4294967295:
                                                    for fi110 := 0; fi110 < 19; fi110++ {
                                                        switch 2147483647 {
                                                        case 1:
                                                            if ((fi110 + fi110) == (fi110 ^ (-(fi110)))) {
                                                                fmt.Println(v18, v2, v2)
                                                                for fi111 := 0; fi111 < 5; fi111++ {
                                                                    var v112 bool = ((v32 >> ((642) & 63)) != v99)
                                                                    v112 = true
                                                                    _ = fi111
                                                                    _ = v112
                                                                }
                                                                v113 := GenS4{F0: v18, F1: 2147483647}
                                                                _ = v113
                                                            } else {
                                                                v32 -= 1
                                                            }
                                                        case 32767:
                                                            switch fi110 {
                                                            case 127:
                                                                v114 := GenS5{F0: true, F1: 14}
                                                                _ = v114
                                                            default:
                                                                for fi115 := 0; fi115 < 2; fi115++ {
                                                                    if (((float64(3.14159) + float64(v19)) >= ((v2 * v2) * (3.14159 * (v2 / 3.14159)))) && ((fi110 >> (((^(v34))) & 63)) <= fi115)) {
                                                                        var v116 string = "foo bar"
                                                                        fmt.Println(v99, v5)
                                                                        v117 := GenS1{F0: (-(fi115)), F1: 18446744073709551615, F2: 100}
                                                                        _ = v116
                                                                        _ = v117
                                                                    } else {
                                                                        v118 := GenS4{F0: 1, F1: 2147483647}
                                                                        _ = v118
                                                                    }
                                                                    if false {
                                                                        var v119 float32 = v19
                                                                        _ = v119
                                                                    }
                                                                    fmt.Println(fi110, v18)
                                                                    _ = fi115
                                                                }
                                                            }
                                                        case 4294967295:
                                                            var v120 float32 = v19
                                                            _ = v120
                                                        }
                                                        if v5 {
                                                            if v5 {
                                                                v18 = v18
                                                                switch fi110 {
                                                                case 0:
                                                                    if false {
                                                                        v32 += (v32 & v32)
                                                                        if (false || true) {
                                                                            var v121 uint32 = 721
                                                                            _ = v121
                                                                        }
                                                                    }
                                                                case -670:
                                                                    var v122 uint16 = 743
                                                                    _ = v122
                                                                }
                                                                for fi123 := 0; fi123 < 20; fi123++ {
                                                                    v124 := GenS6{F0: false, F1: (uint64(v34) < 2)}
                                                                    if v5 {
                                                                        var v125 string = ""
                                                                        _ = v125
                                                                    }
                                                                    _ = fi123
                                                                    _ = v124
                                                                }
                                                            }
                                                        }
                                                        var v126 int64 = (v18 % (((v18 - (v18 + (v18 << (((v34 - uint(v35))) & 63))))) | 1))
                                                        _ = fi110
                                                        _ = v126
                                                    }
                                                case -9223372036854775808:
                                                    v127 := v1[32768]
                                                    _ = v127
                                                default:
                                                    var v128 int = 42
                                                    _ = v128
                                                }
                                            }
                                        case "\n\t":
                                            v32 += 982
                                        }
                                    }
                                    var v129 uint32 = 2147483648
                                    _ = v129
                                }
                                _ = v99
                            }
                        } else {
                            var v130 int8 = -71
                            _ = v130
                        }
                    }
                    v131 := GenS12{F0: 32767}
                    fmt.Println(v34, v2)
                    _ = v131
                }
            case "foo bar":
                v19 += -2.25
            case "hello":
                var v132 float64 = 0.1
                _ = v132
            }
            fmt.Println(v18)
            for fi133 := 0; fi133 < 6; fi133++ {
                _ = len(v1)
                v134 := GenS11{F0: (v32 / (((v32 << ((2) & 63))) | 1)), F1: 127}
                _ = fi133
                _ = v134
            }
        } else {
            _ = len(v4)
            if v5 {
                v5 = (2147483647 >= 42)
                var v135 float32 = ((((float32(v2) + ((1e10 + float32(v2)) + v19)) - v19) / -2.25) + v19)
                _ = v135
            } else {
                if v5 {
                    for fi136 := 0; fi136 < 18; fi136++ {
                        v5 = v5
                        v137 := GenS8{F0: 1000, F1: (genFn0() != uint64(v34)), F2: (v34 | (v34 * (^(v34))))}
                        v19 -= v19
                        _ = fi136
                        _ = v137
                    }
                    fmt.Println(v32)
                }
                v18 += v18
                var v138 int32 = v32
                _ = v138
            }
            var v139 int16 = -1
            _ = v139
        }
        var v140 uint8 = uint8(v35)
        if (!true) {
            var v141 uint16 = 32767
            v142 := GenS13{F0: "a"}
            _ = v141
            _ = v142
        }
        _ = v140
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
