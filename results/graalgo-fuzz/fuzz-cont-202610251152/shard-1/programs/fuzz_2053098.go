package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 32767
    v3 := []string{"a", "a", ("foo bar" + (("\n\t" + "hello") + "hello"))}
    v4 := make(map[int]string)
    v5 := []string{}
    fmt.Println(v2)
    switch -494 {
    case 2147483647:
        switch ("\n\t" + ("" + ("foo bar" + (("a" + "hello") + "")))) {
        case "":
            v2 -= v2
        case "\n\t":
            var v6 uint32 = 42
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: false}
            _ = v7
        default:
            var v8 int = v2
            _ = v8
        }
    case -104:
        v2 += v2
    case 127:
        var v9 uint16 = 65535
        _ = v9
    }
    v10 := GenS1{F0: false}
    if (true && true) {
        switch "" {
        case "hello":
            var v11 bool = false
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint64 = 2
            var v14 int8 = 11
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := v4[v2]
        _ = v15
    } else {
        fmt.Println(v2)
        if (-2147483648 >= -806) {
            var v16 uint16 = 255
            v17 := GenS2{F0: "\n\t", F1: ((!false) || ((!false) || false)), F2: 4294967295}
            _ = v16
            _ = v17
        } else {
            var v18 int = (v2 % ((-90) | 1))
            _ = v18
        }
    }
    if true {
        for fi19 := 0; fi19 < 14; fi19++ {
            if true {
                if false {
                    v2 = (fi19 * v2)
                    for fi20 := 0; fi20 < 2; fi20++ {
                        switch 32768 {
                        case -708:
                            var v21 float64 = 3.14159
                            _ = v21
                        case -61:
                            if (false || false) {
                                fmt.Println(fi19, v2)
                                _ = len(v3)
                            }
                        default:
                            var v22 float32 = 1.5
                            _ = v22
                        }
                        _ = fi20
                    }
                }
                if false {
                    for fi23 := 0; fi23 < 4; fi23++ {
                        switch (("" + "foo bar") + "hello") {
                        case "hello":
                            var v24 float64 = 0.1
                            _ = v24
                        }
                        if true {
                            v25 := GenS2{F0: "\n\t", F1: true, F2: 595}
                            fmt.Println(fi19, fi23, fi23)
                            _ = v25
                        } else {
                            switch fi19 {
                            case 42:
                                v26 := GenS1{F0: ((2147483648 == 2147483647) && false)}
                                _ = v26
                            case 128:
                                v2 = (^(fi19))
                            case 32768:
                                if false {
                                    var v27 bool = ("a" >= "a")
                                    _ = v27
                                }
                            default:
                                if ((false || (!true)) && false) {
                                    var v28 int8 = -128
                                    fmt.Println(v2)
                                    switch "a" {
                                    case "a":
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v28)
                                        case "a":
                                            var v29 uint = 569
                                            _ = v29
                                        }
                                    case "foo bar":
                                        var v30 uint16 = 1
                                        v28 = v28
                                        switch (fi19 >> ((65535) & 63)) {
                                        case 0:
                                            switch -272 {
                                            case -2147483648:
                                                v2 -= 264
                                                v31 := GenS3{F0: 459, F1: 2147483647, F2: 0}
                                                _ = v31
                                            case 1:
                                                v32 := GenS4{F0: v28, F1: -1}
                                                _ = v32
                                            }
                                        case -589:
                                            if len(v3) > 0 {
                                                v3[0] = (("hello" + "") + "hello")
                                            }
                                        case 42:
                                            switch v2 {
                                            case -158:
                                                _ = len(v1)
                                            case 0:
                                                var v33 float64 = 0.0
                                                _ = v33
                                            case 32767:
                                                v30 += (^(v30))
                                            }
                                        default:
                                            switch v2 {
                                            case 32767:
                                                v2 = fi23
                                            case 463:
                                                v2 = v2
                                            case 4294967295:
                                                if true {
                                                    if false {
                                                        var v34 int = (fi19 << ((65535) & 63))
                                                        _ = v34
                                                    } else {
                                                        var v35 uint16 = v30
                                                        switch fi19 {
                                                        case 556:
                                                            fmt.Println(v30, v35, fi19)
                                                        case 127:
                                                            v36 := GenS5{F0: ((4294967295 > 2) || (0 > 271))}
                                                            _ = v36
                                                        default:
                                                            fmt.Println(fi23)
                                                        }
                                                        v37 := GenS6{F0: 778, F1: v30, F2: 127}
                                                        _ = v35
                                                        _ = v37
                                                    }
                                                }
                                            default:
                                                if ((339 > 897) && true) {
                                                    var v38 uint64 = 107
                                                    fmt.Println(fi23, v28, v38)
                                                    v4[(fi19 ^ (fi23 | -9223372036854775808))] = "foo bar"
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v39 := GenS2{F0: "foo bar", F1: false, F2: 1}
                                                        _ = v39
                                                    case "a":
                                                        var v40 uint8 = 255
                                                        _ = v40
                                                    default:
                                                        switch ("\n\t" + "foo bar") {
                                                        case "\n\t":
                                                            if (true && false) {
                                                                var v41 float64 = (0.1 * 1.5)
                                                                v38 -= 1
                                                                _ = v41
                                                            } else {
                                                                var v42 int8 = -1
                                                                v43 := GenS1{F0: (!false)}
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        case "a":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v44 uint64 = v38
                                                                _ = v44
                                                            case "\n\t":
                                                                var v45 uint16 = 2
                                                                _ = v45
                                                            case "":
                                                                v46 := GenS2{F0: "a", F1: false, F2: 167}
                                                                _ = v46
                                                            }
                                                        default:
                                                            switch ("" + "foo bar") {
                                                            case "a":
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v47 uint8 = 89
                                                                    _ = v47
                                                                case "foo bar":
                                                                    var v48 uint16 = v30
                                                                    _ = v48
                                                                case "a":
                                                                    if false {
                                                                        var v49 int32 = -785
                                                                        var v50 bool = (2 < uint8(v38))
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if v50 {
                                                                                v38 += (v38 * (v38 | 42))
                                                                                v51 := GenS7{F0: (v28 >> ((862) & 63))}
                                                                                _ = v51
                                                                            } else {
                                                                                switch (fi23 / ((42) | 1)) {
                                                                                case 32767:
                                                                                    v38 = 934
                                                                                case -819:
                                                                                    var v52 uint16 = 851
                                                                                    _ = v52
                                                                                }
                                                                                v53 := GenS5{F0: v50}
                                                                                switch fi19 {
                                                                                case -2147483648:
                                                                                    _ = len(v5)
                                                                                }
                                                                                _ = v53
                                                                            }
                                                                        case "foo bar":
                                                                            v3 = append(v3, ("\n\t" + "\n\t"))
                                                                            if false {
                                                                                v49 += -2147483648
                                                                            }
                                                                        default:
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                }
                                                                var v54 int16 = -1
                                                                fmt.Println(fi19, v54)
                                                                _ = v54
                                                            }
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v55 int8 = (v28 & v28)
                                                        _ = v55
                                                    case "":
                                                        _ = len(v1)
                                                    }
                                                    if (v30 < 1) {
                                                        v30 = (v30 ^ (^(v30)))
                                                        switch ("a" + "foo bar") {
                                                        case "foo bar":
                                                            var v56 int16 = -32768
                                                            _ = v56
                                                        case "\n\t":
                                                            v4[1] = "\n\t"
                                                            v57 := GenS3{F0: 2147483648, F1: 2147483647, F2: v38}
                                                            v28 = (v28 - 42)
                                                            _ = v57
                                                        }
                                                        if (255 < 127) {
                                                            var v58 bool = false
                                                            if (!v58) {
                                                                v38 = uint64(v30)
                                                            }
                                                            if (v58 || (v58 && (-876 > 314))) {
                                                                var v59 uint32 = 2147483648
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        }
                                                    }
                                                    v30 -= (v30 ^ (v30 % (((v30 - v30)) | 1)))
                                                    _ = v38
                                                } else {
                                                    var v60 int16 = 194
                                                    if (42 <= (v28 << ((558) & 63))) {
                                                        var v61 bool = (((v30 != 1) && true) || true)
                                                        _ = v61
                                                    } else {
                                                        fmt.Println(v28)
                                                    }
                                                    _ = v60
                                                }
                                            }
                                        }
                                        v2 = (fi19 >> ((0) & 63))
                                        var v62 int64 = -188
                                        _ = v30
                                        _ = v62
                                    case "\n\t":
                                        v28 += v28
                                    default:
                                        v63 := GenS8{F0: -345, F1: ("foo bar" + "")}
                                        _ = v63
                                    }
                                    _ = v28
                                }
                            }
                            v2 -= v2
                        }
                        switch v2 {
                        case -9223372036854775808:
                            v64 := GenS7{F0: 1}
                            _ = v64
                        case -197:
                            switch (("a" + "") + "a") {
                            case "foo bar":
                                var v65 int64 = 1
                                _ = v65
                            case "\n\t":
                                v66 := GenS1{F0: (true && true)}
                                _ = v66
                            case "a":
                                v67 := GenS9{F0: 65535}
                                _ = v67
                            }
                            if true {
                                v2 = fi19
                                v2 = fi23
                                var v68 int16 = 32767
                                _ = v68
                            }
                            var v69 uint8 = 128
                            _ = v69
                        }
                        _ = fi23
                    }
                    var v70 uint16 = 255
                    for fi71 := 0; fi71 < 2; fi71++ {
                        v2 += (-(v2))
                        _ = fi71
                    }
                    _ = v70
                }
                fmt.Println(v2, fi19)
            }
            switch v2 {
            case 32768:
                v72 := GenS10{F0: 0, F1: 127}
                _ = v72
            case 127:
                for fi73 := 0; fi73 < 10; fi73++ {
                    switch "foo bar" {
                    case "":
                        v2 = v2
                    case "a":
                        _ = len(v4)
                    default:
                        v74 := GenS3{F0: 65535, F1: 42, F2: 588}
                        _ = v74
                    }
                    var v75 int64 = -9223372036854775808
                    _ = fi73
                    _ = v75
                }
            case 560:
                fmt.Println(fi19, v2)
            default:
                for fi76 := 0; fi76 < 2; fi76++ {
                    fmt.Println(fi76)
                    v2 -= 32767
                    _ = fi76
                }
            }
            v2 += (fi19 * 32768)
            _ = fi19
        }
        var v77 uint16 = 255
        _ = v77
    }
    switch v2 {
    case 469:
        for fi78 := 0; fi78 < 3; fi78++ {
            var v79 uint16 = 318
            _ = fi78
            _ = v79
        }
    case -2147483648:
        var v80 int16 = -32768
        _ = v80
    default:
        for fi81 := 0; fi81 < 12; fi81++ {
            for fi82 := 0; fi82 < 12; fi82++ {
                v2 = fi81
                v83 := v4[(v2 * (-(fi81)))]
                switch ("foo bar" + (v83 + ((v83 + v83) + "foo bar"))) {
                case "\n\t":
                    var v84 int32 = -2147483648
                    _ = v84
                case "foo bar":
                    if true {
                        var v85 uint16 = 1
                        v83 = v83
                        v86 := GenS3{F0: 4294967295, F1: 0, F2: 872}
                        _ = v85
                        _ = v86
                    } else {
                        fmt.Println(fi81)
                        _ = len(v5)
                        v83 = v83
                        var v87 float32 = -2.25
                        _ = v87
                    }
                case "hello":
                    if false {
                        fmt.Println(fi81, fi82)
                        if (-1 != -674) {
                            v2 = (^(fi81))
                            v2 += (fi82 << ((61) & 63))
                            if false {
                                var v88 string = (v83 + v83)
                                _ = v88
                            }
                        }
                    } else {
                        v83 = v83
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v83 = "a"
                    }
                default:
                    switch v83 {
                    case "\n\t":
                        fmt.Println(v2, fi82, v2)
                    }
                }
                _ = fi82
                _ = v83
            }
            switch (fi81 / ((-974) | 1)) {
            case -968:
                v2 += v2
            case 9223372036854775807:
                _ = len(v3)
            }
            var v89 string = "hello"
            _ = fi81
            _ = v89
        }
    }
    v2 -= (-(v2))
    switch ("" + "foo bar") {
    case "\n\t":
        switch (-(v2)) {
        case 155:
            var v90 bool = true
            _ = v90
        case -997:
            v2 += v2
        case -884:
            var v91 string = (("" + "foo bar") + "")
            _ = v91
        default:
            if false {
                var v92 uint = 342
                v93 := GenS6{F0: 1, F1: 0, F2: 128}
                v94 := v4[128]
                _ = v92
                _ = v93
                _ = v94
            }
        }
    default:
        delete(v4, v2)
    }
    if ((0 >= 883) && true) {
        var v95 string = "hello"
        var v96 uint16 = 0
        for fi97 := 0; fi97 < 3; fi97++ {
            v95 = "hello"
            var v98 float32 = float32(1.5)
            var v99 float64 = 0.0
            _ = fi97
            _ = v98
            _ = v99
        }
        _ = v95
        _ = v96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
