package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 int64; F2 int64 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]string)
    var v3 uint32 = 403
    v4 := make(map[int]int)
    if false {
        var v5 uint16 = 1
        v5 -= v5
        var v6 int = 1
        _ = v5
        _ = v6
    }
    v3 = v3
    if ((42 == -29) || (0.0 == 3.14159)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 uint = 385
            if true {
                if (false || (!true)) {
                    v9 := []int{fi7, (-(fi7))}
                    v10 := v4[fi7]
                    _ = v10
                    _ = v9
                } else {
                    _ = len(v2)
                    switch -2147483648 {
                    case 506:
                        fmt.Println(v3, fi7)
                    default:
                        var v11 uint32 = (v3 & 526)
                        _ = v11
                    }
                }
                if (v3 >= uint32(v8)) {
                    var v12 int16 = int16(v8)
                    _ = v12
                } else {
                    delete(v2, (v3 << ((v8) & 63)))
                    v8 += uint(v3)
                }
                var v13 bool = (127 <= fi7)
                _ = v13
            } else {
                if true {
                    fmt.Println(v3, fi7, fi7)
                }
                var v14 string = ""
                _ = v14
            }
            v15 := []int{(fi7 >> (((^(v8))) & 63))}
            _ = fi7
            _ = v15
            _ = v8
        }
    } else {
        switch (("foo bar" + "hello") + ("\n\t" + (("foo bar" + ("" + "hello")) + "hello"))) {
        case "foo bar":
            switch 1 {
            case 0:
                switch -45 {
                case 2147483647:
                    if (("hello" + "\n\t") <= ("" + "a")) {
                        fmt.Println(v1)
                        v16 := GenS1{F0: 0, F1: 1, F2: v1}
                        _ = v16
                    } else {
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v1)
                        default:
                            var v17 uint8 = 128
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 11; fi18++ {
                            v3 -= 4294967295
                            v19 := GenS1{F0: 1, F1: 845, F2: (-(v1))}
                            fmt.Println(v3)
                            _ = fi18
                            _ = v19
                        }
                    }
                case 4294967295:
                    v1 = (v1 >> ((2) & 63))
                default:
                    v1 += (v1 * int16(v3))
                }
            case -169:
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 uint64 = 2
                    v21 -= (^(v21))
                    _ = fi20
                    _ = v21
                }
            case -719:
                v3 = 2147483647
            default:
                var v22 int = 42
                _ = v22
            }
        case "a":
            fmt.Println(v1, v1)
        case "\n\t":
            switch -9223372036854775808 {
            case 0:
                var v23 bool = (758 <= 776)
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS1{F0: 127, F1: 9223372036854775807, F2: v1}
        if (2 <= (v3 / ((v3) | 1))) {
            var v26 int8 = -110
            _ = v26
        }
        if true {
            if ((1.5 / 1.5) > (3.14159 + 0.0)) {
                v27 := GenS1{F0: 118, F1: -1, F2: (v1 + v1)}
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 uint = 0
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                var v30 float32 = float32((0.1 * 1.5))
                v31 := GenS1{F0: 42, F1: 9223372036854775807, F2: (v1 % ((v1) | 1))}
                _ = v30
                _ = v31
            }
            v3 += 489
            if (true && false) {
                if true {
                    v32 := []int{fi24}
                    _ = v32
                }
                v33 := v2[v3]
                _ = v33
            }
        } else {
            var v34 uint = 65535
            _ = v34
        }
        _ = fi24
        _ = v25
    }
    switch "" {
    case "\n\t":
        v35 := []int{4294967295, 127}
        _ = v35
    case "a":
        v36 := GenS2{F0: 2, F1: -128}
        _ = v36
    }
    if true {
        v37 := GenS1{F0: 177, F1: 9223372036854775807, F2: (v1 >> ((42) & 63))}
        _ = v37
    }
    v1 += v1
    for fi38 := 0; fi38 < 15; fi38++ {
        v39 := v4[-319]
        _ = fi38
        _ = v39
    }
    if (!((uint(v3) >= 65535) || false)) {
        if (!(2 > 1)) {
            if (((!true) && (false || (0.0 == (((1.5 - -2.25) / float32(float64((((0.1 + 3.14159) + (((0.1 / 1.5) - (1.5 - 1.5)) + 3.14159)) + ((float32((0.0 + 0.1)) / 0.1) - 0.0))))) + (1e10 - 1.5))))) && (true && (-2.25 <= 0.0))) {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v3 = v3
                    switch "foo bar" {
                    case "":
                        if (-1 <= 42) {
                            var v41 uint64 = 723
                            v42 := v4[(fi40 | int(v3))]
                            _ = v41
                            _ = v42
                        }
                    case "foo bar":
                        if true {
                            var v43 int8 = 127
                            _ = v43
                        } else {
                            if false {
                                if true {
                                    if (true || false) {
                                        var v44 uint16 = 32767
                                        v44 = (v44 % ((v44) | 1))
                                        var v45 float32 = (1e10 - 1.5)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            if (!(-1 < fi40)) {
                                                fmt.Println(fi40, v1, v3)
                                                var v46 uint16 = 2
                                                if false {
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if true {
                                                                v47 := GenS2{F0: 32767, F1: -32}
                                                                _ = v47
                                                            } else {
                                                                var v48 int = fi40
                                                                v49 := []int{}
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        case "\n\t":
                                                            if false {
                                                                for fi50 := 0; fi50 < 15; fi50++ {
                                                                    var v51 int16 = -32768
                                                                    _ = fi50
                                                                    _ = v51
                                                                }
                                                            } else {
                                                                var v52 float32 = 0.1
                                                                v3 += v3
                                                                _ = v52
                                                            }
                                                        case "a":
                                                            var v53 float32 = (-2.25 / ((1.5 + float32(1.5)) / 1e10))
                                                            _ = v53
                                                        default:
                                                            if (false || (42 > 2)) {
                                                                v3 = v3
                                                                v54 := []int{fi40}
                                                                _ = v54
                                                            } else {
                                                                v46 += v46
                                                                v55 := GenS3{F0: 140, F1: "", F2: v1}
                                                                _ = v55
                                                            }
                                                        }
                                                    }
                                                    switch (-(fi40)) {
                                                    case 32768:
                                                        var v56 int = fi40
                                                        _ = v56
                                                    case -2147483648:
                                                        v46 = (v46 ^ v46)
                                                    case -674:
                                                        if ((1.5 / 1.0) != 3.14159) {
                                                            v3 = (^(v3))
                                                            if true {
                                                                v57 := GenS1{F0: 1, F1: 1, F2: v1}
                                                                v4[(^(fi40))] = 423
                                                                _ = v57
                                                            }
                                                            var v58 int = -382
                                                            _ = v58
                                                        } else {
                                                            for fi59 := 0; fi59 < 5; fi59++ {
                                                                var v60 uint16 = v46
                                                                v60 -= 65535
                                                                _ = fi59
                                                                _ = v60
                                                            }
                                                            switch 806 {
                                                            case 1:
                                                                var v61 bool = false
                                                                _ = v61
                                                            }
                                                            var v62 uint = 628
                                                            _ = v62
                                                        }
                                                    }
                                                    for fi63 := 0; fi63 < 1; fi63++ {
                                                        v3 -= 432
                                                        var v64 uint32 = 2
                                                        fmt.Println(v46, fi40)
                                                        _ = fi63
                                                        _ = v64
                                                    }
                                                }
                                                _ = v46
                                            } else {
                                                v4[fi40] = (fi40 ^ fi40)
                                                for fi65 := 0; fi65 < 2; fi65++ {
                                                    fmt.Println(fi65)
                                                    v66 := GenS4{F0: 32767, F1: (fi40 << ((2) & 63))}
                                                    v67 := GenS3{F0: 247, F1: "hello", F2: (v1 | v1)}
                                                    _ = fi65
                                                    _ = v66
                                                    _ = v67
                                                }
                                                for fi68 := 0; fi68 < 9; fi68++ {
                                                    v69 := v4[fi68]
                                                    _ = fi68
                                                    _ = v69
                                                }
                                            }
                                        case "hello":
                                            for fi70 := 0; fi70 < 19; fi70++ {
                                                v3 = (v3 / ((42) | 1))
                                                _ = fi70
                                            }
                                        case "a":
                                            var v71 int32 = -477
                                            _ = v71
                                        }
                                        if (65535 < 2) {
                                            if true {
                                                v2[65535] = ""
                                            }
                                        } else {
                                            switch ("" + ("" + "foo bar")) {
                                            case "foo bar":
                                                var v72 int8 = 64
                                                _ = v72
                                            case "":
                                                fmt.Println(fi40, v1, v3)
                                            case "\n\t":
                                                var v73 string = ((("\n\t" + "hello") + "") + ("a" + ""))
                                                _ = v73
                                            default:
                                                if (!(!false)) {
                                                    v74 := v2[(v3 << ((42) & 63))]
                                                    delete(v2, v3)
                                                    _ = v74
                                                } else {
                                                    v3 = v3
                                                    if false {
                                                        v75 := v4[fi40]
                                                        for fi76 := 0; fi76 < 10; fi76++ {
                                                            v77 := []int{(^(fi40)), (^(v75))}
                                                            fmt.Println(v1)
                                                            fmt.Println(v3)
                                                            v78 := GenS1{F0: 181, F1: int64(v3), F2: v1}
                                                            v75 += 32767
                                                            _ = fi76
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        v3 += (^(v3))
                                                        _ = v75
                                                    } else {
                                                        switch 4294967295 {
                                                        case -9223372036854775808:
                                                            v79 := []int{fi40}
                                                            _ = v79
                                                        case 833:
                                                            var v80 int16 = (^(v1))
                                                            _ = v80
                                                        }
                                                        v3 -= 2147483647
                                                    }
                                                    switch ("foo bar" + ("foo bar" + ("\n\t" + "foo bar"))) {
                                                    case "hello":
                                                        var v81 string = "\n\t"
                                                        _ = v81
                                                    case "\n\t":
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                            v82 := []int{(fi40 >> ((65535) & 63))}
                                            v1 += -1
                                            _ = v82
                                        }
                                        if (false || (!(!("a" == ("foo bar" + "\n\t"))))) {
                                            for fi83 := 0; fi83 < 18; fi83++ {
                                                v3 -= v3
                                                v4[(-(fi83))] = fi40
                                                _ = fi83
                                            }
                                            v1 -= v1
                                        }
                                    }
                                }
                                switch "a" {
                                case "":
                                    fmt.Println(v1)
                                case "a":
                                    v1 -= -1
                                case "\n\t":
                                    v4[fi40] = fi40
                                default:
                                    if (0 > 18446744073709551615) {
                                        fmt.Println(fi40, v3, v3)
                                        for fi84 := 0; fi84 < 15; fi84++ {
                                            v1 = v1
                                            _ = fi84
                                        }
                                    }
                                }
                            }
                        }
                    default:
                        for fi85 := 0; fi85 < 11; fi85++ {
                            switch (fi85 * (fi85 - 4294967295)) {
                            case 9223372036854775807:
                                if (2 >= 96) {
                                    v3 = (^(v3))
                                    var v86 int64 = 42
                                    _ = v86
                                } else {
                                    if true {
                                        v87 := []int{fi85}
                                        v1 = (^(v1))
                                        _ = v87
                                    } else {
                                        _ = len(v4)
                                    }
                                    v88 := GenS5{F0: 1}
                                    _ = v88
                                }
                            case 4294967295:
                                fmt.Println(v3, v1)
                            }
                            _ = fi85
                        }
                    }
                    _ = fi40
                }
                if (2147483647 >= 9223372036854775807) {
                    var v89 uint32 = (^(v3))
                    switch -1 {
                    case 9223372036854775807:
                        var v90 bool = (1 <= (v89 + 42))
                        _ = v90
                    case -712:
                        if true {
                            fmt.Println(v3)
                            v3 = 42
                        } else {
                            var v91 int64 = -9223372036854775808
                            v92 := GenS3{F0: 193, F1: ("a" + ("foo bar" + "")), F2: v1}
                            _ = v91
                            _ = v92
                        }
                    case 500:
                        switch -2147483648 {
                        case 1:
                            v93 := []int{32767}
                            _ = v93
                        }
                    default:
                        var v94 uint64 = 0
                        _ = v94
                    }
                    _ = v89
                }
                for fi95 := 0; fi95 < 13; fi95++ {
                    for fi96 := 0; fi96 < 5; fi96++ {
                        v3 = (^(v3))
                        _ = fi96
                    }
                    v97 := GenS6{F0: 42, F1: -554, F2: 9223372036854775807}
                    _ = fi95
                    _ = v97
                }
            }
        }
        for fi98 := 0; fi98 < 3; fi98++ {
            v3 = 2147483648
            _ = fi98
        }
    } else {
        fmt.Println(v3, v3)
        var v99 float32 = -2.25
        _ = v99
    }
    if (true || false) {
        fmt.Println(v1)
        v1 = 1000
    } else {
        v3 = v3
        switch "foo bar" {
        case "foo bar":
            if false {
                v100 := []int{32767, -1, 1, 4294967295}
                v1 += (v1 + 183)
                v101 := GenS4{F0: 42, F1: -612}
                _ = v100
                _ = v101
            } else {
                for fi102 := 0; fi102 < 10; fi102++ {
                    var v103 int32 = 0
                    _ = fi102
                    _ = v103
                }
                switch "a" {
                case "foo bar":
                    _ = len(v4)
                case "":
                    for fi104 := 0; fi104 < 9; fi104++ {
                        fmt.Println(fi104, fi104, v3)
                        _ = fi104
                    }
                default:
                    var v105 uint = 65535
                    _ = v105
                }
            }
        case "":
            var v106 uint = 1
            _ = v106
        default:
            if false {
                if true {
                    if (2 != 69) {
                        if (100 < 42) {
                            if false {
                                v3 = 58
                                fmt.Println(v1, v3)
                            } else {
                                if (85 < 0) {
                                    var v107 int32 = -1
                                    v3 -= (v3 - (v3 * 2147483647))
                                    v108 := GenS4{F0: 2, F1: 42}
                                    _ = v107
                                    _ = v108
                                }
                            }
                        }
                    } else {
                        fmt.Println(v1, v1)
                        v109 := GenS6{F0: v1, F1: 0, F2: 42}
                        v110 := v2[v3]
                        _ = v109
                        _ = v110
                    }
                    v1 = (v1 >> ((1) & 63))
                } else {
                    var v111 uint = 393
                    v111 -= 18446744073709551615
                    _ = v111
                }
                var v112 float64 = float64(0.0)
                _ = v112
            }
        }
    }
    switch "" {
    case "hello":
        var v113 float32 = (0.1 - float32(float64(1e10)))
        _ = v113
    case "foo bar":
        v114 := []int{-9223372036854775808, 0, 127}
        _ = v114
    default:
        v115 := v2[v3]
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
