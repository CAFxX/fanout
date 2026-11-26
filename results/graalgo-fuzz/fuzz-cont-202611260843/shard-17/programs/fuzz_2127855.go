package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 int; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string; F1 uint }

func main() {
    var v1 uint16 = 1
    var v2 int16 = -1
    var v3 float64 = -2.25
    var v4 uint16 = 0
    var v5 int8 = -97
    v6 := []int{}
    switch "\n\t" {
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = 32768
        }
    case "a":
        fmt.Println(v4, v4)
    case "":
        fmt.Println(v1, v2, v1)
    }
    switch -1 {
    case 9:
        v2 = v2
    default:
        switch ((("hello" + ("a" + ("foo bar" + ""))) + ("foo bar" + ("\n\t" + ("\n\t" + "hello")))) + ("a" + "hello")) {
        case "\n\t":
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v7 uint8 = 1
                    _ = v7
                } else {
                    var v8 int32 = 2147483647
                    _ = v8
                }
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                var v9 uint64 = 0
                _ = v9
            }
        case "hello":
            for fi10 := 0; fi10 < 10; fi10++ {
                v4 += v1
                _ = fi10
            }
        }
    }
    var v11 float64 = 3.14159
    switch 32768 {
    case -9223372036854775808:
        v1 = v4
    case -431:
        for fi12 := 0; fi12 < 5; fi12++ {
            v6 = append(v6, -922)
            _ = fi12
        }
    case 32767:
        if (65535 <= 237) {
            v2 = (v2 % ((v2) | 1))
        }
    }
    if true {
        var v13 uint8 = 128
        switch ("\n\t" + ("a" + "a")) {
        case "\n\t":
            var v14 uint32 = 42
            _ = v14
        case "":
            switch "" {
            case "\n\t":
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 uint16 = 255
                    _ = fi15
                    _ = v16
                }
            case "":
                v6 = append(v6, 340)
            case "foo bar":
                v2 += v2
            default:
                v4 -= 42
            }
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            if true {
                v18 := GenS1{F0: (v13 + (^(v13))), F1: 1, F2: (v2 >> ((1) & 63))}
                if (fi17 == int(v2)) {
                    v2 += (v2 | v2)
                }
                _ = v18
            }
            _ = fi17
        }
        _ = v13
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 bool = ((1e10 / v3) != v3)
            if len(v6) > 0 {
                v6[0] = (-(fi19))
            }
            if v20 {
                var v21 int16 = (v2 - v2)
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        switch "a" {
        case "hello":
            var v22 int64 = int64(v4)
            _ = v22
        case "":
            if (false && false) {
                v1 = (v4 % ((423) | 1))
            }
        case "foo bar":
            fmt.Println(v4)
        default:
            var v23 float64 = (v3 + v3)
            _ = v23
        }
        var v24 float64 = v3
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        switch "foo bar" {
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = 830
            }
        default:
            for fi26 := 0; fi26 < 6; fi26++ {
                switch 0 {
                case 32767:
                    v6 = append(v6, fi25)
                case 0:
                    switch (fi25 & (fi25 & -456)) {
                    case -9223372036854775808:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case 32767:
                        var v27 uint = 18446744073709551615
                        _ = v27
                    default:
                        var v28 uint = 1
                        _ = v28
                    }
                }
                _ = fi26
            }
        }
        if (((!false) && false) && (!(42 == 0))) {
            if (false || true) {
                var v29 string = "a"
                _ = v29
            } else {
                var v30 uint16 = (^(v1))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v31 := GenS1{F0: 128, F1: -542, F2: (-(v2))}
                _ = v30
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            switch ("\n\t" + "") {
            case "foo bar":
                v33 := GenS1{F0: 1, F1: 1, F2: v2}
                _ = v33
            case "a":
                var v34 int64 = -9223372036854775808
                _ = v34
            case "hello":
                v3 -= -2.25
            default:
                v35 := GenS1{F0: uint8(v4), F1: 0, F2: (v2 & v2)}
                _ = v35
            }
            _ = fi32
        }
        _ = fi25
    }
    if (v11 < v11) {
        v5 -= v5
    }
    v36 := GenS2{F0: (v4 << ((18446744073709551615) & 63)), F1: -1, F2: v5}
    switch -371 {
    case -2147483648:
        var v37 int64 = 42
        _ = v37
    case 4294967295:
        var v38 int = 32768
        _ = v38
    }
    v11 = (v11 / 3.14159)
    v6 = append(v6, -2147483648)
    v1 = v1
    if ("" > "foo bar") {
        if true {
            switch -595 {
            case 42:
                if (127 != -1) {
                    if (true && (0.1 <= float64(-2.25))) {
                        var v39 uint32 = 2
                        if len(v6) > 0 {
                            v6[0] = -9223372036854775808
                        }
                        _ = v39
                    } else {
                        if (4294967295 > 2147483647) {
                            switch 0 {
                            case 0:
                                fmt.Println(v1, v4, v3)
                            default:
                                switch 32768 {
                                case 322:
                                    v11 = v11
                                case 9223372036854775807:
                                    v5 -= -128
                                case -2147483648:
                                    v4 -= (v1 ^ (^(v1)))
                                }
                            }
                            fmt.Println(v5, v4)
                            _ = len(v6)
                        }
                    }
                    for fi40 := 0; fi40 < 5; fi40++ {
                        switch (fi40 & fi40) {
                        case 9223372036854775807:
                            switch (fi40 ^ fi40) {
                            case 32767:
                                if (!(!((0.1 != 0.1) || (false && true)))) {
                                    if (!(127 <= 0)) {
                                        if false {
                                            for fi41 := 0; fi41 < 12; fi41++ {
                                                if (42 > 880) {
                                                    if len(v6) > 0 {
                                                        v6[0] = (fi41 * fi41)
                                                    }
                                                    v42 := GenS1{F0: 2, F1: -111, F2: 1000}
                                                    var v43 uint = 65535
                                                    _ = v42
                                                    _ = v43
                                                }
                                                var v44 int32 = 42
                                                v6 = append(v6, fi41)
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = fi41
                                                _ = v44
                                            }
                                            v45 := GenS2{F0: v1, F1: fi40, F2: v5}
                                            _ = v45
                                        } else {
                                            v2 += (^(v2))
                                            if len(v6) > 0 {
                                                v6[0] = (fi40 & fi40)
                                            }
                                        }
                                        var v46 float64 = 0.1
                                        if (-850 != 1) {
                                            var v47 uint = 79
                                            v4 += 32767
                                            v3 = float64(0.1)
                                            _ = v47
                                        } else {
                                            switch ("" + ("\n\t" + ("a" + "\n\t"))) {
                                            case "":
                                                if false {
                                                    switch ("\n\t" + "hello") {
                                                    case "\n\t":
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case "hello":
                                                        switch fi40 {
                                                        case -2147483648:
                                                            v3 = v11
                                                        case 1:
                                                            if true {
                                                                fmt.Println(v3)
                                                                if false {
                                                                    var v48 int64 = 42
                                                                    v49 := GenS3{F0: 1, F1: 815}
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                            } else {
                                                                var v50 uint16 = (^(v1))
                                                                _ = v50
                                                            }
                                                        }
                                                    default:
                                                        var v51 bool = true
                                                        _ = v51
                                                    }
                                                }
                                            case "a":
                                                fmt.Println(v5)
                                            case "\n\t":
                                                fmt.Println(fi40, v5)
                                            }
                                            v52 := GenS4{F0: 0, F1: 63}
                                            var v53 int = (-(fi40))
                                            _ = v52
                                            _ = v53
                                        }
                                        _ = v46
                                    }
                                    switch 128 {
                                    case 32768:
                                        v54 := GenS2{F0: (v4 + 2), F1: int(v5), F2: 1}
                                        _ = v54
                                    case 32767:
                                        v11 = v11
                                    default:
                                        v3 = float64((3.14159 + (0.1 * (float32(v3) / 3.14159))))
                                    }
                                } else {
                                    var v55 int = (fi40 + fi40)
                                    _ = v55
                                }
                            case -1:
                                v56 := GenS5{F0: v2}
                                _ = v56
                            }
                        case 32768:
                            v5 = v5
                        default:
                            for fi57 := 0; fi57 < 8; fi57++ {
                                v11 += v3
                                switch 515 {
                                case 184:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                fmt.Println(v1, v4, fi57)
                                _ = fi57
                            }
                        }
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            fmt.Println(v3, v1, fi40)
                        case "hello":
                            var v58 bool = (float32(v11) > 0.1)
                            _ = v58
                        case "a":
                            v59 := GenS6{F0: 42, F1: (fi40 & fi40), F2: v2}
                            _ = v59
                        }
                        if (true || (4294967295 == 0)) {
                            v4 = v1
                            var v60 uint = 42
                            _ = v60
                        }
                        _ = fi40
                    }
                    v1 = v4
                }
            case 9223372036854775807:
                if false {
                    var v61 string = ""
                    _ = v61
                } else {
                    v62 := GenS7{F0: v2}
                    switch 188 {
                    case -2147483648:
                        switch 4294967295 {
                        case 128:
                            switch "a" {
                            case "a":
                                if false {
                                    _ = len(v6)
                                    if true {
                                        fmt.Println(v4, v1)
                                    } else {
                                        v3 -= 0.1
                                        var v63 uint = 150
                                        _ = v63
                                    }
                                } else {
                                    var v64 uint16 = 65535
                                    for fi65 := 0; fi65 < 15; fi65++ {
                                        for fi66 := 0; fi66 < 17; fi66++ {
                                            var v67 int = (-(fi66))
                                            _ = fi66
                                            _ = v67
                                        }
                                        v68 := GenS5{F0: v2}
                                        v64 = (^(v4))
                                        _ = fi65
                                        _ = v68
                                    }
                                    v69 := GenS8{F0: 127}
                                    _ = v64
                                    _ = v69
                                }
                            case "foo bar":
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    var v71 int16 = -32768
                                    _ = fi70
                                    _ = v71
                                }
                            case "":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        case -907:
                            var v72 float64 = 0.1
                            _ = v72
                        case 1:
                            v73 := GenS8{F0: uint8(v4)}
                            _ = v73
                        }
                    case -386:
                        fmt.Println(v1, v11)
                    default:
                        var v74 float64 = 1e10
                        _ = v74
                    }
                    _ = v62
                }
            case 32767:
                fmt.Println(v5, v5, v4)
            default:
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
            }
            fmt.Println(v1, v3)
        } else {
            v3 = v11
        }
        v4 -= v1
        switch ("a" + "hello") {
        case "\n\t":
            v75 := GenS9{F0: 18446744073709551615}
            _ = v75
        case "foo bar":
            v4 += v1
        case "":
            v3 -= (0.1 / -2.25)
        default:
            if len(v6) > 0 {
                v6[0] = int(v1)
            }
        }
    } else {
        v2 += 1
        v76 := GenS10{F0: 2}
        _ = v76
    }
    var v77 bool = false
    if (255 == 2) {
        if (491 > 1) {
            var v78 int16 = int16(v1)
            _ = v78
        } else {
            v4 -= 255
            fmt.Println(v77, v77, v5)
            var v79 uint8 = 255
            _ = v79
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch 51 {
        case 32767:
            v80 := GenS7{F0: -1}
            _ = v80
        case -9223372036854775808:
            switch "hello" {
            case "foo bar":
                var v81 uint64 = 995
                _ = v81
            case "":
                v1 -= 255
            case "a":
                if (v3 < (1e10 - (3.14159 / 1e10))) {
                    var v82 uint8 = 255
                    fmt.Println(v4, v11)
                    if (!v77) {
                        v83 := GenS11{F0: "foo bar", F1: 255}
                        for fi84 := 0; fi84 < 20; fi84++ {
                            v11 += ((float64(0.0) + v3) / (float64(-2.25) + (1.5 + (v3 + 0.0))))
                            _ = fi84
                        }
                        _ = v83
                    }
                    _ = v82
                }
            }
        case -2147483648:
            v11 = (float64(-2.25) * 0.0)
        default:
            v6 = append(v6, 42)
        }
    } else {
        var v85 uint8 = 255
        switch 977 {
        case 9223372036854775807:
            fmt.Println(v4, v77)
        default:
            for fi86 := 0; fi86 < 10; fi86++ {
                v87 := GenS10{F0: 2}
                _ = fi86
                _ = v87
            }
        }
        v77 = (((v5 + v5) > v5) || (true || (v77 && (!false))))
        _ = v85
    }
    for fi88 := 0; fi88 < 16; fi88++ {
        var v89 bool = ((65535 < 1) && v77)
        _ = fi88
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}
