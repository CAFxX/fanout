package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int32; F1 int32; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 string; F1 int; F2 int32 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 string; F1 string; F2 bool }
type GenS13 struct { F0 int8; F1 uint; F2 uint }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 string; F1 string }
type GenS16 struct { F0 string; F1 bool; F2 int8 }
type GenS17 struct { F0 uint32; F1 string; F2 uint64 }
type GenS18 struct { F0 uint32; F1 uint64 }
type GenS19 struct { F0 uint8; F1 string }
type GenS20 struct { F0 int32; F1 bool }
type GenS21 struct { F0 int64; F1 bool; F2 uint16 }
type GenS22 struct { F0 uint64; F1 int64 }
type GenS23 struct { F0 string; F1 int; F2 uint8 }
type GenS24 struct { F0 uint; F1 int16; F2 uint32 }
type GenS25 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0() int8 {
    var v6 int = 32768
    v6 = v6
    v6 += 4294967295
    v6 = 1
    _ = v6
    return 1
}

func main() {
    var v1 bool = (!(!true))
    var v2 int8 = -128
    var v3 int8 = (-(v2))
    var v4 string = "foo bar"
    var v5 float64 = 0.1
    if v1 {
        var v7 int64 = -534
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(v5, v3, v2)
            v4 = "hello"
            _ = fi8
        }
        _ = v7
    }
    var v9 int = -1
    v10 := GenS1{F0: 1000, F1: v4}
    var v11 bool = (v9 == (v9 - v9))
    v12 := []int{(v9 / ((v9) | 1)), 696}
    fmt.Println(v11, v4)
    v13 := GenS1{F0: 0, F1: (v4 + v4)}
    if v11 {
        if (585 > 1) {
            v11 = v1
            var v14 uint = 4294967295
            switch (v4 + ((v4 + v4) + (("foo bar" + v4) + v4))) {
            case "hello":
                var v15 uint32 = 464
                _ = v15
            case "a":
                fmt.Println(v9)
            }
            _ = v14
        }
        v16 := GenS1{F0: 1, F1: v4}
        _ = v16
    }
    _ = len(v12)
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS1{F0: -32768, F1: v4}
        var v19 float32 = float32(v5)
        var v20 int16 = 677
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    if v1 {
        switch "hello" {
        case "a":
            if true {
                var v21 int16 = 330
                for fi22 := 0; fi22 < 15; fi22++ {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = fi22
                }
                _ = v21
            } else {
                if v1 {
                    if false {
                        if (240 >= 4294967295) {
                            var v23 float32 = 3.14159
                            v24 := GenS1{F0: 0, F1: (v4 + "hello")}
                            switch v9 {
                            case -686:
                                switch (-(v9)) {
                                case 4294967295:
                                    switch (^(v9)) {
                                    case -9223372036854775808:
                                        if true {
                                            var v25 int64 = -9223372036854775808
                                            switch (v4 + v4) {
                                            case "\n\t":
                                                switch "" {
                                                case "a":
                                                    v9 = v9
                                                default:
                                                    if (!(!(v23 == v23))) {
                                                        for fi26 := 0; fi26 < 11; fi26++ {
                                                            if (v25 >= (v25 - v25)) {
                                                                switch fi26 {
                                                                case -9223372036854775808:
                                                                    v25 = (v25 | 0)
                                                                case -517:
                                                                    for fi27 := 0; fi27 < 14; fi27++ {
                                                                        v12 = append(v12, 0)
                                                                        _ = fi27
                                                                    }
                                                                case -443:
                                                                    if len(v12) > 0 {
                                                                        v12[0] = -2147483648
                                                                    }
                                                                default:
                                                                    if (625 < 829) {
                                                                        switch v4 {
                                                                        case "a":
                                                                            v25 += v25
                                                                        case "":
                                                                            var v28 bool = v11
                                                                            _ = v28
                                                                        case "\n\t":
                                                                            switch v9 {
                                                                            case 1:
                                                                                fmt.Println(fi26)
                                                                            case 32768:
                                                                                for fi29 := 0; fi29 < 1; fi29++ {
                                                                                    if v1 {
                                                                                        var v30 int32 = -1
                                                                                        var v31 uint = 396
                                                                                        _ = v30
                                                                                        _ = v31
                                                                                    }
                                                                                    _ = fi29
                                                                                }
                                                                            case -41:
                                                                                var v32 uint16 = 32767
                                                                                _ = v32
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v33 uint32 = 65535
                                                                        fmt.Println(v9, v25)
                                                                        v12 = append(v12, -890)
                                                                        _ = v33
                                                                    }
                                                                }
                                                                if (!v11) {
                                                                    switch fi26 {
                                                                    case 9223372036854775807:
                                                                        var v34 uint32 = 1
                                                                        _ = v34
                                                                    case -9223372036854775808:
                                                                        switch (("hello" + v4) + "foo bar") {
                                                                        case "hello":
                                                                            for fi35 := 0; fi35 < 12; fi35++ {
                                                                                var v36 uint32 = 65535
                                                                                if (genFn0() > -128) {
                                                                                    v5 += 3.14159
                                                                                } else {
                                                                                    if v1 {
                                                                                        if (65535 < 1) {
                                                                                            switch fi35 {
                                                                                            case 0:
                                                                                                if v11 {
                                                                                                    v37 := GenS1{F0: 1000, F1: v4}
                                                                                                    var v38 bool = (18446744073709551615 == 65535)
                                                                                                    _ = v37
                                                                                                    _ = v38
                                                                                                }
                                                                                            case -577:
                                                                                                if v1 {
                                                                                                    v5 += (1.5 + 1.5)
                                                                                                } else {
                                                                                                    v3 += v3
                                                                                                }
                                                                                            default:
                                                                                                v39 := GenS1{F0: -32768, F1: v4}
                                                                                                v25 += (-(v25))
                                                                                                _ = v39
                                                                                            }
                                                                                            var v40 string = (v4 + "\n\t")
                                                                                            v12 = append(v12, v9)
                                                                                            _ = v40
                                                                                        }
                                                                                        var v41 uint8 = 1
                                                                                        _ = v41
                                                                                    }
                                                                                    v1 = false
                                                                                    var v42 float32 = 3.14159
                                                                                    if v11 {
                                                                                        v9 -= (^(v9))
                                                                                        var v43 int8 = genFn0()
                                                                                        v5 += v5
                                                                                        v11 = (18446744073709551615 == 4294967295)
                                                                                        v12 = append(v12, fi26)
                                                                                        _ = v43
                                                                                    } else {
                                                                                        _ = len(v12)
                                                                                        v5 -= (v5 * (1e10 - 1.5))
                                                                                        v44 := GenS2{F0: v9, F1: v1, F2: 920}
                                                                                        v45 := GenS1{F0: 1, F1: v4}
                                                                                        v5 = (float64(0.0) + 1e10)
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                    }
                                                                                    v46 := GenS3{F0: 1, F1: -9223372036854775808, F2: "a"}
                                                                                    _ = v42
                                                                                    _ = v46
                                                                                }
                                                                                _ = fi35
                                                                                _ = v36
                                                                            }
                                                                        default:
                                                                            if (127 != fi26) {
                                                                                v1 = v11
                                                                                fmt.Println(v2, v1)
                                                                            } else {
                                                                                v23 -= (v23 - float32((v5 + ((float64((0.1 + v23)) - v5) - float64(v23)))))
                                                                                if len(v12) > 0 {
                                                                                    v12[0] = v9
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v5, v9)
                                                                } else {
                                                                    for fi47 := 0; fi47 < 2; fi47++ {
                                                                        v48 := GenS3{F0: 326, F1: v25, F2: "a"}
                                                                        _ = fi47
                                                                        _ = v48
                                                                    }
                                                                    var v49 uint8 = 128
                                                                    _ = v49
                                                                }
                                                            } else {
                                                                var v50 int16 = -1
                                                                v51 := GenS4{F0: 252}
                                                                var v52 int8 = (v2 / ((genFn0()) | 1))
                                                                _ = v50
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            var v53 string = v4
                                                            _ = fi26
                                                            _ = v53
                                                        }
                                                        _ = len(v12)
                                                        if len(v12) > 0 {
                                                            v12[0] = (v9 ^ 32768)
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v54 := GenS2{F0: (-(v9)), F1: false, F2: 998}
                                                _ = v54
                                            case "a":
                                                v55 := GenS5{F0: int16(v25), F1: -1}
                                                _ = v55
                                            default:
                                                v4 = "a"
                                            }
                                            _ = v25
                                        }
                                    }
                                default:
                                    fmt.Println(v3)
                                }
                            case 42:
                                for fi56 := 0; fi56 < 9; fi56++ {
                                    for fi57 := 0; fi57 < 7; fi57++ {
                                        v11 = false
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        v58 := GenS4{F0: 32767}
                                        _ = fi57
                                        _ = v58
                                    }
                                    fmt.Println(v23, v23)
                                    _ = fi56
                                }
                            default:
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    var v60 int32 = 42
                                    _ = fi59
                                    _ = v60
                                }
                            }
                            _ = v23
                            _ = v24
                        } else {
                            if (((552 != 32767) && (v5 == (-2.25 + (1.5 - v5)))) || v11) {
                                fmt.Println(v4, v9)
                                for fi61 := 0; fi61 < 8; fi61++ {
                                    for fi62 := 0; fi62 < 3; fi62++ {
                                        v4 = ("a" + "a")
                                        v4 = (v4 + ("foo bar" + (v4 + v4)))
                                        _ = fi62
                                    }
                                    v9 += v9
                                    switch v4 {
                                    case "":
                                        for fi63 := 0; fi63 < 12; fi63++ {
                                            var v64 uint32 = 2147483648
                                            _ = fi63
                                            _ = v64
                                        }
                                    }
                                    _ = fi61
                                }
                            }
                        }
                    }
                    if len(v12) > 0 {
                        v12[0] = (^(v9))
                    }
                } else {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        fmt.Println(v4, fi65, fi65)
                        if v11 {
                            if len(v12) > 0 {
                                v12[0] = v9
                            }
                        } else {
                            v12 = append(v12, fi65)
                            fmt.Println(v3)
                        }
                        switch (fi65 / (((fi65 ^ -910)) | 1)) {
                        case 4294967295:
                            if ((173 > 128) || (!(v1 && ((1e10 / float64((0.0 + 0.1))) < v5)))) {
                                fmt.Println(v5)
                            } else {
                                if len(v12) > 0 {
                                    v12[0] = (v9 / ((v9) | 1))
                                }
                                v1 = (!false)
                                var v66 uint16 = 0
                                _ = v66
                            }
                        case 2147483647:
                            var v67 uint8 = 42
                            _ = v67
                        case 969:
                            v68 := GenS6{F0: -2147483648, F1: -1, F2: 1}
                            _ = v68
                        default:
                            v11 = true
                        }
                        _ = fi65
                    }
                }
                for fi69 := 0; fi69 < 18; fi69++ {
                    if (v5 < (v5 + 0.0)) {
                        v70 := GenS7{F0: 622, F1: 1000}
                        _ = v70
                    } else {
                        if v1 {
                            switch v9 {
                            case -2147483648:
                                for fi71 := 0; fi71 < 18; fi71++ {
                                    v3 -= 24
                                    fmt.Println(v5, v9)
                                    switch 0 {
                                    case -1:
                                        v72 := GenS8{F0: "\n\t", F1: 1000}
                                        _ = v72
                                    case 2147483647:
                                        v2 -= 1
                                    default:
                                        v4 = ((((v4 + v4) + v4) + (v4 + ((((v4 + "\n\t") + "\n\t") + (v4 + (v4 + ("a" + v4)))) + (("\n\t" + v4) + ("hello" + v4))))) + ((v4 + v4) + ((v4 + (v4 + v4)) + "")))
                                    }
                                    _ = fi71
                                }
                            case 194:
                                _ = len(v12)
                            }
                            var v73 int32 = 42
                            if true {
                                if v1 {
                                    fmt.Println(fi69, v3, v73)
                                    for fi74 := 0; fi74 < 16; fi74++ {
                                        fmt.Println(v2, fi74, v5)
                                        _ = fi74
                                    }
                                    switch ((((("" + v4) + v4) + "") + (v4 + "foo bar")) + (v4 + v4)) {
                                    case "\n\t":
                                        var v75 int = fi69
                                        _ = v75
                                    }
                                } else {
                                    v3 -= (v2 ^ 100)
                                }
                                fmt.Println(v73, v73, v1)
                            } else {
                                if true {
                                    v76 := GenS6{F0: v73, F1: v73, F2: 255}
                                    var v77 uint32 = 2147483648
                                    _ = v76
                                    _ = v77
                                } else {
                                    if false {
                                        if (v5 != (-2.25 / 1e10)) {
                                            var v78 bool = (-32768 >= -851)
                                            v12 = append(v12, int(v3))
                                            var v79 bool = (v11 || v1)
                                            _ = v78
                                            _ = v79
                                        }
                                        var v80 bool = v1
                                        _ = v80
                                    } else {
                                        var v81 float64 = 0.0
                                        v73 -= v73
                                        for fi82 := 0; fi82 < 3; fi82++ {
                                            if (!((-(v9)) > (v9 % ((v9) | 1)))) {
                                                v4 = v4
                                                v83 := GenS7{F0: 0, F1: 42}
                                                _ = v83
                                            } else {
                                                fmt.Println(v9)
                                                if v11 {
                                                    v4 = v4
                                                    switch (("hello" + ("" + (v4 + "hello"))) + (((v4 + v4) + "\n\t") + v4)) {
                                                    case "hello":
                                                        switch (fi82 | v9) {
                                                        case -2147483648:
                                                            switch v4 {
                                                            case "\n\t":
                                                                _ = len(v12)
                                                            case "hello":
                                                                if ((!((((v81 - (v81 + v5)) == (1e10 * 1e10)) || ((fi82 ^ 2147483647) <= fi82)) || ((v73 * v73) == 0))) && (1 == 0)) {
                                                                    if (((("hello" + v4) + (v4 + (("\n\t" + v4) + (v4 + "foo bar")))) + (v4 + v4)) < v4) {
                                                                        v84 := GenS2{F0: (fi82 * (v9 % (((fi69 ^ (fi82 << ((2) & 63)))) | 1))), F1: true, F2: 42}
                                                                        _ = v84
                                                                    } else {
                                                                        var v85 uint16 = 65535
                                                                        fmt.Println(fi82, v5, v3)
                                                                        v86 := GenS9{F0: ((v4 + (v4 + ("foo bar" + (v4 + (v4 + "foo bar"))))) + v4), F1: fi69, F2: -2147483648}
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                }
                                                            }
                                                            if ((-2.25 < v5) || (v1 && (1 < 174))) {
                                                                var v87 uint16 = 65535
                                                                fmt.Println(v87, v3)
                                                                _ = v87
                                                            }
                                                            v88 := GenS10{F0: 255, F1: 0}
                                                            if (-2.25 < 1.5) {
                                                                if v1 {
                                                                    var v89 uint32 = 1
                                                                    v90 := GenS2{F0: (fi82 | fi82), F1: ((1 >= 78) || v1), F2: 2}
                                                                    _ = v89
                                                                    _ = v90
                                                                } else {
                                                                    v91 := GenS4{F0: 255}
                                                                    var v92 string = ((v4 + (v4 + "foo bar")) + ("a" + (v4 + ("hello" + ""))))
                                                                    if false {
                                                                        if true {
                                                                            var v93 uint64 = 1
                                                                            if (434 < 2147483648) {
                                                                                switch 1 {
                                                                                case 32767:
                                                                                    var v94 float32 = 0.0
                                                                                    _ = v94
                                                                                case 128:
                                                                                    var v95 int8 = (v2 * 1)
                                                                                    _ = v95
                                                                                default:
                                                                                    if true {
                                                                                        if (v11 && v11) {
                                                                                            v96 := GenS2{F0: (v9 * (fi82 + (v9 & 32768))), F1: true, F2: (^(v93))}
                                                                                            v4 = ((v4 + "foo bar") + v4)
                                                                                            _ = v96
                                                                                        } else {
                                                                                            v97 := GenS11{F0: "\n\t", F1: (134 < 2147483647)}
                                                                                            v9 -= (^(fi82))
                                                                                            _ = v97
                                                                                        }
                                                                                        v92 = "a"
                                                                                        if v11 {
                                                                                            v98 := GenS11{F0: (v4 + ("a" + v92)), F1: v1}
                                                                                            if v1 {
                                                                                                v3 -= v3
                                                                                                switch fi69 {
                                                                                                case 4294967295:
                                                                                                    var v99 float64 = v81
                                                                                                    _ = v99
                                                                                                default:
                                                                                                    v2 = v2
                                                                                                }
                                                                                            }
                                                                                            if (v81 <= float64(0.1)) {
                                                                                                var v100 string = v92
                                                                                                if (-2.25 >= -2.25) {
                                                                                                    if v11 {
                                                                                                        if v1 {
                                                                                                            if len(v12) > 0 {
                                                                                                                v12[0] = 42
                                                                                                            }
                                                                                                        } else {
                                                                                                            v93 += 2
                                                                                                            if len(v12) > 0 {
                                                                                                                v12[0] = (fi82 / ((-610) | 1))
                                                                                                            }
                                                                                                            if (float64(float32(0.0)) == (1e10 * (v5 - (v81 * float64(0.1))))) {
                                                                                                                var v101 uint16 = 255
                                                                                                                _ = v101
                                                                                                            } else {
                                                                                                                v100 = "a"
                                                                                                                if v1 {
                                                                                                                    v3 += int8(fi82)
                                                                                                                } else {
                                                                                                                    if len(v12) > 0 {
                                                                                                                        v12[0] = v9
                                                                                                                    }
                                                                                                                }
                                                                                                                v102 := GenS12{F0: v4, F1: "\n\t", F2: (0 == 255)}
                                                                                                                v103 := GenS13{F0: -1, F1: 4294967295, F2: 42}
                                                                                                                _ = v102
                                                                                                                _ = v103
                                                                                                            }
                                                                                                        }
                                                                                                        var v104 bool = (0.0 <= (1.5 - 0.1))
                                                                                                        _ = v104
                                                                                                    } else {
                                                                                                        v105 := GenS14{F0: v11, F1: 1}
                                                                                                        _ = v105
                                                                                                    }
                                                                                                    v1 = (true && ((v2 | (^(v2))) == -110))
                                                                                                    if false {
                                                                                                        _ = len(v12)
                                                                                                        switch (v9 - 1) {
                                                                                                        case 947:
                                                                                                            v106 := GenS15{F0: "hello", F1: v4}
                                                                                                            _ = v106
                                                                                                        case 1:
                                                                                                            v5 = float64(1.5)
                                                                                                        case -9:
                                                                                                            if len(v12) > 0 {
                                                                                                                v12[0] = fi82
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v2)
                                                                                                        var v107 uint64 = v93
                                                                                                        var v108 int = (^(fi82))
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    }
                                                                                                } else {
                                                                                                    switch (("hello" + (("a" + v100) + ((v92 + "") + (v100 + v100)))) + (v4 + v100)) {
                                                                                                    case "a":
                                                                                                        v100 = ((("\n\t" + v4) + v92) + "")
                                                                                                    }
                                                                                                    fmt.Println(v73, v5, v11)
                                                                                                    var v109 uint64 = 2
                                                                                                    _ = v109
                                                                                                }
                                                                                                v2 -= v3
                                                                                                var v110 string = "a"
                                                                                                var v111 uint64 = (^(v93))
                                                                                                _ = v100
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            } else {
                                                                                                v12 = append(v12, (v9 >> ((0) & 63)))
                                                                                                if len(v12) > 0 {
                                                                                                    v12[0] = fi82
                                                                                                }
                                                                                            }
                                                                                            _ = v98
                                                                                        }
                                                                                    } else {
                                                                                        v112 := GenS16{F0: (v4 + v92), F1: (58 < 255), F2: 37}
                                                                                        v93 -= v93
                                                                                        var v113 uint8 = 131
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    var v114 uint32 = 2
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                            _ = v93
                                                                        }
                                                                        switch (-(fi69)) {
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v5, v11, v5)
                                                                        case 565:
                                                                            var v115 uint16 = 0
                                                                            _ = v115
                                                                        case 42:
                                                                            var v116 uint32 = 242
                                                                            _ = v116
                                                                        default:
                                                                            if v1 {
                                                                                if false {
                                                                                    v73 += (v73 & (-(v73)))
                                                                                } else {
                                                                                    if len(v12) > 0 {
                                                                                        v12[0] = 32768
                                                                                    }
                                                                                    if (42 < -435) {
                                                                                        v2 = (v3 - (v3 % ((v3) | 1)))
                                                                                        if len(v12) >= 2 {
                                                                                            v12 = v12[1:len(v12)]
                                                                                        }
                                                                                    } else {
                                                                                        switch (fi69 * int(v73)) {
                                                                                        case 2147483647:
                                                                                            var v117 int64 = -1
                                                                                            _ = v117
                                                                                        default:
                                                                                            switch fi82 {
                                                                                            case 128:
                                                                                                v118 := GenS17{F0: 0, F1: (v92 + v4), F2: 42}
                                                                                                _ = v118
                                                                                            case -604:
                                                                                                switch (v92 + v4) {
                                                                                                case "hello":
                                                                                                    switch v92 {
                                                                                                    case "a":
                                                                                                        v119 := GenS6{F0: v73, F1: int32(v3), F2: 32767}
                                                                                                        _ = v119
                                                                                                    case "":
                                                                                                        var v120 int64 = 0
                                                                                                        _ = v120
                                                                                                    }
                                                                                                case "a":
                                                                                                    if v11 {
                                                                                                        if (!(!v1)) {
                                                                                                            var v121 int8 = (v3 * v2)
                                                                                                            _ = v121
                                                                                                        } else {
                                                                                                            var v122 string = v4
                                                                                                            if true {
                                                                                                                if (true || v1) {
                                                                                                                    fmt.Println(v73)
                                                                                                                    v123 := GenS15{F0: ("\n\t" + "a"), F1: v92}
                                                                                                                    switch "foo bar" {
                                                                                                                    case "":
                                                                                                                        var v124 int64 = 639
                                                                                                                        _ = v124
                                                                                                                    case "hello":
                                                                                                                        switch "a" {
                                                                                                                        case "":
                                                                                                                            if len(v12) > 0 {
                                                                                                                                v12[0] = (^(v9))
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            v3 -= (v2 * v3)
                                                                                                                        case "hello":
                                                                                                                            if (v92 <= v4) {
                                                                                                                                v12 = append(v12, (v9 - -521))
                                                                                                                                switch v122 {
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(v11)
                                                                                                                                case "foo bar":
                                                                                                                                    if len(v12) >= 2 {
                                                                                                                                        v12 = v12[1:len(v12)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v125 int64 = 473
                                                                                                                                fmt.Println(v125)
                                                                                                                                var v126 string = ("a" + "foo bar")
                                                                                                                                _ = v125
                                                                                                                                _ = v126
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if v1 {
                                                                                                                            v127 := GenS18{F0: 65535, F1: 762}
                                                                                                                            _ = v127
                                                                                                                        } else {
                                                                                                                            v12 = append(v12, 128)
                                                                                                                            var v128 int32 = (v73 / ((v73) | 1))
                                                                                                                            var v129 int16 = -12
                                                                                                                            _ = v128
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        var v130 uint8 = 255
                                                                                                                        _ = v130
                                                                                                                    default:
                                                                                                                        if len(v12) >= 2 {
                                                                                                                            v12 = v12[1:len(v12)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v2 -= genFn0()
                                                                                                                    _ = v123
                                                                                                                } else {
                                                                                                                    var v131 uint8 = 2
                                                                                                                    v132 := GenS19{F0: 42, F1: (v4 + "foo bar")}
                                                                                                                    _ = v131
                                                                                                                    _ = v132
                                                                                                                }
                                                                                                                switch 32767 {
                                                                                                                case -39:
                                                                                                                    var v133 uint64 = 0
                                                                                                                    _ = v133
                                                                                                                case 1:
                                                                                                                    var v134 int = fi69
                                                                                                                    _ = v134
                                                                                                                default:
                                                                                                                    var v135 int = (v9 << ((65535) & 63))
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                                switch (("\n\t" + v92) + ("foo bar" + (("a" + "hello") + (((v4 + "\n\t") + v122) + "a")))) {
                                                                                                                case "\n\t":
                                                                                                                    v136 := GenS20{F0: v73, F1: (!((((0.0 * v5) + (v81 / v81)) / float64((((1e10 * (3.14159 - -2.25)) * 0.1) + 1.5))) >= 0.0))}
                                                                                                                    _ = v136
                                                                                                                case "a":
                                                                                                                    switch fi69 {
                                                                                                                    case 9223372036854775807:
                                                                                                                        if ((((0.1 - 0.1) != (1e10 + ((0.1 / -2.25) + (-2.25 - 1e10)))) || false) || false) {
                                                                                                                            if false {
                                                                                                                                fmt.Println(v92, v1)
                                                                                                                                var v137 float32 = float32(float64(0.0))
                                                                                                                                _ = v137
                                                                                                                            }
                                                                                                                            if v1 {
                                                                                                                                v4 = v122
                                                                                                                                v11 = false
                                                                                                                                var v138 float32 = 0.1
                                                                                                                                switch -2147483648 {
                                                                                                                                case 4294967295:
                                                                                                                                    fmt.Println(v138, v138, fi82)
                                                                                                                                case 369:
                                                                                                                                    fmt.Println(v5, v2)
                                                                                                                                case -2147483648:
                                                                                                                                    if (!true) {
                                                                                                                                        if (!((genFn0() <= (v3 ^ -79)) && (fi82 < (v9 * fi69)))) {
                                                                                                                                            var v139 string = (v4 + (((v4 + "a") + "a") + "hello"))
                                                                                                                                            var v140 float32 = 0.0
                                                                                                                                            v140 -= ((v140 / (((v140 - v138) + (-2.25 + (v140 + (v140 / (v138 - 0.1))))) - v138)) * (0.0 - v138))
                                                                                                                                            _ = v139
                                                                                                                                            _ = v140
                                                                                                                                        } else {
                                                                                                                                            v141 := GenS21{F0: -1, F1: true, F2: 65535}
                                                                                                                                            _ = v141
                                                                                                                                        }
                                                                                                                                        var v142 float32 = float32(v81)
                                                                                                                                        _ = v142
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    _ = len(v12)
                                                                                                                                }
                                                                                                                                _ = v138
                                                                                                                            } else {
                                                                                                                                fmt.Println(v73, v3, v1)
                                                                                                                                v11 = ((v11 || (true || (-126 >= -32768))) && (!v1))
                                                                                                                                switch (v9 | v9) {
                                                                                                                                case 4294967295:
                                                                                                                                    fmt.Println(v81, fi82)
                                                                                                                                default:
                                                                                                                                    if len(v12) >= 2 {
                                                                                                                                        v12 = v12[1:len(v12)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v81 = (v81 / (v5 - (3.14159 / ((0.1 + float64(((1.5 * ((1e10 * float32(float64(-2.25))) + 0.1)) / 1e10))) + float64(((-2.25 - float32(((v5 + v81) + -2.25))) * 0.1))))))
                                                                                                                            }
                                                                                                                            var v143 uint = 4294967295
                                                                                                                            _ = v143
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v144 int64 = 42
                                                                                                                        _ = v144
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v81, v92)
                                                                                                                if v1 {
                                                                                                                    v122 = v4
                                                                                                                }
                                                                                                                var v145 int = (fi69 ^ (^(fi82)))
                                                                                                                v146 := GenS5{F0: 32767, F1: 1}
                                                                                                                _ = v145
                                                                                                                _ = v146
                                                                                                            } else {
                                                                                                                if v11 {
                                                                                                                    if len(v12) > 0 {
                                                                                                                        v12[0] = 128
                                                                                                                    }
                                                                                                                    fmt.Println(v1, v5)
                                                                                                                }
                                                                                                                var v147 bool = true
                                                                                                                _ = v147
                                                                                                            }
                                                                                                            v81 = v81
                                                                                                            var v148 int16 = -32768
                                                                                                            v149 := GenS22{F0: 66, F1: 42}
                                                                                                            _ = v122
                                                                                                            _ = v148
                                                                                                            _ = v149
                                                                                                        }
                                                                                                        v92 = "a"
                                                                                                    } else {
                                                                                                        v150 := GenS23{F0: "\n\t", F1: 9223372036854775807, F2: 255}
                                                                                                        var v151 uint16 = 65535
                                                                                                        _ = v150
                                                                                                        _ = v151
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v11 = false
                                                                                                var v152 int32 = v73
                                                                                                _ = v152
                                                                                            }
                                                                                        }
                                                                                        v1 = (v11 || (!(-2.25 < 0.0)))
                                                                                        v153 := GenS20{F0: v73, F1: v1}
                                                                                        var v154 int = (fi82 / (((fi69 * -9223372036854775808)) | 1))
                                                                                        _ = v153
                                                                                        _ = v154
                                                                                    }
                                                                                    v155 := GenS10{F0: 42, F1: -1}
                                                                                    v73 += v73
                                                                                    _ = v155
                                                                                }
                                                                                v156 := GenS24{F0: 18446744073709551615, F1: -32768, F2: 1}
                                                                                _ = v156
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                var v157 int16 = 0
                                                                _ = v157
                                                            }
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                            _ = v88
                                                        case 1:
                                                            fmt.Println(v73, v3)
                                                        case -1:
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        default:
                                                            fmt.Println(fi69)
                                                            var v158 float64 = ((v5 * 1.5) * ((v5 / -2.25) + ((3.14159 / v81) / float64(((0.0 + (0.1 / 0.1)) + 0.0)))))
                                                            _ = v158
                                                        }
                                                    case "":
                                                        if (v11 && (42 != 42)) {
                                                            if len(v12) > 0 {
                                                                v12[0] = fi82
                                                            }
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        }
                                                    default:
                                                        var v159 uint64 = 629
                                                        v160 := GenS15{F0: v4, F1: v4}
                                                        var v161 uint8 = 1
                                                        _ = v159
                                                        _ = v160
                                                        _ = v161
                                                    }
                                                } else {
                                                    if v1 {
                                                        fmt.Println(v73, v73, fi82)
                                                    } else {
                                                        var v162 int32 = v73
                                                        v162 = (^(v162))
                                                        switch fi69 {
                                                        case -32:
                                                            v163 := GenS6{F0: v162, F1: v162, F2: 32767}
                                                            _ = v163
                                                        }
                                                        _ = v162
                                                    }
                                                }
                                            }
                                            _ = fi82
                                        }
                                        _ = v81
                                    }
                                    for fi164 := 0; fi164 < 14; fi164++ {
                                        var v165 float32 = -2.25
                                        _ = fi164
                                        _ = v165
                                    }
                                }
                            }
                            _ = v73
                        }
                    }
                    for fi166 := 0; fi166 < 20; fi166++ {
                        v167 := GenS25{F0: 2147483648, F1: v2, F2: 2}
                        if v11 {
                            fmt.Println(fi166, v1, fi69)
                            if v11 {
                                switch (fi69 + (^(fi69))) {
                                case -945:
                                    v168 := GenS22{F0: 2, F1: -875}
                                    _ = v168
                                case -135:
                                    if len(v12) > 0 {
                                        v12[0] = v9
                                    }
                                case 32767:
                                    fmt.Println(fi69)
                                }
                                _ = len(v12)
                                if (v11 && (!v11)) {
                                    if len(v12) > 0 {
                                        v12[0] = 32767
                                    }
                                }
                            }
                        }
                        if v11 {
                            var v169 float32 = float32((v5 + v5))
                            v3 = 1
                            switch fi166 {
                            case 667:
                                v11 = (690 <= 42)
                            }
                            _ = v169
                        }
                        _ = fi166
                        _ = v167
                    }
                    _ = fi69
                }
            }
        default:
            for fi170 := 0; fi170 < 19; fi170++ {
                var v171 uint32 = 42
                var v172 string = (v4 + v4)
                _ = fi170
                _ = v171
                _ = v172
            }
        }
    } else {
        if false {
            if (!(v1 && (991 < 42))) {
                v5 = float64((1e10 / -2.25))
            }
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        } else {
            var v173 float32 = 1e10
            _ = v173
        }
    }
    switch v9 {
    case 907:
        v174 := GenS23{F0: "foo bar", F1: (^(v9)), F2: 128}
        _ = v174
    default:
        fmt.Println(v11)
    }
    if len(v12) > 0 {
        v12[0] = 42
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v11 = false
    fmt.Println(v9, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
