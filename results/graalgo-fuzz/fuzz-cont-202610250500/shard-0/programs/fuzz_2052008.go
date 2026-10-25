package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int }
type GenS6 struct { F0 uint16; F1 string; F2 int }
type GenS7 struct { F0 int32; F1 string; F2 uint32 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 int8; F1 uint32; F2 bool }
type GenS11 struct { F0 int32; F1 bool; F2 uint64 }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint, p2 int) uint64 {
    p0 = (p0 ^ 907)
    fmt.Println(p1, p2)
    v5 := GenS1{F0: 1, F1: 2, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}
func genFn1() int {
    var v6 uint = 65535
    v7 := []int{42, 127, -2147483648, 9223372036854775807}
    if false {
        v6 -= (v6 % ((42) | 1))
    }
    _ = v6
    _ = v7
    return -305
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{-948}
    v3 := make(map[int]uint32)
    var v4 int32 = -1
    v8 := GenS2{F0: 0, F1: 1000}
    for fi9 := 0; fi9 < 19; fi9++ {
        if false {
            fmt.Println(v4, fi9, v1)
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                if (634 < int64(fi9)) {
                    v1 = v1
                    if true {
                        v11 := GenS3{F0: (2 > 0), F1: 1, F2: v1}
                        if (false && (!((!true) && false))) {
                            v12 := GenS2{F0: 65535, F1: 32767}
                            var v13 uint8 = 0
                            v14 := GenS3{F0: (genFn0(724, 194, genFn1()) < 4294967295), F1: 561, F2: (v1 + "hello")}
                            v15 := GenS1{F0: -1, F1: 18446744073709551615, F2: 660}
                            if (false && false) {
                                v3[fi10] = 42
                                fmt.Println(fi10, fi10)
                            } else {
                                var v16 int64 = 859
                                switch fi9 {
                                case 9223372036854775807:
                                    fmt.Println(v1, v1)
                                default:
                                    v13 = (v13 | (v13 - (v13 * 128)))
                                }
                                _ = len(v3)
                                _ = v16
                            }
                            switch fi10 {
                            case -71:
                                v1 = "foo bar"
                            default:
                                v13 = (v13 & v13)
                            }
                            fmt.Println(v4)
                            var v17 uint32 = 2147483647
                            if true {
                                var v18 uint = 42
                                v1 = v1
                                v13 += (^(v13))
                                _ = v18
                            } else {
                                var v19 int64 = 51
                                _ = v19
                            }
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v17
                        }
                        var v20 bool = (true || ((3.14159 * 0.1) == 1e10))
                        v1 = "hello"
                        _ = v11
                        _ = v20
                    }
                    v21 := GenS3{F0: (!(true || false)), F1: 65535, F2: v1}
                    if true {
                        var v22 string = ("\n\t" + v1)
                        v1 = "a"
                        _ = v22
                    }
                    if ((608 >= 4294967295) || (18446744073709551615 < 283)) {
                        if (!true) {
                            var v23 int8 = -1
                            switch "a" {
                            case "hello":
                                v4 -= v4
                            }
                            _ = v23
                        }
                    }
                    _ = v21
                } else {
                    v1 = v1
                    if false {
                        v1 = v1
                        var v24 uint16 = 2
                        _ = v24
                    } else {
                        if ("" != (v1 + (v1 + v1))) {
                            v25 := GenS4{F0: 0, F1: -591, F2: 42}
                            var v26 int64 = 0
                            var v27 int16 = 233
                            var v28 uint64 = 818
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            switch v1 {
                            case "foo bar":
                                switch (fi10 + genFn1()) {
                                case 127:
                                    v1 = "a"
                                case 128:
                                    switch 2147483647 {
                                    case 0:
                                        _ = len(v3)
                                    case -678:
                                        v29 := GenS2{F0: 65535, F1: -457}
                                        _ = v29
                                    case 133:
                                        v1 = "\n\t"
                                    default:
                                        var v30 int16 = -32768
                                        _ = v30
                                    }
                                default:
                                    switch v1 {
                                    case "hello":
                                        switch (v1 + (v1 + ((v1 + (v1 + ((v1 + (v1 + ("" + ("foo bar" + v1)))) + ("\n\t" + (v1 + v1))))) + (v1 + (v1 + v1))))) {
                                        case "hello":
                                            if (!((((false || ((^(v4)) >= (v4 | v4))) && true) || true) && false)) {
                                                if true {
                                                    switch fi10 {
                                                    case 42:
                                                        var v31 int = (fi10 % (((fi9 | int(v4))) | 1))
                                                        _ = v31
                                                    case 1:
                                                        var v32 uint32 = 1
                                                        _ = v32
                                                    case -912:
                                                        var v33 uint16 = 32767
                                                        _ = v33
                                                    }
                                                    switch v1 {
                                                    case "foo bar":
                                                        var v34 float64 = 3.14159
                                                        _ = v34
                                                    case "hello":
                                                        v35 := GenS1{F0: 88, F1: 4294967295, F2: 65535}
                                                        _ = v35
                                                    case "":
                                                        delete(v3, fi10)
                                                    }
                                                    v4 -= v4
                                                    v1 = v1
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v36 string = v1
                                                        _ = v36
                                                    case "foo bar":
                                                        var v37 uint32 = 0
                                                        _ = v37
                                                    case "":
                                                        var v38 int32 = v4
                                                        _ = v38
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                } else {
                                                    v3[fi10] = 1
                                                    v4 = v4
                                                    switch genFn1() {
                                                    case -904:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case 5:
                                                        switch "hello" {
                                                        case "a":
                                                            var v39 uint8 = 127
                                                            _ = v39
                                                        case "":
                                                            var v40 uint = 340
                                                            _ = v40
                                                        case "foo bar":
                                                            if true {
                                                                v4 -= -2147483648
                                                                v1 = "hello"
                                                                v41 := GenS1{F0: 42, F1: genFn0(255, 0, (fi10 % ((genFn1()) | 1))), F2: 4294967295}
                                                                _ = v41
                                                            }
                                                        default:
                                                            fmt.Println(fi10)
                                                        }
                                                    case 2147483647:
                                                        v1 = ("a" + v1)
                                                        v42 := GenS5{F0: 4294967295, F1: 65535, F2: fi10}
                                                        _ = v42
                                                    }
                                                }
                                            } else {
                                                v4 -= (v4 ^ (v4 | (v4 * 23)))
                                                var v43 float64 = 1e10
                                                if false {
                                                    v44 := GenS3{F0: (!true), F1: 4294967295, F2: (("a" + (v1 + v1)) + ("\n\t" + v1))}
                                                    fmt.Println(v4, fi10)
                                                    var v45 float64 = 0.0
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    v46 := GenS2{F0: 42, F1: -32768}
                                                    _ = v46
                                                }
                                                _ = v43
                                            }
                                        case "":
                                            switch -1 {
                                            case 128:
                                                var v47 float32 = 3.14159
                                                _ = v47
                                            case 127:
                                                fmt.Println(v1)
                                            }
                                        case "foo bar":
                                            var v48 string = v1
                                            _ = v48
                                        }
                                    default:
                                        var v49 float64 = float64((1.5 - 0.1))
                                        _ = v49
                                    }
                                }
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = genFn1()
                                }
                            default:
                                if len(v2) > 0 {
                                    v2[0] = fi10
                                }
                            }
                            v1 = "a"
                            v4 = (v4 * (v4 >> ((0) & 63)))
                        }
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                        var v50 uint64 = 18446744073709551615
                        _ = v50
                    }
                    if (!((224 < 1) || false)) {
                        var v51 int64 = 42
                        v52 := GenS2{F0: 1, F1: 42}
                        _ = v51
                        _ = v52
                    }
                }
                var v53 string = ("foo bar" + v1)
                if ((fi9 | 2147483647) > genFn1()) {
                    v53 = (("\n\t" + (v1 + "foo bar")) + ("a" + "a"))
                    v1 = "hello"
                    _ = len(v2)
                }
                _ = fi10
                _ = v53
            }
        }
        v4 -= -926
        switch (fi9 * fi9) {
        case 4294967295:
            var v54 uint32 = 65535
            _ = v54
        default:
            fmt.Println(v1, v4, v1)
        }
        _ = fi9
    }
    var v55 float32 = 1.5
    for fi56 := 0; fi56 < 10; fi56++ {
        var v57 int8 = 127
        v3[(fi56 * (-(fi56)))] = 2147483647
        v58 := v3[genFn1()]
        _ = fi56
        _ = v57
        _ = v58
    }
    var v59 bool = true
    var v60 float32 = 1.5
    if v59 {
        v60 = (v55 * v60)
        if true {
            var v61 float64 = float64(-2.25)
            _ = v61
        } else {
            switch 42 {
            case 0:
                v59 = v59
            case -1:
                var v62 float32 = -2.25
                _ = v62
            }
            if (690 >= 9223372036854775807) {
                for fi63 := 0; fi63 < 3; fi63++ {
                    if false {
                        v55 -= ((v60 + ((v60 - v60) - v55)) / v55)
                        var v64 uint16 = 1
                        var v65 uint16 = v64
                        _ = v64
                        _ = v65
                    }
                    _ = fi63
                }
                fmt.Println(v1, v4, v55)
                switch 1 {
                case 128:
                    if ((^(v4)) == (v4 << ((0) & 63))) {
                        switch v1 {
                        case "foo bar":
                            switch genFn1() {
                            case 4294967295:
                                if (true || true) {
                                    v1 = v1
                                }
                            default:
                                if v59 {
                                    for fi66 := 0; fi66 < 9; fi66++ {
                                        var v67 uint16 = 32767
                                        _ = fi66
                                        _ = v67
                                    }
                                    for fi68 := 0; fi68 < 5; fi68++ {
                                        v59 = (0 <= 0)
                                        var v69 string = "hello"
                                        var v70 string = (("a" + v1) + (v1 + v1))
                                        _ = fi68
                                        _ = v69
                                        _ = v70
                                    }
                                    _ = len(v2)
                                } else {
                                    if false {
                                        fmt.Println(v55, v59)
                                        if ((!(2 > 0)) || false) {
                                            v71 := v3[9223372036854775807]
                                            _ = v71
                                        }
                                    } else {
                                        for fi72 := 0; fi72 < 11; fi72++ {
                                            v59 = (true && v59)
                                            _ = fi72
                                        }
                                        for fi73 := 0; fi73 < 16; fi73++ {
                                            fmt.Println(v1)
                                            _ = fi73
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = len(v3)
                                    v4 += (v4 << ((0) & 63))
                                }
                            }
                        case "\n\t":
                            var v74 int8 = 42
                            _ = v74
                        case "hello":
                            v59 = v59
                        }
                        for fi75 := 0; fi75 < 9; fi75++ {
                            switch v1 {
                            case "hello":
                                for fi76 := 0; fi76 < 11; fi76++ {
                                    v77 := GenS2{F0: 4294967295, F1: -1}
                                    var v78 int32 = -2147483648
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                }
                            case "":
                                if false {
                                    var v79 uint32 = 4294967295
                                    switch (fi75 | -1) {
                                    case 127:
                                        v80 := GenS2{F0: 1, F1: 91}
                                        _ = v80
                                    case -768:
                                        v3[fi75] = v79
                                    }
                                    _ = v79
                                } else {
                                    if false {
                                        if (!false) {
                                            if (fi75 != (fi75 << ((4294967295) & 63))) {
                                                var v81 string = ""
                                                v81 = ""
                                                v82 := GenS6{F0: 2, F1: ((("foo bar" + v1) + (v81 + "a")) + v81), F2: 127}
                                                _ = v81
                                                _ = v82
                                            } else {
                                                _ = len(v3)
                                                v2 = append(v2, (fi75 + (fi75 * (fi75 % ((fi75) | 1)))))
                                                v59 = ("a" > v1)
                                            }
                                            v2 = append(v2, fi75)
                                        }
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v55 = 3.14159
                                        if false {
                                            delete(v3, int(v4))
                                            if v59 {
                                                for fi83 := 0; fi83 < 14; fi83++ {
                                                    v1 = v1
                                                    _ = fi83
                                                }
                                                if (755 >= 744) {
                                                    if len(v2) > 0 {
                                                        v2[0] = fi75
                                                    }
                                                }
                                            }
                                            v84 := GenS7{F0: v4, F1: v1, F2: 62}
                                            _ = v84
                                        } else {
                                            var v85 uint32 = 42
                                            switch (fi75 % ((fi75) | 1)) {
                                            case -254:
                                                if ((-(fi75)) >= (fi75 - (fi75 & (fi75 & (fi75 / (((fi75 + (fi75 & fi75))) | 1)))))) {
                                                    v59 = (true || v59)
                                                    if v59 {
                                                        fmt.Println(fi75, v59)
                                                        if false {
                                                            v55 = v55
                                                            var v86 uint16 = 42
                                                            _ = v86
                                                        }
                                                        fmt.Println(v1, fi75, v85)
                                                    }
                                                    for fi87 := 0; fi87 < 15; fi87++ {
                                                        v88 := GenS8{F0: 0, F1: ((v1 + ("foo bar" + ("a" + v1))) + v1)}
                                                        var v89 float32 = (v55 / (v60 / 1e10))
                                                        _ = fi87
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                } else {
                                                    var v90 uint64 = 638
                                                    var v91 float32 = (v60 + (v60 * 1e10))
                                                    v92 := GenS9{F0: 42, F1: 32767}
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                }
                                            }
                                            for fi93 := 0; fi93 < 15; fi93++ {
                                                v4 += v4
                                                _ = fi93
                                            }
                                            _ = v85
                                        }
                                    }
                                }
                            }
                            fmt.Println(fi75, v55, v1)
                            _ = fi75
                        }
                    } else {
                        v94 := GenS10{F0: 100, F1: 0, F2: v59}
                        v95 := GenS4{F0: 32767, F1: 0, F2: 0}
                        _ = v94
                        _ = v95
                    }
                case 4294967295:
                    fmt.Println(v60)
                default:
                    v59 = true
                }
            } else {
                var v96 int32 = v4
                if true {
                    for fi97 := 0; fi97 < 3; fi97++ {
                        for fi98 := 0; fi98 < 11; fi98++ {
                            var v99 int = fi97
                            _ = fi98
                            _ = v99
                        }
                        _ = fi97
                    }
                    var v100 uint8 = 34
                    v101 := GenS11{F0: (v96 >> ((905) & 63)), F1: v59, F2: genFn0(255, 0, -9223372036854775808)}
                    _ = v100
                    _ = v101
                } else {
                    var v102 bool = false
                    _ = v102
                }
                _ = v96
            }
        }
        if (!(99 >= -1)) {
            fmt.Println(v1, v60)
        }
    }
    v55 -= 0.1
    for fi103 := 0; fi103 < 13; fi103++ {
        v4 -= (v4 - (-(v4)))
        if (true && (v59 && true)) {
            var v104 string = (("" + "a") + "a")
            v104 = ""
            _ = v104
        }
        _ = fi103
    }
    if v59 {
        v60 += -2.25
        if true {
            var v105 int8 = int8(v4)
            fmt.Println(v105)
            _ = v105
        } else {
            var v106 int8 = 100
            var v107 int32 = v4
            _ = v106
            _ = v107
        }
    }
    for fi108 := 0; fi108 < 3; fi108++ {
        if v59 {
            v109 := GenS3{F0: (1 > 590), F1: 1, F2: "hello"}
            v110 := GenS9{F0: 1000, F1: 2}
            v111 := v3[fi108]
            _ = v109
            _ = v110
            _ = v111
        }
        if v59 {
            v55 -= (float32(float64(1.5)) / v55)
            var v112 uint64 = 42
            _ = v112
        }
        switch (fi108 + fi108) {
        case 32767:
            v55 += v55
        case 128:
            if v59 {
                v4 = v4
            }
        case -2147483648:
            for fi113 := 0; fi113 < 7; fi113++ {
                switch -289 {
                case -1:
                    v60 = (float32((0.0 * 3.14159)) * 0.1)
                case 883:
                    var v114 bool = (!(v59 || (genFn0(32767, 255, (fi113 & 127)) == 42)))
                    _ = v114
                case 9223372036854775807:
                    switch (-(fi108)) {
                    case 1:
                        if true {
                            v115 := v3[fi108]
                            var v116 uint = 2
                            _ = v115
                            _ = v116
                        } else {
                            var v117 uint32 = 267
                            v118 := GenS11{F0: 1, F1: false, F2: 175}
                            delete(v3, -361)
                            _ = v117
                            _ = v118
                        }
                    case 42:
                        v119 := GenS12{F0: 2}
                        _ = v119
                    }
                default:
                    v60 += v55
                }
                _ = fi113
            }
        }
        _ = fi108
    }
    v59 = v59
    var v120 string = (v1 + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v59
    _ = v60
    _ = v8
}
