package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 int8; F1 int; F2 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int32; F2 string }
type GenS9 struct { F0 int32; F1 uint64; F2 uint }
func genFn0() bool {
    if true {
        var v6 int64 = 0
        _ = v6
    } else {
        var v7 int16 = 32767
        _ = v7
    }
    if false {
        var v8 uint8 = 128
        v9 := GenS1{F0: "a", F1: 4294967295, F2: ""}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: "a", F1: 4294967295, F2: "foo bar"}
        var v12 uint64 = 665
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    _ = v13
    return (((0 == -9223372036854775808) && (0 > 0)) && true)
}
func genFn1(p0 int16, p1 int16) uint32 {
    if ((-2.25 * (0.0 - 1e10)) <= -2.25) {
        var v14 uint = 359
        v15 := GenS2{F0: 127, F1: 336, F2: 0}
        var v16 uint64 = 4294967295
        _ = v14
        _ = v15
        _ = v16
    }
    p1 -= (p1 | (p0 >> ((2) & 63)))
    var v17 float32 = 0.1
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 int8 = 0
            if (131 > 298) {
                var v21 uint32 = 329
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            if true {
                var v23 uint64 = 4294967295
                var v24 string = ""
                _ = v23
                _ = v24
            } else {
                if (!(false && (2 > 0))) {
                    fmt.Println(p0, fi22)
                }
                if true {
                    var v25 uint16 = 65535
                    _ = v25
                } else {
                    v26 := []int{128, 438, 32767}
                    _ = v26
                }
            }
            _ = fi22
        }
        if true {
            v17 -= v17
            p0 = p0
        } else {
            if (122 >= 127) {
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(p0)
                    switch fi18 {
                    case 128:
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            switch (fi18 >> ((915) & 63)) {
                            case 32768:
                                p0 = p0
                            }
                        }
                    case 127:
                        var v28 int16 = (p1 | (p0 * 383))
                        var v29 float64 = 0.0
                        _ = v28
                        _ = v29
                    }
                    if false {
                        switch "hello" {
                        case "":
                            fmt.Println(v17, fi27, fi27)
                            v30 := []int{(fi18 + fi27), (fi18 << ((2) & 63)), (fi18 / (((fi27 / (((fi18 << ((858) & 63))) | 1))) | 1))}
                            v31 := GenS1{F0: "\n\t", F1: 786, F2: ("" + "hello")}
                            switch (fi18 + (^(fi18))) {
                            case 32768:
                                p0 = (p1 << ((1) & 63))
                            case 9223372036854775807:
                                var v32 bool = (true && ((1.5 + (3.14159 + 0.1)) > 3.14159))
                                _ = v32
                            case 0:
                                if false {
                                    var v33 int16 = 140
                                    v34 := GenS2{F0: 0, F1: fi18, F2: 873}
                                    switch fi27 {
                                    case 9223372036854775807:
                                        _ = len(v30)
                                    case 0:
                                        var v35 string = ("foo bar" + ("" + ("foo bar" + "foo bar")))
                                        p0 += 42
                                        _ = v35
                                    default:
                                        if ((1 >= 0) && (v17 < v17)) {
                                            switch -9223372036854775808 {
                                            case 4294967295:
                                                switch "" {
                                                case "a":
                                                    if false {
                                                        if len(v30) > 0 {
                                                            v30[0] = (^(fi27))
                                                        }
                                                    }
                                                case "\n\t":
                                                    v36 := GenS3{F0: (-(fi18)), F1: 18446744073709551615, F2: (0 == -2147483648)}
                                                    _ = v36
                                                case "foo bar":
                                                    var v37 uint8 = 128
                                                    _ = v37
                                                default:
                                                    if (31 >= 127) {
                                                        if false {
                                                            var v38 uint8 = 248
                                                            v39 := GenS4{F0: (fi27 >> ((65535) & 63))}
                                                            if (-2147483648 != 1) {
                                                                var v40 uint8 = v38
                                                                var v41 uint64 = 788
                                                                _ = v40
                                                                _ = v41
                                                            } else {
                                                                fmt.Println(p1, v38, fi18)
                                                            }
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            _ = len(v30)
                                                            switch ("hello" + ("\n\t" + "foo bar")) {
                                                            case "\n\t":
                                                                if len(v30) >= 2 {
                                                                    v30 = v30[1:len(v30)]
                                                                }
                                                            default:
                                                                switch (fi18 ^ fi27) {
                                                                case -2147483648:
                                                                    v33 = (^(v33))
                                                                case -767:
                                                                    v42 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: "hello"}
                                                                    _ = v42
                                                                }
                                                                if len(v30) > 0 {
                                                                    v30[0] = (fi18 & -2147483648)
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            case 42:
                                                v43 := GenS4{F0: (fi18 ^ fi27)}
                                                _ = v43
                                            default:
                                                var v44 int16 = p0
                                                _ = v44
                                            }
                                        } else {
                                            v45 := GenS5{F0: true, F1: 42}
                                            _ = v45
                                        }
                                    }
                                    _ = v33
                                    _ = v34
                                } else {
                                    if len(v30) >= 2 {
                                        v30 = v30[1:len(v30)]
                                    }
                                }
                            default:
                                fmt.Println(v17, fi27, p1)
                            }
                            var v46 string = "foo bar"
                            v46 = v46
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                            var v47 int16 = -1
                            var v48 int64 = -9223372036854775808
                            _ = v30
                            _ = v31
                            _ = v46
                            _ = v47
                            _ = v48
                        case "hello":
                            v17 += (0.1 / v17)
                        }
                        p0 = p1
                        fmt.Println(fi18)
                    }
                    _ = fi27
                }
                fmt.Println(fi18, v17)
            }
        }
        _ = fi18
    }
    for fi49 := 0; fi49 < 15; fi49++ {
        for fi50 := 0; fi50 < 11; fi50++ {
            v51 := []int{}
            _ = fi50
            _ = v51
        }
        if (-2147483648 != -2147483648) {
            v52 := GenS5{F0: (0.0 < 0.1), F1: 950}
            switch 127 {
            case 4294967295:
                v53 := []int{(fi49 * fi49), (fi49 | (-(fi49))), (fi49 >> ((4294967295) & 63))}
                _ = v53
            case 2147483647:
                var v54 uint32 = 170
                _ = v54
            case -9223372036854775808:
                fmt.Println(fi49, v17)
            }
            _ = v52
        }
        _ = fi49
    }
    _ = p0
    _ = p1
    _ = v17
    return 1
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]uint32)
    var v3 uint16 = 42
    v4 := make(map[int]bool)
    var v5 int32 = 0
    for fi55 := 0; fi55 < 19; fi55++ {
        v5 += (^(v5))
        _ = fi55
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        var v57 uint16 = 32767
        if (true && (0 != 65535)) {
            var v58 int = fi56
            _ = v58
        }
        _ = fi56
        _ = v57
    }
    var v59 uint32 = 0
    v5 = (v5 % ((v5) | 1))
    if true {
        switch 9223372036854775807 {
        case 0:
            for fi60 := 0; fi60 < 15; fi60++ {
                for fi61 := 0; fi61 < 10; fi61++ {
                    v5 -= (v5 / ((v5) | 1))
                    if true {
                        fmt.Println(v3)
                        fmt.Println(v5)
                        var v62 uint = 1
                        var v63 int16 = 420
                        _ = v62
                        _ = v63
                    }
                    _ = fi61
                }
                v64 := v2[(fi60 & fi60)]
                _ = fi60
                _ = v64
            }
        default:
            fmt.Println(v5, v3, v1)
        }
        fmt.Println(v5, v59)
    }
    switch (("foo bar" + ("\n\t" + "hello")) + "hello") {
    case "":
        if (v59 == (v59 ^ (v59 - (v59 + (v59 & 85))))) {
            v65 := GenS1{F0: "\n\t", F1: 0, F2: "\n\t"}
            for fi66 := 0; fi66 < 14; fi66++ {
                var v67 uint64 = 1
                switch ((((("\n\t" + ("" + "foo bar")) + (("a" + "") + "\n\t")) + "\n\t") + "\n\t") + "hello") {
                case "\n\t":
                    v1 = v1
                case "":
                    var v68 uint = 0
                    _ = v68
                case "foo bar":
                    var v69 uint32 = 2147483648
                    _ = v69
                }
                v1 = v1
                _ = fi66
                _ = v67
            }
            var v70 int = 4294967295
            _ = v65
            _ = v70
        } else {
            v71 := GenS4{F0: -989}
            _ = v71
        }
    case "\n\t":
        v1 = -1
    }
    var v72 float32 = 3.14159
    for fi73 := 0; fi73 < 11; fi73++ {
        switch "\n\t" {
        case "foo bar":
            for fi74 := 0; fi74 < 20; fi74++ {
                var v75 int8 = v1
                v5 = 0
                _ = fi74
                _ = v75
            }
        case "":
            v76 := GenS3{F0: (fi73 / ((fi73) | 1)), F1: 4294967295, F2: true}
            _ = v76
        case "a":
            v3 = (v3 & v3)
        default:
            v77 := GenS6{F0: 0, F1: 32768}
            _ = v77
        }
        v59 -= v59
        if false {
            if false {
                var v78 float32 = (v72 - 0.0)
                _ = v78
            }
        }
        _ = fi73
    }
    var v79 bool = (128 != 0)
    switch ("\n\t" + "hello") {
    case "\n\t":
        v79 = true
    case "a":
        if (9223372036854775807 == -2147483648) {
            v3 = v3
            var v80 string = ""
            v79 = true
            _ = v80
        } else {
            if v79 {
                v1 = v1
                var v81 int8 = v1
                _ = v81
            } else {
                if v79 {
                    var v82 uint8 = 42
                    _ = v82
                }
                v83 := []int{-9223372036854775808, 9223372036854775807, 790, 522}
                _ = v83
            }
            v79 = ((0 == 28) && v79)
            if false {
                switch "hello" {
                case "":
                    fmt.Println(v59, v72, v5)
                case "a":
                    v84 := []int{-360, -266}
                    _ = v84
                case "foo bar":
                    if ((!(-32768 != 32767)) || false) {
                        fmt.Println(v3, v59, v72)
                    }
                }
                var v85 int = -501
                v5 = -2147483648
                _ = v85
            } else {
                switch -454 {
                case 127:
                    var v86 string = "a"
                    _ = v86
                }
                v87 := GenS7{F0: uint64(v3)}
                if (!false) {
                    v88 := GenS3{F0: 1, F1: 0, F2: v79}
                    switch "a" {
                    case "\n\t":
                        var v89 uint = 1
                        _ = v89
                    case "":
                        var v90 int = -1
                        _ = v90
                    case "a":
                        v91 := GenS8{F0: 465, F1: -2147483648, F2: ""}
                        _ = v91
                    default:
                        switch "\n\t" {
                        case "hello":
                            v59 += 65535
                        case "":
                            fmt.Println(v1, v59, v59)
                        case "foo bar":
                            switch "foo bar" {
                            case "":
                                v92 := GenS9{F0: -577, F1: 42, F2: 4294967295}
                                _ = v92
                            case "a":
                                var v93 uint = 4294967295
                                _ = v93
                            case "\n\t":
                                fmt.Println(v59, v59)
                            default:
                                var v94 uint = 255
                                _ = v94
                            }
                        }
                    }
                    fmt.Println(v72)
                    _ = v88
                }
                _ = v87
            }
        }
    default:
        delete(v4, -986)
    }
    fmt.Println(v3, v59, v3)
    v1 -= -128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v72:", v72)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v72
    _ = v79
}
