package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 string; F2 bool }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 int8; F1 uint }

func main() {
    var v1 int32 = 244
    var v2 int8 = 126
    var v3 uint = 729
    v4 := make(map[string]int)
    v5 := []uint32{114}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v6 int64 = int64(v1)
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        switch "a" {
        case "hello":
            fmt.Println(v2, fi7)
        case "a":
            v3 = v3
        }
        v8 := GenS1{F0: 42, F1: 32767}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if ((v1 ^ v1) == v1) {
            delete(v4, "\n\t")
        } else {
            if (!true) {
                if true {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 int16 = -32768
                        _ = fi10
                        _ = v11
                    }
                    v3 += v3
                    var v12 int8 = 1
                    _ = v12
                } else {
                    v13 := GenS1{F0: 1000, F1: 32767}
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 float32 = float32(1.5)
                        _ = fi14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 7; fi16++ {
                        v17 := GenS2{F0: "", F1: "a"}
                        if ((v3 - (^(v3))) == v3) {
                            var v18 uint16 = 65535
                            v1 = 1
                            var v19 float64 = (float64((1.5 + 1.5)) * 3.14159)
                            switch fi16 {
                            case 2147483647:
                                _ = len(v4)
                            case 127:
                                if (!false) {
                                    v20 := GenS1{F0: 32767, F1: v18}
                                    _ = v20
                                }
                            case 32767:
                                v3 = (v3 ^ (^(v3)))
                            default:
                                var v21 int32 = v1
                                _ = v21
                            }
                            _ = v18
                            _ = v19
                        } else {
                            var v22 float32 = 1.5
                            _ = v22
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = v13
                }
                switch "hello" {
                case "a":
                    var v23 float32 = (-2.25 * 1e10)
                    _ = v23
                }
            } else {
                var v24 uint8 = 1
                if true {
                    var v25 bool = true
                    switch 128 {
                    case 424:
                        v26 := GenS3{F0: (v2 << ((v3) & 63))}
                        _ = v26
                    case -1:
                        v1 += 42
                    }
                    _ = v25
                }
                switch 32767 {
                case 4294967295:
                    switch "hello" {
                    case "foo bar":
                        if false {
                            if (int16(v24) == 42) {
                                var v27 int8 = v2
                                fmt.Println(fi9, v24, v24)
                                if ((((v3 - v3) > v3) && (1e10 < (-2.25 + -2.25))) && (true && (3.14159 == 0.0))) {
                                    v5 = append(v5, 571)
                                } else {
                                    v1 = int32(fi9)
                                    v27 += -121
                                    var v28 int32 = v1
                                    _ = v28
                                }
                                _ = v27
                            }
                        } else {
                            for fi29 := 0; fi29 < 12; fi29++ {
                                var v30 int = fi29
                                v31 := GenS3{F0: v2}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            v1 = v1
                            fmt.Println(fi9, v3, v2)
                        }
                    case "a":
                        v1 += (v1 % ((431) | 1))
                    case "hello":
                        v1 = 1
                    }
                case -9223372036854775808:
                    var v32 float32 = ((float32(1e10) - float32(3.14159)) + ((((1e10 + (-2.25 / 0.1)) - -2.25) + (0.0 - ((3.14159 / 1.5) - 1e10))) * 0.0))
                    _ = v32
                case 2147483647:
                    switch fi9 {
                    case 518:
                        if true {
                            for fi33 := 0; fi33 < 9; fi33++ {
                                _ = len(v5)
                                var v34 int16 = -528
                                var v35 int = fi9
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            if (((float64((1.5 + float32(1e10))) * ((0.0 - 0.0) - (1.5 / float64(-2.25)))) != 0.0) || (32767 > 198)) {
                                v24 = (v24 & v24)
                            }
                        }
                    case -189:
                        for fi36 := 0; fi36 < 4; fi36++ {
                            fmt.Println(v2, fi9)
                            switch fi9 {
                            case 42:
                                fmt.Println(v1, v2, fi9)
                                v2 = v2
                            default:
                                var v37 int = int(v24)
                                _ = v37
                            }
                            _ = fi36
                        }
                    }
                default:
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                }
                _ = v24
            }
            for fi38 := 0; fi38 < 5; fi38++ {
                switch (fi38 & fi9) {
                case 9223372036854775807:
                    if (3.14159 > float64(1e10)) {
                        v2 = v2
                    } else {
                        switch ("" + "a") {
                        case "a":
                            if false {
                                v3 += (v3 / (((v3 & (v3 * (v3 + 4294967295)))) | 1))
                                v1 += -388
                                var v39 int64 = -1
                                _ = v39
                            } else {
                                var v40 int8 = (v2 + -124)
                                _ = len(v4)
                                switch (fi9 + (^(fi9))) {
                                case 628:
                                    switch fi38 {
                                    case 0:
                                        v2 = v2
                                        var v41 uint64 = 1
                                        _ = v41
                                    case 32768:
                                        switch "foo bar" {
                                        case "":
                                            var v42 uint64 = 42
                                            _ = v42
                                        default:
                                            fmt.Println(v3)
                                            v43 := GenS4{F0: (("hello" + "hello") + "foo bar"), F1: 0}
                                            if false {
                                                fmt.Println(v40, v2)
                                            } else {
                                                var v44 int = (^(fi38))
                                                _ = v44
                                            }
                                            var v45 uint = v3
                                            _ = v43
                                            _ = v45
                                        }
                                    case -758:
                                        v2 = (^(v2))
                                    }
                                }
                                v46 := v4[("hello" + "\n\t")]
                                _ = v40
                                _ = v46
                            }
                        case "hello":
                            var v47 float32 = 0.1
                            _ = v47
                        default:
                            v48 := GenS1{F0: -735, F1: 143}
                            _ = v48
                        }
                        v2 = 0
                        var v49 int32 = 1
                        _ = v49
                    }
                default:
                    switch 42 {
                    case 1:
                        v50 := GenS5{F0: v3}
                        _ = v50
                    case 128:
                        if true {
                            v3 += (v3 * v3)
                            if ((v2 * v2) <= -30) {
                                _ = len(v4)
                                fmt.Println(v2, v1, fi9)
                            } else {
                                if false {
                                    v51 := GenS6{F0: v3}
                                    v3 += (v3 + 18446744073709551615)
                                    _ = v51
                                } else {
                                    var v52 uint64 = 2
                                    v5 = append(v5, 341)
                                    v53 := GenS5{F0: 42}
                                    _ = v52
                                    _ = v53
                                }
                                v54 := GenS7{F0: ((v1 + (v1 * 42)) >= 0), F1: "\n\t", F2: ((false && ("foo bar" >= ("a" + "hello"))) && true)}
                                if (true && (!(true && (int(v3) == -517)))) {
                                    v1 += (v1 >> ((v3) & 63))
                                    var v55 int16 = 0
                                    _ = v55
                                } else {
                                    if true {
                                        switch -9223372036854775808 {
                                        case 32767:
                                            if false {
                                                fmt.Println(fi9, v3, fi9)
                                                if (int16(fi9) >= -32768) {
                                                    if (!true) {
                                                        v1 = -1
                                                        if ((v2 << (((v3 - (v3 / ((v3) | 1)))) & 63)) <= (v2 & 8)) {
                                                            var v56 int32 = v1
                                                            _ = len(v4)
                                                            var v57 uint32 = 2147483647
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            if (2147483647 >= 600) {
                                                                switch "" {
                                                                case "foo bar":
                                                                    v1 = int32(fi9)
                                                                case "":
                                                                    var v58 int32 = 515
                                                                    _ = v58
                                                                }
                                                            } else {
                                                                var v59 uint64 = 1
                                                                v4[("hello" + "")] = int(v1)
                                                                v2 = (v2 % ((v2) | 1))
                                                                _ = v59
                                                            }
                                                            v1 = -2147483648
                                                            fmt.Println(v1, v2)
                                                        }
                                                    } else {
                                                        switch ("" + ("foo bar" + "")) {
                                                        case "\n\t":
                                                            var v60 string = ("a" + "")
                                                            _ = v60
                                                        case "a":
                                                            v2 = (v2 >> (((v3 % ((65535) | 1))) & 63))
                                                        case "foo bar":
                                                            switch fi9 {
                                                            case 4294967295:
                                                                v3 = 102
                                                            case -9223372036854775808:
                                                                var v61 float64 = (((0.0 + ((3.14159 + ((0.0 + -2.25) * 1e10)) * 1.5)) / 1.5) + 3.14159)
                                                                _ = v61
                                                            }
                                                        }
                                                        var v62 uint16 = 1
                                                        if ((!false) && true) {
                                                            if (false && (((float64(((-2.25 - ((0.0 * 3.14159) * 0.1)) * 3.14159)) * 0.0) * 0.0) == (3.14159 / 1e10))) {
                                                                v2 += int8(v1)
                                                            }
                                                            v4[""] = fi9
                                                        }
                                                        _ = v62
                                                    }
                                                    var v63 int = -2147483648
                                                    _ = v63
                                                } else {
                                                    fmt.Println(fi38)
                                                    var v64 int8 = (v2 % (((v2 & (v2 % ((118) | 1)))) | 1))
                                                    _ = v64
                                                }
                                                v3 = v3
                                                fmt.Println(v2)
                                                if (200 > -32768) {
                                                    var v65 uint64 = 1
                                                    v2 = -1
                                                    var v66 bool = (-2.25 == 0.1)
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    fmt.Println(v3, fi9)
                                                }
                                                switch ("\n\t" + "a") {
                                                case "foo bar":
                                                    v3 = (v3 / (((v3 & v3)) | 1))
                                                }
                                                switch 42 {
                                                case -1:
                                                    v3 -= (v3 | v3)
                                                case 42:
                                                    v1 -= -487
                                                }
                                                switch "a" {
                                                case "hello":
                                                    var v67 uint8 = 1
                                                    _ = v67
                                                case "foo bar":
                                                    if ((14 >= (v3 / (((v3 ^ (v3 | (v3 | v3)))) | 1))) || (!(61 >= 960))) {
                                                        fmt.Println(v1, fi9, v2)
                                                        if false {
                                                            if false {
                                                                switch (fi38 - fi38) {
                                                                case 608:
                                                                    switch 0 {
                                                                    case 846:
                                                                        v68 := v4["foo bar"]
                                                                        _ = v68
                                                                    }
                                                                case -129:
                                                                    if (float32(float64(0.0)) >= 3.14159) {
                                                                        v3 -= v3
                                                                        v2 += v2
                                                                        v2 -= 36
                                                                    }
                                                                    v2 -= (v2 << (((v3 + v3)) & 63))
                                                                    v1 = (v1 + int32(fi9))
                                                                case -9223372036854775808:
                                                                    v3 -= (v3 ^ (^(v3)))
                                                                }
                                                                v1 = 2147483647
                                                            }
                                                            var v69 uint32 = 2
                                                            if (!false) {
                                                                if (!true) {
                                                                    var v70 uint32 = v69
                                                                    delete(v4, "a")
                                                                    _ = v70
                                                                }
                                                            }
                                                            _ = v69
                                                        } else {
                                                            v71 := GenS8{F0: (v2 << ((v3) & 63)), F1: 0}
                                                            v4["a"] = 32768
                                                            if (-1 > (v1 << ((v3) & 63))) {
                                                                if (("a" < ("foo bar" + "a")) && false) {
                                                                    v3 -= v3
                                                                    v1 = (v1 / ((v1) | 1))
                                                                    var v72 int64 = -9223372036854775808
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        switch "a" {
                                                                        case "hello":
                                                                            switch fi9 {
                                                                            case -635:
                                                                                v1 += int32(v2)
                                                                            }
                                                                        }
                                                                    case "":
                                                                        if false {
                                                                            v72 = 1
                                                                            var v73 uint16 = 556
                                                                            v2 = 96
                                                                            _ = v73
                                                                        } else {
                                                                            if false {
                                                                                var v74 int32 = (v1 % ((v1) | 1))
                                                                                v74 = int32(fi38)
                                                                                fmt.Println(v1, v3)
                                                                                fmt.Println(fi9, fi38, v2)
                                                                                _ = v74
                                                                            } else {
                                                                                v2 -= v2
                                                                                v75 := v4["a"]
                                                                                _ = v75
                                                                            }
                                                                            var v76 uint8 = 42
                                                                            if (!true) {
                                                                                fmt.Println(v3, v76)
                                                                            } else {
                                                                                v2 -= v2
                                                                            }
                                                                            _ = v76
                                                                        }
                                                                    case "a":
                                                                        v1 = v1
                                                                    }
                                                                    _ = v72
                                                                } else {
                                                                    delete(v4, ("" + (((((("a" + (("foo bar" + "") + "foo bar")) + "hello") + "a") + ("a" + "a")) + "hello") + ("\n\t" + "\n\t"))))
                                                                }
                                                                v1 += -1
                                                                if (145 <= 287) {
                                                                    v1 -= (v1 << (((^(v3))) & 63))
                                                                } else {
                                                                    if (-32768 < -1) {
                                                                        switch ("foo bar" + "hello") {
                                                                        case "a":
                                                                            var v77 string = "\n\t"
                                                                            delete(v4, "foo bar")
                                                                            fmt.Println(fi9, fi38)
                                                                            _ = v77
                                                                        case "hello":
                                                                            fmt.Println(fi38)
                                                                        case "\n\t":
                                                                            v78 := GenS9{F0: (^(v2)), F1: 355}
                                                                            _ = v78
                                                                        default:
                                                                            var v79 int8 = (v2 >> ((v3) & 63))
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                    var v80 bool = (false || false)
                                                                    if v80 {
                                                                        switch ("\n\t" + "foo bar") {
                                                                        case "":
                                                                            v81 := GenS7{F0: ((v1 | (v1 / ((v1) | 1))) <= (v1 - v1)), F1: (("\n\t" + "hello") + "a"), F2: (958 <= fi9)}
                                                                            _ = v81
                                                                        default:
                                                                            v5 = append(v5, 2147483647)
                                                                        }
                                                                        if (!false) {
                                                                            if true {
                                                                                var v82 int8 = v2
                                                                                delete(v4, ("" + ("\n\t" + "")))
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        var v83 int8 = v2
                                                                        _ = v83
                                                                    } else {
                                                                        var v84 int32 = v1
                                                                        _ = v84
                                                                    }
                                                                    v2 = v2
                                                                    _ = v80
                                                                }
                                                            }
                                                            _ = v71
                                                        }
                                                        _ = len(v5)
                                                        if false {
                                                            v2 -= v2
                                                            var v85 int64 = -9223372036854775808
                                                            _ = v85
                                                        }
                                                    }
                                                default:
                                                    v5 = append(v5, 42)
                                                }
                                                var v86 float32 = 1.5
                                                _ = v86
                                            }
                                        case 42:
                                            v87 := GenS3{F0: (v2 - v2)}
                                            _ = v87
                                        case -1:
                                            var v88 int = fi38
                                            _ = v88
                                        }
                                        var v89 uint16 = 255
                                        _ = v89
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 2147483647
                                        }
                                        v2 += (v2 | (v2 + v2))
                                        var v90 int16 = 32767
                                        _ = v90
                                    }
                                    fmt.Println(fi9)
                                }
                                _ = v54
                            }
                        }
                    default:
                        v2 = int8(v3)
                    }
                }
                switch ("hello" + "\n\t") {
                case "a":
                    var v91 uint8 = 2
                    _ = v91
                case "\n\t":
                    var v92 int64 = -1
                    _ = v92
                }
                _ = fi38
            }
            v1 -= (-(v1))
        }
        _ = fi9
    }
    var v93 float32 = 0.1
    var v94 float32 = v93
    var v95 string = ("" + "\n\t")
    v96 := GenS3{F0: v2}
    fmt.Println(v94, v2)
    v94 += (v94 + v93)
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    fmt.Println(v2)
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    switch 547 {
    case 676:
        var v97 string = v95
        _ = v97
    case -742:
        fmt.Println(v93, v2, v95)
    case 1:
        v2 += (v2 & (v2 * v2))
    default:
        var v98 uint16 = 211
        _ = v98
    }
    var v99 string = v95
    var v100 int32 = -117
    for fi101 := 0; fi101 < 15; fi101++ {
        fmt.Println(v2, fi101, v1)
        _ = fi101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v99
}
