package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int64; F1 uint; F2 int16 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 int16; F1 string; F2 int64 }
type GenS10 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []string{}
    var v2 int32 = 2147483647
    var v3 int = -9223372036854775808
    var v4 int64 = 807
    var v5 float64 = (-2.25 / 0.1)
    var v6 uint32 = 696
    var v7 string = ("\n\t" + ("hello" + "a"))
    var v8 bool = ((!true) || (42 == 2))
    switch -2147483648 {
    case 0:
        v4 += v4
    case 32767:
        if v8 {
            if true {
                switch ("\n\t" + v7) {
                case "":
                    var v9 float64 = v5
                    _ = v9
                case "a":
                    var v10 bool = v8
                    _ = v10
                case "hello":
                    fmt.Println(v2, v2, v7)
                }
            } else {
                for fi11 := 0; fi11 < 8; fi11++ {
                    v1 = append(v1, "")
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if (((3.14159 + 3.14159) + 0.0) > 0.0) {
                            var v13 int64 = v4
                            var v14 uint16 = 884
                            if v8 {
                                var v15 int64 = (v4 & (v13 ^ v13))
                                if false {
                                    var v16 int8 = 0
                                    var v17 int8 = (v16 - v16)
                                    _ = v16
                                    _ = v17
                                } else {
                                    v5 = ((float64((float32((0.0 * v5)) + 1.5)) + (v5 - (0.1 - v5))) / 1.0)
                                    _ = len(v1)
                                    v5 += (v5 - ((float64(1e10) + v5) * 0.1))
                                    v2 = 2147483647
                                }
                                if (v8 && ((!v8) || v8)) {
                                    var v18 int32 = v2
                                    _ = v18
                                }
                                _ = v15
                            }
                            _ = v13
                            _ = v14
                        }
                        v19 := GenS1{F0: uint8(v6), F1: 127}
                        v20 := GenS1{F0: 127, F1: 159}
                        _ = fi12
                        _ = v19
                        _ = v20
                    }
                    _ = fi11
                }
            }
            v6 = v6
        }
    default:
        v2 = (v2 ^ (v2 & -32))
    }
    _ = len(v1)
    if (0.1 <= float32(v5)) {
        if v8 {
            switch int(v4) {
            case 4294967295:
                switch v3 {
                case 42:
                    v6 = v6
                case 32767:
                    switch (v3 / (((-(v3))) | 1)) {
                    case 9223372036854775807:
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v22 := GenS1{F0: 150, F1: 255}
                            _ = fi21
                            _ = v22
                        }
                    case -1:
                        v23 := GenS2{F0: (v2 + v2), F1: 100}
                        _ = v23
                    default:
                        v3 -= (v3 | v3)
                    }
                case 2147483647:
                    if (!(1 < -1)) {
                        var v24 float64 = ((1.5 * 0.1) + 0.0)
                        switch v7 {
                        case "\n\t":
                            var v25 float32 = 0.1
                            _ = v25
                        }
                        _ = v24
                    } else {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi26
                        }
                        if true {
                            if false {
                                v8 = v8
                            }
                            v2 += v2
                        }
                    }
                }
            case 128:
                var v27 uint16 = 2
                _ = v27
            default:
                v28 := GenS2{F0: (v2 | 2147483647), F1: 29}
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                if v8 {
                    if ((((v4 & (v4 | v4)) <= int64(v3)) || v8) || v8) {
                        if (v8 || true) {
                            v5 = v5
                            switch v7 {
                            case "foo bar":
                                var v30 uint16 = 65535
                                _ = v30
                            case "\n\t":
                                v5 -= ((0.0 * v5) - 1.5)
                            default:
                                var v31 int32 = -1
                                _ = v31
                            }
                            switch v3 {
                            case 0:
                                var v32 int = (-(v3))
                                _ = v32
                            }
                        } else {
                            if (-1 != 10) {
                                if v8 {
                                    v3 = (v3 * 2147483647)
                                    v8 = (971 >= 0)
                                } else {
                                    var v33 int = -385
                                    _ = v33
                                }
                                if (-128 <= 0) {
                                    for fi34 := 0; fi34 < 8; fi34++ {
                                        v7 = v7
                                        _ = fi34
                                    }
                                    _ = len(v1)
                                } else {
                                    switch (v7 + v7) {
                                    case "":
                                        v35 := GenS1{F0: 127, F1: 42}
                                        _ = v35
                                    case "\n\t":
                                        var v36 uint8 = 119
                                        _ = v36
                                    default:
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            if ((1.5 - float64((float32(v5) - (0.0 + (0.1 * 1e10))))) < (((v5 - v5) * 3.14159) - (v5 * v5))) {
                                                fmt.Println(fi37)
                                            }
                                            v38 := GenS3{F0: -32768, F1: 1, F2: 173}
                                            _ = fi37
                                            _ = v38
                                        }
                                    }
                                }
                            }
                            v8 = v8
                        }
                        v5 += 0.0
                        v4 += (v4 % (((v4 - v4)) | 1))
                    }
                } else {
                    var v39 float64 = 1e10
                    if v8 {
                        _ = len(v1)
                        for fi40 := 0; fi40 < 2; fi40++ {
                            switch 32768 {
                            case 103:
                                if (v2 >= int32(v3)) {
                                    v41 := GenS4{F0: 2}
                                    if v8 {
                                        fmt.Println(v39)
                                        if (66 != 142) {
                                            v39 -= v5
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        switch (v3 / ((32767) | 1)) {
                                        case 912:
                                            v8 = v8
                                        case 836:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                    var v42 uint64 = 0
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 bool = false
                                    v44 := GenS5{F0: 523, F1: 1, F2: 0}
                                    _ = v43
                                    _ = v44
                                }
                                fmt.Println(v2)
                                var v45 float32 = 3.14159
                                fmt.Println(v4, v8)
                                if true {
                                    fmt.Println(v3)
                                    fmt.Println(v5, v7)
                                    v46 := GenS6{F0: (v4 << ((319) & 63)), F1: 4294967295, F2: 1}
                                    var v47 float32 = 3.14159
                                    _ = v46
                                    _ = v47
                                }
                                _ = v45
                            case 2147483647:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case -869:
                                v39 = v5
                            }
                            switch v3 {
                            case 32768:
                                if false {
                                    v4 = (-(v4))
                                }
                            case 0:
                                v4 = v4
                            case 9223372036854775807:
                                fmt.Println(v2, v4, fi29)
                                v4 = 42
                                if v8 {
                                    if (!false) {
                                        switch v7 {
                                        case "hello":
                                            switch fi29 {
                                            case 32768:
                                                var v48 uint32 = 4294967295
                                                _ = v48
                                            case 0:
                                                var v49 float64 = ((v39 + v5) / 0.1)
                                                _ = v49
                                            case 32767:
                                                v6 = (v6 % ((42) | 1))
                                            }
                                        case "foo bar":
                                            var v50 int64 = (v4 * v4)
                                            _ = v50
                                        case "a":
                                            if v8 {
                                                var v51 uint32 = 1
                                                v52 := GenS2{F0: 0, F1: 0}
                                                v53 := GenS5{F0: 42, F1: 18446744073709551615, F2: 358}
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                            } else {
                                                var v54 int16 = 1
                                                var v55 bool = true
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    } else {
                                        switch (v7 + v7) {
                                        case "foo bar":
                                            v56 := GenS1{F0: 253, F1: 2}
                                            _ = v56
                                        case "hello":
                                            var v57 bool = (!(v6 != v6))
                                            _ = v57
                                        }
                                        v6 += 65535
                                        v58 := GenS7{F0: 42, F1: v6}
                                        _ = v58
                                    }
                                    var v59 string = v7
                                    fmt.Println(v39, v7)
                                    switch int(v4) {
                                    case 9223372036854775807:
                                        if v8 {
                                            if (0 < 4294967295) {
                                                switch v7 {
                                                case "a":
                                                    if v8 {
                                                        if len(v1) > 0 {
                                                            v1[0] = "\n\t"
                                                        }
                                                        var v60 uint = 288
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v60
                                                    }
                                                default:
                                                    var v61 uint64 = 42
                                                    _ = v61
                                                }
                                                v62 := GenS2{F0: v2, F1: 0}
                                                var v63 int16 = 1000
                                                _ = v62
                                                _ = v63
                                            }
                                            var v64 uint = 148
                                            _ = v64
                                        }
                                    default:
                                        if v8 {
                                            v39 = 0.1
                                            if true {
                                                v2 = v2
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                var v65 uint64 = 42
                                                _ = v65
                                            }
                                        }
                                    }
                                    _ = v59
                                }
                                if (4294967295 == 296) {
                                    _ = len(v1)
                                    var v66 int64 = -1
                                    switch 181 {
                                    case 127:
                                        v3 = (v3 / (((v3 << ((4294967295) & 63))) | 1))
                                    case 2147483647:
                                        var v67 int8 = -128
                                        _ = v67
                                    default:
                                        v5 -= float64(-2.25)
                                    }
                                    _ = v66
                                } else {
                                    v8 = ((v2 / ((v2) | 1)) <= v2)
                                    fmt.Println(v6, fi40, v8)
                                    v7 = v7
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v6 += v6
                                }
                                v39 += v5
                            default:
                                fmt.Println(v2, v8)
                            }
                            switch (-(fi29)) {
                            case 156:
                                var v68 int8 = -128
                                _ = v68
                            case 127:
                                switch "foo bar" {
                                case "\n\t":
                                    v4 += v4
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            _ = fi40
                        }
                        v7 = (v7 + (v7 + v7))
                    } else {
                        if len(v1) > 0 {
                            v1[0] = v7
                        }
                    }
                    var v69 int16 = -168
                    _ = v39
                    _ = v69
                }
                if false {
                    if true {
                        var v70 int8 = 100
                        v5 = 1.5
                        _ = v70
                    }
                }
                _ = fi29
            }
        } else {
            if (1 <= (^(v3))) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch (v3 - v3) {
                case 2147483647:
                    for fi71 := 0; fi71 < 3; fi71++ {
                        v72 := GenS8{F0: v2, F1: 65535}
                        _ = fi71
                        _ = v72
                    }
                }
                for fi73 := 0; fi73 < 3; fi73++ {
                    var v74 int32 = -373
                    switch (fi73 + 42) {
                    case 600:
                        if v8 {
                            var v75 uint = 954
                            var v76 uint32 = v6
                            for fi77 := 0; fi77 < 12; fi77++ {
                                v78 := GenS9{F0: -1, F1: ("" + "foo bar"), F2: -1}
                                switch (fi77 ^ fi77) {
                                case -9223372036854775808:
                                    switch (v7 + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(v7, v3, v76)
                                    case "":
                                        v79 := GenS7{F0: 0, F1: v6}
                                        _ = v79
                                    case "a":
                                        v8 = (v8 && false)
                                    }
                                case 32767:
                                    v1 = append(v1, v7)
                                default:
                                    var v80 int32 = 437
                                    _ = v80
                                }
                                _ = fi77
                                _ = v78
                            }
                            _ = v75
                            _ = v76
                        } else {
                            switch v7 {
                            case "hello":
                                switch "\n\t" {
                                case "\n\t":
                                    if v8 {
                                        if v8 {
                                            var v81 int64 = -1
                                            _ = v81
                                        }
                                        var v82 uint = 4294967295
                                        var v83 uint64 = 858
                                        _ = v82
                                        _ = v83
                                    }
                                default:
                                    v84 := GenS5{F0: 32767, F1: 18446744073709551615, F2: 520}
                                    _ = v84
                                }
                            default:
                                if (!v8) {
                                    v7 = "a"
                                    var v85 uint64 = 1
                                    switch fi73 {
                                    case 32767:
                                        for fi86 := 0; fi86 < 12; fi86++ {
                                            v8 = v8
                                            fmt.Println(v7, v3)
                                            _ = fi86
                                        }
                                    }
                                    _ = v85
                                }
                            }
                            v2 = -1
                        }
                    case 0:
                        for fi87 := 0; fi87 < 20; fi87++ {
                            fmt.Println(fi87, fi73, v7)
                            _ = fi87
                        }
                    default:
                        v88 := GenS10{F0: 363, F1: 1}
                        _ = v88
                    }
                    _ = fi73
                    _ = v74
                }
            }
        }
        var v89 float64 = v5
        _ = v89
    }
    for fi90 := 0; fi90 < 16; fi90++ {
        fmt.Println(fi90, v7, v2)
        _ = fi90
    }
    fmt.Println(v7, v6)
    v6 -= (^(v6))
    if (false && (-14 <= 1)) {
        if len(v1) > 0 {
            v1[0] = ("\n\t" + v7)
        }
    } else {
        switch int(v6) {
        case 0:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -396:
            _ = len(v1)
        }
    }
    var v91 uint32 = (^(v6))
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v91
}
