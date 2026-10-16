package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16; F1 string; F2 int32 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 bool; F1 uint32; F2 int }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 string; F1 uint16; F2 int64 }
type GenS11 struct { F0 bool; F1 int8 }
func genFn0(p0 int, p1 uint16, p2 int16) uint8 {
    p0 = int(p2)
    var v7 uint8 = 42
    p0 = (p0 >> ((65535) & 63))
    var v8 int16 = (p2 - -350)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2147483647
    var v3 float32 = 1e10
    var v4 bool = false
    var v5 int = 32768
    var v6 int8 = -1
    if (-1 == 42) {
        for fi9 := 0; fi9 < 13; fi9++ {
            v1 = append(v1, v2)
            var v10 uint32 = 195
            _ = fi9
            _ = v10
        }
        _ = len(v1)
        var v11 uint64 = 61
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            switch -9223372036854775808 {
            case 127:
                var v13 uint = 255
                _ = v13
            case 392:
                if true {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v15 := GenS1{F0: 940}
                        if false {
                            if (-1 < (^(v5))) {
                                switch (fi12 ^ fi12) {
                                case 127:
                                    var v16 int8 = 81
                                    fmt.Println(fi12)
                                    var v17 uint64 = 18446744073709551615
                                    _ = v16
                                    _ = v17
                                case -604:
                                    switch "" {
                                    case "":
                                        v18 := GenS1{F0: -670}
                                        _ = v18
                                    default:
                                        var v19 uint16 = 2
                                        _ = v19
                                    }
                                case 236:
                                    fmt.Println(fi14, v3, v5)
                                default:
                                    v4 = v4
                                }
                                v20 := GenS1{F0: 377}
                                if v4 {
                                    if false {
                                        var v21 uint = 237
                                        v6 += v6
                                        _ = v21
                                    } else {
                                        if (!v4) {
                                            v22 := GenS1{F0: 0}
                                            switch fi12 {
                                            case -1:
                                                var v23 uint32 = v2
                                                fmt.Println(fi12, fi12, v6)
                                                if v4 {
                                                    v24 := GenS2{F0: (v6 << ((262) & 63))}
                                                    var v25 uint64 = 573
                                                    v23 += v23
                                                    _ = v24
                                                    _ = v25
                                                }
                                                v4 = v4
                                                _ = v23
                                            case -590:
                                                v2 = 444
                                            }
                                            _ = v22
                                        }
                                    }
                                    v26 := GenS2{F0: v6}
                                    _ = v26
                                } else {
                                    var v27 float32 = v3
                                    v5 -= fi14
                                    _ = v27
                                }
                                var v28 bool = ("hello" <= "hello")
                                _ = v20
                                _ = v28
                            }
                        } else {
                            if false {
                                var v29 float32 = v3
                                v30 := GenS2{F0: v6}
                                _ = v29
                                _ = v30
                            }
                            var v31 uint8 = 0
                            _ = v31
                        }
                        _ = fi14
                        _ = v15
                    }
                    if ((v4 && (546 == 255)) || false) {
                        if v4 {
                            switch (v5 ^ (v5 ^ int(v6))) {
                            case 460:
                                v1 = append(v1, 65535)
                            case 853:
                                if true {
                                    switch v5 {
                                    case 42:
                                        if v4 {
                                            fmt.Println(v5)
                                        } else {
                                            v32 := GenS1{F0: 0}
                                            if (false && false) {
                                                if len(v1) > 0 {
                                                    v1[0] = 2
                                                }
                                                switch (fi12 | fi12) {
                                                case -611:
                                                    for fi33 := 0; fi33 < 9; fi33++ {
                                                        v34 := GenS3{F0: 674, F1: 0}
                                                        v35 := GenS1{F0: -1}
                                                        fmt.Println(fi12, v6, fi33)
                                                        switch fi12 {
                                                        case 32767:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                _ = len(v1)
                                                            case "":
                                                                if false {
                                                                    switch (fi33 * (-(v5))) {
                                                                    case 1:
                                                                        v6 = v6
                                                                    case -1:
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            var v36 int32 = 42
                                                                            _ = v36
                                                                        case "":
                                                                            v3 = v3
                                                                        case "a":
                                                                            v5 = fi12
                                                                        default:
                                                                            switch "" {
                                                                            case "":
                                                                                fmt.Println(v2, v4, v6)
                                                                                v4 = v4
                                                                            case "foo bar":
                                                                                if (genFn0((^(fi33)), 42, -812) == 255) {
                                                                                    if (-1 >= 726) {
                                                                                        var v37 int = -274
                                                                                        v38 := GenS4{F0: 1000, F1: "a", F2: 898}
                                                                                        _ = v37
                                                                                        _ = v38
                                                                                    }
                                                                                    switch 4294967295 {
                                                                                    case 0:
                                                                                        var v39 float64 = (0.0 / (3.14159 * (float64((v3 + v3)) - 0.0)))
                                                                                        _ = v39
                                                                                    case -2147483648:
                                                                                        _ = len(v1)
                                                                                    case 32768:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = (v2 >> ((255) & 63))
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    _ = len(v1)
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        v6 = v6
                                                                    }
                                                                    var v40 uint32 = v2
                                                                    _ = v40
                                                                }
                                                            case "hello":
                                                                fmt.Println(fi33)
                                                            }
                                                        }
                                                        switch (v5 << ((65535) & 63)) {
                                                        case 128:
                                                            var v41 uint32 = v2
                                                            var v42 string = "a"
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        _ = fi33
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                case 4294967295:
                                                    switch ("\n\t" + ("\n\t" + ("" + "foo bar"))) {
                                                    case "\n\t":
                                                        v2 = (^(v2))
                                                    case "foo bar":
                                                        var v43 int8 = (v6 + v6)
                                                        _ = v43
                                                    }
                                                default:
                                                    switch "a" {
                                                    case "hello":
                                                        fmt.Println(v2, v5)
                                                    case "":
                                                        if true {
                                                            fmt.Println(v6, fi12)
                                                        } else {
                                                            switch "\n\t" {
                                                            case "":
                                                                v3 += 0.0
                                                            case "\n\t":
                                                                if (v4 && v4) {
                                                                    var v44 int16 = 1000
                                                                    v45 := GenS5{F0: ((v4 && (0.0 != 0.1)) || v4), F1: 127, F2: uint8(v2)}
                                                                    _ = v44
                                                                    _ = v45
                                                                } else {
                                                                    if false {
                                                                        for fi46 := 0; fi46 < 20; fi46++ {
                                                                            v1 = append(v1, v2)
                                                                            _ = fi46
                                                                        }
                                                                        v47 := GenS6{F0: "\n\t"}
                                                                        _ = v47
                                                                    }
                                                                    if v4 {
                                                                        _ = len(v1)
                                                                        for fi48 := 0; fi48 < 20; fi48++ {
                                                                            v2 = (^(v2))
                                                                            if (1e10 != float64((-2.25 + (((((v3 - 1.5) + v3) - v3) - 1e10) * v3)))) {
                                                                                if (0.1 < ((3.14159 / 3.14159) + (1e10 + v3))) {
                                                                                    v6 += (v6 | (v6 - v6))
                                                                                } else {
                                                                                    switch (fi12 | (fi12 % ((32767) | 1))) {
                                                                                    case 534:
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    case -2147483648:
                                                                                        fmt.Println(fi12, v3, v3)
                                                                                    }
                                                                                    var v49 uint = 408
                                                                                    var v50 uint = v49
                                                                                    _ = v49
                                                                                    _ = v50
                                                                                }
                                                                                if v4 {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (v2 / ((v2) | 1))
                                                                                    }
                                                                                    if (v4 || v4) {
                                                                                        if (173 != 42) {
                                                                                            var v51 int32 = 395
                                                                                            fmt.Println(v3, v6, fi12)
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            _ = v51
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (v2 * (^(v2)))
                                                                                    }
                                                                                    var v52 bool = (!v4)
                                                                                    if (357 < 18446744073709551615) {
                                                                                        var v53 float32 = (((v3 / v3) * v3) + (1e10 * float32(((1e10 - 1.5) - 0.0))))
                                                                                        v54 := GenS6{F0: ""}
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                    }
                                                                                    _ = v52
                                                                                }
                                                                            } else {
                                                                                var v55 float32 = -2.25
                                                                                _ = v55
                                                                            }
                                                                            if false {
                                                                                var v56 float64 = 1e10
                                                                                v57 := GenS3{F0: 2, F1: 255}
                                                                                v56 = v56
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                            _ = fi48
                                                                        }
                                                                        var v58 uint32 = (v2 >> ((1) & 63))
                                                                        _ = v58
                                                                    } else {
                                                                        var v59 int16 = 42
                                                                        v59 -= v59
                                                                        v60 := GenS2{F0: 127}
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                    v61 := GenS2{F0: (v6 - v6)}
                                                                    _ = v61
                                                                }
                                                            default:
                                                                switch v5 {
                                                                case 9223372036854775807:
                                                                    var v62 uint64 = 2
                                                                    _ = v62
                                                                case 501:
                                                                    var v63 float32 = (-2.25 + (-2.25 + v3))
                                                                    _ = v63
                                                                case 2147483647:
                                                                    if (!true) {
                                                                        v5 = v5
                                                                        var v64 int = (v5 % ((fi12) | 1))
                                                                        _ = v64
                                                                    }
                                                                default:
                                                                    var v65 int16 = 32767
                                                                    _ = v65
                                                                }
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                fmt.Println(v6, v3)
                                                            default:
                                                                v66 := GenS3{F0: 722, F1: 4294967295}
                                                                _ = v66
                                                            }
                                                        }
                                                    default:
                                                        if ((v2 << ((4294967295) & 63)) == (^(v2))) {
                                                            var v67 int32 = 1
                                                            var v68 int32 = v67
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    }
                                                }
                                            } else {
                                                v3 = v3
                                                if (874 != 18446744073709551615) {
                                                    fmt.Println(v2, v6)
                                                } else {
                                                    var v69 uint16 = uint16(v2)
                                                    v69 = v69
                                                    _ = v69
                                                }
                                            }
                                            var v70 bool = (!(1000 >= -1))
                                            _ = v32
                                            _ = v70
                                        }
                                    case -791:
                                        if true {
                                            fmt.Println(v6)
                                            switch "" {
                                            case "foo bar":
                                                var v71 int64 = 0
                                                _ = v71
                                            default:
                                                v4 = (2 >= 42)
                                            }
                                        }
                                    case 32768:
                                        if v4 {
                                            for fi72 := 0; fi72 < 10; fi72++ {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch ("\n\t" + ("a" + ("a" + "\n\t"))) {
                                                    case "\n\t":
                                                        v73 := GenS7{F0: -819, F1: fi12}
                                                        _ = v73
                                                    case "foo bar":
                                                        if len(v1) > 0 {
                                                            v1[0] = v2
                                                        }
                                                    case "":
                                                        var v74 uint16 = 1
                                                        _ = v74
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            var v75 int32 = 2147483647
                                                            _ = v75
                                                        case "\n\t":
                                                            v2 += 42
                                                        default:
                                                            switch fi12 {
                                                            case 942:
                                                                v6 = 0
                                                            case -9223372036854775808:
                                                                fmt.Println(v4, v2)
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    v6 = 1
                                                    v6 = 0
                                                    v2 = v2
                                                default:
                                                    v76 := GenS6{F0: "a"}
                                                    _ = v76
                                                }
                                                _ = fi72
                                            }
                                        }
                                    default:
                                        v77 := GenS7{F0: -900, F1: fi12}
                                        _ = v77
                                    }
                                    for fi78 := 0; fi78 < 17; fi78++ {
                                        if (!(!(2 >= 74))) {
                                            if (v4 && (((v6 & v6) < (v6 >> ((628) & 63))) && (!false))) {
                                                v79 := GenS8{F0: (-1 == v6), F1: (v2 - 2147483648), F2: (fi78 / ((fi12) | 1))}
                                                v80 := GenS5{F0: v4, F1: 0, F2: 128}
                                                v5 += v5
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch "" {
                                                case "a":
                                                    var v81 uint8 = uint8(v2)
                                                    if (("foo bar" + "a") > ("\n\t" + "")) {
                                                        var v82 uint16 = 32767
                                                        v83 := GenS4{F0: -1, F1: ("a" + "\n\t"), F2: -2147483648}
                                                        var v84 float64 = (((float64((1e10 + float32(-2.25))) / 1.0) / 1.0) * (-2.25 / 1.5))
                                                        var v85 float64 = v84
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = (v2 ^ v2)
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v81
                                                case "hello":
                                                    _ = len(v1)
                                                }
                                                _ = v79
                                                _ = v80
                                            } else {
                                                var v86 float64 = 1.5
                                                fmt.Println(v2, v6, fi78)
                                                fmt.Println(v6, v2, fi78)
                                                var v87 int64 = 42
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        _ = fi78
                                    }
                                } else {
                                    var v88 int32 = 1
                                    switch v5 {
                                    case 127:
                                        var v89 float32 = v3
                                        _ = v89
                                    default:
                                        fmt.Println(fi12, v3, v6)
                                    }
                                    _ = v88
                                }
                            case 4294967295:
                                var v90 uint = 9
                                _ = v90
                            }
                            fmt.Println(v3, fi12, v5)
                            v91 := GenS8{F0: (260 < 677), F1: (v2 | v2), F2: -553}
                            _ = v91
                        } else {
                            v92 := GenS1{F0: 42}
                            var v93 uint = 1
                            _ = v92
                            _ = v93
                        }
                        for fi94 := 0; fi94 < 4; fi94++ {
                            if (!(42 <= 65535)) {
                                switch "" {
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = 788
                                    }
                                case "hello":
                                    switch ("\n\t" + (("hello" + "\n\t") + "")) {
                                    case "hello":
                                        var v95 int8 = -123
                                        _ = v95
                                    case "":
                                        v2 += v2
                                    case "\n\t":
                                        v3 += v3
                                    }
                                case "a":
                                    var v96 uint16 = 840
                                    _ = v96
                                default:
                                    var v97 int16 = -568
                                    _ = v97
                                }
                            } else {
                                v98 := GenS9{F0: 42, F1: fi94}
                                var v99 int32 = 0
                                _ = v98
                                _ = v99
                            }
                            if (v6 != 1) {
                                var v100 uint16 = 907
                                fmt.Println(v100, v4)
                                v101 := GenS6{F0: ("foo bar" + ("foo bar" + "foo bar"))}
                                _ = v100
                                _ = v101
                            } else {
                                v102 := GenS10{F0: "a", F1: 65535, F2: 9223372036854775807}
                                _ = v102
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi94
                        }
                    }
                } else {
                    v103 := GenS11{F0: true, F1: v6}
                    v3 += v3
                    _ = v103
                }
            case 703:
                v3 = ((float32(3.14159) * (((v3 * v3) + ((v3 + v3) + v3)) + (1.5 - v3))) / v3)
            default:
                var v104 string = ""
                _ = v104
            }
            _ = fi12
        }
    }
    var v105 int = -698
    v1 = append(v1, (v2 & 2))
    switch "" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "":
        v6 += -128
    default:
        v106 := GenS10{F0: "", F1: 974, F2: -1}
        _ = v106
    }
    var v107 int16 = 467
    v2 += 2147483647
    for fi108 := 0; fi108 < 19; fi108++ {
        for fi109 := 0; fi109 < 7; fi109++ {
            var v110 uint8 = 255
            _ = fi109
            _ = v110
        }
        _ = fi108
    }
    for fi111 := 0; fi111 < 9; fi111++ {
        var v112 int = 42
        _ = fi111
        _ = v112
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
