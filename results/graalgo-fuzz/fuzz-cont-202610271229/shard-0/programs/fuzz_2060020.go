package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int32; F2 uint16 }
type GenS8 struct { F0 int8; F1 int16; F2 int32 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 int64; F1 int8; F2 uint64 }
type GenS11 struct { F0 int; F1 string; F2 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8; F1 uint8 }
type GenS14 struct { F0 int16; F1 string; F2 int64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int64; F1 uint32 }
type GenS18 struct { F0 uint; F1 string; F2 uint }
type GenS19 struct { F0 int8; F1 uint8 }
type GenS20 struct { F0 uint16; F1 bool }
type GenS21 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint32; F1 int16 }
type GenS24 struct { F0 string; F1 uint }
type GenS25 struct { F0 uint8; F1 int64; F2 int }
func genFn0(p0 int, p1 int8, p2 int) int {
    var v5 uint = 4294967295
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{(p0 + (fi7 + (p0 - (p2 << ((65535) & 63))))), (p0 * p0), (p0 * (fi7 & 42))}
        for fi9 := 0; fi9 < 4; fi9++ {
            switch p0 {
            case 127:
                var v10 uint64 = 18446744073709551615
                _ = v10
            case 9223372036854775807:
                var v11 uint32 = 76
                _ = v11
            case 418:
                switch ("hello" + ("" + "hello")) {
                case "":
                    fmt.Println(fi9, v6)
                case "\n\t":
                    v12 := GenS1{F0: 207, F1: "\n\t", F2: ((v6 + (0.0 * v6)) >= 1e10)}
                    _ = v12
                case "foo bar":
                    v13 := GenS2{F0: (p1 << ((v5) & 63)), F1: v5}
                    _ = v13
                }
            default:
                switch "\n\t" {
                case "a":
                    var v14 int64 = -9223372036854775808
                    _ = v14
                case "\n\t":
                    if len(v8) > 0 {
                        v8[0] = fi7
                    }
                    var v15 int8 = (p1 % (((p1 - (p1 & (p1 + (p1 | -1))))) | 1))
                    _ = v15
                default:
                    v16 := GenS2{F0: (p1 / ((-103) | 1)), F1: 524}
                    _ = v16
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 180
}
func genFn1(p0 bool, p1 uint16, p2 bool) int16 {
    p1 = p1
    if (!(873 <= 1000)) {
        var v17 int64 = 42
        var v18 uint64 = 4294967295
        v18 += (v18 - (v18 & (^(v18))))
        _ = v17
        _ = v18
    } else {
        var v19 uint32 = 72
        p2 = (4294967295 > 18446744073709551615)
        for fi20 := 0; fi20 < 3; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint16 = p1
                v23 := GenS3{F0: 886, F1: 2}
                p1 -= v22
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        _ = v19
    }
    var v24 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return -1
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = 1
    var v3 float64 = 1e10
    var v4 uint16 = 0
    if true {
        v3 += ((v3 / ((v3 * 1.5) - (0.1 - v3))) + (v3 - 0.1))
    }
    var v25 string = "a"
    v26 := GenS1{F0: 44, F1: (v25 + ("foo bar" + "hello")), F2: true}
    var v27 int64 = 42
    if true {
        v28 := GenS2{F0: 1, F1: 255}
        v4 -= (v4 >> ((19) & 63))
        if (1 < 42) {
            var v29 uint = 4294967295
            v30 := GenS2{F0: 127, F1: v29}
            if true {
                if true {
                    v31 := GenS3{F0: 0, F1: uint64(v29)}
                    _ = v31
                } else {
                    if (v3 != v3) {
                        v32 := []int{}
                        for fi33 := 0; fi33 < 2; fi33++ {
                            switch fi33 {
                            case -598:
                                for fi34 := 0; fi34 < 2; fi34++ {
                                    v35 := GenS3{F0: (^(v27)), F1: 2}
                                    v32 = append(v32, 9223372036854775807)
                                    _ = fi34
                                    _ = v35
                                }
                            case 32768:
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    v32 = append(v32, -2147483648)
                                    _ = len(v32)
                                    switch v25 {
                                    case "\n\t":
                                        switch (fi36 << (((v29 / ((uint(v4)) | 1))) & 63)) {
                                        case 9223372036854775807:
                                            var v37 int8 = 1
                                            _ = v37
                                        }
                                    case "":
                                        var v38 uint = v29
                                        _ = v38
                                    case "hello":
                                        fmt.Println(v1)
                                    default:
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                    }
                                    if (false || false) {
                                        var v39 int32 = -1
                                        fmt.Println(fi33, v3, v25)
                                        var v40 int64 = v27
                                        switch (v25 + v25) {
                                        case "":
                                            var v41 int64 = v27
                                            _ = v41
                                        }
                                        if false {
                                            v42 := GenS4{F0: ""}
                                            v1 = (((v1 - 3.14159) + v1) - ((((v1 / 1.0) / v1) - float32(0.1)) - (((1e10 - v1) - (((v1 * 0.0) - (3.14159 * ((0.0 + v1) - float32((v3 - float64((((v1 / 3.14159) * (v1 - (((float32(v3) - float32((float64(0.1) + (((v3 - v3) / 1.0) - (v3 * (v3 / (((float64(v1) * -2.25) - ((0.0 + 0.0) / (float64(v1) * -2.25))) * (0.0 + 1.5)))))))) + (v1 - (float32(v3) + (v1 / 3.14159)))) - v1))) * (v1 + (3.14159 / -2.25))))))))) - ((3.14159 / float32((v3 / 0.1))) + v1))) * v1)))
                                            _ = v42
                                        } else {
                                            switch genFn0((^(fi36)), 127, (fi36 % ((fi33) | 1))) {
                                            case -1:
                                                switch ((v25 + v25) + "a") {
                                                case "\n\t":
                                                    if false {
                                                        fmt.Println(v25)
                                                    }
                                                }
                                            default:
                                                v25 = "\n\t"
                                            }
                                            var v43 uint32 = 2147483647
                                            if false {
                                                switch "foo bar" {
                                                case "":
                                                    if ((^(v4)) < 42) {
                                                        v44 := GenS1{F0: 42, F1: v25, F2: false}
                                                        _ = v44
                                                    }
                                                case "a":
                                                    v32 = append(v32, (fi33 % ((fi33) | 1)))
                                                default:
                                                    var v45 string = "foo bar"
                                                    _ = v45
                                                }
                                            }
                                            _ = v43
                                        }
                                        v2 = v2
                                        var v46 int64 = v40
                                        _ = v39
                                        _ = v40
                                        _ = v46
                                    }
                                    _ = fi36
                                }
                            default:
                                v32 = append(v32, genFn0(-9223372036854775808, -128, fi33))
                            }
                            if true {
                                if len(v32) >= 2 {
                                    v32 = v32[1:len(v32)]
                                }
                                var v47 uint16 = v4
                                _ = v47
                            }
                            _ = fi33
                        }
                        v1 = -2.25
                        _ = v32
                    } else {
                        v48 := []int{307, 2147483647}
                        var v49 int32 = 2147483647
                        var v50 uint64 = 18446744073709551615
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    v51 := []int{4294967295, 754, 1, -730}
                    _ = v51
                }
                for fi52 := 0; fi52 < 15; fi52++ {
                    if (((fi52 * -419) != fi52) && true) {
                        fmt.Println(v1, v25, v27)
                        v1 += v1
                    } else {
                        v29 -= 18446744073709551615
                        switch "foo bar" {
                        case "a":
                            v29 = v29
                        case "hello":
                            switch 32768 {
                            case 1:
                                v29 += v29
                            case 349:
                                if ("foo bar" >= "hello") {
                                    fmt.Println(fi52, v1)
                                } else {
                                    var v53 bool = true
                                    switch fi52 {
                                    case 127:
                                        v54 := []int{fi52, genFn0((fi52 ^ (fi52 & (fi52 / (((^(fi52))) | 1)))), int8(v29), fi52)}
                                        _ = v54
                                    case 0:
                                        if true {
                                            v55 := GenS5{F0: ((v25 + v25) + (((v25 + "\n\t") + v25) + (v25 + v25))), F1: (-(v27)), F2: (v4 ^ (v4 - (v4 / ((1) | 1))))}
                                            v29 = v29
                                            _ = v55
                                        }
                                    case -478:
                                        if ("foo bar" < "") {
                                            if (((((float64(1.5) / 1.0) - v3) / (v3 / 3.14159)) > 0.1) || false) {
                                                v56 := []int{}
                                                v4 = v4
                                                _ = v56
                                            } else {
                                                fmt.Println(v53, v2, v2)
                                                var v57 int8 = 42
                                                _ = v57
                                            }
                                            if (18446744073709551615 == 4294967295) {
                                                var v58 uint = v29
                                                switch v25 {
                                                case "hello":
                                                    v59 := GenS6{F0: (v2 * genFn1((-1 <= (v2 - (v2 >> (((v58 | (v29 / (((v29 % ((v29) | 1))) | 1)))) & 63)))), 32767, (v4 == (v4 * (v4 ^ v4)))))}
                                                    _ = v59
                                                }
                                                v60 := []int{1, genFn0(int(v4), 20, fi52)}
                                                _ = v58
                                                _ = v60
                                            }
                                        } else {
                                            if ((1.5 / v3) <= ((v3 * v3) * v3)) {
                                                for fi61 := 0; fi61 < 9; fi61++ {
                                                    var v62 uint8 = 72
                                                    if (!(26 == 0)) {
                                                        v63 := GenS5{F0: "a", F1: int64(fi52), F2: v4}
                                                        v64 := GenS2{F0: 67, F1: uint(v4)}
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v53 = (!true)
                                                        v2 += (^(v2))
                                                    }
                                                    _ = fi61
                                                    _ = v62
                                                }
                                                switch v25 {
                                                case "a":
                                                    var v65 uint32 = 65535
                                                    _ = v65
                                                case "foo bar":
                                                    var v66 uint8 = 1
                                                    _ = v66
                                                case "\n\t":
                                                    if (v53 || v53) {
                                                        var v67 int = (fi52 * 32768)
                                                        var v68 string = ((v25 + (("foo bar" + v25) + "")) + v25)
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            } else {
                                                switch ("foo bar" + "\n\t") {
                                                case "hello":
                                                    var v69 uint16 = 831
                                                    _ = v69
                                                }
                                                switch (v25 + v25) {
                                                case "hello":
                                                    if v53 {
                                                        if false {
                                                            v1 = (float32(3.14159) + (0.1 * 1.5))
                                                            if true {
                                                                if v53 {
                                                                    v3 = (0.1 + 1.5)
                                                                    var v70 float32 = v1
                                                                    for fi71 := 0; fi71 < 18; fi71++ {
                                                                        var v72 uint = v29
                                                                        _ = fi71
                                                                        _ = v72
                                                                    }
                                                                    _ = v70
                                                                } else {
                                                                    for fi73 := 0; fi73 < 5; fi73++ {
                                                                        v53 = false
                                                                        if false {
                                                                            var v74 bool = (0.1 >= v1)
                                                                            var v75 uint8 = 2
                                                                            v2 -= v2
                                                                            _ = v74
                                                                            _ = v75
                                                                        } else {
                                                                            var v76 int32 = 42
                                                                            if v53 {
                                                                                if (v29 != v29) {
                                                                                    v77 := []int{(fi73 / (((fi52 & fi73)) | 1)), fi73, (fi52 << (((v29 << ((v29) & 63))) & 63)), 9223372036854775807}
                                                                                    v53 = v53
                                                                                    _ = v77
                                                                                }
                                                                                v78 := GenS7{F0: 127, F1: (v76 & -1), F2: 0}
                                                                                _ = v78
                                                                            } else {
                                                                                v79 := []int{fi52, 32767, (fi73 ^ fi52)}
                                                                                if ((1 == (^(v29))) || v53) {
                                                                                    v76 = (v76 << (((v29 | v29)) & 63))
                                                                                    v79 = append(v79, int(v2))
                                                                                } else {
                                                                                    if true {
                                                                                        switch fi52 {
                                                                                        case 32768:
                                                                                            fmt.Println(v3, v29)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v79 = append(v79, fi73)
                                                                                _ = v79
                                                                            }
                                                                            switch fi52 {
                                                                            case 552:
                                                                                v80 := GenS8{F0: 100, F1: 1000, F2: (v76 >> (((v29 * 576)) & 63))}
                                                                                _ = v80
                                                                            default:
                                                                                if false {
                                                                                    fmt.Println(v53)
                                                                                    if v53 {
                                                                                        var v81 float64 = v3
                                                                                        var v82 uint8 = 2
                                                                                        fmt.Println(v29, v3)
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    } else {
                                                                                        v2 = (-(v2))
                                                                                        v4 -= v4
                                                                                        v76 = v76
                                                                                    }
                                                                                } else {
                                                                                    var v83 float64 = ((1e10 + (((v3 - v3) + -2.25) - 3.14159)) + v3)
                                                                                    v84 := GenS9{F0: v25, F1: 988}
                                                                                    v85 := []int{fi73, fi52, (fi73 & genFn0(genFn0(fi52, 100, 42), -30, genFn0(genFn0((fi52 << ((v29) & 63)), 87, (^(fi73))), 42, int(v27)))), (fi52 & fi73)}
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                }
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                        switch (fi73 * 0) {
                                                                        case 808:
                                                                            fmt.Println(v25, v25)
                                                                        case -151:
                                                                            v4 = 615
                                                                        default:
                                                                            var v86 uint32 = 651
                                                                            _ = v86
                                                                        }
                                                                        v3 -= v3
                                                                        _ = fi73
                                                                    }
                                                                    switch fi52 {
                                                                    case -523:
                                                                        for fi87 := 0; fi87 < 2; fi87++ {
                                                                            var v88 int = fi52
                                                                            fmt.Println(v3, v3)
                                                                            _ = fi87
                                                                            _ = v88
                                                                        }
                                                                    default:
                                                                        for fi89 := 0; fi89 < 7; fi89++ {
                                                                            fmt.Println(v25, v1)
                                                                            _ = fi89
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            if v53 {
                                                                v90 := GenS9{F0: "hello", F1: 28}
                                                                v29 = (v29 >> ((0) & 63))
                                                                _ = v90
                                                            } else {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v91 := []int{int(v2)}
                                                                    _ = v91
                                                                default:
                                                                    for fi92 := 0; fi92 < 9; fi92++ {
                                                                        var v93 uint64 = 1
                                                                        if (uint32(v93) != 42) {
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                fmt.Println(v27, v2, v93)
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v94 := []int{int(v2), (fi92 - genFn0(genFn0(fi92, 1, (-(fi92))), -1, fi52)), -29, fi52}
                                                                                    _ = v94
                                                                                case "hello":
                                                                                    var v95 uint64 = (v93 ^ v93)
                                                                                    _ = v95
                                                                                }
                                                                                fmt.Println(v25, v1, v53)
                                                                            case "a":
                                                                                var v96 int16 = -290
                                                                                _ = v96
                                                                            case "hello":
                                                                                v97 := GenS4{F0: v25}
                                                                                _ = v97
                                                                            default:
                                                                                v53 = v53
                                                                            }
                                                                            var v98 uint8 = 42
                                                                            switch (v25 + v25) {
                                                                            case "foo bar":
                                                                                v99 := []int{(fi92 % (((-(fi52))) | 1)), genFn0(genFn0(fi92, 69, 32767), 42, (^(fi52))), 2147483647, (fi52 / ((-9223372036854775808) | 1))}
                                                                                switch int(v4) {
                                                                                case 4294967295:
                                                                                    v100 := GenS5{F0: v25, F1: (v27 & 42), F2: v4}
                                                                                    _ = v100
                                                                                case 9223372036854775807:
                                                                                    v93 += v93
                                                                                }
                                                                                if (v3 >= v3) {
                                                                                    var v101 uint = 199
                                                                                    v98 = v98
                                                                                    var v102 int16 = (v2 >> ((18446744073709551615) & 63))
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                } else {
                                                                                    var v103 uint = (v29 % ((2) | 1))
                                                                                    _ = v103
                                                                                }
                                                                                if false {
                                                                                    var v104 uint32 = 1
                                                                                    if len(v99) > 0 {
                                                                                        v99[0] = 32767
                                                                                    }
                                                                                    v25 = v25
                                                                                    if false {
                                                                                        var v105 int8 = 127
                                                                                        switch v25 {
                                                                                        case "a":
                                                                                            if false {
                                                                                                switch fi52 {
                                                                                                case 32767:
                                                                                                    _ = len(v99)
                                                                                                }
                                                                                            } else {
                                                                                                switch v25 {
                                                                                                case "\n\t":
                                                                                                    if len(v99) > 0 {
                                                                                                        v99[0] = (fi52 % (((^(fi52))) | 1))
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v106 := GenS6{F0: 0}
                                                                                                    _ = v106
                                                                                                default:
                                                                                                    v1 -= ((0.0 / float32(v3)) / float32((((float64(v1) * (v3 * (1e10 / (v3 / (v3 * (0.1 - (v3 * -2.25))))))) / v3) / 1.5)))
                                                                                                }
                                                                                                v107 := GenS7{F0: -55, F1: 0, F2: uint16(v104)}
                                                                                                var v108 uint8 = v98
                                                                                                fmt.Println(v29)
                                                                                                _ = v107
                                                                                                _ = v108
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v109 := GenS8{F0: v105, F1: v2, F2: 0}
                                                                                            _ = v109
                                                                                        case "":
                                                                                            v4 = v4
                                                                                        default:
                                                                                            var v110 int = genFn0((fi52 + (^(fi52))), v105, (-(fi92)))
                                                                                            _ = v110
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            v111 := GenS8{F0: int8(v93), F1: v2, F2: 0}
                                                                                            _ = v111
                                                                                        }
                                                                                        fmt.Println(v27, v25)
                                                                                        var v112 uint16 = (v4 >> (((v29 & 0)) & 63))
                                                                                        _ = v105
                                                                                        _ = v112
                                                                                    }
                                                                                    _ = v104
                                                                                } else {
                                                                                    if v53 {
                                                                                        fmt.Println(v93, v98, v93)
                                                                                        v2 = genFn1((0 <= 4294967295), 423, (true || v53))
                                                                                    }
                                                                                }
                                                                                _ = v99
                                                                            case "\n\t":
                                                                                v93 = 2
                                                                            case "a":
                                                                                var v113 uint64 = v93
                                                                                _ = v113
                                                                            }
                                                                            var v114 int16 = v2
                                                                            fmt.Println(v3, fi92, fi52)
                                                                            _ = v114
                                                                            _ = v98
                                                                        } else {
                                                                            fmt.Println(fi92, v29, v53)
                                                                            v115 := GenS10{F0: (-(v27)), F1: 42, F2: (v93 >> (((v29 >> ((v29) & 63))) & 63))}
                                                                            v116 := GenS1{F0: 37, F1: "a", F2: (!v53)}
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                        _ = fi92
                                                                        _ = v93
                                                                    }
                                                                }
                                                                for fi117 := 0; fi117 < 2; fi117++ {
                                                                    if v53 {
                                                                        fmt.Println(v1)
                                                                        switch v25 {
                                                                        case "\n\t":
                                                                            if (1 <= 42) {
                                                                                v53 = true
                                                                                var v118 int32 = 1
                                                                                _ = v118
                                                                            } else {
                                                                                v29 = (v29 ^ (v29 & (^(v29))))
                                                                                if ((v29 * v29) >= (^(v29))) {
                                                                                    if (2 == 86) {
                                                                                        v119 := GenS6{F0: -1}
                                                                                        var v120 uint64 = 1
                                                                                        switch (((v25 + (v25 + (v25 + "a"))) + ("\n\t" + (v25 + "foo bar"))) + v25) {
                                                                                        case "\n\t":
                                                                                            v121 := GenS11{F0: genFn0((fi117 >> (((^(v29))) & 63)), 17, (fi52 << ((18446744073709551615) & 63))), F1: (v25 + (v25 + (("\n\t" + "hello") + (v25 + "foo bar")))), F2: 0}
                                                                                            _ = v121
                                                                                        case "":
                                                                                            var v122 int64 = (v27 >> ((v29) & 63))
                                                                                            _ = v122
                                                                                        case "a":
                                                                                            fmt.Println(v1, v2, v3)
                                                                                        default:
                                                                                            v27 = (v27 / (((v27 - (v27 | -393))) | 1))
                                                                                        }
                                                                                        _ = v119
                                                                                        _ = v120
                                                                                    }
                                                                                    var v123 float64 = 0.0
                                                                                    _ = v123
                                                                                }
                                                                                fmt.Println(v29)
                                                                            }
                                                                        case "a":
                                                                            v124 := []int{(fi117 >> ((v29) & 63)), fi52, (fi52 - 42)}
                                                                            _ = v124
                                                                        case "hello":
                                                                            switch fi52 {
                                                                            case 415:
                                                                                v125 := []int{genFn0(fi52, -1, (fi52 % ((fi52) | 1))), (fi52 - (-(fi52)))}
                                                                                _ = v125
                                                                            case 2147483647:
                                                                                if v53 {
                                                                                    var v126 int = (fi117 * 127)
                                                                                    var v127 string = v25
                                                                                    v3 -= ((v3 - 1e10) - 0.0)
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                }
                                                                            default:
                                                                                if v53 {
                                                                                    v27 += -1
                                                                                    v128 := GenS6{F0: genFn1((185 <= 0), v4, (-121 <= 1))}
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                        }
                                                                        var v129 uint = v29
                                                                        v27 += 1
                                                                        _ = v129
                                                                    } else {
                                                                        v25 = v25
                                                                        v3 += 1e10
                                                                        v130 := GenS12{F0: v29}
                                                                        if v53 {
                                                                            v131 := GenS13{F0: 51, F1: 248}
                                                                            _ = v131
                                                                        }
                                                                        _ = v130
                                                                    }
                                                                    var v132 uint = v29
                                                                    var v133 uint = (v132 - (^(v132)))
                                                                    _ = fi117
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                                v27 = 677
                                                            }
                                                            for fi134 := 0; fi134 < 7; fi134++ {
                                                                var v135 int16 = 724
                                                                _ = fi134
                                                                _ = v135
                                                            }
                                                            if v53 {
                                                                if false {
                                                                    if (-2.25 < float64(v1)) {
                                                                        var v136 int64 = 42
                                                                        v137 := GenS14{F0: v2, F1: ("hello" + "\n\t"), F2: (v27 * 466)}
                                                                        switch fi52 {
                                                                        case -1:
                                                                            fmt.Println(fi52, v27)
                                                                        case 337:
                                                                            switch ("a" + v25) {
                                                                            case "a":
                                                                                for fi138 := 0; fi138 < 15; fi138++ {
                                                                                    v29 -= v29
                                                                                    _ = fi138
                                                                                }
                                                                            case "foo bar":
                                                                                v139 := GenS7{F0: 0, F1: 497, F2: v4}
                                                                                _ = v139
                                                                            default:
                                                                                var v140 float64 = v3
                                                                                _ = v140
                                                                            }
                                                                        default:
                                                                            var v141 uint16 = uint16(v29)
                                                                            _ = v141
                                                                        }
                                                                        _ = v136
                                                                        _ = v137
                                                                    }
                                                                    for fi142 := 0; fi142 < 6; fi142++ {
                                                                        switch ((v25 + "") + (v25 + ("a" + v25))) {
                                                                        case "hello":
                                                                            fmt.Println(fi142)
                                                                        case "":
                                                                            if true {
                                                                                v27 -= -9223372036854775808
                                                                            }
                                                                        default:
                                                                            v143 := GenS15{F0: 2147483647}
                                                                            _ = v143
                                                                        }
                                                                        _ = fi142
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        for fi144 := 0; fi144 < 20; fi144++ {
                                                            fmt.Println(v27, fi144)
                                                            v145 := GenS6{F0: (v2 >> ((691) & 63))}
                                                            v25 = v25
                                                            _ = fi144
                                                            _ = v145
                                                        }
                                                    } else {
                                                        v25 = v25
                                                        fmt.Println(v29)
                                                    }
                                                case "":
                                                    if (!(v3 == v3)) {
                                                        v146 := GenS16{F0: 152}
                                                        _ = v146
                                                    } else {
                                                        var v147 int32 = 1
                                                        _ = v147
                                                    }
                                                }
                                                v148 := []int{(fi52 * fi52)}
                                                _ = v148
                                            }
                                            v2 -= 0
                                            v149 := GenS16{F0: 111}
                                            _ = v149
                                        }
                                    default:
                                        v150 := []int{fi52, -606, (fi52 - fi52)}
                                        _ = v150
                                    }
                                    v151 := GenS8{F0: -12, F1: genFn1(false, (v4 << (((^(v29))) & 63)), v53), F2: 1}
                                    _ = v151
                                    _ = v53
                                }
                            default:
                                var v152 uint16 = 42
                                _ = v152
                            }
                        default:
                            v153 := []int{(fi52 * 42), -69, (^(fi52))}
                            _ = v153
                        }
                        v154 := GenS17{F0: v27, F1: 657}
                        _ = v154
                    }
                    v155 := []int{0}
                    v156 := GenS18{F0: uint(v4), F1: "hello", F2: uint(v4)}
                    _ = fi52
                    _ = v155
                    _ = v156
                }
            }
            _ = v29
            _ = v30
        } else {
            v4 += (v4 - v4)
        }
        _ = v28
    } else {
        switch "hello" {
        case "":
            var v157 uint16 = 32767
            _ = v157
        case "\n\t":
            var v158 int16 = -1
            _ = v158
        case "a":
            if true {
                for fi159 := 0; fi159 < 4; fi159++ {
                    if (v25 <= "foo bar") {
                        if ((fi159 + (fi159 - -1)) > genFn0((-(fi159)), 1, fi159)) {
                            for fi160 := 0; fi160 < 17; fi160++ {
                                var v161 int = fi159
                                _ = fi160
                                _ = v161
                            }
                            v162 := GenS19{F0: -29, F1: 127}
                            _ = v162
                        } else {
                            switch (fi159 | fi159) {
                            case 42:
                                for fi163 := 0; fi163 < 3; fi163++ {
                                    v164 := []int{}
                                    _ = fi163
                                    _ = v164
                                }
                            case 127:
                                v165 := GenS20{F0: (v4 - (v4 | v4)), F1: (724 == 4294967295)}
                                _ = v165
                            case 0:
                                v3 -= v3
                            }
                        }
                    }
                    v2 = v2
                    _ = fi159
                }
                fmt.Println(v3)
            }
        }
        v166 := []int{9}
        _ = v166
    }
    var v167 float32 = 3.14159
    switch -89 {
    case -11:
        var v168 float32 = v1
        _ = v168
    case 127:
        var v169 bool = false
        _ = v169
    }
    if false {
        for fi170 := 0; fi170 < 19; fi170++ {
            var v171 bool = false
            v172 := []int{(fi170 + 32768), (fi170 % (((fi170 << ((805) & 63))) | 1))}
            v173 := GenS21{F0: 7, F1: 0, F2: 42}
            _ = fi170
            _ = v171
            _ = v172
            _ = v173
        }
    }
    var v174 string = "\n\t"
    for fi175 := 0; fi175 < 8; fi175++ {
        v176 := GenS11{F0: fi175, F1: (v174 + v25), F2: -1}
        v3 = 1e10
        switch (fi175 | fi175) {
        case -475:
            if (!false) {
                v174 = (v25 + (v25 + "\n\t"))
                if ((v1 - v1) != (v1 / float32(1e10))) {
                    var v177 uint = 0
                    v178 := []int{2147483647}
                    var v179 uint64 = 1
                    _ = v177
                    _ = v178
                    _ = v179
                } else {
                    if (!true) {
                        for fi180 := 0; fi180 < 15; fi180++ {
                            v181 := []int{(fi180 | fi180), fi175, (fi175 + genFn0(198, 100, fi180)), fi175}
                            _ = fi180
                            _ = v181
                        }
                    }
                    for fi182 := 0; fi182 < 14; fi182++ {
                        v25 = (v25 + v25)
                        fmt.Println(v27, fi182)
                        v1 = v167
                        _ = fi182
                    }
                    v183 := []int{(fi175 - 0), (fi175 + (-(fi175)))}
                    _ = v183
                }
                if (v3 > v3) {
                    v184 := GenS18{F0: 701, F1: v25, F2: 191}
                    v3 -= (v3 * v3)
                    var v185 float64 = 0.0
                    _ = v184
                    _ = v185
                } else {
                    v25 = (v25 + ("foo bar" + v25))
                }
            }
        case 128:
            if false {
                fmt.Println(v2)
                v25 = "a"
            } else {
                v186 := []int{32768, (fi175 - (fi175 - (^(fi175))))}
                v4 -= (v4 << ((18446744073709551615) & 63))
                _ = v186
            }
        case 380:
            var v187 uint64 = 18446744073709551615
            _ = v187
        default:
            if false {
                switch (("a" + "") + (v174 + v174)) {
                case "hello":
                    for fi188 := 0; fi188 < 13; fi188++ {
                        v2 -= (v2 >> ((18446744073709551615) & 63))
                        _ = fi188
                    }
                default:
                    for fi189 := 0; fi189 < 7; fi189++ {
                        var v190 uint64 = 0
                        fmt.Println(v167, v190, v190)
                        if ("" > "") {
                            switch v25 {
                            case "\n\t":
                                v191 := GenS22{F0: true}
                                _ = v191
                            case "foo bar":
                                v192 := GenS23{F0: 1, F1: (v2 / ((genFn1(true, v4, true)) | 1))}
                                _ = v192
                            case "":
                                v193 := []int{}
                                v194 := GenS24{F0: (v174 + ("foo bar" + (v174 + "hello"))), F1: 928}
                                _ = v193
                                _ = v194
                            }
                            var v195 float32 = -2.25
                            v25 = v25
                            _ = v195
                        } else {
                            var v196 string = v25
                            _ = v196
                        }
                        var v197 float64 = v3
                        _ = fi189
                        _ = v190
                        _ = v197
                    }
                }
            } else {
                v174 = "\n\t"
                if (true && false) {
                    switch "a" {
                    case "foo bar":
                        v198 := GenS9{F0: "foo bar", F1: 18446744073709551615}
                        _ = v198
                    }
                    v199 := []int{}
                    v199 = append(v199, genFn0((fi175 & (-(fi175))), -1, 127))
                    _ = v199
                }
                v200 := GenS25{F0: 42, F1: v27, F2: -752}
                _ = v200
            }
        }
        _ = fi175
        _ = v176
    }
    fmt.Println(v25)
    fmt.Println(v3)
    if false {
        for fi201 := 0; fi201 < 11; fi201++ {
            var v202 bool = (v3 <= (float64(0.1) - v3))
            v167 = 3.14159
            _ = fi201
            _ = v202
        }
        var v203 string = v174
        _ = v203
    }
    v204 := []int{}
    _ = len(v204)
    for fi205 := 0; fi205 < 13; fi205++ {
        v204 = append(v204, fi205)
        _ = fi205
    }
    for fi206 := 0; fi206 < 13; fi206++ {
        if false {
            v27 = v27
        } else {
            if len(v204) > 0 {
                v204[0] = (fi206 << ((255) & 63))
            }
            for fi207 := 0; fi207 < 15; fi207++ {
                if true {
                    v3 = v3
                    var v208 int64 = 647
                    _ = v208
                }
                _ = fi207
            }
        }
        fmt.Println(v3)
        var v209 int16 = genFn1(false, (v4 + v4), (((4294967295 > 277) || true) || ((0 <= 1) || true)))
        _ = fi206
        _ = v209
    }
    fmt.Println("v1:", v1)
    fmt.Println("v167:", v167)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v204:", len(v204), v204)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v167
    _ = v174
    _ = v2
    _ = v204
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
