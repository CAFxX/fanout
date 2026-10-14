package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint16; F2 int }
type GenS6 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint8 = 128
    var v2 int64 = 1
    v3 := make(map[uint32]int)
    v4 := []string{}
    var v5 uint16 = 255
    var v6 bool = true
    var v7 uint8 = (v1 * v1)
    v7 += 72
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float64 = -2.25
        v10 := GenS1{F0: 255}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int16 = -32768
        var v13 int64 = 9223372036854775807
        for fi14 := 0; fi14 < 3; fi14++ {
            v3[42] = fi11
            v2 = 9223372036854775807
            _ = fi14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    if (-112 > 55) {
        fmt.Println(v6, v1, v2)
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint32 = 65535
                _ = fi16
                _ = v17
            }
            fmt.Println(fi15, v5, v2)
            _ = fi15
        }
        fmt.Println(v2)
    } else {
        if v6 {
            v6 = true
        } else {
            if v6 {
                v5 -= v5
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v1, v1)
                case "\n\t":
                    v6 = v6
                default:
                    var v18 uint16 = v5
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 18; fi19++ {
                    v5 += (^(v5))
                    for fi20 := 0; fi20 < 8; fi20++ {
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        fmt.Println(fi19, fi19)
                        if v6 {
                            delete(v3, 42)
                            v4 = append(v4, "\n\t")
                            v2 += v2
                            v6 = (((v7 / (((v7 | (v7 | v1))) | 1)) >= (v7 & v7)) && (!(65535 != 42)))
                        } else {
                            var v21 int32 = 0
                            if ((1 <= 986) && ("\n\t" < (("" + "a") + ("\n\t" + (("" + "a") + ("foo bar" + ("hello" + "\n\t"))))))) {
                                if (!(-2.25 != (float64(0.0) * 0.1))) {
                                    var v22 string = "hello"
                                    _ = v22
                                } else {
                                    v3[65535] = (fi19 | 507)
                                    if v6 {
                                        _ = len(v4)
                                        fmt.Println(v6)
                                    }
                                    v23 := GenS1{F0: 973}
                                    var v24 float64 = (0.0 * -2.25)
                                    var v25 string = "foo bar"
                                    if v6 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if v6 {
                                            if len(v4) > 0 {
                                                v4[0] = v25
                                            }
                                            if v6 {
                                                switch fi20 {
                                                case 42:
                                                    if (!(1 >= 0)) {
                                                        switch v25 {
                                                        case "":
                                                            switch (-(fi19)) {
                                                            case 0:
                                                                fmt.Println(v1)
                                                            case 766:
                                                                v7 = (v1 / (((v1 / (((v1 - (^(v7)))) | 1))) | 1))
                                                            case -506:
                                                                _ = len(v4)
                                                            default:
                                                                if v6 {
                                                                    var v26 float32 = (float32(0.0) + 3.14159)
                                                                    fmt.Println(v26, fi20, v1)
                                                                    _ = v26
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v2, fi20, fi20)
                                                        fmt.Println(v24)
                                                    } else {
                                                        v27 := GenS2{F0: ((v25 + "\n\t") + ("a" + v25)), F1: v5, F2: 396}
                                                        v25 = v25
                                                        _ = v27
                                                    }
                                                }
                                                if true {
                                                    delete(v3, 42)
                                                    var v28 int64 = 42
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v28
                                                } else {
                                                    if (!(!(!((v2 << ((4294967295) & 63)) != 42)))) {
                                                        var v29 int = 0
                                                        v5 = 51
                                                        _ = v29
                                                    } else {
                                                        if v6 {
                                                            fmt.Println(v25)
                                                            if v6 {
                                                                switch v25 {
                                                                case "":
                                                                    if (!v6) {
                                                                        switch v25 {
                                                                        case "\n\t":
                                                                            switch v25 {
                                                                            case "a":
                                                                                var v30 string = v25
                                                                                if true {
                                                                                    if (v5 < 2) {
                                                                                        switch v30 {
                                                                                        case "a":
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = "foo bar"
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v31 uint = 1
                                                                                            _ = v31
                                                                                        case "hello":
                                                                                            v21 += v21
                                                                                        default:
                                                                                            v2 = (v2 - (v2 & (^(v2))))
                                                                                        }
                                                                                    } else {
                                                                                        v21 = v21
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            v21 -= 546
                                                                                        default:
                                                                                            v32 := GenS2{F0: v30, F1: (v5 ^ uint16(v7)), F2: uint64(v7)}
                                                                                            _ = v32
                                                                                        }
                                                                                        var v33 float64 = (0.0 / v24)
                                                                                        _ = v33
                                                                                    }
                                                                                    fmt.Println(v30, v24)
                                                                                } else {
                                                                                    _ = len(v4)
                                                                                    var v34 int16 = 952
                                                                                    _ = v34
                                                                                }
                                                                                _ = v30
                                                                            case "":
                                                                                var v35 uint32 = 2147483647
                                                                                _ = v35
                                                                            default:
                                                                                v36 := GenS1{F0: 255}
                                                                                var v37 uint32 = 2147483647
                                                                                _ = v36
                                                                                _ = v37
                                                                            }
                                                                        }
                                                                        fmt.Println(v6)
                                                                    } else {
                                                                        v38 := GenS3{F0: 0, F1: 2147483647, F2: -849}
                                                                        _ = v38
                                                                    }
                                                                case "a":
                                                                    var v39 int16 = 32767
                                                                    _ = v39
                                                                case "hello":
                                                                    switch v25 {
                                                                    case "hello":
                                                                        switch v25 {
                                                                        case "foo bar":
                                                                            v21 += v21
                                                                        case "":
                                                                            v40 := GenS3{F0: 2, F1: 778, F2: 0}
                                                                            _ = v40
                                                                        case "\n\t":
                                                                            v41 := v3[42]
                                                                            _ = v41
                                                                        default:
                                                                            var v42 int16 = -1
                                                                            _ = v42
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                switch fi19 {
                                                                case 0:
                                                                    switch fi19 {
                                                                    case 127:
                                                                        v43 := GenS3{F0: 4294967295, F1: 47, F2: -298}
                                                                        _ = v43
                                                                    default:
                                                                        fmt.Println(v1, v24, v1)
                                                                    }
                                                                case -899:
                                                                    var v44 float32 = 1e10
                                                                    _ = v44
                                                                case 128:
                                                                    var v45 bool = (0 > v5)
                                                                    _ = v45
                                                                default:
                                                                    var v46 int8 = 42
                                                                    _ = v46
                                                                }
                                                                fmt.Println(fi19, v7)
                                                            }
                                                            var v47 float64 = v24
                                                            _ = v47
                                                        }
                                                        if (((!(v6 && v6)) || false) || (-494 == 340)) {
                                                            var v48 int32 = v21
                                                            fmt.Println(v1, v6, v7)
                                                            v1 = v7
                                                            _ = v48
                                                        }
                                                    }
                                                }
                                            } else {
                                                if (693 <= 2147483647) {
                                                    v49 := GenS4{F0: 55}
                                                    _ = v49
                                                }
                                                switch 0 {
                                                case 948:
                                                    v50 := GenS1{F0: 882}
                                                    _ = v50
                                                case 197:
                                                    v51 := GenS5{F0: 9223372036854775807, F1: 618, F2: int(v5)}
                                                    _ = v51
                                                }
                                            }
                                        }
                                        var v52 uint16 = (v5 * 353)
                                        _ = v52
                                    }
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                var v53 int64 = (v2 ^ v2)
                                _ = v53
                            }
                            v21 += (v21 - v21)
                            _ = v21
                        }
                        _ = fi20
                    }
                    for fi54 := 0; fi54 < 5; fi54++ {
                        v4 = append(v4, "hello")
                        _ = fi54
                    }
                    _ = fi19
                }
            }
            v55 := v3[1]
            _ = v55
        }
    }
    if (v5 == 32767) {
        if true {
            for fi56 := 0; fi56 < 8; fi56++ {
                var v57 uint = 42
                _ = fi56
                _ = v57
            }
            var v58 int16 = -571
            if v6 {
                v7 = v7
                for fi59 := 0; fi59 < 9; fi59++ {
                    if (!v6) {
                        switch 9223372036854775807 {
                        case 1:
                            switch ("foo bar" + (("\n\t" + "\n\t") + "\n\t")) {
                            case "a":
                                delete(v3, 42)
                            case "hello":
                                var v60 float32 = (1e10 * (float32((((3.14159 * ((-2.25 + 1e10) * 1e10)) / 1e10) + (float64(3.14159) + (1e10 / 1.0)))) * 1.5))
                                _ = v60
                            default:
                                v61 := GenS3{F0: 987, F1: 2147483648, F2: v58}
                                _ = v61
                            }
                        case 127:
                            var v62 uint32 = 231
                            _ = v62
                        case -786:
                            v7 += (v1 >> ((749) & 63))
                        }
                        var v63 uint64 = 5
                        _ = v63
                    }
                    v1 -= uint8(v5)
                    _ = fi59
                }
            } else {
                switch ("\n\t" + (("hello" + "") + "")) {
                case "a":
                    for fi64 := 0; fi64 < 15; fi64++ {
                        if (!true) {
                            for fi65 := 0; fi65 < 16; fi65++ {
                                if v6 {
                                    if v6 {
                                        switch ("" + "a") {
                                        case "a":
                                            v6 = false
                                            v66 := GenS5{F0: (fi65 / ((int(v2)) | 1)), F1: 1, F2: fi65}
                                            if (true || v6) {
                                                delete(v3, 1)
                                                var v67 int16 = (v58 / (((v58 << ((4294967295) & 63))) | 1))
                                                _ = v67
                                            } else {
                                                v1 = 1
                                                _ = len(v4)
                                            }
                                            _ = v66
                                        }
                                        switch ("" + "foo bar") {
                                        case "foo bar":
                                            var v68 uint32 = 605
                                            _ = v68
                                        case "\n\t":
                                            fmt.Println(fi64)
                                        default:
                                            var v69 int16 = v58
                                            _ = v69
                                        }
                                    }
                                    var v70 uint8 = (v1 - v7)
                                    var v71 int16 = (v58 | (v58 + v58))
                                    _ = v70
                                    _ = v71
                                }
                                v1 += v1
                                fmt.Println(v7)
                                _ = fi65
                            }
                        }
                        var v72 int8 = 100
                        for fi73 := 0; fi73 < 7; fi73++ {
                            if true {
                                var v74 uint64 = 2
                                var v75 int = fi64
                                _ = v74
                                _ = v75
                            } else {
                                v58 = (v58 - v58)
                                fmt.Println(v58)
                            }
                            _ = fi73
                        }
                        _ = fi64
                        _ = v72
                    }
                case "hello":
                    v76 := GenS6{F0: 2147483647, F1: 42, F2: (^(v1))}
                    _ = v76
                case "foo bar":
                    v77 := GenS5{F0: -650, F1: v5, F2: -506}
                    _ = v77
                default:
                    v78 := GenS6{F0: 4294967295, F1: v58, F2: (v7 + (v1 | 127))}
                    _ = v78
                }
                v7 -= 128
            }
            _ = v58
        } else {
            v6 = true
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            for fi79 := 0; fi79 < 3; fi79++ {
                if (!false) {
                    if false {
                        if v6 {
                            v80 := v3[1]
                            var v81 string = ""
                            _ = v80
                            _ = v81
                        } else {
                            var v82 uint64 = 2
                            switch fi79 {
                            case 65:
                                v82 += (^(v82))
                            case 1:
                                _ = len(v3)
                            case -778:
                                var v83 string = (("\n\t" + "\n\t") + "\n\t")
                                _ = v83
                            default:
                                for fi84 := 0; fi84 < 4; fi84++ {
                                    if (!false) {
                                        v7 += 127
                                    } else {
                                        v82 = 18446744073709551615
                                    }
                                    _ = fi84
                                }
                            }
                            fmt.Println(v7, v1, v1)
                            _ = v82
                        }
                    }
                    switch "a" {
                    case "":
                        v4 = append(v4, ("hello" + ("\n\t" + ((("\n\t" + "") + "\n\t") + ("foo bar" + "foo bar")))))
                    case "foo bar":
                        switch (fi79 >> ((65535) & 63)) {
                        case 211:
                            switch "" {
                            case "foo bar":
                                var v85 uint8 = v7
                                _ = v85
                            case "hello":
                                v1 = v1
                            }
                        case 32767:
                            fmt.Println(v7, v5)
                        case 128:
                            if (-1 != -816) {
                                switch fi79 {
                                case 127:
                                    fmt.Println(fi79, fi79, v6)
                                case 0:
                                    for fi86 := 0; fi86 < 4; fi86++ {
                                        v2 -= v2
                                        v7 += v1
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        v1 = v7
                                        _ = fi86
                                    }
                                case 427:
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                }
                                v87 := v3[344]
                                _ = v87
                            }
                        default:
                            fmt.Println(v1, v6)
                        }
                    }
                    var v88 int32 = -2147483648
                    _ = v88
                }
                _ = fi79
            }
        }
        if (!(939 < 1)) {
            v5 -= (v5 ^ (v5 & 905))
            v89 := GenS3{F0: 42, F1: 65535, F2: 32767}
            var v90 uint8 = uint8(v5)
            _ = v89
            _ = v90
        } else {
            v91 := GenS3{F0: 1, F1: uint32(v5), F2: 42}
            v92 := GenS7{F0: -1, F1: 255}
            _ = v91
            _ = v92
        }
        if false {
            v6 = (0 != (v5 >> ((0) & 63)))
            var v93 float64 = 1e10
            _ = v93
        } else {
            for fi94 := 0; fi94 < 16; fi94++ {
                v4 = append(v4, "foo bar")
                if (uint64(v5) == 4294967295) {
                    v1 = v1
                }
                var v95 int16 = 241
                _ = fi94
                _ = v95
            }
            for fi96 := 0; fi96 < 14; fi96++ {
                if (true || (uint(v5) == 18446744073709551615)) {
                    var v97 uint16 = v5
                    _ = v97
                } else {
                    switch (^(fi96)) {
                    case -2147483648:
                        var v98 int8 = -1
                        _ = v98
                    case 4294967295:
                        var v99 float64 = -2.25
                        _ = v99
                    case 9223372036854775807:
                        for fi100 := 0; fi100 < 17; fi100++ {
                            v2 = v2
                            v3[4294967295] = fi96
                            v2 -= v2
                            _ = fi100
                        }
                    }
                    v7 += (v7 & v7)
                    fmt.Println(fi96)
                }
                var v101 int = fi96
                var v102 int64 = v2
                _ = fi96
                _ = v101
                _ = v102
            }
            for fi103 := 0; fi103 < 12; fi103++ {
                var v104 int32 = 42
                switch (fi103 / ((int(v1)) | 1)) {
                case 128:
                    if (v5 == 351) {
                        var v105 bool = (false && (!(65535 != 65535)))
                        if false {
                            var v106 int = fi103
                            _ = v106
                        }
                        v5 = v5
                        _ = v105
                    }
                }
                _ = fi103
                _ = v104
            }
        }
    }
    v7 = 255
    for fi107 := 0; fi107 < 7; fi107++ {
        for fi108 := 0; fi108 < 20; fi108++ {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            v4 = append(v4, "hello")
            _ = fi108
        }
        switch "foo bar" {
        case "foo bar":
            _ = len(v4)
        case "\n\t":
            v6 = (v5 >= v5)
        case "":
            switch "a" {
            case "a":
                var v109 uint = 1
                _ = v109
            case "hello":
                if (0 != (v2 % ((int64(fi107)) | 1))) {
                    if v6 {
                        var v110 int32 = -2147483648
                        fmt.Println(v1, v2)
                        _ = v110
                    }
                    for fi111 := 0; fi111 < 10; fi111++ {
                        fmt.Println(v2)
                        var v112 uint32 = 0
                        _ = fi111
                        _ = v112
                    }
                } else {
                    v1 -= v1
                }
            case "":
                if (false || false) {
                    var v113 int16 = 0
                    v6 = (!(v1 >= (^(v1))))
                    fmt.Println(v5)
                    _ = v113
                }
            default:
                for fi114 := 0; fi114 < 20; fi114++ {
                    var v115 uint16 = 42
                    v116 := GenS3{F0: 18446744073709551615, F1: uint32(v1), F2: -32768}
                    var v117 float32 = (float32(1.5) * 1e10)
                    _ = fi114
                    _ = v115
                    _ = v116
                    _ = v117
                }
            }
        default:
            v7 += v1
        }
        _ = fi107
    }
    if ("" >= "a") {
        var v118 float32 = 0.1
        var v119 int = int(v7)
        fmt.Println(v7, v1)
        _ = v118
        _ = v119
    } else {
        v6 = v6
    }
    v5 = v5
    v4 = append(v4, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
