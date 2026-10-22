package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 string; F1 int8; F2 int64 }
type GenS6 struct { F0 int8; F1 int32; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 int16; F1 uint8; F2 uint }
type GenS12 struct { F0 int64; F1 int64 }
type GenS13 struct { F0 int64; F1 int16 }
type GenS14 struct { F0 uint }
func genFn0(p0 int64, p1 uint8) int32 {
    fmt.Println(p0, p0)
    v5 := GenS1{F0: 4294967295, F1: 4294967295, F2: 42}
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{}
        for fi8 := 0; fi8 < 12; fi8++ {
            switch 32767 {
            case -817:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case 2147483647:
                v9 := GenS1{F0: 42, F1: 65535, F2: 4294967295}
                _ = v9
            case 42:
                v10 := GenS1{F0: 0, F1: 2, F2: 42}
                _ = v10
            }
            v11 := GenS1{F0: 42, F1: 65535, F2: 42}
            if true {
                if true {
                    fmt.Println(p1, fi8)
                    var v12 uint32 = 4294967295
                    var v13 float32 = ((3.14159 / 1.0) + 1.5)
                    _ = v12
                    _ = v13
                } else {
                    v7 = append(v7, -1)
                }
                switch fi8 {
                case 116:
                    var v14 float32 = 3.14159
                    _ = v14
                }
                v7 = append(v7, fi6)
            } else {
                if false {
                    v15 := GenS2{F0: p0, F1: -389}
                    if len(v7) > 0 {
                        v7[0] = fi8
                    }
                    _ = v15
                }
            }
            _ = fi8
            _ = v11
        }
        switch fi6 {
        case -963:
            v16 := GenS1{F0: 42, F1: 1, F2: uint64(p1)}
            _ = v16
        case 945:
            for fi17 := 0; fi17 < 15; fi17++ {
                if false {
                    var v18 int = (fi17 >> ((18446744073709551615) & 63))
                    v19 := GenS2{F0: (p0 & (p0 ^ 373)), F1: 1000}
                    _ = v18
                    _ = v19
                } else {
                    p0 = (p0 % ((p0) | 1))
                    var v20 int64 = (-(p0))
                    _ = v20
                }
                p0 -= (^(p0))
                var v21 float32 = float32(1.5)
                var v22 uint64 = 42
                _ = fi17
                _ = v21
                _ = v22
            }
        case 0:
            _ = len(v7)
        default:
            fmt.Println(p1, fi6)
        }
        _ = fi6
        _ = v7
    }
    switch -509 {
    case 32768:
        p0 = (p0 * p0)
    }
    p0 = (p0 | 760)
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1() int8 {
    if true {
        if false {
            for fi23 := 0; fi23 < 16; fi23++ {
                for fi24 := 0; fi24 < 19; fi24++ {
                    if ((true && (528 >= 255)) && false) {
                        var v25 uint16 = 65535
                        v25 += v25
                        if ((65535 != 1) && (0 != 65535)) {
                            v25 = v25
                            fmt.Println(fi24, v25)
                        } else {
                            v25 = 65535
                            v25 = v25
                        }
                        _ = v25
                    } else {
                        v26 := GenS1{F0: 2, F1: 2147483648, F2: 18446744073709551615}
                        if true {
                            var v27 bool = (0 <= 42)
                            var v28 uint32 = 65535
                            _ = v27
                            _ = v28
                        }
                        var v29 float32 = float32(0.0)
                        _ = v26
                        _ = v29
                    }
                    v30 := GenS1{F0: 4294967295, F1: 502, F2: 0}
                    v31 := []int{}
                    _ = fi24
                    _ = v30
                    _ = v31
                }
                _ = fi23
            }
            var v32 string = "foo bar"
            _ = v32
        } else {
            var v33 int64 = 1
            v34 := GenS2{F0: (v33 & v33), F1: 154}
            fmt.Println(v33)
            _ = v33
            _ = v34
        }
        var v35 int8 = 42
        v35 -= (v35 / ((v35) | 1))
        _ = v35
    }
    var v36 uint = 42
    fmt.Println(v36)
    for fi37 := 0; fi37 < 10; fi37++ {
        if (964 >= -658) {
            var v38 float64 = (1.5 + -2.25)
            if ("foo bar" > "") {
                v39 := []int{128, (^(fi37)), 1}
                if true {
                    for fi40 := 0; fi40 < 3; fi40++ {
                        v36 -= (v36 >> ((4294967295) & 63))
                        _ = fi40
                    }
                    if len(v39) > 0 {
                        v39[0] = -1
                    }
                }
                _ = v39
            }
            _ = v38
        } else {
            var v41 string = (("hello" + "\n\t") + ("" + ("\n\t" + (("\n\t" + ("" + ("\n\t" + "a"))) + "a"))))
            _ = v41
        }
        v36 += v36
        for fi42 := 0; fi42 < 5; fi42++ {
            v43 := []int{(-(fi37)), -1, (^(fi37)), fi37}
            _ = fi42
            _ = v43
        }
        _ = fi37
    }
    _ = v36
    return 110
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.1
    var v3 bool = false
    v4 := []string{(("\n\t" + "a") + "")}
    if false {
        var v44 uint64 = 42
        if (genFn1() < 42) {
            if v3 {
                var v45 bool = true
                _ = v45
            }
        }
        var v46 float64 = (1e10 + 0.0)
        _ = v44
        _ = v46
    }
    v47 := GenS1{F0: 255, F1: 42, F2: 751}
    for fi48 := 0; fi48 < 10; fi48++ {
        if false {
            var v49 int64 = -9223372036854775808
            switch "\n\t" {
            case "hello":
                var v50 int64 = 11
                _ = v50
            case "a":
                if (genFn1() <= 1) {
                    var v51 uint = 42
                    for fi52 := 0; fi52 < 1; fi52++ {
                        _ = len(v1)
                        var v53 uint8 = 255
                        fmt.Println(fi52, v51, v53)
                        _ = fi52
                        _ = v53
                    }
                    fmt.Println(v49)
                    _ = v51
                } else {
                    v49 += v49
                    v2 = float32(-2.25)
                    v4 = append(v4, "hello")
                }
            }
            if v3 {
                v3 = v3
            }
            _ = v49
        } else {
            var v54 string = ((("foo bar" + "a") + ("\n\t" + ("\n\t" + "\n\t"))) + "hello")
            _ = v54
        }
        if v3 {
            var v55 float64 = (1.5 * 1.5)
            _ = v55
        } else {
            v3 = v3
            v2 -= v2
        }
        if (66 < 0) {
            _ = len(v1)
            for fi56 := 0; fi56 < 5; fi56++ {
                switch fi56 {
                case 1:
                    var v57 bool = v3
                    _ = v57
                default:
                    _ = len(v1)
                }
                var v58 uint32 = 209
                switch "hello" {
                case "\n\t":
                    v59 := GenS1{F0: 2, F1: (v58 ^ (v58 << ((330) & 63))), F2: 18446744073709551615}
                    _ = v59
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "hello":
                    var v60 uint8 = 128
                    _ = v60
                default:
                    v61 := GenS1{F0: 892, F1: v58, F2: 2}
                    _ = v61
                }
                _ = fi56
                _ = v58
            }
        } else {
            for fi62 := 0; fi62 < 1; fi62++ {
                delete(v1, 1)
                v3 = v3
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                _ = fi62
            }
            v3 = (!true)
            if v3 {
                if v3 {
                    fmt.Println(v3, v2, fi48)
                    for fi63 := 0; fi63 < 4; fi63++ {
                        v3 = (!v3)
                        _ = fi63
                    }
                    v1[1] = 32767
                } else {
                    if (1 != 1) {
                        var v64 bool = true
                        for fi65 := 0; fi65 < 5; fi65++ {
                            if v3 {
                                switch "" {
                                case "a":
                                    var v66 int = (fi65 | 32767)
                                    _ = v66
                                default:
                                    var v67 uint8 = 128
                                    _ = v67
                                }
                            } else {
                                if true {
                                    v3 = (74 <= 0)
                                    delete(v1, 65535)
                                    v3 = false
                                    var v68 uint8 = 2
                                    _ = v68
                                } else {
                                    v69 := GenS2{F0: -850, F1: -32768}
                                    var v70 int16 = -32768
                                    _ = v69
                                    _ = v70
                                }
                                switch "a" {
                                case "a":
                                    _ = len(v1)
                                case "hello":
                                    _ = len(v4)
                                default:
                                    switch fi48 {
                                    case -899:
                                        if (71 < 2147483648) {
                                            v71 := GenS3{F0: 653}
                                            _ = v71
                                        }
                                    }
                                }
                                _ = len(v4)
                            }
                            v72 := GenS3{F0: 2147483648}
                            v73 := GenS2{F0: -704, F1: -671}
                            _ = fi65
                            _ = v72
                            _ = v73
                        }
                        _ = v64
                    }
                    v74 := GenS4{F0: 42, F1: "hello"}
                    _ = v74
                }
            } else {
                v75 := GenS5{F0: "", F1: 100, F2: -9223372036854775808}
                if (59 == 206) {
                    for fi76 := 0; fi76 < 4; fi76++ {
                        if (true || true) {
                            v2 = -2.25
                            var v77 string = ("a" + "foo bar")
                            var v78 uint64 = 42
                            _ = v77
                            _ = v78
                        }
                        _ = fi76
                    }
                } else {
                    var v79 int = 1
                    if v3 {
                        v80 := GenS6{F0: 88, F1: genFn0(-1, 255), F2: true}
                        fmt.Println(fi48)
                        var v81 int8 = -49
                        _ = v80
                        _ = v81
                    }
                    fmt.Println(fi48)
                    _ = v79
                }
                switch (("" + "") + ((("a" + "hello") + "\n\t") + "")) {
                case "":
                    var v82 uint = 18446744073709551615
                    _ = v82
                }
                _ = v75
            }
        }
        _ = fi48
    }
    v83 := GenS4{F0: -96, F1: "\n\t"}
    v4 = append(v4, "\n\t")
    v84 := v1[2147483647]
    if (0.0 >= v2) {
        if (49 < 98) {
            switch (v84 ^ v84) {
            case 127:
                v85 := GenS7{F0: int32(v84)}
                _ = v85
            case 42:
                v86 := GenS2{F0: -9223372036854775808, F1: 1}
                _ = v86
            case -2147483648:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            v4 = append(v4, "foo bar")
            for fi87 := 0; fi87 < 1; fi87++ {
                v2 = (v2 + v2)
                for fi88 := 0; fi88 < 5; fi88++ {
                    v89 := GenS8{F0: 32767}
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    var v90 int32 = 2147483647
                    var v91 uint32 = 502
                    _ = fi88
                    _ = v89
                    _ = v90
                    _ = v91
                }
                var v92 uint = 368
                _ = fi87
                _ = v92
            }
            v2 = 1.5
        }
        v93 := GenS6{F0: -70, F1: 330, F2: true}
        v94 := GenS5{F0: "", F1: 1, F2: 42}
        _ = v93
        _ = v94
    } else {
        switch ("" + "a") {
        case "foo bar":
            if (1 > 0) {
                v84 -= v84
            } else {
                var v95 int = (v84 & (-(v84)))
                if (!true) {
                    v2 -= (v2 - float32(0.0))
                }
                _ = v95
            }
        case "\n\t":
            v96 := GenS9{F0: 4294967295, F1: -755}
            _ = v96
        }
        for fi97 := 0; fi97 < 16; fi97++ {
            v4 = append(v4, "foo bar")
            for fi98 := 0; fi98 < 11; fi98++ {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                var v99 uint8 = 176
                fmt.Println(v3, v3, v99)
                v99 = (v99 / ((v99) | 1))
                switch ("" + ("" + ((("a" + "\n\t") + "hello") + ((("\n\t" + "\n\t") + "hello") + "a")))) {
                case "foo bar":
                    switch -9223372036854775808 {
                    case 42:
                        fmt.Println(fi97, v2, v2)
                    default:
                        v100 := GenS10{F0: 106, F1: 326}
                        _ = v100
                    }
                }
                _ = fi98
                _ = v99
            }
            _ = fi97
        }
        for fi101 := 0; fi101 < 6; fi101++ {
            v1[0] = fi101
            switch "a" {
            case "hello":
                switch ("\n\t" + "foo bar") {
                case "":
                    switch (v84 | (v84 & v84)) {
                    case 2147483647:
                        v2 = (float32(0.1) - 3.14159)
                    default:
                        var v102 bool = v3
                        _ = v102
                    }
                }
            case "":
                v103 := GenS11{F0: 1000, F1: 0, F2: 14}
                _ = v103
            case "\n\t":
                for fi104 := 0; fi104 < 1; fi104++ {
                    if v3 {
                        v84 += 683
                        var v105 string = "a"
                        fmt.Println(v84)
                        _ = v105
                    } else {
                        var v106 int8 = 48
                        v2 = (3.14159 - 1.5)
                        v84 = -358
                        _ = v106
                    }
                    _ = fi104
                }
            default:
                v3 = (65535 < 2147483648)
            }
            for fi107 := 0; fi107 < 20; fi107++ {
                v2 -= (v2 * 0.1)
                v108 := GenS8{F0: -1}
                _ = fi107
                _ = v108
            }
            _ = fi101
        }
    }
    var v109 uint16 = 2
    v110 := GenS12{F0: -561, F1: -1}
    v3 = (1 < 42)
    switch v84 {
    case 32768:
        _ = len(v4)
    case 810:
        if len(v4) > 0 {
            v4[0] = ("a" + (("\n\t" + ("foo bar" + "")) + ("\n\t" + "")))
        }
    case 583:
        var v111 uint = 2
        _ = v111
    }
    if ((!(!(0 < 0))) && (-2147483648 >= 1)) {
        if v3 {
            var v112 uint = 925
            _ = v112
        } else {
            v2 += (-2.25 + 1e10)
            v3 = (!v3)
            v113 := GenS13{F0: -815, F1: 1}
            _ = v113
        }
        v109 += (^(v109))
    } else {
        if (!true) {
            if true {
                var v114 uint64 = 18446744073709551615
                v2 += 1.5
                _ = v114
            } else {
                if v3 {
                    delete(v1, 0)
                    _ = len(v4)
                } else {
                    v1[2] = v84
                    v1[512] = (v84 >> ((255) & 63))
                }
            }
        }
        var v115 int8 = -6
        if (uint64(v109) <= 42) {
            switch 4294967295 {
            case -9223372036854775808:
                if ("hello" != ("a" + "")) {
                    v116 := GenS8{F0: -1}
                    fmt.Println(v84)
                    _ = v116
                } else {
                    var v117 int32 = 1
                    v1[2147483647] = 128
                    _ = v117
                }
            case 32767:
                fmt.Println(v109, v115, v115)
            case -2147483648:
                switch (v84 * v84) {
                case -2147483648:
                    v118 := GenS14{F0: 18446744073709551615}
                    _ = v118
                case 32767:
                    for fi119 := 0; fi119 < 14; fi119++ {
                        var v120 uint32 = 210
                        v109 += (v109 & v109)
                        switch "hello" {
                        case "":
                            var v121 int16 = 42
                            _ = v121
                        case "\n\t":
                            switch (v84 << ((0) & 63)) {
                            case 128:
                                var v122 string = "foo bar"
                                _ = v122
                            case 0:
                                var v123 int64 = 42
                                _ = v123
                            case 9223372036854775807:
                                v1[(v120 - v120)] = 1
                            }
                        case "hello":
                            v109 += 255
                        }
                        _ = fi119
                        _ = v120
                    }
                }
            default:
                v2 += (v2 + -2.25)
            }
        }
        _ = v115
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v83
    _ = v84
}
