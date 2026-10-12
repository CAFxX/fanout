package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 uint; F2 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 bool }
func genFn0(p0 uint16) int {
    switch 32768 {
    case -180:
        var v7 float64 = (float64(1e10) * (float64(1e10) + 1.5))
        _ = v7
    default:
        v8 := []int{2147483647}
        _ = v8
    }
    var v9 int8 = -128
    v10 := GenS1{F0: p0, F1: -2147483648, F2: -646}
    _ = p0
    _ = v10
    _ = v9
    return 42
}
func genFn1(p0 int32, p1 int64) int {
    switch "hello" {
    case "":
        var v11 uint32 = 42
        _ = v11
    case "hello":
        switch 157 {
        case -9223372036854775808:
            var v12 int64 = p1
            _ = v12
        case 2147483647:
            var v13 uint8 = 1
            _ = v13
        case -1:
            var v14 int32 = (^(p0))
            _ = v14
        }
    }
    var v15 uint16 = 65535
    var v16 int64 = p1
    var v17 int8 = -12
    p1 = v16
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    return -9223372036854775808
}

func main() {
    var v1 int64 = 519
    v2 := []uint32{683}
    v3 := make(map[uint32]int)
    var v4 bool = false
    var v5 uint = 151
    var v6 uint32 = 689
    var v18 int16 = 485
    if v4 {
        var v19 int32 = 42
        v20 := GenS1{F0: 0, F1: v19, F2: 0}
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int32 = 202
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: 42, F1: -1, F2: -2147483648}
        v6 = 440
        _ = v23
    }
    switch "hello" {
    case "hello":
        switch "hello" {
        case "\n\t":
            v24 := GenS1{F0: 852, F1: 42, F2: 0}
            _ = v24
        default:
            for fi25 := 0; fi25 < 9; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v5 = (v5 % (((v5 ^ (^(v5)))) | 1))
                    delete(v3, v6)
                    if v4 {
                        fmt.Println(v5)
                        v27 := GenS2{F0: 498, F1: uint16(v5), F2: genFn0(0)}
                        _ = v27
                    } else {
                        v4 = true
                        fmt.Println(fi26, v18, fi25)
                        switch "" {
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = (v6 & uint32(v5))
                            }
                            if len(v2) > 0 {
                                v2[0] = (^(v6))
                            }
                        }
                    }
                    if (v4 || (true || (true || v4))) {
                        v28 := GenS3{F0: 0, F1: 2147483647}
                        v1 = 0
                        var v29 uint8 = 1
                        _ = v28
                        _ = v29
                    } else {
                        if (4294967295 > 1) {
                            switch -721 {
                            case 127:
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                            }
                            fmt.Println(v18)
                        }
                        if true {
                            fmt.Println(v18, fi26)
                        } else {
                            var v30 uint16 = 65535
                            v1 = -9223372036854775808
                            if false {
                                delete(v3, 42)
                                v31 := GenS2{F0: 1, F1: (^(v30)), F2: (^(fi26))}
                                _ = v31
                            } else {
                                if v4 {
                                    v32 := v3[65535]
                                    v2 = append(v2, v6)
                                    _ = v32
                                }
                                _ = len(v3)
                            }
                            switch "a" {
                            case "a":
                                var v33 int32 = 627
                                _ = v33
                            case "hello":
                                var v34 uint64 = 845
                                _ = v34
                            default:
                                fmt.Println(v18)
                            }
                            _ = v30
                        }
                    }
                    v5 = (v5 - (v5 + 2))
                    _ = fi26
                }
                delete(v3, v6)
                _ = fi25
            }
        }
    case "\n\t":
        v35 := v3[v6]
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        for fi37 := 0; fi37 < 20; fi37++ {
            if v4 {
                v6 = v6
                v38 := GenS4{F0: v5, F1: (v5 / ((v5) | 1))}
                var v39 float32 = (1.5 + 1.5)
                _ = v38
                _ = v39
            }
            if (v4 && ((("a" + "") + "foo bar") <= ("\n\t" + "a"))) {
                var v40 uint = v5
                fmt.Println(v40, v40)
                _ = v40
            } else {
                var v41 bool = v4
                var v42 uint64 = 251
                _ = v41
                _ = v42
            }
            v5 += v5
            v5 = v5
            _ = fi37
        }
        switch "a" {
        case "foo bar":
            var v43 float64 = 3.14159
            _ = v43
        default:
            v6 = (v6 * (^(v6)))
        }
        _ = fi36
    }
    switch ("hello" + "\n\t") {
    case "":
        v6 += 65535
    default:
        if (-32768 <= -32768) {
            _ = len(v3)
        } else {
            v44 := GenS1{F0: 1, F1: -865, F2: -640}
            v4 = v4
            _ = v44
        }
    }
    v4 = (v4 && (0 > v5))
    if ((v6 / ((485) | 1)) <= (v6 ^ (v6 + v6))) {
        for fi45 := 0; fi45 < 1; fi45++ {
            if (((^(v6)) > (v6 | 2147483648)) || (((!v4) && v4) || v4)) {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v6, v1, v18)
                case "\n\t":
                    v18 = v18
                }
            } else {
                v46 := GenS5{F0: v1, F1: 65535, F2: ((("foo bar" + "foo bar") + "") + "foo bar")}
                _ = v46
            }
            fmt.Println(v1, v18)
            fmt.Println(fi45)
            _ = fi45
        }
        if ((v6 + v6) < 4294967295) {
            switch (("foo bar" + "a") + (("a" + "foo bar") + "\n\t")) {
            case "a":
                v47 := GenS2{F0: 1, F1: 32767, F2: 9223372036854775807}
                _ = v47
            case "foo bar":
                var v48 uint8 = 42
                _ = v48
            case "":
                switch "foo bar" {
                case "hello":
                    v1 += v1
                case "a":
                    if v4 {
                        v49 := GenS3{F0: (^(v5)), F1: (v6 | (v6 << (((^(v5))) & 63)))}
                        _ = v49
                    } else {
                        var v50 int = -1
                        v51 := GenS6{F0: (v50 | v50)}
                        if v4 {
                            var v52 int64 = v1
                            if true {
                                var v53 int32 = -692
                                _ = v53
                            }
                            _ = v52
                        } else {
                            switch v50 {
                            case -703:
                                v5 = v5
                            case -2147483648:
                                var v54 uint64 = 18446744073709551615
                                _ = v54
                            default:
                                v18 += 919
                            }
                            switch (v50 - (v50 << (((v5 << ((v5) & 63))) & 63))) {
                            case -1:
                                v50 -= (v50 ^ genFn0(2))
                            default:
                                v3[(v6 - (v6 >> ((568) & 63)))] = 128
                            }
                        }
                        _ = v50
                        _ = v51
                    }
                case "\n\t":
                    v4 = v4
                default:
                    if v4 {
                        v55 := GenS7{F0: v6}
                        v56 := GenS3{F0: 65535, F1: v6}
                        _ = v55
                        _ = v56
                    }
                }
            default:
                var v57 uint64 = 537
                _ = v57
            }
        }
        for fi58 := 0; fi58 < 13; fi58++ {
            if (0 > v5) {
                if v4 {
                    for fi59 := 0; fi59 < 12; fi59++ {
                        switch "foo bar" {
                        case "hello":
                            if (!true) {
                                v5 = uint(v6)
                                switch ("foo bar" + (("" + "a") + ("hello" + ((((("foo bar" + "") + ("foo bar" + "a")) + "") + "a") + (("foo bar" + (("foo bar" + "a") + (("hello" + "a") + "foo bar"))) + ""))))) {
                                case "":
                                    v5 -= 2
                                case "a":
                                    fmt.Println(v5, fi58)
                                }
                            }
                        case "":
                            v60 := GenS5{F0: -6, F1: 32767, F2: "\n\t"}
                            _ = v60
                        case "\n\t":
                            switch fi58 {
                            case 4294967295:
                                fmt.Println(v4, v6, fi59)
                            case 32768:
                                v61 := GenS2{F0: 18446744073709551615, F1: 1, F2: fi59}
                                _ = v61
                            case -687:
                                var v62 bool = (690 >= 228)
                                _ = v62
                            default:
                                v18 = v18
                            }
                        }
                        v18 -= 42
                        if ((v5 <= (v5 | v5)) && (float32(1.5) != float32((0.1 + (1e10 * 3.14159))))) {
                            switch "hello" {
                            case "hello":
                                var v63 bool = ((127 > 100) && v4)
                                _ = v63
                            case "foo bar":
                                var v64 float64 = (0.1 - (0.0 / 1e10))
                                _ = v64
                            case "a":
                                var v65 float32 = 1.5
                                _ = v65
                            }
                        }
                        _ = fi59
                    }
                    var v66 uint32 = v6
                    _ = v66
                } else {
                    var v67 int8 = 127
                    fmt.Println(v67, v1)
                    switch fi58 {
                    case -1:
                        switch (("hello" + "hello") + "foo bar") {
                        case "hello":
                            v68 := v3[v6]
                            _ = v68
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v5 -= 541
                        }
                    default:
                        _ = len(v2)
                    }
                    _ = v67
                }
            } else {
                for fi69 := 0; fi69 < 20; fi69++ {
                    v70 := GenS1{F0: 65535, F1: 42, F2: int(v6)}
                    delete(v3, 4294967295)
                    switch fi69 {
                    case 127:
                        var v71 uint = (v5 + 99)
                        _ = v71
                    case 9223372036854775807:
                        if (uint8(v5) <= 235) {
                            fmt.Println(v4, fi58)
                            if v4 {
                                var v72 string = "a"
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v72
                            }
                        } else {
                            var v73 uint8 = 42
                            v74 := GenS3{F0: v5, F1: (v6 - v6)}
                            v73 -= v73
                            _ = v73
                            _ = v74
                        }
                    default:
                        if false {
                            switch (("" + "\n\t") + ("\n\t" + "\n\t")) {
                            case "":
                                var v75 uint32 = v6
                                _ = v75
                            default:
                                v6 -= 275
                            }
                            var v76 uint32 = 1
                            _ = v76
                        }
                    }
                    _ = fi69
                    _ = v70
                }
            }
            var v77 float32 = 1e10
            delete(v3, v6)
            _ = fi58
            _ = v77
        }
    }
    v78 := GenS3{F0: 255, F1: 42}
    for fi79 := 0; fi79 < 15; fi79++ {
        switch "\n\t" {
        case "\n\t":
            var v80 int = 2147483647
            _ = v80
        case "a":
            if v4 {
                v81 := GenS8{F0: -1}
                v82 := GenS2{F0: 18446744073709551615, F1: 572, F2: fi79}
                var v83 int = fi79
                _ = v81
                _ = v82
                _ = v83
            }
        default:
            switch "" {
            case "foo bar":
                for fi84 := 0; fi84 < 10; fi84++ {
                    v85 := GenS9{F0: "hello"}
                    if (-461 == 51) {
                        fmt.Println(v5)
                        v86 := v3[(v6 & (v6 & (v6 | 0)))]
                        _ = v86
                    } else {
                        v4 = true
                        v6 = (v6 ^ 65535)
                        var v87 float32 = 1e10
                        _ = v87
                    }
                    _ = fi84
                    _ = v85
                }
            case "hello":
                _ = len(v2)
            }
        }
        _ = fi79
    }
    fmt.Println(v4, v4, v6)
    if (2147483647 != int(v18)) {
        switch int(v5) {
        case 2147483647:
            v1 -= v1
        case 32767:
            var v88 float32 = 1.5
            _ = v88
        case -188:
            v1 = (v1 - (v1 - (v1 ^ v1)))
        default:
            fmt.Println(v4, v6, v6)
        }
        for fi89 := 0; fi89 < 15; fi89++ {
            switch (fi89 >> ((v5) & 63)) {
            case -2147483648:
                v6 += 42
            }
            fmt.Println(v18, v1)
            _ = fi89
        }
        for fi90 := 0; fi90 < 7; fi90++ {
            v91 := GenS10{F0: 1, F1: (v5 - (v5 % ((v5) | 1))), F2: 255}
            v92 := GenS11{F0: 350}
            _ = fi90
            _ = v91
            _ = v92
        }
    }
    switch 4294967295 {
    case 127:
        v18 = int16(v6)
    case -1:
        v93 := GenS12{F0: ("hello" > "")}
        _ = v93
    case 32767:
        v1 = v1
    default:
        v6 = 2
    }
    if (0.1 < (((0.0 + -2.25) + -2.25) * 1e10)) {
        switch ("hello" + "\n\t") {
        case "foo bar":
            v5 += (v5 % ((255) | 1))
        default:
            v1 = 42
        }
    } else {
        for fi94 := 0; fi94 < 4; fi94++ {
            var v95 uint16 = 80
            _ = fi94
            _ = v95
        }
        v4 = false
    }
    v18 += (v18 ^ (v18 * v18))
    if ((0 != 255) || v4) {
        switch ("hello" + "\n\t") {
        case "a":
            v96 := GenS1{F0: 42, F1: 1, F2: 42}
            _ = v96
        default:
            fmt.Println(v4)
        }
        v1 += (v1 | v1)
    }
    var v97 int8 = 26
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v97
}
