package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint; F2 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 string, p1 int8, p2 int32) uint8 {
    p1 = p1
    v7 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 275
    v3 := []uint32{2, 2, 298, 659}
    var v4 uint8 = 255
    v5 := []int{-125, 128, 32767}
    v6 := []uint32{65535, 2147483648, 2147483647, 2}
    var v8 float32 = (0.0 - ((3.14159 - (-2.25 / 0.1)) * 3.14159))
    for fi9 := 0; fi9 < 4; fi9++ {
        switch "\n\t" {
        case "a":
            switch (fi9 / ((-2147483648) | 1)) {
            case 1:
                var v10 bool = (0 >= 1)
                _ = v10
            }
        default:
            for fi11 := 0; fi11 < 16; fi11++ {
                v2 = 791
                _ = fi11
            }
        }
        switch (-(fi9)) {
        case -2147483648:
            for fi12 := 0; fi12 < 11; fi12++ {
                switch 0 {
                case 0:
                    if true {
                        v8 = (v8 + -2.25)
                        if ((2 != 848) || true) {
                            v13 := v1[("a" + "hello")]
                            v13 = "hello"
                            _ = v13
                        } else {
                            switch 2147483647 {
                            case 128:
                                switch fi9 {
                                case -926:
                                    if len(v6) > 0 {
                                        v6[0] = 42
                                    }
                                case 521:
                                    if false {
                                        var v14 uint32 = 4294967295
                                        _ = v14
                                    }
                                default:
                                    v15 := GenS1{F0: "a"}
                                    _ = v15
                                }
                            case 10:
                                if true {
                                    var v16 float64 = 0.0
                                    v16 += (v16 - v16)
                                    switch fi9 {
                                    case 265:
                                        if false {
                                            var v17 int32 = 1
                                            _ = v17
                                        } else {
                                            fmt.Println(v2, v4, v16)
                                            var v18 int = (fi12 ^ fi9)
                                            delete(v1, ("\n\t" + "\n\t"))
                                            _ = v18
                                        }
                                    }
                                    _ = v16
                                }
                            case 1:
                                if (128 == (fi9 >> ((42) & 63))) {
                                    v19 := GenS1{F0: "foo bar"}
                                    v1[("hello" + "")] = "foo bar"
                                    _ = v19
                                } else {
                                    v8 -= 1e10
                                    fmt.Println(v4)
                                    v8 = float32(1.5)
                                }
                            default:
                                v20 := GenS1{F0: ""}
                                _ = v20
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v21 := GenS2{F0: 2}
                            _ = v21
                        }
                    } else {
                        var v22 int = (fi12 - 32768)
                        v22 = fi12
                        v4 -= (v4 + v4)
                        _ = v22
                    }
                    v8 += float32(0.0)
                    v1["hello"] = "foo bar"
                    fmt.Println(v4, v8, v4)
                    v4 = (v4 / (((v4 - (v4 / (((v4 / ((v4) | 1))) | 1)))) | 1))
                case 127:
                    v8 = (v8 + (v8 - v8))
                }
                if true {
                    var v23 float32 = -2.25
                    if true {
                        v24 := GenS2{F0: v2}
                        _ = len(v3)
                        _ = v24
                    } else {
                        v2 -= (^(v2))
                        v2 = 2
                        var v25 uint32 = 2147483648
                        _ = v25
                    }
                    _ = v23
                } else {
                    v2 = v2
                }
                _ = fi12
            }
        default:
            delete(v1, "")
        }
        v1["foo bar"] = "\n\t"
        _ = fi9
    }
    if (v4 <= 2) {
        if false {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v4 = genFn0("foo bar", 5, -416)
            switch 950 {
            case 9223372036854775807:
                v26 := GenS3{F0: 100, F1: -536, F2: 42}
                _ = v26
            case 421:
                var v27 int8 = 100
                _ = v27
            }
        }
    }
    v4 -= v4
    var v28 int8 = -1
    if (!true) {
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 uint64 = v2
            _ = fi29
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 int16 = 0
            v33 := v1["a"]
            v32 += 800
            _ = fi31
            _ = v32
            _ = v33
        }
        if ("hello" == ("a" + "foo bar")) {
            v34 := v1["foo bar"]
            var v35 int32 = 0
            var v36 uint16 = 255
            _ = v34
            _ = v35
            _ = v36
        } else {
            v8 = 1.5
            _ = len(v6)
        }
        v1["a"] = ""
    }
    var v37 string = "a"
    switch -126 {
    case 0:
        fmt.Println(v37, v4)
    case 2147483647:
        v38 := GenS3{F0: v28, F1: 32767, F2: v28}
        _ = v38
    case 128:
        switch ((v37 + v37) + ((v37 + "foo bar") + ((((("" + v37) + v37) + "\n\t") + "a") + v37))) {
        case "a":
            _ = len(v1)
        case "\n\t":
            if false {
                v39 := GenS1{F0: (v37 + ("foo bar" + ("\n\t" + v37)))}
                var v40 uint8 = genFn0(v37, (v28 & (v28 - v28)), -2147483648)
                _ = v39
                _ = v40
            }
        default:
            switch 4294967295 {
            case -1:
                if false {
                    for fi41 := 0; fi41 < 4; fi41++ {
                        if true {
                            fmt.Println(fi41)
                            if false {
                                if true {
                                    for fi42 := 0; fi42 < 8; fi42++ {
                                        fmt.Println(fi41, v4, v4)
                                        _ = fi42
                                    }
                                    if true {
                                        for fi43 := 0; fi43 < 13; fi43++ {
                                            v28 = v28
                                            _ = fi43
                                        }
                                    } else {
                                        if false {
                                            var v44 int8 = v28
                                            v45 := GenS4{F0: 18446744073709551615}
                                            var v46 int16 = -583
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        } else {
                                            if ((1e10 - v8) <= v8) {
                                                var v47 int = -538
                                                switch v37 {
                                                case "hello":
                                                    v2 = v2
                                                case "foo bar":
                                                    switch v47 {
                                                    case -448:
                                                        v48 := GenS5{F0: (^(v4)), F1: 413, F2: 42}
                                                        _ = v48
                                                    default:
                                                        fmt.Println(v47)
                                                    }
                                                case "a":
                                                    switch v37 {
                                                    case "\n\t":
                                                        for fi49 := 0; fi49 < 8; fi49++ {
                                                            var v50 uint64 = v2
                                                            v51 := GenS1{F0: "hello"}
                                                            _ = fi49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    case "hello":
                                                        var v52 int16 = -849
                                                        _ = v52
                                                    default:
                                                        v53 := GenS4{F0: 18446744073709551615}
                                                        _ = v53
                                                    }
                                                default:
                                                    var v54 uint = 1
                                                    _ = v54
                                                }
                                                _ = v47
                                            } else {
                                                v2 -= (v2 << ((195) & 63))
                                                var v55 int = (^(fi41))
                                                v4 += v4
                                                _ = v55
                                            }
                                            v56 := GenS1{F0: "hello"}
                                            _ = v56
                                        }
                                        var v57 uint16 = 32767
                                        for fi58 := 0; fi58 < 17; fi58++ {
                                            if (((false && true) && (673 < 2147483647)) || true) {
                                                fmt.Println(v57, fi41, v37)
                                                switch fi41 {
                                                case -9223372036854775808:
                                                    fmt.Println(v28, v8)
                                                case 0:
                                                    fmt.Println(v28, v4)
                                                case 518:
                                                    if (0.0 != -2.25) {
                                                        v8 += ((v8 + 0.0) - ((((v8 * (v8 - (-2.25 - v8))) - ((v8 - ((1.5 + v8) * (v8 + (float32(1e10) / (v8 - v8))))) * v8)) + 0.0) * (((v8 / v8) + 1.5) + v8)))
                                                        var v59 uint64 = v2
                                                        v8 -= (v8 + v8)
                                                        _ = v59
                                                    } else {
                                                        var v60 float32 = -2.25
                                                        if false {
                                                            v37 = "a"
                                                            _ = len(v1)
                                                        } else {
                                                            var v61 int8 = int8(v2)
                                                            if false {
                                                                if true {
                                                                    var v62 int16 = 1
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v63 string = v37
                                                                        _ = v63
                                                                    case "\n\t":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    case "":
                                                                        var v64 uint32 = 4294967295
                                                                        _ = v64
                                                                    }
                                                                    if false {
                                                                        if true {
                                                                            if (!false) {
                                                                                var v65 uint8 = (v4 >> ((uint(v57)) & 63))
                                                                                _ = v65
                                                                            }
                                                                            v60 = v60
                                                                            var v66 uint8 = genFn0(v37, 60, 42)
                                                                            _ = v66
                                                                        }
                                                                        switch (v37 + ("foo bar" + ("a" + ("\n\t" + v37)))) {
                                                                        case "foo bar":
                                                                            switch (fi58 / (((-(fi58))) | 1)) {
                                                                            case 32767:
                                                                                var v67 uint64 = 528
                                                                                v60 = float32((3.14159 / -2.25))
                                                                                _ = v67
                                                                            case 383:
                                                                                v68 := GenS3{F0: v61, F1: 32768, F2: (v61 % ((100) | 1))}
                                                                                _ = v68
                                                                            case -9223372036854775808:
                                                                                fmt.Println(fi58, v61, v60)
                                                                            default:
                                                                                switch fi41 {
                                                                                case 1:
                                                                                    if false {
                                                                                        var v69 float32 = (v60 - (float32(1.5) - (v8 + (0.0 + v8))))
                                                                                        v4 -= v4
                                                                                        switch (fi58 & (fi41 >> ((0) & 63))) {
                                                                                        case -2147483648:
                                                                                            v61 = 0
                                                                                        }
                                                                                        _ = v69
                                                                                    }
                                                                                case -1:
                                                                                    var v70 int64 = -151
                                                                                    _ = v70
                                                                                case 515:
                                                                                    _ = len(v3)
                                                                                }
                                                                                var v71 uint16 = (v57 % (((v57 % ((v57) | 1))) | 1))
                                                                                fmt.Println(fi41, v62, v60)
                                                                                _ = v71
                                                                            }
                                                                        case "hello":
                                                                            switch (("hello" + v37) + v37) {
                                                                            case "hello":
                                                                                var v72 uint16 = v57
                                                                                _ = v72
                                                                            case "\n\t":
                                                                                v73 := GenS2{F0: v2}
                                                                                _ = v73
                                                                            }
                                                                        default:
                                                                            fmt.Println(v62, v57, v8)
                                                                        }
                                                                        if true {
                                                                            v1[(v37 + "\n\t")] = ((v37 + "foo bar") + (("hello" + v37) + (v37 + v37)))
                                                                            var v74 int64 = -970
                                                                            var v75 uint16 = uint16(v2)
                                                                            v28 = (v61 << ((4294967295) & 63))
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                    } else {
                                                                        _ = len(v1)
                                                                        if (2 < v4) {
                                                                            if true {
                                                                                switch fi58 {
                                                                                case 42:
                                                                                    var v76 int32 = -494
                                                                                    _ = v76
                                                                                case 4294967295:
                                                                                    fmt.Println(fi41)
                                                                                    delete(v1, "a")
                                                                                case 2147483647:
                                                                                    var v77 string = (v37 + "")
                                                                                    _ = v77
                                                                                }
                                                                                v62 += int16(v57)
                                                                                _ = len(v1)
                                                                            }
                                                                            switch (("a" + "\n\t") + v37) {
                                                                            case "hello":
                                                                                if ((-(v62)) != v62) {
                                                                                    var v78 int16 = (v62 & (v62 * v62))
                                                                                    var v79 int8 = (^(v61))
                                                                                    if (!(278 > -442)) {
                                                                                        v3 = append(v3, 4294967295)
                                                                                        fmt.Println(fi58, v28, v28)
                                                                                        switch (fi41 ^ 32768) {
                                                                                        case 127:
                                                                                            if (!(0.0 < 1.5)) {
                                                                                                _ = len(v6)
                                                                                            } else {
                                                                                                if (v62 > v78) {
                                                                                                    v28 = v61
                                                                                                    if (!false) {
                                                                                                        switch fi41 {
                                                                                                        case 0:
                                                                                                            delete(v1, v37)
                                                                                                        default:
                                                                                                            v80 := GenS1{F0: ("foo bar" + ((("foo bar" + v37) + (v37 + v37)) + ""))}
                                                                                                            _ = v80
                                                                                                        }
                                                                                                        v81 := v1["a"]
                                                                                                        _ = v81
                                                                                                    }
                                                                                                    var v82 uint64 = 209
                                                                                                    _ = v82
                                                                                                } else {
                                                                                                    v78 = (v78 * v62)
                                                                                                }
                                                                                                fmt.Println(v78, v28, v4)
                                                                                                switch 2147483647 {
                                                                                                case 21:
                                                                                                    v60 += float32(0.1)
                                                                                                case -97:
                                                                                                    if (-1 > 42) {
                                                                                                        var v83 float32 = v8
                                                                                                        _ = v83
                                                                                                    } else {
                                                                                                        v28 -= v61
                                                                                                        delete(v1, v37)
                                                                                                        var v84 int32 = int32(v57)
                                                                                                        v60 -= v60
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    var v85 uint8 = (v4 / ((v4) | 1))
                                                                                                    _ = v85
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            switch v37 {
                                                                                            case "":
                                                                                                if true {
                                                                                                    var v86 int32 = 2147483647
                                                                                                    v28 = v79
                                                                                                    v87 := v1[("" + (v37 + v37))]
                                                                                                    _ = v86
                                                                                                    _ = v87
                                                                                                } else {
                                                                                                    v88 := GenS6{F0: (v28 << ((992) & 63))}
                                                                                                    var v89 int16 = (v78 + int16(v4))
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v28 = 42
                                                                                            default:
                                                                                                switch "" {
                                                                                                case "":
                                                                                                    var v90 int32 = -2147483648
                                                                                                    _ = v90
                                                                                                case "foo bar":
                                                                                                    v91 := GenS5{F0: (v4 >> ((0) & 63)), F1: 187, F2: -61}
                                                                                                    _ = v91
                                                                                                case "a":
                                                                                                    if (42 < -1) {
                                                                                                        switch "a" {
                                                                                                        case "foo bar":
                                                                                                            var v92 uint16 = 1
                                                                                                            _ = v92
                                                                                                        case "a":
                                                                                                            var v93 int = fi58
                                                                                                            _ = v93
                                                                                                        default:
                                                                                                            v2 = (^(v2))
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (true || (true || true)) {
                                                                                                            switch ("\n\t" + (v37 + v37)) {
                                                                                                            case "hello":
                                                                                                                v37 = "\n\t"
                                                                                                            case "":
                                                                                                                var v94 uint = 897
                                                                                                                _ = v94
                                                                                                            case "foo bar":
                                                                                                                if (42 > -9223372036854775808) {
                                                                                                                    switch (fi41 >> ((0) & 63)) {
                                                                                                                    case 482:
                                                                                                                        v2 = (v2 % ((v2) | 1))
                                                                                                                    }
                                                                                                                    v57 = uint16(v4)
                                                                                                                    var v95 string = v37
                                                                                                                    var v96 uint64 = (v2 >> ((65535) & 63))
                                                                                                                    _ = v95
                                                                                                                    _ = v96
                                                                                                                }
                                                                                                                switch ((v37 + "\n\t") + v37) {
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v57, v28)
                                                                                                                case "hello":
                                                                                                                    v2 -= v2
                                                                                                                default:
                                                                                                                    v97 := GenS6{F0: (v28 * 36)}
                                                                                                                    _ = v97
                                                                                                                }
                                                                                                                switch 42 {
                                                                                                                case 431:
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                case 32768:
                                                                                                                    switch v37 {
                                                                                                                    case "foo bar":
                                                                                                                        v98 := GenS4{F0: 4294967295}
                                                                                                                        _ = v98
                                                                                                                    default:
                                                                                                                        var v99 int16 = 581
                                                                                                                        _ = v99
                                                                                                                    }
                                                                                                                case 4294967295:
                                                                                                                    v4 = (v4 << ((65535) & 63))
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        var v100 bool = (1 <= -396)
                                                                                                                        if ((true && (false && false)) && (false || (!true))) {
                                                                                                                            var v101 uint16 = (v57 & (v57 / ((v57) | 1)))
                                                                                                                            if false {
                                                                                                                                v102 := GenS2{F0: (v2 - (v2 * uint64(v57)))}
                                                                                                                                switch (fi58 + 32767) {
                                                                                                                                case -9223372036854775808:
                                                                                                                                    switch (v37 + v37) {
                                                                                                                                    case "foo bar":
                                                                                                                                        v62 -= (v78 % (((v62 << ((255) & 63))) | 1))
                                                                                                                                        v60 -= (float32(3.14159) / (v8 - float32(0.1)))
                                                                                                                                        var v103 int32 = 238
                                                                                                                                        _ = v103
                                                                                                                                    }
                                                                                                                                case 2147483647:
                                                                                                                                    if (v100 && v100) {
                                                                                                                                        var v104 uint32 = uint32(v57)
                                                                                                                                        switch fi41 {
                                                                                                                                        case -2147483648:
                                                                                                                                            v105 := GenS6{F0: v28}
                                                                                                                                            _ = v105
                                                                                                                                        default:
                                                                                                                                            var v106 uint16 = uint16(v104)
                                                                                                                                            var v107 int = 42
                                                                                                                                            _ = v106
                                                                                                                                            _ = v107
                                                                                                                                        }
                                                                                                                                        _ = v104
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v108 := GenS3{F0: 20, F1: (fi41 | fi58), F2: v61}
                                                                                                                                    _ = v108
                                                                                                                                }
                                                                                                                                var v109 uint16 = v57
                                                                                                                                _ = v102
                                                                                                                                _ = v109
                                                                                                                            }
                                                                                                                            _ = v101
                                                                                                                        }
                                                                                                                        _ = v100
                                                                                                                    } else {
                                                                                                                        switch ("" + v37) {
                                                                                                                        case "hello":
                                                                                                                            v5 = append(v5, fi41)
                                                                                                                        case "\n\t":
                                                                                                                            if true {
                                                                                                                                if (-694 == 1) {
                                                                                                                                    if ((v60 >= (v8 - (((-2.25 - (v60 * (0.0 - v60))) / v60) + (v60 + v8)))) || false) {
                                                                                                                                        v5 = append(v5, (-(fi41)))
                                                                                                                                        if false {
                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                            }
                                                                                                                                            var v110 float32 = (-2.25 - v8)
                                                                                                                                            _ = v110
                                                                                                                                        } else {
                                                                                                                                            var v111 uint32 = 834
                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                v5[0] = fi58
                                                                                                                                            }
                                                                                                                                            fmt.Println(v37, fi41, v78)
                                                                                                                                            var v112 uint = uint(v2)
                                                                                                                                            _ = v111
                                                                                                                                            _ = v112
                                                                                                                                        }
                                                                                                                                        var v113 int8 = v61
                                                                                                                                        _ = v113
                                                                                                                                    }
                                                                                                                                    fmt.Println(v61)
                                                                                                                                    var v114 int = (fi58 / ((int(v57)) | 1))
                                                                                                                                    v62 -= (v78 / ((1) | 1))
                                                                                                                                    fmt.Println(v4)
                                                                                                                                    if len(v3) >= 2 {
                                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                                    }
                                                                                                                                    var v115 int8 = 100
                                                                                                                                    _ = v114
                                                                                                                                    _ = v115
                                                                                                                                } else {
                                                                                                                                    var v116 uint16 = uint16(v4)
                                                                                                                                    if ((true || ((v116 % ((1) | 1)) != v116)) && false) {
                                                                                                                                        if len(v6) > 0 {
                                                                                                                                            v6[0] = 65535
                                                                                                                                        }
                                                                                                                                        v117 := GenS5{F0: genFn0("\n\t", v79, 2147483647), F1: 755, F2: -471}
                                                                                                                                        _ = v117
                                                                                                                                    }
                                                                                                                                    fmt.Println(v37)
                                                                                                                                    _ = v116
                                                                                                                                }
                                                                                                                                v118 := GenS2{F0: (v2 - v2)}
                                                                                                                                _ = v118
                                                                                                                            } else {
                                                                                                                                v57 -= (v57 << ((168) & 63))
                                                                                                                                v2 -= (v2 + (^(v2)))
                                                                                                                                v119 := GenS7{F0: 65535, F1: 0}
                                                                                                                                if len(v6) >= 2 {
                                                                                                                                    v6 = v6[1:len(v6)]
                                                                                                                                }
                                                                                                                                var v120 float32 = v8
                                                                                                                                _ = v119
                                                                                                                                _ = v120
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            if (((-(v79)) != -128) && true) {
                                                                                                                                fmt.Println(fi41)
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            if ((v4 * (v4 - genFn0(v37, int8(v4), 1))) != v4) {
                                                                                                                                var v121 float32 = float32((3.14159 + (float64(v60) / 1e10)))
                                                                                                                                _ = v121
                                                                                                                            }
                                                                                                                        }
                                                                                                                        fmt.Println(v61, v37, v2)
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v4 -= genFn0(v37, v79, -1)
                                                                                                            }
                                                                                                        }
                                                                                                        switch v37 {
                                                                                                        case "hello":
                                                                                                            var v122 string = ((v37 + v37) + v37)
                                                                                                            _ = v122
                                                                                                        case "a":
                                                                                                            var v123 uint16 = v57
                                                                                                            _ = v123
                                                                                                        default:
                                                                                                            var v124 float32 = 3.14159
                                                                                                            _ = v124
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                }
                                                                            case "a":
                                                                                if ((251 != (v4 & v4)) || false) {
                                                                                    v61 -= -1
                                                                                    var v125 uint = uint(v57)
                                                                                    _ = v125
                                                                                }
                                                                            }
                                                                            v3 = append(v3, 2147483647)
                                                                        } else {
                                                                            fmt.Println(v37, v28)
                                                                        }
                                                                    }
                                                                    v126 := GenS8{F0: 65535, F1: "\n\t"}
                                                                    var v127 int32 = 42
                                                                    _ = v126
                                                                    _ = v127
                                                                    _ = v62
                                                                } else {
                                                                    fmt.Println(v57, fi58, v8)
                                                                    if true {
                                                                        fmt.Println(v2)
                                                                        v28 = 42
                                                                    } else {
                                                                        v37 = v37
                                                                    }
                                                                    var v128 uint64 = (v2 | v2)
                                                                    _ = v128
                                                                }
                                                                var v129 uint64 = (v2 - 643)
                                                                _ = v129
                                                            }
                                                            v57 = 65535
                                                            _ = v61
                                                        }
                                                        _ = v60
                                                    }
                                                default:
                                                    v2 -= (v2 & 1)
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                            } else {
                                                v2 += v2
                                            }
                                            v130 := GenS7{F0: 1, F1: (v28 / ((v28) | 1))}
                                            v4 = genFn0(("\n\t" + ""), -128, 47)
                                            _ = fi58
                                            _ = v130
                                        }
                                        _ = v57
                                    }
                                }
                            }
                        }
                        if false {
                            switch 421 {
                            case -1:
                                if (!true) {
                                    var v131 bool = (32767 == 32767)
                                    _ = v131
                                }
                            default:
                                v8 -= v8
                            }
                            v1[v37] = v37
                        } else {
                            fmt.Println(v2, v28, v4)
                            var v132 float64 = ((1e10 - (1.5 * 1.5)) * 1e10)
                            _ = v132
                        }
                        _ = fi41
                    }
                    switch v37 {
                    case "":
                        switch 127 {
                        case 128:
                            var v133 uint = 1
                            _ = v133
                        case 32768:
                            v28 = (^(v28))
                        default:
                            fmt.Println(v8)
                        }
                    case "foo bar":
                        v2 -= (v2 / ((18446744073709551615) | 1))
                    default:
                        fmt.Println(v4, v8)
                    }
                }
            case 32768:
                for fi134 := 0; fi134 < 16; fi134++ {
                    switch fi134 {
                    case 4294967295:
                        v135 := GenS9{F0: 0, F1: 1}
                        _ = v135
                    default:
                        v28 += -128
                    }
                    _ = fi134
                }
            case 449:
                v136 := GenS7{F0: 0, F1: v28}
                _ = v136
            default:
                switch "\n\t" {
                case "a":
                    var v137 int64 = 42
                    _ = v137
                }
            }
        }
    default:
        fmt.Println(v37, v2, v8)
    }
    if (!(0 == 32767)) {
        if ((v8 * v8) != (-2.25 * ((v8 * (v8 * 0.1)) - v8))) {
            v1[v37] = "foo bar"
            v138 := GenS1{F0: (v37 + "\n\t")}
            _ = v138
        } else {
            v139 := v1[("\n\t" + v37)]
            _ = v139
        }
        v140 := GenS10{F0: (^(v2))}
        if (false || false) {
            fmt.Println(v37)
            v2 -= (v2 * v2)
            if ((0.1 + (3.14159 / (3.14159 / 1.0))) > (1.5 - float64(float32(((0.0 + (((((float64(3.14159) - (float64(float32((3.14159 + 1.5))) - 1.5)) / (float64(v8) * float64(-2.25))) * 3.14159) - (1e10 - float64(((((1e10 + v8) / v8) + (float32(0.1) + v8)) + (v8 * v8))))) + 1.5)) - (-2.25 - (float64((v8 / (1e10 * (v8 / 1.0)))) + ((((3.14159 + (float64(v8) + 1.5)) + 3.14159) - 0.0) * (3.14159 - ((3.14159 + (float64((v8 / v8)) + 0.1)) * (((1.5 * (float64((1.5 - (v8 / v8))) * -2.25)) - 1e10) + float64(v8)))))))))))) {
                var v141 uint8 = (^(v4))
                var v142 float32 = (v8 / 3.14159)
                _ = v141
                _ = v142
            }
        }
        _ = v140
    } else {
        if ((!false) || true) {
            v4 -= v4
        } else {
            var v143 string = v37
            if false {
                v8 = float32(0.0)
            } else {
                for fi144 := 0; fi144 < 3; fi144++ {
                    _ = len(v1)
                    _ = fi144
                }
                fmt.Println(v28)
                if (v28 == v28) {
                    v37 = v143
                    v37 = v143
                } else {
                    var v145 int8 = 42
                    fmt.Println(v2, v8)
                    for fi146 := 0; fi146 < 4; fi146++ {
                        v28 = v145
                        v147 := GenS7{F0: 2147483648, F1: v28}
                        _ = fi146
                        _ = v147
                    }
                    _ = v145
                }
            }
            var v148 float64 = (0.1 - 1.5)
            _ = v143
            _ = v148
        }
    }
    v4 = v4
    v149 := v1[((v37 + (v37 + ("a" + (((v37 + (v37 + (v37 + v37))) + "foo bar") + (((v37 + (v37 + (v37 + ((v37 + v37) + "hello")))) + v37) + v37))))) + ("hello" + ((v37 + v37) + v37)))]
    v37 = (v149 + "a")
    if true {
        v150 := GenS10{F0: (^(v2))}
        v151 := GenS11{F0: -2147483648, F1: -243}
        v37 = (v149 + "a")
        _ = v150
        _ = v151
    }
    for fi152 := 0; fi152 < 14; fi152++ {
        var v153 int = (-(fi152))
        _ = fi152
        _ = v153
    }
    v154 := GenS6{F0: (v28 / (((v28 + v28)) | 1))}
    if (v28 > 42) {
        var v155 float64 = -2.25
        v156 := GenS12{F0: 127, F1: 2, F2: 42}
        fmt.Println(v4, v2)
        _ = v155
        _ = v156
    }
    switch 32768 {
    case 639:
        if len(v6) > 0 {
            v6[0] = 349
        }
    case 0:
        v37 = (v37 + (v149 + (v37 + ("foo bar" + ("a" + "hello")))))
    default:
        var v157 uint32 = 4294967295
        _ = v157
    }
    for fi158 := 0; fi158 < 20; fi158++ {
        fmt.Println(v149, v8)
        _ = fi158
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v149
    _ = v154
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
