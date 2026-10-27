package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 uint8; F2 int64 }
type GenS8 struct { F0 bool; F1 int64; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 string; F1 string; F2 uint64 }

func main() {
    v1 := []string{("a" + "\n\t"), ""}
    v2 := make(map[string]uint32)
    var v3 uint64 = 0
    v4 := make(map[string]bool)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 4294967295 {
    case -1:
        var v5 int16 = -1
        _ = v5
    case 128:
        fmt.Println(v3)
    default:
        for fi6 := 0; fi6 < 19; fi6++ {
            delete(v2, "a")
            var v7 float32 = -2.25
            if (1 > -32768) {
                if (32767 > 418) {
                    for fi8 := 0; fi8 < 10; fi8++ {
                        v7 = float32((0.0 * 0.0))
                        _ = fi8
                    }
                    fmt.Println(fi6, v3, v7)
                    fmt.Println(v3, v7, v7)
                }
                switch (^(fi6)) {
                case 127:
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                case 128:
                    switch "" {
                    case "a":
                        for fi9 := 0; fi9 < 7; fi9++ {
                            fmt.Println(fi9)
                            fmt.Println(fi9)
                            if (!(!true)) {
                                v10 := GenS1{F0: 2, F1: 2147483647}
                                if false {
                                    switch "hello" {
                                    case "foo bar":
                                        v11 := GenS2{F0: uint16(v3), F1: 855, F2: 0}
                                        _ = v11
                                    case "\n\t":
                                        if ((2 != 1) && false) {
                                            var v12 int32 = -2147483648
                                            _ = v12
                                        }
                                    case "hello":
                                        switch ("hello" + ("hello" + "")) {
                                        case "hello":
                                            var v13 uint64 = v3
                                            _ = v13
                                        case "foo bar":
                                            var v14 uint64 = (v3 + v3)
                                            _ = v14
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v15 int = -2147483648
                                        _ = v15
                                    default:
                                        v3 += (v3 - (v3 >> ((255) & 63)))
                                    }
                                    var v16 int16 = -308
                                    fmt.Println(fi9, fi9, fi9)
                                    v17 := GenS1{F0: 830, F1: 1}
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch 32768 {
                                    case 32768:
                                        if true {
                                            if true {
                                                var v18 int = fi9
                                                if ((255 < 42) || (513 >= (v3 | v3))) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v19 := GenS1{F0: 32767, F1: 711}
                                                        _ = v19
                                                    default:
                                                        if ((v3 - v3) <= (^(v3))) {
                                                            delete(v2, "\n\t")
                                                            v3 += (v3 - v3)
                                                            switch "foo bar" {
                                                            case "a":
                                                                if ((v3 * 1) != (^(v3))) {
                                                                    switch fi6 {
                                                                    case -2147483648:
                                                                        switch fi9 {
                                                                        case -9223372036854775808:
                                                                            _ = len(v1)
                                                                        case 32768:
                                                                            v3 = 1
                                                                        case -259:
                                                                            if false {
                                                                                var v20 uint = 255
                                                                                _ = v20
                                                                            } else {
                                                                                v21 := GenS3{F0: 0, F1: 42}
                                                                                v7 = ((v7 - v7) + (float32(1.5) * (float32((0.0 - float64(0.1))) - 1.5)))
                                                                                v22 := GenS4{F0: 0}
                                                                                _ = v21
                                                                                _ = v22
                                                                            }
                                                                        default:
                                                                            fmt.Println(v3, fi9)
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v18 = 9223372036854775807
                                                                        if false {
                                                                            v3 += v3
                                                                        }
                                                                        v23 := GenS3{F0: -128, F1: 255}
                                                                        _ = v23
                                                                    }
                                                                    if (4294967295 != 18446744073709551615) {
                                                                        fmt.Println(fi6)
                                                                        v24 := v2["foo bar"]
                                                                        _ = v24
                                                                    } else {
                                                                        v7 = (float32(1e10) - (v7 * v7))
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = ("a" + "hello")
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                if false {
                                                    var v25 uint32 = 384
                                                    v26 := GenS2{F0: 157, F1: -2147483648, F2: 0}
                                                    var v27 float64 = 0.1
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                } else {
                                                    if false {
                                                        v3 += (v3 - 1)
                                                        if false {
                                                            if false {
                                                                v28 := GenS1{F0: 42, F1: 42}
                                                                var v29 uint64 = v3
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v28
                                                                _ = v29
                                                            } else {
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v30 float32 = v7
                                                                    var v31 int16 = -32768
                                                                    v3 = (^(v3))
                                                                    _ = v30
                                                                    _ = v31
                                                                case "a":
                                                                    var v32 int32 = 273
                                                                    _ = v32
                                                                }
                                                            }
                                                            switch fi6 {
                                                            case 42:
                                                                _ = len(v1)
                                                            case -9223372036854775808:
                                                                if ((false || (false || false)) && false) {
                                                                    if true {
                                                                        var v33 uint64 = 4294967295
                                                                        _ = v33
                                                                    } else {
                                                                        v3 = 2
                                                                    }
                                                                    v3 += v3
                                                                    if ((853 <= 1) || (1000 <= 1000)) {
                                                                        v18 += fi9
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v3 += (v3 % ((v3) | 1))
                                                                        case "a":
                                                                            fmt.Println(v18, v18, v7)
                                                                        }
                                                                        v7 = (v7 / v7)
                                                                    }
                                                                } else {
                                                                    var v34 int32 = 42
                                                                    var v35 int16 = 836
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            default:
                                                                if true {
                                                                    var v36 bool = true
                                                                    var v37 uint32 = 2
                                                                    _ = v36
                                                                    _ = v37
                                                                }
                                                            }
                                                            v38 := GenS4{F0: 100}
                                                            _ = v38
                                                        }
                                                        var v39 int = 0
                                                        _ = v39
                                                    } else {
                                                        v3 -= 956
                                                    }
                                                    switch (fi9 >> ((4294967295) & 63)) {
                                                    case -1:
                                                        if (true || true) {
                                                            switch 4294967295 {
                                                            case -1:
                                                                delete(v4, ("foo bar" + "a"))
                                                            case -9223372036854775808:
                                                                var v40 uint8 = 255
                                                                _ = v40
                                                            case 850:
                                                                var v41 uint16 = 868
                                                                _ = v41
                                                            }
                                                            v7 -= 3.14159
                                                            var v42 uint16 = 255
                                                            _ = v42
                                                        }
                                                    case 127:
                                                        switch "hello" {
                                                        case "":
                                                            if false {
                                                                delete(v2, ("foo bar" + "a"))
                                                                v3 = (v3 ^ (v3 & v3))
                                                            }
                                                        case "foo bar":
                                                            if len(v1) > 0 {
                                                                v1[0] = "\n\t"
                                                            }
                                                        case "a":
                                                            var v43 int32 = -105
                                                            _ = v43
                                                        default:
                                                            if (("" + "hello") < "foo bar") {
                                                                if ("hello" != "a") {
                                                                    var v44 uint64 = v3
                                                                    _ = v44
                                                                }
                                                            } else {
                                                                switch -2147483648 {
                                                                case 32768:
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v45 int8 = -1
                                                                        _ = v45
                                                                    }
                                                                case -26:
                                                                    if (-9223372036854775808 <= -9223372036854775808) {
                                                                        switch fi6 {
                                                                        case 127:
                                                                            fmt.Println(v3)
                                                                        default:
                                                                            if true {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ("" + "hello")
                                                                                }
                                                                                fmt.Println(v7, v3)
                                                                            }
                                                                        }
                                                                        fmt.Println(v18, v7, v3)
                                                                    }
                                                                    v3 += v3
                                                                    var v46 int = fi9
                                                                    _ = v46
                                                                }
                                                            }
                                                            v47 := v4[(("" + "a") + "foo bar")]
                                                            v48 := GenS4{F0: -1}
                                                            if false {
                                                                var v49 uint64 = 18446744073709551615
                                                                _ = v49
                                                            } else {
                                                                if (-577 <= 508) {
                                                                    v7 -= ((v7 * (((v7 - 1e10) / v7) * 1e10)) - v7)
                                                                    var v50 string = ((("hello" + "") + "a") + ("foo bar" + "a"))
                                                                    if false {
                                                                        switch fi6 {
                                                                        case 42:
                                                                            v3 = (v3 * (v3 / ((18446744073709551615) | 1)))
                                                                        case 32767:
                                                                            fmt.Println(v7, fi6)
                                                                        case -2147483648:
                                                                            var v51 int64 = -694
                                                                            _ = v51
                                                                        default:
                                                                            if (1.5 > 0.1) {
                                                                                v18 = 32768
                                                                                var v52 uint32 = 1
                                                                                _ = v52
                                                                            } else {
                                                                                _ = len(v4)
                                                                                v47 = false
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi9, v7)
                                                                        v53 := v4[""]
                                                                        if true {
                                                                            var v54 int = 2147483647
                                                                            var v55 float64 = 1e10
                                                                            v47 = v53
                                                                            v56 := GenS3{F0: 100, F1: 0}
                                                                            v55 = (float64(v7) - float64(float32(float64(3.14159))))
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                        _ = v53
                                                                    }
                                                                    _ = v50
                                                                }
                                                                v57 := GenS1{F0: 65535, F1: -724}
                                                                _ = v57
                                                            }
                                                            v3 -= 4294967295
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    default:
                                                        if (1 <= 128) {
                                                            if (333 <= 42) {
                                                                if (v7 == (3.14159 * (v7 - v7))) {
                                                                    v3 -= v3
                                                                } else {
                                                                    fmt.Println(fi6)
                                                                }
                                                            } else {
                                                                v3 = 2
                                                                v58 := GenS5{F0: 77, F1: v3}
                                                                _ = v58
                                                            }
                                                            v7 -= v7
                                                            v7 += (v7 * 1.5)
                                                            var v59 uint64 = 2
                                                            _ = v59
                                                        }
                                                        switch ((("\n\t" + "a") + "") + "\n\t") {
                                                        case "hello":
                                                            v7 = v7
                                                        }
                                                        delete(v2, "foo bar")
                                                    }
                                                }
                                                _ = v18
                                            }
                                            if true {
                                                var v60 float32 = 1e10
                                                v3 -= v3
                                                _ = v60
                                            }
                                            var v61 int8 = 64
                                            _ = v61
                                        } else {
                                            delete(v2, ("a" + "\n\t"))
                                            if true {
                                                v62 := GenS6{F0: -9223372036854775808}
                                                _ = len(v1)
                                                v1 = append(v1, "hello")
                                                _ = v62
                                            }
                                        }
                                    case 2147483647:
                                        var v63 string = "\n\t"
                                        _ = v63
                                    case 0:
                                        fmt.Println(fi9, fi6, fi9)
                                    }
                                    v7 = (0.0 - v7)
                                }
                                var v64 float32 = float32(((1.5 - 1e10) / (((1.5 + (0.1 - 3.14159)) * 1e10) - ((-2.25 - -2.25) / 1.0))))
                                v3 += (v3 << ((2) & 63))
                                var v65 uint = 2
                                _ = v10
                                _ = v64
                                _ = v65
                            } else {
                                v7 += ((v7 + (v7 + 0.0)) + v7)
                                var v66 uint32 = 2147483648
                                switch (fi9 / ((fi9) | 1)) {
                                case -938:
                                    var v67 int8 = -122
                                    _ = v67
                                case 32767:
                                    v68 := GenS3{F0: 75, F1: uint16(v3)}
                                    _ = v68
                                }
                                _ = v66
                            }
                            _ = fi9
                        }
                    }
                }
                var v69 uint64 = v3
                _ = v69
            }
            _ = fi6
            _ = v7
        }
    }
    v3 -= v3
    switch 163 {
    case 127:
        v3 = v3
    case 917:
        for fi70 := 0; fi70 < 19; fi70++ {
            var v71 float32 = 0.0
            _ = fi70
            _ = v71
        }
    case 804:
        switch 9223372036854775807 {
        case 11:
            v3 -= (v3 | v3)
        case 4294967295:
            var v72 int8 = 1
            _ = v72
        default:
            switch "" {
            case "\n\t":
                if false {
                    for fi73 := 0; fi73 < 5; fi73++ {
                        v74 := GenS7{F0: (2147483647 == 2), F1: 21, F2: -167}
                        _ = fi73
                        _ = v74
                    }
                    switch "" {
                    case "a":
                        for fi75 := 0; fi75 < 5; fi75++ {
                            fmt.Println(fi75, fi75)
                            v76 := GenS8{F0: false, F1: 263, F2: 127}
                            _ = fi75
                            _ = v76
                        }
                    case "foo bar":
                        var v77 uint32 = 65535
                        _ = v77
                    case "hello":
                        v3 += 1
                    default:
                        v78 := GenS9{F0: v3, F1: 42}
                        _ = v78
                    }
                    var v79 uint = 18446744073709551615
                    _ = v79
                }
            case "hello":
                var v80 uint8 = 128
                _ = v80
            default:
                var v81 uint64 = v3
                _ = v81
            }
        }
    default:
        var v82 int16 = 1
        _ = v82
    }
    for fi83 := 0; fi83 < 2; fi83++ {
        var v84 bool = false
        var v85 uint64 = (^(v3))
        _ = fi83
        _ = v84
        _ = v85
    }
    v86 := GenS10{F0: "", F1: "\n\t", F2: 388}
    var v87 string = "foo bar"
    for fi88 := 0; fi88 < 16; fi88++ {
        var v89 bool = false
        var v90 float32 = 3.14159
        var v91 int = (fi88 / ((fi88) | 1))
        _ = fi88
        _ = v89
        _ = v90
        _ = v91
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v86
    _ = v87
}
