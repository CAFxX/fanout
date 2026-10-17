package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 bool; F1 uint8; F2 int32 }
type GenS4 struct { F0 bool; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 int; F1 int; F2 int }
type GenS12 struct { F0 uint8; F1 int8 }
type GenS13 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS14 struct { F0 int8; F1 int64; F2 int64 }
type GenS15 struct { F0 uint64; F1 string }
type GenS16 struct { F0 uint32; F1 uint; F2 uint64 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 float64 = 1e10
        if (((((float64(1e10) * (0.0 + v8)) * v8) + (v8 / 3.14159)) / v8) == 0.1) {
            v9 := GenS1{F0: 271, F1: 9, F2: 38}
            _ = v9
        } else {
            v10 := []int{fi7, 1, (fi7 + 120), (fi7 / (((fi7 ^ (fi7 - (fi7 + (fi7 >> ((1) & 63)))))) | 1))}
            v11 := GenS2{F0: 42, F1: -1}
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            v8 += float64(0.1)
            _ = fi12
        }
        _ = fi7
        _ = v8
    }
    v13 := GenS2{F0: 42, F1: 1}
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            if false {
                if ((true || true) || false) {
                    v16 := GenS2{F0: 4294967295, F1: -488}
                    _ = v16
                }
                var v17 uint32 = 2
                var v18 int32 = 406
                _ = v17
                _ = v18
            }
            if ("\n\t" != "") {
                fmt.Println(fi15, fi14)
                switch "" {
                case "foo bar":
                    v19 := []int{fi15, -2147483648, 32768}
                    _ = v19
                case "\n\t":
                    v20 := GenS3{F0: true, F1: 172, F2: 2147483647}
                    _ = v20
                case "a":
                    fmt.Println(fi14, fi14)
                default:
                    switch ("a" + ("hello" + "\n\t")) {
                    case "a":
                        var v21 int = fi15
                        _ = v21
                    case "":
                        var v22 bool = false
                        _ = v22
                    case "\n\t":
                        var v23 float64 = 0.1
                        _ = v23
                    }
                }
                fmt.Println(fi14, fi14)
            }
            if true {
                var v24 int64 = int64(fi15)
                _ = v24
            }
            switch ("hello" + "hello") {
            case "\n\t":
                switch (("hello" + "\n\t") + "\n\t") {
                case "foo bar":
                    v25 := GenS4{F0: false, F1: 0, F2: 520}
                    _ = v25
                case "":
                    var v26 bool = false
                    _ = v26
                }
            case "foo bar":
                switch "" {
                case "foo bar":
                    if (-9223372036854775808 < 42) {
                        if false {
                            if true {
                                var v27 float32 = 0.1
                                v28 := GenS3{F0: (79 > 846), F1: 2, F2: 336}
                                if (9223372036854775807 < 42) {
                                    v27 += v27
                                    if ((!("\n\t" >= "\n\t")) || (!false)) {
                                        v27 += ((float32(0.1) + -2.25) * (v27 / v27))
                                    } else {
                                        var v29 uint8 = 1
                                        v30 := []int{fi14}
                                        v31 := GenS2{F0: 2, F1: 42}
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                                _ = v27
                                _ = v28
                            }
                            fmt.Println(fi15, fi14)
                            if false {
                                fmt.Println(fi14, fi14)
                                fmt.Println(fi15)
                            }
                            v32 := []int{}
                            v33 := GenS5{F0: -1, F1: 42}
                            var v34 int32 = -241
                            var v35 float64 = 3.14159
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            if (!true) {
                                var v36 uint16 = 1
                                v36 = (v36 & v36)
                                v36 -= v36
                                _ = v36
                            }
                        }
                    }
                case "a":
                    var v37 int64 = 369
                    _ = v37
                case "hello":
                    v38 := GenS1{F0: -2147483648, F1: 2, F2: 175}
                    _ = v38
                }
            }
            var v39 float64 = (0.0 - -2.25)
            _ = fi15
            _ = v39
        }
        var v40 uint16 = 255
        v40 -= (v40 + (v40 + 0))
        _ = fi14
        _ = v40
    }
    var v41 uint16 = 65535
    _ = v13
    _ = v41
    return 0
}

func main() {
    v1 := []string{"\n\t", ""}
    var v2 uint64 = 4294967295
    var v3 int = 4294967295
    var v4 float64 = ((1e10 * 1e10) + (3.14159 / 0.1))
    var v5 uint8 = 2
    v6 := []uint32{}
    fmt.Println(v5, v5)
    if (891 > v3) {
        v5 = (^(v5))
        for fi42 := 0; fi42 < 13; fi42++ {
            switch "a" {
            case "foo bar":
                var v43 int = fi42
                _ = v43
            case "a":
                v2 = (v2 / (((^(v2))) | 1))
            case "":
                v4 = v4
            default:
                for fi44 := 0; fi44 < 9; fi44++ {
                    if true {
                        var v45 uint64 = v2
                        v46 := GenS2{F0: v2, F1: 1}
                        switch fi42 {
                        case 695:
                            switch fi44 {
                            case -496:
                                if ((fi44 * v3) < int(v5)) {
                                    if false {
                                        v47 := GenS4{F0: true, F1: 42, F2: 42}
                                        _ = v47
                                    } else {
                                        var v48 float32 = 1.5
                                        _ = len(v6)
                                        v4 -= v4
                                        _ = v48
                                    }
                                } else {
                                    var v49 float64 = v4
                                    fmt.Println(v5)
                                    if (false || true) {
                                        fmt.Println(v5, v5, v45)
                                        var v50 bool = (fi44 < v3)
                                        _ = v50
                                    }
                                    _ = v49
                                }
                            case -9223372036854775808:
                                if (42 > 837) {
                                    fmt.Println(v45, v3)
                                    fmt.Println(v45)
                                    var v51 uint = 255
                                    if ((-1 >= -489) || false) {
                                        var v52 float64 = v4
                                        _ = v52
                                    }
                                    switch "a" {
                                    case "":
                                        switch fi42 {
                                        case 32768:
                                            v3 -= fi42
                                        case 0:
                                            v51 += v51
                                        default:
                                            var v53 int16 = genFn0()
                                            _ = v53
                                        }
                                    case "foo bar":
                                        if false {
                                            switch "" {
                                            case "":
                                                v54 := GenS4{F0: ((v4 == v4) || ((-1 > int8(v45)) || false)), F1: int32(v5), F2: 964}
                                                _ = v54
                                            case "\n\t":
                                                v55 := GenS5{F0: -2147483648, F1: v5}
                                                _ = v55
                                            case "a":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            v45 -= (v45 - v2)
                                            _ = len(v6)
                                        }
                                    }
                                    _ = v51
                                } else {
                                    v56 := GenS6{F0: false, F1: 442}
                                    _ = v56
                                }
                            case 461:
                                var v57 int16 = 149
                                _ = v57
                            }
                        case -413:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v2 = 498
                            if len(v6) > 0 {
                                v6[0] = 4294967295
                            }
                        case 1:
                            if true {
                                v3 = -1
                                switch (v3 & (v3 - -1)) {
                                case 1:
                                    if false {
                                        v58 := GenS7{F0: 65535, F1: v5}
                                        v59 := GenS8{F0: true, F1: -799}
                                        _ = v58
                                        _ = v59
                                    }
                                case 0:
                                    switch "a" {
                                    case "a":
                                        var v60 int = fi42
                                        _ = v60
                                    case "hello":
                                        v1 = append(v1, "hello")
                                    }
                                    fmt.Println(v3, v45)
                                    v61 := GenS2{F0: v45, F1: 0}
                                    _ = v61
                                case -1:
                                    if len(v1) > 0 {
                                        v1[0] = (("foo bar" + "") + ("" + "a"))
                                    }
                                }
                            }
                        }
                        var v62 float32 = 1e10
                        _ = v45
                        _ = v46
                        _ = v62
                    } else {
                        v1 = append(v1, ("\n\t" + ("" + "a")))
                        fmt.Println(fi44, v4, fi42)
                        v3 += fi44
                    }
                    v2 = (^(v2))
                    var v63 int = (-(fi44))
                    _ = fi44
                    _ = v63
                }
            }
            _ = fi42
        }
        fmt.Println(v3, v5)
    }
    switch v3 {
    case 32767:
        switch ("foo bar" + "hello") {
        case "":
            if true {
                if (1 == 2147483647) {
                    fmt.Println(v2)
                }
                if false {
                    v5 += 42
                    for fi64 := 0; fi64 < 14; fi64++ {
                        for fi65 := 0; fi65 < 14; fi65++ {
                            switch "a" {
                            case "foo bar":
                                var v66 float32 = 0.1
                                _ = v66
                            case "":
                                var v67 int = (^(fi65))
                                _ = v67
                            case "\n\t":
                                switch -95 {
                                case 32768:
                                    v1 = append(v1, "hello")
                                case 128:
                                    v4 += ((((v4 + (((3.14159 + v4) - v4) * float64(0.0))) + ((v4 / 1.0) + v4)) * 3.14159) + 0.0)
                                case 1:
                                    v68 := GenS1{F0: 1, F1: v5, F2: (^(v5))}
                                    _ = v68
                                }
                            default:
                                var v69 uint32 = 436
                                _ = v69
                            }
                            var v70 int64 = -141
                            fmt.Println(v5, v4, v5)
                            _ = fi65
                            _ = v70
                        }
                        for fi71 := 0; fi71 < 1; fi71++ {
                            var v72 int = (-(fi64))
                            var v73 int16 = -543
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                        var v74 float64 = float64(1e10)
                        _ = fi64
                        _ = v74
                    }
                }
            } else {
                v1 = append(v1, "foo bar")
                for fi75 := 0; fi75 < 20; fi75++ {
                    var v76 int8 = 0
                    if (!((!false) && true)) {
                        v76 = v76
                    } else {
                        if false {
                            if (750 != 563) {
                                switch (v3 % ((4294967295) | 1)) {
                                case -2147483648:
                                    fmt.Println(v76)
                                case -294:
                                    v77 := GenS1{F0: 42, F1: (^(v5)), F2: (v5 / ((v5) | 1))}
                                    _ = v77
                                case 342:
                                    if false {
                                        v78 := GenS5{F0: -1, F1: 42}
                                        _ = v78
                                    } else {
                                        v79 := GenS1{F0: 206, F1: (v5 * 255), F2: (v5 & 255)}
                                        if (true && false) {
                                            v3 -= v3
                                            v4 += 3.14159
                                            v2 = v2
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v80 int64 = 9223372036854775807
                                                _ = v80
                                            case "":
                                                v81 := GenS3{F0: ((v3 > (fi75 | -212)) || false), F1: (v5 % (((^(v5))) | 1)), F2: 1}
                                                _ = v81
                                            case "a":
                                                if false {
                                                    v82 := GenS6{F0: true, F1: 4294967295}
                                                    _ = v82
                                                } else {
                                                    fmt.Println(v4, fi75)
                                                    for fi83 := 0; fi83 < 6; fi83++ {
                                                        v84 := GenS9{F0: 242, F1: 4294967295}
                                                        v3 = (fi83 ^ (-(fi75)))
                                                        _ = fi83
                                                        _ = v84
                                                    }
                                                }
                                            default:
                                                v76 = (v76 + (v76 * (v76 - (v76 / ((v76) | 1)))))
                                            }
                                            v85 := GenS10{F0: ((uint16(v5) == 912) || (65535 >= 0)), F1: (v2 + v2)}
                                            _ = v85
                                        }
                                        _ = v79
                                    }
                                }
                                if false {
                                    var v86 string = "a"
                                    var v87 bool = true
                                    for fi88 := 0; fi88 < 16; fi88++ {
                                        if (v3 != (v3 | (fi88 / ((32767) | 1)))) {
                                            var v89 uint16 = 1
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            fmt.Println(v3, v76, fi88)
                                            _ = v89
                                        } else {
                                            v90 := GenS11{F0: (fi75 << ((4294967295) & 63)), F1: fi88, F2: (-(v3))}
                                            if v87 {
                                                var v91 bool = v87
                                                fmt.Println(v3, v87, v87)
                                                _ = len(v6)
                                                v92 := GenS7{F0: 42, F1: v5}
                                                _ = v91
                                                _ = v92
                                            }
                                            fmt.Println(v5, v86)
                                            _ = v90
                                        }
                                        _ = fi88
                                    }
                                    _ = v86
                                    _ = v87
                                }
                            }
                            _ = len(v1)
                        }
                        v93 := GenS10{F0: (v2 < v2), F1: 0}
                        _ = v93
                    }
                    v4 = 1e10
                    _ = fi75
                    _ = v76
                }
                for fi94 := 0; fi94 < 4; fi94++ {
                    v5 += (v5 ^ (^(v5)))
                    v95 := GenS12{F0: v5, F1: -124}
                    _ = fi94
                    _ = v95
                }
            }
        default:
            var v96 int64 = -474
            _ = v96
        }
    case 127:
        for fi97 := 0; fi97 < 2; fi97++ {
            if (("\n\t" == "hello") || true) {
                if false {
                    for fi98 := 0; fi98 < 1; fi98++ {
                        v2 += v2
                        var v99 int16 = 542
                        _ = fi98
                        _ = v99
                    }
                    v2 = (v2 + (^(v2)))
                } else {
                    if (true || false) {
                        var v100 int = (fi97 & -1)
                        var v101 int8 = 85
                        fmt.Println(fi97, v2, v101)
                        _ = v100
                        _ = v101
                    }
                    var v102 int8 = 0
                    var v103 uint8 = v5
                    _ = v102
                    _ = v103
                }
            }
            switch (("hello" + "hello") + "") {
            case "a":
                var v104 int8 = -1
                _ = v104
            case "hello":
                for fi105 := 0; fi105 < 4; fi105++ {
                    fmt.Println(v4, fi105)
                    var v106 string = (("" + "\n\t") + "a")
                    v107 := GenS11{F0: (^(fi97)), F1: fi97, F2: -826}
                    _ = fi105
                    _ = v106
                    _ = v107
                }
            default:
                v108 := GenS7{F0: 1, F1: 40}
                _ = v108
            }
            var v109 float32 = 3.14159
            _ = fi97
            _ = v109
        }
    default:
        v110 := GenS13{F0: 556, F1: 65535, F2: 42}
        _ = v110
    }
    v5 = (^(v5))
    if (float32((((v4 * float64((0.0 / (-2.25 - (0.0 / (float32(v4) - (1e10 * -2.25))))))) / (v4 + 1e10)) - ((v4 * (v4 - v4)) / v4))) != (1.5 * -2.25)) {
        v111 := GenS14{F0: 127, F1: 1, F2: 365}
        v6 = append(v6, 1)
        _ = v111
    }
    fmt.Println(v2)
    v6 = append(v6, 769)
    switch (v3 | -12) {
    case 32768:
        v5 = v5
    case 9223372036854775807:
        v2 += 209
    case 0:
        v3 -= v3
    }
    var v112 uint = 0
    v4 = v4
    v113 := GenS15{F0: v2, F1: "foo bar"}
    _ = len(v6)
    v114 := GenS16{F0: 2147483648, F1: (v112 ^ v112), F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
