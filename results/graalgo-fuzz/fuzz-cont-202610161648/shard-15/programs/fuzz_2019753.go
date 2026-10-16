package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 string; F2 string }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint; F2 uint }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 int; F2 uint16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 int16) uint8 {
    var v4 uint64 = 0
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 17; fi6++ {
        if (-128 <= 120) {
            var v7 int16 = p0
            var v8 uint8 = 127
            if true {
                var v9 float32 = 1.5
                var v10 int64 = 42
                _ = v10
                _ = v9
            }
            _ = v7
            _ = v8
        } else {
            fmt.Println(p0, fi6, v5)
            var v11 uint64 = 1
            _ = v11
        }
        switch fi6 {
        case 32767:
            var v12 uint16 = 2
            _ = v12
        }
        _ = fi6
    }
    v13 := []int{}
    _ = p0
    _ = v13
    _ = v4
    _ = v5
    return 213
}

func main() {
    v1 := []int{-631, 329, -1}
    var v2 uint = 65535
    v3 := []string{"foo bar", "a", "\n\t", "hello"}
    for fi14 := 0; fi14 < 6; fi14++ {
        v2 = 2
        var v15 uint64 = 541
        _ = fi14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 1; fi16++ {
            if (!(2147483648 <= 2147483647)) {
                var v17 int64 = -9223372036854775808
                for fi18 := 0; fi18 < 7; fi18++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v17 = (v17 | (-(v17)))
                    _ = fi18
                }
                _ = v17
            } else {
                fmt.Println(fi16, fi16)
                if false {
                    var v19 uint = v2
                    _ = v19
                }
            }
            _ = fi16
        }
        v2 = (v2 ^ v2)
    } else {
        switch "a" {
        case "hello":
            var v20 uint = 2
            _ = v20
        case "":
            for fi21 := 0; fi21 < 14; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 float64 = 1e10
                    v24 := GenS1{F0: 18446744073709551615, F1: (fi22 << (((v2 >> ((1) & 63))) & 63))}
                    var v25 int64 = -784
                    switch (-(fi22)) {
                    case -2147483648:
                        if false {
                            v25 += (v25 << ((18446744073709551615) & 63))
                            _ = len(v3)
                            var v26 uint64 = 504
                            _ = v26
                        } else {
                            var v27 string = "hello"
                            var v28 uint32 = 4294967295
                            v1 = append(v1, 9223372036854775807)
                            _ = v27
                            _ = v28
                        }
                    }
                    var v29 float32 = 1e10
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v29
                }
                for fi30 := 0; fi30 < 5; fi30++ {
                    if (((((1.5 / 3.14159) - 3.14159) - 1.5) == 1.5) && false) {
                        _ = len(v3)
                        if ((0.0 < (3.14159 - float64(1.5))) && (!true)) {
                            v2 -= (v2 & 0)
                            if false {
                                var v31 int16 = 107
                                _ = v31
                            } else {
                                fmt.Println(fi21, fi30)
                                v2 = 42
                            }
                        } else {
                            fmt.Println(v2, v2)
                            v2 = v2
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v2 = 1
                                switch "a" {
                                case "foo bar":
                                    if (!true) {
                                        v2 = v2
                                    }
                                case "":
                                    fmt.Println(fi21, fi21)
                                default:
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v32 int16 = 42
                                            _ = v32
                                        case "a":
                                            v2 = (v2 % ((607) | 1))
                                        case "hello":
                                            fmt.Println(fi30, v2, v2)
                                        default:
                                            switch (fi21 << ((781) & 63)) {
                                            case -334:
                                                v2 -= v2
                                            case 0:
                                                if false {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "\n\t":
                                                        v2 = (v2 & v2)
                                                    case "hello":
                                                        var v33 bool = (!(-1 > 194))
                                                        _ = v33
                                                    }
                                                    var v34 int = fi30
                                                    var v35 int32 = -823
                                                    fmt.Println(v34, fi21)
                                                    _ = v34
                                                    _ = v35
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v36 int8 = -97
                                                    _ = v36
                                                default:
                                                    if (((0.1 + 0.0) - (-2.25 + 3.14159)) <= 0.0) {
                                                        v2 += v2
                                                        v37 := GenS1{F0: 18446744073709551615, F1: (-(fi21))}
                                                        _ = v37
                                                    } else {
                                                        if (0 > 0) {
                                                            var v38 uint32 = 4294967295
                                                            _ = v38
                                                        }
                                                    }
                                                }
                                            case -1:
                                                var v39 uint = 425
                                                _ = v39
                                            default:
                                                switch "a" {
                                                case "":
                                                    v2 = (^(v2))
                                                case "\n\t":
                                                    v40 := GenS1{F0: v2, F1: 127}
                                                    _ = v40
                                                }
                                            }
                                        }
                                        fmt.Println(fi30)
                                    }
                                }
                                if (!((v2 ^ 255) != (v2 - 65535))) {
                                    _ = len(v1)
                                    if true {
                                        if (-1 == -32768) {
                                            switch "hello" {
                                            case "":
                                                v2 += 4294967295
                                            case "a":
                                                if len(v1) > 0 {
                                                    v1[0] = fi21
                                                }
                                            case "foo bar":
                                                v3 = append(v3, "")
                                            default:
                                                v2 -= (v2 | (^(v2)))
                                            }
                                        }
                                        if (((0.0 / 3.14159) - 1.5) >= 0.1) {
                                            if (("hello" + "a") == "hello") {
                                                switch (fi30 - 4294967295) {
                                                case 842:
                                                    switch fi21 {
                                                    case 4294967295:
                                                        v1 = append(v1, (fi30 - fi30))
                                                    case 128:
                                                        switch "a" {
                                                        case "hello":
                                                            if ("\n\t" > "hello") {
                                                                v2 = (v2 / ((v2) | 1))
                                                            }
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = fi21
                                                            }
                                                        }
                                                    }
                                                case -2147483648:
                                                    var v41 uint8 = 68
                                                    _ = v41
                                                case -504:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                _ = len(v1)
                                                switch "\n\t" {
                                                case "hello":
                                                    switch ("" + ("\n\t" + "a")) {
                                                    case "":
                                                        fmt.Println(fi30, fi30)
                                                        v42 := GenS1{F0: (^(v2)), F1: 127}
                                                        _ = v42
                                                    case "\n\t":
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                    default:
                                                        v43 := GenS2{F0: -12, F1: 42}
                                                        _ = v43
                                                    }
                                                    v44 := GenS2{F0: 1, F1: -9223372036854775808}
                                                    _ = v44
                                                }
                                            }
                                            switch "\n\t" {
                                            case "hello":
                                                v2 = 18446744073709551615
                                            case "foo bar":
                                                if (fi21 < fi30) {
                                                    if (uint64(v2) > 619) {
                                                        v2 = v2
                                                    } else {
                                                        if (!(1.5 < 0.1)) {
                                                            if false {
                                                                v2 += (v2 >> (((^(v2))) & 63))
                                                                var v45 int16 = -899
                                                                var v46 int = (fi21 - fi30)
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        } else {
                                                            v47 := GenS2{F0: -1, F1: 9223372036854775807}
                                                            _ = v47
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v48 := GenS3{F0: (fi21 ^ fi30), F1: "foo bar", F2: ""}
                                                        _ = v48
                                                    }
                                                    var v49 int16 = -829
                                                    _ = v49
                                                } else {
                                                    var v50 uint64 = 4294967295
                                                    v51 := GenS1{F0: v2, F1: 127}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v52 uint32 = 671
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi30, v2, v2)
                                        var v53 float32 = ((float32(float64(1.5)) - 1e10) * (1e10 * 1e10))
                                        _ = v53
                                    }
                                }
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = (fi21 / ((193) | 1))
                                }
                                v2 += (v2 % ((v2) | 1))
                                var v54 float64 = -2.25
                                v55 := GenS4{F0: (v2 + (v2 % ((5) | 1))), F1: 740}
                                _ = v54
                                _ = v55
                            }
                        }
                    }
                    v2 -= 18446744073709551615
                    if false {
                        v56 := GenS1{F0: 124, F1: (fi30 >> ((v2) & 63))}
                        fmt.Println(v2, v2, v2)
                        _ = v56
                    } else {
                        if false {
                            _ = len(v3)
                            if (((0.0 > 3.14159) && true) || false) {
                                var v57 int64 = 0
                                _ = v57
                            }
                            v2 += (^(v2))
                            var v58 uint8 = 255
                            _ = v58
                        } else {
                            fmt.Println(fi30, fi30)
                        }
                        var v59 uint = v2
                        _ = v59
                    }
                    _ = fi30
                }
                v2 = 255
                _ = fi21
            }
        }
        var v60 int8 = 127
        _ = v60
    }
    v2 -= (v2 >> ((v2) & 63))
    v61 := GenS2{F0: 111, F1: 0}
    v62 := GenS3{F0: -9223372036854775808, F1: "hello", F2: ("foo bar" + "foo bar")}
    v2 += (v2 - v2)
    if false {
        _ = len(v3)
    } else {
        v2 -= v2
        switch "foo bar" {
        case "hello":
            v1 = append(v1, 128)
        case "\n\t":
            var v63 uint = (v2 % (((v2 * (v2 | v2))) | 1))
            _ = v63
        case "":
            v64 := GenS5{F0: 433, F1: v2, F2: v2}
            _ = v64
        default:
            if true {
                for fi65 := 0; fi65 < 3; fi65++ {
                    v66 := GenS5{F0: -975, F1: (v2 % ((v2) | 1)), F2: v2}
                    _ = fi65
                    _ = v66
                }
            } else {
                v2 += (v2 - v2)
            }
        }
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        if (false && false) {
            if false {
                v2 -= v2
            } else {
                for fi67 := 0; fi67 < 4; fi67++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if ((2 <= genFn0(353)) && ("foo bar" >= "a")) {
                        v2 = v2
                        for fi68 := 0; fi68 < 18; fi68++ {
                            v69 := GenS6{F0: 0, F1: 255}
                            _ = fi68
                            _ = v69
                        }
                    } else {
                        var v70 int = (fi67 + fi67)
                        v71 := GenS7{F0: 42, F1: v2}
                        _ = v70
                        _ = v71
                    }
                    v2 -= v2
                    _ = fi67
                }
                v2 += (v2 ^ 874)
            }
            if (!true) {
                if false {
                    fmt.Println(v2)
                }
                for fi72 := 0; fi72 < 10; fi72++ {
                    fmt.Println(v2, fi72)
                    _ = fi72
                }
            } else {
                fmt.Println(v2)
            }
            v2 = v2
        } else {
            if true {
                v2 = v2
                for fi73 := 0; fi73 < 13; fi73++ {
                    v74 := GenS8{F0: "hello", F1: (true || true)}
                    switch (("a" + "") + ("hello" + "")) {
                    case "\n\t":
                        fmt.Println(fi73, fi73)
                    case "foo bar":
                        v2 -= (v2 | 42)
                    case "":
                        v75 := GenS4{F0: (^(v2)), F1: 1}
                        _ = v75
                    default:
                        if (2 < 18446744073709551615) {
                            for fi76 := 0; fi76 < 12; fi76++ {
                                var v77 float64 = -2.25
                                var v78 bool = (fi73 != (fi76 ^ fi73))
                                v79 := GenS9{F0: 1}
                                _ = fi76
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                            v2 -= 65535
                            v80 := GenS10{F0: -970, F1: fi73, F2: 255}
                            _ = v80
                        }
                    }
                    _ = fi73
                    _ = v74
                }
            }
        }
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        fmt.Println(v2)
    default:
        if (!false) {
            v81 := GenS11{F0: -32768}
            _ = v81
        }
    }
    var v82 int16 = 32767
    for fi83 := 0; fi83 < 15; fi83++ {
        if (-1 < -704) {
            v84 := GenS4{F0: (v2 * 4294967295), F1: 2}
            _ = v84
        } else {
            v82 -= 0
            for fi85 := 0; fi85 < 11; fi85++ {
                var v86 uint8 = 114
                _ = fi85
                _ = v86
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi87 := 0; fi87 < 14; fi87++ {
            if true {
                v82 -= -987
                _ = len(v3)
                if (false && (false && false)) {
                    v88 := GenS7{F0: 1, F1: 18446744073709551615}
                    var v89 uint32 = 65535
                    fmt.Println(v89, fi83, v82)
                    _ = v88
                    _ = v89
                } else {
                    var v90 uint = (v2 * 42)
                    fmt.Println(v2)
                    switch int(v82) {
                    case 9223372036854775807:
                        if true {
                            v91 := GenS7{F0: 2147483648, F1: v90}
                            v82 += v82
                            _ = v91
                        } else {
                            if false {
                                v92 := GenS11{F0: 32767}
                                _ = v92
                            }
                        }
                    case 42:
                        v93 := GenS3{F0: 4294967295, F1: "\n\t", F2: "hello"}
                        _ = v93
                    default:
                        var v94 string = "\n\t"
                        _ = v94
                    }
                    v95 := GenS12{F0: 32767, F1: 491, F2: 369}
                    fmt.Println(fi87, fi83)
                    v96 := GenS8{F0: ("hello" + ""), F1: true}
                    _ = v90
                    _ = v95
                    _ = v96
                }
            } else {
                var v97 int32 = 42
                v98 := GenS10{F0: 42, F1: fi83, F2: 2}
                _ = v97
                _ = v98
            }
            if false {
                fmt.Println(fi87)
            }
            _ = fi87
        }
        _ = fi83
    }
    v2 -= v2
    var v99 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v82:", v82)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v82
    _ = v99
}
