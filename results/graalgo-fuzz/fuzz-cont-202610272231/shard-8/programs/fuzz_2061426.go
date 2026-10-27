package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int; F1 uint; F2 uint16 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS8 struct { F0 uint; F1 uint8; F2 int8 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 uint; F1 uint32; F2 int }
type GenS11 struct { F0 int8 }

func main() {
    v1 := []uint32{1}
    var v2 int64 = 1
    var v3 int = 2147483647
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 -= (v3 % (((v3 ^ 42)) | 1))
    if (128 == 42) {
        if false {
            fmt.Println(v3, v3)
            if ((97 < 113) && ((!(-71 > 27)) && true)) {
                if (2147483647 == 42) {
                    for fi4 := 0; fi4 < 3; fi4++ {
                        var v5 int32 = 1
                        _ = fi4
                        _ = v5
                    }
                    v6 := GenS1{F0: 1, F1: 2, F2: 4294967295}
                    _ = v6
                }
            } else {
                v7 := GenS2{F0: 31, F1: 189, F2: 2}
                if true {
                    switch 9223372036854775807 {
                    case 898:
                        fmt.Println(v3, v3)
                    default:
                        switch v3 {
                        case -1:
                            v3 = v3
                        case 23:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case -834:
                            for fi8 := 0; fi8 < 5; fi8++ {
                                fmt.Println(fi8)
                                _ = fi8
                            }
                        }
                    }
                    _ = len(v1)
                } else {
                    if ((!true) && ((471 >= 2147483647) || (!true))) {
                        v9 := GenS2{F0: 1, F1: 1, F2: 0}
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            fmt.Println(fi10, v2)
                            _ = fi10
                        }
                        switch v3 {
                        case 128:
                            for fi11 := 0; fi11 < 6; fi11++ {
                                var v12 int64 = (v2 % (((v2 + (v2 - (v2 % ((v2) | 1))))) | 1))
                                _ = fi11
                                _ = v12
                            }
                        case -171:
                            if (!((("\n\t" == "") || (0.1 >= 0.0)) && (true || false))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    v2 = 1
                                    v13 := GenS3{F0: -431, F1: -126}
                                    v3 = (-(v3))
                                    _ = v13
                                }
                                if true {
                                    for fi14 := 0; fi14 < 1; fi14++ {
                                        for fi15 := 0; fi15 < 12; fi15++ {
                                            v3 -= v3
                                            switch "hello" {
                                            case "a":
                                                if (float64(-2.25) > (0.0 / 0.1)) {
                                                    var v16 int16 = -32768
                                                    var v17 float32 = 1.5
                                                    fmt.Println(v2, v3, v17)
                                                    _ = v16
                                                    _ = v17
                                                }
                                            case "\n\t":
                                                v18 := GenS3{F0: 2147483647, F1: 42}
                                                _ = v18
                                            case "foo bar":
                                                var v19 bool = true
                                                _ = v19
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            _ = fi15
                                        }
                                        v20 := GenS3{F0: -236, F1: -1}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi14
                                        _ = v20
                                    }
                                    fmt.Println(v3, v2)
                                }
                            } else {
                                v3 = 127
                            }
                        case 42:
                            v3 += (^(v3))
                        default:
                            v2 -= v2
                        }
                        fmt.Println(v3, v2)
                    }
                    switch (-(v3)) {
                    case 32767:
                        v2 += v2
                    default:
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 uint = 0
                            _ = fi21
                            _ = v22
                        }
                    }
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 int = (^(fi23))
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v7
            }
        }
        v3 -= -2147483648
        v1 = append(v1, 305)
    } else {
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS3{F0: -1, F1: 100}
                fmt.Println(v2)
                _ = fi26
                _ = v27
            }
            v3 -= v3
            _ = fi25
        }
        v1 = append(v1, 689)
    }
    if true {
        v2 = 608
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v28 uint16 = 255
        _ = v28
    } else {
        fmt.Println(v3, v2)
        var v29 string = "hello"
        v29 = ("foo bar" + (v29 + "hello"))
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 int16 = 0
        var v32 uint = 684
        for fi33 := 0; fi33 < 13; fi33++ {
            if (285 <= 736) {
                var v34 int = fi30
                fmt.Println(v2)
                if (v2 <= (^(v2))) {
                    v31 -= (v31 ^ -1)
                } else {
                    var v35 string = "hello"
                    _ = len(v1)
                    switch "foo bar" {
                    case "\n\t":
                        switch (fi30 & 32767) {
                        case 4294967295:
                            v35 = "\n\t"
                        case -2147483648:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case "hello":
                        v36 := GenS4{F0: (fi33 << (((v32 * v32)) & 63)), F1: (v32 & 57), F2: 74}
                        _ = v36
                    case "foo bar":
                        v37 := GenS5{F0: 2147483647, F1: 0}
                        _ = v37
                    default:
                        switch (fi30 | (v3 / ((32767) | 1))) {
                        case 42:
                            v31 -= v31
                        default:
                            switch (fi33 | fi33) {
                            case -9223372036854775808:
                                var v38 uint64 = 0
                                _ = v38
                            case 128:
                                if (!false) {
                                    v39 := GenS3{F0: 554, F1: 42}
                                    fmt.Println(v32)
                                    switch "hello" {
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            if (1 == 227) {
                                                v34 -= fi33
                                            } else {
                                                var v40 uint32 = 367
                                                var v41 int16 = v31
                                                _ = v40
                                                _ = v41
                                            }
                                        case "hello":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    case "":
                                        var v42 string = ("foo bar" + "\n\t")
                                        _ = v42
                                    default:
                                        var v43 uint64 = 576
                                        _ = v43
                                    }
                                    _ = v39
                                } else {
                                    var v44 int8 = -1
                                    fmt.Println(v44, v32, v34)
                                    _ = v44
                                }
                            default:
                                v32 -= 133
                            }
                        }
                    }
                    _ = v35
                }
                _ = v34
            }
            var v45 int = v3
            if ("a" != "a") {
                v46 := GenS1{F0: 986, F1: 1, F2: 42}
                _ = v46
            } else {
                v31 = v31
                switch "hello" {
                case "hello":
                    switch fi30 {
                    case -1:
                        var v47 int32 = 0
                        _ = v47
                    case 32767:
                        if true {
                            if false {
                                if (false || false) {
                                    v31 += (v31 | v31)
                                    v3 = fi33
                                } else {
                                    if false {
                                        var v48 int16 = v31
                                        v49 := GenS2{F0: 127, F1: 138, F2: 42}
                                        var v50 float64 = 0.0
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v1 = append(v1, 2)
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                        var v51 float32 = 3.14159
                                        var v52 int32 = -815
                                        switch "\n\t" {
                                        case "":
                                            v52 += (v52 << (((v32 % ((v32) | 1))) & 63))
                                        case "a":
                                            if false {
                                                fmt.Println(v45, v51)
                                            }
                                        case "foo bar":
                                            fmt.Println(v51, v31, v45)
                                        }
                                        fmt.Println(v3)
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = len(v1)
                                }
                            }
                        }
                    case 32768:
                        v53 := GenS5{F0: 2147483647, F1: 42}
                        _ = v53
                    default:
                        if (v2 <= (v2 >> ((1) & 63))) {
                            var v54 uint32 = 2
                            v55 := GenS6{F0: 18446744073709551615, F1: int8(v3)}
                            switch v3 {
                            case -9223372036854775808:
                                var v56 uint16 = 0
                                _ = v56
                            default:
                                v57 := GenS7{F0: 0, F1: 2, F2: 943}
                                _ = v57
                            }
                            _ = v54
                            _ = v55
                        } else {
                            if true {
                                fmt.Println(v3)
                            } else {
                                if ((((fi30 ^ fi33) > v45) || (false || true)) && true) {
                                    v32 = 18446744073709551615
                                }
                            }
                        }
                        v1 = append(v1, 2)
                    }
                case "\n\t":
                    v31 -= v31
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            _ = fi33
            _ = v45
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v3)
    v58 := GenS8{F0: 800, F1: 250, F2: 100}
    if false {
        v2 = v2
        switch "hello" {
        case "hello":
            var v59 int8 = 1
            _ = v59
        case "foo bar":
            v3 = -9223372036854775808
        case "":
            for fi60 := 0; fi60 < 5; fi60++ {
                switch fi60 {
                case 128:
                    if true {
                        v3 = v3
                    }
                case 2147483647:
                    v61 := GenS4{F0: v3, F1: 42, F2: 65535}
                    _ = v61
                case 127:
                    for fi62 := 0; fi62 < 1; fi62++ {
                        if ((38 < -1) || false) {
                            var v63 uint = 4294967295
                            _ = v63
                        }
                        v3 -= (fi62 - (v3 << ((255) & 63)))
                        _ = fi62
                    }
                }
                if (1 >= -1) {
                    fmt.Println(v2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if true {
                        v1 = append(v1, 65535)
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    } else {
                        v2 -= v2
                        var v64 uint32 = 374
                        v64 += (^(v64))
                        _ = v64
                    }
                } else {
                    var v65 int16 = -32768
                    var v66 float32 = ((((1.5 - 0.0) * 1.5) * 0.1) * ((3.14159 / ((1.5 * (-2.25 / 1.5)) * (((3.14159 * 0.1) - (-2.25 - 1.5)) + 0.1))) / float32(float64(0.1))))
                    _ = v65
                    _ = v66
                }
                _ = fi60
            }
        }
        v2 -= -683
    }
    for fi67 := 0; fi67 < 9; fi67++ {
        var v68 int8 = 100
        if (false && (42 >= 826)) {
            switch (fi67 % (((v3 >> ((875) & 63))) | 1)) {
            case -206:
                v69 := GenS6{F0: 405, F1: 8}
                _ = v69
            case 894:
                switch ("hello" + "") {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        v68 += v68
                    case "foo bar":
                        fmt.Println(fi67, fi67, v68)
                    case "a":
                        _ = len(v1)
                    default:
                        switch "a" {
                        case "foo bar":
                            v2 += v2
                        case "\n\t":
                            switch "a" {
                            case "\n\t":
                                switch "" {
                                case "a":
                                    switch (v3 % ((-9223372036854775808) | 1)) {
                                    case 1:
                                        for fi70 := 0; fi70 < 10; fi70++ {
                                            if (1 >= 627) {
                                                var v71 uint32 = 2147483647
                                                _ = v71
                                            } else {
                                                if (272 >= 1) {
                                                    v1 = append(v1, 2)
                                                    var v72 int8 = (v68 | -10)
                                                    var v73 int8 = v72
                                                    _ = v72
                                                    _ = v73
                                                }
                                                switch (v3 | fi67) {
                                                case 937:
                                                    var v74 float64 = 1e10
                                                    _ = v74
                                                case 42:
                                                    v75 := GenS3{F0: 126, F1: v68}
                                                    _ = v75
                                                case -2147483648:
                                                    v2 = v2
                                                default:
                                                    switch int(v2) {
                                                    case 32768:
                                                        v68 = v68
                                                    default:
                                                        v76 := GenS8{F0: 4294967295, F1: 2, F2: (v68 | v68)}
                                                        _ = v76
                                                    }
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 42
                                                }
                                                _ = len(v1)
                                                v77 := GenS5{F0: 1, F1: 595}
                                                fmt.Println(fi70, v3, v68)
                                                var v78 float64 = (((((0.0 * (((float64((((3.14159 / (-2.25 + 3.14159)) / ((((3.14159 * -2.25) - 0.0) / 0.1) / 1.0)) / (3.14159 + 0.0))) - float64(((0.1 / ((1e10 - (0.1 / (-2.25 / (0.1 / 1.0)))) / ((0.0 + 1.5) - 0.1))) + -2.25))) * 0.1) / 1.0)) * ((3.14159 * 0.0) / 1.0)) - (0.0 / -2.25)) * 1.5) / -2.25)
                                                v79 := GenS9{F0: 242, F1: (v3 | 4294967295)}
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                            }
                                            v80 := GenS5{F0: -31, F1: 0}
                                            _ = fi70
                                            _ = v80
                                        }
                                    case 2147483647:
                                        var v81 float64 = float64(1.5)
                                        _ = v81
                                    }
                                }
                            case "hello":
                                v3 -= (v3 * 9223372036854775807)
                            case "foo bar":
                                if (false || false) {
                                    for fi82 := 0; fi82 < 15; fi82++ {
                                        v83 := GenS10{F0: 901, F1: 2147483648, F2: 0}
                                        switch "hello" {
                                        case "hello":
                                            v2 -= (v2 << ((42) & 63))
                                        case "\n\t":
                                            v84 := GenS9{F0: 1, F1: fi67}
                                            _ = v84
                                        default:
                                            if false {
                                                if (673 < 1) {
                                                    v3 += 318
                                                    var v85 int8 = v68
                                                    _ = v85
                                                }
                                                v1 = append(v1, 0)
                                                var v86 uint8 = 255
                                                _ = v86
                                            } else {
                                                switch 32767 {
                                                case 2147483647:
                                                    fmt.Println(v68)
                                                case -743:
                                                    _ = len(v1)
                                                case -84:
                                                    var v87 int = v3
                                                    _ = v87
                                                }
                                                v68 = v68
                                                v88 := GenS11{F0: (v68 << ((1) & 63))}
                                                _ = v88
                                            }
                                        }
                                        _ = fi82
                                        _ = v83
                                    }
                                    fmt.Println(fi67)
                                    fmt.Println(v68, fi67)
                                } else {
                                    if true {
                                        for fi89 := 0; fi89 < 16; fi89++ {
                                            var v90 int32 = -1
                                            if ("hello" <= "hello") {
                                                var v91 int64 = (v2 | (v2 / (((v2 % (((v2 << ((255) & 63))) | 1))) | 1)))
                                                _ = v91
                                            }
                                            v1 = append(v1, 2147483648)
                                            _ = fi89
                                            _ = v90
                                        }
                                    }
                                    v1 = append(v1, 2147483648)
                                    var v92 float64 = 1e10
                                    _ = v92
                                }
                            }
                        case "a":
                            fmt.Println(v68, v2, fi67)
                        }
                    }
                default:
                    var v93 bool = (false && (!((float32(1e10) <= 0.1) && (!false))))
                    _ = v93
                }
            case 9223372036854775807:
                fmt.Println(v3)
            }
        }
        _ = fi67
        _ = v68
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
}
