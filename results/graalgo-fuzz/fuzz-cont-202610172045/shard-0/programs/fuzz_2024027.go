package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int; F1 uint }
func genFn0() uint16 {
    var v7 int8 = -1
    switch -1 {
    case -2147483648:
        for fi8 := 0; fi8 < 1; fi8++ {
            if ((fi8 % (((fi8 << ((18446744073709551615) & 63))) | 1)) < (-(fi8))) {
                for fi9 := 0; fi9 < 10; fi9++ {
                    v7 = (v7 << ((989) & 63))
                    v10 := GenS1{F0: "", F1: (fi8 - fi8)}
                    _ = fi9
                    _ = v10
                }
            } else {
                var v11 uint64 = 42
                _ = v11
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 bool = false
                v14 := []int{(fi8 >> ((1) & 63)), (-(fi8)), fi12, (fi8 % ((32768) | 1))}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
    default:
        for fi15 := 0; fi15 < 15; fi15++ {
            if false {
                var v16 bool = (fi15 >= fi15)
                for fi17 := 0; fi17 < 14; fi17++ {
                    v7 -= v7
                    var v18 int32 = 0
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = fi15
        }
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: ("foo bar" + "a"), F1: 127}
        switch "\n\t" {
        case "":
            var v21 int = fi19
            _ = v21
        case "a":
            v7 = v7
        case "foo bar":
            switch "" {
            case "hello":
                v22 := []int{(^(fi19)), (fi19 + fi19)}
                _ = v22
            case "foo bar":
                var v23 string = "\n\t"
                _ = v23
            case "\n\t":
                var v24 int = (fi19 << ((2) & 63))
                _ = v24
            default:
                if false {
                    if (65535 <= 2) {
                        if ((fi19 & fi19) >= -325) {
                            var v25 uint = 374
                            _ = v25
                        }
                    } else {
                        if ("foo bar" > (((("a" + "\n\t") + "foo bar") + ((("\n\t" + "a") + ("" + "hello")) + ("\n\t" + "\n\t"))) + "hello")) {
                            for fi26 := 0; fi26 < 18; fi26++ {
                                v7 = (v7 / ((v7) | 1))
                                v27 := []int{(fi26 & (fi19 + fi19)), 0}
                                _ = fi26
                                _ = v27
                            }
                            v28 := []int{fi19, 127}
                            var v29 uint64 = 144
                            _ = v28
                            _ = v29
                        } else {
                            if ("a" < "a") {
                                fmt.Println(fi19, fi19)
                                v30 := []int{9223372036854775807, (fi19 >> ((0) & 63))}
                                v7 += (v7 ^ -128)
                                _ = v30
                            } else {
                                v31 := []int{660, (fi19 / (((fi19 % (((fi19 - (^(fi19)))) | 1))) | 1)), fi19}
                                _ = v31
                            }
                            v32 := GenS2{F0: (v7 + v7), F1: 1}
                            _ = v32
                        }
                        var v33 int32 = -1
                        v7 = (v7 * v7)
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 20; fi34++ {
                        v35 := GenS1{F0: ("\n\t" + "hello"), F1: 2147483647}
                        _ = fi34
                        _ = v35
                    }
                    switch "a" {
                    case "foo bar":
                        var v36 float64 = (((1.5 / 3.14159) - ((1e10 - 0.0) / 3.14159)) / -2.25)
                        _ = v36
                    case "a":
                        fmt.Println(v7, fi19)
                    default:
                        v37 := []int{fi19, fi19}
                        _ = v37
                    }
                } else {
                    var v38 uint8 = 42
                    v7 = (v7 - 100)
                    v39 := GenS1{F0: "hello", F1: fi19}
                    _ = v38
                    _ = v39
                }
            }
        default:
            for fi40 := 0; fi40 < 1; fi40++ {
                v7 -= -105
                _ = fi40
            }
        }
        fmt.Println(v7, v7)
        _ = fi19
        _ = v20
    }
    if (-1 > -1) {
        v41 := []int{-732}
        _ = v41
    } else {
        for fi42 := 0; fi42 < 11; fi42++ {
            v7 -= 100
            v43 := GenS3{F0: 701, F1: 451, F2: 42}
            _ = fi42
            _ = v43
        }
        var v44 int16 = 1000
        _ = v44
    }
    v45 := GenS2{F0: 100, F1: 65535}
    _ = v45
    _ = v7
    return 876
}

func main() {
    v1 := []int{-773, 128, 300}
    var v2 int = -2147483648
    v3 := []uint32{722, 1, 196}
    v4 := make(map[uint32]uint32)
    var v5 int = (v2 / (((-(v2))) | 1))
    v6 := []string{"hello", "hello", "\n\t"}
    var v46 int8 = -2
    for fi47 := 0; fi47 < 9; fi47++ {
        if (-1 != 439) {
            v2 += (v5 + (v5 / ((-9223372036854775808) | 1)))
            var v48 uint32 = 421
            v49 := GenS3{F0: 1, F1: 0, F2: 130}
            _ = v48
            _ = v49
        }
        fmt.Println(v2)
        _ = fi47
    }
    v50 := GenS2{F0: v46, F1: 1}
    fmt.Println(v46)
    v46 = 42
    v5 = (v5 / (((v5 + (^(v5)))) | 1))
    v46 += (v46 - v46)
    for fi51 := 0; fi51 < 14; fi51++ {
        for fi52 := 0; fi52 < 3; fi52++ {
            if true {
                if true {
                    _ = len(v4)
                } else {
                    v53 := GenS4{F0: 0}
                    if (true || false) {
                        if (1 < (fi51 << ((255) & 63))) {
                            v54 := GenS3{F0: 42, F1: 2, F2: 0}
                            if true {
                                v5 = v5
                                fmt.Println(v2, v5, v2)
                                fmt.Println(fi51, v2, fi52)
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v54
                        }
                        if (false || (32767 <= 0)) {
                            if true {
                                fmt.Println(fi51)
                                v2 += (fi51 | fi52)
                                if len(v6) > 0 {
                                    v6[0] = "a"
                                }
                            } else {
                                v46 -= v46
                                if false {
                                    switch (v5 >> ((255) & 63)) {
                                    case 845:
                                        if true {
                                            v2 -= (v2 & fi52)
                                            v46 = 0
                                            switch "hello" {
                                            case "a":
                                                var v55 uint64 = 479
                                                if (!false) {
                                                    var v56 uint32 = 65535
                                                    v4[v56] = (^(v56))
                                                    if (v56 >= v56) {
                                                        v57 := v4[237]
                                                        _ = v57
                                                    }
                                                    _ = v56
                                                }
                                                v58 := v4[228]
                                                _ = v55
                                                _ = v58
                                            case "foo bar":
                                                var v59 float64 = ((0.1 + 0.0) * ((float64((0.1 + (1e10 - 0.1))) * 0.0) / 3.14159))
                                                _ = v59
                                            default:
                                                var v60 int64 = -9223372036854775808
                                                _ = v60
                                            }
                                        } else {
                                            v46 -= (v46 / (((^(v46))) | 1))
                                            fmt.Println(fi52, v5)
                                        }
                                    case 60:
                                        var v61 int16 = 42
                                        _ = v61
                                    }
                                    v62 := GenS5{F0: 102}
                                    _ = v62
                                }
                                var v63 float32 = (1e10 - 1e10)
                                _ = v63
                            }
                            v64 := GenS4{F0: 42}
                            _ = v64
                        }
                        var v65 uint8 = 127
                        _ = v65
                    }
                    var v66 uint64 = 775
                    _ = v53
                    _ = v66
                }
                v2 -= fi51
                if (4294967295 != 18446744073709551615) {
                    fmt.Println(v2, fi52)
                    var v67 uint32 = 42
                    _ = v67
                } else {
                    var v68 uint = 2
                    var v69 float64 = float64((((0.1 / 1.0) * float32(1e10)) + -2.25))
                    switch fi51 {
                    case 422:
                        v70 := GenS2{F0: v46, F1: 32767}
                        if false {
                            v2 -= 0
                            var v71 int8 = v46
                            _ = v71
                        } else {
                            v72 := GenS3{F0: 0, F1: 1, F2: 1}
                            switch ("\n\t" + ((("foo bar" + ((("a" + "foo bar") + ("a" + "a")) + ("hello" + (((("hello" + "") + "\n\t") + "hello") + "foo bar")))) + ("foo bar" + "foo bar")) + "foo bar")) {
                            case "a":
                                var v73 uint8 = 255
                                _ = v73
                            default:
                                fmt.Println(v46, fi51)
                            }
                            v5 += (v2 * fi52)
                            fmt.Println(fi51, v68, v5)
                            _ = v72
                        }
                        if (-1 > 596) {
                            var v74 int32 = -1
                            _ = v74
                        }
                        _ = v70
                    case 2147483647:
                        v75 := v4[928]
                        _ = v75
                    case -9223372036854775808:
                        v76 := GenS1{F0: "\n\t", F1: 4294967295}
                        switch 0 {
                        case -9223372036854775808:
                            switch fi51 {
                            case -1:
                                _ = len(v3)
                            case 4294967295:
                                fmt.Println(v5, fi51)
                            default:
                                v68 -= v68
                            }
                        case 42:
                            fmt.Println(v5, v69, fi51)
                        case 32768:
                            var v77 uint = (v68 % ((v68) | 1))
                            _ = len(v6)
                            if true {
                                var v78 float32 = 1.5
                                if true {
                                    var v79 uint32 = 608
                                    _ = v79
                                }
                                if false {
                                    _ = len(v3)
                                    v80 := GenS5{F0: 42}
                                    fmt.Println(v46, v2)
                                    _ = v80
                                } else {
                                    v81 := GenS3{F0: 42, F1: 18446744073709551615, F2: 1}
                                    var v82 string = "a"
                                    _ = v81
                                    _ = v82
                                }
                                _ = v78
                            }
                            switch fi51 {
                            case -1:
                                fmt.Println(v5)
                            case -9223372036854775808:
                                var v83 string = "\n\t"
                                _ = v83
                            case 1:
                                switch (v5 << (((v68 % ((v77) | 1))) & 63)) {
                                case 965:
                                    if true {
                                        switch (v2 + 2147483647) {
                                        case 1:
                                            v46 -= (v46 ^ (v46 / (((v46 << ((v68) & 63))) | 1)))
                                        case 144:
                                            if (-2.25 < v69) {
                                                v68 = v77
                                            } else {
                                                v84 := GenS6{F0: 170, F1: (v68 >> ((4294967295) & 63)), F2: (-(v2))}
                                                var v85 int64 = 333
                                                _ = v84
                                                _ = v85
                                            }
                                        case 4294967295:
                                            switch (fi52 ^ 478) {
                                            case 349:
                                                switch fi52 {
                                                case 2147483647:
                                                    var v86 uint32 = 1
                                                    _ = v86
                                                case 9223372036854775807:
                                                    v68 += 4294967295
                                                case -442:
                                                    v87 := GenS7{F0: 625}
                                                    _ = v87
                                                }
                                            case -321:
                                                fmt.Println(v69, fi51, v68)
                                            case 2147483647:
                                                if (-2147483648 != 42) {
                                                    if true {
                                                        var v88 uint8 = 2
                                                        v89 := GenS7{F0: (v68 * (v77 / ((v68) | 1)))}
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                    delete(v4, 2147483648)
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                } else {
                                                    var v90 uint = v77
                                                    switch "" {
                                                    case "\n\t":
                                                        var v91 int = (-(fi51))
                                                        _ = v91
                                                    case "hello":
                                                        fmt.Println(v5, v5, v69)
                                                    case "foo bar":
                                                        v2 -= 128
                                                    default:
                                                        switch "\n\t" {
                                                        case "":
                                                            if true {
                                                                v5 -= 0
                                                                v92 := GenS8{F0: "", F1: -32768}
                                                                v68 += 4294967295
                                                                _ = v92
                                                            } else {
                                                                var v93 int8 = (v46 & (-(v46)))
                                                                v46 = v46
                                                                _ = v93
                                                            }
                                                        case "\n\t":
                                                            switch 196 {
                                                            case 9223372036854775807:
                                                                v69 += float64(-2.25)
                                                            }
                                                        case "hello":
                                                            if (255 <= (v77 * v68)) {
                                                                v90 = 0
                                                                if (!(!(!(!true)))) {
                                                                    if false {
                                                                        if false {
                                                                            v94 := GenS5{F0: 821}
                                                                            _ = v94
                                                                        } else {
                                                                            v69 -= v69
                                                                            var v95 string = "\n\t"
                                                                            var v96 int64 = 112
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        if len(v6) > 0 {
                                                                            v6[0] = ""
                                                                        }
                                                                    }
                                                                    var v97 int = v2
                                                                    delete(v4, 2147483648)
                                                                    _ = v97
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = 4294967295
                                                                }
                                                            } else {
                                                                var v98 uint64 = 2
                                                                fmt.Println(v77, fi51)
                                                                if (!(uint32(v90) > 221)) {
                                                                    if true {
                                                                        var v99 uint16 = 255
                                                                        _ = v99
                                                                    }
                                                                }
                                                                _ = v98
                                                            }
                                                        default:
                                                            v100 := GenS9{F0: (v77 % ((v77) | 1)), F1: 128}
                                                            _ = v100
                                                        }
                                                    }
                                                    if false {
                                                        v101 := GenS10{F0: (fi52 >> ((v90) & 63)), F1: (v77 + (v90 / (((v68 ^ v77)) | 1)))}
                                                        v69 = (((v69 / 1.0) + v69) * float64(-2.25))
                                                        _ = v101
                                                    } else {
                                                        fmt.Println(v5, v68, v69)
                                                        v6 = append(v6, ("foo bar" + "hello"))
                                                    }
                                                    _ = v90
                                                }
                                            }
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                case 9223372036854775807:
                                    var v102 int = int(v68)
                                    _ = v102
                                }
                            default:
                                _ = len(v3)
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi52 & 127)
                            }
                            _ = v77
                        }
                        _ = v76
                    default:
                        var v103 uint32 = 65535
                        _ = v103
                    }
                    var v104 string = "hello"
                    delete(v4, 4294967295)
                    _ = v104
                    _ = v68
                    _ = v69
                }
            } else {
                v5 = fi52
                var v105 float64 = 0.0
                if (!true) {
                    v5 += fi51
                    fmt.Println(v5, v2)
                } else {
                    var v106 int8 = v46
                    fmt.Println(fi51, v105)
                    _ = v106
                }
                _ = v105
            }
            _ = fi52
        }
        _ = fi51
    }
    var v107 float64 = ((0.0 / 1.0) + float64(float32(0.0)))
    if false {
        v2 = (v5 * v5)
        var v108 uint = 0
        if (!(!true)) {
            v2 = v5
            switch (v5 + -63) {
            case -2147483648:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        _ = v108
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v6
}
