package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 bool; F2 int }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS11 struct { F0 int32; F1 uint32; F2 int32 }
type GenS12 struct { F0 int32; F1 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int8; F1 int64 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int16; F1 string }
type GenS19 struct { F0 int16; F1 int16; F2 int16 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 int }
type GenS23 struct { F0 int16; F1 bool; F2 int8 }
type GenS24 struct { F0 string; F1 int32; F2 uint64 }
type GenS25 struct { F0 uint; F1 uint }
type GenS26 struct { F0 uint64 }
type GenS27 struct { F0 string; F1 uint64; F2 string }
type GenS28 struct { F0 uint16; F1 int64 }
type GenS29 struct { F0 bool }
func genFn0(p0 uint16, p1 uint, p2 uint32) uint16 {
    var v4 bool = false
    fmt.Println(p1, p2)
    switch 678 {
    case 4294967295:
        v5 := []int{32767, 9223372036854775807}
        _ = v5
    default:
        fmt.Println(v4, p1, v4)
    }
    fmt.Println(v4, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (^(p0))
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int = 9223372036854775807
    v3 := []string{""}
    switch v2 {
    case 294:
        if false {
            v6 := GenS1{F0: 42}
            v7 := GenS2{F0: 0, F1: 42, F2: 289}
            if len(v1) > 0 {
                v1[0] = ("\n\t" + "\n\t")
            }
            _ = v6
            _ = v7
        }
    case 2147483647:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi8 := 0; fi8 < 16; fi8++ {
            v2 = -361
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint16 = 135
                if (!false) {
                    switch ("" + "hello") {
                    case "\n\t":
                        fmt.Println(fi8, fi8)
                    case "foo bar":
                        fmt.Println(v10)
                        if true {
                            if true {
                                switch fi8 {
                                case 1:
                                    var v11 bool = true
                                    _ = v11
                                default:
                                    fmt.Println(fi9, v2, v10)
                                }
                                switch "a" {
                                case "foo bar":
                                    var v12 float64 = (((1.5 - float64(float32((0.1 / 1.5)))) * (1e10 * float64(1e10))) + 1.5)
                                    _ = v12
                                case "hello":
                                    if (-519 >= fi8) {
                                        var v13 int = -484
                                        _ = v13
                                    } else {
                                        var v14 float32 = 1e10
                                        v2 -= v2
                                        _ = v14
                                    }
                                case "a":
                                    v2 += (v2 << ((42) & 63))
                                }
                                if ((-2.25 / (-2.25 - ((0.1 + 3.14159) + (-2.25 - (0.1 * (0.0 / 1e10)))))) <= 0.0) {
                                    v15 := GenS2{F0: (v2 % ((fi8) | 1)), F1: 32767, F2: 42}
                                    v10 -= 763
                                    v16 := GenS2{F0: fi9, F1: 0, F2: 661}
                                    _ = v15
                                    _ = v16
                                } else {
                                    _ = len(v3)
                                    var v17 uint64 = 42
                                    if false {
                                        if (!(65535 <= 187)) {
                                            switch "a" {
                                            case "":
                                                var v18 int = 9223372036854775807
                                                _ = v18
                                            case "\n\t":
                                                var v19 float64 = (1.5 + -2.25)
                                                _ = v19
                                            }
                                        } else {
                                            if (!((v2 / ((42) | 1)) > 431)) {
                                                v20 := GenS3{F0: (v17 | (v17 * v17)), F1: 105}
                                                _ = v20
                                            } else {
                                                v21 := GenS1{F0: 2}
                                                fmt.Println(fi9, v2, fi8)
                                                _ = v21
                                            }
                                            fmt.Println(fi9, fi9)
                                        }
                                    } else {
                                        if true {
                                            var v22 bool = (!(("" + "hello") < "foo bar"))
                                            _ = v22
                                        }
                                        v1 = append(v1, "")
                                    }
                                    v23 := GenS1{F0: 65535}
                                    if true {
                                        if true {
                                            if (true && (-681 <= 1)) {
                                                if false {
                                                    fmt.Println(v2, v17)
                                                    v24 := GenS1{F0: 65535}
                                                    fmt.Println(fi9, fi9, v2)
                                                    var v25 uint = 4294967295
                                                    if (!false) {
                                                        v2 = 611
                                                        v10 = 32767
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    var v26 int8 = 1
                                                    var v27 int32 = 1
                                                    switch ("foo bar" + "") {
                                                    case "foo bar":
                                                        v28 := GenS3{F0: v17, F1: -1}
                                                        _ = v28
                                                    case "hello":
                                                        var v29 bool = true
                                                        _ = v29
                                                    default:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            switch (-(fi8)) {
                                                            case 128:
                                                                var v30 uint8 = 42
                                                                _ = v30
                                                            case -513:
                                                                if false {
                                                                    var v31 int64 = 9223372036854775807
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "foo bar"
                                                                    }
                                                                    _ = len(v3)
                                                                    _ = v31
                                                                } else {
                                                                    v32 := GenS2{F0: (-(fi8)), F1: 1, F2: v17}
                                                                    _ = v32
                                                                }
                                                            case -9223372036854775808:
                                                                v33 := GenS2{F0: v2, F1: -445, F2: v17}
                                                                _ = v33
                                                            }
                                                        case "":
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        v17 += v17
                                                        fmt.Println(v25, v10, fi9)
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v10 -= 1
                                                        case "\n\t":
                                                            fmt.Println(v17)
                                                            var v34 float32 = float32((((1.5 * (0.1 - (3.14159 * 0.0))) + (3.14159 * 1.5)) - ((0.1 - 0.1) / 3.14159)))
                                                            var v35 string = ""
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                    }
                                                    if true {
                                                        fmt.Println(v27, fi8)
                                                    } else {
                                                        v1 = append(v1, "")
                                                        v3 = append(v3, "a")
                                                        v17 -= v17
                                                    }
                                                    var v36 uint8 = uint8(v17)
                                                    if ((1.5 - (0.1 + ((-2.25 * 1e10) + 1e10))) == 0.0) {
                                                        v26 += (v26 * v26)
                                                    }
                                                    _ = v24
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                    _ = v36
                                                }
                                            } else {
                                                var v37 float32 = (((1.5 * (3.14159 * -2.25)) * 1e10) / 3.14159)
                                                fmt.Println(v17, v2)
                                                _ = v37
                                            }
                                        }
                                    } else {
                                        var v38 int8 = 34
                                        if (false && true) {
                                            _ = len(v1)
                                            var v39 string = ""
                                            switch 497 {
                                            case -861:
                                                fmt.Println(v39)
                                            }
                                            _ = v39
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v40 uint32 = 500
                                            _ = v40
                                        }
                                        var v41 float64 = 1e10
                                        _ = v38
                                        _ = v41
                                    }
                                    fmt.Println(v2)
                                    _ = v17
                                    _ = v23
                                }
                            }
                            v10 = v10
                        }
                        if false {
                            fmt.Println(fi9)
                        } else {
                            v42 := GenS1{F0: 333}
                            _ = v42
                        }
                    case "a":
                        v10 = (v10 % ((genFn0(42, 255, 675)) | 1))
                    default:
                        var v43 float32 = -2.25
                        _ = v43
                    }
                }
                v44 := GenS4{F0: -107}
                if true {
                    v45 := GenS5{F0: "\n\t", F1: false, F2: fi9}
                    v2 -= fi9
                    _ = v45
                }
                if ((!true) && ((v2 * 96) > (^(v2)))) {
                    switch 127 {
                    case -1:
                        v2 = (v2 & 0)
                    case 582:
                        var v46 int = -671
                        _ = v46
                    case 4294967295:
                        if (4294967295 < 4294967295) {
                            var v47 uint32 = 210
                            var v48 bool = (0.0 <= 3.14159)
                            if false {
                                if (!(1.5 > 0.1)) {
                                    fmt.Println(fi8)
                                }
                            } else {
                                if false {
                                    if ((!(413 >= 42)) && v48) {
                                        var v49 string = ""
                                        var v50 int = (-(fi8))
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v2 += fi9
                                        var v51 float32 = (1.5 - (((0.1 * (float32(float64(3.14159)) + ((float32(float64(0.0)) / (-2.25 / (0.0 - 3.14159))) - 1.5))) + 3.14159) + 0.0))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v51
                                    }
                                    var v52 uint32 = 4294967295
                                    fmt.Println(v10, v48, v47)
                                    if (0 == 1) {
                                        switch "foo bar" {
                                        case "hello":
                                            v53 := GenS6{F0: "", F1: -26}
                                            _ = v53
                                        default:
                                            fmt.Println(fi8)
                                        }
                                        var v54 uint = 982
                                        var v55 uint64 = 1
                                        _ = v54
                                        _ = v55
                                    } else {
                                        var v56 float64 = 1e10
                                        v48 = v48
                                        _ = v56
                                    }
                                    _ = v52
                                } else {
                                    v47 -= v47
                                    if true {
                                        var v57 float32 = float32(((float64((1.5 / 1e10)) - 3.14159) / 1.5))
                                        switch "a" {
                                        case "":
                                            v48 = (42 == 42)
                                        default:
                                            v57 = 1e10
                                            var v58 uint32 = (v47 / (((v47 / (((v47 - 2147483648)) | 1))) | 1))
                                            if (220 < 42) {
                                                v10 = 533
                                                switch "\n\t" {
                                                case "":
                                                    v59 := GenS1{F0: v58}
                                                    v60 := GenS7{F0: 201, F1: 0, F2: 0}
                                                    _ = v59
                                                    _ = v60
                                                case "a":
                                                    fmt.Println(fi9, v47)
                                                }
                                                switch (fi9 | (v2 & fi9)) {
                                                case -1:
                                                    v61 := GenS3{F0: 2, F1: 100}
                                                    _ = v61
                                                case 32767:
                                                    v10 = (^(v10))
                                                default:
                                                    var v62 string = "foo bar"
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v62
                                                }
                                                if true {
                                                    if true {
                                                        var v63 float32 = float32(float64((1e10 - (1.5 - v57))))
                                                        if v48 {
                                                            _ = len(v1)
                                                        } else {
                                                            switch 11 {
                                                            case 4294967295:
                                                                if v48 {
                                                                    _ = len(v3)
                                                                    v64 := GenS7{F0: uint8(v10), F1: 249, F2: 2}
                                                                    _ = v64
                                                                }
                                                            case -1000:
                                                                v1 = append(v1, "")
                                                            default:
                                                                if (42 >= 18446744073709551615) {
                                                                    var v65 int = (fi9 << ((18446744073709551615) & 63))
                                                                    fmt.Println(v10, v57)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "foo bar"
                                                                    }
                                                                    _ = v65
                                                                } else {
                                                                    v48 = (398 > 4294967295)
                                                                    var v66 int32 = 42
                                                                    fmt.Println(v2, v58)
                                                                    _ = v66
                                                                }
                                                            }
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v67 uint8 = 2
                                                                _ = v67
                                                            case "hello":
                                                                var v68 uint16 = v10
                                                                _ = v68
                                                            }
                                                        }
                                                        v58 -= (^(v47))
                                                        _ = v63
                                                    }
                                                }
                                            } else {
                                                var v69 int8 = -113
                                                switch "" {
                                                case "a":
                                                    switch "" {
                                                    case "":
                                                        var v70 uint = 4294967295
                                                        _ = v70
                                                    case "a":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            fmt.Println(v10, v57)
                                                        } else {
                                                            fmt.Println(v57, v58, v10)
                                                            var v71 uint16 = v10
                                                            if (!v48) {
                                                                switch v2 {
                                                                case 0:
                                                                    if (1 >= 2) {
                                                                        fmt.Println(v69)
                                                                        var v72 float32 = (float32(0.0) / v57)
                                                                        _ = len(v1)
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                            _ = v71
                                                        }
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                case "\n\t":
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(v69)
                                                    case "\n\t":
                                                        v1 = append(v1, "\n\t")
                                                    case "":
                                                        v73 := GenS4{F0: 36}
                                                        _ = v73
                                                    }
                                                }
                                                v57 -= (((3.14159 * 1e10) / v57) / v57)
                                                _ = v69
                                            }
                                            _ = v58
                                        }
                                        _ = v57
                                    } else {
                                        var v74 bool = true
                                        switch "hello" {
                                        case "foo bar":
                                            v48 = (!((364 > 4294967295) && (!v74)))
                                        case "":
                                            switch "" {
                                            case "hello":
                                                v2 += fi9
                                            case "a":
                                                switch "" {
                                                case "foo bar":
                                                    v1 = append(v1, "")
                                                case "hello":
                                                    v10 -= genFn0(v10, 790, v47)
                                                case "\n\t":
                                                    if false {
                                                        v2 += fi8
                                                        if v74 {
                                                            if false {
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("a" + (("" + ("\n\t" + "hello")) + "\n\t"))
                                                                }
                                                            } else {
                                                                switch ("a" + "\n\t") {
                                                                case "a":
                                                                    v75 := GenS2{F0: fi9, F1: -1, F2: 2}
                                                                    _ = v75
                                                                default:
                                                                    if (-810 <= 42) {
                                                                        if (1e10 == 0.1) {
                                                                            if true {
                                                                                var v76 uint = 0
                                                                                _ = v76
                                                                            } else {
                                                                                fmt.Println(v48, fi8, fi9)
                                                                                v74 = v48
                                                                                var v77 int8 = 127
                                                                                _ = v77
                                                                            }
                                                                            v2 = v2
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                switch ("\n\t" + ("" + ("foo bar" + ""))) {
                                                                                case "\n\t":
                                                                                    switch (v2 >> ((65535) & 63)) {
                                                                                    case -316:
                                                                                        fmt.Println(fi8)
                                                                                    default:
                                                                                        if v48 {
                                                                                            var v78 bool = (224 != 0)
                                                                                            if (!(127 >= 128)) {
                                                                                                v47 += 4294967295
                                                                                                v74 = (false || (!(0 < 508)))
                                                                                            }
                                                                                            v79 := GenS8{F0: (fi8 << ((233) & 63)), F1: "a"}
                                                                                            var v80 bool = (!(1 > 2))
                                                                                            if false {
                                                                                                fmt.Println(v47, v2, v47)
                                                                                            }
                                                                                            fmt.Println(v10)
                                                                                            fmt.Println(v74, v48)
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    v81 := GenS9{F0: -892, F1: 1}
                                                                                    _ = v81
                                                                                case "a":
                                                                                    if (!v48) {
                                                                                        v82 := GenS10{F0: 42, F1: 1, F2: (v10 / (((v10 & 255)) | 1))}
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            var v83 uint64 = 18446744073709551615
                                                                                            _ = v83
                                                                                        case "foo bar":
                                                                                            var v84 uint64 = 402
                                                                                            _ = v84
                                                                                        case "a":
                                                                                            v48 = v48
                                                                                        default:
                                                                                            var v85 uint = 42
                                                                                            _ = v85
                                                                                        }
                                                                                        switch v2 {
                                                                                        case 127:
                                                                                            var v86 float64 = (1.5 * -2.25)
                                                                                            _ = v86
                                                                                        case 0:
                                                                                            var v87 string = "\n\t"
                                                                                            _ = v87
                                                                                        case -2147483648:
                                                                                            v88 := GenS2{F0: 127, F1: -380, F2: 1}
                                                                                            _ = v88
                                                                                        default:
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = "\n\t"
                                                                                            }
                                                                                        }
                                                                                        _ = v82
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                v48 = v74
                                                                            }
                                                                            var v89 int = v2
                                                                            _ = v89
                                                                        } else {
                                                                            var v90 string = "hello"
                                                                            switch (fi9 % ((int(v10)) | 1)) {
                                                                            case 4294967295:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            case 2147483647:
                                                                                switch fi8 {
                                                                                case 950:
                                                                                    v91 := GenS7{F0: 42, F1: -1, F2: 116}
                                                                                    _ = v91
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    var v92 uint32 = (v47 >> ((18446744073709551615) & 63))
                                                                                    _ = v92
                                                                                case "a":
                                                                                    if true {
                                                                                        v93 := GenS5{F0: ("foo bar" + "\n\t"), F1: v48, F2: v2}
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v47 = v47
                                                                                        v94 := GenS11{F0: -2147483648, F1: v47, F2: 2147483647}
                                                                                        switch (v90 + v90) {
                                                                                        case "":
                                                                                            var v95 int = 205
                                                                                            _ = v95
                                                                                        case "a":
                                                                                            fmt.Println(v10)
                                                                                        }
                                                                                        v90 = v90
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                }
                                                                                v96 := GenS4{F0: 1}
                                                                                _ = v96
                                                                            }
                                                                            _ = v90
                                                                        }
                                                                        v48 = v48
                                                                        v2 = v2
                                                                    }
                                                                }
                                                            }
                                                            var v97 uint = 2
                                                            if (((fi8 + (-(fi9))) >= int(v47)) || (9223372036854775807 == -1)) {
                                                                v74 = v74
                                                                if v74 {
                                                                    if (!v48) {
                                                                        switch (v2 + (fi8 << ((v97) & 63))) {
                                                                        case 1:
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                fmt.Println(v74, v74, v10)
                                                                            default:
                                                                                var v98 int16 = 1
                                                                                _ = v98
                                                                            }
                                                                        case 9223372036854775807:
                                                                            if false {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                        case -2147483648:
                                                                            var v99 int8 = 1
                                                                            _ = v99
                                                                        }
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    } else {
                                                                        v10 = (v10 | (v10 & 318))
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("\n\t" + ("" + ("foo bar" + "hello")))
                                                                        }
                                                                        var v100 string = ("a" + "a")
                                                                        fmt.Println(v10, v2)
                                                                        var v101 int = v2
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                    v102 := GenS1{F0: (v47 % (((v47 ^ 42)) | 1))}
                                                                    v103 := GenS5{F0: "", F1: ((73 == -49) && (3.14159 > 0.0)), F2: fi8}
                                                                    _ = v102
                                                                    _ = v103
                                                                }
                                                            }
                                                            _ = v97
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = ("a" + "")
                                                            }
                                                            switch (^(fi9)) {
                                                            case 9223372036854775807:
                                                                v74 = v74
                                                            default:
                                                                if false {
                                                                    fmt.Println(fi8, v48, fi8)
                                                                    var v104 int64 = 9223372036854775807
                                                                    v105 := GenS12{F0: 42, F1: (("" + "foo bar") + "\n\t")}
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    switch v2 {
                                                                    case 32767:
                                                                        var v106 uint16 = (v10 / (((v10 - v10)) | 1))
                                                                        _ = v106
                                                                    case 483:
                                                                        var v107 int32 = 819
                                                                        _ = v107
                                                                    default:
                                                                        var v108 uint64 = 18446744073709551615
                                                                        _ = v108
                                                                    }
                                                                    if v74 {
                                                                        switch (fi8 ^ fi8) {
                                                                        case 2147483647:
                                                                            v47 += (v47 * 2147483648)
                                                                        default:
                                                                            if (1000 >= 42) {
                                                                                switch "hello" {
                                                                                case "hello":
                                                                                    var v109 string = "hello"
                                                                                    _ = v109
                                                                                }
                                                                                switch "" {
                                                                                case "a":
                                                                                    v110 := GenS13{F0: (42 == 32767)}
                                                                                    _ = v110
                                                                                case "foo bar":
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ""
                                                                                    }
                                                                                    if (v48 && (v48 || (("" <= "") && false))) {
                                                                                        v104 = 42
                                                                                    }
                                                                                    var v111 string = "a"
                                                                                    _ = v111
                                                                                }
                                                                                if true {
                                                                                    var v112 uint16 = v10
                                                                                    var v113 uint8 = 2
                                                                                    v47 -= (v47 / ((uint32(v10)) | 1))
                                                                                    _ = v112
                                                                                    _ = v113
                                                                                }
                                                                                if (!(!(!v74))) {
                                                                                    if v74 {
                                                                                        var v114 float32 = (0.1 / -2.25)
                                                                                        _ = len(v3)
                                                                                        fmt.Println(v10, v48)
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v114
                                                                                    } else {
                                                                                        var v115 uint64 = 1
                                                                                        v48 = v74
                                                                                        if v74 {
                                                                                            v47 += v47
                                                                                        } else {
                                                                                            fmt.Println(fi8)
                                                                                            var v116 bool = v74
                                                                                            v117 := GenS14{F0: -66, F1: v104}
                                                                                            switch fi8 {
                                                                                            case -2147483648:
                                                                                                v10 = (v10 % ((v10) | 1))
                                                                                            case -1:
                                                                                                if ((v10 < uint16(v115)) || v74) {
                                                                                                    v10 += (v10 % ((genFn0((^(v10)), 7, (v47 << ((4294967295) & 63)))) | 1))
                                                                                                    v3 = append(v3, "a")
                                                                                                }
                                                                                            default:
                                                                                                v118 := GenS15{F0: 100}
                                                                                                _ = v118
                                                                                            }
                                                                                            fmt.Println(v74, fi8, v2)
                                                                                            var v119 uint32 = 2147483648
                                                                                            if ((float64(float32(-2.25)) > 1.5) && ((v115 - v115) >= (^(v115)))) {
                                                                                                if false {
                                                                                                    switch (-(fi8)) {
                                                                                                    case 9223372036854775807:
                                                                                                        var v120 uint16 = v10
                                                                                                        _ = v120
                                                                                                    case -2147483648:
                                                                                                        v121 := GenS16{F0: ("" <= ("\n\t" + "\n\t"))}
                                                                                                        _ = v121
                                                                                                    default:
                                                                                                        v104 = (v104 >> ((836) & 63))
                                                                                                    }
                                                                                                }
                                                                                                if (v2 != (fi8 * 32768)) {
                                                                                                    switch (("foo bar" + ("\n\t" + "hello")) + "hello") {
                                                                                                    case "\n\t":
                                                                                                        var v122 float64 = ((0.1 / float64(0.1)) - 1e10)
                                                                                                        _ = v122
                                                                                                    case "hello":
                                                                                                        var v123 float64 = (1e10 + (3.14159 / 3.14159))
                                                                                                        _ = v123
                                                                                                    case "":
                                                                                                        var v124 uint = 18446744073709551615
                                                                                                        _ = v124
                                                                                                    default:
                                                                                                        v125 := GenS5{F0: "foo bar", F1: (42 == 87), F2: fi8}
                                                                                                        _ = v125
                                                                                                    }
                                                                                                    v126 := GenS17{F0: v48}
                                                                                                    var v127 uint64 = 1
                                                                                                    _ = v126
                                                                                                    _ = v127
                                                                                                } else {
                                                                                                    fmt.Println(v47)
                                                                                                    v128 := GenS17{F0: v116}
                                                                                                    var v129 uint32 = (v47 & 732)
                                                                                                    _ = v128
                                                                                                    _ = v129
                                                                                                }
                                                                                                var v130 uint64 = (v115 << ((170) & 63))
                                                                                                fmt.Println(v10, fi8)
                                                                                                _ = v130
                                                                                            }
                                                                                            _ = v116
                                                                                            _ = v117
                                                                                            _ = v119
                                                                                        }
                                                                                        _ = v115
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v131 int8 = 42
                                                                            _ = v131
                                                                        case "foo bar":
                                                                            switch -780 {
                                                                            case -9223372036854775808:
                                                                                fmt.Println(v74, v10)
                                                                            case 2147483647:
                                                                                if (!v74) {
                                                                                    if v74 {
                                                                                        fmt.Println(fi9)
                                                                                    } else {
                                                                                        v3 = append(v3, "a")
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ("hello" + "")
                                                                                        }
                                                                                        fmt.Println(v74, v104, v74)
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        v1 = append(v1, "foo bar")
                                                                                    }
                                                                                } else {
                                                                                    switch ("" + "") {
                                                                                    case "a":
                                                                                        v2 -= -1
                                                                                    case "":
                                                                                        var v132 uint8 = 127
                                                                                        _ = v132
                                                                                    case "\n\t":
                                                                                        v47 -= 612
                                                                                    }
                                                                                    fmt.Println(v47, v47, v74)
                                                                                    switch (fi8 | (fi9 - fi8)) {
                                                                                    case -1:
                                                                                        fmt.Println(fi9)
                                                                                    case 128:
                                                                                        if (((3.14159 / (-2.25 / 1.5)) * 0.1) != 0.1) {
                                                                                            fmt.Println(v2)
                                                                                        } else {
                                                                                            var v133 int64 = (-(v104))
                                                                                            var v134 int32 = 0
                                                                                            _ = v133
                                                                                            _ = v134
                                                                                        }
                                                                                    default:
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "\n\t"
                                                                                        }
                                                                                        fmt.Println(v47, v47, fi9)
                                                                                        switch ("" + "\n\t") {
                                                                                        case "a":
                                                                                            var v135 float32 = -2.25
                                                                                            _ = v135
                                                                                        case "foo bar":
                                                                                            v1 = append(v1, "")
                                                                                        case "":
                                                                                            switch v2 {
                                                                                            case -1:
                                                                                                v48 = (!(v2 <= v2))
                                                                                            case -2147483648:
                                                                                                v104 += (v104 ^ (^(v104)))
                                                                                            case -76:
                                                                                                v74 = v48
                                                                                            default:
                                                                                                var v136 uint16 = v10
                                                                                                _ = v136
                                                                                            }
                                                                                        default:
                                                                                            if false {
                                                                                                if (!((v10 >> ((4294967295) & 63)) == (v10 >> ((uint(v47)) & 63)))) {
                                                                                                    if true {
                                                                                                        v2 += (v2 % ((fi9) | 1))
                                                                                                        v47 = v47
                                                                                                    } else {
                                                                                                        if (1000 <= -22) {
                                                                                                            var v137 int16 = 32767
                                                                                                            var v138 float64 = ((float64(3.14159) + (0.1 + (1e10 - -2.25))) + 3.14159)
                                                                                                            if ((v2 >> ((65535) & 63)) != 4294967295) {
                                                                                                                var v139 bool = (int8(v104) == -128)
                                                                                                                _ = v139
                                                                                                            } else {
                                                                                                                fmt.Println(fi9)
                                                                                                                fmt.Println(fi9, v10, v2)
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                            _ = v137
                                                                                                            _ = v138
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v140 := GenS14{F0: -61, F1: (v104 << ((780) & 63))}
                                                                                                _ = v140
                                                                                            } else {
                                                                                                switch ("\n\t" + "\n\t") {
                                                                                                case "\n\t":
                                                                                                    if v48 {
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = ("" + "")
                                                                                                        }
                                                                                                        v2 = (^(fi9))
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v47 -= v47
                                                                                                default:
                                                                                                    v141 := GenS13{F0: false}
                                                                                                    _ = v141
                                                                                                }
                                                                                                v10 += (v10 * v10)
                                                                                                v2 -= (-(fi9))
                                                                                                var v142 uint64 = 1
                                                                                                v10 = (v10 / ((v10) | 1))
                                                                                                _ = v142
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v74 = v74
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = ""
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v1 = append(v1, ("foo bar" + "foo bar"))
                                                                                    default:
                                                                                        _ = len(v1)
                                                                                    }
                                                                                }
                                                                            }
                                                                            if v74 {
                                                                                v3 = append(v3, "hello")
                                                                                var v143 int = (fi8 % (((fi8 & fi8)) | 1))
                                                                                _ = v143
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        default:
                                                                            if v48 {
                                                                                switch fi8 {
                                                                                case 4294967295:
                                                                                    v2 -= fi9
                                                                                case 42:
                                                                                    fmt.Println(v104)
                                                                                default:
                                                                                    v144 := GenS18{F0: 1000, F1: ("hello" + ("\n\t" + "\n\t"))}
                                                                                    _ = v144
                                                                                }
                                                                                v145 := GenS8{F0: fi8, F1: "a"}
                                                                                v10 += genFn0((v10 << ((684) & 63)), 401, 2147483648)
                                                                                var v146 uint32 = (^(v47))
                                                                                v147 := GenS19{F0: 713, F1: 32767, F2: -1}
                                                                                v104 = (v104 ^ v104)
                                                                                _ = v145
                                                                                _ = v146
                                                                                _ = v147
                                                                            } else {
                                                                                var v148 int8 = -1
                                                                                v104 = (v104 * (v104 / ((-1) | 1)))
                                                                                _ = v148
                                                                            }
                                                                        }
                                                                        v149 := GenS16{F0: true}
                                                                        var v150 uint8 = 127
                                                                        _ = v149
                                                                        _ = v150
                                                                    }
                                                                    if v48 {
                                                                        if v74 {
                                                                            var v151 uint32 = v47
                                                                            _ = len(v1)
                                                                            var v152 uint16 = (v10 | (v10 >> ((106) & 63)))
                                                                            _ = v151
                                                                            _ = v152
                                                                        } else {
                                                                            if false {
                                                                                var v153 int32 = 42
                                                                                if (!(v2 > fi9)) {
                                                                                    var v154 int32 = (v153 ^ 0)
                                                                                    v47 = (v47 << ((uint(v10)) & 63))
                                                                                    var v155 float32 = ((-2.25 / float32(0.1)) * float32(0.0))
                                                                                    _ = v154
                                                                                    _ = v155
                                                                                } else {
                                                                                    v2 = (-(v2))
                                                                                    v48 = false
                                                                                    var v156 string = "hello"
                                                                                    _ = v156
                                                                                }
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                _ = v153
                                                                            } else {
                                                                                v157 := GenS20{F0: v47}
                                                                                if v74 {
                                                                                    var v158 float64 = 0.1
                                                                                    _ = v158
                                                                                } else {
                                                                                    switch v2 {
                                                                                    case -249:
                                                                                        v2 = (v2 >> ((4294967295) & 63))
                                                                                    case 127:
                                                                                        var v159 int8 = 100
                                                                                        _ = v159
                                                                                    default:
                                                                                        _ = len(v3)
                                                                                    }
                                                                                }
                                                                                v160 := GenS21{F0: v47}
                                                                                fmt.Println(v47)
                                                                                _ = v157
                                                                                _ = v160
                                                                            }
                                                                            fmt.Println(v10, fi9, v2)
                                                                        }
                                                                    }
                                                                    var v161 uint8 = uint8(v10)
                                                                    _ = v104
                                                                    _ = v105
                                                                    _ = v161
                                                                }
                                                                var v162 float64 = 3.14159
                                                                _ = v162
                                                            }
                                                        }
                                                        v10 = 65535
                                                        if v74 {
                                                            if v74 {
                                                                var v163 int = (fi9 * 670)
                                                                _ = v163
                                                            } else {
                                                                var v164 float32 = float32(float64(1.5))
                                                                v2 += fi8
                                                                v165 := GenS8{F0: (v2 % (((fi9 % (((-(v2))) | 1))) | 1)), F1: "\n\t"}
                                                                _ = v164
                                                                _ = v165
                                                            }
                                                            v166 := GenS22{F0: v2}
                                                            var v167 uint16 = 840
                                                            _ = v166
                                                            _ = v167
                                                        }
                                                        var v168 int32 = -2147483648
                                                        _ = v168
                                                    } else {
                                                        v47 -= (^(v47))
                                                        v169 := GenS1{F0: (v47 - v47)}
                                                        v10 = v10
                                                        _ = v169
                                                    }
                                                }
                                            case "\n\t":
                                                v48 = (42 < 1)
                                            }
                                        case "hello":
                                            v47 = 461
                                        default:
                                            var v170 int64 = 575
                                            _ = v170
                                        }
                                        if (1.5 != 3.14159) {
                                            if v48 {
                                                fmt.Println(v2, v2, fi9)
                                                v171 := GenS23{F0: 42, F1: v74, F2: -128}
                                                _ = v171
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    if (-128 >= 127) {
                                                        if v48 {
                                                            v47 -= (v47 - 2147483648)
                                                            var v172 int8 = 0
                                                            if (3.14159 <= 0.1) {
                                                                var v173 uint8 = 42
                                                                var v174 uint32 = 2
                                                                _ = v173
                                                                _ = v174
                                                            } else {
                                                                var v175 uint16 = v10
                                                                _ = v175
                                                            }
                                                            _ = v172
                                                        } else {
                                                            switch (("hello" + "a") + ((("hello" + "\n\t") + ("foo bar" + (("foo bar" + "") + (("foo bar" + "foo bar") + "\n\t")))) + ("a" + (("a" + "foo bar") + "foo bar")))) {
                                                            case "":
                                                                fmt.Println(v2, v47, v47)
                                                            case "hello":
                                                                switch "a" {
                                                                case "hello":
                                                                    switch fi9 {
                                                                    case -347:
                                                                        fmt.Println(v47, fi8)
                                                                    default:
                                                                        var v176 uint64 = 523
                                                                        _ = v176
                                                                    }
                                                                case "\n\t":
                                                                    v2 = fi9
                                                                default:
                                                                    var v177 uint64 = 43
                                                                    _ = v177
                                                                }
                                                            default:
                                                                var v178 int = v2
                                                                _ = v178
                                                            }
                                                            v47 += (^(v47))
                                                        }
                                                    }
                                                case "":
                                                    v48 = (0 < 0)
                                                case "\n\t":
                                                    v179 := GenS16{F0: v74}
                                                    _ = v179
                                                default:
                                                    switch (("a" + (("foo bar" + "foo bar") + (("foo bar" + "") + ""))) + "foo bar") {
                                                    case "":
                                                        if (-1 >= 0) {
                                                            var v180 int64 = 680
                                                            fmt.Println(v10, v180)
                                                            fmt.Println(v48)
                                                            _ = v180
                                                        }
                                                    case "a":
                                                        fmt.Println(fi9)
                                                    case "foo bar":
                                                        if true {
                                                            var v181 int64 = 0
                                                            v47 = v47
                                                            _ = v181
                                                        }
                                                    }
                                                }
                                                var v182 int = -358
                                                _ = v182
                                            }
                                            v183 := GenS24{F0: "a", F1: 1, F2: 0}
                                            var v184 int32 = -309
                                            v185 := GenS25{F0: 65535, F1: 1}
                                            switch ((("hello" + ("a" + (("\n\t" + "hello") + "a"))) + ("" + "")) + "\n\t") {
                                            case "foo bar":
                                                var v186 uint16 = v10
                                                _ = v186
                                            case "hello":
                                                var v187 int64 = int64(fi8)
                                                _ = v187
                                            }
                                            _ = v183
                                            _ = v184
                                            _ = v185
                                        } else {
                                            v74 = v48
                                            var v188 int8 = -1
                                            v10 = (^(v10))
                                            _ = v188
                                        }
                                        var v189 string = "foo bar"
                                        _ = v189
                                        _ = v74
                                    }
                                    if (((float32(3.14159) / float32(float64(float32((3.14159 * (((1.5 * 3.14159) * (-2.25 / (1e10 / -2.25))) * 1e10)))))) < 1.5) && v48) {
                                        if true {
                                            _ = len(v1)
                                        }
                                        if true {
                                            switch "\n\t" {
                                            case "hello":
                                                v2 += fi9
                                            case "\n\t":
                                                var v190 uint32 = (v47 - (^(v47)))
                                                _ = v190
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            if v48 {
                                                var v191 int8 = 0
                                                var v192 uint64 = 1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v191
                                                _ = v192
                                            }
                                            var v193 bool = ((-2.25 * 1.5) == 3.14159)
                                            _ = v193
                                        } else {
                                            v10 -= v10
                                            v2 = 1
                                        }
                                    } else {
                                        var v194 int32 = 969
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v195 := GenS21{F0: (v47 << ((18446744073709551615) & 63))}
                                        _ = v194
                                        _ = v195
                                    }
                                }
                                var v196 int32 = 1
                                _ = v196
                            }
                            v47 = 608
                            v10 -= 879
                            _ = v47
                            _ = v48
                        }
                    default:
                        var v197 uint32 = 2147483647
                        if false {
                            if false {
                                var v198 int8 = -1
                                _ = v198
                            } else {
                                v10 = 32767
                                v199 := GenS9{F0: 1000, F1: -128}
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                _ = v199
                            }
                            fmt.Println(v197)
                        } else {
                            var v200 uint = 18446744073709551615
                            var v201 int = (fi9 - -9223372036854775808)
                            _ = v200
                            _ = v201
                        }
                        if (true && false) {
                            switch "" {
                            case "hello":
                                v10 = (v10 >> ((65535) & 63))
                            case "":
                                fmt.Println(fi9, v10, v197)
                            }
                            switch fi9 {
                            case 491:
                                if true {
                                    v202 := GenS6{F0: ("foo bar" + "foo bar"), F1: -1}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v202
                                } else {
                                    v203 := GenS24{F0: "", F1: -2147483648, F2: 1}
                                    _ = v203
                                }
                            }
                        }
                        _ = v197
                    }
                    var v204 int = -1
                    if (true && true) {
                        v3 = append(v3, "a")
                        v205 := GenS26{F0: 2}
                        var v206 uint8 = 128
                        _ = v205
                        _ = v206
                    }
                    _ = v204
                } else {
                    if ((!(1.5 < 0.1)) || (2147483647 <= (fi9 - int(v10)))) {
                        v3 = append(v3, ("\n\t" + ""))
                        v3 = append(v3, "")
                    }
                    switch fi9 {
                    case 670:
                        fmt.Println(fi9, v10)
                    case -1:
                        v10 = genFn0(32767, 1, 2147483648)
                    }
                }
                var v207 uint8 = 223
                v3 = append(v3, ("" + (("hello" + ("foo bar" + "")) + ("hello" + "foo bar"))))
                v2 += fi9
                _ = len(v1)
                _ = fi9
                _ = v10
                _ = v207
                _ = v44
            }
            _ = fi8
        }
    }
    v208 := GenS27{F0: "foo bar", F1: 2, F2: "hello"}
    v2 = (v2 & (v2 & v2))
    _ = len(v1)
    v2 -= v2
    switch (("\n\t" + (("\n\t" + "foo bar") + "hello")) + "a") {
    case "a":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case "\n\t":
        v2 -= (v2 - v2)
    case "":
        if true {
            for fi209 := 0; fi209 < 20; fi209++ {
                if false {
                    switch "hello" {
                    case "hello":
                        switch ("foo bar" + (((("" + ((("hello" + "a") + ("a" + ((("hello" + ("foo bar" + "a")) + ("" + ((("foo bar" + "\n\t") + "a") + "foo bar"))) + ""))) + "a")) + (("" + ("hello" + "\n\t")) + ("hello" + "a"))) + "") + ("foo bar" + ""))) {
                        case "":
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                        case "\n\t":
                            var v210 float64 = 3.14159
                            _ = v210
                        case "a":
                            fmt.Println(v2, fi209)
                        }
                    case "a":
                        if true {
                            var v211 float32 = ((-2.25 * -2.25) + -2.25)
                            _ = v211
                        } else {
                            v212 := GenS28{F0: 2, F1: 0}
                            _ = len(v1)
                            _ = v212
                        }
                    default:
                        if ((!(4294967295 <= 42)) && (!true)) {
                            v2 = (v2 - v2)
                        }
                    }
                    fmt.Println(fi209, v2)
                    v3 = append(v3, "\n\t")
                }
                for fi213 := 0; fi213 < 3; fi213++ {
                    if false {
                        v214 := GenS29{F0: false}
                        _ = v214
                    } else {
                        switch -287 {
                        case 127:
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            var v215 bool = (false && true)
                            _ = v215
                        case 9223372036854775807:
                            var v216 int = fi213
                            if true {
                                v216 -= -1
                                if (true || (2147483647 <= -2147483648)) {
                                    switch "a" {
                                    case "a":
                                        var v217 int16 = 42
                                        _ = v217
                                    case "\n\t":
                                        v218 := GenS3{F0: 4294967295, F1: 0}
                                        _ = v218
                                    }
                                    fmt.Println(v2, fi213)
                                }
                                var v219 float32 = (1.5 / -2.25)
                                _ = v219
                            } else {
                                var v220 int = (fi209 + fi209)
                                _ = v220
                            }
                            var v221 uint16 = 1
                            _ = v216
                            _ = v221
                        default:
                            fmt.Println(v2)
                        }
                        v222 := GenS22{F0: v2}
                        if (65535 >= 2147483647) {
                            var v223 uint16 = 96
                            _ = v223
                        }
                        _ = v222
                    }
                    if (-1 == 1) {
                        var v224 float32 = 1e10
                        _ = v224
                    } else {
                        if ((fi213 >> ((4294967295) & 63)) >= fi209) {
                            fmt.Println(v2, fi209, fi213)
                        } else {
                            var v225 uint8 = 2
                            _ = v225
                        }
                        fmt.Println(v2, v2)
                        fmt.Println(v2, fi209, v2)
                    }
                    if false {
                        var v226 int = -253
                        _ = v226
                    } else {
                        fmt.Println(fi213)
                    }
                    _ = fi213
                }
                _ = fi209
            }
            switch (v2 - v2) {
            case -9223372036854775808:
                var v227 uint64 = 1
                _ = v227
            }
        }
    }
    var v228 int = (-(v2))
    var v229 float64 = float64(3.14159)
    var v230 bool = ((0 > 289) || false)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v228:", v228)
    fmt.Println("v229:", v229)
    fmt.Println("v230:", v230)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v208
    _ = v228
    _ = v229
    _ = v230
    _ = v3
}
