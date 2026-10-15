package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int64; F1 uint32; F2 int16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 int64; F1 string; F2 uint32 }

func main() {
    var v1 uint32 = 846
    var v2 uint64 = 146
    v3 := []uint32{(^(v1)), v1, 42}
    var v4 int32 = 934
    fmt.Println(v4, v1)
    v5 := GenS1{F0: v4, F1: -913, F2: 0}
    if (0 == 65535) {
        var v6 uint = 1
        var v7 float64 = ((0.1 + 1.5) + ((3.14159 - float64((1.5 / (float32((0.1 - 0.0)) - float32(1.5))))) - 1.5))
        for fi8 := 0; fi8 < 14; fi8++ {
            switch 32768 {
            case 894:
                switch "foo bar" {
                case "foo bar":
                    v9 := GenS2{F0: (^(v6)), F1: 1000}
                    _ = v9
                case "hello":
                    if false {
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v11 := GenS1{F0: (-(v4)), F1: v4, F2: 42}
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        var v12 int16 = 32767
                        switch "a" {
                        case "a":
                            switch "\n\t" {
                            case "":
                                var v13 bool = true
                                _ = v13
                            case "hello":
                                if (v2 != (v2 >> ((v6) & 63))) {
                                    switch ("hello" + ("foo bar" + ("a" + ""))) {
                                    case "":
                                        v3 = append(v3, (^(v1)))
                                    }
                                }
                            case "foo bar":
                                switch "" {
                                case "hello":
                                    var v14 uint32 = v1
                                    _ = v14
                                case "foo bar":
                                    _ = len(v3)
                                case "a":
                                    var v15 int8 = 100
                                    _ = v15
                                default:
                                    for fi16 := 0; fi16 < 12; fi16++ {
                                        var v17 bool = false
                                        var v18 bool = (true && ("hello" <= "hello"))
                                        _ = fi16
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            }
                        case "hello":
                            fmt.Println(v1, fi8)
                        case "foo bar":
                            if true {
                                switch "a" {
                                case "":
                                    var v19 uint16 = 2
                                    _ = v19
                                default:
                                    if (440 < (^(v2))) {
                                        var v20 float32 = (((((((-2.25 - ((1.5 / 0.1) + 1e10)) * ((1e10 * 0.0) - (3.14159 - 0.0))) - 3.14159) + 1e10) + (float32((v7 - v7)) + 1.5)) / -2.25) * 1e10)
                                        if false {
                                            fmt.Println(v7, v12, v12)
                                        } else {
                                            fmt.Println(v7, v2)
                                            switch "a" {
                                            case "hello":
                                                var v21 uint64 = 1
                                                _ = v21
                                            case "foo bar":
                                                switch ("a" + ("\n\t" + (("a" + "") + "hello"))) {
                                                case "\n\t":
                                                    var v22 uint8 = 2
                                                    _ = v22
                                                }
                                            case "":
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = 4294967295
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = v1
                                                    }
                                                    var v23 uint64 = (v2 >> ((0) & 63))
                                                    _ = v23
                                                } else {
                                                    switch (fi8 | (fi8 - fi8)) {
                                                    case 1:
                                                        if len(v3) > 0 {
                                                            v3[0] = 42
                                                        }
                                                    case -2147483648:
                                                        for fi24 := 0; fi24 < 13; fi24++ {
                                                            fmt.Println(v7)
                                                            switch fi24 {
                                                            case -9223372036854775808:
                                                                v12 = v12
                                                            case -405:
                                                                fmt.Println(v2, v6)
                                                                if (44 <= -124) {
                                                                    fmt.Println(v20)
                                                                    v7 += v7
                                                                    v25 := GenS3{F0: 127, F1: fi24}
                                                                    _ = v25
                                                                } else {
                                                                    var v26 uint8 = 1
                                                                    _ = v26
                                                                }
                                                                switch fi8 {
                                                                case 128:
                                                                    v20 = float32(0.0)
                                                                default:
                                                                    fmt.Println(fi24)
                                                                }
                                                                if ((!true) && false) {
                                                                    var v27 uint16 = 65535
                                                                    _ = v27
                                                                } else {
                                                                    v1 = (v1 * (v1 - v1))
                                                                    if ((1.5 / -2.25) != (3.14159 * 1e10)) {
                                                                        v7 = v7
                                                                        v6 += v6
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                if true {
                                                                    v28 := GenS4{F0: 42, F1: 184, F2: 1}
                                                                    var v29 uint64 = (v2 & 42)
                                                                    _ = v28
                                                                    _ = v29
                                                                } else {
                                                                    fmt.Println(v2, v20)
                                                                }
                                                                if ((!false) || (!(false || false))) {
                                                                    _ = len(v3)
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        fmt.Println(v20)
                                                                    default:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v1 >> (((^(v6))) & 63))
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v7 += 3.14159
                                                                        v4 = v4
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v1
                                                                        }
                                                                        _ = len(v3)
                                                                    } else {
                                                                        v20 = v20
                                                                    }
                                                                }
                                                            }
                                                            var v30 bool = (42 <= 40)
                                                            v30 = (!((v4 + v4) == 42))
                                                            if v30 {
                                                                if (!v30) {
                                                                    v6 -= (v6 * 18446744073709551615)
                                                                    v30 = v30
                                                                    var v31 uint = v6
                                                                    _ = v31
                                                                }
                                                            }
                                                            if (18446744073709551615 >= (v2 | (v2 + (v2 ^ v2)))) {
                                                                var v32 int = fi24
                                                                v3 = append(v3, v1)
                                                                v33 := GenS5{F0: (v4 ^ v4), F1: 42}
                                                                _ = v32
                                                                _ = v33
                                                            } else {
                                                                _ = len(v3)
                                                                var v34 uint = 4294967295
                                                                v4 -= -961
                                                                _ = v34
                                                            }
                                                            var v35 string = ("a" + (("\n\t" + "a") + "foo bar"))
                                                            if (-1 == 100) {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v36 uint64 = 4294967295
                                                                    _ = v36
                                                                case "a":
                                                                    v37 := GenS4{F0: 1, F1: v1, F2: v12}
                                                                    _ = v37
                                                                case "":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v1 >> ((65535) & 63))
                                                                    }
                                                                default:
                                                                    v4 -= v4
                                                                }
                                                            } else {
                                                                v3 = append(v3, (v1 * 2))
                                                                if (true || v30) {
                                                                    v38 := GenS6{F0: 0, F1: v1, F2: (v1 - 42)}
                                                                    var v39 uint8 = 1
                                                                    if (v30 || (-1 <= 891)) {
                                                                        v6 = v6
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v1 << (((v6 - v6)) & 63))
                                                                        }
                                                                        var v40 int = 58
                                                                        _ = v40
                                                                    } else {
                                                                        v41 := GenS5{F0: (v4 << ((v6) & 63)), F1: (v39 ^ v39)}
                                                                        var v42 float32 = 1e10
                                                                        var v43 uint32 = v1
                                                                        _ = v41
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                            }
                                                            _ = fi24
                                                            _ = v30
                                                            _ = v35
                                                        }
                                                    }
                                                    switch ("foo bar" + "hello") {
                                                    case "\n\t":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "hello":
                                                        v44 := GenS5{F0: 2147483647, F1: 127}
                                                        _ = v44
                                                    default:
                                                        if false {
                                                            for fi45 := 0; fi45 < 20; fi45++ {
                                                                v46 := GenS7{F0: (fi8 & (fi45 | (fi45 << ((255) & 63)))), F1: v2, F2: v1}
                                                                _ = fi45
                                                                _ = v46
                                                            }
                                                            if (!false) {
                                                                v47 := GenS3{F0: 1, F1: (fi8 | 127)}
                                                                _ = v47
                                                            } else {
                                                                if ((v1 >= (v1 - v1)) && ((v6 - (v6 - (v6 * (v6 + v6)))) != 1)) {
                                                                    var v48 float32 = (v20 / v20)
                                                                    if (-9223372036854775808 == 432) {
                                                                        for fi49 := 0; fi49 < 20; fi49++ {
                                                                            if true {
                                                                                var v50 uint = v6
                                                                                v51 := GenS4{F0: -280, F1: (^(v1)), F2: int16(v4)}
                                                                                v52 := GenS8{F0: 242, F1: uint16(v1)}
                                                                                _ = v50
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v53 := GenS3{F0: 127, F1: (-(fi8))}
                                                                                _ = v53
                                                                            }
                                                                            var v54 int64 = -9223372036854775808
                                                                            _ = fi49
                                                                            _ = v54
                                                                        }
                                                                        v1 -= v1
                                                                        var v55 float32 = ((v20 + v20) - float32(v7))
                                                                        _ = v55
                                                                    }
                                                                    var v56 float64 = (v7 + -2.25)
                                                                    _ = v48
                                                                    _ = v56
                                                                }
                                                                v1 = 42
                                                            }
                                                            v2 += (^(v2))
                                                        } else {
                                                            for fi57 := 0; fi57 < 17; fi57++ {
                                                                if len(v3) > 0 {
                                                                    v3[0] = v1
                                                                }
                                                                fmt.Println(v12)
                                                                v20 = v20
                                                                _ = fi57
                                                            }
                                                            if false {
                                                                if ((v6 * v6) <= v6) {
                                                                    v6 += 17
                                                                }
                                                                for fi58 := 0; fi58 < 7; fi58++ {
                                                                    v20 = (v20 * (v20 / (float32((1.5 * v7)) * v20)))
                                                                    var v59 int64 = -9223372036854775808
                                                                    _ = fi58
                                                                    _ = v59
                                                                }
                                                                v60 := GenS9{F0: true, F1: fi8}
                                                                _ = v60
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                v7 += ((1e10 - v7) - float64(v20))
                                            }
                                            v1 += 4294967295
                                        }
                                        _ = v20
                                    } else {
                                        v12 += (v12 ^ 1000)
                                        switch ((("\n\t" + "") + "a") + ("hello" + (("hello" + ("foo bar" + "")) + ""))) {
                                        case "a":
                                            v4 = (v4 << ((738) & 63))
                                        case "hello":
                                            var v61 int = (-(fi8))
                                            _ = v61
                                        case "foo bar":
                                            for fi62 := 0; fi62 < 9; fi62++ {
                                                v63 := GenS5{F0: (v4 ^ 2147483647), F1: 1}
                                                v1 += 1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = fi62
                                                _ = v63
                                            }
                                        default:
                                            for fi64 := 0; fi64 < 18; fi64++ {
                                                v1 += (v1 | v1)
                                                _ = fi64
                                            }
                                        }
                                        var v65 float32 = 1e10
                                        _ = v65
                                    }
                                }
                                if ((true && false) || ((v2 + v2) > 1)) {
                                    if len(v3) > 0 {
                                        v3[0] = (v1 & (v1 + (v1 << ((v6) & 63))))
                                    }
                                    if true {
                                        if ((1.5 >= -2.25) || true) {
                                            if false {
                                                v6 = (v6 / ((uint(v2)) | 1))
                                                v66 := GenS8{F0: 0, F1: 293}
                                                _ = v66
                                            } else {
                                                fmt.Println(v2, v6)
                                            }
                                        }
                                        if false {
                                            if ((-(fi8)) <= (^(fi8))) {
                                                v3 = append(v3, v1)
                                            } else {
                                                for fi67 := 0; fi67 < 11; fi67++ {
                                                    var v68 bool = true
                                                    if len(v3) > 0 {
                                                        v3[0] = (v1 % (((^(v1))) | 1))
                                                    }
                                                    if v68 {
                                                        v68 = (true || (v6 < (v6 >> ((v6) & 63))))
                                                        v69 := GenS4{F0: -9223372036854775808, F1: (v1 - 34), F2: v12}
                                                        _ = len(v3)
                                                        _ = v69
                                                    }
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                if true {
                                                    for fi70 := 0; fi70 < 15; fi70++ {
                                                        if (!(42 <= 1)) {
                                                            v7 -= float64(0.1)
                                                            if (!false) {
                                                                var v71 uint8 = 255
                                                                var v72 int32 = -2147483648
                                                                fmt.Println(v6)
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                        }
                                                        fmt.Println(v4, fi70)
                                                        v4 += (v4 / ((v4) | 1))
                                                        _ = fi70
                                                    }
                                                    fmt.Println(v6, fi8)
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                } else {
                                                    v3 = append(v3, (v1 << ((v6) & 63)))
                                                    for fi73 := 0; fi73 < 2; fi73++ {
                                                        if (9223372036854775807 >= 9223372036854775807) {
                                                            if true {
                                                                var v74 int8 = int8(v6)
                                                                _ = v74
                                                            } else {
                                                                var v75 uint64 = (v2 + v2)
                                                                _ = v75
                                                            }
                                                            v2 -= (v2 / ((0) | 1))
                                                            v6 -= v6
                                                            var v76 int16 = -1
                                                            var v77 int32 = v4
                                                            v6 += v6
                                                            v1 = v1
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        _ = fi73
                                                    }
                                                }
                                            }
                                            v6 -= 255
                                            v1 = v1
                                        } else {
                                            var v78 uint8 = 214
                                            for fi79 := 0; fi79 < 2; fi79++ {
                                                v80 := GenS10{F0: (v4 % ((v4) | 1)), F1: 9223372036854775807}
                                                if false {
                                                    v81 := GenS6{F0: -12, F1: v1, F2: v1}
                                                    _ = v81
                                                } else {
                                                    var v82 int16 = (v12 << (((v6 - 42)) & 63))
                                                    _ = v82
                                                }
                                                var v83 uint8 = (v78 | (v78 >> (((^(v6))) & 63)))
                                                _ = fi79
                                                _ = v80
                                                _ = v83
                                            }
                                            _ = v78
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v84 int32 = v4
                                    _ = v84
                                } else {
                                    switch "hello" {
                                    case "":
                                        v2 -= (v2 & v2)
                                    case "foo bar":
                                        v1 = v1
                                    case "hello":
                                        var v85 uint8 = 128
                                        _ = v85
                                    }
                                }
                                fmt.Println(v12)
                            } else {
                                var v86 int = (fi8 % ((fi8) | 1))
                                var v87 int = 32767
                                v12 = (v12 ^ (v12 | (v12 & v12)))
                                _ = v86
                                _ = v87
                            }
                        default:
                            v2 += (v2 + 663)
                        }
                        v12 = v12
                        _ = v12
                    }
                case "\n\t":
                    v88 := GenS4{F0: 9223372036854775807, F1: (v1 & 2147483648), F2: -552}
                    _ = v88
                }
            }
            _ = fi8
        }
        _ = v6
        _ = v7
    } else {
        switch "\n\t" {
        case "":
            switch -842 {
            case 128:
                if (true && (!(42 > 950))) {
                    if true {
                        fmt.Println(v1)
                        var v89 int32 = 353
                        _ = v89
                    }
                }
            case -486:
                var v90 float64 = (0.0 + ((1e10 / 1.0) * (1.5 + 1e10)))
                _ = v90
            }
        case "a":
            fmt.Println(v4, v4, v1)
        }
    }
    fmt.Println(v4)
    v2 = (v2 * (v2 % (((v2 ^ (v2 >> ((603) & 63)))) | 1)))
    if ((v1 * v1) >= 0) {
        if true {
            v3 = append(v3, v1)
            fmt.Println(v4, v1, v2)
        }
        if (!(-90 < (v4 * v4))) {
            v2 = (v2 << ((4294967295) & 63))
            if len(v3) > 0 {
                v3[0] = uint32(v2)
            }
        } else {
            for fi91 := 0; fi91 < 20; fi91++ {
                if (!(9223372036854775807 != 472)) {
                    v4 = (v4 * (v4 / ((v4) | 1)))
                    fmt.Println(v4, v1, v1)
                    fmt.Println(v2)
                } else {
                    var v92 bool = true
                    _ = v92
                }
                v1 = 2147483647
                v1 -= v1
                _ = fi91
            }
            var v93 uint8 = 58
            _ = v93
        }
    }
    v94 := GenS11{F0: 113, F1: "a", F2: (v1 << ((42) & 63))}
    var v95 uint = 59
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v94
    _ = v95
}
