package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint32; F1 string; F2 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 uint8; F2 int }
type GenS11 struct { F0 int32; F1 string }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 int32; F1 uint }
type GenS14 struct { F0 int8; F1 uint16; F2 string }
type GenS15 struct { F0 uint; F1 int8; F2 uint }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int16; F1 uint8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 string; F1 int8; F2 uint }
type GenS21 struct { F0 int; F1 uint32; F2 int32 }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 uint32; F1 string }
type GenS25 struct { F0 uint32; F1 uint8 }
type GenS26 struct { F0 uint; F1 int8; F2 int }
type GenS27 struct { F0 int }
type GenS28 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS29 struct { F0 bool; F1 int32; F2 int16 }
type GenS30 struct { F0 bool }
type GenS31 struct { F0 string; F1 uint64; F2 int32 }
type GenS32 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS33 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64, p1 uint8, p2 bool) bool {
    if p2 {
        v7 := []int{-9223372036854775808, 1, -1, 127}
        if (p2 && ("hello" == "hello")) {
            var v8 uint16 = 0
            p0 = p0
            if false {
                if (!(p0 <= 9223372036854775807)) {
                    fmt.Println(v8, p1, v8)
                    if true {
                        if true {
                            p0 = (p0 ^ (p0 * (p0 + 0)))
                            var v9 uint = 0
                            if (!p2) {
                                v10 := GenS1{F0: v9}
                                var v11 bool = false
                                var v12 uint64 = 234
                                _ = v10
                                _ = v11
                                _ = v12
                            }
                            _ = v9
                        }
                        var v13 float32 = ((float32(1.5) - (((0.0 + -2.25) * 3.14159) / (((0.1 * ((0.0 - 1e10) - float32(3.14159))) / -2.25) * ((3.14159 + ((0.1 / 1.5) + float32(float64(0.1)))) - 1e10)))) / -2.25)
                        var v14 uint = 934
                        _ = v13
                        _ = v14
                    } else {
                        if (1 == 0) {
                            switch 127 {
                            case 4294967295:
                                for fi15 := 0; fi15 < 7; fi15++ {
                                    fmt.Println(p0, fi15)
                                    _ = fi15
                                }
                            case -2147483648:
                                switch ("a" + "hello") {
                                case "\n\t":
                                    for fi16 := 0; fi16 < 1; fi16++ {
                                        v17 := GenS1{F0: 642}
                                        for fi18 := 0; fi18 < 4; fi18++ {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            _ = fi18
                                        }
                                        _ = len(v7)
                                        _ = fi16
                                        _ = v17
                                    }
                                case "foo bar":
                                    var v19 int32 = -1
                                    _ = v19
                                default:
                                    if p2 {
                                        var v20 int = 42
                                        _ = v20
                                    }
                                }
                            case 0:
                                var v21 int32 = 0
                                _ = v21
                            default:
                                fmt.Println(p0)
                            }
                            _ = len(v7)
                        }
                        var v22 string = "\n\t"
                        _ = v22
                    }
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if false {
                    var v23 bool = false
                    p1 -= 1
                    _ = v23
                }
            }
            _ = v8
        }
        _ = v7
    } else {
        if p2 {
            var v24 uint32 = 206
            fmt.Println(p2, p1, p0)
            _ = v24
        }
    }
    switch -355 {
    case 127:
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS1{F0: 18446744073709551615}
            p1 = (p1 << ((304) & 63))
            if false {
                p2 = (2 < 1)
                v27 := []int{fi25, fi25, (-(fi25)), fi25}
                _ = v27
            } else {
                v28 := []int{}
                var v29 int = (^(fi25))
                var v30 string = "a"
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi25
            _ = v26
        }
    }
    p2 = true
    var v31 int8 = -128
    v32 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    return (581 != 283)
}
func genFn1(p0 uint16) uint {
    var v33 int = 9223372036854775807
    fmt.Println(p0)
    p0 = (^(p0))
    fmt.Println(p0, v33)
    _ = p0
    _ = v33
    return 42
}

func main() {
    var v1 uint64 = 581
    v2 := []string{"\n\t", ("foo bar" + ""), "hello", ""}
    var v3 uint = 797
    var v4 uint16 = 787
    v5 := make(map[string]bool)
    var v6 int16 = 32767
    var v34 float64 = 3.14159
    v1 += (v1 + v1)
    delete(v5, "")
    if (-9223372036854775808 == 1) {
        for fi35 := 0; fi35 < 3; fi35++ {
            if false {
                switch 4294967295 {
                case 128:
                    var v36 uint32 = 65535
                    _ = v36
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = ("a" + "foo bar")
                }
                switch "a" {
                case "":
                    var v37 uint32 = 42
                    _ = v37
                case "hello":
                    v4 += v4
                case "a":
                    for fi38 := 0; fi38 < 5; fi38++ {
                        fmt.Println(v4)
                        _ = fi38
                    }
                }
                if false {
                    v5[("hello" + "\n\t")] = true
                } else {
                    fmt.Println(v4, v3, v4)
                    v1 = (v1 & (v1 % ((4294967295) | 1)))
                }
            }
            switch fi35 {
            case 9223372036854775807:
                for fi39 := 0; fi39 < 18; fi39++ {
                    switch (fi39 << ((genFn1(665)) & 63)) {
                    case 9223372036854775807:
                        if (127 > 127) {
                            _ = len(v2)
                            fmt.Println(v3, fi35)
                        }
                    default:
                        if false {
                            v6 = v6
                        }
                    }
                    var v40 uint32 = 1
                    v4 = 623
                    var v41 uint16 = (v4 & (v4 - uint16(v40)))
                    v4 = v4
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            default:
                var v42 int16 = 32767
                _ = v42
            }
            _ = fi35
        }
        v6 = v6
        if ((true || (1e10 > float32(v34))) && true) {
            var v43 int64 = -9223372036854775808
            _ = v43
        }
    } else {
        v4 += v4
        v6 -= -32768
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        if false {
            if true {
                var v45 int = (fi44 + (fi44 | 231))
                v46 := GenS2{F0: genFn1(v4)}
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v47 := GenS3{F0: true, F1: -1}
                    for fi48 := 0; fi48 < 16; fi48++ {
                        if true {
                            v4 -= (v4 << ((v3) & 63))
                            if (true || false) {
                                if (true || false) {
                                    fmt.Println(fi44, fi44)
                                    switch int(v1) {
                                    case 0:
                                        delete(v5, "hello")
                                    }
                                    v6 = (-(v6))
                                } else {
                                    if false {
                                        v5["a"] = (v34 >= ((3.14159 + -2.25) * (-2.25 * v34)))
                                    } else {
                                        v5[(("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + "")] = false
                                        switch fi44 {
                                        case -641:
                                            if (684 == 9223372036854775807) {
                                                var v49 uint64 = 1
                                                fmt.Println(v45, v45)
                                                _ = v49
                                            }
                                        case 0:
                                            fmt.Println(v1)
                                        case 2147483647:
                                            if (-1 < 42) {
                                                v50 := GenS1{F0: 42}
                                                switch fi44 {
                                                case 799:
                                                    var v51 uint32 = 177
                                                    _ = v51
                                                }
                                                fmt.Println(v1, v6)
                                                _ = v50
                                            }
                                        }
                                    }
                                    _ = len(v5)
                                }
                            } else {
                                v52 := GenS4{F0: v4}
                                _ = len(v5)
                                v34 = 1.5
                                _ = v52
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = ("" + "")
                            }
                            v4 -= 0
                            if (!((0 > 128) && true)) {
                                v6 = v6
                                if ((v4 | 65535) != 0) {
                                    v53 := GenS5{F0: (((("\n\t" + "a") + "\n\t") + "") + ("" + "")), F1: (^(v3)), F2: (v6 % (((v6 - (v6 - 524))) | 1))}
                                    v34 = v34
                                    _ = v53
                                }
                                var v54 int = -2147483648
                                _ = v54
                            } else {
                                v6 = (^(v6))
                                if true {
                                    var v55 uint8 = 255
                                    _ = v55
                                } else {
                                    var v56 int = (fi48 + fi44)
                                    if (4294967295 <= uint32(v4)) {
                                        _ = len(v5)
                                    }
                                    var v57 uint8 = 128
                                    var v58 bool = (v3 >= v3)
                                    var v59 int32 = 1
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                var v60 int64 = 797
                                _ = v60
                            }
                        }
                        _ = fi48
                    }
                    v61 := GenS1{F0: (^(v3))}
                    _ = v47
                    _ = v61
                }
                _ = v45
                _ = v46
            } else {
                fmt.Println(v6)
            }
        }
        if false {
            var v62 float32 = (0.1 / 1e10)
            switch fi44 {
            case -9223372036854775808:
                v62 -= v62
            case -2147483648:
                v6 = v6
            case 128:
                switch "a" {
                case "foo bar":
                    for fi63 := 0; fi63 < 3; fi63++ {
                        if true {
                            v64 := GenS2{F0: genFn1(42)}
                            switch "foo bar" {
                            case "a":
                                v2 = append(v2, "a")
                            case "":
                                v65 := GenS2{F0: v3}
                                _ = v65
                            case "\n\t":
                                v1 -= v1
                            default:
                                v5["hello"] = true
                            }
                            _ = v64
                        } else {
                            var v66 uint32 = 2147483648
                            v3 += v3
                            v62 -= v62
                            _ = v66
                        }
                        switch fi63 {
                        case 211:
                            v67 := GenS6{F0: 1}
                            _ = v67
                        }
                        v3 -= genFn1((v4 + (v4 ^ (^(v4)))))
                        _ = fi63
                    }
                case "a":
                    v68 := GenS1{F0: v3}
                    _ = v68
                case "":
                    var v69 int64 = 978
                    _ = v69
                default:
                    var v70 float32 = v62
                    _ = v70
                }
            }
            v71 := GenS4{F0: (v4 / ((uint16(v1)) | 1))}
            _ = v62
            _ = v71
        }
        switch fi44 {
        case 284:
            _ = len(v2)
        case 128:
            if false {
                v72 := GenS7{F0: fi44, F1: 42}
                if false {
                    if ((v4 <= (v4 << ((uint(v1)) & 63))) && true) {
                        v3 += (v3 | v3)
                        v34 += 3.14159
                    }
                }
                var v73 uint32 = 42
                _ = v72
                _ = v73
            } else {
                v34 -= v34
                v4 = 1
                if true {
                    if (1 >= 0) {
                        if (!(uint(v4) <= (v3 * 2))) {
                            delete(v5, "hello")
                            _ = len(v5)
                            for fi74 := 0; fi74 < 6; fi74++ {
                                v75 := v5[((("a" + "a") + "") + "a")]
                                var v76 uint32 = 660
                                v5[("foo bar" + ("foo bar" + ""))] = ((-(fi44)) > fi74)
                                v77 := GenS8{F0: 4294967295, F1: "hello", F2: (true || ((v75 || v75) || (true && (v75 || false))))}
                                _ = fi74
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                        } else {
                            v1 += (v1 + (^(v1)))
                            v78 := GenS9{F0: ("a" + "hello")}
                            _ = v78
                        }
                    } else {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v6 = v6
                        for fi79 := 0; fi79 < 19; fi79++ {
                            v80 := GenS5{F0: "foo bar", F1: 4294967295, F2: v6}
                            _ = fi79
                            _ = v80
                        }
                    }
                    if (!true) {
                        fmt.Println(v34, v6)
                        var v81 int8 = -57
                        var v82 int16 = v6
                        _ = v81
                        _ = v82
                    } else {
                        if false {
                            var v83 int32 = 0
                            if (v3 <= v3) {
                                var v84 string = "a"
                                _ = v84
                            }
                            var v85 int64 = -9223372036854775808
                            _ = v83
                            _ = v85
                        } else {
                            if true {
                                var v86 int32 = 0
                                if false {
                                    for fi87 := 0; fi87 < 13; fi87++ {
                                        fmt.Println(fi44)
                                        var v88 string = "a"
                                        _ = fi87
                                        _ = v88
                                    }
                                    for fi89 := 0; fi89 < 4; fi89++ {
                                        v3 -= genFn1((v4 - (v4 << ((v3) & 63))))
                                        _ = fi89
                                    }
                                    var v90 float64 = v34
                                    _ = v90
                                }
                                v4 -= 1
                                _ = v86
                            } else {
                                switch "a" {
                                case "foo bar":
                                    if (("hello" + "hello") == "foo bar") {
                                        var v91 uint16 = (v4 - (v4 * v4))
                                        switch "foo bar" {
                                        case "hello":
                                            var v92 bool = (v34 == v34)
                                            _ = v92
                                        case "foo bar":
                                            _ = len(v2)
                                        }
                                        _ = v91
                                    }
                                case "":
                                    if true {
                                        var v93 int = (fi44 * fi44)
                                        _ = v93
                                    } else {
                                        switch (fi44 & -9223372036854775808) {
                                        case 2147483647:
                                            v94 := GenS1{F0: genFn1(v4)}
                                            _ = v94
                                        }
                                    }
                                default:
                                    if false {
                                        if false {
                                            var v95 uint = v3
                                            if (!(!false)) {
                                                if false {
                                                    v1 = v1
                                                    if true {
                                                        for fi96 := 0; fi96 < 1; fi96++ {
                                                            var v97 float32 = (0.0 / 1.0)
                                                            _ = fi96
                                                            _ = v97
                                                        }
                                                    }
                                                    for fi98 := 0; fi98 < 17; fi98++ {
                                                        v99 := GenS10{F0: 0, F1: 1, F2: 782}
                                                        if true {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v4 = (v4 ^ (v4 ^ v4))
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        } else {
                                                            v100 := GenS9{F0: ((("\n\t" + "hello") + "\n\t") + ("a" + "hello"))}
                                                            _ = v100
                                                        }
                                                        if false {
                                                            v4 -= v4
                                                            var v101 float32 = (1e10 / 1e10)
                                                            _ = v101
                                                        } else {
                                                            var v102 uint64 = 18446744073709551615
                                                            if false {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    fmt.Println(v34)
                                                                default:
                                                                    if true {
                                                                        var v103 float64 = (0.0 * 1e10)
                                                                        _ = v103
                                                                    }
                                                                }
                                                            }
                                                            _ = len(v2)
                                                            if ((v4 % ((v4) | 1)) < v4) {
                                                                v6 = (-(v6))
                                                                switch "\n\t" {
                                                                case "":
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                            if (((v3 - v95) > (v3 - uint(v1))) && false) {
                                                                fmt.Println(v95, v34)
                                                                if true {
                                                                    v2 = append(v2, "\n\t")
                                                                } else {
                                                                    switch ("a" + (((("a" + "foo bar") + (((("foo bar" + "hello") + ("\n\t" + "hello")) + "hello") + "\n\t")) + "") + ("" + "a"))) {
                                                                    case "":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    default:
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v104 float32 = -2.25
                                                                            _ = v104
                                                                        case "a":
                                                                            v105 := GenS11{F0: 1, F1: ("a" + "a")}
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    v4 = uint16(v1)
                                                                }
                                                                v3 = (v95 % ((v95) | 1))
                                                            }
                                                            _ = v102
                                                        }
                                                        if (81 >= 562) {
                                                            v4 -= (v4 + (v4 + v4))
                                                        }
                                                        if (v1 >= v1) {
                                                            v4 -= v4
                                                            fmt.Println(v95, v1, v95)
                                                        } else {
                                                            if true {
                                                                v5[""] = (true || true)
                                                            }
                                                            fmt.Println(v4, v4, fi98)
                                                            v5[("foo bar" + "\n\t")] = true
                                                        }
                                                        _ = fi98
                                                        _ = v99
                                                    }
                                                } else {
                                                    v5["hello"] = true
                                                    if (128 > 1) {
                                                        v6 = 42
                                                        var v106 int16 = 1000
                                                        _ = v106
                                                    }
                                                }
                                                if (!(true || (fi44 != fi44))) {
                                                    v5["a"] = false
                                                }
                                            }
                                            var v107 int32 = 2147483647
                                            _ = v107
                                            _ = v95
                                        } else {
                                            switch (fi44 | fi44) {
                                            case 197:
                                                for fi108 := 0; fi108 < 17; fi108++ {
                                                    v6 = v6
                                                    _ = fi108
                                                }
                                            case 373:
                                                v109 := GenS7{F0: 604, F1: -1}
                                                _ = v109
                                            default:
                                                var v110 float32 = (0.1 - 3.14159)
                                                _ = v110
                                            }
                                            if true {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if false {
                                                    var v111 uint32 = 4294967295
                                                    var v112 int8 = 1
                                                    v113 := GenS10{F0: v1, F1: 1, F2: int(v4)}
                                                    _ = v111
                                                    _ = v112
                                                    _ = v113
                                                } else {
                                                    if false {
                                                        if false {
                                                            switch fi44 {
                                                            case 544:
                                                                var v114 uint8 = 0
                                                                _ = v114
                                                            case 42:
                                                                switch int(v4) {
                                                                case -936:
                                                                    if false {
                                                                        switch "a" {
                                                                        case "":
                                                                            delete(v5, "foo bar")
                                                                        case "a":
                                                                            var v115 int16 = 32767
                                                                            _ = v115
                                                                        case "hello":
                                                                            if (3.14159 >= 1e10) {
                                                                                var v116 float64 = 1e10
                                                                                if false {
                                                                                    for fi117 := 0; fi117 < 1; fi117++ {
                                                                                        v1 += v1
                                                                                        _ = fi117
                                                                                    }
                                                                                    var v118 float32 = (1e10 / 0.1)
                                                                                    _ = v118
                                                                                }
                                                                                _ = v116
                                                                            } else {
                                                                                fmt.Println(v3, v3, v1)
                                                                                v3 -= (v3 >> ((v3) & 63))
                                                                                v34 -= (((v34 - v34) * -2.25) - (v34 / v34))
                                                                            }
                                                                        }
                                                                        var v119 int16 = (v6 / (((v6 * (v6 - v6))) | 1))
                                                                        _ = v119
                                                                    }
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                            case -275:
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    fmt.Println(v3)
                                                                case "foo bar":
                                                                    for fi120 := 0; fi120 < 17; fi120++ {
                                                                        var v121 string = "hello"
                                                                        _ = fi120
                                                                        _ = v121
                                                                    }
                                                                case "a":
                                                                    for fi122 := 0; fi122 < 9; fi122++ {
                                                                        fmt.Println(v3, v34)
                                                                        switch fi44 {
                                                                        case 32768:
                                                                            if true {
                                                                                if (((v34 > 3.14159) || false) && (127 >= 211)) {
                                                                                    if (true && true) {
                                                                                        var v123 float32 = (float32(-2.25) - ((0.0 - 0.0) * float32(1e10)))
                                                                                        v6 = (v6 / ((484) | 1))
                                                                                        delete(v5, "hello")
                                                                                        _ = v123
                                                                                    }
                                                                                    var v124 bool = (!true)
                                                                                    _ = v124
                                                                                }
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v125 := GenS12{F0: v1, F1: -798}
                                                                                    _ = v125
                                                                                case "a":
                                                                                    switch fi44 {
                                                                                    case 962:
                                                                                        v1 += (v1 % ((uint64(v3)) | 1))
                                                                                    }
                                                                                case "":
                                                                                    fmt.Println(v6, fi44, fi44)
                                                                                default:
                                                                                    switch fi44 {
                                                                                    case 2147483647:
                                                                                        v1 += v1
                                                                                    }
                                                                                    v1 = v1
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v126 int = 32768
                                                                                        _ = v126
                                                                                    }
                                                                                    fmt.Println(fi122)
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        _ = len(v2)
                                                                                    }
                                                                                    v127 := GenS13{F0: 960, F1: 4294967295}
                                                                                    _ = v127
                                                                                }
                                                                                v128 := GenS14{F0: 100, F1: 32, F2: "a"}
                                                                                _ = v128
                                                                            }
                                                                        case 9223372036854775807:
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                var v129 uint16 = (v4 + (v4 + (v4 - (v4 + 690))))
                                                                                _ = v129
                                                                            case "":
                                                                                v4 = (^(v4))
                                                                            case "foo bar":
                                                                                var v130 uint64 = v1
                                                                                _ = v130
                                                                            default:
                                                                                v131 := GenS3{F0: ((3.14159 / -2.25) >= ((0.0 * (3.14159 + float32((float64(1e10) / v34)))) - 1.5)), F1: 1}
                                                                                _ = v131
                                                                            }
                                                                        case -2147483648:
                                                                            var v132 uint16 = v4
                                                                            _ = v132
                                                                        }
                                                                        _ = fi122
                                                                    }
                                                                }
                                                            }
                                                            for fi133 := 0; fi133 < 4; fi133++ {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v134 int64 = int64(v1)
                                                                _ = fi133
                                                                _ = v134
                                                            }
                                                        }
                                                    } else {
                                                        v4 = v4
                                                    }
                                                    for fi135 := 0; fi135 < 18; fi135++ {
                                                        v34 -= ((((v34 - float64((1.5 + 0.0))) + (float64(1e10) / float64(float32((1.5 * float64((-2.25 * 3.14159))))))) / v34) + v34)
                                                        switch "a" {
                                                        case "":
                                                            if true {
                                                                var v136 float64 = v34
                                                                if false {
                                                                    v137 := v5["\n\t"]
                                                                    v5["\n\t"] = v137
                                                                    if v137 {
                                                                        v3 += 0
                                                                        var v138 uint = (^(v3))
                                                                        v5["\n\t"] = ("hello" == "hello")
                                                                        _ = v138
                                                                    }
                                                                    _ = v137
                                                                } else {
                                                                    v139 := v5[""]
                                                                    _ = v139
                                                                }
                                                                if false {
                                                                    v34 = (float64(1.5) - ((-2.25 - v34) * v136))
                                                                    var v140 int64 = -1
                                                                    _ = v140
                                                                } else {
                                                                    v3 = (v3 + (v3 + v3))
                                                                    if true {
                                                                        switch "hello" {
                                                                        case "":
                                                                            v141 := GenS15{F0: (v3 ^ (v3 | v3)), F1: -20, F2: (v3 << ((v3) & 63))}
                                                                            _ = v141
                                                                        case "hello":
                                                                            if ((false || true) || (!true)) {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    switch ("foo bar" + "") {
                                                                                    case "hello":
                                                                                        v4 -= uint16(v1)
                                                                                    case "foo bar":
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            v4 = (^(v4))
                                                                                        case "a":
                                                                                            var v142 int64 = -9223372036854775808
                                                                                            _ = v142
                                                                                        }
                                                                                    default:
                                                                                        var v143 bool = false
                                                                                        _ = v143
                                                                                    }
                                                                                case "hello":
                                                                                    v34 = v34
                                                                                }
                                                                                v6 -= v6
                                                                                v3 = (^(v3))
                                                                            }
                                                                        }
                                                                        var v144 uint32 = 2147483647
                                                                        v145 := v5[((("" + "hello") + ("hello" + "a")) + "\n\t")]
                                                                        _ = v144
                                                                        _ = v145
                                                                    } else {
                                                                        if ("a" == ("\n\t" + ("foo bar" + ("a" + "")))) {
                                                                            var v146 uint64 = v1
                                                                            fmt.Println(v3, v1, v4)
                                                                            _ = v146
                                                                        }
                                                                        var v147 int64 = 1
                                                                        if (!(true && true)) {
                                                                            v148 := GenS15{F0: v3, F1: -1, F2: 42}
                                                                            v147 += (v147 % (((v147 - v147)) | 1))
                                                                            _ = v148
                                                                        }
                                                                        _ = v147
                                                                    }
                                                                    fmt.Println(v4, v4, v34)
                                                                }
                                                                _ = v136
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v1)
                                                        case "foo bar":
                                                            v3 = (v3 - 104)
                                                        }
                                                        _ = fi135
                                                    }
                                                }
                                                v149 := GenS15{F0: (v3 << (((v3 + (v3 & 99))) & 63)), F1: -128, F2: 18446744073709551615}
                                                _ = v149
                                            }
                                        }
                                    } else {
                                        v6 += (v6 - v6)
                                    }
                                }
                            }
                            if (true || false) {
                                fmt.Println(v1)
                                var v150 string = ("foo bar" + "hello")
                                switch v150 {
                                case "":
                                    for fi151 := 0; fi151 < 18; fi151++ {
                                        _ = len(v5)
                                        _ = fi151
                                    }
                                }
                                _ = v150
                            } else {
                                for fi152 := 0; fi152 < 3; fi152++ {
                                    switch (fi44 >> ((v3) & 63)) {
                                    case -574:
                                        fmt.Println(fi44, v34, v34)
                                    default:
                                        if true {
                                            fmt.Println(v3)
                                            if (!(true && (!(v34 == (v34 / v34))))) {
                                                var v153 int8 = 100
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi152, v153, v6)
                                                case "a":
                                                    var v154 bool = false
                                                    _ = v154
                                                case "":
                                                    fmt.Println(v1)
                                                default:
                                                    switch (fi152 & (^(fi44))) {
                                                    case -2147483648:
                                                        var v155 int32 = int32(v3)
                                                        _ = v155
                                                    case 9223372036854775807:
                                                        v4 = (v4 / ((v4) | 1))
                                                    case 1:
                                                        fmt.Println(v34)
                                                    default:
                                                        var v156 int64 = -1
                                                        _ = v156
                                                    }
                                                }
                                                _ = v153
                                            }
                                        } else {
                                            switch (("" + (((("\n\t" + ("a" + "a")) + "\n\t") + "") + "")) + "a") {
                                            case "hello":
                                                v157 := GenS2{F0: (v3 & (v3 + (^(v3))))}
                                                _ = v157
                                            case "":
                                                fmt.Println(fi152, v6)
                                            case "\n\t":
                                                var v158 bool = (!true)
                                                _ = v158
                                            default:
                                                fmt.Println(v6)
                                            }
                                            switch fi152 {
                                            case 713:
                                                if true {
                                                    if (uint8(v3) >= 127) {
                                                        fmt.Println(fi44, v4)
                                                        var v159 uint8 = 128
                                                        var v160 uint = v3
                                                        _ = v159
                                                        _ = v160
                                                    } else {
                                                        var v161 int16 = 0
                                                        var v162 uint8 = 1
                                                        _ = v161
                                                        _ = v162
                                                    }
                                                    var v163 uint64 = (v1 & (v1 & v1))
                                                    v164 := GenS12{F0: (v163 << ((65535) & 63)), F1: 1}
                                                    _ = v163
                                                    _ = v164
                                                } else {
                                                    fmt.Println(v6, v1)
                                                    var v165 int = 2147483647
                                                    fmt.Println(v4)
                                                    _ = v165
                                                }
                                            case 32767:
                                                fmt.Println(v6, v34)
                                            case -9223372036854775808:
                                                v166 := GenS5{F0: "foo bar", F1: v3, F2: 32767}
                                                _ = v166
                                            }
                                        }
                                    }
                                    v3 = genFn1(v4)
                                    if false {
                                        var v167 float64 = -2.25
                                        if (v4 == (v4 - v4)) {
                                            v168 := GenS4{F0: (v4 | (v4 * v4))}
                                            switch "\n\t" {
                                            case "":
                                                v6 -= v6
                                            case "a":
                                                v6 -= v6
                                                v34 = (1.5 * 3.14159)
                                                fmt.Println(v6, v1)
                                            }
                                            _ = len(v2)
                                            _ = v168
                                        } else {
                                            if false {
                                                if true {
                                                    v5[("foo bar" + "\n\t")] = false
                                                } else {
                                                    v3 += (v3 ^ (v3 * (v3 >> (((^(v3))) & 63))))
                                                    v3 -= (v3 / ((v3) | 1))
                                                    v169 := GenS8{F0: 2147483647, F1: (("hello" + "hello") + "\n\t"), F2: (180 == 127)}
                                                    _ = v169
                                                }
                                            }
                                            var v170 bool = true
                                            _ = v170
                                        }
                                        _ = v167
                                    }
                                    _ = fi152
                                }
                                var v171 string = "\n\t"
                                _ = v171
                            }
                        }
                        if (!false) {
                            switch (fi44 >> ((4294967295) & 63)) {
                            case -860:
                                var v172 int8 = 100
                                _ = v172
                            case 32767:
                                var v173 int = fi44
                                _ = v173
                            case 127:
                                if false {
                                    v4 = (v4 + (v4 ^ 476))
                                    var v174 float32 = 0.0
                                    _ = v174
                                } else {
                                    switch (fi44 | fi44) {
                                    case 42:
                                        var v175 bool = true
                                        _ = v175
                                    case 987:
                                        for fi176 := 0; fi176 < 10; fi176++ {
                                            v177 := GenS7{F0: (fi176 | (fi44 + (fi176 - fi176))), F1: -128}
                                            var v178 int16 = v6
                                            _ = fi176
                                            _ = v177
                                            _ = v178
                                        }
                                    default:
                                        v3 = v3
                                    }
                                }
                            }
                            var v179 bool = true
                            var v180 uint8 = 128
                            _ = v179
                            _ = v180
                        }
                    }
                } else {
                    v34 += 0.0
                    var v181 int64 = 1
                    _ = v181
                }
            }
        default:
            switch "" {
            case "a":
                v182 := GenS10{F0: (v1 + v1), F1: 77, F2: (fi44 % ((42) | 1))}
                _ = v182
            }
        }
        _ = fi44
    }
    for fi183 := 0; fi183 < 2; fi183++ {
        switch "a" {
        case "hello":
            v184 := GenS8{F0: 2, F1: "", F2: true}
            _ = v184
        case "":
            if true {
                var v185 int8 = 25
                if (v6 <= (v6 >> (((v3 << ((genFn1(v4)) & 63))) & 63))) {
                    v186 := GenS12{F0: v1, F1: 639}
                    switch ("hello" + "\n\t") {
                    case "a":
                        var v187 int32 = 42
                        _ = v187
                    }
                    _ = v186
                } else {
                    v6 -= (-(v6))
                    v5["\n\t"] = ((!(v185 != 88)) && (65535 >= 1))
                    if true {
                        for fi188 := 0; fi188 < 18; fi188++ {
                            switch fi188 {
                            case -850:
                                v3 = 18446744073709551615
                            case 644:
                                if (!((true || false) && (true && (2 == 2)))) {
                                    var v189 int64 = -326
                                    if true {
                                        var v190 bool = false
                                        if (uint8(v3) < 128) {
                                            fmt.Println(v3, v3)
                                        } else {
                                            fmt.Println(v4, v185, v1)
                                            var v191 uint = genFn1((v4 + 167))
                                            v2 = append(v2, ("foo bar" + (("\n\t" + "a") + "hello")))
                                            _ = v191
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "\n\t"
                                        }
                                        _ = v190
                                    } else {
                                        var v192 int32 = 709
                                        _ = v192
                                    }
                                    v34 += ((v34 - v34) / float64(1.5))
                                    v2 = append(v2, "hello")
                                    _ = v189
                                }
                            case 127:
                                if true {
                                    if (-2.25 > 0.0) {
                                        v4 += 272
                                        var v193 bool = (110 < 42)
                                        fmt.Println(v3)
                                        _ = v193
                                    }
                                    v194 := GenS4{F0: (v4 & (v4 + (v4 - v4)))}
                                    v195 := GenS16{F0: 255}
                                    _ = v194
                                    _ = v195
                                } else {
                                    var v196 float64 = (v34 - 0.1)
                                    switch -856 {
                                    case -372:
                                        if (!true) {
                                            v1 = 4294967295
                                            v1 -= 2
                                            if false {
                                                var v197 int64 = 42
                                                _ = v197
                                            }
                                            fmt.Println(fi188, fi183, v4)
                                            var v198 int64 = 9223372036854775807
                                            fmt.Println(fi188, fi183)
                                            v196 = 1e10
                                            if (571 <= 42) {
                                                v3 = genFn1(v4)
                                                v196 -= v196
                                            }
                                            _ = v198
                                        } else {
                                            if false {
                                                v199 := GenS17{F0: v1}
                                                var v200 int64 = -68
                                                v185 = -43
                                                _ = v199
                                                _ = v200
                                            }
                                            v201 := GenS2{F0: 0}
                                            var v202 string = "\n\t"
                                            fmt.Println(fi183)
                                            _ = v201
                                            _ = v202
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        fmt.Println(v185)
                                        if (48 == 36) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v2 = append(v2, "")
                                            case "hello":
                                                if false {
                                                    v185 = (-(v185))
                                                    var v203 uint = 85
                                                    delete(v5, "foo bar")
                                                    if false {
                                                        v1 = (v1 & v1)
                                                        if len(v2) > 0 {
                                                            v2[0] = "foo bar"
                                                        }
                                                        if (!true) {
                                                            switch (^(fi183)) {
                                                            case 766:
                                                                v4 -= v4
                                                            case -87:
                                                                v204 := GenS8{F0: 4294967295, F1: "", F2: false}
                                                                _ = v204
                                                            case 1:
                                                                var v205 float32 = (-2.25 - 0.1)
                                                                v6 -= int16(fi188)
                                                                switch fi183 {
                                                                case -2147483648:
                                                                    v3 = 255
                                                                    fmt.Println(v3, v203)
                                                                    fmt.Println(v4)
                                                                case 32767:
                                                                    v34 -= ((0.1 + (v196 + (-2.25 + (float64(v205) - float64(float32((v34 - v196))))))) - v34)
                                                                case 4294967295:
                                                                    if false {
                                                                        fmt.Println(v1, v185, v6)
                                                                        v203 = genFn1((v4 + 32767))
                                                                        v205 = float32(v34)
                                                                        fmt.Println(v205, v1, v196)
                                                                        v4 = 42
                                                                    } else {
                                                                        v206 := GenS14{F0: 65, F1: v4, F2: ("foo bar" + "hello")}
                                                                        _ = v206
                                                                    }
                                                                default:
                                                                    _ = len(v5)
                                                                    delete(v5, "")
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                _ = v205
                                                            default:
                                                                v3 = 394
                                                            }
                                                            if false {
                                                                v4 = v4
                                                                v207 := GenS4{F0: 1}
                                                                var v208 int64 = int64(v203)
                                                                _ = v207
                                                                _ = v208
                                                            }
                                                            var v209 float32 = float32(1.5)
                                                            _ = v209
                                                        }
                                                    } else {
                                                        v196 = v34
                                                    }
                                                    var v210 uint8 = 128
                                                    _ = v203
                                                    _ = v210
                                                } else {
                                                    v1 -= 18446744073709551615
                                                    fmt.Println(v1)
                                                    switch 139 {
                                                    case 0:
                                                        v211 := GenS18{F0: -625, F1: 128}
                                                        _ = v211
                                                    default:
                                                        v6 = (v6 >> (((v3 | genFn1(v4))) & 63))
                                                    }
                                                }
                                            case "foo bar":
                                                v185 -= int8(v1)
                                            }
                                            var v212 int32 = 163
                                            _ = v212
                                        } else {
                                            if (-195 > -568) {
                                                if true {
                                                    v213 := GenS18{F0: 1000, F1: 62}
                                                    if false {
                                                        v214 := GenS19{F0: (v4 + (v4 * 400))}
                                                        _ = v214
                                                    }
                                                    _ = v213
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v215 uint32 = 65535
                                                        _ = v215
                                                    case "a":
                                                        var v216 int16 = (v6 << (((v3 + genFn1((v4 | 647)))) & 63))
                                                        _ = v216
                                                    case "hello":
                                                        var v217 uint32 = 172
                                                        _ = v217
                                                    default:
                                                        if true {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v6 = (^(v6))
                                                            default:
                                                                v218 := GenS16{F0: 42}
                                                                _ = v218
                                                            }
                                                            fmt.Println(v196, v4)
                                                            switch 127 {
                                                            case -1:
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v219 float32 = 1e10
                                                                    _ = v219
                                                                case "a":
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "hello"
                                                                    }
                                                                case "foo bar":
                                                                    v6 = 1000
                                                                default:
                                                                    v34 += 0.0
                                                                    v185 = v185
                                                                }
                                                                v4 = (^(v4))
                                                            case 2147483647:
                                                                fmt.Println(v185, v4, v185)
                                                            case 341:
                                                                var v220 int32 = 1
                                                                _ = v220
                                                            }
                                                        } else {
                                                            var v221 int = fi183
                                                            switch int(v185) {
                                                            case -962:
                                                                v222 := GenS14{F0: (v185 ^ (^(v185))), F1: v4, F2: "hello"}
                                                                _ = v222
                                                            case -9223372036854775808:
                                                                var v223 int8 = v185
                                                                _ = v223
                                                            default:
                                                                v6 -= (^(v6))
                                                            }
                                                            _ = v221
                                                        }
                                                    }
                                                    v1 -= v1
                                                    fmt.Println(v3, v185, v1)
                                                }
                                            }
                                            v196 -= v34
                                            if true {
                                                var v224 int64 = 9223372036854775807
                                                _ = v224
                                            } else {
                                                var v225 uint32 = 42
                                                if true {
                                                    var v226 int32 = -528
                                                    var v227 bool = true
                                                    fmt.Println(v196, v1, v196)
                                                    _ = v226
                                                    _ = v227
                                                } else {
                                                    v3 = v3
                                                    if (v6 >= (v6 * (v6 | int16(fi183)))) {
                                                        var v228 uint16 = (v4 / (((v4 * 453)) | 1))
                                                        switch fi188 {
                                                        case -549:
                                                            if false {
                                                                var v229 int16 = v6
                                                                if false {
                                                                    fmt.Println(v196, fi183, v196)
                                                                    if true {
                                                                        v1 += (^(v1))
                                                                        var v230 int16 = v229
                                                                        v196 += 3.14159
                                                                        if true {
                                                                            var v231 int16 = 162
                                                                            _ = v231
                                                                        }
                                                                        _ = v230
                                                                    }
                                                                } else {
                                                                    fmt.Println(v225, v196)
                                                                    if false {
                                                                        var v232 int8 = -81
                                                                        v225 -= (v225 - v225)
                                                                        v232 += v185
                                                                        _ = v232
                                                                    }
                                                                    if ((v185 | (v185 + v185)) <= int8(v1)) {
                                                                        v2 = append(v2, "hello")
                                                                    } else {
                                                                        v34 += v34
                                                                        switch fi188 {
                                                                        case 382:
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        case 950:
                                                                            if (!false) {
                                                                                if true {
                                                                                    v1 -= (v1 >> (((v3 >> ((v3) & 63))) & 63))
                                                                                } else {
                                                                                    if false {
                                                                                        v233 := GenS18{F0: (v6 * v6), F1: 1}
                                                                                        _ = v233
                                                                                    }
                                                                                }
                                                                                var v234 int8 = (v185 >> (((v3 | (^(v3)))) & 63))
                                                                                fmt.Println(v4, v1, v3)
                                                                                _ = v234
                                                                            }
                                                                        case 0:
                                                                            _ = len(v5)
                                                                        }
                                                                        fmt.Println(fi188, v229)
                                                                    }
                                                                    v6 = (-(v6))
                                                                }
                                                                var v235 string = "a"
                                                                _ = v229
                                                                _ = v235
                                                            }
                                                        default:
                                                            v236 := GenS1{F0: (^(v3))}
                                                            _ = v236
                                                        }
                                                        _ = v228
                                                    }
                                                    v6 = int16(v4)
                                                    var v237 uint16 = uint16(v3)
                                                    v238 := GenS20{F0: ("hello" + "a"), F1: v185, F2: genFn1((v237 ^ (v237 % (((v237 | (^(v4)))) | 1))))}
                                                    if (v6 < v6) {
                                                        v225 = (^(v225))
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            switch "foo bar" {
                                                            case "":
                                                                if false {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "\n\t"
                                                                    }
                                                                }
                                                                if (85 != 255) {
                                                                    v3 = 425
                                                                    v2 = append(v2, ("a" + "hello"))
                                                                    fmt.Println(v185)
                                                                    var v239 uint16 = 2
                                                                    _ = v239
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "\n\t"
                                                                    }
                                                                    v225 += 2
                                                                    var v240 int32 = 494
                                                                    _ = v240
                                                                }
                                                                v4 += (v4 + (v4 << ((255) & 63)))
                                                            case "\n\t":
                                                                var v241 float32 = (1.5 / 1.5)
                                                                _ = v241
                                                            }
                                                        case "\n\t":
                                                            v185 += v185
                                                        case "hello":
                                                            v1 = v1
                                                        }
                                                    } else {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v242 bool = (true && true)
                                                            _ = v242
                                                        default:
                                                            delete(v5, ("a" + (("\n\t" + "") + ((("hello" + ((("" + "foo bar") + ((("foo bar" + "foo bar") + "\n\t") + "\n\t")) + "foo bar")) + "hello") + "hello"))))
                                                        }
                                                        if ((!true) || ((-1 >= (v185 | v185)) || true)) {
                                                            if false {
                                                                if ((v237 | v4) > (v237 ^ v4)) {
                                                                    if true {
                                                                        v6 += (v6 >> ((42) & 63))
                                                                        delete(v5, "")
                                                                        var v243 int32 = 1
                                                                        var v244 int16 = v6
                                                                        var v245 uint16 = (v237 % (((v237 << ((genFn1(v237)) & 63))) | 1))
                                                                        v245 = (v245 & v245)
                                                                        if true {
                                                                            fmt.Println(v34, fi183)
                                                                        }
                                                                        if true {
                                                                            v246 := v5[("foo bar" + "")]
                                                                            _ = v246
                                                                        }
                                                                        var v247 int32 = 42
                                                                        _ = v243
                                                                        _ = v244
                                                                        _ = v245
                                                                        _ = v247
                                                                    } else {
                                                                        var v248 uint16 = (v4 % ((uint16(v1)) | 1))
                                                                        if len(v2) > 0 {
                                                                            v2[0] = ("\n\t" + "")
                                                                        }
                                                                        v249 := GenS16{F0: uint8(v225)}
                                                                        _ = v248
                                                                        _ = v249
                                                                    }
                                                                    v250 := GenS11{F0: 2147483647, F1: ""}
                                                                    _ = v250
                                                                } else {
                                                                    if true {
                                                                        v185 = 100
                                                                        var v251 uint64 = 1
                                                                        _ = v251
                                                                    } else {
                                                                        if (!false) {
                                                                            v252 := GenS21{F0: fi188, F1: v225, F2: 0}
                                                                            _ = v252
                                                                        }
                                                                        v253 := GenS22{F0: (v225 + v225)}
                                                                        _ = v253
                                                                    }
                                                                    v237 -= (v4 - (v4 % ((v4) | 1)))
                                                                    if (v185 == (-(v185))) {
                                                                        v5["\n\t"] = true
                                                                        v254 := v5["a"]
                                                                        var v255 uint = (v3 / ((v3) | 1))
                                                                        _ = v254
                                                                        _ = v255
                                                                    }
                                                                    var v256 uint = v3
                                                                    v185 = v185
                                                                    v257 := GenS9{F0: "a"}
                                                                    _ = v256
                                                                    _ = v257
                                                                }
                                                                if false {
                                                                    v1 += v1
                                                                }
                                                                delete(v5, (("foo bar" + (("foo bar" + "foo bar") + ("hello" + "foo bar"))) + ("hello" + "foo bar")))
                                                            } else {
                                                                if false {
                                                                    v258 := GenS8{F0: v225, F1: ("a" + ((("foo bar" + "a") + "\n\t") + "a")), F2: (426 < 255)}
                                                                    v4 = (v4 + v4)
                                                                    _ = v258
                                                                }
                                                            }
                                                            v259 := GenS19{F0: v237}
                                                            v260 := GenS5{F0: "hello", F1: v3, F2: (v6 >> ((v3) & 63))}
                                                            _ = v259
                                                            _ = v260
                                                        }
                                                    }
                                                    _ = v237
                                                    _ = v238
                                                }
                                                var v261 int32 = 42
                                                _ = v225
                                                _ = v261
                                            }
                                            if (!(false || (true || ((v185 << (((v3 | v3)) & 63)) >= (v185 | v185))))) {
                                                v1 = 4294967295
                                                switch "hello" {
                                                case "a":
                                                    v34 -= (v34 + v196)
                                                case "\n\t":
                                                    if (65535 < 1) {
                                                        var v262 float32 = -2.25
                                                        v263 := GenS23{F0: 0}
                                                        v1 = v1
                                                        _ = v262
                                                        _ = v263
                                                    }
                                                case "hello":
                                                    v264 := GenS24{F0: 1, F1: ("foo bar" + ("" + "foo bar"))}
                                                    _ = v264
                                                }
                                            } else {
                                                switch 1 {
                                                case 494:
                                                    v265 := GenS13{F0: -1, F1: v3}
                                                    _ = v265
                                                default:
                                                    v266 := GenS25{F0: 0, F1: 127}
                                                    _ = v266
                                                }
                                                var v267 uint64 = (^(v1))
                                                var v268 uint64 = v1
                                                _ = v267
                                                _ = v268
                                            }
                                            switch "a" {
                                            case "a":
                                                var v269 float64 = ((v34 - 0.0) * v34)
                                                _ = v269
                                            default:
                                                var v270 bool = true
                                                v271 := GenS14{F0: -128, F1: v4, F2: "a"}
                                                _ = v270
                                                _ = v271
                                            }
                                            v272 := GenS10{F0: 1, F1: 42, F2: (fi188 * fi183)}
                                            v4 = v4
                                            switch 4294967295 {
                                            case 9223372036854775807:
                                                v273 := GenS26{F0: v3, F1: -1, F2: (fi188 << (((v3 * v3)) & 63))}
                                                _ = v273
                                            case 1:
                                                v6 -= (v6 << ((v3) & 63))
                                            case 128:
                                                if false {
                                                    var v274 uint32 = 65535
                                                    v196 = v34
                                                    v275 := GenS27{F0: int(v274)}
                                                    v6 -= v6
                                                    _ = v274
                                                    _ = v275
                                                }
                                            }
                                            _ = v272
                                        }
                                        var v276 float32 = float32(v34)
                                        _ = v276
                                    case 127:
                                        switch 1 {
                                        case 0:
                                            fmt.Println(fi188)
                                        case 67:
                                            var v277 uint8 = 0
                                            _ = v277
                                        case -528:
                                            var v278 uint8 = uint8(v4)
                                            _ = v278
                                        default:
                                            v34 = ((0.0 * (((float64(-2.25) / 0.1) + (((1.5 * float64(3.14159)) - (float64(1.5) / v196)) - v34)) + 1e10)) + (float64((((0.1 * float32(v196)) - ((0.1 - float32(v34)) * 1e10)) + -2.25)) * v34))
                                        }
                                    case 32768:
                                        fmt.Println(v196)
                                    }
                                    _ = v196
                                }
                            default:
                                switch fi188 {
                                case 234:
                                    var v279 float64 = (1e10 / (-2.25 * 0.1))
                                    _ = v279
                                }
                            }
                            var v280 int32 = -771
                            switch "foo bar" {
                            case "":
                                var v281 float64 = 3.14159
                                _ = v281
                            case "hello":
                                if ((float64((0.0 - ((float32(1.5) * (0.1 * 0.1)) * (3.14159 / ((1e10 / float32(v34)) * 1.5))))) + float64(3.14159)) > v34) {
                                    v282 := GenS28{F0: 4294967295, F1: v3, F2: (v4 << ((v3) & 63))}
                                    v185 += 127
                                    _ = v282
                                } else {
                                    v1 = v1
                                    switch ("hello" + "hello") {
                                    case "a":
                                        v283 := GenS25{F0: 2147483647, F1: 133}
                                        _ = v283
                                    }
                                }
                            }
                            _ = fi188
                            _ = v280
                        }
                        for fi284 := 0; fi284 < 18; fi284++ {
                            var v285 float64 = v34
                            fmt.Println(v4)
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            _ = fi284
                            _ = v285
                        }
                    }
                }
                v185 = 24
                _ = v185
            }
        case "\n\t":
            switch (("hello" + "\n\t") + "foo bar") {
            case "a":
                fmt.Println(v1, v1)
            case "":
                v5["foo bar"] = ((^(v1)) <= uint64(v3))
            case "hello":
                v286 := GenS29{F0: true, F1: -1, F2: v6}
                _ = v286
            }
        }
        _ = fi183
    }
    v3 -= v3
    v287 := GenS30{F0: (2147483647 != 767)}
    v288 := v5[""]
    var v289 uint8 = 128
    var v290 bool = true
    v1 = v1
    v288 = (0 >= (v4 << ((255) & 63)))
    v291 := GenS31{F0: ("hello" + "\n\t"), F1: v1, F2: -1}
    var v292 bool = v290
    _ = len(v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -875 {
    case 1:
        for fi293 := 0; fi293 < 5; fi293++ {
            var v294 uint = 4294967295
            v295 := GenS32{F0: 42, F1: 944, F2: 750}
            _ = fi293
            _ = v294
            _ = v295
        }
    case -137:
        if (!true) {
            for fi296 := 0; fi296 < 10; fi296++ {
                var v297 int32 = int32(v1)
                switch "\n\t" {
                case "":
                    var v298 int8 = int8(v289)
                    _ = v298
                case "a":
                    v299 := GenS33{F0: v1, F1: 42}
                    _ = v299
                case "hello":
                    _ = len(v5)
                }
                _ = fi296
                _ = v297
            }
            v289 -= v289
            for fi300 := 0; fi300 < 14; fi300++ {
                var v301 uint32 = 4294967295
                _ = fi300
                _ = v301
            }
        } else {
            fmt.Println(v6, v1, v4)
        }
    case -1:
        v289 = 255
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v288:", v288)
    fmt.Println("v289:", v289)
    fmt.Println("v290:", v290)
    fmt.Println("v292:", v292)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v287
    _ = v288
    _ = v289
    _ = v290
    _ = v291
    _ = v292
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
