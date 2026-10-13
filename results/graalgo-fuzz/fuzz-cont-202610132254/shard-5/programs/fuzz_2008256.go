package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int32; F2 int32 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS12 struct { F0 uint64; F1 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int16; F1 int8 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int; F1 int32 }
type GenS18 struct { F0 int32; F1 int64 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int64; F1 uint8 }
type GenS23 struct { F0 uint64; F1 int16 }
type GenS24 struct { F0 uint16 }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 uint16 }
func genFn0(p0 int32, p1 int64) uint8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 -= 42
        p0 -= p0
        _ = fi4
    }
    var v5 float64 = 1.5
    _ = p0
    _ = p1
    _ = v5
    return 48
}
func genFn1(p0 uint64) uint64 {
    var v6 uint64 = (^(p0))
    var v7 int = 961
    v8 := []int{(v7 | 0), (v7 / ((v7) | 1))}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v6 - p0)
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 42
    var v3 bool = true
    var v9 int32 = -2147483648
    if v3 {
        v2 += v2
        v9 = v9
        v10 := GenS1{F0: 2}
        _ = v10
    }
    v11 := GenS1{F0: 1}
    v12 := []int{-1}
    _ = len(v12)
    v13 := GenS2{F0: "hello", F1: 290}
    fmt.Println(v1, v3)
    v14 := GenS2{F0: ("" + ""), F1: 0}
    var v15 uint8 = 128
    var v16 uint8 = 2
    var v17 float64 = ((1e10 + float64((-2.25 / (((1.5 * (3.14159 + ((0.1 / 0.1) + 3.14159))) - 0.0) / 0.1)))) / -2.25)
    v18 := GenS1{F0: (^(v15))}
    if v3 {
        if (18446744073709551615 <= uint64(v2)) {
            fmt.Println(v9)
            if v3 {
                v19 := GenS3{F0: 42}
                var v20 uint64 = 2
                v21 := GenS4{F0: 1}
                _ = v19
                _ = v20
                _ = v21
            } else {
                if (1000 >= 0) {
                    if true {
                        fmt.Println(v9)
                        if (!v3) {
                            var v22 int32 = (v9 << ((1) & 63))
                            _ = v22
                        } else {
                            v23 := GenS5{F0: v2, F1: false}
                            if false {
                                fmt.Println(v15)
                            } else {
                                switch ("foo bar" + "") {
                                case "a":
                                    v16 = genFn0(1, -9223372036854775808)
                                case "":
                                    var v24 int8 = v1
                                    _ = v24
                                case "hello":
                                    switch 313 {
                                    case 4294967295:
                                        if v3 {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            var v25 float64 = 3.14159
                                            _ = v25
                                        }
                                    }
                                default:
                                    for fi26 := 0; fi26 < 7; fi26++ {
                                        var v27 string = ((("a" + "") + "foo bar") + ("foo bar" + ("a" + "hello")))
                                        var v28 uint64 = 167
                                        v12 = append(v12, fi26)
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            _ = v23
                        }
                    }
                    v15 += (v15 | genFn0(1, 9223372036854775807))
                    switch "hello" {
                    case "hello":
                        switch "\n\t" {
                        case "":
                            for fi29 := 0; fi29 < 6; fi29++ {
                                if true {
                                    for fi30 := 0; fi30 < 7; fi30++ {
                                        v9 += -84
                                        switch (fi30 - (fi29 - fi29)) {
                                        case 32767:
                                            v31 := GenS6{F0: (^(v2)), F1: 249}
                                            _ = v31
                                        default:
                                            v2 += (v2 / ((4294967295) | 1))
                                        }
                                        if true {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v32 := GenS7{F0: v1}
                                                _ = v32
                                            case "\n\t":
                                                if len(v12) > 0 {
                                                    v12[0] = 914
                                                }
                                            case "a":
                                                _ = len(v12)
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    if (2 != uint(v15)) {
                                                        var v33 bool = v3
                                                        fmt.Println(fi30)
                                                        _ = v33
                                                    } else {
                                                        switch ((("\n\t" + ("\n\t" + "\n\t")) + "\n\t") + "\n\t") {
                                                        case "a":
                                                            if len(v12) > 0 {
                                                                v12[0] = 9223372036854775807
                                                            }
                                                        case "foo bar":
                                                            v1 += -5
                                                        case "hello":
                                                            v9 = (v9 / ((2147483647) | 1))
                                                        default:
                                                            var v34 bool = true
                                                            _ = v34
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v15, fi29, fi30)
                                                }
                                            }
                                        } else {
                                            switch fi29 {
                                            case 89:
                                                v17 -= float64(1e10)
                                            case -863:
                                                v15 -= (v16 % ((128) | 1))
                                            default:
                                                if false {
                                                    v35 := GenS5{F0: (v2 - v2), F1: (127 >= 2147483647)}
                                                    if (!(v3 && true)) {
                                                        switch "a" {
                                                        case "":
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                fmt.Println(fi29, v2)
                                                            case "\n\t":
                                                                v36 := GenS5{F0: (v2 >> ((2) & 63)), F1: true}
                                                                _ = v36
                                                            }
                                                        case "foo bar":
                                                            switch -2147483648 {
                                                            case 447:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if v3 {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v37 uint8 = genFn0(1, 1)
                                                                            _ = v37
                                                                        case "foo bar":
                                                                            if (v1 != 0) {
                                                                                v38 := GenS4{F0: 926}
                                                                                switch fi29 {
                                                                                case 1:
                                                                                    fmt.Println(fi29)
                                                                                case -131:
                                                                                    v15 = v15
                                                                                case 2147483647:
                                                                                    var v39 float64 = (0.1 * (((v17 - -2.25) + v17) / v17))
                                                                                    _ = v39
                                                                                default:
                                                                                    var v40 string = "hello"
                                                                                    _ = v40
                                                                                }
                                                                                if false {
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        v12 = append(v12, -156)
                                                                                        if v3 {
                                                                                            switch 1 {
                                                                                            case -185:
                                                                                                v41 := GenS8{F0: (v15 << ((65535) & 63)), F1: v9, F2: (v2 ^ 634)}
                                                                                                _ = v41
                                                                                            }
                                                                                            v17 = (float64(1e10) + v17)
                                                                                            if len(v12) > 0 {
                                                                                                v12[0] = fi29
                                                                                            }
                                                                                        } else {
                                                                                            var v42 int32 = (v9 >> ((0) & 63))
                                                                                            if len(v12) > 0 {
                                                                                                v12[0] = (fi30 | fi30)
                                                                                            }
                                                                                            var v43 int16 = 1
                                                                                            _ = v42
                                                                                            _ = v43
                                                                                        }
                                                                                    case "hello":
                                                                                        fmt.Println(v17, v16)
                                                                                        if len(v12) > 0 {
                                                                                            v12[0] = fi29
                                                                                        }
                                                                                        var v44 int32 = -2147483648
                                                                                        _ = v44
                                                                                    case "\n\t":
                                                                                        v45 := GenS3{F0: 4294967295}
                                                                                        _ = v45
                                                                                    default:
                                                                                        var v46 int16 = -32768
                                                                                        _ = v46
                                                                                    }
                                                                                    v12 = append(v12, 32767)
                                                                                    var v47 float32 = 3.14159
                                                                                    _ = v47
                                                                                }
                                                                                _ = v38
                                                                            } else {
                                                                                if len(v12) >= 2 {
                                                                                    v12 = v12[1:len(v12)]
                                                                                }
                                                                                var v48 float64 = 1e10
                                                                                v3 = v3
                                                                                _ = v48
                                                                            }
                                                                        case "hello":
                                                                            v3 = (uint32(v15) <= (v2 & v2))
                                                                        default:
                                                                            if v3 {
                                                                                switch fi30 {
                                                                                case 424:
                                                                                    if ((("hello" + "a") + "foo bar") <= "hello") {
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            var v49 int32 = 1
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                if (((0.0 - float32(v17)) + -2.25) <= 1e10) {
                                                                                                    _ = len(v12)
                                                                                                    v16 = (v15 / (((v16 + uint8(v2))) | 1))
                                                                                                    if v3 {
                                                                                                        fmt.Println(v16, v16, v3)
                                                                                                    }
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v16 -= (^(v15))
                                                                                            default:
                                                                                                var v50 bool = v3
                                                                                                _ = v50
                                                                                            }
                                                                                            fmt.Println(v3)
                                                                                            var v51 float32 = -2.25
                                                                                            var v52 uint64 = 2
                                                                                            _ = v49
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                        default:
                                                                                            v16 = 255
                                                                                        }
                                                                                        if v3 {
                                                                                            var v53 float32 = float32(v17)
                                                                                            if (v1 == 0) {
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    v1 = v1
                                                                                                    v2 = v2
                                                                                                    var v54 uint64 = 42
                                                                                                    _ = v54
                                                                                                case "hello":
                                                                                                    if len(v12) > 0 {
                                                                                                        v12[0] = fi30
                                                                                                    }
                                                                                                case "a":
                                                                                                    v15 += (v16 | (^(v16)))
                                                                                                }
                                                                                            }
                                                                                            v55 := GenS8{F0: v15, F1: (v9 * (v9 ^ v9)), F2: v2}
                                                                                            _ = v53
                                                                                            _ = v55
                                                                                        } else {
                                                                                            v56 := GenS8{F0: 255, F1: -2147483648, F2: (v2 | (^(v2)))}
                                                                                            v57 := GenS5{F0: v2, F1: false}
                                                                                            var v58 uint = 1
                                                                                            _ = v56
                                                                                            _ = v57
                                                                                            _ = v58
                                                                                        }
                                                                                    } else {
                                                                                        v1 = (v1 - -1)
                                                                                        if true {
                                                                                            v59 := GenS2{F0: "\n\t", F1: 4294967295}
                                                                                            var v60 string = ("a" + (("hello" + "a") + "foo bar"))
                                                                                            if (false || (32767 != 698)) {
                                                                                                fmt.Println(v15, v1, v9)
                                                                                                if (!(v3 && (-458 != 9223372036854775807))) {
                                                                                                    if ((1 == -501) && ((v15 ^ (v16 >> ((953) & 63))) != (v15 % ((0) | 1)))) {
                                                                                                        v12 = append(v12, (fi29 - (fi30 - 2147483647)))
                                                                                                        v60 = "a"
                                                                                                    } else {
                                                                                                        v3 = (v3 || (v3 || v3))
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(fi29, v16)
                                                                                            }
                                                                                            _ = v59
                                                                                            _ = v60
                                                                                        } else {
                                                                                            var v61 bool = ((v2 + v2) == v2)
                                                                                            var v62 int16 = 0
                                                                                            v1 = (v1 & (v1 * v1))
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                        }
                                                                                        v12 = append(v12, (fi29 & (fi29 * 128)))
                                                                                        v15 = v16
                                                                                        v63 := GenS8{F0: (^(v16)), F1: v9, F2: 2147483647}
                                                                                        var v64 string = ("a" + "foo bar")
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    }
                                                                                }
                                                                                if len(v12) > 0 {
                                                                                    v12[0] = fi29
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch 127 {
                                                                        case 42:
                                                                            v16 = (v15 % (((v15 >> ((12) & 63))) | 1))
                                                                        case -2147483648:
                                                                            var v65 string = "foo bar"
                                                                            _ = v65
                                                                        default:
                                                                            var v66 float32 = 3.14159
                                                                            _ = v66
                                                                        }
                                                                        var v67 int = (fi29 & fi30)
                                                                        var v68 string = "a"
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                case "foo bar":
                                                                    if len(v12) >= 2 {
                                                                        v12 = v12[1:len(v12)]
                                                                    }
                                                                case "":
                                                                    switch ("hello" + "\n\t") {
                                                                    case "foo bar":
                                                                        v16 = (v16 | (v16 << ((33) & 63)))
                                                                    case "hello":
                                                                        var v69 uint8 = 0
                                                                        _ = v69
                                                                    case "\n\t":
                                                                        v9 -= (v9 % (((v9 % ((v9) | 1))) | 1))
                                                                    default:
                                                                        var v70 uint8 = v16
                                                                        _ = v70
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = len(v12)
                                                    }
                                                    fmt.Println(v2, v17)
                                                    _ = v35
                                                } else {
                                                    v71 := GenS2{F0: "hello", F1: 18446744073709551615}
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                    v17 -= ((v17 / 0.1) / (v17 - (v17 * (((3.14159 / v17) - -2.25) / ((v17 * v17) - v17)))))
                                                    _ = v71
                                                }
                                            }
                                        }
                                        _ = fi30
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = (^(fi29))
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = (fi29 ^ (fi29 % (((fi29 % ((fi29) | 1))) | 1)))
                                    }
                                }
                                for fi72 := 0; fi72 < 7; fi72++ {
                                    switch ("foo bar" + ("a" + "a")) {
                                    case "a":
                                        if v3 {
                                            var v73 uint16 = 832
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            var v74 uint16 = (^(v73))
                                            _ = v73
                                            _ = v74
                                        } else {
                                            if (v3 || v3) {
                                                v17 = (0.1 / 1e10)
                                                if (!v3) {
                                                    var v75 int64 = -1
                                                    if (float32(v17) >= 1e10) {
                                                        if ((fi72 % (((fi29 % (((fi72 >> ((1) & 63))) | 1))) | 1)) == fi29) {
                                                            if v3 {
                                                                if (!false) {
                                                                    v76 := GenS5{F0: (v2 % (((^(v2))) | 1)), F1: ((false && (v3 && (fi29 <= (fi29 / (((fi29 % (((fi29 & 526)) | 1))) | 1))))) || (!("a" != "")))}
                                                                    v2 = v2
                                                                    _ = v76
                                                                }
                                                                var v77 int64 = (v75 << ((255) & 63))
                                                                _ = v77
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    v15 = (v15 | (^(v16)))
                                                                case "":
                                                                    var v78 string = "\n\t"
                                                                    _ = v78
                                                                default:
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        var v79 int16 = 1000
                                                                        _ = v79
                                                                    case "":
                                                                        var v80 float32 = (1e10 - -2.25)
                                                                        _ = v80
                                                                    case "\n\t":
                                                                        switch ((("\n\t" + "hello") + "a") + ("hello" + "foo bar")) {
                                                                        case "":
                                                                            v1 = v1
                                                                        case "\n\t":
                                                                            if len(v12) > 0 {
                                                                                v12[0] = 938
                                                                            }
                                                                        case "a":
                                                                            v81 := GenS5{F0: v2, F1: ((v1 ^ v1) != v1)}
                                                                            _ = v81
                                                                        default:
                                                                            var v82 uint16 = 255
                                                                            _ = v82
                                                                        }
                                                                    default:
                                                                        v16 -= 255
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            if v3 {
                                                                v3 = v3
                                                                var v83 int32 = (v9 | v9)
                                                                _ = v83
                                                            }
                                                            _ = len(v12)
                                                            v84 := GenS9{F0: 763, F1: v9, F2: (^(v9))}
                                                            _ = v84
                                                        }
                                                        var v85 uint16 = 32767
                                                        if ((!false) || (0.1 >= 1.5)) {
                                                            _ = len(v12)
                                                            var v86 float64 = v17
                                                            _ = v86
                                                        }
                                                        _ = v85
                                                    }
                                                    _ = v75
                                                } else {
                                                    v87 := GenS10{F0: 757, F1: 1}
                                                    var v88 uint32 = 42
                                                    _ = v87
                                                    _ = v88
                                                }
                                                if v3 {
                                                    v16 -= v16
                                                    v12 = append(v12, int(v2))
                                                    _ = len(v12)
                                                } else {
                                                    fmt.Println(v16)
                                                    v89 := GenS11{F0: (v2 - v2), F1: true, F2: v15}
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v90 uint16 = 42
                                                        fmt.Println(v2, fi29)
                                                        _ = v90
                                                    case "":
                                                        var v91 int64 = 551
                                                        _ = v91
                                                    case "hello":
                                                        if len(v12) > 0 {
                                                            v12[0] = 42
                                                        }
                                                    }
                                                    _ = v89
                                                }
                                            } else {
                                                v15 = 1
                                                fmt.Println(v17, v2)
                                            }
                                            var v92 int32 = (v9 ^ -799)
                                            switch 1 {
                                            case 42:
                                                var v93 string = "\n\t"
                                                _ = v93
                                            case 1:
                                                v2 = 817
                                            case 128:
                                                if false {
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                    v17 -= 3.14159
                                                }
                                            default:
                                                switch "\n\t" {
                                                case "a":
                                                    if v3 {
                                                        v15 -= uint8(v2)
                                                    } else {
                                                        v94 := GenS12{F0: 42, F1: (((v2 - v2) > v2) && false)}
                                                        if (v9 < v9) {
                                                            if len(v12) > 0 {
                                                                v12[0] = (fi29 & int(v9))
                                                            }
                                                            v15 -= v16
                                                        } else {
                                                            v95 := GenS4{F0: 18446744073709551615}
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    }
                                                case "hello":
                                                    v96 := GenS13{F0: 18446744073709551615}
                                                    _ = v96
                                                default:
                                                    var v97 int8 = 127
                                                    _ = v97
                                                }
                                            }
                                            _ = v92
                                        }
                                    case "":
                                        if true {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            if false {
                                                var v98 uint = 653
                                                _ = v98
                                            }
                                            v9 = v9
                                        } else {
                                            if v3 {
                                                fmt.Println(v17, fi72, v17)
                                                if v3 {
                                                    v3 = v3
                                                } else {
                                                    switch fi29 {
                                                    case -9223372036854775808:
                                                        if (2147483647 > (v2 | 913)) {
                                                            if true {
                                                                var v99 int = (^(fi72))
                                                                if (42 == 32767) {
                                                                    if len(v12) > 0 {
                                                                        v12[0] = 128
                                                                    }
                                                                    v17 = 1.5
                                                                    if ((false && (!v3)) && (319 <= 1)) {
                                                                        _ = len(v12)
                                                                        v100 := GenS14{F0: 1}
                                                                        v17 += 0.0
                                                                        fmt.Println(fi72, v2)
                                                                        _ = v100
                                                                    } else {
                                                                        var v101 uint16 = 188
                                                                        _ = v101
                                                                    }
                                                                }
                                                                fmt.Println(v15)
                                                                _ = len(v12)
                                                                _ = v99
                                                            } else {
                                                                if false {
                                                                    if true {
                                                                        v17 -= v17
                                                                        var v102 uint16 = 32767
                                                                        if v3 {
                                                                            v15 += (v15 / ((v15) | 1))
                                                                            var v103 int8 = (v1 * v1)
                                                                            v104 := GenS1{F0: 127}
                                                                            v3 = (genFn0((^(v9)), 42) >= (v15 >> ((401) & 63)))
                                                                            if v3 {
                                                                                v15 -= v16
                                                                                v105 := GenS10{F0: v102, F1: 1}
                                                                                v102 = uint16(v16)
                                                                                v103 = (v1 ^ 52)
                                                                                _ = v105
                                                                            }
                                                                            _ = v103
                                                                            _ = v104
                                                                        }
                                                                        _ = v102
                                                                    } else {
                                                                        var v106 uint = 42
                                                                        fmt.Println(v3, v15, v106)
                                                                        var v107 uint = (v106 >> ((v106) & 63))
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                    v108 := GenS15{F0: 42, F1: v1}
                                                                    _ = v108
                                                                }
                                                                v12 = append(v12, (^(fi72)))
                                                            }
                                                        } else {
                                                            fmt.Println(fi29, fi29, fi29)
                                                        }
                                                    case -702:
                                                        _ = len(v12)
                                                    }
                                                }
                                            } else {
                                                v109 := GenS16{F0: (v1 / ((v1) | 1))}
                                                var v110 uint8 = v16
                                                v111 := GenS17{F0: -9223372036854775808, F1: -507}
                                                _ = v109
                                                _ = v110
                                                _ = v111
                                            }
                                            v12 = append(v12, (fi72 / ((fi29) | 1)))
                                            var v112 float32 = 1e10
                                            _ = v112
                                        }
                                    case "\n\t":
                                        v2 += 42
                                    default:
                                        v17 += v17
                                    }
                                    if false {
                                        v113 := GenS18{F0: (v9 ^ 42), F1: 843}
                                        v16 = genFn0(v9, 959)
                                        _ = v113
                                    } else {
                                        switch ("a" + "") {
                                        case "":
                                            var v114 uint8 = v15
                                            _ = v114
                                        case "a":
                                            switch (("foo bar" + "") + "\n\t") {
                                            case "hello":
                                                v17 = (0.1 + 3.14159)
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    switch fi72 {
                                                    case -9223372036854775808:
                                                        if ((0 == 4294967295) && false) {
                                                            var v115 int32 = 948
                                                            v116 := GenS13{F0: 18446744073709551615}
                                                            _ = v115
                                                            _ = v116
                                                        } else {
                                                            var v117 bool = (32767 >= -1)
                                                            v117 = v117
                                                            var v118 uint32 = 855
                                                            if (((float64(3.14159) + v17) >= 0.1) || true) {
                                                                fmt.Println(v3, v16, v1)
                                                                if true {
                                                                    if (true || (-9223372036854775808 >= int64(v2))) {
                                                                        v119 := GenS19{F0: (v9 / ((1) | 1))}
                                                                        var v120 int8 = (v1 % ((v1) | 1))
                                                                        var v121 int16 = 371
                                                                        _ = v119
                                                                        _ = v120
                                                                        _ = v121
                                                                    }
                                                                    v118 = (v2 >> ((255) & 63))
                                                                    v15 -= (^(v16))
                                                                } else {
                                                                    if v117 {
                                                                        v117 = ((^(v9)) < -2147483648)
                                                                        v12 = append(v12, fi29)
                                                                    } else {
                                                                        if len(v12) >= 2 {
                                                                            v12 = v12[1:len(v12)]
                                                                        }
                                                                        if v3 {
                                                                            var v122 int8 = (v1 + (-(v1)))
                                                                            if v117 {
                                                                                if ("foo bar" < "\n\t") {
                                                                                    v123 := GenS3{F0: 42}
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v124 float64 = (0.1 + v17)
                                                                                        _ = v124
                                                                                    case "hello":
                                                                                        v125 := GenS20{F0: -571}
                                                                                        _ = v125
                                                                                    }
                                                                                    _ = v123
                                                                                }
                                                                                var v126 uint = 381
                                                                                _ = v126
                                                                            }
                                                                            var v127 bool = (false || ((!("\n\t" > "\n\t")) || v3))
                                                                            _ = v122
                                                                            _ = v127
                                                                        }
                                                                        if false {
                                                                            v3 = (v3 && ((32767 < 0) && ((^(v15)) <= (^(v15)))))
                                                                        }
                                                                    }
                                                                    if len(v12) >= 2 {
                                                                        v12 = v12[1:len(v12)]
                                                                    }
                                                                    var v128 uint = 0
                                                                    if (v3 || v3) {
                                                                        v129 := GenS19{F0: (^(v9))}
                                                                        var v130 uint16 = 42
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    v128 = v128
                                                                    _ = v128
                                                                }
                                                            } else {
                                                                if v3 {
                                                                    v16 += (v15 >> ((255) & 63))
                                                                    v131 := GenS21{F0: 0}
                                                                    if v117 {
                                                                        var v132 uint64 = 18446744073709551615
                                                                        v12 = append(v12, (fi29 >> ((485) & 63)))
                                                                        v9 = v9
                                                                        v16 += v16
                                                                        _ = v132
                                                                    } else {
                                                                        fmt.Println(v118, v117, v1)
                                                                    }
                                                                    _ = v131
                                                                } else {
                                                                    var v133 float32 = ((((-2.25 * float32(v17)) * 0.0) + (-2.25 - 3.14159)) / float32((((1.5 + (((v17 * float64(-2.25)) - 1e10) - ((v17 * v17) - (v17 + 0.0)))) - ((1e10 / v17) * 1.5)) * -2.25)))
                                                                    var v134 int32 = v9
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                                if len(v12) > 0 {
                                                                    v12[0] = fi72
                                                                }
                                                                var v135 float32 = 0.0
                                                                _ = v135
                                                            }
                                                            v136 := GenS14{F0: v1}
                                                            if (-360 >= 0) {
                                                                if false {
                                                                    v2 = (v2 >> ((2) & 63))
                                                                    if true {
                                                                        if len(v12) >= 2 {
                                                                            v12 = v12[1:len(v12)]
                                                                        }
                                                                        if ("\n\t" >= "a") {
                                                                            v12 = append(v12, (fi72 & (fi72 ^ (fi29 | (fi29 - fi72)))))
                                                                            if ((!true) || (v3 && v3)) {
                                                                                var v137 uint64 = genFn1(4294967295)
                                                                                v138 := GenS22{F0: 9223372036854775807, F1: (v16 / ((v15) | 1))}
                                                                                _ = v137
                                                                                _ = v138
                                                                            } else {
                                                                                var v139 uint8 = v15
                                                                                if (v9 == v9) {
                                                                                    fmt.Println(v9)
                                                                                    v117 = v3
                                                                                } else {
                                                                                    var v140 string = "foo bar"
                                                                                    v141 := GenS3{F0: uint(v16)}
                                                                                    v117 = ((18446744073709551615 != 1) && (("foo bar" + ((v140 + v140) + (v140 + v140))) <= v140))
                                                                                    v142 := GenS23{F0: 2, F1: 193}
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                    _ = v142
                                                                                }
                                                                                _ = v139
                                                                            }
                                                                            v143 := GenS14{F0: (v1 >> ((760) & 63))}
                                                                            _ = v143
                                                                        }
                                                                    }
                                                                    v144 := GenS24{F0: 195}
                                                                    _ = v144
                                                                }
                                                                v145 := GenS9{F0: 2, F1: v9, F2: v9}
                                                                _ = v145
                                                            }
                                                            _ = v117
                                                            _ = v118
                                                            _ = v136
                                                        }
                                                    default:
                                                        var v146 uint64 = 2
                                                        _ = v146
                                                    }
                                                case "":
                                                    v16 -= v16
                                                case "foo bar":
                                                    v3 = ((("hello" + (("foo bar" + ("" + "a")) + "a")) + "a") < "hello")
                                                default:
                                                    if v3 {
                                                        if (!true) {
                                                            _ = len(v12)
                                                        }
                                                        fmt.Println(v16, fi72, fi29)
                                                    } else {
                                                        fmt.Println(fi72)
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            v15 = (v16 / ((27) | 1))
                                        default:
                                            if (1e10 != 0.1) {
                                                v3 = v3
                                                switch fi72 {
                                                case 0:
                                                    v1 += -128
                                                }
                                            }
                                        }
                                        v15 = v16
                                        switch (fi29 ^ fi29) {
                                        case 9223372036854775807:
                                            if false {
                                                var v147 string = "\n\t"
                                                v17 = v17
                                                v148 := GenS25{F0: 1}
                                                _ = v147
                                                _ = v148
                                            }
                                        }
                                        v12 = append(v12, int(v9))
                                        fmt.Println(v17)
                                        fmt.Println(fi29, v1, v1)
                                        fmt.Println(fi29, v1, v1)
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                    switch 32767 {
                                    case 127:
                                        if ((v3 && v3) && v3) {
                                            v149 := GenS10{F0: uint16(v2), F1: 707}
                                            fmt.Println(v15)
                                            v15 = (v16 - (v15 & v16))
                                            var v150 uint16 = 939
                                            _ = v149
                                            _ = v150
                                        }
                                    case -96:
                                        v2 -= 42
                                    case 42:
                                        fmt.Println(v15, v15)
                                    default:
                                        v1 = v1
                                    }
                                    _ = fi72
                                }
                                v151 := GenS26{F0: 0}
                                _ = fi29
                                _ = v151
                            }
                        default:
                            v2 = v2
                        }
                    default:
                        if (v3 || v3) {
                            for fi152 := 0; fi152 < 3; fi152++ {
                                v153 := GenS8{F0: 128, F1: v9, F2: (^(v2))}
                                if (true || (true && (!(!false)))) {
                                    fmt.Println(v1, v2, v15)
                                    var v154 int32 = -2147483648
                                    _ = v154
                                }
                                _ = fi152
                                _ = v153
                            }
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                    }
                } else {
                    var v155 int16 = -224
                    if (!(!((((v17 - v17) / -2.25) * float64((0.1 - 0.0))) != v17))) {
                        fmt.Println(v9, v155, v9)
                        v156 := GenS20{F0: -465}
                        _ = v156
                    }
                    _ = v155
                }
            }
            v157 := GenS17{F0: -1, F1: (v9 - v9)}
            _ = v157
        } else {
            var v158 float32 = (1.5 / float32(v17))
            _ = v158
        }
        for fi159 := 0; fi159 < 17; fi159++ {
            if (583 != 42) {
                _ = len(v12)
            } else {
                for fi160 := 0; fi160 < 7; fi160++ {
                    if false {
                        fmt.Println(v1, fi159)
                    }
                    _ = fi160
                }
                switch fi159 {
                case 32767:
                    v16 = v16
                case 0:
                    fmt.Println(fi159, v1, fi159)
                case 728:
                    v9 -= v9
                }
            }
            v2 = 1
            _ = fi159
        }
        v16 = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
