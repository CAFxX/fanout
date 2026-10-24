package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 int8; F2 int32 }
type GenS9 struct { F0 uint32; F1 int; F2 uint }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int8 }

func main() {
    var v1 uint16 = 42
    var v2 string = ("a" + "\n\t")
    var v3 int8 = 42
    var v4 float32 = 0.1
    v5 := []string{}
    v6 := make(map[string]int)
    v6["\n\t"] = 2147483647
    var v7 float32 = ((float32(1e10) + v4) * v4)
    _ = len(v5)
    var v8 int32 = -442
    if (!true) {
        fmt.Println(v2, v8)
        v5 = append(v5, (v2 + ""))
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = -1
        _ = fi9
        _ = v10
    }
    var v11 int16 = -32768
    var v12 float32 = (v7 + (float32(3.14159) * -2.25))
    var v13 int8 = -128
    if (1 <= 2) {
        var v14 bool = (false && ((-(v11)) == 1))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v4 = 1.5
        if false {
            var v16 int32 = (v8 & (-(v8)))
            _ = v16
        } else {
            switch (^(fi15)) {
            case -750:
                var v17 float32 = -2.25
                _ = v17
            case -602:
                if ((fi15 - (fi15 / (((fi15 ^ 9223372036854775807)) | 1))) >= 0) {
                    v6[v2] = -7
                    switch 4294967295 {
                    case 2147483647:
                        if true {
                            if false {
                                v18 := GenS1{F0: (v8 | (v8 * (v8 ^ -2147483648))), F1: v2}
                                var v19 int8 = v3
                                _ = v18
                                _ = v19
                            } else {
                                switch ((v2 + v2) + v2) {
                                case "":
                                    v20 := GenS2{F0: 255, F1: v13}
                                    _ = v20
                                case "\n\t":
                                    v21 := GenS2{F0: (^(v1)), F1: (v3 + (v3 * (v3 >> ((263) & 63))))}
                                    _ = v21
                                case "foo bar":
                                    switch fi15 {
                                    case -9:
                                        if len(v5) > 0 {
                                            v5[0] = ("" + (v2 + v2))
                                        }
                                    case 32767:
                                        for fi22 := 0; fi22 < 19; fi22++ {
                                            v23 := GenS3{F0: v11, F1: 1, F2: 2}
                                            _ = fi22
                                            _ = v23
                                        }
                                    case -1:
                                        v3 += v3
                                    default:
                                        if (42 <= 18446744073709551615) {
                                            if (v8 < -1) {
                                                if true {
                                                    var v24 int16 = v11
                                                    v3 -= (v13 >> ((507) & 63))
                                                    _ = v24
                                                }
                                            }
                                            var v25 bool = (!(255 == v1))
                                            switch (fi15 & (fi15 & fi15)) {
                                            case -543:
                                                fmt.Println(v4, v13, v7)
                                            case 33:
                                                switch "\n\t" {
                                                case "hello":
                                                    for fi26 := 0; fi26 < 15; fi26++ {
                                                        v27 := GenS4{F0: 255, F1: (v8 + v8), F2: 42}
                                                        v25 = v25
                                                        var v28 int = fi15
                                                        _ = fi26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                case "a":
                                                    var v29 uint16 = v1
                                                    _ = v29
                                                }
                                            default:
                                                v2 = (v2 + "\n\t")
                                            }
                                            _ = v25
                                        } else {
                                            v13 += v13
                                            var v30 int64 = 1
                                            var v31 uint32 = 2147483648
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                }
                            }
                        } else {
                            switch v2 {
                            case "foo bar":
                                v8 -= v8
                            }
                        }
                    }
                    switch ("foo bar" + v2) {
                    case "a":
                        var v32 string = (v2 + "a")
                        _ = v32
                    case "hello":
                        if (false && (-1 >= (^(v3)))) {
                            if true {
                                if true {
                                    if len(v5) > 0 {
                                        v5[0] = ("" + ("\n\t" + "foo bar"))
                                    }
                                }
                                switch (fi15 << ((255) & 63)) {
                                case -261:
                                    if false {
                                        if true {
                                            var v33 int8 = (-(v3))
                                            v33 = v3
                                            _ = v33
                                        }
                                        switch -2147483648 {
                                        case 175:
                                            if true {
                                                var v34 int8 = v3
                                                v34 += v13
                                                fmt.Println(v3, v4, v13)
                                                _ = v34
                                            } else {
                                                v1 = 42
                                                var v35 int16 = (v11 - v11)
                                                if (v2 >= (v2 + "foo bar")) {
                                                    switch (v2 + v2) {
                                                    case "foo bar":
                                                        v36 := GenS2{F0: v1, F1: (v13 * 100)}
                                                        _ = v36
                                                    default:
                                                        var v37 uint32 = 927
                                                        _ = v37
                                                    }
                                                    var v38 int8 = v3
                                                    _ = v38
                                                } else {
                                                    var v39 int64 = 319
                                                    var v40 int64 = v39
                                                    _ = len(v6)
                                                    _ = v39
                                                    _ = v40
                                                }
                                                _ = v35
                                            }
                                        default:
                                            fmt.Println(v1)
                                        }
                                    }
                                default:
                                    var v41 float32 = float32(0.0)
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 1; fi42++ {
                                    if false {
                                        if len(v5) > 0 {
                                            v5[0] = v2
                                        }
                                        switch v2 {
                                        case "a":
                                            v5 = append(v5, v2)
                                        default:
                                            v2 = ((v2 + "foo bar") + "a")
                                        }
                                        v43 := GenS2{F0: v1, F1: 100}
                                        var v44 uint8 = 145
                                        _ = v43
                                        _ = v44
                                    }
                                    v6[v2] = fi42
                                    _ = fi42
                                }
                            }
                        }
                    default:
                        for fi45 := 0; fi45 < 2; fi45++ {
                            v13 = v3
                            _ = fi45
                        }
                    }
                } else {
                    switch v2 {
                    case "hello":
                        for fi46 := 0; fi46 < 6; fi46++ {
                            v11 = v11
                            _ = fi46
                        }
                    case "a":
                        for fi47 := 0; fi47 < 4; fi47++ {
                            switch fi47 {
                            case 931:
                                var v48 int = fi47
                                _ = v48
                            case -364:
                                var v49 bool = ((true && true) && true)
                                _ = v49
                            case 32768:
                                var v50 uint16 = v1
                                _ = v50
                            }
                            _ = fi47
                        }
                    case "":
                        switch "a" {
                        case "foo bar":
                            v13 -= 116
                        default:
                            var v51 int32 = 2147483647
                            _ = v51
                        }
                    default:
                        v6["hello"] = 0
                    }
                }
            case -330:
                fmt.Println(fi15, v11, v1)
            }
        }
        v52 := GenS5{F0: 665}
        _ = fi15
        _ = v52
    }
    for fi53 := 0; fi53 < 18; fi53++ {
        if ((^(v8)) >= (-(v8))) {
            v13 += (v3 >> ((0) & 63))
            if false {
                fmt.Println(v8)
            } else {
                v54 := GenS5{F0: (fi53 >> ((42) & 63))}
                v55 := GenS6{F0: v11, F1: 42}
                for fi56 := 0; fi56 < 10; fi56++ {
                    v3 += (v3 | (v3 << ((382) & 63)))
                    switch (^(fi53)) {
                    case 0:
                        var v57 uint = 264
                        _ = v57
                    case -373:
                        switch fi56 {
                        case 4294967295:
                            switch "foo bar" {
                            case "\n\t":
                                v58 := GenS5{F0: (-(fi53))}
                                _ = v58
                            default:
                                if true {
                                    var v59 int16 = (v11 / ((v11) | 1))
                                    var v60 uint64 = 18446744073709551615
                                    switch -141 {
                                    case 32767:
                                        if ((!true) || (-1 >= v13)) {
                                            var v61 bool = false
                                            if len(v5) > 0 {
                                                v5[0] = v2
                                            }
                                            v61 = ((v61 || ((173 > (fi56 << ((1) & 63))) && true)) || (212 < -9223372036854775808))
                                            _ = v61
                                        } else {
                                            var v62 int16 = v11
                                            if (v1 > (^(v1))) {
                                                v12 = v4
                                                if false {
                                                    if true {
                                                        fmt.Println(v62)
                                                        var v63 float32 = (((((v12 - 1.5) * v7) + v12) * (1e10 + v4)) + float32(3.14159))
                                                        v59 = v59
                                                        _ = v63
                                                    }
                                                } else {
                                                    v64 := GenS2{F0: v1, F1: (v13 >> ((674) & 63))}
                                                    _ = v64
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v7 = -2.25
                                                case "a":
                                                    var v65 uint64 = v60
                                                    _ = v65
                                                default:
                                                    switch fi56 {
                                                    case 2147483647:
                                                        v66 := GenS7{F0: (v1 | v1)}
                                                        _ = v66
                                                    }
                                                }
                                                v60 = (^(v60))
                                            } else {
                                                v67 := GenS4{F0: 1, F1: (-(v8)), F2: 87}
                                                if (!false) {
                                                    if (v11 == -194) {
                                                        var v68 uint8 = 113
                                                        v4 -= v4
                                                        _ = v68
                                                    } else {
                                                        _ = len(v6)
                                                        v69 := GenS3{F0: (v11 & (v59 / (((v62 ^ (^(v62)))) | 1))), F1: 65535, F2: v60}
                                                        _ = v69
                                                    }
                                                    v8 -= (v8 << ((42) & 63))
                                                }
                                                _ = v67
                                            }
                                            _ = v62
                                        }
                                    case 492:
                                        var v70 float32 = v7
                                        _ = v70
                                    case 9223372036854775807:
                                        v2 = v2
                                    }
                                    _ = v59
                                    _ = v60
                                } else {
                                    v71 := GenS3{F0: int16(v1), F1: 1, F2: 63}
                                    v72 := GenS4{F0: 255, F1: (v8 + v8), F2: 42}
                                    var v73 int32 = 1
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                            }
                        case 487:
                            if true {
                                if false {
                                    if (!(32767 == (v11 + 32767))) {
                                        var v74 int = fi53
                                        var v75 int64 = 9223372036854775807
                                        _ = v74
                                        _ = v75
                                    } else {
                                        if (336 >= 321) {
                                            var v76 int8 = v3
                                            var v77 uint8 = 0
                                            fmt.Println(fi56, v76)
                                            _ = v76
                                            _ = v77
                                        }
                                        if ((516 < 2) || true) {
                                            if false {
                                                v13 = -128
                                                v4 = v7
                                            }
                                        } else {
                                            v3 = v3
                                            switch v2 {
                                            case "hello":
                                                var v78 uint64 = 2
                                                _ = v78
                                            default:
                                                fmt.Println(fi53, v3)
                                            }
                                        }
                                    }
                                    fmt.Println(fi56, v3)
                                    var v79 uint32 = 2
                                    var v80 uint16 = v1
                                    if (true || (!((v11 >> ((255) & 63)) <= v11))) {
                                        var v81 int64 = 200
                                        fmt.Println(v7)
                                        if ("hello" != ((("a" + "\n\t") + "hello") + (v2 + ("" + v2)))) {
                                            if false {
                                                var v82 int16 = 1
                                                v3 = v13
                                                _ = v82
                                            }
                                            var v83 bool = true
                                            v84 := GenS8{F0: true, F1: (v3 + v3), F2: v8}
                                            _ = v83
                                            _ = v84
                                        } else {
                                            if (((((((v12 / v4) - (v12 / 1.5)) + v12) - 1e10) + (1.5 * v4)) == v12) && (v79 > 2)) {
                                                if false {
                                                    fmt.Println(v3)
                                                }
                                            } else {
                                                if (!((v3 >> ((255) & 63)) < (v3 * (v13 >> ((402) & 63))))) {
                                                    var v85 uint = 512
                                                    if (uint32(v85) < 766) {
                                                        v86 := GenS7{F0: 32767}
                                                        _ = v86
                                                    } else {
                                                        switch (v2 + v2) {
                                                        case "foo bar":
                                                            fmt.Println(v81)
                                                        case "\n\t":
                                                            var v87 bool = (!true)
                                                            v11 -= v11
                                                            _ = v87
                                                        case "a":
                                                            v7 -= 1.5
                                                        default:
                                                            if (true || false) {
                                                                fmt.Println(v3, v13)
                                                                switch ((v2 + v2) + v2) {
                                                                case "foo bar":
                                                                    v2 = "a"
                                                                case "":
                                                                    fmt.Println(v85, v11)
                                                                }
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    switch int(v8) {
                                                                    case 846:
                                                                        fmt.Println(v2)
                                                                    case 0:
                                                                        var v88 int32 = v8
                                                                        _ = v88
                                                                    case -2147483648:
                                                                        var v89 uint32 = v79
                                                                        _ = v89
                                                                    }
                                                                case "a":
                                                                    fmt.Println(v11)
                                                                }
                                                                var v90 float32 = (v7 * v4)
                                                                var v91 float32 = (((v7 - (v4 + 0.0)) - v90) - 1.5)
                                                                v7 += 3.14159
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                        }
                                                        fmt.Println(v85, v81)
                                                    }
                                                    if true {
                                                        v92 := GenS4{F0: 41, F1: 1, F2: 0}
                                                        var v93 float32 = float32((1e10 + ((1e10 * (3.14159 + 1e10)) / float64((1e10 / v7)))))
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                    _ = v85
                                                }
                                                var v94 float32 = 1.5
                                                _ = v94
                                            }
                                        }
                                        _ = v81
                                    } else {
                                        var v95 uint32 = v79
                                        v96 := GenS9{F0: (v95 % (((v95 % (((v79 * v95)) | 1))) | 1)), F1: fi53, F2: 0}
                                        if (-1 >= (-(v8))) {
                                            if ((v11 * (-(v11))) > (v11 | (^(v11)))) {
                                                fmt.Println(v3, fi53)
                                                v97 := GenS9{F0: 315, F1: fi53, F2: 849}
                                                var v98 uint64 = 0
                                                _ = v97
                                                _ = v98
                                            }
                                            v3 -= v13
                                        } else {
                                            fmt.Println(fi53)
                                            v79 = v95
                                        }
                                        _ = v95
                                        _ = v96
                                    }
                                    _ = v79
                                    _ = v80
                                }
                            }
                        default:
                            v99 := GenS1{F0: (v8 & (v8 << ((802) & 63))), F1: ""}
                            _ = v99
                        }
                    case 42:
                        v1 = 467
                    }
                    var v100 bool = false
                    v7 += (v7 / (v12 / 1.0))
                    v101 := GenS5{F0: (fi56 % (((^(fi53))) | 1))}
                    fmt.Println(v3)
                    v102 := GenS10{F0: 58, F1: v1}
                    _ = fi56
                    _ = v100
                    _ = v101
                    _ = v102
                }
                _ = v54
                _ = v55
            }
            if true {
                if (521 == 1) {
                    switch (v2 + v2) {
                    case "foo bar":
                        for fi103 := 0; fi103 < 18; fi103++ {
                            v104 := GenS11{F0: (-(v13))}
                            _ = fi103
                            _ = v104
                        }
                    }
                    switch fi53 {
                    case -9223372036854775808:
                        switch "a" {
                        case "a":
                            if false {
                                v105 := v6[(("hello" + v2) + v2)]
                                var v106 int = (v105 << ((382) & 63))
                                fmt.Println(v8)
                                _ = v105
                                _ = v106
                            }
                        case "foo bar":
                            v7 = 3.14159
                        }
                    }
                    if (65535 != 4294967295) {
                        v11 = (v11 & 0)
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v8, v13)
                        case "\n\t":
                            v107 := GenS6{F0: (v11 + -32768), F1: 1}
                            _ = v107
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = v2
                            }
                        }
                        v108 := GenS10{F0: 1, F1: 65535}
                        _ = v108
                    }
                } else {
                    switch (-(fi53)) {
                    case 32767:
                        var v109 float32 = ((v4 - v12) / (-2.25 - 1.5))
                        _ = v109
                    case 4294967295:
                        var v110 int32 = (v8 | (v8 << ((0) & 63)))
                        _ = v110
                    default:
                        v111 := GenS7{F0: v1}
                        _ = v111
                    }
                }
            }
        }
        var v112 int32 = v8
        _ = fi53
        _ = v112
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
