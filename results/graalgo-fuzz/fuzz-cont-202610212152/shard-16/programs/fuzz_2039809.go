package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    var v3 int8 = -128
    v4 := []string{("\n\t" + "hello")}
    v5 := []uint32{410, 65535, 112, 295}
    var v6 int64 = 0
    v3 = (v3 * v3)
    var v7 uint = 42
    v8 := GenS1{F0: ("a" + ("" + ("a" + ("hello" + "a")))), F1: 32767}
    for fi9 := 0; fi9 < 7; fi9++ {
        v7 += (v7 + (v7 << (((v7 | 255)) & 63)))
        _ = fi9
    }
    v1[("\n\t" + ("foo bar" + (("foo bar" + ("a" + "foo bar")) + "hello")))] = 2147483647
    var v10 int64 = -860
    fmt.Println(v7, v10)
    if false {
        if (-2.25 == 1.5) {
            var v11 float32 = float32((1.5 + (((float64((1.5 + 1e10)) / 3.14159) - (((float64(0.0) * 0.1) / (((1.5 / 1.5) + (float64((1.5 / (((0.0 * 0.0) + 1.5) / (((3.14159 + 3.14159) - 0.1) - 3.14159)))) * (0.1 / 1.5))) * 0.1)) - 0.0)) * (0.0 - 0.0))))
            _ = v11
        } else {
            var v12 float64 = (float64(-2.25) * 0.0)
            v13 := GenS1{F0: "foo bar", F1: -623}
            v3 = (v3 / ((v3) | 1))
            _ = v12
            _ = v13
        }
        switch "\n\t" {
        case "\n\t":
            if false {
                var v14 float64 = (float64(0.1) / (0.1 + -2.25))
                _ = len(v4)
                v15 := GenS1{F0: ("" + "\n\t"), F1: 1}
                _ = v14
                _ = v15
            }
        case "":
            if true {
                var v16 int32 = 2147483647
                _ = v16
            } else {
                var v17 int32 = 42
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 uint64 = 192
                    for fi20 := 0; fi20 < 15; fi20++ {
                        fmt.Println(v17)
                        _ = fi20
                    }
                    _ = fi18
                    _ = v19
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 int32 = (v17 % (((v17 - (v17 - 283))) | 1))
                    v23 := v1["foo bar"]
                    if true {
                        v6 = v6
                    } else {
                        for fi24 := 0; fi24 < 20; fi24++ {
                            switch (v23 | (-(fi24))) {
                            case -517:
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                            case 42:
                                if len(v5) > 0 {
                                    v5[0] = 42
                                }
                            case -9223372036854775808:
                                v22 += -2147483648
                            default:
                                var v25 int32 = v22
                                _ = v25
                            }
                            _ = fi24
                        }
                        _ = len(v1)
                        fmt.Println(v22, v22)
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v17
            }
        case "foo bar":
            v6 += 9223372036854775807
        default:
            fmt.Println(v6, v10, v10)
        }
    } else {
        v1["hello"] = 32767
        v3 += -128
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 string = "foo bar"
            if (!(!true)) {
                v28 := GenS1{F0: "foo bar", F1: 42}
                var v29 int8 = v3
                fmt.Println(v27, v7)
                _ = v28
                _ = v29
            } else {
                var v30 uint8 = 128
                var v31 float64 = 0.0
                if false {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 int64 = 0
                        v1["hello"] = (fi32 * 504)
                        v31 -= v31
                        v10 -= (v10 >> ((v7) & 63))
                        v3 += v3
                        _ = fi32
                        _ = v33
                    }
                    var v34 uint32 = 1
                    fmt.Println(v3)
                    _ = v34
                }
                _ = v30
                _ = v31
            }
            for fi35 := 0; fi35 < 1; fi35++ {
                if (3.14159 >= float32((((1e10 - -2.25) * 0.0) - ((-2.25 * (((((0.0 - ((3.14159 + (-2.25 / 1.5)) / (3.14159 + 1e10))) * 1e10) + ((float64(0.1) / 1.5) * 3.14159)) * (0.0 + 1.5)) + -2.25)) / 1.0)))) {
                    v36 := GenS1{F0: v27, F1: -10}
                    if false {
                        v7 -= v7
                        delete(v1, "")
                        if true {
                            var v37 uint16 = v2
                            _ = v37
                        }
                    }
                    switch int(v7) {
                    case 4294967295:
                        var v38 uint = 247
                        _ = v38
                    }
                    _ = v36
                }
                v39 := v1["\n\t"]
                v39 = (fi35 ^ -9223372036854775808)
                _ = fi35
                _ = v39
            }
            _ = fi26
            _ = v27
        }
    }
    v3 -= 126
    switch 0 {
    case 1:
        v3 = v3
    case 9223372036854775807:
        fmt.Println(v7)
    case 2147483647:
        for fi40 := 0; fi40 < 16; fi40++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi40
        }
    default:
        if true {
            v2 = 1
            var v41 bool = (-128 != (v3 / (((-(v3))) | 1)))
            v42 := GenS2{F0: 72, F1: (v2 + (v2 % ((v2) | 1))), F2: 128}
            _ = v41
            _ = v42
        } else {
            v6 += v10
        }
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        v6 -= 42
        v3 += (v3 + v3)
        v44 := GenS2{F0: 119, F1: (v2 ^ 619), F2: 177}
        _ = fi43
        _ = v44
    }
    _ = len(v4)
    if (!(914 == 1)) {
        switch 959 {
        case -244:
            fmt.Println(v2, v10, v7)
        case 2147483647:
            if (true && false) {
                if false {
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        if true {
                            for fi45 := 0; fi45 < 20; fi45++ {
                                if (-609 > 577) {
                                    switch ((("a" + ("a" + "\n\t")) + ("a" + "hello")) + "\n\t") {
                                    case "\n\t":
                                        if false {
                                            var v46 int16 = 32767
                                            v2 = v2
                                            v5 = append(v5, 2147483648)
                                            _ = v46
                                        } else {
                                            _ = len(v1)
                                            v6 = (v6 + v10)
                                            if (!true) {
                                                for fi47 := 0; fi47 < 14; fi47++ {
                                                    fmt.Println(v6)
                                                    var v48 int64 = (v6 / (((v10 ^ (v10 & (v10 / ((9223372036854775807) | 1))))) | 1))
                                                    switch int(v7) {
                                                    case -9223372036854775808:
                                                        v7 += (v7 & v7)
                                                        fmt.Println(v48)
                                                    case -464:
                                                        fmt.Println(v7)
                                                        v3 = v3
                                                        var v49 uint64 = 491
                                                        _ = v49
                                                    case 32767:
                                                        fmt.Println(v2, fi45)
                                                    default:
                                                        delete(v1, ("a" + ("\n\t" + (("hello" + "hello") + "foo bar"))))
                                                    }
                                                    _ = fi47
                                                    _ = v48
                                                }
                                                if false {
                                                    v50 := v1["foo bar"]
                                                    v10 += (^(v10))
                                                    _ = v50
                                                } else {
                                                    v51 := GenS2{F0: 158, F1: (v2 / ((v2) | 1)), F2: 255}
                                                    _ = v51
                                                }
                                                fmt.Println(v10, fi45)
                                            } else {
                                                if (((0.0 - (1e10 + -2.25)) < (float32(((1e10 + 0.0) + 0.0)) - -2.25)) || false) {
                                                    _ = len(v4)
                                                    var v52 uint16 = 42
                                                    _ = v52
                                                } else {
                                                    var v53 string = ("hello" + ("" + "\n\t"))
                                                    fmt.Println(v2, fi45)
                                                    v10 -= 9223372036854775807
                                                    _ = v53
                                                }
                                                switch -824 {
                                                case 9223372036854775807:
                                                    switch (-(fi45)) {
                                                    case 32768:
                                                        _ = len(v5)
                                                    case 32767:
                                                        v6 += 1
                                                    }
                                                default:
                                                    if false {
                                                        v54 := GenS2{F0: 140, F1: v2, F2: 216}
                                                        fmt.Println(v3, v10)
                                                        if ((2 == 110) && true) {
                                                            v2 = v2
                                                        } else {
                                                            var v55 float64 = ((-2.25 + 1e10) + 0.1)
                                                            _ = v55
                                                        }
                                                        _ = v54
                                                    }
                                                }
                                                fmt.Println(v3, v2)
                                            }
                                        }
                                    case "a":
                                        v7 = (^(v7))
                                    case "hello":
                                        switch fi45 {
                                        case 32768:
                                            fmt.Println(v7, v2)
                                        case 711:
                                            fmt.Println(fi45, fi45, v2)
                                        case 0:
                                            switch 4294967295 {
                                            case -858:
                                                v56 := GenS2{F0: 236, F1: (v2 & 443), F2: 42}
                                                _ = v56
                                            case 1:
                                                v10 -= v6
                                            case 42:
                                                v10 = (v6 % ((v10) | 1))
                                            }
                                        }
                                    }
                                    switch 0 {
                                    case -1:
                                        v3 = (^(v3))
                                    case 171:
                                        v57 := v1[("" + ((("\n\t" + ("foo bar" + "a")) + ("a" + "hello")) + "\n\t"))]
                                        _ = v57
                                    case 727:
                                        var v58 float64 = float64(((3.14159 + 1e10) - 3.14159))
                                        _ = v58
                                    default:
                                        _ = len(v5)
                                    }
                                }
                                _ = fi45
                            }
                            if false {
                                if (!false) {
                                    for fi59 := 0; fi59 < 13; fi59++ {
                                        if (false && ((fi59 != (fi59 >> (((v7 >> (((^(v7))) & 63))) & 63))) || true)) {
                                            fmt.Println(v6, v3, v10)
                                            for fi60 := 0; fi60 < 16; fi60++ {
                                                if (!true) {
                                                    v10 = (^(v6))
                                                    switch "" {
                                                    case "hello":
                                                        v61 := GenS3{F0: 32767, F1: -60}
                                                        _ = v61
                                                    case "a":
                                                        var v62 bool = true
                                                        _ = v62
                                                    case "\n\t":
                                                        switch ("a" + "\n\t") {
                                                        case "foo bar":
                                                            v2 = v2
                                                        }
                                                    }
                                                } else {
                                                    var v63 float64 = 1.5
                                                    _ = v63
                                                }
                                                _ = fi60
                                            }
                                            v64 := GenS4{F0: (v6 << ((0) & 63))}
                                            _ = v64
                                        }
                                        _ = fi59
                                    }
                                } else {
                                    if (65535 == v7) {
                                        v1["foo bar"] = -70
                                        v65 := GenS3{F0: 714, F1: 358}
                                        _ = v65
                                    } else {
                                        v66 := GenS5{F0: ("" + ("\n\t" + "")), F1: 0}
                                        if false {
                                            for fi67 := 0; fi67 < 16; fi67++ {
                                                fmt.Println(v3, v7)
                                                if true {
                                                    v3 += int8(fi67)
                                                    var v68 uint16 = v2
                                                    _ = v68
                                                }
                                                var v69 int8 = v3
                                                _ = fi67
                                                _ = v69
                                            }
                                        } else {
                                            switch 2147483647 {
                                            case -800:
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = "\n\t"
                                                    }
                                                    fmt.Println(v10, v3, v3)
                                                }
                                            }
                                            fmt.Println(v7)
                                        }
                                        _ = v66
                                    }
                                }
                            }
                        } else {
                            for fi70 := 0; fi70 < 14; fi70++ {
                                fmt.Println(v2, fi70)
                                _ = fi70
                            }
                            if (1 > 47) {
                                v7 += (v7 >> ((4294967295) & 63))
                                switch 4294967295 {
                                case 128:
                                    for fi71 := 0; fi71 < 5; fi71++ {
                                        _ = len(v1)
                                        for fi72 := 0; fi72 < 14; fi72++ {
                                            v10 = v10
                                            var v73 uint8 = 42
                                            _ = fi72
                                            _ = v73
                                        }
                                        _ = fi71
                                    }
                                default:
                                    switch 42 {
                                    case 121:
                                        var v74 float32 = (1.5 * float32(float64(-2.25)))
                                        _ = v74
                                    case -257:
                                        if true {
                                            var v75 bool = true
                                            if true {
                                                v7 += (v7 << (((v7 / ((v7) | 1))) & 63))
                                                fmt.Println(v10)
                                            } else {
                                                for fi76 := 0; fi76 < 19; fi76++ {
                                                    for fi77 := 0; fi77 < 14; fi77++ {
                                                        v78 := GenS6{F0: 18446744073709551615}
                                                        _ = fi77
                                                        _ = v78
                                                    }
                                                    v4 = append(v4, ("" + "hello"))
                                                    v10 = v6
                                                    _ = fi76
                                                }
                                                v6 = v6
                                            }
                                            fmt.Println(v75, v3)
                                            _ = v75
                                        } else {
                                            var v79 uint16 = (v2 / (((v2 & 42)) | 1))
                                            _ = v79
                                        }
                                    default:
                                        v1["\n\t"] = 127
                                    }
                                }
                            } else {
                                if (!(((0.0 - 1.5) != 0.0) && (42 < 2))) {
                                    v80 := GenS6{F0: (v7 + v7)}
                                    delete(v1, "hello")
                                    var v81 int16 = 668
                                    _ = v80
                                    _ = v81
                                }
                                switch -1 {
                                case 128:
                                    v82 := GenS1{F0: "\n\t", F1: 0}
                                    _ = v82
                                case 32768:
                                    for fi83 := 0; fi83 < 14; fi83++ {
                                        for fi84 := 0; fi84 < 18; fi84++ {
                                            if false {
                                                if (65535 <= (^(v2))) {
                                                    v85 := v1[("" + ((("foo bar" + "\n\t") + "") + ""))]
                                                    v86 := GenS7{F0: 2147483648}
                                                    _ = v85
                                                    _ = v86
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v87 uint32 = 65535
                                                var v88 uint32 = (^(v87))
                                                if true {
                                                    if (222 == 694) {
                                                        var v89 uint16 = 960
                                                        if (((736 >= 0) && false) || (true || ((v2 | 633) < 224))) {
                                                            fmt.Println(v2)
                                                            fmt.Println(v6, v88, fi83)
                                                        } else {
                                                            var v90 float32 = 3.14159
                                                            v3 -= v3
                                                            _ = v90
                                                        }
                                                        _ = v89
                                                    }
                                                    switch (fi83 + 9223372036854775807) {
                                                    case -764:
                                                        switch int(v87) {
                                                        case 9223372036854775807:
                                                            var v91 int8 = v3
                                                            _ = v91
                                                        }
                                                    default:
                                                        v10 = v10
                                                    }
                                                    var v92 int = (fi83 - fi83)
                                                    v7 = v7
                                                    v2 -= (^(v2))
                                                    v1["\n\t"] = -1
                                                    _ = v92
                                                }
                                                _ = v87
                                                _ = v88
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = 42
                                                }
                                            }
                                            v3 = 100
                                            v7 += 18446744073709551615
                                            _ = fi84
                                        }
                                        var v93 uint32 = 503
                                        _ = len(v1)
                                        _ = fi83
                                        _ = v93
                                    }
                                case -254:
                                    _ = len(v4)
                                }
                            }
                        }
                    }
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        if true {
                            v7 = (v7 + v7)
                            var v94 int16 = 32767
                            var v95 float32 = -2.25
                            _ = v94
                            _ = v95
                        }
                    }
                    switch -1 {
                    case -9223372036854775808:
                        var v96 bool = (true && false)
                        _ = v96
                    case 127:
                        var v97 uint32 = 0
                        _ = v97
                    }
                } else {
                    v7 -= (v7 ^ (v7 - 65535))
                    var v98 uint16 = v2
                    v7 += v7
                    _ = v98
                }
                switch "foo bar" {
                case "a":
                    v2 = 654
                case "hello":
                    v6 = v6
                default:
                    v7 += v7
                }
                switch 1 {
                case 4294967295:
                    for fi99 := 0; fi99 < 16; fi99++ {
                        v2 = 32767
                        var v100 uint32 = 42
                        _ = len(v1)
                        _ = fi99
                        _ = v100
                    }
                case 376:
                    var v101 uint16 = 300
                    _ = v101
                case 548:
                    v102 := GenS3{F0: 32767, F1: 42}
                    _ = v102
                }
            }
        default:
            v2 += (v2 << (((v7 - (v7 & 2))) & 63))
        }
        var v103 float64 = (0.1 - (1.5 / 0.1))
        var v104 float32 = -2.25
        _ = v103
        _ = v104
    }
    v105 := GenS1{F0: "hello", F1: 1000}
    if false {
        for fi106 := 0; fi106 < 11; fi106++ {
            v6 = (v10 % ((1) | 1))
            _ = fi106
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
