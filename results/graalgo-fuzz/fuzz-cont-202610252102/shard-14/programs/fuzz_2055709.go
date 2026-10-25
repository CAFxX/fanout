package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 int64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint32; F1 int32; F2 int64 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 int32; F1 int64 }
type GenS12 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS13 struct { F0 uint8; F1 string; F2 uint16 }
type GenS14 struct { F0 string; F1 string; F2 uint8 }
type GenS15 struct { F0 int; F1 int16 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 bool; F1 int16; F2 int8 }
func genFn0(p0 uint32) uint64 {
    v7 := GenS1{F0: 1, F1: 983, F2: 42}
    p0 = p0
    var v8 int16 = 0
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := []int{fi9, fi9, (fi9 | 128), fi9}
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint = 4294967295
            fmt.Println(v12)
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1() int {
    var v13 int32 = 617
    v13 = v13
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint64 = 1
        v16 := []int{-62, (fi14 | (fi14 | int(v13))), (fi14 >> ((255) & 63)), 226}
        var v17 int64 = -1
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    _ = v13
    return -884
}

func main() {
    v1 := []string{"foo bar", "\n\t", (("foo bar" + "") + ("foo bar" + ("foo bar" + "hello")))}
    var v2 int8 = -62
    v3 := []string{}
    v4 := []int{}
    var v5 uint64 = 326
    var v6 int = -9223372036854775808
    v5 = v5
    var v18 int16 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v19 := GenS2{F0: v6, F1: v18}
    v2 -= v2
    switch "hello" {
    case "foo bar":
        v2 += v2
    case "a":
        var v20 int16 = v18
        _ = v20
    case "hello":
        if false {
            v21 := GenS2{F0: -2147483648, F1: (v18 / ((v18) | 1))}
            _ = v21
        }
    }
    fmt.Println(v18)
    v22 := GenS1{F0: genFn0(9), F1: v18, F2: v18}
    if (-2.25 == (1e10 - 0.1)) {
        var v23 float32 = float32(((1.5 / (float64(float32(3.14159)) / 1.0)) - (1.5 * 0.0)))
        _ = v23
    }
    v24 := GenS3{F0: -750}
    var v25 uint64 = 1
    v26 := GenS1{F0: v25, F1: 32767, F2: v18}
    var v27 float32 = 3.14159
    v6 = (-(v6))
    fmt.Println(v6, v25, v27)
    if (((1.5 - (-2.25 / 1e10)) / 1.5) > 3.14159) {
        v25 -= (^(v5))
        if (1 > v2) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS3{F0: 720}
                for fi30 := 0; fi30 < 1; fi30++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        var v31 float64 = (3.14159 + -2.25)
                        fmt.Println(v18)
                        _ = v31
                    } else {
                        v25 += (v5 * (v5 * (v5 ^ (v25 >> ((754) & 63)))))
                        v32 := GenS4{F0: -1, F1: -1, F2: -9223372036854775808}
                        _ = v32
                    }
                    var v33 uint = uint(v25)
                    v6 = (v6 * genFn1())
                    _ = fi30
                    _ = v33
                }
                if ((244 >= 42) || (1e10 >= ((-2.25 / (float64((3.14159 / v27)) * -2.25)) + -2.25))) {
                    for fi34 := 0; fi34 < 14; fi34++ {
                        if (-2147483648 <= (fi28 >> ((uint(v5)) & 63))) {
                            fmt.Println(fi34)
                            v25 = (v5 % ((v5) | 1))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v35 := GenS3{F0: -105}
                            _ = v35
                        }
                        v1 = append(v1, "")
                        if true {
                            var v36 int16 = (v18 / ((v18) | 1))
                            switch "foo bar" {
                            case "a":
                                if (1e10 < (((-2.25 / 3.14159) * 1.5) * float64((1e10 * 1e10)))) {
                                    if ((false && (42 > 128)) && false) {
                                        switch (^(fi34)) {
                                        case 32767:
                                            _ = len(v3)
                                        case 32768:
                                            if ("foo bar" != ("" + ("a" + "foo bar"))) {
                                                switch ("\n\t" + "foo bar") {
                                                case "a":
                                                    if true {
                                                        if (65535 == 65535) {
                                                            fmt.Println(fi34, v2)
                                                            if false {
                                                                if true {
                                                                    var v37 uint32 = 4294967295
                                                                    v38 := GenS3{F0: -1}
                                                                    _ = v37
                                                                    _ = v38
                                                                } else {
                                                                    v6 += (^(fi28))
                                                                    v4 = append(v4, fi34)
                                                                }
                                                                if (v27 != 1.5) {
                                                                    fmt.Println(v6)
                                                                    var v39 float64 = 1.5
                                                                    var v40 uint16 = 616
                                                                    _ = v39
                                                                    _ = v40
                                                                } else {
                                                                    v2 += 42
                                                                    v41 := GenS1{F0: (v25 & (v5 / (((^(v25))) | 1))), F1: (^(v18)), F2: (v18 * v18)}
                                                                    v18 -= (v36 & -690)
                                                                    _ = v41
                                                                }
                                                                v36 = 549
                                                                if len(v1) > 0 {
                                                                    v1[0] = "hello"
                                                                }
                                                                v27 += v27
                                                            } else {
                                                                v42 := GenS3{F0: 42}
                                                                var v43 float32 = v27
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                            _ = len(v3)
                                                        } else {
                                                            switch v6 {
                                                            case 128:
                                                                var v44 uint64 = (v5 * v5)
                                                                _ = v44
                                                            case 32767:
                                                                if (!true) {
                                                                    switch fi34 {
                                                                    case 2147483647:
                                                                        v6 = 227
                                                                    default:
                                                                        var v45 uint = 18446744073709551615
                                                                        _ = v45
                                                                    }
                                                                } else {
                                                                    v27 = 1.5
                                                                    switch "hello" {
                                                                    case "":
                                                                        if (true && (false || false)) {
                                                                            v46 := GenS1{F0: 42, F1: (v18 | 0), F2: (-(v18))}
                                                                            v47 := GenS3{F0: 42}
                                                                            var v48 int64 = -581
                                                                            _ = v46
                                                                            _ = v47
                                                                            _ = v48
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v49 := GenS4{F0: 42, F1: 42, F2: int64(v2)}
                                                                                _ = v49
                                                                            case "hello":
                                                                                switch "" {
                                                                                case "":
                                                                                    v1 = append(v1, "hello")
                                                                                case "hello":
                                                                                    if true {
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            v6 += 0
                                                                                        case "\n\t":
                                                                                            var v50 uint16 = 32767
                                                                                            _ = v50
                                                                                        default:
                                                                                            var v51 uint = 531
                                                                                            _ = v51
                                                                                        }
                                                                                        v52 := GenS5{F0: (v36 >> ((739) & 63)), F1: v2}
                                                                                        v25 += (v25 << ((1) & 63))
                                                                                        v27 += 1e10
                                                                                        if (false || (!true)) {
                                                                                            _ = len(v3)
                                                                                        }
                                                                                        _ = v52
                                                                                    } else {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "\n\t"
                                                                                        }
                                                                                        var v53 int8 = v2
                                                                                        _ = v53
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if (!(435 > 1)) {
                                                                                    var v54 uint = 4294967295
                                                                                    _ = len(v4)
                                                                                    v2 = (v2 * (v2 * (v2 - v2)))
                                                                                    _ = v54
                                                                                }
                                                                            }
                                                                            var v55 int64 = 42
                                                                            _ = v55
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi34, fi28, fi34)
                                                                    }
                                                                }
                                                                v27 += v27
                                                                v18 -= v36
                                                            case -580:
                                                                v5 -= v25
                                                            }
                                                        }
                                                        v56 := GenS6{F0: 1, F1: v2, F2: 2}
                                                        fmt.Println(v25, fi28, v5)
                                                        v27 -= 0.1
                                                        var v57 uint = 636
                                                        var v58 uint = v57
                                                        fmt.Println(fi28, v2)
                                                        var v59 int16 = int16(v6)
                                                        if true {
                                                            _ = len(v4)
                                                            v27 += 0.0
                                                        } else {
                                                            v4 = append(v4, fi34)
                                                            var v60 int16 = int16(v2)
                                                            if (!((352 <= 42) && true)) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = "\n\t"
                                                                }
                                                                var v61 int16 = (-(v36))
                                                                fmt.Println(v2, fi34, v27)
                                                                _ = v61
                                                            } else {
                                                                v5 = (v25 << ((v57) & 63))
                                                                _ = len(v1)
                                                            }
                                                            _ = v60
                                                        }
                                                        if (true || false) {
                                                            if (v57 >= v57) {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if true {
                                                                        var v62 int32 = -108
                                                                        v4 = append(v4, (^(fi28)))
                                                                        fmt.Println(v18, v58, v57)
                                                                        var v63 uint8 = uint8(v58)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        if ((v18 << ((v58) & 63)) == v18) {
                                                                            v64 := GenS6{F0: 32767, F1: (v2 * v2), F2: 1}
                                                                            _ = v64
                                                                        } else {
                                                                            if ((v2 << ((uint(v25)) & 63)) != v2) {
                                                                                v4 = append(v4, -867)
                                                                            } else {
                                                                                v62 -= (v62 / ((v62) | 1))
                                                                                fmt.Println(v63, fi34)
                                                                            }
                                                                        }
                                                                        _ = v62
                                                                        _ = v63
                                                                    } else {
                                                                        switch 128 {
                                                                        case -2147483648:
                                                                            switch v6 {
                                                                            case 128:
                                                                                var v65 int16 = (v59 | (v59 & 474))
                                                                                _ = v65
                                                                            case 42:
                                                                                fmt.Println(v18, v25, v57)
                                                                                switch fi34 {
                                                                                case -1:
                                                                                    var v66 uint16 = 2
                                                                                    _ = v66
                                                                                case -174:
                                                                                    v27 = v27
                                                                                }
                                                                                fmt.Println(v5, v58, v36)
                                                                            default:
                                                                                v57 = (v58 | (v57 / ((v58) | 1)))
                                                                                v27 = (v27 / 1.5)
                                                                                var v67 int = 2147483647
                                                                                fmt.Println(v18, v27)
                                                                                _ = v67
                                                                            }
                                                                        default:
                                                                            var v68 int8 = 42
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    fmt.Println(v27, fi34)
                                                                case "foo bar":
                                                                    fmt.Println(v5, v27)
                                                                default:
                                                                    switch (v6 >> ((v58) & 63)) {
                                                                    case 0:
                                                                        var v69 uint8 = 217
                                                                        _ = v69
                                                                    case 9223372036854775807:
                                                                        if (3.14159 <= 1e10) {
                                                                            v59 += (v18 >> ((v57) & 63))
                                                                            var v70 uint16 = 1
                                                                            _ = v70
                                                                        }
                                                                    case -509:
                                                                        fmt.Println(v25, fi34)
                                                                    default:
                                                                        fmt.Println(fi34, v2, v25)
                                                                    }
                                                                }
                                                                var v71 uint16 = 255
                                                                fmt.Println(v25, v27)
                                                                _ = v71
                                                            }
                                                            v6 = -178
                                                        }
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                case "\n\t":
                                                    v72 := GenS7{F0: 635, F1: v2}
                                                    _ = v72
                                                }
                                                fmt.Println(v27, v27)
                                                switch "hello" {
                                                case "foo bar":
                                                    var v73 float64 = -2.25
                                                    _ = v73
                                                default:
                                                    v74 := GenS5{F0: (v36 >> ((1) & 63)), F1: v2}
                                                    _ = v74
                                                }
                                                v75 := GenS1{F0: 18446744073709551615, F1: (v36 - 0), F2: (v18 << ((18446744073709551615) & 63))}
                                                _ = v75
                                            }
                                        default:
                                            v2 -= v2
                                        }
                                        fmt.Println(v18, v6)
                                        if (1e10 > (1e10 * (-2.25 * 1e10))) {
                                            if (false || ((0.0 - -2.25) < (float64(v27) / 0.1))) {
                                                v76 := GenS8{F0: uint32(v5), F1: -1, F2: -1}
                                                v6 = genFn1()
                                                _ = v76
                                            } else {
                                                v2 += v2
                                                var v77 float64 = (1e10 + 0.1)
                                                v18 += v36
                                                _ = v77
                                            }
                                            if false {
                                                if false {
                                                    v78 := GenS9{F0: ((("hello" + "a") + ("" + "a")) + ("\n\t" + ("a" + ("foo bar" + (("hello" + ("\n\t" + (("hello" + "a") + ("" + "a")))) + ("foo bar" + ((("hello" + "foo bar") + "") + "a"))))))), F1: (!((1.5 > 0.1) && true))}
                                                    v36 += v36
                                                    _ = v78
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        } else {
                                            fmt.Println(v5)
                                            if (v27 == -2.25) {
                                                v2 = int8(fi28)
                                                v79 := GenS10{F0: -347, F1: 65535}
                                                fmt.Println(v27, v27)
                                                _ = v79
                                            } else {
                                                v80 := GenS9{F0: (("" + (("a" + ((((("hello" + (("" + "\n\t") + "")) + "a") + ("" + "\n\t")) + ("\n\t" + (("hello" + "foo bar") + ("" + "\n\t")))) + "hello")) + "")) + "hello"), F1: true}
                                                if (!false) {
                                                    if true {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            if true {
                                                                if false {
                                                                    if (false || true) {
                                                                        v25 -= 4294967295
                                                                        v2 += (^(v2))
                                                                        v1 = append(v1, "")
                                                                    } else {
                                                                        var v81 int16 = (v18 << ((1) & 63))
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v6
                                                                        }
                                                                        _ = v81
                                                                    }
                                                                    v18 = (v36 >> ((18446744073709551615) & 63))
                                                                } else {
                                                                    var v82 int8 = -97
                                                                    fmt.Println(v6, v5)
                                                                    var v83 uint = 42
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                                if true {
                                                                    fmt.Println(fi28, v18, fi34)
                                                                    if (true || (1 > -2147483648)) {
                                                                        fmt.Println(v27)
                                                                    } else {
                                                                        var v84 int32 = 2147483647
                                                                        v85 := GenS5{F0: v18, F1: v2}
                                                                        var v86 string = "foo bar"
                                                                        _ = v84
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    var v87 bool = false
                                                                    _ = v87
                                                                } else {
                                                                    var v88 uint32 = 1
                                                                    fmt.Println(fi34)
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v89 := GenS3{F0: 536}
                                                                        if (1 >= 0) {
                                                                            _ = len(v1)
                                                                            if false {
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = genFn1()
                                                                                }
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 2147483647
                                                                                }
                                                                                v1 = append(v1, ("a" + "foo bar"))
                                                                            }
                                                                            v90 := GenS11{F0: -2147483648, F1: -9223372036854775808}
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                fmt.Println(v88, v6, fi28)
                                                                            default:
                                                                                v36 -= (v36 | (v36 >> ((0) & 63)))
                                                                            }
                                                                            _ = v90
                                                                        } else {
                                                                            v25 += genFn0((v88 % ((2147483648) | 1)))
                                                                        }
                                                                        _ = v89
                                                                    case "a":
                                                                        if (!true) {
                                                                            var v91 uint64 = v5
                                                                            _ = v91
                                                                        } else {
                                                                            v92 := GenS7{F0: -1, F1: int8(fi34)}
                                                                            var v93 float32 = v27
                                                                            if (fi28 == fi34) {
                                                                                v88 -= (v88 - (v88 + (^(v88))))
                                                                            } else {
                                                                                if true {
                                                                                    v36 += (v36 >> ((0) & 63))
                                                                                }
                                                                                v88 -= (v88 - 4294967295)
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (-(v6))
                                                                            }
                                                                            var v94 uint = 721
                                                                            _ = v92
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                    case "":
                                                                        _ = len(v1)
                                                                        if (128 != 2) {
                                                                            var v95 int = fi28
                                                                            _ = v95
                                                                        } else {
                                                                            v5 = genFn0(v88)
                                                                        }
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            switch (v6 >> ((18446744073709551615) & 63)) {
                                                                            case 988:
                                                                                v96 := GenS6{F0: 290, F1: (v2 % ((v2) | 1)), F2: 32767}
                                                                                _ = v96
                                                                            case 0:
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    _ = len(v1)
                                                                                case "":
                                                                                    if false {
                                                                                        fmt.Println(v6)
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(v36, v2)
                                                                                        case "hello":
                                                                                            if false {
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = "a"
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            if true {
                                                                                                v4 = append(v4, genFn1())
                                                                                                var v97 int8 = (v2 + (v2 ^ v2))
                                                                                                v5 += (^(v25))
                                                                                                _ = v97
                                                                                            }
                                                                                        }
                                                                                        var v98 uint32 = v88
                                                                                        v88 = 1
                                                                                        _ = v98
                                                                                    }
                                                                                case "hello":
                                                                                    v36 = v18
                                                                                default:
                                                                                    var v99 uint8 = 127
                                                                                    _ = v99
                                                                                }
                                                                            case 975:
                                                                                v100 := GenS6{F0: 619, F1: -128, F2: 32767}
                                                                                _ = v100
                                                                            }
                                                                        case "a":
                                                                            fmt.Println(fi28, v36)
                                                                        }
                                                                    }
                                                                    _ = v88
                                                                }
                                                                v101 := GenS9{F0: ("\n\t" + "a"), F1: true}
                                                                fmt.Println(v6, v18, v6)
                                                                var v102 int16 = -587
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v101
                                                                _ = v102
                                                            } else {
                                                                var v103 float64 = 0.1
                                                                if (("\n\t" + (((("a" + ("\n\t" + "")) + ("\n\t" + ("\n\t" + ("\n\t" + ("a" + ("a" + ("" + "hello"))))))) + "a") + "a")) != "") {
                                                                    switch fi34 {
                                                                    case 32768:
                                                                        switch (fi34 / ((fi28) | 1)) {
                                                                        case 42:
                                                                            if true {
                                                                                v18 -= v18
                                                                                var v104 float32 = v27
                                                                                _ = v104
                                                                            }
                                                                        case -9223372036854775808:
                                                                            v27 = (float32(0.0) / (v27 - 0.0))
                                                                        }
                                                                        var v105 uint8 = 127
                                                                        var v106 string = "foo bar"
                                                                        _ = v105
                                                                        _ = v106
                                                                    case -82:
                                                                        v27 -= 0.0
                                                                        if (-554 <= -32768) {
                                                                            v6 = (v6 & int(v36))
                                                                            v1 = append(v1, ((("foo bar" + ("a" + "foo bar")) + "hello") + "a"))
                                                                            switch fi28 {
                                                                            case 0:
                                                                                if (42 <= -1) {
                                                                                    fmt.Println(v18, v25, v18)
                                                                                    var v107 uint = 42
                                                                                    _ = v107
                                                                                }
                                                                            case -685:
                                                                                v108 := GenS12{F0: 0, F1: 1, F2: 1}
                                                                                _ = v108
                                                                            case 32767:
                                                                                var v109 float64 = v103
                                                                                _ = v109
                                                                            default:
                                                                                v6 += fi28
                                                                            }
                                                                        }
                                                                        v103 = 0.1
                                                                        if false {
                                                                            var v110 bool = (!(!true))
                                                                            if (v110 || (97 >= 2)) {
                                                                                fmt.Println(v6, v27)
                                                                                var v111 int32 = -2147483648
                                                                                fmt.Println(v2, v5)
                                                                                if v110 {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = "foo bar"
                                                                                    }
                                                                                    if v110 {
                                                                                        v25 -= v25
                                                                                    } else {
                                                                                        var v112 uint64 = 1
                                                                                        v113 := GenS9{F0: "", F1: true}
                                                                                        var v114 float64 = v103
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    }
                                                                                    v3 = append(v3, "")
                                                                                }
                                                                                v36 -= v36
                                                                                var v115 float32 = v27
                                                                                fmt.Println(v103, v25, v36)
                                                                                _ = v111
                                                                                _ = v115
                                                                            }
                                                                            v5 += v25
                                                                            _ = v110
                                                                        } else {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                switch genFn1() {
                                                                                case 660:
                                                                                    v116 := GenS12{F0: 65535, F1: 255, F2: -1}
                                                                                    _ = v116
                                                                                case -2147483648:
                                                                                    v25 += (v25 * 248)
                                                                                }
                                                                            case "foo bar":
                                                                                if true {
                                                                                    var v117 float32 = 0.0
                                                                                    var v118 int64 = -9223372036854775808
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                }
                                                                                v25 = genFn0(1)
                                                                                var v119 uint8 = 108
                                                                                _ = v119
                                                                            }
                                                                            if (0 < uint16(v5)) {
                                                                                switch (fi28 - (v6 << ((18446744073709551615) & 63))) {
                                                                                case -1:
                                                                                    var v120 int32 = 0
                                                                                    _ = v120
                                                                                case 4294967295:
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "foo bar"
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v121 := GenS13{F0: 90, F1: ("" + "foo bar"), F2: 32767}
                                                                                        _ = v121
                                                                                    case "hello":
                                                                                        v122 := GenS14{F0: "hello", F1: "foo bar", F2: 255}
                                                                                        _ = v122
                                                                                    default:
                                                                                        v123 := GenS15{F0: genFn1(), F1: -32768}
                                                                                        _ = v123
                                                                                    }
                                                                                case 840:
                                                                                    v27 -= (v27 * v27)
                                                                                }
                                                                                switch genFn1() {
                                                                                case 1:
                                                                                    v6 -= 32767
                                                                                default:
                                                                                    if true {
                                                                                        v124 := GenS16{F0: ""}
                                                                                        var v125 bool = true
                                                                                        var v126 int16 = 0
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                        _ = v126
                                                                                    }
                                                                                    if (((true && ((0.0 < 1.5) || (42 == 32767))) && true) && false) {
                                                                                        v127 := GenS8{F0: 1, F1: -2147483648, F2: 1}
                                                                                        if (v103 >= 0.0) {
                                                                                            if true {
                                                                                                if false {
                                                                                                    v27 -= 0.0
                                                                                                } else {
                                                                                                    if (910 != 863) {
                                                                                                        var v128 string = (("a" + "hello") + ("" + ""))
                                                                                                        v18 = 1
                                                                                                        _ = v128
                                                                                                    }
                                                                                                    var v129 uint64 = (v5 + (v5 << ((65535) & 63)))
                                                                                                    v6 = -580
                                                                                                    _ = v129
                                                                                                }
                                                                                                if false {
                                                                                                    v130 := GenS17{F0: true, F1: v36, F2: -128}
                                                                                                    fmt.Println(v27, fi28, v6)
                                                                                                    _ = v130
                                                                                                }
                                                                                            } else {
                                                                                                v36 = 730
                                                                                                var v131 int64 = 42
                                                                                                _ = v131
                                                                                            }
                                                                                        } else {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        }
                                                                                        _ = v127
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (("\n\t" + ("hello" + "hello")) == ("" + "foo bar")) {
                                                                            v25 = (v5 % (((v5 % ((v5) | 1))) | 1))
                                                                        } else {
                                                                            v18 = (v18 | v36)
                                                                        }
                                                                    }
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v103
                                                            }
                                                        }
                                                        fmt.Println(v6)
                                                        v132 := GenS6{F0: 65535, F1: (-(v2)), F2: 65535}
                                                        _ = v132
                                                    }
                                                }
                                                var v133 uint64 = (v5 * v25)
                                                v2 += v2
                                                _ = v133
                                                _ = v80
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi28, v2, fi28)
                                    }
                                    if true {
                                        _ = len(v3)
                                    }
                                    v2 = (v2 ^ (v2 - -23))
                                }
                            }
                            _ = v36
                        }
                        _ = fi34
                    }
                    if (127 <= 127) {
                        if false {
                            v134 := GenS9{F0: ((("\n\t" + "\n\t") + "foo bar") + (("a" + "hello") + (("\n\t" + "foo bar") + ""))), F1: true}
                            _ = v134
                        } else {
                            var v135 int = genFn1()
                            _ = v135
                        }
                        if ((!false) && (2 != 250)) {
                            if (4294967295 == (-(v6))) {
                                v136 := GenS16{F0: ""}
                                var v137 uint64 = 515
                                fmt.Println(v2)
                                _ = v136
                                _ = v137
                            } else {
                                for fi138 := 0; fi138 < 19; fi138++ {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v25, v27)
                                    case "hello":
                                        v139 := GenS12{F0: 2147483648, F1: 560, F2: 0}
                                        _ = len(v4)
                                        v18 += (v18 / (((^(v18))) | 1))
                                        var v140 int = fi138
                                        v25 = v25
                                        if true {
                                            _ = len(v1)
                                            var v141 int = 128
                                            _ = v141
                                        }
                                        switch 127 {
                                        case 128:
                                            v3 = append(v3, "\n\t")
                                        case 962:
                                            v6 += v6
                                        case 0:
                                            var v142 int64 = 141
                                            _ = v142
                                        default:
                                            v6 = genFn1()
                                        }
                                        _ = v139
                                        _ = v140
                                    }
                                    var v143 float32 = 1e10
                                    var v144 int = (-(v6))
                                    _ = fi138
                                    _ = v143
                                    _ = v144
                                }
                                if ((("foo bar" + "hello") == ("a" + ("\n\t" + ""))) || false) {
                                    v25 += 18446744073709551615
                                } else {
                                    var v145 string = ("\n\t" + "\n\t")
                                    _ = v145
                                }
                            }
                        } else {
                            var v146 int8 = (v2 >> ((65535) & 63))
                            _ = v146
                        }
                        _ = len(v4)
                    } else {
                        v2 = v2
                        if ((((-2.25 / v27) + -2.25) - v27) != v27) {
                            if ((!false) && true) {
                                var v147 float32 = ((float32(float64(((-2.25 * float32((3.14159 - 3.14159))) * (v27 / (float32((3.14159 - (3.14159 / (0.1 + -2.25)))) + v27))))) / (float32(0.0) * (v27 * v27))) + v27)
                                var v148 uint64 = genFn0(614)
                                v149 := GenS1{F0: v5, F1: v18, F2: (v18 * 1000)}
                                _ = v147
                                _ = v148
                                _ = v149
                            } else {
                                var v150 uint = 65535
                                fmt.Println(fi28)
                                _ = v150
                            }
                        }
                        v18 -= v18
                    }
                } else {
                    for fi151 := 0; fi151 < 4; fi151++ {
                        var v152 int16 = (v18 << ((805) & 63))
                        _ = fi151
                        _ = v152
                    }
                }
                _ = fi28
                _ = v29
            }
            if (42 == 185) {
                if ((!("a" >= "foo bar")) || (255 != 194)) {
                    for fi153 := 0; fi153 < 14; fi153++ {
                        for fi154 := 0; fi154 < 13; fi154++ {
                            var v155 int8 = (v2 << ((65535) & 63))
                            _ = fi154
                            _ = v155
                        }
                        v2 -= v2
                        var v156 float32 = 0.1
                        _ = fi153
                        _ = v156
                    }
                    var v157 uint16 = 740
                    _ = v157
                }
                var v158 uint = 0
                fmt.Println(v25)
                _ = v158
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
