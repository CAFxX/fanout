package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint8; F2 string }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 string; F2 uint16 }
type GenS10 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 int = 1
    var v2 uint32 = 2147483647
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    var v6 bool = true
    switch "a" {
    case "\n\t":
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 42, F1: v1, F2: 1}
            switch -829 {
            case -9223372036854775808:
                for fi9 := 0; fi9 < 14; fi9++ {
                    v2 = v2
                    if v6 {
                        fmt.Println(v2, v6)
                        var v10 int8 = 67
                        _ = v10
                    } else {
                        if v6 {
                            if false {
                                v11 := GenS1{F0: 525, F1: v1, F2: 0}
                                _ = v11
                            } else {
                                if true {
                                    switch "hello" {
                                    case "":
                                        v12 := GenS1{F0: 2, F1: (-(fi9)), F2: 1}
                                        _ = v12
                                    case "hello":
                                        if v6 {
                                            var v13 float64 = 1e10
                                            _ = v13
                                        } else {
                                            var v14 int8 = 43
                                            if (50 > 174) {
                                                if (92 != 2147483647) {
                                                    v15 := GenS2{F0: 255}
                                                    v16 := v3["\n\t"]
                                                    _ = v15
                                                    _ = v16
                                                }
                                            }
                                            var v17 bool = v6
                                            _ = v14
                                            _ = v17
                                        }
                                    case "\n\t":
                                        var v18 uint16 = 14
                                        _ = v18
                                    }
                                    var v19 float64 = (1e10 / 1.5)
                                    _ = v19
                                } else {
                                    var v20 int32 = 0
                                    var v21 uint8 = 192
                                    if (42 != 42) {
                                        v22 := GenS1{F0: uint(v2), F1: (v1 % ((fi9) | 1)), F2: v21}
                                        _ = v22
                                    }
                                    _ = v20
                                    _ = v21
                                }
                            }
                        } else {
                            v5[("" + "hello")] = "hello"
                            v23 := []int{195, (fi9 & 128), (^(fi9))}
                            v2 += 694
                            _ = v23
                        }
                        v6 = (!v6)
                    }
                    switch "foo bar" {
                    case "foo bar":
                        delete(v3, ("hello" + ("a" + "")))
                    case "":
                        v1 = fi9
                    default:
                        if ((!(436 <= 2)) || false) {
                            delete(v4, (fi7 / ((170) | 1)))
                            switch -961 {
                            case 1:
                                switch ("hello" + ("\n\t" + ("" + "hello"))) {
                                case "a":
                                    var v24 uint32 = v2
                                    _ = v24
                                case "":
                                    v25 := GenS2{F0: 42}
                                    _ = v25
                                case "foo bar":
                                    switch fi7 {
                                    case 136:
                                        _ = len(v4)
                                    case -1:
                                        v2 = v2
                                    }
                                    v6 = v6
                                }
                            case -48:
                                fmt.Println(v6)
                                var v26 uint32 = (^(v2))
                                v26 += 42
                                if (248 > 42) {
                                    v6 = v6
                                    switch (fi7 | v1) {
                                    case -255:
                                        var v27 float32 = 3.14159
                                        switch fi9 {
                                        case 42:
                                            switch (v1 ^ v1) {
                                            case 82:
                                                v28 := GenS3{F0: 42, F1: 300}
                                                _ = v28
                                            }
                                        case 9223372036854775807:
                                            var v29 string = ("hello" + ("a" + "\n\t"))
                                            _ = v29
                                        default:
                                            v30 := GenS4{F0: 1}
                                            _ = v30
                                        }
                                        var v31 int32 = 256
                                        v26 = v26
                                        _ = v27
                                        _ = v31
                                    case -734:
                                        var v32 int64 = 1
                                        _ = v32
                                    default:
                                        if ((3.14159 - (0.0 / 3.14159)) <= 1.5) {
                                            v2 = v26
                                            switch fi7 {
                                            case 740:
                                                fmt.Println(v1)
                                                fmt.Println(v6)
                                            case 127:
                                                var v33 uint64 = 695
                                                _ = v33
                                            }
                                        } else {
                                            v34 := []int{5, 127, v1}
                                            v35 := GenS5{F0: 0}
                                            if (234 == 54) {
                                                if (!v6) {
                                                    fmt.Println(v26, v2)
                                                    var v36 uint = 1
                                                    v6 = false
                                                    _ = v36
                                                } else {
                                                    var v37 string = ("a" + ("foo bar" + "hello"))
                                                    var v38 float64 = (0.0 + (1.5 + 0.1))
                                                    var v39 bool = false
                                                    v40 := GenS1{F0: 65535, F1: 9223372036854775807, F2: 255}
                                                    v41 := GenS5{F0: -2147483648}
                                                    _ = len(v34)
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                }
                                                v1 = fi9
                                                v26 -= (v2 & v2)
                                            }
                                            v42 := GenS2{F0: 588}
                                            _ = v34
                                            _ = v35
                                            _ = v42
                                        }
                                    }
                                } else {
                                    v43 := GenS3{F0: 4294967295, F1: 18446744073709551615}
                                    switch fi7 {
                                    case 9223372036854775807:
                                        var v44 uint64 = 4294967295
                                        _ = v44
                                    case 611:
                                        var v45 string = ("foo bar" + "foo bar")
                                        _ = v45
                                    }
                                    if v6 {
                                        var v46 int = 4294967295
                                        v47 := []int{128, v1, fi9}
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = v43
                                }
                                _ = v26
                            case 32767:
                                v48 := GenS4{F0: 255}
                                _ = v48
                            }
                        }
                    }
                    _ = fi9
                }
            case 0:
                var v49 int8 = -1
                _ = v49
            }
            _ = fi7
            _ = v8
        }
    default:
        v50 := []int{v1, v1, (v1 | (v1 * (v1 + v1)))}
        _ = v50
    }
    v6 = (v6 && v6)
    switch (v1 + (v1 - (v1 * v1))) {
    case 240:
        if (-2147483648 >= (^(v1))) {
            for fi51 := 0; fi51 < 4; fi51++ {
                v1 += 9
                var v52 int = fi51
                _ = fi51
                _ = v52
            }
        }
    case 128:
        v53 := GenS4{F0: 255}
        _ = v53
    }
    _ = len(v5)
    if (!(-1 >= 1000)) {
        switch ((("a" + "") + "hello") + "hello") {
        case "foo bar":
            fmt.Println(v1, v1)
        case "\n\t":
            v54 := []int{-33, (v1 >> ((866) & 63)), 1}
            _ = v54
        default:
            for fi55 := 0; fi55 < 2; fi55++ {
                var v56 uint8 = 42
                var v57 float64 = float64((1.5 + float32(1e10)))
                switch ("" + ("\n\t" + "hello")) {
                case "a":
                    if (18446744073709551615 >= 42) {
                        v58 := GenS3{F0: 146, F1: 42}
                        _ = v58
                    }
                case "hello":
                    v59 := GenS1{F0: 271, F1: fi55, F2: (v56 << ((uint(v2)) & 63))}
                    _ = v59
                }
                _ = fi55
                _ = v56
                _ = v57
            }
        }
        v2 -= 0
    } else {
        if false {
            switch "a" {
            case "a":
                for fi60 := 0; fi60 < 2; fi60++ {
                    if false {
                        v1 = (^(fi60))
                    } else {
                        fmt.Println(fi60, fi60, v1)
                        for fi61 := 0; fi61 < 6; fi61++ {
                            var v62 uint8 = 128
                            _ = fi61
                            _ = v62
                        }
                    }
                    var v63 int8 = 1
                    _ = fi60
                    _ = v63
                }
            case "\n\t":
                fmt.Println(v2, v1, v2)
            default:
                v1 += (v1 / (((-(v1))) | 1))
            }
            var v64 uint8 = 43
            v65 := []int{128, (v1 >> ((555) & 63))}
            _ = v64
            _ = v65
        }
    }
    v66 := GenS6{F0: 64, F1: 1, F2: "foo bar"}
    switch (v1 * v1) {
    case -2147483648:
        for fi67 := 0; fi67 < 1; fi67++ {
            var v68 int16 = -1
            for fi69 := 0; fi69 < 5; fi69++ {
                var v70 uint = 0
                if (uint64(v70) <= 1) {
                    var v71 int32 = -1
                    v72 := []int{9223372036854775807}
                    v70 = (v70 + v70)
                    _ = v71
                    _ = v72
                } else {
                    var v73 uint64 = 0
                    _ = v73
                }
                if (fi69 > fi69) {
                    v74 := GenS4{F0: uint(v2)}
                    var v75 uint = 58
                    if v6 {
                        v76 := v3[("\n\t" + "a")]
                        var v77 uint8 = 128
                        _ = v76
                        _ = v77
                    } else {
                        v2 -= (v2 - uint32(v75))
                        v6 = (65535 >= v70)
                        if (!((!v6) || (!(v6 || (!true))))) {
                            v70 = (v70 | v75)
                            v70 += v70
                        } else {
                            v78 := GenS7{F0: v1, F1: 42}
                            _ = v78
                        }
                        var v79 float64 = float64((0.0 * 3.14159))
                        var v80 string = ""
                        _ = v79
                        _ = v80
                    }
                    v68 = v68
                    v2 += (v2 - (v2 + v2))
                    switch (fi67 >> (((v75 / (((v70 | v70)) | 1))) & 63)) {
                    case 42:
                        if v6 {
                            v81 := GenS8{F0: 127}
                            _ = v81
                        }
                    case -9223372036854775808:
                        var v82 int16 = 1000
                        _ = v82
                    }
                    v70 = v75
                    if (!(3.14159 == 0.0)) {
                        switch 128 {
                        case 32767:
                            _ = len(v3)
                        case 830:
                            v83 := v3[(("" + "\n\t") + "a")]
                            _ = v83
                        case 2147483647:
                            fmt.Println(fi69, v68)
                            if false {
                                var v84 uint64 = 18446744073709551615
                                v75 += (v75 * 255)
                                fmt.Println(v70, v75, v68)
                                _ = v84
                            }
                            v85 := GenS9{F0: (-80 != 1), F1: ("" + ("" + "a")), F2: 12}
                            var v86 uint32 = v2
                            _ = v85
                            _ = v86
                        }
                        var v87 float32 = float32(1.5)
                        v6 = (v6 || (false && ((v68 << (((^(v70))) & 63)) >= 1000)))
                        _ = v87
                    }
                    v2 += v2
                    _ = v74
                    _ = v75
                }
                _ = fi69
                _ = v70
            }
            if v6 {
                for fi88 := 0; fi88 < 18; fi88++ {
                    var v89 uint16 = 65535
                    _ = fi88
                    _ = v89
                }
            }
            _ = fi67
            _ = v68
        }
    case 0:
        switch ("a" + ("hello" + (("foo bar" + "a") + "foo bar"))) {
        case "a":
            if (!((42 <= -2147483648) && (v6 && false))) {
                var v90 float64 = -2.25
                var v91 uint = 18446744073709551615
                _ = v90
                _ = v91
            }
        case "":
            v1 = v1
        case "\n\t":
            v92 := GenS10{F0: 0, F1: v2}
            _ = v92
        }
    default:
        for fi93 := 0; fi93 < 10; fi93++ {
            switch "foo bar" {
            case "foo bar":
                v94 := GenS4{F0: 18446744073709551615}
                _ = v94
            default:
                v2 -= (v2 / ((v2) | 1))
            }
            switch ("" + "hello") {
            case "\n\t":
                switch "\n\t" {
                case "":
                    var v95 uint = 879
                    _ = v95
                case "a":
                    v6 = false
                case "hello":
                    if (1 != 0) {
                        v6 = true
                    }
                default:
                    v96 := GenS9{F0: false, F1: "foo bar", F2: 32767}
                    _ = v96
                }
            case "foo bar":
                v6 = false
            default:
                v6 = (false || (((69 != -126) && v6) && (746 > 42)))
            }
            _ = fi93
        }
    }
    fmt.Println(v6)
    v97 := GenS3{F0: 4294967295, F1: 4294967295}
    var v98 string = (("foo bar" + "a") + "foo bar")
    v99 := v3["\n\t"]
    var v100 uint = 777
    v4[-134] = v98
    v6 = false
    _ = len(v4)
    var v101 int8 = -128
    v102 := GenS3{F0: 2, F1: uint64(v2)}
    v103 := v4[(v1 * v1)]
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
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
    _ = v66
    _ = v97
    _ = v98
    _ = v99
}
