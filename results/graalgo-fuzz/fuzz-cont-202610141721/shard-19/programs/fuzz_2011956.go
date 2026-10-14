package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0(p0 uint64, p1 int64) uint32 {
    p0 += (^(p0))
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 250, F1: -76, F2: (p1 / ((p1) | 1))}
        p1 = (p1 | (p1 & p1))
        _ = fi6
        _ = v7
    }
    if (211 >= 268) {
        p0 = 1
    }
    v8 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = v8
    return 259
}
func genFn1() int32 {
    var v9 uint8 = 2
    v9 += (^(v9))
    v9 = v9
    v9 += v9
    _ = v9
    return 924
}

func main() {
    var v1 int16 = -4
    v2 := []string{"", "", ("" + "a")}
    var v3 int16 = v1
    var v4 int8 = -128
    v5 := []uint32{42, 0, 4294967295}
    if (false || true) {
        var v10 float32 = (((3.14159 - 1e10) / 1.0) * 0.1)
        _ = v10
    } else {
        var v11 uint8 = 255
        _ = v11
    }
    v4 -= (v4 * v4)
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int8 = (v4 + (-(v4)))
        _ = fi12
        _ = v13
    }
    v5 = append(v5, 170)
    var v14 uint32 = 809
    var v15 int8 = int8(v1)
    v16 := GenS1{F0: -9223372036854775808, F1: -17, F2: 42}
    switch "" {
    case "foo bar":
        v17 := GenS1{F0: 127, F1: -128, F2: -536}
        _ = v17
    case "hello":
        if true {
            if true {
                var v18 uint = 507
                _ = v18
            } else {
                v14 = (v14 ^ (v14 | v14))
                v19 := GenS1{F0: -1, F1: v4, F2: -1}
                _ = v19
            }
        }
    }
    if len(v5) > 0 {
        v5[0] = genFn0(18446744073709551615, 1)
    }
    var v20 uint8 = 2
    if false {
        v15 = (^(v4))
    }
    if true {
        fmt.Println(v15, v15)
        switch 247 {
        case 4294967295:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        var v21 float64 = float64(3.14159)
        fmt.Println(v21)
        _ = v21
    }
    if true {
        switch (("a" + (("\n\t" + ("foo bar" + "hello")) + "a")) + ("foo bar" + "hello")) {
        case "":
            v22 := GenS2{F0: -2147483648, F1: 65535, F2: (v14 & 65535)}
            _ = v22
        case "foo bar":
            v2 = append(v2, "foo bar")
        case "hello":
            var v23 uint32 = v14
            _ = v23
        default:
            if (true && ((v1 % (((v1 << ((42) & 63))) | 1)) >= v3)) {
                v4 = 1
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 int16 = v1
                    var v26 uint64 = 0
                    v26 = 526
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v3 = v3
            } else {
                if ((v15 < 88) || true) {
                    switch -366 {
                    case 32768:
                        for fi27 := 0; fi27 < 7; fi27++ {
                            if ((v4 % (((v15 * (v4 & v4))) | 1)) != (-(v4))) {
                                v14 += (v14 | 42)
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    fmt.Println(v14, v20, v15)
                                    v29 := GenS1{F0: int(v3), F1: (v4 ^ 127), F2: 9223372036854775807}
                                    v2 = append(v2, "foo bar")
                                    if len(v5) > 0 {
                                        v5[0] = v14
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v30 int64 = -9223372036854775808
                                        _ = v30
                                    default:
                                        v4 += v4
                                    }
                                    v14 = v14
                                    _ = fi28
                                    _ = v29
                                }
                            } else {
                                var v31 string = ("foo bar" + "\n\t")
                                fmt.Println(v14, v4, v14)
                                _ = v31
                            }
                            if ((v20 ^ (v20 % ((128) | 1))) == (v20 / ((v20) | 1))) {
                                var v32 bool = false
                                v32 = v32
                                _ = v32
                            }
                            _ = fi27
                        }
                    case 2147483647:
                        var v33 uint16 = 255
                        _ = v33
                    }
                } else {
                    var v34 float32 = 3.14159
                    _ = v34
                }
                if false {
                    var v35 uint8 = 2
                    _ = v35
                }
                var v36 uint32 = (^(v14))
                _ = v36
            }
        }
    } else {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := GenS2{F0: (fi37 & 32767), F1: 1, F2: (v14 & 112)}
            var v39 int32 = -1
            var v40 float32 = 1e10
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    }
    var v41 uint = 79
    for fi42 := 0; fi42 < 17; fi42++ {
        var v43 float64 = (float64(0.1) + 1.5)
        if (!true) {
            for fi44 := 0; fi44 < 3; fi44++ {
                v45 := GenS2{F0: fi42, F1: (v41 & (^(v41))), F2: 1}
                if true {
                    v3 -= (v3 >> ((v41) & 63))
                    if len(v2) > 0 {
                        v2[0] = ("a" + "a")
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi44
                _ = v45
            }
            if false {
                for fi46 := 0; fi46 < 19; fi46++ {
                    fmt.Println(v14)
                    _ = fi46
                }
                for fi47 := 0; fi47 < 5; fi47++ {
                    if false {
                        _ = len(v5)
                    }
                    v14 += v14
                    _ = fi47
                }
            } else {
                if false {
                    for fi48 := 0; fi48 < 2; fi48++ {
                        switch "hello" {
                        case "":
                            if len(v5) > 0 {
                                v5[0] = 2147483647
                            }
                        case "foo bar":
                            var v49 float64 = 0.1
                            _ = v49
                        case "\n\t":
                            if false {
                                switch "\n\t" {
                                case "a":
                                    switch fi42 {
                                    case 128:
                                        v14 -= (v14 / ((v14) | 1))
                                    }
                                case "":
                                    v50 := GenS2{F0: (fi48 ^ fi48), F1: (v41 - (v41 & v41)), F2: 4294967295}
                                    _ = v50
                                case "foo bar":
                                    var v51 int64 = -9223372036854775808
                                    _ = v51
                                default:
                                    switch "" {
                                    case "hello":
                                        v52 := GenS1{F0: 1, F1: v15, F2: -88}
                                        _ = v52
                                    case "":
                                        var v53 int = 9223372036854775807
                                        _ = v53
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            if (2 >= 850) {
                                                switch ("a" + ("\n\t" + "hello")) {
                                                case "\n\t":
                                                    if (0 > v41) {
                                                        if false {
                                                            v54 := GenS2{F0: int(v14), F1: (v41 << ((v41) & 63)), F2: 65535}
                                                            _ = v54
                                                        }
                                                    } else {
                                                        v55 := GenS2{F0: int(v20), F1: (v41 & v41), F2: genFn0(18446744073709551615, 132)}
                                                        v20 += 255
                                                        if true {
                                                            var v56 bool = false
                                                            if true {
                                                                v57 := GenS3{F0: (fi48 % ((fi48) | 1)), F1: (fi48 * (fi42 ^ 2147483647))}
                                                                v58 := GenS4{F0: 0}
                                                                _ = v57
                                                                _ = v58
                                                            } else {
                                                                switch (fi42 | (^(fi48))) {
                                                                case 1:
                                                                    switch (fi42 ^ int(v15)) {
                                                                    case -1:
                                                                        if v56 {
                                                                            var v59 uint64 = 0
                                                                            if true {
                                                                                v43 += ((v43 - 1.5) * 3.14159)
                                                                                v59 = 1
                                                                                switch (fi48 / (((fi42 % ((-363) | 1))) | 1)) {
                                                                                case 128:
                                                                                    v60 := GenS5{F0: v3}
                                                                                    _ = v60
                                                                                }
                                                                            } else {
                                                                                var v61 uint64 = v59
                                                                                v62 := GenS5{F0: v3}
                                                                                v2 = append(v2, "hello")
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                            _ = v59
                                                                        }
                                                                    case 32767:
                                                                        v4 = (v15 % (((v4 << ((v41) & 63))) | 1))
                                                                    }
                                                                default:
                                                                    v63 := GenS3{F0: fi48, F1: (fi48 ^ (-(fi42)))}
                                                                    _ = v63
                                                                }
                                                                var v64 uint = 255
                                                                fmt.Println(v43, v4, v43)
                                                                switch fi42 {
                                                                case 0:
                                                                    if ((1 < v14) || (true && true)) {
                                                                        var v65 uint16 = 65535
                                                                        _ = v65
                                                                    } else {
                                                                        if v56 {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ("" + "a")
                                                                            }
                                                                            var v66 uint16 = 210
                                                                            v43 += -2.25
                                                                            _ = v66
                                                                        }
                                                                        var v67 int = (^(fi42))
                                                                        v68 := GenS6{F0: true, F1: -2147483648}
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                }
                                                                _ = v64
                                                            }
                                                            var v69 int8 = 100
                                                            if false {
                                                                if true {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if (fi48 <= fi48) {
                                                                        v70 := GenS3{F0: int(v1), F1: fi48}
                                                                        v20 += 1
                                                                        _ = v70
                                                                    }
                                                                    var v71 uint64 = 0
                                                                    _ = v71
                                                                }
                                                                v2 = append(v2, "hello")
                                                            }
                                                            _ = v56
                                                            _ = v69
                                                        } else {
                                                            var v72 int32 = int32(v3)
                                                            _ = v72
                                                        }
                                                        _ = v55
                                                    }
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v2 = append(v2, "")
                                            }
                                        case "hello":
                                            if false {
                                                v3 -= (v3 % ((827) | 1))
                                                if (0 > 475) {
                                                    switch fi42 {
                                                    case -726:
                                                        if (2 <= 42) {
                                                            if (2 != 1) {
                                                                var v73 float64 = (3.14159 * v43)
                                                                if false {
                                                                    if true {
                                                                        v43 += ((((1e10 - (v43 - 0.1)) - ((0.0 + v73) * (v73 * float64(1.5)))) * v73) / v73)
                                                                    } else {
                                                                        v3 += 32767
                                                                    }
                                                                    var v74 uint64 = 18446744073709551615
                                                                    var v75 uint8 = (v20 | v20)
                                                                    v41 = (v41 << (((v41 ^ (v41 << (((v41 | 820)) & 63)))) & 63))
                                                                    v73 += (3.14159 - ((float64(3.14159) + v43) / ((float64(float32((((v43 + v43) / ((((1e10 * (v73 + v73)) + ((float64(1.5) * 1e10) + (v73 - (((1e10 / (v73 + ((float64(1.5) * v73) + v73))) / 1.0) - v73)))) - -2.25) + v73)) + 1e10))) - (((float64(0.1) + v73) / (-2.25 - 0.1)) / (1e10 / v73))) * 0.0)))
                                                                    if (!((-2.25 - 3.14159) >= (3.14159 + 1e10))) {
                                                                        fmt.Println(v41)
                                                                        var v76 int8 = v4
                                                                        _ = v76
                                                                    } else {
                                                                        var v77 int16 = (^(v1))
                                                                        v3 = (v77 * (v3 | v1))
                                                                        _ = v77
                                                                    }
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                                v78 := GenS2{F0: fi48, F1: v41, F2: genFn0(2, 42)}
                                                                _ = v73
                                                                _ = v78
                                                            }
                                                            switch (fi48 - fi48) {
                                                            case 4294967295:
                                                                var v79 uint32 = (v14 + (v14 << ((uint(v14)) & 63)))
                                                                _ = v79
                                                            default:
                                                                v5 = append(v5, v14)
                                                            }
                                                            v80 := GenS3{F0: fi42, F1: (fi42 % (((fi42 ^ -931)) | 1))}
                                                            _ = v80
                                                        }
                                                    case -627:
                                                        v81 := GenS5{F0: -864}
                                                        _ = v81
                                                    }
                                                }
                                            } else {
                                                v1 += v1
                                                v1 = 787
                                                switch ("a" + "foo bar") {
                                                case "hello":
                                                    fmt.Println(v20, v43, v4)
                                                }
                                            }
                                        case "\n\t":
                                            var v82 int32 = 150
                                            var v83 uint = (v41 * v41)
                                            _ = v82
                                            _ = v83
                                        }
                                    }
                                }
                                var v84 string = (("a" + "hello") + ("a" + "\n\t"))
                                if true {
                                    var v85 int = fi48
                                    v43 -= float64((-2.25 * (3.14159 * 1e10)))
                                    var v86 int32 = -212
                                    _ = v85
                                    _ = v86
                                } else {
                                    v87 := GenS5{F0: -32768}
                                    _ = v87
                                }
                                _ = v84
                            } else {
                                if ("foo bar" <= (("\n\t" + ("\n\t" + "foo bar")) + ("" + "a"))) {
                                    v88 := GenS3{F0: (-(fi48)), F1: (fi48 ^ int(v41))}
                                    switch ((("" + ("foo bar" + "")) + ("foo bar" + "")) + ("" + "a")) {
                                    case "\n\t":
                                        _ = len(v5)
                                    default:
                                        fmt.Println(fi42, v43)
                                    }
                                    _ = v88
                                } else {
                                    var v89 uint64 = 448
                                    v89 += (v89 << ((v41) & 63))
                                    _ = v89
                                }
                                if true {
                                    fmt.Println(v15, v1, fi42)
                                } else {
                                    var v90 int32 = 42
                                    fmt.Println(v20, v14, fi42)
                                    _ = v90
                                }
                            }
                        default:
                            v4 = v4
                        }
                        v43 = ((v43 / v43) * float64(0.1))
                        _ = fi48
                    }
                    fmt.Println(v20, v15)
                }
            }
            v91 := GenS7{F0: (v3 | (v3 * v1)), F1: 2, F2: (^(v20))}
            _ = v91
        }
        fmt.Println(v20, v3, v43)
        _ = fi42
        _ = v43
    }
    v4 = 124
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
