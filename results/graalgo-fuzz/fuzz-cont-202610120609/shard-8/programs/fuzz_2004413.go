package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8; F1 uint; F2 int8 }
func genFn0() uint64 {
    v5 := GenS1{F0: 658, F1: -2147483648}
    var v6 uint16 = 0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint64, p1 int16) uint8 {
    p0 = (p0 << ((237) & 63))
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 4; fi8++ {
        if (924 == 0) {
            switch "hello" {
            case "foo bar":
                for fi9 := 0; fi9 < 3; fi9++ {
                    p0 = p0
                    var v10 string = "hello"
                    p0 -= (p0 >> ((13) & 63))
                    _ = fi9
                    _ = v10
                }
            case "a":
                p1 = (p1 ^ 1000)
            }
            fmt.Println(p1, v7)
        }
        v11 := []int{(fi8 / ((fi8) | 1)), fi8, (v7 << ((4294967295) & 63))}
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 128
}

func main() {
    var v1 bool = (!true)
    var v2 int = -2147483648
    var v3 uint16 = 32767
    v4 := make(map[uint32]uint32)
    switch "foo bar" {
    case "hello":
        fmt.Println(v3, v1, v2)
    case "":
        var v12 uint64 = 42
        _ = v12
    }
    v1 = ((!true) || true)
    switch "\n\t" {
    case "\n\t":
        v3 += (v3 >> ((18446744073709551615) & 63))
    default:
        v13 := []int{v2}
        _ = v13
    }
    if v1 {
        for fi14 := 0; fi14 < 19; fi14++ {
            v4[0] = 4294967295
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 float64 = 1e10
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    } else {
        var v17 bool = (464 < 181)
        var v18 int32 = 0
        _ = v17
        _ = v18
    }
    var v19 int32 = 42
    v19 = (v19 ^ (v19 - (v19 | v19)))
    for fi20 := 0; fi20 < 8; fi20++ {
        switch 9223372036854775807 {
        case 2147483647:
            if false {
                if (v1 || (!v1)) {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        switch "a" {
                        case "hello":
                            v22 := []int{}
                            _ = v22
                        default:
                            fmt.Println(fi20)
                        }
                        _ = fi21
                    }
                    v2 = v2
                } else {
                    var v23 int8 = -2
                    var v24 string = "a"
                    _ = v23
                    _ = v24
                }
                v2 -= (fi20 * -81)
                v2 -= fi20
            }
        case 127:
            if true {
                fmt.Println(fi20, v2, v3)
                v25 := GenS1{F0: 127, F1: v2}
                _ = v25
            } else {
                var v26 uint8 = 2
                v3 -= v3
                if v1 {
                    if true {
                        switch "a" {
                        case "a":
                            v27 := GenS1{F0: (v2 - fi20), F1: fi20}
                            _ = v27
                        case "":
                            if ((!("hello" != ("a" + "\n\t"))) && false) {
                                if v1 {
                                    if (!(0 < 18446744073709551615)) {
                                        v28 := GenS2{F0: "", F1: 42}
                                        _ = v28
                                    } else {
                                        v29 := []int{fi20}
                                        var v30 string = "a"
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = len(v4)
                                }
                                if (v2 < 9223372036854775807) {
                                    if (!(!v1)) {
                                        v31 := GenS2{F0: "", F1: 2}
                                        _ = v31
                                    } else {
                                        v32 := []int{(v2 + fi20), fi20, -258, v2}
                                        v19 += (v19 % ((42) | 1))
                                        _ = v32
                                    }
                                    v33 := []int{(-(v2)), v2, fi20, (fi20 / (((fi20 >> ((42) & 63))) | 1))}
                                    _ = v33
                                } else {
                                    if v1 {
                                        fmt.Println(v19, v19)
                                    } else {
                                        v34 := []int{40, (v2 | v2), fi20}
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v19, v19)
                                        default:
                                            for fi35 := 0; fi35 < 7; fi35++ {
                                                var v36 uint = 0
                                                v37 := GenS1{F0: 32767, F1: fi35}
                                                _ = fi35
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                        for fi38 := 0; fi38 < 20; fi38++ {
                                            v1 = true
                                            _ = fi38
                                        }
                                        _ = v34
                                    }
                                    v26 -= 0
                                    for fi39 := 0; fi39 < 18; fi39++ {
                                        var v40 int16 = -861
                                        if (1.5 < 3.14159) {
                                            if (!true) {
                                                var v41 uint16 = (v3 & (v3 * 2))
                                                v42 := GenS3{F0: -2147483648, F1: 42, F2: v3}
                                                switch ("" + "hello") {
                                                case "\n\t":
                                                    var v43 int64 = -523
                                                    _ = v43
                                                }
                                                if v1 {
                                                    var v44 uint64 = 320
                                                    _ = v44
                                                }
                                                _ = v41
                                                _ = v42
                                            } else {
                                                if false {
                                                    var v45 int = (fi39 << ((65535) & 63))
                                                    _ = len(v4)
                                                    _ = v45
                                                } else {
                                                    if (v1 || ((0.0 + (0.1 * (3.14159 / -2.25))) == 0.0)) {
                                                        v46 := GenS4{F0: (v26 <= genFn1(genFn0(), (v40 | v40))), F1: 0, F2: 42}
                                                        _ = v46
                                                    }
                                                }
                                            }
                                            v47 := GenS5{F0: 2, F1: ("" + "hello")}
                                            v48 := GenS5{F0: 18446744073709551615, F1: ("hello" + "\n\t")}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            var v49 bool = (("foo bar" + "\n\t") != "")
                                            _ = v49
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                }
                                v50 := []int{128, (fi20 | fi20)}
                                _ = v50
                            }
                        case "hello":
                            for fi51 := 0; fi51 < 1; fi51++ {
                                _ = len(v4)
                                v1 = ((3.14159 - (0.1 - (1.5 / float64(3.14159)))) > 0.0)
                                delete(v4, 4294967295)
                                _ = fi51
                            }
                        default:
                            for fi52 := 0; fi52 < 4; fi52++ {
                                v2 += (fi20 % ((fi20) | 1))
                                _ = fi52
                            }
                        }
                        var v53 bool = (false && v1)
                        for fi54 := 0; fi54 < 1; fi54++ {
                            if v1 {
                                v26 = (v26 | 1)
                            }
                            _ = fi54
                        }
                        _ = v53
                    }
                    v19 += (-(v19))
                    delete(v4, 0)
                } else {
                    v19 = int32(v26)
                }
                _ = v26
            }
        default:
            for fi55 := 0; fi55 < 4; fi55++ {
                v2 = (fi20 * 9223372036854775807)
                if (2 == 65535) {
                    var v56 float32 = 0.0
                    var v57 float64 = 1.5
                    switch ("\n\t" + "") {
                    case "":
                        v4[65535] = 42
                    case "foo bar":
                        v1 = (v1 || (128 != 23))
                    default:
                        var v58 float64 = ((-2.25 + float64((v56 * 3.14159))) * v57)
                        _ = v58
                    }
                    _ = v56
                    _ = v57
                }
                v19 = v19
                _ = fi55
            }
        }
        switch "" {
        case "":
            v19 = v19
        default:
            switch fi20 {
            case 32768:
                _ = len(v4)
            case -1:
                var v59 uint8 = 255
                _ = v59
            }
        }
        _ = fi20
    }
    var v60 float32 = 3.14159
    switch (("a" + "hello") + ("" + "foo bar")) {
    case "a":
        var v61 float64 = (1e10 + 3.14159)
        _ = v61
    case "hello":
        v62 := GenS4{F0: (!v1), F1: -9223372036854775808, F2: 65535}
        _ = v62
    default:
        for fi63 := 0; fi63 < 4; fi63++ {
            v64 := []int{v2, v2}
            v65 := GenS1{F0: -2147483648, F1: v2}
            var v66 uint32 = 42
            _ = fi63
            _ = v64
            _ = v65
            _ = v66
        }
    }
    if false {
        for fi67 := 0; fi67 < 11; fi67++ {
            if (v1 || (!(-32768 > 347))) {
                switch "" {
                case "":
                    if (-2.25 >= ((1e10 - float64(((v60 + (v60 / 1.0)) + float32(((0.0 * (float64((((float32((0.0 / 1.0)) / v60) + float32(0.1)) + (0.1 / (v60 - float32(3.14159))))) + -2.25)) / -2.25))))) + 0.0)) {
                        for fi68 := 0; fi68 < 8; fi68++ {
                            var v69 int16 = 1000
                            _ = fi68
                            _ = v69
                        }
                    } else {
                        v1 = v1
                    }
                case "a":
                    switch int(v19) {
                    case -1:
                        v70 := GenS2{F0: "hello", F1: uint(v3)}
                        _ = v70
                    }
                case "\n\t":
                    v19 = v19
                default:
                    var v71 int16 = -1
                    _ = v71
                }
                if true {
                    v72 := GenS2{F0: ("foo bar" + (((("hello" + "") + "hello") + "") + "a")), F1: 2}
                    v73 := []int{fi67, (fi67 >> ((1) & 63)), v2, fi67}
                    _ = v72
                    _ = v73
                }
                v74 := []int{-555, fi67, fi67}
                _ = v74
            } else {
                v75 := GenS6{F0: 42}
                switch fi67 {
                case -1:
                    v76 := GenS7{F0: ("hello" + "hello"), F1: 1}
                    _ = v76
                default:
                    switch "hello" {
                    case "foo bar":
                        v2 += v2
                    case "hello":
                        for fi77 := 0; fi77 < 12; fi77++ {
                            var v78 float32 = v60
                            fmt.Println(v60, v19)
                            v4[803] = 0
                            _ = fi77
                            _ = v78
                        }
                    case "a":
                        v2 -= v2
                    default:
                        var v79 uint32 = 4294967295
                        _ = v79
                    }
                }
                _ = v75
            }
            v80 := GenS7{F0: "hello", F1: 2}
            for fi81 := 0; fi81 < 6; fi81++ {
                v2 += -949
                _ = fi81
            }
            _ = fi67
            _ = v80
        }
        var v82 uint16 = (^(v3))
        _ = v82
    } else {
        switch "" {
        case "\n\t":
            v19 = (-(v19))
        case "foo bar":
            switch (v2 * 127) {
            case 32768:
                switch "" {
                case "":
                    var v83 int8 = 42
                    _ = v83
                case "\n\t":
                    v84 := []int{v2, v2, (-(v2)), v2}
                    _ = v84
                case "a":
                    v60 = float32((3.14159 + (0.1 - (1e10 / -2.25))))
                default:
                    v85 := []int{}
                    _ = v85
                }
            default:
                if v1 {
                    var v86 bool = (!true)
                    var v87 uint32 = 837
                    switch "a" {
                    case "a":
                        for fi88 := 0; fi88 < 20; fi88++ {
                            fmt.Println(v60)
                            if v1 {
                                v89 := GenS5{F0: 945, F1: (("foo bar" + "") + ("\n\t" + "a"))}
                                var v90 uint16 = (v3 / (((v3 - (v3 + (v3 ^ (v3 + (v3 << ((1) & 63))))))) | 1))
                                _ = v89
                                _ = v90
                            } else {
                                var v91 int64 = 1
                                _ = v91
                            }
                            switch (fi88 >> ((18446744073709551615) & 63)) {
                            case 2147483647:
                                for fi92 := 0; fi92 < 19; fi92++ {
                                    v3 = (v3 - (^(v3)))
                                    _ = fi92
                                }
                            case -9223372036854775808:
                                for fi93 := 0; fi93 < 14; fi93++ {
                                    var v94 uint32 = (v87 | (v87 | v87))
                                    if (v19 > (-(v19))) {
                                        var v95 uint64 = 42
                                        var v96 string = (("\n\t" + ("a" + ("a" + ("foo bar" + "\n\t")))) + "a")
                                        var v97 string = "a"
                                        var v98 int16 = -1
                                        _ = v95
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                    }
                                    _ = fi93
                                    _ = v94
                                }
                            default:
                                v99 := GenS6{F0: -53}
                                _ = v99
                            }
                            _ = fi88
                        }
                    case "foo bar":
                        v60 -= float32(1.5)
                    default:
                        if (1 >= 100) {
                            fmt.Println(v60, v19)
                        } else {
                            v1 = (42 < 357)
                            var v100 uint16 = (v3 / (((v3 & (v3 % (((v3 + (v3 & (^(v3))))) | 1)))) | 1))
                            v2 = (^(v2))
                            _ = v100
                        }
                    }
                    _ = v86
                    _ = v87
                } else {
                    var v101 uint8 = 127
                    for fi102 := 0; fi102 < 20; fi102++ {
                        v103 := GenS8{F0: 42}
                        v104 := []int{}
                        _ = fi102
                        _ = v103
                        _ = v104
                    }
                    _ = v101
                }
            }
        }
        fmt.Println(v2, v1, v60)
    }
    var v105 int = (v2 * (v2 + v2))
    var v106 string = "a"
    switch (v2 * v105) {
    case -2147483648:
        v3 = 255
    case 0:
        for fi107 := 0; fi107 < 19; fi107++ {
            for fi108 := 0; fi108 < 4; fi108++ {
                switch v106 {
                case "\n\t":
                    var v109 bool = v1
                    _ = v109
                }
                v110 := GenS9{F0: 32767}
                if false {
                    v111 := GenS7{F0: (v106 + v106), F1: 4294967295}
                    var v112 int32 = v19
                    _ = v111
                    _ = v112
                }
                var v113 uint = 65535
                if ((fi107 - -1) < v105) {
                    v114 := GenS10{F0: -42, F1: 42, F2: 0}
                    _ = v114
                } else {
                    v2 += fi108
                    v4[42] = 880
                }
                delete(v4, uint32(v3))
                _ = fi108
                _ = v110
                _ = v113
            }
            _ = fi107
        }
    case 1:
        v106 = v106
    default:
        var v115 uint8 = 42
        _ = v115
    }
    fmt.Println(v105, v60)
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v105
    _ = v106
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
