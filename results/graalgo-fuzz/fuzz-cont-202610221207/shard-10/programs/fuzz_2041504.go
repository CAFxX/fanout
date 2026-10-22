package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int16; F2 string }
type GenS5 struct { F0 int; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint8; F1 string; F2 uint }
type GenS7 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 uint, p1 bool) uint32 {
    v7 := []int{873, -358}
    if len(v7) > 0 {
        v7[0] = 42
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 65535
}
func genFn1(p0 bool) uint16 {
    var v8 float64 = (float64(0.0) - -2.25)
    for fi9 := 0; fi9 < 14; fi9++ {
        p0 = true
        fmt.Println(v8)
        fmt.Println(p0)
        _ = fi9
    }
    if (42 != 255) {
        if true {
            v8 = v8
            if (!(26 >= 1)) {
                v10 := []int{32767}
                v11 := GenS1{F0: 268, F1: 1, F2: -1}
                if len(v10) > 0 {
                    v10[0] = -822
                }
                _ = v10
                _ = v11
            }
        } else {
            switch 32767 {
            case -2147483648:
                p0 = ((4294967295 <= 65535) && false)
            case 32767:
                var v12 int32 = 2147483647
                _ = v12
            case 128:
                if ((-2.25 + v8) < ((v8 - float64(0.1)) + (((v8 + (float64(-2.25) * ((0.0 - 3.14159) / v8))) * v8) * (((v8 / 0.1) / ((((v8 / v8) - v8) * v8) - -2.25)) / float64(-2.25))))) {
                    v13 := []int{9223372036854775807, 42}
                    p0 = ((("" + "\n\t") + "a") == "\n\t")
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                for fi15 := 0; fi15 < 8; fi15++ {
                    fmt.Println(p0, p0, fi15)
                    _ = fi15
                }
                _ = fi14
            }
        }
    }
    var v16 int = 127
    _ = p0
    _ = v16
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", "foo bar", "\n\t"}
    var v3 int32 = 2147483647
    v4 := make(map[int]string)
    v5 := []string{"hello", "hello"}
    var v6 int64 = -887
    v3 -= v3
    if (1 <= 434) {
        v6 -= v6
    }
    var v17 float32 = 3.14159
    var v18 float64 = ((-2.25 * (0.0 * 0.0)) * 1.5)
    v17 = ((3.14159 / v17) + v17)
    var v19 string = ("foo bar" + "")
    var v20 uint32 = 2147483647
    var v21 int64 = v6
    var v22 uint64 = 295
    v23 := GenS1{F0: 65535, F1: 1000, F2: (^(v3))}
    if true {
        if (v3 > v3) {
            var v24 int32 = (v3 << ((65535) & 63))
            if (false && true) {
                fmt.Println(v19, v17)
                v18 -= -2.25
            }
            if false {
                var v25 int8 = -16
                _ = v25
            } else {
                var v26 int32 = -411
                v22 -= (v22 << ((0) & 63))
                fmt.Println(v24)
                _ = v26
            }
            _ = v24
        } else {
            for fi27 := 0; fi27 < 12; fi27++ {
                if true {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        v20 -= v20
                        v3 = (v3 >> ((uint(v20)) & 63))
                        v18 -= (float64(v17) + -2.25)
                        _ = fi28
                    }
                    switch "\n\t" {
                    case "a":
                        for fi29 := 0; fi29 < 20; fi29++ {
                            fmt.Println(v6)
                            var v30 uint16 = 613
                            if false {
                                if ((fi27 <= fi27) || true) {
                                    var v31 int64 = v6
                                    _ = v31
                                }
                                v22 = v22
                                v22 -= 0
                            } else {
                                fmt.Println(v17, v3, fi29)
                            }
                            _ = fi29
                            _ = v30
                        }
                    default:
                        for fi32 := 0; fi32 < 2; fi32++ {
                            var v33 uint16 = genFn1(false)
                            fmt.Println(v3, v33)
                            if false {
                                if (18446744073709551615 < 9) {
                                    if false {
                                        switch (((v19 + "") + "a") + v19) {
                                        case "a":
                                            var v34 uint64 = v22
                                            _ = v34
                                        case "hello":
                                            fmt.Println(v20, fi32)
                                            if len(v2) > 0 {
                                                v2[0] = v19
                                            }
                                            v35 := GenS1{F0: (v33 - (v33 * (v33 % ((32767) | 1)))), F1: 361, F2: 42}
                                            if ((float32(-2.25) < (0.1 + v17)) || false) {
                                                v21 -= v21
                                                var v36 uint64 = v22
                                                _ = v36
                                            }
                                            _ = v35
                                        case "foo bar":
                                            if (((v3 % ((v3) | 1)) < (v3 - v3)) && (false || true)) {
                                                v21 = v21
                                                switch "a" {
                                                case "":
                                                    v2 = append(v2, v19)
                                                case "foo bar":
                                                    switch v19 {
                                                    case "foo bar":
                                                        v37 := GenS1{F0: genFn1(false), F1: 1000, F2: (v3 % ((v3) | 1))}
                                                        _ = v37
                                                    case "hello":
                                                        var v38 int32 = v3
                                                        _ = v38
                                                    case "a":
                                                        v22 += (v22 << ((uint(v33)) & 63))
                                                    }
                                                case "hello":
                                                    var v39 int64 = (v6 % ((42) | 1))
                                                    if (0 <= 4294967295) {
                                                        if ((v22 < 0) && false) {
                                                            if ((true || (false && (43 > 255))) && (-36 != -128)) {
                                                                v40 := GenS2{F0: v22, F1: 21}
                                                                _ = v40
                                                            } else {
                                                                v41 := GenS1{F0: 1, F1: 42, F2: v3}
                                                                v20 = (v20 - v20)
                                                                v2 = append(v2, v19)
                                                                _ = v41
                                                            }
                                                        }
                                                        var v42 uint8 = 0
                                                        _ = v42
                                                    }
                                                    v22 = (v22 / ((1) | 1))
                                                    _ = v39
                                                default:
                                                    v17 += (v17 * ((v17 - float32((v18 - (v18 * 1e10)))) + v17))
                                                    var v43 uint = 18446744073709551615
                                                    v6 = (v6 - (^(v6)))
                                                    _ = v43
                                                }
                                                v6 += (^(v21))
                                            } else {
                                                v44 := GenS2{F0: (v22 >> ((111) & 63)), F1: -1}
                                                _ = v44
                                            }
                                        default:
                                            var v45 int32 = v3
                                            _ = v45
                                        }
                                        var v46 uint = 540
                                        var v47 int64 = 9223372036854775807
                                        _ = v46
                                        _ = v47
                                    } else {
                                        var v48 int8 = 1
                                        _ = v48
                                    }
                                    fmt.Println(fi32, v21)
                                } else {
                                    fmt.Println(v20, fi27)
                                    v19 = ("\n\t" + (v19 + v19))
                                }
                            } else {
                                fmt.Println(v20, v19, v21)
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = (("\n\t" + v19) + "")
                                    }
                                } else {
                                    _ = len(v1)
                                }
                            }
                            var v49 int32 = v3
                            var v50 int16 = 1
                            _ = fi32
                            _ = v33
                            _ = v49
                            _ = v50
                        }
                    }
                    switch "foo bar" {
                    case "":
                        v51 := GenS3{F0: 189}
                        _ = v51
                    case "\n\t":
                        for fi52 := 0; fi52 < 13; fi52++ {
                            switch v19 {
                            case "foo bar":
                                v53 := GenS2{F0: v22, F1: 7}
                                v20 -= (v20 & 479)
                                _ = v53
                            }
                            _ = fi52
                        }
                    }
                } else {
                    v18 = 0.0
                    if false {
                        if false {
                            v54 := GenS4{F0: 0, F1: 32767, F2: ((v19 + v19) + (v19 + v19))}
                            var v55 uint64 = v22
                            if (127 == 42) {
                                switch fi27 {
                                case 278:
                                    if true {
                                        v55 = v55
                                    }
                                case -949:
                                    switch fi27 {
                                    case 634:
                                        v19 = "a"
                                    case 674:
                                        v19 = ((v19 + (v19 + v19)) + v19)
                                    default:
                                        var v56 uint8 = uint8(v22)
                                        _ = v56
                                    }
                                case 176:
                                    switch v19 {
                                    case "foo bar":
                                        v57 := GenS5{F0: fi27, F1: 0, F2: -1}
                                        _ = v57
                                    }
                                default:
                                    for fi58 := 0; fi58 < 10; fi58++ {
                                        var v59 uint32 = v20
                                        fmt.Println(v20, v19, v6)
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                                fmt.Println(v3, v22)
                            }
                            _ = v54
                            _ = v55
                        } else {
                            v22 -= (v22 << ((0) & 63))
                            switch "\n\t" {
                            case "a":
                                if (true && (4294967295 == 665)) {
                                    for fi60 := 0; fi60 < 11; fi60++ {
                                        fmt.Println(v3, v6, v22)
                                        v17 = v17
                                        _ = fi60
                                    }
                                }
                            default:
                                if (true && false) {
                                    var v61 int64 = (^(v6))
                                    fmt.Println(v22, v18, v6)
                                    var v62 int = fi27
                                    _ = v61
                                    _ = v62
                                } else {
                                    for fi63 := 0; fi63 < 13; fi63++ {
                                        fmt.Println(v18, v3)
                                        if ((v3 | int32(v22)) > v3) {
                                            v64 := GenS1{F0: genFn1((68 > 127)), F1: 509, F2: (^(v3))}
                                            _ = len(v1)
                                            var v65 int = int(v22)
                                            _ = v64
                                            _ = v65
                                        }
                                        var v66 uint16 = 32767
                                        _ = fi63
                                        _ = v66
                                    }
                                    v20 += genFn0(42, (!(!false)))
                                    for fi67 := 0; fi67 < 20; fi67++ {
                                        var v68 uint64 = (v22 & (v22 / ((18446744073709551615) | 1)))
                                        if true {
                                            v22 -= (v68 / (((^(v22))) | 1))
                                            v69 := GenS4{F0: -224, F1: -1000, F2: (v19 + v19)}
                                            if (!(!((v3 >> ((1) & 63)) != -1))) {
                                                v70 := GenS6{F0: 128, F1: v19, F2: 4294967295}
                                                if ((127 == 85) && ((0 < 1) && (2 >= 255))) {
                                                    var v71 int = -994
                                                    _ = v71
                                                }
                                                v20 = (^(v20))
                                                _ = v70
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = ("foo bar" + "")
                                                }
                                                v6 = v6
                                                v22 += v22
                                            }
                                            _ = v69
                                        }
                                        fmt.Println(v17, v20, v68)
                                        if (127 < 127) {
                                            var v72 uint16 = genFn1((false && (v17 == 0.0)))
                                            fmt.Println(fi27)
                                            var v73 uint16 = (^(v72))
                                            _ = v72
                                            _ = v73
                                        } else {
                                            var v74 uint32 = genFn0(0, false)
                                            fmt.Println(v17)
                                            switch "" {
                                            case "\n\t":
                                                v75 := GenS7{F0: 530, F1: (v68 << ((42) & 63)), F2: 211}
                                                _ = v75
                                            case "foo bar":
                                                v17 += v17
                                            case "":
                                                var v76 uint = 42
                                                _ = v76
                                            }
                                            _ = v74
                                        }
                                        v77 := GenS3{F0: 1}
                                        _ = fi67
                                        _ = v68
                                        _ = v77
                                    }
                                }
                            }
                            v78 := GenS2{F0: v22, F1: 100}
                            _ = v78
                        }
                        for fi79 := 0; fi79 < 18; fi79++ {
                            v80 := v1[v19]
                            if (-1 >= v3) {
                                v81 := GenS3{F0: 2}
                                fmt.Println(fi27, v22, fi79)
                                _ = v81
                            } else {
                                fmt.Println(v19)
                                var v82 int16 = 735
                                _ = v82
                            }
                            v83 := GenS8{F0: (^(v20))}
                            _ = fi79
                            _ = v80
                            _ = v83
                        }
                    } else {
                        for fi84 := 0; fi84 < 1; fi84++ {
                            v18 += v18
                            _ = fi84
                        }
                        delete(v4, (-(fi27)))
                    }
                }
                var v85 int32 = v3
                _ = fi27
                _ = v85
            }
            if len(v2) > 0 {
                v2[0] = ((v19 + "\n\t") + "hello")
            }
            if (false || false) {
                v86 := v4[32768]
                _ = v86
            } else {
                delete(v1, v19)
                var v87 float64 = 0.1
                var v88 int32 = (v3 % ((2147483647) | 1))
                _ = v87
                _ = v88
            }
        }
    }
    for fi89 := 0; fi89 < 9; fi89++ {
        switch (("" + v19) + ((v19 + v19) + v19)) {
        case "foo bar":
            v20 = (v20 / ((4294967295) | 1))
        case "a":
            v1["foo bar"] = (v20 % ((v20) | 1))
        }
        if (false && false) {
            v90 := v1[v19]
            _ = v90
        }
        _ = fi89
    }
    fmt.Println(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
