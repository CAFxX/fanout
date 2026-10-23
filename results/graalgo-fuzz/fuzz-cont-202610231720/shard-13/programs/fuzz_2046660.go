package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int8; F2 int }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{2147483647, 878, 465, 65535}
    var v3 uint16 = 0
    var v4 int = -921
    v4 = v4
    var v5 uint8 = 183
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            if ((1000 >= 42) || false) {
                var v8 int16 = -32768
                _ = v8
            } else {
                switch "\n\t" {
                case "foo bar":
                    v4 += fi6
                case "a":
                    _ = len(v1)
                case "":
                    var v9 int = fi6
                    _ = v9
                default:
                    var v10 uint32 = 847
                    _ = v10
                }
                v11 := GenS1{F0: 0, F1: 594, F2: 42}
                var v12 int64 = 1
                _ = v11
                _ = v12
            }
            v4 -= fi7
            v13 := GenS1{F0: 2, F1: 0, F2: 9223372036854775807}
            _ = fi7
            _ = v13
        }
        var v14 uint32 = 68
        _ = fi6
        _ = v14
    }
    v4 += v4
    var v15 int32 = 293
    v4 = (v4 * 9223372036854775807)
    switch v4 {
    case 4294967295:
        if (!false) {
            var v16 float64 = float64(((1.5 / (-2.25 * (float32(1e10) / (1e10 / 1.0)))) + 0.0))
            var v17 int8 = 1
            _ = v16
            _ = v17
        }
    case 199:
        if false {
            for fi18 := 0; fi18 < 20; fi18++ {
                v15 = v15
                v19 := GenS1{F0: 911, F1: uint(v5), F2: 600}
                v4 -= (fi18 << ((964) & 63))
                _ = fi18
                _ = v19
            }
        }
    case 9223372036854775807:
        if true {
            for fi20 := 0; fi20 < 16; fi20++ {
                for fi21 := 0; fi21 < 17; fi21++ {
                    delete(v1, 238)
                    _ = fi21
                }
                fmt.Println(v15, v4)
                _ = fi20
            }
            var v22 int = v4
            _ = v22
        } else {
            switch v4 {
            case 1:
                fmt.Println(v15, v3, v5)
            case 994:
                if (v5 == 255) {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 int16 = 42
                        _ = len(v1)
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 17; fi25++ {
                        for fi26 := 0; fi26 < 15; fi26++ {
                            if false {
                                fmt.Println(fi26)
                                v5 -= v5
                                v27 := GenS1{F0: 42, F1: 4294967295, F2: -1}
                                _ = v27
                            }
                            switch ("a" + "hello") {
                            case "":
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    fmt.Println(v5)
                                default:
                                    v28 := v1[177]
                                    _ = v28
                                }
                                v4 = fi26
                                if false {
                                    var v29 uint64 = 110
                                    var v30 float64 = (((1.5 + 0.0) / 0.1) - float64(0.1))
                                    fmt.Println(v15)
                                    var v31 int = fi26
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            case "a":
                                fmt.Println(v3)
                            case "\n\t":
                                var v32 uint8 = (v5 & (v5 | (v5 & v5)))
                                _ = v32
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    v33 := GenS1{F0: 804, F1: 941, F2: -48}
                                    _ = v33
                                case "":
                                    _ = len(v1)
                                }
                            }
                            v3 = v3
                            var v34 uint = 375
                            v2 = append(v2, 521)
                            _ = fi26
                            _ = v34
                        }
                        var v35 int = (fi25 - 1)
                        if (false || true) {
                            _ = len(v1)
                            if (!(60 > 127)) {
                                for fi36 := 0; fi36 < 9; fi36++ {
                                    var v37 float64 = -2.25
                                    switch (v35 << ((255) & 63)) {
                                    case 2147483647:
                                        switch "\n\t" {
                                        case "":
                                            var v38 uint16 = (v3 ^ v3)
                                            _ = v38
                                        }
                                    case 9223372036854775807:
                                        var v39 uint32 = 65535
                                        _ = v39
                                    case 4294967295:
                                        var v40 bool = (("hello" + "hello") <= "hello")
                                        _ = v40
                                    default:
                                        v35 += v35
                                    }
                                    _ = fi36
                                    _ = v37
                                }
                            } else {
                                fmt.Println(v4, v4, v5)
                                v3 = 1
                                var v41 int16 = 32767
                                _ = v41
                            }
                        }
                        _ = fi25
                        _ = v35
                    }
                    if ((true || (65535 >= 2147483647)) && false) {
                        var v42 uint = 4294967295
                        v42 += v42
                        _ = v42
                    } else {
                        v4 += v4
                        var v43 string = ("" + ("\n\t" + "\n\t"))
                        if true {
                            if true {
                                if (!true) {
                                    fmt.Println(v43, v4, v5)
                                } else {
                                    if (7 >= v5) {
                                        v44 := GenS1{F0: 0, F1: 688, F2: int64(v4)}
                                        for fi45 := 0; fi45 < 20; fi45++ {
                                            v3 = 32767
                                            _ = fi45
                                        }
                                        v1[4294967295] = (v43 != v43)
                                        _ = v44
                                    }
                                    switch v4 {
                                    case 128:
                                        if true {
                                            var v46 int = v4
                                            _ = v46
                                        }
                                    case -9223372036854775808:
                                        var v47 bool = false
                                        _ = v47
                                    default:
                                        for fi48 := 0; fi48 < 7; fi48++ {
                                            var v49 string = (((v43 + v43) + v43) + "hello")
                                            if false {
                                                v43 = v43
                                            } else {
                                                v50 := GenS2{F0: -905}
                                                fmt.Println(v3, fi48, v3)
                                                _ = v50
                                            }
                                            for fi51 := 0; fi51 < 8; fi51++ {
                                                if (v5 < (v5 % ((36) | 1))) {
                                                    var v52 uint = 42
                                                    v49 = "foo bar"
                                                    _ = v52
                                                }
                                                fmt.Println(v4, v5)
                                                _ = fi51
                                            }
                                            _ = fi48
                                            _ = v49
                                        }
                                    }
                                }
                            }
                            switch v4 {
                            case 791:
                                if ((0.1 + (((-2.25 + 0.1) * 0.1) / 1.0)) >= -2.25) {
                                    if (42 != 479) {
                                        if (false || (!false)) {
                                            for fi53 := 0; fi53 < 1; fi53++ {
                                                _ = len(v1)
                                                v43 = v43
                                                _ = fi53
                                            }
                                            v54 := GenS1{F0: 0, F1: 65535, F2: 923}
                                            for fi55 := 0; fi55 < 17; fi55++ {
                                                fmt.Println(v5, v5, v43)
                                                _ = fi55
                                            }
                                            _ = v54
                                        }
                                    }
                                    v56 := v1[2147483648]
                                    v57 := GenS1{F0: 255, F1: 1, F2: 1}
                                    _ = v56
                                    _ = v57
                                }
                            case -2147483648:
                                v4 = -534
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    for fi58 := 0; fi58 < 10; fi58++ {
                                        switch v43 {
                                        case "hello":
                                            for fi59 := 0; fi59 < 5; fi59++ {
                                                if len(v2) > 0 {
                                                    v2[0] = 985
                                                }
                                                v3 += (v3 ^ (v3 ^ v3))
                                                fmt.Println(v5, v43, v4)
                                                _ = fi59
                                            }
                                        case "":
                                            fmt.Println(v43, v3, v4)
                                        default:
                                            var v60 int16 = -32768
                                            _ = v60
                                        }
                                        fmt.Println(fi58)
                                        _ = fi58
                                    }
                                case "hello":
                                    for fi61 := 0; fi61 < 4; fi61++ {
                                        v3 = 65535
                                        _ = fi61
                                    }
                                }
                            }
                            v62 := GenS3{F0: 42, F1: (v4 ^ 32767), F2: (v15 ^ v15)}
                            _ = v62
                        } else {
                            if (v15 == (v15 >> ((4294967295) & 63))) {
                                v4 += v4
                                fmt.Println(v43, v43, v5)
                            } else {
                                v4 += v4
                                var v63 int = 128
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 9; fi64++ {
                                var v65 int64 = int64(v15)
                                v66 := GenS2{F0: -516}
                                for fi67 := 0; fi67 < 20; fi67++ {
                                    v15 -= v15
                                    if true {
                                        if (true || (-1 <= fi64)) {
                                            v68 := GenS4{F0: -1, F1: 720, F2: -961}
                                            if true {
                                                var v69 int8 = -1
                                                if true {
                                                    var v70 string = "foo bar"
                                                    fmt.Println(v15)
                                                    v1[65535] = false
                                                    _ = v70
                                                }
                                                _ = v69
                                            }
                                            var v71 int8 = -128
                                            _ = v68
                                            _ = v71
                                        } else {
                                            delete(v1, uint32(v3))
                                            var v72 int64 = 390
                                            v73 := GenS5{F0: v15}
                                            var v74 uint16 = (v3 ^ 32767)
                                            v75 := GenS6{F0: (fi64 > fi67), F1: 0, F2: v4}
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                    fmt.Println(v65, v5, v65)
                                    v76 := GenS5{F0: (v15 & v15)}
                                    _ = fi67
                                    _ = v76
                                }
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        }
                        _ = v43
                    }
                } else {
                    v3 = v3
                }
            default:
                if true {
                    v15 -= v15
                    fmt.Println(v15)
                }
            }
        }
    default:
        if false {
            switch ("foo bar" + "foo bar") {
            case "a":
                v15 = 2147483647
            case "":
                fmt.Println(v4, v15)
            }
            for fi77 := 0; fi77 < 6; fi77++ {
                var v78 uint16 = (v3 / ((v3) | 1))
                _ = fi77
                _ = v78
            }
        }
    }
    fmt.Println(v15)
    v2 = append(v2, 0)
    fmt.Println(v3, v4)
    var v79 int8 = 42
    v5 = 83
    if ((0.0 - (((0.0 / 1e10) - -2.25) * -2.25)) <= 1.5) {
        for fi80 := 0; fi80 < 9; fi80++ {
            v81 := GenS1{F0: 4294967295, F1: 756, F2: -1}
            _ = fi80
            _ = v81
        }
    } else {
        if false {
            v5 = 127
            v5 = 127
        }
        for fi82 := 0; fi82 < 14; fi82++ {
            v1[2] = ((v5 == v5) || (338 >= -9223372036854775808))
            switch (fi82 % ((4) | 1)) {
            case 4294967295:
                for fi83 := 0; fi83 < 19; fi83++ {
                    v84 := GenS7{F0: 864, F1: true}
                    v3 -= (v3 | (v3 * 2))
                    v85 := GenS7{F0: 711, F1: (false || false)}
                    _ = fi83
                    _ = v84
                    _ = v85
                }
            case 429:
                v4 = fi82
            default:
                for fi86 := 0; fi86 < 18; fi86++ {
                    v79 = 127
                    var v87 uint = 42
                    v1[0] = ((^(fi82)) <= (-(fi82)))
                    _ = fi86
                    _ = v87
                }
            }
            _ = fi82
        }
        v88 := v1[2]
        _ = v88
    }
    if ((v79 / ((v79) | 1)) != 42) {
        v5 = v5
    }
    switch 782 {
    case 2147483647:
        var v89 bool = (674 == 65535)
        _ = v89
    case -2147483648:
        if (42 < 963) {
            v3 = 65535
            for fi90 := 0; fi90 < 12; fi90++ {
                fmt.Println(v3, v3, v15)
                v4 = (fi90 & int(v79))
                _ = fi90
            }
            v2 = append(v2, 42)
        } else {
            fmt.Println(v15)
            switch "hello" {
            case "a":
                for fi91 := 0; fi91 < 5; fi91++ {
                    v92 := GenS2{F0: 32767}
                    var v93 int16 = -32768
                    _ = fi91
                    _ = v92
                    _ = v93
                }
            case "foo bar":
                if true {
                    var v94 uint8 = v5
                    v4 = (v4 * int(v3))
                    for fi95 := 0; fi95 < 1; fi95++ {
                        var v96 float32 = ((3.14159 * 3.14159) / (1e10 / ((((-2.25 + 0.0) + (0.1 * -2.25)) * -2.25) / (((0.0 - 1e10) * 1.5) - (0.0 + 0.1)))))
                        v97 := GenS8{F0: fi95}
                        _ = fi95
                        _ = v96
                        _ = v97
                    }
                    _ = v94
                } else {
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                }
            default:
                for fi98 := 0; fi98 < 14; fi98++ {
                    if true {
                        v4 += fi98
                        switch "\n\t" {
                        case "":
                            var v99 float64 = (((-2.25 + (float64(float32(((3.14159 - (-2.25 - 3.14159)) + float64((0.1 + (1.5 - ((float32(1.5) * (float32(1.5) + 0.0)) * 1e10))))))) / (0.1 * 1.5))) - 0.1) * 0.1)
                            _ = v99
                        case "a":
                            for fi100 := 0; fi100 < 20; fi100++ {
                                v2 = append(v2, 337)
                                fmt.Println(fi98, v15)
                                _ = fi100
                            }
                        default:
                            var v101 int = (v4 * fi98)
                            _ = v101
                        }
                    } else {
                        v79 = (v79 << ((18446744073709551615) & 63))
                        v1[1] = (0 > -32768)
                        var v102 int32 = v15
                        _ = v102
                    }
                    v3 -= v3
                    for fi103 := 0; fi103 < 3; fi103++ {
                        v104 := GenS9{F0: 1, F1: (v5 / ((v5) | 1))}
                        if false {
                            var v105 string = "\n\t"
                            var v106 int16 = -1
                            _ = v105
                            _ = v106
                        }
                        v107 := GenS10{F0: 4294967295}
                        _ = fi103
                        _ = v104
                        _ = v107
                    }
                    _ = fi98
                }
            }
            if false {
                v79 -= 100
            }
        }
    case 409:
        if (!((42 < 0) && false)) {
            fmt.Println(v5, v3)
            v108 := GenS8{F0: 0}
            fmt.Println(v15, v79)
            _ = v108
        }
    }
    var v109 float32 = ((0.0 / 1e10) * (((((float32(((0.0 + 3.14159) * 1.5)) / 1.0) - 1e10) / (-2.25 - 1e10)) - (0.0 - 0.0)) + 1.5))
    var v110 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v109
    _ = v110
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
}
