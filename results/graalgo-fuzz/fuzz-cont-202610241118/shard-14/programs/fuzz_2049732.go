package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 int32; F1 uint16; F2 int64 }

func main() {
    var v1 float64 = 1e10
    var v2 float64 = float64((-2.25 / 3.14159))
    v3 := []int{-419}
    v4 := make(map[int]string)
    switch 4294967295 {
    case -2147483648:
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 uint32 = 979
            var v7 int16 = 175
            switch (fi5 - (fi5 << ((671) & 63))) {
            case -9223372036854775808:
                v6 += v6
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    case 32767:
        var v8 int64 = 0
        _ = v8
    }
    var v9 uint8 = 103
    for fi10 := 0; fi10 < 18; fi10++ {
        delete(v4, (fi10 * (fi10 << ((0) & 63))))
        v11 := GenS1{F0: 2, F1: fi10}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if true {
            v1 += 3.14159
        } else {
            v1 += v1
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int = fi13
                fmt.Println(fi13)
                switch v14 {
                case 32767:
                    if false {
                        switch "foo bar" {
                        case "":
                            var v15 int8 = 42
                            v1 += (0.1 / (-2.25 + float64(((((0.0 * ((((-2.25 * (1e10 * (float32((-2.25 + 3.14159)) + (1e10 + float32(v1))))) + float32(((3.14159 + (1e10 + v1)) / 1e10))) - (-2.25 - 1e10)) - ((1e10 * (0.0 * 0.1)) * 3.14159))) + (((0.1 / -2.25) * 1e10) / (((3.14159 / ((-2.25 + 1e10) * float32(float64(float32(v2))))) + 0.0) * (1.5 + 3.14159)))) - (0.1 - (1.5 / -2.25))) - -2.25))))
                            _ = v15
                        default:
                            var v16 int64 = 42
                            var v17 uint8 = v9
                            _ = v16
                            _ = v17
                        }
                    }
                case 32768:
                    fmt.Println(fi13, fi12)
                case 0:
                    v18 := GenS1{F0: 255, F1: fi12}
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (620 == 944) {
                            if (false || (!false)) {
                                v4[fi13] = "hello"
                                if (-911 >= 0) {
                                    v19 := GenS1{F0: 65535, F1: (v14 / ((v14) | 1))}
                                    var v20 int16 = 1000
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 string = "a"
                                    v4[v14] = v21
                                    _ = v21
                                }
                            } else {
                                v14 += v14
                                v14 = (v14 ^ (-(fi13)))
                            }
                        }
                    }
                    _ = v18
                }
                fmt.Println(fi12, v1)
                switch (v14 ^ (fi13 & fi12)) {
                case -9223372036854775808:
                    fmt.Println(v2, v1)
                default:
                    if (uint(v9) <= 255) {
                        var v22 int8 = 68
                        v3 = append(v3, fi13)
                        _ = v22
                    }
                }
                _ = fi13
                _ = v14
            }
        }
        v1 += float64((-2.25 - (0.1 / 1.5)))
        _ = fi12
    }
    if true {
        switch 32768 {
        case -9223372036854775808:
            v23 := GenS2{F0: 242}
            _ = v23
        case 963:
            for fi24 := 0; fi24 < 12; fi24++ {
                for fi25 := 0; fi25 < 20; fi25++ {
                    fmt.Println(v1)
                    if false {
                        var v26 uint16 = 626
                        v27 := GenS2{F0: 32767}
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            var v28 int64 = -9223372036854775808
                            _ = v28
                        case "a":
                            switch fi25 {
                            case -2147483648:
                                v29 := GenS1{F0: (v26 | (^(v26))), F1: (fi24 & (fi24 * fi24))}
                                _ = v29
                            case -866:
                                switch -1 {
                                case 1:
                                    _ = len(v3)
                                case 197:
                                    if (!(-1 > 42)) {
                                        if (false || false) {
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v9)
                                                var v30 uint16 = v26
                                                _ = v30
                                            case "":
                                                v31 := GenS1{F0: (v26 >> ((18446744073709551615) & 63)), F1: 2147483647}
                                                _ = v31
                                            default:
                                                fmt.Println(fi24)
                                            }
                                        } else {
                                            var v32 int16 = 0
                                            var v33 bool = (-937 != -330)
                                            _ = v32
                                            _ = v33
                                        }
                                        v2 += 0.0
                                        v34 := GenS1{F0: (v26 >> ((252) & 63)), F1: 107}
                                        _ = v34
                                    } else {
                                        if true {
                                            switch -2147483648 {
                                            case 2147483647:
                                                v35 := GenS3{F0: "hello", F1: 41, F2: 0}
                                                var v36 bool = ((0.1 - float64(3.14159)) != v1)
                                                var v37 uint64 = 479
                                                v38 := v4[(fi25 >> ((18446744073709551615) & 63))]
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            case -9223372036854775808:
                                                var v39 uint = 255
                                                _ = v39
                                            }
                                            var v40 int = (fi25 & fi24)
                                            v9 = v9
                                            _ = v40
                                        }
                                        var v41 uint32 = 0
                                        if (int32(v41) == -67) {
                                            if false {
                                                var v42 float64 = v2
                                                v43 := GenS4{F0: 42, F1: v9, F2: (^(v41))}
                                                switch ("\n\t" + "hello") {
                                                case "":
                                                    fmt.Println(v2, v42, fi25)
                                                case "a":
                                                    v41 = (v41 / ((42) | 1))
                                                case "\n\t":
                                                    if false {
                                                        switch (fi24 - fi24) {
                                                        case 493:
                                                            if false {
                                                                var v44 bool = (true && true)
                                                                var v45 int64 = 42
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                fmt.Println(v2, fi24, v2)
                                                            }
                                                        case 32768:
                                                            v46 := GenS1{F0: v26, F1: -9223372036854775808}
                                                            _ = v46
                                                        case -9223372036854775808:
                                                            v47 := GenS1{F0: (^(v26)), F1: (fi24 & 9223372036854775807)}
                                                            switch (^(fi25)) {
                                                            case -612:
                                                                delete(v4, int(v9))
                                                            }
                                                            _ = v47
                                                        }
                                                    }
                                                    switch (fi25 ^ fi25) {
                                                    case 127:
                                                        if false {
                                                            v48 := GenS5{F0: 916}
                                                            _ = v48
                                                        } else {
                                                            if ((v41 | v41) >= (v41 + 2147483647)) {
                                                                var v49 float64 = ((3.14159 - v1) / v42)
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi24
                                                                }
                                                                v9 -= 0
                                                                _ = v49
                                                            } else {
                                                                v50 := GenS5{F0: 65535}
                                                                v41 = v41
                                                                _ = v50
                                                            }
                                                            v26 = v26
                                                            v2 = 3.14159
                                                        }
                                                    default:
                                                        if false {
                                                            v41 += v41
                                                            v51 := GenS1{F0: v26, F1: (fi25 << ((4294967295) & 63))}
                                                            switch "hello" {
                                                            case "":
                                                                v1 -= v2
                                                            case "hello":
                                                                var v52 int32 = 927
                                                                _ = v52
                                                            case "\n\t":
                                                                v53 := v4[fi24]
                                                                _ = v53
                                                            }
                                                            v42 = v1
                                                            var v54 bool = (!true)
                                                            v3 = append(v3, fi24)
                                                            _ = v51
                                                            _ = v54
                                                        } else {
                                                            var v55 uint16 = v26
                                                            delete(v4, fi24)
                                                            v1 += (0.1 + v1)
                                                            _ = v55
                                                        }
                                                    }
                                                default:
                                                    v1 -= (v2 - 1.5)
                                                }
                                                if false {
                                                    v41 = v41
                                                } else {
                                                    if true {
                                                        v2 -= (float64(1e10) - (1e10 - ((v2 - float64(((1.5 / (0.1 * 0.1)) / (3.14159 - 0.1)))) - 3.14159)))
                                                        var v56 int = 128
                                                        if true {
                                                            fmt.Println(fi25)
                                                            fmt.Println(v41)
                                                            v9 = 91
                                                            var v57 int8 = -1
                                                            fmt.Println(v42, v57, v9)
                                                            _ = v57
                                                        }
                                                        if false {
                                                            v56 = (-(fi25))
                                                            _ = len(v4)
                                                        }
                                                        _ = v56
                                                    }
                                                    v41 -= v41
                                                    v42 = (v1 / -2.25)
                                                    var v58 int16 = 1
                                                    _ = v58
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                        } else {
                                            v59 := GenS1{F0: v26, F1: (-(fi24))}
                                            fmt.Println(v26, v1)
                                            switch (^(fi25)) {
                                            case 32768:
                                                v60 := GenS2{F0: 1}
                                                _ = v60
                                            }
                                            _ = v59
                                        }
                                        if (!((7 > 538) || ((true || (!(!true))) || false))) {
                                            _ = len(v3)
                                            v1 -= -2.25
                                        } else {
                                            _ = len(v4)
                                            v61 := GenS6{F0: v26}
                                            _ = v61
                                        }
                                        _ = v41
                                    }
                                case 0:
                                    v2 = 0.1
                                }
                                fmt.Println(v2)
                            }
                        case "":
                            if (1 > -32768) {
                                switch 267 {
                                case 32767:
                                    fmt.Println(v2)
                                    var v62 int32 = 0
                                    _ = v62
                                case -196:
                                    v9 = 1
                                default:
                                    v2 += 0.1
                                }
                                if true {
                                    v1 += v2
                                }
                                var v63 uint32 = 65535
                                _ = v63
                            }
                        default:
                            v64 := GenS5{F0: 9}
                            _ = v64
                        }
                        _ = v26
                        _ = v27
                    }
                    switch "" {
                    case "":
                        v3 = append(v3, -2147483648)
                    case "foo bar":
                        if true {
                            if len(v3) > 0 {
                                v3[0] = fi25
                            }
                            _ = len(v4)
                            var v65 bool = (32767 <= 255)
                            _ = v65
                        } else {
                            if (!false) {
                                v1 += (v2 - (float64(1e10) / v2))
                                var v66 int16 = 32767
                                if (469 <= 18446744073709551615) {
                                    v67 := GenS6{F0: 0}
                                    v68 := GenS7{F0: 1}
                                    var v69 int8 = -128
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                }
                                _ = v66
                            }
                            switch "a" {
                            case "a":
                                var v70 uint = 0
                                _ = v70
                            case "\n\t":
                                var v71 int32 = 1
                                _ = v71
                            case "hello":
                                delete(v4, (^(fi25)))
                            default:
                                var v72 int32 = 42
                                _ = v72
                            }
                            v2 += ((-2.25 - (float64(-2.25) - v2)) / ((0.0 * (v2 + (1.5 / v2))) - (v2 + (((v1 - ((0.1 / v2) / v1)) / (3.14159 * (float64(1e10) + v1))) - v2))))
                        }
                    case "a":
                        var v73 uint = 0
                        _ = v73
                    default:
                        v2 = 0.1
                    }
                    _ = fi25
                }
                if (-1 >= 0) {
                    v9 = v9
                    for fi74 := 0; fi74 < 7; fi74++ {
                        _ = len(v4)
                        v1 += v1
                        _ = fi74
                    }
                    _ = len(v4)
                } else {
                    if ((true && (true || ((false && (true && false)) && true))) && (1.5 != float64(((-2.25 - (1e10 + 0.1)) * 1.5)))) {
                        for fi75 := 0; fi75 < 20; fi75++ {
                            switch (fi24 - (fi24 ^ 42)) {
                            case -762:
                                v1 -= (v1 / ((((v2 / 1e10) * v2) / (3.14159 * v1)) - float64(0.0)))
                            case 4294967295:
                                _ = len(v4)
                            default:
                                v76 := GenS8{F0: (9223372036854775807 < (fi75 ^ (fi75 & fi75))), F1: 1}
                                _ = v76
                            }
                            _ = fi75
                        }
                        for fi77 := 0; fi77 < 5; fi77++ {
                            fmt.Println(fi77, v2, fi24)
                            _ = fi77
                        }
                        fmt.Println(v2, v1)
                    } else {
                        v2 -= 1e10
                    }
                    v1 = v2
                    switch "foo bar" {
                    case "foo bar":
                        var v78 uint = 626
                        _ = v78
                    }
                }
                fmt.Println(v9, fi24, v9)
                _ = fi24
            }
        case 470:
            var v79 uint8 = v9
            _ = v79
        }
        _ = len(v3)
        v80 := GenS8{F0: ((v2 * 0.1) >= ((v2 * v1) * v1)), F1: 4294967295}
        _ = v80
    }
    for fi81 := 0; fi81 < 2; fi81++ {
        v1 = (v2 * float64(-2.25))
        v2 = float64(1.5)
        _ = fi81
    }
    v2 += (0.1 * v2)
    if true {
        v2 = (v2 + v2)
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi82 := 0; fi82 < 17; fi82++ {
                var v83 int = fi82
                _ = fi82
                _ = v83
            }
            v84 := GenS9{F0: 449, F1: "hello"}
            _ = v84
        }
    } else {
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        if false {
            v2 += (float64(1e10) - (float64(1e10) / v1))
        }
        for fi85 := 0; fi85 < 19; fi85++ {
            for fi86 := 0; fi86 < 10; fi86++ {
                v4[(fi86 ^ fi86)] = "a"
                if len(v3) > 0 {
                    v3[0] = 32767
                }
                _ = fi86
            }
            if true {
                v1 -= ((float64((-2.25 - 0.0)) * v2) + v1)
            } else {
                for fi87 := 0; fi87 < 15; fi87++ {
                    v1 = 1.5
                    v88 := GenS10{F0: -1, F1: 730, F2: 42}
                    switch (fi85 >> ((1) & 63)) {
                    case -369:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -1:
                        v2 += (v1 / v1)
                    case 0:
                        fmt.Println(v9, fi85, v1)
                    }
                    _ = fi87
                    _ = v88
                }
                v1 = v2
            }
            _ = fi85
        }
    }
    v9 += (v9 >> ((270) & 63))
    v1 += (0.1 - float64((1.5 + -2.25)))
    v9 = (v9 % ((v9) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi89 := 0; fi89 < 20; fi89++ {
        switch (("foo bar" + "foo bar") + ((("hello" + ("hello" + "")) + ((("foo bar" + "\n\t") + "hello") + "hello")) + "a")) {
        case "":
            fmt.Println(fi89)
        case "foo bar":
            var v90 int64 = 774
            _ = v90
        }
        var v91 uint = uint(v9)
        _ = fi89
        _ = v91
    }
    v2 += (v2 * v1)
    for fi92 := 0; fi92 < 4; fi92++ {
        v1 = 0.0
        var v93 int32 = 2147483647
        _ = fi92
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
