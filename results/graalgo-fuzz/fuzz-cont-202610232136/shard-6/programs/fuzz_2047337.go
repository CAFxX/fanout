package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    v1 := []int{128, 32768, 0, 63}
    var v2 uint32 = 525
    var v3 int16 = 768
    var v4 uint64 = 0
    var v5 uint32 = 426
    var v6 uint64 = v4
    for fi7 := 0; fi7 < 5; fi7++ {
        v1 = append(v1, 546)
        if (42 < -2147483648) {
            v5 -= v5
            var v8 int64 = 577
            _ = v8
        }
        switch fi7 {
        case 9223372036854775807:
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 bool = ((v6 ^ (v4 & (v6 ^ (v4 * (v6 ^ 1))))) == 648)
                v4 += (v4 | v4)
                _ = fi9
                _ = v10
            }
        case -602:
            switch "hello" {
            case "\n\t":
                _ = len(v1)
            default:
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 float64 = 0.1
                    _ = fi11
                    _ = v12
                }
            }
        case 32767:
            if ((-1 != -2147483648) && true) {
                v4 += 42
                for fi13 := 0; fi13 < 18; fi13++ {
                    _ = len(v1)
                    _ = fi13
                }
                if (0.0 > (1.5 * float32((float64(0.1) * 0.1)))) {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 float64 = (-2.25 * 1.5)
                        if ((fi14 ^ (fi7 & (fi7 & fi14))) != fi7) {
                            var v16 bool = true
                            v4 = 58
                            _ = v16
                        } else {
                            v15 -= float64((1e10 + 0.1))
                            var v17 int = -9223372036854775808
                            if (!true) {
                                v17 -= 32768
                                if false {
                                    v18 := GenS1{F0: (v3 | -1), F1: false, F2: 255}
                                    if (((3.14159 + v15) >= 0.0) || false) {
                                        var v19 int32 = 641
                                        var v20 uint64 = v4
                                        v21 := GenS1{F0: (-(v3)), F1: false, F2: 77}
                                        fmt.Println(v6, v17, v5)
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                    switch v17 {
                                    case -2147483648:
                                        if len(v1) > 0 {
                                            v1[0] = (fi7 & (v17 + 0))
                                        }
                                    case 42:
                                        v1 = append(v1, 495)
                                    case 128:
                                        v4 = 4294967295
                                    }
                                    fmt.Println(v15, v3, v2)
                                    _ = v18
                                }
                                var v22 string = "hello"
                                _ = v22
                            }
                            fmt.Println(v4, v17)
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                    v6 = (^(v6))
                    if false {
                        v23 := GenS1{F0: 722, F1: (!false), F2: 2}
                        _ = v23
                    } else {
                        if true {
                            fmt.Println(v3)
                            var v24 int8 = 1
                            var v25 uint8 = 128
                            _ = v24
                            _ = v25
                        } else {
                            switch "" {
                            case "\n\t":
                                for fi26 := 0; fi26 < 7; fi26++ {
                                    var v27 uint = 946
                                    var v28 uint8 = 0
                                    var v29 int32 = -937
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            case "foo bar":
                                var v30 int32 = -556
                                _ = v30
                            case "hello":
                                for fi31 := 0; fi31 < 12; fi31++ {
                                    fmt.Println(v2, v5)
                                    if len(v1) > 0 {
                                        v1[0] = fi7
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = fi7
                                    }
                                    v2 -= (v5 << ((4294967295) & 63))
                                    _ = fi31
                                }
                            }
                            fmt.Println(v5, v5)
                            var v32 int16 = -32768
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 2; fi33++ {
                            v34 := GenS1{F0: v3, F1: true, F2: 0}
                            var v35 bool = false
                            var v36 int64 = 0
                            var v37 int8 = -128
                            if (!false) {
                                fmt.Println(v6, fi7)
                                switch fi7 {
                                case 4294967295:
                                    var v38 float64 = 0.1
                                    _ = v38
                                case 9223372036854775807:
                                    if (!((v35 && false) || ((v37 << ((484) & 63)) < (v37 + -128)))) {
                                        fmt.Println(v3, v37, v2)
                                        fmt.Println(v6, v2)
                                        var v39 bool = false
                                        _ = v39
                                    } else {
                                        var v40 int64 = v36
                                        var v41 string = (((("hello" + (("" + "\n\t") + "")) + "") + ("" + "")) + "")
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        v1 = append(v1, (fi7 & 2147483647))
                                    case "foo bar":
                                        v6 -= 692
                                        v4 -= (v4 + v4)
                                        v1 = append(v1, fi33)
                                    default:
                                        v1 = append(v1, (fi7 << ((4294967295) & 63)))
                                    }
                                    v4 = 181
                                }
                                v4 -= 2
                            }
                            _ = fi33
                            _ = v34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                }
            } else {
                if true {
                    var v42 bool = true
                    _ = v42
                } else {
                    switch "a" {
                    case "\n\t":
                        for fi43 := 0; fi43 < 8; fi43++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi43
                        }
                    default:
                        if (!true) {
                            for fi44 := 0; fi44 < 3; fi44++ {
                                v4 -= 4294967295
                                if true {
                                    if (2 <= 18446744073709551615) {
                                        if (837 > v6) {
                                            switch fi44 {
                                            case 94:
                                                var v45 float32 = 1.5
                                                _ = v45
                                            }
                                            v2 = (v5 & v2)
                                            v4 = (^(v4))
                                            v6 += 0
                                        }
                                        fmt.Println(v3)
                                    } else {
                                        if true {
                                            var v46 float64 = (float64(float32(1e10)) * -2.25)
                                            if true {
                                                fmt.Println(v4)
                                                var v47 uint16 = 748
                                                _ = v47
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v46 -= 1.5
                                                switch "" {
                                                case "foo bar":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "hello":
                                                    var v48 int32 = -439
                                                    _ = v48
                                                }
                                            }
                                            _ = v46
                                        } else {
                                            v49 := GenS1{F0: v3, F1: true, F2: 26}
                                            _ = v49
                                        }
                                        if true {
                                            switch "hello" {
                                            case "":
                                                v6 += v4
                                            case "a":
                                                v6 = (^(v4))
                                            case "\n\t":
                                                fmt.Println(fi44, fi44, v5)
                                            default:
                                                v3 -= v3
                                            }
                                        } else {
                                            v4 += v4
                                            v4 += 962
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch (fi44 % ((0) | 1)) {
                                        case -618:
                                            var v50 int64 = -1
                                            _ = v50
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v4)
                                            case "":
                                                var v51 int = 1
                                                _ = v51
                                            }
                                        }
                                    }
                                } else {
                                    v1 = append(v1, (fi44 ^ (fi44 - fi7)))
                                    fmt.Println(fi44, v5)
                                    var v52 int64 = -489
                                    v2 = (^(v2))
                                    v2 = (v5 << ((uint(v4)) & 63))
                                    v3 = 0
                                    var v53 string = "\n\t"
                                    _ = v52
                                    _ = v53
                                }
                                _ = len(v1)
                                _ = fi44
                            }
                            fmt.Println(v3, v5)
                            if false {
                                if true {
                                    var v54 float64 = 3.14159
                                    _ = v54
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v55 bool = true
                                v56 := GenS1{F0: (v3 ^ v3), F1: v55, F2: uint8(v4)}
                                _ = v55
                                _ = v56
                            }
                        }
                    }
                }
                fmt.Println(fi7, v4, v6)
                if true {
                    if (!true) {
                        if false {
                            if true {
                                for fi57 := 0; fi57 < 20; fi57++ {
                                    switch ((("\n\t" + "a") + "\n\t") + ("a" + "\n\t")) {
                                    case "":
                                        fmt.Println(fi57, v6, v4)
                                    case "hello":
                                        v58 := GenS2{F0: "\n\t"}
                                        _ = v58
                                    }
                                    _ = fi57
                                }
                            } else {
                                v59 := GenS1{F0: (v3 - (v3 << ((373) & 63))), F1: ((v3 / (((v3 + v3)) | 1)) <= v3), F2: 1}
                                _ = v59
                            }
                            fmt.Println(v3, v2)
                        }
                        v60 := GenS3{F0: 18446744073709551615, F1: 295, F2: 42}
                        fmt.Println(fi7)
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 1; fi61++ {
                        v3 += 42
                        _ = fi61
                    }
                    var v62 float64 = float64(3.14159)
                    _ = v62
                } else {
                    var v63 bool = (true && (true && false))
                    _ = v63
                }
            }
        }
        _ = fi7
    }
    v3 -= (v3 << ((979) & 63))
    var v64 uint64 = (v6 + 1)
    v65 := GenS4{F0: (v5 - 2147483647)}
    v66 := GenS4{F0: v2}
    fmt.Println(v3, v2)
    switch "foo bar" {
    case "foo bar":
        if false {
            for fi67 := 0; fi67 < 5; fi67++ {
                switch "foo bar" {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    v68 := GenS5{F0: ("" + ("hello" + "foo bar")), F1: 255, F2: 485}
                    _ = v68
                case "a":
                    for fi69 := 0; fi69 < 17; fi69++ {
                        switch fi69 {
                        case -266:
                            var v70 bool = true
                            _ = v70
                        case 0:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case -868:
                            var v71 uint8 = 0
                            _ = v71
                        default:
                            v4 = (v64 ^ v6)
                        }
                        _ = fi69
                    }
                default:
                    var v72 uint64 = (v6 | 1)
                    _ = v72
                }
                v73 := GenS3{F0: 42, F1: 42, F2: 0}
                _ = fi67
                _ = v73
            }
        } else {
            fmt.Println(v3)
            var v74 uint32 = (v2 / ((v5) | 1))
            for fi75 := 0; fi75 < 15; fi75++ {
                var v76 int64 = 42
                var v77 float64 = 0.0
                v74 = v2
                _ = fi75
                _ = v76
                _ = v77
            }
            _ = v74
        }
    case "hello":
        v3 -= (v3 * (^(v3)))
    case "":
        var v78 int16 = -727
        _ = v78
    }
    v1 = append(v1, -395)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
