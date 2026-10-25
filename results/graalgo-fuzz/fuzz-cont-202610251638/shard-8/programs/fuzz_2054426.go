package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 uint8; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 int8; F1 uint; F2 int64 }
type GenS10 struct { F0 int32; F1 int32 }
func genFn0(p0 int16, p1 int, p2 bool) bool {
    var v5 uint8 = 255
    var v6 int64 = 9223372036854775807
    switch ((("a" + "") + "foo bar") + ("foo bar" + "")) {
    case "hello":
        var v7 uint = 223
        _ = v7
    case "\n\t":
        switch (p1 << ((42) & 63)) {
        case 32767:
            if false {
                v8 := GenS1{F0: (p2 || (-119 <= -10)), F1: 42}
                fmt.Println(v6)
                for fi9 := 0; fi9 < 18; fi9++ {
                    fmt.Println(fi9, p0)
                    var v10 int16 = p0
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                var v11 bool = ((0.1 + 0.1) >= 0.0)
                var v12 float32 = (0.1 - -2.25)
                var v13 int64 = -1
                _ = v11
                _ = v12
                _ = v13
            }
        case -528:
            fmt.Println(p2, p1)
        case 42:
            fmt.Println(v5, p2)
        default:
            if p2 {
                if p2 {
                    var v14 int64 = (v6 - -9223372036854775808)
                    var v15 int32 = 346
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v17 := []int{(p1 % ((-73) | 1)), (-(p1))}
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                v17 = append(v17, fi16)
                            }
                            v18 := GenS1{F0: false, F1: 100}
                            _ = v18
                        }
                        v19 := GenS2{F0: p1, F1: (((v5 << ((42) & 63)) <= (v5 / ((1) | 1))) && (((1 < 0) && false) || (0.1 <= float64((((0.1 + ((float32((1.5 / 1e10)) * (0.0 + (((-2.25 / 3.14159) + 0.0) * (((1e10 + 3.14159) * (0.0 * (0.0 * 0.1))) * ((1.5 * 0.1) * (float32(((float64(0.0) * 3.14159) * 3.14159)) - ((float32(1e10) + 0.0) * 0.0))))))) * -2.25)) / 0.1) - 0.0))))), F2: 65535}
                        _ = fi16
                        _ = v17
                        _ = v19
                    }
                    _ = v14
                    _ = v15
                }
            } else {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := []int{(p1 / ((fi20) | 1)), -890, fi20}
                    p1 = 2147483647
                    v5 += v5
                    _ = fi20
                    _ = v21
                }
            }
        }
    case "foo bar":
        var v22 uint64 = 2
        _ = v22
    }
    switch ("hello" + "hello") {
    case "":
        v23 := []int{p1, (p1 | p1)}
        _ = v23
    default:
        if true {
            v6 += (v6 & (v6 & (v6 << ((146) & 63))))
            var v24 bool = ((^(p0)) != p0)
            _ = v24
        }
    }
    switch (^(p1)) {
    case 32767:
        v25 := GenS3{F0: v6}
        _ = v25
    case 128:
        p1 = int(p0)
    default:
        var v26 int32 = 1
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (1 > 42)
}

func main() {
    var v1 uint16 = 1
    var v2 int16 = 42
    var v3 int32 = -8
    var v4 int = -9223372036854775808
    var v27 bool = (("foo bar" + ("hello" + "foo bar")) >= ("hello" + ""))
    switch v4 {
    case 2147483647:
        v28 := []int{v4}
        _ = v28
    }
    if (!(false || (((float64(-2.25) / 1.0) != (1.5 * (0.1 + -2.25))) && (false || (42 <= 128))))) {
        v29 := []int{}
        switch v4 {
        case 294:
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v4)
            case "hello":
                switch (v4 / ((4294967295) | 1)) {
                case -1:
                    var v30 string = "a"
                    _ = v30
                case -9223372036854775808:
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        switch v4 {
                        case 192:
                            var v31 string = ""
                            _ = v31
                        default:
                            for fi32 := 0; fi32 < 19; fi32++ {
                                switch fi32 {
                                case 32767:
                                    v2 += v2
                                case -9223372036854775808:
                                    var v33 int64 = 763
                                    _ = v33
                                case 2147483647:
                                    v2 = v2
                                }
                                _ = fi32
                            }
                        }
                    case "":
                        if len(v29) > 0 {
                            v29[0] = v4
                        }
                    case "foo bar":
                        v27 = v27
                    }
                default:
                    if (v27 && v27) {
                        switch 9223372036854775807 {
                        case 127:
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                        case 32768:
                            v3 += 2147483647
                        case 42:
                            for fi34 := 0; fi34 < 9; fi34++ {
                                v35 := GenS4{F0: (v4 & (-(v4))), F1: 9223372036854775807}
                                _ = fi34
                                _ = v35
                            }
                        default:
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v29)
                        case "":
                            fmt.Println(v3, v27, v1)
                        case "foo bar":
                            var v36 float64 = (-2.25 + (0.1 * float64(3.14159)))
                            _ = v36
                        default:
                            if true {
                                var v37 int8 = 115
                                _ = v37
                            } else {
                                v38 := GenS4{F0: v4, F1: 825}
                                if (v27 && true) {
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    v2 = int16(v4)
                                    var v39 uint16 = 1
                                    _ = v39
                                } else {
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    if v27 {
                                        for fi40 := 0; fi40 < 18; fi40++ {
                                            var v41 float64 = 1.5
                                            switch v4 {
                                            case 40:
                                                var v42 int8 = -1
                                                _ = v42
                                            case 127:
                                                var v43 uint64 = 1
                                                _ = v43
                                            case 32768:
                                                var v44 uint64 = 889
                                                _ = v44
                                            default:
                                                var v45 uint = 255
                                                _ = v45
                                            }
                                            if (984 == 4294967295) {
                                                v46 := GenS4{F0: fi40, F1: -1}
                                                _ = v46
                                            } else {
                                                var v47 int = 32767
                                                for fi48 := 0; fi48 < 10; fi48++ {
                                                    fmt.Println(v1)
                                                    _ = fi48
                                                }
                                                v49 := GenS4{F0: -2147483648, F1: -500}
                                                _ = v47
                                                _ = v49
                                            }
                                            _ = fi40
                                            _ = v41
                                        }
                                    }
                                }
                                for fi50 := 0; fi50 < 19; fi50++ {
                                    v51 := GenS4{F0: v4, F1: 9223372036854775807}
                                    _ = fi50
                                    _ = v51
                                }
                                _ = v38
                            }
                        }
                    }
                }
            case "a":
                v2 += (v2 % (((^(v2))) | 1))
            default:
                if (42 != 18446744073709551615) {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        v29 = append(v29, 127)
                        var v53 int64 = 730
                        var v54 int8 = -1
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    if (1e10 < 0.0) {
                        v2 += (v2 ^ (v2 / ((v2) | 1)))
                        switch "hello" {
                        case "a":
                            switch v4 {
                            case 0:
                                var v55 int32 = (v3 | v3)
                                _ = v55
                            case -822:
                                var v56 int32 = (v3 + v3)
                                _ = v56
                            case 9223372036854775807:
                                fmt.Println(v27, v3)
                            default:
                                v2 = 1
                            }
                        case "foo bar":
                            for fi57 := 0; fi57 < 19; fi57++ {
                                v4 = 32767
                                if v27 {
                                    for fi58 := 0; fi58 < 8; fi58++ {
                                        v29 = append(v29, fi58)
                                        v3 = 1
                                        _ = fi58
                                    }
                                    v59 := GenS1{F0: true, F1: 1}
                                    fmt.Println(v4)
                                    _ = v59
                                }
                                _ = fi57
                            }
                        case "\n\t":
                            v60 := GenS5{F0: 334}
                            _ = v60
                        }
                    }
                }
            }
        case -205:
            switch ("a" + "a") {
            case "":
                v61 := GenS6{F0: 65535}
                _ = v61
            case "a":
                var v62 uint64 = 333
                _ = v62
            case "foo bar":
                v4 = (v4 % ((v4) | 1))
            default:
                if len(v29) > 0 {
                    v29[0] = v4
                }
            }
        default:
            fmt.Println(v1, v1, v2)
        }
        _ = v29
    }
    switch v4 {
    case 1:
        v4 += 712
    }
    switch ("a" + "a") {
    case "foo bar":
        v3 += v3
    case "a":
        for fi63 := 0; fi63 < 12; fi63++ {
            if v27 {
                v64 := []int{v4, (v4 * (fi63 + 509))}
                v64 = append(v64, (fi63 % (((fi63 + (v4 ^ 0))) | 1)))
                if v27 {
                    if len(v64) > 0 {
                        v64[0] = -75
                    }
                } else {
                    for fi65 := 0; fi65 < 1; fi65++ {
                        if ("\n\t" == "\n\t") {
                            switch 32768 {
                            case -2147483648:
                                var v66 uint16 = 65535
                                _ = v66
                            }
                        } else {
                            v2 -= v2
                            v67 := GenS7{F0: -51, F1: 182, F2: int16(v1)}
                            _ = v67
                        }
                        switch ("foo bar" + "hello") {
                        case "hello":
                            v68 := GenS4{F0: (v4 >> ((18446744073709551615) & 63)), F1: -1}
                            _ = v68
                        case "a":
                            fmt.Println(v2)
                        default:
                            v69 := GenS8{F0: 805, F1: 470}
                            _ = v69
                        }
                        v3 -= v3
                        _ = fi65
                    }
                }
                _ = v64
            }
            switch "hello" {
            case "hello":
                v70 := []int{(v4 * (v4 | v4)), v4, v4, v4}
                _ = v70
            case "":
                if (v4 != (v4 << ((42) & 63))) {
                    v71 := GenS3{F0: 0}
                    _ = v71
                }
            default:
                fmt.Println(v2)
            }
            if v27 {
                fmt.Println(v4, v1, v4)
            } else {
                v1 -= v1
                for fi72 := 0; fi72 < 5; fi72++ {
                    v73 := []int{332, -962}
                    _ = fi72
                    _ = v73
                }
                v2 = v2
            }
            _ = fi63
        }
    case "hello":
        if v27 {
            v4 += (v4 >> ((42) & 63))
            var v74 int8 = -85
            v74 += v74
            _ = v74
        }
    default:
        v2 -= v2
    }
    var v75 float64 = float64(-2.25)
    v1 = 65535
    v3 = v3
    fmt.Println(v75)
    fmt.Println(v2)
    v4 = v4
    for fi76 := 0; fi76 < 20; fi76++ {
        v77 := GenS7{F0: 1, F1: 107, F2: -32768}
        v3 = int32(fi76)
        fmt.Println(v27)
        _ = fi76
        _ = v77
    }
    v78 := GenS8{F0: 2147483648, F1: 18446744073709551615}
    for fi79 := 0; fi79 < 8; fi79++ {
        switch 2147483647 {
        case -1:
            if true {
                var v80 bool = v27
                v4 -= v4
                v80 = ((^(v1)) > (^(v1)))
                _ = v80
            } else {
                v4 -= v4
                v81 := GenS9{F0: 59, F1: 42, F2: 1}
                v3 = 1
                _ = v81
            }
        case 869:
            v82 := []int{2147483647}
            _ = v82
        case 742:
            v83 := []int{}
            _ = v83
        default:
            v2 += 167
        }
        _ = fi79
    }
    v84 := GenS10{F0: 0, F1: (v3 + (v3 & 726))}
    v3 = (v3 + 243)
    switch (("hello" + "hello") + "foo bar") {
    case "hello":
        v85 := []int{v4, -141, v4}
        _ = v85
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v75
    _ = v78
    _ = v84
}
