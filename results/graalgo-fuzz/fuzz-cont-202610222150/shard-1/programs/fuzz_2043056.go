package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int; F2 uint }
type GenS6 struct { F0 int16; F1 uint; F2 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 string; F2 int8 }
type GenS10 struct { F0 uint64; F1 uint8; F2 uint }
type GenS11 struct { F0 int32; F1 int32; F2 uint32 }
type GenS12 struct { F0 uint8; F1 uint64 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint16; F1 string }
type GenS16 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS17 struct { F0 int64; F1 string; F2 int64 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 uint16; F1 bool }
type GenS21 struct { F0 int16 }

func main() {
    var v1 float32 = float32((3.14159 + float64(1.5)))
    var v2 int32 = -1
    v3 := []uint32{}
    fmt.Println(v2, v1)
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v4 float64 = (-2.25 - float64((((0.1 * 0.1) + v1) * 1.5)))
    var v5 float64 = 3.14159
    fmt.Println(v5, v2)
    var v6 uint16 = 50
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int32 = 2147483647
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int32 = -175
            if (255 == 255) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v2 = (v2 | v2)
            } else {
                v5 = v4
                v2 = v10
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    switch 765 {
    case 2147483647:
        switch ("foo bar" + ("" + "hello")) {
        case "\n\t":
            if (2 <= 183) {
                v6 -= v6
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        switch "foo bar" {
                        case "hello":
                            v1 = v1
                        default:
                            v2 = v2
                        }
                        v13 := GenS1{F0: 1000}
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            }
        }
    case 127:
        var v14 float32 = v1
        _ = v14
    case -660:
        v15 := GenS1{F0: 42}
        _ = v15
    default:
        fmt.Println(v4, v5, v5)
    }
    var v16 uint32 = 2
    switch ("foo bar" + ("hello" + "a")) {
    case "":
        v6 = v6
    default:
        if (true && (2 <= 1)) {
            switch 463 {
            case 9223372036854775807:
                for fi17 := 0; fi17 < 10; fi17++ {
                    if true {
                        v4 += ((v5 / 0.1) + v4)
                        if false {
                            v18 := GenS2{F0: -1}
                            for fi19 := 0; fi19 < 2; fi19++ {
                                if ((1.5 + 3.14159) >= (-2.25 - (-2.25 - v4))) {
                                    fmt.Println(v16, v2, v4)
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    _ = len(v3)
                                }
                                if true {
                                    if false {
                                        fmt.Println(fi17)
                                        if ("foo bar" >= "hello") {
                                            _ = len(v3)
                                            if false {
                                                v6 = 255
                                                switch fi17 {
                                                case 32768:
                                                    v6 += v6
                                                case -9223372036854775808:
                                                    v20 := GenS2{F0: 99}
                                                    _ = v20
                                                }
                                                v3 = append(v3, 1)
                                            } else {
                                                var v21 string = ("\n\t" + ("foo bar" + "hello"))
                                                if false {
                                                    var v22 float64 = v4
                                                    if false {
                                                        if (!(!(v1 >= float32(v5)))) {
                                                            v6 = (v6 >> ((4294967295) & 63))
                                                            v23 := GenS3{F0: 0, F1: uint64(v16), F2: 42}
                                                            v24 := GenS2{F0: 0}
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                    }
                                                    var v25 int32 = int32(v6)
                                                    _ = v22
                                                    _ = v25
                                                } else {
                                                    if false {
                                                        if (fi19 >= -9223372036854775808) {
                                                            v21 = v21
                                                        } else {
                                                            v3 = append(v3, (v16 - v16))
                                                            v6 -= 672
                                                            if true {
                                                                var v26 int32 = 540
                                                                _ = v26
                                                            } else {
                                                                switch ("a" + ((v21 + "hello") + v21)) {
                                                                case "a":
                                                                    fmt.Println(fi19, v16)
                                                                case "\n\t":
                                                                    switch (^(fi17)) {
                                                                    case -40:
                                                                        switch fi17 {
                                                                        case 2147483647:
                                                                            v27 := GenS1{F0: 32767}
                                                                            _ = v27
                                                                        case -2147483648:
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        case 1:
                                                                            var v28 float32 = ((v1 - (float32(3.14159) + -2.25)) * v1)
                                                                            _ = v28
                                                                        }
                                                                    case 9223372036854775807:
                                                                        _ = len(v3)
                                                                    }
                                                                case "foo bar":
                                                                    v21 = v21
                                                                }
                                                                v21 = (v21 + v21)
                                                                var v29 float32 = (0.1 - v1)
                                                                _ = v29
                                                            }
                                                        }
                                                        v2 = v2
                                                    }
                                                }
                                                var v30 float64 = -2.25
                                                _ = v21
                                                _ = v30
                                            }
                                            fmt.Println(v1, v5, v6)
                                        } else {
                                            v31 := GenS2{F0: 0}
                                            _ = v31
                                        }
                                        v3 = append(v3, v16)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = len(v3)
                                }
                                _ = fi19
                            }
                            _ = v18
                        } else {
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                                for fi32 := 0; fi32 < 2; fi32++ {
                                    v6 = v6
                                    switch "hello" {
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v4, v5, fi32)
                                            if (!(1 <= 42)) {
                                                v33 := GenS4{F0: 42}
                                                v34 := GenS5{F0: 719, F1: (^(fi17)), F2: 18446744073709551615}
                                                var v35 int8 = -50
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            } else {
                                                fmt.Println(v4, v4)
                                                if (true || (!true)) {
                                                    if true {
                                                        v36 := GenS6{F0: -332, F1: 1, F2: 1}
                                                        _ = v36
                                                    } else {
                                                        var v37 bool = (!true)
                                                        v3 = append(v3, 2)
                                                        if v37 {
                                                            var v38 bool = (!true)
                                                            var v39 uint64 = 2
                                                            v3 = append(v3, (v16 + v16))
                                                            if (1 <= 65535) {
                                                                var v40 bool = (v4 <= (v4 / 0.1))
                                                                v41 := GenS4{F0: 42}
                                                                if (v39 >= 2) {
                                                                    v39 += v39
                                                                    switch fi32 {
                                                                    case 32768:
                                                                        if true {
                                                                            _ = len(v3)
                                                                            v42 := GenS7{F0: 32767}
                                                                            _ = v42
                                                                        } else {
                                                                            v1 -= (v1 - (((1e10 / 1e10) + (((v1 + (1.5 * v1)) * v1) - 0.1)) * -2.25))
                                                                            v37 = ((v2 / (((^(v2))) | 1)) == (v2 * -1))
                                                                            v2 += (v2 + v2)
                                                                        }
                                                                    case -9223372036854775808:
                                                                        var v43 int32 = 892
                                                                        _ = v43
                                                                    case 429:
                                                                        if (!(uint(v6) >= 0)) {
                                                                            if v40 {
                                                                                v4 = 3.14159
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            if v40 {
                                                                                if (((370 < (^(v2))) || true) || (v37 || ((v40 && false) && true))) {
                                                                                    var v44 uint32 = 706
                                                                                    var v45 bool = (!true)
                                                                                    var v46 float64 = (v4 * 0.0)
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                    _ = v46
                                                                                }
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                v4 += (float64(v1) / v5)
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v47 uint64 = (v39 & v39)
                                                                    if v40 {
                                                                        _ = len(v3)
                                                                        v4 = 1e10
                                                                        v5 = 1.5
                                                                    } else {
                                                                        v4 = (float64(v1) * v5)
                                                                        var v48 int64 = -1
                                                                        v48 -= v48
                                                                        if v40 {
                                                                            var v49 uint64 = v39
                                                                            var v50 uint = uint(v6)
                                                                            v51 := GenS8{F0: 127}
                                                                            _ = v49
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                        v48 += v48
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v52 := GenS8{F0: 32}
                                                                            _ = v52
                                                                        default:
                                                                            if ((!v40) || v38) {
                                                                                fmt.Println(v48, v40)
                                                                                v39 = (v39 & (^(v47)))
                                                                            }
                                                                        }
                                                                        v16 = uint32(v39)
                                                                        _ = v48
                                                                    }
                                                                    _ = v47
                                                                }
                                                                fmt.Println(v39)
                                                                _ = v40
                                                                _ = v41
                                                            } else {
                                                                switch -210 {
                                                                case 20:
                                                                    var v53 bool = v37
                                                                    _ = v53
                                                                case 42:
                                                                    fmt.Println(v39, fi17, v2)
                                                                case 128:
                                                                    fmt.Println(v4)
                                                                default:
                                                                    switch ("hello" + (((("" + "hello") + ("hello" + (("foo bar" + "") + "foo bar"))) + "hello") + "hello")) {
                                                                    case "":
                                                                        if v37 {
                                                                            if ((2 < 1) || false) {
                                                                                var v54 uint = 4294967295
                                                                                v3 = append(v3, 4294967295)
                                                                                fmt.Println(v6, fi32)
                                                                                _ = v54
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        var v55 int = (fi17 ^ fi32)
                                                                        _ = v55
                                                                    }
                                                                }
                                                                _ = len(v3)
                                                                switch (("foo bar" + "") + "a") {
                                                                case "hello":
                                                                    _ = len(v3)
                                                                case "foo bar":
                                                                    v5 = (v4 / 3.14159)
                                                                case "a":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        _ = v37
                                                    }
                                                    if (!((v6 * 32767) == (v6 + v6))) {
                                                        v56 := GenS9{F0: -64, F1: ("\n\t" + ("foo bar" + "hello")), F2: 127}
                                                        v57 := GenS8{F0: -59}
                                                        v16 += v16
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        switch fi32 {
                                                        case -2147483648:
                                                            v16 += 65535
                                                        case -9223372036854775808:
                                                            switch fi32 {
                                                            case 4294967295:
                                                                if len(v3) > 0 {
                                                                    v3[0] = (^(v16))
                                                                }
                                                            case 0:
                                                                var v58 string = ((("hello" + "") + "foo bar") + (((("" + (("foo bar" + "\n\t") + "foo bar")) + "") + ("" + "foo bar")) + ""))
                                                                _ = v58
                                                            case 990:
                                                                if false {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if ((v6 & (^(v6))) > (v6 ^ v6)) {
                                                                            switch -1 {
                                                                            case 127:
                                                                                var v59 int16 = 860
                                                                                _ = v59
                                                                            case -446:
                                                                                switch (fi17 / ((-134) | 1)) {
                                                                                case 4294967295:
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                case 0:
                                                                                    v60 := GenS7{F0: (v6 * v6)}
                                                                                    _ = v60
                                                                                }
                                                                            default:
                                                                                v61 := GenS10{F0: 1, F1: 42, F2: 98}
                                                                                _ = v61
                                                                            }
                                                                            v62 := GenS1{F0: 1000}
                                                                            v2 -= v2
                                                                            _ = v62
                                                                        }
                                                                    case "":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 42
                                                                        }
                                                                    case "hello":
                                                                        var v63 float32 = v1
                                                                        _ = v63
                                                                    default:
                                                                        var v64 bool = (true || (!true))
                                                                        _ = v64
                                                                    }
                                                                    v65 := GenS9{F0: -496, F1: "a", F2: 127}
                                                                    v66 := GenS11{F0: (^(v2)), F1: v2, F2: 1}
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                _ = len(v3)
                                                                if false {
                                                                    v1 = v1
                                                                    if (!true) {
                                                                        var v67 int32 = v2
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v16 - v16)
                                                                        }
                                                                        if true {
                                                                            var v68 int8 = -57
                                                                            _ = v68
                                                                        } else {
                                                                            v69 := GenS12{F0: 255, F1: 1}
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 2
                                                                            }
                                                                            var v70 float64 = (((v5 / v4) * ((-2.25 + (v4 * v4)) + (v5 + (-2.25 / (v4 - 0.1))))) - 1.5)
                                                                            _ = v69
                                                                            _ = v70
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                    _ = len(v3)
                                                                } else {
                                                                    v71 := GenS3{F0: -933, F1: 0, F2: 330}
                                                                    v1 -= v1
                                                                    fmt.Println(fi32, v1, v6)
                                                                    _ = v71
                                                                }
                                                            default:
                                                                if (("hello" > ("foo bar" + "")) || (!(4294967295 <= 1))) {
                                                                    switch fi32 {
                                                                    case -319:
                                                                        switch int(v2) {
                                                                        case -925:
                                                                            v72 := GenS10{F0: uint64(v16), F1: 128, F2: 42}
                                                                            _ = v72
                                                                        default:
                                                                            switch "hello" {
                                                                            case "":
                                                                                switch (^(fi17)) {
                                                                                case -9223372036854775808:
                                                                                    var v73 uint64 = 423
                                                                                    _ = v73
                                                                                case 4294967295:
                                                                                    v1 += v1
                                                                                }
                                                                            case "foo bar":
                                                                                if ((false && false) || false) {
                                                                                    v4 -= v5
                                                                                    switch (-(fi32)) {
                                                                                    case -389:
                                                                                        v1 = float32((v4 - 1e10))
                                                                                    case 900:
                                                                                        v3 = append(v3, 503)
                                                                                    default:
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            fmt.Println(fi32, v5, fi32)
                                                                                        case "hello":
                                                                                            v6 = v6
                                                                                        default:
                                                                                            var v74 int64 = 0
                                                                                            _ = v74
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if ((true || true) && false) {
                                                                                    fmt.Println(v2, v5)
                                                                                    v6 -= (v6 ^ 509)
                                                                                } else {
                                                                                    v3 = append(v3, v16)
                                                                                    v75 := GenS11{F0: 2147483647, F1: -254, F2: (v16 << ((255) & 63))}
                                                                                    v76 := GenS11{F0: v2, F1: (v2 << ((765) & 63)), F2: (v16 + (v16 / ((v16) | 1)))}
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch ((("foo bar" + "a") + ("a" + "foo bar")) + (((("foo bar" + "") + "\n\t") + "\n\t") + ("foo bar" + "hello"))) {
                                                                        case "hello":
                                                                            if false {
                                                                                v6 += (v6 ^ v6)
                                                                            } else {
                                                                                _ = len(v3)
                                                                                var v77 uint64 = 456
                                                                                _ = v77
                                                                            }
                                                                        default:
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v3 = append(v3, v16)
                                                                            case "hello":
                                                                                v2 = v2
                                                                            case "a":
                                                                                v16 = (^(v16))
                                                                            default:
                                                                                var v78 float32 = v1
                                                                                _ = v78
                                                                            }
                                                                            var v79 string = (("hello" + "a") + ((("" + "\n\t") + "\n\t") + "\n\t"))
                                                                            if true {
                                                                                v5 -= (1.5 + 0.1)
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v80 uint = 121
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                switch (^(fi32)) {
                                                                                case -536:
                                                                                    if (0 != 126) {
                                                                                        var v81 int8 = 1
                                                                                        if (v81 != (v81 * v81)) {
                                                                                            switch v79 {
                                                                                            case "a":
                                                                                                var v82 uint32 = v16
                                                                                                _ = v82
                                                                                            case "foo bar":
                                                                                                if false {
                                                                                                    var v83 int32 = 2147483647
                                                                                                    switch (fi32 + 127) {
                                                                                                    case 127:
                                                                                                        v84 := GenS10{F0: 18446744073709551615, F1: 2, F2: (v80 << ((v80) & 63))}
                                                                                                        _ = v84
                                                                                                    case -1:
                                                                                                        var v85 int = -717
                                                                                                        var v86 float32 = v1
                                                                                                        v81 = -1
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    case 32767:
                                                                                                        var v87 int64 = 42
                                                                                                        _ = v87
                                                                                                    default:
                                                                                                        v2 = (-(v2))
                                                                                                        if true {
                                                                                                            v88 := GenS13{F0: v81}
                                                                                                            v16 = (^(v16))
                                                                                                            _ = v88
                                                                                                        } else {
                                                                                                            v81 += v81
                                                                                                        }
                                                                                                    }
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = v16
                                                                                                    }
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    v81 = (^(v81))
                                                                                                }
                                                                                            case "\n\t":
                                                                                                fmt.Println(fi32)
                                                                                            }
                                                                                            v89 := GenS14{F0: v81}
                                                                                            _ = v89
                                                                                        }
                                                                                        if true {
                                                                                            fmt.Println(v6, v4)
                                                                                            v90 := GenS6{F0: 42, F1: (v80 / (((v80 & uint(v6))) | 1)), F2: 113}
                                                                                            _ = v90
                                                                                        }
                                                                                        _ = v81
                                                                                    } else {
                                                                                        v91 := GenS15{F0: 0, F1: v79}
                                                                                        v92 := GenS7{F0: v6}
                                                                                        var v93 float64 = (1e10 - v4)
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    v3 = append(v3, v16)
                                                                                case 838:
                                                                                    fmt.Println(v79, v6)
                                                                                    v94 := GenS16{F0: -670, F1: 255, F2: v16}
                                                                                    _ = v94
                                                                                }
                                                                                v6 += (^(v6))
                                                                                _ = v80
                                                                            }
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                    v5 += (v5 / (float64(1e10) - v4))
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 65535
                                                                    }
                                                                }
                                                            }
                                                        case -853:
                                                            v95 := GenS1{F0: 42}
                                                            _ = v95
                                                        }
                                                        var v96 uint64 = 42
                                                        var v97 uint = 4294967295
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                }
                                                var v98 string = "\n\t"
                                                _ = v98
                                            }
                                        }
                                    case "a":
                                        v4 -= v5
                                    default:
                                        v99 := GenS13{F0: -56}
                                        _ = v99
                                    }
                                    v100 := GenS17{F0: 0, F1: ((("\n\t" + "\n\t") + "hello") + "foo bar"), F2: 0}
                                    _ = fi32
                                    _ = v100
                                }
                            } else {
                                v6 = v6
                                v101 := GenS17{F0: 1, F1: "", F2: 1000}
                                _ = v101
                            }
                            v102 := GenS18{F0: 1000}
                            _ = v102
                        }
                        var v103 uint32 = (v16 + v16)
                        _ = v103
                    } else {
                        v6 += v6
                    }
                    _ = fi17
                }
            case -583:
                v1 = 0.1
            }
            if (!((false || false) || true)) {
                switch 32767 {
                case 32767:
                    var v104 string = (("\n\t" + "a") + (("foo bar" + "hello") + ("\n\t" + "")))
                    _ = v104
                case -911:
                    if true {
                        var v105 float32 = v1
                        fmt.Println(v2)
                        var v106 uint64 = 2
                        _ = v105
                        _ = v106
                    } else {
                        if len(v3) > 0 {
                            v3[0] = uint32(v6)
                        }
                    }
                case 238:
                    switch "" {
                    case "hello":
                        v107 := GenS4{F0: -128}
                        _ = v107
                    case "":
                        var v108 uint16 = 1
                        _ = v108
                    }
                }
            }
            for fi109 := 0; fi109 < 12; fi109++ {
                v110 := GenS13{F0: 0}
                switch (("a" + "hello") + ("foo bar" + "hello")) {
                case "foo bar":
                    for fi111 := 0; fi111 < 5; fi111++ {
                        switch "a" {
                        case "foo bar":
                            var v112 bool = ((!(816 != 1)) || true)
                            _ = v112
                        case "a":
                            v113 := GenS13{F0: 0}
                            _ = v113
                        }
                        _ = fi111
                    }
                case "hello":
                    var v114 float64 = (float64(((v1 / 1.0) * v1)) / (v5 * 3.14159))
                    _ = v114
                case "":
                    fmt.Println(fi109)
                }
                v5 = ((v4 - (v5 * float64(v1))) + 3.14159)
                _ = fi109
                _ = v110
            }
        } else {
            var v115 uint = 0
            var v116 uint16 = v6
            if (42 != (^(v115))) {
                v6 = v116
                switch (("hello" + "foo bar") + "a") {
                case "\n\t":
                    var v117 uint8 = 2
                    _ = v117
                case "a":
                    switch -553 {
                    case 1:
                        if (true && false) {
                            fmt.Println(v16, v115, v116)
                            if false {
                                for fi118 := 0; fi118 < 2; fi118++ {
                                    var v119 uint64 = 0
                                    var v120 int32 = 550
                                    _ = fi118
                                    _ = v119
                                    _ = v120
                                }
                                v121 := GenS19{F0: -1}
                                _ = v121
                            }
                        }
                    case 118:
                        switch 128 {
                        case 127:
                            for fi122 := 0; fi122 < 14; fi122++ {
                                fmt.Println(v16)
                                var v123 float64 = 0.0
                                _ = fi122
                                _ = v123
                            }
                        case 32767:
                            for fi124 := 0; fi124 < 2; fi124++ {
                                if (true || true) {
                                    var v125 bool = ((!((false || ((-249 < (v2 >> (((v115 & v115)) & 63))) || (!(127 != 183)))) || ((fi124 / ((fi124) | 1)) != (-(fi124))))) && true)
                                    for fi126 := 0; fi126 < 17; fi126++ {
                                        _ = len(v3)
                                        _ = fi126
                                    }
                                    _ = v125
                                }
                                var v127 float32 = -2.25
                                _ = fi124
                                _ = v127
                            }
                        default:
                            v128 := GenS20{F0: 2, F1: (!(true || false))}
                            _ = v128
                        }
                    }
                }
                v129 := GenS21{F0: 1000}
                _ = v129
            } else {
                var v130 uint8 = 191
                fmt.Println(v16, v115)
                _ = v130
            }
            _ = v115
            _ = v116
        }
    }
    fmt.Println(v6, v4, v2)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
