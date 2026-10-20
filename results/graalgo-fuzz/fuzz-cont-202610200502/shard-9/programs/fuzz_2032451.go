package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 int64; F2 uint16 }
type GenS12 struct { F0 int; F1 uint }

func main() {
    var v1 string = "a"
    var v2 bool = (!(true && (0 >= 255)))
    var v3 uint = 500
    var v4 uint = (v3 ^ (v3 + v3))
    var v5 uint32 = 65535
    var v6 bool = (-2.25 < float32((1e10 / 1e10)))
    v7 := GenS1{F0: 183, F1: v1, F2: 0}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint8 = 136
        _ = fi8
        _ = v9
    }
    fmt.Println(v6, v6, v5)
    switch v1 {
    case "a":
        var v10 bool = (v6 && (v2 || true))
        _ = v10
    case "\n\t":
        fmt.Println(v1, v5, v6)
    default:
        fmt.Println(v3, v4)
    }
    if (2 > 183) {
        var v11 int = -1
        var v12 int64 = -9223372036854775808
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: v1, F2: 1000}
    switch (v1 + v1) {
    case "hello":
        v2 = v2
    case "foo bar":
        switch 241 {
        case 9223372036854775807:
            if false {
                var v14 uint16 = 423
                _ = v14
            }
        }
    }
    v2 = (-723 >= -2147483648)
    v5 = v5
    switch -1 {
    case -1:
        if false {
            if v2 {
                v4 = v4
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 uint64 = 1
                    v17 := GenS2{F0: (((("foo bar" + "hello") + v1) + "") + ((v1 + "") + ("a" + (((v1 + v1) + "hello") + "")))), F1: 240, F2: 0}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            var v18 string = (((v1 + ("\n\t" + v1)) + ("\n\t" + v1)) + ("foo bar" + v1))
            v2 = false
            _ = v18
        } else {
            switch v1 {
            case "":
                var v19 int8 = 100
                _ = v19
            case "foo bar":
                v2 = (4294967295 != 1)
            case "a":
                for fi20 := 0; fi20 < 5; fi20++ {
                    switch (fi20 / (((fi20 + 127)) | 1)) {
                    case -637:
                        if (!(65535 != 870)) {
                            var v21 bool = (-1 != 0)
                            fmt.Println(v2, v2)
                            _ = v21
                        }
                    }
                    var v22 uint64 = uint64(v5)
                    _ = fi20
                    _ = v22
                }
            default:
                fmt.Println(v6)
            }
            fmt.Println(v6)
        }
    }
    v2 = false
    v23 := []int{}
    v2 = false
    switch 32768 {
    case 785:
        v2 = true
    }
    v6 = (0 != 8)
    if (v1 > "") {
        var v24 int16 = 0
        if false {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            if false {
                var v25 int64 = 1
                var v26 bool = (((float32(0.0) / 1.5) - (0.1 - 1.5)) < -2.25)
                v24 = v24
                _ = v25
                _ = v26
            } else {
                switch 32768 {
                case 566:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        var v28 int = 4294967295
                        _ = fi27
                        _ = v28
                    }
                case 1:
                    if (v2 && v6) {
                        v29 := GenS3{F0: v3, F1: 1}
                        for fi30 := 0; fi30 < 2; fi30++ {
                            v31 := GenS4{F0: v6}
                            var v32 float64 = (1.5 - float64((0.1 + float32(3.14159))))
                            v6 = ((v3 >= v4) || ((v24 & v24) <= v24))
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        switch -466 {
                        case 344:
                            switch "a" {
                            case "a":
                                var v33 int8 = 100
                                _ = v33
                            case "foo bar":
                                v23 = append(v23, 2147483647)
                            case "":
                                var v34 uint16 = 109
                                _ = v34
                            default:
                                var v35 uint32 = (v5 & (v5 / ((v5) | 1)))
                                _ = v35
                            }
                        case 517:
                            fmt.Println(v1, v5)
                        }
                        _ = v29
                    }
                case 127:
                    for fi36 := 0; fi36 < 1; fi36++ {
                        if v2 {
                            v5 -= (v5 * (v5 & 65535))
                        }
                        _ = fi36
                    }
                }
                var v37 string = ""
                if v2 {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    _ = len(v23)
                }
                _ = v37
            }
            v1 = ("" + "")
        } else {
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 float32 = 1e10
                switch 183 {
                case 42:
                    fmt.Println(v3, v5)
                case 9223372036854775807:
                    if (3.14159 > 0.1) {
                        switch fi38 {
                        case 32767:
                            var v40 uint = (v4 / ((42) | 1))
                            _ = v40
                        }
                        if (0 <= 255) {
                            if v6 {
                                for fi41 := 0; fi41 < 18; fi41++ {
                                    v42 := GenS5{F0: 33}
                                    if true {
                                        if false {
                                            v43 := GenS2{F0: ((v1 + ((v1 + "\n\t") + ((v1 + ("foo bar" + (v1 + v1))) + (v1 + v1)))) + v1), F1: 2, F2: 80}
                                            if v2 {
                                                v1 = v1
                                                var v44 uint64 = 18446744073709551615
                                                var v45 bool = v2
                                                _ = v44
                                                _ = v45
                                            } else {
                                                var v46 int = (^(fi41))
                                                var v47 float64 = 1e10
                                                fmt.Println(fi41, v6, v47)
                                                _ = v46
                                                _ = v47
                                            }
                                            _ = len(v23)
                                            _ = v43
                                        } else {
                                            v48 := GenS1{F0: 2, F1: v1, F2: -32768}
                                            switch (v1 + v1) {
                                            case "hello":
                                                v49 := GenS6{F0: 4294967295, F1: -66, F2: v5}
                                                _ = v49
                                            default:
                                                var v50 int64 = 42
                                                _ = v50
                                            }
                                            _ = v48
                                        }
                                        if v2 {
                                            _ = len(v23)
                                            var v51 uint16 = 2
                                            var v52 int = 32767
                                            v23 = append(v23, fi41)
                                            _ = v51
                                            _ = v52
                                        } else {
                                            fmt.Println(v4)
                                            v3 = 853
                                        }
                                        v6 = true
                                    }
                                    _ = fi41
                                    _ = v42
                                }
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            }
                            if v2 {
                                fmt.Println(v4, v5, v3)
                                if true {
                                    if len(v23) > 0 {
                                        v23[0] = fi38
                                    }
                                    v53 := GenS7{F0: 128}
                                    v1 = "\n\t"
                                    _ = v53
                                } else {
                                    v23 = append(v23, 4294967295)
                                    for fi54 := 0; fi54 < 13; fi54++ {
                                        v2 = (v39 <= (v39 + (-2.25 - ((v39 / -2.25) - 3.14159))))
                                        if v2 {
                                            switch int(v3) {
                                            case 52:
                                                v39 = ((float32((3.14159 + 1.5)) * (v39 * (v39 * v39))) / v39)
                                            case 32768:
                                                var v55 int64 = -1
                                                _ = v55
                                            default:
                                                var v56 int64 = 173
                                                _ = v56
                                            }
                                        } else {
                                            if v6 {
                                                v39 = (v39 - (1e10 - v39))
                                                v57 := GenS8{F0: v2, F1: v5}
                                                switch -2147483648 {
                                                case -9223372036854775808:
                                                    switch (fi38 << (((v4 * 2)) & 63)) {
                                                    case 32767:
                                                        fmt.Println(fi38)
                                                        v58 := GenS9{F0: 762}
                                                        switch v1 {
                                                        case "hello":
                                                            if (!v6) {
                                                                var v59 bool = (!(((v39 / v39) > 0.0) && v6))
                                                                fmt.Println(v3)
                                                                _ = v59
                                                            } else {
                                                                if (-2147483648 != -799) {
                                                                    if len(v23) >= 2 {
                                                                        v23 = v23[1:len(v23)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v58
                                                    case 2147483647:
                                                        switch fi54 {
                                                        case -273:
                                                            var v60 uint16 = 0
                                                            _ = v60
                                                        case 803:
                                                            fmt.Println(fi54, v5)
                                                        case 32768:
                                                            if (fi54 <= (fi54 & fi38)) {
                                                                v23 = append(v23, 42)
                                                                var v61 string = (v1 + v1)
                                                                var v62 int32 = 602
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                v63 := GenS10{F0: (v24 / ((v24) | 1))}
                                                                var v64 int16 = v24
                                                                var v65 uint16 = 637
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                        }
                                                    }
                                                case 263:
                                                    var v66 int32 = 415
                                                    _ = v66
                                                case 0:
                                                    switch fi38 {
                                                    case 1:
                                                        v4 += v4
                                                    case -2147483648:
                                                        var v67 int64 = 0
                                                        _ = v67
                                                    case 4294967295:
                                                        if false {
                                                            v1 = ("\n\t" + "foo bar")
                                                            v5 += uint32(v3)
                                                        }
                                                    }
                                                }
                                                _ = v57
                                            }
                                        }
                                        var v68 float64 = 1e10
                                        _ = fi54
                                        _ = v68
                                    }
                                }
                            } else {
                                var v69 int32 = int32(v24)
                                v70 := GenS11{F0: -362, F1: -585, F2: 567}
                                fmt.Println(v39, v2)
                                _ = v69
                                _ = v70
                            }
                        } else {
                            v23 = append(v23, (fi38 + fi38))
                            fmt.Println(v1, v39, v1)
                            if len(v23) > 0 {
                                v23[0] = fi38
                            }
                        }
                        v71 := GenS5{F0: 0}
                        _ = v71
                    } else {
                        switch "" {
                        case "":
                            if v6 {
                                v24 -= 32767
                                for fi72 := 0; fi72 < 10; fi72++ {
                                    fmt.Println(v24, v2, fi72)
                                    v23 = append(v23, int(v24))
                                    var v73 uint8 = 1
                                    _ = fi72
                                    _ = v73
                                }
                            } else {
                                var v74 float64 = 0.1
                                if ("foo bar" >= ((v1 + ((v1 + ("foo bar" + v1)) + (v1 + ("" + (v1 + v1))))) + v1)) {
                                    for fi75 := 0; fi75 < 7; fi75++ {
                                        if len(v23) > 0 {
                                            v23[0] = int(v24)
                                        }
                                        _ = fi75
                                    }
                                    v74 = v74
                                } else {
                                    if v2 {
                                        switch (-(fi38)) {
                                        case 786:
                                            v3 += v3
                                        case -488:
                                            for fi76 := 0; fi76 < 14; fi76++ {
                                                fmt.Println(v6, v1)
                                                v77 := GenS12{F0: fi38, F1: v4}
                                                if len(v23) > 0 {
                                                    v23[0] = (^(fi38))
                                                }
                                                _ = fi76
                                                _ = v77
                                            }
                                        default:
                                            for fi78 := 0; fi78 < 19; fi78++ {
                                                var v79 int16 = (v24 / (((v24 | (v24 | v24))) | 1))
                                                if (!(489 <= 65535)) {
                                                    var v80 uint16 = 32767
                                                    _ = v80
                                                } else {
                                                    v5 = 4294967295
                                                }
                                                v79 = 928
                                                _ = fi78
                                                _ = v79
                                            }
                                        }
                                        for fi81 := 0; fi81 < 12; fi81++ {
                                            fmt.Println(v1)
                                            _ = fi81
                                        }
                                    }
                                    v82 := GenS3{F0: v4, F1: 147}
                                    var v83 int = (fi38 ^ 284)
                                    _ = v82
                                    _ = v83
                                }
                                fmt.Println(v5)
                                _ = v74
                            }
                        case "hello":
                            var v84 uint = v4
                            _ = v84
                        }
                        if v2 {
                            var v85 int = fi38
                            for fi86 := 0; fi86 < 7; fi86++ {
                                var v87 int = (^(fi86))
                                if (!(!v6)) {
                                    switch (v1 + (v1 + (("a" + v1) + "\n\t"))) {
                                    case "":
                                        switch (-(v87)) {
                                        case 32767:
                                            if v6 {
                                                switch "\n\t" {
                                                case "a":
                                                    if ((!false) && v2) {
                                                        switch v1 {
                                                        case "\n\t":
                                                            v88 := GenS6{F0: 421, F1: 42, F2: v5}
                                                            _ = v88
                                                        default:
                                                            v89 := GenS1{F0: 121, F1: v1, F2: -1}
                                                            _ = v89
                                                        }
                                                        v3 += 42
                                                    } else {
                                                        v90 := GenS8{F0: ((v5 & v5) > 2147483647), F1: (^(v5))}
                                                        v23 = append(v23, fi86)
                                                        var v91 int32 = 42
                                                        var v92 int = (-(v85))
                                                        _ = v90
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                default:
                                                    if true {
                                                        switch (fi86 << (((^(v3))) & 63)) {
                                                        case 4294967295:
                                                            var v93 uint64 = 42
                                                            _ = v93
                                                        }
                                                        if len(v23) > 0 {
                                                            v23[0] = 330
                                                        }
                                                        switch ((((("foo bar" + v1) + v1) + (v1 + v1)) + ("" + (v1 + ((("hello" + "foo bar") + ((v1 + v1) + (v1 + ((v1 + v1) + "hello")))) + "foo bar")))) + v1) {
                                                        case "\n\t":
                                                            v94 := GenS7{F0: 0}
                                                            _ = v94
                                                        default:
                                                            v1 = ((v1 + "\n\t") + v1)
                                                        }
                                                    } else {
                                                        var v95 uint16 = 255
                                                        _ = v95
                                                    }
                                                }
                                                var v96 int16 = v24
                                                fmt.Println(v39)
                                                _ = v96
                                            }
                                        case 127:
                                            v85 += fi38
                                        default:
                                            _ = len(v23)
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            if v6 {
                                                var v97 uint64 = 0
                                                switch "\n\t" {
                                                case "a":
                                                    v5 += v5
                                                default:
                                                    var v98 int8 = int8(v5)
                                                    _ = v98
                                                }
                                                _ = v97
                                            }
                                        }
                                    default:
                                        var v99 string = (v1 + (v1 + v1))
                                        _ = v99
                                    }
                                }
                                _ = fi86
                                _ = v87
                            }
                            _ = v85
                        }
                    }
                }
                _ = fi38
                _ = v39
            }
        }
        _ = v24
    } else {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v100 bool = false
        v4 -= (v4 / (((v3 >> (((v4 / ((v4) | 1))) & 63))) | 1))
        _ = v100
    }
    var v101 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v101
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
