package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int8; F1 uint; F2 uint8 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int64; F2 int }
type GenS9 struct { F0 uint32; F1 int64; F2 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS12 struct { F0 uint16; F1 int16 }
type GenS13 struct { F0 uint; F1 int8; F2 uint8 }
type GenS14 struct { F0 bool; F1 int }
func genFn0(p0 uint32) uint16 {
    if (0.1 >= 0.0) {
        v7 := []int{2147483647, 4294967295}
        switch "\n\t" {
        case "hello":
            if true {
                var v8 int16 = -1
                switch -1 {
                case 32768:
                    var v9 string = "foo bar"
                    _ = v9
                case -598:
                    if (42 != 65535) {
                        fmt.Println(p0, v8)
                        p0 = 2147483647
                        v10 := GenS1{F0: 43, F1: "foo bar"}
                        _ = v10
                    }
                case 32767:
                    var v11 float32 = float32((-2.25 * 0.0))
                    _ = v11
                default:
                    var v12 float64 = (float64(0.1) - 1.5)
                    _ = v12
                }
                _ = v8
            } else {
                fmt.Println(p0)
                for fi13 := 0; fi13 < 19; fi13++ {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS1{F0: 68, F1: "hello"}
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = fi14
                        _ = v15
                    }
                    var v16 int16 = 988
                    v16 += 873
                    _ = fi13
                    _ = v16
                }
            }
        default:
            var v17 float64 = -2.25
            _ = v17
        }
        _ = v7
    } else {
        p0 = p0
        p0 += 2
        v18 := GenS2{F0: -1}
        _ = v18
    }
    v19 := GenS1{F0: 100, F1: "\n\t"}
    fmt.Println(p0)
    _ = p0
    _ = v19
    return 1
}
func genFn1(p0 string, p1 uint16, p2 int32) uint32 {
    fmt.Println(p2, p1)
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return 650
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 42
    v3 := []uint32{164, uint32(v2)}
    var v4 bool = (72 != 65535)
    v5 := make(map[int]uint32)
    var v6 int32 = 715
    if v4 {
        switch "\n\t" {
        case "\n\t":
            v4 = (18446744073709551615 >= 0)
        case "":
            for fi20 := 0; fi20 < 2; fi20++ {
                v4 = v4
                var v21 uint16 = 42
                v3 = append(v3, 65535)
                _ = fi20
                _ = v21
            }
        case "a":
            if (("hello" <= ("hello" + "")) || true) {
                var v22 uint8 = 69
                _ = v22
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    v4 = true
                    for fi24 := 0; fi24 < 4; fi24++ {
                        if (true && true) {
                            v2 = v2
                            switch ("foo bar" + "") {
                            case "a":
                                if (32767 == -227) {
                                    v3 = append(v3, 65535)
                                    _ = len(v3)
                                } else {
                                    v6 -= -695
                                }
                            default:
                                v2 = v2
                            }
                            v1 += v1
                        } else {
                            var v25 uint32 = 0
                            _ = v25
                        }
                        fmt.Println(v1, v2, v1)
                        v4 = ((fi23 & (fi24 ^ fi23)) == (^(fi24)))
                        _ = fi24
                    }
                    v4 = (true && (!v4))
                    _ = fi23
                }
                fmt.Println(v1, v1, v1)
                var v26 float32 = ((1e10 / 3.14159) + float32(1e10))
                _ = v26
            }
        }
    } else {
        _ = len(v5)
        switch ("a" + ((("a" + "foo bar") + "foo bar") + "")) {
        case "a":
            v4 = (-740 <= 192)
        case "foo bar":
            _ = len(v5)
        case "hello":
            if true {
                fmt.Println(v6, v4)
            } else {
                if v4 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v27 uint64 = v2
                    _ = v27
                } else {
                    if false {
                        for fi28 := 0; fi28 < 15; fi28++ {
                            var v29 int = fi28
                            _ = fi28
                            _ = v29
                        }
                    }
                }
                if false {
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v31 := GenS1{F0: 30, F1: "a"}
                        v32 := GenS2{F0: int16(v2)}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                } else {
                    var v33 uint64 = (^(v2))
                    _ = v33
                }
            }
        default:
            for fi34 := 0; fi34 < 16; fi34++ {
                v6 = -27
                _ = fi34
            }
        }
    }
    if (!(9223372036854775807 >= 42)) {
        v35 := GenS2{F0: 32767}
        _ = v35
    }
    switch (("hello" + "\n\t") + "hello") {
    case "\n\t":
        delete(v5, 128)
    default:
        var v36 int32 = v6
        _ = v36
    }
    if (4294967295 > 65535) {
        if false {
            var v37 uint32 = 2
            if false {
                if v4 {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        switch fi38 {
                        case 246:
                            v5[fi38] = (v37 - v37)
                        case 27:
                            v39 := GenS3{F0: "a", F1: 470}
                            _ = v39
                        }
                        switch (fi38 << ((65535) & 63)) {
                        case 128:
                            if true {
                                v40 := GenS4{F0: 367, F1: "foo bar"}
                                _ = v40
                            } else {
                                v1 = v1
                                for fi41 := 0; fi41 < 19; fi41++ {
                                    switch "hello" {
                                    case "":
                                        if (!(255 >= 119)) {
                                            v42 := v5[-464]
                                            _ = v42
                                        } else {
                                            _ = len(v3)
                                            switch "a" {
                                            case "a":
                                                if len(v3) > 0 {
                                                    v3[0] = genFn1("", genFn0(v37), 2147483647)
                                                }
                                            case "hello":
                                                v1 = (v1 / 3.14159)
                                            case "":
                                                if v4 {
                                                    _ = len(v5)
                                                }
                                                v37 -= (v37 + v37)
                                            }
                                            fmt.Println(v6, fi38, v1)
                                        }
                                    case "a":
                                        fmt.Println(v37, v2, v6)
                                    case "hello":
                                        v43 := GenS4{F0: 32767, F1: ("\n\t" + ((("a" + "a") + (("a" + "hello") + "foo bar")) + "hello"))}
                                        _ = v43
                                    }
                                    switch "foo bar" {
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "hello":
                                        var v44 uint = 0
                                        _ = v44
                                    case "foo bar":
                                        v45 := GenS4{F0: 579, F1: ((("foo bar" + "") + "\n\t") + ((("" + "hello") + "hello") + ""))}
                                        _ = v45
                                    default:
                                        var v46 uint8 = 42
                                        _ = v46
                                    }
                                    switch (fi41 / (((fi41 + (fi41 - fi38))) | 1)) {
                                    case 128:
                                        switch "" {
                                        case "foo bar":
                                            if (42 != uint16(v2)) {
                                                v6 += (^(v6))
                                                var v47 float64 = 3.14159
                                                _ = v47
                                            } else {
                                                fmt.Println(fi38)
                                                v2 += v2
                                            }
                                        }
                                    case 953:
                                        v48 := GenS5{F0: 100, F1: 18446744073709551615, F2: 42}
                                        _ = v48
                                    case 4294967295:
                                        if v4 {
                                            if (!(3.14159 == (-2.25 / 1.5))) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch (fi41 / (((fi38 ^ fi41)) | 1)) {
                                                    case -2147483648:
                                                        fmt.Println(v4, v6)
                                                    case -1:
                                                        v49 := v5[759]
                                                        var v50 uint64 = v2
                                                        _ = v49
                                                        _ = v50
                                                    case 1:
                                                        fmt.Println(v2, v2)
                                                    }
                                                case "a":
                                                    var v51 int64 = 1
                                                    v2 += 0
                                                    _ = v51
                                                case "\n\t":
                                                    var v52 float64 = 1e10
                                                    _ = v52
                                                }
                                                v5[854] = v37
                                            }
                                            fmt.Println(v2, v6, v37)
                                            var v53 int64 = 365
                                            _ = v53
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = (v37 * (v37 - v37))
                                            }
                                            var v54 int32 = v6
                                            _ = v54
                                        }
                                    default:
                                        var v55 int64 = 95
                                        _ = v55
                                    }
                                    _ = fi41
                                }
                                if v4 {
                                    for fi56 := 0; fi56 < 3; fi56++ {
                                        if (uint16(v37) > 42) {
                                            v57 := GenS2{F0: -1}
                                            _ = v57
                                        } else {
                                            if v4 {
                                                fmt.Println(fi56)
                                                var v58 string = "a"
                                                _ = v58
                                            } else {
                                                v1 = -2.25
                                                if (((v37 - 42) <= (v37 ^ (v37 | (^(v37))))) && (true || (-1 <= 741))) {
                                                    var v59 int32 = 64
                                                    _ = v59
                                                }
                                                var v60 string = "hello"
                                                _ = v60
                                            }
                                        }
                                        _ = len(v3)
                                        _ = fi56
                                    }
                                }
                            }
                        case -2147483648:
                            var v61 uint64 = (^(v2))
                            _ = v61
                        default:
                            for fi62 := 0; fi62 < 6; fi62++ {
                                if true {
                                    switch "" {
                                    case "a":
                                        fmt.Println(v37, v4)
                                    case "":
                                        fmt.Println(v6, v2, v4)
                                    case "hello":
                                        fmt.Println(v37, fi38)
                                    default:
                                        var v63 string = ""
                                        _ = v63
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if true {
                                        fmt.Println(v37)
                                        switch (-(fi38)) {
                                        case -9223372036854775808:
                                            var v64 float64 = ((-2.25 - float64(v1)) / 3.14159)
                                            _ = v64
                                        default:
                                            if true {
                                                fmt.Println(v37)
                                                var v65 uint16 = 861
                                                var v66 uint = 65535
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                        fmt.Println(fi62)
                                    } else {
                                        if true {
                                            _ = len(v3)
                                        }
                                        if ((v37 + 0) >= (^(v37))) {
                                            v67 := GenS4{F0: 0, F1: "\n\t"}
                                            v4 = false
                                            fmt.Println(fi38, fi62, v4)
                                            _ = v67
                                        } else {
                                            v2 += (v2 + (v2 / (((v2 * v2)) | 1)))
                                        }
                                        var v68 float32 = (3.14159 + v1)
                                        _ = v68
                                    }
                                    var v69 int16 = 42
                                    var v70 int8 = 42
                                    fmt.Println(fi38, fi62)
                                    v2 = v2
                                    if (v4 && v4) {
                                        if (v4 && v4) {
                                            _ = len(v3)
                                        } else {
                                            v37 -= v37
                                            var v71 string = "a"
                                            var v72 string = (v71 + v71)
                                            var v73 uint8 = 2
                                            if (0.0 == 3.14159) {
                                                if (!false) {
                                                    v4 = v4
                                                    switch (-(fi38)) {
                                                    case 127:
                                                        fmt.Println(v72, fi38, v73)
                                                    case 9223372036854775807:
                                                        v74 := GenS1{F0: 91, F1: ""}
                                                        _ = v74
                                                    case 128:
                                                        if len(v3) > 0 {
                                                            v3[0] = v37
                                                        }
                                                    default:
                                                        v75 := GenS3{F0: ((v72 + v71) + ("" + v72)), F1: (v69 + v69)}
                                                        _ = v75
                                                    }
                                                    v76 := GenS5{F0: int8(v37), F1: 255, F2: (v73 + uint8(v2))}
                                                    delete(v5, (-(fi38)))
                                                    _ = len(v5)
                                                    _ = v76
                                                }
                                            }
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        }
                                        var v77 uint = 42
                                        v78 := GenS6{F0: -2147483648, F1: v70}
                                        _ = v77
                                        _ = v78
                                    } else {
                                        var v79 uint64 = (v2 - (v2 ^ v2))
                                        switch ("foo bar" + ("" + "")) {
                                        case "a":
                                            v80 := GenS5{F0: -1, F1: 1, F2: 255}
                                            var v81 uint8 = 128
                                            if (!false) {
                                                if (-1 >= 0) {
                                                    v70 += 100
                                                } else {
                                                    var v82 uint16 = 42
                                                    _ = len(v3)
                                                    var v83 string = "hello"
                                                    var v84 int8 = int8(v6)
                                                    v69 -= -1
                                                    fmt.Println(v79)
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                }
                                                var v85 float32 = (0.0 - v1)
                                                _ = v85
                                            } else {
                                                if v4 {
                                                    fmt.Println(v1, v4, v81)
                                                    fmt.Println(fi38, v4, fi38)
                                                }
                                                var v86 uint64 = (v2 & (^(v79)))
                                                v86 = (v86 / (((^(v2))) | 1))
                                                v69 = 42
                                                _ = v86
                                            }
                                            _ = v80
                                            _ = v81
                                        case "foo bar":
                                            if v4 {
                                                var v87 int = (-(fi62))
                                                switch "" {
                                                case "\n\t":
                                                    switch fi62 {
                                                    case 1:
                                                        var v88 uint16 = 2
                                                        _ = v88
                                                    case 774:
                                                        fmt.Println(fi38, v37, v6)
                                                    default:
                                                        if len(v3) > 0 {
                                                            v3[0] = (v37 * genFn1("foo bar", genFn0(4294967295), (v6 * 689)))
                                                        }
                                                    }
                                                default:
                                                    switch fi38 {
                                                    case 127:
                                                        fmt.Println(v70, v6)
                                                    case 206:
                                                        fmt.Println(v1, fi62)
                                                    case 32768:
                                                        var v89 int = -1
                                                        _ = v89
                                                    default:
                                                        delete(v5, fi62)
                                                    }
                                                    v70 -= 1
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v90 string = ("\n\t" + ("\n\t" + "a"))
                                                        _ = v90
                                                    case "a":
                                                        var v91 float32 = float32(-2.25)
                                                        _ = v91
                                                    }
                                                }
                                                _ = v87
                                            }
                                        case "":
                                            switch "foo bar" {
                                            case "foo bar":
                                                v92 := GenS7{F0: v70}
                                                _ = v92
                                            case "":
                                                var v93 int = 9223372036854775807
                                                _ = v93
                                            default:
                                                switch "\n\t" {
                                                case "":
                                                    v1 += 0.1
                                                case "foo bar":
                                                    v2 = v2
                                                }
                                            }
                                        default:
                                            v6 += (v6 << ((11) & 63))
                                        }
                                        _ = v79
                                    }
                                    _ = v69
                                    _ = v70
                                } else {
                                    if false {
                                        v1 -= v1
                                        v4 = ((v1 == float32((1.5 + 0.0))) && (1 >= -359))
                                    } else {
                                        v4 = false
                                    }
                                    if true {
                                        v2 -= (v2 - 1)
                                    } else {
                                        if true {
                                            fmt.Println(v6, v6)
                                        } else {
                                            v94 := v5[(fi38 | 4294967295)]
                                            var v95 int = 946
                                            _ = v94
                                            _ = v95
                                        }
                                        if (!true) {
                                            v2 = (v2 - v2)
                                            var v96 uint64 = v2
                                            v97 := GenS8{F0: 209, F1: 42, F2: (fi62 + fi38)}
                                            _ = v96
                                            _ = v97
                                        }
                                        v1 = (float32(0.0) * 1.5)
                                    }
                                }
                                if false {
                                    if true {
                                        var v98 uint8 = 255
                                        v4 = v4
                                        v99 := GenS6{F0: (-(fi62)), F1: -1}
                                        v37 += (v37 % (((v37 * (v37 - 295))) | 1))
                                        _ = v98
                                        _ = v99
                                    } else {
                                        if v4 {
                                            var v100 bool = v4
                                            v101 := GenS9{F0: (v37 >> ((278) & 63)), F1: -957, F2: (^(v6))}
                                            fmt.Println(v4, v1)
                                            _ = v100
                                            _ = v101
                                        }
                                    }
                                    fmt.Println(v4)
                                    v37 = 565
                                } else {
                                    v3 = append(v3, v37)
                                }
                                var v102 int16 = 1
                                switch "\n\t" {
                                case "foo bar":
                                    v102 = (v102 >> ((2) & 63))
                                case "hello":
                                    v103 := GenS10{F0: -271}
                                    _ = v103
                                default:
                                    var v104 int64 = 937
                                    _ = v104
                                }
                                var v105 int = fi62
                                _ = fi62
                                _ = v102
                                _ = v105
                            }
                        }
                        v106 := GenS1{F0: -128, F1: ("a" + ((("foo bar" + "a") + "foo bar") + ""))}
                        _ = fi38
                        _ = v106
                    }
                    var v107 string = ("a" + ("foo bar" + "a"))
                    if true {
                        var v108 int32 = (v6 / ((0) | 1))
                        for fi109 := 0; fi109 < 20; fi109++ {
                            for fi110 := 0; fi110 < 2; fi110++ {
                                if (genFn1("\n\t", 65535, (^(v6))) >= (^(v37))) {
                                    v107 = ("hello" + v107)
                                    if true {
                                        var v111 uint32 = 0
                                        switch fi109 {
                                        case -344:
                                            var v112 int32 = v108
                                            _ = v112
                                        case -9223372036854775808:
                                            v3 = append(v3, v111)
                                            v107 = v107
                                        case -887:
                                            var v113 int = int(v108)
                                            _ = v113
                                        default:
                                            var v114 int32 = (v108 * (v108 + v6))
                                            _ = v114
                                        }
                                        _ = v111
                                    }
                                } else {
                                    var v115 int8 = 1
                                    v2 -= v2
                                    v5[(^(fi109))] = 405
                                    _ = v115
                                }
                                v108 += v108
                                v37 = v37
                                _ = fi110
                            }
                            v2 = v2
                            if len(v3) > 0 {
                                v3[0] = v37
                            }
                            _ = fi109
                        }
                        var v116 int8 = 127
                        _ = v108
                        _ = v116
                    } else {
                        for fi117 := 0; fi117 < 1; fi117++ {
                            var v118 uint16 = 255
                            v119 := GenS5{F0: -128, F1: 377, F2: 227}
                            _ = fi117
                            _ = v118
                            _ = v119
                        }
                        fmt.Println(v37, v2, v37)
                        v107 = v107
                    }
                    _ = v107
                } else {
                    fmt.Println(v1, v2)
                }
                v37 = genFn1("\n\t", genFn0((v37 ^ uint32(v2))), -1)
            } else {
                switch "hello" {
                case "":
                    v2 = v2
                case "\n\t":
                    var v120 float32 = float32(1.5)
                    _ = v120
                default:
                    for fi121 := 0; fi121 < 9; fi121++ {
                        var v122 uint32 = genFn1("foo bar", 255, (v6 - v6))
                        var v123 int = 32767
                        v124 := GenS5{F0: 22, F1: 18446744073709551615, F2: 255}
                        _ = fi121
                        _ = v122
                        _ = v123
                        _ = v124
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                v125 := GenS7{F0: -1}
                _ = v125
            }
            v126 := GenS4{F0: genFn0((v37 + (v37 / (((v37 % ((v37) | 1))) | 1)))), F1: "foo bar"}
            _ = v126
            _ = v37
        }
        var v127 uint32 = genFn1("", 255, v6)
        _ = v127
    }
    switch -801 {
    case 127:
        fmt.Println(v1)
    case 42:
        switch -1 {
        case -2147483648:
            v1 += (v1 - v1)
        case 32767:
            fmt.Println(v6)
        default:
            v128 := GenS11{F0: genFn1((("a" + (("\n\t" + (("\n\t" + "foo bar") + (("\n\t" + "\n\t") + "foo bar"))) + "foo bar")) + "hello"), 65535, -2147483648), F1: -9, F2: 65535}
            _ = v128
        }
    default:
        switch "foo bar" {
        case "foo bar":
            v129 := GenS12{F0: 1000, F1: 0}
            _ = v129
        case "\n\t":
            fmt.Println(v4, v4)
        case "":
            v130 := GenS6{F0: -9223372036854775808, F1: 0}
            _ = v130
        default:
            switch 128 {
            case 2147483647:
                var v131 uint32 = 2
                _ = v131
            case -9223372036854775808:
                v3 = append(v3, 0)
            default:
                if (float32((3.14159 * float64(float32(float64(3.14159))))) == float32(0.0)) {
                    v132 := GenS12{F0: 0, F1: -32768}
                    if true {
                        v133 := GenS13{F0: 255, F1: 38, F2: 127}
                        _ = v133
                    }
                    for fi134 := 0; fi134 < 4; fi134++ {
                        v135 := GenS14{F0: v4, F1: (fi134 >> ((220) & 63))}
                        if v4 {
                            switch ((("" + "") + (("" + ((("foo bar" + ("\n\t" + "foo bar")) + ("" + "foo bar")) + "hello")) + "foo bar")) + "\n\t") {
                            case "a":
                                fmt.Println(v4, v4)
                            }
                            fmt.Println(v1, fi134, v2)
                            var v136 uint16 = 0
                            _ = v136
                        } else {
                            v5[(fi134 >> ((1) & 63))] = 1
                        }
                        v4 = ((9223372036854775807 == -9223372036854775808) && v4)
                        _ = fi134
                        _ = v135
                    }
                    _ = v132
                }
            }
        }
    }
    if (v4 && true) {
        var v137 int = 32768
        _ = v137
    }
    for fi138 := 0; fi138 < 6; fi138++ {
        v139 := GenS11{F0: uint32(v2), F1: -1, F2: 0}
        v4 = false
        _ = fi138
        _ = v139
    }
    v4 = (v4 && v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
