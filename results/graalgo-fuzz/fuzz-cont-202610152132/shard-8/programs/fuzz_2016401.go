package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint16; F2 string }
type GenS5 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int32; F1 int; F2 string }
type GenS7 struct { F0 uint8; F1 string; F2 uint8 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 uint64) string {
    var v7 string = "foo bar"
    v8 := GenS1{F0: (42 == 0), F1: -718, F2: 555}
    _ = p0
    _ = v7
    _ = v8
    return ("" + ((v7 + v7) + v7))
}
func genFn1(p0 string, p1 int64, p2 bool) uint8 {
    p2 = p2
    for fi9 := 0; fi9 < 11; fi9++ {
        p0 = (p0 + p0)
        v10 := []int{0, (-(fi9))}
        _ = fi9
        _ = v10
    }
    if (p2 && p2) {
        v11 := []int{1, 42}
        _ = v11
    }
    v12 := GenS1{F0: false, F1: int32(p1), F2: 355}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 127
}

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 1
    var v3 uint64 = 42
    var v4 uint32 = 2
    var v5 uint = 255
    var v6 uint8 = 2
    v13 := GenS2{F0: -790, F1: "a"}
    if false {
        if (42 == 9223372036854775807) {
            if (true || ((v5 ^ (v5 ^ v5)) > v5)) {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v1 = float32(0.1)
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v3 = 4294967295
                        var v16 int = 4294967295
                        switch (fi15 & (^(v16))) {
                        case 2147483647:
                            if (0.0 != 0.1) {
                                v6 -= (v6 ^ (v6 | genFn1(("\n\t" + ("hello" + "hello")), 655, (int16(v16) < 472))))
                                switch "hello" {
                                case "foo bar":
                                    switch (^(fi15)) {
                                    case -678:
                                        var v17 int = (v16 + (v16 << (((v5 / (((v5 - (v5 ^ v5))) | 1))) & 63)))
                                        _ = v17
                                    }
                                }
                                if (42 < 950) {
                                    fmt.Println(v16, v16, v5)
                                    v18 := []int{-9223372036854775808, 128}
                                    var v19 uint32 = v4
                                    var v20 uint32 = v4
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    v21 := GenS2{F0: 0, F1: "foo bar"}
                                    v22 := []int{int(v4), fi14, fi15, (fi14 & v16)}
                                    v22 = append(v22, (v16 - fi15))
                                    if true {
                                        fmt.Println(v6, v4, v5)
                                        v5 -= (^(v5))
                                    } else {
                                        _ = len(v22)
                                        if true {
                                            v23 := GenS3{F0: 2147483647}
                                            if len(v22) > 0 {
                                                v22[0] = int(v2)
                                            }
                                            fmt.Println(v6)
                                            _ = v23
                                        }
                                    }
                                    fmt.Println(v2)
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                v24 := GenS3{F0: 1}
                                switch "" {
                                case "":
                                    if true {
                                        v4 = (v4 & v4)
                                        v25 := GenS1{F0: true, F1: 0, F2: v5}
                                        _ = v25
                                    }
                                case "foo bar":
                                    var v26 float32 = 0.0
                                    fmt.Println(v26)
                                    v27 := []int{}
                                    var v28 uint16 = 1
                                    if (!true) {
                                        if (v5 != (v5 - 42)) {
                                            switch -2147483648 {
                                            case -505:
                                                v29 := GenS1{F0: false, F1: 958, F2: uint(v3)}
                                                _ = v29
                                            case 128:
                                                var v30 int32 = int32(v2)
                                                _ = v30
                                            }
                                            v31 := GenS3{F0: -97}
                                            fmt.Println(fi14, v5)
                                            _ = v31
                                        }
                                    }
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                case "a":
                                    v32 := GenS4{F0: int16(fi15), F1: 32767, F2: ("foo bar" + "hello")}
                                    _ = v32
                                }
                                v1 -= float32((1.5 / (3.14159 + -2.25)))
                                _ = v24
                            }
                        default:
                            var v33 uint16 = 326
                            _ = v33
                        }
                        _ = fi15
                        _ = v16
                    }
                    if false {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            v4 = uint32(v3)
                            if true {
                                var v35 int8 = v2
                                _ = v35
                            }
                            var v36 float64 = 1.5
                            if (false && ((3.14159 - v36) >= v36)) {
                                if (v1 != v1) {
                                    switch 1 {
                                    case 4294967295:
                                        var v37 bool = true
                                        _ = v37
                                    }
                                    fmt.Println(v6, v4, v1)
                                    v6 = v6
                                    v38 := GenS5{F0: 197, F1: -9223372036854775808, F2: 32767}
                                    fmt.Println(fi14, v3, v5)
                                    _ = v38
                                } else {
                                    v2 = v2
                                    var v39 int16 = -32768
                                    _ = v39
                                }
                                fmt.Println(v1, v2, v3)
                            } else {
                                if (!true) {
                                    v5 += (v5 | v5)
                                }
                                var v40 string = "foo bar"
                                v41 := []int{(fi14 | fi34), -9223372036854775808, (fi14 ^ fi34)}
                                _ = v40
                                _ = v41
                            }
                            v42 := []int{}
                            _ = fi34
                            _ = v36
                            _ = v42
                        }
                        switch "a" {
                        case "foo bar":
                            if true {
                                if (int16(v2) != 1) {
                                    var v43 int = -9223372036854775808
                                    if false {
                                        if true {
                                            for fi44 := 0; fi44 < 5; fi44++ {
                                                v4 = uint32(v5)
                                                v2 -= (v2 / ((v2) | 1))
                                                _ = fi44
                                            }
                                        } else {
                                            v43 = (v43 << (((v5 | (v5 << ((4294967295) & 63)))) & 63))
                                            for fi45 := 0; fi45 < 6; fi45++ {
                                                var v46 uint8 = (^(v6))
                                                v47 := []int{fi45, v43}
                                                _ = fi45
                                                _ = v46
                                                _ = v47
                                            }
                                            if false {
                                                v5 -= v5
                                            } else {
                                                if false {
                                                    for fi48 := 0; fi48 < 5; fi48++ {
                                                        v49 := []int{879, (fi48 / ((fi14) | 1)), (^(fi14)), fi14}
                                                        v50 := GenS6{F0: -945, F1: (fi48 ^ (fi14 & v43)), F2: ("foo bar" + "foo bar")}
                                                        _ = fi48
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                }
                                            }
                                        }
                                        switch "hello" {
                                        case "":
                                            v51 := []int{fi14, (fi14 + (v43 << (((v5 / ((42) | 1))) & 63)))}
                                            _ = v51
                                        }
                                    } else {
                                        if true {
                                            var v52 int = (v43 * fi14)
                                            _ = v52
                                        }
                                        v53 := GenS7{F0: uint8(v5), F1: (("" + ("\n\t" + "")) + ("\n\t" + ((("" + ("hello" + "hello")) + (("hello" + ("foo bar" + "a")) + "a")) + ((("foo bar" + "foo bar") + ("a" + "hello")) + ((("foo bar" + ("foo bar" + "a")) + (("" + ("" + "foo bar")) + "")) + "foo bar"))))), F2: (v6 >> (((v5 / (((v5 / ((18446744073709551615) | 1))) | 1))) & 63))}
                                        v3 = (v3 | (v3 >> ((v5) & 63)))
                                        _ = v53
                                    }
                                    _ = v43
                                }
                                if ("\n\t" >= "a") {
                                    var v54 uint8 = v6
                                    _ = v54
                                }
                                v3 = uint64(v6)
                            }
                        case "\n\t":
                            v55 := []int{}
                            _ = v55
                        default:
                            switch ("" + "foo bar") {
                            case "a":
                                v6 = (v6 ^ v6)
                            }
                        }
                    } else {
                        var v56 float64 = (((-2.25 * 3.14159) + (float64(3.14159) * ((1.5 - 0.0) - (1e10 - 1e10)))) + 1.5)
                        v57 := []int{-2147483648, (fi14 / ((fi14) | 1)), (-(fi14)), (fi14 % (((^(fi14))) | 1))}
                        fmt.Println(v4, v6)
                        _ = v56
                        _ = v57
                    }
                    _ = fi14
                }
            }
        }
        v58 := []int{-9223372036854775808}
        _ = v58
    } else {
        v59 := GenS8{F0: (false || false), F1: v2}
        _ = v59
    }
    if (!(-835 <= 190)) {
        if (false || (128 <= 128)) {
            v60 := []int{32767, 32767, int(v4), -9223372036854775808}
            var v61 uint = (^(v5))
            fmt.Println(v6, v6)
            _ = v60
            _ = v61
        } else {
            for fi62 := 0; fi62 < 15; fi62++ {
                v63 := []int{fi62, fi62}
                var v64 string = "a"
                _ = fi62
                _ = v63
                _ = v64
            }
            switch "foo bar" {
            case "\n\t":
                v5 -= uint(v3)
            case "hello":
                v6 -= (v6 * (v6 + v6))
            }
        }
        var v65 bool = true
        for fi66 := 0; fi66 < 8; fi66++ {
            for fi67 := 0; fi67 < 15; fi67++ {
                fmt.Println(fi66)
                if v65 {
                    var v68 int = 4294967295
                    _ = v68
                }
                _ = fi67
            }
            v69 := GenS9{F0: 4294967295}
            _ = fi66
            _ = v69
        }
        _ = v65
    }
    v4 = (v4 | v4)
    if true {
        v1 = 1e10
    } else {
        v70 := []int{-2147483648, int(v3), -597, 1}
        _ = v70
    }
    var v71 uint16 = 2
    var v72 string = "hello"
    if (((((v72 + (v72 + ("\n\t" + "\n\t"))) + "\n\t") + v72) + (v72 + "a")) < (v72 + (v72 + "a"))) {
        v73 := []int{17, 32767, 127, 127}
        if true {
            switch ((v72 + v72) + v72) {
            case "foo bar":
                var v74 float32 = ((1.5 * v1) * ((1.5 * v1) - 3.14159))
                _ = v74
            case "hello":
                fmt.Println(v1)
            default:
                for fi75 := 0; fi75 < 12; fi75++ {
                    fmt.Println(v3)
                    _ = fi75
                }
            }
            for fi76 := 0; fi76 < 17; fi76++ {
                v2 += v2
                var v77 int64 = -567
                v78 := GenS9{F0: v5}
                _ = fi76
                _ = v77
                _ = v78
            }
            var v79 uint16 = 65535
            _ = v79
        } else {
            v80 := GenS10{F0: (v71 + 0)}
            v72 = "a"
            _ = v80
        }
        _ = v73
    } else {
        var v81 int8 = 1
        _ = v81
    }
    v5 = (^(v5))
    var v82 int = 2147483647
    for fi83 := 0; fi83 < 6; fi83++ {
        if (false || false) {
            var v84 uint16 = v71
            v4 -= (^(v4))
            v85 := GenS4{F0: 1, F1: (^(v84)), F2: v72}
            _ = v84
            _ = v85
        } else {
            var v86 bool = false
            fmt.Println(v2, v3, v72)
            v3 = uint64(v4)
            _ = v86
        }
        v4 = (v4 >> ((v5) & 63))
        _ = fi83
    }
    fmt.Println(v6, v5)
    v87 := GenS11{F0: (v71 / ((136) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
    _ = v82
    _ = v87
}
