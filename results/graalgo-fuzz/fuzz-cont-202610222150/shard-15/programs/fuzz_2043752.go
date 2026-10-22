package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 uint; F1 bool; F2 uint64 }
type GenS10 struct { F0 int; F1 int; F2 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 int; F1 uint; F2 uint }
type GenS14 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS15 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS16 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint64; F1 int8 }
func genFn0() uint {
    var v4 uint8 = 255
    v4 += (v4 >> ((209) & 63))
    v5 := GenS1{F0: 2, F1: 1000}
    _ = v4
    _ = v5
    return 418
}
func genFn1(p0 uint64) uint {
    if true {
        p0 -= 18446744073709551615
        v6 := GenS1{F0: 0, F1: 1}
        for fi7 := 0; fi7 < 3; fi7++ {
            p0 = p0
            _ = fi7
        }
        _ = v6
    }
    var v8 int16 = -189
    switch 87 {
    case 1:
        var v9 int64 = 9223372036854775807
        _ = v9
    case 0:
        v8 = 1000
    case -648:
        switch 32767 {
        case -429:
            var v10 uint8 = 98
            _ = v10
        }
    }
    v11 := []int{0, 127, 32767}
    _ = p0
    _ = v11
    _ = v8
    return 255
}

func main() {
    v1 := []string{(((("a" + "hello") + "hello") + ("foo bar" + "\n\t")) + ""), (("" + ((("\n\t" + ("\n\t" + "\n\t")) + "hello") + "a")) + "a")}
    var v2 uint32 = 1
    var v3 string = "a"
    var v12 uint64 = 18446744073709551615
    switch -9223372036854775808 {
    case 128:
        if len(v1) > 0 {
            v1[0] = v3
        }
    case -2147483648:
        var v13 uint8 = 14
        _ = v13
    case 42:
        var v14 bool = (false || true)
        _ = v14
    default:
        var v15 int64 = 0
        _ = v15
    }
    fmt.Println(v3)
    if false {
        for fi16 := 0; fi16 < 1; fi16++ {
            v3 = "a"
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                if (255 <= 278) {
                    if false {
                        var v18 int = 354
                        if (((v12 - (^(v12))) < v12) || (1000 <= 32767)) {
                            switch "a" {
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = v3
                                }
                                _ = len(v1)
                            case "a":
                                v18 = (fi17 % ((fi16) | 1))
                            case "":
                                v3 = "\n\t"
                            }
                        }
                        v2 += (^(v2))
                        _ = v18
                    } else {
                        var v19 string = v3
                        v12 += (v12 - (v12 & uint64(v2)))
                        switch (fi16 - (fi17 * -659)) {
                        case -9223372036854775808:
                            fmt.Println(fi16, fi16, v2)
                        case 680:
                            v20 := GenS1{F0: 1, F1: 32767}
                            _ = v20
                        }
                        _ = v19
                    }
                } else {
                    v12 += v12
                    if false {
                        fmt.Println(fi16, fi17, v2)
                    }
                }
                _ = len(v1)
                _ = fi17
            }
            _ = fi16
        }
        v1 = append(v1, v3)
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) > 0 {
            v1[0] = v3
        }
        v1 = append(v1, (v3 + ("foo bar" + (v3 + v3))))
        _ = fi21
    }
    if (!false) {
        if true {
            v2 -= (v2 & v2)
        } else {
            v22 := GenS1{F0: genFn0(), F1: -1}
            _ = len(v1)
            for fi23 := 0; fi23 < 5; fi23++ {
                v1 = append(v1, v3)
                switch (((v3 + "") + ((v3 + "foo bar") + "")) + v3) {
                case "hello":
                    v24 := GenS2{F0: (("foo bar" + v3) > "a")}
                    _ = v24
                case "foo bar":
                    v12 = v12
                case "\n\t":
                    v3 = ("\n\t" + v3)
                default:
                    v1 = append(v1, ((v3 + v3) + v3))
                }
                _ = fi23
            }
            _ = v22
        }
        fmt.Println(v12, v2)
        if (0 <= -75) {
            switch 127 {
            case 4294967295:
                if (255 > 127) {
                    switch (v3 + v3) {
                    case "\n\t":
                        switch "a" {
                        case "a":
                            switch "foo bar" {
                            case "":
                                v25 := GenS3{F0: "hello", F1: 437}
                                _ = v25
                            case "hello":
                                fmt.Println(v2, v3)
                            case "a":
                                if true {
                                    if (true && ((0.0 + 1.5) != ((float32(float64(0.0)) * (-2.25 / 0.1)) * ((1e10 + (float32((3.14159 * (3.14159 - float64(float32(0.1))))) / -2.25)) - -2.25)))) {
                                        var v26 int = -803
                                        _ = v26
                                    }
                                    switch v3 {
                                    case "\n\t":
                                        _ = len(v1)
                                    default:
                                        for fi27 := 0; fi27 < 7; fi27++ {
                                            v2 = (v2 | v2)
                                            v28 := GenS4{F0: 65535, F1: (255 == 80)}
                                            _ = fi27
                                            _ = v28
                                        }
                                    }
                                    var v29 int16 = -372
                                    _ = v29
                                } else {
                                    if (!true) {
                                        if true {
                                            v2 = 42
                                            var v30 uint64 = 858
                                            _ = v30
                                        } else {
                                            var v31 int64 = 9223372036854775807
                                            _ = v31
                                        }
                                        v12 -= (v12 | 42)
                                        switch v3 {
                                        case "":
                                            v2 -= (v2 / ((v2) | 1))
                                        case "foo bar":
                                            v3 = ((v3 + v3) + "")
                                        }
                                    }
                                    if true {
                                        var v32 uint = 497
                                        v3 = (v3 + v3)
                                        _ = v32
                                    }
                                }
                            }
                        case "foo bar":
                            var v33 int = 956
                            _ = v33
                        case "hello":
                            switch v3 {
                            case "foo bar":
                                v2 = (v2 - (v2 | (v2 / ((720) | 1))))
                            case "hello":
                                _ = len(v1)
                            }
                        }
                    default:
                        var v34 int = 202
                        _ = v34
                    }
                }
            case -9223372036854775808:
                v3 = ((v3 + v3) + v3)
            case 127:
                if (!true) {
                    switch 0 {
                    case 215:
                        v12 -= (v12 ^ v12)
                    case 377:
                        if true {
                            v12 -= v12
                        } else {
                            for fi35 := 0; fi35 < 10; fi35++ {
                                var v36 int64 = 42
                                v37 := GenS1{F0: 42, F1: -32768}
                                v3 = v3
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            switch int(v2) {
                            case 147:
                                if (v12 < 328) {
                                    fmt.Println(v12, v2)
                                    if (4294967295 >= (^(v12))) {
                                        if true {
                                            var v38 uint8 = 0
                                            if (-32768 < 1000) {
                                                if false {
                                                    switch -2147483648 {
                                                    case 32768:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case -2147483648:
                                                        if (0 == -1) {
                                                            if false {
                                                                var v39 uint = 255
                                                                switch 249 {
                                                                case -985:
                                                                    for fi40 := 0; fi40 < 7; fi40++ {
                                                                        _ = len(v1)
                                                                        var v41 uint16 = 65535
                                                                        if false {
                                                                            for fi42 := 0; fi42 < 8; fi42++ {
                                                                                switch fi42 {
                                                                                case 4294967295:
                                                                                    switch (-(fi40)) {
                                                                                    case -1:
                                                                                        v2 -= v2
                                                                                    case -115:
                                                                                        v43 := GenS5{F0: -239}
                                                                                        _ = v43
                                                                                    case 2147483647:
                                                                                        fmt.Println(v3, v2)
                                                                                    }
                                                                                    var v44 uint16 = (v41 / ((v41) | 1))
                                                                                    _ = v44
                                                                                default:
                                                                                    fmt.Println(fi42, v38, fi42)
                                                                                    switch fi40 {
                                                                                    case 4294967295:
                                                                                        v2 += v2
                                                                                    case 803:
                                                                                        _ = len(v1)
                                                                                    case 1:
                                                                                        v41 += v41
                                                                                        v41 += (v41 - v41)
                                                                                    }
                                                                                }
                                                                                v45 := GenS6{F0: v41, F1: fi42}
                                                                                v46 := GenS7{F0: (v38 * 0), F1: 127}
                                                                                _ = fi42
                                                                                _ = v45
                                                                                _ = v46
                                                                            }
                                                                            switch (fi40 - fi40) {
                                                                            case -37:
                                                                                for fi47 := 0; fi47 < 3; fi47++ {
                                                                                    v12 -= 4294967295
                                                                                    switch (fi47 & fi47) {
                                                                                    case -2147483648:
                                                                                        v48 := GenS1{F0: (v39 % (((v39 >> ((genFn0()) & 63))) | 1)), F1: -1}
                                                                                        _ = v48
                                                                                    case -60:
                                                                                        if false {
                                                                                            if true {
                                                                                                fmt.Println(fi40)
                                                                                            }
                                                                                            var v49 float64 = 0.0
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            _ = v49
                                                                                        } else {
                                                                                            fmt.Println(v41, fi40, fi47)
                                                                                            fmt.Println(v38, fi47)
                                                                                            fmt.Println(v2, v39, v3)
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v12, v39)
                                                                                    }
                                                                                    _ = fi47
                                                                                }
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v38, v12)
                                                                            default:
                                                                                if (v2 < 228) {
                                                                                    fmt.Println(v41, fi40, v39)
                                                                                    for fi50 := 0; fi50 < 14; fi50++ {
                                                                                        var v51 float32 = -2.25
                                                                                        v52 := GenS8{F0: 1000, F1: 19}
                                                                                        v2 = (v2 + (^(v2)))
                                                                                        _ = fi50
                                                                                        _ = v51
                                                                                        _ = v52
                                                                                    }
                                                                                } else {
                                                                                    v53 := GenS1{F0: uint(v12), F1: 1000}
                                                                                    fmt.Println(v3, fi40)
                                                                                    _ = v53
                                                                                }
                                                                            }
                                                                            v2 = (^(v2))
                                                                        } else {
                                                                            for fi54 := 0; fi54 < 20; fi54++ {
                                                                                var v55 float64 = (float64(1.5) / (((1e10 + (((0.1 / 1.5) - 3.14159) + 3.14159)) - 1e10) + ((-2.25 * (1.5 + (3.14159 - 0.1))) - float64(3.14159))))
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if (42 > 2147483647) {
                                                                                    var v56 int = -403
                                                                                    var v57 uint8 = 255
                                                                                    v58 := GenS1{F0: 1, F1: 32767}
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                                _ = fi54
                                                                                _ = v55
                                                                            }
                                                                            var v59 float64 = 3.14159
                                                                            _ = v59
                                                                        }
                                                                        _ = fi40
                                                                        _ = v41
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v60 uint8 = (v38 >> ((uint(v38)) & 63))
                                                                    _ = v60
                                                                case -647:
                                                                    var v61 uint8 = v38
                                                                    _ = v61
                                                                default:
                                                                    for fi62 := 0; fi62 < 20; fi62++ {
                                                                        for fi63 := 0; fi63 < 11; fi63++ {
                                                                            fmt.Println(fi63, v2, v39)
                                                                            fmt.Println(v12, v3)
                                                                            var v64 int8 = 42
                                                                            var v65 uint64 = (^(v12))
                                                                            v64 += v64
                                                                            _ = fi63
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                        if true {
                                                                            for fi66 := 0; fi66 < 7; fi66++ {
                                                                                v67 := GenS4{F0: 4294967295, F1: false}
                                                                                _ = fi66
                                                                                _ = v67
                                                                            }
                                                                            fmt.Println(v12, v39)
                                                                        }
                                                                        for fi68 := 0; fi68 < 14; fi68++ {
                                                                            v69 := GenS8{F0: -32768, F1: -1}
                                                                            _ = fi68
                                                                            _ = v69
                                                                        }
                                                                        _ = fi62
                                                                    }
                                                                }
                                                                _ = v39
                                                            }
                                                            var v70 float32 = (1e10 + (0.0 * 1.5))
                                                            _ = v70
                                                        }
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = "hello"
                                                    }
                                                    fmt.Println(v2, v3, v3)
                                                }
                                                for fi71 := 0; fi71 < 14; fi71++ {
                                                    for fi72 := 0; fi72 < 15; fi72++ {
                                                        switch fi72 {
                                                        case 269:
                                                            v73 := GenS9{F0: 545, F1: ((0.0 - 1.5) > (3.14159 / float32(3.14159))), F2: v12}
                                                            _ = v73
                                                        case 0:
                                                            fmt.Println(v38)
                                                        case 2147483647:
                                                            var v74 int64 = 161
                                                            _ = v74
                                                        default:
                                                            v75 := GenS10{F0: (fi72 << ((167) & 63)), F1: fi71, F2: 2}
                                                            _ = v75
                                                        }
                                                        v2 += (v2 << ((233) & 63))
                                                        if false {
                                                            var v76 uint64 = v12
                                                            switch "foo bar" {
                                                            case "a":
                                                                if (255 == 204) {
                                                                    fmt.Println(v3, v76)
                                                                    if false {
                                                                        if (true || false) {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v77 float32 = 0.1
                                                                            _ = v77
                                                                        } else {
                                                                            fmt.Println(fi71, fi72, v2)
                                                                            if false {
                                                                                fmt.Println(fi71, v38, v12)
                                                                                v3 = "hello"
                                                                                switch 42 {
                                                                                case 32767:
                                                                                    switch v3 {
                                                                                    case "hello":
                                                                                        v38 = (v38 * v38)
                                                                                    case "":
                                                                                        _ = len(v1)
                                                                                    case "a":
                                                                                        _ = len(v1)
                                                                                    }
                                                                                default:
                                                                                    switch v3 {
                                                                                    case "a":
                                                                                        _ = len(v1)
                                                                                    case "":
                                                                                        switch fi72 {
                                                                                        case 24:
                                                                                            if (465 >= 18446744073709551615) {
                                                                                                v38 += (v38 / ((163) | 1))
                                                                                            } else {
                                                                                                _ = len(v1)
                                                                                            }
                                                                                        case 1:
                                                                                            if false {
                                                                                                var v78 int = fi72
                                                                                                v12 = (v76 / ((v76) | 1))
                                                                                                _ = v78
                                                                                            } else {
                                                                                                v76 += v12
                                                                                                v79 := GenS5{F0: 0}
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = (v3 + (((v3 + "a") + "hello") + ("foo bar" + ((v3 + "hello") + (("" + v3) + "foo bar")))))
                                                                                                }
                                                                                                _ = v79
                                                                                            }
                                                                                        case -31:
                                                                                            fmt.Println(v2, v2, v12)
                                                                                        }
                                                                                    default:
                                                                                        v12 = v12
                                                                                    }
                                                                                }
                                                                            }
                                                                            v38 = v38
                                                                            if (-462 < 56) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v3
                                                                                }
                                                                            } else {
                                                                                v2 = v2
                                                                            }
                                                                            var v80 uint8 = v38
                                                                            _ = v80
                                                                        }
                                                                        if ((-2147483648 < 560) || (42 < 32767)) {
                                                                            v81 := GenS2{F0: ("foo bar" > (((v3 + ("hello" + "hello")) + v3) + "hello"))}
                                                                            v1 = append(v1, (v3 + v3))
                                                                            _ = v81
                                                                        } else {
                                                                            switch v3 {
                                                                            case "":
                                                                                v3 = v3
                                                                            }
                                                                            fmt.Println(v76, v76)
                                                                            var v82 float64 = (0.0 + (3.14159 / 1e10))
                                                                            _ = v82
                                                                        }
                                                                    } else {
                                                                        switch fi72 {
                                                                        case 32767:
                                                                            if true {
                                                                                v76 = (v76 + (v12 + v12))
                                                                                var v83 float32 = (1e10 - ((3.14159 / -2.25) / 1e10))
                                                                                switch fi72 {
                                                                                case 32768:
                                                                                    v76 -= v76
                                                                                    fmt.Println(v83, fi72, v12)
                                                                                    v3 = "foo bar"
                                                                                case 32767:
                                                                                    var v84 uint = uint(v76)
                                                                                    _ = v84
                                                                                default:
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v3
                                                                                    }
                                                                                }
                                                                                v85 := GenS11{F0: -2147483648}
                                                                                _ = v83
                                                                                _ = v85
                                                                            }
                                                                        case 42:
                                                                            if (!(float32((1.5 + 1.5)) != (0.0 * float32(0.1)))) {
                                                                                if ((2 <= uint8(v2)) || ((1e10 >= (0.1 / ((0.0 - 1e10) - 1e10))) || true)) {
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        switch -1 {
                                                                                        case -9223372036854775808:
                                                                                            var v86 int16 = int16(v38)
                                                                                            _ = v86
                                                                                        case 4294967295:
                                                                                            v12 -= (v76 * (v12 | (v12 + (v12 | (v76 ^ v12)))))
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    if (-344 <= -1) {
                                                                                        v87 := GenS9{F0: genFn1((v12 << ((831) & 63))), F1: (!((!(!true)) && true)), F2: uint64(v2)}
                                                                                        _ = v87
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    v88 := GenS1{F0: 985, F1: 0}
                                                                                    _ = v88
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi72)
                                                                            v38 = uint8(v76)
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            v89 := GenS1{F0: 4294967295, F1: 32767}
                                                                            v76 = 362
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                v90 := GenS12{F0: -100, F1: (v2 & v2)}
                                                                _ = v90
                                                            }
                                                            _ = v76
                                                        } else {
                                                            var v91 int16 = 1000
                                                            _ = v91
                                                        }
                                                        _ = fi72
                                                    }
                                                    v38 = v38
                                                    var v92 int8 = -128
                                                    _ = fi71
                                                    _ = v92
                                                }
                                            } else {
                                                for fi93 := 0; fi93 < 1; fi93++ {
                                                    if false {
                                                        _ = len(v1)
                                                    }
                                                    v2 += v2
                                                    _ = fi93
                                                }
                                                var v94 int32 = -671
                                                for fi95 := 0; fi95 < 8; fi95++ {
                                                    for fi96 := 0; fi96 < 13; fi96++ {
                                                        fmt.Println(fi96, fi95)
                                                        if false {
                                                            if false {
                                                                v2 += v2
                                                                v38 = (v38 << ((96) & 63))
                                                                var v97 string = "foo bar"
                                                                _ = v97
                                                            }
                                                            v3 = "\n\t"
                                                        } else {
                                                            var v98 uint8 = v38
                                                            v99 := GenS13{F0: fi95, F1: 381, F2: 800}
                                                            var v100 uint64 = (v12 >> ((4294967295) & 63))
                                                            _ = v100
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                        switch 1 {
                                                        case 32768:
                                                            _ = len(v1)
                                                        case -2147483648:
                                                            if (1e10 == (float32(-2.25) * 3.14159)) {
                                                                v101 := GenS6{F0: 682, F1: fi95}
                                                                _ = v101
                                                            } else {
                                                                v1 = append(v1, v3)
                                                                switch (fi95 | fi95) {
                                                                case 9223372036854775807:
                                                                    var v102 uint16 = 679
                                                                    _ = v102
                                                                }
                                                            }
                                                        case 0:
                                                            var v103 uint16 = 0
                                                            _ = v103
                                                        }
                                                        _ = fi96
                                                    }
                                                    if (!(!false)) {
                                                        var v104 uint32 = v2
                                                        _ = v104
                                                    } else {
                                                        if true {
                                                            v12 += 971
                                                            v105 := GenS14{F0: (v12 % ((929) | 1)), F1: 0, F2: 4294967295}
                                                            for fi106 := 0; fi106 < 16; fi106++ {
                                                                if true {
                                                                    var v107 uint = 702
                                                                    if ("hello" != v3) {
                                                                        var v108 int = -733
                                                                        _ = v108
                                                                    }
                                                                    _ = v107
                                                                }
                                                                _ = fi106
                                                            }
                                                            _ = v105
                                                        }
                                                    }
                                                    _ = fi95
                                                }
                                                _ = v94
                                            }
                                            _ = v38
                                        }
                                        v109 := GenS13{F0: 4294967295, F1: 42, F2: 255}
                                        _ = v109
                                    } else {
                                        if false {
                                            var v110 float64 = 0.0
                                            var v111 uint32 = (^(v2))
                                            _ = v110
                                            _ = v111
                                        } else {
                                            var v112 uint = 2
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v112
                                        }
                                        _ = len(v1)
                                        v1 = append(v1, "a")
                                    }
                                }
                            case 0:
                                for fi113 := 0; fi113 < 3; fi113++ {
                                    var v114 bool = (0 != 42)
                                    switch v3 {
                                    case "hello":
                                        var v115 float64 = 0.1
                                        _ = v115
                                    case "\n\t":
                                        for fi116 := 0; fi116 < 15; fi116++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch (fi113 - fi113) {
                                            case 32768:
                                                v117 := GenS9{F0: 0, F1: (1 != -808), F2: v12}
                                                _ = v117
                                            default:
                                                v12 -= 4294967295
                                            }
                                            _ = fi116
                                        }
                                    case "":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = fi113
                                    _ = v114
                                }
                            case 700:
                                var v118 bool = (65535 < 255)
                                _ = v118
                            }
                            _ = len(v1)
                        }
                    case -584:
                        v3 = v3
                    }
                    v12 -= (v12 / (((v12 * 1)) | 1))
                    v119 := GenS15{F0: (v2 / ((v2) | 1)), F1: genFn0(), F2: 127}
                    _ = v119
                }
            }
            var v120 string = ((v3 + v3) + "\n\t")
            var v121 string = v3
            _ = v120
            _ = v121
        } else {
            v12 = v12
            for fi122 := 0; fi122 < 2; fi122++ {
                switch fi122 {
                case 42:
                    v12 -= (v12 - (^(v12)))
                }
                _ = fi122
            }
        }
    }
    switch -1 {
    case 9223372036854775807:
        if true {
            v12 -= (v12 + (v12 << ((genFn1((v12 ^ (v12 ^ v12)))) & 63)))
        }
    default:
        for fi123 := 0; fi123 < 17; fi123++ {
            var v124 bool = true
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if v124 {
                var v125 uint16 = 42
                _ = v125
            } else {
                v124 = ("a" > ("foo bar" + ("\n\t" + (((v3 + (("" + v3) + v3)) + v3) + "a"))))
                switch v3 {
                case "hello":
                    v124 = false
                }
            }
            _ = fi123
            _ = v124
        }
    }
    v12 = (^(v12))
    if len(v1) > 0 {
        v1[0] = ((((v3 + v3) + v3) + "hello") + v3)
    }
    for fi126 := 0; fi126 < 20; fi126++ {
        for fi127 := 0; fi127 < 13; fi127++ {
            if true {
                v128 := GenS16{F0: 742, F1: (v2 * (v2 % ((v2) | 1))), F2: -112}
                var v129 uint64 = 475
                switch v3 {
                case "a":
                    v129 += v12
                case "":
                    var v130 string = v3
                    _ = v130
                case "foo bar":
                    if ((v2 / ((v2) | 1)) == v2) {
                        if false {
                            v129 = v129
                        } else {
                            fmt.Println(v2, v3, v12)
                        }
                    } else {
                        v3 = v3
                        switch ("foo bar" + (v3 + v3)) {
                        case "":
                            v131 := GenS12{F0: -2147483648, F1: 2}
                            _ = v131
                        case "foo bar":
                            var v132 int16 = 32767
                            _ = v132
                        case "\n\t":
                            switch (fi127 / ((fi126) | 1)) {
                            case -527:
                                v12 = v12
                            case -2147483648:
                                fmt.Println(fi126, v129, v12)
                            case 934:
                                var v133 string = "\n\t"
                                _ = v133
                            default:
                                var v134 uint64 = v129
                                _ = v134
                            }
                        default:
                            var v135 int = (^(fi126))
                            _ = v135
                        }
                    }
                }
                _ = v128
                _ = v129
            }
            _ = fi127
        }
        _ = fi126
    }
    _ = len(v1)
    v3 = "foo bar"
    for fi136 := 0; fi136 < 15; fi136++ {
        if (!false) {
            var v137 float32 = 1.5
            _ = v137
        }
        _ = fi136
    }
    for fi138 := 0; fi138 < 3; fi138++ {
        v139 := GenS17{F0: (false || true)}
        var v140 int16 = 552
        var v141 float32 = (-2.25 + 0.0)
        _ = fi138
        _ = v139
        _ = v140
        _ = v141
    }
    v142 := GenS18{F0: (v12 * (v12 >> ((56) & 63))), F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v142
    _ = v2
    _ = v3
}
