package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int64; F1 int16; F2 int16 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int; F1 int; F2 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint; F1 string }
type GenS15 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint) uint16 {
    switch "foo bar" {
    case "\n\t":
        p0 += 65535
    case "a":
        fmt.Println(p0, p0)
    default:
        fmt.Println(p0)
    }
    v7 := []int{42, -2147483648, -2147483648}
    _ = p0
    _ = p1
    _ = v7
    return 913
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := []uint32{2}
    v4 := make(map[int]uint32)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]int)
    switch (("\n\t" + "hello") + ("hello" + ("" + (((("hello" + "\n\t") + "") + "") + ("hello" + "foo bar"))))) {
    case "\n\t":
        _ = len(v3)
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch "hello" {
            case "a":
                v5 -= v5
            case "hello":
                v10 := GenS1{F0: 1000, F1: 564, F2: -1}
                _ = v10
            default:
                var v11 int16 = 32767
                fmt.Println(fi8)
                if false {
                    v11 -= 32767
                    v5 = (^(v5))
                    var v12 int64 = 1
                    var v13 string = "\n\t"
                    _ = v12
                    _ = v13
                }
                _ = v11
            }
            var v14 int64 = 0
            var v15 bool = false
            _ = fi9
            _ = v14
            _ = v15
        }
        v5 -= (v5 * 684)
        v5 += v5
        _ = fi8
    }
    delete(v1, 128)
    v5 = v5
    v16 := v1[390]
    v5 = v5
    _ = len(v3)
    v17 := v6["hello"]
    var v18 uint = 18446744073709551615
    v19 := GenS2{F0: 1}
    v18 = (^(v18))
    if false {
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS3{F0: 28, F1: (v16 + v16), F2: v16}
            v22 := v2[uint32(v5)]
            fmt.Println(v17, v22, v5)
            _ = fi20
            _ = v21
            _ = v22
        }
        if (0 != 1) {
            switch v17 {
            case 9223372036854775807:
                fmt.Println(v5, v5)
            case 43:
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := GenS2{F0: -1}
                    if ("" < "a") {
                        v25 := GenS4{F0: 517, F1: 0}
                        if (0.1 != -2.25) {
                            v17 -= v17
                            switch v17 {
                            case 2147483647:
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v5 = v5
                                        if false {
                                            var v27 int16 = 32767
                                            v27 = v27
                                            _ = len(v3)
                                            _ = v27
                                        } else {
                                            fmt.Println(fi23)
                                            v28 := v2[423]
                                            v28 = v28
                                            _ = v28
                                        }
                                    }
                                    delete(v2, (v16 << ((v18) & 63)))
                                    v29 := v1[(fi26 ^ (v17 >> ((82) & 63)))]
                                    _ = fi26
                                    _ = v29
                                }
                            }
                        } else {
                            v30 := GenS5{F0: 0, F1: 1000, F2: -32768}
                            v31 := GenS1{F0: -32768, F1: 42, F2: 108}
                            v17 = (^(v17))
                            _ = v30
                            _ = v31
                        }
                        _ = v25
                    } else {
                        if false {
                            if false {
                                v16 = v16
                                var v32 float32 = ((1e10 * 1e10) + -2.25)
                                v33 := GenS6{F0: v5, F1: v18}
                                _ = v32
                                _ = v33
                            } else {
                                v18 += (v18 / ((v18) | 1))
                            }
                            v34 := GenS3{F0: 2, F1: 65535, F2: (v16 % (((^(v16))) | 1))}
                            _ = v34
                        }
                        switch (("" + "a") + "a") {
                        case "a":
                            for fi35 := 0; fi35 < 3; fi35++ {
                                switch (^(fi23)) {
                                case 127:
                                    v36 := v2[(v16 / (((v16 / (((^(v16))) | 1))) | 1))]
                                    _ = v36
                                }
                                _ = fi35
                            }
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    _ = fi23
                    _ = v24
                }
            default:
                if (-1 < -776) {
                    var v37 float64 = float64(1e10)
                    if len(v3) > 0 {
                        v3[0] = (v16 & (v16 >> ((v18) & 63)))
                    }
                    if true {
                        if (("" + ("" + "")) == "foo bar") {
                            switch 9223372036854775807 {
                            case 4294967295:
                                v38 := GenS4{F0: 0, F1: (v18 - v18)}
                                _ = v38
                            case 175:
                                if false {
                                    v18 += v18
                                    if (!("\n\t" != ("a" + "a"))) {
                                        v17 -= v17
                                    } else {
                                        fmt.Println(v37, v17)
                                        var v39 int32 = 0
                                        _ = v39
                                    }
                                    var v40 int8 = -1
                                    _ = v40
                                } else {
                                    if false {
                                        switch v17 {
                                        case 2147483647:
                                            switch ((("a" + (("\n\t" + "foo bar") + "\n\t")) + "\n\t") + ("a" + ("foo bar" + ("\n\t" + "\n\t")))) {
                                            case "\n\t":
                                                var v41 uint = v18
                                                _ = v41
                                            default:
                                                fmt.Println(v5)
                                            }
                                        default:
                                            var v42 float64 = ((1e10 / v37) + 1e10)
                                            _ = v42
                                        }
                                    } else {
                                        var v43 uint = (v18 & v18)
                                        _ = v43
                                    }
                                }
                            case 9223372036854775807:
                                v5 -= (^(v5))
                            default:
                                var v44 int = v17
                                _ = v44
                            }
                            switch ("foo bar" + (("hello" + "hello") + "")) {
                            case "\n\t":
                                for fi45 := 0; fi45 < 6; fi45++ {
                                    v5 = 42
                                    _ = fi45
                                }
                            case "hello":
                                var v46 int32 = 0
                                _ = v46
                            case "":
                                _ = len(v1)
                            default:
                                if true {
                                    var v47 int = v17
                                    _ = v47
                                }
                            }
                        } else {
                            if false {
                                for fi48 := 0; fi48 < 12; fi48++ {
                                    v49 := GenS5{F0: 1, F1: 32767, F2: -32768}
                                    fmt.Println(v18)
                                    _ = fi48
                                    _ = v49
                                }
                                if false {
                                    for fi50 := 0; fi50 < 6; fi50++ {
                                        var v51 uint32 = v16
                                        _ = fi50
                                        _ = v51
                                    }
                                } else {
                                    switch -880 {
                                    case 307:
                                        v52 := GenS7{F0: v5, F1: (v5 + (v5 ^ (v5 / ((v5) | 1))))}
                                        _ = v52
                                    }
                                    if false {
                                        v18 -= 42
                                        for fi53 := 0; fi53 < 18; fi53++ {
                                            v54 := GenS8{F0: 1}
                                            v17 -= 32767
                                            _ = fi53
                                            _ = v54
                                        }
                                        for fi55 := 0; fi55 < 17; fi55++ {
                                            v16 += v16
                                            v56 := GenS3{F0: 2, F1: (v16 - (v16 & v16)), F2: v16}
                                            _ = fi55
                                            _ = v56
                                        }
                                    }
                                }
                                if (!true) {
                                    for fi57 := 0; fi57 < 8; fi57++ {
                                        v37 = (0.1 + -2.25)
                                        _ = fi57
                                    }
                                    if (!true) {
                                        var v58 uint32 = (v16 + (v16 + (^(v16))))
                                        fmt.Println(v17)
                                        _ = v58
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v59 uint8 = 255
                                        _ = v59
                                    }
                                    var v60 string = "\n\t"
                                    _ = v60
                                }
                            } else {
                                v5 = v5
                                if true {
                                    var v61 uint32 = v16
                                    _ = v61
                                } else {
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        v4[-1] = (v16 >> ((255) & 63))
                                        var v62 uint = v18
                                        v62 = v18
                                        _ = v62
                                    }
                                    var v63 int16 = 42
                                    _ = v63
                                }
                                _ = len(v3)
                            }
                        }
                        if false {
                            for fi64 := 0; fi64 < 15; fi64++ {
                                _ = len(v3)
                                for fi65 := 0; fi65 < 20; fi65++ {
                                    var v66 int16 = -32768
                                    switch (("" + "") + ("a" + (((("" + "hello") + "") + "hello") + "hello"))) {
                                    case "":
                                        v67 := v6[(((("foo bar" + "foo bar") + "foo bar") + "hello") + "hello")]
                                        _ = v67
                                    case "\n\t":
                                        switch ("\n\t" + "foo bar") {
                                        case "foo bar":
                                            v68 := GenS6{F0: v5, F1: 1}
                                            _ = v68
                                        case "":
                                            fmt.Println(v16, v17)
                                        default:
                                            v69 := GenS7{F0: 4294967295, F1: (v5 / ((745) | 1))}
                                            _ = v69
                                        }
                                    default:
                                        v66 -= v66
                                    }
                                    _ = len(v4)
                                    _ = fi65
                                    _ = v66
                                }
                                _ = fi64
                            }
                        }
                    } else {
                        _ = len(v6)
                        v70 := GenS5{F0: -1, F1: 0, F2: -32768}
                        v1[v17] = (v16 + 2147483648)
                        _ = v70
                    }
                    _ = v37
                }
            }
            if false {
                switch (^(v17)) {
                case 9223372036854775807:
                    switch v17 {
                    case -296:
                        v18 -= (v18 << ((v18) & 63))
                    }
                }
            }
        }
    } else {
        for fi71 := 0; fi71 < 13; fi71++ {
            var v72 uint = 85
            var v73 int32 = 2147483647
            if (!(32767 <= 209)) {
                if (false || (151 == -9223372036854775808)) {
                    for fi74 := 0; fi74 < 15; fi74++ {
                        var v75 float32 = ((0.0 - -2.25) / 1e10)
                        var v76 float64 = float64(v75)
                        v16 -= v16
                        var v77 float32 = float32(v76)
                        _ = fi74
                        _ = v75
                        _ = v76
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 11; fi78++ {
                        if false {
                            var v79 string = ""
                            switch (v79 + "\n\t") {
                            case "\n\t":
                                if ((v16 >> ((42) & 63)) < v16) {
                                    v80 := GenS7{F0: 42, F1: v5}
                                    v72 -= 799
                                    if (520 >= (v73 ^ (-(v73)))) {
                                        v17 = fi78
                                        if (309 == -1) {
                                            v18 = v72
                                        }
                                        if false {
                                            if (255 <= 2) {
                                                var v81 int32 = (v73 % ((2147483647) | 1))
                                                if (65535 <= (v18 % ((v72) | 1))) {
                                                    v5 = (v5 | v5)
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                    var v82 int16 = -32768
                                                    _ = v82
                                                }
                                                _ = v81
                                            }
                                        }
                                        var v83 int = fi71
                                        v84 := GenS9{F0: 658, F1: (-(v73))}
                                        _ = v83
                                        _ = v84
                                    } else {
                                        v17 = v17
                                        v18 = v18
                                    }
                                    _ = v80
                                } else {
                                    var v85 uint32 = 2147483647
                                    if false {
                                        v79 = ""
                                    } else {
                                        v18 += (^(v72))
                                    }
                                    _ = v85
                                }
                            case "":
                                var v86 float32 = (0.1 - 3.14159)
                                _ = v86
                            case "foo bar":
                                if (!(((fi71 & fi71) != -1) || false)) {
                                    if len(v3) > 0 {
                                        v3[0] = (^(v16))
                                    }
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = v16
                                        }
                                    } else {
                                        v72 = 0
                                    }
                                    switch v79 {
                                    case "hello":
                                        if false {
                                            var v87 bool = ((1 < 1) || (v5 >= (^(v5))))
                                            _ = v87
                                        }
                                    case "":
                                        v18 -= (v72 / ((v18) | 1))
                                    default:
                                        var v88 uint8 = uint8(v5)
                                        _ = v88
                                    }
                                } else {
                                    fmt.Println(v73)
                                    v4[fi71] = (v16 ^ (v16 * (^(v16))))
                                    var v89 int16 = -427
                                    var v90 float32 = 1e10
                                    var v91 int64 = 0
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                            }
                            var v92 float32 = 1e10
                            _ = v79
                            _ = v92
                        } else {
                            v93 := GenS10{F0: (v17 | v17), F1: -327, F2: (v18 % ((4294967295) | 1))}
                            _ = v93
                        }
                        fmt.Println(v16)
                        _ = fi78
                    }
                } else {
                    for fi94 := 0; fi94 < 8; fi94++ {
                        if ((!(18446744073709551615 > (v5 | 42))) || true) {
                            v95 := GenS4{F0: 1000, F1: 4294967295}
                            v5 = v5
                            if true {
                                v96 := GenS11{F0: (v73 << ((2) & 63))}
                                delete(v1, (-(fi94)))
                                v18 = (v18 ^ v18)
                                v97 := GenS12{F0: uint32(v5), F1: v16}
                                switch "" {
                                case "hello":
                                    var v98 int = (fi71 >> ((44) & 63))
                                    _ = v98
                                }
                                _ = v96
                                _ = v97
                            } else {
                                fmt.Println(v18, fi71)
                            }
                            var v99 string = "a"
                            var v100 int = 9223372036854775807
                            switch "\n\t" {
                            case "hello":
                                v101 := GenS12{F0: v16, F1: v16}
                                _ = v101
                            }
                            var v102 int16 = 42
                            switch fi94 {
                            case 1:
                                if (115 == 128) {
                                    if ((v99 + ((("foo bar" + "foo bar") + v99) + (("foo bar" + (v99 + v99)) + "hello"))) >= v99) {
                                        _ = len(v3)
                                    } else {
                                        v102 += (v102 << (((v18 + v18)) & 63))
                                        v99 = v99
                                        v103 := GenS8{F0: -40}
                                        fmt.Println(v73)
                                        _ = v103
                                    }
                                } else {
                                    v104 := GenS13{F0: v73}
                                    _ = v104
                                }
                            case 127:
                                if true {
                                    delete(v6, "\n\t")
                                    var v105 uint32 = (v16 & 2)
                                    switch v99 {
                                    case "\n\t":
                                        var v106 uint64 = (v5 >> ((v18) & 63))
                                        _ = v106
                                    }
                                    _ = v105
                                }
                            case -9223372036854775808:
                                fmt.Println(fi94, v73, v17)
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v107 := v1[(fi94 & fi94)]
                            _ = v100
                            _ = v102
                            _ = v107
                            _ = v95
                            _ = v99
                        } else {
                            v72 -= v18
                            v17 = (-(fi94))
                        }
                        switch v17 {
                        case -2147483648:
                            if ((1 > 42) || ((v73 | v73) == 42)) {
                                switch "" {
                                case "a":
                                    switch (^(fi94)) {
                                    case -889:
                                        var v108 int = fi71
                                        _ = v108
                                    default:
                                        delete(v2, 563)
                                    }
                                default:
                                    v73 = (v73 - -1)
                                }
                                var v109 bool = (((^(v5)) >= 18446744073709551615) && false)
                                _ = v109
                            }
                        case 4294967295:
                            switch fi94 {
                            case 127:
                                v110 := GenS12{F0: 0, F1: (v16 << ((v72) & 63))}
                                _ = v110
                            default:
                                if true {
                                    var v111 float64 = ((((((((-2.25 + float64(-2.25)) * float64(0.1)) / float64(-2.25)) * 1e10) / 1.0) - (float64(1e10) / (0.1 + float64(1e10)))) / 0.1) / 1.0)
                                    v17 = fi94
                                    var v112 uint32 = 161
                                    v111 += 1e10
                                    v111 = v111
                                    _ = v111
                                    _ = v112
                                }
                            }
                        }
                        _ = fi94
                    }
                    for fi113 := 0; fi113 < 4; fi113++ {
                        var v114 float64 = 3.14159
                        var v115 string = "hello"
                        v17 -= (fi71 * (-(fi113)))
                        var v116 int8 = 42
                        v18 += v18
                        _ = fi113
                        _ = v114
                        _ = v115
                        _ = v116
                    }
                    for fi117 := 0; fi117 < 3; fi117++ {
                        v16 = 4294967295
                        v118 := GenS9{F0: -1, F1: (v73 ^ v73)}
                        _ = fi117
                        _ = v118
                    }
                }
                var v119 float32 = 1e10
                _ = v119
            }
            _ = fi71
            _ = v72
            _ = v73
        }
    }
    v120 := GenS14{F0: (v18 / ((762) | 1)), F1: ("hello" + "foo bar")}
    for fi121 := 0; fi121 < 8; fi121++ {
        v18 = v18
        var v122 int8 = 1
        v123 := GenS15{F0: genFn0(v18, 4294967295)}
        _ = fi121
        _ = v122
        _ = v123
    }
    delete(v4, v17)
    v5 = v5
    var v124 bool = (v5 == (v5 >> ((v18) & 63)))
    v16 -= v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v120
    _ = v124
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
