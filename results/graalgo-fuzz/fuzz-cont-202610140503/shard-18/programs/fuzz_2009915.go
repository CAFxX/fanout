package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int64; F1 int8; F2 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int8; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.1
    var v3 int = 259
    var v4 uint32 = 65535
    var v5 uint32 = v4
    var v6 int8 = 107
    v4 -= v5
    switch int(v5) {
    case -151:
        fmt.Println(v6)
    case -2147483648:
        for fi7 := 0; fi7 < 19; fi7++ {
            v2 += v2
            if true {
                v6 += 42
            } else {
                v3 += v3
                if ("\n\t" != ("\n\t" + ("hello" + "\n\t"))) {
                    if true {
                        var v8 uint8 = 170
                        var v9 int32 = 42
                        v10 := GenS1{F0: 18446744073709551615}
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                }
            }
            v1[581] = (("foo bar" >= ("\n\t" + ("hello" + ("hello" + ("\n\t" + ""))))) || false)
            _ = fi7
        }
    case -9223372036854775808:
        v4 = v4
    }
    switch v3 {
    case 1:
        v5 -= v5
    case 127:
        switch v3 {
        case -597:
            for fi11 := 0; fi11 < 8; fi11++ {
                if ((9223372036854775807 < -900) || false) {
                    switch ("\n\t" + "hello") {
                    case "":
                        v12 := []int{}
                        _ = v12
                    case "\n\t":
                        var v13 int8 = -64
                        _ = v13
                    case "foo bar":
                        fmt.Println(v4, v6, v3)
                    }
                } else {
                    fmt.Println(v4, v5)
                }
                _ = fi11
            }
        case 127:
            var v14 int32 = 1
            _ = v14
        case -374:
            v15 := GenS1{F0: 651}
            _ = v15
        default:
            v5 = 774
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int8 = 127
        v3 -= -655
        v18 := GenS2{F0: -1, F1: v5}
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if (!false) {
            if (2 != 209) {
                switch "hello" {
                case "":
                    for fi20 := 0; fi20 < 13; fi20++ {
                        if (v4 != (v4 & (v5 & 2))) {
                            _ = len(v1)
                        } else {
                            var v21 bool = false
                            if (v21 || v21) {
                                v22 := GenS1{F0: 4294967295}
                                fmt.Println(fi19, v21)
                                v1[(v5 + v5)] = v21
                                fmt.Println(v6, v2)
                                v23 := GenS2{F0: 2147483647, F1: (v5 ^ 315)}
                                _ = v22
                                _ = v23
                            } else {
                                v5 = (^(v5))
                                var v24 string = ("\n\t" + ("" + "a"))
                                v25 := []int{}
                                _ = v24
                                _ = v25
                            }
                            _ = v21
                        }
                        _ = fi20
                    }
                default:
                    var v26 uint8 = 109
                    _ = v26
                }
            }
            switch "\n\t" {
            case "a":
                v5 += 4294967295
            case "\n\t":
                v2 = (float32(3.14159) + (v2 + v2))
            case "foo bar":
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint = 4294967295
                    _ = fi27
                    _ = v28
                }
            }
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            switch fi19 {
            case 0:
                v6 += 127
            case 127:
                var v30 uint8 = 255
                _ = v30
            case -454:
                var v31 uint32 = 259
                _ = v31
            default:
                fmt.Println(v5)
            }
            if false {
                var v32 uint32 = 952
                v3 -= (fi19 / ((540) | 1))
                _ = v32
            }
            fmt.Println(v4, v5)
            _ = fi29
        }
        _ = fi19
    }
    if false {
        if true {
            var v33 uint16 = 440
            v4 = 221
            v2 += (1.5 - 0.1)
            _ = v33
        } else {
            v2 = v2
            v6 += (v6 >> ((2) & 63))
            switch (("a" + "\n\t") + "") {
            case "a":
                for fi34 := 0; fi34 < 7; fi34++ {
                    if true {
                        v2 = float32(3.14159)
                        fmt.Println(v4, v6, v6)
                        switch ("foo bar" + ("foo bar" + "hello")) {
                        case "foo bar":
                            v6 = 127
                        case "hello":
                            var v35 int8 = v6
                            _ = v35
                        }
                    }
                    v36 := GenS3{F0: 699, F1: (fi34 / ((2147483647) | 1)), F2: 65535}
                    v37 := GenS3{F0: 2147483647, F1: v3, F2: 1}
                    _ = fi34
                    _ = v36
                    _ = v37
                }
            case "":
                fmt.Println(v2)
            case "hello":
                for fi38 := 0; fi38 < 19; fi38++ {
                    v39 := GenS2{F0: -261, F1: v4}
                    for fi40 := 0; fi40 < 20; fi40++ {
                        var v41 int64 = 0
                        v42 := []int{}
                        v43 := GenS2{F0: 977, F1: (v5 / (((^(v5))) | 1))}
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    v44 := []int{-9223372036854775808, (fi38 - -2147483648), fi38, (-(v3))}
                    _ = fi38
                    _ = v39
                    _ = v44
                }
            default:
                switch "\n\t" {
                case "\n\t":
                    if true {
                        v3 = 716
                    } else {
                        v45 := GenS4{F0: 42, F1: 154}
                        var v46 float64 = (((-2.25 - 0.1) + (-2.25 - ((0.1 * 1e10) - 0.1))) / 1.5)
                        _ = v45
                        _ = v46
                    }
                case "hello":
                    if true {
                        fmt.Println(v4, v6, v4)
                        fmt.Println(v2)
                    }
                case "foo bar":
                    if true {
                        var v47 float64 = (-2.25 / (3.14159 / 1e10))
                        delete(v1, 2147483648)
                        _ = v47
                    } else {
                        if (!((!true) || false)) {
                            var v48 int32 = -1
                            _ = v48
                        } else {
                            switch v3 {
                            case -9223372036854775808:
                                for fi49 := 0; fi49 < 18; fi49++ {
                                    var v50 int8 = v6
                                    if (1.5 > -2.25) {
                                        if (275 != 441) {
                                            for fi51 := 0; fi51 < 8; fi51++ {
                                                var v52 int8 = v50
                                                v53 := GenS2{F0: -1, F1: v5}
                                                v52 += v6
                                                v2 += float32(float64(0.1))
                                                _ = fi51
                                                _ = v52
                                                _ = v53
                                            }
                                            var v54 uint8 = 0
                                            _ = v54
                                        } else {
                                            v6 += v6
                                            v4 = (v4 * 1)
                                        }
                                        v4 = 42
                                    } else {
                                        var v55 bool = ((float64((v2 + 1.5)) + ((((0.0 / (float64((v2 * (v2 / (((1.5 * v2) * (0.1 / (float32(0.1) * v2))) - ((float32(0.1) - v2) - v2))))) / -2.25)) / float64((v2 + (v2 + float32((-2.25 - 3.14159)))))) - ((-2.25 * (((float64((((v2 - v2) - v2) - v2)) - 3.14159) * 0.1) + 1e10)) + 3.14159)) + float64(v2))) > (0.1 - (float64(((float32(0.1) / v2) - float32(1.5))) / 1.0)))
                                        v56 := GenS3{F0: 0, F1: 128, F2: 222}
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 int8 = v6
                                    _ = fi49
                                    _ = v50
                                    _ = v57
                                }
                            case 42:
                                if (!(122 > 878)) {
                                    for fi58 := 0; fi58 < 18; fi58++ {
                                        if (v5 >= v4) {
                                            v1[(v5 & (v4 & 65535))] = (3.14159 < float64((v2 / 1.5)))
                                        }
                                        _ = fi58
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        for fi59 := 0; fi59 < 8; fi59++ {
                                            var v60 uint16 = 296
                                            delete(v1, (v5 & (v4 * (v4 / ((v4) | 1)))))
                                            for fi61 := 0; fi61 < 4; fi61++ {
                                                fmt.Println(v60, v4)
                                                _ = fi61
                                            }
                                            _ = fi59
                                            _ = v60
                                        }
                                    default:
                                        fmt.Println(v5, v4)
                                    }
                                } else {
                                    var v62 string = "hello"
                                    v63 := GenS2{F0: 42, F1: 42}
                                    _ = v62
                                    _ = v63
                                }
                            default:
                                v4 = (v5 & (v5 + v4))
                            }
                            if ("a" != ("hello" + "a")) {
                                fmt.Println(v4, v5)
                                if (v2 != (v2 * (((3.14159 / ((3.14159 + v2) / v2)) - v2) * (((v2 / 3.14159) - ((0.1 - v2) * 0.0)) * v2)))) {
                                    v1[(v4 + (v4 + v4))] = (true && (-2147483648 >= v3))
                                    for fi64 := 0; fi64 < 16; fi64++ {
                                        var v65 int32 = -2147483648
                                        fmt.Println(v6, fi64)
                                        v4 = 340
                                        _ = fi64
                                        _ = v65
                                    }
                                    if (int32(v4) != 0) {
                                        v66 := v1[(v5 * (v5 >> ((1) & 63)))]
                                        for fi67 := 0; fi67 < 6; fi67++ {
                                            v68 := GenS5{F0: 42, F1: 100, F2: 42}
                                            _ = fi67
                                            _ = v68
                                        }
                                        _ = v66
                                    } else {
                                        var v69 int64 = 1
                                        _ = v69
                                    }
                                } else {
                                    fmt.Println(v3, v6)
                                    _ = len(v1)
                                }
                            }
                            fmt.Println(v5)
                        }
                        v70 := GenS1{F0: 4294967295}
                        if (-32768 == 0) {
                            if ((!true) && (1 == 0)) {
                                switch ("a" + "\n\t") {
                                case "a":
                                    for fi71 := 0; fi71 < 17; fi71++ {
                                        var v72 float32 = 3.14159
                                        for fi73 := 0; fi73 < 12; fi73++ {
                                            if ((false || (!true)) || (((fi71 * (v3 & -1)) == (v3 << ((4294967295) & 63))) || false)) {
                                                if ((42 > 1) && false) {
                                                    fmt.Println(v3)
                                                    v3 -= (^(v3))
                                                }
                                                switch ("a" + "a") {
                                                case "":
                                                    v4 = 2147483648
                                                case "hello":
                                                    if (185 < 340) {
                                                        v2 -= 0.0
                                                        v3 = fi73
                                                        v74 := GenS1{F0: 18446744073709551615}
                                                        _ = v74
                                                    } else {
                                                        var v75 float64 = 1e10
                                                        _ = v75
                                                    }
                                                case "foo bar":
                                                    v76 := GenS6{F0: 42}
                                                    if (true || (!(!((v3 ^ v3) != (fi71 & (fi73 + 4294967295)))))) {
                                                        v77 := []int{-2147483648, v3, v3, (v3 | (fi71 / (((-(fi71))) | 1)))}
                                                        if false {
                                                            v78 := GenS2{F0: 1, F1: v5}
                                                            v79 := GenS7{F0: 4294967295, F1: v6, F2: -2147483648}
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            _ = len(v1)
                                                            if (false || (1 >= 966)) {
                                                                v77 = append(v77, (v3 >> ((2) & 63)))
                                                                var v80 int = (fi71 - fi71)
                                                                _ = v80
                                                            }
                                                            if false {
                                                                v2 += (float32(1e10) / float32(-2.25))
                                                            } else {
                                                                if (42 <= 42) {
                                                                    var v81 float64 = 3.14159
                                                                    v2 = float32(v81)
                                                                    _ = len(v77)
                                                                    _ = v81
                                                                } else {
                                                                    fmt.Println(v3, v6, fi71)
                                                                    _ = len(v77)
                                                                    v4 = v5
                                                                }
                                                                v2 = (v2 - v2)
                                                                if len(v77) > 0 {
                                                                    v77[0] = -2147483648
                                                                }
                                                            }
                                                        }
                                                        _ = v77
                                                    }
                                                    _ = v76
                                                default:
                                                    if false {
                                                        fmt.Println(v3)
                                                        var v82 uint = 255
                                                        _ = v82
                                                    }
                                                }
                                            } else {
                                                v83 := []int{}
                                                if true {
                                                    v84 := GenS8{F0: 32767}
                                                    _ = v84
                                                }
                                                _ = v83
                                            }
                                            _ = fi73
                                        }
                                        v85 := GenS9{F0: 42}
                                        _ = fi71
                                        _ = v72
                                        _ = v85
                                    }
                                case "hello":
                                    var v86 uint64 = 4294967295
                                    _ = v86
                                }
                            } else {
                                fmt.Println(v5)
                                var v87 int8 = v6
                                _ = v87
                            }
                            if true {
                                if false {
                                    fmt.Println(v5)
                                } else {
                                    _ = len(v1)
                                    v88 := v1[v4]
                                    fmt.Println(v2)
                                    _ = v88
                                }
                            } else {
                                for fi89 := 0; fi89 < 12; fi89++ {
                                    v4 -= (v5 | v4)
                                    for fi90 := 0; fi90 < 1; fi90++ {
                                        var v91 uint16 = 42
                                        v5 -= v5
                                        var v92 int16 = -1
                                        _ = fi90
                                        _ = v91
                                        _ = v92
                                    }
                                    _ = fi89
                                }
                            }
                            delete(v1, v4)
                        }
                        _ = v70
                    }
                }
            }
        }
        switch (v3 - (v3 | v3)) {
        case 2147483647:
            var v93 bool = ((true || false) || true)
            _ = v93
        case -9223372036854775808:
            v94 := GenS10{F0: 67, F1: 65535, F2: 2}
            _ = v94
        }
        v6 = (-(v6))
    }
    if false {
        v6 = v6
    } else {
        v6 = (v6 | v6)
        if false {
            if ("" == "hello") {
                var v95 int64 = 9223372036854775807
                switch "a" {
                case "":
                    v4 -= v4
                case "a":
                    v3 -= 352
                case "foo bar":
                    var v96 uint32 = 42
                    _ = v96
                default:
                    v97 := GenS4{F0: 42, F1: 42}
                    _ = v97
                }
                v4 = (v4 + v5)
                _ = v95
            }
        }
    }
    fmt.Println(v4)
    var v98 uint8 = 42
    fmt.Println(v6, v5, v4)
    if false {
        v4 += (v4 / ((734) | 1))
        delete(v1, 145)
        v98 += v98
    } else {
        v98 = v98
        _ = len(v1)
        v99 := GenS3{F0: -28, F1: v3, F2: 410}
        _ = v99
    }
    var v100 int = (v3 | v3)
    v101 := []int{v100, v3, (v100 / (((v3 * (v3 / ((v100) | 1)))) | 1))}
    var v102 uint8 = (v98 / ((v98) | 1))
    v6 = (v6 & (^(v6)))
    var v103 int16 = 42
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
