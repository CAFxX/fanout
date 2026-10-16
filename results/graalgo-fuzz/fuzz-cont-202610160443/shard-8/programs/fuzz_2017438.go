package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 string }
type GenS6 struct { F0 int8; F1 string; F2 int }
type GenS7 struct { F0 bool; F1 uint32; F2 int32 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 uint; F2 bool }
func genFn0() uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 0, F1: 0, F2: (!false)}
        var v6 int16 = 42
        v6 = v6
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 100, F2: (42 <= 68)}
    _ = v7
    return 239
}
func genFn1(p0 string, p1 int64) int8 {
    if true {
        var v8 float32 = -2.25
        _ = v8
    }
    var v9 int64 = (p1 << ((18446744073709551615) & 63))
    p1 = (v9 + v9)
    v10 := []int{0, -1}
    for fi11 := 0; fi11 < 17; fi11++ {
        if (!(true && true)) {
            if len(v10) > 0 {
                v10[0] = (fi11 & (fi11 * fi11))
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 100
}

func main() {
    var v1 bool = false
    var v2 uint = 1
    var v3 int64 = -187
    v1 = v1
    if (0 != 255) {
        var v12 int64 = v3
        v13 := []int{734, 751, -27}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                switch fi15 {
                case 127:
                    v2 -= 255
                case 762:
                    v16 := []int{(fi14 * fi15), (^(fi15))}
                    _ = v16
                case -9223372036854775808:
                    var v17 uint64 = 42
                    _ = v17
                default:
                    var v18 bool = false
                    _ = v18
                }
                v3 -= (v3 ^ v3)
                _ = fi15
            }
            v19 := []int{}
            switch "" {
            case "\n\t":
                switch fi14 {
                case 32767:
                    v3 -= v3
                case -2147483648:
                    v1 = ((((3.14159 + 1e10) - float64(((1e10 - 1e10) + (-2.25 - 1.5)))) <= (0.1 - 1e10)) || (v1 && (1 <= genFn1("a", v3))))
                }
            case "":
                var v20 float64 = (-2.25 * -2.25)
                _ = v20
            case "foo bar":
                switch (-(fi14)) {
                case 9223372036854775807:
                    var v21 uint = v2
                    _ = v21
                case -2147483648:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        fmt.Println(fi22)
                        var v23 int32 = 1
                        if v1 {
                            switch (fi22 - fi14) {
                            case 42:
                                var v24 float32 = (0.1 / (0.0 - ((1.5 + (1e10 - (float32(0.0) + -2.25))) - (1e10 + 1e10))))
                                _ = v24
                            case -1:
                                var v25 int = (-(fi14))
                                _ = v25
                            case 1:
                                var v26 int32 = (v23 ^ 0)
                                _ = v26
                            default:
                                var v27 uint8 = 158
                                _ = v27
                            }
                            switch "" {
                            case "\n\t":
                                fmt.Println(fi14)
                            case "foo bar":
                                switch "a" {
                                case "":
                                    v1 = true
                                case "\n\t":
                                    var v28 uint8 = 154
                                    _ = v28
                                case "foo bar":
                                    v29 := GenS1{F0: -32768, F1: 100, F2: true}
                                    v3 += 1
                                    v1 = (2 == 702)
                                    _ = v29
                                default:
                                    fmt.Println(v23, v1, v23)
                                    switch "hello" {
                                    case "foo bar":
                                        v2 = 248
                                        if (!true) {
                                            v23 = 2147483647
                                            if len(v19) > 0 {
                                                v19[0] = fi14
                                            }
                                        }
                                        switch (fi14 * fi14) {
                                        case 32767:
                                            v23 -= 2147483647
                                        case 0:
                                            if v1 {
                                                fmt.Println(v3)
                                                v23 = (v23 & (v23 + v23))
                                            } else {
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                var v30 float64 = 1.5
                                                _ = v30
                                            }
                                        default:
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                        }
                                        v31 := GenS1{F0: 0, F1: int8(v3), F2: (563 < 0)}
                                        switch (fi22 ^ fi22) {
                                        case 796:
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                        default:
                                            fmt.Println(fi14)
                                        }
                                        v32 := GenS1{F0: 1, F1: -1, F2: v1}
                                        v23 = (v23 / (((v23 - (v23 - v23))) | 1))
                                        if v1 {
                                            if len(v19) > 0 {
                                                v19[0] = (fi22 ^ (fi22 / ((fi22) | 1)))
                                            }
                                            var v33 float64 = (((float64(1.5) * 3.14159) - 3.14159) * 0.1)
                                            _ = v33
                                        }
                                        _ = len(v19)
                                        _ = v31
                                        _ = v32
                                    case "":
                                        var v34 int16 = 42
                                        _ = v34
                                    case "hello":
                                        v35 := GenS2{F0: 0, F1: 528, F2: (^(v23))}
                                        _ = v35
                                    default:
                                        if (!("foo bar" >= "a")) {
                                            var v36 uint32 = 990
                                            switch ("hello" + "foo bar") {
                                            case "foo bar":
                                                switch "a" {
                                                case "a":
                                                    var v37 int32 = v23
                                                    _ = v37
                                                case "\n\t":
                                                    v38 := GenS3{F0: 1000, F1: v36, F2: -128}
                                                    _ = v38
                                                }
                                            }
                                            var v39 uint = v2
                                            _ = v36
                                            _ = v39
                                        } else {
                                            v40 := GenS4{F0: 1}
                                            fmt.Println(v3, fi22)
                                            fmt.Println(v2)
                                            _ = v40
                                        }
                                    }
                                }
                            default:
                                v41 := GenS5{F0: 1, F1: 283, F2: ("\n\t" + "")}
                                _ = v41
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                default:
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                }
            default:
                v42 := GenS4{F0: 1000}
                _ = v42
            }
            _ = fi14
            _ = v19
        }
        if false {
            v43 := []int{-1, -2147483648, -1}
            _ = v43
        }
    }
    v3 = v3
    v44 := GenS1{F0: 32767, F1: -128, F2: (1 >= 65535)}
    if (((-2.25 / 3.14159) > float64(float32(float64(1e10)))) || true) {
        if ((v3 != v3) || true) {
            for fi45 := 0; fi45 < 7; fi45++ {
                switch ("\n\t" + "hello") {
                case "hello":
                    v2 -= v2
                case "foo bar":
                    switch 32767 {
                    case 42:
                        if true {
                            v2 += genFn0()
                            for fi46 := 0; fi46 < 20; fi46++ {
                                if true {
                                    v2 += 0
                                    switch -9223372036854775808 {
                                    case -22:
                                        v47 := GenS5{F0: 1, F1: 321, F2: (("\n\t" + "hello") + ((("\n\t" + ("hello" + "foo bar")) + "a") + ("a" + ("\n\t" + (("foo bar" + "\n\t") + "")))))}
                                        _ = v47
                                    case -585:
                                        switch (("" + "") + (((("\n\t" + ("foo bar" + ("a" + "foo bar"))) + ("" + ("\n\t" + "a"))) + "hello") + "")) {
                                        case "\n\t":
                                            v48 := GenS2{F0: 1, F1: 42, F2: -2147483648}
                                            _ = v48
                                        }
                                    default:
                                        v49 := GenS1{F0: -1, F1: -31, F2: (v1 || (0 != 141))}
                                        _ = v49
                                    }
                                    v50 := GenS1{F0: -32768, F1: 42, F2: ((-2.25 - (1.5 - (0.0 + (1e10 / ((3.14159 + 3.14159) / ((3.14159 / 3.14159) / 1e10)))))) > 1.5)}
                                    _ = v50
                                } else {
                                    v1 = v1
                                }
                                if (true || v1) {
                                    var v51 int = 2147483647
                                    _ = v51
                                }
                                _ = fi46
                            }
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                v52 := GenS4{F0: -32768}
                                _ = v52
                            case "":
                                if false {
                                    var v53 uint = genFn0()
                                    var v54 uint8 = 128
                                    switch "hello" {
                                    case "a":
                                        if (v1 && true) {
                                            if v1 {
                                                var v55 string = ("\n\t" + "")
                                                v55 = v55
                                                _ = v55
                                            }
                                        } else {
                                            fmt.Println(v3, v53, v2)
                                            for fi56 := 0; fi56 < 11; fi56++ {
                                                if ((v3 + v3) < (^(v3))) {
                                                    v57 := []int{fi56, fi45, (^(fi45)), (-(fi56))}
                                                    v57 = append(v57, (fi56 % (((fi45 << ((2) & 63))) | 1)))
                                                    var v58 int32 = -2147483648
                                                    _ = v57
                                                    _ = v58
                                                }
                                                var v59 uint32 = 2147483648
                                                _ = fi56
                                                _ = v59
                                            }
                                            switch 335 {
                                            case 127:
                                                v1 = (461 == 18446744073709551615)
                                            }
                                        }
                                    case "foo bar":
                                        v53 += (v53 / (((v2 % ((v2) | 1))) | 1))
                                    case "\n\t":
                                        v2 = (v2 * (v2 >> ((4294967295) & 63)))
                                    default:
                                        var v60 uint32 = 0
                                        _ = v60
                                    }
                                    _ = v53
                                    _ = v54
                                }
                            default:
                                v61 := GenS6{F0: 1, F1: "\n\t", F2: -9223372036854775808}
                                _ = v61
                            }
                        } else {
                            var v62 float32 = (((1.5 / ((0.0 + float32((0.1 / -2.25))) * ((float32(1.5) / 1e10) / (0.1 - float32(1e10))))) - 0.0) - 0.0)
                            var v63 uint32 = 42
                            v64 := GenS4{F0: -1}
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                    case 317:
                        fmt.Println(v3, v2, v3)
                    case 32767:
                        for fi65 := 0; fi65 < 14; fi65++ {
                            var v66 int = (fi65 >> ((v2) & 63))
                            var v67 int16 = 677
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                    default:
                        if v1 {
                            var v68 int8 = -128
                            v69 := GenS2{F0: -77, F1: -1, F2: 2147483647}
                            _ = v68
                            _ = v69
                        }
                    }
                }
                for fi70 := 0; fi70 < 15; fi70++ {
                    v3 -= (v3 & (v3 + (v3 - (v3 + v3))))
                    switch (fi45 & fi45) {
                    case -1:
                        switch "foo bar" {
                        case "":
                            v71 := []int{32768, (fi45 ^ -9223372036854775808)}
                            _ = v71
                        }
                    default:
                        v72 := GenS7{F0: v1, F1: 0, F2: 2147483647}
                        _ = v72
                    }
                    var v73 uint8 = 128
                    _ = fi70
                    _ = v73
                }
                fmt.Println(v3, v3, fi45)
                _ = fi45
            }
            var v74 int32 = 2147483647
            v75 := GenS8{F0: 0, F1: 0}
            _ = v74
            _ = v75
        } else {
            for fi76 := 0; fi76 < 2; fi76++ {
                switch "\n\t" {
                case "\n\t":
                    for fi77 := 0; fi77 < 4; fi77++ {
                        fmt.Println(v1, fi76, v1)
                        _ = fi77
                    }
                case "foo bar":
                    fmt.Println(v2)
                default:
                    var v78 uint16 = 412
                    _ = v78
                }
                var v79 uint32 = 2147483648
                _ = fi76
                _ = v79
            }
            v1 = (false && v1)
            v80 := GenS3{F0: 1000, F1: 4294967295, F2: 127}
            _ = v80
        }
        v81 := []int{9223372036854775807, 42, 4294967295, 128}
        v2 = (v2 - (v2 + v2))
        _ = v81
    }
    v82 := []int{42, 2147483647, 0, -400}
    var v83 float32 = (3.14159 + 0.1)
    _ = len(v82)
    var v84 uint8 = 1
    switch "\n\t" {
    case "a":
        switch 125 {
        case 351:
            var v85 uint = v2
            _ = v85
        default:
            v86 := GenS9{F0: v1, F1: 529}
            _ = v86
        }
    case "foo bar":
        v1 = v1
    case "\n\t":
        switch 893 {
        case 2147483647:
            var v87 uint = 108
            _ = v87
        }
    }
    for fi88 := 0; fi88 < 7; fi88++ {
        var v89 float32 = (v83 + ((float32(3.14159) + (0.0 * float32(1e10))) / (((v83 * v83) / -2.25) * v83)))
        v90 := GenS10{F0: 42}
        _ = fi88
        _ = v89
        _ = v90
    }
    switch 1 {
    case 85:
        v83 -= ((v83 / float32(float64(v83))) * -2.25)
    case 877:
        var v91 float64 = (0.0 + (0.0 / 1.5))
        _ = v91
    default:
        v92 := GenS11{F0: 938, F1: (^(v2)), F2: ((v1 || (18446744073709551615 > 42)) || v1)}
        _ = v92
    }
    for fi93 := 0; fi93 < 8; fi93++ {
        var v94 int16 = -32768
        _ = fi93
        _ = v94
    }
    var v95 int32 = 42
    v95 = (v95 % ((1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v82:", len(v82), v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v82
    _ = v83
    _ = v84
    _ = v95
}
