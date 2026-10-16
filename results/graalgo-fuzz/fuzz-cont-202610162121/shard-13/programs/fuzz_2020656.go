package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS8 struct { F0 int32; F1 bool; F2 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 uint16 }
func genFn0() int16 {
    switch "foo bar" {
    case "foo bar":
        if false {
            var v5 bool = (false || (!(!false)))
            _ = v5
        } else {
            var v6 int32 = -125
            v7 := GenS1{F0: 517, F1: -1}
            _ = v6
            _ = v7
        }
    case "a":
        var v8 int8 = 0
        _ = v8
    }
    switch "hello" {
    case "hello":
        var v9 int = -2147483648
        _ = v9
    }
    return 32767
}
func genFn1(p0 int16, p1 uint32, p2 int8) int32 {
    v10 := []int{0}
    var v11 bool = false
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := []string{("a" + "\n\t"), (("foo bar" + "\n\t") + (("" + ("\n\t" + "hello")) + "\n\t")), "a", (("" + "\n\t") + ("" + (("" + "") + (("foo bar" + "\n\t") + "\n\t"))))}
    var v4 int32 = -2147483648
    v12 := GenS2{F0: 0, F1: 94, F2: 689}
    var v13 uint = 0
    v14 := GenS3{F0: -1, F1: (v4 ^ v4), F2: 31}
    switch "a" {
    case "foo bar":
        var v15 bool = false
        _ = v15
    case "\n\t":
        var v16 uint32 = 4294967295
        _ = v16
    default:
        var v17 bool = true
        _ = v17
    }
    var v18 int8 = 127
    v19 := GenS3{F0: 1, F1: genFn1(-347, 2147483647, v18), F2: 83}
    v20 := GenS1{F0: 42, F1: 55}
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 int8 = (v18 & v18)
        switch ((("hello" + "a") + ((("\n\t" + "hello") + "") + ("hello" + ("" + "hello")))) + ("a" + ("foo bar" + ("hello" + "")))) {
        case "":
            for fi23 := 0; fi23 < 6; fi23++ {
                switch "" {
                case "":
                    switch fi21 {
                    case -2147483648:
                        if true {
                            switch (fi21 + (fi21 - (fi21 & (fi21 / ((fi23) | 1))))) {
                            case -70:
                                if (false || true) {
                                    v22 -= v18
                                    v24 := GenS4{F0: 42, F1: v13, F2: 1}
                                    v25 := v2[1]
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 int8 = 21
                                    var v27 int = fi21
                                    _ = v26
                                    _ = v27
                                }
                            case 127:
                                if (true || false) {
                                    var v28 float32 = ((float32(1.5) * (0.0 - 0.1)) + (0.0 * float32(3.14159)))
                                    v18 += 74
                                    switch 32 {
                                    case 519:
                                        if (false && false) {
                                            fmt.Println(v28, v18)
                                        } else {
                                            var v29 int32 = -1
                                            v29 = (v29 & genFn1(genFn0(), 4294967295, v22))
                                            _ = v29
                                        }
                                    case -157:
                                        v28 = 0.0
                                        if true {
                                            var v30 int64 = 42
                                            var v31 bool = (((true && (0.0 != 0.1)) || true) || true)
                                            _ = v30
                                            _ = v31
                                        } else {
                                            _ = len(v1)
                                            if len(v3) > 0 {
                                                v3[0] = "a"
                                            }
                                        }
                                        v32 := GenS3{F0: (v4 / ((v4) | 1)), F1: (^(v4)), F2: 2}
                                        _ = v32
                                    case 32767:
                                        v4 = 2147483647
                                    default:
                                        v3 = append(v3, ("\n\t" + "a"))
                                    }
                                    _ = v28
                                }
                            default:
                                v18 = v18
                            }
                        }
                    case 265:
                        var v33 float64 = ((3.14159 + 1e10) + (float64(1e10) / float64(1.5)))
                        _ = v33
                    default:
                        var v34 float64 = (1e10 / 1e10)
                        _ = v34
                    }
                    v35 := v2[4294967295]
                    _ = len(v2)
                    _ = v35
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    var v36 int64 = 1
                    _ = v36
                default:
                    v22 = -1
                }
                switch 127 {
                case 0:
                    v22 = -128
                case 32767:
                    v37 := GenS3{F0: v4, F1: (v4 | v4), F2: 128}
                    _ = v37
                case -818:
                    var v38 int32 = v4
                    _ = v38
                }
                v39 := v2[733]
                _ = fi23
                _ = v39
            }
        case "a":
            fmt.Println(v22, v4, fi21)
        default:
            v22 = v22
        }
        _ = fi21
        _ = v22
    }
    switch 4294967295 {
    case 128:
        v4 += (v4 % (((v4 & (v4 * (-(v4))))) | 1))
    case 32767:
        if (float64(0.1) > (0.1 / 0.1)) {
            switch -245 {
            case -1:
                switch 128 {
                case -2147483648:
                    for fi40 := 0; fi40 < 9; fi40++ {
                        fmt.Println(fi40, v13)
                        _ = fi40
                    }
                case -25:
                    for fi41 := 0; fi41 < 17; fi41++ {
                        var v42 float32 = (3.14159 - 1e10)
                        v43 := GenS2{F0: 1, F1: (v18 << (((v13 * 65535)) & 63)), F2: 348}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                case 1:
                    fmt.Println(v18)
                default:
                    if false {
                        var v44 uint64 = 4294967295
                        _ = v44
                    } else {
                        var v45 uint64 = 42
                        if (417 <= 42) {
                            var v46 uint16 = 65535
                            switch -373 {
                            case 127:
                                if (65535 > 1) {
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    v4 = v4
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                }
                            default:
                                var v47 int64 = 0
                                _ = v47
                            }
                            var v48 float32 = 0.0
                            _ = v46
                            _ = v48
                        }
                        _ = v45
                    }
                }
            default:
                for fi49 := 0; fi49 < 17; fi49++ {
                    var v50 bool = (!false)
                    _ = fi49
                    _ = v50
                }
            }
            switch "" {
            case "":
                v4 = (v4 - (v4 + genFn1(0, 777, -128)))
            default:
                v51 := GenS4{F0: 595, F1: v13, F2: 42}
                _ = v51
            }
            v4 = -1
        }
    case -605:
        v52 := GenS1{F0: 614, F1: 1}
        _ = v52
    default:
        for fi53 := 0; fi53 < 8; fi53++ {
            switch (fi53 * fi53) {
            case 0:
                switch "a" {
                case "a":
                    v54 := GenS5{F0: 1, F1: true}
                    _ = v54
                case "\n\t":
                    var v55 string = ("" + "foo bar")
                    _ = v55
                default:
                    var v56 bool = true
                    _ = v56
                }
            case -61:
                var v57 uint = v13
                _ = v57
            default:
                switch (fi53 >> ((65535) & 63)) {
                case 0:
                    v13 = 42
                default:
                    var v58 int64 = -9223372036854775808
                    _ = v58
                }
            }
            switch (fi53 - (fi53 ^ (fi53 - (-(fi53))))) {
            case 4294967295:
                v13 = (^(v13))
            default:
                _ = len(v3)
            }
            _ = fi53
        }
    }
    if true {
        v4 += 0
        var v59 uint = (^(v13))
        _ = v59
    }
    if (!true) {
        if (float64(0.0) < 3.14159) {
            v18 += v18
            for fi60 := 0; fi60 < 13; fi60++ {
                v18 -= 39
                v4 -= 0
                _ = fi60
            }
            v3 = append(v3, "hello")
        }
        switch (("hello" + "\n\t") + "foo bar") {
        case "":
            v61 := GenS1{F0: 32767, F1: 1}
            _ = v61
        }
        v13 += v13
    }
    for fi62 := 0; fi62 < 2; fi62++ {
        for fi63 := 0; fi63 < 17; fi63++ {
            v64 := GenS6{F0: v4}
            switch ("hello" + "\n\t") {
            case "hello":
                v13 -= (^(v13))
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                fmt.Println(v13, v13)
            }
            v18 -= v18
            _ = fi63
            _ = v64
        }
        _ = fi62
    }
    for fi65 := 0; fi65 < 10; fi65++ {
        fmt.Println(fi65, fi65, fi65)
        for fi66 := 0; fi66 < 13; fi66++ {
            if false {
                switch 127 {
                case 127:
                    var v67 uint32 = 375
                    _ = v67
                case 9223372036854775807:
                    var v68 bool = (true || false)
                    _ = v68
                case 32767:
                    var v69 float32 = (0.1 - (float32(-2.25) * 1.5))
                    _ = v69
                }
            } else {
                v4 = genFn1(-32768, 86, 42)
            }
            var v70 uint32 = 2147483648
            var v71 uint8 = 127
            _ = fi66
            _ = v70
            _ = v71
        }
        _ = fi65
    }
    if (!(-2147483648 < 9223372036854775807)) {
        if false {
            for fi72 := 0; fi72 < 8; fi72++ {
                if false {
                    var v73 uint16 = 311
                    var v74 uint = v13
                    _ = v73
                    _ = v74
                } else {
                    switch int(v4) {
                    case -1:
                        v75 := GenS7{F0: 65535, F1: 42, F2: 676}
                        _ = v75
                    case -916:
                        delete(v1, 4294967295)
                    default:
                        if false {
                            var v76 string = "foo bar"
                            switch "" {
                            case "\n\t":
                                switch (fi72 ^ fi72) {
                                case -890:
                                    fmt.Println(v4)
                                case -2147483648:
                                    var v77 int16 = 1
                                    _ = v77
                                case 210:
                                    for fi78 := 0; fi78 < 14; fi78++ {
                                        v18 -= -81
                                        if ((true || false) && true) {
                                            if len(v3) > 0 {
                                                v3[0] = v76
                                            }
                                            if (fi78 != fi72) {
                                                v13 -= v13
                                                fmt.Println(fi78)
                                            } else {
                                                if ((v4 | 2147483647) >= (-(v4))) {
                                                    v18 -= v18
                                                    v79 := GenS8{F0: 0, F1: false, F2: 127}
                                                    _ = v79
                                                }
                                            }
                                        }
                                        var v80 bool = true
                                        _ = fi78
                                        _ = v80
                                    }
                                }
                            case "":
                                var v81 int8 = (-(v18))
                                _ = v81
                            default:
                                v82 := GenS3{F0: v4, F1: v4, F2: 1}
                                _ = v82
                            }
                            _ = v76
                        }
                    }
                    var v83 uint = v13
                    for fi84 := 0; fi84 < 15; fi84++ {
                        var v85 int8 = int8(v83)
                        v83 = 951
                        switch (("hello" + "foo bar") + "\n\t") {
                        case "foo bar":
                            if true {
                                v86 := GenS6{F0: 1}
                                var v87 uint = 255
                                _ = v86
                                _ = v87
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    var v88 int64 = 1
                                    _ = v88
                                case "\n\t":
                                    v3 = append(v3, "\n\t")
                                case "hello":
                                    switch ("" + ("a" + "")) {
                                    case "":
                                        v18 = (v18 ^ 56)
                                    case "hello":
                                        if false {
                                            var v89 uint8 = 126
                                            _ = v89
                                        } else {
                                            var v90 uint8 = 223
                                            delete(v1, 707)
                                            _ = v90
                                        }
                                    case "\n\t":
                                        v13 = v83
                                    }
                                default:
                                    v91 := GenS9{F0: 4294967295}
                                    _ = v91
                                }
                                v3 = append(v3, (("foo bar" + "\n\t") + ("foo bar" + ("foo bar" + "\n\t"))))
                                v85 = 58
                            }
                        }
                        _ = fi84
                        _ = v85
                    }
                    _ = v83
                }
                switch (-(fi72)) {
                case 127:
                    for fi92 := 0; fi92 < 4; fi92++ {
                        fmt.Println(v13, v18)
                        _ = fi92
                    }
                case -1:
                    if ((!true) && (true || (255 < 278))) {
                        v18 = 0
                        for fi93 := 0; fi93 < 9; fi93++ {
                            var v94 int8 = (v18 ^ 42)
                            _ = fi93
                            _ = v94
                        }
                        v4 = (^(v4))
                    }
                case -708:
                    v95 := v2[0]
                    _ = v95
                }
                for fi96 := 0; fi96 < 10; fi96++ {
                    var v97 uint8 = 27
                    v98 := GenS10{F0: 0, F1: 65535}
                    switch 542 {
                    case -563:
                        v99 := GenS6{F0: v4}
                        _ = v99
                    case 128:
                        var v100 uint32 = 2147483648
                        _ = v100
                    case -1:
                        var v101 string = "foo bar"
                        _ = v101
                    }
                    _ = fi96
                    _ = v97
                    _ = v98
                }
                _ = fi72
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v102 bool = true
            _ = v102
        }
        v103 := GenS5{F0: 128, F1: (!true)}
        delete(v1, 65535)
        _ = v103
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
