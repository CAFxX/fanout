package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 int16; F2 int64 }
type GenS4 struct { F0 int64; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 string; F1 int; F2 int8 }
type GenS8 struct { F0 string; F1 bool }
func genFn0(p0 uint16, p1 string) uint {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint = 2
            fmt.Println(fi8, p1, v9)
            switch fi8 {
            case -9223372036854775808:
                var v10 int16 = 460
                _ = v10
            case 230:
                p0 += (^(p0))
            default:
                p1 = ("foo bar" + ((p1 + "hello") + p1))
            }
            if (!("foo bar" <= (p1 + p1))) {
                var v11 string = p1
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        var v12 float64 = float64(3.14159)
        _ = fi7
        _ = v12
    }
    var v13 float64 = 0.1
    v14 := GenS1{F0: 42}
    var v15 int32 = 62
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := []uint32{2147483647, 42}
    var v2 float32 = float32(((float64(0.1) / (0.1 - 1e10)) + (((0.1 / (((0.1 * float64(-2.25)) + 1e10) * ((0.0 / 1.5) - -2.25))) - -2.25) - 1.5)))
    var v3 int = 32768
    var v4 int8 = 106
    v5 := []int{(v3 >> ((42) & 63)), 4294967295, -9223372036854775808}
    var v6 float64 = (0.0 + (-2.25 / (3.14159 - ((((1e10 / -2.25) / (1.5 + float64(v2))) - ((0.1 * 1.5) - ((3.14159 + (1e10 - 1e10)) / 3.14159))) + 0.0))))
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint8 = 26
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 2147483648
    v19 := GenS2{F0: 42, F1: -1, F2: -583}
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS2{F0: 152, F1: 42, F2: 76}
            fmt.Println(fi21, fi21, v2)
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v18 += (v18 % (((v18 & (^(v18)))) | 1))
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v18, fi23)
        switch fi23 {
        case 42:
            var v24 int16 = 220
            _ = v24
        case 32768:
            v25 := GenS1{F0: -9223372036854775808}
            _ = v25
        case -581:
            _ = len(v1)
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if ((v18 << ((42) & 63)) != v18) {
            v5 = append(v5, -812)
            v2 += (v2 / (0.0 - v2))
        } else {
            switch fi26 {
            case 128:
                switch v3 {
                case 128:
                    if ((v18 | (v18 >> ((uint(v18)) & 63))) != (v18 % ((65535) | 1))) {
                        var v27 int32 = 589
                        _ = v27
                    }
                default:
                    if false {
                        var v28 int64 = -9223372036854775808
                        if true {
                            v2 = -2.25
                            var v29 float64 = 1e10
                            v30 := GenS3{F0: -1, F1: 42, F2: v28}
                            _ = v29
                            _ = v30
                        }
                        _ = v28
                    } else {
                        if false {
                            if (0.0 >= 1.5) {
                                v5 = append(v5, (^(v3)))
                                var v31 string = "hello"
                                _ = v31
                            } else {
                                v4 = v4
                                var v32 string = "a"
                                _ = v32
                            }
                            v18 = 311
                            var v33 int16 = -1
                            _ = v33
                        }
                    }
                }
            case 0:
                v18 -= 977
            default:
                v4 = (v4 & 0)
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v18 = v18
        }
        if ((0 == 32767) || false) {
            if false {
                v34 := GenS3{F0: 2147483647, F1: -1, F2: -9223372036854775808}
                switch "foo bar" {
                case "a":
                    var v35 float64 = v6
                    _ = v35
                case "":
                    v18 = (v18 & (v18 % ((v18) | 1)))
                case "hello":
                    switch (("a" + "\n\t") + "foo bar") {
                    case "a":
                        switch "a" {
                        case "foo bar":
                            if (!false) {
                                var v36 int8 = (v4 << ((339) & 63))
                                if (528 > 1000) {
                                    if (true || true) {
                                        if true {
                                            var v37 string = (("foo bar" + "foo bar") + "foo bar")
                                            v38 := GenS4{F0: 1, F1: -2147483648, F2: 128}
                                            _ = v37
                                            _ = v38
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = v18
                                            }
                                            if (623 >= uint64(v18)) {
                                                if (9223372036854775807 >= -925) {
                                                    switch ("a" + "") {
                                                    case "\n\t":
                                                        switch 32768 {
                                                        case 127:
                                                            fmt.Println(v36)
                                                        }
                                                    case "hello":
                                                        v1 = append(v1, (v18 ^ 769))
                                                    case "a":
                                                        if (true && true) {
                                                            var v39 int16 = 66
                                                            switch "" {
                                                            case "a":
                                                                if (827 > -824) {
                                                                    var v40 uint32 = v18
                                                                    _ = v40
                                                                } else {
                                                                    var v41 uint64 = 2
                                                                    _ = v41
                                                                }
                                                            case "":
                                                                if (42 != 1) {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    for fi42 := 0; fi42 < 19; fi42++ {
                                                                        var v43 uint16 = 0
                                                                        fmt.Println(v3, v18)
                                                                        var v44 uint8 = 255
                                                                        _ = fi42
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                } else {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    var v45 float64 = v6
                                                                    _ = v45
                                                                }
                                                            case "\n\t":
                                                                v6 = 3.14159
                                                            }
                                                            _ = v39
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            switch "hello" {
                                                            case "hello":
                                                                switch -716 {
                                                                case -183:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (^(v18))
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    for fi46 := 0; fi46 < 8; fi46++ {
                                                        if true {
                                                            var v47 uint32 = 2147483647
                                                            _ = v47
                                                        } else {
                                                            v1 = append(v1, (^(v18)))
                                                        }
                                                        _ = fi46
                                                    }
                                                    for fi48 := 0; fi48 < 9; fi48++ {
                                                        var v49 float32 = v2
                                                        fmt.Println(v2)
                                                        v4 = (v36 - v36)
                                                        _ = fi48
                                                        _ = v49
                                                    }
                                                } else {
                                                    fmt.Println(v4, v36, v4)
                                                }
                                                fmt.Println(v2, v18, v2)
                                                for fi50 := 0; fi50 < 7; fi50++ {
                                                    v3 += fi50
                                                    _ = len(v5)
                                                    _ = fi50
                                                }
                                            }
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if (255 > 20) {
                                            fmt.Println(v36)
                                            fmt.Println(v4, v3, v36)
                                            var v51 float32 = float32((v6 / float64(v2)))
                                            _ = v51
                                        }
                                    } else {
                                        var v52 uint32 = v18
                                        _ = v52
                                    }
                                }
                                v53 := GenS3{F0: -1, F1: 1000, F2: 9223372036854775807}
                                _ = v36
                                _ = v53
                            } else {
                                switch (v3 * (-(v3))) {
                                case 127:
                                    var v54 int8 = (v4 + v4)
                                    _ = v54
                                case 2147483647:
                                    var v55 string = "\n\t"
                                    _ = v55
                                case -840:
                                    var v56 int16 = 1000
                                    _ = v56
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v57 float64 = ((1.5 * float64(float32(v6))) / v6)
                                    _ = v57
                                }
                            }
                        case "hello":
                            fmt.Println(v2, v2)
                        case "":
                            var v58 float32 = 3.14159
                            _ = v58
                        }
                    }
                }
                _ = v34
            }
        } else {
            if (18446744073709551615 < 1) {
                fmt.Println(v3)
                var v59 int16 = 763
                if len(v5) > 0 {
                    v5[0] = (fi26 ^ fi26)
                }
                _ = v59
            }
            if false {
                v3 -= fi26
            }
        }
        v18 = (v18 | 0)
        _ = fi26
    }
    switch (v3 << ((301) & 63)) {
    case 128:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (!(!((3.14159 == 1.5) && true))) {
                if len(v1) > 0 {
                    v1[0] = v18
                }
                var v60 int16 = -417
                v61 := GenS4{F0: -9223372036854775808, F1: 1, F2: 190}
                _ = v60
                _ = v61
            }
        } else {
            fmt.Println(v6, v6, v2)
            v62 := GenS5{F0: 62, F1: -1, F2: 0}
            var v63 int64 = 0
            _ = v62
            _ = v63
        }
    }
    v5 = append(v5, (v3 ^ v3))
    _ = len(v1)
    for fi64 := 0; fi64 < 8; fi64++ {
        v18 -= 1
        switch (v3 / ((994) | 1)) {
        case -1:
            for fi65 := 0; fi65 < 9; fi65++ {
                var v66 string = ""
                var v67 uint16 = 2
                _ = fi65
                _ = v66
                _ = v67
            }
        case 36:
            if true {
                fmt.Println(v2)
                var v68 uint32 = 64
                if false {
                    for fi69 := 0; fi69 < 18; fi69++ {
                        v70 := GenS2{F0: -1, F1: 9223372036854775807, F2: -38}
                        var v71 int8 = (-(v4))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    v72 := GenS6{F0: 756, F1: 0}
                    _ = v72
                } else {
                    v18 = 2147483647
                }
                _ = v68
            } else {
                for fi73 := 0; fi73 < 18; fi73++ {
                    if (!(1e10 == v6)) {
                        v6 = 0.0
                    } else {
                        v74 := GenS7{F0: "\n\t", F1: (fi73 & int(v4)), F2: (-(v4))}
                        switch "hello" {
                        case "hello":
                            if ((fi73 & fi64) < -464) {
                                v18 += v18
                                switch -2147483648 {
                                case -135:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v2 = (v2 + 1.5)
                            } else {
                                if true {
                                    fmt.Println(v18, v3, fi64)
                                    v4 = (v4 ^ (v4 + 127))
                                    var v75 uint = 2
                                    _ = v75
                                }
                            }
                        case "a":
                            if (true || false) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case "\n\t":
                            switch (fi73 >> ((42) & 63)) {
                            case -9223372036854775808:
                                fmt.Println(v3)
                            case 768:
                                switch "a" {
                                case "":
                                    var v76 int16 = 1000
                                    _ = v76
                                case "hello":
                                    fmt.Println(v2, v18, v3)
                                case "a":
                                    var v77 int8 = v4
                                    _ = v77
                                default:
                                    v18 = (v18 * (v18 * (v18 ^ (v18 & v18))))
                                }
                            case 1:
                                v5 = append(v5, 373)
                            }
                        default:
                            switch fi73 {
                            case 42:
                                fmt.Println(v4, v4)
                            case 32767:
                                if true {
                                    var v78 uint64 = 0
                                    v79 := GenS3{F0: 0, F1: -556, F2: 1}
                                    if (((float32(0.0) + float32(1.5)) / v2) > v2) {
                                        if false {
                                            v2 = (v2 * (v2 * 0.1))
                                            switch int(v18) {
                                            case 42:
                                                var v80 int32 = 914
                                                _ = v80
                                            case -1:
                                                var v81 float64 = v6
                                                _ = v81
                                            case 2147483647:
                                                fmt.Println(v4)
                                            default:
                                                if false {
                                                    switch "hello" {
                                                    case "a":
                                                        var v82 bool = true
                                                        _ = v82
                                                    default:
                                                        v18 = (v18 << ((65535) & 63))
                                                    }
                                                }
                                            }
                                        } else {
                                            var v83 float64 = (-2.25 - v6)
                                            _ = v83
                                        }
                                        var v84 string = "foo bar"
                                        switch "foo bar" {
                                        case "":
                                            v85 := GenS5{F0: 2147483647, F1: 920, F2: 1}
                                            _ = v85
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = 32767
                                            }
                                        }
                                        v1 = append(v1, v18)
                                        _ = v84
                                    }
                                    _ = v78
                                    _ = v79
                                }
                            case 32768:
                                v86 := GenS8{F0: "\n\t", F1: false}
                                _ = v86
                            }
                        }
                        var v87 int8 = (^(v4))
                        _ = v74
                        _ = v87
                    }
                    var v88 bool = true
                    v4 -= (v4 - v4)
                    _ = fi73
                    _ = v88
                }
                var v89 string = ("\n\t" + ("hello" + ("hello" + ("\n\t" + "\n\t"))))
                _ = v89
            }
        default:
            v90 := GenS5{F0: (v18 % ((v18) | 1)), F1: -9223372036854775808, F2: 127}
            _ = v90
        }
        _ = fi64
    }
    var v91 bool = true
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
