package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := []string{"\n\t", "\n\t", ("a" + "\n\t")}
    v5 := make(map[int]uint32)
    var v6 float64 = 1.5
    var v7 float64 = v6
    for fi8 := 0; fi8 < 13; fi8++ {
        v7 -= (3.14159 - -2.25)
        v9 := GenS1{F0: 1, F1: 169, F2: 2}
        _ = fi8
        _ = v9
    }
    fmt.Println(v6)
    for fi10 := 0; fi10 < 8; fi10++ {
        v6 += v7
        _ = fi10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v7 += (v7 / v7)
        _ = fi11
    }
    if ("a" > "") {
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(v1, v7)
            v13 := GenS1{F0: 801, F1: 695, F2: 255}
            for fi14 := 0; fi14 < 12; fi14++ {
                switch (("" + "hello") + "foo bar") {
                case "foo bar":
                    var v15 float64 = (float64(v1) * 0.1)
                    _ = v15
                case "hello":
                    var v16 int = (-(fi14))
                    _ = v16
                }
                switch (fi12 + fi14) {
                case 548:
                    v1 += v1
                case 127:
                    if false {
                        if true {
                            if false {
                                v1 = float32((v6 * (v6 / -2.25)))
                                _ = len(v3)
                                v17 := GenS2{F0: 0, F1: "a"}
                                _ = v17
                            } else {
                                if (!(1 != 42)) {
                                    switch "" {
                                    case "hello":
                                        v1 = 1e10
                                    }
                                    if (-377 > 1) {
                                        fmt.Println(v6, v7, fi12)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    var v18 int64 = -9223372036854775808
                                    _ = v18
                                }
                            }
                            var v19 float32 = float32((float64(3.14159) * ((v6 * v6) * v6)))
                            _ = v19
                        } else {
                            _ = len(v4)
                            var v20 float32 = v1
                            var v21 uint = 4294967295
                            v22 := GenS2{F0: 2, F1: "hello"}
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        var v23 int64 = 994
                        var v24 int8 = 1
                        _ = v23
                        _ = v24
                    }
                case 32768:
                    if (!(987 > 0)) {
                        v6 -= (v7 + v6)
                    } else {
                        v7 = v7
                    }
                default:
                    v1 += (v1 / v1)
                }
                if false {
                    v6 = 1e10
                    switch "\n\t" {
                    case "foo bar":
                        var v25 uint = 42
                        _ = v25
                    case "":
                        var v26 int32 = -1
                        _ = v26
                    }
                } else {
                    v7 = float64(v1)
                }
                v7 = (float64(v1) + float64(v1))
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        for fi27 := 0; fi27 < 12; fi27++ {
            fmt.Println(v7, v1, v7)
            if ("\n\t" < "a") {
                v7 = v6
            }
            _ = fi27
        }
        switch ("foo bar" + ("" + "a")) {
        case "\n\t":
            var v28 string = "foo bar"
            _ = v28
        case "hello":
            switch 42 {
            case -123:
                if false {
                    delete(v2, 2147483647)
                }
            case 32767:
                v29 := GenS2{F0: 43, F1: "a"}
                _ = v29
            default:
                switch "foo bar" {
                case "foo bar":
                    var v30 int64 = 1
                    _ = v30
                case "hello":
                    if true {
                        if (191 >= 210) {
                            v31 := GenS2{F0: 18446744073709551615, F1: "a"}
                            for fi32 := 0; fi32 < 17; fi32++ {
                                for fi33 := 0; fi33 < 7; fi33++ {
                                    v34 := GenS1{F0: 18446744073709551615, F1: 0, F2: 0}
                                    _ = fi33
                                    _ = v34
                                }
                                v35 := GenS1{F0: 945, F1: 65535, F2: 42}
                                switch fi32 {
                                case -289:
                                    switch 0 {
                                    case 9223372036854775807:
                                        v36 := GenS3{F0: 582, F1: false}
                                        _ = v36
                                    }
                                case 2147483647:
                                    v37 := GenS4{F0: (^(fi32)), F1: 186, F2: 0}
                                    _ = v37
                                default:
                                    var v38 int16 = 974
                                    _ = v38
                                }
                                _ = fi32
                                _ = v35
                            }
                            if (1 == 0) {
                                switch -742 {
                                case -2147483648:
                                    for fi39 := 0; fi39 < 6; fi39++ {
                                        if (false && (-1 > 1000)) {
                                            v7 = (float64(0.0) + (v6 * v7))
                                        }
                                        _ = fi39
                                    }
                                case 4294967295:
                                    fmt.Println(v1, v7)
                                default:
                                    for fi40 := 0; fi40 < 2; fi40++ {
                                        switch fi40 {
                                        case 1:
                                            if false {
                                                var v41 float32 = 0.0
                                                for fi42 := 0; fi42 < 13; fi42++ {
                                                    v41 = v41
                                                    var v43 float32 = ((-2.25 - ((float32(1e10) + (((v1 / v41) * 0.0) / v1)) - 1e10)) / (v1 * (v1 / float32(v7))))
                                                    _ = fi42
                                                    _ = v43
                                                }
                                                v44 := v2[14]
                                                _ = v41
                                                _ = v44
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        case 42:
                                            v6 += (((v6 - 1.5) + 3.14159) * (v6 - v7))
                                        case -9223372036854775808:
                                            v1 = v1
                                        default:
                                            var v45 int64 = -9223372036854775808
                                            _ = v45
                                        }
                                        _ = fi40
                                    }
                                }
                                fmt.Println(v1, v6)
                            }
                            _ = v31
                        } else {
                            v6 += v7
                            v46 := GenS5{F0: -1, F1: 0}
                            _ = v46
                        }
                        var v47 int16 = 506
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v47
                    } else {
                        v7 = (0.1 + (3.14159 * v7))
                        switch "foo bar" {
                        case "\n\t":
                            var v48 uint64 = 4294967295
                            _ = v48
                        case "hello":
                            v7 += (0.1 * v7)
                        }
                        if true {
                            switch "" {
                            case "a":
                                fmt.Println(v7, v1)
                            }
                            v49 := GenS4{F0: -2147483648, F1: 0, F2: 239}
                            var v50 int8 = -50
                            _ = v49
                            _ = v50
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v7 = ((float64(1e10) - 3.14159) + float64(((0.0 * (1e10 / v1)) * (3.14159 - (1e10 / (3.14159 + v1))))))
                            case "foo bar":
                                var v51 float32 = (float32(v6) / v1)
                                _ = v51
                            case "hello":
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v7, v1, v1)
                                case "":
                                    v6 -= (v6 * v6)
                                case "foo bar":
                                    var v52 bool = true
                                    _ = v52
                                }
                            }
                            var v53 uint16 = 42
                            var v54 string = ("\n\t" + "hello")
                            _ = v53
                            _ = v54
                        }
                    }
                }
            }
        case "foo bar":
            var v55 int32 = -1
            _ = v55
        default:
            v7 = v6
        }
    } else {
        v56 := GenS1{F0: 426, F1: 65535, F2: 4294967295}
        _ = len(v4)
        fmt.Println(v1)
        _ = v56
    }
    v1 = float32(v7)
    fmt.Println(v1)
    for fi57 := 0; fi57 < 15; fi57++ {
        if true {
            if (false && false) {
                if false {
                    var v58 int8 = 0
                    var v59 float64 = ((3.14159 + (0.1 * (3.14159 / (1.5 / -2.25)))) * 0.1)
                    for fi60 := 0; fi60 < 2; fi60++ {
                        var v61 int8 = v58
                        v7 -= v59
                        var v62 uint32 = 2
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    _ = v58
                    _ = v59
                } else {
                    v1 += (((v1 * (0.1 - 0.1)) - (1.5 - v1)) + (v1 - (v1 + float32(float64((1e10 / v1))))))
                }
                v6 -= v7
                switch "foo bar" {
                case "a":
                    for fi63 := 0; fi63 < 5; fi63++ {
                        switch "\n\t" {
                        case "":
                            v1 = (float32((0.1 + ((v6 + (v6 + v6)) / 1.5))) + (-2.25 + v1))
                        case "foo bar":
                            var v64 int64 = -1
                            _ = v64
                        default:
                            var v65 uint = 65535
                            _ = v65
                        }
                        fmt.Println(fi63, v1)
                        _ = fi63
                    }
                }
            }
            var v66 bool = true
            v66 = v66
            _ = v66
        } else {
            var v67 int32 = -824
            _ = v67
        }
        for fi68 := 0; fi68 < 14; fi68++ {
            _ = len(v4)
            var v69 int64 = 795
            _ = fi68
            _ = v69
        }
        _ = fi57
    }
    for fi70 := 0; fi70 < 16; fi70++ {
        v71 := GenS6{F0: 65535}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi70
        _ = v71
    }
    v7 -= float64(v1)
    v3[750] = 2
    if true {
        var v72 string = ((("" + "hello") + "foo bar") + ("a" + "a"))
        for fi73 := 0; fi73 < 11; fi73++ {
            var v74 uint32 = 1
            switch v72 {
            case "hello":
                delete(v5, 326)
            default:
                if true {
                    switch (fi73 - 128) {
                    case -2147483648:
                        v1 -= float32(v6)
                    case 613:
                        var v75 int = (fi73 & 2147483647)
                        _ = v75
                    case 42:
                        v1 = 0.0
                    default:
                        if (663 >= 4294967295) {
                            v7 -= (float64(float32(v7)) + 3.14159)
                            var v76 int = (fi73 % ((fi73) | 1))
                            v6 += v7
                            _ = v76
                        } else {
                            v72 = v72
                        }
                    }
                } else {
                    if true {
                        if (99 != 127) {
                            var v77 uint64 = 4294967295
                            v78 := GenS6{F0: 355}
                            v79 := v5[(fi73 * (fi73 + (fi73 << ((73) & 63))))]
                            _ = v77
                            _ = v78
                            _ = v79
                        }
                    } else {
                        for fi80 := 0; fi80 < 15; fi80++ {
                            var v81 string = (v72 + ((v72 + v72) + ((v72 + (v72 + v72)) + "foo bar")))
                            v82 := GenS6{F0: 349}
                            _ = fi80
                            _ = v81
                            _ = v82
                        }
                        if true {
                            switch "" {
                            case "a":
                                fmt.Println(v7, v72, v1)
                            }
                            if false {
                                var v83 int16 = 0
                                var v84 float64 = 3.14159
                                v85 := GenS2{F0: 18446744073709551615, F1: "hello"}
                                _ = v83
                                _ = v84
                                _ = v85
                            }
                            switch -818 {
                            case -258:
                                switch 2147483647 {
                                case 32767:
                                    delete(v3, 132)
                                default:
                                    v86 := GenS4{F0: fi73, F1: 128, F2: 42}
                                    _ = v86
                                }
                            case -9223372036854775808:
                                for fi87 := 0; fi87 < 6; fi87++ {
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                    if false {
                                        v7 -= -2.25
                                        v1 = v1
                                        v7 += 1.5
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                    }
                                    fmt.Println(fi87)
                                    _ = fi87
                                }
                            }
                        }
                    }
                }
            }
            _ = fi73
            _ = v74
        }
        v7 -= float64(v1)
        _ = v72
    } else {
        for fi88 := 0; fi88 < 5; fi88++ {
            for fi89 := 0; fi89 < 20; fi89++ {
                v1 -= v1
                _ = fi89
            }
            for fi90 := 0; fi90 < 6; fi90++ {
                v6 += 0.0
                switch (("hello" + "foo bar") + ((("" + "") + ((("a" + "\n\t") + ("hello" + "")) + "a")) + ("a" + "a"))) {
                case "foo bar":
                    if ("hello" < "hello") {
                        fmt.Println(v6, v1)
                    }
                case "\n\t":
                    if false {
                        switch fi88 {
                        case 127:
                            _ = len(v5)
                        case 32767:
                            v91 := GenS6{F0: 4294967295}
                            _ = v91
                        case -587:
                            v6 -= (v7 * ((((v7 - (float64(-2.25) + v7)) + v6) / (((v6 + 0.1) - 0.1) * (-2.25 / (1.5 * (v7 / v7))))) + float64(-2.25)))
                        }
                        if (0 >= -2147483648) {
                            var v92 int32 = 1
                            _ = v92
                        }
                        var v93 int16 = 975
                        _ = v93
                    }
                }
                fmt.Println(v1, fi88, v6)
                _ = fi90
            }
            v1 = (v1 - 0.1)
            _ = fi88
        }
    }
    v94 := GenS6{F0: 858}
    v7 -= v7
    var v95 int64 = 42
    v96 := GenS1{F0: 736, F1: 893, F2: 42}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v94
    _ = v95
    _ = v96
}
