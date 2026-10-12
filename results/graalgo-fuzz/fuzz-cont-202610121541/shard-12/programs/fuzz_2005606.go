package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 int, p1 uint16, p2 bool) int32 {
    var v7 uint = 788
    v7 -= v7
    for fi8 := 0; fi8 < 1; fi8++ {
        p2 = false
        v9 := GenS1{F0: 65535, F1: 42, F2: "a"}
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int32 = 391
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 int64 = -738
    p2 = (!false)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 uint32, p2 int64) int64 {
    if true {
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS1{F0: (p1 + p1), F1: 2, F2: "foo bar"}
                if false {
                    switch "a" {
                    case "":
                        p2 = (p2 << ((4294967295) & 63))
                    }
                    p2 -= p2
                    fmt.Println(p1, fi13, p2)
                } else {
                    var v16 uint32 = p1
                    if (fi14 == fi13) {
                        v17 := []int{-100}
                        _ = v17
                    }
                    p2 = p2
                    v18 := []int{}
                    _ = v16
                    _ = v18
                }
                if (false || true) {
                    fmt.Println(fi14)
                    var v19 int64 = (p2 ^ 42)
                    v20 := GenS2{F0: true, F1: 2}
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS1{F0: p1, F1: 2, F2: "hello"}
                    v22 := GenS3{F0: (p0 - 127), F1: "hello"}
                    fmt.Println(fi13, p2, p1)
                    _ = v21
                    _ = v22
                }
                _ = fi14
                _ = v15
            }
            for fi23 := 0; fi23 < 10; fi23++ {
                if ((p2 <= p2) && false) {
                    v24 := []int{(^(fi23)), (fi23 << ((255) & 63)), (fi13 / ((fi23) | 1)), fi13}
                    p2 -= (p2 ^ int64(fi23))
                    _ = v24
                }
                switch (fi23 ^ 1) {
                case 2147483647:
                    var v25 int64 = (p2 & 781)
                    _ = v25
                case 4294967295:
                    p1 = p1
                case 128:
                    v26 := []int{101}
                    _ = v26
                }
                _ = fi23
            }
            _ = fi13
        }
        p2 += 420
        switch "" {
        case "":
            v27 := GenS3{F0: (p0 * -1), F1: "a"}
            _ = v27
        }
    } else {
        if false {
            v28 := GenS1{F0: (^(p1)), F1: 42, F2: "foo bar"}
            if false {
                for fi29 := 0; fi29 < 12; fi29++ {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        p1 = p1
                        if (fi29 > (fi29 >> ((2) & 63))) {
                            if (65535 < p1) {
                                var v31 int64 = 9223372036854775807
                                var v32 float64 = 3.14159
                                if false {
                                    switch fi29 {
                                    case -944:
                                        v33 := []int{4294967295, (fi29 << ((562) & 63)), (fi30 << ((1) & 63))}
                                        _ = v33
                                    case 620:
                                        v32 = v32
                                    default:
                                        switch fi30 {
                                        case 2147483647:
                                            var v34 uint8 = 149
                                            _ = v34
                                        case -35:
                                            p2 += p2
                                        case 9223372036854775807:
                                            p2 += (v31 - (^(v31)))
                                        default:
                                            v35 := []int{int(p0)}
                                            _ = v35
                                        }
                                    }
                                    v36 := GenS4{F0: -1}
                                    v37 := GenS5{F0: (fi29 >> ((65535) & 63)), F1: 4294967295}
                                    _ = v36
                                    _ = v37
                                }
                                _ = v31
                                _ = v32
                            }
                            if (4294967295 != 446) {
                                var v38 float64 = 0.1
                                if true {
                                    v39 := []int{(fi30 * 9223372036854775807), 32767}
                                    var v40 uint64 = 2
                                    switch "hello" {
                                    case "\n\t":
                                        switch ("hello" + ("" + "a")) {
                                        case "foo bar":
                                            p1 += (p1 | p1)
                                        }
                                    case "a":
                                        v41 := GenS1{F0: (p1 * p1), F1: (v40 * v40), F2: (("\n\t" + ("a" + "\n\t")) + "a")}
                                        _ = v41
                                    case "hello":
                                        v42 := GenS5{F0: (^(fi29)), F1: (v40 / ((18446744073709551615) | 1))}
                                        _ = v42
                                    default:
                                        if (true || true) {
                                            fmt.Println(v40, p1, v38)
                                            v39 = append(v39, (fi29 + 280))
                                            v43 := GenS3{F0: 124, F1: "\n\t"}
                                            _ = v43
                                        }
                                    }
                                    p1 += p1
                                    _ = v39
                                    _ = v40
                                }
                                _ = v38
                            }
                            p1 = p1
                        } else {
                            switch (((("" + "") + ("" + "")) + (((("\n\t" + "a") + "") + "foo bar") + "a")) + ("a" + "a")) {
                            case "foo bar":
                                var v44 uint16 = 263
                                _ = v44
                            }
                        }
                        v45 := []int{(fi29 & (fi30 + (fi30 * int(p0)))), (fi29 / ((fi29) | 1)), (fi29 ^ 79), (fi30 % ((32768) | 1))}
                        _ = fi30
                        _ = v45
                    }
                    _ = fi29
                }
                for fi46 := 0; fi46 < 19; fi46++ {
                    var v47 int64 = p2
                    v48 := GenS2{F0: true, F1: 51}
                    for fi49 := 0; fi49 < 13; fi49++ {
                        p1 -= 65535
                        if true {
                            switch "foo bar" {
                            case "hello":
                                var v50 uint64 = 4294967295
                                _ = v50
                            case "a":
                                v51 := []int{fi46, fi46, fi49, fi49}
                                _ = v51
                            case "":
                                if false {
                                    p2 = p2
                                }
                            default:
                                p0 = p0
                            }
                            if (((2 > 65535) && (32767 >= 597)) || true) {
                                if ("hello" < "foo bar") {
                                    switch (fi49 % ((fi49) | 1)) {
                                    case 9223372036854775807:
                                        var v52 uint = 872
                                        _ = v52
                                    }
                                } else {
                                    v53 := []int{(fi49 << ((65535) & 63)), (fi46 << ((263) & 63)), -2147483648}
                                    if len(v53) > 0 {
                                        v53[0] = fi49
                                    }
                                    _ = v53
                                }
                            }
                            p2 -= (v47 >> ((65535) & 63))
                        }
                        var v54 uint16 = 65535
                        _ = fi49
                        _ = v54
                    }
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                fmt.Println(p0, p0, p1)
            } else {
                for fi55 := 0; fi55 < 8; fi55++ {
                    for fi56 := 0; fi56 < 7; fi56++ {
                        v57 := GenS3{F0: (p0 >> ((0) & 63)), F1: ""}
                        v58 := []int{128}
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    if false {
                        v59 := GenS2{F0: false, F1: 127}
                        _ = v59
                    }
                    var v60 uint32 = p1
                    _ = fi55
                    _ = v60
                }
                switch 9223372036854775807 {
                case 42:
                    fmt.Println(p0)
                case 128:
                    for fi61 := 0; fi61 < 6; fi61++ {
                        var v62 uint16 = 0
                        _ = fi61
                        _ = v62
                    }
                default:
                    for fi63 := 0; fi63 < 3; fi63++ {
                        for fi64 := 0; fi64 < 14; fi64++ {
                            var v65 int8 = -1
                            if false {
                                if false {
                                    v66 := GenS6{F0: 255, F1: 127}
                                    _ = v66
                                } else {
                                    if (255 == 127) {
                                        p2 = 1
                                        var v67 float32 = float32(3.14159)
                                        if true {
                                            switch 975 {
                                            case 127:
                                                v68 := GenS1{F0: (p1 - (p1 ^ (p1 % ((0) | 1)))), F1: 1, F2: "hello"}
                                                _ = v68
                                            case -1:
                                                var v69 uint16 = 2
                                                _ = v69
                                            case 2147483647:
                                                if true {
                                                    p1 += (p1 % ((2) | 1))
                                                    v70 := []int{(fi64 % (((fi64 - fi64)) | 1)), (^(fi64))}
                                                    switch ("\n\t" + "a") {
                                                    case "foo bar":
                                                        v71 := GenS3{F0: -128, F1: ("a" + "foo bar")}
                                                        _ = v71
                                                    default:
                                                        p1 = 42
                                                    }
                                                    _ = v70
                                                }
                                            default:
                                                fmt.Println(p0)
                                                var v72 int8 = v65
                                                _ = v72
                                            }
                                        } else {
                                            v73 := []int{}
                                            p1 = (p1 + p1)
                                            _ = len(v73)
                                            _ = v73
                                        }
                                        _ = v67
                                    }
                                    switch fi63 {
                                    case -550:
                                        fmt.Println(v65)
                                    case 586:
                                        if (p0 < p0) {
                                            p2 -= 920
                                            v74 := []int{}
                                            v75 := GenS4{F0: p2}
                                            p2 = p2
                                            _ = v74
                                            _ = v75
                                        }
                                    default:
                                        fmt.Println(p1, v65, fi63)
                                    }
                                    v76 := GenS4{F0: p2}
                                    _ = v76
                                }
                                if (!(!(true || false))) {
                                    if true {
                                        var v77 int8 = v65
                                        switch "" {
                                        case "foo bar":
                                            switch (^(fi63)) {
                                            case -9223372036854775808:
                                                p1 -= p1
                                            case -2147483648:
                                                v78 := GenS4{F0: p2}
                                                _ = v78
                                            case -1:
                                                var v79 uint32 = (p1 + (p1 >> ((18446744073709551615) & 63)))
                                                _ = v79
                                            }
                                        case "\n\t":
                                            v80 := GenS3{F0: (p0 & 81), F1: (("" + "hello") + "")}
                                            _ = v80
                                        default:
                                            v81 := []int{fi64, fi64}
                                            _ = v81
                                        }
                                        _ = v77
                                    }
                                } else {
                                    var v82 int64 = (p2 ^ p2)
                                    _ = v82
                                }
                            }
                            p2 += p2
                            _ = fi64
                            _ = v65
                        }
                        _ = fi63
                    }
                }
            }
            _ = v28
        } else {
            v83 := []int{}
            fmt.Println(p2, p0)
            _ = v83
        }
    }
    v84 := []int{0, -335, 1}
    v85 := GenS2{F0: false, F1: 2}
    switch -980 {
    case -177:
        if (4294967295 != -982) {
            v86 := GenS4{F0: (^(p2))}
            fmt.Println(p0, p1, p2)
            _ = len(v84)
            _ = v86
        }
    case -112:
        var v87 float32 = 3.14159
        _ = v87
    case 127:
        p1 += p1
    default:
        fmt.Println(p2, p2)
    }
    switch "\n\t" {
    case "":
        p2 -= p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v84
    _ = v85
    return -1
}

func main() {
    var v1 float64 = (1.5 / (3.14159 * 1e10))
    var v2 uint64 = 42
    var v3 int16 = 0
    var v4 uint32 = 311
    v5 := make(map[uint32]uint32)
    v6 := []uint32{(v4 >> ((435) & 63)), 42}
    v2 = v2
    v88 := v5[v4]
    fmt.Println(v3, v88, v3)
    fmt.Println(v2)
    v1 += (-2.25 * (v1 + v1))
    var v89 uint = uint(v88)
    v3 -= -201
    switch "a" {
    case "\n\t":
        if (true && ((v2 | (v2 - (v2 / ((uint64(v4)) | 1)))) >= (^(v2)))) {
            var v90 int64 = 240
            v88 += 1
            v91 := GenS7{F0: (v3 ^ (v3 % (((v3 % ((-32768) | 1))) | 1))), F1: 65535}
            _ = v90
            _ = v91
        }
    case "foo bar":
        for fi92 := 0; fi92 < 8; fi92++ {
            var v93 int8 = 42
            if false {
                v94 := v5[v4]
                fmt.Println(v2, v1, v89)
                _ = v94
            }
            _ = fi92
            _ = v93
        }
    case "":
        fmt.Println(v88, v2)
    default:
        v4 += 4294967295
    }
    switch "hello" {
    case "a":
        v95 := GenS6{F0: (v89 + (v89 | (v89 << ((uint(v88)) & 63)))), F1: 42}
        _ = v95
    case "\n\t":
        v89 += v89
    }
    switch ("" + "hello") {
    case "":
        for fi96 := 0; fi96 < 8; fi96++ {
            var v97 int = (^(fi96))
            fmt.Println(v2, fi96, v89)
            var v98 int8 = 42
            _ = fi96
            _ = v97
            _ = v98
        }
    case "hello":
        var v99 int16 = (v3 * (v3 / (((v3 % ((v3) | 1))) | 1)))
        _ = v99
    case "\n\t":
        v100 := GenS8{F0: 42, F1: (v4 % (((v4 | 65535)) | 1)), F2: v4}
        _ = v100
    default:
        v4 += 131
    }
    v89 += (v89 & 0)
    for fi101 := 0; fi101 < 12; fi101++ {
        var v102 uint = (v89 - v89)
        var v103 int64 = -1
        _ = fi101
        _ = v102
        _ = v103
    }
    v88 = (v4 ^ v88)
    for fi104 := 0; fi104 < 19; fi104++ {
        for fi105 := 0; fi105 < 18; fi105++ {
            var v106 float64 = ((v1 - 3.14159) + v1)
            _ = fi105
            _ = v106
        }
        _ = fi104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v89
}
