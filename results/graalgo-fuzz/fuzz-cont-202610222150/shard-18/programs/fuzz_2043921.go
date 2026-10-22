package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint64; F1 string }
type GenS9 struct { F0 int; F1 int16; F2 uint32 }
type GenS10 struct { F0 int64; F1 bool; F2 uint32 }
type GenS11 struct { F0 bool }
func genFn0(p0 uint16, p1 int32) int16 {
    v5 := GenS1{F0: p1, F1: -848, F2: -1}
    switch "" {
    case "\n\t":
        v6 := []int{-210, 308, 996}
        _ = v6
    default:
        var v7 uint = 18446744073709551615
        _ = v7
    }
    if false {
        switch 127 {
        case 42:
            v8 := []int{1, 32, 128}
            _ = v8
        case 9223372036854775807:
            var v9 int64 = 9223372036854775807
            _ = v9
        default:
            var v10 int64 = 9223372036854775807
            _ = v10
        }
        var v11 bool = true
        _ = v11
    }
    p0 -= (p0 % ((p0) | 1))
    if false {
        v12 := []int{-2147483648, 32767}
        _ = v12
    } else {
        var v13 string = (((("\n\t" + "hello") + "foo bar") + "a") + "\n\t")
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 int, p1 uint16, p2 uint64) uint16 {
    p0 += -568
    for fi14 := 0; fi14 < 9; fi14++ {
        fmt.Println(p2)
        for fi15 := 0; fi15 < 3; fi15++ {
            p1 -= (p1 / (((p1 % ((p1) | 1))) | 1))
            var v16 string = "a"
            p0 -= 12
            if ((!(0.1 > float32(float64((float32(-2.25) / -2.25))))) && true) {
                if ((p1 / (((p1 / ((p1) | 1))) | 1)) == p1) {
                    p0 = 9223372036854775807
                    v17 := []int{fi15, fi14}
                    v18 := GenS1{F0: -1, F1: -2147483648, F2: -32768}
                    p0 += fi14
                    var v19 int16 = 1
                    _ = v17
                    _ = v18
                    _ = v19
                }
                p2 -= p2
            } else {
                v20 := []int{(fi14 & (-(fi15))), 2147483647, fi14, int(p2)}
                v20 = append(v20, (fi15 & int(p2)))
                _ = v20
            }
            if (false && (2 == 42)) {
                if false {
                    p0 += (fi14 | 4294967295)
                    v21 := GenS2{F0: false, F1: -757, F2: 2}
                    _ = v21
                } else {
                    p1 += 311
                    v22 := []int{(fi14 | (fi15 * (fi14 | fi15))), (fi14 & (^(fi14))), (^(p0))}
                    p0 += fi15
                    _ = v22
                }
                p0 = fi14
                var v23 uint64 = p2
                _ = v23
            } else {
                switch (v16 + v16) {
                case "foo bar":
                    var v24 int = fi15
                    _ = v24
                case "a":
                    if (!false) {
                        p0 = fi14
                        fmt.Println(fi15, fi15, p2)
                        v25 := GenS1{F0: -1, F1: p0, F2: 32767}
                        _ = v25
                    }
                }
                var v26 uint8 = 42
                _ = v26
            }
            _ = fi15
            _ = v16
        }
        v27 := GenS1{F0: 1, F1: (fi14 & fi14), F2: -1}
        _ = fi14
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 float32 = 1.5
    var v2 uint = 65535
    v3 := []uint32{796, 0, 4294967295}
    var v4 float64 = (((-2.25 + 1.5) * 3.14159) * 1.5)
    switch "" {
    case "a":
        v1 += 3.14159
    default:
        fmt.Println(v2, v4)
    }
    _ = len(v3)
    if false {
        var v28 int = 2147483647
        for fi29 := 0; fi29 < 13; fi29++ {
            fmt.Println(v4, v2)
            _ = fi29
        }
        if (("hello" + "a") != "a") {
            var v30 bool = false
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                v32 := GenS1{F0: -2147483648, F1: (fi31 * 311), F2: 32767}
                fmt.Println(v28, v30)
                _ = fi31
                _ = v32
            }
            _ = v30
        }
        _ = v28
    }
    _ = len(v3)
    v33 := GenS2{F0: false, F1: 40, F2: 0}
    v1 -= float32(float64(v1))
    v34 := GenS2{F0: (128 < 128), F1: -289, F2: 128}
    _ = len(v3)
    if (-9223372036854775808 > 0) {
        v4 -= v4
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
    } else {
        fmt.Println(v4, v4, v2)
    }
    v1 = float32(v4)
    v35 := GenS3{F0: 2147483647, F1: 42}
    _ = len(v3)
    var v36 int8 = 1
    v4 = 0.0
    switch (("hello" + "a") + "a") {
    case "":
        switch 32767 {
        case 0:
            for fi37 := 0; fi37 < 16; fi37++ {
                switch (fi37 ^ fi37) {
                case 264:
                    if (-9223372036854775808 < 9223372036854775807) {
                        if (false && ((v36 % ((v36) | 1)) != v36)) {
                            var v38 float64 = v4
                            v39 := GenS4{F0: 65535}
                            _ = v38
                            _ = v39
                        }
                        if false {
                            var v40 uint16 = 1
                            for fi41 := 0; fi41 < 2; fi41++ {
                                fmt.Println(v36)
                                var v42 uint16 = 969
                                _ = fi41
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 17; fi43++ {
                                if false {
                                    fmt.Println(v36)
                                    v4 -= (v4 / v4)
                                    if len(v3) > 0 {
                                        v3[0] = 2
                                    }
                                } else {
                                    var v44 int8 = 1
                                    v3 = append(v3, 593)
                                    var v45 uint = (v2 << (((v2 + (^(v2)))) & 63))
                                    switch (("" + ("\n\t" + ("a" + "\n\t"))) + "a") {
                                    case "a":
                                        var v46 float32 = 3.14159
                                        _ = v46
                                    case "foo bar":
                                        v4 = float64((1.5 - 0.1))
                                    case "":
                                        var v47 int64 = 204
                                        _ = v47
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            if len(v3) > 0 {
                                                v3[0] = 905
                                            }
                                        case "":
                                            var v48 int = (fi43 ^ (fi37 * (fi43 ^ fi43)))
                                            var v49 int64 = -1
                                            v50 := GenS4{F0: v40}
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                    v1 += v1
                                    _ = v44
                                    _ = v45
                                }
                                var v51 bool = false
                                _ = fi43
                                _ = v51
                            }
                            _ = v40
                        } else {
                            switch "" {
                            case "":
                                v3 = append(v3, 2147483647)
                            case "foo bar":
                                if ((true || true) || (!false)) {
                                    for fi52 := 0; fi52 < 9; fi52++ {
                                        if false {
                                            if (((float32(-2.25) - ((v1 - ((((v1 * 1e10) + ((v1 / v1) / float32(((float64(0.0) / 1.5) - v4)))) - (float32(v4) - (float32(float64(((-2.25 * v1) - float32((float64(1e10) + ((v4 + v4) * (3.14159 * (1.5 - (0.0 - v4))))))))) - (((v1 - v1) + 0.0) + (v1 + (((v1 + 0.1) - v1) * v1)))))) * 0.0)) / ((v1 + 1.5) - v1))) - v1) == (v1 - v1)) {
                                                v4 = v4
                                                switch "a" {
                                                case "hello":
                                                    fmt.Println(v1, v4)
                                                case "foo bar":
                                                    v53 := GenS2{F0: (true && true), F1: fi52, F2: 42}
                                                    _ = v53
                                                default:
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483648
                                                        }
                                                    }
                                                }
                                                v36 = -128
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                        }
                                        if true {
                                            fmt.Println(v4, v2)
                                            var v54 float32 = v1
                                            fmt.Println(v36)
                                            _ = v54
                                        }
                                        _ = fi52
                                    }
                                    for fi55 := 0; fi55 < 14; fi55++ {
                                        if (!true) {
                                            var v56 uint32 = 0
                                            v1 = (0.1 * -2.25)
                                            if (42 <= 0) {
                                                var v57 uint32 = 959
                                                switch (("hello" + "\n\t") + "\n\t") {
                                                case "foo bar":
                                                    var v58 int32 = -2147483648
                                                    _ = v58
                                                default:
                                                    v3 = append(v3, v57)
                                                }
                                                var v59 int8 = (v36 - (v36 * 0))
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v57
                                                _ = v59
                                            } else {
                                                var v60 int64 = 1
                                                var v61 float64 = float64(v1)
                                                if true {
                                                    _ = len(v3)
                                                    fmt.Println(v61)
                                                    v62 := GenS5{F0: 255, F1: -128}
                                                    _ = v62
                                                }
                                                _ = v60
                                                _ = v61
                                            }
                                            _ = v56
                                        } else {
                                            if true {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        }
                                        var v63 uint32 = 265
                                        switch (^(fi37)) {
                                        case 32768:
                                            if (true || false) {
                                                if (!(497 == genFn1(fi37, 42, 339))) {
                                                    v1 += 1.5
                                                    var v64 int32 = 1
                                                    _ = v64
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        if (-107 == 0) {
                                                            v65 := GenS5{F0: 0, F1: (v36 ^ (v36 << ((v2) & 63)))}
                                                            v2 += v2
                                                            _ = v65
                                                        }
                                                    case "":
                                                        var v66 uint8 = 0
                                                        var v67 uint32 = (^(v63))
                                                        var v68 uint8 = v66
                                                        v2 -= (^(v2))
                                                        var v69 uint32 = (v67 ^ 350)
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    default:
                                                        switch fi37 {
                                                        case -9223372036854775808:
                                                            if false {
                                                                v63 += 2147483647
                                                            }
                                                        case 32767:
                                                            if (-681 >= -851) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 2
                                                                }
                                                                v36 = -88
                                                                if (v1 <= v1) {
                                                                    v3 = append(v3, 1)
                                                                    var v70 int8 = (v36 >> ((18446744073709551615) & 63))
                                                                    _ = v70
                                                                }
                                                            }
                                                        default:
                                                            if false {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v63 += (^(v63))
                                                                }
                                                                v71 := GenS1{F0: -1, F1: (fi55 + fi55), F2: 30}
                                                                v36 = -1
                                                                v72 := GenS2{F0: false, F1: -579, F2: 2}
                                                                if (!false) {
                                                                    fmt.Println(v1)
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v63
                                                                    }
                                                                    fmt.Println(fi55, v63)
                                                                }
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                        }
                                                    }
                                                    if ((!(!false)) || (18446744073709551615 != 952)) {
                                                        fmt.Println(v2, fi37, fi37)
                                                    } else {
                                                        switch 1 {
                                                        case 9223372036854775807:
                                                            v73 := GenS5{F0: 236, F1: -1}
                                                            v2 = v2
                                                            if (!(uint(v63) == 18446744073709551615)) {
                                                                if true {
                                                                    var v74 int = fi37
                                                                    _ = v74
                                                                } else {
                                                                    v75 := GenS6{F0: "\n\t"}
                                                                    if ((-1 == -2147483648) || (-2.25 <= v1)) {
                                                                        var v76 int = fi55
                                                                        var v77 int16 = 1
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    v4 = float64(float32((v4 - ((v4 + (v4 - ((v4 + 1.5) + float64(float32(v4))))) - 3.14159))))
                                                                    var v78 int8 = (^(v36))
                                                                    v79 := GenS2{F0: true, F1: int(v63), F2: 0}
                                                                    _ = v75
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                fmt.Println(fi37, fi55)
                                                            } else {
                                                                v2 -= (v2 ^ v2)
                                                                if (((v4 - v4) / v4) > 3.14159) {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v63
                                                                    }
                                                                    var v80 float32 = 1e10
                                                                    _ = v80
                                                                } else {
                                                                    var v81 int64 = 9223372036854775807
                                                                    v1 = float32(((v4 * v4) + v4))
                                                                    switch 9223372036854775807 {
                                                                    case 128:
                                                                        if false {
                                                                            v2 -= (v2 + (v2 / (((v2 + (^(v2)))) | 1)))
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v63
                                                                            }
                                                                            if true {
                                                                                switch "hello" {
                                                                                case "":
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                case "a":
                                                                                    v82 := GenS6{F0: "a"}
                                                                                    _ = v82
                                                                                default:
                                                                                    v36 = (-(v36))
                                                                                }
                                                                                var v83 int32 = 0
                                                                                _ = v83
                                                                            }
                                                                        }
                                                                    case -1:
                                                                        var v84 string = ("hello" + "")
                                                                        _ = v84
                                                                    case 9223372036854775807:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    _ = v81
                                                                }
                                                            }
                                                            _ = v73
                                                        case 2147483647:
                                                            var v85 uint = (v2 - 996)
                                                            _ = v85
                                                        case -9223372036854775808:
                                                            if false {
                                                                v86 := GenS2{F0: false, F1: (fi37 & fi37), F2: 127}
                                                                v87 := GenS5{F0: 0, F1: v36}
                                                                fmt.Println(v63, fi37, v63)
                                                                _ = v86
                                                                _ = v87
                                                            } else {
                                                                var v88 float64 = v4
                                                                v88 -= v88
                                                                var v89 int8 = v36
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        default:
                                                            var v90 string = "\n\t"
                                                            _ = v90
                                                        }
                                                        var v91 string = ("" + "\n\t")
                                                        v92 := GenS3{F0: -588, F1: 2147483647}
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                    switch (fi37 - 1) {
                                                    case -649:
                                                        v63 -= (v63 / ((v63) | 1))
                                                    case 867:
                                                        v4 -= (v4 - -2.25)
                                                    case 128:
                                                        v36 += 85
                                                    default:
                                                        v1 = (v1 - v1)
                                                    }
                                                }
                                                if true {
                                                    var v93 int32 = 1
                                                    if (((1e10 * ((v4 - v4) + ((v4 * float64(((v1 - v1) / (1.5 * -2.25)))) * (0.1 - 0.1)))) + v4) < 3.14159) {
                                                        _ = len(v3)
                                                    }
                                                    var v94 string = "hello"
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    v95 := GenS7{F0: 255, F1: v36, F2: 42}
                                                    v2 = (v2 * (v2 * v2))
                                                    if true {
                                                        v96 := GenS8{F0: 2, F1: "foo bar"}
                                                        _ = v96
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (v63 * 418)
                                                        }
                                                        if true {
                                                            fmt.Println(v4)
                                                            fmt.Println(fi37)
                                                            var v97 string = "foo bar"
                                                            var v98 int8 = (v36 & v36)
                                                            _ = v97
                                                            _ = v98
                                                        } else {
                                                            if (32767 >= -145) {
                                                                var v99 bool = (false || true)
                                                                _ = v99
                                                            }
                                                            v1 = (1e10 / float32(v4))
                                                            v4 -= v4
                                                        }
                                                    }
                                                    _ = v95
                                                }
                                                if (-337 >= -1) {
                                                    var v100 uint64 = 42
                                                    _ = v100
                                                } else {
                                                    var v101 bool = (!false)
                                                    _ = v101
                                                }
                                                v3 = append(v3, (v63 * 169))
                                            }
                                        case -9223372036854775808:
                                            var v102 bool = true
                                            _ = v102
                                        default:
                                            v103 := GenS8{F0: 2, F1: "\n\t"}
                                            _ = v103
                                        }
                                        _ = fi55
                                        _ = v63
                                    }
                                }
                            case "a":
                                var v104 uint32 = 485
                                _ = v104
                            }
                        }
                        v36 = v36
                    }
                case 32768:
                    if (!(0 != -9223372036854775808)) {
                        v105 := GenS9{F0: 42, F1: -601, F2: 42}
                        _ = v105
                    } else {
                        v106 := GenS7{F0: 32767, F1: (v36 + 42), F2: 42}
                        switch "foo bar" {
                        case "\n\t":
                            for fi107 := 0; fi107 < 8; fi107++ {
                                switch fi37 {
                                case 9223372036854775807:
                                    var v108 uint64 = 42
                                    if (!(-2.25 != 0.1)) {
                                        v108 -= (v108 | 42)
                                        v36 = v36
                                        if true {
                                            if (255 == 87) {
                                                var v109 int8 = 52
                                                _ = v109
                                            } else {
                                                var v110 uint8 = 221
                                                fmt.Println(fi107, v1)
                                                _ = v110
                                            }
                                            if true {
                                                v111 := GenS10{F0: 661, F1: true, F2: 1}
                                                v112 := GenS6{F0: "\n\t"}
                                                fmt.Println(v1)
                                                _ = v111
                                                _ = v112
                                            }
                                            fmt.Println(v36, v1, fi107)
                                            var v113 int16 = 0
                                            _ = v113
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 536
                                            }
                                        }
                                    }
                                    switch ("" + (("\n\t" + ("foo bar" + "\n\t")) + "")) {
                                    case "hello":
                                        var v114 float32 = float32((v4 / v4))
                                        _ = v114
                                    case "":
                                        var v115 uint = v2
                                        v116 := GenS4{F0: 179}
                                        _ = v115
                                        _ = v116
                                    case "foo bar":
                                        v4 = ((v4 + v4) - v4)
                                    default:
                                        var v117 uint16 = 0
                                        _ = v117
                                    }
                                    _ = v108
                                case -9223372036854775808:
                                    v118 := GenS11{F0: (("hello" + "foo bar") != "hello")}
                                    _ = v118
                                case 4294967295:
                                    v4 -= (v4 * v4)
                                }
                                var v119 uint64 = 327
                                _ = fi107
                                _ = v119
                            }
                        default:
                            switch 32768 {
                            case 127:
                                v36 = v36
                            case -907:
                                if false {
                                    var v120 uint32 = 65535
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v120
                                }
                            case -1:
                                var v121 float32 = v1
                                _ = v121
                            default:
                                var v122 int16 = 32767
                                _ = v122
                            }
                        }
                        _ = v106
                    }
                }
                fmt.Println(v4, v1)
                v1 += float32(v4)
                _ = fi37
            }
        case 384:
            var v123 int = -866
            _ = v123
        }
    default:
        v4 += (v4 / 1.5)
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    var v124 uint16 = genFn1(838, 1, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
