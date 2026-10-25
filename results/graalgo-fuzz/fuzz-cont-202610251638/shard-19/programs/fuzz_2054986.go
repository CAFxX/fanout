package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 uint32 = 4294967295
    var v4 string = "foo bar"
    v5 := GenS1{F0: 9223372036854775807, F1: v2, F2: 0}
    v6 := []int{-9223372036854775808, 1, int(v3)}
    switch -1 {
    case 127:
        if true {
            v3 = v3
            var v7 int16 = 32767
            v7 = (v7 % ((v7) | 1))
            _ = v7
        }
    case -9223372036854775808:
        for fi8 := 0; fi8 < 3; fi8++ {
            v3 = v3
            _ = fi8
        }
    case 128:
        v2 += (v2 | (v2 & 670))
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        switch (fi9 | (^(fi9))) {
        case 0:
            var v10 int32 = 438
            _ = v10
        default:
            var v11 uint64 = (v2 | v2)
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            switch fi12 {
            case 4294967295:
                if false {
                    _ = len(v6)
                } else {
                    var v14 string = "hello"
                    switch v14 {
                    case "hello":
                        v15 := GenS1{F0: 1, F1: (v2 ^ 236), F2: 191}
                        _ = v15
                    case "foo bar":
                        _ = len(v6)
                        switch (v4 + (("a" + v14) + ("foo bar" + v4))) {
                        case "hello":
                            v16 := GenS1{F0: -1, F1: v2, F2: 127}
                            _ = v16
                        case "foo bar":
                            if (682 > 0) {
                                v14 = "foo bar"
                                v3 = v3
                            }
                        }
                    case "":
                        switch (-(fi13)) {
                        case -9223372036854775808:
                            if (1000 <= 1) {
                                switch (fi13 & (^(fi12))) {
                                case 9223372036854775807:
                                    fmt.Println(v14)
                                case -517:
                                    var v17 bool = true
                                    _ = v17
                                default:
                                    var v18 bool = (1e10 == (0.0 - ((1e10 + -2.25) / 1.0)))
                                    _ = v18
                                }
                            } else {
                                v4 = v14
                            }
                            fmt.Println(v3, fi12)
                        case 4294967295:
                            switch -924 {
                            case 357:
                                v19 := GenS1{F0: 1, F1: v2, F2: 61}
                                _ = v19
                            case 0:
                                _ = len(v1)
                            default:
                                v20 := GenS1{F0: 1, F1: (v2 & 18446744073709551615), F2: 111}
                                _ = v20
                            }
                        }
                    }
                    _ = v14
                }
                v1[((((((v4 + "\n\t") + v4) + v4) + v4) + "a") + ((v4 + (((v4 + v4) + "a") + "a")) + "hello"))] = (-(fi13))
            case 32767:
                switch (((("hello" + v4) + v4) + v4) + "foo bar") {
                case "foo bar":
                    v3 += (^(v3))
                case "hello":
                    _ = len(v1)
                default:
                    v21 := GenS1{F0: -1, F1: (v2 & (v2 | v2)), F2: 42}
                    _ = v21
                }
            default:
                if ((!true) && (v4 > v4)) {
                    var v22 string = ("hello" + "")
                    v6 = append(v6, (fi13 ^ (fi13 + 32767)))
                    var v23 uint32 = 0
                    _ = v22
                    _ = v23
                } else {
                    v4 = v4
                }
            }
            v4 = "hello"
            _ = fi13
        }
        _ = fi12
    }
    v4 = v4
    if false {
        switch "" {
        case "":
            v3 -= v3
        case "\n\t":
            if true {
                for fi24 := 0; fi24 < 9; fi24++ {
                    v25 := GenS1{F0: -334, F1: v2, F2: 151}
                    _ = fi24
                    _ = v25
                }
            } else {
                v3 = (v3 | (v3 ^ v3))
                v3 = 1
                switch "a" {
                case "hello":
                    _ = len(v1)
                default:
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 float32 = 1e10
                        _ = fi26
                        _ = v27
                    }
                }
            }
        default:
            if (uint(v2) == 42) {
                v28 := GenS2{F0: 1, F1: 4294967295, F2: 90}
                v4 = "\n\t"
                for fi29 := 0; fi29 < 16; fi29++ {
                    v6 = append(v6, 32768)
                    var v30 uint = 500
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi29
                    _ = v30
                }
                _ = v28
            } else {
                var v31 int8 = 42
                var v32 int16 = 32767
                switch 533 {
                case 32767:
                    v32 += (^(v32))
                case -9223372036854775808:
                    var v33 string = (v4 + v4)
                    _ = v33
                case 4294967295:
                    for fi34 := 0; fi34 < 8; fi34++ {
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                        _ = fi34
                    }
                default:
                    v31 -= v31
                }
                _ = v31
                _ = v32
            }
        }
        if (!true) {
            var v35 uint32 = v3
            v36 := v1["hello"]
            _ = v35
            _ = v36
        }
        if true {
            if len(v6) > 0 {
                v6[0] = 0
            }
            var v37 uint = 42
            _ = v37
        }
    } else {
        _ = len(v6)
        v2 = (^(v2))
        if false {
            v6 = append(v6, -464)
            for fi38 := 0; fi38 < 3; fi38++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v2 = (v2 ^ v2)
                _ = fi38
            }
            if true {
                if ((0 >= 0) || (!((100 >= 100) || ((255 != 4294967295) || (v4 == v4))))) {
                    switch v4 {
                    case "foo bar":
                        for fi39 := 0; fi39 < 9; fi39++ {
                            v3 -= v3
                            v4 = v4
                            var v40 int = (fi39 + fi39)
                            _ = fi39
                            _ = v40
                        }
                    default:
                        if false {
                            v3 = v3
                            if (!false) {
                                v41 := v1[(v4 + (v4 + ""))]
                                v3 += v3
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 12; fi42++ {
                                if false {
                                    v6 = append(v6, 0)
                                    v1[(v4 + v4)] = fi42
                                    var v43 uint64 = v2
                                    _ = v43
                                } else {
                                    v44 := GenS3{F0: 65535, F1: v3, F2: (^(fi42))}
                                    _ = v44
                                }
                                if ((0 >= 9223372036854775807) && true) {
                                    if false {
                                        for fi45 := 0; fi45 < 15; fi45++ {
                                            fmt.Println(fi42, v4)
                                            if false {
                                                switch "hello" {
                                                case "":
                                                    if true {
                                                        fmt.Println(v2)
                                                        var v46 uint16 = 0
                                                        _ = len(v1)
                                                        _ = v46
                                                    }
                                                case "a":
                                                    var v47 int64 = 9223372036854775807
                                                    _ = v47
                                                default:
                                                    var v48 uint = 395
                                                    _ = v48
                                                }
                                            }
                                            if true {
                                                switch (fi42 * fi45) {
                                                case -2147483648:
                                                    if (!true) {
                                                        var v49 uint64 = 625
                                                        var v50 int16 = 1000
                                                        var v51 int64 = 9223372036854775807
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        fmt.Println(fi45, v4)
                                                    }
                                                case 32767:
                                                    v52 := v1[v4]
                                                    _ = v52
                                                case 803:
                                                    var v53 string = v4
                                                    _ = v53
                                                default:
                                                    v2 = (v2 | 42)
                                                }
                                            } else {
                                                var v54 int = (fi42 & 1)
                                                _ = v54
                                            }
                                            v55 := GenS1{F0: -1, F1: 530, F2: 0}
                                            var v56 int64 = -1
                                            _ = fi45
                                            _ = v55
                                            _ = v56
                                        }
                                    } else {
                                        switch "a" {
                                        case "\n\t":
                                            var v57 uint32 = 2147483648
                                            _ = v57
                                        case "a":
                                            var v58 int = (fi42 >> ((65535) & 63))
                                            _ = v58
                                        }
                                        for fi59 := 0; fi59 < 18; fi59++ {
                                            var v60 uint16 = 60
                                            var v61 uint16 = uint16(v2)
                                            _ = fi59
                                            _ = v60
                                            _ = v61
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = fi42
                                        }
                                    }
                                    fmt.Println(fi42, v2)
                                } else {
                                    var v62 string = v4
                                    var v63 int = fi42
                                    if (32767 != 152) {
                                        if (!false) {
                                            for fi64 := 0; fi64 < 15; fi64++ {
                                                var v65 int16 = int16(v3)
                                                v65 += v65
                                                _ = fi64
                                                _ = v65
                                            }
                                            fmt.Println(v62, v4)
                                        }
                                        for fi66 := 0; fi66 < 16; fi66++ {
                                            switch v4 {
                                            case "hello":
                                                fmt.Println(fi66, v2, v2)
                                            default:
                                                var v67 int64 = 0
                                                _ = v67
                                            }
                                            _ = fi66
                                        }
                                    }
                                    _ = v62
                                    _ = v63
                                }
                                v68 := GenS1{F0: -9223372036854775808, F1: 1, F2: 2}
                                _ = fi42
                                _ = v68
                            }
                        }
                    }
                    fmt.Println(v3)
                } else {
                    for fi69 := 0; fi69 < 1; fi69++ {
                        if true {
                            var v70 uint = 42
                            v71 := GenS2{F0: 127, F1: (v3 | (v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))), F2: 232}
                            _ = v70
                            _ = v71
                        }
                        var v72 int8 = 100
                        switch 437 {
                        case 841:
                            v2 = 18446744073709551615
                        case 128:
                            v2 -= 427
                        case 711:
                            var v73 uint32 = v3
                            _ = v73
                        default:
                            switch fi69 {
                            case 42:
                                switch "" {
                                case "hello":
                                    var v74 bool = (0 != 0)
                                    _ = v74
                                case "":
                                    if true {
                                        var v75 float64 = 3.14159
                                        if true {
                                            for fi76 := 0; fi76 < 20; fi76++ {
                                                if ((v3 * 4294967295) == v3) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v77 := GenS4{F0: 625, F1: 299}
                                                        _ = len(v1)
                                                        v2 = (v2 % (((v2 / (((v2 | v2)) | 1))) | 1))
                                                        _ = v77
                                                    case "":
                                                        var v78 float64 = (float64(1e10) + 1e10)
                                                        _ = v78
                                                    default:
                                                        if false {
                                                            v1[("\n\t" + v4)] = fi69
                                                            v79 := GenS1{F0: -1, F1: v2, F2: 144}
                                                            _ = v79
                                                        }
                                                    }
                                                    fmt.Println(fi76)
                                                    switch v4 {
                                                    case "a":
                                                        var v80 uint = 1
                                                        _ = v80
                                                    case "":
                                                        _ = len(v6)
                                                    default:
                                                        v81 := GenS2{F0: 1, F1: 254, F2: 127}
                                                        _ = v81
                                                    }
                                                }
                                                var v82 int16 = 1
                                                var v83 int16 = (v82 << ((18446744073709551615) & 63))
                                                _ = fi76
                                                _ = v82
                                                _ = v83
                                            }
                                        } else {
                                            v84 := GenS3{F0: 32767, F1: v3, F2: fi69}
                                            for fi85 := 0; fi85 < 15; fi85++ {
                                                if true {
                                                    v86 := GenS5{F0: 2, F1: 65535}
                                                    _ = v86
                                                } else {
                                                    v87 := GenS6{F0: fi85, F1: (v72 - v72)}
                                                    _ = v87
                                                }
                                                _ = fi85
                                            }
                                            _ = v84
                                        }
                                        _ = v75
                                    }
                                }
                            case -2147483648:
                                if ((fi69 | 42) < fi69) {
                                    v88 := GenS7{F0: 42}
                                    for fi89 := 0; fi89 < 3; fi89++ {
                                        if len(v6) > 0 {
                                            v6[0] = (fi69 / ((91) | 1))
                                        }
                                        v4 = v4
                                        _ = fi89
                                    }
                                    _ = v88
                                } else {
                                    var v90 uint8 = 226
                                    v72 -= (v72 >> ((529) & 63))
                                    v91 := GenS8{F0: 9223372036854775807}
                                    _ = v90
                                    _ = v91
                                }
                            default:
                                for fi92 := 0; fi92 < 4; fi92++ {
                                    fmt.Println(fi69, v4, v2)
                                    delete(v1, "\n\t")
                                    v6 = append(v6, 127)
                                    var v93 int32 = -1
                                    _ = fi92
                                    _ = v93
                                }
                            }
                        }
                        _ = fi69
                        _ = v72
                    }
                }
                switch 723 {
                case 0:
                    var v94 int = 42
                    _ = v94
                default:
                    v6 = append(v6, 128)
                }
            } else {
                switch 617 {
                case -70:
                    v2 = 680
                case 42:
                    switch 702 {
                    case 42:
                        v95 := GenS3{F0: 981, F1: v3, F2: 0}
                        _ = v95
                    case 401:
                        for fi96 := 0; fi96 < 16; fi96++ {
                            v3 = (v3 - v3)
                            var v97 int8 = -1
                            _ = fi96
                            _ = v97
                        }
                    }
                case 6:
                    v98 := v1[(("\n\t" + v4) + v4)]
                    _ = v98
                }
            }
        } else {
            _ = len(v6)
        }
    }
    v99 := v1[v4]
    v1[v4] = (-(v99))
    var v100 float64 = 0.1
    v101 := GenS6{F0: (v99 + v99), F1: 38}
    if false {
        v2 = v2
        if (false || (!true)) {
            _ = len(v6)
        }
    } else {
        switch v4 {
        case "\n\t":
            v102 := v1[(v4 + v4)]
            _ = v102
        case "a":
            v3 = (v3 ^ v3)
        }
        v1[v4] = -9223372036854775808
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v103 uint16 = 32767
    switch 549 {
    case 32768:
        switch (v99 / ((v99) | 1)) {
        case 2147483647:
            v104 := GenS8{F0: 42}
            _ = v104
        case 9223372036854775807:
            if false {
                if (!true) {
                    v3 += (^(v3))
                    v1["\n\t"] = v99
                    switch (v4 + (((v4 + v4) + (v4 + ("hello" + ((v4 + (v4 + v4)) + v4)))) + v4)) {
                    case "\n\t":
                        fmt.Println(v4, v4)
                    }
                } else {
                    v1["hello"] = v99
                }
            }
        default:
            switch (v4 + v4) {
            case "a":
                _ = len(v6)
            case "hello":
                if (!true) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    fmt.Println(v2, v2)
                    v4 = v4
                } else {
                    for fi105 := 0; fi105 < 17; fi105++ {
                        delete(v1, "a")
                        _ = fi105
                    }
                    v103 = v103
                }
            case "foo bar":
                fmt.Println(v103, v99, v2)
            default:
                fmt.Println(v103, v3, v4)
            }
        }
    case 128:
        v106 := GenS7{F0: v3}
        _ = v106
    case 9223372036854775807:
        var v107 int32 = 0
        _ = v107
    }
    v3 -= v3
    v4 = v4
    switch v99 {
    case -517:
        v100 = v100
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
