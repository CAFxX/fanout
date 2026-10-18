package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int8; F1 int64 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]int)
    v3 := []string{"hello", "foo bar", ("a" + "a")}
    var v4 int = 128
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint16 = 32767
        v7 := GenS1{F0: 2, F1: 4294967295}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: 70, F1: 0}
            _ = fi8
            _ = v9
        }
    case "\n\t":
        var v10 int8 = -1
        _ = v10
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint = 4294967295
            _ = fi11
            _ = v12
        }
    }
    switch (v4 & 4294967295) {
    case 42:
        if (0 >= 1) {
            switch "" {
            case "a":
                fmt.Println(v4, v4)
            case "hello":
                v1 = 0.0
            }
            var v13 uint32 = 42
            if true {
                var v14 uint = 4294967295
                for fi15 := 0; fi15 < 5; fi15++ {
                    v13 = (v13 | v13)
                    _ = fi15
                }
                if (true && false) {
                    if (2 != 773) {
                        v13 = 2147483647
                    }
                    fmt.Println(v13, v13)
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v4 = fi16
                        var v17 uint32 = 65535
                        _ = fi16
                        _ = v17
                    }
                } else {
                    fmt.Println(v4, v13)
                    var v18 int = (v4 / ((v4) | 1))
                    _ = v18
                }
                _ = v14
            } else {
                var v19 int64 = -141
                var v20 bool = false
                _ = len(v2)
                _ = v19
                _ = v20
            }
            _ = v13
        } else {
            var v21 float64 = 0.0
            _ = v21
        }
    case 127:
        var v22 int16 = -416
        _ = v22
    case 9223372036854775807:
        if false {
            if (1e10 > v1) {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v1 = float64(3.14159)
                    _ = fi23
                }
                if true {
                    var v24 float64 = 0.1
                    switch v4 {
                    case -773:
                        var v25 string = (("a" + ("foo bar" + ("a" + "hello"))) + "foo bar")
                        _ = v25
                    }
                    switch (v4 - (v4 % (((v4 << ((2) & 63))) | 1))) {
                    case 1:
                        var v26 int = (^(v4))
                        _ = v26
                    default:
                        v24 = v24
                    }
                    _ = v24
                } else {
                    var v27 float32 = float32(v1)
                    fmt.Println(v1, v4, v4)
                    switch ("" + "a") {
                    case "a":
                        if (true && true) {
                            v1 = 0.0
                            for fi28 := 0; fi28 < 2; fi28++ {
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v27)
                                case "a":
                                    if (-15 >= 84) {
                                        if (((!(true && true)) && true) || false) {
                                            var v29 int32 = 0
                                            switch (v4 | v4) {
                                            case 0:
                                                for fi30 := 0; fi30 < 11; fi30++ {
                                                    fmt.Println(v1, fi28)
                                                    if (!true) {
                                                        if false {
                                                            v29 -= v29
                                                            if (748 != 466) {
                                                                var v31 int = v4
                                                                switch -88 {
                                                                case -1:
                                                                    v32 := GenS1{F0: 651, F1: 489}
                                                                    _ = v32
                                                                }
                                                                _ = v31
                                                            } else {
                                                                v29 -= v29
                                                                v33 := GenS2{F0: (v4 % (((fi28 / ((fi28) | 1))) | 1)), F1: ("a" + ("foo bar" + "a"))}
                                                                if true {
                                                                    v29 = v29
                                                                    _ = len(v3)
                                                                    v29 = 2147483647
                                                                    var v34 uint16 = 295
                                                                    _ = v34
                                                                }
                                                                _ = v33
                                                            }
                                                            v27 = 0.0
                                                        }
                                                        _ = len(v3)
                                                    }
                                                    _ = fi30
                                                }
                                            case -1:
                                                if false {
                                                    for fi35 := 0; fi35 < 16; fi35++ {
                                                        v36 := GenS2{F0: (fi28 << ((18446744073709551615) & 63)), F1: ""}
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v27 = v27
                                                        }
                                                        _ = fi35
                                                        _ = v36
                                                    }
                                                    switch (v4 << ((1) & 63)) {
                                                    case 2147483647:
                                                        if (!true) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            fmt.Println(v27, fi28, v27)
                                                        }
                                                    case 32767:
                                                        if true {
                                                            fmt.Println(v29, v29, v29)
                                                            v37 := v2[42]
                                                            _ = v37
                                                        }
                                                    default:
                                                        fmt.Println(v1, fi28)
                                                    }
                                                }
                                            }
                                            _ = v29
                                        }
                                        v38 := GenS3{F0: 42, F1: -96, F2: 65535}
                                        var v39 int16 = 1
                                        _ = v38
                                        _ = v39
                                    }
                                case "\n\t":
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        v40 := v2[32768]
                                        _ = v40
                                    case 2147483647:
                                        if ((32767 != 556) || (1 > 9223372036854775807)) {
                                            fmt.Println(v27, v1, v27)
                                            if (!false) {
                                                v41 := GenS4{F0: 65535, F1: true, F2: 908}
                                                _ = v41
                                            }
                                            var v42 uint = 951
                                            _ = v42
                                        } else {
                                            fmt.Println(v1, v1)
                                            if (954 < 18446744073709551615) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                v43 := GenS3{F0: 2147483648, F1: 42, F2: 0}
                                                v27 += (0.0 + v27)
                                                v27 = (v27 - -2.25)
                                                _ = v43
                                            }
                                        }
                                    }
                                default:
                                    for fi44 := 0; fi44 < 4; fi44++ {
                                        var v45 uint64 = 440
                                        var v46 int = (fi28 * (-(fi28)))
                                        var v47 float64 = (0.0 * v1)
                                        fmt.Println(v27)
                                        var v48 uint = 255
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                _ = fi28
                            }
                        } else {
                            fmt.Println(v4, v1, v4)
                            fmt.Println(v1, v27)
                            var v49 uint8 = 182
                            _ = v49
                        }
                    }
                    _ = v27
                }
                switch -653 {
                case 4294967295:
                    v1 = 1.5
                case 74:
                    v1 += 0.1
                case 413:
                    var v50 int8 = 1
                    _ = v50
                default:
                    var v51 bool = ((-2.25 - (((-2.25 * -2.25) + (float32(v1) / 1e10)) - 1e10)) != 0.1)
                    _ = v51
                }
            } else {
                v52 := GenS3{F0: 821, F1: 0, F2: 362}
                for fi53 := 0; fi53 < 3; fi53++ {
                    for fi54 := 0; fi54 < 15; fi54++ {
                        var v55 int32 = -189
                        v56 := GenS3{F0: 0, F1: 127, F2: 32767}
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v2)
                                _ = len(v2)
                            case "hello":
                                var v57 string = "\n\t"
                                _ = v57
                            case "a":
                                v55 = (v55 >> ((2) & 63))
                            default:
                                delete(v2, v4)
                            }
                            switch ("" + "hello") {
                            case "hello":
                                v1 -= 0.0
                            case "":
                                v55 += v55
                            case "foo bar":
                                fmt.Println(v4, v55, fi53)
                                if (211 >= 1) {
                                    fmt.Println(v4)
                                    var v58 uint8 = 255
                                    fmt.Println(v55)
                                    _ = v58
                                } else {
                                    if true {
                                        v55 = v55
                                        var v59 int16 = 1000
                                        _ = v59
                                    }
                                    v1 = v1
                                }
                                fmt.Println(fi54, fi53)
                            }
                            v55 = v55
                        }
                        v3 = append(v3, (("a" + "\n\t") + "a"))
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    fmt.Println(v1, v1, v4)
                    v1 = (((v1 - v1) + (v1 + v1)) / ((-2.25 / v1) * float64(0.0)))
                    _ = fi53
                }
                if false {
                    var v60 int16 = 1000
                    var v61 string = "foo bar"
                    v61 = (v61 + v61)
                    _ = v60
                    _ = v61
                }
                _ = v52
            }
            v4 -= v4
        }
    default:
        var v62 int8 = -45
        _ = v62
    }
    switch "" {
    case "\n\t":
        for fi63 := 0; fi63 < 13; fi63++ {
            v2[-1] = fi63
            _ = fi63
        }
    case "hello":
        var v64 int16 = 140
        _ = v64
    case "":
        v65 := GenS1{F0: 42, F1: 4294967295}
        _ = v65
    }
    v4 -= v4
    var v66 uint16 = 255
    if true {
        v67 := GenS3{F0: 416, F1: 100, F2: (v66 * (v66 << ((8) & 63)))}
        _ = v67
    } else {
        _ = len(v2)
        v66 -= v66
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    var v68 uint8 = 2
    var v69 int32 = -2147483648
    _ = len(v2)
    switch ("a" + "\n\t") {
    case "a":
        if (!(!false)) {
            var v70 bool = false
            _ = v70
        } else {
            v66 -= 32767
        }
    case "hello":
        v71 := GenS2{F0: (-(v4)), F1: ("a" + ("foo bar" + "a"))}
        _ = v71
    case "":
        if true {
            for fi72 := 0; fi72 < 18; fi72++ {
                fmt.Println(v4, v4)
                _ = fi72
            }
            v68 = 128
            for fi73 := 0; fi73 < 19; fi73++ {
                _ = len(v2)
                _ = fi73
            }
        }
    default:
        v2[(v4 % (((^(v4))) | 1))] = v4
    }
    var v74 int8 = -73
    switch v4 {
    case 42:
        for fi75 := 0; fi75 < 2; fi75++ {
            for fi76 := 0; fi76 < 13; fi76++ {
                switch "foo bar" {
                case "foo bar":
                    v77 := GenS5{F0: (^(v68)), F1: v69}
                    _ = v77
                case "hello":
                    v78 := GenS6{F0: (42 > -1), F1: v69}
                    _ = v78
                }
                _ = fi76
            }
            v79 := v2[(v4 | (fi75 % ((v4) | 1)))]
            _ = fi75
            _ = v79
        }
    case 177:
        switch "a" {
        case "foo bar":
            if ("a" > "a") {
                for fi80 := 0; fi80 < 8; fi80++ {
                    switch "hello" {
                    case "foo bar":
                        if false {
                            if (true || false) {
                                switch int(v66) {
                                case 42:
                                    for fi81 := 0; fi81 < 5; fi81++ {
                                        fmt.Println(v69)
                                        if (!true) {
                                            var v82 int64 = 1
                                            v4 = fi81
                                            v83 := v2[(v4 & 128)]
                                            _ = v82
                                            _ = v83
                                        } else {
                                            if false {
                                                var v84 string = ""
                                                v68 = 255
                                                _ = v84
                                            }
                                            fmt.Println(v68, v74)
                                        }
                                        if (v69 != v69) {
                                            var v85 uint = 65535
                                            v86 := GenS7{F0: 42, F1: (v68 % (((v68 & v68)) | 1))}
                                            _ = v85
                                            _ = v86
                                        }
                                        _ = fi81
                                    }
                                case 4294967295:
                                    v87 := GenS5{F0: 220, F1: -1}
                                    _ = v87
                                default:
                                    var v88 float32 = float32(v1)
                                    _ = v88
                                }
                            }
                        }
                    case "hello":
                        var v89 float32 = (3.14159 / 0.1)
                        _ = v89
                    }
                    _ = fi80
                }
                v1 += (-2.25 + v1)
            } else {
                var v90 int8 = v74
                switch ("a" + "") {
                case "foo bar":
                    var v91 int32 = v69
                    _ = v91
                case "a":
                    var v92 int8 = v90
                    _ = v92
                default:
                    var v93 float64 = (1e10 / 1.0)
                    _ = v93
                }
                var v94 int32 = 0
                _ = v90
                _ = v94
            }
        case "hello":
            fmt.Println(v1, v66, v68)
        default:
            var v95 uint = 1
            _ = v95
        }
    case 289:
        if ((v68 + v68) >= 24) {
            for fi96 := 0; fi96 < 20; fi96++ {
                for fi97 := 0; fi97 < 18; fi97++ {
                    v3 = append(v3, ("" + "\n\t"))
                    v4 += (v4 + (-(fi96)))
                    _ = fi97
                }
                _ = fi96
            }
            var v98 uint64 = 760
            fmt.Println(v69, v66)
            _ = v98
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch "foo bar" {
    case "foo bar":
        v99 := GenS8{F0: 101, F1: 9223372036854775807}
        _ = v99
    case "\n\t":
        var v100 int64 = -787
        _ = v100
    default:
        v66 -= uint16(v68)
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v68
    _ = v69
    _ = v74
}
