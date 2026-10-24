package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64; F1 string; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 int; F1 int; F2 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int8; F1 uint16 }

func main() {
    var v1 string = "hello"
    v2 := []string{((v1 + "hello") + (("a" + v1) + v1)), "hello"}
    var v3 uint8 = 2
    var v4 uint64 = 31
    v5 := GenS1{F0: 2}
    switch v1 {
    case "\n\t":
        var v6 uint16 = 508
        _ = v6
    case "a":
        v7 := GenS1{F0: 82}
        _ = v7
    }
    fmt.Println(v3, v3, v1)
    var v8 string = "a"
    v3 += v3
    if false {
        v9 := GenS2{F0: 659}
        switch (v8 + v1) {
        case "":
            if false {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        var v10 bool = (v3 <= v3)
                        for fi11 := 0; fi11 < 14; fi11++ {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                if v10 {
                                    fmt.Println(v10)
                                    v13 := GenS3{F0: 0, F1: 915}
                                    v8 = v8
                                    _ = v13
                                } else {
                                    v8 = ("foo bar" + v8)
                                    var v14 int32 = 1
                                    var v15 int16 = -1
                                    _ = v14
                                    _ = v15
                                }
                                v3 = 128
                                _ = fi12
                            }
                            v10 = v10
                            _ = fi11
                        }
                        for fi16 := 0; fi16 < 11; fi16++ {
                            switch v8 {
                            case "foo bar":
                                v4 = 1
                            }
                            if (v1 > v8) {
                                var v17 uint64 = v4
                                fmt.Println(v10, fi16, v17)
                                var v18 uint8 = v3
                                _ = v17
                                _ = v18
                            } else {
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    var v20 int32 = -1
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    switch (fi19 ^ 127) {
                                    case 2147483647:
                                        if v10 {
                                            if ((1000 > 0) && v10) {
                                                v21 := GenS1{F0: 1}
                                                var v22 uint64 = (v4 + 2)
                                                _ = v21
                                                _ = v22
                                            } else {
                                                v4 = v4
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v23 int64 = -9223372036854775808
                                            _ = v23
                                        }
                                    case 128:
                                        fmt.Println(fi16, v10, v20)
                                    default:
                                        fmt.Println(fi19, v4, v3)
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                var v24 uint8 = v3
                                var v25 int32 = 572
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS3{F0: 768, F1: 2147483647}
                            _ = fi16
                            _ = v26
                        }
                        _ = v10
                    }
                    var v27 int32 = -1
                    _ = v27
                }
            } else {
                switch "\n\t" {
                case "":
                    if ((^(v4)) == (v4 ^ v4)) {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            for fi29 := 0; fi29 < 14; fi29++ {
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                                v2 = append(v2, v8)
                                fmt.Println(v4)
                                _ = fi29
                            }
                            switch fi28 {
                            case 1:
                                if len(v2) > 0 {
                                    v2[0] = v1
                                }
                            case 2147483647:
                                v30 := GenS3{F0: fi28, F1: 2147483648}
                                _ = v30
                            case 32767:
                                for fi31 := 0; fi31 < 10; fi31++ {
                                    var v32 string = "hello"
                                    v1 = v8
                                    switch v32 {
                                    case "foo bar":
                                        switch v32 {
                                        case "hello":
                                            if true {
                                                var v33 int32 = 42
                                                _ = v33
                                            }
                                        case "a":
                                            var v34 uint16 = 349
                                            _ = v34
                                        case "foo bar":
                                            switch v1 {
                                            case "\n\t":
                                                var v35 uint16 = 0
                                                _ = v35
                                            case "foo bar":
                                                v36 := GenS4{F0: false, F1: 42, F2: 42}
                                                _ = v36
                                            default:
                                                fmt.Println(fi28)
                                            }
                                        default:
                                            if ((42 != -26) && false) {
                                                if len(v2) > 0 {
                                                    v2[0] = (v1 + ((v1 + v1) + ""))
                                                }
                                                switch v8 {
                                                case "hello":
                                                    if (2147483647 <= 2147483647) {
                                                        if false {
                                                            v3 = v3
                                                        } else {
                                                            var v37 string = (v1 + v8)
                                                            _ = v37
                                                        }
                                                        fmt.Println(v3, v4)
                                                        var v38 int16 = -1
                                                        _ = v38
                                                    }
                                                case "foo bar":
                                                    v2 = append(v2, "a")
                                                case "a":
                                                    var v39 int32 = 0
                                                    _ = v39
                                                default:
                                                    v40 := GenS5{F0: 696, F1: "", F2: 255}
                                                    _ = v40
                                                }
                                                var v41 float64 = (3.14159 / float64((float32(1e10) + float32((0.1 + (0.1 + ((0.0 - 3.14159) * (-2.25 * 1.5))))))))
                                                v42 := GenS4{F0: (!(1 <= 225)), F1: 412, F2: -1}
                                                v4 = v4
                                                _ = v41
                                                _ = v42
                                            }
                                            var v43 float32 = 1.5
                                            if true {
                                                v44 := GenS3{F0: (-(fi28)), F1: 42}
                                                if true {
                                                    var v45 float64 = 0.1
                                                    var v46 string = v1
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v47 bool = true
                                                    if len(v2) > 0 {
                                                        v2[0] = "foo bar"
                                                    }
                                                    if v47 {
                                                        fmt.Println(v1)
                                                    } else {
                                                        if (-9223372036854775808 < -1) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v2 = append(v2, "\n\t")
                                                            if false {
                                                                switch "hello" {
                                                                case "":
                                                                    var v48 bool = v47
                                                                    if (v46 <= (v32 + (((("" + v8) + v32) + (v32 + ((((v1 + (v32 + (v1 + v46))) + "foo bar") + "\n\t") + ((v8 + ("" + v8)) + v1)))) + ""))) {
                                                                        if false {
                                                                            switch -1 {
                                                                            case -9223372036854775808:
                                                                                switch v8 {
                                                                                case "":
                                                                                    if v48 {
                                                                                        if true {
                                                                                            var v49 uint = 2
                                                                                            v43 += (v43 / v43)
                                                                                            _ = len(v2)
                                                                                            _ = v49
                                                                                        } else {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = ""
                                                                                            }
                                                                                            v50 := GenS6{F0: v46}
                                                                                            _ = v50
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    _ = len(v2)
                                                                                case "a":
                                                                                    v51 := GenS4{F0: false, F1: 0, F2: 0}
                                                                                    _ = v51
                                                                                }
                                                                                if (fi28 == 128) {
                                                                                    var v52 uint64 = (v4 * (v4 ^ v4))
                                                                                    switch v8 {
                                                                                    case "":
                                                                                        fmt.Println(fi28)
                                                                                    default:
                                                                                        v53 := GenS5{F0: -9223372036854775808, F1: "foo bar", F2: 65535}
                                                                                        _ = v53
                                                                                    }
                                                                                    _ = v52
                                                                                } else {
                                                                                    if true {
                                                                                        switch v1 {
                                                                                        case "hello":
                                                                                            var v54 int = fi28
                                                                                            _ = v54
                                                                                        case "a":
                                                                                            v2 = append(v2, "foo bar")
                                                                                        case "":
                                                                                            v55 := GenS2{F0: 18446744073709551615}
                                                                                            _ = v55
                                                                                        }
                                                                                        fmt.Println(fi31)
                                                                                    } else {
                                                                                        v45 = (1.5 * v45)
                                                                                    }
                                                                                    v48 = false
                                                                                }
                                                                                fmt.Println(v48, fi28, v46)
                                                                            case -237:
                                                                                v56 := GenS6{F0: "foo bar"}
                                                                                _ = v56
                                                                            case 2147483647:
                                                                                v57 := GenS7{F0: 42, F1: (^(v4))}
                                                                                _ = v57
                                                                            default:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = v1
                                                                                }
                                                                            }
                                                                            v58 := GenS8{F0: (fi31 ^ fi28), F1: fi31, F2: 1000}
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                    var v59 int64 = -1
                                                                    _ = v48
                                                                    _ = v59
                                                                case "foo bar":
                                                                    var v60 uint = 1
                                                                    _ = v60
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                if (v47 || (false && (fi28 != 4294967295))) {
                                                                    var v61 uint = 2
                                                                    _ = v61
                                                                }
                                                            }
                                                        }
                                                        _ = len(v2)
                                                        v62 := GenS8{F0: fi28, F1: 42, F2: 42}
                                                        if (2 <= 71) {
                                                            var v63 uint16 = 278
                                                            fmt.Println(v45, v63, v46)
                                                            v64 := GenS9{F0: 123}
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            if (true && v47) {
                                                                var v65 float64 = (v45 / v45)
                                                                _ = v65
                                                            }
                                                            v2 = append(v2, v1)
                                                            fmt.Println(fi28, v45, v3)
                                                            var v66 float64 = v45
                                                            _ = v66
                                                        }
                                                        if v47 {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v67 int8 = -54
                                                                _ = v67
                                                            }
                                                            var v68 string = "hello"
                                                            _ = v68
                                                        }
                                                        _ = len(v2)
                                                        _ = v62
                                                    }
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    var v69 int8 = 100
                                                    var v70 uint64 = (^(v4))
                                                    v70 -= 2
                                                    _ = v69
                                                    _ = v70
                                                }
                                                _ = v44
                                            }
                                            _ = v43
                                        }
                                    case "":
                                        v1 = v32
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            default:
                                switch fi28 {
                                case 873:
                                    v71 := GenS3{F0: fi28, F1: 2}
                                    _ = v71
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                }
                            }
                            for fi72 := 0; fi72 < 9; fi72++ {
                                if false {
                                    var v73 uint32 = 2147483648
                                    v1 = ((v8 + v1) + ((v8 + v8) + v8))
                                    v74 := GenS2{F0: 42}
                                    if true {
                                        v75 := GenS6{F0: "\n\t"}
                                        v76 := GenS4{F0: false, F1: v73, F2: 0}
                                        _ = v75
                                        _ = v76
                                    } else {
                                        var v77 int = fi28
                                        _ = v77
                                    }
                                    v73 = (v73 * (v73 & v73))
                                    _ = v73
                                    _ = v74
                                }
                                _ = fi72
                            }
                            _ = fi28
                        }
                        switch v1 {
                        case "\n\t":
                            if false {
                                if true {
                                    var v78 uint = 2
                                    _ = v78
                                } else {
                                    v8 = (("foo bar" + v8) + v1)
                                    if (!true) {
                                        v79 := GenS10{F0: (v3 / (((v3 * 255)) | 1)), F1: v4, F2: 1}
                                        _ = v79
                                    }
                                }
                            } else {
                                var v80 uint8 = (^(v3))
                                if (((-1 == 100) || false) || false) {
                                    switch (v8 + "a") {
                                    case "\n\t":
                                        _ = len(v2)
                                    case "foo bar":
                                        v81 := GenS6{F0: v8}
                                        _ = v81
                                    }
                                } else {
                                    fmt.Println(v8)
                                    if false {
                                        v4 = 2
                                        if len(v2) > 0 {
                                            v2[0] = v1
                                        }
                                    } else {
                                        v2 = append(v2, v8)
                                        switch v1 {
                                        case "hello":
                                            var v82 int64 = 9223372036854775807
                                            _ = v82
                                        case "foo bar":
                                            fmt.Println(v8, v80)
                                        default:
                                            fmt.Println(v80)
                                        }
                                        v1 = v8
                                    }
                                }
                                _ = v80
                            }
                        case "":
                            var v83 int64 = -464
                            _ = v83
                        case "foo bar":
                            _ = len(v2)
                        }
                        for fi84 := 0; fi84 < 10; fi84++ {
                            for fi85 := 0; fi85 < 9; fi85++ {
                                switch fi85 {
                                case 2147483647:
                                    var v86 uint8 = (v3 | v3)
                                    _ = v86
                                case 61:
                                    fmt.Println(v1)
                                case 1:
                                    switch v1 {
                                    case "a":
                                        switch "a" {
                                        case "a":
                                            switch (fi85 & fi85) {
                                            case 4294967295:
                                                if len(v2) > 0 {
                                                    v2[0] = (((v1 + v8) + v8) + "a")
                                                }
                                            case 1:
                                                fmt.Println(v1, v8, fi84)
                                            default:
                                                v8 = ((v8 + v1) + "")
                                            }
                                        case "hello":
                                            v8 = (((v1 + v1) + "hello") + v1)
                                        }
                                    default:
                                        v4 += (v4 % ((42) | 1))
                                    }
                                default:
                                    var v87 uint = 42
                                    _ = v87
                                }
                                if (false && true) {
                                    v88 := GenS11{F0: v8, F1: 0}
                                    var v89 int = (fi85 - fi85)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v88
                                    _ = v89
                                }
                                v90 := GenS11{F0: "\n\t", F1: 9223372036854775807}
                                _ = fi85
                                _ = v90
                            }
                            if ((!(true && true)) && true) {
                                var v91 float64 = 3.14159
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v92 := GenS12{F0: v3}
                                _ = v91
                                _ = v92
                            }
                            var v93 uint32 = 1
                            _ = fi84
                            _ = v93
                        }
                    }
                case "a":
                    for fi94 := 0; fi94 < 11; fi94++ {
                        for fi95 := 0; fi95 < 20; fi95++ {
                            if true {
                                v96 := GenS13{F0: 100, F1: 1}
                                if false {
                                    var v97 int8 = -128
                                    var v98 uint32 = 4294967295
                                    v2 = append(v2, (v8 + ("foo bar" + v1)))
                                    _ = v97
                                    _ = v98
                                } else {
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        fmt.Println(fi95, v4, fi94)
                                        var v99 int32 = -2147483648
                                        var v100 int = fi95
                                        _ = v100
                                        _ = v99
                                    }
                                }
                                _ = v96
                            }
                            _ = fi95
                        }
                        _ = fi94
                    }
                case "foo bar":
                    v8 = "foo bar"
                default:
                    v1 = v1
                }
                v8 = v8
                var v101 uint32 = 0
                _ = v101
            }
        }
        _ = v9
    } else {
        switch (("\n\t" + v1) + ("foo bar" + "a")) {
        case "a":
            var v102 float64 = -2.25
            _ = v102
        case "":
            for fi103 := 0; fi103 < 10; fi103++ {
                switch v1 {
                case "\n\t":
                    v3 -= (v3 & (^(v3)))
                case "foo bar":
                    var v104 uint = 1
                    _ = v104
                }
                var v105 int16 = 140
                for fi106 := 0; fi106 < 6; fi106++ {
                    fmt.Println(v8)
                    _ = fi106
                }
                _ = fi103
                _ = v105
            }
        case "hello":
            if false {
                fmt.Println(v8, v4)
            }
        default:
            v3 -= 2
        }
    }
    for fi107 := 0; fi107 < 4; fi107++ {
        if false {
            if true {
                switch (v1 + "foo bar") {
                case "foo bar":
                    if true {
                        _ = len(v2)
                        v108 := GenS4{F0: true, F1: uint32(v4), F2: -798}
                        _ = len(v2)
                        _ = v108
                    }
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = ("hello" + (("foo bar" + (((v1 + (v8 + ("hello" + v8))) + v1) + v1)) + v1))
                    }
                }
                switch (-(fi107)) {
                case 590:
                    if (true && (!false)) {
                        v8 = v1
                    }
                case -869:
                    var v109 int64 = 1
                    _ = v109
                default:
                    if true {
                        var v110 string = (v8 + v1)
                        switch v1 {
                        case "hello":
                            v110 = v1
                        case "\n\t":
                            var v111 float64 = (0.1 / 1.5)
                            _ = v111
                        case "a":
                            v112 := GenS10{F0: (v3 - (v3 * 128)), F1: v4, F2: 949}
                            _ = v112
                        default:
                            var v113 int32 = -1
                            _ = v113
                        }
                        _ = v110
                    } else {
                        v4 = v4
                        v1 = (v8 + (("foo bar" + v1) + v8))
                        switch 32767 {
                        case -1:
                            var v114 int8 = 4
                            _ = v114
                        case 4294967295:
                            v1 = ("a" + "")
                        }
                    }
                }
            }
            v3 = 91
        }
        var v115 float64 = float64((1.5 / 3.14159))
        _ = fi107
        _ = v115
    }
    v116 := GenS12{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
